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

// checks if the StaticGroupSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticGroupSummary{}

// StaticGroupSummary struct for StaticGroupSummary
type StaticGroupSummary struct {
	GroupId *string `json:"groupId,omitempty"`
	GroupName *string `json:"groupName,omitempty"`
	SiteId *string `json:"siteId,omitempty"`
	// membership count
	Count *int64 `json:"count,omitempty"`
}

// NewStaticGroupSummary instantiates a new StaticGroupSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticGroupSummary() *StaticGroupSummary {
	this := StaticGroupSummary{}
	return &this
}

// NewStaticGroupSummaryWithDefaults instantiates a new StaticGroupSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticGroupSummaryWithDefaults() *StaticGroupSummary {
	this := StaticGroupSummary{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *StaticGroupSummary) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticGroupSummary) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *StaticGroupSummary) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *StaticGroupSummary) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *StaticGroupSummary) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticGroupSummary) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *StaticGroupSummary) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *StaticGroupSummary) SetGroupName(v string) {
	o.GroupName = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *StaticGroupSummary) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticGroupSummary) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *StaticGroupSummary) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *StaticGroupSummary) SetSiteId(v string) {
	o.SiteId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *StaticGroupSummary) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticGroupSummary) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *StaticGroupSummary) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *StaticGroupSummary) SetCount(v int64) {
	o.Count = &v
}

func (o StaticGroupSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticGroupSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableStaticGroupSummary struct {
	value *StaticGroupSummary
	isSet bool
}

func (v NullableStaticGroupSummary) Get() *StaticGroupSummary {
	return v.value
}

func (v *NullableStaticGroupSummary) Set(val *StaticGroupSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticGroupSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticGroupSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticGroupSummary(val *StaticGroupSummary) *NullableStaticGroupSummary {
	return &NullableStaticGroupSummary{value: val, isSet: true}
}

func (v NullableStaticGroupSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticGroupSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

