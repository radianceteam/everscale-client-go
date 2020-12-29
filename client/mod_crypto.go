package client

// CryptoRegisterSigningBox Register an application implemented signing box.
func (c *Client) CryptoRegisterSigningBox() (*RegisteredSigningBox, error) {
	response := new(RegisteredSigningBox)
	err := c.dllClient.waitErrorOrResultUnmarshal("crypto.register_signing_box", nil, response)

	return response, err
}
