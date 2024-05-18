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

// checks if the EquityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquityResponse{}

// EquityResponse Quote info of Equity security
type EquityResponse struct {
	AssetMainType *AssetMainType `json:"assetMainType,omitempty"`
	AssetSubType NullableEquityAssetSubType `json:"assetSubType,omitempty"`
	// SSID of instrument
	Ssid *int64 `json:"ssid,omitempty"`
	// Symbol of instrument
	Symbol *string `json:"symbol,omitempty"`
	// is quote realtime
	Realtime *bool `json:"realtime,omitempty"`
	QuoteType NullableQuoteType `json:"quoteType,omitempty"`
	Extended *ExtendedMarket `json:"extended,omitempty"`
	Fundamental *Fundamental `json:"fundamental,omitempty"`
	Quote *QuoteEquity `json:"quote,omitempty"`
	Reference *ReferenceEquity `json:"reference,omitempty"`
	Regular *RegularMarket `json:"regular,omitempty"`
}

// NewEquityResponse instantiates a new EquityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityResponse() *EquityResponse {
	this := EquityResponse{}
	return &this
}

// NewEquityResponseWithDefaults instantiates a new EquityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityResponseWithDefaults() *EquityResponse {
	this := EquityResponse{}
	return &this
}

// GetAssetMainType returns the AssetMainType field value if set, zero value otherwise.
func (o *EquityResponse) GetAssetMainType() AssetMainType {
	if o == nil || IsNil(o.AssetMainType) {
		var ret AssetMainType
		return ret
	}
	return *o.AssetMainType
}

// GetAssetMainTypeOk returns a tuple with the AssetMainType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityResponse) GetAssetMainTypeOk() (*AssetMainType, bool) {
	if o == nil || IsNil(o.AssetMainType) {
		return nil, false
	}
	return o.AssetMainType, true
}

// HasAssetMainType returns a boolean if a field has been set.
func (o *EquityResponse) HasAssetMainType() bool {
	if o != nil && !IsNil(o.AssetMainType) {
		return true
	}

	return false
}

// SetAssetMainType gets a reference to the given AssetMainType and assigns it to the AssetMainType field.
func (o *EquityResponse) SetAssetMainType(v AssetMainType) {
	o.AssetMainType = &v
}

// GetAssetSubType returns the AssetSubType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityResponse) GetAssetSubType() EquityAssetSubType {
	if o == nil || IsNil(o.AssetSubType.Get()) {
		var ret EquityAssetSubType
		return ret
	}
	return *o.AssetSubType.Get()
}

// GetAssetSubTypeOk returns a tuple with the AssetSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityResponse) GetAssetSubTypeOk() (*EquityAssetSubType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetSubType.Get(), o.AssetSubType.IsSet()
}

// HasAssetSubType returns a boolean if a field has been set.
func (o *EquityResponse) HasAssetSubType() bool {
	if o != nil && o.AssetSubType.IsSet() {
		return true
	}

	return false
}

// SetAssetSubType gets a reference to the given NullableEquityAssetSubType and assigns it to the AssetSubType field.
func (o *EquityResponse) SetAssetSubType(v EquityAssetSubType) {
	o.AssetSubType.Set(&v)
}
// SetAssetSubTypeNil sets the value for AssetSubType to be an explicit nil
func (o *EquityResponse) SetAssetSubTypeNil() {
	o.AssetSubType.Set(nil)
}

// UnsetAssetSubType ensures that no value is present for AssetSubType, not even an explicit nil
func (o *EquityResponse) UnsetAssetSubType() {
	o.AssetSubType.Unset()
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *EquityResponse) GetSsid() int64 {
	if o == nil || IsNil(o.Ssid) {
		var ret int64
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityResponse) GetSsidOk() (*int64, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *EquityResponse) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int64 and assigns it to the Ssid field.
func (o *EquityResponse) SetSsid(v int64) {
	o.Ssid = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *EquityResponse) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityResponse) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *EquityResponse) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *EquityResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetRealtime returns the Realtime field value if set, zero value otherwise.
func (o *EquityResponse) GetRealtime() bool {
	if o == nil || IsNil(o.Realtime) {
		var ret bool
		return ret
	}
	return *o.Realtime
}

// GetRealtimeOk returns a tuple with the Realtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityResponse) GetRealtimeOk() (*bool, bool) {
	if o == nil || IsNil(o.Realtime) {
		return nil, false
	}
	return o.Realtime, true
}

// HasRealtime returns a boolean if a field has been set.
func (o *EquityResponse) HasRealtime() bool {
	if o != nil && !IsNil(o.Realtime) {
		return true
	}

	return false
}

// SetRealtime gets a reference to the given bool and assigns it to the Realtime field.
func (o *EquityResponse) SetRealtime(v bool) {
	o.Realtime = &v
}

// GetQuoteType returns the QuoteType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquityResponse) GetQuoteType() QuoteType {
	if o == nil || IsNil(o.QuoteType.Get()) {
		var ret QuoteType
		return ret
	}
	return *o.QuoteType.Get()
}

// GetQuoteTypeOk returns a tuple with the QuoteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquityResponse) GetQuoteTypeOk() (*QuoteType, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuoteType.Get(), o.QuoteType.IsSet()
}

