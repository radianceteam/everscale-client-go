package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 10 Mar 21 13:54 UTC
//
// Mod tvm
//

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/volatiletech/null"
)

const (
	CanNotReadTransactionTvmErrorCode      = 401
	CanNotReadBlockchainConfigTvmErrorCode = 402
	TransactionAbortedTvmErrorCode         = 403
	InternalErrorTvmErrorCode              = 404
	ActionPhaseFailedTvmErrorCode          = 405
	AccountCodeMissingTvmErrorCode         = 406
	LowBalanceTvmErrorCode                 = 407
	AccountFrozenOrDeletedTvmErrorCode     = 408
	AccountMissingTvmErrorCode             = 409
	UnknownExecutionErrorTvmErrorCode      = 410
	InvalidInputStackTvmErrorCode          = 411
	InvalidAccountBocTvmErrorCode          = 412
	InvalidMessageTypeTvmErrorCode         = 413
	ContractExecutionErrorTvmErrorCode     = 414
)

func init() { // nolint gochecknoinits
	errorCodesToErrorTypes[CanNotReadTransactionTvmErrorCode] = "CanNotReadTransactionTvmErrorCode"
	errorCodesToErrorTypes[CanNotReadBlockchainConfigTvmErrorCode] = "CanNotReadBlockchainConfigTvmErrorCode"
	errorCodesToErrorTypes[TransactionAbortedTvmErrorCode] = "TransactionAbortedTvmErrorCode"
	errorCodesToErrorTypes[InternalErrorTvmErrorCode] = "InternalErrorTvmErrorCode"
	errorCodesToErrorTypes[ActionPhaseFailedTvmErrorCode] = "ActionPhaseFailedTvmErrorCode"
	errorCodesToErrorTypes[AccountCodeMissingTvmErrorCode] = "AccountCodeMissingTvmErrorCode"
	errorCodesToErrorTypes[LowBalanceTvmErrorCode] = "LowBalanceTvmErrorCode"
	errorCodesToErrorTypes[AccountFrozenOrDeletedTvmErrorCode] = "AccountFrozenOrDeletedTvmErrorCode"
	errorCodesToErrorTypes[AccountMissingTvmErrorCode] = "AccountMissingTvmErrorCode"
	errorCodesToErrorTypes[UnknownExecutionErrorTvmErrorCode] = "UnknownExecutionErrorTvmErrorCode"
	errorCodesToErrorTypes[InvalidInputStackTvmErrorCode] = "InvalidInputStackTvmErrorCode"
	errorCodesToErrorTypes[InvalidAccountBocTvmErrorCode] = "InvalidAccountBocTvmErrorCode"
	errorCodesToErrorTypes[InvalidMessageTypeTvmErrorCode] = "InvalidMessageTypeTvmErrorCode"
	errorCodesToErrorTypes[ContractExecutionErrorTvmErrorCode] = "ContractExecutionErrorTvmErrorCode"
}

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

// Non-existing account to run a creation internal message. Should be used with `skip_transaction_check = true` if the message has no deploy data since transactions on the uninitialized account are always aborted.
type NoneAccountForExecutor struct{}

// Emulate uninitialized account to run deploy message.
type UninitAccountForExecutor struct{}

// Account state to run message.
type AccountAccountForExecutor struct {
	// Account BOC.
	// Encoded as base64.
	Boc string `json:"boc"`
	// Flag for running account with the unlimited balance.
	// Can be used to calculate transaction fees without balance check.
	UnlimitedBalance null.Bool `json:"unlimited_balance"` // optional
}

type AccountForExecutor struct {
	// Should be any of
	// NoneAccountForExecutor
	// UninitAccountForExecutor
	// AccountAccountForExecutor
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *AccountForExecutor) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case NoneAccountForExecutor:
		return json.Marshal(struct {
			NoneAccountForExecutor
			Type string `json:"type"`
		}{
			value,
			"None",
		})

	case UninitAccountForExecutor:
		return json.Marshal(struct {
			UninitAccountForExecutor
			Type string `json:"type"`
		}{
			value,
			"Uninit",
		})

	case AccountAccountForExecutor:
		return json.Marshal(struct {
			AccountAccountForExecutor
			Type string `json:"type"`
		}{
			value,
			"Account",
		})

	default:
		return nil, fmt.Errorf("unsupported type for AccountForExecutor %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *AccountForExecutor) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "None":
		var enumTypeValue NoneAccountForExecutor
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Uninit":
		var enumTypeValue UninitAccountForExecutor
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Account":
		var enumTypeValue AccountAccountForExecutor
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for AccountForExecutor %v", typeDescriptor.Type)
	}

	return nil
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
	// Cache type to put the result.
	// The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
	// Return updated account flag.
	// Empty string is returned if the flag is `false`.
	ReturnUpdatedAccount null.Bool `json:"return_updated_account"` // optional
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
	// Contract ABI for decoding output messages.
	Abi *Abi `json:"abi"` // optional
	// Cache type to put the result.
	// The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
	// Return updated account flag.
	// Empty string is returned if the flag is `false`.
	ReturnUpdatedAccount null.Bool `json:"return_updated_account"` // optional
}

