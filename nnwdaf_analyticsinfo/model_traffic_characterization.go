/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TrafficCharacterization struct for TrafficCharacterization
type TrafficCharacterization struct {
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	AppId *string `json:"appId,omitempty"`
	FDescs []IpEthFlowDescription `json:"fDescs,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	UlVol *int64 `json:"ulVol,omitempty"`
	UlVolVariance *float32 `json:"ulVolVariance,omitempty"`
	// Unsigned integer identifying a volume in units of bytes.
	DlVol *int64 `json:"dlVol,omitempty"`
	DlVolVariance *float32 `json:"dlVolVariance,omitempty"`
}

// NewTrafficCharacterization instantiates a new TrafficCharacterization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficCharacterization() *TrafficCharacterization {
	this := TrafficCharacterization{}
	return &this
}

// NewTrafficCharacterizationWithDefaults instantiates a new TrafficCharacterization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficCharacterizationWithDefaults() *TrafficCharacterization {
	this := TrafficCharacterization{}
	return &this
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *TrafficCharacterization) GetDnn() string {
	if o == nil || o.Dnn == nil {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCharacterization) GetDnnOk() (*string, bool) {
	if o == nil || o.Dnn == nil {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *TrafficCharacterization) HasDnn() bool {
	if o != nil && o.Dnn != nil {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *TrafficCharacterization) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *TrafficCharacterization) GetSnssai() Snssai {
	if o == nil || o.Snssai == nil {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCharacterization) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || o.Snssai == nil {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *TrafficCharacterization) HasSnssai() bool {
	if o != nil && o.Snssai != nil {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *TrafficCharacterization) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *TrafficCharacterization) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCharacterization) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *TrafficCharacterization) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *TrafficCharacterization) SetAppId(v string) {
	o.AppId = &v
}

// GetFDescs returns the FDescs field value if set, zero value otherwise.
func (o *TrafficCharacterization) GetFDescs() []IpEthFlowDescription {
	if o == nil || o.FDescs == nil {
		var ret []IpEthFlowDescription
		return ret
	}
	return o.FDescs
}

// GetFDescsOk returns a tuple with the FDescs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCharacterization) GetFDescsOk() ([]IpEthFlowDescription, bool) {
	if o == nil || o.FDescs == nil {
		return nil, false
	}
	return o.FDescs, true
}

// HasFDescs returns a boolean if a field has been set.
func (o *TrafficCharacterization) HasFDescs() bool {
	if o != nil && o.FDescs != nil {
		return true
	}

	return false
}

// SetFDescs gets a reference to the given []IpEthFlowDescription and assigns it to the FDescs field.
func (o *TrafficCharacterization) SetFDescs(v []IpEthFlowDescription) {
	o.FDescs = v
}

// GetUlVol returns the UlVol field value if set, zero value otherwise.
func (o *TrafficCharacterization) GetUlVol() int64 {
	if o == nil || o.UlVol == nil {
		var ret int64
		return ret
	}
	return *o.UlVol
}

// GetUlVolOk returns a tuple with the UlVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCharacterization) GetUlVolOk() (*int64, bool) {
	if o == nil || o.UlVol == nil {
		return nil, false
	}
	return o.UlVol, true
}

// HasUlVol returns a boolean if a field has been set.
func (o *TrafficCharacterization) HasUlVol() bool {
	if o != nil && o.UlVol != nil {
		return true
	}

	return false
}

// SetUlVol gets a reference to the given int64 and assigns it to the UlVol field.
func (o *TrafficCharacterization) SetUlVol(v int64) {
	o.UlVol = &v
}

// GetUlVolVariance returns the UlVolVariance field value if set, zero value otherwise.
func (o *TrafficCharacterization) GetUlVolVariance() float32 {
	if o == nil || o.UlVolVariance == nil {
		var ret float32
		return ret
	}
	return *o.UlVolVariance
}

// GetUlVolVarianceOk returns a tuple with the UlVolVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCharacterization) GetUlVolVarianceOk() (*float32, bool) {
	if o == nil || o.UlVolVariance == nil {
		return nil, false
	}
	return o.UlVolVariance, true
}

// HasUlVolVariance returns a boolean if a field has been set.
func (o *TrafficCharacterization) HasUlVolVariance() bool {
	if o != nil && o.UlVolVariance != nil {
		return true
	}

	return false
}

// SetUlVolVariance gets a reference to the given float32 and assigns it to the UlVolVariance field.
func (o *TrafficCharacterization) SetUlVolVariance(v float32) {
	o.UlVolVariance = &v
}

// GetDlVol returns the DlVol field value if set, zero value otherwise.
func (o *TrafficCharacterization) GetDlVol() int64 {
	if o == nil || o.DlVol == nil {
		var ret int64
		return ret
	}
	return *o.DlVol
}

// GetDlVolOk returns a tuple with the DlVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCharacterization) GetDlVolOk() (*int64, bool) {
	if o == nil || o.DlVol == nil {
		return nil, false
	}
	return o.DlVol, true
}

// HasDlVol returns a boolean if a field has been set.
func (o *TrafficCharacterization) HasDlVol() bool {
	if o != nil && o.DlVol != nil {
		return true
	}

	return false
}

// SetDlVol gets a reference to the given int64 and assigns it to the DlVol field.
func (o *TrafficCharacterization) SetDlVol(v int64) {
	o.DlVol = &v
}

// GetDlVolVariance returns the DlVolVariance field value if set, zero value otherwise.
func (o *TrafficCharacterization) GetDlVolVariance() float32 {
	if o == nil || o.DlVolVariance == nil {
		var ret float32
		return ret
	}
	return *o.DlVolVariance
}

// GetDlVolVarianceOk returns a tuple with the DlVolVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCharacterization) GetDlVolVarianceOk() (*float32, bool) {
	if o == nil || o.DlVolVariance == nil {
		return nil, false
	}
	return o.DlVolVariance, true
}

// HasDlVolVariance returns a boolean if a field has been set.
func (o *TrafficCharacterization) HasDlVolVariance() bool {
	if o != nil && o.DlVolVariance != nil {
		return true
	}

	return false
}

// SetDlVolVariance gets a reference to the given float32 and assigns it to the DlVolVariance field.
func (o *TrafficCharacterization) SetDlVolVariance(v float32) {
	o.DlVolVariance = &v
}

func (o TrafficCharacterization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dnn != nil {
		toSerialize["dnn"] = o.Dnn
	}
	if o.Snssai != nil {
		toSerialize["snssai"] = o.Snssai
	}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.FDescs != nil {
		toSerialize["fDescs"] = o.FDescs
	}
	if o.UlVol != nil {
		toSerialize["ulVol"] = o.UlVol
	}
	if o.UlVolVariance != nil {
		toSerialize["ulVolVariance"] = o.UlVolVariance
	}
	if o.DlVol != nil {
		toSerialize["dlVol"] = o.DlVol
	}
	if o.DlVolVariance != nil {
		toSerialize["dlVolVariance"] = o.DlVolVariance
	}
	return json.Marshal(toSerialize)
}

type NullableTrafficCharacterization struct {
	value *TrafficCharacterization
	isSet bool
}

func (v NullableTrafficCharacterization) Get() *TrafficCharacterization {
	return v.value
}

func (v *NullableTrafficCharacterization) Set(val *TrafficCharacterization) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficCharacterization) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficCharacterization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficCharacterization(val *TrafficCharacterization) *NullableTrafficCharacterization {
	return &NullableTrafficCharacterization{value: val, isSet: true}
}

func (v NullableTrafficCharacterization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficCharacterization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


