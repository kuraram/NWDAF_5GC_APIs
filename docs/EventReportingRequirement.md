# EventReportingRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accuracy** | Pointer to [**Accuracy**](Accuracy.md) |  | [optional] 
**StartTs** | Pointer to **time.Time** |  | [optional] 
**EndTs** | Pointer to **time.Time** |  | [optional] 
**SampRatio** | Pointer to **int32** |  | [optional] 
**MaxObjectNbr** | Pointer to **int32** |  | [optional] 
**MaxSupiNbr** | Pointer to **int32** |  | [optional] 
**TimeAnaNeeded** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEventReportingRequirement

`func NewEventReportingRequirement() *EventReportingRequirement`

NewEventReportingRequirement instantiates a new EventReportingRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventReportingRequirementWithDefaults

`func NewEventReportingRequirementWithDefaults() *EventReportingRequirement`

NewEventReportingRequirementWithDefaults instantiates a new EventReportingRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracy

`func (o *EventReportingRequirement) GetAccuracy() Accuracy`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *EventReportingRequirement) GetAccuracyOk() (*Accuracy, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *EventReportingRequirement) SetAccuracy(v Accuracy)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *EventReportingRequirement) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetStartTs

`func (o *EventReportingRequirement) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *EventReportingRequirement) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *EventReportingRequirement) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *EventReportingRequirement) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *EventReportingRequirement) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *EventReportingRequirement) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *EventReportingRequirement) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *EventReportingRequirement) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetSampRatio

`func (o *EventReportingRequirement) GetSampRatio() int32`

GetSampRatio returns the SampRatio field if non-nil, zero value otherwise.

### GetSampRatioOk

`func (o *EventReportingRequirement) GetSampRatioOk() (*int32, bool)`

GetSampRatioOk returns a tuple with the SampRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampRatio

`func (o *EventReportingRequirement) SetSampRatio(v int32)`

SetSampRatio sets SampRatio field to given value.

### HasSampRatio

`func (o *EventReportingRequirement) HasSampRatio() bool`

HasSampRatio returns a boolean if a field has been set.

### GetMaxObjectNbr

`func (o *EventReportingRequirement) GetMaxObjectNbr() int32`

GetMaxObjectNbr returns the MaxObjectNbr field if non-nil, zero value otherwise.

### GetMaxObjectNbrOk

`func (o *EventReportingRequirement) GetMaxObjectNbrOk() (*int32, bool)`

GetMaxObjectNbrOk returns a tuple with the MaxObjectNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxObjectNbr

`func (o *EventReportingRequirement) SetMaxObjectNbr(v int32)`

SetMaxObjectNbr sets MaxObjectNbr field to given value.

### HasMaxObjectNbr

`func (o *EventReportingRequirement) HasMaxObjectNbr() bool`

HasMaxObjectNbr returns a boolean if a field has been set.

### GetMaxSupiNbr

`func (o *EventReportingRequirement) GetMaxSupiNbr() int32`

GetMaxSupiNbr returns the MaxSupiNbr field if non-nil, zero value otherwise.

### GetMaxSupiNbrOk

`func (o *EventReportingRequirement) GetMaxSupiNbrOk() (*int32, bool)`

GetMaxSupiNbrOk returns a tuple with the MaxSupiNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSupiNbr

`func (o *EventReportingRequirement) SetMaxSupiNbr(v int32)`

SetMaxSupiNbr sets MaxSupiNbr field to given value.

### HasMaxSupiNbr

`func (o *EventReportingRequirement) HasMaxSupiNbr() bool`

HasMaxSupiNbr returns a boolean if a field has been set.

### GetTimeAnaNeeded

`func (o *EventReportingRequirement) GetTimeAnaNeeded() time.Time`

GetTimeAnaNeeded returns the TimeAnaNeeded field if non-nil, zero value otherwise.

### GetTimeAnaNeededOk

`func (o *EventReportingRequirement) GetTimeAnaNeededOk() (*time.Time, bool)`

GetTimeAnaNeededOk returns a tuple with the TimeAnaNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAnaNeeded

`func (o *EventReportingRequirement) SetTimeAnaNeeded(v time.Time)`

SetTimeAnaNeeded sets TimeAnaNeeded field to given value.

### HasTimeAnaNeeded

`func (o *EventReportingRequirement) HasTimeAnaNeeded() bool`

HasTimeAnaNeeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


