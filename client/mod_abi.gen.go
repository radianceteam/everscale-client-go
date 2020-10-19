package client

// DON'T EDIT THIS FILE is generated 2020-10-19 10:03:07.682726 +0000 UTC
// Mod abi
//  Provides message encoding and decoding according to the ABI
//  Provides message encoding and decoding according to the ABI
//  specification.

import (
	"github.com/shopspring/decimal"
)

type Abi interface{}

type AbiHandle struct {
	int `json:""`
}

type FunctionHeader struct {
	Expire *int             `json:"expire,omitempty"`
	Time   *decimal.Decimal `json:"time,omitempty"`
	Pubkey *string          `json:"pubkey,omitempty"`
}

type CallSet struct {
	FunctionName string          `json:"function_name"`
	Header       *FunctionHeader `json:"header,omitempty"`
	Input        interface{}     `json:"input,omitempty"`
}

type DeploySet struct {
	Tvc         string      `json:"tvc"`
	WorkchainID *int        `json:"workchain_id,omitempty"`
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
	Abi                Abi     `json:"abi"`
	CallSet            CallSet `json:"call_set"`
	IsInternal         bool    `json:"is_internal"`
	Signer             Signer  `json:"signer"`
	ProcessingTryIndex *int    `json:"processing_try_index,omitempty"`
}

type ResultOfEncodeMessageBody struct {
	Body       string  `json:"body"`
	DataToSign *string `json:"data_to_sign,omitempty"`
}

type ParamsOfAttachSignatureToMessageBody struct {
	Abi       Abi    `json:"abi"`
	PublicKey string `json:"public_key"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type ResultOfAttachSignatureToMessageBody struct {
	Body string `json:"body"`
}

type ParamsOfEncodeMessage struct {
	Abi                Abi        `json:"abi"`
	Address            *string    `json:"address,omitempty"`
	DeploySet          *DeploySet `json:"deploy_set,omitempty"`
	CallSet            *CallSet   `json:"call_set,omitempty"`
	Signer             Signer     `json:"signer"`
	ProcessingTryIndex *int       `json:"processing_try_index,omitempty"`
}

type ResultOfEncodeMessage struct {
	Message    string  `json:"message"`
	DataToSign *string `json:"data_to_sign,omitempty"`
	Address    string  `json:"address"`
	MessageID  string  `json:"message_id"`
}

type ParamsOfAttachSignature struct {
	Abi       Abi    `json:"abi"`
	PublicKey string `json:"public_key"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type ResultOfAttachSignature struct {
	Message   string `json:"message"`
	MessageID string `json:"message_id"`
}

type ParamsOfDecodeMessage struct {
	Abi     Abi    `json:"abi"`
	Message string `json:"message"`
}

type DecodedMessageBody struct {
	BodyType MessageBodyType `json:"body_type"`
	Name     string          `json:"name"`
	Value    interface{}     `json:"value,omitempty"`
	Header   *FunctionHeader `json:"header,omitempty"`
}

type ParamsOfDecodeMessageBody struct {
	Abi        Abi    `json:"abi"`
	Body       string `json:"body"`
	IsInternal bool   `json:"is_internal"`
}

type ParamsOfEncodeAccount struct {
	StateInit   StateInitSource  `json:"state_init"`
	Balance     *decimal.Decimal `json:"balance,omitempty"`
	LastTransLt *decimal.Decimal `json:"last_trans_lt,omitempty"`
	LastPaid    *int             `json:"last_paid,omitempty"`
}

type ResultOfEncodeAccount struct {
	Account string `json:"account"`
	ID      string `json:"id"`
}
