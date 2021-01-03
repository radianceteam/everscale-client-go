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
	AppType         string
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
	`func (c *Client) {{.Name}}( {{if ne .ParamType ""}} p *{{.ParamType}}, {{end}} app {{.AppType}}) (*{{.ResultType}}, error)  {
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

	go func() {
		for r := range responses {
			if r.Code == ResponseCodeAppRequest {
				c.dispatchRequest{{.Name}}(r.Data, app)
			}
			if r.Code == ResponseCodeAppNotify {
				c.dispatchNotify{{.Name}}(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (c *Client) dispatchRequest{{.Name}}(payload []byte, app {{.AppType}}) {
	var appRequest ParamsOfAppRequest
	var appParams {{.AppObjectFirst}}
	err := json.Unmarshal(payload, &appRequest)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(appRequest.RequestData, &appParams)
	if err != nil {
		panic(err)
	}
	appResponse, err := app.Request(appParams)
	appRequestResult := AppRequestResult{}
	if err != nil {
		appRequestResult.Type = ErrorAppRequestResultType
		appRequestResult.Text = err.Error()
	} else {
		appRequestResult.Type = OkAppRequestResultType
		appRequestResult.Result, _ = json.Marshal(appResponse)
	}
	err = c.ClientResolveAppRequest(&ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result: appRequestResult,
	})
	if err != nil {
		panic(err)
	}
}

func (c *Client) dispatchNotify{{.Name}}(payload []byte, app {{.AppType}}) {
	var appParams {{.AppObjectFirst}}
	err := json.Unmarshal(payload, &appParams)
	if err != nil {
		panic(err)
	}
	app.Notify(appParams)
}
`))
