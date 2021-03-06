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

// ContextElement Contains context information corresponding with a specific context identifier.
type ContextElement struct {
	ContextId AnalyticsContextIdentifier `json:"contextId"`
	// Output analytics for the analytics subscription which have not yet been sent to the analytics consumer.
	PendAnalytics []EventNotification `json:"pendAnalytics,omitempty"`
	// Historical output analytics.
	HistAnalytics []EventNotification `json:"histAnalytics,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LastOutputTime *time.Time `json:"lastOutputTime,omitempty"`
	// Information about analytics subscriptions that the NWDAF has with other NWDAFs to perform aggregation.
	AggrSubs []SpecificAnalyticsSubscription `json:"aggrSubs,omitempty"`
	// Historical data related to the analytics subscription.
	HistData []HistoricalData `json:"histData,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AdrfId *string `json:"adrfId,omitempty"`
	// Type(s) of data stored in the ADRF by the NWDAF.
	AdrfDataTypes []AdrfDataType `json:"adrfDataTypes,omitempty"`
	// NWDAF identifiers of NWDAF instances used by the NWDAF service consumer when aggregating multiple analytics subscriptions.
	AggrNwdafIds []string `json:"aggrNwdafIds,omitempty"`
	// Identifiers of NWDAFs that provide ML models used by the NF service consumer.
	ModelProvIds []string `json:"modelProvIds,omitempty"`
}

// NewContextElement instantiates a new ContextElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextElement(contextId AnalyticsContextIdentifier) *ContextElement {
	this := ContextElement{}
	this.ContextId = contextId
	return &this
}

// NewContextElementWithDefaults instantiates a new ContextElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextElementWithDefaults() *ContextElement {
	this := ContextElement{}
	return &this
}

// GetContextId returns the ContextId field value
func (o *ContextElement) GetContextId() AnalyticsContextIdentifier {
	if o == nil {
		var ret AnalyticsContextIdentifier
		return ret
	}

	return o.ContextId
}

// GetContextIdOk returns a tuple with the ContextId field value
// and a boolean to check if the value has been set.
func (o *ContextElement) GetContextIdOk() (*AnalyticsContextIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContextId, true
}

// SetContextId sets field value
func (o *ContextElement) SetContextId(v AnalyticsContextIdentifier) {
	o.ContextId = v
}

// GetPendAnalytics returns the PendAnalytics field value if set, zero value otherwise.
func (o *ContextElement) GetPendAnalytics() []EventNotification {
	if o == nil || o.PendAnalytics == nil {
		var ret []EventNotification
		return ret
	}
	return o.PendAnalytics
}

// GetPendAnalyticsOk returns a tuple with the PendAnalytics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextElement) GetPendAnalyticsOk() ([]EventNotification, bool) {
	if o == nil || o.PendAnalytics == nil {
		return nil, false
	}
	return o.PendAnalytics, true
}

// HasPendAnalytics returns a boolean if a field has been set.
func (o *ContextElement) HasPendAnalytics() bool {
	if o != nil && o.PendAnalytics != nil {
		return true
	}

	return false
}

// SetPendAnalytics gets a reference to the given []EventNotification and assigns it to the PendAnalytics field.
func (o *ContextElement) SetPendAnalytics(v []EventNotification) {
	o.PendAnalytics = v
}

// GetHistAnalytics returns the HistAnalytics field value if set, zero value otherwise.
func (o *ContextElement) GetHistAnalytics() []EventNotification {
	if o == nil || o.HistAnalytics == nil {
		var ret []EventNotification
		return ret
	}
	return o.HistAnalytics
}

// GetHistAnalyticsOk returns a tuple with the HistAnalytics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextElement) GetHistAnalyticsOk() ([]EventNotification, bool) {
	if o == nil || o.HistAnalytics == nil {
		return nil, false
	}
	return o.HistAnalytics, true
}

// HasHistAnalytics returns a boolean if a field has been set.
func (o *ContextElement) HasHistAnalytics() bool {
	if o != nil && o.HistAnalytics != nil {
		return true
	}

	return false
}

