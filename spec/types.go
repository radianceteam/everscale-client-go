package spec

type TypeName string

// TypeName different values.
const (
	Ref          TypeName = "Ref"
	Optional     TypeName = "Optional"
	String       TypeName = "String"
	Number       TypeName = "Number"
	Value        TypeName = "Value"
	Struct       TypeName = "Struct"
	None         TypeName = "None"
	BigInt       TypeName = "BigInt"
	Boolean      TypeName = "Boolean"
	Array        TypeName = "Array"
	Generic      TypeName = "Generic"
	EnumOfTypes  TypeName = "EnumOfTypes"
	EnumOfConsts TypeName = "EnumOfConsts"
)

type NumberType string

const (
	NTFloat NumberType = "Float"
	NTInt   NumberType = "Int"
	NTUint  NumberType = "UInt"
)

var ignoredTypesByName = map[string]bool{
	"ResultOfGetApiReference": true,
	"ClientError":             true,
	"ClientConfig":            true,
	"Abi":                     true,
	"AbiHandle":               true,
	"SigningBoxHandle":        true,
	"MessageSource":           true,
}

var ignoredFunctionsByName = map[string]bool{
	"get_api_reference":    true,
	"send_message":         true,
	"wait_for_transaction": true,
	"process_message":      true,
	"subscribe_collection": true,
}
