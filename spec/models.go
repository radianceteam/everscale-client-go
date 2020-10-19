package spec

import "strings"

type Description struct {
	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	GoComment   string // populated manually
	ConstName   string // name for enum consts
}

func (d *Description) ToComment() string {
	r := ""
	if !strings.Contains(d.Description, d.Summary) {
		r += "// " + prepareMultiline(d.Summary)
		if !strings.HasSuffix(r, ".") {
			r += "."
		}
		r += "\n"
	}

	if d.Description != "" {
		r += "// " + prepareMultiline(d.Description)
		if !strings.HasSuffix(r, ".") {
			r += "."
		}
		r += "\n"
	}

	return r
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
	Version string   `json:"version"`
	Modules []Module `json:"modules"`
}
