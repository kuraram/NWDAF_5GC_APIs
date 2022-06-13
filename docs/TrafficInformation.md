# TrafficInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UplinkRate** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**DownlinkRate** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**UplinkVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**DownlinkVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**TotalVolume** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 

## Methods

### NewTrafficInformation

`func NewTrafficInformation() *TrafficInformation`

NewTrafficInformation instantiates a new TrafficInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficInformationWithDefaults

`func NewTrafficInformationWithDefaults() *TrafficInformation`

NewTrafficInformationWithDefaults instantiates a new TrafficInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplinkRate

`func (o *TrafficInformation) GetUplinkRate() string`

GetUplinkRate returns the UplinkRate field if non-nil, zero value otherwise.

### GetUplinkRateOk

`func (o *TrafficInformation) GetUplinkRateOk() (*string, bool)`

GetUplinkRateOk returns a tuple with the UplinkRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkRate

`func (o *TrafficInformation) SetUplinkRate(v string)`

SetUplinkRate sets UplinkRate field to given value.

### HasUplinkRate

`func (o *TrafficInformation) HasUplinkRate() bool`

HasUplinkRate returns a boolean if a field has been set.

### GetDownlinkRate

`func (o *TrafficInformation) GetDownlinkRate() string`

GetDownlinkRate returns the DownlinkRate field if non-nil, zero value otherwise.

### GetDownlinkRateOk

`func (o *TrafficInformation) GetDownlinkRateOk() (*string, bool)`

GetDownlinkRateOk returns a tuple with the DownlinkRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkRate

`func (o *TrafficInformation) SetDownlinkRate(v string)`

SetDownlinkRate sets DownlinkRate field to given value.

### HasDownlinkRate

`func (o *TrafficInformation) HasDownlinkRate() bool`

HasDownlinkRate returns a boolean if a field has been set.

### GetUplinkVolume

`func (o *TrafficInformation) GetUplinkVolume() int64`

GetUplinkVolume returns the UplinkVolume field if non-nil, zero value otherwise.

### GetUplinkVolumeOk

`func (o *TrafficInformation) GetUplinkVolumeOk() (*int64, bool)`

GetUplinkVolumeOk returns a tuple with the UplinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkVolume

`func (o *TrafficInformation) SetUplinkVolume(v int64)`

SetUplinkVolume sets UplinkVolume field to given value.

### HasUplinkVolume

`func (o *TrafficInformation) HasUplinkVolume() bool`

HasUplinkVolume returns a boolean if a field has been set.

### GetDownlinkVolume

`func (o *TrafficInformation) GetDownlinkVolume() int64`

GetDownlinkVolume returns the DownlinkVolume field if non-nil, zero value otherwise.

### GetDownlinkVolumeOk

`func (o *TrafficInformation) GetDownlinkVolumeOk() (*int64, bool)`

GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkVolume

`func (o *TrafficInformation) SetDownlinkVolume(v int64)`

SetDownlinkVolume sets DownlinkVolume field to given value.

### HasDownlinkVolume

`func (o *TrafficInformation) HasDownlinkVolume() bool`

HasDownlinkVolume returns a boolean if a field has been set.

### GetTotalVolume

`func (o *TrafficInformation) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *TrafficInformation) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *TrafficInformation) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *TrafficInformation) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


