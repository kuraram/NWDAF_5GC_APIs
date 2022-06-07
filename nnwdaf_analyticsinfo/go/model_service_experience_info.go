/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ServiceExperienceInfo struct {

	SvcExprc *SvcExperience `json:"svcExprc"`

	SvcExprcVariance float32 `json:"svcExprcVariance,omitempty"`

	Supis []string `json:"supis,omitempty"`

	Snssai *Snssai `json:"snssai,omitempty"`

	AppId string `json:"appId,omitempty"`

	Confidence int32 `json:"confidence,omitempty"`

	Dnn string `json:"dnn,omitempty"`

	NetworkArea *NetworkAreaInfo `json:"networkArea,omitempty"`

	NsiId string `json:"nsiId,omitempty"`

	Ratio int32 `json:"ratio,omitempty"`
}
