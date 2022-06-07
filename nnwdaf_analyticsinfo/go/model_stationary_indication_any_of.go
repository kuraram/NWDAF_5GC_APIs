/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type StationaryIndicationAnyOf string

// List of StationaryIndicationAnyOf
const (
	STATIONARY StationaryIndicationAnyOf = "STATIONARY"
	MOBILE StationaryIndicationAnyOf = "MOBILE"
)

// AssertStationaryIndicationAnyOfRequired checks if the required fields are not zero-ed
func AssertStationaryIndicationAnyOfRequired(obj StationaryIndicationAnyOf) error {
	return nil
}

// AssertRecurseStationaryIndicationAnyOfRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of StationaryIndicationAnyOf (e.g. [][]StationaryIndicationAnyOf), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseStationaryIndicationAnyOfRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aStationaryIndicationAnyOf, ok := obj.(StationaryIndicationAnyOf)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertStationaryIndicationAnyOfRequired(aStationaryIndicationAnyOf)
	})
}
