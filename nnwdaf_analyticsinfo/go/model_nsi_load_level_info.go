/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NsiLoadLevelInfo - Represents the network slice and optionally the associated network slice instance and the load level information.
type NsiLoadLevelInfo struct {

	// Load level information of the network slice and the optionally associated network slice instance.
	LoadLevelInformation int32 `json:"loadLevelInformation"`

	Snssai Snssai `json:"snssai"`

	NsiId string `json:"nsiId,omitempty"`
}

// AssertNsiLoadLevelInfoRequired checks if the required fields are not zero-ed
func AssertNsiLoadLevelInfoRequired(obj NsiLoadLevelInfo) error {
	elements := map[string]interface{}{
		"loadLevelInformation": obj.LoadLevelInformation,
		"snssai": obj.Snssai,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	return nil
}

// AssertRecurseNsiLoadLevelInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of NsiLoadLevelInfo (e.g. [][]NsiLoadLevelInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseNsiLoadLevelInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aNsiLoadLevelInfo, ok := obj.(NsiLoadLevelInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertNsiLoadLevelInfoRequired(aNsiLoadLevelInfo)
	})
}
