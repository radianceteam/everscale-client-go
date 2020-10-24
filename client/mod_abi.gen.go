package client

// DON'T EDIT THIS FILE is generated 24 Oct 20 12:36 UTC
//
// Mod abi
//
// Provides message encoding and decoding according to the ABI
// specification.

import (
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

type Abi interface{}

type AbiHandle struct {
	int `json:""`
}

type FunctionHeader struct {
	// Message expiration time in seconds.
	Expire null.Int `json:"expire"` // optional
	// Message creation time in milliseconds.
	Time decimal.NullDecimal `json:"time"` // optional
	// Public key used to sign message. Encoded with `hex`.
	Pubkey null.String `json:"pubkey"` // optional
}

type CallSet struct {
	// Function name that is being called.
	FunctionName string `json:"function_name"`
	// Function header.
	//
	// If an application omits some header parameters required by the
	// contract's ABI, the library will set the default values for
	// them.
	Header *FunctionHeader `json:"header"` // optional
	// Function input parameters according to ABI.
	Input interface{} `json:"input"` // optional
}

type DeploySet struct {
	// Content of TVC file encoded in `base64`.
	Tvc string `json:"tvc"`
	// Target workchain for destination address. Default is `0`.
	WorkchainID null.Int `json:"workchain_id"` // optional
	// List of initial values for contract's public variables.
	InitialData interface{} `json:"initial_data"` // optional
}

type Signer interface{}

type MessageBodyType string

const (
	Input          MessageBodyType = "Input"
	Output         MessageBodyType = "Output"
	InternalOutput MessageBodyType = "InternalOutput"
	Event          MessageBodyType = "Event"
)

type StateInitSource interface{}

type StateInitParams struct {
	Abi   Abi         `json:"abi"`
	Value interface{} `json:"value"`
}

type MessageSource interface{}

type ParamsOfEncodeMessageBody struct {
	// Contract ABI.
	Abi Abi `json:"abi"`
	// Function call parameters.
	//
	// Must be specified in non deploy message.
	//
	// In case of deploy message contains parameters of constructor.
	CallSet CallSet `json:"call_set"`
	// True if internal message body must be encoded.
	IsInternal bool `json:"is_internal"`
	// Signing parameters.
	Signer Signer `json:"signer"`
	// Processing try index.
	//
	// Used in message processing with retries.
	//
	// Encoder uses the provided try index to calculate message
	// expiration time.
	//
	// Expiration timeouts will grow with every retry.
	//
	// Default value is 0.
	ProcessingTryIndex null.Int `json:"processing_try_index"` // optional
}

type ResultOfEncodeMessageBody struct {
	// Message body BOC encoded with `base64`.
	Body string `json:"body"`
	// Optional data to sign. Encoded with `base64`.
	//
	// Presents when `message` is unsigned. Can be used for external
	// message signing. Is this case you need to sing this data and
	// produce signed message using `abi.attach_signature`.
	DataToSign null.String `json:"data_to_sign"` // optional
}

type ParamsOfAttachSignatureToMessageBody struct {
	// Contract ABI.
	Abi Abi `json:"abi"`
	// Public key. Must be encoded with `hex`.
	PublicKey string `json:"public_key"`
	// Unsigned message BOC. Must be encoded with `base64`.
	Message string `json:"message"`
	// Signature. Must be encoded with `hex`.
	Signature string `json:"signature"`
}

type ResultOfAttachSignatureToMessageBody struct {
	Body string `json:"body"`
}

type ParamsOfEncodeMessage struct {
	// Contract ABI.
	Abi Abi `json:"abi"`
	// Target address the message will be sent to.
	//
	// Must be specified in case of non-deploy message.
	Address null.String `json:"address"` // optional
	// Deploy parameters.
	//
	// Must be specified in case of deploy message.
	DeploySet *DeploySet `json:"deploy_set"` // optional
	// Function call parameters.
	//
	// Must be specified in case of non-deploy message.
	//
	// In case of deploy message it is optional and contains parameters
	// of the functions that will to be called upon deploy transaction.
	CallSet *CallSet `json:"call_set"` // optional
	// Signing parameters.
	Signer Signer `json:"signer"`
	// Processing try index.
	//
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
	ProcessingTryIndex null.Int `json:"processing_try_index"` // optional
}

type ResultOfEncodeMessage struct {
	// Message BOC encoded with `base64`.
	Message string `json:"message"`
	// Optional data to be signed encoded in `base64`.
	//
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
	Value interface{} `json:"value"` // optional
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
	Balance decimal.NullDecimal `json:"balance"` // optional
	// Initial value for the `last_trans_lt`.
	LastTransLt decimal.NullDecimal `json:"last_trans_lt"` // optional
	// Initial value for the `last_paid`.
	LastPaid null.Int `json:"last_paid"` // optional
}

type ResultOfEncodeAccount struct {
	// Account BOC encoded in `base64`.
	Account string `json:"account"`
	// Account ID  encoded in `hex`.
	ID string `json:"id"`
}

// Encodes message body according to ABI function call.
func (c *Client) AbiEncodeMessageBody(p *ParamsOfEncodeMessageBody) (*ResultOfEncodeMessageBody, error) {
	response := new(ResultOfEncodeMessageBody)
	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_message_body", p, response)

	return response, err
}

func (c *Client) AbiAttachSignatureToMessageBody(p *ParamsOfAttachSignatureToMessageBody) (*ResultOfAttachSignatureToMessageBody, error) {
	response := new(ResultOfAttachSignatureToMessageBody)
	err := c.dllClient.waitErrorOrResultUnmarshal("abi.attach_signature_to_message_body", p, response)

	return response, err
}

// Encodes an ABI-compatible message
//
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
	response := new(ResultOfEncodeMessage)
	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_message", p, response)

	return response, err
}

// Combines `hex`-encoded `signature` with `base64`-encoded `unsigned_message`.
// Returns signed message encoded in `base64`.
func (c *Client) AbiAttachSignature(p *ParamsOfAttachSignature) (*ResultOfAttachSignature, error) {
	response := new(ResultOfAttachSignature)
	err := c.dllClient.waitErrorOrResultUnmarshal("abi.attach_signature", p, response)

	return response, err
}

// Decodes message body using provided message BOC and ABI.
func (c *Client) AbiDecodeMessage(p *ParamsOfDecodeMessage) (*DecodedMessageBody, error) {
	response := new(DecodedMessageBody)
	err := c.dllClient.waitErrorOrResultUnmarshal("abi.decode_message", p, response)

	return response, err
}

// Decodes message body using provided body BOC and ABI.
func (c *Client) AbiDecodeMessageBody(p *ParamsOfDecodeMessageBody) (*DecodedMessageBody, error) {
	response := new(DecodedMessageBody)
	err := c.dllClient.waitErrorOrResultUnmarshal("abi.decode_message_body", p, response)

	return response, err
}

// Creates account state BOC
//
// Creates account state provided with one of these sets of data :
// 1. BOC of code, BOC of data, BOC of library
// 2. TVC (string in `base64`), keys, init params.
func (c *Client) AbiEncodeAccount(p *ParamsOfEncodeAccount) (*ResultOfEncodeAccount, error) {
	response := new(ResultOfEncodeAccount)
	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_account", p, response)

	return response, err
}
