package spec

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var headerTmpl = template.Must(template.New("header").Parse(
	`package client
// Mod {{.Name}}
// {{.Summary}}
// {{.Description.Description}}

import (
	"github.com/shopspring/decimal"
)

`))

var enumTmpl = template.Must(template.New("enum").Parse(
	`
// ENUM =========================
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
	return strings.Join(strings.Split(data, "\n"), "\n// ")
}

func GenModule(dir string, m Module) error {
	f, err := os.Create(dir + "/mod_" + m.Name + ".gen.go")
	if err != nil {
		return err
	}
	m.Description.Description = prepareMultiline(m.Description.Description)
	if err = headerTmpl.Execute(f, m); err != nil {
		return err
	}
	for _, t := range m.Types {
		_, err = f.WriteString(GenerateType("", t))
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateType(prefix string, t Type) string {
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
			r = "* " + GenerateType(prefix+"	", *t.OptionalInner)
		} else {
			r = GenerateType(prefix+"	", *t.OptionalInner)
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
		r = "type " + t.Name + " struct {\n"
		for _, f := range t.StructFields {
			r += "	" + strcase.ToCamel(f.Name) + " " + GenerateType("", f) + " `json:\"" + f.Name + "\"`\n"
		}
		r += "}\n\n"
	case BigInt:
		r = "decimal.Decimal"
	case Array:
		r = "[]" + GenerateType("", *t.ArrayItem)
	case Boolean:
		r = "bool"
	case Generic:
		r = "GENERIC"
	case EnumOfType:
		r = "type " + t.Name + " interface{}\n\n"
	case EnumOfConsts:
		for i := range t.EnumConsts {
			t.EnumConsts[i].ConstName = strcase.ToScreamingSnake(t.EnumConsts[i].Name)
		}
		r = genEnum(t)
	}

	return r
}
