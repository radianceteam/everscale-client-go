package client

import "encoding/json"

type Client interface {
	NetQueryCollection(NetQueryCollectionParams) ([]byte, error)
	CryptoGenerateRandomSignKeys() (*CryptoKeyPair, error)
	Close()
}

type tonClient struct {
	dllClient
}

func (c *tonClient) Close() {
	c.dllClient.Close()
}

type ConfigNetwork struct {
	ServerAddress string `json:"server_address"`
}

type Config struct{
	Network ConfigNetwork `json:"network"`
	Crypto struct {
		
	} `json:"crypto"`
	ABI struct {
		
	} `json:"abi"`
}

func NewClient(config Config) (Client, error) {
	rawConfig, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	dllClient, err := NewDLLClient(rawConfig)
	if err != nil {
		return nil, err
	}

	return &tonClient{dllClient: dllClient}, nil
}
