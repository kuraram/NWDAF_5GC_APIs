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

type UtraLocation struct {

	Cgi *CellGlobalId `json:"cgi,omitempty"`

	Sai *ServiceAreaId `json:"sai,omitempty"`

	Lai *LocationAreaId `json:"lai,omitempty"`

	Rai *RoutingAreaId `json:"rai,omitempty"`

	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`

	UeLocationTimestamp *time.Time `json:"ueLocationTimestamp,omitempty"`

	GeographicalInformation string `json:"geographicalInformation,omitempty"`

	GeodeticInformation string `json:"geodeticInformation,omitempty"`
}