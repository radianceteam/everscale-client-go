package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:03:07.68314 +0000 UTC
// Mod processing
//  Message processing module.
//  Message processing module.
//
//  This module incorporates functions related to complex message
//  processing scenarios.

type ProcessingEvent interface{}

type ResultOfProcessMessage struct {
	Transaction interface{}    `json:"transaction"`
	OutMessages []interface{}  `json:"out_messages"`
	Decoded     *DecodedOutput `json:"decoded,omitempty"`
}

type DecodedOutput struct {
	OutMessages []*DecodedMessageBody `json:"out_messages"`
	Output      interface{}           `json:"output,omitempty"`
}

type ParamsOfSendMessage struct {
	Message    string `json:"message"`
	Abi        *Abi   `json:"abi,omitempty"`
	SendEvents bool   `json:"send_events"`
}

type ResultOfSendMessage struct {
	ShardBlockID string `json:"shard_block_id"`
}

type ParamsOfWaitForTransaction struct {
	Abi          *Abi   `json:"abi,omitempty"`
	Message      string `json:"message"`
	ShardBlockID string `json:"shard_block_id"`
	SendEvents   bool   `json:"send_events"`
}

type ParamsOfProcessMessage struct {
	Message    MessageSource `json:"message"`
	SendEvents bool          `json:"send_events"`
}
