package spec

import "strings"

type Description struct {
	Name        string `json:"name,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
	ConstName   string `json:"-"` // populated manually, used in templates only
	GoComment   string `json:"-"` // populated manually, used in templates only
}

// ToComment generates comment for generate code from spec.
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

// Type description from JSON spec.
type Type struct {
	Description
	Type          TypeName      `json:"type,omitempty"`
	RefName       string        `json:"ref_name,omitempty"`
	GenericName   string        `json:"generic_name,omitempty"`
	NumberType    string        `json:"number_type,omitempty"`
	NumberSize    uint          `json:"number_size,omitempty"`
	GenericArgs   []Type        `json:"generic_args,omitempty"`
	OptionalInner *Type         `json:"optional_inner,omitempty"`
	ArrayItem     *Type         `json:"array_item,omitempty"`
	StructFields  []Type        `json:"struct_fields,omitempty"`
	EnumConsts    []Description `json:"enum_consts,omitempty"`
	EnumTypes     []Type        `json:"enum_types"`
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
