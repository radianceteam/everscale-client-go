package client

// DON'T EDIT THIS FILE is generated 20 Oct 20 13:40 UTC
//
// Mod boc
//
// BOC manipulation module.

type ParamsOfParse struct {
	// BOC encoded as base64.
	Boc string `json:"boc"`
}

type ResultOfParse struct {
	// JSON containing parsed BOC.
	Parsed interface{} `json:"parsed"`
}

type ParamsOfGetBlockchainConfig struct {
	// Key block BOC encoded as base64.
	BlockBoc string `json:"block_boc"`
}

type ResultOfGetBlockchainConfig struct {
	// Blockchain config BOC encoded as base64.
	ConfigBoc string `json:"config_boc"`
}

func (c *Client) BocParseMessage(p *ParamsOfParse) (*ResultOfParse, error) {
	response := new(ResultOfParse)
	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_message", p, response)

	return response, err
}

func (c *Client) BocParseTransaction(p *ParamsOfParse) (*ResultOfParse, error) {
	response := new(ResultOfParse)
	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_transaction", p, response)

	return response, err
}

func (c *Client) BocParseAccount(p *ParamsOfParse) (*ResultOfParse, error) {
	response := new(ResultOfParse)
	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_account", p, response)

	return response, err
}

func (c *Client) BocParseBlock(p *ParamsOfParse) (*ResultOfParse, error) {
	response := new(ResultOfParse)
	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_block", p, response)

	return response, err
}

func (c *Client) BocGetBlockchainConfig(p *ParamsOfGetBlockchainConfig) (*ResultOfGetBlockchainConfig, error) {
	response := new(ResultOfGetBlockchainConfig)
	err := c.dllClient.waitErrorOrResultUnmarshal("boc.get_blockchain_config", p, response)

	return response, err
}
