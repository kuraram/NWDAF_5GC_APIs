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

// SpecificAnalyticsSubscription1 Represents an existing subscription for a specific type of analytics to a specific NWDAF.
type SpecificAnalyticsSubscription1 struct {
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	ProducerId *string `json:"producerId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"  set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or \"set<SetID>.  <NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with <MCC> encoded as defined in clause 5.4.2  (\"Mcc\" data type definition) <MNC> encoded as defined in clause 5.4.2 (\"Mnc\" data type  definition) <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but  with lower case characters <Set ID> encoded as a string of characters consisting of alphabetic  characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end with either an  alphabetic character or a digit.  
	ProducerSetId *string `json:"producerSetId,omitempty"`
	NwdafEvSub *NnwdafEventsSubscription `json:"nwdafEvSub,omitempty"`
}

// NewSpecificAnalyticsSubscription1 instantiates a new SpecificAnalyticsSubscription1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecificAnalyticsSubscription1() *SpecificAnalyticsSubscription1 {
	this := SpecificAnalyticsSubscription1{}
	return &this
}

// NewSpecificAnalyticsSubscription1WithDefaults instantiates a new SpecificAnalyticsSubscription1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecificAnalyticsSubscription1WithDefaults() *SpecificAnalyticsSubscription1 {
	this := SpecificAnalyticsSubscription1{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SpecificAnalyticsSubscription1) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecificAnalyticsSubscription1) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SpecificAnalyticsSubscription1) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *SpecificAnalyticsSubscription1) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetProducerId returns the ProducerId field value if set, zero value otherwise.
func (o *SpecificAnalyticsSubscription1) GetProducerId() string {
	if o == nil || o.ProducerId == nil {
		var ret string
		return ret
	}
	return *o.ProducerId
}

// GetProducerIdOk returns a tuple with the ProducerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecificAnalyticsSubscription1) GetProducerIdOk() (*string, bool) {
	if o == nil || o.ProducerId == nil {
		return nil, false
	}
	return o.ProducerId, true
}

// HasProducerId returns a boolean if a field has been set.
func (o *SpecificAnalyticsSubscription1) HasProducerId() bool {
	if o != nil && o.ProducerId != nil {
		return true
	}

	return false
}

// SetProducerId gets a reference to the given string and assigns it to the ProducerId field.
func (o *SpecificAnalyticsSubscription1) SetProducerId(v string) {
	o.ProducerId = &v
}

// GetProducerSetId returns the ProducerSetId field value if set, zero value otherwise.
func (o *SpecificAnalyticsSubscription1) GetProducerSetId() string {
	if o == nil || o.ProducerSetId == nil {
		var ret string
		return ret
	}
	return *o.ProducerSetId
}

// GetProducerSetIdOk returns a tuple with the ProducerSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecificAnalyticsSubscription1) GetProducerSetIdOk() (*string, bool) {
	if o == nil || o.ProducerSetId == nil {
		return nil, false
	}
	return o.ProducerSetId, true
}

// HasProducerSetId returns a boolean if a field has been set.
func (o *SpecificAnalyticsSubscription1) HasProducerSetId() bool {
	if o != nil && o.ProducerSetId != nil {
		return true
	}

	return false
}

// SetProducerSetId gets a reference to the given string and assigns it to the ProducerSetId field.
func (o *SpecificAnalyticsSubscription1) SetProducerSetId(v string) {
	o.ProducerSetId = &v
}

// GetNwdafEvSub returns the NwdafEvSub field value if set, zero value otherwise.
func (o *SpecificAnalyticsSubscription1) GetNwdafEvSub() NnwdafEventsSubscription {
	if o == nil || o.NwdafEvSub == nil {
		var ret NnwdafEventsSubscription
		return ret
	}
	return *o.NwdafEvSub
}

// GetNwdafEvSubOk returns a tuple with the NwdafEvSub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecificAnalyticsSubscription1) GetNwdafEvSubOk() (*NnwdafEventsSubscription, bool) {
	if o == nil || o.NwdafEvSub == nil {
		return nil, false
	}
	return o.NwdafEvSub, true
}

// HasNwdafEvSub returns a boolean if a field has been set.
func (o *SpecificAnalyticsSubscription1) HasNwdafEvSub() bool {
	if o != nil && o.NwdafEvSub != nil {
		return true
	}

	return false
}

// SetNwdafEvSub gets a reference to the given NnwdafEventsSubscription and assigns it to the NwdafEvSub field.
func (o *SpecificAnalyticsSubscription1) SetNwdafEvSub(v NnwdafEventsSubscription) {
	o.NwdafEvSub = &v
}

func (o SpecificAnalyticsSubscription1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.ProducerId != nil {
		toSerialize["producerId"] = o.ProducerId
	}
	if o.ProducerSetId != nil {
		toSerialize["producerSetId"] = o.ProducerSetId
	}
	if o.NwdafEvSub != nil {
		toSerialize["nwdafEvSub"] = o.NwdafEvSub
	}
	return json.Marshal(toSerialize)
}

type NullableSpecificAnalyticsSubscription1 struct {
	value *SpecificAnalyticsSubscription1
	isSet bool
}

func (v NullableSpecificAnalyticsSubscription1) Get() *SpecificAnalyticsSubscription1 {
	return v.value
}

func (v *NullableSpecificAnalyticsSubscription1) Set(val *SpecificAnalyticsSubscription1) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecificAnalyticsSubscription1) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecificAnalyticsSubscription1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecificAnalyticsSubscription1(val *SpecificAnalyticsSubscription1) *NullableSpecificAnalyticsSubscription1 {
	return &NullableSpecificAnalyticsSubscription1{value: val, isSet: true}
}

func (v NullableSpecificAnalyticsSubscription1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecificAnalyticsSubscription1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


