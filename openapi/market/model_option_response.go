/*
Market Data

Trader API - Market data

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package market

import (
	"encoding/json"
)

// checks if the OptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionResponse{}

// OptionResponse Quote info of Option security
type OptionResponse struct {
	AssetMainType *AssetMainType `json:"assetMainType,omitempty"`
	// SSID of instrument
	Ssid *int64 `json:"ssid,omitempty"`
	// Symbol of instrument
	Symbol *string `json:"symbol,omitempty"`
	// is quote realtime
	Realtime *bool `json:"realtime,omitempty"`
	Quote *QuoteOption `json:"quote,omitempty"`
	Reference *ReferenceOption `json:"reference,omitempty"`
}

// NewOptionResponse instantiates a new OptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionResponse() *OptionResponse {
	this := OptionResponse{}
	return &this
}

// NewOptionResponseWithDefaults instantiates a new OptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionResponseWithDefaults() *OptionResponse {
	this := OptionResponse{}
	return &this
}

// GetAssetMainType returns the AssetMainType field value if set, zero value otherwise.
func (o *OptionResponse) GetAssetMainType() AssetMainType {
	if o == nil || IsNil(o.AssetMainType) {
		var ret AssetMainType
		return ret
	}
	return *o.AssetMainType
}

// GetAssetMainTypeOk returns a tuple with the AssetMainType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionResponse) GetAssetMainTypeOk() (*AssetMainType, bool) {
	if o == nil || IsNil(o.AssetMainType) {
		return nil, false
	}
	return o.AssetMainType, true
}

// HasAssetMainType returns a boolean if a field has been set.
func (o *OptionResponse) HasAssetMainType() bool {
	if o != nil && !IsNil(o.AssetMainType) {
		return true
	}

	return false
}

// SetAssetMainType gets a reference to the given AssetMainType and assigns it to the AssetMainType field.
func (o *OptionResponse) SetAssetMainType(v AssetMainType) {
	o.AssetMainType = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *OptionResponse) GetSsid() int64 {
	if o == nil || IsNil(o.Ssid) {
		var ret int64
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionResponse) GetSsidOk() (*int64, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *OptionResponse) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int64 and assigns it to the Ssid field.
func (o *OptionResponse) SetSsid(v int64) {
	o.Ssid = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OptionResponse) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionResponse) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OptionResponse) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OptionResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetRealtime returns the Realtime field value if set, zero value otherwise.
func (o *OptionResponse) GetRealtime() bool {
	if o == nil || IsNil(o.Realtime) {
		var ret bool
		return ret
	}
	return *o.Realtime
}

// GetRealtimeOk returns a tuple with the Realtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionResponse) GetRealtimeOk() (*bool, bool) {
	if o == nil || IsNil(o.Realtime) {
		return nil, false
	}
	return o.Realtime, true
}

// HasRealtime returns a boolean if a field has been set.
func (o *OptionResponse) HasRealtime() bool {
	if o != nil && !IsNil(o.Realtime) {
		return true
	}

	return false
}

// SetRealtime gets a reference to the given bool and assigns it to the Realtime field.
func (o *OptionResponse) SetRealtime(v bool) {
	o.Realtime = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *OptionResponse) GetQuote() QuoteOption {
	if o == nil || IsNil(o.Quote) {
		var ret QuoteOption
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionResponse) GetQuoteOk() (*QuoteOption, bool) {
	if o == nil || IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *OptionResponse) HasQuote() bool {
	if o != nil && !IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given QuoteOption and assigns it to the Quote field.
func (o *OptionResponse) SetQuote(v QuoteOption) {
	o.Quote = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *OptionResponse) GetReference() ReferenceOption {
	if o == nil || IsNil(o.Reference) {
		var ret ReferenceOption
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionResponse) GetReferenceOk() (*ReferenceOption, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *OptionResponse) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given ReferenceOption and assigns it to the Reference field.
func (o *OptionResponse) SetReference(v ReferenceOption) {
	o.Reference = &v
}

func (o OptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetMainType) {
		toSerialize["assetMainType"] = o.AssetMainType
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Realtime) {
		toSerialize["realtime"] = o.Realtime
	}
	if !IsNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullableOptionResponse struct {
	value *OptionResponse
	isSet bool
}

func (v NullableOptionResponse) Get() *OptionResponse {
	return v.value
}

func (v *NullableOptionResponse) Set(val *OptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionResponse(val *OptionResponse) *NullableOptionResponse {
	return &NullableOptionResponse{value: val, isSet: true}
}

func (v NullableOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


