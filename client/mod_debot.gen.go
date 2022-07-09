package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 09 Jul 22 15:07 UTC
//
// Mod debot
//
// [UNSTABLE](UNSTABLE.md) Module for working with debot.

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/volatiletech/null"
)

const (
	DebotStartFailedDebotErrorCode           = 801
	DebotFetchFailedDebotErrorCode           = 802
	DebotExecutionFailedDebotErrorCode       = 803
	DebotInvalidHandleDebotErrorCode         = 804
	DebotInvalidJSONParamsDebotErrorCode     = 805
	DebotInvalidFunctionIDDebotErrorCode     = 806
	DebotInvalidAbiDebotErrorCode            = 807
	DebotGetMethodFailedDebotErrorCode       = 808
	DebotInvalidMsgDebotErrorCode            = 809
	DebotExternalCallFailedDebotErrorCode    = 810
	DebotBrowserCallbackFailedDebotErrorCode = 811
	DebotOperationRejectedDebotErrorCode     = 812
	DebotNoCodeDebotErrorCode                = 813
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
	errorCodesToErrorTypes[DebotBrowserCallbackFailedDebotErrorCode] = "DebotBrowserCallbackFailedDebotErrorCode"
	errorCodesToErrorTypes[DebotOperationRejectedDebotErrorCode] = "DebotOperationRejectedDebotErrorCode"
	errorCodesToErrorTypes[DebotNoCodeDebotErrorCode] = "DebotNoCodeDebotErrorCode"
}

type DebotHandle uint32

// [UNSTABLE](UNSTABLE.md) Describes a debot action in a Debot Context.
type DebotAction struct {
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

// [UNSTABLE](UNSTABLE.md) Describes DeBot metadata.
type DebotInfo struct {
	// DeBot short name.
	Name null.String `json:"name"` // optional
	// DeBot semantic version.
	Version null.String `json:"version"` // optional
	// The name of DeBot deployer.
	Publisher null.String `json:"publisher"` // optional
	// Short info about DeBot.
	Caption null.String `json:"caption"` // optional
	// The name of DeBot developer.
	Author null.String `json:"author"` // optional
	// TON address of author for questions and donations.
	Support null.String `json:"support"` // optional
	// String with the first messsage from DeBot.
	Hello null.String `json:"hello"` // optional
	// String with DeBot interface language (ISO-639).
	Language null.String `json:"language"` // optional
	// String with DeBot ABI.
	Dabi null.String `json:"dabi"` // optional
	// DeBot icon.
	Icon null.String `json:"icon"` // optional
	// Vector with IDs of DInterfaces used by DeBot.
	Interfaces []string `json:"interfaces"`
	// ABI version ("x.y") supported by DeBot.
	DabiVersion string `json:"dabiVersion"`
}

// [UNSTABLE](UNSTABLE.md) Describes the operation that the DeBot wants to perform.

// DeBot wants to create new transaction in blockchain.
type TransactionDebotActivity struct {
	// External inbound message BOC.
	Msg string `json:"msg"`
	// Target smart contract address.
	Dst string `json:"dst"`
	// List of spendings as a result of transaction.
	Out []Spending `json:"out"`
	// Transaction total fee.
	Fee big.Int `json:"fee"`
	// Indicates if target smart contract updates its code.
	Setcode bool `json:"setcode"`
	// Public key from keypair that was used to sign external message.
	Signkey string `json:"signkey"`
	// Signing box handle used to sign external message.
	SigningBoxHandle uint32 `json:"signing_box_handle"`
}

type DebotActivity struct {
	// Should be any of
	// TransactionDebotActivity
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *DebotActivity) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case TransactionDebotActivity:
		return json.Marshal(struct {
			TransactionDebotActivity
			Type string `json:"type"`
		}{
			value,
			"Transaction",
		})

	default:
		return nil, fmt.Errorf("unsupported type for DebotActivity %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *DebotActivity) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "Transaction":
		var enumTypeValue TransactionDebotActivity
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for DebotActivity %v", typeDescriptor.Type)
	}

	return nil
}

// [UNSTABLE](UNSTABLE.md) Describes how much funds will be debited from the target  contract balance as a result of the transaction.
type Spending struct {
	// Amount of nanotokens that will be sent to `dst` address.
	Amount big.Int `json:"amount"`
	// Destination address of recipient of funds.
	Dst string `json:"dst"`
}

// [UNSTABLE](UNSTABLE.md) Parameters to init DeBot.
type ParamsOfInit struct {
	// Debot smart contract address.
	Address string `json:"address"`
}

// [UNSTABLE](UNSTABLE.md) Structure for storing debot handle returned from `init` function.
type RegisteredDebot struct {
	// Debot handle which references an instance of debot engine.
	DebotHandle DebotHandle `json:"debot_handle"`
	// Debot abi as json string.
	DebotAbi string `json:"debot_abi"`
	// Debot metadata.
	Info DebotInfo `json:"info"`
}

// [UNSTABLE](UNSTABLE.md) Debot Browser callbacks.
// Called by debot engine to communicate with debot browser.

