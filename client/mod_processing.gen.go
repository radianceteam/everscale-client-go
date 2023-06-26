package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 26 Jun 23 09:46 UTC
//
// Mod processing
//
// Message processing module.
// This module incorporates functions related to complex message
// processing scenarios.

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/volatiletech/null"
)

const (
	MessageAlreadyExpiredProcessingErrorCode           = 501
	MessageHasNotDestinationAddressProcessingErrorCode = 502
	CanNotBuildMessageCellProcessingErrorCode          = 503
	FetchBlockFailedProcessingErrorCode                = 504
	SendMessageFailedProcessingErrorCode               = 505
	InvalidMessageBocProcessingErrorCode               = 506
	MessageExpiredProcessingErrorCode                  = 507
	TransactionWaitTimeoutProcessingErrorCode          = 508
	InvalidBlockReceivedProcessingErrorCode            = 509
	CanNotCheckBlockShardProcessingErrorCode           = 510
	BlockNotFoundProcessingErrorCode                   = 511
	InvalidDataProcessingErrorCode                     = 512
	ExternalSignerMustNotBeUsedProcessingErrorCode     = 513
	MessageRejectedProcessingErrorCode                 = 514
	InvalidRempStatusProcessingErrorCode               = 515
	NextRempStatusTimeoutProcessingErrorCode           = 516
)

func init() { // nolint gochecknoinits
	errorCodesToErrorTypes[MessageAlreadyExpiredProcessingErrorCode] = "MessageAlreadyExpiredProcessingErrorCode"
	errorCodesToErrorTypes[MessageHasNotDestinationAddressProcessingErrorCode] = "MessageHasNotDestinationAddressProcessingErrorCode"
	errorCodesToErrorTypes[CanNotBuildMessageCellProcessingErrorCode] = "CanNotBuildMessageCellProcessingErrorCode"
	errorCodesToErrorTypes[FetchBlockFailedProcessingErrorCode] = "FetchBlockFailedProcessingErrorCode"
	errorCodesToErrorTypes[SendMessageFailedProcessingErrorCode] = "SendMessageFailedProcessingErrorCode"
	errorCodesToErrorTypes[InvalidMessageBocProcessingErrorCode] = "InvalidMessageBocProcessingErrorCode"
	errorCodesToErrorTypes[MessageExpiredProcessingErrorCode] = "MessageExpiredProcessingErrorCode"
	errorCodesToErrorTypes[TransactionWaitTimeoutProcessingErrorCode] = "TransactionWaitTimeoutProcessingErrorCode"
	errorCodesToErrorTypes[InvalidBlockReceivedProcessingErrorCode] = "InvalidBlockReceivedProcessingErrorCode"
	errorCodesToErrorTypes[CanNotCheckBlockShardProcessingErrorCode] = "CanNotCheckBlockShardProcessingErrorCode"
	errorCodesToErrorTypes[BlockNotFoundProcessingErrorCode] = "BlockNotFoundProcessingErrorCode"
	errorCodesToErrorTypes[InvalidDataProcessingErrorCode] = "InvalidDataProcessingErrorCode"
	errorCodesToErrorTypes[ExternalSignerMustNotBeUsedProcessingErrorCode] = "ExternalSignerMustNotBeUsedProcessingErrorCode"
	errorCodesToErrorTypes[MessageRejectedProcessingErrorCode] = "MessageRejectedProcessingErrorCode"
	errorCodesToErrorTypes[InvalidRempStatusProcessingErrorCode] = "InvalidRempStatusProcessingErrorCode"
	errorCodesToErrorTypes[NextRempStatusTimeoutProcessingErrorCode] = "NextRempStatusTimeoutProcessingErrorCode"
}

// Notifies the application that the account's current shard block will be fetched from the network. This step is performed before the message sending so that sdk knows starting from which block it will search for the transaction.
// Fetched block will be used later in waiting phase.
type WillFetchFirstBlockProcessingEvent struct {
	MessageID  string `json:"message_id"`
	MessageDst string `json:"message_dst"`
}

