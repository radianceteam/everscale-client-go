package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 28 Aug 23 13:53 UTC
//
// Mod abi
//
// Provides message encoding and decoding according to the ABI specification.

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/volatiletech/null"
)

const (
	RequiredAddressMissingForEncodeMessageAbiErrorCode    = 301
	RequiredCallSetMissingForEncodeMessageAbiErrorCode    = 302
	InvalidJSONAbiErrorCode                               = 303
	InvalidMessageAbiErrorCode                            = 304
	EncodeDeployMessageFailedAbiErrorCode                 = 305
	EncodeRunMessageFailedAbiErrorCode                    = 306
	AttachSignatureFailedAbiErrorCode                     = 307
	InvalidTvcImageAbiErrorCode                           = 308
	RequiredPublicKeyMissingForFunctionHeaderAbiErrorCode = 309
	InvalidSignerAbiErrorCode                             = 310
	InvalidAbiAbiErrorCode                                = 311
	InvalidFunctionIDAbiErrorCode                         = 312
	InvalidDataAbiErrorCode                               = 313
	EncodeInitialDataFailedAbiErrorCode                   = 314
	InvalidFunctionNameAbiErrorCode                       = 315
)

func init() { // nolint gochecknoinits
	errorCodesToErrorTypes[RequiredAddressMissingForEncodeMessageAbiErrorCode] = "RequiredAddressMissingForEncodeMessageAbiErrorCode"
	errorCodesToErrorTypes[RequiredCallSetMissingForEncodeMessageAbiErrorCode] = "RequiredCallSetMissingForEncodeMessageAbiErrorCode"
	errorCodesToErrorTypes[InvalidJSONAbiErrorCode] = "InvalidJSONAbiErrorCode"
	errorCodesToErrorTypes[InvalidMessageAbiErrorCode] = "InvalidMessageAbiErrorCode"
	errorCodesToErrorTypes[EncodeDeployMessageFailedAbiErrorCode] = "EncodeDeployMessageFailedAbiErrorCode"
	errorCodesToErrorTypes[EncodeRunMessageFailedAbiErrorCode] = "EncodeRunMessageFailedAbiErrorCode"
	errorCodesToErrorTypes[AttachSignatureFailedAbiErrorCode] = "AttachSignatureFailedAbiErrorCode"
	errorCodesToErrorTypes[InvalidTvcImageAbiErrorCode] = "InvalidTvcImageAbiErrorCode"
	errorCodesToErrorTypes[RequiredPublicKeyMissingForFunctionHeaderAbiErrorCode] = "RequiredPublicKeyMissingForFunctionHeaderAbiErrorCode"
	errorCodesToErrorTypes[InvalidSignerAbiErrorCode] = "InvalidSignerAbiErrorCode"
	errorCodesToErrorTypes[InvalidAbiAbiErrorCode] = "InvalidAbiAbiErrorCode"
	errorCodesToErrorTypes[InvalidFunctionIDAbiErrorCode] = "InvalidFunctionIDAbiErrorCode"
	errorCodesToErrorTypes[InvalidDataAbiErrorCode] = "InvalidDataAbiErrorCode"
	errorCodesToErrorTypes[EncodeInitialDataFailedAbiErrorCode] = "EncodeInitialDataFailedAbiErrorCode"
	errorCodesToErrorTypes[InvalidFunctionNameAbiErrorCode] = "InvalidFunctionNameAbiErrorCode"
}

type AbiHandle uint32

// The ABI function header.
// Includes several hidden function parameters that contract
// uses for security, message delivery monitoring and replay protection reasons.
//
// The actual set of header fields depends on the contract's ABI.
// If a contract's ABI does not include some headers, then they are not filled.
type FunctionHeader struct {
	// Message expiration timestamp (UNIX time) in seconds.
	// If not specified - calculated automatically from message_expiration_timeout(),
	// try_index and message_expiration_timeout_grow_factor() (if ABI includes `expire` header).
	Expire null.Uint32 `json:"expire"` // optional
	// Message creation time in milliseconds.
	// If not specified, `now` is used (if ABI includes `time` header).
	Time *big.Int `json:"time"` // optional
	// Public key is used by the contract to check the signature.
	// Encoded in `hex`. If not specified, method fails with exception (if ABI includes `pubkey` header)..
	Pubkey null.String `json:"pubkey"` // optional
}

