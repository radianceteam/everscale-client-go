package client

// DON'T EDIT THIS FILE is generated 20 Oct 20 13:40 UTC
//
// Mod utils
//
// Misc utility Functions.

type AddressStringFormat interface{}

type ParamsOfConvertAddress struct {
	// Account address in any format.
	Address string `json:"address"`
	// Specify the format to convert to.
	OutputFormat AddressStringFormat `json:"output_format"`
}

type ResultOfConvertAddress struct {
	// address in the specified format.
	Address string `json:"address"`
}

// Sends message to the network and monitors network for a result of
// message processing.
func (c *Client) UtilsConvertAddress(p *ParamsOfConvertAddress) (*ResultOfConvertAddress, error) {
	response := new(ResultOfConvertAddress)
	err := c.dllClient.waitErrorOrResultUnmarshal("utils.convert_address", p, response)

	return response, err
}
