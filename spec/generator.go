package spec

import (
	"bytes"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
)

var funcMap = template.FuncMap{
	"now": time.Now,
}

var NewLiner = regexp.MustCompile("(?m)^.*$")

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
	parts := strings.Split(data, "\n")

	for i := range parts {
		parts[i] = strings.Trim(parts[i], " ")
	}

	return strings.Join(parts, "\n// ")
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
	m.Description.Summary = prepareMultiline(m.Description.Summary)
	if err = headerTmpl.Execute(f, m); err != nil {
		return err
	}
	for _, t := range m.Types {
		if ignoredTypesByName[t.Name] {
			continue
		}
		_, err = f.WriteString(GenerateType(t))
		if err != nil {
			return err
		}
	}

	return nil
}

func genStruct(t Type) string {
	r := "type " + t.Name + " struct {\n"
	for _, f := range t.StructFields {
		if !strings.Contains(f.Description.Description, f.Description.Summary) {
			r += "// " + prepareMultiline(f.Description.Summary) + "\n"
		}
		if f.Description.Description != "" {
			r += "// " + prepareMultiline(f.Description.Description) + "\n"
		}
		r += "	" + toGoName(f.Name) + " " + GenerateType(f)
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
