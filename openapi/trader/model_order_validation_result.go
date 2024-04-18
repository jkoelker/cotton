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

// checks if the OrderValidationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderValidationResult{}

// OrderValidationResult struct for OrderValidationResult
type OrderValidationResult struct {
	Alerts []OrderValidationDetail `json:"alerts,omitempty"`
	Accepts []OrderValidationDetail `json:"accepts,omitempty"`
	Rejects []OrderValidationDetail `json:"rejects,omitempty"`
	Reviews []OrderValidationDetail `json:"reviews,omitempty"`
	Warns []OrderValidationDetail `json:"warns,omitempty"`
}

// NewOrderValidationResult instantiates a new OrderValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderValidationResult() *OrderValidationResult {
	this := OrderValidationResult{}
	return &this
}

// NewOrderValidationResultWithDefaults instantiates a new OrderValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderValidationResultWithDefaults() *OrderValidationResult {
	this := OrderValidationResult{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *OrderValidationResult) GetAlerts() []OrderValidationDetail {
	if o == nil || IsNil(o.Alerts) {
		var ret []OrderValidationDetail
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderValidationResult) GetAlertsOk() ([]OrderValidationDetail, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *OrderValidationResult) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []OrderValidationDetail and assigns it to the Alerts field.
func (o *OrderValidationResult) SetAlerts(v []OrderValidationDetail) {
	o.Alerts = v
}

// GetAccepts returns the Accepts field value if set, zero value otherwise.
func (o *OrderValidationResult) GetAccepts() []OrderValidationDetail {
	if o == nil || IsNil(o.Accepts) {
		var ret []OrderValidationDetail
		return ret
	}
	return o.Accepts
}

// GetAcceptsOk returns a tuple with the Accepts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderValidationResult) GetAcceptsOk() ([]OrderValidationDetail, bool) {
	if o == nil || IsNil(o.Accepts) {
		return nil, false
	}
	return o.Accepts, true
}

// HasAccepts returns a boolean if a field has been set.
func (o *OrderValidationResult) HasAccepts() bool {
	if o != nil && !IsNil(o.Accepts) {
		return true
	}

	return false
}

// SetAccepts gets a reference to the given []OrderValidationDetail and assigns it to the Accepts field.
func (o *OrderValidationResult) SetAccepts(v []OrderValidationDetail) {
	o.Accepts = v
}

// GetRejects returns the Rejects field value if set, zero value otherwise.
func (o *OrderValidationResult) GetRejects() []OrderValidationDetail {
	if o == nil || IsNil(o.Rejects) {
		var ret []OrderValidationDetail
		return ret
	}
	return o.Rejects
}

// GetRejectsOk returns a tuple with the Rejects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderValidationResult) GetRejectsOk() ([]OrderValidationDetail, bool) {
	if o == nil || IsNil(o.Rejects) {
		return nil, false
	}
	return o.Rejects, true
}

// HasRejects returns a boolean if a field has been set.
func (o *OrderValidationResult) HasRejects() bool {
	if o != nil && !IsNil(o.Rejects) {
		return true
	}

	return false
}

// SetRejects gets a reference to the given []OrderValidationDetail and assigns it to the Rejects field.
func (o *OrderValidationResult) SetRejects(v []OrderValidationDetail) {
	o.Rejects = v
}

// GetReviews returns the Reviews field value if set, zero value otherwise.
func (o *OrderValidationResult) GetReviews() []OrderValidationDetail {
	if o == nil || IsNil(o.Reviews) {
		var ret []OrderValidationDetail
		return ret
	}
	return o.Reviews
}

// GetReviewsOk returns a tuple with the Reviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderValidationResult) GetReviewsOk() ([]OrderValidationDetail, bool) {
	if o == nil || IsNil(o.Reviews) {
		return nil, false
	}
	return o.Reviews, true
}

// HasReviews returns a boolean if a field has been set.
func (o *OrderValidationResult) HasReviews() bool {
	if o != nil && !IsNil(o.Reviews) {
		return true
	}

	return false
}

// SetReviews gets a reference to the given []OrderValidationDetail and assigns it to the Reviews field.
func (o *OrderValidationResult) SetReviews(v []OrderValidationDetail) {
	o.Reviews = v
}

// GetWarns returns the Warns field value if set, zero value otherwise.
func (o *OrderValidationResult) GetWarns() []OrderValidationDetail {
	if o == nil || IsNil(o.Warns) {
		var ret []OrderValidationDetail
		return ret
	}
	return o.Warns
}

// GetWarnsOk returns a tuple with the Warns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderValidationResult) GetWarnsOk() ([]OrderValidationDetail, bool) {
	if o == nil || IsNil(o.Warns) {
		return nil, false
	}
	return o.Warns, true
}

// HasWarns returns a boolean if a field has been set.
func (o *OrderValidationResult) HasWarns() bool {
	if o != nil && !IsNil(o.Warns) {
		return true
	}

	return false
}

// SetWarns gets a reference to the given []OrderValidationDetail and assigns it to the Warns field.
func (o *OrderValidationResult) SetWarns(v []OrderValidationDetail) {
	o.Warns = v
}

func (o OrderValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderValidationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.Accepts) {
		toSerialize["accepts"] = o.Accepts
	}
	if !IsNil(o.Rejects) {
		toSerialize["rejects"] = o.Rejects
	}
	if !IsNil(o.Reviews) {
		toSerialize["reviews"] = o.Reviews
	}
	if !IsNil(o.Warns) {
		toSerialize["warns"] = o.Warns
	}
	return toSerialize, nil
}

type NullableOrderValidationResult struct {
	value *OrderValidationResult
	isSet bool
}

func (v NullableOrderValidationResult) Get() *OrderValidationResult {
	return v.value
}

func (v *NullableOrderValidationResult) Set(val *OrderValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderValidationResult(val *OrderValidationResult) *NullableOrderValidationResult {
	return &NullableOrderValidationResult{value: val, isSet: true}
}

func (v NullableOrderValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


