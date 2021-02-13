package spec

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

func genEnumOfConsts(t Type) string {
	t.Name = withTypeAlias(t.Name)
	for i := range t.EnumConsts {
		if t.EnumConsts[i].Type == None {
			t.GoType = "string"
			t.EnumConsts[i].Value = strconv.Quote(withTypeAlias(t.EnumConsts[i].Name))
		}

		t.EnumConsts[i].Name = withTypeAlias(t.EnumConsts[i].Name)
		t.EnumConsts[i].GoComment = t.EnumConsts[i].ToComment()
		t.EnumConsts[i].ConstName = toGoName(strcase.ToSnake(t.EnumConsts[i].ConstName)) + t.Name
	}
	if strings.HasSuffix(t.Name, "ErrorCode") {
		t.Name = "" // constants without type for error-codes
	}
	var tpl bytes.Buffer
	if err := enumTmpl.Execute(&tpl, t); err != nil {
		panic(err)
	}

	return tpl.String()
}

// genEnumOfTypes - generates enum of types with special enumOfConsts helper type.
func genEnumOfTypes(m Module, t Type) string {
	enumHelperType := Type{Type: EnumOfConsts, GoType: "string", Description: Description{Name: t.Name + "Type"}}
	fields := make(map[string]int)
	structFields := make([]Type, 1)
	structFields[0] = Type{Type: Ref, RefName: enumHelperType.Name, Description: Description{Name: "type"}}
	fields["type"] = 0
	for _, et := range t.EnumTypes {
		enumHelperType.EnumConsts = append(enumHelperType.EnumConsts, Type{
			Description: Description{
				Name:        et.Name,
				Summary:     et.Summary,
				Description: et.Description.Description,
				ConstName:   et.Name,
			},
			Value: strconv.Quote(et.Name),
		})
		if et.Type != Struct {
			panic("EnumOfTypes only supports structs " + et.Name)
		}
		for _, sf := range et.StructFields {
			index, ok := fields[sf.Name]
			if ok && structFields[index].Type != sf.Type {
				panic("type mismatch for enum " + et.Name + ":" + sf.Name)
			}

			if !ok {
				index = len(structFields)
				fields[sf.Name] = index
				structFields = append(structFields, sf)
				structFields[index].Description.Description += " presented in types:"
			}
			structFields[index].Description.Description += "\n\"" + et.Name + "\""
		}
	}
	t.StructFields = structFields
	return genEnumOfConsts(enumHelperType) + "\n" + genStruct(m, t)
}
