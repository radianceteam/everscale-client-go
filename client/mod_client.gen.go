package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 27 Feb 21 21:40 UTC
//
// Mod client
//
// Provides information about library.

import (
	"encoding/json"

	"github.com/volatiletech/null"
)

const (
	NotImplementedErrorCode                      = 1
	InvalidHexErrorCode                          = 2
	InvalidBase64ErrorCode                       = 3
	InvalidAddressErrorCode                      = 4
	CallbackParamsCantBeConvertedToJSONErrorCode = 5
	WebsocketConnectErrorErrorCode               = 6
	WebsocketReceiveErrorErrorCode               = 7
	WebsocketSendErrorErrorCode                  = 8
	HTTPClientCreateErrorErrorCode               = 9
	HTTPRequestCreateErrorErrorCode              = 10
	HTTPRequestSendErrorErrorCode                = 11
	HTTPRequestParseErrorErrorCode               = 12
	CallbackNotRegisteredErrorCode               = 13
	NetModuleNotInitErrorCode                    = 14
	InvalidConfigErrorCode                       = 15
	CannotCreateRuntimeErrorCode                 = 16
	InvalidContextHandleErrorCode                = 17
	CannotSerializeResultErrorCode               = 18
	CannotSerializeErrorErrorCode                = 19
	CannotConvertJsValueToJSONErrorCode          = 20
	CannotReceiveSpawnedResultErrorCode          = 21
	SetTimerErrorErrorCode                       = 22
	InvalidParamsErrorCode                       = 23
	ContractsAddressConversionFailedErrorCode    = 24
	UnknownFunctionErrorCode                     = 25
	AppRequestErrorErrorCode                     = 26
	NoSuchRequestErrorCode                       = 27
	CanNotSendRequestResultErrorCode             = 28
	CanNotReceiveRequestResultErrorCode          = 29
	CanNotParseRequestResultErrorCode            = 30
	UnexpectedCallbackResponseErrorCode          = 31
	CanNotParseNumberErrorCode                   = 32
	InternalErrorErrorCode                       = 33
)

func init() { // nolint gochecknoinits
	errorCodesToErrorTypes[NotImplementedErrorCode] = "NotImplementedErrorCode"
	errorCodesToErrorTypes[InvalidHexErrorCode] = "InvalidHexErrorCode"
	errorCodesToErrorTypes[InvalidBase64ErrorCode] = "InvalidBase64ErrorCode"
	errorCodesToErrorTypes[InvalidAddressErrorCode] = "InvalidAddressErrorCode"
	errorCodesToErrorTypes[CallbackParamsCantBeConvertedToJSONErrorCode] = "CallbackParamsCantBeConvertedToJSONErrorCode"
	errorCodesToErrorTypes[WebsocketConnectErrorErrorCode] = "WebsocketConnectErrorErrorCode"
	errorCodesToErrorTypes[WebsocketReceiveErrorErrorCode] = "WebsocketReceiveErrorErrorCode"
	errorCodesToErrorTypes[WebsocketSendErrorErrorCode] = "WebsocketSendErrorErrorCode"
	errorCodesToErrorTypes[HTTPClientCreateErrorErrorCode] = "HTTPClientCreateErrorErrorCode"
	errorCodesToErrorTypes[HTTPRequestCreateErrorErrorCode] = "HTTPRequestCreateErrorErrorCode"
	errorCodesToErrorTypes[HTTPRequestSendErrorErrorCode] = "HTTPRequestSendErrorErrorCode"
	errorCodesToErrorTypes[HTTPRequestParseErrorErrorCode] = "HTTPRequestParseErrorErrorCode"
	errorCodesToErrorTypes[CallbackNotRegisteredErrorCode] = "CallbackNotRegisteredErrorCode"
	errorCodesToErrorTypes[NetModuleNotInitErrorCode] = "NetModuleNotInitErrorCode"
	errorCodesToErrorTypes[InvalidConfigErrorCode] = "InvalidConfigErrorCode"
	errorCodesToErrorTypes[CannotCreateRuntimeErrorCode] = "CannotCreateRuntimeErrorCode"
	errorCodesToErrorTypes[InvalidContextHandleErrorCode] = "InvalidContextHandleErrorCode"
	errorCodesToErrorTypes[CannotSerializeResultErrorCode] = "CannotSerializeResultErrorCode"
	errorCodesToErrorTypes[CannotSerializeErrorErrorCode] = "CannotSerializeErrorErrorCode"
	errorCodesToErrorTypes[CannotConvertJsValueToJSONErrorCode] = "CannotConvertJsValueToJSONErrorCode"
	errorCodesToErrorTypes[CannotReceiveSpawnedResultErrorCode] = "CannotReceiveSpawnedResultErrorCode"
	errorCodesToErrorTypes[SetTimerErrorErrorCode] = "SetTimerErrorErrorCode"
	errorCodesToErrorTypes[InvalidParamsErrorCode] = "InvalidParamsErrorCode"
	errorCodesToErrorTypes[ContractsAddressConversionFailedErrorCode] = "ContractsAddressConversionFailedErrorCode"
	errorCodesToErrorTypes[UnknownFunctionErrorCode] = "UnknownFunctionErrorCode"
	errorCodesToErrorTypes[AppRequestErrorErrorCode] = "AppRequestErrorErrorCode"
	errorCodesToErrorTypes[NoSuchRequestErrorCode] = "NoSuchRequestErrorCode"
	errorCodesToErrorTypes[CanNotSendRequestResultErrorCode] = "CanNotSendRequestResultErrorCode"
	errorCodesToErrorTypes[CanNotReceiveRequestResultErrorCode] = "CanNotReceiveRequestResultErrorCode"
	errorCodesToErrorTypes[CanNotParseRequestResultErrorCode] = "CanNotParseRequestResultErrorCode"
	errorCodesToErrorTypes[UnexpectedCallbackResponseErrorCode] = "UnexpectedCallbackResponseErrorCode"
	errorCodesToErrorTypes[CanNotParseNumberErrorCode] = "CanNotParseNumberErrorCode"
	errorCodesToErrorTypes[InternalErrorErrorCode] = "InternalErrorErrorCode"
}

