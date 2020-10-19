package spec

type TypeName string

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
	EnumOfType   TypeName = "EnumOfType"
	EnumOfConsts TypeName = "EnymOfConsts"
)
