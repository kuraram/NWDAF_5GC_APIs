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

// NetworkPerfInfo Represents the network performance information.
type NetworkPerfInfo struct {
	NetworkArea NetworkAreaInfo `json:"networkArea"`
	NwPerfType NetworkPerfType `json:"nwPerfType"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	RelativeRatio *int32 `json:"relativeRatio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AbsoluteNum *int32 `json:"absoluteNum,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
}

// NewNetworkPerfInfo instantiates a new NetworkPerfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPerfInfo(networkArea NetworkAreaInfo, nwPerfType NetworkPerfType) *NetworkPerfInfo {
	this := NetworkPerfInfo{}
	this.NetworkArea = networkArea
	this.NwPerfType = nwPerfType
	return &this
}

// NewNetworkPerfInfoWithDefaults instantiates a new NetworkPerfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPerfInfoWithDefaults() *NetworkPerfInfo {
	this := NetworkPerfInfo{}
	return &this
}

// GetNetworkArea returns the NetworkArea field value
func (o *NetworkPerfInfo) GetNetworkArea() NetworkAreaInfo {
	if o == nil {
		var ret NetworkAreaInfo
		return ret
	}

	return o.NetworkArea
}

// GetNetworkAreaOk returns a tuple with the NetworkArea field value
// and a boolean to check if the value has been set.
func (o *NetworkPerfInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkArea, true
}

// SetNetworkArea sets field value
func (o *NetworkPerfInfo) SetNetworkArea(v NetworkAreaInfo) {
	o.NetworkArea = v
}

// GetNwPerfType returns the NwPerfType field value
func (o *NetworkPerfInfo) GetNwPerfType() NetworkPerfType {
	if o == nil {
		var ret NetworkPerfType
		return ret
	}

	return o.NwPerfType
}

// GetNwPerfTypeOk returns a tuple with the NwPerfType field value
// and a boolean to check if the value has been set.
func (o *NetworkPerfInfo) GetNwPerfTypeOk() (*NetworkPerfType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NwPerfType, true
}

// SetNwPerfType sets field value
func (o *NetworkPerfInfo) SetNwPerfType(v NetworkPerfType) {
	o.NwPerfType = v
}

// GetRelativeRatio returns the RelativeRatio field value if set, zero value otherwise.
func (o *NetworkPerfInfo) GetRelativeRatio() int32 {
	if o == nil || o.RelativeRatio == nil {
		var ret int32
		return ret
	}
	return *o.RelativeRatio
}

// GetRelativeRatioOk returns a tuple with the RelativeRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPerfInfo) GetRelativeRatioOk() (*int32, bool) {
	if o == nil || o.RelativeRatio == nil {
		return nil, false
	}
	return o.RelativeRatio, true
}

// HasRelativeRatio returns a boolean if a field has been set.
func (o *NetworkPerfInfo) HasRelativeRatio() bool {
	if o != nil && o.RelativeRatio != nil {
		return true
	}

	return false
}

// SetRelativeRatio gets a reference to the given int32 and assigns it to the RelativeRatio field.
func (o *NetworkPerfInfo) SetRelativeRatio(v int32) {
	o.RelativeRatio = &v
}

// GetAbsoluteNum returns the AbsoluteNum field value if set, zero value otherwise.
func (o *NetworkPerfInfo) GetAbsoluteNum() int32 {
	if o == nil || o.AbsoluteNum == nil {
		var ret int32
		return ret
	}
	return *o.AbsoluteNum
}

// GetAbsoluteNumOk returns a tuple with the AbsoluteNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPerfInfo) GetAbsoluteNumOk() (*int32, bool) {
	if o == nil || o.AbsoluteNum == nil {
		return nil, false
	}
	return o.AbsoluteNum, true
}

// HasAbsoluteNum returns a boolean if a field has been set.
func (o *NetworkPerfInfo) HasAbsoluteNum() bool {
	if o != nil && o.AbsoluteNum != nil {
		return true
	}

	return false
}

// SetAbsoluteNum gets a reference to the given int32 and assigns it to the AbsoluteNum field.
func (o *NetworkPerfInfo) SetAbsoluteNum(v int32) {
	o.AbsoluteNum = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *NetworkPerfInfo) GetConfidence() int32 {
	if o == nil || o.Confidence == nil {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPerfInfo) GetConfidenceOk() (*int32, bool) {
	if o == nil || o.Confidence == nil {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *NetworkPerfInfo) HasConfidence() bool {
	if o != nil && o.Confidence != nil {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *NetworkPerfInfo) SetConfidence(v int32) {
	o.Confidence = &v
}

func (o NetworkPerfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["networkArea"] = o.NetworkArea
	}
	if true {
		toSerialize["nwPerfType"] = o.NwPerfType
	}
	if o.RelativeRatio != nil {
		toSerialize["relativeRatio"] = o.RelativeRatio
	}
	if o.AbsoluteNum != nil {
		toSerialize["absoluteNum"] = o.AbsoluteNum
	}
	if o.Confidence != nil {
		toSerialize["confidence"] = o.Confidence
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkPerfInfo struct {
	value *NetworkPerfInfo
	isSet bool
}

func (v NullableNetworkPerfInfo) Get() *NetworkPerfInfo {
	return v.value
}

func (v *NullableNetworkPerfInfo) Set(val *NetworkPerfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPerfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPerfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPerfInfo(val *NetworkPerfInfo) *NullableNetworkPerfInfo {
	return &NullableNetworkPerfInfo{value: val, isSet: true}
}

func (v NullableNetworkPerfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPerfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


