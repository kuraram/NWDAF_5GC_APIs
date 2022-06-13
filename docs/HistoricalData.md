# HistoricalData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**EndTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SubsWithSources** | Pointer to **[]string** | Information about subscriptions with the data sources. | [optional] 
**Data** | **[]string** | Historical data related to the analytics. | 

## Methods

### NewHistoricalData

`func NewHistoricalData(data []string, ) *HistoricalData`

NewHistoricalData instantiates a new HistoricalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalDataWithDefaults

`func NewHistoricalDataWithDefaults() *HistoricalData`

NewHistoricalDataWithDefaults instantiates a new HistoricalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *HistoricalData) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HistoricalData) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HistoricalData) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HistoricalData) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *HistoricalData) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HistoricalData) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HistoricalData) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HistoricalData) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetSubsWithSources

`func (o *HistoricalData) GetSubsWithSources() []string`

GetSubsWithSources returns the SubsWithSources field if non-nil, zero value otherwise.

### GetSubsWithSourcesOk

`func (o *HistoricalData) GetSubsWithSourcesOk() (*[]string, bool)`

GetSubsWithSourcesOk returns a tuple with the SubsWithSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsWithSources

`func (o *HistoricalData) SetSubsWithSources(v []string)`

SetSubsWithSources sets SubsWithSources field to given value.

### HasSubsWithSources

`func (o *HistoricalData) HasSubsWithSources() bool`

HasSubsWithSources returns a boolean if a field has been set.

### GetData

`func (o *HistoricalData) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoricalData) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoricalData) SetData(v []string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


