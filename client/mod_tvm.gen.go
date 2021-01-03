package client

// DON'T EDIT THIS FILE is generated 03 Jan 21 17:31 UTC
//
// Mod tvm
//

import (
	"encoding/json"
	"math/big"

	"github.com/volatiletech/null"
)

type ExecutionOptions struct {
	// boc with config.
	BlockchainConfig null.String `json:"blockchain_config"` // optional
	// time that is used as transaction time.
	BlockTime null.Uint32 `json:"block_time"` // optional
	// block logical time.
	BlockLt *big.Int `json:"block_lt"` // optional
	// transaction logical time.
	TransactionLt *big.Int `json:"transaction_lt"` // optional
}

type AccountForExecutorType string

const (

	// Non-existing account to run a creation internal message. Should be used with `skip_transaction_check = true` if the message has no deploy data since transactions on the uninitialized account are always aborted.
	NoneAccountForExecutorType AccountForExecutorType = "None"
	// Emulate uninitialized account to run deploy message.
	UninitAccountForExecutorType AccountForExecutorType = "Uninit"
	// Account state to run message.
	AccountAccountForExecutorType AccountForExecutorType = "Account"
)

type AccountForExecutor struct {
	Type AccountForExecutorType `json:"type"`
	// Account BOC.
	// Encoded as base64. presented in types:
	// "Account".
	Boc string `json:"boc"`
	// Flag for running account with the unlimited balance.
	// Can be used to calculatetransaction fees without balance check presented in types:
	// "Account".
	UnlimitedBalance null.Bool `json:"unlimited_balance"` // optional
}

type TransactionFees struct {
	InMsgFwdFee      big.Int `json:"in_msg_fwd_fee"`
	StorageFee       big.Int `json:"storage_fee"`
	GasFee           big.Int `json:"gas_fee"`
	OutMsgsFwdFee    big.Int `json:"out_msgs_fwd_fee"`
	TotalAccountFees big.Int `json:"total_account_fees"`
	TotalOutput      big.Int `json:"total_output"`
}

type ParamsOfRunExecutor struct {
	// Input message BOC.
	// Must be encoded as base64.
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
	// In addition to the regular transaction fields there is a
	// `boc` field encoded with `base64` which contains source
	// transaction BOC.
	Transaction json.RawMessage `json:"transaction"`
	// List of output messages' BOCs.
	// Encoded as `base64`.
	OutMessages []string `json:"out_messages"`
	// Optional decoded message bodies according to the optional `abi` parameter.
	Decoded *DecodedOutput `json:"decoded"` // optional
	// Updated account state BOC.
	// Encoded as `base64`.
	Account string `json:"account"`
	// Transaction fees.
	Fees TransactionFees `json:"fees"`
}

type ParamsOfRunTvm struct {
	// Input message BOC.
	// Must be encoded as base64.
	Message string `json:"message"`
	// Account BOC.
	// Must be encoded as base64.
	Account string `json:"account"`
	// Execution options.
	ExecutionOptions *ExecutionOptions `json:"execution_options"` // optional
	// Contract ABI for dedcoding output messages.
	Abi *Abi `json:"abi"` // optional
}

type ResultOfRunTvm struct {
	// List of output messages' BOCs.
	// Encoded as `base64`.
	OutMessages []string `json:"out_messages"`
	// Optional decoded message bodies according to the optional `abi` parameter.
	Decoded *DecodedOutput `json:"decoded"` // optional
	// Updated account state BOC.
	// Encoded as `base64`.Attention! Only `account_state.storage.state.data` part of the boc is updated.
	Account string `json:"account"`
}

type ParamsOfRunGet struct {
	// Account BOC in `base64`.
	Account string `json:"account"`
	// Function name.
	FunctionName string `json:"function_name"`
	// Input parameters.
	Input            json.RawMessage   `json:"input"`             // optional
	ExecutionOptions *ExecutionOptions `json:"execution_options"` // optional
}

type ResultOfRunGet struct {
	// Values returned by getmethod on stack.
	Output json.RawMessage `json:"output"`
}

// Emulates all the phases of contract execution locally.
// Performs all the phases of contract execution on Transaction Executor -
// the same component that is used on Validator Nodes.
//
// Can be used for contract debug, to find out the reason of message unsuccessful
// delivery - as Validators just throw away failed transactions, here you can catch it.
//
// Another use case is to estimate fees for message execution. Set  `AccountForExecutor::Account.unlimited_balance`
// to `true` so that emulation will not depend on the actual balance.
//
// One more use case - you can procude the sequence of operations,
// thus emulating the multiple contract calls locally.
// And so on.
//
// To get the account boc (bag of cells) - use `net.query` method to download it from graphql api
// (field `boc` of `account`) or generate it with `abi.encode_account method`.
// To get the message boc - use `abi.encode_message` or prepare it any other way, for instance, with Fift script.
//
// If you need this emulation to be as precise as possible then specify `ParamsOfRunExecutor` parameter.
// If you need to see the aborted transaction as a result, not as an error, set `skip_transaction_check` to `true`.
func (c *Client) TvmRunExecutor(p *ParamsOfRunExecutor) (*ResultOfRunExecutor, error) {
	result := new(ResultOfRunExecutor)

	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.run_executor", p, result)

	return result, err
}

// Executes get methods of ABI-compatible contracts.
// Performs only a part of compute phase of transaction execution
// that is used to run get-methods of ABI-compatible contracts.
//
// If you try to run get methods with `run_executor` you will get an error, because it checks ACCEPT and exits
// if there is none, which is actually true for get methods.
//
// To get the account boc (bag of cells) - use `net.query` method to download it from graphql api
// (field `boc` of `account`) or generate it with `abi.encode_account method`.
// To get the message boc - use `abi.encode_message` or prepare it any other way, for instance, with Fift script.
//
// Attention! Updated account state is produces as well, but only
// `account_state.storage.state.data`  part of the boc is updated.
func (c *Client) TvmRunTvm(p *ParamsOfRunTvm) (*ResultOfRunTvm, error) {
	result := new(ResultOfRunTvm)

	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.run_tvm", p, result)

	return result, err
}

// Executes a getmethod of FIFT contract that fulfills the smc-guidelines https://test.ton.org/smc-guidelines.txt
// and returns the result data from TVM's stack.
func (c *Client) TvmRunGet(p *ParamsOfRunGet) (*ResultOfRunGet, error) {
	result := new(ResultOfRunGet)

	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.run_get", p, result)

	return result, err
}
