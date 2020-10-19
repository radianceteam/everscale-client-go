package client

import "encoding/json"

type Client struct {
	dllClient
}

func (c *Client) Close() {
	if c == nil {
		return
	}
	c.dllClient.close()
}

type Config struct {
	Crypto  CryptoConfig  `json:"crypto"`
	ABI     AbiConfig     `json:"abi"`
	Network NetworkConfig `json:"network"`
}

func NewClient(config Config) (*Client, error) {
	rawConfig, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	dllClient, err := newDLLClient(rawConfig)
	if err != nil {
		return nil, err
	}

	return &Client{dllClient: dllClient}, nil
}
