package client

// DON'T EDIT THIS FILE! It is generated via 'task generate' at 28 Feb 21 17:43 UTC
//
// Mod utils
//
// Misc utility Functions.

import (
	"encoding/json"
	"fmt"
)

type AccountIDAddressStringFormat struct{}

type HexAddressStringFormat struct{}

type Base64AddressStringFormat struct {
	URL    bool `json:"url"`
	Test   bool `json:"test"`
	Bounce bool `json:"bounce"`
}

type AddressStringFormat struct {
	EnumTypeValue interface{} // any of AccountIDAddressStringFormat, HexAddressStringFormat, Base64AddressStringFormat,
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

// Converts address from any TON format to any TON format.
func (c *Client) UtilsConvertAddress(p *ParamsOfConvertAddress) (*ResultOfConvertAddress, error) {
	result := new(ResultOfConvertAddress)

	err := c.dllClient.waitErrorOrResultUnmarshal("utils.convert_address", p, result)

	return result, err
}
