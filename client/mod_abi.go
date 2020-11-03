package client

// Manually implemented methods for mod_abi.

type AbiType string

const (
	SerializedAbiType AbiType = "Serialized"
	HandleAbiType     AbiType = "Handle"
)

type Abi struct {
	Type  AbiType     `json:"type"`
	Value interface{} `json:"value"`
}

type SigningBoxHandle int

type MessageSourceType string

const (
	EncodedMessageSourceType        MessageSourceType = "Encoded"
	EncodingParamsMessageSourceType MessageSourceType = "EncodingParams"
)

type MessageSource struct {
	Type MessageSourceType `json:"type"`
	// presented in types:
	// "Encoded".
	Message string `json:"message"`
	// presented in types:
	// "EncodingParams", but abi field also for "Encoded"
	ParamsOfEncodeMessage
}
