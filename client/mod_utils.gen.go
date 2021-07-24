package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 22 Jul 21 08:39 UTC
//
// Mod utils
//
// Misc utility Functions.

import (
	"encoding/json"
	"fmt"

	"github.com/volatiletech/null"
)

type AccountIDAddressStringFormat struct{}

type HexAddressStringFormat struct{}

type Base64AddressStringFormat struct {
	URL    bool `json:"url"`
	Test   bool `json:"test"`
	Bounce bool `json:"bounce"`
}

type AddressStringFormat struct {
	// Should be any of
	// AccountIDAddressStringFormat
	// HexAddressStringFormat
	// Base64AddressStringFormat
	EnumTypeValue interface{}
}

// MarshalJSON implements custom marshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *AddressStringFormat) MarshalJSON() ([]byte, error) { // nolint funlen
	switch value := (p.EnumTypeValue).(type) {
	case AccountIDAddressStringFormat:
		return json.Marshal(struct {
			AccountIDAddressStringFormat
			Type string `json:"type"`
		}{
			value,
			"AccountId",
		})

	case HexAddressStringFormat:
		return json.Marshal(struct {
			HexAddressStringFormat
			Type string `json:"type"`
		}{
			value,
			"Hex",
		})

	case Base64AddressStringFormat:
		return json.Marshal(struct {
			Base64AddressStringFormat
			Type string `json:"type"`
		}{
			value,
			"Base64",
		})

	default:
		return nil, fmt.Errorf("unsupported type for AddressStringFormat %v", p.EnumTypeValue)
	}
}

// UnmarshalJSON implements custom unmarshalling for rust
// directive #[serde(tag="type")] for enum of types.
func (p *AddressStringFormat) UnmarshalJSON(b []byte) error { // nolint funlen
	var typeDescriptor EnumOfTypesDescriptor
	if err := json.Unmarshal(b, &typeDescriptor); err != nil {
		return err
	}
	switch typeDescriptor.Type {
	case "AccountId":
		var enumTypeValue AccountIDAddressStringFormat
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Hex":
		var enumTypeValue HexAddressStringFormat
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	case "Base64":
		var enumTypeValue Base64AddressStringFormat
		if err := json.Unmarshal(b, &enumTypeValue); err != nil {
			return err
		}
		p.EnumTypeValue = enumTypeValue

	default:
		return fmt.Errorf("unsupported type for AddressStringFormat %v", typeDescriptor.Type)
	}

	return nil
}

type AccountAddressType string

const (
	AccountIDAccountAddressType AccountAddressType = "AccountId"
	HexAccountAddressType       AccountAddressType = "Hex"
	Base64AccountAddressType    AccountAddressType = "Base64"
)

type ParamsOfConvertAddress struct {
	// Account address in any TON format.
	Address string `json:"address"`
	// Specify the format to convert to.
	OutputFormat AddressStringFormat `json:"output_format"`
}

type ResultOfConvertAddress struct {
	// Address in the specified format.
	Address string `json:"address"`
}

type ParamsOfGetAddressType struct {
	// Account address in any TON format.
	Address string `json:"address"`
}

type ResultOfGetAddressType struct {
	// Account address type.
	AddressType AccountAddressType `json:"address_type"`
}

type ParamsOfCalcStorageFee struct {
	Account string `json:"account"`
	Period  uint32 `json:"period"`
}

type ResultOfCalcStorageFee struct {
	Fee string `json:"fee"`
}

type ParamsOfCompressZstd struct {
	// Uncompressed data.
	// Must be encoded as base64.
	Uncompressed string `json:"uncompressed"`
	// Compression level, from 1 to 21. Where: 1 - lowest compression level (fastest compression); 21 - highest compression level (slowest compression). If level is omitted, the default compression level is used (currently `3`).
	Level null.Int32 `json:"level"` // optional
}

type ResultOfCompressZstd struct {
	// Compressed data.
	// Must be encoded as base64.
	Compressed string `json:"compressed"`
}

type ParamsOfDecompressZstd struct {
	// Compressed data.
	// Must be encoded as base64.
	Compressed string `json:"compressed"`
}

type ResultOfDecompressZstd struct {
	// Decompressed data.
	// Must be encoded as base64.
	Decompressed string `json:"decompressed"`
}

// Converts address from any TON format to any TON format.
func (c *Client) UtilsConvertAddress(p *ParamsOfConvertAddress) (*ResultOfConvertAddress, error) {
	result := new(ResultOfConvertAddress)

	err := c.dllClient.waitErrorOrResultUnmarshal("utils.convert_address", p, result)

	return result, err
}

// Validates and returns the type of any TON address.
// Address types are the following
//
// `0:919db8e740d50bf349df2eea03fa30c385d846b991ff5542e67098ee833fc7f7` - standart TON address most
// commonly used in all cases. Also called as hex address
// `919db8e740d50bf349df2eea03fa30c385d846b991ff5542e67098ee833fc7f7` - account ID. A part of full
// address. Identifies account inside particular workchain
// `EQCRnbjnQNUL80nfLuoD+jDDhdhGuZH/VULmcJjugz/H9wam` - base64 address. Also called "user-friendly".
// Was used at the beginning of TON. Now it is supported for compatibility.
func (c *Client) UtilsGetAddressType(p *ParamsOfGetAddressType) (*ResultOfGetAddressType, error) {
	result := new(ResultOfGetAddressType)

	err := c.dllClient.waitErrorOrResultUnmarshal("utils.get_address_type", p, result)

	return result, err
}

// Calculates storage fee for an account over a specified time period.
func (c *Client) UtilsCalcStorageFee(p *ParamsOfCalcStorageFee) (*ResultOfCalcStorageFee, error) {
	result := new(ResultOfCalcStorageFee)

	err := c.dllClient.waitErrorOrResultUnmarshal("utils.calc_storage_fee", p, result)

	return result, err
}

// Compresses data using Zstandard algorithm.
func (c *Client) UtilsCompressZstd(p *ParamsOfCompressZstd) (*ResultOfCompressZstd, error) {
	result := new(ResultOfCompressZstd)

	err := c.dllClient.waitErrorOrResultUnmarshal("utils.compress_zstd", p, result)

	return result, err
}

// Decompresses data using Zstandard algorithm.
func (c *Client) UtilsDecompressZstd(p *ParamsOfDecompressZstd) (*ResultOfDecompressZstd, error) {
	result := new(ResultOfDecompressZstd)

	err := c.dllClient.waitErrorOrResultUnmarshal("utils.decompress_zstd", p, result)

	return result, err
}
