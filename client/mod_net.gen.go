package client

// DON'T EDIT THIS FILE is generated 31 Jan 21 10:48 UTC
//
// Mod net
//
// Network access.

import (
	"encoding/json"

	"github.com/volatiletech/null"
)

type NetErrorCode string

const (
	QueryFailedNetErrorCode                 NetErrorCode = "QueryFailed"
	SubscribeFailedNetErrorCode             NetErrorCode = "SubscribeFailed"
	WaitForFailedNetErrorCode               NetErrorCode = "WaitForFailed"
	GetSubscriptionResultFailedNetErrorCode NetErrorCode = "GetSubscriptionResultFailed"
	InvalidServerResponseNetErrorCode       NetErrorCode = "InvalidServerResponse"
	ClockOutOfSyncNetErrorCode              NetErrorCode = "ClockOutOfSync"
	WaitForTimeoutNetErrorCode              NetErrorCode = "WaitForTimeout"
	GraphqlErrorNetErrorCode                NetErrorCode = "GraphqlError"
	NetworkModuleSuspendedNetErrorCode      NetErrorCode = "NetworkModuleSuspended"
	WebsocketDisconnectedNetErrorCode       NetErrorCode = "WebsocketDisconnected"
	NotSupportedNetErrorCode                NetErrorCode = "NotSupported"
	NoEndpointsProvidedNetErrorCode         NetErrorCode = "NoEndpointsProvided"
	GraphqlWebsocketInitErrorNetErrorCode   NetErrorCode = "GraphqlWebsocketInitError"
	NetworkModuleResumedNetErrorCode        NetErrorCode = "NetworkModuleResumed"
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
	// Must be a map with named values that can be used in query.
	Variables json.RawMessage `json:"variables"` // optional
}

type ResultOfQuery struct {
	// Result provided by DAppServer.
	Result json.RawMessage `json:"result"`
}

type ParamsOfBatchQuery struct {
	// List of query operations that must be performed per single fetch.
	Operations []ParamsOfQueryOperation `json:"operations"`
}

type ResultOfBatchQuery struct {
	// Result values for batched queries.
	// Returns an array of values. Each value corresponds to `queries` item.
	Results []json.RawMessage `json:"results"`
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

type ParamsOfAggregateCollection struct {
	// Collection name (accounts, blocks, transactions, messages, block_signatures).
	Collection string `json:"collection"`
	// Collection filter.
	Filter json.RawMessage `json:"filter"` // optional
	// Projection (result) string.
	Fields []FieldAggregation `json:"fields"` // optional
}

type ResultOfAggregateCollection struct {
	// Values for requested fields.
	// Returns an array of strings. Each string refers to the corresponding `fields` item.
	// Numeric value is returned as a decimal string representations.
	Values json.RawMessage `json:"values"`
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

type EndpointsSet struct {
	// List of endpoints provided by server.
	Endpoints []string `json:"endpoints"`
}

// Performs DAppServer GraphQL query.
func (c *Client) NetQuery(p *ParamsOfQuery) (*ResultOfQuery, error) {
	result := new(ResultOfQuery)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.query", p, result)

	return result, err
}

// Performs multiple queries per single fetch.
func (c *Client) NetBatchQuery(p *ParamsOfBatchQuery) (*ResultOfBatchQuery, error) {
	result := new(ResultOfBatchQuery)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.batch_query", p, result)

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

// Aggregates collection data.
// Aggregates values from the specified `fields` for records
// that satisfies the `filter` conditions,.
func (c *Client) NetAggregateCollection(p *ParamsOfAggregateCollection) (*ResultOfAggregateCollection, error) {
	result := new(ResultOfAggregateCollection)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.aggregate_collection", p, result)

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

// Requests the list of alternative endpoints from server.
func (c *Client) NetFetchEndpoints() (*EndpointsSet, error) {
	result := new(EndpointsSet)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.fetch_endpoints", nil, result)

	return result, err
}

// Sets the list of endpoints to use on reinit.
func (c *Client) NetSetEndpoints(p *EndpointsSet) error {
	_, err := c.dllClient.waitErrorOrResult("net.set_endpoints", p)

	return err
}