// Notifies the app that the client has failed to fetch the account's current shard block.
// This may happen due to the network issues. Receiving this event means that message processing will not proceed -
// message was not sent, and Developer can try to run `process_message` again,
// in the hope that the connection is restored.
type FetchFirstBlockFailedProcessingEvent struct {
	Error      Error  `json:"error"`
	MessageID  string `json:"message_id"`
	MessageDst string `json:"message_dst"`
}

// Notifies the app that the message will be sent to the network. This event means that the account's current shard block was successfully fetched and the message was successfully created (`abi.encode_message` function was executed successfully).
type WillSendProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	MessageDst   string `json:"message_dst"`
	Message      string `json:"message"`
}

// Notifies the app that the message was sent to the network, i.e `processing.send_message` was successfully executed. Now, the message is in the blockchain. If Application exits at this phase, Developer needs to proceed with processing after the application is restored with `wait_for_transaction` function, passing shard_block_id and message from this event.
// Do not forget to specify abi of your contract as well, it is crucial for processing. See `processing.wait_for_transaction` documentation.
type DidSendProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	MessageDst   string `json:"message_dst"`
	Message      string `json:"message"`
}

// Notifies the app that the sending operation was failed with network error.
// Nevertheless the processing will be continued at the waiting
// phase because the message possibly has been delivered to the
// node.
// If Application exits at this phase, Developer needs to proceed with processing
// after the application is restored with `wait_for_transaction` function, passing
// shard_block_id and message from this event. Do not forget to specify abi of your contract
// as well, it is crucial for processing. See `processing.wait_for_transaction` documentation.
type SendFailedProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	MessageDst   string `json:"message_dst"`
	Message      string `json:"message"`
	Error        Error  `json:"error"`
}

// Notifies the app that the next shard block will be fetched from the network.
// Event can occurs more than one time due to block walking
// procedure.
// If Application exits at this phase, Developer needs to proceed with processing
// after the application is restored with `wait_for_transaction` function, passing
// shard_block_id and message from this event. Do not forget to specify abi of your contract
// as well, it is crucial for processing. See `processing.wait_for_transaction` documentation.
type WillFetchNextBlockProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	MessageDst   string `json:"message_dst"`
	Message      string `json:"message"`
}

// Notifies the app that the next block can't be fetched.
// If no block was fetched within `NetworkConfig.wait_for_timeout` then processing stops.
// This may happen when the shard stops, or there are other network issues.
// In this case Developer should resume message processing with `wait_for_transaction`, passing shard_block_id,
// message and contract abi to it. Note that passing ABI is crucial, because it will influence the processing strategy.
//
// Another way to tune this is to specify long timeout in `NetworkConfig.wait_for_timeout`.
type FetchNextBlockFailedProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	MessageDst   string `json:"message_dst"`
	Message      string `json:"message"`
	Error        Error  `json:"error"`
}

// Notifies the app that the message was not executed within expire timeout on-chain and will never be because it is already expired. The expiration timeout can be configured with `AbiConfig` parameters.
// This event occurs only for the contracts which ABI includes "expire" header.
//
// If Application specifies `NetworkConfig.message_retries_count` > 0, then `process_message`
// will perform retries: will create a new message and send it again and repeat it until it reaches
// the maximum retries count or receives a successful result.  All the processing
// events will be repeated.
type MessageExpiredProcessingEvent struct {
	MessageID  string `json:"message_id"`
	MessageDst string `json:"message_dst"`
	Message    string `json:"message"`
	Error      Error  `json:"error"`
}

// Notifies the app that the message has been delivered to the thread's validators.
type RempSentToValidatorsProcessingEvent struct {
	MessageID  string          `json:"message_id"`
	MessageDst string          `json:"message_dst"`
	Timestamp  big.Int         `json:"timestamp"`
	JSON       json.RawMessage `json:"json"`
}

// Notifies the app that the message has been successfully included into a block candidate by the thread's collator.
type RempIncludedIntoBlockProcessingEvent struct {
	MessageID  string          `json:"message_id"`
	MessageDst string          `json:"message_dst"`
	Timestamp  big.Int         `json:"timestamp"`
	JSON       json.RawMessage `json:"json"`
}

