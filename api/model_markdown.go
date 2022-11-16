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

// Markdown struct for Markdown
type Markdown struct {
	Markdown *string `json:"markdown,omitempty"`
}

// NewMarkdown instantiates a new Markdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkdown() *Markdown {
	this := Markdown{}
	return &this
}

// NewMarkdownWithDefaults instantiates a new Markdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkdownWithDefaults() *Markdown {
	this := Markdown{}
	return &this
}

// GetMarkdown returns the Markdown field value if set, zero value otherwise.
func (o *Markdown) GetMarkdown() string {
	if o == nil || o.Markdown == nil {
		var ret string
		return ret
	}
	return *o.Markdown
}

// GetMarkdownOk returns a tuple with the Markdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Markdown) GetMarkdownOk() (*string, bool) {
	if o == nil || o.Markdown == nil {
		return nil, false
	}
	return o.Markdown, true
}

// HasMarkdown returns a boolean if a field has been set.
func (o *Markdown) HasMarkdown() bool {
	if o != nil && o.Markdown != nil {
		return true
	}

	return false
}

// SetMarkdown gets a reference to the given string and assigns it to the Markdown field.
func (o *Markdown) SetMarkdown(v string) {
	o.Markdown = &v
}

func (o Markdown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Markdown != nil {
		toSerialize["markdown"] = o.Markdown
	}
	return json.Marshal(toSerialize)
}

type NullableMarkdown struct {
	value *Markdown
	isSet bool
}

func (v NullableMarkdown) Get() *Markdown {
	return v.value
}

func (v *NullableMarkdown) Set(val *Markdown) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkdown) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkdown(val *Markdown) *NullableMarkdown {
	return &NullableMarkdown{value: val, isSet: true}
}

func (v NullableMarkdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

