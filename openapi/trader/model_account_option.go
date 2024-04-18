/*
Trader API - Account Access and User Preferences

Schwab Trader API access to Account, Order entry and User Preferences

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trader

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AccountOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountOption{}

// AccountOption struct for AccountOption
type AccountOption struct {
	OptionDeliverables []AccountAPIOptionDeliverable `json:"optionDeliverables,omitempty"`
	PutCall *string `json:"putCall,omitempty"`
	OptionMultiplier *int32 `json:"optionMultiplier,omitempty"`
	Type *string `json:"type,omitempty"`
	UnderlyingSymbol *string `json:"underlyingSymbol,omitempty"`
	AssetType string `json:"assetType"`
	Cusip *string `json:"cusip,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Description *string `json:"description,omitempty"`
	InstrumentId *int64 `json:"instrumentId,omitempty"`
	NetChange *float64 `json:"netChange,omitempty"`
}

type _AccountOption AccountOption

// NewAccountOption instantiates a new AccountOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountOption(assetType string) *AccountOption {
	this := AccountOption{}
	this.AssetType = assetType
	return &this
}

// NewAccountOptionWithDefaults instantiates a new AccountOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountOptionWithDefaults() *AccountOption {
	this := AccountOption{}
	return &this
}

// GetOptionDeliverables returns the OptionDeliverables field value if set, zero value otherwise.
func (o *AccountOption) GetOptionDeliverables() []AccountAPIOptionDeliverable {
	if o == nil || IsNil(o.OptionDeliverables) {
		var ret []AccountAPIOptionDeliverable
		return ret
	}
	return o.OptionDeliverables
}

// GetOptionDeliverablesOk returns a tuple with the OptionDeliverables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetOptionDeliverablesOk() ([]AccountAPIOptionDeliverable, bool) {
	if o == nil || IsNil(o.OptionDeliverables) {
		return nil, false
	}
	return o.OptionDeliverables, true
}

// HasOptionDeliverables returns a boolean if a field has been set.
func (o *AccountOption) HasOptionDeliverables() bool {
	if o != nil && !IsNil(o.OptionDeliverables) {
		return true
	}

	return false
}

// SetOptionDeliverables gets a reference to the given []AccountAPIOptionDeliverable and assigns it to the OptionDeliverables field.
func (o *AccountOption) SetOptionDeliverables(v []AccountAPIOptionDeliverable) {
	o.OptionDeliverables = v
}

// GetPutCall returns the PutCall field value if set, zero value otherwise.
func (o *AccountOption) GetPutCall() string {
	if o == nil || IsNil(o.PutCall) {
		var ret string
		return ret
	}
	return *o.PutCall
}

// GetPutCallOk returns a tuple with the PutCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetPutCallOk() (*string, bool) {
	if o == nil || IsNil(o.PutCall) {
		return nil, false
	}
	return o.PutCall, true
}

// HasPutCall returns a boolean if a field has been set.
func (o *AccountOption) HasPutCall() bool {
	if o != nil && !IsNil(o.PutCall) {
		return true
	}

	return false
}

// SetPutCall gets a reference to the given string and assigns it to the PutCall field.
func (o *AccountOption) SetPutCall(v string) {
	o.PutCall = &v
}

// GetOptionMultiplier returns the OptionMultiplier field value if set, zero value otherwise.
func (o *AccountOption) GetOptionMultiplier() int32 {
	if o == nil || IsNil(o.OptionMultiplier) {
		var ret int32
		return ret
	}
	return *o.OptionMultiplier
}

// GetOptionMultiplierOk returns a tuple with the OptionMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetOptionMultiplierOk() (*int32, bool) {
	if o == nil || IsNil(o.OptionMultiplier) {
		return nil, false
	}
	return o.OptionMultiplier, true
}

// HasOptionMultiplier returns a boolean if a field has been set.
func (o *AccountOption) HasOptionMultiplier() bool {
	if o != nil && !IsNil(o.OptionMultiplier) {
		return true
	}

	return false
}

// SetOptionMultiplier gets a reference to the given int32 and assigns it to the OptionMultiplier field.
func (o *AccountOption) SetOptionMultiplier(v int32) {
	o.OptionMultiplier = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountOption) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountOption) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountOption) SetType(v string) {
	o.Type = &v
}

// GetUnderlyingSymbol returns the UnderlyingSymbol field value if set, zero value otherwise.
func (o *AccountOption) GetUnderlyingSymbol() string {
	if o == nil || IsNil(o.UnderlyingSymbol) {
		var ret string
		return ret
	}
	return *o.UnderlyingSymbol
}

// GetUnderlyingSymbolOk returns a tuple with the UnderlyingSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetUnderlyingSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.UnderlyingSymbol) {
		return nil, false
	}
	return o.UnderlyingSymbol, true
}

// HasUnderlyingSymbol returns a boolean if a field has been set.
func (o *AccountOption) HasUnderlyingSymbol() bool {
	if o != nil && !IsNil(o.UnderlyingSymbol) {
		return true
	}

	return false
}

// SetUnderlyingSymbol gets a reference to the given string and assigns it to the UnderlyingSymbol field.
func (o *AccountOption) SetUnderlyingSymbol(v string) {
	o.UnderlyingSymbol = &v
}

// GetAssetType returns the AssetType field value
func (o *AccountOption) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *AccountOption) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *AccountOption) SetAssetType(v string) {
	o.AssetType = v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *AccountOption) GetCusip() string {
	if o == nil || IsNil(o.Cusip) {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetCusipOk() (*string, bool) {
	if o == nil || IsNil(o.Cusip) {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *AccountOption) HasCusip() bool {
	if o != nil && !IsNil(o.Cusip) {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *AccountOption) SetCusip(v string) {
	o.Cusip = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AccountOption) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AccountOption) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AccountOption) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccountOption) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountOption) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccountOption) SetDescription(v string) {
	o.Description = &v
}

// GetInstrumentId returns the InstrumentId field value if set, zero value otherwise.
func (o *AccountOption) GetInstrumentId() int64 {
	if o == nil || IsNil(o.InstrumentId) {
		var ret int64
		return ret
	}
	return *o.InstrumentId
}

// GetInstrumentIdOk returns a tuple with the InstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetInstrumentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.InstrumentId) {
		return nil, false
	}
	return o.InstrumentId, true
}

// HasInstrumentId returns a boolean if a field has been set.
func (o *AccountOption) HasInstrumentId() bool {
	if o != nil && !IsNil(o.InstrumentId) {
		return true
	}

	return false
}

// SetInstrumentId gets a reference to the given int64 and assigns it to the InstrumentId field.
func (o *AccountOption) SetInstrumentId(v int64) {
	o.InstrumentId = &v
}

// GetNetChange returns the NetChange field value if set, zero value otherwise.
func (o *AccountOption) GetNetChange() float64 {
	if o == nil || IsNil(o.NetChange) {
		var ret float64
		return ret
	}
	return *o.NetChange
}

// GetNetChangeOk returns a tuple with the NetChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetNetChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.NetChange) {
		return nil, false
	}
	return o.NetChange, true
}

// HasNetChange returns a boolean if a field has been set.
func (o *AccountOption) HasNetChange() bool {
	if o != nil && !IsNil(o.NetChange) {
		return true
	}

	return false
}

// SetNetChange gets a reference to the given float64 and assigns it to the NetChange field.
func (o *AccountOption) SetNetChange(v float64) {
	o.NetChange = &v
}

func (o AccountOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionDeliverables) {
		toSerialize["optionDeliverables"] = o.OptionDeliverables
	}
	if !IsNil(o.PutCall) {
		toSerialize["putCall"] = o.PutCall
	}
	if !IsNil(o.OptionMultiplier) {
		toSerialize["optionMultiplier"] = o.OptionMultiplier
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UnderlyingSymbol) {
		toSerialize["underlyingSymbol"] = o.UnderlyingSymbol
	}
	toSerialize["assetType"] = o.AssetType
	if !IsNil(o.Cusip) {
		toSerialize["cusip"] = o.Cusip
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InstrumentId) {
		toSerialize["instrumentId"] = o.InstrumentId
	}
	if !IsNil(o.NetChange) {
		toSerialize["netChange"] = o.NetChange
	}
	return toSerialize, nil
}

func (o *AccountOption) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assetType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccountOption := _AccountOption{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountOption)

	if err != nil {
		return err
	}

	*o = AccountOption(varAccountOption)

	return err
}

type NullableAccountOption struct {
	value *AccountOption
	isSet bool
}

func (v NullableAccountOption) Get() *AccountOption {
	return v.value
}

func (v *NullableAccountOption) Set(val *AccountOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountOption(val *AccountOption) *NullableAccountOption {
	return &NullableAccountOption{value: val, isSet: true}
}

func (v NullableAccountOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


