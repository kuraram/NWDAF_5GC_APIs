/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ReportingInformation Represents the type of reporting that the subscription requires.
type ReportingInformation struct {
	ImmRep *bool `json:"immRep,omitempty"`
	NotifMethod *NotificationMethod1 `json:"notifMethod,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReportNbr *int32 `json:"maxReportNbr,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	MonDur *time.Time `json:"monDur,omitempty"`
	// indicating a time in seconds.
	RepPeriod *int32 `json:"repPeriod,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	SampRatio *int32 `json:"sampRatio,omitempty"`
	// Criteria for partitioning the UEs before applying the sampling ratio.
	PartitionCriteria []PartitioningCriteria `json:"partitionCriteria,omitempty"`
	// indicating a time in seconds.
	GrpRepTime *int32 `json:"grpRepTime,omitempty"`
	NotifFlag *NotificationFlag `json:"notifFlag,omitempty"`
}

// NewReportingInformation instantiates a new ReportingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingInformation() *ReportingInformation {
	this := ReportingInformation{}
	return &this
}

// NewReportingInformationWithDefaults instantiates a new ReportingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingInformationWithDefaults() *ReportingInformation {
	this := ReportingInformation{}
	return &this
}

// GetImmRep returns the ImmRep field value if set, zero value otherwise.
func (o *ReportingInformation) GetImmRep() bool {
	if o == nil || o.ImmRep == nil {
		var ret bool
		return ret
	}
	return *o.ImmRep
}

// GetImmRepOk returns a tuple with the ImmRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetImmRepOk() (*bool, bool) {
	if o == nil || o.ImmRep == nil {
		return nil, false
	}
	return o.ImmRep, true
}

// HasImmRep returns a boolean if a field has been set.
func (o *ReportingInformation) HasImmRep() bool {
	if o != nil && o.ImmRep != nil {
		return true
	}

	return false
}

// SetImmRep gets a reference to the given bool and assigns it to the ImmRep field.
func (o *ReportingInformation) SetImmRep(v bool) {
	o.ImmRep = &v
}

// GetNotifMethod returns the NotifMethod field value if set, zero value otherwise.
func (o *ReportingInformation) GetNotifMethod() NotificationMethod1 {
	if o == nil || o.NotifMethod == nil {
		var ret NotificationMethod1
		return ret
	}
	return *o.NotifMethod
}

// GetNotifMethodOk returns a tuple with the NotifMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetNotifMethodOk() (*NotificationMethod1, bool) {
	if o == nil || o.NotifMethod == nil {
		return nil, false
	}
	return o.NotifMethod, true
}

// HasNotifMethod returns a boolean if a field has been set.
func (o *ReportingInformation) HasNotifMethod() bool {
	if o != nil && o.NotifMethod != nil {
		return true
	}

	return false
}

// SetNotifMethod gets a reference to the given NotificationMethod1 and assigns it to the NotifMethod field.
func (o *ReportingInformation) SetNotifMethod(v NotificationMethod1) {
	o.NotifMethod = &v
}

// GetMaxReportNbr returns the MaxReportNbr field value if set, zero value otherwise.
func (o *ReportingInformation) GetMaxReportNbr() int32 {
	if o == nil || o.MaxReportNbr == nil {
		var ret int32
		return ret
	}
	return *o.MaxReportNbr
}

// GetMaxReportNbrOk returns a tuple with the MaxReportNbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetMaxReportNbrOk() (*int32, bool) {
	if o == nil || o.MaxReportNbr == nil {
		return nil, false
	}
	return o.MaxReportNbr, true
}

// HasMaxReportNbr returns a boolean if a field has been set.
func (o *ReportingInformation) HasMaxReportNbr() bool {
	if o != nil && o.MaxReportNbr != nil {
		return true
	}

	return false
}

// SetMaxReportNbr gets a reference to the given int32 and assigns it to the MaxReportNbr field.
func (o *ReportingInformation) SetMaxReportNbr(v int32) {
	o.MaxReportNbr = &v
}

// GetMonDur returns the MonDur field value if set, zero value otherwise.
func (o *ReportingInformation) GetMonDur() time.Time {
	if o == nil || o.MonDur == nil {
		var ret time.Time
		return ret
	}
	return *o.MonDur
}

// GetMonDurOk returns a tuple with the MonDur field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetMonDurOk() (*time.Time, bool) {
	if o == nil || o.MonDur == nil {
		return nil, false
	}
	return o.MonDur, true
}

// HasMonDur returns a boolean if a field has been set.
func (o *ReportingInformation) HasMonDur() bool {
	if o != nil && o.MonDur != nil {
		return true
	}

	return false
}

// SetMonDur gets a reference to the given time.Time and assigns it to the MonDur field.
func (o *ReportingInformation) SetMonDur(v time.Time) {
	o.MonDur = &v
}

// GetRepPeriod returns the RepPeriod field value if set, zero value otherwise.
func (o *ReportingInformation) GetRepPeriod() int32 {
	if o == nil || o.RepPeriod == nil {
		var ret int32
		return ret
	}
	return *o.RepPeriod
}

// GetRepPeriodOk returns a tuple with the RepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetRepPeriodOk() (*int32, bool) {
	if o == nil || o.RepPeriod == nil {
		return nil, false
	}
	return o.RepPeriod, true
}

// HasRepPeriod returns a boolean if a field has been set.
func (o *ReportingInformation) HasRepPeriod() bool {
	if o != nil && o.RepPeriod != nil {
		return true
	}

	return false
}

// SetRepPeriod gets a reference to the given int32 and assigns it to the RepPeriod field.
func (o *ReportingInformation) SetRepPeriod(v int32) {
	o.RepPeriod = &v
}

// GetSampRatio returns the SampRatio field value if set, zero value otherwise.
func (o *ReportingInformation) GetSampRatio() int32 {
	if o == nil || o.SampRatio == nil {
		var ret int32
		return ret
	}
	return *o.SampRatio
}

// GetSampRatioOk returns a tuple with the SampRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetSampRatioOk() (*int32, bool) {
	if o == nil || o.SampRatio == nil {
		return nil, false
	}
	return o.SampRatio, true
}

// HasSampRatio returns a boolean if a field has been set.
func (o *ReportingInformation) HasSampRatio() bool {
	if o != nil && o.SampRatio != nil {
		return true
	}

	return false
}

// SetSampRatio gets a reference to the given int32 and assigns it to the SampRatio field.
func (o *ReportingInformation) SetSampRatio(v int32) {
	o.SampRatio = &v
}

// GetPartitionCriteria returns the PartitionCriteria field value if set, zero value otherwise.
func (o *ReportingInformation) GetPartitionCriteria() []PartitioningCriteria {
	if o == nil || o.PartitionCriteria == nil {
		var ret []PartitioningCriteria
		return ret
	}
	return o.PartitionCriteria
}

// GetPartitionCriteriaOk returns a tuple with the PartitionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetPartitionCriteriaOk() ([]PartitioningCriteria, bool) {
	if o == nil || o.PartitionCriteria == nil {
		return nil, false
	}
	return o.PartitionCriteria, true
}

// HasPartitionCriteria returns a boolean if a field has been set.
func (o *ReportingInformation) HasPartitionCriteria() bool {
	if o != nil && o.PartitionCriteria != nil {
		return true
	}

	return false
}

// SetPartitionCriteria gets a reference to the given []PartitioningCriteria and assigns it to the PartitionCriteria field.
func (o *ReportingInformation) SetPartitionCriteria(v []PartitioningCriteria) {
	o.PartitionCriteria = v
}

// GetGrpRepTime returns the GrpRepTime field value if set, zero value otherwise.
func (o *ReportingInformation) GetGrpRepTime() int32 {
	if o == nil || o.GrpRepTime == nil {
		var ret int32
		return ret
	}
	return *o.GrpRepTime
}

// GetGrpRepTimeOk returns a tuple with the GrpRepTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetGrpRepTimeOk() (*int32, bool) {
	if o == nil || o.GrpRepTime == nil {
		return nil, false
	}
	return o.GrpRepTime, true
}

// HasGrpRepTime returns a boolean if a field has been set.
func (o *ReportingInformation) HasGrpRepTime() bool {
	if o != nil && o.GrpRepTime != nil {
		return true
	}

	return false
}

// SetGrpRepTime gets a reference to the given int32 and assigns it to the GrpRepTime field.
func (o *ReportingInformation) SetGrpRepTime(v int32) {
	o.GrpRepTime = &v
}

// GetNotifFlag returns the NotifFlag field value if set, zero value otherwise.
func (o *ReportingInformation) GetNotifFlag() NotificationFlag {
	if o == nil || o.NotifFlag == nil {
		var ret NotificationFlag
		return ret
	}
	return *o.NotifFlag
}

// GetNotifFlagOk returns a tuple with the NotifFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingInformation) GetNotifFlagOk() (*NotificationFlag, bool) {
	if o == nil || o.NotifFlag == nil {
		return nil, false
	}
	return o.NotifFlag, true
}

// HasNotifFlag returns a boolean if a field has been set.
func (o *ReportingInformation) HasNotifFlag() bool {
	if o != nil && o.NotifFlag != nil {
		return true
	}

	return false
}

// SetNotifFlag gets a reference to the given NotificationFlag and assigns it to the NotifFlag field.
func (o *ReportingInformation) SetNotifFlag(v NotificationFlag) {
	o.NotifFlag = &v
}

func (o ReportingInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImmRep != nil {
		toSerialize["immRep"] = o.ImmRep
	}
	if o.NotifMethod != nil {
		toSerialize["notifMethod"] = o.NotifMethod
	}
	if o.MaxReportNbr != nil {
		toSerialize["maxReportNbr"] = o.MaxReportNbr
	}
	if o.MonDur != nil {
		toSerialize["monDur"] = o.MonDur
	}
	if o.RepPeriod != nil {
		toSerialize["repPeriod"] = o.RepPeriod
	}
	if o.SampRatio != nil {
		toSerialize["sampRatio"] = o.SampRatio
	}
	if o.PartitionCriteria != nil {
		toSerialize["partitionCriteria"] = o.PartitionCriteria
	}
	if o.GrpRepTime != nil {
		toSerialize["grpRepTime"] = o.GrpRepTime
	}
	if o.NotifFlag != nil {
		toSerialize["notifFlag"] = o.NotifFlag
	}
	return json.Marshal(toSerialize)
}

type NullableReportingInformation struct {
	value *ReportingInformation
	isSet bool
}

func (v NullableReportingInformation) Get() *ReportingInformation {
	return v.value
}

func (v *NullableReportingInformation) Set(val *ReportingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingInformation(val *ReportingInformation) *NullableReportingInformation {
	return &NullableReportingInformation{value: val, isSet: true}
}

func (v NullableReportingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


