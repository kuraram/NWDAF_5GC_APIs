# ProblemDetailsAnalyticsInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | String providing an URI formatted according to RFC 3986 | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**Instance** | Pointer to **string** | String providing an URI formatted according to RFC 3986 | [optional] 
**Cause** | Pointer to **string** | A machine-readable application error cause specific to this occurrence of the problem.  This IE should be present and provide application-related error information, if available.  | [optional] 
**InvalidParams** | Pointer to [**[]InvalidParam**](InvalidParam.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AccessTokenError** | Pointer to [**AccessTokenErr**](AccessTokenErr.md) |  | [optional] 
**AccessTokenRequest** | Pointer to [**AccessTokenReq**](AccessTokenReq.md) |  | [optional] 
**NrfId** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**RvWaitTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewProblemDetailsAnalyticsInfoRequest

`func NewProblemDetailsAnalyticsInfoRequest() *ProblemDetailsAnalyticsInfoRequest`

NewProblemDetailsAnalyticsInfoRequest instantiates a new ProblemDetailsAnalyticsInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemDetailsAnalyticsInfoRequestWithDefaults

`func NewProblemDetailsAnalyticsInfoRequestWithDefaults() *ProblemDetailsAnalyticsInfoRequest`

NewProblemDetailsAnalyticsInfoRequestWithDefaults instantiates a new ProblemDetailsAnalyticsInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProblemDetailsAnalyticsInfoRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProblemDetailsAnalyticsInfoRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProblemDetailsAnalyticsInfoRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *ProblemDetailsAnalyticsInfoRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProblemDetailsAnalyticsInfoRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProblemDetailsAnalyticsInfoRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *ProblemDetailsAnalyticsInfoRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProblemDetailsAnalyticsInfoRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProblemDetailsAnalyticsInfoRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *ProblemDetailsAnalyticsInfoRequest) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ProblemDetailsAnalyticsInfoRequest) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ProblemDetailsAnalyticsInfoRequest) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *ProblemDetailsAnalyticsInfoRequest) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ProblemDetailsAnalyticsInfoRequest) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ProblemDetailsAnalyticsInfoRequest) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCause

`func (o *ProblemDetailsAnalyticsInfoRequest) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ProblemDetailsAnalyticsInfoRequest) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *ProblemDetailsAnalyticsInfoRequest) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetInvalidParams

`func (o *ProblemDetailsAnalyticsInfoRequest) GetInvalidParams() []InvalidParam`

GetInvalidParams returns the InvalidParams field if non-nil, zero value otherwise.

### GetInvalidParamsOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetInvalidParamsOk() (*[]InvalidParam, bool)`

GetInvalidParamsOk returns a tuple with the InvalidParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParams

`func (o *ProblemDetailsAnalyticsInfoRequest) SetInvalidParams(v []InvalidParam)`

SetInvalidParams sets InvalidParams field to given value.

### HasInvalidParams

`func (o *ProblemDetailsAnalyticsInfoRequest) HasInvalidParams() bool`

HasInvalidParams returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ProblemDetailsAnalyticsInfoRequest) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ProblemDetailsAnalyticsInfoRequest) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ProblemDetailsAnalyticsInfoRequest) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAccessTokenError

`func (o *ProblemDetailsAnalyticsInfoRequest) GetAccessTokenError() AccessTokenErr`

GetAccessTokenError returns the AccessTokenError field if non-nil, zero value otherwise.

### GetAccessTokenErrorOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetAccessTokenErrorOk() (*AccessTokenErr, bool)`

GetAccessTokenErrorOk returns a tuple with the AccessTokenError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenError

`func (o *ProblemDetailsAnalyticsInfoRequest) SetAccessTokenError(v AccessTokenErr)`

SetAccessTokenError sets AccessTokenError field to given value.

### HasAccessTokenError

`func (o *ProblemDetailsAnalyticsInfoRequest) HasAccessTokenError() bool`

HasAccessTokenError returns a boolean if a field has been set.

### GetAccessTokenRequest

`func (o *ProblemDetailsAnalyticsInfoRequest) GetAccessTokenRequest() AccessTokenReq`

GetAccessTokenRequest returns the AccessTokenRequest field if non-nil, zero value otherwise.

### GetAccessTokenRequestOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetAccessTokenRequestOk() (*AccessTokenReq, bool)`

GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenRequest

`func (o *ProblemDetailsAnalyticsInfoRequest) SetAccessTokenRequest(v AccessTokenReq)`

SetAccessTokenRequest sets AccessTokenRequest field to given value.

### HasAccessTokenRequest

`func (o *ProblemDetailsAnalyticsInfoRequest) HasAccessTokenRequest() bool`

HasAccessTokenRequest returns a boolean if a field has been set.

### GetNrfId

`func (o *ProblemDetailsAnalyticsInfoRequest) GetNrfId() string`

GetNrfId returns the NrfId field if non-nil, zero value otherwise.

### GetNrfIdOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetNrfIdOk() (*string, bool)`

GetNrfIdOk returns a tuple with the NrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfId

`func (o *ProblemDetailsAnalyticsInfoRequest) SetNrfId(v string)`

SetNrfId sets NrfId field to given value.

### HasNrfId

`func (o *ProblemDetailsAnalyticsInfoRequest) HasNrfId() bool`

HasNrfId returns a boolean if a field has been set.

### GetRvWaitTime

`func (o *ProblemDetailsAnalyticsInfoRequest) GetRvWaitTime() int32`

GetRvWaitTime returns the RvWaitTime field if non-nil, zero value otherwise.

### GetRvWaitTimeOk

`func (o *ProblemDetailsAnalyticsInfoRequest) GetRvWaitTimeOk() (*int32, bool)`

GetRvWaitTimeOk returns a tuple with the RvWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRvWaitTime

`func (o *ProblemDetailsAnalyticsInfoRequest) SetRvWaitTime(v int32)`

SetRvWaitTime sets RvWaitTime field to given value.

### HasRvWaitTime

`func (o *ProblemDetailsAnalyticsInfoRequest) HasRvWaitTime() bool`

HasRvWaitTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