// Print message to user.
type LogParamsOfAppDebotBrowser struct {
	// A string that must be printed to user.
	Msg string `json:"msg"`
}

// Switch debot to another context (menu).
type SwitchParamsOfAppDebotBrowser struct {
	// Debot context ID to which debot is switched.
	ContextID uint8 `json:"context_id"`
}

// Notify browser that all context actions are shown.
type SwitchCompletedParamsOfAppDebotBrowser struct{}

// Show action to the user. Called after `switch` for each action in context.
type ShowActionParamsOfAppDebotBrowser struct {
	// Debot action that must be shown to user as menu item. At least `description` property must be shown from [DebotAction] structure.
	Action DebotAction `json:"action"`
}

// Request user input.
type InputParamsOfAppDebotBrowser struct {
	// A prompt string that must be printed to user before input request.
	Prompt string `json:"prompt"`
}

// Get signing box to sign data.
// Signing box returned is owned and disposed by debot engine.
type GetSigningBoxParamsOfAppDebotBrowser struct{}

// Execute action of another debot.
type InvokeDebotParamsOfAppDebotBrowser struct {
	// Address of debot in blockchain.
	DebotAddr string `json:"debot_addr"`
	// Debot action to execute.
	Action DebotAction `json:"action"`
}

// Used by Debot to call DInterface implemented by Debot Browser.
type SendParamsOfAppDebotBrowser struct {
	// Internal message to DInterface address.
	// Message body contains interface function and parameters.
	Message string `json:"message"`
}

// Requests permission from DeBot Browser to execute DeBot operation.
type ApproveParamsOfAppDebotBrowser struct {
	// DeBot activity details.
	Activity DebotActivity `json:"activity"`
}

type ParamsOfAppDebotBrowser struct {
	// Should be any of
	// LogParamsOfAppDebotBrowser
	// SwitchParamsOfAppDebotBrowser
	// SwitchCompletedParamsOfAppDebotBrowser
	// ShowActionParamsOfAppDebotBrowser
	// InputParamsOfAppDebotBrowser
	// GetSigningBoxParamsOfAppDebotBrowser
	// InvokeDebotParamsOfAppDebotBrowser
	// SendParamsOfAppDebotBrowser
	// ApproveParamsOfAppDebotBrowser
	EnumTypeValue interface{}
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

	case ApproveParamsOfAppDebotBrowser:
		return json.Marshal(struct {
			ApproveParamsOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"Approve",
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

	case "Approve":
		var enumTypeValue ApproveParamsOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ParamsOfAppDebotBrowser %v", typeDescriptor.Type)
	}

	return nil
}

// [UNSTABLE](UNSTABLE.md) Returning values from Debot Browser callbacks.

// Result of user input.
type InputResultOfAppDebotBrowser struct {
	// String entered by user.
	Value string `json:"value"`
}

// Result of getting signing box.
type GetSigningBoxResultOfAppDebotBrowser struct {
	// Signing box for signing data requested by debot engine.
	// Signing box is owned and disposed by debot engine.
	SigningBox SigningBoxHandle `json:"signing_box"`
}

// Result of debot invoking.
type InvokeDebotResultOfAppDebotBrowser struct{}

// Result of `approve` callback.
type ApproveResultOfAppDebotBrowser struct {
	// Indicates whether the DeBot is allowed to perform the specified operation.
	Approved bool `json:"approved"`
}

type ResultOfAppDebotBrowser struct {
	// Should be any of
	// InputResultOfAppDebotBrowser
	// GetSigningBoxResultOfAppDebotBrowser
	// InvokeDebotResultOfAppDebotBrowser
	// ApproveResultOfAppDebotBrowser
	EnumTypeValue interface{}
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

	case ApproveResultOfAppDebotBrowser:
		return json.Marshal(struct {
			ApproveResultOfAppDebotBrowser
			Type string `json:"type"`
		}{
			value,
			"Approve",
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

	case "Approve":
		var enumTypeValue ApproveResultOfAppDebotBrowser
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ResultOfAppDebotBrowser %v", typeDescriptor.Type)
	}

	return nil
}

// [UNSTABLE](UNSTABLE.md) Parameters to start DeBot. DeBot must be already initialized with init() function.
type ParamsOfStart struct {
	// Debot handle which references an instance of debot engine.
	DebotHandle DebotHandle `json:"debot_handle"`
}

// [UNSTABLE](UNSTABLE.md) Parameters to fetch DeBot metadata.
type ParamsOfFetch struct {
	// Debot smart contract address.
	Address string `json:"address"`
}

// [UNSTABLE](UNSTABLE.md).
type ResultOfFetch struct {
	// Debot metadata.
	Info DebotInfo `json:"info"`
}

// [UNSTABLE](UNSTABLE.md) Parameters for executing debot action.
type ParamsOfExecute struct {
	// Debot handle which references an instance of debot engine.
	DebotHandle DebotHandle `json:"debot_handle"`
	// Debot Action that must be executed.
	Action DebotAction `json:"action"`
}

// [UNSTABLE](UNSTABLE.md) Parameters of `send` function.
type ParamsOfSend struct {
	// Debot handle which references an instance of debot engine.
	DebotHandle DebotHandle `json:"debot_handle"`
	// BOC of internal message to debot encoded in base64 format.
	Message string `json:"message"`
}

// [UNSTABLE](UNSTABLE.md).
type ParamsOfRemove struct {
	// Debot handle which references an instance of debot engine.
	DebotHandle DebotHandle `json:"debot_handle"`
}

// [UNSTABLE](UNSTABLE.md) Creates and instance of DeBot.
// Downloads debot smart contract (code and data) from blockchain and creates
// an instance of Debot Engine for it.
//
// # Remarks
// It does not switch debot to context 0. Browser Callbacks are not called.

func (c *Client) DebotInit(p *ParamsOfInit, app AppDebotBrowser) (*RegisteredDebot, error) { // nolint dupl
	result := new(RegisteredDebot)
	responses, err := c.dllClient.resultsChannel("debot.init", p)
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
				c.dispatchRequestDebotInit(r.Data, app)
			}

			if r.Code == ResponseCodeAppNotify {
				c.dispatchNotifyDebotInit(r.Data, app)
			}
		}
	}()

	return result, nil
}

