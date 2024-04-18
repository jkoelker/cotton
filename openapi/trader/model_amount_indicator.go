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

// AmountIndicator the model 'AmountIndicator'
type AmountIndicator string

// List of amountIndicator
const (
	AMOUNTINDICATOR_DOLLARS AmountIndicator = "DOLLARS"
	AMOUNTINDICATOR_SHARES AmountIndicator = "SHARES"
	AMOUNTINDICATOR_ALL_SHARES AmountIndicator = "ALL_SHARES"
	AMOUNTINDICATOR_PERCENTAGE AmountIndicator = "PERCENTAGE"
	AMOUNTINDICATOR_UNKNOWN AmountIndicator = "UNKNOWN"
)

// All allowed values of AmountIndicator enum
var AllowedAmountIndicatorEnumValues = []AmountIndicator{
	"DOLLARS",
	"SHARES",
	"ALL_SHARES",
	"PERCENTAGE",
	"UNKNOWN",
}

func (v *AmountIndicator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AmountIndicator(value)
	for _, existing := range AllowedAmountIndicatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AmountIndicator", value)
}

// NewAmountIndicatorFromValue returns a pointer to a valid AmountIndicator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAmountIndicatorFromValue(v string) (*AmountIndicator, error) {
	ev := AmountIndicator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AmountIndicator: valid values are %v", v, AllowedAmountIndicatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AmountIndicator) IsValid() bool {
	for _, existing := range AllowedAmountIndicatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to amountIndicator value
func (v AmountIndicator) Ptr() *AmountIndicator {
	return &v
}

type NullableAmountIndicator struct {
	value *AmountIndicator
	isSet bool
}

func (v NullableAmountIndicator) Get() *AmountIndicator {
	return v.value
}

func (v *NullableAmountIndicator) Set(val *AmountIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableAmountIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableAmountIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmountIndicator(val *AmountIndicator) *NullableAmountIndicator {
	return &NullableAmountIndicator{value: val, isSet: true}
}

func (v NullableAmountIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmountIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

