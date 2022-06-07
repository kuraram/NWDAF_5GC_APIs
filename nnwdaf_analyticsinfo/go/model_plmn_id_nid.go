/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PlmnIdNid struct {

	Mcc Object `json:"mcc"`

	Mnc Object `json:"mnc"`

	Nid Object `json:"nid,omitempty"`
}

// AssertPlmnIdNidRequired checks if the required fields are not zero-ed
func AssertPlmnIdNidRequired(obj PlmnIdNid) error {
	elements := map[string]interface{}{
		"mcc": obj.Mcc,
		"mnc": obj.Mnc,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecursePlmnIdNidRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PlmnIdNid (e.g. [][]PlmnIdNid), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePlmnIdNidRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPlmnIdNid, ok := obj.(PlmnIdNid)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPlmnIdNidRequired(aPlmnIdNid)
	})
}
