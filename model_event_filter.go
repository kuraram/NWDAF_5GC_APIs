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

// EventFilter Represents the event filters used to identify the requested analytics.
type EventFilter struct {
	// FALSE represents not applicable for all slices. TRUE represents applicable for all slices.
	AnySlice *bool `json:"anySlice,omitempty"`
	// Identification(s) of network slice to which the subscription belongs.
	Snssais []Snssai `json:"snssais,omitempty"`
	AppIds []string `json:"appIds,omitempty"`
	Dnns []string `json:"dnns,omitempty"`
	Dnais []string `json:"dnais,omitempty"`
	NetworkArea *NetworkAreaInfo `json:"networkArea,omitempty"`
	NfInstanceIds []string `json:"nfInstanceIds,omitempty"`
	NfSetIds []string `json:"nfSetIds,omitempty"`
	NfTypes []NFType `json:"nfTypes,omitempty"`
	NsiIdInfos []NsiIdInfo `json:"nsiIdInfos,omitempty"`
	QosRequ *QosRequirement `json:"qosRequ,omitempty"`
	NwPerfTypes []NetworkPerfType `json:"nwPerfTypes,omitempty"`
	BwRequs []BwRequirement `json:"bwRequs,omitempty"`
	ExcepIds []ExceptionId `json:"excepIds,omitempty"`
	ExptAnaType *ExpectedAnalyticsType `json:"exptAnaType,omitempty"`
	ExptUeBehav *ExpectedUeBehaviourData `json:"exptUeBehav,omitempty"`
}

// NewEventFilter instantiates a new EventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFilter() *EventFilter {
	this := EventFilter{}
	return &this
}

// NewEventFilterWithDefaults instantiates a new EventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFilterWithDefaults() *EventFilter {
	this := EventFilter{}
	return &this
}

// GetAnySlice returns the AnySlice field value if set, zero value otherwise.
func (o *EventFilter) GetAnySlice() bool {
	if o == nil || o.AnySlice == nil {
		var ret bool
		return ret
	}
	return *o.AnySlice
}

// GetAnySliceOk returns a tuple with the AnySlice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetAnySliceOk() (*bool, bool) {
	if o == nil || o.AnySlice == nil {
		return nil, false
	}
	return o.AnySlice, true
}

// HasAnySlice returns a boolean if a field has been set.
func (o *EventFilter) HasAnySlice() bool {
	if o != nil && o.AnySlice != nil {
		return true
	}

	return false
}

// SetAnySlice gets a reference to the given bool and assigns it to the AnySlice field.
func (o *EventFilter) SetAnySlice(v bool) {
	o.AnySlice = &v
}

// GetSnssais returns the Snssais field value if set, zero value otherwise.
func (o *EventFilter) GetSnssais() []Snssai {
	if o == nil || o.Snssais == nil {
		var ret []Snssai
		return ret
	}
	return o.Snssais
}

// GetSnssaisOk returns a tuple with the Snssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetSnssaisOk() ([]Snssai, bool) {
	if o == nil || o.Snssais == nil {
		return nil, false
	}
	return o.Snssais, true
}

// HasSnssais returns a boolean if a field has been set.
func (o *EventFilter) HasSnssais() bool {
	if o != nil && o.Snssais != nil {
		return true
	}

	return false
}