type CallSet struct {
	// Function name that is being called. Or function id encoded as string in hex (starting with 0x).
	FunctionName string `json:"function_name"`
	// Function header.
	// If an application omits some header parameters required by the
	// contract's ABI, the library will set the default values for
	// them.
	Header *FunctionHeader `json:"header"` // optional
	// Function input parameters according to ABI.
	Input json.RawMessage `json:"input"` // optional
}

type DeploySet struct {
	// Content of TVC file encoded in `base64`. For compatibility reason this field can contain an encoded  `StateInit`.
	Tvc null.String `json:"tvc"` // optional
	// Contract code BOC encoded with base64.
	Code null.String `json:"code"` // optional
	// State init BOC encoded with base64.
	StateInit null.String `json:"state_init"` // optional
	// Target workchain for destination address.
	// Default is `0`.
	WorkchainID null.Int32 `json:"workchain_id"` // optional
	// List of initial values for contract's public variables.
	InitialData json.RawMessage `json:"initial_data"` // optional
	// Optional public key that can be provided in deploy set in order to substitute one in TVM file or provided by Signer.
	// Public key resolving priority:
	// 1. Public key from deploy set.
	// 2. Public key, specified in TVM file.
	// 3. Public key, provided by Signer.
	InitialPubkey null.String `json:"initial_pubkey"` // optional
}

// No keys are provided.
// Creates an unsigned message.
type NoneSigner struct{}

// Only public key is provided in unprefixed hex string format to generate unsigned message and `data_to_sign` which can be signed later.
type ExternalSigner struct {
	PublicKey string `json:"public_key"`
}

// Key pair is provided for signing.
type KeysSigner struct {
	Keys KeyPair `json:"keys"`
}

// Signing Box interface is provided for signing, allows Dapps to sign messages using external APIs, such as HSM, cold wallet, etc.
type SigningBoxSigner struct {
	Handle SigningBoxHandle `json:"handle"`
}

type Signer struct {
	// Should be any of
	// NoneSigner
	// ExternalSigner
	// KeysSigner
	// SigningBoxSigner
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *Signer) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case NoneSigner:
		return json.Marshal(struct {
			NoneSigner
			Type string `json:"type"`
		}{
			value,
			"None",
		})

	case ExternalSigner:
		return json.Marshal(struct {
			ExternalSigner
			Type string `json:"type"`
		}{
			value,
			"External",
		})

	case KeysSigner:
		return json.Marshal(struct {
			KeysSigner
			Type string `json:"type"`
		}{
			value,
			"Keys",
		})

	case SigningBoxSigner:
		return json.Marshal(struct {
			SigningBoxSigner
			Type string `json:"type"`
		}{
			value,
			"SigningBox",
		})

	default:
		return nil, fmt.Errorf("unsupported type for Signer %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *Signer) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "None":
		var enumTypeValue NoneSigner
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "External":
		var enumTypeValue ExternalSigner
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Keys":
		var enumTypeValue KeysSigner
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "SigningBox":
		var enumTypeValue SigningBoxSigner
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for Signer %v", typeDescriptor.Type)
	}

	return nil
}

type MessageBodyType string

const (

	// Message contains the input of the ABI function.
	InputMessageBodyType MessageBodyType = "Input"
	// Message contains the output of the ABI function.
	OutputMessageBodyType MessageBodyType = "Output"
	// Message contains the input of the imported ABI function.
	// Occurs when contract sends an internal message to other
	// contract.
	InternalOutputMessageBodyType MessageBodyType = "InternalOutput"
	// Message contains the input of the ABI event.
	EventMessageBodyType MessageBodyType = "Event"
)

// Deploy message.
type MessageStateInitSource struct {
	Source MessageSource `json:"source"`
}

