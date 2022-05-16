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

var ignoredTypesByName = map[string]bool{
	"ResultOfGetApiReference": true, // implemented via spec from this package
	"Abi":                     true,
}

var ignoredFunctionsByName = map[string]bool{
	"client.get_api_reference":        true,
	"processing.send_message":         true,
	"processing.wait_for_transaction": true,
	"processing.process_message":      true,
	"net.subscribe_collection":        true,
	"net.subscribe":                   true,
}
