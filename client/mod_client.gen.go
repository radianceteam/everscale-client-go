package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 09 Jul 22 15:07 UTC
//
// Mod client
//
// Provides information about library.

import (
	"encoding/json"
	"fmt"

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
	InvalidHandleErrorCode                       = 34
	LocalStorageErrorErrorCode                   = 35
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
	errorCodesToErrorTypes[InvalidHandleErrorCode] = "InvalidHandleErrorCode"
	errorCodesToErrorTypes[LocalStorageErrorErrorCode] = "LocalStorageErrorErrorCode"
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
	Proofs  *ProofsConfig  `json:"proofs"`  // optional
	// For file based storage is a folder name where SDK will store its data. For browser based is a browser async storage key prefix. Default (recommended) value is "~/.tonclient" for native environments and ".tonclient" for web-browser.
	LocalStoragePath null.String `json:"local_storage_path"` // optional
}

type NetworkConfig struct {
	// **This field is deprecated, but left for backward-compatibility.** DApp Server public address.
	ServerAddress null.String `json:"server_address"` // optional
	// List of DApp Server addresses.
	// Any correct URL format can be specified, including IP addresses. This parameter is prevailing over `server_address`.
	// Check the full list of [supported network endpoints](../ton-os-api/networks.md).
	Endpoints []string `json:"endpoints"` // optional
	// Deprecated.
	// You must use `network.max_reconnect_timeout` that allows to specify maximum network resolving timeout.
	NetworkRetriesCount null.Int8 `json:"network_retries_count"` // optional
	// Maximum time for sequential reconnections.
	// Must be specified in milliseconds. Default is 120000 (2 min).
	MaxReconnectTimeout null.Uint32 `json:"max_reconnect_timeout"` // optional
	// Deprecated.
	ReconnectTimeout null.Uint32 `json:"reconnect_timeout"` // optional
	// The number of automatic message processing retries that SDK performs in case of `Message Expired (507)` error - but only for those messages which local emulation was successful or failed with replay protection error.
	// Default is 5.
	MessageRetriesCount null.Int8 `json:"message_retries_count"` // optional
	// Timeout that is used to process message delivery for the contracts which ABI does not include "expire" header. If the message is not delivered within the specified timeout the appropriate error occurs.
	// Must be specified in milliseconds. Default is 40000 (40 sec).
	MessageProcessingTimeout null.Uint32 `json:"message_processing_timeout"` // optional
	// Maximum timeout that is used for query response.
	// Must be specified in milliseconds. Default is 40000 (40 sec).
	WaitForTimeout null.Uint32 `json:"wait_for_timeout"` // optional
	// Maximum time difference between server and client.
	// If client's device time is out of sync and difference is more than the threshold then error will occur. Also an error will occur if the specified threshold is more than
	// `message_processing_timeout/2`.
	//
	// Must be specified in milliseconds. Default is 15000 (15 sec).
	OutOfSyncThreshold null.Uint32 `json:"out_of_sync_threshold"` // optional
	// Maximum number of randomly chosen endpoints the library uses to broadcast a message.
	// Default is 1.
	SendingEndpointCount null.Uint8 `json:"sending_endpoint_count"` // optional
	// Frequency of sync latency detection.
	// Library periodically checks the current endpoint for blockchain data synchronization latency.
	// If the latency (time-lag) is less then `NetworkConfig.max_latency`
	// then library selects another endpoint.
	//
	// Must be specified in milliseconds. Default is 60000 (1 min).
	LatencyDetectionInterval null.Uint32 `json:"latency_detection_interval"` // optional
	// Maximum value for the endpoint's blockchain data synchronization latency (time-lag). Library periodically checks the current endpoint for blockchain data synchronization latency. If the latency (time-lag) is less then `NetworkConfig.max_latency` then library selects another endpoint.
	// Must be specified in milliseconds. Default is 60000 (1 min).
	MaxLatency null.Uint32 `json:"max_latency"` // optional
	// Default timeout for http requests.
	// Is is used when no timeout specified for the request to limit the answer waiting time. If no answer received during the timeout requests ends with
	// error.
	//
	// Must be specified in milliseconds. Default is 60000 (1 min).
	QueryTimeout null.Uint32 `json:"query_timeout"` // optional
	// Queries protocol.
	// `HTTP` or `WS`.
	// Default is `HTTP`.
	QueriesProtocol *NetworkQueriesProtocol `json:"queries_protocol"` // optional
	// UNSTABLE.
	// First REMP status awaiting timeout. If no status received during the timeout than fallback transaction scenario is activated.
	//
	// Must be specified in milliseconds. Default is 1000 (1 sec).
	FirstRempStatusTimeout null.Uint32 `json:"first_remp_status_timeout"` // optional
	// UNSTABLE.
	// Subsequent REMP status awaiting timeout. If no status received during the timeout than fallback transaction scenario is activated.
	//
	// Must be specified in milliseconds. Default is 5000 (5 sec).
	NextRempStatusTimeout null.Uint32 `json:"next_remp_status_timeout"` // optional
	// Access key to GraphQL API.
	// At the moment is not used in production.
	AccessKey null.String `json:"access_key"` // optional
}

type NetworkQueriesProtocol string

const (

	// Each GraphQL query uses separate HTTP request.
	HTTPNetworkQueriesProtocol NetworkQueriesProtocol = "HTTP"
	// All GraphQL queries will be served using single web socket connection.
	WsNetworkQueriesProtocol NetworkQueriesProtocol = "WS"
)

// Crypto config.
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

type ProofsConfig struct {
	// Cache proofs in the local storage.
	// Default is `true`. If this value is set to `true`, downloaded proofs and master-chain BOCs are saved into the
	// persistent local storage (e.g. file system for native environments or browser's IndexedDB
	// for the web); otherwise all the data is cached only in memory in current client's context
	// and will be lost after destruction of the client.
	CacheInLocalStorage null.Bool `json:"cache_in_local_storage"` // optional
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

// Error occurred during request processing.
type ErrorAppRequestResult struct {
	// Error description.
	Text string `json:"text"`
}

// Request processed successfully.
type OkAppRequestResult struct {
	// Request processing result.
	Result json.RawMessage `json:"result"`
}

type AppRequestResult struct {
	// Should be any of
	// ErrorAppRequestResult
	// OkAppRequestResult
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *AppRequestResult) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case ErrorAppRequestResult:
		return json.Marshal(struct {
			ErrorAppRequestResult
			Type string `json:"type"`
		}{
			value,
			"Error",
		})

	case OkAppRequestResult:
		return json.Marshal(struct {
			OkAppRequestResult
			Type string `json:"type"`
		}{
			value,
			"Ok",
		})

	default:
		return nil, fmt.Errorf("unsupported type for AppRequestResult %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *AppRequestResult) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "Error":
		var enumTypeValue ErrorAppRequestResult
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Ok":
		var enumTypeValue OkAppRequestResult
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for AppRequestResult %v", typeDescriptor.Type)
	}

	return nil
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

// Returns Core Library API reference.
func (c *Client) ClientConfig() (*Config, error) {
	result := new(Config)

	err := c.dllClient.waitErrorOrResultUnmarshal("client.config", nil, result)

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
