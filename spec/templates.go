package spec

import (
	"text/template"
	"time"
)

var funcMap = template.FuncMap{
	"now": func() interface{} {
		return time.Now().UTC().Format("02 Jan 06 15:04 MST")
	},
}

const emptyInterface = "interface{}"

var headerTmpl = template.Must(template.New("header").Funcs(funcMap).Parse(
	`package client
// DON'T EDIT THIS FILE is generated {{now}}
//
// Mod {{.Name}}
//
{{.GoComment}}

import (
	"math/big"

	"github.com/volatiletech/null"
)

`))

var enumTmpl = template.Must(template.New("enum").Parse(
	`
type {{.Name}} string
const (
{{range $e := .EnumConsts}} 
	{{$e.GoComment}} {{$e.ConstName}} {{$.Name}} = "{{$e.Name}}"{{end}}
)
`))

type funcContent struct {
	ResultType string
	ParamType  string
	Name       string
	MethodName string
}

var funcTemplate = template.Must(template.New("funcTemplate").Parse(
	`func (c *Client) {{.Name}}( {{if ne .ParamType ""}} p *{{.ParamType}} {{end}} ) {{if eq .ResultType ""}} error {{else}} (*{{.ResultType}}, error) {{end}} {
	{{if ne .ResultType ""}} response := new({{.ResultType}}) {{end}}
	{{if eq .ResultType "" }}
	_, err := c.dllClient.waitErrorOrResult("{{.MethodName}}", {{if eq .ParamType "" }} nil {{else}} p {{end}})
	{{else}}
	err := c.dllClient.waitErrorOrResultUnmarshal("{{.MethodName}}", {{if eq .ParamType "" }} nil {{else}} p {{end}}, response)
	{{end}}

	return {{if ne .ResultType "" }} response, {{ end }} err
}
`))
