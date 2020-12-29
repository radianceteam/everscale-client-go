package client

// DON'T EDIT THIS FILE is generated 29 Dec 20 20:08 UTC
//
// Mod debot
//
// [UNSTABLE](UNSTABLE.md) Module for working with debot.

type DebotHandle uint32

type DebotAction struct {
	// A short action description.
	// Should be used by Debot Browser as name ofmenu item.
	Description string `json:"description"`
	// Depends on action type.
	// Can be a debot function name or a print string(for Print Action).
	Name string `json:"name"`
	// Action type.
	ActionType uint8 `json:"action_type"`
	// ID of debot context to switch after action execution.
	To uint8 `json:"to"`
	// Action attributes.
	// In the form of "param=value,flag".attribute example: instant, args, fargs, sign.
	Attributes string `json:"attributes"`
	// Some internal action data.
	// Used by debot only.
	Misc string `json:"misc"`
}

type ParamsOfStart struct {
	// Debot smart contract address.
	Address string `json:"address"`
}

type RegisteredDebot struct {
	// Debot handle which references an instance of debot engine.
	DebotHandle DebotHandle `json:"debot_handle"`
}

type ParamsOfAppDebotBrowserType string

const (

	// Print message to user.
	LogParamsOfAppDebotBrowserType ParamsOfAppDebotBrowserType = "Log"
	// Switch debot to another context (menu).
	SwitchParamsOfAppDebotBrowserType ParamsOfAppDebotBrowserType = "Switch"
	// Notify browser that all context actions are shown.
	SwitchCompletedParamsOfAppDebotBrowserType ParamsOfAppDebotBrowserType = "SwitchCompleted"
	// Show action to the user. Called after `switch` for each action in context.
	ShowActionParamsOfAppDebotBrowserType ParamsOfAppDebotBrowserType = "ShowAction"
	// Request user input.
	InputParamsOfAppDebotBrowserType ParamsOfAppDebotBrowserType = "Input"
	// Get signing box to sign data.
	// Signing box returned is owned and disposed by debot engine.
	GetSigningBoxParamsOfAppDebotBrowserType ParamsOfAppDebotBrowserType = "GetSigningBox"
	// Execute action of another debot.
	InvokeDebotParamsOfAppDebotBrowserType ParamsOfAppDebotBrowserType = "InvokeDebot"
)

type ParamsOfAppDebotBrowser struct {
	Type ParamsOfAppDebotBrowserType `json:"type"`
	// A string that must be printed to user.
	// presented in types:
	// "Log".
	Msg string `json:"msg"`
	// Debot context ID to which debot is switched.
	// presented in types:
	// "Switch".
	ContextID uint8 `json:"context_id"`
	// Debot action that must be shown to user as menu item. At least `description` property must be shown from [DebotAction] structure.
	// presented in types:
	// "ShowAction"
	// "InvokeDebot".
	Action DebotAction `json:"action"`
	// A prompt string that must be printed to user before input request.
	// presented in types:
	// "Input".
	Prompt string `json:"prompt"`
	// Address of debot in blockchain.
	// presented in types:
	// "InvokeDebot".
	DebotAddr string `json:"debot_addr"`
}

type ResultOfAppDebotBrowserType string

const (

	// Result of user input.
	InputResultOfAppDebotBrowserType ResultOfAppDebotBrowserType = "Input"
	// Result of getting signing box.
	GetSigningBoxResultOfAppDebotBrowserType ResultOfAppDebotBrowserType = "GetSigningBox"
	// Result of debot invoking.
	InvokeDebotResultOfAppDebotBrowserType ResultOfAppDebotBrowserType = "InvokeDebot"
)

type ResultOfAppDebotBrowser struct {
	Type ResultOfAppDebotBrowserType `json:"type"`
	// String entered by user.
	// presented in types:
	// "Input".
	Value string `json:"value"`
	// Signing box for signing data requested by debot engine.
	// Signing box is owned and disposed by debot engine presented in types:
	// "GetSigningBox".
	SigningBox SigningBoxHandle `json:"signing_box"`
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

// [UNSTABLE](UNSTABLE.md) Destroys debot handle.
// Removes handle from Client Context and drops debot engine referenced by that handle.
func (c *Client) DebotRemove(p *RegisteredDebot) error {
	_, err := c.dllClient.waitErrorOrResult("debot.remove", p)

	return err
}
