package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:03:07.683298 +0000 UTC
// Mod utils
//  Misc utility Functions.
//  Misc utility Functions.

type AddressStringFormat interface{}

type ParamsOfConvertAddress struct {
	Address      string              `json:"address"`
	OutputFormat AddressStringFormat `json:"output_format"`
}

type ResultOfConvertAddress struct {
	Address string `json:"address"`
}
