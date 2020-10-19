package client

import "encoding/json"

// Client - go client for TON-SDK dll.
type Client interface {
	// mod_client
	// ClientVersion returns version of TON-SDK dll.
	ClientVersion() (string, error)
	// ClientGetAPIReference returns bytes representing json docs.
	ClientGetAPIReference() ([]byte, error)

	// mod_net
	// NetQueryCollection queries net and returns json bytes response.
	NetQueryCollection(NetQueryCollectionParams) ([]byte, error)

	// mod_crypto
	// CryptoGenerateRandomSignKeys generates random ed25519 key pair.
	CryptoGenerateRandomSignKeys() (*CryptoKeyPair, error)

	// Close freeze resources.
	Close()
}

type tonClient struct {
	dllClient
}

func (c *tonClient) Close() {
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

func NewClient(config Config) (Client, error) {
	rawConfig, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	dllClient, err := newDLLClient(rawConfig)
	if err != nil {
		return nil, err
	}

	return &tonClient{dllClient: dllClient}, nil
}
