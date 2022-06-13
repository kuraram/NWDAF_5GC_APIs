# DispersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TsStart** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**TsDuration** | **int32** | indicating a time in seconds. | 
**DisperCollects** | [**[]DispersionCollection**](DispersionCollection.md) |  | 
**DisperType** | [**DispersionType**](DispersionType.md) |  | 

## Methods

### NewDispersionInfo

`func NewDispersionInfo(tsStart time.Time, tsDuration int32, disperCollects []DispersionCollection, disperType DispersionType, ) *DispersionInfo`

NewDispersionInfo instantiates a new DispersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDispersionInfoWithDefaults

`func NewDispersionInfoWithDefaults() *DispersionInfo`

NewDispersionInfoWithDefaults instantiates a new DispersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTsStart

`func (o *DispersionInfo) GetTsStart() time.Time`

GetTsStart returns the TsStart field if non-nil, zero value otherwise.

### GetTsStartOk

`func (o *DispersionInfo) GetTsStartOk() (*time.Time, bool)`

GetTsStartOk returns a tuple with the TsStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsStart

`func (o *DispersionInfo) SetTsStart(v time.Time)`

SetTsStart sets TsStart field to given value.


### GetTsDuration

`func (o *DispersionInfo) GetTsDuration() int32`

GetTsDuration returns the TsDuration field if non-nil, zero value otherwise.

### GetTsDurationOk

`func (o *DispersionInfo) GetTsDurationOk() (*int32, bool)`

GetTsDurationOk returns a tuple with the TsDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsDuration

`func (o *DispersionInfo) SetTsDuration(v int32)`

SetTsDuration sets TsDuration field to given value.


### GetDisperCollects

`func (o *DispersionInfo) GetDisperCollects() []DispersionCollection`

GetDisperCollects returns the DisperCollects field if non-nil, zero value otherwise.

### GetDisperCollectsOk

`func (o *DispersionInfo) GetDisperCollectsOk() (*[]DispersionCollection, bool)`

GetDisperCollectsOk returns a tuple with the DisperCollects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperCollects

`func (o *DispersionInfo) SetDisperCollects(v []DispersionCollection)`

SetDisperCollects sets DisperCollects field to given value.


### GetDisperType

`func (o *DispersionInfo) GetDisperType() DispersionType`

GetDisperType returns the DisperType field if non-nil, zero value otherwise.

### GetDisperTypeOk

`func (o *DispersionInfo) GetDisperTypeOk() (*DispersionType, bool)`

GetDisperTypeOk returns a tuple with the DisperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperType

`func (o *DispersionInfo) SetDisperType(v DispersionType)`

SetDisperType sets DisperType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


