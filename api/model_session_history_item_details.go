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

// checks if the SessionHistoryItemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionHistoryItemDetails{}

// SessionHistoryItemDetails struct for SessionHistoryItemDetails
type SessionHistoryItemDetails struct {
	FileTransferItemList []FileTransferItem `json:"fileTransferItemList,omitempty"`
}

// NewSessionHistoryItemDetails instantiates a new SessionHistoryItemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionHistoryItemDetails() *SessionHistoryItemDetails {
	this := SessionHistoryItemDetails{}
	return &this
}

// NewSessionHistoryItemDetailsWithDefaults instantiates a new SessionHistoryItemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionHistoryItemDetailsWithDefaults() *SessionHistoryItemDetails {
	this := SessionHistoryItemDetails{}
	return &this
}

// GetFileTransferItemList returns the FileTransferItemList field value if set, zero value otherwise.
func (o *SessionHistoryItemDetails) GetFileTransferItemList() []FileTransferItem {
	if o == nil || IsNil(o.FileTransferItemList) {
		var ret []FileTransferItem
		return ret
	}
	return o.FileTransferItemList
}

// GetFileTransferItemListOk returns a tuple with the FileTransferItemList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionHistoryItemDetails) GetFileTransferItemListOk() ([]FileTransferItem, bool) {
	if o == nil || IsNil(o.FileTransferItemList) {
		return nil, false
	}
	return o.FileTransferItemList, true
}

// HasFileTransferItemList returns a boolean if a field has been set.
func (o *SessionHistoryItemDetails) HasFileTransferItemList() bool {
	if o != nil && !IsNil(o.FileTransferItemList) {
		return true
	}

	return false
}

// SetFileTransferItemList gets a reference to the given []FileTransferItem and assigns it to the FileTransferItemList field.
func (o *SessionHistoryItemDetails) SetFileTransferItemList(v []FileTransferItem) {
	o.FileTransferItemList = v
}

func (o SessionHistoryItemDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionHistoryItemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileTransferItemList) {
		toSerialize["fileTransferItemList"] = o.FileTransferItemList
	}
	return toSerialize, nil
}

type NullableSessionHistoryItemDetails struct {
	value *SessionHistoryItemDetails
	isSet bool
}

func (v NullableSessionHistoryItemDetails) Get() *SessionHistoryItemDetails {
	return v.value
}

func (v *NullableSessionHistoryItemDetails) Set(val *SessionHistoryItemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionHistoryItemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionHistoryItemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionHistoryItemDetails(val *SessionHistoryItemDetails) *NullableSessionHistoryItemDetails {
	return &NullableSessionHistoryItemDetails{value: val, isSet: true}
}

func (v NullableSessionHistoryItemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionHistoryItemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

