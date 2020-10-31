package spec

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

func genEnum(t Type) string {
	var tpl bytes.Buffer
	if err := enumTmpl.Execute(&tpl, t); err != nil {
		panic(err)
	}

	return tpl.String()
}

func prepareMultiline(data string) string {
	parts := strings.Split(data, "\n")

	for i := range parts {
		parts[i] = strings.Trim(parts[i], " ")
	}

	return strings.Join(parts, "\n// ")
}

// ExceptionForCamelCase - camel case names exception.
var ExceptionForCamelCase = map[string]string{
	"id":   "ID",
	"xprv": "XPrv",
	"url":  "URL",
}

func toGoName(name string) string {
	if cameled, ok := ExceptionForCamelCase[name]; ok {
		return cameled
	}

	cameled := strcase.ToCamel(name)
	if strings.Contains(name, "id") {
		cameled = strings.ReplaceAll(cameled, "Id", "ID")
	}

	return cameled
}

func genFunc(m Module, f Function) string {
	params := make([]Type, 0, 2)
	for _, p := range f.Params {
		if p.Name == "context" || p.Name == "_context" {
			continue
		}
		params = append(params, p)
	}
	if len(params) > 1 {
		fmt.Println("WARNING: ignored function", len(params), m.Name, f.Name)

		return ""
	}

	var b bytes.Buffer
	content := funcContent{
		Name:       strcase.ToCamel(m.Name + "_" + f.Name),
		MethodName: m.Name + "." + f.Name,
		ResultType: f.Result.GenericArgs[0].RefName,
	}

	var tmpl *template.Template
	if len(params) == 0 {
		tmpl = withoutParamFunc
	} else if len(params) == 1 {
		content.ParamType = params[0].RefName
		if content.ResultType != "" {
			tmpl = singleParamFunc
		} else {
			tmpl = singleParamWithoutResultFunc
		}
	}

	if err := tmpl.Execute(&b, content); err != nil {
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
		_, err = file.WriteString(GenerateType(t))
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

func genStruct(t Type) string {
	r := "type " + t.Name + " struct {\n"
	for _, f := range t.StructFields {
		if f.Name == "" {
			fmt.Println("WARNING: skip struct field without name", t.Type, t.Name, f)

			continue
		}
		r += "\t" + f.ToComment() + "	" + toGoName(f.Name) + " " + GenerateType(f)
		if f.Type == Optional {
			r += " `json:\"" + f.Name + "\"` // optional \n"
		} else {
			r += " `json:\"" + f.Name + "\"`\n"
		}
	}
	r += "}\n\n"

	return r
}

func genEnumOfTypes(t Type) string {
	fields := make(map[string]TypeName)
	structFields := make([]Type, 1)
	structFields[0] = Type{Type: String, Description: Description{Name: "type"}}
	fields["type"] = String
	for _, t := range t.EnumTypes {
		if t.Type != Struct {
			panic("EnumOfTypes only supports structs " + t.Name)
		}
		for _, sf := range t.StructFields {
			prevName, ok := fields[sf.Name]
			if ok && prevName != sf.Type {
				panic("type mismatch for enum " + t.Name + ":" + sf.Name)
			}
			fields[sf.Name] = sf.Type
			if !ok {
				structFields = append(structFields, sf)
			}
		}
	}
	t.StructFields = structFields

	return genStruct(t)
}

func GenerateOptionalType(t Type) string {
	switch t.Type { // nolint exhaustive
	case Ref:
		if t.RefName == "Value" {
			return emptyInterface
		}

		return "*" + t.RefName
	case String:
		return "null.String"
	case Value:
		return emptyInterface
	case Number:
		return "null.Int"
	case BigInt:
		return "decimal.NullDecimal"
	case None:
		return ""
	case Array:
		return "[]" + GenerateType(*t.ArrayItem)
	case Boolean:
		return "null.Bool"
	default:
		return "* " + GenerateType(t)
	}
}

func GenerateType(t Type) string {
	r := "NotFound::" + string(t.Type) + "::"
	switch t.Type {
	case Ref:
		if t.RefName == "Value" {
			r = emptyInterface
		} else {
			r = t.RefName + ""
		}
	case Optional:
		r = GenerateOptionalType(*t.OptionalInner)
	case String:
		r = "string"
	case Value:
		r = emptyInterface
	case Number:
		r = "int"
	case None:
		r = ""
	case Struct:
		r = genStruct(t)
	case BigInt:
		r = "decimal.Decimal"
	case Array:
		r = "[]" + GenerateType(*t.ArrayItem)
	case Boolean:
		r = "bool"
	case Generic:
		r = "GENERIC"
	case EnumOfTypes:
		r = genEnumOfTypes(t)
	case EnumOfConsts:
		for i := range t.EnumConsts {
			t.EnumConsts[i].ConstName = strcase.ToCamel(t.EnumConsts[i].Name)
		}
		r = genEnum(t)
	}

	return r
}
