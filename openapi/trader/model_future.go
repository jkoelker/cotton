/*
Trader API - Account Access and User Preferences

Schwab Trader API access to Account, Order entry and User Preferences

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trader

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Future type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Future{}

// Future struct for Future
type Future struct {
	TransactionInstrument
	ActiveContract *bool `json:"activeContract,omitempty"`
	Type *string `json:"type,omitempty"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	LastTradingDate *time.Time `json:"lastTradingDate,omitempty"`
	FirstNoticeDate *time.Time `json:"firstNoticeDate,omitempty"`
	Multiplier *float64 `json:"multiplier,omitempty"`
	AssetType string `json:"assetType"`
	Cusip *string `json:"cusip,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Description *string `json:"description,omitempty"`
	InstrumentId *int64 `json:"instrumentId,omitempty"`
	NetChange *float64 `json:"netChange,omitempty"`
}

type _Future Future

// NewFuture instantiates a new Future object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuture(assetType string) *Future {
	this := Future{}
	this.AssetType = assetType
	var activeContract bool = false
	this.ActiveContract = &activeContract
	return &this
}

// NewFutureWithDefaults instantiates a new Future object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFutureWithDefaults() *Future {
	this := Future{}
	var activeContract bool = false
	this.ActiveContract = &activeContract
	return &this
}

// GetActiveContract returns the ActiveContract field value if set, zero value otherwise.
func (o *Future) GetActiveContract() bool {
	if o == nil || IsNil(o.ActiveContract) {
		var ret bool
		return ret
	}
	return *o.ActiveContract
}

// GetActiveContractOk returns a tuple with the ActiveContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetActiveContractOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveContract) {
		return nil, false
	}
	return o.ActiveContract, true
}

// HasActiveContract returns a boolean if a field has been set.
func (o *Future) HasActiveContract() bool {
	if o != nil && !IsNil(o.ActiveContract) {
		return true
	}

	return false
}

// SetActiveContract gets a reference to the given bool and assigns it to the ActiveContract field.
func (o *Future) SetActiveContract(v bool) {
	o.ActiveContract = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Future) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Future) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Future) SetType(v string) {
	o.Type = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *Future) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *Future) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *Future) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetLastTradingDate returns the LastTradingDate field value if set, zero value otherwise.
func (o *Future) GetLastTradingDate() time.Time {
	if o == nil || IsNil(o.LastTradingDate) {
		var ret time.Time
		return ret
	}
	return *o.LastTradingDate
}

// GetLastTradingDateOk returns a tuple with the LastTradingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetLastTradingDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastTradingDate) {
		return nil, false
	}
	return o.LastTradingDate, true
}

// HasLastTradingDate returns a boolean if a field has been set.
func (o *Future) HasLastTradingDate() bool {
	if o != nil && !IsNil(o.LastTradingDate) {
		return true
	}

	return false
}

// SetLastTradingDate gets a reference to the given time.Time and assigns it to the LastTradingDate field.
func (o *Future) SetLastTradingDate(v time.Time) {
	o.LastTradingDate = &v
}

// GetFirstNoticeDate returns the FirstNoticeDate field value if set, zero value otherwise.
func (o *Future) GetFirstNoticeDate() time.Time {
	if o == nil || IsNil(o.FirstNoticeDate) {
		var ret time.Time
		return ret
	}
	return *o.FirstNoticeDate
}

// GetFirstNoticeDateOk returns a tuple with the FirstNoticeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetFirstNoticeDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FirstNoticeDate) {
		return nil, false
	}
	return o.FirstNoticeDate, true
}

// HasFirstNoticeDate returns a boolean if a field has been set.
func (o *Future) HasFirstNoticeDate() bool {
	if o != nil && !IsNil(o.FirstNoticeDate) {
		return true
	}

	return false
}

// SetFirstNoticeDate gets a reference to the given time.Time and assigns it to the FirstNoticeDate field.
func (o *Future) SetFirstNoticeDate(v time.Time) {
	o.FirstNoticeDate = &v
}

// GetMultiplier returns the Multiplier field value if set, zero value otherwise.
func (o *Future) GetMultiplier() float64 {
	if o == nil || IsNil(o.Multiplier) {
		var ret float64
		return ret
	}
	return *o.Multiplier
}

// GetMultiplierOk returns a tuple with the Multiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetMultiplierOk() (*float64, bool) {
	if o == nil || IsNil(o.Multiplier) {
		return nil, false
	}
	return o.Multiplier, true
}

// HasMultiplier returns a boolean if a field has been set.
func (o *Future) HasMultiplier() bool {
	if o != nil && !IsNil(o.Multiplier) {
		return true
	}

	return false
}

// SetMultiplier gets a reference to the given float64 and assigns it to the Multiplier field.
func (o *Future) SetMultiplier(v float64) {
	o.Multiplier = &v
}

// GetAssetType returns the AssetType field value
func (o *Future) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *Future) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *Future) SetAssetType(v string) {
	o.AssetType = v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *Future) GetCusip() string {
	if o == nil || IsNil(o.Cusip) {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetCusipOk() (*string, bool) {
	if o == nil || IsNil(o.Cusip) {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *Future) HasCusip() bool {
	if o != nil && !IsNil(o.Cusip) {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *Future) SetCusip(v string) {
	o.Cusip = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Future) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Future) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Future) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Future) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Future) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Future) SetDescription(v string) {
	o.Description = &v
}

// GetInstrumentId returns the InstrumentId field value if set, zero value otherwise.
func (o *Future) GetInstrumentId() int64 {
	if o == nil || IsNil(o.InstrumentId) {
		var ret int64
		return ret
	}
	return *o.InstrumentId
}

// GetInstrumentIdOk returns a tuple with the InstrumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetInstrumentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.InstrumentId) {
		return nil, false
	}
	return o.InstrumentId, true
}

// HasInstrumentId returns a boolean if a field has been set.
func (o *Future) HasInstrumentId() bool {
	if o != nil && !IsNil(o.InstrumentId) {
		return true
	}

	return false
}

// SetInstrumentId gets a reference to the given int64 and assigns it to the InstrumentId field.
func (o *Future) SetInstrumentId(v int64) {
	o.InstrumentId = &v
}

// GetNetChange returns the NetChange field value if set, zero value otherwise.
func (o *Future) GetNetChange() float64 {
	if o == nil || IsNil(o.NetChange) {
		var ret float64
		return ret
	}
	return *o.NetChange
}

// GetNetChangeOk returns a tuple with the NetChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Future) GetNetChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.NetChange) {
		return nil, false
	}
	return o.NetChange, true
}

// HasNetChange returns a boolean if a field has been set.
func (o *Future) HasNetChange() bool {
	if o != nil && !IsNil(o.NetChange) {
		return true
	}

	return false
}

// SetNetChange gets a reference to the given float64 and assigns it to the NetChange field.
func (o *Future) SetNetChange(v float64) {
	o.NetChange = &v
}

func (o Future) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Future) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTransactionInstrument, errTransactionInstrument := json.Marshal(o.TransactionInstrument)
	if errTransactionInstrument != nil {
		return map[string]interface{}{}, errTransactionInstrument
	}
	errTransactionInstrument = json.Unmarshal([]byte(serializedTransactionInstrument), &toSerialize)
	if errTransactionInstrument != nil {
		return map[string]interface{}{}, errTransactionInstrument
	}
	if !IsNil(o.ActiveContract) {
		toSerialize["activeContract"] = o.ActiveContract
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.LastTradingDate) {
		toSerialize["lastTradingDate"] = o.LastTradingDate
	}
	if !IsNil(o.FirstNoticeDate) {
		toSerialize["firstNoticeDate"] = o.FirstNoticeDate
	}
	if !IsNil(o.Multiplier) {
		toSerialize["multiplier"] = o.Multiplier
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

func (o *Future) UnmarshalJSON(data []byte) (err error) {
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

	varFuture := _Future{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFuture)

	if err != nil {
		return err
	}

	*o = Future(varFuture)

	return err
}

type NullableFuture struct {
	value *Future
	isSet bool
}

func (v NullableFuture) Get() *Future {
	return v.value
}

func (v *NullableFuture) Set(val *Future) {
	v.value = val
	v.isSet = true
}

func (v NullableFuture) IsSet() bool {
	return v.isSet
}

func (v *NullableFuture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuture(val *Future) *NullableFuture {
	return &NullableFuture{value: val, isSet: true}
}

func (v NullableFuture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


