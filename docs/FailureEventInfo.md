# FailureEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**NwdafEvent**](NwdafEvent.md) |  | 
**FailureCode** | [**NwdafFailureCode**](NwdafFailureCode.md) |  | 

## Methods

### NewFailureEventInfo

`func NewFailureEventInfo(event NwdafEvent, failureCode NwdafFailureCode, ) *FailureEventInfo`

NewFailureEventInfo instantiates a new FailureEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureEventInfoWithDefaults

`func NewFailureEventInfoWithDefaults() *FailureEventInfo`

NewFailureEventInfoWithDefaults instantiates a new FailureEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *FailureEventInfo) GetEvent() NwdafEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *FailureEventInfo) GetEventOk() (*NwdafEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *FailureEventInfo) SetEvent(v NwdafEvent)`

SetEvent sets Event field to given value.


### GetFailureCode

`func (o *FailureEventInfo) GetFailureCode() NwdafFailureCode`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *FailureEventInfo) GetFailureCodeOk() (*NwdafFailureCode, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *FailureEventInfo) SetFailureCode(v NwdafFailureCode)`

SetFailureCode sets FailureCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


