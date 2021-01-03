package client

// DON'T EDIT THIS FILE is generated 03 Jan 21 17:49 UTC
//
// Mod client
//
// Provides information about library.

import (
	"encoding/json"

	"github.com/volatiletech/null"
)

type ErrorCode string

const (
	NotImplementedErrorCode                      ErrorCode = "NotImplemented"
	InvalidHexErrorCode                          ErrorCode = "InvalidHex"
	InvalidBase64ErrorCode                       ErrorCode = "InvalidBase64"
	InvalidAddressErrorCode                      ErrorCode = "InvalidAddress"
	CallbackParamsCantBeConvertedToJSONErrorCode ErrorCode = "CallbackParamsCantBeConvertedToJson"
	WebsocketConnectErrorErrorCode               ErrorCode = "WebsocketConnectError"
	WebsocketReceiveErrorErrorCode               ErrorCode = "WebsocketReceiveError"
	WebsocketSendErrorErrorCode                  ErrorCode = "WebsocketSendError"
	HTTPClientCreateErrorErrorCode               ErrorCode = "HttpClientCreateError"
	HTTPRequestCreateErrorErrorCode              ErrorCode = "HttpRequestCreateError"
	HTTPRequestSendErrorErrorCode                ErrorCode = "HttpRequestSendError"
	HTTPRequestParseErrorErrorCode               ErrorCode = "HttpRequestParseError"
	CallbackNotRegisteredErrorCode               ErrorCode = "CallbackNotRegistered"
	NetModuleNotInitErrorCode                    ErrorCode = "NetModuleNotInit"
	InvalidConfigErrorCode                       ErrorCode = "InvalidConfig"
	CannotCreateRuntimeErrorCode                 ErrorCode = "CannotCreateRuntime"
	InvalidContextHandleErrorCode                ErrorCode = "InvalidContextHandle"
	CannotSerializeResultErrorCode               ErrorCode = "CannotSerializeResult"
	CannotSerializeErrorErrorCode                ErrorCode = "CannotSerializeError"
	CannotConvertJsValueToJSONErrorCode          ErrorCode = "CannotConvertJsValueToJson"
	CannotReceiveSpawnedResultErrorCode          ErrorCode = "CannotReceiveSpawnedResult"
	SetTimerErrorErrorCode                       ErrorCode = "SetTimerError"
	InvalidParamsErrorCode                       ErrorCode = "InvalidParams"
	ContractsAddressConversionFailedErrorCode    ErrorCode = "ContractsAddressConversionFailed"
	UnknownFunctionErrorCode                     ErrorCode = "UnknownFunction"
	AppRequestErrorErrorCode                     ErrorCode = "AppRequestError"
	NoSuchRequestErrorCode                       ErrorCode = "NoSuchRequest"
	CanNotSendRequestResultErrorCode             ErrorCode = "CanNotSendRequestResult"
	CanNotReceiveRequestResultErrorCode          ErrorCode = "CanNotReceiveRequestResult"
	CanNotParseRequestResultErrorCode            ErrorCode = "CanNotParseRequestResult"
	UnexpectedCallbackResponseErrorCode          ErrorCode = "UnexpectedCallbackResponse"
	CanNotParseNumberErrorCode                   ErrorCode = "CanNotParseNumber"
	InternalErrorErrorCode                       ErrorCode = "InternalError"
)

type NetworkConfig struct {
	// DApp Server public address. For instance, for `net.ton.dev/graphql` GraphQL endpoint the server address will be net.ton.dev.
	ServerAddress null.String `json:"server_address"` // optional
	// List of DApp Server addresses.
	// Any correct URL format can be specified, including IP addresses.
	Endpoints []string `json:"endpoints"` // optional
	// The number of automatic network retries that SDK performs in case of connection problems The default value is 5.
	NetworkRetriesCount null.Int8 `json:"network_retries_count"` // optional
	// The number of automatic message processing retries that SDK performs in case of `Message Expired (507)` error - but only for those messages which local emulation was successfull or failed with replay protection error. The default value is 5.
	MessageRetriesCount null.Int8 `json:"message_retries_count"` // optional
	// Timeout that is used to process message delivery for the contracts which ABI does not include "expire" header. If the message is not delivered within the speficied timeout the appropriate error occurs.
	MessageProcessingTimeout null.Uint32 `json:"message_processing_timeout"` // optional
	// Maximum timeout that is used for query response. The default value is 40 sec.
	WaitForTimeout null.Uint32 `json:"wait_for_timeout"` // optional
	// Maximum time difference between server and client.
	// If client's device time is out of sink and difference is more thanthe threshhold then error will occur. Also the error will occur if the specified threshhold is more than
	// `message_processing_timeout/2`.
	// The default value is 15 sec.
	OutOfSyncThreshold null.Uint32 `json:"out_of_sync_threshold"` // optional
	// Timeout between reconnect attempts.
	ReconnectTimeout null.Uint32 `json:"reconnect_timeout"` // optional
	// Access key to GraphQL API.
	// At the moment is not used in production.
	AccessKey null.String `json:"access_key"` // optional
}

type CryptoConfig struct {
	// Mnemonic dictionary that will be used by default in crypto funcions. If not specified, 1 dictionary will be used.
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
