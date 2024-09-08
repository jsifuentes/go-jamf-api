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

// checks if the PatchPolicyLogDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchPolicyLogDetail{}

// PatchPolicyLogDetail struct for PatchPolicyLogDetail
type PatchPolicyLogDetail struct {
	Id *string `json:"id,omitempty"`
	AttemptNumber *int64 `json:"attemptNumber,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	Actions []PatchPolicyLogDetailAction `json:"actions,omitempty"`
}

// NewPatchPolicyLogDetail instantiates a new PatchPolicyLogDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPolicyLogDetail() *PatchPolicyLogDetail {
	this := PatchPolicyLogDetail{}
	return &this
}

// NewPatchPolicyLogDetailWithDefaults instantiates a new PatchPolicyLogDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPolicyLogDetailWithDefaults() *PatchPolicyLogDetail {
	this := PatchPolicyLogDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchPolicyLogDetail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicyLogDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchPolicyLogDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchPolicyLogDetail) SetId(v string) {
	o.Id = &v
}

// GetAttemptNumber returns the AttemptNumber field value if set, zero value otherwise.
func (o *PatchPolicyLogDetail) GetAttemptNumber() int64 {
	if o == nil || IsNil(o.AttemptNumber) {
		var ret int64
		return ret
	}
	return *o.AttemptNumber
}

// GetAttemptNumberOk returns a tuple with the AttemptNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicyLogDetail) GetAttemptNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.AttemptNumber) {
		return nil, false
	}
	return o.AttemptNumber, true
}

// HasAttemptNumber returns a boolean if a field has been set.
func (o *PatchPolicyLogDetail) HasAttemptNumber() bool {
	if o != nil && !IsNil(o.AttemptNumber) {
		return true
	}

	return false
}

// SetAttemptNumber gets a reference to the given int64 and assigns it to the AttemptNumber field.
func (o *PatchPolicyLogDetail) SetAttemptNumber(v int64) {
	o.AttemptNumber = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *PatchPolicyLogDetail) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicyLogDetail) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *PatchPolicyLogDetail) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *PatchPolicyLogDetail) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *PatchPolicyLogDetail) GetActions() []PatchPolicyLogDetailAction {
	if o == nil || IsNil(o.Actions) {
		var ret []PatchPolicyLogDetailAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPolicyLogDetail) GetActionsOk() ([]PatchPolicyLogDetailAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *PatchPolicyLogDetail) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []PatchPolicyLogDetailAction and assigns it to the Actions field.
func (o *PatchPolicyLogDetail) SetActions(v []PatchPolicyLogDetailAction) {
	o.Actions = v
}

func (o PatchPolicyLogDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchPolicyLogDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AttemptNumber) {
		toSerialize["attemptNumber"] = o.AttemptNumber
	}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	return toSerialize, nil
}

type NullablePatchPolicyLogDetail struct {
	value *PatchPolicyLogDetail
	isSet bool
}

func (v NullablePatchPolicyLogDetail) Get() *PatchPolicyLogDetail {
	return v.value
}

func (v *NullablePatchPolicyLogDetail) Set(val *PatchPolicyLogDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPolicyLogDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPolicyLogDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPolicyLogDetail(val *PatchPolicyLogDetail) *NullablePatchPolicyLogDetail {
	return &NullablePatchPolicyLogDetail{value: val, isSet: true}
}

func (v NullablePatchPolicyLogDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPolicyLogDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

