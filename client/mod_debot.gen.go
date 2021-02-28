package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 28 Feb 21 15:56 UTC
//
// Mod debot
//
// [UNSTABLE](UNSTABLE.md) Module for working with debot.

import (
	"encoding/json"
	"fmt"
)

const (
	DebotStartFailedDebotErrorCode        = 801
	DebotFetchFailedDebotErrorCode        = 802
	DebotExecutionFailedDebotErrorCode    = 803
	DebotInvalidHandleDebotErrorCode      = 804
	DebotInvalidJSONParamsDebotErrorCode  = 805
	DebotInvalidFunctionIDDebotErrorCode  = 806
	DebotInvalidAbiDebotErrorCode         = 807
	DebotGetMethodFailedDebotErrorCode    = 808
	DebotInvalidMsgDebotErrorCode         = 809
	DebotExternalCallFailedDebotErrorCode = 810
)

func init() { // nolint gochecknoinits
	errorCodesToErrorTypes[DebotStartFailedDebotErrorCode] = "DebotStartFailedDebotErrorCode"
	errorCodesToErrorTypes[DebotFetchFailedDebotErrorCode] = "DebotFetchFailedDebotErrorCode"
	errorCodesToErrorTypes[DebotExecutionFailedDebotErrorCode] = "DebotExecutionFailedDebotErrorCode"
	errorCodesToErrorTypes[DebotInvalidHandleDebotErrorCode] = "DebotInvalidHandleDebotErrorCode"
	errorCodesToErrorTypes[DebotInvalidJSONParamsDebotErrorCode] = "DebotInvalidJSONParamsDebotErrorCode"
	errorCodesToErrorTypes[DebotInvalidFunctionIDDebotErrorCode] = "DebotInvalidFunctionIDDebotErrorCode"
	errorCodesToErrorTypes[DebotInvalidAbiDebotErrorCode] = "DebotInvalidAbiDebotErrorCode"
	errorCodesToErrorTypes[DebotGetMethodFailedDebotErrorCode] = "DebotGetMethodFailedDebotErrorCode"
	errorCodesToErrorTypes[DebotInvalidMsgDebotErrorCode] = "DebotInvalidMsgDebotErrorCode"
	errorCodesToErrorTypes[DebotExternalCallFailedDebotErrorCode] = "DebotExternalCallFailedDebotErrorCode"
}

type (
	DebotHandle uint32
	DebotAction struct {
		// A short action description.
		// Should be used by Debot Browser as name of menu item.
		Description string `json:"description"`
		// Depends on action type.
		// Can be a debot function name or a print string (for Print Action).
		Name string `json:"name"`
		// Action type.
		ActionType uint8 `json:"action_type"`
		// ID of debot context to switch after action execution.
		To uint8 `json:"to"`
		// Action attributes.
		// In the form of "param=value,flag". attribute example: instant, args, fargs, sign.
		Attributes string `json:"attributes"`
		// Some internal action data.
		// Used by debot only.
		Misc string `json:"misc"`
	}
)

type ParamsOfStart struct {
	// Debot smart contract address.
	Address string `json:"address"`
}

type RegisteredDebot struct {
	// Debot handle which references an instance of debot engine.
	DebotHandle DebotHandle `json:"debot_handle"`
}

type LogParamsOfAppDebotBrowser struct {
	// A string that must be printed to user.
	Msg string `json:"msg"`
}

type SwitchParamsOfAppDebotBrowser struct {
	// Debot context ID to which debot is switched.
	ContextID uint8 `json:"context_id"`
}

type SwitchCompletedParamsOfAppDebotBrowser struct {
}

type ShowActionParamsOfAppDebotBrowser struct {
	// Debot action that must be shown to user as menu item. At least `description` property must be shown from [DebotAction] structure.
	Action DebotAction `json:"action"`
}

type InputParamsOfAppDebotBrowser struct {
	// A prompt string that must be printed to user before input request.
	Prompt string `json:"prompt"`
}

type GetSigningBoxParamsOfAppDebotBrowser struct {
}

type InvokeDebotParamsOfAppDebotBrowser struct {
	// Address of debot in blockchain.
	DebotAddr string `json:"debot_addr"`
	// Debot action to execute.
	Action DebotAction `json:"action"`
}

type SendParamsOfAppDebotBrowser struct {
	// Internal message to DInterface address.
	// Message body contains interface function and parameters.
	Message string `json:"message"`
}

