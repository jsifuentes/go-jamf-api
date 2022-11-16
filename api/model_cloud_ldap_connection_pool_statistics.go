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

// CloudLdapConnectionPoolStatistics Ldap Cloud Identity Provider conection pool statistics
type CloudLdapConnectionPoolStatistics struct {
	NumConnectionsClosedDefunct *int64 `json:"numConnectionsClosedDefunct,omitempty"`
	NumConnectionsClosedExpired *int64 `json:"numConnectionsClosedExpired,omitempty"`
	NumConnectionsClosedUnneeded *int64 `json:"numConnectionsClosedUnneeded,omitempty"`
	NumFailedCheckouts *int64 `json:"numFailedCheckouts,omitempty"`
	NumFailedConnectionAttempts *int64 `json:"numFailedConnectionAttempts,omitempty"`
	NumReleasedValid *int64 `json:"numReleasedValid,omitempty"`
	NumSuccessfulCheckouts *int64 `json:"numSuccessfulCheckouts,omitempty"`
	NumSuccessfulCheckoutsNewConnection *int64 `json:"numSuccessfulCheckoutsNewConnection,omitempty"`
	NumSuccessfulConnectionAttempts *int64 `json:"numSuccessfulConnectionAttempts,omitempty"`
	MaximumAvailableConnections *int64 `json:"maximumAvailableConnections,omitempty"`
	NumSuccessfulCheckoutsWithoutWaiting *int64 `json:"numSuccessfulCheckoutsWithoutWaiting,omitempty"`
	NumSuccessfulCheckoutsAfterWaiting *int64 `json:"numSuccessfulCheckoutsAfterWaiting,omitempty"`
	NumAvailableConnections *int64 `json:"numAvailableConnections,omitempty"`
}

// NewCloudLdapConnectionPoolStatistics instantiates a new CloudLdapConnectionPoolStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudLdapConnectionPoolStatistics() *CloudLdapConnectionPoolStatistics {
	this := CloudLdapConnectionPoolStatistics{}
	return &this
}

// NewCloudLdapConnectionPoolStatisticsWithDefaults instantiates a new CloudLdapConnectionPoolStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudLdapConnectionPoolStatisticsWithDefaults() *CloudLdapConnectionPoolStatistics {
	this := CloudLdapConnectionPoolStatistics{}
	return &this
}

// GetNumConnectionsClosedDefunct returns the NumConnectionsClosedDefunct field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedDefunct() int64 {
	if o == nil || o.NumConnectionsClosedDefunct == nil {
		var ret int64
		return ret
	}
	return *o.NumConnectionsClosedDefunct
}

// GetNumConnectionsClosedDefunctOk returns a tuple with the NumConnectionsClosedDefunct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedDefunctOk() (*int64, bool) {
	if o == nil || o.NumConnectionsClosedDefunct == nil {
		return nil, false
	}
	return o.NumConnectionsClosedDefunct, true
}

// HasNumConnectionsClosedDefunct returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumConnectionsClosedDefunct() bool {
	if o != nil && o.NumConnectionsClosedDefunct != nil {
		return true
	}

	return false
}

// SetNumConnectionsClosedDefunct gets a reference to the given int64 and assigns it to the NumConnectionsClosedDefunct field.
func (o *CloudLdapConnectionPoolStatistics) SetNumConnectionsClosedDefunct(v int64) {
	o.NumConnectionsClosedDefunct = &v
}

// GetNumConnectionsClosedExpired returns the NumConnectionsClosedExpired field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedExpired() int64 {
	if o == nil || o.NumConnectionsClosedExpired == nil {
		var ret int64
		return ret
	}
	return *o.NumConnectionsClosedExpired
}

// GetNumConnectionsClosedExpiredOk returns a tuple with the NumConnectionsClosedExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedExpiredOk() (*int64, bool) {
	if o == nil || o.NumConnectionsClosedExpired == nil {
		return nil, false
	}
	return o.NumConnectionsClosedExpired, true
}

