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

// checks if the PatchSoftwareTitleConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchSoftwareTitleConfiguration{}

// PatchSoftwareTitleConfiguration struct for PatchSoftwareTitleConfiguration
type PatchSoftwareTitleConfiguration struct {
	DisplayName string `json:"displayName"`
	CategoryId *string `json:"categoryId,omitempty"`
	SiteId *string `json:"siteId,omitempty"`
	UiNotifications *bool `json:"uiNotifications,omitempty"`
	EmailNotifications *bool `json:"emailNotifications,omitempty"`
	SoftwareTitleId string `json:"softwareTitleId"`
	JamfOfficial *bool `json:"jamfOfficial,omitempty"`
	ExtensionAttributes []PatchSoftwareTitleConfigurationExtensionAttributes `json:"extensionAttributes,omitempty"`
	SoftwareTitleName *string `json:"softwareTitleName,omitempty"`
	SoftwareTitleNameId *string `json:"softwareTitleNameId,omitempty"`
	SoftwareTitlePublisher *string `json:"softwareTitlePublisher,omitempty"`
	PatchSourceName *string `json:"patchSourceName,omitempty"`
	PatchSourceEnabled *bool `json:"patchSourceEnabled,omitempty"`
	Id *string `json:"id,omitempty"`
	Packages []PatchSoftwareTitlePackages `json:"packages,omitempty"`
}

type _PatchSoftwareTitleConfiguration PatchSoftwareTitleConfiguration

// NewPatchSoftwareTitleConfiguration instantiates a new PatchSoftwareTitleConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchSoftwareTitleConfiguration(displayName string, softwareTitleId string) *PatchSoftwareTitleConfiguration {
	this := PatchSoftwareTitleConfiguration{}
	this.DisplayName = displayName
	var categoryId string = "-1"
	this.CategoryId = &categoryId
	var siteId string = "-1"
	this.SiteId = &siteId
	var uiNotifications bool = false
	this.UiNotifications = &uiNotifications
	var emailNotifications bool = false
	this.EmailNotifications = &emailNotifications
	this.SoftwareTitleId = softwareTitleId
	return &this
}

// NewPatchSoftwareTitleConfigurationWithDefaults instantiates a new PatchSoftwareTitleConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchSoftwareTitleConfigurationWithDefaults() *PatchSoftwareTitleConfiguration {
	this := PatchSoftwareTitleConfiguration{}
	var categoryId string = "-1"
	this.CategoryId = &categoryId
	var siteId string = "-1"
	this.SiteId = &siteId
	var uiNotifications bool = false
	this.UiNotifications = &uiNotifications
	var emailNotifications bool = false
	this.EmailNotifications = &emailNotifications
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *PatchSoftwareTitleConfiguration) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PatchSoftwareTitleConfiguration) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *PatchSoftwareTitleConfiguration) SetCategoryId(v string) {
	o.CategoryId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *PatchSoftwareTitleConfiguration) SetSiteId(v string) {
	o.SiteId = &v
}

// GetUiNotifications returns the UiNotifications field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetUiNotifications() bool {
	if o == nil || IsNil(o.UiNotifications) {
		var ret bool
		return ret
	}
	return *o.UiNotifications
}

// GetUiNotificationsOk returns a tuple with the UiNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetUiNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.UiNotifications) {
		return nil, false
	}
	return o.UiNotifications, true
}

// HasUiNotifications returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasUiNotifications() bool {
	if o != nil && !IsNil(o.UiNotifications) {
		return true
	}

	return false
}

// SetUiNotifications gets a reference to the given bool and assigns it to the UiNotifications field.
func (o *PatchSoftwareTitleConfiguration) SetUiNotifications(v bool) {
	o.UiNotifications = &v
}

// GetEmailNotifications returns the EmailNotifications field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetEmailNotifications() bool {
	if o == nil || IsNil(o.EmailNotifications) {
		var ret bool
		return ret
	}
	return *o.EmailNotifications
}

// GetEmailNotificationsOk returns a tuple with the EmailNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetEmailNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailNotifications) {
		return nil, false
	}
	return o.EmailNotifications, true
}

// HasEmailNotifications returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasEmailNotifications() bool {
	if o != nil && !IsNil(o.EmailNotifications) {
		return true
	}

	return false
}

// SetEmailNotifications gets a reference to the given bool and assigns it to the EmailNotifications field.
func (o *PatchSoftwareTitleConfiguration) SetEmailNotifications(v bool) {
	o.EmailNotifications = &v
}

// GetSoftwareTitleId returns the SoftwareTitleId field value
func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SoftwareTitleId
}

// GetSoftwareTitleIdOk returns a tuple with the SoftwareTitleId field value
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SoftwareTitleId, true
}

// SetSoftwareTitleId sets field value
func (o *PatchSoftwareTitleConfiguration) SetSoftwareTitleId(v string) {
	o.SoftwareTitleId = v
}

