package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 28 Aug 21 05:19 UTC
//
// Mod processing
//
// Message processing module.
// This module incorporates functions related to complex message
// processing scenarios.

import (
	"encoding/json"
	"fmt"
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
}

// Notifies the application that the account's current shard block will be fetched from the network. This step is performed before the message sending so that sdk knows starting from which block it will search for the transaction.
// Fetched block will be used later in waiting phase.
type WillFetchFirstBlockProcessingEvent struct{}

// Notifies the app that the client has failed to fetch the account's current shard block.
// This may happen due to the network issues. Receiving this event means that message processing will not proceed -
// message was not sent, and Developer can try to run `process_message` again,
// in the hope that the connection is restored.
type FetchFirstBlockFailedProcessingEvent struct {
	Error Error `json:"error"`
}

// Notifies the app that the message will be sent to the network. This event means that the account's current shard block was successfully fetched and the message was successfully created (`abi.encode_message` function was executed successfully).
type WillSendProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	Message      string `json:"message"`
}

// Notifies the app that the message was sent to the network, i.e `processing.send_message` was successfully executed. Now, the message is in the blockchain. If Application exits at this phase, Developer needs to proceed with processing after the application is restored with `wait_for_transaction` function, passing shard_block_id and message from this event.
// Do not forget to specify abi of your contract as well, it is crucial for processing. See `processing.wait_for_transaction` documentation.
type DidSendProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
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
	MessageID string `json:"message_id"`
	Message   string `json:"message"`
	Error     Error  `json:"error"`
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
