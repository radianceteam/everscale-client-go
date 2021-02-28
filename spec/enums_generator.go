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
	if err := enumOfConstsTpl.Execute(&tpl, t); err != nil {
		panic(err)
	}

	return tpl.String()
}

// genEnumOfTypes - generates enum of types with special enumOfConsts helper type.
func genEnumOfTypes(m Module, t Type) string {
	var tpl bytes.Buffer
	for i, et := range t.EnumTypes {
		if et.Type != Struct && et.Type != Ref {
			panic("EnumOfTypes only supports structs " + et.Name)
		}
		if et.Type == Struct {
			typeName := toGoName(strcase.ToSnake(et.Name + t.Name))
			t.EnumTypes[i].GoType = typeName
			et.Name = typeName
			tpl.WriteString("\n" + genStruct(m, et))
		} else {
			t.EnumTypes[i].GoType = toTypeName(et.RefName)
		}
	}

	if err := enumOfTypesTpl.Execute(&tpl, t); err != nil {
		panic(err)
	}

	return t.ToComment() + tpl.String()
}
