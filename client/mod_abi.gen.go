package client

// DON'T EDIT THIS FILE is generated 03 Jan 21 17:49 UTC
//
// Mod abi
//
// Provides message encoding and decoding according to the ABI specification.

import (
	"encoding/json"
	"math/big"

	"github.com/volatiletech/null"
)

type AbiErrorCode string

const (
	RequiredAddressMissingForEncodeMessageAbiErrorCode    AbiErrorCode = "RequiredAddressMissingForEncodeMessage"
	RequiredCallSetMissingForEncodeMessageAbiErrorCode    AbiErrorCode = "RequiredCallSetMissingForEncodeMessage"
	InvalidJSONAbiErrorCode                               AbiErrorCode = "InvalidJson"
	InvalidMessageAbiErrorCode                            AbiErrorCode = "InvalidMessage"
	EncodeDeployMessageFailedAbiErrorCode                 AbiErrorCode = "EncodeDeployMessageFailed"
	EncodeRunMessageFailedAbiErrorCode                    AbiErrorCode = "EncodeRunMessageFailed"
	AttachSignatureFailedAbiErrorCode                     AbiErrorCode = "AttachSignatureFailed"
	InvalidTvcImageAbiErrorCode                           AbiErrorCode = "InvalidTvcImage"
	RequiredPublicKeyMissingForFunctionHeaderAbiErrorCode AbiErrorCode = "RequiredPublicKeyMissingForFunctionHeader"
	InvalidSignerAbiErrorCode                             AbiErrorCode = "InvalidSigner"
	InvalidAbiAbiErrorCode                                AbiErrorCode = "InvalidAbi"
)

type FunctionHeader struct {
	// Message expiration time in seconds. If not specified - calculated automatically from message_expiration_timeout(), try_index and message_expiration_timeout_grow_factor() (if ABI includes `expire` header).
	Expire null.Uint32 `json:"expire"` // optional
	// Message creation time in milliseconds.
	// If not specified, `now` is used(if ABI includes `time` header).
	Time *big.Int `json:"time"` // optional
	// Public key is used by the contract to check the signature.
	// Encoded in `hex`.If not specified, method fails with exception (if ABI includes `pubkey` header)..
	Pubkey null.String `json:"pubkey"` // optional
}

type CallSet struct {
	// Function name that is being called.
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
	// Content of TVC file encoded in `base64`.
	Tvc string `json:"tvc"`
	// Target workchain for destination address.
	// Default is `0`.
	WorkchainID null.Int32 `json:"workchain_id"` // optional
	// List of initial values for contract's public variables.
	InitialData json.RawMessage `json:"initial_data"` // optional
}

type SignerType string

const (

	// No keys are provided.
	// Creates an unsigned message.
	NoneSignerType SignerType = "None"
	// Only public key is provided in unprefixed hex string format to generate unsigned message and `data_to_sign` which can be signed later.
	ExternalSignerType SignerType = "External"
	// Key pair is provided for signing.
	KeysSignerType SignerType = "Keys"
	// Signing Box interface is provided for signing, allows Dapps to sign messages using external APIs, such as HSM, cold wallet, etc.
	SigningBoxSignerType SignerType = "SigningBox"
)

type Signer struct {
	Type SignerType `json:"type"`
	// presented in types:
	// "External".
	PublicKey string `json:"public_key"`
	// presented in types:
	// "Keys".
	Keys KeyPair `json:"keys"`
	// presented in types:
	// "SigningBox".
	Handle SigningBoxHandle `json:"handle"`
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

type StateInitSourceType string

const (

	// Deploy message.
	MessageStateInitSourceType StateInitSourceType = "Message"
	// State init data.
	StateInitStateInitSourceType StateInitSourceType = "StateInit"
	// Content of the TVC file.
	// Encoded in `base64`.
	TvcStateInitSourceType StateInitSourceType = "Tvc"
)

type StateInitSource struct {
	Type StateInitSourceType `json:"type"`
	// presented in types:
	// "Message".
	Source MessageSource `json:"source"`
	// Code BOC.
	// Encoded in `base64`. presented in types:
	// "StateInit".
	Code string `json:"code"`
	// Data BOC.
	// Encoded in `base64`. presented in types:
	// "StateInit".
	Data string `json:"data"`
	// Library BOC.
	// Encoded in `base64`. presented in types:
	// "StateInit".
	Library null.String `json:"library"` // optional
	// presented in types:
	// "Tvc".
	Tvc string `json:"tvc"`
	// presented in types:
	// "Tvc".
	PublicKey null.String `json:"public_key"` // optional
	// presented in types:
	// "Tvc".
	InitParams *StateInitParams `json:"init_params"` // optional
}

type StateInitParams struct {
	Abi   Abi             `json:"abi"`
	Value json.RawMessage `json:"value"`
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
	Key        big.Int    `json:"key"`
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
	Header     []string      `json:"header"`      // optional
	Functions  []AbiFunction `json:"functions"`   // optional
	Events     []AbiEvent    `json:"events"`      // optional
	Data       []AbiData     `json:"data"`        // optional
}

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
}

type ResultOfEncodeAccount struct {
	// Account BOC encoded in `base64`.
	Account string `json:"account"`
	// Account ID  encoded in `hex`.
	ID string `json:"id"`
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
// [SOON] `Signer::SigningBox` Allows using a special interface to imlepement signing
// without private key disclosure to SDK. For instance, in case of using a cold wallet or HSM,
// when application calls some API to sign data.
func (c *Client) AbiEncodeMessage(p *ParamsOfEncodeMessage) (*ResultOfEncodeMessage, error) {
	result := new(ResultOfEncodeMessage)

	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_message", p, result)

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
