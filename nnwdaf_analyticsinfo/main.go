/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	NWDAFAnalyticsDocumentApiService := openapi.NewNWDAFAnalyticsDocumentApiService()
	NWDAFAnalyticsDocumentApiController := openapi.NewNWDAFAnalyticsDocumentApiController(NWDAFAnalyticsDocumentApiService)

	router := openapi.NewRouter(NWDAFAnalyticsDocumentApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
