/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// EventFilter - Represents the event filters used to identify the requested analytics.
type EventFilter struct {

	// FALSE represents not applicable for all slices. TRUE represents applicable for all slices.
	AnySlice bool `json:"anySlice,omitempty"`

	// Identification(s) of network slice to which the subscription belongs.
	Snssais []Snssai `json:"snssais,omitempty"`

	AppIds []string `json:"appIds,omitempty"`

	Dnns []string `json:"dnns,omitempty"`

	Dnais []string `json:"dnais,omitempty"`

	NetworkArea NetworkAreaInfo `json:"networkArea,omitempty"`

	NfInstanceIds []string `json:"nfInstanceIds,omitempty"`

	NfSetIds []string `json:"nfSetIds,omitempty"`

	NfTypes []NfType `json:"nfTypes,omitempty"`

	NsiIdInfos []NsiIdInfo `json:"nsiIdInfos,omitempty"`

	QosRequ QosRequirement `json:"qosRequ,omitempty"`

	NwPerfTypes []NetworkPerfType `json:"nwPerfTypes,omitempty"`

	BwRequs []BwRequirement `json:"bwRequs,omitempty"`

	ExcepIds []ExceptionId `json:"excepIds,omitempty"`

	ExptAnaType ExpectedAnalyticsType `json:"exptAnaType,omitempty"`

	ExptUeBehav ExpectedUeBehaviourData `json:"exptUeBehav,omitempty"`
}

// AssertEventFilterRequired checks if the required fields are not zero-ed
func AssertEventFilterRequired(obj EventFilter) error {
	for _, el := range obj.Snssais {
		if err := AssertSnssaiRequired(el); err != nil {
			return err
		}
	}
	if err := AssertNetworkAreaInfoRequired(obj.NetworkArea); err != nil {
		return err
	}
	for _, el := range obj.NfTypes {
		if err := AssertNfTypeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NsiIdInfos {
		if err := AssertNsiIdInfoRequired(el); err != nil {
			return err
		}
	}
	if err := AssertQosRequirementRequired(obj.QosRequ); err != nil {
		return err
	}
	for _, el := range obj.NwPerfTypes {
		if err := AssertNetworkPerfTypeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.BwRequs {
		if err := AssertBwRequirementRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ExcepIds {
		if err := AssertExceptionIdRequired(el); err != nil {
			return err
		}
	}
	if err := AssertExpectedAnalyticsTypeRequired(obj.ExptAnaType); err != nil {
		return err
	}
	if err := AssertExpectedUeBehaviourDataRequired(obj.ExptUeBehav); err != nil {
		return err
	}
	return nil
}

// AssertRecurseEventFilterRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of EventFilter (e.g. [][]EventFilter), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseEventFilterRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aEventFilter, ok := obj.(EventFilter)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertEventFilterRequired(aEventFilter)
	})
}
