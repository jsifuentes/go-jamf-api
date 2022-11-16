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

// ApiError struct for ApiError
type ApiError struct {
	// HTTP status of the response
	HttpStatus *int32 `json:"httpStatus,omitempty"`
	Errors []ApiErrorCause `json:"errors,omitempty"`
}

// NewApiError instantiates a new ApiError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiError() *ApiError {
	this := ApiError{}
	return &this
}

// NewApiErrorWithDefaults instantiates a new ApiError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorWithDefaults() *ApiError {
	this := ApiError{}
	return &this
}

// GetHttpStatus returns the HttpStatus field value if set, zero value otherwise.
func (o *ApiError) GetHttpStatus() int32 {
	if o == nil || o.HttpStatus == nil {
		var ret int32
		return ret
	}
	return *o.HttpStatus
}

// GetHttpStatusOk returns a tuple with the HttpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetHttpStatusOk() (*int32, bool) {
	if o == nil || o.HttpStatus == nil {
		return nil, false
	}
	return o.HttpStatus, true
}

// HasHttpStatus returns a boolean if a field has been set.
func (o *ApiError) HasHttpStatus() bool {
	if o != nil && o.HttpStatus != nil {
		return true
	}

	return false
}

// SetHttpStatus gets a reference to the given int32 and assigns it to the HttpStatus field.
func (o *ApiError) SetHttpStatus(v int32) {
	o.HttpStatus = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ApiError) GetErrors() []ApiErrorCause {
	if o == nil || o.Errors == nil {
		var ret []ApiErrorCause
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetErrorsOk() ([]ApiErrorCause, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ApiError) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ApiErrorCause and assigns it to the Errors field.
func (o *ApiError) SetErrors(v []ApiErrorCause) {
	o.Errors = v
}

func (o ApiError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpStatus != nil {
		toSerialize["httpStatus"] = o.HttpStatus
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableApiError struct {
	value *ApiError
	isSet bool
}

func (v NullableApiError) Get() *ApiError {
	return v.value
}

func (v *NullableApiError) Set(val *ApiError) {
	v.value = val
	v.isSet = true
}

func (v NullableApiError) IsSet() bool {
	return v.isSet
}

func (v *NullableApiError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiError(val *ApiError) *NullableApiError {
	return &NullableApiError{value: val, isSet: true}
}

func (v NullableApiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

