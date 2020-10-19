package spec

import (
	"bytes"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
)

var funcMap = template.FuncMap{
	"now": time.Now,
}

var headerTmpl = template.Must(template.New("header").Funcs(funcMap).Parse(
	`package client
// DON'T EDIT THIS FILE is generated {{now.UTC}}
//
// Mod {{.Name}}
//
{{.GoComment}}

import (
	"github.com/shopspring/decimal"
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

func genFunc(f Function) string {
	return f.ToComment() + "func (c *Client) " + strcase.ToCamel(f.Name) + "() {}\n"
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
		_, err = file.WriteString(genFunc(f))
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
			r += " `json:\"" + f.Name + ",omitempty\"`\n"
		} else {
			r += " `json:\"" + f.Name + "\"`\n"
		}
	}
	r += "}\n\n"

	return r
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
		if t.OptionalInner.Type != Ref || t.OptionalInner.RefName != "Value" {
			r = "* " + GenerateType(*t.OptionalInner)
		} else {
			r = GenerateType(*t.OptionalInner)
		}
	case String:
		r = "string"
	case Value:
		r = "interface{}"
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
