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

// Duration the model 'Duration'
type Duration string

// List of duration
const (
	DURATION_DAY Duration = "DAY"
	DURATION_GOOD_TILL_CANCEL Duration = "GOOD_TILL_CANCEL"
	DURATION_FILL_OR_KILL Duration = "FILL_OR_KILL"
	DURATION_IMMEDIATE_OR_CANCEL Duration = "IMMEDIATE_OR_CANCEL"
	DURATION_END_OF_WEEK Duration = "END_OF_WEEK"
	DURATION_END_OF_MONTH Duration = "END_OF_MONTH"
	DURATION_NEXT_END_OF_MONTH Duration = "NEXT_END_OF_MONTH"
	DURATION_UNKNOWN Duration = "UNKNOWN"
)

// All allowed values of Duration enum
var AllowedDurationEnumValues = []Duration{
	"DAY",
	"GOOD_TILL_CANCEL",
	"FILL_OR_KILL",
	"IMMEDIATE_OR_CANCEL",
	"END_OF_WEEK",
	"END_OF_MONTH",
	"NEXT_END_OF_MONTH",
	"UNKNOWN",
}

func (v *Duration) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Duration(value)
	for _, existing := range AllowedDurationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Duration", value)
}

// NewDurationFromValue returns a pointer to a valid Duration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDurationFromValue(v string) (*Duration, error) {
	ev := Duration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Duration: valid values are %v", v, AllowedDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Duration) IsValid() bool {
	for _, existing := range AllowedDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to duration value
func (v Duration) Ptr() *Duration {
	return &v
}

type NullableDuration struct {
	value *Duration
	isSet bool
}

func (v NullableDuration) Get() *Duration {
	return v.value
}

func (v *NullableDuration) Set(val *Duration) {
	v.value = val
	v.isSet = true
}

func (v NullableDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuration(val *Duration) *NullableDuration {
	return &NullableDuration{value: val, isSet: true}
}

func (v NullableDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

