# AccessTokenErr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**ErrorUri** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessTokenErr

`func NewAccessTokenErr(error_ string, ) *AccessTokenErr`

NewAccessTokenErr instantiates a new AccessTokenErr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenErrWithDefaults

`func NewAccessTokenErrWithDefaults() *AccessTokenErr`

NewAccessTokenErrWithDefaults instantiates a new AccessTokenErr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AccessTokenErr) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AccessTokenErr) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AccessTokenErr) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *AccessTokenErr) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *AccessTokenErr) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *AccessTokenErr) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *AccessTokenErr) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorUri

`func (o *AccessTokenErr) GetErrorUri() string`

GetErrorUri returns the ErrorUri field if non-nil, zero value otherwise.

### GetErrorUriOk

`func (o *AccessTokenErr) GetErrorUriOk() (*string, bool)`

GetErrorUriOk returns a tuple with the ErrorUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUri

`func (o *AccessTokenErr) SetErrorUri(v string)`

SetErrorUri sets ErrorUri field to given value.

### HasErrorUri

`func (o *AccessTokenErr) HasErrorUri() bool`

HasErrorUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


