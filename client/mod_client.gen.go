package client

// DON'T EDIT THIS FILE is generated 24 Oct 20 12:36 UTC
//
// Mod client
//
// Provides information about library.

import (
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

type NetworkConfig struct {
	ServerAddress            null.String         `json:"server_address"`             // optional
	NetworkRetriesCount      null.Int            `json:"network_retries_count"`      // optional
	MessageRetriesCount      null.Int            `json:"message_retries_count"`      // optional
	MessageProcessingTimeout null.Int            `json:"message_processing_timeout"` // optional
	WaitForTimeout           null.Int            `json:"wait_for_timeout"`           // optional
	OutOfSyncThreshold       decimal.NullDecimal `json:"out_of_sync_threshold"`      // optional
	AccessKey                null.String         `json:"access_key"`                 // optional
}

type CryptoConfig struct {
	MnemonicDictionary  null.Int    `json:"mnemonic_dictionary"`   // optional
	MnemonicWordCount   null.Int    `json:"mnemonic_word_count"`   // optional
	HdkeyDerivationPath null.String `json:"hdkey_derivation_path"` // optional
	HdkeyCompliant      null.Bool   `json:"hdkey_compliant"`       // optional
}

type AbiConfig struct {
	Workchain                          null.Int `json:"workchain"`                              // optional
	MessageExpirationTimeout           null.Int `json:"message_expiration_timeout"`             // optional
	MessageExpirationTimeoutGrowFactor null.Int `json:"message_expiration_timeout_grow_factor"` // optional
}

type ResultOfVersion struct {
	// Core Library version.
	Version string `json:"version"`
}

type ResultOfBuildInfo struct {
	BuildInfo interface{} `json:"build_info"`
}

// Returns Core Library version.
func (c *Client) ClientVersion() (*ResultOfVersion, error) {
	response := new(ResultOfVersion)
	err := c.dllClient.waitErrorOrResultUnmarshal("client.version", nil, response)

	return response, err
}

func (c *Client) ClientBuildInfo() (*ResultOfBuildInfo, error) {
	response := new(ResultOfBuildInfo)
	err := c.dllClient.waitErrorOrResultUnmarshal("client.build_info", nil, response)

	return response, err
}
