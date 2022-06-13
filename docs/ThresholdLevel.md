# ThresholdLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CongLevel** | Pointer to **int32** |  | [optional] 
**NfLoadLevel** | Pointer to **int32** |  | [optional] 
**NfCpuUsage** | Pointer to **int32** |  | [optional] 
**NfMemoryUsage** | Pointer to **int32** |  | [optional] 
**NfStorageUsage** | Pointer to **int32** |  | [optional] 
**AvgTrafficRate** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**MaxTrafficRate** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**AvgPacketDelay** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. | [optional] 
**MaxPacketDelay** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. | [optional] 
**AvgPacketLossRate** | Pointer to **int32** | Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. | [optional] 

## Methods

### NewThresholdLevel

`func NewThresholdLevel() *ThresholdLevel`

NewThresholdLevel instantiates a new ThresholdLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdLevelWithDefaults

`func NewThresholdLevelWithDefaults() *ThresholdLevel`

NewThresholdLevelWithDefaults instantiates a new ThresholdLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCongLevel

`func (o *ThresholdLevel) GetCongLevel() int32`

GetCongLevel returns the CongLevel field if non-nil, zero value otherwise.

### GetCongLevelOk

`func (o *ThresholdLevel) GetCongLevelOk() (*int32, bool)`

GetCongLevelOk returns a tuple with the CongLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongLevel

`func (o *ThresholdLevel) SetCongLevel(v int32)`

SetCongLevel sets CongLevel field to given value.

### HasCongLevel

`func (o *ThresholdLevel) HasCongLevel() bool`

HasCongLevel returns a boolean if a field has been set.

### GetNfLoadLevel

`func (o *ThresholdLevel) GetNfLoadLevel() int32`

GetNfLoadLevel returns the NfLoadLevel field if non-nil, zero value otherwise.

### GetNfLoadLevelOk

`func (o *ThresholdLevel) GetNfLoadLevelOk() (*int32, bool)`

GetNfLoadLevelOk returns a tuple with the NfLoadLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfLoadLevel

`func (o *ThresholdLevel) SetNfLoadLevel(v int32)`

SetNfLoadLevel sets NfLoadLevel field to given value.

### HasNfLoadLevel

`func (o *ThresholdLevel) HasNfLoadLevel() bool`

HasNfLoadLevel returns a boolean if a field has been set.

### GetNfCpuUsage

`func (o *ThresholdLevel) GetNfCpuUsage() int32`

GetNfCpuUsage returns the NfCpuUsage field if non-nil, zero value otherwise.

### GetNfCpuUsageOk

`func (o *ThresholdLevel) GetNfCpuUsageOk() (*int32, bool)`

GetNfCpuUsageOk returns a tuple with the NfCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfCpuUsage

`func (o *ThresholdLevel) SetNfCpuUsage(v int32)`

SetNfCpuUsage sets NfCpuUsage field to given value.

### HasNfCpuUsage

`func (o *ThresholdLevel) HasNfCpuUsage() bool`

HasNfCpuUsage returns a boolean if a field has been set.

### GetNfMemoryUsage

`func (o *ThresholdLevel) GetNfMemoryUsage() int32`

GetNfMemoryUsage returns the NfMemoryUsage field if non-nil, zero value otherwise.

### GetNfMemoryUsageOk

`func (o *ThresholdLevel) GetNfMemoryUsageOk() (*int32, bool)`

GetNfMemoryUsageOk returns a tuple with the NfMemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfMemoryUsage

`func (o *ThresholdLevel) SetNfMemoryUsage(v int32)`

SetNfMemoryUsage sets NfMemoryUsage field to given value.

### HasNfMemoryUsage

`func (o *ThresholdLevel) HasNfMemoryUsage() bool`

HasNfMemoryUsage returns a boolean if a field has been set.

### GetNfStorageUsage

`func (o *ThresholdLevel) GetNfStorageUsage() int32`

GetNfStorageUsage returns the NfStorageUsage field if non-nil, zero value otherwise.

### GetNfStorageUsageOk

`func (o *ThresholdLevel) GetNfStorageUsageOk() (*int32, bool)`

