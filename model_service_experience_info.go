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

// ServiceExperienceInfo Represents service experience information.
type ServiceExperienceInfo struct {
	SvcExprc SvcExperience `json:"svcExprc"`
	// string with format 'float' as defined in OpenAPI.
	SvcExprcVariance *float32 `json:"svcExprcVariance,omitempty"`
	Supis []string `json:"supis,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// String providing an application identifier.
	AppId *string `json:"appId,omitempty"`
	SrvExpcType *ServiceExperienceType `json:"srvExpcType,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \"Label1.Label2.Label3\").
	Dnn *string `json:"dnn,omitempty"`
	NetworkArea *NetworkAreaInfo `json:"networkArea,omitempty"`
	// Contains the Identifier of the selected Network Slice instance
	NsiId *string `json:"nsiId,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio *int32 `json:"ratio,omitempty"`
	RatType *RatType `json:"ratType,omitempty"`
	// Integer value indicating the ARFCN applicable for a downlink, uplink or bi-directional (TDD) NR global frequency raster, as definition of \"ARFCN-ValueNR\" IE in clause 6.3.2 of 3GPP TS 38.331. 
	Frequency *int32 `json:"frequency,omitempty"`
}

// NewServiceExperienceInfo instantiates a new ServiceExperienceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceExperienceInfo(svcExprc SvcExperience) *ServiceExperienceInfo {
	this := ServiceExperienceInfo{}
	this.SvcExprc = svcExprc
	return &this
}

// NewServiceExperienceInfoWithDefaults instantiates a new ServiceExperienceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceExperienceInfoWithDefaults() *ServiceExperienceInfo {
	this := ServiceExperienceInfo{}
	return &this
}

// GetSvcExprc returns the SvcExprc field value
func (o *ServiceExperienceInfo) GetSvcExprc() SvcExperience {
	if o == nil {
		var ret SvcExperience
		return ret
	}

	return o.SvcExprc
}

// GetSvcExprcOk returns a tuple with the SvcExprc field value
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSvcExprcOk() (*SvcExperience, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SvcExprc, true
}

// SetSvcExprc sets field value
func (o *ServiceExperienceInfo) SetSvcExprc(v SvcExperience) {
	o.SvcExprc = v
}

// GetSvcExprcVariance returns the SvcExprcVariance field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetSvcExprcVariance() float32 {
	if o == nil || o.SvcExprcVariance == nil {
		var ret float32
		return ret
	}
	return *o.SvcExprcVariance
}

// GetSvcExprcVarianceOk returns a tuple with the SvcExprcVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSvcExprcVarianceOk() (*float32, bool) {
	if o == nil || o.SvcExprcVariance == nil {
		return nil, false
	}
	return o.SvcExprcVariance, true
}

// HasSvcExprcVariance returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasSvcExprcVariance() bool {
	if o != nil && o.SvcExprcVariance != nil {
		return true
	}

	return false
}

// SetSvcExprcVariance gets a reference to the given float32 and assigns it to the SvcExprcVariance field.
func (o *ServiceExperienceInfo) SetSvcExprcVariance(v float32) {
	o.SvcExprcVariance = &v
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetSupis() []string {
	if o == nil || o.Supis == nil {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSupisOk() ([]string, bool) {
	if o == nil || o.Supis == nil {
		return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasSupis() bool {
	if o != nil && o.Supis != nil {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *ServiceExperienceInfo) SetSupis(v []string) {
	o.Supis = v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetSnssai() Snssai {
	if o == nil || o.Snssai == nil {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || o.Snssai == nil {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasSnssai() bool {
	if o != nil && o.Snssai != nil {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *ServiceExperienceInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ServiceExperienceInfo) SetAppId(v string) {
	o.AppId = &v
}

// GetSrvExpcType returns the SrvExpcType field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetSrvExpcType() ServiceExperienceType {
	if o == nil || o.SrvExpcType == nil {
		var ret ServiceExperienceType
		return ret
	}
	return *o.SrvExpcType
}

// GetSrvExpcTypeOk returns a tuple with the SrvExpcType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetSrvExpcTypeOk() (*ServiceExperienceType, bool) {
	if o == nil || o.SrvExpcType == nil {
		return nil, false
	}
	return o.SrvExpcType, true
}

// HasSrvExpcType returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasSrvExpcType() bool {
	if o != nil && o.SrvExpcType != nil {
		return true
	}

	return false
}

// SetSrvExpcType gets a reference to the given ServiceExperienceType and assigns it to the SrvExpcType field.
func (o *ServiceExperienceInfo) SetSrvExpcType(v ServiceExperienceType) {
	o.SrvExpcType = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetConfidence() int32 {
	if o == nil || o.Confidence == nil {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetConfidenceOk() (*int32, bool) {
	if o == nil || o.Confidence == nil {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasConfidence() bool {
	if o != nil && o.Confidence != nil {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *ServiceExperienceInfo) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetDnn() string {
	if o == nil || o.Dnn == nil {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetDnnOk() (*string, bool) {
	if o == nil || o.Dnn == nil {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasDnn() bool {
	if o != nil && o.Dnn != nil {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *ServiceExperienceInfo) SetDnn(v string) {
	o.Dnn = &v
}

// GetNetworkArea returns the NetworkArea field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetNetworkArea() NetworkAreaInfo {
	if o == nil || o.NetworkArea == nil {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.NetworkArea
}

// GetNetworkAreaOk returns a tuple with the NetworkArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil || o.NetworkArea == nil {
		return nil, false
	}
	return o.NetworkArea, true
}

// HasNetworkArea returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasNetworkArea() bool {
	if o != nil && o.NetworkArea != nil {
		return true
	}

	return false
}

// SetNetworkArea gets a reference to the given NetworkAreaInfo and assigns it to the NetworkArea field.
func (o *ServiceExperienceInfo) SetNetworkArea(v NetworkAreaInfo) {
	o.NetworkArea = &v
}

// GetNsiId returns the NsiId field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetNsiId() string {
	if o == nil || o.NsiId == nil {
		var ret string
		return ret
	}
	return *o.NsiId
}

// GetNsiIdOk returns a tuple with the NsiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetNsiIdOk() (*string, bool) {
	if o == nil || o.NsiId == nil {
		return nil, false
	}
	return o.NsiId, true
}

// HasNsiId returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasNsiId() bool {
	if o != nil && o.NsiId != nil {
		return true
	}

	return false
}

// SetNsiId gets a reference to the given string and assigns it to the NsiId field.
func (o *ServiceExperienceInfo) SetNsiId(v string) {
	o.NsiId = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetRatio() int32 {
	if o == nil || o.Ratio == nil {
		var ret int32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetRatioOk() (*int32, bool) {
	if o == nil || o.Ratio == nil {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasRatio() bool {
	if o != nil && o.Ratio != nil {
		return true
	}

	return false
}

// SetRatio gets a reference to the given int32 and assigns it to the Ratio field.
func (o *ServiceExperienceInfo) SetRatio(v int32) {
	o.Ratio = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetRatType() RatType {
	if o == nil || o.RatType == nil {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetRatTypeOk() (*RatType, bool) {
	if o == nil || o.RatType == nil {
		return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasRatType() bool {
	if o != nil && o.RatType != nil {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *ServiceExperienceInfo) SetRatType(v RatType) {
	o.RatType = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *ServiceExperienceInfo) GetFrequency() int32 {
	if o == nil || o.Frequency == nil {
		var ret int32
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceExperienceInfo) GetFrequencyOk() (*int32, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *ServiceExperienceInfo) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given int32 and assigns it to the Frequency field.
func (o *ServiceExperienceInfo) SetFrequency(v int32) {
	o.Frequency = &v
}

func (o ServiceExperienceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["svcExprc"] = o.SvcExprc
	}
	if o.SvcExprcVariance != nil {
		toSerialize["svcExprcVariance"] = o.SvcExprcVariance
	}
	if o.Supis != nil {
		toSerialize["supis"] = o.Supis
	}
	if o.Snssai != nil {
		toSerialize["snssai"] = o.Snssai
	}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.SrvExpcType != nil {
		toSerialize["srvExpcType"] = o.SrvExpcType
	}
	if o.Confidence != nil {
		toSerialize["confidence"] = o.Confidence
	}
	if o.Dnn != nil {
		toSerialize["dnn"] = o.Dnn
	}
	if o.NetworkArea != nil {
		toSerialize["networkArea"] = o.NetworkArea
	}
	if o.NsiId != nil {
		toSerialize["nsiId"] = o.NsiId
	}
	if o.Ratio != nil {
		toSerialize["ratio"] = o.Ratio
	}
	if o.RatType != nil {
		toSerialize["ratType"] = o.RatType
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	return json.Marshal(toSerialize)
}

type NullableServiceExperienceInfo struct {
	value *ServiceExperienceInfo
	isSet bool
}

func (v NullableServiceExperienceInfo) Get() *ServiceExperienceInfo {
	return v.value
}

func (v *NullableServiceExperienceInfo) Set(val *ServiceExperienceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceExperienceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceExperienceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceExperienceInfo(val *ServiceExperienceInfo) *NullableServiceExperienceInfo {
	return &NullableServiceExperienceInfo{value: val, isSet: true}
}

func (v NullableServiceExperienceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceExperienceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


