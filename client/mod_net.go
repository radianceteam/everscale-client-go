package client

import "encoding/json"

type NetQueryCollectionParams struct {
	Collection string `json:"collection"`
	Result string `json:"result"`
	Filter string `json:"filter,omitempty"`
	Limit uint `json:"limit,omitempty"`
}

func (c *tonClient) NetQueryCollection(p NetQueryCollectionParams) ([]byte, error)  {
	rawParams, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return c.dllClient.Request("net.query_collection", rawParams)
}