GetNfStorageUsageOk returns a tuple with the NfStorageUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfStorageUsage

`func (o *ThresholdLevel) SetNfStorageUsage(v int32)`

SetNfStorageUsage sets NfStorageUsage field to given value.

### HasNfStorageUsage

`func (o *ThresholdLevel) HasNfStorageUsage() bool`

HasNfStorageUsage returns a boolean if a field has been set.

### GetAvgTrafficRate

`func (o *ThresholdLevel) GetAvgTrafficRate() string`

GetAvgTrafficRate returns the AvgTrafficRate field if non-nil, zero value otherwise.

### GetAvgTrafficRateOk

`func (o *ThresholdLevel) GetAvgTrafficRateOk() (*string, bool)`

GetAvgTrafficRateOk returns a tuple with the AvgTrafficRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTrafficRate

`func (o *ThresholdLevel) SetAvgTrafficRate(v string)`

SetAvgTrafficRate sets AvgTrafficRate field to given value.

### HasAvgTrafficRate

`func (o *ThresholdLevel) HasAvgTrafficRate() bool`

HasAvgTrafficRate returns a boolean if a field has been set.

### GetMaxTrafficRate

`func (o *ThresholdLevel) GetMaxTrafficRate() string`

GetMaxTrafficRate returns the MaxTrafficRate field if non-nil, zero value otherwise.

### GetMaxTrafficRateOk

`func (o *ThresholdLevel) GetMaxTrafficRateOk() (*string, bool)`

GetMaxTrafficRateOk returns a tuple with the MaxTrafficRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrafficRate

`func (o *ThresholdLevel) SetMaxTrafficRate(v string)`

SetMaxTrafficRate sets MaxTrafficRate field to given value.

### HasMaxTrafficRate

`func (o *ThresholdLevel) HasMaxTrafficRate() bool`

HasMaxTrafficRate returns a boolean if a field has been set.

### GetAvgPacketDelay

`func (o *ThresholdLevel) GetAvgPacketDelay() int32`

GetAvgPacketDelay returns the AvgPacketDelay field if non-nil, zero value otherwise.

### GetAvgPacketDelayOk

`func (o *ThresholdLevel) GetAvgPacketDelayOk() (*int32, bool)`

GetAvgPacketDelayOk returns a tuple with the AvgPacketDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPacketDelay

`func (o *ThresholdLevel) SetAvgPacketDelay(v int32)`

SetAvgPacketDelay sets AvgPacketDelay field to given value.

### HasAvgPacketDelay

`func (o *ThresholdLevel) HasAvgPacketDelay() bool`

HasAvgPacketDelay returns a boolean if a field has been set.

### GetMaxPacketDelay

`func (o *ThresholdLevel) GetMaxPacketDelay() int32`

GetMaxPacketDelay returns the MaxPacketDelay field if non-nil, zero value otherwise.

### GetMaxPacketDelayOk

`func (o *ThresholdLevel) GetMaxPacketDelayOk() (*int32, bool)`

GetMaxPacketDelayOk returns a tuple with the MaxPacketDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketDelay

`func (o *ThresholdLevel) SetMaxPacketDelay(v int32)`

SetMaxPacketDelay sets MaxPacketDelay field to given value.

### HasMaxPacketDelay

`func (o *ThresholdLevel) HasMaxPacketDelay() bool`

HasMaxPacketDelay returns a boolean if a field has been set.

### GetAvgPacketLossRate

`func (o *ThresholdLevel) GetAvgPacketLossRate() int32`

GetAvgPacketLossRate returns the AvgPacketLossRate field if non-nil, zero value otherwise.

### GetAvgPacketLossRateOk

`func (o *ThresholdLevel) GetAvgPacketLossRateOk() (*int32, bool)`

GetAvgPacketLossRateOk returns a tuple with the AvgPacketLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPacketLossRate

`func (o *ThresholdLevel) SetAvgPacketLossRate(v int32)`

SetAvgPacketLossRate sets AvgPacketLossRate field to given value.

### HasAvgPacketLossRate

`func (o *ThresholdLevel) HasAvgPacketLossRate() bool`

HasAvgPacketLossRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