// HasNumConnectionsClosedExpired returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumConnectionsClosedExpired() bool {
	if o != nil && o.NumConnectionsClosedExpired != nil {
		return true
	}

	return false
}

// SetNumConnectionsClosedExpired gets a reference to the given int64 and assigns it to the NumConnectionsClosedExpired field.
func (o *CloudLdapConnectionPoolStatistics) SetNumConnectionsClosedExpired(v int64) {
	o.NumConnectionsClosedExpired = &v
}

// GetNumConnectionsClosedUnneeded returns the NumConnectionsClosedUnneeded field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedUnneeded() int64 {
	if o == nil || o.NumConnectionsClosedUnneeded == nil {
		var ret int64
		return ret
	}
	return *o.NumConnectionsClosedUnneeded
}

// GetNumConnectionsClosedUnneededOk returns a tuple with the NumConnectionsClosedUnneeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumConnectionsClosedUnneededOk() (*int64, bool) {
	if o == nil || o.NumConnectionsClosedUnneeded == nil {
		return nil, false
	}
	return o.NumConnectionsClosedUnneeded, true
}

// HasNumConnectionsClosedUnneeded returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumConnectionsClosedUnneeded() bool {
	if o != nil && o.NumConnectionsClosedUnneeded != nil {
		return true
	}

	return false
}

// SetNumConnectionsClosedUnneeded gets a reference to the given int64 and assigns it to the NumConnectionsClosedUnneeded field.
func (o *CloudLdapConnectionPoolStatistics) SetNumConnectionsClosedUnneeded(v int64) {
	o.NumConnectionsClosedUnneeded = &v
}

// GetNumFailedCheckouts returns the NumFailedCheckouts field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumFailedCheckouts() int64 {
	if o == nil || o.NumFailedCheckouts == nil {
		var ret int64
		return ret
	}
	return *o.NumFailedCheckouts
}

// GetNumFailedCheckoutsOk returns a tuple with the NumFailedCheckouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumFailedCheckoutsOk() (*int64, bool) {
	if o == nil || o.NumFailedCheckouts == nil {
		return nil, false
	}
	return o.NumFailedCheckouts, true
}

// HasNumFailedCheckouts returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumFailedCheckouts() bool {
	if o != nil && o.NumFailedCheckouts != nil {
		return true
	}

	return false
}

// SetNumFailedCheckouts gets a reference to the given int64 and assigns it to the NumFailedCheckouts field.
func (o *CloudLdapConnectionPoolStatistics) SetNumFailedCheckouts(v int64) {
	o.NumFailedCheckouts = &v
}

// GetNumFailedConnectionAttempts returns the NumFailedConnectionAttempts field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumFailedConnectionAttempts() int64 {
	if o == nil || o.NumFailedConnectionAttempts == nil {
		var ret int64
		return ret
	}
	return *o.NumFailedConnectionAttempts
}

// GetNumFailedConnectionAttemptsOk returns a tuple with the NumFailedConnectionAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumFailedConnectionAttemptsOk() (*int64, bool) {
	if o == nil || o.NumFailedConnectionAttempts == nil {
		return nil, false
	}
	return o.NumFailedConnectionAttempts, true
}

// HasNumFailedConnectionAttempts returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumFailedConnectionAttempts() bool {
	if o != nil && o.NumFailedConnectionAttempts != nil {
		return true
	}

	return false
}

// SetNumFailedConnectionAttempts gets a reference to the given int64 and assigns it to the NumFailedConnectionAttempts field.
func (o *CloudLdapConnectionPoolStatistics) SetNumFailedConnectionAttempts(v int64) {
	o.NumFailedConnectionAttempts = &v
}

// GetNumReleasedValid returns the NumReleasedValid field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumReleasedValid() int64 {
	if o == nil || o.NumReleasedValid == nil {
		var ret int64
		return ret
	}
	return *o.NumReleasedValid
}

// GetNumReleasedValidOk returns a tuple with the NumReleasedValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumReleasedValidOk() (*int64, bool) {
	if o == nil || o.NumReleasedValid == nil {
		return nil, false
	}
	return o.NumReleasedValid, true
}

