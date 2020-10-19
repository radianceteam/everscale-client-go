package client

type versionResponse struct {
	Version string `json:"version"`
}

func (c *Client) ClientVersion() (string, error) {
	var version versionResponse
	err := c.dllClient.waitErrorOrResultUnmarshal("client.version", nil, &version)

	return version.Version, err
}

func (c *Client) ClientGetAPIReference() ([]byte, error) {
	return c.dllClient.waitErrorOrResult("client.get_api_reference", nil)
}
