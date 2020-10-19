package client

// DON'T EDIT THIS FILE is generated 2020-10-19 13:51:06.679007 +0000 UTC
//
// Mod abi
//
// Provides message encoding and decoding according to the ABI
// specification.

import (
	"github.com/shopspring/decimal"
)

type Abi interface{}

type AbiHandle struct {
	int `json:""`
}

type FunctionHeader struct {
	// Message expiration time in seconds.
	Expire *int `json:"expire,omitempty"`
	// Message creation time in milliseconds.
	Time *decimal.Decimal `json:"time,omitempty"`
	// Public key used to sign message. Encoded with `hex`.
	Pubkey *string `json:"pubkey,omitempty"`
}

type CallSet struct {
	// Function name.
	FunctionName string `json:"function_name"`
	// Function header.
	//
	// If an application omit some parameters required by the
	// contract's ABI, the library will set the default values for
	// it.
	Header *FunctionHeader `json:"header,omitempty"`
	// Function input according to ABI.
	Input interface{} `json:"input,omitempty"`
}

type DeploySet struct {
	// Content of TVC file. Must be encoded with `base64`.
	Tvc string `json:"tvc"`
	// Target workchain for destination address. Default is `0`.
	WorkchainID *int `json:"workchain_id,omitempty"`
	// List of initial values for contract's public variables.
	InitialData interface{} `json:"initial_data,omitempty"`
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
	ProcessingTryIndex *int `json:"processing_try_index,omitempty"`
}

type ResultOfEncodeMessageBody struct {
	// Message body BOC encoded with `base64`.
	Body string `json:"body"`
	// Optional data to sign. Encoded with `base64`.
	//
	// Presents when `message` is unsigned. Can be used for external
	// message signing. Is this case you need to sing this data and
	// produce signed message using `abi.attach_signature`.
	DataToSign *string `json:"data_to_sign,omitempty"`
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
	// Contract address.
	//
	// Must be specified in case of non deploy message.
	Address *string `json:"address,omitempty"`
	// Deploy parameters.
	//
	// Must be specified in case of deploy message.
	DeploySet *DeploySet `json:"deploy_set,omitempty"`
	// Function call parameters.
	//
	// Must be specified in non deploy message.
	//
	// In case of deploy message contains parameters of constructor.
	CallSet *CallSet `json:"call_set,omitempty"`
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
	ProcessingTryIndex *int `json:"processing_try_index,omitempty"`
}

type ResultOfEncodeMessage struct {
	// Message BOC encoded with `base64`.
	Message string `json:"message"`
	// Optional data to sign. Encoded with `base64`.
	//
	// Presents when `message` is unsigned. Can be used for external
	// message signing. Is this case you need to sing this data and
	// produce signed message using `abi.attach_signature`.
	DataToSign *string `json:"data_to_sign,omitempty"`
	// Destination address.
	Address string `json:"address"`
	// Message id.
	MessageID string `json:"message_id"`
}

type ParamsOfAttachSignature struct {
	// Contract ABI.
	Abi Abi `json:"abi"`
	// Public key. Must be encoded with `hex`.
	PublicKey string `json:"public_key"`
	// Unsigned message BOC. Must be encoded with `base64`.
	Message string `json:"message"`
	// Signature. Must be encoded with `hex`.
	Signature string `json:"signature"`
}

type ResultOfAttachSignature struct {
	Message   string `json:"message"`
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
	Value interface{} `json:"value,omitempty"`
	// Function header.
	Header *FunctionHeader `json:"header,omitempty"`
}

type ParamsOfDecodeMessageBody struct {
	// Contract ABI used to decode.
	Abi Abi `json:"abi"`
	// Message body BOC. Must be encoded with `base64`.
	Body string `json:"body"`
	// True if the body belongs to the internal message.
	IsInternal bool `json:"is_internal"`
}

type ParamsOfEncodeAccount struct {
	// Source of the account state init.
	StateInit StateInitSource `json:"state_init"`
	// Initial balance.
	Balance *decimal.Decimal `json:"balance,omitempty"`
	// Initial value for the `last_trans_lt`.
	LastTransLt *decimal.Decimal `json:"last_trans_lt,omitempty"`
	// Initial value for the `last_paid`.
	LastPaid *int `json:"last_paid,omitempty"`
}

type ResultOfEncodeAccount struct {
	// Account BOC. Encoded with `base64`.
	Account string `json:"account"`
	// Account id. Encoded with `hex`.
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

func (c *Client) AbiEncodeMessage(p *ParamsOfEncodeMessage) (*ResultOfEncodeMessage, error) {
	response := new(ResultOfEncodeMessage)
	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_message", p, response)

	return response, err
}

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

// Encodes account state as it will be.
func (c *Client) AbiEncodeAccount(p *ParamsOfEncodeAccount) (*ResultOfEncodeAccount, error) {
	response := new(ResultOfEncodeAccount)
	err := c.dllClient.waitErrorOrResultUnmarshal("abi.encode_account", p, response)

	return response, err
}
