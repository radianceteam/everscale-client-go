package client

// DON'T EDIT THIS FILE is generated 03 Jan 21 17:19 UTC
//
// Mod client
//
// Provides information about library.

import (
	"encoding/json"

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
}

type AbiConfig struct {
	Workchain                          null.Int32   `json:"workchain"`                              // optional
	MessageExpirationTimeout           null.Uint32  `json:"message_expiration_timeout"`             // optional
	MessageExpirationTimeoutGrowFactor null.Float32 `json:"message_expiration_timeout_grow_factor"` // optional
}

type BuildInfoDependency struct {
	// Dependency name.
	// Usually it is a crate name.
	Name string `json:"name"`
	// Git commit hash of the related repository.
	GitCommit string `json:"git_commit"`
}

type ParamsOfAppRequest struct {
	// Request ID.
	// Should be used in `resolve_app_request` call.
	AppRequestID uint32 `json:"app_request_id"`
	// Request describing data.
	RequestData json.RawMessage `json:"request_data"`
}

type AppRequestResultType string

const (

	// Error occured during request processing.
	ErrorAppRequestResultType AppRequestResultType = "Error"
	// Request processed successfully.
	OkAppRequestResultType AppRequestResultType = "Ok"
)

type AppRequestResult struct {
	Type AppRequestResultType `json:"type"`
	// Error description.
	// presented in types:
	// "Error".
	Text string `json:"text"`
	// Request processing result.
	// presented in types:
	// "Ok".
	Result json.RawMessage `json:"result"`
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

type ParamsOfResolveAppRequest struct {
	// Request ID received from SDK.
	AppRequestID uint32 `json:"app_request_id"`
	// Result of request processing.
	Result AppRequestResult `json:"result"`
}

// Returns Core Library version.
func (c *Client) ClientVersion() (*ResultOfVersion, error) {
	result := new(ResultOfVersion)

	err := c.dllClient.waitErrorOrResultUnmarshal("client.version", nil, result)

	return result, err
}

// Returns detailed information about this build.
func (c *Client) ClientBuildInfo() (*ResultOfBuildInfo, error) {
	result := new(ResultOfBuildInfo)

	err := c.dllClient.waitErrorOrResultUnmarshal("client.build_info", nil, result)

	return result, err
}

// Resolves application request processing result.
func (c *Client) ClientResolveAppRequest(p *ParamsOfResolveAppRequest) error {
	_, err := c.dllClient.waitErrorOrResult("client.resolve_app_request", p)

	return err
}
