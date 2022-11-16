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

// EnrollmentCustomizationPanelTextAllOf struct for EnrollmentCustomizationPanelTextAllOf
type EnrollmentCustomizationPanelTextAllOf struct {
	Body string `json:"body"`
	Subtext *string `json:"subtext,omitempty"`
	Title string `json:"title"`
	BackButtonText string `json:"backButtonText"`
	ContinueButtonText string `json:"continueButtonText"`
}

// NewEnrollmentCustomizationPanelTextAllOf instantiates a new EnrollmentCustomizationPanelTextAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentCustomizationPanelTextAllOf(body string, title string, backButtonText string, continueButtonText string) *EnrollmentCustomizationPanelTextAllOf {
	this := EnrollmentCustomizationPanelTextAllOf{}
	this.Body = body
	this.Title = title
	this.BackButtonText = backButtonText
	this.ContinueButtonText = continueButtonText
	return &this
}

// NewEnrollmentCustomizationPanelTextAllOfWithDefaults instantiates a new EnrollmentCustomizationPanelTextAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentCustomizationPanelTextAllOfWithDefaults() *EnrollmentCustomizationPanelTextAllOf {
	this := EnrollmentCustomizationPanelTextAllOf{}
	return &this
}

// GetBody returns the Body field value
func (o *EnrollmentCustomizationPanelTextAllOf) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelTextAllOf) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *EnrollmentCustomizationPanelTextAllOf) SetBody(v string) {
	o.Body = v
}

// GetSubtext returns the Subtext field value if set, zero value otherwise.
func (o *EnrollmentCustomizationPanelTextAllOf) GetSubtext() string {
	if o == nil || o.Subtext == nil {
		var ret string
		return ret
	}
	return *o.Subtext
}

// GetSubtextOk returns a tuple with the Subtext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelTextAllOf) GetSubtextOk() (*string, bool) {
	if o == nil || o.Subtext == nil {
		return nil, false
	}
	return o.Subtext, true
}

// HasSubtext returns a boolean if a field has been set.
func (o *EnrollmentCustomizationPanelTextAllOf) HasSubtext() bool {
	if o != nil && o.Subtext != nil {
		return true
	}

	return false
}

// SetSubtext gets a reference to the given string and assigns it to the Subtext field.
func (o *EnrollmentCustomizationPanelTextAllOf) SetSubtext(v string) {
	o.Subtext = &v
}

// GetTitle returns the Title field value
func (o *EnrollmentCustomizationPanelTextAllOf) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelTextAllOf) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *EnrollmentCustomizationPanelTextAllOf) SetTitle(v string) {
	o.Title = v
}

// GetBackButtonText returns the BackButtonText field value
func (o *EnrollmentCustomizationPanelTextAllOf) GetBackButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackButtonText
}

// GetBackButtonTextOk returns a tuple with the BackButtonText field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelTextAllOf) GetBackButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackButtonText, true
}

// SetBackButtonText sets field value
func (o *EnrollmentCustomizationPanelTextAllOf) SetBackButtonText(v string) {
	o.BackButtonText = v
}

// GetContinueButtonText returns the ContinueButtonText field value
func (o *EnrollmentCustomizationPanelTextAllOf) GetContinueButtonText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContinueButtonText
}

// GetContinueButtonTextOk returns a tuple with the ContinueButtonText field value
// and a boolean to check if the value has been set.
func (o *EnrollmentCustomizationPanelTextAllOf) GetContinueButtonTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContinueButtonText, true
}

// SetContinueButtonText sets field value
func (o *EnrollmentCustomizationPanelTextAllOf) SetContinueButtonText(v string) {
	o.ContinueButtonText = v
}

func (o EnrollmentCustomizationPanelTextAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["body"] = o.Body
	}
	if o.Subtext != nil {
		toSerialize["subtext"] = o.Subtext
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["backButtonText"] = o.BackButtonText
	}
	if true {
		toSerialize["continueButtonText"] = o.ContinueButtonText
	}
	return json.Marshal(toSerialize)
}

type NullableEnrollmentCustomizationPanelTextAllOf struct {
	value *EnrollmentCustomizationPanelTextAllOf
	isSet bool
}

func (v NullableEnrollmentCustomizationPanelTextAllOf) Get() *EnrollmentCustomizationPanelTextAllOf {
	return v.value
}

func (v *NullableEnrollmentCustomizationPanelTextAllOf) Set(val *EnrollmentCustomizationPanelTextAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentCustomizationPanelTextAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentCustomizationPanelTextAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentCustomizationPanelTextAllOf(val *EnrollmentCustomizationPanelTextAllOf) *NullableEnrollmentCustomizationPanelTextAllOf {
	return &NullableEnrollmentCustomizationPanelTextAllOf{value: val, isSet: true}
}

func (v NullableEnrollmentCustomizationPanelTextAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentCustomizationPanelTextAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

