package client

// DON'T EDIT THIS FILE is generated 06 Nov 20 19:56 UTC
//
// Mod net
//
// Network access.

import (
	"github.com/volatiletech/null"
)

type OrderBy struct {
	Path      string        `json:"path"`
	Direction SortDirection `json:"direction"`
}

type SortDirection string

const (
	AscSortDirection  SortDirection = "ASC"
	DescSortDirection SortDirection = "DESC"
)

type ParamsOfQueryCollection struct {
	// Collection name (accounts, blocks, transactions, messages, block_signatures).
	Collection string `json:"collection"`
	// Collection filter.
	Filter interface{} `json:"filter"` // optional
	// Projection (result) string.
	Result string `json:"result"`
	// Sorting order.
	Order []OrderBy `json:"order"` // optional
	// Number of documents to return.
	Limit null.Uint32 `json:"limit"` // optional
}

type ResultOfQueryCollection struct {
	// Objects that match the provided criteria.
	Result []interface{} `json:"result"`
}

type ParamsOfWaitForCollection struct {
	// Collection name (accounts, blocks, transactions, messages, block_signatures).
	Collection string `json:"collection"`
	// Collection filter.
	Filter interface{} `json:"filter"` // optional
	// Projection (result) string.
	Result string `json:"result"`
	// Query timeout.
	Timeout null.Uint32 `json:"timeout"` // optional
}

type ResultOfWaitForCollection struct {
	// First found object that matches the provided criteria.
	Result interface{} `json:"result"`
}

type ResultOfSubscribeCollection struct {
	// Subscription handle. Must be closed with `unsubscribe`.
	Handle uint32 `json:"handle"`
}

type ParamsOfSubscribeCollection struct {
	// Collection name (accounts, blocks, transactions, messages, block_signatures).
	Collection string `json:"collection"`
	// Collection filter.
	Filter interface{} `json:"filter"` // optional
	// Projection (result) string.
	Result string `json:"result"`
}

// Queries collection data
//
// Queries data that satisfies the `filter` conditions,
// limits the number of returned records and orders them.
// The projection fields are limited to `result` fields.
func (c *Client) NetQueryCollection(p *ParamsOfQueryCollection) (*ResultOfQueryCollection, error) {
	response := new(ResultOfQueryCollection)
	err := c.dllClient.waitErrorOrResultUnmarshal("net.query_collection", p, response)

	return response, err
}

// Returns an object that fulfills the conditions or waits for its appearance
//
// Triggers only once.
// If object that satisfies the `filter` conditions
// already exists - returns it immediately.
// If not - waits for insert/update of data within the specified `timeout`,
// and returns it.
// The projection fields are limited to `result` fields.
func (c *Client) NetWaitForCollection(p *ParamsOfWaitForCollection) (*ResultOfWaitForCollection, error) {
	response := new(ResultOfWaitForCollection)
	err := c.dllClient.waitErrorOrResultUnmarshal("net.wait_for_collection", p, response)

	return response, err
}

// Cancels a subscription
//
// Cancels a subscription specified by its handle.
func (c *Client) NetUnsubscribe(p *ResultOfSubscribeCollection) error {
	_, err := c.dllClient.waitErrorOrResult("net.unsubscribe", p)

	return err
}
