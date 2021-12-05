package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 05 Dec 21 04:40 UTC
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

type ParamsOfCreateBlockIterator struct {
	// Starting time to iterate from.
	// If the application specifies this parameter then the iteration
	// includes blocks with `gen_utime` >= `start_time`.
	// Otherwise the iteration starts from zero state.
	//
	// Must be specified in seconds.
	StartTime null.Uint32 `json:"start_time"` // optional
	// Optional end time to iterate for.
	// If the application specifies this parameter then the iteration
	// includes blocks with `gen_utime` < `end_time`.
	// Otherwise the iteration never stops.
	//
	// Must be specified in seconds.
	EndTime null.Uint32 `json:"end_time"` // optional
	// Shard prefix filter.
	// If the application specifies this parameter and it is not the empty array
	// then the iteration will include items related to accounts that belongs to
	// the specified shard prefixes.
	// Shard prefix must be represented as a string "workchain:prefix".
	// Where `workchain` is a signed integer and the `prefix` if a hexadecimal
	// representation if the 64-bit unsigned integer with tagged shard prefix.
	// For example: "0:3800000000000000".
	ShardFilter []string `json:"shard_filter"` // optional
	// Projection (result) string.
	// List of the fields that must be returned for iterated items.
	// This field is the same as the `result` parameter of
	// the `query_collection` function.
	// Note that iterated items can contains additional fields that are
	// not requested in the `result`.
	Result null.String `json:"result"` // optional
}

type RegisteredIterator struct {
	// Iterator handle.
	// Must be removed using `remove_iterator`
	// when it is no more needed for the application.
	Handle uint32 `json:"handle"`
}

type ParamsOfResumeBlockIterator struct {
	// Iterator state from which to resume.
	// Same as value returned from `iterator_next`.
	ResumeState json.RawMessage `json:"resume_state"`
}

type ParamsOfCreateTransactionIterator struct {
	// Starting time to iterate from.
	// If the application specifies this parameter then the iteration
	// includes blocks with `gen_utime` >= `start_time`.
	// Otherwise the iteration starts from zero state.
	//
	// Must be specified in seconds.
	StartTime null.Uint32 `json:"start_time"` // optional
	// Optional end time to iterate for.
	// If the application specifies this parameter then the iteration
	// includes blocks with `gen_utime` < `end_time`.
	// Otherwise the iteration never stops.
	//
	// Must be specified in seconds.
	EndTime null.Uint32 `json:"end_time"` // optional
	// Shard prefix filters.
	// If the application specifies this parameter and it is not an empty array
	// then the iteration will include items related to accounts that belongs to
	// the specified shard prefixes.
	// Shard prefix must be represented as a string "workchain:prefix".
	// Where `workchain` is a signed integer and the `prefix` if a hexadecimal
	// representation if the 64-bit unsigned integer with tagged shard prefix.
	// For example: "0:3800000000000000".
	// Account address conforms to the shard filter if
	// it belongs to the filter workchain and the first bits of address match to
	// the shard prefix. Only transactions with suitable account addresses are iterated.
	ShardFilter []string `json:"shard_filter"` // optional
	// Account address filter.
	// Application can specify the list of accounts for which
	// it wants to iterate transactions.
	//
	// If this parameter is missing or an empty list then the library iterates
	// transactions for all accounts that pass the shard filter.
	//
	// Note that the library doesn't detect conflicts between the account filter and the shard filter
	// if both are specified.
	// So it is an application responsibility to specify the correct filter combination.
	AccountsFilter []string `json:"accounts_filter"` // optional
	// Projection (result) string.
	// List of the fields that must be returned for iterated items.
	// This field is the same as the `result` parameter of
	// the `query_collection` function.
	// Note that iterated items can contain additional fields that are
	// not requested in the `result`.
	Result null.String `json:"result"` // optional
	// Include `transfers` field in iterated transactions.
	// If this parameter is `true` then each transaction contains field
	// `transfers` with list of transfer. See more about this structure in function description.
	IncludeTransfers null.Bool `json:"include_transfers"` // optional
}

type ParamsOfResumeTransactionIterator struct {
	// Iterator state from which to resume.
	// Same as value returned from `iterator_next`.
	ResumeState json.RawMessage `json:"resume_state"`
	// Account address filter.
	// Application can specify the list of accounts for which
	// it wants to iterate transactions.
	//
	// If this parameter is missing or an empty list then the library iterates
	// transactions for all accounts that passes the shard filter.
	//
	// Note that the library doesn't detect conflicts between the account filter and the shard filter
	// if both are specified.
	// So it is the application's responsibility to specify the correct filter combination.
	AccountsFilter []string `json:"accounts_filter"` // optional
}

