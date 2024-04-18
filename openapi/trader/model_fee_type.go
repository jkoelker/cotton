/*
Trader API - Account Access and User Preferences

Schwab Trader API access to Account, Order entry and User Preferences

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trader

import (
	"encoding/json"
	"fmt"
)

// FeeType the model 'FeeType'
type FeeType string

// List of FeeType
const (
	FEETYPE_COMMISSION FeeType = "COMMISSION"
	FEETYPE_SEC_FEE FeeType = "SEC_FEE"
	FEETYPE_STR_FEE FeeType = "STR_FEE"
	FEETYPE_R_FEE FeeType = "R_FEE"
	FEETYPE_CDSC_FEE FeeType = "CDSC_FEE"
	FEETYPE_OPT_REG_FEE FeeType = "OPT_REG_FEE"
	FEETYPE_ADDITIONAL_FEE FeeType = "ADDITIONAL_FEE"
	FEETYPE_MISCELLANEOUS_FEE FeeType = "MISCELLANEOUS_FEE"
	FEETYPE_FTT FeeType = "FTT"
	FEETYPE_FUTURES_CLEARING_FEE FeeType = "FUTURES_CLEARING_FEE"
	FEETYPE_FUTURES_DESK_OFFICE_FEE FeeType = "FUTURES_DESK_OFFICE_FEE"
	FEETYPE_FUTURES_EXCHANGE_FEE FeeType = "FUTURES_EXCHANGE_FEE"
	FEETYPE_FUTURES_GLOBEX_FEE FeeType = "FUTURES_GLOBEX_FEE"
	FEETYPE_FUTURES_NFA_FEE FeeType = "FUTURES_NFA_FEE"
	FEETYPE_FUTURES_PIT_BROKERAGE_FEE FeeType = "FUTURES_PIT_BROKERAGE_FEE"
	FEETYPE_FUTURES_TRANSACTION_FEE FeeType = "FUTURES_TRANSACTION_FEE"
	FEETYPE_LOW_PROCEEDS_COMMISSION FeeType = "LOW_PROCEEDS_COMMISSION"
	FEETYPE_BASE_CHARGE FeeType = "BASE_CHARGE"
	FEETYPE_GENERAL_CHARGE FeeType = "GENERAL_CHARGE"
	FEETYPE_GST_FEE FeeType = "GST_FEE"
	FEETYPE_TAF_FEE FeeType = "TAF_FEE"
	FEETYPE_INDEX_OPTION_FEE FeeType = "INDEX_OPTION_FEE"
	FEETYPE_TEFRA_TAX FeeType = "TEFRA_TAX"
	FEETYPE_STATE_TAX FeeType = "STATE_TAX"
	FEETYPE_UNKNOWN FeeType = "UNKNOWN"
)

// All allowed values of FeeType enum
var AllowedFeeTypeEnumValues = []FeeType{
	"COMMISSION",
	"SEC_FEE",
	"STR_FEE",
	"R_FEE",
	"CDSC_FEE",
	"OPT_REG_FEE",
	"ADDITIONAL_FEE",
	"MISCELLANEOUS_FEE",
	"FTT",
	"FUTURES_CLEARING_FEE",
	"FUTURES_DESK_OFFICE_FEE",
	"FUTURES_EXCHANGE_FEE",
	"FUTURES_GLOBEX_FEE",
	"FUTURES_NFA_FEE",
	"FUTURES_PIT_BROKERAGE_FEE",
	"FUTURES_TRANSACTION_FEE",
	"LOW_PROCEEDS_COMMISSION",
	"BASE_CHARGE",
	"GENERAL_CHARGE",
	"GST_FEE",
	"TAF_FEE",
	"INDEX_OPTION_FEE",
	"TEFRA_TAX",
	"STATE_TAX",
	"UNKNOWN",
}

func (v *FeeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeeType(value)
	for _, existing := range AllowedFeeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeeType", value)
}

// NewFeeTypeFromValue returns a pointer to a valid FeeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeeTypeFromValue(v string) (*FeeType, error) {
	ev := FeeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeeType: valid values are %v", v, AllowedFeeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeeType) IsValid() bool {
	for _, existing := range AllowedFeeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeeType value
func (v FeeType) Ptr() *FeeType {
	return &v
}

type NullableFeeType struct {
	value *FeeType
	isSet bool
}

func (v NullableFeeType) Get() *FeeType {
	return v.value
}

func (v *NullableFeeType) Set(val *FeeType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeType(val *FeeType) *NullableFeeType {
	return &NullableFeeType{value: val, isSet: true}
}

func (v NullableFeeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

