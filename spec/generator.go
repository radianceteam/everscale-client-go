package spec

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

func prepareMultiline(data string) string {
	parts := strings.Split(data, "\n")

	for i := range parts {
		parts[i] = strings.Trim(parts[i], " ")
	}

	return strings.Join(parts, "\n// ")
}

// ExceptionForCamelCase - camel case names exception.
var ExceptionForCamelCase = map[string]string{
	"id":                        "ID",
	"xprv":                      "XPrv",
	"url":                       "URL",
	"account_id_address_string": "AccountIDAddressString",
}

func withTypeAlias(name string) string {
	if name == "ClientErrorCode" {
		return "ErrorCode"
	}
	if name == "ClientConfig" {
		return "Config"
	}
	if name == "ClientError" {
		return "Error"
	}
	return name
}

func toGoName(name string) string {
	if cameled, ok := ExceptionForCamelCase[name]; ok {
		return cameled
	}

	cameled := strcase.ToCamel(name)
	if strings.Contains(name, "id") {
		cameled = strings.ReplaceAll(cameled, "Id", "ID")
	}
	if strings.Contains(name, "http") {
		cameled = strings.ReplaceAll(cameled, "Http", "HTTP")
	}
	if strings.Contains(name, "json") {
		cameled = strings.ReplaceAll(cameled, "Json", "JSON")
	}

	return cameled
}

func toTypeName(name string) string {
	parts := strings.Split(name, ".")
	if len(parts) != 1 {
		return parts[1]
	}

	return name
}

func findAppObject(params []Type) *Type {
	for _, p := range params {
		if p.Type == "Generic" && p.GenericName == "AppObject" {
			return &p
		}
	}
	return nil
}

func genFunc(m Module, f Function) string {
	params := make([]Type, 0, 2)
	for _, p := range f.Params {
		if p.Name == "context" || p.Name == "_context" {
			continue // always skips implicit context parameter
		}
		params = append(params, p)
	}
	appObject := findAppObject(params)
	if len(params) > 1 && appObject == nil {
		fmt.Println("WARNING: ignored function", len(params), m.Name, f.Name)

		return ""
	}

	var b bytes.Buffer
	content := funcContent{
		Name:       strcase.ToCamel(m.Name + "_" + f.Name),
		MethodName: m.Name + "." + f.Name,
		ResultType: toTypeName(f.Result.GenericArgs[0].RefName),
	}

	if len(params) == 1 || len(params) == 2 && appObject != nil {
		content.ParamType = toTypeName(params[0].RefName)
	}

	var err error
	if appObject == nil {
		err = funcTemplate.Execute(&b, content)
	} else {
		paramsAppObjectType := toTypeName(appObject.GenericArgs[0].RefName)
		resultAppObjectType := toTypeName(appObject.GenericArgs[1].RefName)
		err = funcTemplateWithAppObject.Execute(&b, funcWithAppObjectContent{
			funcContent:     content,
			AppType:         strings.TrimPrefix(paramsAppObjectType, "ParamsOf"),
			AppObjectFirst:  paramsAppObjectType,
			AppObjectSecond: resultAppObjectType,
		})
	}

	if err != nil {
		panic(err)
	}

	return "\n" + f.ToComment() + b.String()
}

func GenModule(dir string, m Module) error {
	file, err := os.Create(dir + "/mod_" + m.Name + ".gen.go")
	if err != nil {
		return err
	}
	m.Description.GoComment = m.Description.ToComment()
	if err = headerTmpl.Execute(file, m); err != nil {
		return err
	}
	for _, t := range m.Types {
		if ignoredTypesByName[t.Name] {
			continue
		}
		_, err = file.WriteString(GenerateAnyType(m, t, true))
		if err != nil {
			return err
		}
	}
	for _, f := range m.Functions {
		if ignoredFunctionsByName[f.Name] {
			continue
		}
		_, err = file.WriteString(genFunc(m, f))
		if err != nil {
			return err
		}
	}

	return nil
}

// genStruct - generates struct with each field specified.
func genStruct(m Module, t Type) string {
	r := "type " + withTypeAlias(toTypeName(t.Name)) + " struct {\n"
	for _, f := range t.StructFields {
		if f.Name == "" {
			fmt.Println("WARNING: add struct field with empty name", t.Type, t.Name, f)
			f.Name = "value"
		}
		r += "\t" + f.ToComment() + "	" + toGoName(f.Name) + " " + GenerateAnyType(m, f, false)
		if f.Type == Optional {
			r += " `json:\"" + f.Name + "\"` // optional \n"
		} else {
			r += " `json:\"" + f.Name + "\"`\n"
		}
	}
	r += "}\n\n"

	return r
}

// GenerateOptionalType - pointer or null simple type.
func GenerateOptionalType(m Module, t Type) string {
	switch t.Type { // nolint exhaustive
	case Ref:
		if t.RefName == "Value" {
			return emptyInterface
		}

		return "*" + toTypeName(t.RefName)
	case String:
		return "null.String"
	case Value:
		return emptyInterface
	case Number:
		return genNumber(t, true)
	case BigInt:
		return "*big.Int"
	case None:
		return ""
	case Array:
		return "[]" + GenerateAnyType(m, *t.ArrayItem, false)
	case Boolean:
		return "null.Bool"
	default:
		return "* " + GenerateAnyType(m, t, false)
	}
}

// GenerateAnyType - root generator function for type.
func GenerateAnyType(m Module, t Type, isRoot bool) string {
	r := "NotFound::" + string(t.Type) + "::" // easy to find in generated code
	switch t.Type {
	case Ref:
		if t.RefName == "Value" {
			r = emptyInterface
		} else {
			r = withTypeAlias(toTypeName(t.RefName))
		}
	case Optional:
		r = GenerateOptionalType(m, *t.OptionalInner)
	case String:
		r = "string"
	case Value:
		r = emptyInterface
	case Number:
		r = genNumber(t, false)
		if isRoot {
			r = "type " + t.Name + " " + r + "\n"
		}
	case None:
		r = ""
	case Struct:
		r = genStruct(m, t)
	case BigInt:
		r = "big.Int"
	case Array:
		r = "[]" + GenerateAnyType(m, *t.ArrayItem, false)
	case Boolean:
		r = "bool"
	case Generic:
		r = "GENERIC"
	case EnumOfTypes:
		r = genEnumOfTypes(m, t)
	case EnumOfConsts:
		for i := range t.EnumConsts {
			t.EnumConsts[i].ConstName = strcase.ToCamel(t.EnumConsts[i].Name)
		}
		r = genEnumOfConsts(t)
	}

	return r
}
