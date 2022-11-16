/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Reenrollment struct for Reenrollment
type Reenrollment struct {
	IsFlushPolicyHistoryEnabled *bool `json:"isFlushPolicyHistoryEnabled,omitempty"`
	IsFlushLocationInformationEnabled *bool `json:"isFlushLocationInformationEnabled,omitempty"`
	IsFlushLocationInformationHistoryEnabled *bool `json:"isFlushLocationInformationHistoryEnabled,omitempty"`
	IsFlushExtensionAttributesEnabled *bool `json:"isFlushExtensionAttributesEnabled,omitempty"`
	FlushMDMQueue string `json:"flushMDMQueue"`
}

// NewReenrollment instantiates a new Reenrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReenrollment(flushMDMQueue string) *Reenrollment {
	this := Reenrollment{}
	var isFlushPolicyHistoryEnabled bool = false
	this.IsFlushPolicyHistoryEnabled = &isFlushPolicyHistoryEnabled
	var isFlushLocationInformationEnabled bool = false
	this.IsFlushLocationInformationEnabled = &isFlushLocationInformationEnabled
	var isFlushLocationInformationHistoryEnabled bool = false
	this.IsFlushLocationInformationHistoryEnabled = &isFlushLocationInformationHistoryEnabled
	var isFlushExtensionAttributesEnabled bool = false
	this.IsFlushExtensionAttributesEnabled = &isFlushExtensionAttributesEnabled
	this.FlushMDMQueue = flushMDMQueue
	return &this
}

// NewReenrollmentWithDefaults instantiates a new Reenrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReenrollmentWithDefaults() *Reenrollment {
	this := Reenrollment{}
	var isFlushPolicyHistoryEnabled bool = false
	this.IsFlushPolicyHistoryEnabled = &isFlushPolicyHistoryEnabled
	var isFlushLocationInformationEnabled bool = false
	this.IsFlushLocationInformationEnabled = &isFlushLocationInformationEnabled
	var isFlushLocationInformationHistoryEnabled bool = false
	this.IsFlushLocationInformationHistoryEnabled = &isFlushLocationInformationHistoryEnabled
	var isFlushExtensionAttributesEnabled bool = false
	this.IsFlushExtensionAttributesEnabled = &isFlushExtensionAttributesEnabled
	return &this
}

// GetIsFlushPolicyHistoryEnabled returns the IsFlushPolicyHistoryEnabled field value if set, zero value otherwise.
func (o *Reenrollment) GetIsFlushPolicyHistoryEnabled() bool {
	if o == nil || o.IsFlushPolicyHistoryEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsFlushPolicyHistoryEnabled
}

// GetIsFlushPolicyHistoryEnabledOk returns a tuple with the IsFlushPolicyHistoryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reenrollment) GetIsFlushPolicyHistoryEnabledOk() (*bool, bool) {
	if o == nil || o.IsFlushPolicyHistoryEnabled == nil {
		return nil, false
	}
	return o.IsFlushPolicyHistoryEnabled, true
}

// HasIsFlushPolicyHistoryEnabled returns a boolean if a field has been set.
func (o *Reenrollment) HasIsFlushPolicyHistoryEnabled() bool {
	if o != nil && o.IsFlushPolicyHistoryEnabled != nil {
		return true
	}

	return false
}

// SetIsFlushPolicyHistoryEnabled gets a reference to the given bool and assigns it to the IsFlushPolicyHistoryEnabled field.
func (o *Reenrollment) SetIsFlushPolicyHistoryEnabled(v bool) {
	o.IsFlushPolicyHistoryEnabled = &v
}

// GetIsFlushLocationInformationEnabled returns the IsFlushLocationInformationEnabled field value if set, zero value otherwise.
func (o *Reenrollment) GetIsFlushLocationInformationEnabled() bool {
	if o == nil || o.IsFlushLocationInformationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsFlushLocationInformationEnabled
}

// GetIsFlushLocationInformationEnabledOk returns a tuple with the IsFlushLocationInformationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reenrollment) GetIsFlushLocationInformationEnabledOk() (*bool, bool) {
	if o == nil || o.IsFlushLocationInformationEnabled == nil {
		return nil, false
	}
	return o.IsFlushLocationInformationEnabled, true
}

// HasIsFlushLocationInformationEnabled returns a boolean if a field has been set.
func (o *Reenrollment) HasIsFlushLocationInformationEnabled() bool {
	if o != nil && o.IsFlushLocationInformationEnabled != nil {
		return true
	}

	return false
}

// SetIsFlushLocationInformationEnabled gets a reference to the given bool and assigns it to the IsFlushLocationInformationEnabled field.
func (o *Reenrollment) SetIsFlushLocationInformationEnabled(v bool) {
	o.IsFlushLocationInformationEnabled = &v
}

