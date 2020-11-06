package client

// DON'T EDIT THIS FILE is generated 06 Nov 20 19:25 UTC
//
// Mod utils
//
// Misc utility Functions.

type AddressStringFormatType string

const (
	AccountIDAddressStringFormatType AddressStringFormatType = "AccountId"
	HexAddressStringFormatType       AddressStringFormatType = "Hex"
	Base64AddressStringFormatType    AddressStringFormatType = "Base64"
)

type AddressStringFormat struct {
	Type AddressStringFormatType `json:"type"`
	// presented in types:
	// "Base64".
	URL bool `json:"url"`
	// presented in types:
	// "Base64".
	Test bool `json:"test"`
	// presented in types:
	// "Base64".
	Bounce bool `json:"bounce"`
}

type ParamsOfConvertAddress struct {
	// Account address in any TON format.
	Address string `json:"address"`
	// Specify the format to convert to.
	OutputFormat AddressStringFormat `json:"output_format"`
}

type ResultOfConvertAddress struct {
	// Address in the specified format.
	Address string `json:"address"`
}

// Converts address from any TON format to any TON format.
func (c *Client) UtilsConvertAddress(p *ParamsOfConvertAddress) (*ResultOfConvertAddress, error) {
	response := new(ResultOfConvertAddress)
	err := c.dllClient.waitErrorOrResultUnmarshal("utils.convert_address", p, response)

	return response, err
}
