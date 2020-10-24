package client

// DON'T EDIT THIS FILE is generated 24 Oct 20 12:36 UTC
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
	// List of output messages' BOCs. Encoded as `base64`.
	OutMessages []string `json:"out_messages"`
	// Optional decoded message bodies according to the optional
	// `abi` parameter.
	Decoded *DecodedOutput `json:"decoded"` // optional
	// Transaction fees.
	Fees TransactionFees `json:"fees"`
}

type DecodedOutput struct {
	// Decoded bodies of the out messages.
	//
	// If the message can't be decoded then `None` will be stored in
	// the appropriate position.
	OutMessages []*DecodedMessageBody `json:"out_messages"`
	// Decoded body of the function output message.
	Output interface{} `json:"output"` // optional
}

type ParamsOfSendMessage struct {
	// Message BOC.
	Message string `json:"message"`
	// Optional message ABI.
	//
	// If this parameter is specified and the message has the
	// `expire` header then expiration time will be checked against
	// the current time to prevent an unnecessary sending of already expired message.
	//
	// The `message already expired` error will be returned in this
	// case.
	//
	// Note that specifying `abi` for ABI compliant contracts is
	// strongly recommended due to choosing proper processing
	// strategy.
	Abi *Abi `json:"abi"` // optional
	// Flag for requesting events sending.
	SendEvents bool `json:"send_events"`
}

type ResultOfSendMessage struct {
	// The last generated shard block of the message destination account before the
	// message was sent.
	//
	// This block id must be used as a parameter of the
	// `wait_for_transaction`.
	ShardBlockID string `json:"shard_block_id"`
}

type ParamsOfWaitForTransaction struct {
	// Optional ABI for decoding the transaction result.
	//
	// If it is specified then the output messages' bodies will be
	// decoded according to this ABI.
	//
	// The `abi_decoded` result field will be filled out.
	Abi *Abi `json:"abi"` // optional
	// Message BOC. Encoded with `base64`.
	Message string `json:"message"`
	// The last generated block id of the destination account shard before the message was sent.
	//
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