type Error struct {
	Code    uint32          `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type Config struct {
	Network *NetworkConfig `json:"network"` // optional
	Crypto  *CryptoConfig  `json:"crypto"`  // optional
	Abi     *AbiConfig     `json:"abi"`     // optional
	Boc     *BocConfig     `json:"boc"`     // optional
}

type NetworkConfig struct {
	// DApp Server public address. For instance, for `net.ton.dev/graphql` GraphQL endpoint the server address will be net.ton.dev.
	ServerAddress null.String `json:"server_address"` // optional
	// List of DApp Server addresses.
	// Any correct URL format can be specified, including IP addresses This parameter is prevailing over `server_address`.
	Endpoints []string `json:"endpoints"` // optional
	// Deprecated.
	// You must use `network.max_reconnect_timeout` that allows to specify maximum network resolving timeout.
	NetworkRetriesCount null.Int8 `json:"network_retries_count"` // optional
	// Maximum time for sequential reconnections in ms.
	// Default value is 120000 (2 min).
	MaxReconnectTimeout null.Uint32 `json:"max_reconnect_timeout"` // optional
	// Deprecated.
	ReconnectTimeout null.Uint32 `json:"reconnect_timeout"` // optional
	// The number of automatic message processing retries that SDK performs in case of `Message Expired (507)` error - but only for those messages which local emulation was successful or failed with replay protection error. The default value is 5.
	MessageRetriesCount null.Int8 `json:"message_retries_count"` // optional
	// Timeout that is used to process message delivery for the contracts which ABI does not include "expire" header. If the message is not delivered within the specified timeout the appropriate error occurs.
	MessageProcessingTimeout null.Uint32 `json:"message_processing_timeout"` // optional
	// Maximum timeout that is used for query response. The default value is 40 sec.
	WaitForTimeout null.Uint32 `json:"wait_for_timeout"` // optional
	// Maximum time difference between server and client.
	// If client's device time is out of sync and difference is more than the threshold then error will occur. Also an error will occur if the specified threshold is more than
	// `message_processing_timeout/2`.
	// The default value is 15 sec.
	OutOfSyncThreshold null.Uint32 `json:"out_of_sync_threshold"` // optional
	// Access key to GraphQL API.
	// At the moment is not used in production.
	AccessKey null.String `json:"access_key"` // optional
}

type CryptoConfig struct {
	// Mnemonic dictionary that will be used by default in crypto functions. If not specified, 1 dictionary will be used.
	MnemonicDictionary null.Uint8 `json:"mnemonic_dictionary"` // optional
	// Mnemonic word count that will be used by default in crypto functions. If not specified the default value will be 12.
	MnemonicWordCount null.Uint8 `json:"mnemonic_word_count"` // optional
	// Derivation path that will be used by default in crypto functions. If not specified `m/44'/396'/0'/0/0` will be used.
	HdkeyDerivationPath null.String `json:"hdkey_derivation_path"` // optional
}

type AbiConfig struct {
	// Workchain id that is used by default in DeploySet.
	Workchain null.Int32 `json:"workchain"` // optional
	// Message lifetime for contracts which ABI includes "expire" header. The default value is 40 sec.
	MessageExpirationTimeout null.Uint32 `json:"message_expiration_timeout"` // optional
	// Factor that increases the expiration timeout for each retry The default value is 1.5.
	MessageExpirationTimeoutGrowFactor null.Float32 `json:"message_expiration_timeout_grow_factor"` // optional
}

type BocConfig struct {
	// Maximum BOC cache size in kilobytes.
	// Default is 10 MB.
	CacheMaxSize null.Uint32 `json:"cache_max_size"` // optional
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

	// Error occurred during request processing.
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
