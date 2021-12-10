package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 10 Dec 21 06:20 UTC
//
// Mod boc
//
// BOC manipulation module.

import (
	"encoding/json"
	"fmt"

	"github.com/volatiletech/null"
)

// Pin the BOC with `pin` name.
// Such BOC will not be removed from cache until it is unpinned.
type PinnedBocCacheType struct {
	Pin string `json:"pin"`
}

// .
type UnpinnedBocCacheType struct{}

type BocCacheType struct {
	// Should be any of
	// PinnedBocCacheType
	// UnpinnedBocCacheType
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *BocCacheType) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case PinnedBocCacheType:
		return json.Marshal(struct {
			PinnedBocCacheType
			Type string `json:"type"`
		}{
			value,
			"Pinned",
		})

	case UnpinnedBocCacheType:
		return json.Marshal(struct {
			UnpinnedBocCacheType
			Type string `json:"type"`
		}{
			value,
			"Unpinned",
		})

	default:
		return nil, fmt.Errorf("unsupported type for BocCacheType %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *BocCacheType) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "Pinned":
		var enumTypeValue PinnedBocCacheType
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Unpinned":
		var enumTypeValue UnpinnedBocCacheType
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for BocCacheType %v", typeDescriptor.Type)
	}

	return nil
}

const (
	InvalidBocBocErrorCode            = 201
	SerializationErrorBocErrorCode    = 202
	InappropriateBlockBocErrorCode    = 203
	MissingSourceBocBocErrorCode      = 204
	InsufficientCacheSizeBocErrorCode = 205
	BocRefNotFoundBocErrorCode        = 206
	InvalidBocRefBocErrorCode         = 207
)

func init() { // nolint gochecknoinits
	errorCodesToErrorTypes[InvalidBocBocErrorCode] = "InvalidBocBocErrorCode"
	errorCodesToErrorTypes[SerializationErrorBocErrorCode] = "SerializationErrorBocErrorCode"
	errorCodesToErrorTypes[InappropriateBlockBocErrorCode] = "InappropriateBlockBocErrorCode"
	errorCodesToErrorTypes[MissingSourceBocBocErrorCode] = "MissingSourceBocBocErrorCode"
	errorCodesToErrorTypes[InsufficientCacheSizeBocErrorCode] = "InsufficientCacheSizeBocErrorCode"
	errorCodesToErrorTypes[BocRefNotFoundBocErrorCode] = "BocRefNotFoundBocErrorCode"
	errorCodesToErrorTypes[InvalidBocRefBocErrorCode] = "InvalidBocRefBocErrorCode"
}

type ParamsOfParse struct {
	// BOC encoded as base64.
	Boc string `json:"boc"`
}

type ResultOfParse struct {
	// JSON containing parsed BOC.
	Parsed json.RawMessage `json:"parsed"`
}

type ParamsOfParseShardstate struct {
	// BOC encoded as base64.
	Boc string `json:"boc"`
	// Shardstate identificator.
	ID string `json:"id"`
	// Workchain shardstate belongs to.
	WorkchainID int32 `json:"workchain_id"`
}

type ParamsOfGetBlockchainConfig struct {
	// Key block BOC or zerostate BOC encoded as base64.
	BlockBoc string `json:"block_boc"`
}

type ResultOfGetBlockchainConfig struct {
	// Blockchain config BOC encoded as base64.
	ConfigBoc string `json:"config_boc"`
}

type ParamsOfGetBocHash struct {
	// BOC encoded as base64 or BOC handle.
	Boc string `json:"boc"`
}

type ResultOfGetBocHash struct {
	// BOC root hash encoded with hex.
	Hash string `json:"hash"`
}

type ParamsOfGetBocDepth struct {
	// BOC encoded as base64 or BOC handle.
	Boc string `json:"boc"`
}

type ResultOfGetBocDepth struct {
	// BOC root cell depth.
	Depth uint32 `json:"depth"`
}

type ParamsOfGetCodeFromTvc struct {
	// Contract TVC image or image BOC handle.
	Tvc string `json:"tvc"`
}

type ResultOfGetCodeFromTvc struct {
	// Contract code encoded as base64.
	Code string `json:"code"`
}

type ParamsOfBocCacheGet struct {
	// Reference to the cached BOC.
	BocRef string `json:"boc_ref"`
}

