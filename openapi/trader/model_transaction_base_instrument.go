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

// checks if the TransactionBaseInstrument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionBaseInstrument{}

// TransactionBaseInstrument struct for TransactionBaseInstrument
type TransactionBaseInstrument struct {
	AssetType string `json:"assetType"`
	Cusip *string `json:"cusip,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Description *string `json:"description,omitempty"`
	InstrumentId *int64 `json:"instrumentId,omitempty"`
	NetChange *float64 `json:"netChange,omitempty"`
}

type _TransactionBaseInstrument TransactionBaseInstrument

// NewTransactionBaseInstrument instantiates a new TransactionBaseInstrument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionBaseInstrument(assetType string) *TransactionBaseInstrument {
	this := TransactionBaseInstrument{}
	this.AssetType = assetType
	return &this
}

// NewTransactionBaseInstrumentWithDefaults instantiates a new TransactionBaseInstrument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionBaseInstrumentWithDefaults() *TransactionBaseInstrument {
	this := TransactionBaseInstrument{}
	return &this
}

// GetAssetType returns the AssetType field value
func (o *TransactionBaseInstrument) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *TransactionBaseInstrument) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *TransactionBaseInstrument) SetAssetType(v string) {
	o.AssetType = v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *TransactionBaseInstrument) GetCusip() string {
	if o == nil || IsNil(o.Cusip) {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionBaseInstrument) GetCusipOk() (*string, bool) {
	if o == nil || IsNil(o.Cusip) {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *TransactionBaseInstrument) HasCusip() bool {
	if o != nil && !IsNil(o.Cusip) {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *TransactionBaseInstrument) SetCusip(v string) {
	o.Cusip = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TransactionBaseInstrument) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionBaseInstrument) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TransactionBaseInstrument) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TransactionBaseInstrument) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransactionBaseInstrument) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionBaseInstrument) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransactionBaseInstrument) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransactionBaseInstrument) SetDescription(v string) {
	o.Description = &v
}

// GetInstrumentId returns the InstrumentId field value if set, zero value otherwise.
func (o *TransactionBaseInstrument) GetInstrumentId() int64 {
	if o == nil || IsNil(o.InstrumentId) {
		var ret int64
		return ret
	}
	return *o.InstrumentId
}

// GetInstrumentIdOk returns a tuple with the InstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionBaseInstrument) GetInstrumentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.InstrumentId) {
		return nil, false
	}
	return o.InstrumentId, true
}

// HasInstrumentId returns a boolean if a field has been set.
func (o *TransactionBaseInstrument) HasInstrumentId() bool {
	if o != nil && !IsNil(o.InstrumentId) {
		return true
	}

	return false
}

// SetInstrumentId gets a reference to the given int64 and assigns it to the InstrumentId field.
func (o *TransactionBaseInstrument) SetInstrumentId(v int64) {
	o.InstrumentId = &v
}

// GetNetChange returns the NetChange field value if set, zero value otherwise.
func (o *TransactionBaseInstrument) GetNetChange() float64 {
	if o == nil || IsNil(o.NetChange) {
		var ret float64
		return ret
	}
	return *o.NetChange
}

// GetNetChangeOk returns a tuple with the NetChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionBaseInstrument) GetNetChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.NetChange) {
		return nil, false
	}
	return o.NetChange, true
}

// HasNetChange returns a boolean if a field has been set.
func (o *TransactionBaseInstrument) HasNetChange() bool {
	if o != nil && !IsNil(o.NetChange) {
		return true
	}

	return false
}

// SetNetChange gets a reference to the given float64 and assigns it to the NetChange field.
func (o *TransactionBaseInstrument) SetNetChange(v float64) {
	o.NetChange = &v
}

func (o TransactionBaseInstrument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionBaseInstrument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *TransactionBaseInstrument) UnmarshalJSON(data []byte) (err error) {
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

	varTransactionBaseInstrument := _TransactionBaseInstrument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionBaseInstrument)

	if err != nil {
		return err
	}

	*o = TransactionBaseInstrument(varTransactionBaseInstrument)

	return err
}

type NullableTransactionBaseInstrument struct {
	value *TransactionBaseInstrument
	isSet bool
}

func (v NullableTransactionBaseInstrument) Get() *TransactionBaseInstrument {
	return v.value
}

func (v *NullableTransactionBaseInstrument) Set(val *TransactionBaseInstrument) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionBaseInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionBaseInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionBaseInstrument(val *TransactionBaseInstrument) *NullableTransactionBaseInstrument {
	return &NullableTransactionBaseInstrument{value: val, isSet: true}
}

func (v NullableTransactionBaseInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionBaseInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