type ParamsOfAppDebotBrowser struct {
	EnumTypeValue interface{} // any of LogParamsOfAppDebotBrowser, SwitchParamsOfAppDebotBrowser, SwitchCompletedParamsOfAppDebotBrowser, ShowActionParamsOfAppDebotBrowser, InputParamsOfAppDebotBrowser, GetSigningBoxParamsOfAppDebotBrowser, InvokeDebotParamsOfAppDebotBrowser, SendParamsOfAppDebotBrowser,
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ParamsOfAppDebotBrowser) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case LogParamsOfAppDebotBrowser:
		return json.Marshal(struct {
			LogParamsOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"Log",
		})

	case SwitchParamsOfAppDebotBrowser:
		return json.Marshal(struct {
			SwitchParamsOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"Switch",
		})

	case SwitchCompletedParamsOfAppDebotBrowser:
		return json.Marshal(struct {
			SwitchCompletedParamsOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"SwitchCompleted",
		})

	case ShowActionParamsOfAppDebotBrowser:
		return json.Marshal(struct {
			ShowActionParamsOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"ShowAction",
		})

	case InputParamsOfAppDebotBrowser:
		return json.Marshal(struct {
			InputParamsOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"Input",
		})

	case GetSigningBoxParamsOfAppDebotBrowser:
		return json.Marshal(struct {
			GetSigningBoxParamsOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"GetSigningBox",
		})

	case InvokeDebotParamsOfAppDebotBrowser:
		return json.Marshal(struct {
			InvokeDebotParamsOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"InvokeDebot",
		})

	case SendParamsOfAppDebotBrowser:
		return json.Marshal(struct {
			SendParamsOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"Send",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ParamsOfAppDebotBrowser %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ParamsOfAppDebotBrowser) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "Log":
		var enumTypeValue LogParamsOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Switch":
		var enumTypeValue SwitchParamsOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "SwitchCompleted":
		var enumTypeValue SwitchCompletedParamsOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "ShowAction":
		var enumTypeValue ShowActionParamsOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Input":
		var enumTypeValue InputParamsOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "GetSigningBox":
		var enumTypeValue GetSigningBoxParamsOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "InvokeDebot":
		var enumTypeValue InvokeDebotParamsOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Send":
		var enumTypeValue SendParamsOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ParamsOfAppDebotBrowser %v", typeDescriptor.Type)
	}

	return nil
}

type InputResultOfAppDebotBrowser struct {
	// String entered by user.
	Value string `json:"value"`
}

type GetSigningBoxResultOfAppDebotBrowser struct {
	// Signing box for signing data requested by debot engine.
	// Signing box is owned and disposed by debot engine.
	SigningBox SigningBoxHandle `json:"signing_box"`
}

type InvokeDebotResultOfAppDebotBrowser struct {
}

type ResultOfAppDebotBrowser struct {
	EnumTypeValue interface{} // any of InputResultOfAppDebotBrowser, GetSigningBoxResultOfAppDebotBrowser, InvokeDebotResultOfAppDebotBrowser,
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ResultOfAppDebotBrowser) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case InputResultOfAppDebotBrowser:
		return json.Marshal(struct {
			InputResultOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"Input",
		})

	case GetSigningBoxResultOfAppDebotBrowser:
		return json.Marshal(struct {
			GetSigningBoxResultOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"GetSigningBox",
		})

	case InvokeDebotResultOfAppDebotBrowser:
		return json.Marshal(struct {
			InvokeDebotResultOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"InvokeDebot",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ResultOfAppDebotBrowser %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ResultOfAppDebotBrowser) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "Input":
		var enumTypeValue InputResultOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "GetSigningBox":
		var enumTypeValue GetSigningBoxResultOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "InvokeDebot":
		var enumTypeValue InvokeDebotResultOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ResultOfAppDebotBrowser %v", typeDescriptor.Type)
	}

	return nil
}

type ParamsOfFetch struct {
	// Debot smart contract address.
	Address string `json:"address"`
}

type ParamsOfExecute struct {
	// Debot handle which references an instance of debot engine.
	DebotHandle DebotHandle `json:"debot_handle"`
	// Debot Action that must be executed.
	Action DebotAction `json:"action"`
}

type ParamsOfSend struct {
	// Debot handle which references an instance of debot engine.
	DebotHandle DebotHandle `json:"debot_handle"`
	// Std address of interface or debot.
	Source string `json:"source"`
	// Function Id to call.
	FuncID uint32 `json:"func_id"`
	// Json string with parameters.
	Params string `json:"params"`
}