type ResultOfBocCacheGet struct {
	// BOC encoded as base64.
	Boc null.String `json:"boc"` // optional
}

type ParamsOfBocCacheSet struct {
	// BOC encoded as base64 or BOC reference.
	Boc string `json:"boc"`
	// Cache type.
	CacheType BocCacheType `json:"cache_type"`
}

type ResultOfBocCacheSet struct {
	// Reference to the cached BOC.
	BocRef string `json:"boc_ref"`
}

type ParamsOfBocCacheUnpin struct {
	// Pinned name.
	Pin string `json:"pin"`
	// Reference to the cached BOC.
	// If it is provided then only referenced BOC is unpinned.
	BocRef null.String `json:"boc_ref"` // optional
}

// Cell builder operation.

// Append integer to cell data.
type IntegerBuilderOp struct {
	// Bit size of the value.
	Size uint32 `json:"size"`
	// Value: - `Number` containing integer number.
	// e.g. `123`, `-123`. - Decimal string. e.g. `"123"`, `"-123"`.
	// - `0x` prefixed hexadecimal string.
	// e.g `0x123`, `0X123`, `-0x123`.
	Value json.RawMessage `json:"value"`
}

// Append bit string to cell data.
type BitStringBuilderOp struct {
	// Bit string content using bitstring notation. See `TON VM specification` 1.0.
	// Contains hexadecimal string representation:
	// - Can end with `_` tag.
	// - Can be prefixed with `x` or `X`.
	// - Can be prefixed with `x{` or `X{` and ended with `}`.
	//
	// Contains binary string represented as a sequence
	// of `0` and `1` prefixed with `n` or `N`.
	//
	// Examples:
	// `1AB`, `x1ab`, `X1AB`, `x{1abc}`, `X{1ABC}`
	// `2D9_`, `x2D9_`, `X2D9_`, `x{2D9_}`, `X{2D9_}`
	// `n00101101100`, `N00101101100`.
	Value string `json:"value"`
}

// Append ref to nested cells.
type CellBuilderOp struct {
	// Nested cell builder.
	Builder []BuilderOp `json:"builder"`
}

// Append ref to nested cell.
type CellBocBuilderOp struct {
	// Nested cell BOC encoded with `base64` or BOC cache key.
	Boc string `json:"boc"`
}

type BuilderOp struct {
	// Should be any of
	// IntegerBuilderOp
	// BitStringBuilderOp
	// CellBuilderOp
	// CellBocBuilderOp
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *BuilderOp) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case IntegerBuilderOp:
		return json.Marshal(struct {
			IntegerBuilderOp
			Type string `json:"type"`
		}{
			value,
			"Integer",
		})

	case BitStringBuilderOp:
		return json.Marshal(struct {
			BitStringBuilderOp
			Type string `json:"type"`
		}{
			value,
			"BitString",
		})

	case CellBuilderOp:
		return json.Marshal(struct {
			CellBuilderOp
			Type string `json:"type"`
		}{
			value,
			"Cell",
		})

	case CellBocBuilderOp:
		return json.Marshal(struct {
			CellBocBuilderOp
			Type string `json:"type"`
		}{
			value,
			"CellBoc",
		})

	default:
		return nil, fmt.Errorf("unsupported type for BuilderOp %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *BuilderOp) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "Integer":
		var enumTypeValue IntegerBuilderOp
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "BitString":
		var enumTypeValue BitStringBuilderOp
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Cell":
		var enumTypeValue CellBuilderOp
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "CellBoc":
		var enumTypeValue CellBocBuilderOp
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for BuilderOp %v", typeDescriptor.Type)
	}

	return nil
}

type ParamsOfEncodeBoc struct {
	// Cell builder operations.
	Builder []BuilderOp `json:"builder"`
	// Cache type to put the result. The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
}

type ResultOfEncodeBoc struct {
	// Encoded cell BOC or BOC cache key.
	Boc string `json:"boc"`
}

type ParamsOfGetCodeSalt struct {
	// Contract code BOC encoded as base64 or code BOC handle.
	Code string `json:"code"`
	// Cache type to put the result. The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
}

type ResultOfGetCodeSalt struct {
	// Contract code salt if present.
	// BOC encoded as base64 or BOC handle.
	Salt null.String `json:"salt"` // optional
}