type ResultOfRunTvm struct {
	// List of output messages' BOCs.
	// Encoded as `base64`.
	OutMessages []string `json:"out_messages"`
	// Optional decoded message bodies according to the optional `abi` parameter.
	Decoded *DecodedOutput `json:"decoded"` // optional
	// Updated account state BOC.
	// Encoded as `base64`. Attention! Only `account_state.storage.state.data` part of the BOC is updated.
	Account string `json:"account"`
}

type ParamsOfRunGet struct {
	// Account BOC in `base64`.
	Account string `json:"account"`
	// Function name.
	FunctionName string `json:"function_name"`
	// Input parameters.
	Input json.RawMessage `json:"input"` // optional
	// Execution options.
	ExecutionOptions *ExecutionOptions `json:"execution_options"` // optional
	// Convert lists based on nested tuples in the **result** into plain arrays.
	// Default is `false`. Input parameters may use any of lists representations
	// If you receive this error on Web: "Runtime error. Unreachable code should not be executed...",
	// set this flag to true.
	// This may happen, for example, when elector contract contains too many participants.
	TupleListAsArray null.Bool `json:"tuple_list_as_array"` // optional
}

type ResultOfRunGet struct {
	// Values returned by get-method on stack.
	Output json.RawMessage `json:"output"`
}

// Emulates all the phases of contract execution locally.
// Performs all the phases of contract execution on Transaction Executor -
// the same component that is used on Validator Nodes.
//
// Can be used for contract debugginh, to find out the reason why message was not delivered successfully
// - because Validators just throw away the failed external inbound messages, here you can catch them.
//
// Another use case is to estimate fees for message execution. Set  `AccountForExecutor::Account.unlimited_balance`
// to `true` so that emulation will not depend on the actual balance.
//
// One more use case - you can produce the sequence of operations,
// thus emulating the multiple contract calls locally.
// And so on.
//
// To get the account BOC (bag of cells) - use `net.query` method to download it from GraphQL API
// (field `boc` of `account`) or generate it with `abi.encode_account` method.
// To get the message BOC - use `abi.encode_message` or prepare it any other way, for instance, with FIFT script.
//
// If you need this emulation to be as precise as possible then specify `ParamsOfRunExecutor` parameter.
// If you need to see the aborted transaction as a result, not as an error, set `skip_transaction_check` to `true`.
func (c *Client) TvmRunExecutor(p *ParamsOfRunExecutor) (*ResultOfRunExecutor, error) {
	result := new(ResultOfRunExecutor)

	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.run_executor", p, result)

	return result, err
}

// Executes get-methods of ABI-compatible contracts.
// Performs only a part of compute phase of transaction execution
// that is used to run get-methods of ABI-compatible contracts.
//
// If you try to run get-methods with `run_executor` you will get an error, because it checks ACCEPT and exits
// if there is none, which is actually true for get-methods.
//
// To get the account BOC (bag of cells) - use `net.query` method to download it from GraphQL API
// (field `boc` of `account`) or generate it with `abi.encode_account method`.
// To get the message BOC - use `abi.encode_message` or prepare it any other way, for instance, with FIFT script.
//
// Attention! Updated account state is produces as well, but only
// `account_state.storage.state.data`  part of the BOC is updated.
func (c *Client) TvmRunTvm(p *ParamsOfRunTvm) (*ResultOfRunTvm, error) {
	result := new(ResultOfRunTvm)

	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.run_tvm", p, result)

	return result, err
}

// Executes a get-method of FIFT contract that fulfills the smc-guidelines https://test.ton.org/smc-guidelines.txt
// and returns the result data from TVM's stack.
func (c *Client) TvmRunGet(p *ParamsOfRunGet) (*ResultOfRunGet, error) {
	result := new(ResultOfRunGet)

	err := c.dllClient.waitErrorOrResultUnmarshal("tvm.run_get", p, result)

	return result, err
}
