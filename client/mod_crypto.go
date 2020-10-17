package client

import "encoding/json"

type CryptoKeyPair struct {
	Public string `json:"public"`
	Secret string `json:"secret"`
}

func (c *tonClient) CryptoGenerateRandomSignKeys() (*CryptoKeyPair, error) {
	responses := c.dllClient.Request("crypto.generate_random_sign_keys", nil)
	rawData, err := getFirstErrorOrResult(responses)
	if err != nil {
		return nil, err
	}

	var keyPair CryptoKeyPair
	if err := json.Unmarshal(rawData, &keyPair); err != nil {
		return nil, err
	}

	return &keyPair, nil
}
