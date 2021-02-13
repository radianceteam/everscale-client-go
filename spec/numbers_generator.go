package spec

import (
	"fmt"
	"strconv"
)

var optionNumberToGoType = map[string]string{
	"Int8":    "null.Int8",
	"Int16":   "null.Int16",
	"Int32":   "null.Int32",
	"Int64":   "null.Int64",
	"UInt8":   "null.Uint8",
	"UInt16":  "null.Uint16",
	"UInt32":  "null.Uint32",
	"UInt64":  "null.Uint64",
	"Float32": "null.Float32",
}

var numberToGoType = map[string]string{
	"Int8":   "int8",
	"Int16":  "int16",
	"Int32":  "int32",
	"Int64":  "int64",
	"UInt8":  "uint8",
	"UInt16": "uint16",
	"UInt32": "uint32",
	"UInt64": "uint64",
}

func genNumber(t Type, isOptional bool) string {
	if isOptional {
		numberType, found := optionNumberToGoType[t.NumberType+strconv.Itoa(t.NumberSize)]
		if found {
			return numberType
		}
		panic(fmt.Sprintf("opt number type not found %s %d", t.NumberType, t.NumberSize))
	}
	numberType, found := numberToGoType[t.NumberType+strconv.Itoa(t.NumberSize)]
	if found {
		return numberType
	}
	panic(fmt.Sprintf("req number type not found %s %d", t.NumberType, t.NumberSize))
}
