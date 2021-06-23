package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 23 Jun 21 21:13 UTC
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

// Notifies the app that the current shard block will be fetched from the network.
// Fetched block will be used later in waiting phase.
type WillFetchFirstBlockProcessingEvent struct{}

// Notifies the app that the client has failed to fetch current shard block.
// Message processing has finished.
type FetchFirstBlockFailedProcessingEvent struct {
	Error Error `json:"error"`
}

// Notifies the app that the message will be sent to the network.
type WillSendProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	Message      string `json:"message"`
}

// Notifies the app that the message was sent to the network.
type DidSendProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	Message      string `json:"message"`
}

// Notifies the app that the sending operation was failed with network error.
// Nevertheless the processing will be continued at the waiting
// phase because the message possibly has been delivered to the
// node.
type SendFailedProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	Message      string `json:"message"`
	Error        Error  `json:"error"`
}

// Notifies the app that the next shard block will be fetched from the network.
// Event can occurs more than one time due to block walking
// procedure.
type WillFetchNextBlockProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	Message      string `json:"message"`
}

// Notifies the app that the next block can't be fetched due to error.
// Processing will be continued after `network_resume_timeout`.
type FetchNextBlockFailedProcessingEvent struct {
	ShardBlockID string `json:"shard_block_id"`
	MessageID    string `json:"message_id"`
	Message      string `json:"message"`
	Error        Error  `json:"error"`
}

// Notifies the app that the message was expired.
// Event occurs for contracts which ABI includes header "expire"
//
// Processing will be continued from encoding phase after
// `expiration_retries_timeout`.
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
