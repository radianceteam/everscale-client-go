package client

import (
	"encoding/json"

	"github.com/radianceteam/ton-client-go/spec"
)

func I(v int) *int {
	r := new(int)
	*r = v

	return r
}

func S(v string) *string {
	r := new(string)
	*r = v

	return r
}

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

func (c *Client) ClientGetAPIReference() (*spec.APIReference, error) {
	response := new(spec.APIReference)
	err := c.dllClient.waitErrorOrResultUnmarshal("client.get_api_reference", nil, response)

	return response, err
}
