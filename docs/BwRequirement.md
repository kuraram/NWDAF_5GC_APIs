# BwRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | String providing an application identifier. | 
**MarBwDl** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**MarBwUl** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**MirBwDl** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**MirBwUl** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 

## Methods

### NewBwRequirement

`func NewBwRequirement(appId string, ) *BwRequirement`

NewBwRequirement instantiates a new BwRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBwRequirementWithDefaults

`func NewBwRequirementWithDefaults() *BwRequirement`

NewBwRequirementWithDefaults instantiates a new BwRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *BwRequirement) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *BwRequirement) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *BwRequirement) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetMarBwDl

`func (o *BwRequirement) GetMarBwDl() string`

GetMarBwDl returns the MarBwDl field if non-nil, zero value otherwise.

### GetMarBwDlOk

`func (o *BwRequirement) GetMarBwDlOk() (*string, bool)`

GetMarBwDlOk returns a tuple with the MarBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwDl

`func (o *BwRequirement) SetMarBwDl(v string)`

SetMarBwDl sets MarBwDl field to given value.

### HasMarBwDl

`func (o *BwRequirement) HasMarBwDl() bool`

HasMarBwDl returns a boolean if a field has been set.

### GetMarBwUl

`func (o *BwRequirement) GetMarBwUl() string`

GetMarBwUl returns the MarBwUl field if non-nil, zero value otherwise.

### GetMarBwUlOk

`func (o *BwRequirement) GetMarBwUlOk() (*string, bool)`

GetMarBwUlOk returns a tuple with the MarBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwUl

`func (o *BwRequirement) SetMarBwUl(v string)`

SetMarBwUl sets MarBwUl field to given value.

### HasMarBwUl

`func (o *BwRequirement) HasMarBwUl() bool`

HasMarBwUl returns a boolean if a field has been set.

### GetMirBwDl

`func (o *BwRequirement) GetMirBwDl() string`

GetMirBwDl returns the MirBwDl field if non-nil, zero value otherwise.

### GetMirBwDlOk

`func (o *BwRequirement) GetMirBwDlOk() (*string, bool)`

GetMirBwDlOk returns a tuple with the MirBwDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirBwDl

`func (o *BwRequirement) SetMirBwDl(v string)`

SetMirBwDl sets MirBwDl field to given value.

### HasMirBwDl

`func (o *BwRequirement) HasMirBwDl() bool`

HasMirBwDl returns a boolean if a field has been set.

### GetMirBwUl

`func (o *BwRequirement) GetMirBwUl() string`

GetMirBwUl returns the MirBwUl field if non-nil, zero value otherwise.

### GetMirBwUlOk

`func (o *BwRequirement) GetMirBwUlOk() (*string, bool)`

GetMirBwUlOk returns a tuple with the MirBwUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirBwUl

`func (o *BwRequirement) SetMirBwUl(v string)`

SetMirBwUl sets MirBwUl field to given value.

### HasMirBwUl

`func (o *BwRequirement) HasMirBwUl() bool`

HasMirBwUl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


