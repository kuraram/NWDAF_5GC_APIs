# UmtTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeOfDay** | **string** | String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC). | 
**DayOfWeek** | **int32** | integer between and including 1 and 7 denoting a weekday. 1 shall indicate Monday, and the subsequent weekdays  shall be indicated with the next higher numbers. 7 shall indicate Sunday.  | 

## Methods

### NewUmtTime

`func NewUmtTime(timeOfDay string, dayOfWeek int32, ) *UmtTime`

NewUmtTime instantiates a new UmtTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmtTimeWithDefaults

`func NewUmtTimeWithDefaults() *UmtTime`

NewUmtTimeWithDefaults instantiates a new UmtTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeOfDay

`func (o *UmtTime) GetTimeOfDay() string`

GetTimeOfDay returns the TimeOfDay field if non-nil, zero value otherwise.

### GetTimeOfDayOk

`func (o *UmtTime) GetTimeOfDayOk() (*string, bool)`

GetTimeOfDayOk returns a tuple with the TimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDay

`func (o *UmtTime) SetTimeOfDay(v string)`

SetTimeOfDay sets TimeOfDay field to given value.


### GetDayOfWeek

`func (o *UmtTime) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *UmtTime) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *UmtTime) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


