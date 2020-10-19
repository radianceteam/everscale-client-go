package client

// DON'T EDIT THIS FILE is generated 2020-10-19 11:18:28.294357 +0000 UTC
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

func (c *Client) ParseMessage()        {}
func (c *Client) ParseTransaction()    {}
func (c *Client) ParseAccount()        {}
func (c *Client) ParseBlock()          {}
func (c *Client) GetBlockchainConfig() {}
