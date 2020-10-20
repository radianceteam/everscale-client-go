package spec

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
)

var funcMap = template.FuncMap{
	"now": time.Now,
}

const emptyInterface = "interface{}"

var headerTmpl = template.Must(template.New("header").Funcs(funcMap).Parse(
	`package client
// DON'T EDIT THIS FILE is generated {{now.UTC}}
//
// Mod {{.Name}}
//
{{.GoComment}}

import (
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

`))

var enumTmpl = template.Must(template.New("enum").Parse(
	`
type {{.Name}} string
const (
{{range $e := .EnumConsts}} 
	{{$e.ConstName}} {{$.Name}} = "{{$e.Name}}"{{end}}
)
`))

type funcContent struct {
	ResultType string
	ParamType  string
	Name       string
	MethodName string
}

var withoutParamFunc = template.Must(template.New("withoutParamFunc").Parse(
	`func (c *Client) {{.Name}} () (*{{.ResultType}}, error) {
	response := new({{.ResultType}})
	err := c.dllClient.waitErrorOrResultUnmarshal("{{.MethodName}}", nil, response)

	return response, err
}
`))

var singleParamFunc = template.Must(template.New("singleParamFunc").Parse(
	`func (c *Client) {{.Name}} (p *{{.ParamType}}) (*{{.ResultType}}, error) {
	response := new({{.ResultType}})
	err := c.dllClient.waitErrorOrResultUnmarshal("{{.MethodName}}", p, response)

	return response, err
}
`))

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

func toGoName(name string) string {
	if name == "id" {
		return "ID"
	} else if name == "xprv" {
		return "XPrv"
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
		fmt.Println("WARNING: ignored function", len(params), f.Name)

		return ""
	}

	var b bytes.Buffer
	content := funcContent{
		ResultType: f.Result.GenericArgs[0].RefName,
		Name:       strcase.ToCamel(m.Name + "_" + f.Name),
		MethodName: m.Name + "." + f.Name,
	}
	if len(params) == 0 {
		if err := withoutParamFunc.Execute(&b, content); err != nil {
			panic(err)
		}
	} else if len(params) == 1 {
		content.ParamType = params[0].RefName
		if err := singleParamFunc.Execute(&b, content); err != nil {
			panic(err)
		}
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
		r += f.ToComment() + "	" + toGoName(f.Name) + " " + GenerateType(f)
		if f.Type == Optional {
			r += " `json:\"" + f.Name + "\"` // optional \n"
		} else {
			r += " `json:\"" + f.Name + "\"`\n"
		}
	}
	r += "}\n\n"

	return r
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
			r = "interface{}"
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
		r = "type " + t.Name + " interface{}\n\n"
	case EnumOfConsts:
		for i := range t.EnumConsts {
			t.EnumConsts[i].ConstName = strcase.ToCamel(t.EnumConsts[i].Name)
		}
		r = genEnum(t)
	}

	return r
}
