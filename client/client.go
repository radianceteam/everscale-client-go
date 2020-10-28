package client

import (
	"encoding/json"

	"github.com/radianceteam/ton-client-go/spec"
)

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
	Crypto  *CryptoConfig  `json:"crypto,omitempty"`
	ABI     *AbiConfig     `json:"abi,omitempty"`
	Network *NetworkConfig `json:"network,omitempty"`
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

type TransactionFees interface{}

// ClientGetAPIReference loads and parses JSON API spec.
func (c *Client) ClientGetAPIReference() (*spec.APIReference, error) {
	response := new(spec.APIReference)
	err := c.dllClient.waitErrorOrResultUnmarshal("client.get_api_reference", nil, response)

	return response, err
}

// NetQueryCollectionRaw gives ability to unmarshall raw data yourself.
func (c *Client) NetQueryCollectionRaw(p *ParamsOfQueryCollection) ([]byte, error) {
	return c.dllClient.waitErrorOrResult("net.query_collection", p)
}
