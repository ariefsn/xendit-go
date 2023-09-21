/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.5.0
*/


package invoice

import (
	"encoding/json"
	
	"fmt"
)

// EwalletType Representing the available eWallet channels used for invoice-related transactions.
type EwalletType string

// List of EwalletType
const (
	EWALLETTYPE_OVO EwalletType = "OVO"
	EWALLETTYPE_DANA EwalletType = "DANA"
	EWALLETTYPE_LINKAJA EwalletType = "LINKAJA"
	EWALLETTYPE_PAYMAYA EwalletType = "PAYMAYA"
	EWALLETTYPE_SHOPEEPAY EwalletType = "SHOPEEPAY"
	EWALLETTYPE_GCASH EwalletType = "GCASH"
	EWALLETTYPE_GRABPAY EwalletType = "GRABPAY"
)

// All allowed values of EwalletType enum
var AllowedEwalletTypeEnumValues = []EwalletType{
	"OVO",
	"DANA",
	"LINKAJA",
	"PAYMAYA",
	"SHOPEEPAY",
	"GCASH",
	"GRABPAY",
}

func (v *EwalletType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EwalletType(value)
	for _, existing := range AllowedEwalletTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EwalletType", value)
}

// NewEwalletTypeFromValue returns a pointer to a valid EwalletType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEwalletTypeFromValue(v string) (*EwalletType, error) {
	ev := EwalletType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EwalletType: valid values are %v", v, AllowedEwalletTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EwalletType) IsValid() bool {
	for _, existing := range AllowedEwalletTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

func (v EwalletType) String() string {
	return string(v)
}

// Ptr returns reference to EwalletType value
func (v EwalletType) Ptr() *EwalletType {
	return &v
}

type NullableEwalletType struct {
	value *EwalletType
	isSet bool
}

func (v NullableEwalletType) Get() *EwalletType {
	return v.value
}

func (v *NullableEwalletType) Set(val *EwalletType) {
	v.value = val
	v.isSet = true
}

func (v NullableEwalletType) IsSet() bool {
	return v.isSet
}

func (v *NullableEwalletType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEwalletType(val *EwalletType) *NullableEwalletType {
	return &NullableEwalletType{value: val, isSet: true}
}

func (v NullableEwalletType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEwalletType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
