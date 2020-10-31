package client

// DON'T EDIT THIS FILE is generated 31 Oct 20 20:13 UTC
//
// Mod utils
//
// Misc utility Functions.

type AddressStringFormat struct {
	Type   string `json:"type"`
	URL    bool   `json:"url"`
	Test   bool   `json:"test"`
	Bounce bool   `json:"bounce"`
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
