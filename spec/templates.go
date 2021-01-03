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

const emptyInterface = "json.RawMessage"

var headerTmpl = template.Must(template.New("header").Funcs(funcMap).Parse(
	`package client
// DON'T EDIT THIS FILE is generated {{now}}
//
// Mod {{.Name}}
//
{{.GoComment}}

import (
	"math/big"
	"encoding/json"

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

type funcWithAppObjectContent struct {
	funcContent
	AppObjectFirst  string
	AppObjectSecond string
}

var funcTemplate = template.Must(template.New("funcTemplate").Parse(
	`func (c *Client) {{.Name}}( {{if ne .ParamType ""}} p *{{.ParamType}} {{end}} ) {{if eq .ResultType ""}} error {{else}} (*{{.ResultType}}, error) {{end}} {
	{{if ne .ResultType ""}} result := new({{.ResultType}}) {{end}}
	{{if eq .ResultType "" }}
	_, err := c.dllClient.waitErrorOrResult("{{.MethodName}}", {{if eq .ParamType "" }} nil {{else}} p {{end}})
	{{else}}
	err := c.dllClient.waitErrorOrResultUnmarshal("{{.MethodName}}", {{if eq .ParamType "" }} nil {{else}} p {{end}}, result)
	{{end}}

	return {{if ne .ResultType "" }} result, {{ end }} err
}
`))

var funcTemplateWithAppObject = template.Must(template.New("funcTemplate").Parse(
	`func (c *Client) {{.Name}}( {{if ne .ParamType ""}} p *{{.ParamType}} {{end}} ) (*{{.ResultType}}, error)  {
	result := new({{.ResultType}}) 
	responses, err := c.dllClient.resultsChannel("{{.MethodName}}", {{if ne .ParamType ""}} p {{else}} nil {{end}})
	if err != nil {
		return nil, err
	}

	response := <- responses
	if response.Code == ResponseCodeError {
		return nil, response.Error
	}

	if err := json.Unmarshal(response.Data, result); err != nil {
		return nil, err
	}
	
	// result - is populated

    // first = {{.AppObjectFirst}}
    // second = {{.AppObjectSecond}}

	return result, nil
}
`))
