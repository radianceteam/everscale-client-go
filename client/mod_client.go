package client

import "encoding/json"

type versionResponse struct {
	Version string `json:"version"`
}

func (c *tonClient) ClientVersion() (string, error) {
	responses := c.dllClient.Request("client.version", nil)
	rawData, err := getFirstErrorOrResult(responses)
	if err != nil {
		return "", err
	}

	var version versionResponse
	if err := json.Unmarshal(rawData, &version); err != nil {
		return "", err
	}

	return version.Version, nil
}

func (c *tonClient) ClientGetAPIReference() ([]byte, error) {
	responses := c.dllClient.Request("client.get_api_reference", nil)

	return getFirstErrorOrResult(responses)
}
