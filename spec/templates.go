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
// DON'T EDIT THIS FILE! It is generated via 'task generate' at {{now}}
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

var enumOfConstsTpl = template.Must(template.New("enumOfConstsTpl").Parse(
	`
{{if ne .Name ""}}
type {{.Name}} {{.GoType}}
{{ end }}

const (
{{range $e := .EnumConsts}} 
	{{$e.GoComment}} {{$e.ConstName}} {{$.Name}} = {{$e.Value}}{{end}}
)

{{if eq .Name ""}}
func init() { // nolint gochecknoinits {{range $e := .EnumConsts}} 
	 errorCodesToErrorTypes[{{$e.ConstName}}] = "{{$e.ConstName}}"{{end}}
}
{{ end }}

`))

var enumOfTypesTpl = template.Must(template.New("enumOfTypesTpl").Parse(
	`
type {{.Name}} struct {
	EnumTypeValue interface{} // any of {{range $e := .EnumTypes}}{{$e.GoType}}, {{end}}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *{{.Name}}) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	{{range $e := .EnumTypes}}
	case {{$e.GoType}}:
		return json.Marshal(struct {
			{{$e.GoType}}
			Type string ` + "`json:\"type\"`" + `
		}{
			value,
			"{{$e.Name}}",
		})
	{{end}}

	default:
		return nil, fmt.Errorf("unsupported type for {{.Name}} %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *{{.Name}}) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
        return err
    }
	switch typeDescriptor.Type {
	{{range $e := .EnumTypes}}
	case "{{$e.Name}}":
		var enumTypeValue {{$e.GoType}}
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
        	return err
    	}
		p.EnumTypeValue = enumTypeValue
	{{end}}

	default:
		return fmt.Errorf("unsupported type for {{.Name}} %v", typeDescriptor.Type)
	}

	return nil
}
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
	AppObjectParams string
	AppObjectResult string
	Requests        []string
	Notifications   []string
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

var funcTemplateWithAppObject = template.Must(template.New("funcTemplateWithAppObject").Parse(
	`
func (c *Client) {{.Name}}( {{if ne .ParamType ""}} p *{{.ParamType}}, {{end}} app {{.AppType}}) (*{{.ResultType}}, error)  { // nolint dupl
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
{{if .Notifications}} 
			if r.Code == ResponseCodeAppNotify {
				c.dispatchNotify{{.Name}}(r.Data, app)
			}
{{end}}
		}
	}()

	return result, nil
}

func (c *Client) dispatchRequest{{.Name}}(payload []byte, app {{.AppType}}) { // nolint dupl
	var appRequest ParamsOfAppRequest
	var appParams {{.AppObjectParams}}
	err := json.Unmarshal(payload, &appRequest)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(appRequest.RequestData, &appParams)
	if err != nil {
		panic(err)
	}
	var appResponse interface{}
	// appResponse, err := app.Request(appParams)

		switch value := (appParams.EnumTypeValue).(type) {
	{{range $r := .Requests}} 
	case {{$r}}{{$.AppObjectParams}}:
		appResponse, err = app.{{$r}}Request(value)
	{{end}}

	default:
		err = fmt.Errorf("unsupported type for request %v", appParams.EnumTypeValue)
	}

	appRequestResult := AppRequestResult{}
	if err != nil {
		appRequestResult.EnumTypeValue = ErrorAppRequestResult{Text:err.Error()}
	} else {
		marshalled, _:= json.Marshal(&{{.AppObjectResult}}{EnumTypeValue: appResponse})
		appRequestResult.EnumTypeValue = OkAppRequestResult{Result:marshalled}
	}
	err = c.ClientResolveAppRequest(&ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result: appRequestResult,
	})
	if err != nil {
		panic(err)
	}
}

{{if .Notifications}} 
func (c *Client) dispatchNotify{{.Name}}(payload []byte, app {{.AppType}}) { // nolint dupl
	var appParams {{.AppObjectParams}}
	err := json.Unmarshal(payload, &appParams)
	if err != nil {
		panic(err)
	}

	switch value := (appParams.EnumTypeValue).(type) {
	{{range $r := .Notifications}} 
	case {{$r}}{{$.AppObjectParams}}:
		app.{{$r}}Notify(value)
	{{end}}
	default:
		panic(fmt.Errorf("unsupported type for request %v", appParams.EnumTypeValue))
	}
}
{{end}}
`))
