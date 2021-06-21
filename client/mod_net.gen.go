package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 21 Jun 21 14:33 UTC
//
// Mod net
//
// Network access.

import (
	"encoding/json"
	"fmt"

	"github.com/volatiletech/null"
)

const (
	QueryFailedNetErrorCode                 = 601
	SubscribeFailedNetErrorCode             = 602
	WaitForFailedNetErrorCode               = 603
	GetSubscriptionResultFailedNetErrorCode = 604
	InvalidServerResponseNetErrorCode       = 605
	ClockOutOfSyncNetErrorCode              = 606
	WaitForTimeoutNetErrorCode              = 607
	GraphqlErrorNetErrorCode                = 608
	NetworkModuleSuspendedNetErrorCode      = 609
	WebsocketDisconnectedNetErrorCode       = 610
	NotSupportedNetErrorCode                = 611
	NoEndpointsProvidedNetErrorCode         = 612
	GraphqlWebsocketInitErrorNetErrorCode   = 613
	NetworkModuleResumedNetErrorCode        = 614
)

func init() { // nolint gochecknoinits
	errorCodesToErrorTypes[QueryFailedNetErrorCode] = "QueryFailedNetErrorCode"
	errorCodesToErrorTypes[SubscribeFailedNetErrorCode] = "SubscribeFailedNetErrorCode"
	errorCodesToErrorTypes[WaitForFailedNetErrorCode] = "WaitForFailedNetErrorCode"
	errorCodesToErrorTypes[GetSubscriptionResultFailedNetErrorCode] = "GetSubscriptionResultFailedNetErrorCode"
	errorCodesToErrorTypes[InvalidServerResponseNetErrorCode] = "InvalidServerResponseNetErrorCode"
	errorCodesToErrorTypes[ClockOutOfSyncNetErrorCode] = "ClockOutOfSyncNetErrorCode"
	errorCodesToErrorTypes[WaitForTimeoutNetErrorCode] = "WaitForTimeoutNetErrorCode"
	errorCodesToErrorTypes[GraphqlErrorNetErrorCode] = "GraphqlErrorNetErrorCode"
	errorCodesToErrorTypes[NetworkModuleSuspendedNetErrorCode] = "NetworkModuleSuspendedNetErrorCode"
	errorCodesToErrorTypes[WebsocketDisconnectedNetErrorCode] = "WebsocketDisconnectedNetErrorCode"
	errorCodesToErrorTypes[NotSupportedNetErrorCode] = "NotSupportedNetErrorCode"
	errorCodesToErrorTypes[NoEndpointsProvidedNetErrorCode] = "NoEndpointsProvidedNetErrorCode"
	errorCodesToErrorTypes[GraphqlWebsocketInitErrorNetErrorCode] = "GraphqlWebsocketInitErrorNetErrorCode"
	errorCodesToErrorTypes[NetworkModuleResumedNetErrorCode] = "NetworkModuleResumedNetErrorCode"
}

type OrderBy struct {
	Path      string        `json:"path"`
	Direction SortDirection `json:"direction"`
}

type SortDirection string

const (
	AscSortDirection  SortDirection = "ASC"
	DescSortDirection SortDirection = "DESC"
)

type ParamsOfQueryOperation struct {
	// Should be any of
	// ParamsOfQueryCollection
	// ParamsOfWaitForCollection
	// ParamsOfAggregateCollection
	// ParamsOfQueryCounterparties
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ParamsOfQueryOperation) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case ParamsOfQueryCollection:
		return json.Marshal(struct {
			ParamsOfQueryCollection
			Type string `json:"type"`
		}{
			value,
			"QueryCollection",
		})

	case ParamsOfWaitForCollection:
		return json.Marshal(struct {
			ParamsOfWaitForCollection
			Type string `json:"type"`
		}{
			value,
			"WaitForCollection",
		})

	case ParamsOfAggregateCollection:
		return json.Marshal(struct {
			ParamsOfAggregateCollection
			Type string `json:"type"`
		}{
			value,
			"AggregateCollection",
		})

	case ParamsOfQueryCounterparties:
		return json.Marshal(struct {
			ParamsOfQueryCounterparties
			Type string `json:"type"`
		}{
			value,
			"QueryCounterparties",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ParamsOfQueryOperation %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ParamsOfQueryOperation) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "QueryCollection":
		var enumTypeValue ParamsOfQueryCollection
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "WaitForCollection":
		var enumTypeValue ParamsOfWaitForCollection
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "AggregateCollection":
		var enumTypeValue ParamsOfAggregateCollection
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "QueryCounterparties":
		var enumTypeValue ParamsOfQueryCounterparties
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ParamsOfQueryOperation %v", typeDescriptor.Type)
	}

	return nil
}

type FieldAggregation struct {
	// Dot separated path to the field.
	Field string `json:"field"`
	// Aggregation function that must be applied to field values.
	Fn AggregationFn `json:"fn"`
}

type AggregationFn string

