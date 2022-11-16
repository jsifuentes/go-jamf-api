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

// DeprecatedConfigurationUpdate An old Cloud Identity Provider configuration for updates
type DeprecatedConfigurationUpdate struct {
	Server DeprecatedServerUpdate `json:"server"`
	Mappings *CloudLdapMappingsRequest `json:"mappings,omitempty"`
}

// NewDeprecatedConfigurationUpdate instantiates a new DeprecatedConfigurationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedConfigurationUpdate(server DeprecatedServerUpdate) *DeprecatedConfigurationUpdate {
	this := DeprecatedConfigurationUpdate{}
	this.Server = server
	return &this
}

// NewDeprecatedConfigurationUpdateWithDefaults instantiates a new DeprecatedConfigurationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedConfigurationUpdateWithDefaults() *DeprecatedConfigurationUpdate {
	this := DeprecatedConfigurationUpdate{}
	return &this
}

// GetServer returns the Server field value
func (o *DeprecatedConfigurationUpdate) GetServer() DeprecatedServerUpdate {
	if o == nil {
		var ret DeprecatedServerUpdate
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *DeprecatedConfigurationUpdate) GetServerOk() (*DeprecatedServerUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *DeprecatedConfigurationUpdate) SetServer(v DeprecatedServerUpdate) {
	o.Server = v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *DeprecatedConfigurationUpdate) GetMappings() CloudLdapMappingsRequest {
	if o == nil || o.Mappings == nil {
		var ret CloudLdapMappingsRequest
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedConfigurationUpdate) GetMappingsOk() (*CloudLdapMappingsRequest, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *DeprecatedConfigurationUpdate) HasMappings() bool {
	if o != nil && o.Mappings != nil {
		return true
	}

	return false
}

// SetMappings gets a reference to the given CloudLdapMappingsRequest and assigns it to the Mappings field.
func (o *DeprecatedConfigurationUpdate) SetMappings(v CloudLdapMappingsRequest) {
	o.Mappings = &v
}

func (o DeprecatedConfigurationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["server"] = o.Server
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	return json.Marshal(toSerialize)
}

type NullableDeprecatedConfigurationUpdate struct {
	value *DeprecatedConfigurationUpdate
	isSet bool
}

func (v NullableDeprecatedConfigurationUpdate) Get() *DeprecatedConfigurationUpdate {
	return v.value
}

func (v *NullableDeprecatedConfigurationUpdate) Set(val *DeprecatedConfigurationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedConfigurationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedConfigurationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedConfigurationUpdate(val *DeprecatedConfigurationUpdate) *NullableDeprecatedConfigurationUpdate {
	return &NullableDeprecatedConfigurationUpdate{value: val, isSet: true}
}

func (v NullableDeprecatedConfigurationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedConfigurationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