// HasQuoteType returns a boolean if a field has been set.
func (o *EquityResponse) HasQuoteType() bool {
	if o != nil && o.QuoteType.IsSet() {
		return true
	}

	return false
}

// SetQuoteType gets a reference to the given NullableQuoteType and assigns it to the QuoteType field.
func (o *EquityResponse) SetQuoteType(v QuoteType) {
	o.QuoteType.Set(&v)
}
// SetQuoteTypeNil sets the value for QuoteType to be an explicit nil
func (o *EquityResponse) SetQuoteTypeNil() {
	o.QuoteType.Set(nil)
}

// UnsetQuoteType ensures that no value is present for QuoteType, not even an explicit nil
func (o *EquityResponse) UnsetQuoteType() {
	o.QuoteType.Unset()
}

// GetExtended returns the Extended field value if set, zero value otherwise.
func (o *EquityResponse) GetExtended() ExtendedMarket {
	if o == nil || IsNil(o.Extended) {
		var ret ExtendedMarket
		return ret
	}
	return *o.Extended
}

// GetExtendedOk returns a tuple with the Extended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityResponse) GetExtendedOk() (*ExtendedMarket, bool) {
	if o == nil || IsNil(o.Extended) {
		return nil, false
	}
	return o.Extended, true
}

// HasExtended returns a boolean if a field has been set.
func (o *EquityResponse) HasExtended() bool {
	if o != nil && !IsNil(o.Extended) {
		return true
	}

	return false
}

// SetExtended gets a reference to the given ExtendedMarket and assigns it to the Extended field.
func (o *EquityResponse) SetExtended(v ExtendedMarket) {
	o.Extended = &v
}

// GetFundamental returns the Fundamental field value if set, zero value otherwise.
func (o *EquityResponse) GetFundamental() Fundamental {
	if o == nil || IsNil(o.Fundamental) {
		var ret Fundamental
		return ret
	}
	return *o.Fundamental
}

// GetFundamentalOk returns a tuple with the Fundamental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityResponse) GetFundamentalOk() (*Fundamental, bool) {
	if o == nil || IsNil(o.Fundamental) {
		return nil, false
	}
	return o.Fundamental, true
}

// HasFundamental returns a boolean if a field has been set.
func (o *EquityResponse) HasFundamental() bool {
	if o != nil && !IsNil(o.Fundamental) {
		return true
	}

	return false
}

// SetFundamental gets a reference to the given Fundamental and assigns it to the Fundamental field.
func (o *EquityResponse) SetFundamental(v Fundamental) {
	o.Fundamental = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *EquityResponse) GetQuote() QuoteEquity {
	if o == nil || IsNil(o.Quote) {
		var ret QuoteEquity
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityResponse) GetQuoteOk() (*QuoteEquity, bool) {
	if o == nil || IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *EquityResponse) HasQuote() bool {
	if o != nil && !IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given QuoteEquity and assigns it to the Quote field.
func (o *EquityResponse) SetQuote(v QuoteEquity) {
	o.Quote = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *EquityResponse) GetReference() ReferenceEquity {
	if o == nil || IsNil(o.Reference) {
		var ret ReferenceEquity
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityResponse) GetReferenceOk() (*ReferenceEquity, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *EquityResponse) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given ReferenceEquity and assigns it to the Reference field.
func (o *EquityResponse) SetReference(v ReferenceEquity) {
	o.Reference = &v
}

// GetRegular returns the Regular field value if set, zero value otherwise.
func (o *EquityResponse) GetRegular() RegularMarket {
	if o == nil || IsNil(o.Regular) {
		var ret RegularMarket
		return ret
	}
	return *o.Regular
}

// GetRegularOk returns a tuple with the Regular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityResponse) GetRegularOk() (*RegularMarket, bool) {
	if o == nil || IsNil(o.Regular) {
		return nil, false
	}
	return o.Regular, true
}

// HasRegular returns a boolean if a field has been set.
func (o *EquityResponse) HasRegular() bool {
	if o != nil && !IsNil(o.Regular) {
		return true
	}

	return false
}

// SetRegular gets a reference to the given RegularMarket and assigns it to the Regular field.
func (o *EquityResponse) SetRegular(v RegularMarket) {
	o.Regular = &v
}

func (o EquityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetMainType) {
		toSerialize["assetMainType"] = o.AssetMainType
	}
	if o.AssetSubType.IsSet() {
		toSerialize["assetSubType"] = o.AssetSubType.Get()
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
	if o.QuoteType.IsSet() {
		toSerialize["quoteType"] = o.QuoteType.Get()
	}
	if !IsNil(o.Extended) {
		toSerialize["extended"] = o.Extended
	}
	if !IsNil(o.Fundamental) {
		toSerialize["fundamental"] = o.Fundamental
	}
	if !IsNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.Regular) {
		toSerialize["regular"] = o.Regular
	}
	return toSerialize, nil
}

type NullableEquityResponse struct {
	value *EquityResponse
	isSet bool
}

func (v NullableEquityResponse) Get() *EquityResponse {
	return v.value
}

func (v *NullableEquityResponse) Set(val *EquityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityResponse(val *EquityResponse) *NullableEquityResponse {
	return &NullableEquityResponse{value: val, isSet: true}
}

func (v NullableEquityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

