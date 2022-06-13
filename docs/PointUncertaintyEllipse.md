# PointUncertaintyEllipse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Point** | [**GeographicalCoordinates**](GeographicalCoordinates.md) |  | 
**UncertaintyEllipse** | [**UncertaintyEllipse**](UncertaintyEllipse.md) |  | 
**Confidence** | **int32** | Indicates value of confidence. | 

## Methods

### NewPointUncertaintyEllipse

`func NewPointUncertaintyEllipse(point GeographicalCoordinates, uncertaintyEllipse UncertaintyEllipse, confidence int32, ) *PointUncertaintyEllipse`

NewPointUncertaintyEllipse instantiates a new PointUncertaintyEllipse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointUncertaintyEllipseWithDefaults

`func NewPointUncertaintyEllipseWithDefaults() *PointUncertaintyEllipse`

NewPointUncertaintyEllipseWithDefaults instantiates a new PointUncertaintyEllipse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoint

`func (o *PointUncertaintyEllipse) GetPoint() GeographicalCoordinates`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *PointUncertaintyEllipse) GetPointOk() (*GeographicalCoordinates, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *PointUncertaintyEllipse) SetPoint(v GeographicalCoordinates)`

SetPoint sets Point field to given value.


### GetUncertaintyEllipse

`func (o *PointUncertaintyEllipse) GetUncertaintyEllipse() UncertaintyEllipse`

GetUncertaintyEllipse returns the UncertaintyEllipse field if non-nil, zero value otherwise.

### GetUncertaintyEllipseOk

`func (o *PointUncertaintyEllipse) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool)`

GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyEllipse

`func (o *PointUncertaintyEllipse) SetUncertaintyEllipse(v UncertaintyEllipse)`

SetUncertaintyEllipse sets UncertaintyEllipse field to given value.


### GetConfidence

`func (o *PointUncertaintyEllipse) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *PointUncertaintyEllipse) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *PointUncertaintyEllipse) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