type ParamsOfIteratorNext struct {
	// Iterator handle.
	Iterator uint32 `json:"iterator"`
	// Maximum count of the returned items.
	// If value is missing or is less than 1 the library uses 1.
	Limit null.Uint32 `json:"limit"` // optional
	// Indicates that function must return the iterator state that can be used for resuming iteration.
	ReturnResumeState null.Bool `json:"return_resume_state"` // optional
}

type ResultOfIteratorNext struct {
	// Next available items.
	// Note that `iterator_next` can return an empty items and `has_more` equals to `true`.
	// In this case the application have to continue iteration.
	// Such situation can take place when there is no data yet but
	// the requested `end_time` is not reached.
	Items []json.RawMessage `json:"items"`
	// Indicates that there are more available items in iterated range.
	HasMore bool `json:"has_more"`
	// Optional iterator state that can be used for resuming iteration.
	// This field is returned only if the `return_resume_state` parameter
	// is specified.
	//
	// Note that `resume_state` corresponds to the iteration position
	// after the returned items.
	ResumeState json.RawMessage `json:"resume_state"` // optional
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

// Returns a tree of transactions triggered by a specific message.
// Performs recursive retrieval of a transactions tree produced by a specific message:
// in_msg -> dst_transaction -> out_messages -> dst_transaction -> ...
// If the chain of transactions execution is in progress while the function is running,
// it will wait for the next transactions to appear until the full tree or more than 50 transactions
// are received.
//
// All the retrieved messages and transactions are included
// into `result.messages` and `result.transactions` respectively.
//
// Function reads transactions layer by layer, by pages of 20 transactions.
//
// The retrieval prosess goes like this:
// Let's assume we have an infinite chain of transactions and each transaction generates 5 messages.
// 1. Retrieve 1st message (input parameter) and corresponding transaction - put it into result.
// It is the first level of the tree of transactions - its root.
// Retrieve 5 out message ids from the transaction for next steps.
// 2. Retrieve 5 messages and corresponding transactions on the 2nd layer. Put them into result.
// Retrieve 5*5 out message ids from these transactions for next steps
// 3. Retrieve 20 (size of the page) messages and transactions (3rd layer) and 20*5=100 message ids (4th layer).
// 4. Retrieve the last 5 messages and 5 transactions on the 3rd layer + 15 messages and transactions (of 100) from the 4th layer
// + 25 message ids of the 4th layer + 75 message ids of the 5th layer.
// 5. Retrieve 20 more messages and 20 more transactions of the 4th layer + 100 more message ids of the 5th layer.
// 6. Now we have 1+5+20+20+20 = 66 transactions, which is more than 50. Function exits with the tree of
// 1m->1t->5m->5t->25m->25t->35m->35t. If we see any message ids in the last transactions out_msgs, which don't have
// corresponding messages in the function result, it means that the full tree was not received and we need to continue iteration.
//
// To summarize, it is guaranteed that each message in `result.messages` has the corresponding transaction
// in the `result.transactions`.
// But there is no guarantee that all messages from transactions `out_msgs` are
// presented in `result.messages`.
// So the application has to continue retrieval for missing messages if it requires.
func (c *Client) NetQueryTransactionTree(p *ParamsOfQueryTransactionTree) (*ResultOfQueryTransactionTree, error) {
	result := new(ResultOfQueryTransactionTree)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.query_transaction_tree", p, result)

	return result, err
}

// Creates block iterator.
// Block iterator uses robust iteration methods that guaranties that every
// block in the specified range isn't missed or iterated twice.
//
// Iterated range can be reduced with some filters:
// - `start_time` – the bottom time range. Only blocks with `gen_utime`
// more or equal to this value is iterated. If this parameter is omitted then there is
// no bottom time edge, so all blocks since zero state is iterated.
// - `end_time` – the upper time range. Only blocks with `gen_utime`
// less then this value is iterated. If this parameter is omitted then there is
// no upper time edge, so iterator never finishes.
// - `shard_filter` – workchains and shard prefixes that reduce the set of interesting
// blocks. Block conforms to the shard filter if it belongs to the filter workchain
// and the first bits of block's `shard` fields matches to the shard prefix.
// Only blocks with suitable shard are iterated.
//
// Items iterated is a JSON objects with block data. The minimal set of returned
// fields is:
// ```text
// id
// gen_utime
// workchain_id
// shard
// after_split
// after_merge
// prev_ref {
// root_hash
// }
// prev_alt_ref {
// root_hash
// }
// ```
// Application can request additional fields in the `result` parameter.
//
// Application should call the `remove_iterator` when iterator is no longer required.
func (c *Client) NetCreateBlockIterator(p *ParamsOfCreateBlockIterator) (*RegisteredIterator, error) {
	result := new(RegisteredIterator)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.create_block_iterator", p, result)

	return result, err
}

// Resumes block iterator.
// The iterator stays exactly at the same position where the `resume_state` was catched.
//
// Application should call the `remove_iterator` when iterator is no longer required.
func (c *Client) NetResumeBlockIterator(p *ParamsOfResumeBlockIterator) (*RegisteredIterator, error) {
	result := new(RegisteredIterator)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.resume_block_iterator", p, result)

	return result, err
}

// Creates transaction iterator.
// Transaction iterator uses robust iteration methods that guaranty that every
// transaction in the specified range isn't missed or iterated twice.
//
// Iterated range can be reduced with some filters:
// - `start_time` – the bottom time range. Only transactions with `now`
// more or equal to this value are iterated. If this parameter is omitted then there is
// no bottom time edge, so all the transactions since zero state are iterated.
// - `end_time` – the upper time range. Only transactions with `now`
// less then this value are iterated. If this parameter is omitted then there is
// no upper time edge, so iterator never finishes.
// - `shard_filter` – workchains and shard prefixes that reduce the set of interesting
// accounts. Account address conforms to the shard filter if
// it belongs to the filter workchain and the first bits of address match to
// the shard prefix. Only transactions with suitable account addresses are iterated.
// - `accounts_filter` – set of account addresses whose transactions must be iterated.
// Note that accounts filter can conflict with shard filter so application must combine
// these filters carefully.
//
// Iterated item is a JSON objects with transaction data. The minimal set of returned
// fields is:
// ```text
// id
// account_addr
// now
// balance_delta(format:DEC)
// bounce { bounce_type }
// in_message {
// id
// value(format:DEC)
// msg_type
// src
// }
// out_messages {
// id
// value(format:DEC)
// msg_type
// dst
// }
// ```
// Application can request an additional fields in the `result` parameter.
//
// Another parameter that affects on the returned fields is the `include_transfers`.
// When this parameter is `true` the iterator computes and adds `transfer` field containing
// list of the useful `TransactionTransfer` objects.
// Each transfer is calculated from the particular message related to the transaction
// and has the following structure:
// - message – source message identifier.
// - isBounced – indicates that the transaction is bounced, which means the value will be returned back to the sender.
// - isDeposit – indicates that this transfer is the deposit (true) or withdraw (false).
// - counterparty – account address of the transfer source or destination depending on `isDeposit`.
// - value – amount of nano tokens transferred. The value is represented as a decimal string
// because the actual value can be more precise than the JSON number can represent. Application
// must use this string carefully – conversion to number can follow to loose of precision.
//
// Application should call the `remove_iterator` when iterator is no longer required.
func (c *Client) NetCreateTransactionIterator(p *ParamsOfCreateTransactionIterator) (*RegisteredIterator, error) {
	result := new(RegisteredIterator)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.create_transaction_iterator", p, result)

	return result, err
}

// Resumes transaction iterator.
// The iterator stays exactly at the same position where the `resume_state` was caught.
// Note that `resume_state` doesn't store the account filter. If the application requires
// to use the same account filter as it was when the iterator was created then the application
// must pass the account filter again in `accounts_filter` parameter.
//
// Application should call the `remove_iterator` when iterator is no longer required.
func (c *Client) NetResumeTransactionIterator(p *ParamsOfResumeTransactionIterator) (*RegisteredIterator, error) {
	result := new(RegisteredIterator)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.resume_transaction_iterator", p, result)

	return result, err
}

// Returns next available items.
// In addition to available items this function returns the `has_more` flag
// indicating that the iterator isn't reach the end of the iterated range yet.
//
// This function can return the empty list of available items but
// indicates that there are more items is available.
// This situation appears when the iterator doesn't reach iterated range
// but database doesn't contains available items yet.
//
// If application requests resume state in `return_resume_state` parameter
// then this function returns `resume_state` that can be used later to
// resume the iteration from the position after returned items.
//
// The structure of the items returned depends on the iterator used.
// See the description to the appropriated iterator creation function.
func (c *Client) NetIteratorNext(p *ParamsOfIteratorNext) (*ResultOfIteratorNext, error) {
	result := new(ResultOfIteratorNext)

	err := c.dllClient.waitErrorOrResultUnmarshal("net.iterator_next", p, result)

	return result, err
}

// Removes an iterator.
// Frees all resources allocated in library to serve iterator.
//
// Application always should call the `remove_iterator` when iterator
// is no longer required.
func (c *Client) NetRemoveIterator(p *RegisteredIterator) error {
	_, err := c.dllClient.waitErrorOrResult("net.remove_iterator", p)

	return err
}
