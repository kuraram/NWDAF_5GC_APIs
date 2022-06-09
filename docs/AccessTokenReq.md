# AccessTokenReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** |  | 
**NfInstanceId** | **string** |  | 
**NfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**TargetNfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**Scope** | **string** |  | 
**TargetNfInstanceId** | Pointer to **string** |  | [optional] 
**RequesterPlmn** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**RequesterPlmnList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**RequesterSnssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**RequesterFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**RequesterSnpnList** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**TargetPlmn** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**TargetSnssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**TargetNsiList** | Pointer to **[]string** |  | [optional] 
**TargetNfSetId** | Pointer to **string** |  | [optional] 
**TargetNfServiceSetId** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessTokenReq

`func NewAccessTokenReq(grantType string, nfInstanceId string, scope string, ) *AccessTokenReq`

NewAccessTokenReq instantiates a new AccessTokenReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenReqWithDefaults

`func NewAccessTokenReqWithDefaults() *AccessTokenReq`

NewAccessTokenReqWithDefaults instantiates a new AccessTokenReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *AccessTokenReq) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *AccessTokenReq) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *AccessTokenReq) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetNfInstanceId

`func (o *AccessTokenReq) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *AccessTokenReq) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *AccessTokenReq) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetNfType

`func (o *AccessTokenReq) GetNfType() NFType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *AccessTokenReq) GetNfTypeOk() (*NFType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *AccessTokenReq) SetNfType(v NFType)`

SetNfType sets NfType field to given value.

### HasNfType

`func (o *AccessTokenReq) HasNfType() bool`

HasNfType returns a boolean if a field has been set.

### GetTargetNfType

`func (o *AccessTokenReq) GetTargetNfType() NFType`

GetTargetNfType returns the TargetNfType field if non-nil, zero value otherwise.

### GetTargetNfTypeOk

`func (o *AccessTokenReq) GetTargetNfTypeOk() (*NFType, bool)`

GetTargetNfTypeOk returns a tuple with the TargetNfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfType

`func (o *AccessTokenReq) SetTargetNfType(v NFType)`

SetTargetNfType sets TargetNfType field to given value.

### HasTargetNfType

`func (o *AccessTokenReq) HasTargetNfType() bool`

HasTargetNfType returns a boolean if a field has been set.

### GetScope

`func (o *AccessTokenReq) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccessTokenReq) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccessTokenReq) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTargetNfInstanceId

`func (o *AccessTokenReq) GetTargetNfInstanceId() string`

GetTargetNfInstanceId returns the TargetNfInstanceId field if non-nil, zero value otherwise.

### GetTargetNfInstanceIdOk

`func (o *AccessTokenReq) GetTargetNfInstanceIdOk() (*string, bool)`

GetTargetNfInstanceIdOk returns a tuple with the TargetNfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfInstanceId

`func (o *AccessTokenReq) SetTargetNfInstanceId(v string)`

SetTargetNfInstanceId sets TargetNfInstanceId field to given value.

### HasTargetNfInstanceId

`func (o *AccessTokenReq) HasTargetNfInstanceId() bool`

HasTargetNfInstanceId returns a boolean if a field has been set.

### GetRequesterPlmn

`func (o *AccessTokenReq) GetRequesterPlmn() PlmnId`

GetRequesterPlmn returns the RequesterPlmn field if non-nil, zero value otherwise.

### GetRequesterPlmnOk

`func (o *AccessTokenReq) GetRequesterPlmnOk() (*PlmnId, bool)`

GetRequesterPlmnOk returns a tuple with the RequesterPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterPlmn

`func (o *AccessTokenReq) SetRequesterPlmn(v PlmnId)`

SetRequesterPlmn sets RequesterPlmn field to given value.

### HasRequesterPlmn

`func (o *AccessTokenReq) HasRequesterPlmn() bool`

HasRequesterPlmn returns a boolean if a field has been set.

### GetRequesterPlmnList

`func (o *AccessTokenReq) GetRequesterPlmnList() []PlmnId`

GetRequesterPlmnList returns the RequesterPlmnList field if non-nil, zero value otherwise.

### GetRequesterPlmnListOk

`func (o *AccessTokenReq) GetRequesterPlmnListOk() (*[]PlmnId, bool)`

GetRequesterPlmnListOk returns a tuple with the RequesterPlmnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterPlmnList

`func (o *AccessTokenReq) SetRequesterPlmnList(v []PlmnId)`

SetRequesterPlmnList sets RequesterPlmnList field to given value.

### HasRequesterPlmnList

`func (o *AccessTokenReq) HasRequesterPlmnList() bool`

HasRequesterPlmnList returns a boolean if a field has been set.

### GetRequesterSnssaiList

`func (o *AccessTokenReq) GetRequesterSnssaiList() []Snssai`

GetRequesterSnssaiList returns the RequesterSnssaiList field if non-nil, zero value otherwise.

### GetRequesterSnssaiListOk

`func (o *AccessTokenReq) GetRequesterSnssaiListOk() (*[]Snssai, bool)`

GetRequesterSnssaiListOk returns a tuple with the RequesterSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterSnssaiList

`func (o *AccessTokenReq) SetRequesterSnssaiList(v []Snssai)`

SetRequesterSnssaiList sets RequesterSnssaiList field to given value.