type ParamsOfSetCodeSalt struct {
	// Contract code BOC encoded as base64 or code BOC handle.
	Code string `json:"code"`
	// Code salt to set.
	// BOC encoded as base64 or BOC handle.
	Salt string `json:"salt"`
	// Cache type to put the result. The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
}

type ResultOfSetCodeSalt struct {
	// Contract code with salt set.
	// BOC encoded as base64 or BOC handle.
	Code string `json:"code"`
}

type ParamsOfDecodeTvc struct {
	// Contract TVC image BOC encoded as base64 or BOC handle.
	Tvc string `json:"tvc"`
	// Cache type to put the result. The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
}

type ResultOfDecodeTvc struct {
	// Contract code BOC encoded as base64 or BOC handle.
	Code null.String `json:"code"` // optional
	// Contract code hash.
	CodeHash null.String `json:"code_hash"` // optional
	// Contract code depth.
	CodeDepth null.Uint32 `json:"code_depth"` // optional
	// Contract data BOC encoded as base64 or BOC handle.
	Data null.String `json:"data"` // optional
	// Contract data hash.
	DataHash null.String `json:"data_hash"` // optional
	// Contract data depth.
	DataDepth null.Uint32 `json:"data_depth"` // optional
	// Contract library BOC encoded as base64 or BOC handle.
	Library null.String `json:"library"` // optional
	// `special.tick` field.
	// Specifies the contract ability to handle tick transactions.
	Tick null.Bool `json:"tick"` // optional
	// `special.tock` field.
	// Specifies the contract ability to handle tock transactions.
	Tock null.Bool `json:"tock"` // optional
	// Is present and non-zero only in instances of large smart contracts.
	SplitDepth null.Uint32 `json:"split_depth"` // optional
	// Compiler version, for example 'sol 0.49.0'.
	CompilerVersion null.String `json:"compiler_version"` // optional
}

type ParamsOfEncodeTvc struct {
	// Contract code BOC encoded as base64 or BOC handle.
	Code null.String `json:"code"` // optional
	// Contract data BOC encoded as base64 or BOC handle.
	Data null.String `json:"data"` // optional
	// Contract library BOC encoded as base64 or BOC handle.
	Library null.String `json:"library"` // optional
	// `special.tick` field.
	// Specifies the contract ability to handle tick transactions.
	Tick null.Bool `json:"tick"` // optional
	// `special.tock` field.
	// Specifies the contract ability to handle tock transactions.
	Tock null.Bool `json:"tock"` // optional
	// Is present and non-zero only in instances of large smart contracts.
	SplitDepth null.Uint32 `json:"split_depth"` // optional
	// Cache type to put the result. The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
}

type ResultOfEncodeTvc struct {
	// Contract TVC image BOC encoded as base64 or BOC handle of boc_cache parameter was specified.
	Tvc string `json:"tvc"`
}

type ParamsOfGetCompilerVersion struct {
	// Contract code BOC encoded as base64 or code BOC handle.
	Code string `json:"code"`
}

type ResultOfGetCompilerVersion struct {
	// Compiler version, for example 'sol 0.49.0'.
	Version null.String `json:"version"` // optional
}

// Parses message boc into a JSON.
// JSON structure is compatible with GraphQL API message object.
func (c *Client) BocParseMessage(p *ParamsOfParse) (*ResultOfParse, error) {
	result := new(ResultOfParse)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_message", p, result)

	return result, err
}

// Parses transaction boc into a JSON.
// JSON structure is compatible with GraphQL API transaction object.
func (c *Client) BocParseTransaction(p *ParamsOfParse) (*ResultOfParse, error) {
	result := new(ResultOfParse)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_transaction", p, result)

	return result, err
}

// Parses account boc into a JSON.
// JSON structure is compatible with GraphQL API account object.
func (c *Client) BocParseAccount(p *ParamsOfParse) (*ResultOfParse, error) {
	result := new(ResultOfParse)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_account", p, result)

	return result, err
}

// Parses block boc into a JSON.
// JSON structure is compatible with GraphQL API block object.
func (c *Client) BocParseBlock(p *ParamsOfParse) (*ResultOfParse, error) {
	result := new(ResultOfParse)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_block", p, result)

	return result, err
}

