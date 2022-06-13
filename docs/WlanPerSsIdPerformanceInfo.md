# WlanPerSsIdPerformanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsId** | **string** |  | 
**WlanPerTsInfos** | [**[]WlanPerTsPerformanceInfo**](WlanPerTsPerformanceInfo.md) |  | 

## Methods

### NewWlanPerSsIdPerformanceInfo

`func NewWlanPerSsIdPerformanceInfo(ssId string, wlanPerTsInfos []WlanPerTsPerformanceInfo, ) *WlanPerSsIdPerformanceInfo`

NewWlanPerSsIdPerformanceInfo instantiates a new WlanPerSsIdPerformanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanPerSsIdPerformanceInfoWithDefaults

`func NewWlanPerSsIdPerformanceInfoWithDefaults() *WlanPerSsIdPerformanceInfo`

NewWlanPerSsIdPerformanceInfoWithDefaults instantiates a new WlanPerSsIdPerformanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsId

`func (o *WlanPerSsIdPerformanceInfo) GetSsId() string`

GetSsId returns the SsId field if non-nil, zero value otherwise.

### GetSsIdOk

`func (o *WlanPerSsIdPerformanceInfo) GetSsIdOk() (*string, bool)`

GetSsIdOk returns a tuple with the SsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsId

`func (o *WlanPerSsIdPerformanceInfo) SetSsId(v string)`

SetSsId sets SsId field to given value.


### GetWlanPerTsInfos

`func (o *WlanPerSsIdPerformanceInfo) GetWlanPerTsInfos() []WlanPerTsPerformanceInfo`

GetWlanPerTsInfos returns the WlanPerTsInfos field if non-nil, zero value otherwise.

### GetWlanPerTsInfosOk

`func (o *WlanPerSsIdPerformanceInfo) GetWlanPerTsInfosOk() (*[]WlanPerTsPerformanceInfo, bool)`

GetWlanPerTsInfosOk returns a tuple with the WlanPerTsInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanPerTsInfos

`func (o *WlanPerSsIdPerformanceInfo) SetWlanPerTsInfos(v []WlanPerTsPerformanceInfo)`

SetWlanPerTsInfos sets WlanPerTsInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