const (

	// Returns count of filtered record.
	CountAggregationFn AggregationFn = "COUNT"
	// Returns the minimal value for a field in filtered records.
	MinAggregationFn AggregationFn = "MIN"
	// Returns the maximal value for a field in filtered records.
	MaxAggregationFn AggregationFn = "MAX"
	// Returns a sum of values for a field in filtered records.
	SumAggregationFn AggregationFn = "SUM"
	// Returns an average value for a field in filtered records.
	AverageAggregationFn AggregationFn = "AVERAGE"
)

type TransactionNode struct {
	// Transaction id.
	ID string `json:"id"`
	// In message id.
	InMsg string `json:"in_msg"`
	// Out message ids.
	OutMsgs []string `json:"out_msgs"`
	// Account address.
	AccountAddr string `json:"account_addr"`
	// Transactions total fees.
	TotalFees string `json:"total_fees"`
	// Aborted flag.
	Aborted bool `json:"aborted"`
	// Compute phase exit code.
	ExitCode null.Uint32 `json:"exit_code"` // optional
}

type MessageNode struct {
	// Message id.
	ID string `json:"id"`
	// Source transaction id.
	// This field is missing for an external inbound messages.
	SrcTransactionID null.String `json:"src_transaction_id"` // optional
	// Destination transaction id.
	// This field is missing for an external outbound messages.
	DstTransactionID null.String `json:"dst_transaction_id"` // optional
	// Source address.
	Src null.String `json:"src"` // optional
	// Destination address.
	Dst null.String `json:"dst"` // optional
	// Transferred tokens value.
	Value null.String `json:"value"` // optional
	// Bounce flag.
	Bounce bool `json:"bounce"`
	// Decoded body.
	// Library tries to decode message body using provided `params.abi_registry`.
	// This field will be missing if none of the provided abi can be used to decode.
	DecodedBody *DecodedMessageBody `json:"decoded_body"` // optional
}

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

type ResultOfGetEndpoints struct {
	// Current query endpoint.
	Query string `json:"query"`
	// List of all endpoints used by client.
	Endpoints []string `json:"endpoints"`
}

type ParamsOfQueryCounterparties struct {
	// Account address.
	Account string `json:"account"`
	// Projection (result) string.
	Result string `json:"result"`
	// Number of counterparties to return.
	First null.Uint32 `json:"first"` // optional
	// `cursor` field of the last received result.
	After null.String `json:"after"` // optional
}

type ParamsOfQueryTransactionTree struct {
	// Input message id.
	InMsg string `json:"in_msg"`
	// List of contract ABIs that will be used to decode message bodies. Library will try to decode each returned message body using any ABI from the registry.
	AbiRegistry []Abi `json:"abi_registry"` // optional
	// Timeout used to limit waiting time for the missing messages and transaction.
	// If some of the following messages and transactions are missing yet
	// The maximum waiting time is regulated by this option.
	//
	// Default value is 60000 (1 min).
	Timeout null.Uint32 `json:"timeout"` // optional
}

type ResultOfQueryTransactionTree struct {
	// Messages.
	Messages []MessageNode `json:"messages"`
	// Transactions.
	Transactions []TransactionNode `json:"transactions"`
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

// Requests the list of alternative endpoints from server.
func (c *Client) NetGetEndpoints() (*ResultOfGetEndpoints, error) {
	result := new(ResultOfGetEndpoints)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.get_endpoints", nil, result)

	return result, err
}

// Allows to query and paginate through the list of accounts that the specified account has interacted with, sorted by the time of the last internal message between accounts.
// *Attention* this query retrieves data from 'Counterparties' service which is not supported in
// the opensource version of DApp Server (and will not be supported) as well as in TON OS SE (will be supported in SE in future),
// but is always accessible via [TON OS Devnet/Mainnet Clouds](https://docs.ton.dev/86757ecb2/p/85c869-networks).
func (c *Client) NetQueryCounterparties(p *ParamsOfQueryCounterparties) (*ResultOfQueryCollection, error) {
	result := new(ResultOfQueryCollection)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.query_counterparties", p, result)

	return result, err
}

// Returns transactions tree for specific message.
// Performs recursive retrieval of the transactions tree produced by the specific message:
// in_msg -> dst_transaction -> out_messages -> dst_transaction -> ...
//
// All retrieved messages and transactions will be included
// into `result.messages` and `result.transactions` respectively.
//
// The retrieval process will stop when the retrieved transaction count is more than 50.
//
// It is guaranteed that each message in `result.messages` has the corresponding transaction
// in the `result.transactions`.
//
// But there are no guaranties that all messages from transactions `out_msgs` are
// presented in `result.messages`.
// So the application have to continue retrieval for missing messages if it requires.
func (c *Client) NetQueryTransactionTree(p *ParamsOfQueryTransactionTree) (*ResultOfQueryTransactionTree, error) {
	result := new(ResultOfQueryTransactionTree)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.query_transaction_tree", p, result)

	return result, err
}
