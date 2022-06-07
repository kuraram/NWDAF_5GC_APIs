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
	"context"
	"net/http"
	"errors"
)

// NWDAFAnalyticsDocumentApiService is a service that implements the logic for the NWDAFAnalyticsDocumentApiServicer
// This service should implement the business logic for every endpoint for the NWDAFAnalyticsDocumentApi API.
// Include any external packages or services that will be required by this service.
type NWDAFAnalyticsDocumentApiService struct {
}

// NewNWDAFAnalyticsDocumentApiService creates a default api service
func NewNWDAFAnalyticsDocumentApiService() NWDAFAnalyticsDocumentApiServicer {
	return &NWDAFAnalyticsDocumentApiService{}
}

// GetNWDAFAnalytics - Read a NWDAF Analytics
func (s *NWDAFAnalyticsDocumentApiService) GetNWDAFAnalytics(ctx context.Context, eventId EventId, anaReq EventReportingRequirement, eventFilter EventFilter, supportedFeatures string, tgtUe TargetUeInformation) (ImplResponse, error) {
	// TODO - update GetNWDAFAnalytics with the required logic for this service method.
	// Add api_nwdaf_analytics_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, AnalyticsData{}) or use other options such as http.Ok ...
	//return Response(200, AnalyticsData{}), nil

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(400, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(401, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(401, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(403, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(404, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(406, {}) or use other options such as http.Ok ...
	//return Response(406, nil),nil

	//TODO: Uncomment the next line to return response Response(414, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(414, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(429, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(429, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(500, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(503, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(503, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	//return Response(0, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetNWDAFAnalytics method not implemented")
}