package client

import "encoding/json"

type CryptoKeyPair struct {
	Public string `json:"public"`
	Secret string `json:"secret"`
}

func (c *tonClient) CryptoGenerateRandomSignKeys() (*CryptoKeyPair, error) {
	rawData, err := c.dllClient.Request("crypto.generate_random_sign_keys", nil)
	if err != nil {
		return nil, err
	}
	var result CryptoKeyPair
	if err := json.Unmarshal(rawData, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
