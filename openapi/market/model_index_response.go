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

// checks if the IndexResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexResponse{}

// IndexResponse Quote info of Index security
type IndexResponse struct {
	AssetMainType *AssetMainType `json:"assetMainType,omitempty"`
	// SSID of instrument
	Ssid *int64 `json:"ssid,omitempty"`
	// Symbol of instrument
	Symbol *string `json:"symbol,omitempty"`
	// is quote realtime
	Realtime *bool `json:"realtime,omitempty"`
	Quote *QuoteIndex `json:"quote,omitempty"`
	Reference *ReferenceIndex `json:"reference,omitempty"`
}

// NewIndexResponse instantiates a new IndexResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexResponse() *IndexResponse {
	this := IndexResponse{}
	return &this
}

// NewIndexResponseWithDefaults instantiates a new IndexResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexResponseWithDefaults() *IndexResponse {
	this := IndexResponse{}
	return &this
}

// GetAssetMainType returns the AssetMainType field value if set, zero value otherwise.
func (o *IndexResponse) GetAssetMainType() AssetMainType {
	if o == nil || IsNil(o.AssetMainType) {
		var ret AssetMainType
		return ret
	}
	return *o.AssetMainType
}

// GetAssetMainTypeOk returns a tuple with the AssetMainType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResponse) GetAssetMainTypeOk() (*AssetMainType, bool) {
	if o == nil || IsNil(o.AssetMainType) {
		return nil, false
	}
	return o.AssetMainType, true
}

// HasAssetMainType returns a boolean if a field has been set.
func (o *IndexResponse) HasAssetMainType() bool {
	if o != nil && !IsNil(o.AssetMainType) {
		return true
	}

	return false
}

// SetAssetMainType gets a reference to the given AssetMainType and assigns it to the AssetMainType field.
func (o *IndexResponse) SetAssetMainType(v AssetMainType) {
	o.AssetMainType = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *IndexResponse) GetSsid() int64 {
	if o == nil || IsNil(o.Ssid) {
		var ret int64
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResponse) GetSsidOk() (*int64, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *IndexResponse) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int64 and assigns it to the Ssid field.
func (o *IndexResponse) SetSsid(v int64) {
	o.Ssid = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *IndexResponse) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResponse) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *IndexResponse) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *IndexResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetRealtime returns the Realtime field value if set, zero value otherwise.
func (o *IndexResponse) GetRealtime() bool {
	if o == nil || IsNil(o.Realtime) {
		var ret bool
		return ret
	}
	return *o.Realtime
}

// GetRealtimeOk returns a tuple with the Realtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResponse) GetRealtimeOk() (*bool, bool) {
	if o == nil || IsNil(o.Realtime) {
		return nil, false
	}
	return o.Realtime, true
}

// HasRealtime returns a boolean if a field has been set.
func (o *IndexResponse) HasRealtime() bool {
	if o != nil && !IsNil(o.Realtime) {
		return true
	}

	return false
}

// SetRealtime gets a reference to the given bool and assigns it to the Realtime field.
func (o *IndexResponse) SetRealtime(v bool) {
	o.Realtime = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *IndexResponse) GetQuote() QuoteIndex {
	if o == nil || IsNil(o.Quote) {
		var ret QuoteIndex
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResponse) GetQuoteOk() (*QuoteIndex, bool) {
	if o == nil || IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *IndexResponse) HasQuote() bool {
	if o != nil && !IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given QuoteIndex and assigns it to the Quote field.
func (o *IndexResponse) SetQuote(v QuoteIndex) {
	o.Quote = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *IndexResponse) GetReference() ReferenceIndex {
	if o == nil || IsNil(o.Reference) {
		var ret ReferenceIndex
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexResponse) GetReferenceOk() (*ReferenceIndex, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *IndexResponse) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given ReferenceIndex and assigns it to the Reference field.
func (o *IndexResponse) SetReference(v ReferenceIndex) {
	o.Reference = &v
}

func (o IndexResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexResponse) ToMap() (map[string]interface{}, error) {
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

type NullableIndexResponse struct {
	value *IndexResponse
	isSet bool
}

func (v NullableIndexResponse) Get() *IndexResponse {
	return v.value
}

func (v *NullableIndexResponse) Set(val *IndexResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexResponse(val *IndexResponse) *NullableIndexResponse {
	return &NullableIndexResponse{value: val, isSet: true}
}

func (v NullableIndexResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

