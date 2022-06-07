/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SliceLoadLevelInformation struct {

	// Load level information of the network slice and the optionally associated network slice instance.
	LoadLevelInformation int32 `json:"loadLevelInformation"`

	// Identification(s) of network slice to which the subscription applies.
	Snssais []Snssai `json:"snssais"`
}

// AssertSliceLoadLevelInformationRequired checks if the required fields are not zero-ed
func AssertSliceLoadLevelInformationRequired(obj SliceLoadLevelInformation) error {
	elements := map[string]interface{}{
		"loadLevelInformation": obj.LoadLevelInformation,
		"snssais": obj.Snssais,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Snssais {
		if err := AssertSnssaiRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseSliceLoadLevelInformationRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SliceLoadLevelInformation (e.g. [][]SliceLoadLevelInformation), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSliceLoadLevelInformationRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSliceLoadLevelInformation, ok := obj.(SliceLoadLevelInformation)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSliceLoadLevelInformationRequired(aSliceLoadLevelInformation)
	})
}