// Notifies the app that the block candidate with the message has been accepted by the thread's validators.
type RempIncludedIntoAcceptedBlockProcessingEvent struct {
	MessageID  string          `json:"message_id"`
	MessageDst string          `json:"message_dst"`
	Timestamp  big.Int         `json:"timestamp"`
	JSON       json.RawMessage `json:"json"`
}

// Notifies the app about some other minor REMP statuses occurring during message processing.
type RempOtherProcessingEvent struct {
	MessageID  string          `json:"message_id"`
	MessageDst string          `json:"message_dst"`
	Timestamp  big.Int         `json:"timestamp"`
	JSON       json.RawMessage `json:"json"`
}

// Notifies the app about any problem that has occurred in REMP processing - in this case library switches to the fallback transaction awaiting scenario (sequential block reading).
type RempErrorProcessingEvent struct {
	MessageID  string `json:"message_id"`
	MessageDst string `json:"message_dst"`
	Error      Error  `json:"error"`
}

type ProcessingEvent struct {
	// Should be any of
	// WillFetchFirstBlockProcessingEvent
	// FetchFirstBlockFailedProcessingEvent
	// WillSendProcessingEvent
	// DidSendProcessingEvent
	// SendFailedProcessingEvent
	// WillFetchNextBlockProcessingEvent
	// FetchNextBlockFailedProcessingEvent
	// MessageExpiredProcessingEvent
	// RempSentToValidatorsProcessingEvent
	// RempIncludedIntoBlockProcessingEvent
	// RempIncludedIntoAcceptedBlockProcessingEvent
	// RempOtherProcessingEvent
	// RempErrorProcessingEvent
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ProcessingEvent) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case WillFetchFirstBlockProcessingEvent:
		return json.Marshal(struct {
			WillFetchFirstBlockProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"WillFetchFirstBlock",
		})

	case FetchFirstBlockFailedProcessingEvent:
		return json.Marshal(struct {
			FetchFirstBlockFailedProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"FetchFirstBlockFailed",
		})

	case WillSendProcessingEvent:
		return json.Marshal(struct {
			WillSendProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"WillSend",
		})

	case DidSendProcessingEvent:
		return json.Marshal(struct {
			DidSendProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"DidSend",
		})

	case SendFailedProcessingEvent:
		return json.Marshal(struct {
			SendFailedProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"SendFailed",
		})

	case WillFetchNextBlockProcessingEvent:
		return json.Marshal(struct {
			WillFetchNextBlockProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"WillFetchNextBlock",
		})

	case FetchNextBlockFailedProcessingEvent:
		return json.Marshal(struct {
			FetchNextBlockFailedProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"FetchNextBlockFailed",
		})

	case MessageExpiredProcessingEvent:
		return json.Marshal(struct {
			MessageExpiredProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"MessageExpired",
		})

	case RempSentToValidatorsProcessingEvent:
		return json.Marshal(struct {
			RempSentToValidatorsProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"RempSentToValidators",
		})

	case RempIncludedIntoBlockProcessingEvent:
		return json.Marshal(struct {
			RempIncludedIntoBlockProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"RempIncludedIntoBlock",
		})

	case RempIncludedIntoAcceptedBlockProcessingEvent:
		return json.Marshal(struct {
			RempIncludedIntoAcceptedBlockProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"RempIncludedIntoAcceptedBlock",
		})

	case RempOtherProcessingEvent:
		return json.Marshal(struct {
			RempOtherProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"RempOther",
		})

	case RempErrorProcessingEvent:
		return json.Marshal(struct {
			RempErrorProcessingEvent
			Type string `json:"type"`
		}{
			value,
			"RempError",
		})

	default:
		return nil, fmt.Errorf("unsupported type for ProcessingEvent %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *ProcessingEvent) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "WillFetchFirstBlock":
		var enumTypeValue WillFetchFirstBlockProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "FetchFirstBlockFailed":
		var enumTypeValue FetchFirstBlockFailedProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "WillSend":
		var enumTypeValue WillSendProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "DidSend":
		var enumTypeValue DidSendProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "SendFailed":
		var enumTypeValue SendFailedProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "WillFetchNextBlock":
		var enumTypeValue WillFetchNextBlockProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "FetchNextBlockFailed":
		var enumTypeValue FetchNextBlockFailedProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "MessageExpired":
		var enumTypeValue MessageExpiredProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "RempSentToValidators":
		var enumTypeValue RempSentToValidatorsProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "RempIncludedIntoBlock":
		var enumTypeValue RempIncludedIntoBlockProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "RempIncludedIntoAcceptedBlock":
		var enumTypeValue RempIncludedIntoAcceptedBlockProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "RempOther":
		var enumTypeValue RempOtherProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "RempError":
		var enumTypeValue RempErrorProcessingEvent
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for ProcessingEvent %v", typeDescriptor.Type)
	}

	return nil
}

