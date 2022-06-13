# WlanPerformanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkArea** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**WlanPerSsidInfos** | [**[]WlanPerSsIdPerformanceInfo**](WlanPerSsIdPerformanceInfo.md) |  | 

## Methods

### NewWlanPerformanceInfo

`func NewWlanPerformanceInfo(wlanPerSsidInfos []WlanPerSsIdPerformanceInfo, ) *WlanPerformanceInfo`

NewWlanPerformanceInfo instantiates a new WlanPerformanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanPerformanceInfoWithDefaults

`func NewWlanPerformanceInfoWithDefaults() *WlanPerformanceInfo`

NewWlanPerformanceInfoWithDefaults instantiates a new WlanPerformanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkArea

`func (o *WlanPerformanceInfo) GetNetworkArea() NetworkAreaInfo`

GetNetworkArea returns the NetworkArea field if non-nil, zero value otherwise.

### GetNetworkAreaOk

`func (o *WlanPerformanceInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool)`

GetNetworkAreaOk returns a tuple with the NetworkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkArea

`func (o *WlanPerformanceInfo) SetNetworkArea(v NetworkAreaInfo)`

SetNetworkArea sets NetworkArea field to given value.

### HasNetworkArea

`func (o *WlanPerformanceInfo) HasNetworkArea() bool`

HasNetworkArea returns a boolean if a field has been set.

### GetWlanPerSsidInfos

`func (o *WlanPerformanceInfo) GetWlanPerSsidInfos() []WlanPerSsIdPerformanceInfo`

GetWlanPerSsidInfos returns the WlanPerSsidInfos field if non-nil, zero value otherwise.

### GetWlanPerSsidInfosOk

`func (o *WlanPerformanceInfo) GetWlanPerSsidInfosOk() (*[]WlanPerSsIdPerformanceInfo, bool)`

GetWlanPerSsidInfosOk returns a tuple with the WlanPerSsidInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanPerSsidInfos

`func (o *WlanPerformanceInfo) SetWlanPerSsidInfos(v []WlanPerSsIdPerformanceInfo)`

SetWlanPerSsidInfos sets WlanPerSsidInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


