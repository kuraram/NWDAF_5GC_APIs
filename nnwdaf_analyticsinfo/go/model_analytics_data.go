/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type AnalyticsData struct {

	Expiry time.Time `json:"expiry,omitempty"`

	TimeStampGen time.Time `json:"timeStampGen,omitempty"`

	// The slices and their load level information.
	SliceLoadLevelInfos []SliceLoadLevelInformation `json:"sliceLoadLevelInfos,omitempty"`

	NsiLoadLevelInfos []NsiLoadLevelInfo `json:"nsiLoadLevelInfos,omitempty"`

	NfLoadLevelInfos []NfLoadLevelInformation `json:"nfLoadLevelInfos,omitempty"`

	NwPerfs []NetworkPerfInfo `json:"nwPerfs,omitempty"`

	SvcExps []ServiceExperienceInfo `json:"svcExps,omitempty"`

	QosSustainInfos []QosSustainabilityInfo `json:"qosSustainInfos,omitempty"`

	UeMobs []UeMobility `json:"ueMobs,omitempty"`

	UeComms []UeCommunication `json:"ueComms,omitempty"`

	UserDataCongInfos []UserDataCongestionInfo `json:"userDataCongInfos,omitempty"`

	AbnorBehavrs []AbnormalBehaviour `json:"abnorBehavrs,omitempty"`

	SuppFeat string `json:"suppFeat,omitempty"`
}

// AssertAnalyticsDataRequired checks if the required fields are not zero-ed
func AssertAnalyticsDataRequired(obj AnalyticsData) error {
	for _, el := range obj.SliceLoadLevelInfos {
		if err := AssertSliceLoadLevelInformationRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NsiLoadLevelInfos {
		if err := AssertNsiLoadLevelInfoRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NfLoadLevelInfos {
		if err := AssertNfLoadLevelInformationRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NwPerfs {
		if err := AssertNetworkPerfInfoRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.SvcExps {
		if err := AssertServiceExperienceInfoRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.QosSustainInfos {
		if err := AssertQosSustainabilityInfoRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.UeMobs {
		if err := AssertUeMobilityRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.UeComms {
		if err := AssertUeCommunicationRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.UserDataCongInfos {
		if err := AssertUserDataCongestionInfoRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AbnorBehavrs {
		if err := AssertAbnormalBehaviourRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseAnalyticsDataRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AnalyticsData (e.g. [][]AnalyticsData), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAnalyticsDataRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAnalyticsData, ok := obj.(AnalyticsData)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAnalyticsDataRequired(aAnalyticsData)
	})
}