type ResultOfProcessMessage struct {
	// Parsed transaction.
	// In addition to the regular transaction fields there is a
	// `boc` field encoded with `base64` which contains source
	// transaction BOC.
	Transaction json.RawMessage `json:"transaction"`
	// List of output messages' BOCs.
	// Encoded as `base64`.
	OutMessages []string `json:"out_messages"`
	// Optional decoded message bodies according to the optional `abi` parameter.
	Decoded *DecodedOutput `json:"decoded"` // optional
	// Transaction fees.
	Fees TransactionFees `json:"fees"`
}

type DecodedOutput struct {
	// Decoded bodies of the out messages.
	// If the message can't be decoded, then `None` will be stored in
	// the appropriate position.
	OutMessages []*DecodedMessageBody `json:"out_messages"`
	// Decoded body of the function output message.
	Output json.RawMessage `json:"output"` // optional
}

type MessageMonitoringTransactionCompute struct {
	// Compute phase exit code.
	ExitCode int32 `json:"exit_code"`
}

type MessageMonitoringTransaction struct {
	// Hash of the transaction. Present if transaction was included into the blocks. When then transaction was emulated this field will be missing.
	Hash null.String `json:"hash"` // optional
	// Aborted field of the transaction.
	Aborted bool `json:"aborted"`
	// Optional information about the compute phase of the transaction.
	Compute *MessageMonitoringTransactionCompute `json:"compute"` // optional
}

type MessageMonitoringParams struct {
	// Monitored message identification. Can be provided as a message's BOC or (hash, address) pair. BOC is a preferable way because it helps to determine possible error reason (using TVM execution of the message).
	Message MonitoredMessage `json:"message"`
	// Block time Must be specified as a UNIX timestamp in seconds.
	WaitUntil uint32 `json:"wait_until"`
	// User defined data associated with this message. Helps to identify this message when user received `MessageMonitoringResult`.
	UserData json.RawMessage `json:"user_data"` // optional
}

type MessageMonitoringResult struct {
	// Hash of the message.
	Hash string `json:"hash"`
	// Processing status.
	Status MessageMonitoringStatus `json:"status"`
	// In case of `Finalized` the transaction is extracted from the block. In case of `Timeout` the transaction is emulated using the last known account state.
	Transaction *MessageMonitoringTransaction `json:"transaction"` // optional
	// In case of `Timeout` contains possible error reason.
	Error null.String `json:"error"` // optional
	// User defined data related to this message. This is the same value as passed before with `MessageMonitoringParams` or `SendMessageParams`.
	UserData json.RawMessage `json:"user_data"` // optional
}

type MonitorFetchWaitMode string

const (

	// If there are no resolved results yet, then monitor awaits for the next resolved result.
	AtLeastOneMonitorFetchWaitMode MonitorFetchWaitMode = "AtLeastOne"
	// Monitor waits until all unresolved messages will be resolved. If there are no unresolved messages then monitor will wait.
	AllMonitorFetchWaitMode    MonitorFetchWaitMode = "All"
	NoWaitMonitorFetchWaitMode MonitorFetchWaitMode = "NoWait"
)

// BOC of the message.
type BocMonitoredMessage struct {
	Boc string `json:"boc"`
}

// Message's hash and destination address.
type HashAddressMonitoredMessage struct {
	// Hash of the message.
	Hash string `json:"hash"`
	// Destination address of the message.
	Address string `json:"address"`
}

