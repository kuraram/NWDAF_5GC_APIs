# UserDataCongestionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkArea** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**CongestionInfo** | Pointer to [**CongestionInfo**](CongestionInfo.md) |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 

## Methods

### NewUserDataCongestionInfo

`func NewUserDataCongestionInfo() *UserDataCongestionInfo`

NewUserDataCongestionInfo instantiates a new UserDataCongestionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataCongestionInfoWithDefaults

`func NewUserDataCongestionInfoWithDefaults() *UserDataCongestionInfo`

NewUserDataCongestionInfoWithDefaults instantiates a new UserDataCongestionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkArea

`func (o *UserDataCongestionInfo) GetNetworkArea() NetworkAreaInfo`

GetNetworkArea returns the NetworkArea field if non-nil, zero value otherwise.

### GetNetworkAreaOk

`func (o *UserDataCongestionInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool)`

GetNetworkAreaOk returns a tuple with the NetworkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkArea

`func (o *UserDataCongestionInfo) SetNetworkArea(v NetworkAreaInfo)`

SetNetworkArea sets NetworkArea field to given value.

### HasNetworkArea

`func (o *UserDataCongestionInfo) HasNetworkArea() bool`

HasNetworkArea returns a boolean if a field has been set.

### GetCongestionInfo

`func (o *UserDataCongestionInfo) GetCongestionInfo() CongestionInfo`

GetCongestionInfo returns the CongestionInfo field if non-nil, zero value otherwise.

### GetCongestionInfoOk

`func (o *UserDataCongestionInfo) GetCongestionInfoOk() (*CongestionInfo, bool)`

GetCongestionInfoOk returns a tuple with the CongestionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongestionInfo

`func (o *UserDataCongestionInfo) SetCongestionInfo(v CongestionInfo)`

SetCongestionInfo sets CongestionInfo field to given value.

### HasCongestionInfo

`func (o *UserDataCongestionInfo) HasCongestionInfo() bool`

HasCongestionInfo returns a boolean if a field has been set.

### GetSnssai

`func (o *UserDataCongestionInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *UserDataCongestionInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *UserDataCongestionInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *UserDataCongestionInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


