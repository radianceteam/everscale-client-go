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

var singleParamWithoutResultFunc = template.Must(template.New("singleParamWithoutResultFunc").Parse(
	`func (c *Client) {{.Name}} (p *{{.ParamType}}) error {
	_, err := c.dllClient.waitErrorOrResult("{{.MethodName}}", p)

	return  err
}
`))