// SetSnssais gets a reference to the given []Snssai and assigns it to the Snssais field.
func (o *EventFilter) SetSnssais(v []Snssai) {
	o.Snssais = v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *EventFilter) GetAppIds() []string {
	if o == nil || o.AppIds == nil {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetAppIdsOk() ([]string, bool) {
	if o == nil || o.AppIds == nil {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *EventFilter) HasAppIds() bool {
	if o != nil && o.AppIds != nil {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *EventFilter) SetAppIds(v []string) {
	o.AppIds = v
}

// GetDnns returns the Dnns field value if set, zero value otherwise.
func (o *EventFilter) GetDnns() []string {
	if o == nil || o.Dnns == nil {
		var ret []string
		return ret
	}
	return o.Dnns
}

// GetDnnsOk returns a tuple with the Dnns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetDnnsOk() ([]string, bool) {
	if o == nil || o.Dnns == nil {
		return nil, false
	}
	return o.Dnns, true
}

// HasDnns returns a boolean if a field has been set.
func (o *EventFilter) HasDnns() bool {
	if o != nil && o.Dnns != nil {
		return true
	}

	return false
}

// SetDnns gets a reference to the given []string and assigns it to the Dnns field.
func (o *EventFilter) SetDnns(v []string) {
	o.Dnns = v
}

// GetDnais returns the Dnais field value if set, zero value otherwise.
func (o *EventFilter) GetDnais() []string {
	if o == nil || o.Dnais == nil {
		var ret []string
		return ret
	}
	return o.Dnais
}

// GetDnaisOk returns a tuple with the Dnais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetDnaisOk() ([]string, bool) {
	if o == nil || o.Dnais == nil {
		return nil, false
	}
	return o.Dnais, true
}

// HasDnais returns a boolean if a field has been set.
func (o *EventFilter) HasDnais() bool {
	if o != nil && o.Dnais != nil {
		return true
	}

	return false
}

// SetDnais gets a reference to the given []string and assigns it to the Dnais field.
func (o *EventFilter) SetDnais(v []string) {
	o.Dnais = v
}

// GetNetworkArea returns the NetworkArea field value if set, zero value otherwise.
func (o *EventFilter) GetNetworkArea() NetworkAreaInfo {
	if o == nil || o.NetworkArea == nil {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.NetworkArea
}

// GetNetworkAreaOk returns a tuple with the NetworkArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNetworkAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil || o.NetworkArea == nil {
		return nil, false
	}
	return o.NetworkArea, true
}

// HasNetworkArea returns a boolean if a field has been set.
func (o *EventFilter) HasNetworkArea() bool {
	if o != nil && o.NetworkArea != nil {
		return true
	}

	return false
}

// SetNetworkArea gets a reference to the given NetworkAreaInfo and assigns it to the NetworkArea field.
func (o *EventFilter) SetNetworkArea(v NetworkAreaInfo) {
	o.NetworkArea = &v
}

// GetNfInstanceIds returns the NfInstanceIds field value if set, zero value otherwise.
func (o *EventFilter) GetNfInstanceIds() []string {
	if o == nil || o.NfInstanceIds == nil {
		var ret []string
		return ret
	}
	return o.NfInstanceIds
}

// GetNfInstanceIdsOk returns a tuple with the NfInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNfInstanceIdsOk() ([]string, bool) {
	if o == nil || o.NfInstanceIds == nil {
		return nil, false
	}
	return o.NfInstanceIds, true
}

// HasNfInstanceIds returns a boolean if a field has been set.
func (o *EventFilter) HasNfInstanceIds() bool {
	if o != nil && o.NfInstanceIds != nil {
		return true
	}

	return false
}

// SetNfInstanceIds gets a reference to the given []string and assigns it to the NfInstanceIds field.
func (o *EventFilter) SetNfInstanceIds(v []string) {
	o.NfInstanceIds = v
}

// GetNfSetIds returns the NfSetIds field value if set, zero value otherwise.
func (o *EventFilter) GetNfSetIds() []string {
	if o == nil || o.NfSetIds == nil {
		var ret []string
		return ret
	}
	return o.NfSetIds
}

// GetNfSetIdsOk returns a tuple with the NfSetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNfSetIdsOk() ([]string, bool) {
	if o == nil || o.NfSetIds == nil {
		return nil, false
	}
	return o.NfSetIds, true
}

// HasNfSetIds returns a boolean if a field has been set.
func (o *EventFilter) HasNfSetIds() bool {
	if o != nil && o.NfSetIds != nil {
		return true
	}

	return false
}

// SetNfSetIds gets a reference to the given []string and assigns it to the NfSetIds field.
func (o *EventFilter) SetNfSetIds(v []string) {
	o.NfSetIds = v
}

// GetNfTypes returns the NfTypes field value if set, zero value otherwise.
func (o *EventFilter) GetNfTypes() []NFType {
	if o == nil || o.NfTypes == nil {
		var ret []NFType
		return ret
	}
	return o.NfTypes
}

// GetNfTypesOk returns a tuple with the NfTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNfTypesOk() ([]NFType, bool) {
	if o == nil || o.NfTypes == nil {
		return nil, false
	}
	return o.NfTypes, true
}

// HasNfTypes returns a boolean if a field has been set.
func (o *EventFilter) HasNfTypes() bool {
	if o != nil && o.NfTypes != nil {
		return true
	}

	return false
}

// SetNfTypes gets a reference to the given []NFType and assigns it to the NfTypes field.
func (o *EventFilter) SetNfTypes(v []NFType) {
	o.NfTypes = v
}