// [UNSTABLE](UNSTABLE.md) Starts an instance of debot.
// Downloads debot smart contract from blockchain and switches it to
// context zero.
// Returns a debot handle which can be used later in `execute` function.
// This function must be used by Debot Browser to start a dialog with debot.
// While the function is executing, several Browser Callbacks can be called,
// since the debot tries to display all actions from the context 0 to the user.
//
// # Remarks
// `start` is equivalent to `fetch` + switch to context 0.
func (c *Client) DebotStart(p *ParamsOfStart, app AppDebotBrowser) (*RegisteredDebot, error) {
	result := new(RegisteredDebot)
	responses, err := c.dllClient.resultsChannel("debot.start", p)
	if err != nil {
		return nil, err
	}

	response := <-responses
	if response.Code == ResponseCodeError {
		return nil, response.Error
	}

	if err := json.Unmarshal(response.Data, result); err != nil {
		return nil, err
	}

	go func() {
		for r := range responses {
			if r.Code == ResponseCodeAppRequest {
				c.dispatchRequestDebotStart(r.Data, app)
			}
			if r.Code == ResponseCodeAppNotify {
				c.dispatchNotifyDebotStart(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (c *Client) dispatchRequestDebotStart(payload []byte, app AppDebotBrowser) {
	var appRequest ParamsOfAppRequest
	var appParams ParamsOfAppDebotBrowser
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
		appRequestResult.EnumTypeValue = ErrorAppRequestResult{Text: err.Error()}
	} else {
		marshalled, _ := json.Marshal(&appResponse)
		appRequestResult.EnumTypeValue = OkAppRequestResult{Result: marshalled}
	}
	err = c.ClientResolveAppRequest(&ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result:       appRequestResult,
	})
	if err != nil {
		panic(err)
	}
}

func (c *Client) dispatchNotifyDebotStart(payload []byte, app AppDebotBrowser) {
	var appParams ParamsOfAppDebotBrowser
	err := json.Unmarshal(payload, &appParams)
	if err != nil {
		panic(err)
	}
	app.Notify(appParams)
}

// [UNSTABLE](UNSTABLE.md) Fetches debot from blockchain.
// Downloads debot smart contract (code and data) from blockchain and creates
// an instance of Debot Engine for it.
//
// # Remarks
// It does not switch debot to context 0. Browser Callbacks are not called.
func (c *Client) DebotFetch(p *ParamsOfFetch, app AppDebotBrowser) (*RegisteredDebot, error) {
	result := new(RegisteredDebot)
	responses, err := c.dllClient.resultsChannel("debot.fetch", p)
	if err != nil {
		return nil, err
	}

	response := <-responses
	if response.Code == ResponseCodeError {
		return nil, response.Error
	}

	if err := json.Unmarshal(response.Data, result); err != nil {
		return nil, err
	}

	go func() {
		for r := range responses {
			if r.Code == ResponseCodeAppRequest {
				c.dispatchRequestDebotFetch(r.Data, app)
			}
			if r.Code == ResponseCodeAppNotify {
				c.dispatchNotifyDebotFetch(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (c *Client) dispatchRequestDebotFetch(payload []byte, app AppDebotBrowser) {
	var appRequest ParamsOfAppRequest
	var appParams ParamsOfAppDebotBrowser
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
		appRequestResult.EnumTypeValue = ErrorAppRequestResult{Text: err.Error()}
	} else {
		marshalled, _ := json.Marshal(&appResponse)
		appRequestResult.EnumTypeValue = OkAppRequestResult{Result: marshalled}
	}
	err = c.ClientResolveAppRequest(&ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result:       appRequestResult,
	})
	if err != nil {
		panic(err)
	}
}

func (c *Client) dispatchNotifyDebotFetch(payload []byte, app AppDebotBrowser) {
	var appParams ParamsOfAppDebotBrowser
	err := json.Unmarshal(payload, &appParams)
	if err != nil {
		panic(err)
	}
	app.Notify(appParams)
}

// [UNSTABLE](UNSTABLE.md) Executes debot action.
// Calls debot engine referenced by debot handle to execute input action.
// Calls Debot Browser Callbacks if needed.
//
// # Remarks
// Chain of actions can be executed if input action generates a list of subactions.
func (c *Client) DebotExecute(p *ParamsOfExecute) error {
	_, err := c.dllClient.waitErrorOrResult("debot.execute", p)

	return err
}

// [UNSTABLE](UNSTABLE.md) Sends message to Debot.
// Used by Debot Browser to send response on Dinterface call or from other Debots.
func (c *Client) DebotSend(p *ParamsOfSend) error {
	_, err := c.dllClient.waitErrorOrResult("debot.send", p)

	return err
}

// [UNSTABLE](UNSTABLE.md) Destroys debot handle.
// Removes handle from Client Context and drops debot engine referenced by that handle.
func (c *Client) DebotRemove(p *RegisteredDebot) error {
	_, err := c.dllClient.waitErrorOrResult("debot.remove", p)

	return err
}
