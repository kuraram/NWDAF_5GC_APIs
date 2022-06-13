# UeMobility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RecurringTime** | Pointer to [**ScheduledCommunicationTime**](ScheduledCommunicationTime.md) |  | [optional] 
**Duration** | **int32** | indicating a time in seconds. | 
**DurationVariance** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 
**LocInfos** | [**[]LocationInfo**](LocationInfo.md) |  | 

## Methods

### NewUeMobility

`func NewUeMobility(duration int32, locInfos []LocationInfo, ) *UeMobility`

NewUeMobility instantiates a new UeMobility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeMobilityWithDefaults

`func NewUeMobilityWithDefaults() *UeMobility`

NewUeMobilityWithDefaults instantiates a new UeMobility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *UeMobility) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *UeMobility) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *UeMobility) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *UeMobility) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetRecurringTime

`func (o *UeMobility) GetRecurringTime() ScheduledCommunicationTime`

GetRecurringTime returns the RecurringTime field if non-nil, zero value otherwise.

### GetRecurringTimeOk

`func (o *UeMobility) GetRecurringTimeOk() (*ScheduledCommunicationTime, bool)`

GetRecurringTimeOk returns a tuple with the RecurringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTime

`func (o *UeMobility) SetRecurringTime(v ScheduledCommunicationTime)`

SetRecurringTime sets RecurringTime field to given value.

### HasRecurringTime

`func (o *UeMobility) HasRecurringTime() bool`

HasRecurringTime returns a boolean if a field has been set.

### GetDuration

`func (o *UeMobility) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UeMobility) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UeMobility) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetDurationVariance

`func (o *UeMobility) GetDurationVariance() float32`

GetDurationVariance returns the DurationVariance field if non-nil, zero value otherwise.

### GetDurationVarianceOk

`func (o *UeMobility) GetDurationVarianceOk() (*float32, bool)`

GetDurationVarianceOk returns a tuple with the DurationVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationVariance

`func (o *UeMobility) SetDurationVariance(v float32)`

SetDurationVariance sets DurationVariance field to given value.

### HasDurationVariance

`func (o *UeMobility) HasDurationVariance() bool`

HasDurationVariance returns a boolean if a field has been set.

### GetLocInfos

`func (o *UeMobility) GetLocInfos() []LocationInfo`

GetLocInfos returns the LocInfos field if non-nil, zero value otherwise.

### GetLocInfosOk

`func (o *UeMobility) GetLocInfosOk() (*[]LocationInfo, bool)`

GetLocInfosOk returns a tuple with the LocInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInfos

`func (o *UeMobility) SetLocInfos(v []LocationInfo)`

SetLocInfos sets LocInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