// GetNsiIdInfos returns the NsiIdInfos field value if set, zero value otherwise.
func (o *EventFilter) GetNsiIdInfos() []NsiIdInfo {
	if o == nil || o.NsiIdInfos == nil {
		var ret []NsiIdInfo
		return ret
	}
	return o.NsiIdInfos
}

// GetNsiIdInfosOk returns a tuple with the NsiIdInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNsiIdInfosOk() ([]NsiIdInfo, bool) {
	if o == nil || o.NsiIdInfos == nil {
		return nil, false
	}
	return o.NsiIdInfos, true
}

// HasNsiIdInfos returns a boolean if a field has been set.
func (o *EventFilter) HasNsiIdInfos() bool {
	if o != nil && o.NsiIdInfos != nil {
		return true
	}

	return false
}

// SetNsiIdInfos gets a reference to the given []NsiIdInfo and assigns it to the NsiIdInfos field.
func (o *EventFilter) SetNsiIdInfos(v []NsiIdInfo) {
	o.NsiIdInfos = v
}

// GetQosRequ returns the QosRequ field value if set, zero value otherwise.
func (o *EventFilter) GetQosRequ() QosRequirement {
	if o == nil || o.QosRequ == nil {
		var ret QosRequirement
		return ret
	}
	return *o.QosRequ
}

// GetQosRequOk returns a tuple with the QosRequ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetQosRequOk() (*QosRequirement, bool) {
	if o == nil || o.QosRequ == nil {
		return nil, false
	}
	return o.QosRequ, true
}

// HasQosRequ returns a boolean if a field has been set.
func (o *EventFilter) HasQosRequ() bool {
	if o != nil && o.QosRequ != nil {
		return true
	}

	return false
}

// SetQosRequ gets a reference to the given QosRequirement and assigns it to the QosRequ field.
func (o *EventFilter) SetQosRequ(v QosRequirement) {
	o.QosRequ = &v
}

// GetNwPerfTypes returns the NwPerfTypes field value if set, zero value otherwise.
func (o *EventFilter) GetNwPerfTypes() []NetworkPerfType {
	if o == nil || o.NwPerfTypes == nil {
		var ret []NetworkPerfType
		return ret
	}
	return o.NwPerfTypes
}

// GetNwPerfTypesOk returns a tuple with the NwPerfTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetNwPerfTypesOk() ([]NetworkPerfType, bool) {
	if o == nil || o.NwPerfTypes == nil {
		return nil, false
	}
	return o.NwPerfTypes, true
}

// HasNwPerfTypes returns a boolean if a field has been set.
func (o *EventFilter) HasNwPerfTypes() bool {
	if o != nil && o.NwPerfTypes != nil {
		return true
	}

	return false
}

// SetNwPerfTypes gets a reference to the given []NetworkPerfType and assigns it to the NwPerfTypes field.
func (o *EventFilter) SetNwPerfTypes(v []NetworkPerfType) {
	o.NwPerfTypes = v
}

// GetBwRequs returns the BwRequs field value if set, zero value otherwise.
func (o *EventFilter) GetBwRequs() []BwRequirement {
	if o == nil || o.BwRequs == nil {
		var ret []BwRequirement
		return ret
	}
	return o.BwRequs
}

// GetBwRequsOk returns a tuple with the BwRequs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetBwRequsOk() ([]BwRequirement, bool) {
	if o == nil || o.BwRequs == nil {
		return nil, false
	}
	return o.BwRequs, true
}

// HasBwRequs returns a boolean if a field has been set.
func (o *EventFilter) HasBwRequs() bool {
	if o != nil && o.BwRequs != nil {
		return true
	}

	return false
}

// SetBwRequs gets a reference to the given []BwRequirement and assigns it to the BwRequs field.
func (o *EventFilter) SetBwRequs(v []BwRequirement) {
	o.BwRequs = v
}

// GetExcepIds returns the ExcepIds field value if set, zero value otherwise.
func (o *EventFilter) GetExcepIds() []ExceptionId {
	if o == nil || o.ExcepIds == nil {
		var ret []ExceptionId
		return ret
	}
	return o.ExcepIds
}

// GetExcepIdsOk returns a tuple with the ExcepIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetExcepIdsOk() ([]ExceptionId, bool) {
	if o == nil || o.ExcepIds == nil {
		return nil, false
	}
	return o.ExcepIds, true
}

// HasExcepIds returns a boolean if a field has been set.
func (o *EventFilter) HasExcepIds() bool {
	if o != nil && o.ExcepIds != nil {
		return true
	}

	return false
}

