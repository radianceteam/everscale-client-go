package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:03:07.683423 +0000 UTC
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
	BlockchainConfig *string          `json:"blockchain_config,omitempty"`
	BlockTime        *int             `json:"block_time,omitempty"`
	BlockLt          *decimal.Decimal `json:"block_lt,omitempty"`
	TransactionLt    *decimal.Decimal `json:"transaction_lt,omitempty"`
}

type ParamsOfExecuteMessage struct {
	Message          MessageSource     `json:"message"`
	Account          string            `json:"account"`
	Mode             ExecutionMode     `json:"mode"`
	ExecutionOptions *ExecutionOptions `json:"execution_options,omitempty"`
}

type ResultOfExecuteMessage struct {
	Transaction interface{}    `json:"transaction,omitempty"`
	OutMessages []interface{}  `json:"out_messages"`
	Decoded     *DecodedOutput `json:"decoded,omitempty"`
	Account     interface{}    `json:"account,omitempty"`
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
