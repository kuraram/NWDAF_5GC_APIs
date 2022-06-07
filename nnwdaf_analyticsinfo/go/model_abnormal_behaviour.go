/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AbnormalBehaviour struct {

	Supis []string `json:"supis,omitempty"`

	Excep *Exception `json:"excep"`

	Dnn string `json:"dnn,omitempty"`

	Snssai *Snssai `json:"snssai,omitempty"`

	Ratio int32 `json:"ratio,omitempty"`

	Confidence int32 `json:"confidence,omitempty"`

	AddtMeasInfo *AdditionalMeasurement `json:"addtMeasInfo,omitempty"`
}