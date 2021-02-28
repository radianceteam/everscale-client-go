package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 28 Feb 21 17:43 UTC
//
// Mod boc
//
// BOC manipulation module.

import (
	"encoding/json"
	"fmt"

	"github.com/volatiletech/null"
)

type PinnedBocCacheType struct {
	Pin string `json:"pin"`
}

type UnpinnedBocCacheType struct{}

type BocCacheType struct {
	EnumTypeValue interface{} // any of PinnedBocCacheType, UnpinnedBocCacheType,
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
	// Key block BOC encoded as base64.
	BlockBoc string `json:"block_boc"`
}

type ResultOfGetBlockchainConfig struct {
	// Blockchain config BOC encoded as base64.
	ConfigBoc string `json:"config_boc"`
}

type ParamsOfGetBocHash struct {
	// BOC encoded as base64.
	Boc string `json:"boc"`
}

type ResultOfGetBocHash struct {
	// BOC root hash encoded with hex.
	Hash string `json:"hash"`
}

type ParamsOfGetCodeFromTvc struct {
	// Contract TVC image encoded as base64.
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
