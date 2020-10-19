package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:03:07.683006 +0000 UTC
// Mod boc
//  BOC manipulation module.
//  BOC manipulation module.

type ParamsOfParse struct {
	Boc string `json:"boc"`
}

type ResultOfParse struct {
	Parsed interface{} `json:"parsed"`
}

type ParamsOfGetBlockchainConfig struct {
	BlockBoc string `json:"block_boc"`
}

type ResultOfGetBlockchainConfig struct {
	ConfigBoc string `json:"config_boc"`
}