// SetExcepIds gets a reference to the given []ExceptionId and assigns it to the ExcepIds field.
func (o *EventFilter) SetExcepIds(v []ExceptionId) {
	o.ExcepIds = v
}

// GetExptAnaType returns the ExptAnaType field value if set, zero value otherwise.
func (o *EventFilter) GetExptAnaType() ExpectedAnalyticsType {
	if o == nil || o.ExptAnaType == nil {
		var ret ExpectedAnalyticsType
		return ret
	}
	return *o.ExptAnaType
}

// GetExptAnaTypeOk returns a tuple with the ExptAnaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetExptAnaTypeOk() (*ExpectedAnalyticsType, bool) {
	if o == nil || o.ExptAnaType == nil {
		return nil, false
	}
	return o.ExptAnaType, true
}

// HasExptAnaType returns a boolean if a field has been set.
func (o *EventFilter) HasExptAnaType() bool {
	if o != nil && o.ExptAnaType != nil {
		return true
	}

	return false
}

// SetExptAnaType gets a reference to the given ExpectedAnalyticsType and assigns it to the ExptAnaType field.
func (o *EventFilter) SetExptAnaType(v ExpectedAnalyticsType) {
	o.ExptAnaType = &v
}

// GetExptUeBehav returns the ExptUeBehav field value if set, zero value otherwise.
func (o *EventFilter) GetExptUeBehav() ExpectedUeBehaviourData {
	if o == nil || o.ExptUeBehav == nil {
		var ret ExpectedUeBehaviourData
		return ret
	}
	return *o.ExptUeBehav
}

// GetExptUeBehavOk returns a tuple with the ExptUeBehav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetExptUeBehavOk() (*ExpectedUeBehaviourData, bool) {
	if o == nil || o.ExptUeBehav == nil {
		return nil, false
	}
	return o.ExptUeBehav, true
}

// HasExptUeBehav returns a boolean if a field has been set.
func (o *EventFilter) HasExptUeBehav() bool {
	if o != nil && o.ExptUeBehav != nil {
		return true
	}

	return false
}

// SetExptUeBehav gets a reference to the given ExpectedUeBehaviourData and assigns it to the ExptUeBehav field.
func (o *EventFilter) SetExptUeBehav(v ExpectedUeBehaviourData) {
	o.ExptUeBehav = &v
}

func (o EventFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnySlice != nil {
		toSerialize["anySlice"] = o.AnySlice
	}
	if o.Snssais != nil {
		toSerialize["snssais"] = o.Snssais
	}
	if o.AppIds != nil {
		toSerialize["appIds"] = o.AppIds
	}
	if o.Dnns != nil {
		toSerialize["dnns"] = o.Dnns
	}
	if o.Dnais != nil {
		toSerialize["dnais"] = o.Dnais
	}
	if o.NetworkArea != nil {
		toSerialize["networkArea"] = o.NetworkArea
	}
	if o.NfInstanceIds != nil {
		toSerialize["nfInstanceIds"] = o.NfInstanceIds
	}
	if o.NfSetIds != nil {
		toSerialize["nfSetIds"] = o.NfSetIds
	}
	if o.NfTypes != nil {
		toSerialize["nfTypes"] = o.NfTypes
	}
	if o.NsiIdInfos != nil {
		toSerialize["nsiIdInfos"] = o.NsiIdInfos
	}
	if o.QosRequ != nil {
		toSerialize["qosRequ"] = o.QosRequ
	}
	if o.NwPerfTypes != nil {
		toSerialize["nwPerfTypes"] = o.NwPerfTypes
	}
	if o.BwRequs != nil {
		toSerialize["bwRequs"] = o.BwRequs
	}
	if o.ExcepIds != nil {
		toSerialize["excepIds"] = o.ExcepIds
	}
	if o.ExptAnaType != nil {
		toSerialize["exptAnaType"] = o.ExptAnaType
	}
	if o.ExptUeBehav != nil {
		toSerialize["exptUeBehav"] = o.ExptUeBehav
	}
	return json.Marshal(toSerialize)
}

type NullableEventFilter struct {
	value *EventFilter
	isSet bool
}

func (v NullableEventFilter) Get() *EventFilter {
	return v.value
}

func (v *NullableEventFilter) Set(val *EventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFilter(val *EventFilter) *NullableEventFilter {
	return &NullableEventFilter{value: val, isSet: true}
}

func (v NullableEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