### HasRequesterSnssaiList

`func (o *AccessTokenReq) HasRequesterSnssaiList() bool`

HasRequesterSnssaiList returns a boolean if a field has been set.

### GetRequesterFqdn

`func (o *AccessTokenReq) GetRequesterFqdn() string`

GetRequesterFqdn returns the RequesterFqdn field if non-nil, zero value otherwise.

### GetRequesterFqdnOk

`func (o *AccessTokenReq) GetRequesterFqdnOk() (*string, bool)`

GetRequesterFqdnOk returns a tuple with the RequesterFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFqdn

`func (o *AccessTokenReq) SetRequesterFqdn(v string)`

SetRequesterFqdn sets RequesterFqdn field to given value.

### HasRequesterFqdn

`func (o *AccessTokenReq) HasRequesterFqdn() bool`

HasRequesterFqdn returns a boolean if a field has been set.

### GetRequesterSnpnList

`func (o *AccessTokenReq) GetRequesterSnpnList() []PlmnIdNid`

GetRequesterSnpnList returns the RequesterSnpnList field if non-nil, zero value otherwise.

### GetRequesterSnpnListOk

`func (o *AccessTokenReq) GetRequesterSnpnListOk() (*[]PlmnIdNid, bool)`

GetRequesterSnpnListOk returns a tuple with the RequesterSnpnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterSnpnList

`func (o *AccessTokenReq) SetRequesterSnpnList(v []PlmnIdNid)`

SetRequesterSnpnList sets RequesterSnpnList field to given value.

### HasRequesterSnpnList

`func (o *AccessTokenReq) HasRequesterSnpnList() bool`

HasRequesterSnpnList returns a boolean if a field has been set.

### GetTargetPlmn

`func (o *AccessTokenReq) GetTargetPlmn() PlmnId`

GetTargetPlmn returns the TargetPlmn field if non-nil, zero value otherwise.

### GetTargetPlmnOk

`func (o *AccessTokenReq) GetTargetPlmnOk() (*PlmnId, bool)`

GetTargetPlmnOk returns a tuple with the TargetPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlmn

`func (o *AccessTokenReq) SetTargetPlmn(v PlmnId)`

SetTargetPlmn sets TargetPlmn field to given value.

### HasTargetPlmn

`func (o *AccessTokenReq) HasTargetPlmn() bool`

HasTargetPlmn returns a boolean if a field has been set.

### GetTargetSnssaiList

`func (o *AccessTokenReq) GetTargetSnssaiList() []Snssai`

GetTargetSnssaiList returns the TargetSnssaiList field if non-nil, zero value otherwise.

### GetTargetSnssaiListOk

`func (o *AccessTokenReq) GetTargetSnssaiListOk() (*[]Snssai, bool)`

GetTargetSnssaiListOk returns a tuple with the TargetSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSnssaiList

`func (o *AccessTokenReq) SetTargetSnssaiList(v []Snssai)`

SetTargetSnssaiList sets TargetSnssaiList field to given value.

### HasTargetSnssaiList

`func (o *AccessTokenReq) HasTargetSnssaiList() bool`

HasTargetSnssaiList returns a boolean if a field has been set.

### GetTargetNsiList

`func (o *AccessTokenReq) GetTargetNsiList() []string`

GetTargetNsiList returns the TargetNsiList field if non-nil, zero value otherwise.

### GetTargetNsiListOk

`func (o *AccessTokenReq) GetTargetNsiListOk() (*[]string, bool)`

GetTargetNsiListOk returns a tuple with the TargetNsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNsiList

`func (o *AccessTokenReq) SetTargetNsiList(v []string)`

SetTargetNsiList sets TargetNsiList field to given value.

### HasTargetNsiList

`func (o *AccessTokenReq) HasTargetNsiList() bool`

HasTargetNsiList returns a boolean if a field has been set.

### GetTargetNfSetId

`func (o *AccessTokenReq) GetTargetNfSetId() string`

GetTargetNfSetId returns the TargetNfSetId field if non-nil, zero value otherwise.

### GetTargetNfSetIdOk

`func (o *AccessTokenReq) GetTargetNfSetIdOk() (*string, bool)`

GetTargetNfSetIdOk returns a tuple with the TargetNfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfSetId

`func (o *AccessTokenReq) SetTargetNfSetId(v string)`

SetTargetNfSetId sets TargetNfSetId field to given value.

### HasTargetNfSetId

`func (o *AccessTokenReq) HasTargetNfSetId() bool`

HasTargetNfSetId returns a boolean if a field has been set.

### GetTargetNfServiceSetId

`func (o *AccessTokenReq) GetTargetNfServiceSetId() string`

GetTargetNfServiceSetId returns the TargetNfServiceSetId field if non-nil, zero value otherwise.

### GetTargetNfServiceSetIdOk

`func (o *AccessTokenReq) GetTargetNfServiceSetIdOk() (*string, bool)`

GetTargetNfServiceSetIdOk returns a tuple with the TargetNfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfServiceSetId

`func (o *AccessTokenReq) SetTargetNfServiceSetId(v string)`

SetTargetNfServiceSetId sets TargetNfServiceSetId field to given value.

### HasTargetNfServiceSetId

`func (o *AccessTokenReq) HasTargetNfServiceSetId() bool`

HasTargetNfServiceSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


