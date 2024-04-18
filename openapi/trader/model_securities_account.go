/*
Trader API - Account Access and User Preferences

Schwab Trader API access to Account, Order entry and User Preferences

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trader

import (
	"encoding/json"
)

// checks if the SecuritiesAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecuritiesAccount{}

// SecuritiesAccount struct for SecuritiesAccount
type SecuritiesAccount struct {
	Type *string `json:"type,omitempty"`
	AccountNumber *string `json:"accountNumber,omitempty"`
	RoundTrips *int32 `json:"roundTrips,omitempty"`
	IsDayTrader *bool `json:"isDayTrader,omitempty"`
	IsClosingOnlyRestricted *bool `json:"isClosingOnlyRestricted,omitempty"`
	PfcbFlag *bool `json:"pfcbFlag,omitempty"`
	Positions []Position `json:"positions,omitempty"`
}

// NewSecuritiesAccount instantiates a new SecuritiesAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecuritiesAccount() *SecuritiesAccount {
	this := SecuritiesAccount{}
	var isDayTrader bool = false
	this.IsDayTrader = &isDayTrader
	var isClosingOnlyRestricted bool = false
	this.IsClosingOnlyRestricted = &isClosingOnlyRestricted
	var pfcbFlag bool = false
	this.PfcbFlag = &pfcbFlag
	return &this
}

// NewSecuritiesAccountWithDefaults instantiates a new SecuritiesAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecuritiesAccountWithDefaults() *SecuritiesAccount {
	this := SecuritiesAccount{}
	var isDayTrader bool = false
	this.IsDayTrader = &isDayTrader
	var isClosingOnlyRestricted bool = false
	this.IsClosingOnlyRestricted = &isClosingOnlyRestricted
	var pfcbFlag bool = false
	this.PfcbFlag = &pfcbFlag
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecuritiesAccount) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccount) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecuritiesAccount) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecuritiesAccount) SetType(v string) {
	o.Type = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *SecuritiesAccount) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccount) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *SecuritiesAccount) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *SecuritiesAccount) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetRoundTrips returns the RoundTrips field value if set, zero value otherwise.
func (o *SecuritiesAccount) GetRoundTrips() int32 {
	if o == nil || IsNil(o.RoundTrips) {
		var ret int32
		return ret
	}
	return *o.RoundTrips
}

// GetRoundTripsOk returns a tuple with the RoundTrips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccount) GetRoundTripsOk() (*int32, bool) {
	if o == nil || IsNil(o.RoundTrips) {
		return nil, false
	}
	return o.RoundTrips, true
}

// HasRoundTrips returns a boolean if a field has been set.
func (o *SecuritiesAccount) HasRoundTrips() bool {
	if o != nil && !IsNil(o.RoundTrips) {
		return true
	}

	return false
}

// SetRoundTrips gets a reference to the given int32 and assigns it to the RoundTrips field.
func (o *SecuritiesAccount) SetRoundTrips(v int32) {
	o.RoundTrips = &v
}

// GetIsDayTrader returns the IsDayTrader field value if set, zero value otherwise.
func (o *SecuritiesAccount) GetIsDayTrader() bool {
	if o == nil || IsNil(o.IsDayTrader) {
		var ret bool
		return ret
	}
	return *o.IsDayTrader
}

// GetIsDayTraderOk returns a tuple with the IsDayTrader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccount) GetIsDayTraderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDayTrader) {
		return nil, false
	}
	return o.IsDayTrader, true
}

// HasIsDayTrader returns a boolean if a field has been set.
func (o *SecuritiesAccount) HasIsDayTrader() bool {
	if o != nil && !IsNil(o.IsDayTrader) {
		return true
	}

	return false
}

// SetIsDayTrader gets a reference to the given bool and assigns it to the IsDayTrader field.
func (o *SecuritiesAccount) SetIsDayTrader(v bool) {
	o.IsDayTrader = &v
}

// GetIsClosingOnlyRestricted returns the IsClosingOnlyRestricted field value if set, zero value otherwise.
func (o *SecuritiesAccount) GetIsClosingOnlyRestricted() bool {
	if o == nil || IsNil(o.IsClosingOnlyRestricted) {
		var ret bool
		return ret
	}
	return *o.IsClosingOnlyRestricted
}

// GetIsClosingOnlyRestrictedOk returns a tuple with the IsClosingOnlyRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccount) GetIsClosingOnlyRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsClosingOnlyRestricted) {
		return nil, false
	}
	return o.IsClosingOnlyRestricted, true
}

// HasIsClosingOnlyRestricted returns a boolean if a field has been set.
func (o *SecuritiesAccount) HasIsClosingOnlyRestricted() bool {
	if o != nil && !IsNil(o.IsClosingOnlyRestricted) {
		return true
	}

	return false
}

// SetIsClosingOnlyRestricted gets a reference to the given bool and assigns it to the IsClosingOnlyRestricted field.
func (o *SecuritiesAccount) SetIsClosingOnlyRestricted(v bool) {
	o.IsClosingOnlyRestricted = &v
}

// GetPfcbFlag returns the PfcbFlag field value if set, zero value otherwise.
func (o *SecuritiesAccount) GetPfcbFlag() bool {
	if o == nil || IsNil(o.PfcbFlag) {
		var ret bool
		return ret
	}
	return *o.PfcbFlag
}

// GetPfcbFlagOk returns a tuple with the PfcbFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccount) GetPfcbFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.PfcbFlag) {
		return nil, false
	}
	return o.PfcbFlag, true
}

// HasPfcbFlag returns a boolean if a field has been set.
func (o *SecuritiesAccount) HasPfcbFlag() bool {
	if o != nil && !IsNil(o.PfcbFlag) {
		return true
	}

	return false
}

// SetPfcbFlag gets a reference to the given bool and assigns it to the PfcbFlag field.
func (o *SecuritiesAccount) SetPfcbFlag(v bool) {
	o.PfcbFlag = &v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *SecuritiesAccount) GetPositions() []Position {
	if o == nil || IsNil(o.Positions) {
		var ret []Position
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecuritiesAccount) GetPositionsOk() ([]Position, bool) {
	if o == nil || IsNil(o.Positions) {
		return nil, false
	}
	return o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *SecuritiesAccount) HasPositions() bool {
	if o != nil && !IsNil(o.Positions) {
		return true
	}

	return false
}

// SetPositions gets a reference to the given []Position and assigns it to the Positions field.
func (o *SecuritiesAccount) SetPositions(v []Position) {
	o.Positions = v
}

func (o SecuritiesAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecuritiesAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.RoundTrips) {
		toSerialize["roundTrips"] = o.RoundTrips
	}
	if !IsNil(o.IsDayTrader) {
		toSerialize["isDayTrader"] = o.IsDayTrader
	}
	if !IsNil(o.IsClosingOnlyRestricted) {
		toSerialize["isClosingOnlyRestricted"] = o.IsClosingOnlyRestricted
	}
	if !IsNil(o.PfcbFlag) {
		toSerialize["pfcbFlag"] = o.PfcbFlag
	}
	if !IsNil(o.Positions) {
		toSerialize["positions"] = o.Positions
	}
	return toSerialize, nil
}

type NullableSecuritiesAccount struct {
	value *SecuritiesAccount
	isSet bool
}

func (v NullableSecuritiesAccount) Get() *SecuritiesAccount {
	return v.value
}

func (v *NullableSecuritiesAccount) Set(val *SecuritiesAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSecuritiesAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSecuritiesAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecuritiesAccount(val *SecuritiesAccount) *NullableSecuritiesAccount {
	return &NullableSecuritiesAccount{value: val, isSet: true}
}

func (v NullableSecuritiesAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecuritiesAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


