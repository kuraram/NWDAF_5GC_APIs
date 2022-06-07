/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type QosRequirement struct {

	Var5qi int32 `json:"5qi,omitempty"`

	GfbrUl string `json:"gfbrUl,omitempty"`

	GfbrDl string `json:"gfbrDl,omitempty"`

	ResType *QosResourceType `json:"resType,omitempty"`

	Pdb int32 `json:"pdb,omitempty"`

	Per string `json:"per,omitempty"`
}