# PerfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgTrafficRate** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**MaxTrafficRate** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**AvePacketDelay** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. | [optional] 
**MaxPacketDelay** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. | [optional] 
**AvgPacketLossRate** | Pointer to **int32** | Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. | [optional] 

## Methods

### NewPerfData

`func NewPerfData() *PerfData`

NewPerfData instantiates a new PerfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfDataWithDefaults

`func NewPerfDataWithDefaults() *PerfData`

NewPerfDataWithDefaults instantiates a new PerfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgTrafficRate

`func (o *PerfData) GetAvgTrafficRate() string`

GetAvgTrafficRate returns the AvgTrafficRate field if non-nil, zero value otherwise.

### GetAvgTrafficRateOk

`func (o *PerfData) GetAvgTrafficRateOk() (*string, bool)`

GetAvgTrafficRateOk returns a tuple with the AvgTrafficRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTrafficRate

`func (o *PerfData) SetAvgTrafficRate(v string)`

SetAvgTrafficRate sets AvgTrafficRate field to given value.

### HasAvgTrafficRate

`func (o *PerfData) HasAvgTrafficRate() bool`

HasAvgTrafficRate returns a boolean if a field has been set.

### GetMaxTrafficRate

`func (o *PerfData) GetMaxTrafficRate() string`

GetMaxTrafficRate returns the MaxTrafficRate field if non-nil, zero value otherwise.

### GetMaxTrafficRateOk

`func (o *PerfData) GetMaxTrafficRateOk() (*string, bool)`

GetMaxTrafficRateOk returns a tuple with the MaxTrafficRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrafficRate

`func (o *PerfData) SetMaxTrafficRate(v string)`

SetMaxTrafficRate sets MaxTrafficRate field to given value.

### HasMaxTrafficRate

`func (o *PerfData) HasMaxTrafficRate() bool`

HasMaxTrafficRate returns a boolean if a field has been set.

### GetAvePacketDelay

`func (o *PerfData) GetAvePacketDelay() int32`

GetAvePacketDelay returns the AvePacketDelay field if non-nil, zero value otherwise.

### GetAvePacketDelayOk

`func (o *PerfData) GetAvePacketDelayOk() (*int32, bool)`

GetAvePacketDelayOk returns a tuple with the AvePacketDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvePacketDelay

`func (o *PerfData) SetAvePacketDelay(v int32)`

SetAvePacketDelay sets AvePacketDelay field to given value.

### HasAvePacketDelay

`func (o *PerfData) HasAvePacketDelay() bool`

HasAvePacketDelay returns a boolean if a field has been set.

### GetMaxPacketDelay

`func (o *PerfData) GetMaxPacketDelay() int32`

GetMaxPacketDelay returns the MaxPacketDelay field if non-nil, zero value otherwise.

### GetMaxPacketDelayOk

`func (o *PerfData) GetMaxPacketDelayOk() (*int32, bool)`

GetMaxPacketDelayOk returns a tuple with the MaxPacketDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketDelay

`func (o *PerfData) SetMaxPacketDelay(v int32)`

SetMaxPacketDelay sets MaxPacketDelay field to given value.

### HasMaxPacketDelay

`func (o *PerfData) HasMaxPacketDelay() bool`

HasMaxPacketDelay returns a boolean if a field has been set.

### GetAvgPacketLossRate

`func (o *PerfData) GetAvgPacketLossRate() int32`

GetAvgPacketLossRate returns the AvgPacketLossRate field if non-nil, zero value otherwise.

### GetAvgPacketLossRateOk

`func (o *PerfData) GetAvgPacketLossRateOk() (*int32, bool)`

GetAvgPacketLossRateOk returns a tuple with the AvgPacketLossRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPacketLossRate

`func (o *PerfData) SetAvgPacketLossRate(v int32)`

SetAvgPacketLossRate sets AvgPacketLossRate field to given value.

### HasAvgPacketLossRate

`func (o *PerfData) HasAvgPacketLossRate() bool`

HasAvgPacketLossRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