// GetJamfOfficial returns the JamfOfficial field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetJamfOfficial() bool {
	if o == nil || IsNil(o.JamfOfficial) {
		var ret bool
		return ret
	}
	return *o.JamfOfficial
}

// GetJamfOfficialOk returns a tuple with the JamfOfficial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetJamfOfficialOk() (*bool, bool) {
	if o == nil || IsNil(o.JamfOfficial) {
		return nil, false
	}
	return o.JamfOfficial, true
}

// HasJamfOfficial returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasJamfOfficial() bool {
	if o != nil && !IsNil(o.JamfOfficial) {
		return true
	}

	return false
}

// SetJamfOfficial gets a reference to the given bool and assigns it to the JamfOfficial field.
func (o *PatchSoftwareTitleConfiguration) SetJamfOfficial(v bool) {
	o.JamfOfficial = &v
}

// GetExtensionAttributes returns the ExtensionAttributes field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetExtensionAttributes() []PatchSoftwareTitleConfigurationExtensionAttributes {
	if o == nil || IsNil(o.ExtensionAttributes) {
		var ret []PatchSoftwareTitleConfigurationExtensionAttributes
		return ret
	}
	return o.ExtensionAttributes
}

// GetExtensionAttributesOk returns a tuple with the ExtensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetExtensionAttributesOk() ([]PatchSoftwareTitleConfigurationExtensionAttributes, bool) {
	if o == nil || IsNil(o.ExtensionAttributes) {
		return nil, false
	}
	return o.ExtensionAttributes, true
}

// HasExtensionAttributes returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasExtensionAttributes() bool {
	if o != nil && !IsNil(o.ExtensionAttributes) {
		return true
	}

	return false
}

// SetExtensionAttributes gets a reference to the given []PatchSoftwareTitleConfigurationExtensionAttributes and assigns it to the ExtensionAttributes field.
func (o *PatchSoftwareTitleConfiguration) SetExtensionAttributes(v []PatchSoftwareTitleConfigurationExtensionAttributes) {
	o.ExtensionAttributes = v
}

// GetSoftwareTitleName returns the SoftwareTitleName field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleName() string {
	if o == nil || IsNil(o.SoftwareTitleName) {
		var ret string
		return ret
	}
	return *o.SoftwareTitleName
}

// GetSoftwareTitleNameOk returns a tuple with the SoftwareTitleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleNameOk() (*string, bool) {
	if o == nil || IsNil(o.SoftwareTitleName) {
		return nil, false
	}
	return o.SoftwareTitleName, true
}

// HasSoftwareTitleName returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasSoftwareTitleName() bool {
	if o != nil && !IsNil(o.SoftwareTitleName) {
		return true
	}

	return false
}

// SetSoftwareTitleName gets a reference to the given string and assigns it to the SoftwareTitleName field.
func (o *PatchSoftwareTitleConfiguration) SetSoftwareTitleName(v string) {
	o.SoftwareTitleName = &v
}

// GetSoftwareTitleNameId returns the SoftwareTitleNameId field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleNameId() string {
	if o == nil || IsNil(o.SoftwareTitleNameId) {
		var ret string
		return ret
	}
	return *o.SoftwareTitleNameId
}

// GetSoftwareTitleNameIdOk returns a tuple with the SoftwareTitleNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitleNameIdOk() (*string, bool) {
	if o == nil || IsNil(o.SoftwareTitleNameId) {
		return nil, false
	}
	return o.SoftwareTitleNameId, true
}

// HasSoftwareTitleNameId returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasSoftwareTitleNameId() bool {
	if o != nil && !IsNil(o.SoftwareTitleNameId) {
		return true
	}

	return false
}

// SetSoftwareTitleNameId gets a reference to the given string and assigns it to the SoftwareTitleNameId field.
func (o *PatchSoftwareTitleConfiguration) SetSoftwareTitleNameId(v string) {
	o.SoftwareTitleNameId = &v
}

// GetSoftwareTitlePublisher returns the SoftwareTitlePublisher field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitlePublisher() string {
	if o == nil || IsNil(o.SoftwareTitlePublisher) {
		var ret string
		return ret
	}
	return *o.SoftwareTitlePublisher
}

// GetSoftwareTitlePublisherOk returns a tuple with the SoftwareTitlePublisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetSoftwareTitlePublisherOk() (*string, bool) {
	if o == nil || IsNil(o.SoftwareTitlePublisher) {
		return nil, false
	}
	return o.SoftwareTitlePublisher, true
}

// HasSoftwareTitlePublisher returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasSoftwareTitlePublisher() bool {
	if o != nil && !IsNil(o.SoftwareTitlePublisher) {
		return true
	}

	return false
}

// SetSoftwareTitlePublisher gets a reference to the given string and assigns it to the SoftwareTitlePublisher field.
func (o *PatchSoftwareTitleConfiguration) SetSoftwareTitlePublisher(v string) {
	o.SoftwareTitlePublisher = &v
}