// Parses shardstate boc into a JSON.
// JSON structure is compatible with GraphQL API shardstate object.
func (c *Client) BocParseShardstate(p *ParamsOfParseShardstate) (*ResultOfParse, error) {
	result := new(ResultOfParse)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.parse_shardstate", p, result)

	return result, err
}

// Extract blockchain configuration from key block and also from zerostate.
func (c *Client) BocGetBlockchainConfig(p *ParamsOfGetBlockchainConfig) (*ResultOfGetBlockchainConfig, error) {
	result := new(ResultOfGetBlockchainConfig)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.get_blockchain_config", p, result)

	return result, err
}

// Calculates BOC root hash.
func (c *Client) BocGetBocHash(p *ParamsOfGetBocHash) (*ResultOfGetBocHash, error) {
	result := new(ResultOfGetBocHash)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.get_boc_hash", p, result)

	return result, err
}

// Calculates BOC depth.
func (c *Client) BocGetBocDepth(p *ParamsOfGetBocDepth) (*ResultOfGetBocDepth, error) {
	result := new(ResultOfGetBocDepth)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.get_boc_depth", p, result)

	return result, err
}

// Extracts code from TVC contract image.
func (c *Client) BocGetCodeFromTvc(p *ParamsOfGetCodeFromTvc) (*ResultOfGetCodeFromTvc, error) {
	result := new(ResultOfGetCodeFromTvc)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.get_code_from_tvc", p, result)

	return result, err
}

// Get BOC from cache.
func (c *Client) BocCacheGet(p *ParamsOfBocCacheGet) (*ResultOfBocCacheGet, error) {
	result := new(ResultOfBocCacheGet)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.cache_get", p, result)

	return result, err
}

// Save BOC into cache.
func (c *Client) BocCacheSet(p *ParamsOfBocCacheSet) (*ResultOfBocCacheSet, error) {
	result := new(ResultOfBocCacheSet)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.cache_set", p, result)

	return result, err
}

// Unpin BOCs with specified pin.
// BOCs which don't have another pins will be removed from cache.
func (c *Client) BocCacheUnpin(p *ParamsOfBocCacheUnpin) error {
	_, err := c.dllClient.waitErrorOrResult("boc.cache_unpin", p)

	return err
}

// Encodes bag of cells (BOC) with builder operations. This method provides the same functionality as Solidity TvmBuilder. Resulting BOC of this method can be passed into Solidity and C++ contracts as TvmCell type.
func (c *Client) BocEncodeBoc(p *ParamsOfEncodeBoc) (*ResultOfEncodeBoc, error) {
	result := new(ResultOfEncodeBoc)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.encode_boc", p, result)

	return result, err
}

// Returns the contract code's salt if it is present.
func (c *Client) BocGetCodeSalt(p *ParamsOfGetCodeSalt) (*ResultOfGetCodeSalt, error) {
	result := new(ResultOfGetCodeSalt)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.get_code_salt", p, result)

	return result, err
}

// Sets new salt to contract code.
// Returns the new contract code with salt.
func (c *Client) BocSetCodeSalt(p *ParamsOfSetCodeSalt) (*ResultOfSetCodeSalt, error) {
	result := new(ResultOfSetCodeSalt)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.set_code_salt", p, result)

	return result, err
}

// Decodes tvc into code, data, libraries and special options.
func (c *Client) BocDecodeTvc(p *ParamsOfDecodeTvc) (*ResultOfDecodeTvc, error) {
	result := new(ResultOfDecodeTvc)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.decode_tvc", p, result)

	return result, err
}

// Encodes tvc from code, data, libraries ans special options (see input params).
func (c *Client) BocEncodeTvc(p *ParamsOfEncodeTvc) (*ResultOfEncodeTvc, error) {
	result := new(ResultOfEncodeTvc)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.encode_tvc", p, result)

	return result, err
}

// Returns the compiler version used to compile the code.
func (c *Client) BocGetCompilerVersion(p *ParamsOfGetCompilerVersion) (*ResultOfGetCompilerVersion, error) {
	result := new(ResultOfGetCompilerVersion)

	err := c.dllClient.waitErrorOrResultUnmarshal("boc.get_compiler_version", p, result)

	return result, err
}