// State init data.
type StateInitStateInitSource struct {
	// Code BOC.
	// Encoded in `base64`.
	Code string `json:"code"`
	// Data BOC.
	// Encoded in `base64`.
	Data string `json:"data"`
	// Library BOC.
	// Encoded in `base64`.
	Library null.String `json:"library"` // optional
}

// Content of the TVC file.
// Encoded in `base64`.
type TvcStateInitSource struct {
	Tvc        string           `json:"tvc"`
	PublicKey  null.String      `json:"public_key"`  // optional
	InitParams *StateInitParams `json:"init_params"` // optional
}

type StateInitSource struct {
	// Should be any of
	// MessageStateInitSource
	// StateInitStateInitSource
	// TvcStateInitSource
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *StateInitSource) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case MessageStateInitSource:
		return json.Marshal(struct {
			MessageStateInitSource
			Type string `json:"type"`
		}{
			value,
			"Message",
		})

	case StateInitStateInitSource:
		return json.Marshal(struct {
			StateInitStateInitSource
			Type string `json:"type"`
		}{
			value,
			"StateInit",
		})

	case TvcStateInitSource:
		return json.Marshal(struct {
			TvcStateInitSource
			Type string `json:"type"`
		}{
			value,
			"Tvc",
		})

	default:
		return nil, fmt.Errorf("unsupported type for StateInitSource %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *StateInitSource) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "Message":
		var enumTypeValue MessageStateInitSource
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "StateInit":
		var enumTypeValue StateInitStateInitSource
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Tvc":
		var enumTypeValue TvcStateInitSource
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for StateInitSource %v", typeDescriptor.Type)
	}

	return nil
}

type StateInitParams struct {
	Abi   Abi             `json:"abi"`
	Value json.RawMessage `json:"value"`
}

type EncodedMessageSource struct {
	Message string `json:"message"`
	Abi     *Abi   `json:"abi"` // optional
}

type MessageSource struct {
	// Should be any of
	// EncodedMessageSource
	// ParamsOfEncodeMessage
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *MessageSource) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case EncodedMessageSource:
		return json.Marshal(struct {
			EncodedMessageSource
			Type string `json:"type"`
		}{
			value,
			"Encoded",
		})

	case ParamsOfEncodeMessage:
		return json.Marshal(struct {
			ParamsOfEncodeMessage
			Type string `json:"type"`
		}{
			value,
			"EncodingParams",
		})

	default:
		return nil, fmt.Errorf("unsupported type for MessageSource %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *MessageSource) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "Encoded":
		var enumTypeValue EncodedMessageSource
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "EncodingParams":
		var enumTypeValue ParamsOfEncodeMessage
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for MessageSource %v", typeDescriptor.Type)
	}

	return nil
}

type AbiParam struct {
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Components []AbiParam `json:"components"` // optional
}

type AbiEvent struct {
	Name   string       `json:"name"`
	Inputs []AbiParam   `json:"inputs"`
	ID     *null.String `json:"id"` // optional
}

