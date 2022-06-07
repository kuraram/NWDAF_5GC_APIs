# Exception

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcepId** | [**ExceptionId**](ExceptionId.md) |  | 
**ExcepLevel** | Pointer to **int32** |  | [optional] 
**ExcepTrend** | Pointer to [**ExceptionTrend**](ExceptionTrend.md) |  | [optional] 

## Methods

### NewException

`func NewException(excepId ExceptionId, ) *Exception`

NewException instantiates a new Exception object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionWithDefaults

`func NewExceptionWithDefaults() *Exception`

NewExceptionWithDefaults instantiates a new Exception object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcepId

`func (o *Exception) GetExcepId() ExceptionId`

GetExcepId returns the ExcepId field if non-nil, zero value otherwise.

### GetExcepIdOk

`func (o *Exception) GetExcepIdOk() (*ExceptionId, bool)`

GetExcepIdOk returns a tuple with the ExcepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepId

`func (o *Exception) SetExcepId(v ExceptionId)`

SetExcepId sets ExcepId field to given value.


### GetExcepLevel

`func (o *Exception) GetExcepLevel() int32`

GetExcepLevel returns the ExcepLevel field if non-nil, zero value otherwise.

### GetExcepLevelOk

`func (o *Exception) GetExcepLevelOk() (*int32, bool)`

GetExcepLevelOk returns a tuple with the ExcepLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepLevel

`func (o *Exception) SetExcepLevel(v int32)`

SetExcepLevel sets ExcepLevel field to given value.

### HasExcepLevel

`func (o *Exception) HasExcepLevel() bool`

HasExcepLevel returns a boolean if a field has been set.

### GetExcepTrend

`func (o *Exception) GetExcepTrend() ExceptionTrend`

GetExcepTrend returns the ExcepTrend field if non-nil, zero value otherwise.

### GetExcepTrendOk

`func (o *Exception) GetExcepTrendOk() (*ExceptionTrend, bool)`

GetExcepTrendOk returns a tuple with the ExcepTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepTrend

`func (o *Exception) SetExcepTrend(v ExceptionTrend)`

SetExcepTrend sets ExcepTrend field to given value.

### HasExcepTrend

`func (o *Exception) HasExcepTrend() bool`

HasExcepTrend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