// SetHistAnalytics gets a reference to the given []EventNotification and assigns it to the HistAnalytics field.
func (o *ContextElement) SetHistAnalytics(v []EventNotification) {
	o.HistAnalytics = v
}

// GetLastOutputTime returns the LastOutputTime field value if set, zero value otherwise.
func (o *ContextElement) GetLastOutputTime() time.Time {
	if o == nil || o.LastOutputTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastOutputTime
}

// GetLastOutputTimeOk returns a tuple with the LastOutputTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextElement) GetLastOutputTimeOk() (*time.Time, bool) {
	if o == nil || o.LastOutputTime == nil {
		return nil, false
	}
	return o.LastOutputTime, true
}

// HasLastOutputTime returns a boolean if a field has been set.
func (o *ContextElement) HasLastOutputTime() bool {
	if o != nil && o.LastOutputTime != nil {
		return true
	}

	return false
}

// SetLastOutputTime gets a reference to the given time.Time and assigns it to the LastOutputTime field.
func (o *ContextElement) SetLastOutputTime(v time.Time) {
	o.LastOutputTime = &v
}

// GetAggrSubs returns the AggrSubs field value if set, zero value otherwise.
func (o *ContextElement) GetAggrSubs() []SpecificAnalyticsSubscription {
	if o == nil || o.AggrSubs == nil {
		var ret []SpecificAnalyticsSubscription
		return ret
	}
	return o.AggrSubs
}

// GetAggrSubsOk returns a tuple with the AggrSubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextElement) GetAggrSubsOk() ([]SpecificAnalyticsSubscription, bool) {
	if o == nil || o.AggrSubs == nil {
		return nil, false
	}
	return o.AggrSubs, true
}

// HasAggrSubs returns a boolean if a field has been set.
func (o *ContextElement) HasAggrSubs() bool {
	if o != nil && o.AggrSubs != nil {
		return true
	}

	return false
}

// SetAggrSubs gets a reference to the given []SpecificAnalyticsSubscription and assigns it to the AggrSubs field.
func (o *ContextElement) SetAggrSubs(v []SpecificAnalyticsSubscription) {
	o.AggrSubs = v
}

// GetHistData returns the HistData field value if set, zero value otherwise.
func (o *ContextElement) GetHistData() []HistoricalData {
	if o == nil || o.HistData == nil {
		var ret []HistoricalData
		return ret
	}
	return o.HistData
}

// GetHistDataOk returns a tuple with the HistData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextElement) GetHistDataOk() ([]HistoricalData, bool) {
	if o == nil || o.HistData == nil {
		return nil, false
	}
	return o.HistData, true
}

// HasHistData returns a boolean if a field has been set.
func (o *ContextElement) HasHistData() bool {
	if o != nil && o.HistData != nil {
		return true
	}

	return false
}

// SetHistData gets a reference to the given []HistoricalData and assigns it to the HistData field.
func (o *ContextElement) SetHistData(v []HistoricalData) {
	o.HistData = v
}

// GetAdrfId returns the AdrfId field value if set, zero value otherwise.
func (o *ContextElement) GetAdrfId() string {
	if o == nil || o.AdrfId == nil {
		var ret string
		return ret
	}
	return *o.AdrfId
}

// GetAdrfIdOk returns a tuple with the AdrfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextElement) GetAdrfIdOk() (*string, bool) {
	if o == nil || o.AdrfId == nil {
		return nil, false
	}
	return o.AdrfId, true
}

// HasAdrfId returns a boolean if a field has been set.
func (o *ContextElement) HasAdrfId() bool {
	if o != nil && o.AdrfId != nil {
		return true
	}

	return false
}

// SetAdrfId gets a reference to the given string and assigns it to the AdrfId field.
func (o *ContextElement) SetAdrfId(v string) {
	o.AdrfId = &v
}

// GetAdrfDataTypes returns the AdrfDataTypes field value if set, zero value otherwise.
func (o *ContextElement) GetAdrfDataTypes() []AdrfDataType {
	if o == nil || o.AdrfDataTypes == nil {
		var ret []AdrfDataType
		return ret
	}
	return o.AdrfDataTypes
}

