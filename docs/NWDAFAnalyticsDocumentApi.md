# \NWDAFAnalyticsDocumentApi

All URIs are relative to *https://example.com/nnwdaf-analyticsinfo/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNWDAFAnalytics**](NWDAFAnalyticsDocumentApi.md#GetNWDAFAnalytics) | **Get** /analytics | Read a NWDAF Analytics



## GetNWDAFAnalytics

> AnalyticsData GetNWDAFAnalytics(ctx).EventId(eventId).AnaReq(anaReq).EventFilter(eventFilter).SupportedFeatures(supportedFeatures).TgtUe(tgtUe).Execute()

Read a NWDAF Analytics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    eventId := *openapiclient.NewEventId() // EventId | Identify the analytics.
    anaReq := map[string][]openapiclient.EventReportingRequirement{ ... } // EventReportingRequirement | Identifies the analytics reporting requirement information. (optional)
    eventFilter := map[string][]openapiclient.EventFilter{ ... } // EventFilter | Identify the analytics. (optional)
    supportedFeatures := "supportedFeatures_example" // string | To filter irrelevant responses related to unsupported features (optional)
    tgtUe := map[string][]openapiclient.TargetUeInformation{ ... } // TargetUeInformation | Identify the target UE information. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NWDAFAnalyticsDocumentApi.GetNWDAFAnalytics(context.Background()).EventId(eventId).AnaReq(anaReq).EventFilter(eventFilter).SupportedFeatures(supportedFeatures).TgtUe(tgtUe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NWDAFAnalyticsDocumentApi.GetNWDAFAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNWDAFAnalytics`: AnalyticsData
    fmt.Fprintf(os.Stdout, "Response from `NWDAFAnalyticsDocumentApi.GetNWDAFAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNWDAFAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventId** | [**EventId**](EventId.md) | Identify the analytics. | 
 **anaReq** | [**EventReportingRequirement**](EventReportingRequirement.md) | Identifies the analytics reporting requirement information. | 
 **eventFilter** | [**EventFilter**](EventFilter.md) | Identify the analytics. | 
 **supportedFeatures** | **string** | To filter irrelevant responses related to unsupported features | 
 **tgtUe** | [**TargetUeInformation**](TargetUeInformation.md) | Identify the target UE information. | 

### Return type

[**AnalyticsData**](AnalyticsData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

