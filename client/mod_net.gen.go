package client

// DON'T EDIT THIS FILE is generated 20 Oct 20 13:40 UTC
//
// Mod net
//
// Network access.

import (
	"gopkg.in/guregu/null.v4"
)

type OrderBy struct {
	Path      string        `json:"path"`
	Direction SortDirection `json:"direction"`
}

type SortDirection string

const (
	ASC  SortDirection = "ASC"
	DESC SortDirection = "DESC"
)

type ParamsOfQueryCollection struct {
	// collection name (accounts, blocks, transactions, messages, block_signatures).
	Collection string `json:"collection"`
	// collection filter.
	Filter interface{} `json:"filter"` // optional
	// projection (result) string.
	Result string `json:"result"`
	// sorting order.
	Order []OrderBy `json:"order"` // optional
	// number of documents to return.
	Limit null.Int `json:"limit"` // optional
}

type ResultOfQueryCollection struct {
	// objects that match provided criteria.
	Result []interface{} `json:"result"`
}

type ParamsOfWaitForCollection struct {
	// collection name (accounts, blocks, transactions, messages, block_signatures).
	Collection string `json:"collection"`
	// collection filter.
	Filter interface{} `json:"filter"` // optional
	// projection (result) string.
	Result string `json:"result"`
	// query timeout.
	Timeout null.Int `json:"timeout"` // optional
}

type ResultOfWaitForCollection struct {
	// first found object that match provided criteria.
	Result interface{} `json:"result"`
}

type ResultOfSubscribeCollection struct {
	// handle to subscription. It then can be used in `get_next_subscription_data` function
	// and must be closed with `unsubscribe`.
	Handle int `json:"handle"`
}

type ParamsOfSubscribeCollection struct {
	// collection name (accounts, blocks, transactions, messages, block_signatures).
	Collection string `json:"collection"`
	// collection filter.
	Filter interface{} `json:"filter"` // optional
	// projection (result) string.
	Result string `json:"result"`
}

func (c *Client) NetQueryCollection(p *ParamsOfQueryCollection) (*ResultOfQueryCollection, error) {
	response := new(ResultOfQueryCollection)
	err := c.dllClient.waitErrorOrResultUnmarshal("net.query_collection", p, response)

	return response, err
}

func (c *Client) NetWaitForCollection(p *ParamsOfWaitForCollection) (*ResultOfWaitForCollection, error) {
	response := new(ResultOfWaitForCollection)
	err := c.dllClient.waitErrorOrResultUnmarshal("net.wait_for_collection", p, response)

	return response, err
}

func (c *Client) NetUnsubscribe(p *ResultOfSubscribeCollection) error {
	_, err := c.dllClient.waitErrorOrResult("net.unsubscribe", p)

	return err
}
