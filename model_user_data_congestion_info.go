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

// UserDataCongestionInfo Represents the user data congestion information.
type UserDataCongestionInfo struct {
	NetworkArea *NetworkAreaInfo `json:"networkArea,omitempty"`
	CongestionInfo *CongestionInfo `json:"congestionInfo,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
}

// NewUserDataCongestionInfo instantiates a new UserDataCongestionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataCongestionInfo() *UserDataCongestionInfo {
	this := UserDataCongestionInfo{}
	return &this
}

// NewUserDataCongestionInfoWithDefaults instantiates a new UserDataCongestionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataCongestionInfoWithDefaults() *UserDataCongestionInfo {
	this := UserDataCongestionInfo{}
	return &this
}

// GetNetworkArea returns the NetworkArea field value if set, zero value otherwise.
func (o *UserDataCongestionInfo) GetNetworkArea() NetworkAreaInfo {
	if o == nil || o.NetworkArea == nil {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.NetworkArea
}

// GetNetworkAreaOk returns a tuple with the NetworkArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataCongestionInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil || o.NetworkArea == nil {
		return nil, false
	}
	return o.NetworkArea, true
}

// HasNetworkArea returns a boolean if a field has been set.
func (o *UserDataCongestionInfo) HasNetworkArea() bool {
	if o != nil && o.NetworkArea != nil {
		return true
	}

	return false
}

// SetNetworkArea gets a reference to the given NetworkAreaInfo and assigns it to the NetworkArea field.
func (o *UserDataCongestionInfo) SetNetworkArea(v NetworkAreaInfo) {
	o.NetworkArea = &v
}

// GetCongestionInfo returns the CongestionInfo field value if set, zero value otherwise.
func (o *UserDataCongestionInfo) GetCongestionInfo() CongestionInfo {
	if o == nil || o.CongestionInfo == nil {
		var ret CongestionInfo
		return ret
	}
	return *o.CongestionInfo
}

// GetCongestionInfoOk returns a tuple with the CongestionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataCongestionInfo) GetCongestionInfoOk() (*CongestionInfo, bool) {
	if o == nil || o.CongestionInfo == nil {
		return nil, false
	}
	return o.CongestionInfo, true
}

// HasCongestionInfo returns a boolean if a field has been set.
func (o *UserDataCongestionInfo) HasCongestionInfo() bool {
	if o != nil && o.CongestionInfo != nil {
		return true
	}

	return false
}

// SetCongestionInfo gets a reference to the given CongestionInfo and assigns it to the CongestionInfo field.
func (o *UserDataCongestionInfo) SetCongestionInfo(v CongestionInfo) {
	o.CongestionInfo = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *UserDataCongestionInfo) GetSnssai() Snssai {
	if o == nil || o.Snssai == nil {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataCongestionInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || o.Snssai == nil {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *UserDataCongestionInfo) HasSnssai() bool {
	if o != nil && o.Snssai != nil {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *UserDataCongestionInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

func (o UserDataCongestionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkArea != nil {
		toSerialize["networkArea"] = o.NetworkArea
	}
	if o.CongestionInfo != nil {
		toSerialize["congestionInfo"] = o.CongestionInfo
	}
	if o.Snssai != nil {
		toSerialize["snssai"] = o.Snssai
	}
	return json.Marshal(toSerialize)
}

type NullableUserDataCongestionInfo struct {
	value *UserDataCongestionInfo
	isSet bool
}

func (v NullableUserDataCongestionInfo) Get() *UserDataCongestionInfo {
	return v.value
}

func (v *NullableUserDataCongestionInfo) Set(val *UserDataCongestionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataCongestionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataCongestionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataCongestionInfo(val *UserDataCongestionInfo) *NullableUserDataCongestionInfo {
	return &NullableUserDataCongestionInfo{value: val, isSet: true}
}

func (v NullableUserDataCongestionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataCongestionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


