/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AccuracyAnyOf the model 'AccuracyAnyOf'
type AccuracyAnyOf string

// List of Accuracy_anyOf
const (
	LOW AccuracyAnyOf = "LOW"
	HIGH AccuracyAnyOf = "HIGH"
)

// All allowed values of AccuracyAnyOf enum
var AllowedAccuracyAnyOfEnumValues = []AccuracyAnyOf{
	"LOW",
	"HIGH",
}

func (v *AccuracyAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccuracyAnyOf(value)
	for _, existing := range AllowedAccuracyAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccuracyAnyOf", value)
}

// NewAccuracyAnyOfFromValue returns a pointer to a valid AccuracyAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccuracyAnyOfFromValue(v string) (*AccuracyAnyOf, error) {
	ev := AccuracyAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccuracyAnyOf: valid values are %v", v, AllowedAccuracyAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccuracyAnyOf) IsValid() bool {
	for _, existing := range AllowedAccuracyAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Accuracy_anyOf value
func (v AccuracyAnyOf) Ptr() *AccuracyAnyOf {
	return &v
}

type NullableAccuracyAnyOf struct {
	value *AccuracyAnyOf
	isSet bool
}

func (v NullableAccuracyAnyOf) Get() *AccuracyAnyOf {
	return v.value
}

func (v *NullableAccuracyAnyOf) Set(val *AccuracyAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccuracyAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccuracyAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccuracyAnyOf(val *AccuracyAnyOf) *NullableAccuracyAnyOf {
	return &NullableAccuracyAnyOf{value: val, isSet: true}
}

func (v NullableAccuracyAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccuracyAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

