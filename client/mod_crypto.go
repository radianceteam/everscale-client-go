package client

type CryptoKeyPair struct {
	Public string `json:"public"`
	Secret string `json:"secret"`
}

func (c *Client) CryptoGenerateRandomSignKeys() (*CryptoKeyPair, error) {
	var keyPair CryptoKeyPair

	return &keyPair, c.dllClient.waitErrorOrResultUnmarshal("crypto.generate_random_sign_keys", nil, &keyPair)
}
