/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Exception struct {

	ExcepId *ExceptionId `json:"excepId"`

	ExcepLevel int32 `json:"excepLevel,omitempty"`

	ExcepTrend *ExceptionTrend `json:"excepTrend,omitempty"`
}