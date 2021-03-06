/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TopApplication Top application that contributes the most to the traffic.
type TopApplication struct {
	// String providing an application identifier.
	AppId *string `json:"appId,omitempty"`
	IpTrafficFilter *FlowInfo `json:"ipTrafficFilter,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio *int32 `json:"ratio,omitempty"`
}

// NewTopApplication instantiates a new TopApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopApplication() *TopApplication {
	this := TopApplication{}
	return &this
}

// NewTopApplicationWithDefaults instantiates a new TopApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopApplicationWithDefaults() *TopApplication {
	this := TopApplication{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *TopApplication) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopApplication) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *TopApplication) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *TopApplication) SetAppId(v string) {
	o.AppId = &v
}

// GetIpTrafficFilter returns the IpTrafficFilter field value if set, zero value otherwise.
func (o *TopApplication) GetIpTrafficFilter() FlowInfo {
	if o == nil || o.IpTrafficFilter == nil {
		var ret FlowInfo
		return ret
	}
	return *o.IpTrafficFilter
}

// GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopApplication) GetIpTrafficFilterOk() (*FlowInfo, bool) {
	if o == nil || o.IpTrafficFilter == nil {
		return nil, false
	}
	return o.IpTrafficFilter, true
}

// HasIpTrafficFilter returns a boolean if a field has been set.
func (o *TopApplication) HasIpTrafficFilter() bool {
	if o != nil && o.IpTrafficFilter != nil {
		return true
	}

	return false
}

// SetIpTrafficFilter gets a reference to the given FlowInfo and assigns it to the IpTrafficFilter field.
func (o *TopApplication) SetIpTrafficFilter(v FlowInfo) {
	o.IpTrafficFilter = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *TopApplication) GetRatio() int32 {
	if o == nil || o.Ratio == nil {
		var ret int32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopApplication) GetRatioOk() (*int32, bool) {
	if o == nil || o.Ratio == nil {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *TopApplication) HasRatio() bool {
	if o != nil && o.Ratio != nil {
		return true
	}

	return false
}

// SetRatio gets a reference to the given int32 and assigns it to the Ratio field.
func (o *TopApplication) SetRatio(v int32) {
	o.Ratio = &v
}

func (o TopApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.IpTrafficFilter != nil {
		toSerialize["ipTrafficFilter"] = o.IpTrafficFilter
	}
	if o.Ratio != nil {
		toSerialize["ratio"] = o.Ratio
	}
	return json.Marshal(toSerialize)
}

type NullableTopApplication struct {
	value *TopApplication
	isSet bool
}

func (v NullableTopApplication) Get() *TopApplication {
	return v.value
}

func (v *NullableTopApplication) Set(val *TopApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableTopApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableTopApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopApplication(val *TopApplication) *NullableTopApplication {
	return &NullableTopApplication{value: val, isSet: true}
}

func (v NullableTopApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


