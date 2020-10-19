package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:44:28.807055 +0000 UTC
// Mod boc
// BOC manipulation module.
// BOC manipulation module.

type ParamsOfParse struct {
	// BOC encoded as base64
	Boc string `json:"boc"`
}

type ResultOfParse struct {
	// JSON containing parsed BOC
	Parsed interface{} `json:"parsed"`
}

type ParamsOfGetBlockchainConfig struct {
	// Key block BOC encoded as base64
	BlockBoc string `json:"block_boc"`
}

type ResultOfGetBlockchainConfig struct {
	// Blockchain config BOC encoded as base64
	ConfigBoc string `json:"config_boc"`
}
