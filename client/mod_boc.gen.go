package client

// DON'T EDIT THIS FILE is generated 28 Oct 20 08:23 UTC
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

// Parses message boc into a JSON
//
// JSON structure is compatible with GraphQL API message object.
func (c *Client) BocParseMessage(p *ParamsOfParse) (*ResultOfParse, error) {
	response := new(ResultOfParse)
	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_message", p, response)

	return response, err
}

// Parses transaction boc into a JSON
//
// JSON structure is compatible with GraphQL API transaction object.
func (c *Client) BocParseTransaction(p *ParamsOfParse) (*ResultOfParse, error) {
	response := new(ResultOfParse)
	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_transaction", p, response)

	return response, err
}

// Parses account boc into a JSON
//
// JSON structure is compatible with GraphQL API account object.
func (c *Client) BocParseAccount(p *ParamsOfParse) (*ResultOfParse, error) {
	response := new(ResultOfParse)
	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_account", p, response)

	return response, err
}

// Parses block boc into a JSON
//
// JSON structure is compatible with GraphQL API block object.
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
