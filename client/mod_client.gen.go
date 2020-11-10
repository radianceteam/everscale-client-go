package client

// DON'T EDIT THIS FILE is generated 10 Nov 20 06:44 UTC
//
// Mod client
//
// Provides information about library.

import (
	"github.com/volatiletech/null"
)

type NetworkConfig struct {
	ServerAddress            string      `json:"server_address"`
	NetworkRetriesCount      null.Int8   `json:"network_retries_count"`      // optional
	MessageRetriesCount      null.Int8   `json:"message_retries_count"`      // optional
	MessageProcessingTimeout null.Uint32 `json:"message_processing_timeout"` // optional
	WaitForTimeout           null.Uint32 `json:"wait_for_timeout"`           // optional
	OutOfSyncThreshold       null.Uint32 `json:"out_of_sync_threshold"`      // optional
	AccessKey                null.String `json:"access_key"`                 // optional
}

type CryptoConfig struct {
	MnemonicDictionary  null.Uint8  `json:"mnemonic_dictionary"`   // optional
	MnemonicWordCount   null.Uint8  `json:"mnemonic_word_count"`   // optional
	HdkeyDerivationPath null.String `json:"hdkey_derivation_path"` // optional
	HdkeyCompliant      null.Bool   `json:"hdkey_compliant"`       // optional
}

type AbiConfig struct {
	Workchain                          null.Int32   `json:"workchain"`                              // optional
	MessageExpirationTimeout           null.Uint32  `json:"message_expiration_timeout"`             // optional
	MessageExpirationTimeoutGrowFactor null.Float32 `json:"message_expiration_timeout_grow_factor"` // optional
}

type BuildInfoDependency struct {
	// Dependency name. Usually it is a crate name.
	Name string `json:"name"`
	// Git commit hash of the related repository.
	GitCommit string `json:"git_commit"`
}

type ResultOfVersion struct {
	// Core Library version.
	Version string `json:"version"`
}

type ResultOfBuildInfo struct {
	// Build number assigned to this build by the CI.
	BuildNumber uint32 `json:"build_number"`
	// Fingerprint of the most important dependencies.
	Dependencies []BuildInfoDependency `json:"dependencies"`
}

// Returns Core Library version.
func (c *Client) ClientVersion() (*ResultOfVersion, error) {
	response := new(ResultOfVersion)
	err := c.dllClient.waitErrorOrResultUnmarshal("client.version", nil, response)

	return response, err
}

// Returns detailed information about this build.
func (c *Client) ClientBuildInfo() (*ResultOfBuildInfo, error) {
	response := new(ResultOfBuildInfo)
	err := c.dllClient.waitErrorOrResultUnmarshal("client.build_info", nil, response)

	return response, err
}
