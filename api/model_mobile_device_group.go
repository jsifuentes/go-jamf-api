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

// MobileDeviceGroup struct for MobileDeviceGroup
type MobileDeviceGroup struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	IsSmartGroup *bool `json:"isSmartGroup,omitempty"`
}

// NewMobileDeviceGroup instantiates a new MobileDeviceGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileDeviceGroup() *MobileDeviceGroup {
	this := MobileDeviceGroup{}
	return &this
}

// NewMobileDeviceGroupWithDefaults instantiates a new MobileDeviceGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileDeviceGroupWithDefaults() *MobileDeviceGroup {
	this := MobileDeviceGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MobileDeviceGroup) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceGroup) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MobileDeviceGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MobileDeviceGroup) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MobileDeviceGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MobileDeviceGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MobileDeviceGroup) SetName(v string) {
	o.Name = &v
}

// GetIsSmartGroup returns the IsSmartGroup field value if set, zero value otherwise.
func (o *MobileDeviceGroup) GetIsSmartGroup() bool {
	if o == nil || o.IsSmartGroup == nil {
		var ret bool
		return ret
	}
	return *o.IsSmartGroup
}

// GetIsSmartGroupOk returns a tuple with the IsSmartGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileDeviceGroup) GetIsSmartGroupOk() (*bool, bool) {
	if o == nil || o.IsSmartGroup == nil {
		return nil, false
	}
	return o.IsSmartGroup, true
}

// HasIsSmartGroup returns a boolean if a field has been set.
func (o *MobileDeviceGroup) HasIsSmartGroup() bool {
	if o != nil && o.IsSmartGroup != nil {
		return true
	}

	return false
}

// SetIsSmartGroup gets a reference to the given bool and assigns it to the IsSmartGroup field.
func (o *MobileDeviceGroup) SetIsSmartGroup(v bool) {
	o.IsSmartGroup = &v
}

func (o MobileDeviceGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.IsSmartGroup != nil {
		toSerialize["isSmartGroup"] = o.IsSmartGroup
	}
	return json.Marshal(toSerialize)
}

type NullableMobileDeviceGroup struct {
	value *MobileDeviceGroup
	isSet bool
}

func (v NullableMobileDeviceGroup) Get() *MobileDeviceGroup {
	return v.value
}

func (v *NullableMobileDeviceGroup) Set(val *MobileDeviceGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileDeviceGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileDeviceGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileDeviceGroup(val *MobileDeviceGroup) *NullableMobileDeviceGroup {
	return &NullableMobileDeviceGroup{value: val, isSet: true}
}

func (v NullableMobileDeviceGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileDeviceGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