// GetAdrfDataTypesOk returns a tuple with the AdrfDataTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextElement) GetAdrfDataTypesOk() ([]AdrfDataType, bool) {
	if o == nil || o.AdrfDataTypes == nil {
		return nil, false
	}
	return o.AdrfDataTypes, true
}

// HasAdrfDataTypes returns a boolean if a field has been set.
func (o *ContextElement) HasAdrfDataTypes() bool {
	if o != nil && o.AdrfDataTypes != nil {
		return true
	}

	return false
}

// SetAdrfDataTypes gets a reference to the given []AdrfDataType and assigns it to the AdrfDataTypes field.
func (o *ContextElement) SetAdrfDataTypes(v []AdrfDataType) {
	o.AdrfDataTypes = v
}

// GetAggrNwdafIds returns the AggrNwdafIds field value if set, zero value otherwise.
func (o *ContextElement) GetAggrNwdafIds() []string {
	if o == nil || o.AggrNwdafIds == nil {
		var ret []string
		return ret
	}
	return o.AggrNwdafIds
}

// GetAggrNwdafIdsOk returns a tuple with the AggrNwdafIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextElement) GetAggrNwdafIdsOk() ([]string, bool) {
	if o == nil || o.AggrNwdafIds == nil {
		return nil, false
	}
	return o.AggrNwdafIds, true
}

// HasAggrNwdafIds returns a boolean if a field has been set.
func (o *ContextElement) HasAggrNwdafIds() bool {
	if o != nil && o.AggrNwdafIds != nil {
		return true
	}

	return false
}

// SetAggrNwdafIds gets a reference to the given []string and assigns it to the AggrNwdafIds field.
func (o *ContextElement) SetAggrNwdafIds(v []string) {
	o.AggrNwdafIds = v
}

// GetModelProvIds returns the ModelProvIds field value if set, zero value otherwise.
func (o *ContextElement) GetModelProvIds() []string {
	if o == nil || o.ModelProvIds == nil {
		var ret []string
		return ret
	}
	return o.ModelProvIds
}

// GetModelProvIdsOk returns a tuple with the ModelProvIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextElement) GetModelProvIdsOk() ([]string, bool) {
	if o == nil || o.ModelProvIds == nil {
		return nil, false
	}
	return o.ModelProvIds, true
}

// HasModelProvIds returns a boolean if a field has been set.
func (o *ContextElement) HasModelProvIds() bool {
	if o != nil && o.ModelProvIds != nil {
		return true
	}

	return false
}

// SetModelProvIds gets a reference to the given []string and assigns it to the ModelProvIds field.
func (o *ContextElement) SetModelProvIds(v []string) {
	o.ModelProvIds = v
}

func (o ContextElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contextId"] = o.ContextId
	}
	if o.PendAnalytics != nil {
		toSerialize["pendAnalytics"] = o.PendAnalytics
	}
	if o.HistAnalytics != nil {
		toSerialize["histAnalytics"] = o.HistAnalytics
	}
	if o.LastOutputTime != nil {
		toSerialize["lastOutputTime"] = o.LastOutputTime
	}
	if o.AggrSubs != nil {
		toSerialize["aggrSubs"] = o.AggrSubs
	}
	if o.HistData != nil {
		toSerialize["histData"] = o.HistData
	}
	if o.AdrfId != nil {
		toSerialize["adrfId"] = o.AdrfId
	}
	if o.AdrfDataTypes != nil {
		toSerialize["adrfDataTypes"] = o.AdrfDataTypes
	}
	if o.AggrNwdafIds != nil {
		toSerialize["aggrNwdafIds"] = o.AggrNwdafIds
	}
	if o.ModelProvIds != nil {
		toSerialize["modelProvIds"] = o.ModelProvIds
	}
	return json.Marshal(toSerialize)
}

type NullableContextElement struct {
	value *ContextElement
	isSet bool
}

func (v NullableContextElement) Get() *ContextElement {
	return v.value
}

func (v *NullableContextElement) Set(val *ContextElement) {
	v.value = val
	v.isSet = true
}

func (v NullableContextElement) IsSet() bool {
	return v.isSet
}

func (v *NullableContextElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextElement(val *ContextElement) *NullableContextElement {
	return &NullableContextElement{value: val, isSet: true}
}

func (v NullableContextElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