// HasNumReleasedValid returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumReleasedValid() bool {
	if o != nil && o.NumReleasedValid != nil {
		return true
	}

	return false
}

// SetNumReleasedValid gets a reference to the given int64 and assigns it to the NumReleasedValid field.
func (o *CloudLdapConnectionPoolStatistics) SetNumReleasedValid(v int64) {
	o.NumReleasedValid = &v
}

// GetNumSuccessfulCheckouts returns the NumSuccessfulCheckouts field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckouts() int64 {
	if o == nil || o.NumSuccessfulCheckouts == nil {
		var ret int64
		return ret
	}
	return *o.NumSuccessfulCheckouts
}

// GetNumSuccessfulCheckoutsOk returns a tuple with the NumSuccessfulCheckouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsOk() (*int64, bool) {
	if o == nil || o.NumSuccessfulCheckouts == nil {
		return nil, false
	}
	return o.NumSuccessfulCheckouts, true
}

// HasNumSuccessfulCheckouts returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumSuccessfulCheckouts() bool {
	if o != nil && o.NumSuccessfulCheckouts != nil {
		return true
	}

	return false
}

// SetNumSuccessfulCheckouts gets a reference to the given int64 and assigns it to the NumSuccessfulCheckouts field.
func (o *CloudLdapConnectionPoolStatistics) SetNumSuccessfulCheckouts(v int64) {
	o.NumSuccessfulCheckouts = &v
}

// GetNumSuccessfulCheckoutsNewConnection returns the NumSuccessfulCheckoutsNewConnection field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsNewConnection() int64 {
	if o == nil || o.NumSuccessfulCheckoutsNewConnection == nil {
		var ret int64
		return ret
	}
	return *o.NumSuccessfulCheckoutsNewConnection
}

// GetNumSuccessfulCheckoutsNewConnectionOk returns a tuple with the NumSuccessfulCheckoutsNewConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsNewConnectionOk() (*int64, bool) {
	if o == nil || o.NumSuccessfulCheckoutsNewConnection == nil {
		return nil, false
	}
	return o.NumSuccessfulCheckoutsNewConnection, true
}

// HasNumSuccessfulCheckoutsNewConnection returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumSuccessfulCheckoutsNewConnection() bool {
	if o != nil && o.NumSuccessfulCheckoutsNewConnection != nil {
		return true
	}

	return false
}

// SetNumSuccessfulCheckoutsNewConnection gets a reference to the given int64 and assigns it to the NumSuccessfulCheckoutsNewConnection field.
func (o *CloudLdapConnectionPoolStatistics) SetNumSuccessfulCheckoutsNewConnection(v int64) {
	o.NumSuccessfulCheckoutsNewConnection = &v
}

// GetNumSuccessfulConnectionAttempts returns the NumSuccessfulConnectionAttempts field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulConnectionAttempts() int64 {
	if o == nil || o.NumSuccessfulConnectionAttempts == nil {
		var ret int64
		return ret
	}
	return *o.NumSuccessfulConnectionAttempts
}

// GetNumSuccessfulConnectionAttemptsOk returns a tuple with the NumSuccessfulConnectionAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulConnectionAttemptsOk() (*int64, bool) {
	if o == nil || o.NumSuccessfulConnectionAttempts == nil {
		return nil, false
	}
	return o.NumSuccessfulConnectionAttempts, true
}

// HasNumSuccessfulConnectionAttempts returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumSuccessfulConnectionAttempts() bool {
	if o != nil && o.NumSuccessfulConnectionAttempts != nil {
		return true
	}

	return false
}

// SetNumSuccessfulConnectionAttempts gets a reference to the given int64 and assigns it to the NumSuccessfulConnectionAttempts field.
func (o *CloudLdapConnectionPoolStatistics) SetNumSuccessfulConnectionAttempts(v int64) {
	o.NumSuccessfulConnectionAttempts = &v
}

