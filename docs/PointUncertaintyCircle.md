# PointUncertaintyCircle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Point** | [**GeographicalCoordinates**](GeographicalCoordinates.md) |  | 
**Uncertainty** | **float32** | Indicates value of uncertainty. | 

## Methods

### NewPointUncertaintyCircle

`func NewPointUncertaintyCircle(point GeographicalCoordinates, uncertainty float32, ) *PointUncertaintyCircle`

NewPointUncertaintyCircle instantiates a new PointUncertaintyCircle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointUncertaintyCircleWithDefaults

`func NewPointUncertaintyCircleWithDefaults() *PointUncertaintyCircle`

NewPointUncertaintyCircleWithDefaults instantiates a new PointUncertaintyCircle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoint

`func (o *PointUncertaintyCircle) GetPoint() GeographicalCoordinates`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *PointUncertaintyCircle) GetPointOk() (*GeographicalCoordinates, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *PointUncertaintyCircle) SetPoint(v GeographicalCoordinates)`

SetPoint sets Point field to given value.


### GetUncertainty

`func (o *PointUncertaintyCircle) GetUncertainty() float32`

GetUncertainty returns the Uncertainty field if non-nil, zero value otherwise.

### GetUncertaintyOk

`func (o *PointUncertaintyCircle) GetUncertaintyOk() (*float32, bool)`

GetUncertaintyOk returns a tuple with the Uncertainty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertainty

`func (o *PointUncertaintyCircle) SetUncertainty(v float32)`

SetUncertainty sets Uncertainty field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


