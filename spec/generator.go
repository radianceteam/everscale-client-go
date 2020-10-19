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
// Mod {{.Name}}
// {{.Summary}}
// {{.Description.Description}}

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
	return strings.Join(strings.Split(data, "\n"), "\n// ")
}

func toGoName(name string) string {
	if name == "id" {
		return "ID"
	}
	cameled := strcase.ToCamel(name)
	if strings.Contains(name, "id") {
		cameled = strings.ReplaceAll(cameled, "Id", "ID")
	}

	return cameled
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
		if ignoredTypesByName[t.Name] {
			continue
		}
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
			r += "	" + toGoName(f.Name) + " " + GenerateType("", f)
			if f.Type == Optional {
				r += " `json:\"" + f.Name + ",omitempty\"`\n"
			} else {
				r += " `json:\"" + f.Name + "\"`\n"
			}
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
