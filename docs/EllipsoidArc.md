# EllipsoidArc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Point** | [**GeographicalCoordinates**](GeographicalCoordinates.md) |  | 
**InnerRadius** | **int32** | Indicates value of the inner radius. | 
**UncertaintyRadius** | **float32** | Indicates value of uncertainty. | 
**OffsetAngle** | **int32** | Indicates value of angle. | 
**IncludedAngle** | **int32** | Indicates value of angle. | 
**Confidence** | **int32** | Indicates value of confidence. | 

## Methods

### NewEllipsoidArc

`func NewEllipsoidArc(point GeographicalCoordinates, innerRadius int32, uncertaintyRadius float32, offsetAngle int32, includedAngle int32, confidence int32, ) *EllipsoidArc`

NewEllipsoidArc instantiates a new EllipsoidArc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEllipsoidArcWithDefaults

`func NewEllipsoidArcWithDefaults() *EllipsoidArc`

NewEllipsoidArcWithDefaults instantiates a new EllipsoidArc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoint

`func (o *EllipsoidArc) GetPoint() GeographicalCoordinates`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *EllipsoidArc) GetPointOk() (*GeographicalCoordinates, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *EllipsoidArc) SetPoint(v GeographicalCoordinates)`

SetPoint sets Point field to given value.


### GetInnerRadius

`func (o *EllipsoidArc) GetInnerRadius() int32`

GetInnerRadius returns the InnerRadius field if non-nil, zero value otherwise.

### GetInnerRadiusOk

`func (o *EllipsoidArc) GetInnerRadiusOk() (*int32, bool)`

GetInnerRadiusOk returns a tuple with the InnerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerRadius

`func (o *EllipsoidArc) SetInnerRadius(v int32)`

SetInnerRadius sets InnerRadius field to given value.


### GetUncertaintyRadius

`func (o *EllipsoidArc) GetUncertaintyRadius() float32`

GetUncertaintyRadius returns the UncertaintyRadius field if non-nil, zero value otherwise.

### GetUncertaintyRadiusOk

`func (o *EllipsoidArc) GetUncertaintyRadiusOk() (*float32, bool)`

GetUncertaintyRadiusOk returns a tuple with the UncertaintyRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyRadius

`func (o *EllipsoidArc) SetUncertaintyRadius(v float32)`

SetUncertaintyRadius sets UncertaintyRadius field to given value.


### GetOffsetAngle

`func (o *EllipsoidArc) GetOffsetAngle() int32`

GetOffsetAngle returns the OffsetAngle field if non-nil, zero value otherwise.

### GetOffsetAngleOk

`func (o *EllipsoidArc) GetOffsetAngleOk() (*int32, bool)`

GetOffsetAngleOk returns a tuple with the OffsetAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetAngle

`func (o *EllipsoidArc) SetOffsetAngle(v int32)`

SetOffsetAngle sets OffsetAngle field to given value.


### GetIncludedAngle

`func (o *EllipsoidArc) GetIncludedAngle() int32`

GetIncludedAngle returns the IncludedAngle field if non-nil, zero value otherwise.

### GetIncludedAngleOk

`func (o *EllipsoidArc) GetIncludedAngleOk() (*int32, bool)`

GetIncludedAngleOk returns a tuple with the IncludedAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedAngle

`func (o *EllipsoidArc) SetIncludedAngle(v int32)`

SetIncludedAngle sets IncludedAngle field to given value.


### GetConfidence

`func (o *EllipsoidArc) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *EllipsoidArc) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *EllipsoidArc) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


