# ExpectedUeBehaviourData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StationaryIndication** | Pointer to [**StationaryIndication**](StationaryIndication.md) |  | [optional] 
**CommunicationDurationTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**PeriodicTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**ScheduledCommunicationTime** | Pointer to [**ScheduledCommunicationTime1**](ScheduledCommunicationTime1.md) |  | [optional] 
**ScheduledCommunicationType** | Pointer to [**ScheduledCommunicationType**](ScheduledCommunicationType.md) |  | [optional] 
**ExpectedUmts** | Pointer to [**[]LocationArea**](LocationArea.md) | Identifies the UE&#39;s expected geographical movement. The attribute is only applicable in 5G. | [optional] 
**TrafficProfile** | Pointer to [**TrafficProfile**](TrafficProfile.md) |  | [optional] 
**BatteryIndication** | Pointer to [**BatteryIndication**](BatteryIndication.md) |  | [optional] 
**ValidityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewExpectedUeBehaviourData

`func NewExpectedUeBehaviourData() *ExpectedUeBehaviourData`

NewExpectedUeBehaviourData instantiates a new ExpectedUeBehaviourData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpectedUeBehaviourDataWithDefaults

`func NewExpectedUeBehaviourDataWithDefaults() *ExpectedUeBehaviourData`

NewExpectedUeBehaviourDataWithDefaults instantiates a new ExpectedUeBehaviourData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStationaryIndication

`func (o *ExpectedUeBehaviourData) GetStationaryIndication() StationaryIndication`

GetStationaryIndication returns the StationaryIndication field if non-nil, zero value otherwise.

### GetStationaryIndicationOk

`func (o *ExpectedUeBehaviourData) GetStationaryIndicationOk() (*StationaryIndication, bool)`

GetStationaryIndicationOk returns a tuple with the StationaryIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationaryIndication

`func (o *ExpectedUeBehaviourData) SetStationaryIndication(v StationaryIndication)`

SetStationaryIndication sets StationaryIndication field to given value.

### HasStationaryIndication

`func (o *ExpectedUeBehaviourData) HasStationaryIndication() bool`

HasStationaryIndication returns a boolean if a field has been set.

### GetCommunicationDurationTime

`func (o *ExpectedUeBehaviourData) GetCommunicationDurationTime() int32`

GetCommunicationDurationTime returns the CommunicationDurationTime field if non-nil, zero value otherwise.

### GetCommunicationDurationTimeOk

`func (o *ExpectedUeBehaviourData) GetCommunicationDurationTimeOk() (*int32, bool)`

GetCommunicationDurationTimeOk returns a tuple with the CommunicationDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDurationTime

`func (o *ExpectedUeBehaviourData) SetCommunicationDurationTime(v int32)`

SetCommunicationDurationTime sets CommunicationDurationTime field to given value.

### HasCommunicationDurationTime

`func (o *ExpectedUeBehaviourData) HasCommunicationDurationTime() bool`

HasCommunicationDurationTime returns a boolean if a field has been set.

### GetPeriodicTime

`func (o *ExpectedUeBehaviourData) GetPeriodicTime() int32`

GetPeriodicTime returns the PeriodicTime field if non-nil, zero value otherwise.

### GetPeriodicTimeOk

`func (o *ExpectedUeBehaviourData) GetPeriodicTimeOk() (*int32, bool)`

GetPeriodicTimeOk returns a tuple with the PeriodicTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicTime

`func (o *ExpectedUeBehaviourData) SetPeriodicTime(v int32)`

SetPeriodicTime sets PeriodicTime field to given value.

### HasPeriodicTime

`func (o *ExpectedUeBehaviourData) HasPeriodicTime() bool`

HasPeriodicTime returns a boolean if a field has been set.

### GetScheduledCommunicationTime

`func (o *ExpectedUeBehaviourData) GetScheduledCommunicationTime() ScheduledCommunicationTime1`

GetScheduledCommunicationTime returns the ScheduledCommunicationTime field if non-nil, zero value otherwise.

### GetScheduledCommunicationTimeOk

