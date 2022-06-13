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

// AdrfDataTypeAnyOf the model 'AdrfDataTypeAnyOf'
type AdrfDataTypeAnyOf string

// List of AdrfDataType_anyOf
const (
	ANALYTICS AdrfDataTypeAnyOf = "HISTORICAL_ANALYTICS"
	DATA AdrfDataTypeAnyOf = "HISTORICAL_DATA"
)

// All allowed values of AdrfDataTypeAnyOf enum
var AllowedAdrfDataTypeAnyOfEnumValues = []AdrfDataTypeAnyOf{
	"HISTORICAL_ANALYTICS",
	"HISTORICAL_DATA",
}

func (v *AdrfDataTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdrfDataTypeAnyOf(value)
	for _, existing := range AllowedAdrfDataTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdrfDataTypeAnyOf", value)
}

// NewAdrfDataTypeAnyOfFromValue returns a pointer to a valid AdrfDataTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdrfDataTypeAnyOfFromValue(v string) (*AdrfDataTypeAnyOf, error) {
	ev := AdrfDataTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdrfDataTypeAnyOf: valid values are %v", v, AllowedAdrfDataTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdrfDataTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedAdrfDataTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdrfDataType_anyOf value
func (v AdrfDataTypeAnyOf) Ptr() *AdrfDataTypeAnyOf {
	return &v
}

type NullableAdrfDataTypeAnyOf struct {
	value *AdrfDataTypeAnyOf
	isSet bool
}

func (v NullableAdrfDataTypeAnyOf) Get() *AdrfDataTypeAnyOf {
	return v.value
}

func (v *NullableAdrfDataTypeAnyOf) Set(val *AdrfDataTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAdrfDataTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAdrfDataTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdrfDataTypeAnyOf(val *AdrfDataTypeAnyOf) *NullableAdrfDataTypeAnyOf {
	return &NullableAdrfDataTypeAnyOf{value: val, isSet: true}
}

func (v NullableAdrfDataTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdrfDataTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

