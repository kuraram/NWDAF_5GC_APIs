/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ExpectedAnalyticsType Possible values are   - MOBILITY: Mobility related abnormal behaviour analytics is expected by the consumer.   - COMMUN: Communication related abnormal behaviour analytics is expected by the consumer.   - MOBILITY_AND_COMMUN: Both mobility and communication related abnormal behaviour analytics is expected by the consumer. 
type ExpectedAnalyticsType struct {
	ExpectedAnalyticsTypeAnyOf *ExpectedAnalyticsTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ExpectedAnalyticsType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ExpectedAnalyticsTypeAnyOf
	err = json.Unmarshal(data, &dst.ExpectedAnalyticsTypeAnyOf);
	if err == nil {
		jsonExpectedAnalyticsTypeAnyOf, _ := json.Marshal(dst.ExpectedAnalyticsTypeAnyOf)
		if string(jsonExpectedAnalyticsTypeAnyOf) == "{}" { // empty struct
			dst.ExpectedAnalyticsTypeAnyOf = nil
		} else {
			return nil // data stored in dst.ExpectedAnalyticsTypeAnyOf, return on the first match
		}
	} else {
		dst.ExpectedAnalyticsTypeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(ExpectedAnalyticsType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ExpectedAnalyticsType) MarshalJSON() ([]byte, error) {
	if src.ExpectedAnalyticsTypeAnyOf != nil {
		return json.Marshal(&src.ExpectedAnalyticsTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableExpectedAnalyticsType struct {
	value *ExpectedAnalyticsType
	isSet bool
}

func (v NullableExpectedAnalyticsType) Get() *ExpectedAnalyticsType {
	return v.value
}

func (v *NullableExpectedAnalyticsType) Set(val *ExpectedAnalyticsType) {
	v.value = val
	v.isSet = true
}

func (v NullableExpectedAnalyticsType) IsSet() bool {
	return v.isSet
}

func (v *NullableExpectedAnalyticsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpectedAnalyticsType(val *ExpectedAnalyticsType) *NullableExpectedAnalyticsType {
	return &NullableExpectedAnalyticsType{value: val, isSet: true}
}

func (v NullableExpectedAnalyticsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpectedAnalyticsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


