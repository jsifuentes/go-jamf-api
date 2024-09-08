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

// checks if the ApiRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRoleRequest{}

// ApiRoleRequest struct for ApiRoleRequest
type ApiRoleRequest struct {
	DisplayName string `json:"displayName"`
	Privileges []string `json:"privileges"`
}

type _ApiRoleRequest ApiRoleRequest

// NewApiRoleRequest instantiates a new ApiRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRoleRequest(displayName string, privileges []string) *ApiRoleRequest {
	this := ApiRoleRequest{}
	this.DisplayName = displayName
	this.Privileges = privileges
	return &this
}

// NewApiRoleRequestWithDefaults instantiates a new ApiRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRoleRequestWithDefaults() *ApiRoleRequest {
	this := ApiRoleRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *ApiRoleRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ApiRoleRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ApiRoleRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPrivileges returns the Privileges field value
func (o *ApiRoleRequest) GetPrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *ApiRoleRequest) GetPrivilegesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Privileges, true
}

// SetPrivileges sets field value
func (o *ApiRoleRequest) SetPrivileges(v []string) {
	o.Privileges = v
}

func (o ApiRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["privileges"] = o.Privileges
	return toSerialize, nil
}

func (o *ApiRoleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"privileges",
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

	varApiRoleRequest := _ApiRoleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiRoleRequest)

	if err != nil {
		return err
	}

	*o = ApiRoleRequest(varApiRoleRequest)

	return err
}

type NullableApiRoleRequest struct {
	value *ApiRoleRequest
	isSet bool
}

func (v NullableApiRoleRequest) Get() *ApiRoleRequest {
	return v.value
}

func (v *NullableApiRoleRequest) Set(val *ApiRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRoleRequest(val *ApiRoleRequest) *NullableApiRoleRequest {
	return &NullableApiRoleRequest{value: val, isSet: true}
}

func (v NullableApiRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

