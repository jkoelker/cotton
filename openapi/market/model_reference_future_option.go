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

// checks if the ReferenceFutureOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceFutureOption{}

// ReferenceFutureOption Reference data of Future Option security
type ReferenceFutureOption struct {
	ContractType *ContractType `json:"contractType,omitempty"`
	// Description of Instrument
	Description *string `json:"description,omitempty"`
	// Exchange Code
	Exchange *string `json:"exchange,omitempty"`
	// Exchange Name
	ExchangeName *string `json:"exchangeName,omitempty"`
	// Option multiplier
	Multiplier *float64 `json:"multiplier,omitempty"`
	// date of expiration in long
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// Style of expiration
	ExpirationStyle *string `json:"expirationStyle,omitempty"`
	// Strike Price
	StrikePrice *float64 `json:"strikePrice,omitempty"`
	// A company, index or fund name
	Underlying *string `json:"underlying,omitempty"`
}

// NewReferenceFutureOption instantiates a new ReferenceFutureOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceFutureOption() *ReferenceFutureOption {
	this := ReferenceFutureOption{}
	return &this
}

// NewReferenceFutureOptionWithDefaults instantiates a new ReferenceFutureOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceFutureOptionWithDefaults() *ReferenceFutureOption {
	this := ReferenceFutureOption{}
	return &this
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *ReferenceFutureOption) GetContractType() ContractType {
	if o == nil || IsNil(o.ContractType) {
		var ret ContractType
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceFutureOption) GetContractTypeOk() (*ContractType, bool) {
	if o == nil || IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *ReferenceFutureOption) HasContractType() bool {
	if o != nil && !IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given ContractType and assigns it to the ContractType field.
func (o *ReferenceFutureOption) SetContractType(v ContractType) {
	o.ContractType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReferenceFutureOption) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceFutureOption) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReferenceFutureOption) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReferenceFutureOption) SetDescription(v string) {
	o.Description = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *ReferenceFutureOption) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceFutureOption) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *ReferenceFutureOption) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *ReferenceFutureOption) SetExchange(v string) {
	o.Exchange = &v
}

// GetExchangeName returns the ExchangeName field value if set, zero value otherwise.
func (o *ReferenceFutureOption) GetExchangeName() string {
	if o == nil || IsNil(o.ExchangeName) {
		var ret string
		return ret
	}
	return *o.ExchangeName
}

// GetExchangeNameOk returns a tuple with the ExchangeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceFutureOption) GetExchangeNameOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeName) {
		return nil, false
	}
	return o.ExchangeName, true
}

// HasExchangeName returns a boolean if a field has been set.
func (o *ReferenceFutureOption) HasExchangeName() bool {
	if o != nil && !IsNil(o.ExchangeName) {
		return true
	}

	return false
}

// SetExchangeName gets a reference to the given string and assigns it to the ExchangeName field.
func (o *ReferenceFutureOption) SetExchangeName(v string) {
	o.ExchangeName = &v
}

// GetMultiplier returns the Multiplier field value if set, zero value otherwise.
func (o *ReferenceFutureOption) GetMultiplier() float64 {
	if o == nil || IsNil(o.Multiplier) {
		var ret float64
		return ret
	}
	return *o.Multiplier
}

// GetMultiplierOk returns a tuple with the Multiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceFutureOption) GetMultiplierOk() (*float64, bool) {
	if o == nil || IsNil(o.Multiplier) {
		return nil, false
	}
	return o.Multiplier, true
}

// HasMultiplier returns a boolean if a field has been set.
func (o *ReferenceFutureOption) HasMultiplier() bool {
	if o != nil && !IsNil(o.Multiplier) {
		return true
	}

	return false
}

// SetMultiplier gets a reference to the given float64 and assigns it to the Multiplier field.
func (o *ReferenceFutureOption) SetMultiplier(v float64) {
	o.Multiplier = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ReferenceFutureOption) GetExpirationDate() int64 {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceFutureOption) GetExpirationDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ReferenceFutureOption) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *ReferenceFutureOption) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetExpirationStyle returns the ExpirationStyle field value if set, zero value otherwise.
func (o *ReferenceFutureOption) GetExpirationStyle() string {
	if o == nil || IsNil(o.ExpirationStyle) {
		var ret string
		return ret
	}
	return *o.ExpirationStyle
}

// GetExpirationStyleOk returns a tuple with the ExpirationStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceFutureOption) GetExpirationStyleOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationStyle) {
		return nil, false
	}
	return o.ExpirationStyle, true
}

// HasExpirationStyle returns a boolean if a field has been set.
func (o *ReferenceFutureOption) HasExpirationStyle() bool {
	if o != nil && !IsNil(o.ExpirationStyle) {
		return true
	}

	return false
}

// SetExpirationStyle gets a reference to the given string and assigns it to the ExpirationStyle field.
func (o *ReferenceFutureOption) SetExpirationStyle(v string) {
	o.ExpirationStyle = &v
}

// GetStrikePrice returns the StrikePrice field value if set, zero value otherwise.
func (o *ReferenceFutureOption) GetStrikePrice() float64 {
	if o == nil || IsNil(o.StrikePrice) {
		var ret float64
		return ret
	}
	return *o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceFutureOption) GetStrikePriceOk() (*float64, bool) {
	if o == nil || IsNil(o.StrikePrice) {
		return nil, false
	}
	return o.StrikePrice, true
}

// HasStrikePrice returns a boolean if a field has been set.
func (o *ReferenceFutureOption) HasStrikePrice() bool {
	if o != nil && !IsNil(o.StrikePrice) {
		return true
	}

	return false
}

// SetStrikePrice gets a reference to the given float64 and assigns it to the StrikePrice field.
func (o *ReferenceFutureOption) SetStrikePrice(v float64) {
	o.StrikePrice = &v
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *ReferenceFutureOption) GetUnderlying() string {
	if o == nil || IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceFutureOption) GetUnderlyingOk() (*string, bool) {
	if o == nil || IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *ReferenceFutureOption) HasUnderlying() bool {
	if o != nil && !IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *ReferenceFutureOption) SetUnderlying(v string) {
	o.Underlying = &v
}

func (o ReferenceFutureOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceFutureOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.ExchangeName) {
		toSerialize["exchangeName"] = o.ExchangeName
	}
	if !IsNil(o.Multiplier) {
		toSerialize["multiplier"] = o.Multiplier
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.ExpirationStyle) {
		toSerialize["expirationStyle"] = o.ExpirationStyle
	}
	if !IsNil(o.StrikePrice) {
		toSerialize["strikePrice"] = o.StrikePrice
	}
	if !IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	return toSerialize, nil
}

type NullableReferenceFutureOption struct {
	value *ReferenceFutureOption
	isSet bool
}

func (v NullableReferenceFutureOption) Get() *ReferenceFutureOption {
	return v.value
}

func (v *NullableReferenceFutureOption) Set(val *ReferenceFutureOption) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceFutureOption) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceFutureOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceFutureOption(val *ReferenceFutureOption) *NullableReferenceFutureOption {
	return &NullableReferenceFutureOption{value: val, isSet: true}
}

func (v NullableReferenceFutureOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceFutureOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


