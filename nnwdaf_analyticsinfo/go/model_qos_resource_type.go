/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type QosResourceType struct {
}

// AssertQosResourceTypeRequired checks if the required fields are not zero-ed
func AssertQosResourceTypeRequired(obj QosResourceType) error {
	return nil
}

// AssertRecurseQosResourceTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of QosResourceType (e.g. [][]QosResourceType), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseQosResourceTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aQosResourceType, ok := obj.(QosResourceType)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertQosResourceTypeRequired(aQosResourceType)
	})
}