// GetMaximumAvailableConnections returns the MaximumAvailableConnections field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetMaximumAvailableConnections() int64 {
	if o == nil || o.MaximumAvailableConnections == nil {
		var ret int64
		return ret
	}
	return *o.MaximumAvailableConnections
}

// GetMaximumAvailableConnectionsOk returns a tuple with the MaximumAvailableConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetMaximumAvailableConnectionsOk() (*int64, bool) {
	if o == nil || o.MaximumAvailableConnections == nil {
		return nil, false
	}
	return o.MaximumAvailableConnections, true
}

// HasMaximumAvailableConnections returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasMaximumAvailableConnections() bool {
	if o != nil && o.MaximumAvailableConnections != nil {
		return true
	}

	return false
}

// SetMaximumAvailableConnections gets a reference to the given int64 and assigns it to the MaximumAvailableConnections field.
func (o *CloudLdapConnectionPoolStatistics) SetMaximumAvailableConnections(v int64) {
	o.MaximumAvailableConnections = &v
}

// GetNumSuccessfulCheckoutsWithoutWaiting returns the NumSuccessfulCheckoutsWithoutWaiting field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsWithoutWaiting() int64 {
	if o == nil || o.NumSuccessfulCheckoutsWithoutWaiting == nil {
		var ret int64
		return ret
	}
	return *o.NumSuccessfulCheckoutsWithoutWaiting
}

// GetNumSuccessfulCheckoutsWithoutWaitingOk returns a tuple with the NumSuccessfulCheckoutsWithoutWaiting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsWithoutWaitingOk() (*int64, bool) {
	if o == nil || o.NumSuccessfulCheckoutsWithoutWaiting == nil {
		return nil, false
	}
	return o.NumSuccessfulCheckoutsWithoutWaiting, true
}

// HasNumSuccessfulCheckoutsWithoutWaiting returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumSuccessfulCheckoutsWithoutWaiting() bool {
	if o != nil && o.NumSuccessfulCheckoutsWithoutWaiting != nil {
		return true
	}

	return false
}

// SetNumSuccessfulCheckoutsWithoutWaiting gets a reference to the given int64 and assigns it to the NumSuccessfulCheckoutsWithoutWaiting field.
func (o *CloudLdapConnectionPoolStatistics) SetNumSuccessfulCheckoutsWithoutWaiting(v int64) {
	o.NumSuccessfulCheckoutsWithoutWaiting = &v
}

// GetNumSuccessfulCheckoutsAfterWaiting returns the NumSuccessfulCheckoutsAfterWaiting field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsAfterWaiting() int64 {
	if o == nil || o.NumSuccessfulCheckoutsAfterWaiting == nil {
		var ret int64
		return ret
	}
	return *o.NumSuccessfulCheckoutsAfterWaiting
}

// GetNumSuccessfulCheckoutsAfterWaitingOk returns a tuple with the NumSuccessfulCheckoutsAfterWaiting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumSuccessfulCheckoutsAfterWaitingOk() (*int64, bool) {
	if o == nil || o.NumSuccessfulCheckoutsAfterWaiting == nil {
		return nil, false
	}
	return o.NumSuccessfulCheckoutsAfterWaiting, true
}

// HasNumSuccessfulCheckoutsAfterWaiting returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumSuccessfulCheckoutsAfterWaiting() bool {
	if o != nil && o.NumSuccessfulCheckoutsAfterWaiting != nil {
		return true
	}

	return false
}

// SetNumSuccessfulCheckoutsAfterWaiting gets a reference to the given int64 and assigns it to the NumSuccessfulCheckoutsAfterWaiting field.
func (o *CloudLdapConnectionPoolStatistics) SetNumSuccessfulCheckoutsAfterWaiting(v int64) {
	o.NumSuccessfulCheckoutsAfterWaiting = &v
}

// GetNumAvailableConnections returns the NumAvailableConnections field value if set, zero value otherwise.
func (o *CloudLdapConnectionPoolStatistics) GetNumAvailableConnections() int64 {
	if o == nil || o.NumAvailableConnections == nil {
		var ret int64
		return ret
	}
	return *o.NumAvailableConnections
}