func (c *Client) dispatchRequestDebotInit(payload []byte, app AppDebotBrowser) { // nolint dupl
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
	var appResponse interface{}
	// appResponse, err := app.Request(appParams)

	switch value := (appParams.EnumTypeValue).(type) {
	case InputParamsOfAppDebotBrowser:
		appResponse, err = app.InputRequest(value)

	case GetSigningBoxParamsOfAppDebotBrowser:
		appResponse, err = app.GetSigningBoxRequest(value)

	case InvokeDebotParamsOfAppDebotBrowser:
		appResponse, err = app.InvokeDebotRequest(value)

	case ApproveParamsOfAppDebotBrowser:
		appResponse, err = app.ApproveRequest(value)

	default:
		err = fmt.Errorf("unsupported type for request %v", appParams.EnumTypeValue)
	}

	appRequestResult := AppRequestResult{}
	if err != nil {
		appRequestResult.EnumTypeValue = ErrorAppRequestResult{Text: err.Error()}
	} else {
		marshalled, _ := json.Marshal(&ResultOfAppDebotBrowser{EnumTypeValue: appResponse})
		appRequestResult.EnumTypeValue = OkAppRequestResult{Result: marshalled}
	}
	err = c.ClientResolveAppRequest(&ParamsOfResolveAppRequest{
		AppRequestID: appRequest.AppRequestID,
		Result:       appRequestResult,
	})
	if err == nil || errors.Is(err, ErrContextIsClosed) {
		return
	}
	panic(err)
}

func (c *Client) dispatchNotifyDebotInit(payload []byte, app AppDebotBrowser) { // nolint dupl
	var appParams ParamsOfAppDebotBrowser
	err := json.Unmarshal(payload, &appParams)
	if err != nil {
		panic(err)
	}

	switch value := (appParams.EnumTypeValue).(type) {
	case LogParamsOfAppDebotBrowser:
		app.LogNotify(value)

	case SwitchParamsOfAppDebotBrowser:
		app.SwitchNotify(value)

	case SwitchCompletedParamsOfAppDebotBrowser:
		app.SwitchCompletedNotify(value)

	case ShowActionParamsOfAppDebotBrowser:
		app.ShowActionNotify(value)

	case SendParamsOfAppDebotBrowser:
		app.SendNotify(value)

	default:
		panic(fmt.Errorf("unsupported type for request %v", appParams.EnumTypeValue))
	}
}

// [UNSTABLE](UNSTABLE.md) Starts the DeBot.
// Downloads debot smart contract from blockchain and switches it to
// context zero.
//
// This function must be used by Debot Browser to start a dialog with debot.
// While the function is executing, several Browser Callbacks can be called,
// since the debot tries to display all actions from the context 0 to the user.
//
// When the debot starts SDK registers `BrowserCallbacks` AppObject.
// Therefore when `debote.remove` is called the debot is being deleted and the callback is called
// with `finish`=`true` which indicates that it will never be used again.
func (c *Client) DebotStart(p *ParamsOfStart) error {
	_, err := c.dllClient.waitErrorOrResult("debot.start", p)

	return err
}

// [UNSTABLE](UNSTABLE.md) Fetches DeBot metadata from blockchain.
// Downloads DeBot from blockchain and creates and fetches its metadata.
func (c *Client) DebotFetch(p *ParamsOfFetch) (*ResultOfFetch, error) {
	result := new(ResultOfFetch)

	err := c.dllClient.waitErrorOrResultUnmarshal("debot.fetch", p, result)

	return result, err
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
func (c *Client) DebotRemove(p *ParamsOfRemove) error {
	_, err := c.dllClient.waitErrorOrResult("debot.remove", p)

	return err
}
