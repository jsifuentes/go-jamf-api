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

// checks if the ManagedSoftwareUpdatePlanEventStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedSoftwareUpdatePlanEventStore{}

// ManagedSoftwareUpdatePlanEventStore struct for ManagedSoftwareUpdatePlanEventStore
type ManagedSoftwareUpdatePlanEventStore struct {
	Events *string `json:"events,omitempty"`
}

// NewManagedSoftwareUpdatePlanEventStore instantiates a new ManagedSoftwareUpdatePlanEventStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedSoftwareUpdatePlanEventStore() *ManagedSoftwareUpdatePlanEventStore {
	this := ManagedSoftwareUpdatePlanEventStore{}
	return &this
}

// NewManagedSoftwareUpdatePlanEventStoreWithDefaults instantiates a new ManagedSoftwareUpdatePlanEventStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedSoftwareUpdatePlanEventStoreWithDefaults() *ManagedSoftwareUpdatePlanEventStore {
	this := ManagedSoftwareUpdatePlanEventStore{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ManagedSoftwareUpdatePlanEventStore) GetEvents() string {
	if o == nil || IsNil(o.Events) {
		var ret string
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedSoftwareUpdatePlanEventStore) GetEventsOk() (*string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ManagedSoftwareUpdatePlanEventStore) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given string and assigns it to the Events field.
func (o *ManagedSoftwareUpdatePlanEventStore) SetEvents(v string) {
	o.Events = &v
}

func (o ManagedSoftwareUpdatePlanEventStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedSoftwareUpdatePlanEventStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableManagedSoftwareUpdatePlanEventStore struct {
	value *ManagedSoftwareUpdatePlanEventStore
	isSet bool
}

func (v NullableManagedSoftwareUpdatePlanEventStore) Get() *ManagedSoftwareUpdatePlanEventStore {
	return v.value
}

func (v *NullableManagedSoftwareUpdatePlanEventStore) Set(val *ManagedSoftwareUpdatePlanEventStore) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedSoftwareUpdatePlanEventStore) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedSoftwareUpdatePlanEventStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedSoftwareUpdatePlanEventStore(val *ManagedSoftwareUpdatePlanEventStore) *NullableManagedSoftwareUpdatePlanEventStore {
	return &NullableManagedSoftwareUpdatePlanEventStore{value: val, isSet: true}
}

func (v NullableManagedSoftwareUpdatePlanEventStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedSoftwareUpdatePlanEventStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