type AbiData struct {
	Key        uint32     `json:"key"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Components []AbiParam `json:"components"` // optional
}

type AbiFunction struct {
	Name    string       `json:"name"`
	Inputs  []AbiParam   `json:"inputs"`
	Outputs []AbiParam   `json:"outputs"`
	ID      *null.String `json:"id"` // optional
}

type AbiContract struct {
	ABIVersion null.Uint32   `json:"ABI version"` // optional
	AbiVersion null.Uint32   `json:"abi_version"` // optional
	Version    *null.String  `json:"version"`     // optional
	Header     []string      `json:"header"`      // optional
	Functions  []AbiFunction `json:"functions"`   // optional
	Events     []AbiEvent    `json:"events"`      // optional
	Data       []AbiData     `json:"data"`        // optional
	Fields     []AbiParam    `json:"fields"`      // optional
}

type DataLayout string

const (

	// Decode message body as function input parameters.
	InputDataLayout DataLayout = "Input"
	// Decode message body as function output.
	OutputDataLayout DataLayout = "Output"
)

type ParamsOfEncodeMessageBody struct {
	// Contract ABI.
	Abi Abi `json:"abi"`
	// Function call parameters.
	// Must be specified in non deploy message.
	//
	// In case of deploy message contains parameters of constructor.
	CallSet CallSet `json:"call_set"`
	// True if internal message body must be encoded.
	IsInternal bool `json:"is_internal"`
	// Signing parameters.
	Signer Signer `json:"signer"`
	// Processing try index.
	// Used in message processing with retries.
	//
	// Encoder uses the provided try index to calculate message
	// expiration time.
	//
	// Expiration timeouts will grow with every retry.
	//
	// Default value is 0.
	ProcessingTryIndex null.Uint8 `json:"processing_try_index"` // optional
	// Destination address of the message.
	// Since ABI version 2.3 destination address of external inbound message is used in message
	// body signature calculation. Should be provided when signed external inbound message body is
	// created. Otherwise can be omitted.
	Address null.String `json:"address"` // optional
	// Signature ID to be used in data to sign preparing when CapSignatureWithId capability is enabled.
	SignatureID null.Int32 `json:"signature_id"` // optional
}

type ResultOfEncodeMessageBody struct {
	// Message body BOC encoded with `base64`.
	Body string `json:"body"`
	// Optional data to sign.
	// Encoded with `base64`.
	// Presents when `message` is unsigned. Can be used for external
	// message signing. Is this case you need to sing this data and
	// produce signed message using `abi.attach_signature`.
	DataToSign null.String `json:"data_to_sign"` // optional
}

type ParamsOfAttachSignatureToMessageBody struct {
	// Contract ABI.
	Abi Abi `json:"abi"`
	// Public key.
	// Must be encoded with `hex`.
	PublicKey string `json:"public_key"`
	// Unsigned message body BOC.
	// Must be encoded with `base64`.
	Message string `json:"message"`
	// Signature.
	// Must be encoded with `hex`.
	Signature string `json:"signature"`
}

type ResultOfAttachSignatureToMessageBody struct {
	Body string `json:"body"`
}

type ParamsOfEncodeMessage struct {
	// Contract ABI.
	Abi Abi `json:"abi"`
	// Target address the message will be sent to.
	// Must be specified in case of non-deploy message.
	Address null.String `json:"address"` // optional
	// Deploy parameters.
	// Must be specified in case of deploy message.
	DeploySet *DeploySet `json:"deploy_set"` // optional
	// Function call parameters.
	// Must be specified in case of non-deploy message.
	//
	// In case of deploy message it is optional and contains parameters
	// of the functions that will to be called upon deploy transaction.
	CallSet *CallSet `json:"call_set"` // optional
	// Signing parameters.
	Signer Signer `json:"signer"`
	// Processing try index.
	// Used in message processing with retries (if contract's ABI includes "expire" header).
	//
	// Encoder uses the provided try index to calculate message
	// expiration time. The 1st message expiration time is specified in
	// Client config.
	//
	// Expiration timeouts will grow with every retry.
	// Retry grow factor is set in Client config:
	// <.....add config parameter with default value here>
	//
	// Default value is 0.
	ProcessingTryIndex null.Uint8 `json:"processing_try_index"` // optional
	// Signature ID to be used in data to sign preparing when CapSignatureWithId capability is enabled.
	SignatureID null.Int32 `json:"signature_id"` // optional
}

type ResultOfEncodeMessage struct {
	// Message BOC encoded with `base64`.
	Message string `json:"message"`
	// Optional data to be signed encoded in `base64`.
	// Returned in case of `Signer::External`. Can be used for external
	// message signing. Is this case you need to use this data to create signature and
	// then produce signed message using `abi.attach_signature`.
	DataToSign null.String `json:"data_to_sign"` // optional
	// Destination address.
	Address string `json:"address"`
	// Message id.
	MessageID string `json:"message_id"`
}

type ParamsOfEncodeInternalMessage struct {
	// Contract ABI.
	// Can be None if both deploy_set and call_set are None.
	Abi *Abi `json:"abi"` // optional
	// Target address the message will be sent to.
	// Must be specified in case of non-deploy message.
	Address null.String `json:"address"` // optional
	// Source address of the message.
	SrcAddress null.String `json:"src_address"` // optional
	// Deploy parameters.
	// Must be specified in case of deploy message.
	DeploySet *DeploySet `json:"deploy_set"` // optional
	// Function call parameters.
	// Must be specified in case of non-deploy message.
	//
	// In case of deploy message it is optional and contains parameters
	// of the functions that will to be called upon deploy transaction.
	CallSet *CallSet `json:"call_set"` // optional
	// Value in nanotokens to be sent with message.
	Value string `json:"value"`
	// Flag of bounceable message.
	// Default is true.
	Bounce null.Bool `json:"bounce"` // optional
	// Enable Instant Hypercube Routing for the message.
	// Default is false.
	EnableIhr null.Bool `json:"enable_ihr"` // optional
}

type ResultOfEncodeInternalMessage struct {
	// Message BOC encoded with `base64`.
	Message string `json:"message"`
	// Destination address.
	Address string `json:"address"`
	// Message id.
	MessageID string `json:"message_id"`
}

type ParamsOfAttachSignature struct {
	// Contract ABI.
	Abi Abi `json:"abi"`
	// Public key encoded in `hex`.
	PublicKey string `json:"public_key"`
	// Unsigned message BOC encoded in `base64`.
	Message string `json:"message"`
	// Signature encoded in `hex`.
	Signature string `json:"signature"`
}

type ResultOfAttachSignature struct {
	// Signed message BOC.
	Message string `json:"message"`
	// Message ID.
	MessageID string `json:"message_id"`
}

type ParamsOfDecodeMessage struct {
	// contract ABI.
	Abi Abi `json:"abi"`
	// Message BOC.
	Message string `json:"message"`
	// Flag allowing partial BOC decoding when ABI doesn't describe the full body BOC. Controls decoder behaviour when after decoding all described in ABI params there are some data left in BOC: `true` - return decoded values `false` - return error of incomplete BOC deserialization (default).
	AllowPartial null.Bool `json:"allow_partial"` // optional
	// Function name or function id if is known in advance.
	FunctionName null.String `json:"function_name"` // optional
	DataLayout   *DataLayout `json:"data_layout"`   // optional
}

type DecodedMessageBody struct {
	// Type of the message body content.
	BodyType MessageBodyType `json:"body_type"`
	// Function or event name.
	Name string `json:"name"`
	// Parameters or result value.
	Value json.RawMessage `json:"value"` // optional
	// Function header.
	Header *FunctionHeader `json:"header"` // optional
}

type ParamsOfDecodeMessageBody struct {
	// Contract ABI used to decode.
	Abi Abi `json:"abi"`
	// Message body BOC encoded in `base64`.
	Body string `json:"body"`
	// True if the body belongs to the internal message.
	IsInternal bool `json:"is_internal"`
	// Flag allowing partial BOC decoding when ABI doesn't describe the full body BOC. Controls decoder behaviour when after decoding all described in ABI params there are some data left in BOC: `true` - return decoded values `false` - return error of incomplete BOC deserialization (default).
	AllowPartial null.Bool `json:"allow_partial"` // optional
	// Function name or function id if is known in advance.
	FunctionName null.String `json:"function_name"` // optional
	DataLayout   *DataLayout `json:"data_layout"`   // optional
}

type ParamsOfEncodeAccount struct {
	// Source of the account state init.
	StateInit StateInitSource `json:"state_init"`
	// Initial balance.
	Balance *big.Int `json:"balance"` // optional
	// Initial value for the `last_trans_lt`.
	LastTransLt *big.Int `json:"last_trans_lt"` // optional
	// Initial value for the `last_paid`.
	LastPaid null.Uint32 `json:"last_paid"` // optional
	// Cache type to put the result.
	// The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
}

type ResultOfEncodeAccount struct {
	// Account BOC encoded in `base64`.
	Account string `json:"account"`
	// Account ID  encoded in `hex`.
	ID string `json:"id"`
}

type ParamsOfDecodeAccountData struct {
	// Contract ABI.
	Abi Abi `json:"abi"`
	// Data BOC or BOC handle.
	Data string `json:"data"`
	// Flag allowing partial BOC decoding when ABI doesn't describe the full body BOC. Controls decoder behaviour when after decoding all described in ABI params there are some data left in BOC: `true` - return decoded values `false` - return error of incomplete BOC deserialization (default).
	AllowPartial null.Bool `json:"allow_partial"` // optional
}

type ResultOfDecodeAccountData struct {
	// Decoded data as a JSON structure.
	Data json.RawMessage `json:"data"`
}

type ParamsOfUpdateInitialData struct {
	// Contract ABI.
	Abi *Abi `json:"abi"` // optional
	// Data BOC or BOC handle.
	Data string `json:"data"`
	// List of initial values for contract's static variables.
	// `abi` parameter should be provided to set initial data.
	InitialData json.RawMessage `json:"initial_data"` // optional
	// Initial account owner's public key to set into account data.
	InitialPubkey null.String `json:"initial_pubkey"` // optional
	// Cache type to put the result. The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
}

type ResultOfUpdateInitialData struct {
	// Updated data BOC or BOC handle.
	Data string `json:"data"`
}

type ParamsOfEncodeInitialData struct {
	// Contract ABI.
	Abi *Abi `json:"abi"` // optional
	// List of initial values for contract's static variables.
	// `abi` parameter should be provided to set initial data.
	InitialData json.RawMessage `json:"initial_data"` // optional
	// Initial account owner's public key to set into account data.
	InitialPubkey null.String `json:"initial_pubkey"` // optional
	// Cache type to put the result. The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
}

type ResultOfEncodeInitialData struct {
	// Updated data BOC or BOC handle.
	Data string `json:"data"`
}

type ParamsOfDecodeInitialData struct {
	// Contract ABI.
	// Initial data is decoded if this parameter is provided.
	Abi *Abi `json:"abi"` // optional
	// Data BOC or BOC handle.
	Data string `json:"data"`
	// Flag allowing partial BOC decoding when ABI doesn't describe the full body BOC. Controls decoder behaviour when after decoding all described in ABI params there are some data left in BOC: `true` - return decoded values `false` - return error of incomplete BOC deserialization (default).
	AllowPartial null.Bool `json:"allow_partial"` // optional
}

type ResultOfDecodeInitialData struct {
	// List of initial values of contract's public variables.
	// Initial data is decoded if `abi` input parameter is provided.
	InitialData json.RawMessage `json:"initial_data"` // optional
	// Initial account owner's public key.
	InitialPubkey string `json:"initial_pubkey"`
}

type ParamsOfDecodeBoc struct {
	// Parameters to decode from BOC.
	Params []AbiParam `json:"params"`
	// Data BOC or BOC handle.
	Boc          string `json:"boc"`
	AllowPartial bool   `json:"allow_partial"`
}

type ResultOfDecodeBoc struct {
	// Decoded data as a JSON structure.
	Data json.RawMessage `json:"data"`
}

type ParamsOfAbiEncodeBoc struct {
	// Parameters to encode into BOC.
	Params []AbiParam `json:"params"`
	// Parameters and values as a JSON structure.
	Data json.RawMessage `json:"data"`
	// Cache type to put the result.
	// The BOC itself returned if no cache type provided.
	BocCache *BocCacheType `json:"boc_cache"` // optional
}

type ResultOfAbiEncodeBoc struct {
	// BOC encoded as base64.
	Boc string `json:"boc"`
}

type ParamsOfCalcFunctionId struct {
	// Contract ABI.
	Abi Abi `json:"abi"`
	// Contract function name.
	FunctionName string `json:"function_name"`
	// If set to `true` output function ID will be returned which is used in contract response. Default is `false`.
	Output null.Bool `json:"output"` // optional
}

type ResultOfCalcFunctionId struct {
	// Contract function ID.
	FunctionID uint32 `json:"function_id"`
}

type ParamsOfGetSignatureData struct {
	// Contract ABI used to decode.
	Abi Abi `json:"abi"`
	// Message BOC encoded in `base64`.
	Message string `json:"message"`
	// Signature ID to be used in unsigned data preparing when CapSignatureWithId capability is enabled.
	SignatureID null.Int32 `json:"signature_id"` // optional
}

type ResultOfGetSignatureData struct {
	// Signature from the message in `hex`.
	Signature string `json:"signature"`
	// Data to verify the signature in `base64`.
	Unsigned string `json:"unsigned"`
}

// Encodes message body according to ABI function call.
func (c *Client) AbiEncodeMessageBody(p *ParamsOfEncodeMessageBody) (*ResultOfEncodeMessageBody, error) {
	result := new(ResultOfEncodeMessageBody)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_message_body", p, result)

	return result, err
}

func (c *Client) AbiAttachSignatureToMessageBody(p *ParamsOfAttachSignatureToMessageBody) (*ResultOfAttachSignatureToMessageBody, error) {
	result := new(ResultOfAttachSignatureToMessageBody)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.attach_signature_to_message_body", p, result)

	return result, err
}

// Encodes an ABI-compatible message.
// Allows to encode deploy and function call messages,
// both signed and unsigned.
//
// Use cases include messages of any possible type:
// - deploy with initial function call (i.e. `constructor` or any other function that is used for some kind
// of initialization);
// - deploy without initial function call;
// - signed/unsigned + data for signing.
//
// `Signer` defines how the message should or shouldn't be signed:
//
// `Signer::None` creates an unsigned message. This may be needed in case of some public methods,
// that do not require authorization by pubkey.
//
// `Signer::External` takes public key and returns `data_to_sign` for later signing.
// Use `attach_signature` method with the result signature to get the signed message.
//
// `Signer::Keys` creates a signed message with provided key pair.
//
// [SOON] `Signer::SigningBox` Allows using a special interface to implement signing
// without private key disclosure to SDK. For instance, in case of using a cold wallet or HSM,
// when application calls some API to sign data.
//
// There is an optional public key can be provided in deploy set in order to substitute one
// in TVM file.
//
// Public key resolving priority:
// 1. Public key from deploy set.
// 2. Public key, specified in TVM file.
// 3. Public key, provided by signer.
func (c *Client) AbiEncodeMessage(p *ParamsOfEncodeMessage) (*ResultOfEncodeMessage, error) {
	result := new(ResultOfEncodeMessage)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_message", p, result)

	return result, err
}

// Encodes an internal ABI-compatible message.
// Allows to encode deploy and function call messages.
//
// Use cases include messages of any possible type:
// - deploy with initial function call (i.e. `constructor` or any other function that is used for some kind
// of initialization);
// - deploy without initial function call;
// - simple function call
//
// There is an optional public key can be provided in deploy set in order to substitute one
// in TVM file.
//
// Public key resolving priority:
// 1. Public key from deploy set.
// 2. Public key, specified in TVM file.
func (c *Client) AbiEncodeInternalMessage(p *ParamsOfEncodeInternalMessage) (*ResultOfEncodeInternalMessage, error) {
	result := new(ResultOfEncodeInternalMessage)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_internal_message", p, result)

	return result, err
}

// Combines `hex`-encoded `signature` with `base64`-encoded `unsigned_message`. Returns signed message encoded in `base64`.
func (c *Client) AbiAttachSignature(p *ParamsOfAttachSignature) (*ResultOfAttachSignature, error) {
	result := new(ResultOfAttachSignature)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.attach_signature", p, result)

	return result, err
}

// Decodes message body using provided message BOC and ABI.
func (c *Client) AbiDecodeMessage(p *ParamsOfDecodeMessage) (*DecodedMessageBody, error) {
	result := new(DecodedMessageBody)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.decode_message", p, result)

	return result, err
}

// Decodes message body using provided body BOC and ABI.
func (c *Client) AbiDecodeMessageBody(p *ParamsOfDecodeMessageBody) (*DecodedMessageBody, error) {
	result := new(DecodedMessageBody)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.decode_message_body", p, result)

	return result, err
}

// Creates account state BOC.
// Creates account state provided with one of these sets of data :
// 1. BOC of code, BOC of data, BOC of library
// 2. TVC (string in `base64`), keys, init params.
func (c *Client) AbiEncodeAccount(p *ParamsOfEncodeAccount) (*ResultOfEncodeAccount, error) {
	result := new(ResultOfEncodeAccount)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_account", p, result)

	return result, err
}

// Decodes account data using provided data BOC and ABI.
// Note: this feature requires ABI 2.1 or higher.
func (c *Client) AbiDecodeAccountData(p *ParamsOfDecodeAccountData) (*ResultOfDecodeAccountData, error) {
	result := new(ResultOfDecodeAccountData)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.decode_account_data", p, result)

	return result, err
}

// Updates initial account data with initial values for the contract's static variables and owner's public key. This operation is applicable only for initial account data (before deploy). If the contract is already deployed, its data doesn't contain this data section any more.
func (c *Client) AbiUpdateInitialData(p *ParamsOfUpdateInitialData) (*ResultOfUpdateInitialData, error) {
	result := new(ResultOfUpdateInitialData)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.update_initial_data", p, result)

	return result, err
}

// Encodes initial account data with initial values for the contract's static variables and owner's public key into a data BOC that can be passed to `encode_tvc` function afterwards.
// This function is analogue of `tvm.buildDataInit` function in Solidity.
func (c *Client) AbiEncodeInitialData(p *ParamsOfEncodeInitialData) (*ResultOfEncodeInitialData, error) {
	result := new(ResultOfEncodeInitialData)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_initial_data", p, result)

	return result, err
}

// Decodes initial values of a contract's static variables and owner's public key from account initial data This operation is applicable only for initial account data (before deploy). If the contract is already deployed, its data doesn't contain this data section any more.
func (c *Client) AbiDecodeInitialData(p *ParamsOfDecodeInitialData) (*ResultOfDecodeInitialData, error) {
	result := new(ResultOfDecodeInitialData)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.decode_initial_data", p, result)

	return result, err
}

// Decodes BOC into JSON as a set of provided parameters.
// Solidity functions use ABI types for [builder encoding](https://github.com/tonlabs/TON-Solidity-Compiler/blob/master/API.md#tvmbuilderstore).
// The simplest way to decode such a BOC is to use ABI decoding.
// ABI has it own rules for fields layout in cells so manually encoded
// BOC can not be described in terms of ABI rules.
//
// To solve this problem we introduce a new ABI type `Ref(<ParamType>)`
// which allows to store `ParamType` ABI parameter in cell reference and, thus,
// decode manually encoded BOCs. This type is available only in `decode_boc` function
// and will not be available in ABI messages encoding until it is included into some ABI revision.
//
// Such BOC descriptions covers most users needs. If someone wants to decode some BOC which
// can not be described by these rules (i.e. BOC with TLB containing constructors of flags
// defining some parsing conditions) then they can decode the fields up to fork condition,
// check the parsed data manually, expand the parsing schema and then decode the whole BOC
// with the full schema.
func (c *Client) AbiDecodeBoc(p *ParamsOfDecodeBoc) (*ResultOfDecodeBoc, error) {
	result := new(ResultOfDecodeBoc)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.decode_boc", p, result)

	return result, err
}

// Encodes given parameters in JSON into a BOC using param types from ABI.
func (c *Client) AbiEncodeBoc(p *ParamsOfAbiEncodeBoc) (*ResultOfAbiEncodeBoc, error) {
	result := new(ResultOfAbiEncodeBoc)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_boc", p, result)

	return result, err
}

// Calculates contract function ID by contract ABI.
func (c *Client) AbiCalcFunctionId(p *ParamsOfCalcFunctionId) (*ResultOfCalcFunctionId, error) {
	result := new(ResultOfCalcFunctionId)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.calc_function_id", p, result)

	return result, err
}

// Extracts signature from message body and calculates hash to verify the signature.
func (c *Client) AbiGetSignatureData(p *ParamsOfGetSignatureData) (*ResultOfGetSignatureData, error) {
	result := new(ResultOfGetSignatureData)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.get_signature_data", p, result)

	return result, err
}
