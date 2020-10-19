package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:49:51.496215 +0000 UTC
// Mod tvm
//
//

import (
	"github.com/shopspring/decimal"
)

type ExecutionMode string

const (
	Full    ExecutionMode = "Full"
	TvmOnly ExecutionMode = "TvmOnly"
)

type ExecutionOptions struct {
	// boc with config
	BlockchainConfig *string `json:"blockchain_config,omitempty"`
	// time that is used as transaction time
	BlockTime *int `json:"block_time,omitempty"`
	// block logical time
	BlockLt *decimal.Decimal `json:"block_lt,omitempty"`
	// transaction logical time
	TransactionLt *decimal.Decimal `json:"transaction_lt,omitempty"`
}

type ParamsOfExecuteMessage struct {
	// Input message.
	Message MessageSource `json:"message"`
	// Account BOC. Must be encoded as base64.
	Account string `json:"account"`
	// Execution mode.
	Mode ExecutionMode `json:"mode"`
	// Execution options.
	ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
}

type ResultOfExecuteMessage struct {
	// Parsed transaction.
	//
	// In addition to the regular transaction fields there is a
	// `boc` field encoded with `base64` which contains source
	// transaction BOC.
	Transaction interface{} `json:"transaction,omitempty"`
	// List of parsed output messages.
	//
	// Similar to the `transaction` each message contains the `boc`
	// field.
	OutMessages []interface{} `json:"out_messages"`
	// Optional decoded message bodies according to the optional
	// `abi` parameter.
	Decoded *DecodedOutput `json:"decoded,omitempty"`
	// JSON with parsed updated account state. Attention! When used in
	// `TvmOnly` mode only data in account state is updated.
	Account interface{} `json:"account,omitempty"`
}

type ParamsOfExecuteGet struct {
	Account          string            `json:"account"`
	FunctionName     string            `json:"function_name"`
	Input            interface{}       `json:"input,omitempty"`
	ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
}

type ResultOfExecuteGet struct {
	Output interface{} `json:"output"`
}
