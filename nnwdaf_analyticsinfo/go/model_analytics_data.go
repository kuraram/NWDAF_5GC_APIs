/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type AnalyticsData struct {

	Expiry *time.Time `json:"expiry,omitempty"`

	TimeStampGen *time.Time `json:"timeStampGen,omitempty"`
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
