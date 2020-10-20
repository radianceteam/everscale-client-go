package client

// DON'T EDIT THIS FILE is generated 2020-10-20 09:37:55.817726 +0000 UTC
//
// Mod tvm
//

import (
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

type ExecutionMode string

const (
	Full    ExecutionMode = "Full"
	TvmOnly ExecutionMode = "TvmOnly"
)

type ExecutionOptions struct {
	// boc with config.
	BlockchainConfig null.String `json:"blockchain_config"` // optional
	// time that is used as transaction time.
	BlockTime null.Int `json:"block_time"` // optional
	// block logical time.
	BlockLt decimal.NullDecimal `json:"block_lt"` // optional
	// transaction logical time.
	TransactionLt decimal.NullDecimal `json:"transaction_lt"` // optional
}

type ParamsOfExecuteMessage struct {
	// Input message.
	Message MessageSource `json:"message"`
	// Account BOC. Must be encoded as base64.
	Account string `json:"account"`
	// Execution mode.
	Mode ExecutionMode `json:"mode"`
	// Execution options.
	ExecutionOptions *ExecutionOptions `json:"execution_options"` // optional
}

type ResultOfExecuteMessage struct {
	// Parsed transaction.
	//
	// In addition to the regular transaction fields there is a
	// `boc` field encoded with `base64` which contains source
	// transaction BOC.
	Transaction interface{} `json:"transaction"` // optional
	// List of parsed output messages.
	//
	// Similar to the `transaction` each message contains the `boc`
	// field.
	OutMessages []interface{} `json:"out_messages"`
	// Optional decoded message bodies according to the optional
	// `abi` parameter.
	Decoded *DecodedOutput `json:"decoded"` // optional
	// JSON with parsed updated account state. Attention! When used in
	// `TvmOnly` mode only data in account state is updated.
	Account interface{} `json:"account"` // optional
}

type ParamsOfExecuteGet struct {
	Account          string            `json:"account"`
	FunctionName     string            `json:"function_name"`
	Input            interface{}       `json:"input"`             // optional
	ExecutionOptions *ExecutionOptions `json:"execution_options"` // optional
}

type ResultOfExecuteGet struct {
	Output interface{} `json:"output"`
}

func (c *Client) TvmExecuteMessage(p *ParamsOfExecuteMessage) (*ResultOfExecuteMessage, error) {
	response := new(ResultOfExecuteMessage)
	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.execute_message", p, response)

	return response, err
}

func (c *Client) TvmExecuteGet(p *ParamsOfExecuteGet) (*ResultOfExecuteGet, error) {
	response := new(ResultOfExecuteGet)
	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.execute_get", p, response)

	return response, err
}
