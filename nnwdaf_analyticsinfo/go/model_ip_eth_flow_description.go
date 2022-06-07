/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IpEthFlowDescription struct {

	// Defines a packet filter of an IP flow.
	IpTrafficFilter string `json:"ipTrafficFilter,omitempty"`

	EthTrafficFilter EthFlowDescription `json:"ethTrafficFilter,omitempty"`
}

// AssertIpEthFlowDescriptionRequired checks if the required fields are not zero-ed
func AssertIpEthFlowDescriptionRequired(obj IpEthFlowDescription) error {
	if err := AssertEthFlowDescriptionRequired(obj.EthTrafficFilter); err != nil {
		return err
	}
	return nil
}

// AssertRecurseIpEthFlowDescriptionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IpEthFlowDescription (e.g. [][]IpEthFlowDescription), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIpEthFlowDescriptionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIpEthFlowDescription, ok := obj.(IpEthFlowDescription)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIpEthFlowDescriptionRequired(aIpEthFlowDescription)
	})
}
