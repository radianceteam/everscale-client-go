package client

type NetQueryCollectionParams struct {
	Collection string `json:"collection"`
	Result     string `json:"result"`
	Filter     string `json:"filter,omitempty"`
	Limit      uint   `json:"limit,omitempty"`
}

func (c *Client) NetQueryCollection(p NetQueryCollectionParams) ([]byte, error) {
	return c.dllClient.waitErrorOrResult("net.query_collection", p)
}
