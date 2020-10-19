package spec

type Description struct {
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	ConstName   string // name for enum consts
}

type Type struct {
	Description
	Type          TypeName      `json:"type"`
	RefName       string        `json:"ref_name"`
	GenericName   string        `json:"generic_name"`
	GenericArgs   []Type        `json:"generic_args"`
	OptionalInner *Type         `json:"optional_inner"`
	ArrayItem     *Type         `json:"array_item"`
	StructFields  []Type        `json:"struct_fields"`
	EnumConsts    []Description `json:"enum_consts"`
}

type Function struct {
	Description
	Params []Type `json:"params"`
	Result Type   `json:"result"`
}

type Module struct {
	Description
	Types     []Type     `json:"types"`
	Functions []Function `json:"functions"`
}

type APIReference struct {
	API struct {
		Version string   `json:"version"`
		Modules []Module `json:"modules"`
	} `json:"api"`
}
