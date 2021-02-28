package client

// Manually implemented methods for mod_abi.

type AbiType string

const (
	ContractAbiType AbiType = "Contract"
	HandleAbiType   AbiType = "Handle"
	JSONAbiType     AbiType = "Json"
)

type Abi struct {
	Type  AbiType     `json:"type"`
	Value interface{} `json:"value"`
}
