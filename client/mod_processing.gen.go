package client

// DON'T EDIT THIS FILE is generated 03 Jan 21 12:23 UTC
//
// Mod processing
//
// Message processing module.
// This module incorporates functions related to complex message
// processing scenarios.

import (
	"encoding/json"
)

type ProcessingEventType string

const (

	// Notifies the app that the current shard block will be fetched from the network.
	// Fetched block will be used later in waiting phase.
	WillFetchFirstBlockProcessingEventType ProcessingEventType = "WillFetchFirstBlock"
	// Notifies the app that the client has failed to fetch current shard block.
	// Message processing has finished.
	FetchFirstBlockFailedProcessingEventType ProcessingEventType = "FetchFirstBlockFailed"
	// Notifies the app that the message will be sent to the network.
	WillSendProcessingEventType ProcessingEventType = "WillSend"
	// Notifies the app that the message was sent to the network.
	DidSendProcessingEventType ProcessingEventType = "DidSend"
	// Notifies the app that the sending operation was failed with network error.
	// Nevertheless the processing will be continued at the waiting
	// phase because the message possibly has been delivered to the
	// node.
	SendFailedProcessingEventType ProcessingEventType = "SendFailed"
	// Notifies the app that the next shard block will be fetched from the network.
	// Event can occurs more than one time due to block walking
	// procedure.
	WillFetchNextBlockProcessingEventType ProcessingEventType = "WillFetchNextBlock"
	// Notifies the app that the next block can't be fetched due to error.
	// Processing will be continued after `network_resume_timeout`.
	FetchNextBlockFailedProcessingEventType ProcessingEventType = "FetchNextBlockFailed"
	// Notifies the app that the message was expired.
	// Event occurs for contracts which ABI includes header "expire"
	//
	// Processing will be continued from encoding phase after
	// `expiration_retries_timeout`.
	MessageExpiredProcessingEventType ProcessingEventType = "MessageExpired"
)

type ProcessingEvent struct {
	Type ProcessingEventType `json:"type"`
	// presented in types:
	// "FetchFirstBlockFailed"
	// "SendFailed"
	// "FetchNextBlockFailed"
	// "MessageExpired".
	Error ClientError `json:"error"`
	// presented in types:
	// "WillSend"
	// "DidSend"
	// "SendFailed"
	// "WillFetchNextBlock"
	// "FetchNextBlockFailed".
	ShardBlockID string `json:"shard_block_id"`
	// presented in types:
	// "WillSend"
	// "DidSend"
	// "SendFailed"
	// "WillFetchNextBlock"
	// "FetchNextBlockFailed"
	// "MessageExpired".
	MessageID string `json:"message_id"`
	// presented in types:
	// "WillSend"
	// "DidSend"
	// "SendFailed"
	// "WillFetchNextBlock"
	// "FetchNextBlockFailed"
	// "MessageExpired".
	Message string `json:"message"`
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
}

type ParamsOfProcessMessage struct {
	// Message encode parameters.
	MessageEncodeParams ParamsOfEncodeMessage `json:"message_encode_params"`
	// Flag for requesting events sending.
	SendEvents bool `json:"send_events"`
}
