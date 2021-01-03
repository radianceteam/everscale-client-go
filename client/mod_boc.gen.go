package client

// DON'T EDIT THIS FILE is generated 03 Jan 21 17:19 UTC
//
// Mod boc
//
// BOC manipulation module.

import (
	"encoding/json"
)

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
