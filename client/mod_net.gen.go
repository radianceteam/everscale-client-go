package client

// DON'T EDIT THIS FILE is generated 03 Jan 21 17:31 UTC
//
// Mod net
//
// Network access.

import (
	"encoding/json"

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

type ParamsOfQuery struct {
	// GraphQL query text.
	Query string `json:"query"`
	// Variables used in query.
	// Must be a map with named values thatcan be used in query.
	Variables json.RawMessage `json:"variables"` // optional
}

type ResultOfQuery struct {
	// Result provided by DAppServer.
	Result json.RawMessage `json:"result"`
}

type ParamsOfQueryCollection struct {
	// Collection name (accounts, blocks, transactions, messages, block_signatures).
	Collection string `json:"collection"`
	// Collection filter.
	Filter json.RawMessage `json:"filter"` // optional
	// Projection (result) string.
	Result string `json:"result"`
	// Sorting order.
	Order []OrderBy `json:"order"` // optional
	// Number of documents to return.
	Limit null.Uint32 `json:"limit"` // optional
}

type ResultOfQueryCollection struct {
	// Objects that match the provided criteria.
	Result []json.RawMessage `json:"result"`
}

type ParamsOfWaitForCollection struct {
	// Collection name (accounts, blocks, transactions, messages, block_signatures).
	Collection string `json:"collection"`
	// Collection filter.
	Filter json.RawMessage `json:"filter"` // optional
	// Projection (result) string.
	Result string `json:"result"`
	// Query timeout.
	Timeout null.Uint32 `json:"timeout"` // optional
}

type ResultOfWaitForCollection struct {
	// First found object that matches the provided criteria.
	Result json.RawMessage `json:"result"`
}

type ResultOfSubscribeCollection struct {
	// Subscription handle.
	// Must be closed with `unsubscribe`.
	Handle uint32 `json:"handle"`
}

type ParamsOfSubscribeCollection struct {
	// Collection name (accounts, blocks, transactions, messages, block_signatures).
	Collection string `json:"collection"`
	// Collection filter.
	Filter json.RawMessage `json:"filter"` // optional
	// Projection (result) string.
	Result string `json:"result"`
}

type ParamsOfFindLastShardBlock struct {
	// Account address.
	Address string `json:"address"`
}

type ResultOfFindLastShardBlock struct {
	// Account shard last block ID.
	BlockID string `json:"block_id"`
}

// Performs DAppServer GraphQL query.
func (c *Client) NetQuery(p *ParamsOfQuery) (*ResultOfQuery, error) {
	result := new(ResultOfQuery)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.query", p, result)

	return result, err
}

// Queries collection data.
// Queries data that satisfies the `filter` conditions,
// limits the number of returned records and orders them.
// The projection fields are limited to `result` fields.
func (c *Client) NetQueryCollection(p *ParamsOfQueryCollection) (*ResultOfQueryCollection, error) {
	result := new(ResultOfQueryCollection)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.query_collection", p, result)

	return result, err
}

// Returns an object that fulfills the conditions or waits for its appearance.
// Triggers only once.
// If object that satisfies the `filter` conditions
// already exists - returns it immediately.
// If not - waits for insert/update of data within the specified `timeout`,
// and returns it.
// The projection fields are limited to `result` fields.
func (c *Client) NetWaitForCollection(p *ParamsOfWaitForCollection) (*ResultOfWaitForCollection, error) {
	result := new(ResultOfWaitForCollection)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.wait_for_collection", p, result)

	return result, err
}

// Cancels a subscription specified by its handle.
func (c *Client) NetUnsubscribe(p *ResultOfSubscribeCollection) error {
	_, err := c.dllClient.waitErrorOrResult("net.unsubscribe", p)

	return err
}

// Suspends network module to stop any network activity.
func (c *Client) NetSuspend() error {
	_, err := c.dllClient.waitErrorOrResult("net.suspend", nil)

	return err
}

// Resumes network module to enable network activity.
func (c *Client) NetResume() error {
	_, err := c.dllClient.waitErrorOrResult("net.resume", nil)

	return err
}

// Returns ID of the last block in a specified account shard.
func (c *Client) NetFindLastShardBlock(p *ParamsOfFindLastShardBlock) (*ResultOfFindLastShardBlock, error) {
	result := new(ResultOfFindLastShardBlock)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.find_last_shard_block", p, result)

	return result, err
}
