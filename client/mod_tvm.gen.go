package client

// DON'T EDIT THIS FILE is generated 03 Nov 20 15:00 UTC
//
// Mod tvm
//

import (
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
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

type AccountForExecutorType string

const (

	// Non-existing account to run a creation internal message.
	// Should be used with `skip_transaction_check = true` if the message has no deploy data
	// since transaction on the unitialized account are always aborted.
	NoneAccountForExecutorType AccountForExecutorType = "None"
	// Emulate unitialized account to run deploy message.
	UninitAccountForExecutorType AccountForExecutorType = "Uninit"
	// Account state to run message.
	AccountAccountForExecutorType AccountForExecutorType = "Account"
)

type AccountForExecutor struct {
	Type AccountForExecutorType `json:"type"`
	// Account BOC. Encoded as base64. presented in types:
	// "Account".
	Boc string `json:"boc"`
	// Flag for running account with the unlimited balance. Can be used to calculate
	// transaction fees without balance check presented in types:
	// "Account".
	UnlimitedBalance null.Bool `json:"unlimited_balance"` // optional
}

type ParamsOfRunExecutor struct {
	// Input message BOC. Must be encoded as base64.
	Message string `json:"message"`
	// Account to run on executor.
	Account AccountForExecutor `json:"account"`
	// Execution options.
	ExecutionOptions *ExecutionOptions `json:"execution_options"` // optional
	// Contract ABI for decoding output messages.
	Abi *Abi `json:"abi"` // optional
	// Skip transaction check flag.
	SkipTransactionCheck null.Bool `json:"skip_transaction_check"` // optional
}

type ResultOfRunExecutor struct {
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
	// Updated account state BOC. Encoded as `base64`.
	Account string `json:"account"`
	// Transaction fees.
	Fees TransactionFees `json:"fees"`
}

type ParamsOfRunTvm struct {
	// Input message BOC. Must be encoded as base64.
	Message string `json:"message"`
	// Account BOC. Must be encoded as base64.
	Account string `json:"account"`
	// Execution options.
	ExecutionOptions *ExecutionOptions `json:"execution_options"` // optional
	// Contract ABI for dedcoding output messages.
	Abi *Abi `json:"abi"` // optional
}

type ResultOfRunTvm struct {
	// List of output messages' BOCs. Encoded as `base64`.
	OutMessages []string `json:"out_messages"`
	// Optional decoded message bodies according to the optional
	// `abi` parameter.
	Decoded *DecodedOutput `json:"decoded"` // optional
	// Updated account state BOC. Encoded as `base64`.
	// Attention! Only data in account state is updated.
	Account string `json:"account"`
}

type ParamsOfRunGet struct {
	// Account BOC in `base64`.
	Account string `json:"account"`
	// Function name.
	FunctionName string `json:"function_name"`
	// Input parameters.
	Input            interface{}       `json:"input"`             // optional
	ExecutionOptions *ExecutionOptions `json:"execution_options"` // optional
}

type ResultOfRunGet struct {
	// Values returned by getmethod on stack.
	Output interface{} `json:"output"`
}

func (c *Client) TvmRunExecutor(p *ParamsOfRunExecutor) (*ResultOfRunExecutor, error) {
	response := new(ResultOfRunExecutor)
	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.run_executor", p, response)

	return response, err
}

func (c *Client) TvmRunTvm(p *ParamsOfRunTvm) (*ResultOfRunTvm, error) {
	response := new(ResultOfRunTvm)
	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.run_tvm", p, response)

	return response, err
}

// Executes getmethod and returns data from TVM stack.
func (c *Client) TvmRunGet(p *ParamsOfRunGet) (*ResultOfRunGet, error) {
	response := new(ResultOfRunGet)
	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.run_get", p, response)

	return response, err
}