// GetNumAvailableConnectionsOk returns a tuple with the NumAvailableConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudLdapConnectionPoolStatistics) GetNumAvailableConnectionsOk() (*int64, bool) {
	if o == nil || o.NumAvailableConnections == nil {
		return nil, false
	}
	return o.NumAvailableConnections, true
}

// HasNumAvailableConnections returns a boolean if a field has been set.
func (o *CloudLdapConnectionPoolStatistics) HasNumAvailableConnections() bool {
	if o != nil && o.NumAvailableConnections != nil {
		return true
	}

	return false
}

// SetNumAvailableConnections gets a reference to the given int64 and assigns it to the NumAvailableConnections field.
func (o *CloudLdapConnectionPoolStatistics) SetNumAvailableConnections(v int64) {
	o.NumAvailableConnections = &v
}

func (o CloudLdapConnectionPoolStatistics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumConnectionsClosedDefunct != nil {
		toSerialize["numConnectionsClosedDefunct"] = o.NumConnectionsClosedDefunct
	}
	if o.NumConnectionsClosedExpired != nil {
		toSerialize["numConnectionsClosedExpired"] = o.NumConnectionsClosedExpired
	}
	if o.NumConnectionsClosedUnneeded != nil {
		toSerialize["numConnectionsClosedUnneeded"] = o.NumConnectionsClosedUnneeded
	}
	if o.NumFailedCheckouts != nil {
		toSerialize["numFailedCheckouts"] = o.NumFailedCheckouts
	}
	if o.NumFailedConnectionAttempts != nil {
		toSerialize["numFailedConnectionAttempts"] = o.NumFailedConnectionAttempts
	}
	if o.NumReleasedValid != nil {
		toSerialize["numReleasedValid"] = o.NumReleasedValid
	}
	if o.NumSuccessfulCheckouts != nil {
		toSerialize["numSuccessfulCheckouts"] = o.NumSuccessfulCheckouts
	}
	if o.NumSuccessfulCheckoutsNewConnection != nil {
		toSerialize["numSuccessfulCheckoutsNewConnection"] = o.NumSuccessfulCheckoutsNewConnection
	}
	if o.NumSuccessfulConnectionAttempts != nil {
		toSerialize["numSuccessfulConnectionAttempts"] = o.NumSuccessfulConnectionAttempts
	}
	if o.MaximumAvailableConnections != nil {
		toSerialize["maximumAvailableConnections"] = o.MaximumAvailableConnections
	}
	if o.NumSuccessfulCheckoutsWithoutWaiting != nil {
		toSerialize["numSuccessfulCheckoutsWithoutWaiting"] = o.NumSuccessfulCheckoutsWithoutWaiting
	}
	if o.NumSuccessfulCheckoutsAfterWaiting != nil {
		toSerialize["numSuccessfulCheckoutsAfterWaiting"] = o.NumSuccessfulCheckoutsAfterWaiting
	}
	if o.NumAvailableConnections != nil {
		toSerialize["numAvailableConnections"] = o.NumAvailableConnections
	}
	return json.Marshal(toSerialize)
}

type NullableCloudLdapConnectionPoolStatistics struct {
	value *CloudLdapConnectionPoolStatistics
	isSet bool
}

func (v NullableCloudLdapConnectionPoolStatistics) Get() *CloudLdapConnectionPoolStatistics {
	return v.value
}

func (v *NullableCloudLdapConnectionPoolStatistics) Set(val *CloudLdapConnectionPoolStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudLdapConnectionPoolStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudLdapConnectionPoolStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudLdapConnectionPoolStatistics(val *CloudLdapConnectionPoolStatistics) *NullableCloudLdapConnectionPoolStatistics {
	return &NullableCloudLdapConnectionPoolStatistics{value: val, isSet: true}
}

func (v NullableCloudLdapConnectionPoolStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudLdapConnectionPoolStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

