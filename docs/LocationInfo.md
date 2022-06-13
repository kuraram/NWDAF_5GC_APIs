# LocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loc** | [**UserLocation**](UserLocation.md) |  | 
**Ratio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewLocationInfo

`func NewLocationInfo(loc UserLocation, ) *LocationInfo`

NewLocationInfo instantiates a new LocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationInfoWithDefaults

`func NewLocationInfoWithDefaults() *LocationInfo`

NewLocationInfoWithDefaults instantiates a new LocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoc

`func (o *LocationInfo) GetLoc() UserLocation`

GetLoc returns the Loc field if non-nil, zero value otherwise.

### GetLocOk

`func (o *LocationInfo) GetLocOk() (*UserLocation, bool)`

GetLocOk returns a tuple with the Loc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoc

`func (o *LocationInfo) SetLoc(v UserLocation)`

SetLoc sets Loc field to given value.


### GetRatio

`func (o *LocationInfo) GetRatio() int32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *LocationInfo) GetRatioOk() (*int32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *LocationInfo) SetRatio(v int32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *LocationInfo) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetConfidence

`func (o *LocationInfo) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *LocationInfo) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *LocationInfo) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *LocationInfo) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


