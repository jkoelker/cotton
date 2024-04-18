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

// StopType the model 'StopType'
type StopType string

// List of stopType
const (
	STOPTYPE_STANDARD StopType = "STANDARD"
	STOPTYPE_BID StopType = "BID"
	STOPTYPE_ASK StopType = "ASK"
	STOPTYPE_LAST StopType = "LAST"
	STOPTYPE_MARK StopType = "MARK"
)

// All allowed values of StopType enum
var AllowedStopTypeEnumValues = []StopType{
	"STANDARD",
	"BID",
	"ASK",
	"LAST",
	"MARK",
}

func (v *StopType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StopType(value)
	for _, existing := range AllowedStopTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StopType", value)
}

// NewStopTypeFromValue returns a pointer to a valid StopType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStopTypeFromValue(v string) (*StopType, error) {
	ev := StopType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StopType: valid values are %v", v, AllowedStopTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StopType) IsValid() bool {
	for _, existing := range AllowedStopTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stopType value
func (v StopType) Ptr() *StopType {
	return &v
}

type NullableStopType struct {
	value *StopType
	isSet bool
}

func (v NullableStopType) Get() *StopType {
	return v.value
}

func (v *NullableStopType) Set(val *StopType) {
	v.value = val
	v.isSet = true
}

func (v NullableStopType) IsSet() bool {
	return v.isSet
}

func (v *NullableStopType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopType(val *StopType) *NullableStopType {
	return &NullableStopType{value: val, isSet: true}
}

func (v NullableStopType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
