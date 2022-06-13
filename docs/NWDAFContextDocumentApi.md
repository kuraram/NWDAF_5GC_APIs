# \NWDAFContextDocumentApi

All URIs are relative to *https://example.com/nnwdaf-analyticsinfo/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNwdafContext**](NWDAFContextDocumentApi.md#GetNwdafContext) | **Get** /context | Get context information related to analytics subscriptions.



## GetNwdafContext

> ContextData GetNwdafContext(ctx).ContextIds(contextIds).ReqContext(reqContext).Execute()

Get context information related to analytics subscriptions.

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
    contextIds := map[string][]openapiclient.ContextIdList{ ... } // ContextIdList | Identifies specific context information related to analytics subscriptions.
    reqContext := map[string][]openapiclient.RequestedContext{ ... } // RequestedContext | Identfies the type(s) of the analytics context information the consumer wishes to receive. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NWDAFContextDocumentApi.GetNwdafContext(context.Background()).ContextIds(contextIds).ReqContext(reqContext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NWDAFContextDocumentApi.GetNwdafContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNwdafContext`: ContextData
    fmt.Fprintf(os.Stdout, "Response from `NWDAFContextDocumentApi.GetNwdafContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNwdafContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextIds** | [**ContextIdList**](ContextIdList.md) | Identifies specific context information related to analytics subscriptions. | 
 **reqContext** | [**RequestedContext**](RequestedContext.md) | Identfies the type(s) of the analytics context information the consumer wishes to receive. | 

### Return type

[**ContextData**](ContextData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