// GetIsFlushLocationInformationHistoryEnabled returns the IsFlushLocationInformationHistoryEnabled field value if set, zero value otherwise.
func (o *Reenrollment) GetIsFlushLocationInformationHistoryEnabled() bool {
	if o == nil || o.IsFlushLocationInformationHistoryEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsFlushLocationInformationHistoryEnabled
}

// GetIsFlushLocationInformationHistoryEnabledOk returns a tuple with the IsFlushLocationInformationHistoryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reenrollment) GetIsFlushLocationInformationHistoryEnabledOk() (*bool, bool) {
	if o == nil || o.IsFlushLocationInformationHistoryEnabled == nil {
		return nil, false
	}
	return o.IsFlushLocationInformationHistoryEnabled, true
}

// HasIsFlushLocationInformationHistoryEnabled returns a boolean if a field has been set.
func (o *Reenrollment) HasIsFlushLocationInformationHistoryEnabled() bool {
	if o != nil && o.IsFlushLocationInformationHistoryEnabled != nil {
		return true
	}

	return false
}

// SetIsFlushLocationInformationHistoryEnabled gets a reference to the given bool and assigns it to the IsFlushLocationInformationHistoryEnabled field.
func (o *Reenrollment) SetIsFlushLocationInformationHistoryEnabled(v bool) {
	o.IsFlushLocationInformationHistoryEnabled = &v
}

// GetIsFlushExtensionAttributesEnabled returns the IsFlushExtensionAttributesEnabled field value if set, zero value otherwise.
func (o *Reenrollment) GetIsFlushExtensionAttributesEnabled() bool {
	if o == nil || o.IsFlushExtensionAttributesEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsFlushExtensionAttributesEnabled
}

// GetIsFlushExtensionAttributesEnabledOk returns a tuple with the IsFlushExtensionAttributesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Reenrollment) GetIsFlushExtensionAttributesEnabledOk() (*bool, bool) {
	if o == nil || o.IsFlushExtensionAttributesEnabled == nil {
		return nil, false
	}
	return o.IsFlushExtensionAttributesEnabled, true
}

// HasIsFlushExtensionAttributesEnabled returns a boolean if a field has been set.
func (o *Reenrollment) HasIsFlushExtensionAttributesEnabled() bool {
	if o != nil && o.IsFlushExtensionAttributesEnabled != nil {
		return true
	}

	return false
}

// SetIsFlushExtensionAttributesEnabled gets a reference to the given bool and assigns it to the IsFlushExtensionAttributesEnabled field.
func (o *Reenrollment) SetIsFlushExtensionAttributesEnabled(v bool) {
	o.IsFlushExtensionAttributesEnabled = &v
}

// GetFlushMDMQueue returns the FlushMDMQueue field value
func (o *Reenrollment) GetFlushMDMQueue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlushMDMQueue
}

// GetFlushMDMQueueOk returns a tuple with the FlushMDMQueue field value
// and a boolean to check if the value has been set.
func (o *Reenrollment) GetFlushMDMQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlushMDMQueue, true
}

// SetFlushMDMQueue sets field value
func (o *Reenrollment) SetFlushMDMQueue(v string) {
	o.FlushMDMQueue = v
}

func (o Reenrollment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsFlushPolicyHistoryEnabled != nil {
		toSerialize["isFlushPolicyHistoryEnabled"] = o.IsFlushPolicyHistoryEnabled
	}
	if o.IsFlushLocationInformationEnabled != nil {
		toSerialize["isFlushLocationInformationEnabled"] = o.IsFlushLocationInformationEnabled
	}
	if o.IsFlushLocationInformationHistoryEnabled != nil {
		toSerialize["isFlushLocationInformationHistoryEnabled"] = o.IsFlushLocationInformationHistoryEnabled
	}
	if o.IsFlushExtensionAttributesEnabled != nil {
		toSerialize["isFlushExtensionAttributesEnabled"] = o.IsFlushExtensionAttributesEnabled
	}
	if true {
		toSerialize["flushMDMQueue"] = o.FlushMDMQueue
	}
	return json.Marshal(toSerialize)
}

type NullableReenrollment struct {
	value *Reenrollment
	isSet bool
}

func (v NullableReenrollment) Get() *Reenrollment {
	return v.value
}

func (v *NullableReenrollment) Set(val *Reenrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableReenrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableReenrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReenrollment(val *Reenrollment) *NullableReenrollment {
	return &NullableReenrollment{value: val, isSet: true}
}

func (v NullableReenrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReenrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

