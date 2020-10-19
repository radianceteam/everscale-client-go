package client

// DON'T EDIT THIS FILE is generated 2020-10-19 11:18:28.294501 +0000 UTC
//
// Mod processing
//
// Message processing module.
//
// This module incorporates functions related to complex message
// processing scenarios.

type ProcessingEvent interface{}

type ResultOfProcessMessage struct {
	// Parsed transaction.
	//
	// In addition to the regular transaction fields there is a
	// `boc` field encoded with `base64` which contains source
	// transaction BOC.
	Transaction interface{} `json:"transaction"`
	// List of parsed output messages.
	//
	// Similar to the `transaction` each message contains the `boc`
	// field.
	OutMessages []interface{} `json:"out_messages"`
	// Optional decoded message bodies according to the optional
	// `abi` parameter.
	Decoded *DecodedOutput `json:"decoded,omitempty"`
}

type DecodedOutput struct {
	// Decoded bodies of the out messages.
	//
	// If the message can't be decoded then `None` will be stored in
	// the appropriate position.
	OutMessages []*DecodedMessageBody `json:"out_messages"`
	// Decoded body of the function output message.
	Output interface{} `json:"output,omitempty"`
}

type ParamsOfSendMessage struct {
	// Message BOC.
	Message string `json:"message"`
	// Optional message ABI.
	//
	// If this parameter is specified and the message has the
	// `expire` header then expiration time will be checked against
	// the current time to prevent an unnecessary sending.
	//
	// The `message already expired` error will be returned in this
	// case.
	//
	// Note that specifying `abi` for ABI compliant contracts is
	// strongly recommended due to choosing proper processing
	// strategy.
	Abi *Abi `json:"abi,omitempty"`
	// Flag for requesting events sending.
	SendEvents bool `json:"send_events"`
}

type ResultOfSendMessage struct {
	// Shard block related to the message dst account before the
	// message had been sent.
	//
	// This block id must be used as a parameter of the
	// `wait_for_transaction`.
	ShardBlockID string `json:"shard_block_id"`
}

type ParamsOfWaitForTransaction struct {
	// Optional ABI for decoding transaction results.
	//
	// If it is specified then the output messages bodies will be
	// decoded according to this ABI.
	//
	// The `abi_decoded` result field will be filled out.
	Abi *Abi `json:"abi,omitempty"`
	// Message BOC. Encoded with `base64`.
	Message string `json:"message"`
	// Dst account shard block id before the message had been sent.
	//
	// You must provide the same value as the `send_message` has
	// returned.
	ShardBlockID string `json:"shard_block_id"`
	// Flag for requesting events sending.
	SendEvents bool `json:"send_events"`
}

type ParamsOfProcessMessage struct {
	// Message source.
	Message MessageSource `json:"message"`
	// Flag for requesting events sending.
	SendEvents bool `json:"send_events"`
}

func (c *Client) SendMessage() {}

// Performs monitoring of the network for a results of the external
// inbound message processing.
//
// Note that presence of the `abi` parameter is critical for ABI
// compliant contracts. Message processing uses drastically
// different strategy for processing message with an ABI expiration
// replay protection.
//
// When the ABI header `expire` is present, the processing uses
// `message expiration` strategy:
// - The maximum block gen time is set to
// `message_expiration_time + transaction_wait_timeout`.
// - When maximum block gen time is reached the processing will
// be finished with `MessageExpired` error.
//
// When the ABI header `expire` isn't present or `abi` parameter
// isn't specified, the processing uses `transaction waiting`
// strategy:
// - The maximum block gen time is set to
// `now() + transaction_wait_timeout`.
// - When maximum block gen time is reached the processing will
// be finished with `Incomplete` result.
func (c *Client) WaitForTransaction() {}

// Sends message to the network and monitors network for a result of
// message processing.
func (c *Client) ProcessMessage() {}