`func (o *ExpectedUeBehaviourData) GetScheduledCommunicationTimeOk() (*ScheduledCommunicationTime1, bool)`

GetScheduledCommunicationTimeOk returns a tuple with the ScheduledCommunicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCommunicationTime

`func (o *ExpectedUeBehaviourData) SetScheduledCommunicationTime(v ScheduledCommunicationTime1)`

SetScheduledCommunicationTime sets ScheduledCommunicationTime field to given value.

### HasScheduledCommunicationTime

`func (o *ExpectedUeBehaviourData) HasScheduledCommunicationTime() bool`

HasScheduledCommunicationTime returns a boolean if a field has been set.

### GetScheduledCommunicationType

`func (o *ExpectedUeBehaviourData) GetScheduledCommunicationType() ScheduledCommunicationType`

GetScheduledCommunicationType returns the ScheduledCommunicationType field if non-nil, zero value otherwise.

### GetScheduledCommunicationTypeOk

`func (o *ExpectedUeBehaviourData) GetScheduledCommunicationTypeOk() (*ScheduledCommunicationType, bool)`

GetScheduledCommunicationTypeOk returns a tuple with the ScheduledCommunicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledCommunicationType

`func (o *ExpectedUeBehaviourData) SetScheduledCommunicationType(v ScheduledCommunicationType)`

SetScheduledCommunicationType sets ScheduledCommunicationType field to given value.

### HasScheduledCommunicationType

`func (o *ExpectedUeBehaviourData) HasScheduledCommunicationType() bool`

HasScheduledCommunicationType returns a boolean if a field has been set.

### GetExpectedUmts

`func (o *ExpectedUeBehaviourData) GetExpectedUmts() []LocationArea`

GetExpectedUmts returns the ExpectedUmts field if non-nil, zero value otherwise.

### GetExpectedUmtsOk

`func (o *ExpectedUeBehaviourData) GetExpectedUmtsOk() (*[]LocationArea, bool)`

GetExpectedUmtsOk returns a tuple with the ExpectedUmts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUmts

`func (o *ExpectedUeBehaviourData) SetExpectedUmts(v []LocationArea)`

SetExpectedUmts sets ExpectedUmts field to given value.

### HasExpectedUmts

`func (o *ExpectedUeBehaviourData) HasExpectedUmts() bool`

HasExpectedUmts returns a boolean if a field has been set.

### GetTrafficProfile

`func (o *ExpectedUeBehaviourData) GetTrafficProfile() TrafficProfile`

GetTrafficProfile returns the TrafficProfile field if non-nil, zero value otherwise.

### GetTrafficProfileOk

`func (o *ExpectedUeBehaviourData) GetTrafficProfileOk() (*TrafficProfile, bool)`

GetTrafficProfileOk returns a tuple with the TrafficProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficProfile

`func (o *ExpectedUeBehaviourData) SetTrafficProfile(v TrafficProfile)`

SetTrafficProfile sets TrafficProfile field to given value.

### HasTrafficProfile

`func (o *ExpectedUeBehaviourData) HasTrafficProfile() bool`

HasTrafficProfile returns a boolean if a field has been set.

### GetBatteryIndication

`func (o *ExpectedUeBehaviourData) GetBatteryIndication() BatteryIndication`

GetBatteryIndication returns the BatteryIndication field if non-nil, zero value otherwise.

### GetBatteryIndicationOk

`func (o *ExpectedUeBehaviourData) GetBatteryIndicationOk() (*BatteryIndication, bool)`

GetBatteryIndicationOk returns a tuple with the BatteryIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatteryIndication

`func (o *ExpectedUeBehaviourData) SetBatteryIndication(v BatteryIndication)`

SetBatteryIndication sets BatteryIndication field to given value.

### HasBatteryIndication

`func (o *ExpectedUeBehaviourData) HasBatteryIndication() bool`

HasBatteryIndication returns a boolean if a field has been set.

### GetValidityTime

`func (o *ExpectedUeBehaviourData) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *ExpectedUeBehaviourData) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *ExpectedUeBehaviourData) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *ExpectedUeBehaviourData) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