type MonitoredMessage struct {
	// Should be any of
	// BocMonitoredMessage
	// HashAddressMonitoredMessage
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *MonitoredMessage) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case BocMonitoredMessage:
		return json.Marshal(struct {
			BocMonitoredMessage
			Type string `json:"type"`
		}{
			value,
			"Boc",
		})

	case HashAddressMonitoredMessage:
		return json.Marshal(struct {
			HashAddressMonitoredMessage
			Type string `json:"type"`
		}{
			value,
			"HashAddress",
		})

	default:
		return nil, fmt.Errorf("unsupported type for MonitoredMessage %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *MonitoredMessage) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "Boc":
		var enumTypeValue BocMonitoredMessage
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "HashAddress":
		var enumTypeValue HashAddressMonitoredMessage
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for MonitoredMessage %v", typeDescriptor.Type)
	}

	return nil
}

type MessageMonitoringStatus string

const (

	// Returned when the messages was processed and included into finalized block before `wait_until` block time.
	FinalizedMessageMonitoringStatus MessageMonitoringStatus = "Finalized"
	// Returned when the message was not processed until `wait_until` block time.
	TimeoutMessageMonitoringStatus MessageMonitoringStatus = "Timeout"
	// Reserved for future statuses.
	// Is never returned. Application should wait for one of the `Finalized` or `Timeout` statuses.
	// All other statuses are intermediate.
	ReservedMessageMonitoringStatus MessageMonitoringStatus = "Reserved"
)

type MessageSendingParams struct {
	// BOC of the message, that must be sent to the blockchain.
	Boc string `json:"boc"`
	// Expiration time of the message. Must be specified as a UNIX timestamp in seconds.
	WaitUntil uint32 `json:"wait_until"`
	// User defined data associated with this message. Helps to identify this message when user received `MessageMonitoringResult`.
	UserData json.RawMessage `json:"user_data"` // optional
}

type ParamsOfMonitorMessages struct {
	// Name of the monitoring queue.
	Queue string `json:"queue"`
	// Messages to start monitoring for.
	Messages []MessageMonitoringParams `json:"messages"`
}

type ParamsOfGetMonitorInfo struct {
	// Name of the monitoring queue.
	Queue string `json:"queue"`
}

type MonitoringQueueInfo struct {
	// Count of the unresolved messages.
	Unresolved uint32 `json:"unresolved"`
	// Count of resolved results.
	Resolved uint32 `json:"resolved"`
}

type ParamsOfFetchNextMonitorResults struct {
	// Name of the monitoring queue.
	Queue string `json:"queue"`
	// Wait mode.
	// Default is `NO_WAIT`.
	WaitMode *MonitorFetchWaitMode `json:"wait_mode"` // optional
}

type ResultOfFetchNextMonitorResults struct {
	// List of the resolved results.
	Results []MessageMonitoringResult `json:"results"`
}

type ParamsOfCancelMonitor struct {
	// Name of the monitoring queue.
	Queue string `json:"queue"`
}

type ParamsOfSendMessages struct {
	// Messages that must be sent to the blockchain.
	Messages []MessageSendingParams `json:"messages"`
	// Optional message monitor queue that starts monitoring for the processing results for sent messages.
	MonitorQueue null.String `json:"monitor_queue"` // optional
}

type ResultOfSendMessages struct {
	// Messages that was sent to the blockchain for execution.
	Messages []MessageMonitoringParams `json:"messages"`
}

type ParamsOfSendMessage struct {
	// Message BOC.
	Message string `json:"message"`
	// Optional message ABI.
	// If this parameter is specified and the message has the
	// `expire` header then expiration time will be checked against
	// the current time to prevent unnecessary sending of already expired message.
	//
	// The `message already expired` error will be returned in this
	// case.
	//
	// Note, that specifying `abi` for ABI compliant contracts is
	// strongly recommended, so that proper processing strategy can be
	// chosen.
	Abi *Abi `json:"abi"` // optional
	// Flag for requesting events sending.
	SendEvents bool `json:"send_events"`
}

