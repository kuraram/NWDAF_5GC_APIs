/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LocationArea struct {
	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []Ts29572NlmfLocationYamlcomponentsschemasGeographicArea `json:"geographicAreas,omitempty"`
	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []Ts29572NlmfLocationYamlcomponentsschemasCivicAddress `json:"civicAddresses,omitempty"`

	NwAreaInfo *Ts29503NudmPpYamlcomponentsschemasNetworkAreaInfo `json:"nwAreaInfo,omitempty"`
}
