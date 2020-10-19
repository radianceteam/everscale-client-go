package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:44:28.807383 +0000 UTC
// Mod utils
// Misc utility Functions.
// Misc utility Functions.

type AddressStringFormat interface{}

type ParamsOfConvertAddress struct {
	// Account address in any format.
	Address string `json:"address"`
	// Specify the format to convert to.
	OutputFormat AddressStringFormat `json:"output_format"`
}

type ResultOfConvertAddress struct {
	// address in the specified format
	Address string `json:"address"`
}
