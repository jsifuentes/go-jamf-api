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

// checks if the DashboardSetupFeatureOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardSetupFeatureOptions{}

// DashboardSetupFeatureOptions struct for DashboardSetupFeatureOptions
type DashboardSetupFeatureOptions struct {
	Feature *string `json:"feature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DashboardSetupFeatureOptions DashboardSetupFeatureOptions

// NewDashboardSetupFeatureOptions instantiates a new DashboardSetupFeatureOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardSetupFeatureOptions() *DashboardSetupFeatureOptions {
	this := DashboardSetupFeatureOptions{}
	return &this
}

// NewDashboardSetupFeatureOptionsWithDefaults instantiates a new DashboardSetupFeatureOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardSetupFeatureOptionsWithDefaults() *DashboardSetupFeatureOptions {
	this := DashboardSetupFeatureOptions{}
	return &this
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *DashboardSetupFeatureOptions) GetFeature() string {
	if o == nil || IsNil(o.Feature) {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSetupFeatureOptions) GetFeatureOk() (*string, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *DashboardSetupFeatureOptions) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *DashboardSetupFeatureOptions) SetFeature(v string) {
	o.Feature = &v
}

func (o DashboardSetupFeatureOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardSetupFeatureOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DashboardSetupFeatureOptions) UnmarshalJSON(data []byte) (err error) {
	varDashboardSetupFeatureOptions := _DashboardSetupFeatureOptions{}

	err = json.Unmarshal(data, &varDashboardSetupFeatureOptions)

	if err != nil {
		return err
	}

	*o = DashboardSetupFeatureOptions(varDashboardSetupFeatureOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feature")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDashboardSetupFeatureOptions struct {
	value *DashboardSetupFeatureOptions
	isSet bool
}

func (v NullableDashboardSetupFeatureOptions) Get() *DashboardSetupFeatureOptions {
	return v.value
}

func (v *NullableDashboardSetupFeatureOptions) Set(val *DashboardSetupFeatureOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardSetupFeatureOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardSetupFeatureOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardSetupFeatureOptions(val *DashboardSetupFeatureOptions) *NullableDashboardSetupFeatureOptions {
	return &NullableDashboardSetupFeatureOptions{value: val, isSet: true}
}

func (v NullableDashboardSetupFeatureOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardSetupFeatureOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

