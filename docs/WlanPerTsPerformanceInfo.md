# WlanPerTsPerformanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TsStart** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**TsDuration** | **int32** | indicating a time in seconds. | 
**Rssi** | Pointer to **int32** |  | [optional] 
**Rtt** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TrafficInfo** | Pointer to [**TrafficInformation**](TrafficInformation.md) |  | [optional] 
**NumberOfUes** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewWlanPerTsPerformanceInfo

`func NewWlanPerTsPerformanceInfo(tsStart time.Time, tsDuration int32, ) *WlanPerTsPerformanceInfo`

NewWlanPerTsPerformanceInfo instantiates a new WlanPerTsPerformanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanPerTsPerformanceInfoWithDefaults

`func NewWlanPerTsPerformanceInfoWithDefaults() *WlanPerTsPerformanceInfo`

NewWlanPerTsPerformanceInfoWithDefaults instantiates a new WlanPerTsPerformanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTsStart

`func (o *WlanPerTsPerformanceInfo) GetTsStart() time.Time`

GetTsStart returns the TsStart field if non-nil, zero value otherwise.

### GetTsStartOk

`func (o *WlanPerTsPerformanceInfo) GetTsStartOk() (*time.Time, bool)`

GetTsStartOk returns a tuple with the TsStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsStart

`func (o *WlanPerTsPerformanceInfo) SetTsStart(v time.Time)`

SetTsStart sets TsStart field to given value.


### GetTsDuration

`func (o *WlanPerTsPerformanceInfo) GetTsDuration() int32`

GetTsDuration returns the TsDuration field if non-nil, zero value otherwise.

### GetTsDurationOk

`func (o *WlanPerTsPerformanceInfo) GetTsDurationOk() (*int32, bool)`

GetTsDurationOk returns a tuple with the TsDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsDuration

`func (o *WlanPerTsPerformanceInfo) SetTsDuration(v int32)`

SetTsDuration sets TsDuration field to given value.


### GetRssi

`func (o *WlanPerTsPerformanceInfo) GetRssi() int32`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *WlanPerTsPerformanceInfo) GetRssiOk() (*int32, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *WlanPerTsPerformanceInfo) SetRssi(v int32)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *WlanPerTsPerformanceInfo) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetRtt

`func (o *WlanPerTsPerformanceInfo) GetRtt() int32`

GetRtt returns the Rtt field if non-nil, zero value otherwise.

### GetRttOk

`func (o *WlanPerTsPerformanceInfo) GetRttOk() (*int32, bool)`

GetRttOk returns a tuple with the Rtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtt

`func (o *WlanPerTsPerformanceInfo) SetRtt(v int32)`

SetRtt sets Rtt field to given value.

### HasRtt

`func (o *WlanPerTsPerformanceInfo) HasRtt() bool`

HasRtt returns a boolean if a field has been set.

### GetTrafficInfo

`func (o *WlanPerTsPerformanceInfo) GetTrafficInfo() TrafficInformation`

GetTrafficInfo returns the TrafficInfo field if non-nil, zero value otherwise.

### GetTrafficInfoOk

`func (o *WlanPerTsPerformanceInfo) GetTrafficInfoOk() (*TrafficInformation, bool)`

GetTrafficInfoOk returns a tuple with the TrafficInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficInfo

`func (o *WlanPerTsPerformanceInfo) SetTrafficInfo(v TrafficInformation)`

SetTrafficInfo sets TrafficInfo field to given value.

### HasTrafficInfo

`func (o *WlanPerTsPerformanceInfo) HasTrafficInfo() bool`

HasTrafficInfo returns a boolean if a field has been set.

### GetNumberOfUes

`func (o *WlanPerTsPerformanceInfo) GetNumberOfUes() int32`

GetNumberOfUes returns the NumberOfUes field if non-nil, zero value otherwise.

### GetNumberOfUesOk

`func (o *WlanPerTsPerformanceInfo) GetNumberOfUesOk() (*int32, bool)`

GetNumberOfUesOk returns a tuple with the NumberOfUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfUes

`func (o *WlanPerTsPerformanceInfo) SetNumberOfUes(v int32)`

SetNumberOfUes sets NumberOfUes field to given value.

### HasNumberOfUes

`func (o *WlanPerTsPerformanceInfo) HasNumberOfUes() bool`

HasNumberOfUes returns a boolean if a field has been set.

### GetConfidence

`func (o *WlanPerTsPerformanceInfo) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *WlanPerTsPerformanceInfo) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *WlanPerTsPerformanceInfo) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *WlanPerTsPerformanceInfo) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