type ResultOfSendMessage struct {
	// The last generated shard block of the message destination account before the message was sent.
	// This block id must be used as a parameter of the
	// `wait_for_transaction`.
	ShardBlockID string `json:"shard_block_id"`
	// The list of endpoints to which the message was sent.
	// This list id must be used as a parameter of the
	// `wait_for_transaction`.
	SendingEndpoints []string `json:"sending_endpoints"`
}

type ParamsOfWaitForTransaction struct {
	// Optional ABI for decoding the transaction result.
	// If it is specified, then the output messages' bodies will be
	// decoded according to this ABI.
	//
	// The `abi_decoded` result field will be filled out.
	Abi *Abi `json:"abi"` // optional
	// Message BOC.
	// Encoded with `base64`.
	Message string `json:"message"`
	// The last generated block id of the destination account shard before the message was sent.
	// You must provide the same value as the `send_message` has returned.
	ShardBlockID string `json:"shard_block_id"`
	// Flag that enables/disables intermediate events.
	SendEvents bool `json:"send_events"`
	// The list of endpoints to which the message was sent.
	// Use this field to get more informative errors.
	// Provide the same value as the `send_message` has returned.
	// If the message was not delivered (expired), SDK will log the endpoint URLs, used for its sending.
	SendingEndpoints []string `json:"sending_endpoints"` // optional
}

type ParamsOfProcessMessage struct {
	// Message encode parameters.
	MessageEncodeParams ParamsOfEncodeMessage `json:"message_encode_params"`
	// Flag for requesting events sending.
	SendEvents bool `json:"send_events"`
}

// Starts monitoring for the processing results of the specified messages.
// Message monitor performs background monitoring for a message processing results
// for the specified set of messages.
//
// Message monitor can serve several isolated monitoring queues.
// Each monitor queue has a unique application defined identifier (or name) used
// to separate several queue's.
//
// There are two important lists inside of the monitoring queue:
//
// - unresolved messages: contains messages requested by the application for monitoring
// and not yet resolved;
//
// - resolved results: contains resolved processing results for monitored messages.
//
// Each monitoring queue tracks own unresolved and resolved lists.
// Application can add more messages to the monitoring queue at any time.
//
// Message monitor accumulates resolved results.
// Application should fetch this results with `fetchNextMonitorResults` function.
//
// When both unresolved and resolved lists becomes empty, monitor stops any background activity
// and frees all allocated internal memory.
//
// If monitoring queue with specified name already exists then messages will be added
// to the unresolved list.
//
// If monitoring queue with specified name does not exist then monitoring queue will be created
// with specified unresolved messages.
func (c *Client) ProcessingMonitorMessages(p *ParamsOfMonitorMessages) error {
	_, err := c.dllClient.waitErrorOrResult("processing.monitor_messages", p)

	return err
}

// Returns summary information about current state of the specified monitoring queue.
func (c *Client) ProcessingGetMonitorInfo(p *ParamsOfGetMonitorInfo) (*MonitoringQueueInfo, error) {
	result := new(MonitoringQueueInfo)

	err := c.dllClient.waitErrorOrResultUnmarshal("processing.get_monitor_info", p, result)

	return result, err
}

// Fetches next resolved results from the specified monitoring queue.
// Results and waiting options are depends on the `wait` parameter.
// All returned results will be removed from the queue's resolved list.
func (c *Client) ProcessingFetchNextMonitorResults(p *ParamsOfFetchNextMonitorResults) (*ResultOfFetchNextMonitorResults, error) {
	result := new(ResultOfFetchNextMonitorResults)

	err := c.dllClient.waitErrorOrResultUnmarshal("processing.fetch_next_monitor_results", p, result)

	return result, err
}

// Cancels all background activity and releases all allocated system resources for the specified monitoring queue.
func (c *Client) ProcessingCancelMonitor(p *ParamsOfCancelMonitor) error {
	_, err := c.dllClient.waitErrorOrResult("processing.cancel_monitor", p)

	return err
}

// Sends specified messages to the blockchain.
func (c *Client) ProcessingSendMessages(p *ParamsOfSendMessages) (*ResultOfSendMessages, error) {
	result := new(ResultOfSendMessages)

	err := c.dllClient.waitErrorOrResultUnmarshal("processing.send_messages", p, result)

	return result, err
}