// GetPatchSourceName returns the PatchSourceName field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetPatchSourceName() string {
	if o == nil || IsNil(o.PatchSourceName) {
		var ret string
		return ret
	}
	return *o.PatchSourceName
}

// GetPatchSourceNameOk returns a tuple with the PatchSourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetPatchSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.PatchSourceName) {
		return nil, false
	}
	return o.PatchSourceName, true
}

// HasPatchSourceName returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasPatchSourceName() bool {
	if o != nil && !IsNil(o.PatchSourceName) {
		return true
	}

	return false
}

// SetPatchSourceName gets a reference to the given string and assigns it to the PatchSourceName field.
func (o *PatchSoftwareTitleConfiguration) SetPatchSourceName(v string) {
	o.PatchSourceName = &v
}

// GetPatchSourceEnabled returns the PatchSourceEnabled field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetPatchSourceEnabled() bool {
	if o == nil || IsNil(o.PatchSourceEnabled) {
		var ret bool
		return ret
	}
	return *o.PatchSourceEnabled
}

// GetPatchSourceEnabledOk returns a tuple with the PatchSourceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetPatchSourceEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PatchSourceEnabled) {
		return nil, false
	}
	return o.PatchSourceEnabled, true
}

// HasPatchSourceEnabled returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasPatchSourceEnabled() bool {
	if o != nil && !IsNil(o.PatchSourceEnabled) {
		return true
	}

	return false
}

// SetPatchSourceEnabled gets a reference to the given bool and assigns it to the PatchSourceEnabled field.
func (o *PatchSoftwareTitleConfiguration) SetPatchSourceEnabled(v bool) {
	o.PatchSourceEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchSoftwareTitleConfiguration) SetId(v string) {
	o.Id = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *PatchSoftwareTitleConfiguration) GetPackages() []PatchSoftwareTitlePackages {
	if o == nil || IsNil(o.Packages) {
		var ret []PatchSoftwareTitlePackages
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchSoftwareTitleConfiguration) GetPackagesOk() ([]PatchSoftwareTitlePackages, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *PatchSoftwareTitleConfiguration) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []PatchSoftwareTitlePackages and assigns it to the Packages field.
func (o *PatchSoftwareTitleConfiguration) SetPackages(v []PatchSoftwareTitlePackages) {
	o.Packages = v
}

func (o PatchSoftwareTitleConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchSoftwareTitleConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.CategoryId) {
		toSerialize["categoryId"] = o.CategoryId
	}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	if !IsNil(o.UiNotifications) {
		toSerialize["uiNotifications"] = o.UiNotifications
	}
	if !IsNil(o.EmailNotifications) {
		toSerialize["emailNotifications"] = o.EmailNotifications
	}
	toSerialize["softwareTitleId"] = o.SoftwareTitleId
	if !IsNil(o.JamfOfficial) {
		toSerialize["jamfOfficial"] = o.JamfOfficial
	}
	if !IsNil(o.ExtensionAttributes) {
		toSerialize["extensionAttributes"] = o.ExtensionAttributes
	}
	if !IsNil(o.SoftwareTitleName) {
		toSerialize["softwareTitleName"] = o.SoftwareTitleName
	}
	if !IsNil(o.SoftwareTitleNameId) {
		toSerialize["softwareTitleNameId"] = o.SoftwareTitleNameId
	}
	if !IsNil(o.SoftwareTitlePublisher) {
		toSerialize["softwareTitlePublisher"] = o.SoftwareTitlePublisher
	}
	if !IsNil(o.PatchSourceName) {
		toSerialize["patchSourceName"] = o.PatchSourceName
	}
	if !IsNil(o.PatchSourceEnabled) {
		toSerialize["patchSourceEnabled"] = o.PatchSourceEnabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Packages) {
		toSerialize["packages"] = o.Packages
	}
	return toSerialize, nil
}

func (o *PatchSoftwareTitleConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"softwareTitleId",
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

	varPatchSoftwareTitleConfiguration := _PatchSoftwareTitleConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatchSoftwareTitleConfiguration)

	if err != nil {
		return err
	}

	*o = PatchSoftwareTitleConfiguration(varPatchSoftwareTitleConfiguration)

	return err
}

type NullablePatchSoftwareTitleConfiguration struct {
	value *PatchSoftwareTitleConfiguration
	isSet bool
}

func (v NullablePatchSoftwareTitleConfiguration) Get() *PatchSoftwareTitleConfiguration {
	return v.value
}

func (v *NullablePatchSoftwareTitleConfiguration) Set(val *PatchSoftwareTitleConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchSoftwareTitleConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchSoftwareTitleConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchSoftwareTitleConfiguration(val *PatchSoftwareTitleConfiguration) *NullablePatchSoftwareTitleConfiguration {
	return &NullablePatchSoftwareTitleConfiguration{value: val, isSet: true}
}

func (v NullablePatchSoftwareTitleConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchSoftwareTitleConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

