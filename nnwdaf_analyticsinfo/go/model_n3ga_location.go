/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type N3gaLocation struct {

	N3gppTai *Tai `json:"n3gppTai,omitempty"`

	N3IwfId string `json:"n3IwfId,omitempty"`

	UeIpv4Addr string `json:"ueIpv4Addr,omitempty"`

	UeIpv6Addr *Ipv6Addr `json:"ueIpv6Addr,omitempty"`

	PortNumber int32 `json:"portNumber,omitempty"`

	TnapId *TnapId `json:"tnapId,omitempty"`

	Protocol *TransportProtocol `json:"protocol,omitempty"`

	TwapId *TwapId `json:"twapId,omitempty"`

	HfcNodeId *HfcNodeId `json:"hfcNodeId,omitempty"`

	Gli *Bytes `json:"gli,omitempty"`

	W5gbanLineType *LineType `json:"w5gbanLineType,omitempty"`

	Gci string `json:"gci,omitempty"`
}
