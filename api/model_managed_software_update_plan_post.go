/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ManagedSoftwareUpdatePlanPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedSoftwareUpdatePlanPost{}

// ManagedSoftwareUpdatePlanPost struct for ManagedSoftwareUpdatePlanPost
type ManagedSoftwareUpdatePlanPost struct {
	Devices []PlanDevicePost `json:"devices"`
	Config PlanConfigurationPost `json:"config"`
}

type _ManagedSoftwareUpdatePlanPost ManagedSoftwareUpdatePlanPost

// NewManagedSoftwareUpdatePlanPost instantiates a new ManagedSoftwareUpdatePlanPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedSoftwareUpdatePlanPost(devices []PlanDevicePost, config PlanConfigurationPost) *ManagedSoftwareUpdatePlanPost {
	this := ManagedSoftwareUpdatePlanPost{}
	this.Devices = devices
	this.Config = config
	return &this
}

// NewManagedSoftwareUpdatePlanPostWithDefaults instantiates a new ManagedSoftwareUpdatePlanPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedSoftwareUpdatePlanPostWithDefaults() *ManagedSoftwareUpdatePlanPost {
	this := ManagedSoftwareUpdatePlanPost{}
	return &this
}

// GetDevices returns the Devices field value
func (o *ManagedSoftwareUpdatePlanPost) GetDevices() []PlanDevicePost {
	if o == nil {
		var ret []PlanDevicePost
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *ManagedSoftwareUpdatePlanPost) GetDevicesOk() ([]PlanDevicePost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *ManagedSoftwareUpdatePlanPost) SetDevices(v []PlanDevicePost) {
	o.Devices = v
}

// GetConfig returns the Config field value
func (o *ManagedSoftwareUpdatePlanPost) GetConfig() PlanConfigurationPost {
	if o == nil {
		var ret PlanConfigurationPost
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ManagedSoftwareUpdatePlanPost) GetConfigOk() (*PlanConfigurationPost, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ManagedSoftwareUpdatePlanPost) SetConfig(v PlanConfigurationPost) {
	o.Config = v
}

func (o ManagedSoftwareUpdatePlanPost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedSoftwareUpdatePlanPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["devices"] = o.Devices
	toSerialize["config"] = o.Config
	return toSerialize, nil
}

func (o *ManagedSoftwareUpdatePlanPost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"devices",
		"config",
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

	varManagedSoftwareUpdatePlanPost := _ManagedSoftwareUpdatePlanPost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varManagedSoftwareUpdatePlanPost)

	if err != nil {
		return err
	}

	*o = ManagedSoftwareUpdatePlanPost(varManagedSoftwareUpdatePlanPost)

	return err
}

type NullableManagedSoftwareUpdatePlanPost struct {
	value *ManagedSoftwareUpdatePlanPost
	isSet bool
}

func (v NullableManagedSoftwareUpdatePlanPost) Get() *ManagedSoftwareUpdatePlanPost {
	return v.value
}

func (v *NullableManagedSoftwareUpdatePlanPost) Set(val *ManagedSoftwareUpdatePlanPost) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedSoftwareUpdatePlanPost) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedSoftwareUpdatePlanPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedSoftwareUpdatePlanPost(val *ManagedSoftwareUpdatePlanPost) *NullableManagedSoftwareUpdatePlanPost {
	return &NullableManagedSoftwareUpdatePlanPost{value: val, isSet: true}
}

func (v NullableManagedSoftwareUpdatePlanPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedSoftwareUpdatePlanPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

