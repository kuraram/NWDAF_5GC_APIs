# GeographicArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shape** | [**SupportedGADShapes**](SupportedGADShapes.md) |  | 
**Point** | [**GeographicalCoordinates**](GeographicalCoordinates.md) |  | 
**Uncertainty** | **float32** | Indicates value of uncertainty. | 
**UncertaintyEllipse** | [**UncertaintyEllipse**](UncertaintyEllipse.md) |  | 
**Confidence** | **int32** | Indicates value of confidence. | 
**PointList** | [**[]GeographicalCoordinates**](GeographicalCoordinates.md) | List of points. | 
**Altitude** | **float64** | Indicates value of altitude. | 
**UncertaintyAltitude** | **float32** | Indicates value of uncertainty. | 
**InnerRadius** | **int32** | Indicates value of the inner radius. | 
**UncertaintyRadius** | **float32** | Indicates value of uncertainty. | 
**OffsetAngle** | **int32** | Indicates value of angle. | 
**IncludedAngle** | **int32** | Indicates value of angle. | 

## Methods

### NewGeographicArea

`func NewGeographicArea(shape SupportedGADShapes, point GeographicalCoordinates, uncertainty float32, uncertaintyEllipse UncertaintyEllipse, confidence int32, pointList []GeographicalCoordinates, altitude float64, uncertaintyAltitude float32, innerRadius int32, uncertaintyRadius float32, offsetAngle int32, includedAngle int32, ) *GeographicArea`

NewGeographicArea instantiates a new GeographicArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeographicAreaWithDefaults

`func NewGeographicAreaWithDefaults() *GeographicArea`

NewGeographicAreaWithDefaults instantiates a new GeographicArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShape

`func (o *GeographicArea) GetShape() SupportedGADShapes`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *GeographicArea) GetShapeOk() (*SupportedGADShapes, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *GeographicArea) SetShape(v SupportedGADShapes)`

SetShape sets Shape field to given value.


### GetPoint

`func (o *GeographicArea) GetPoint() GeographicalCoordinates`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *GeographicArea) GetPointOk() (*GeographicalCoordinates, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *GeographicArea) SetPoint(v GeographicalCoordinates)`

SetPoint sets Point field to given value.


### GetUncertainty

`func (o *GeographicArea) GetUncertainty() float32`

GetUncertainty returns the Uncertainty field if non-nil, zero value otherwise.

### GetUncertaintyOk

`func (o *GeographicArea) GetUncertaintyOk() (*float32, bool)`

GetUncertaintyOk returns a tuple with the Uncertainty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertainty

`func (o *GeographicArea) SetUncertainty(v float32)`

SetUncertainty sets Uncertainty field to given value.


### GetUncertaintyEllipse

`func (o *GeographicArea) GetUncertaintyEllipse() UncertaintyEllipse`

GetUncertaintyEllipse returns the UncertaintyEllipse field if non-nil, zero value otherwise.

### GetUncertaintyEllipseOk

`func (o *GeographicArea) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool)`

GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyEllipse

`func (o *GeographicArea) SetUncertaintyEllipse(v UncertaintyEllipse)`

SetUncertaintyEllipse sets UncertaintyEllipse field to given value.


### GetConfidence

`func (o *GeographicArea) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *GeographicArea) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *GeographicArea) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.


### GetPointList

`func (o *GeographicArea) GetPointList() []GeographicalCoordinates`

GetPointList returns the PointList field if non-nil, zero value otherwise.

### GetPointListOk

`func (o *GeographicArea) GetPointListOk() (*[]GeographicalCoordinates, bool)`

GetPointListOk returns a tuple with the PointList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointList

`func (o *GeographicArea) SetPointList(v []GeographicalCoordinates)`

SetPointList sets PointList field to given value.


### GetAltitude

`func (o *GeographicArea) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *GeographicArea) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *GeographicArea) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.


### GetUncertaintyAltitude

`func (o *GeographicArea) GetUncertaintyAltitude() float32`

GetUncertaintyAltitude returns the UncertaintyAltitude field if non-nil, zero value otherwise.

### GetUncertaintyAltitudeOk

`func (o *GeographicArea) GetUncertaintyAltitudeOk() (*float32, bool)`

GetUncertaintyAltitudeOk returns a tuple with the UncertaintyAltitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyAltitude

`func (o *GeographicArea) SetUncertaintyAltitude(v float32)`

SetUncertaintyAltitude sets UncertaintyAltitude field to given value.


### GetInnerRadius

`func (o *GeographicArea) GetInnerRadius() int32`

GetInnerRadius returns the InnerRadius field if non-nil, zero value otherwise.

### GetInnerRadiusOk

`func (o *GeographicArea) GetInnerRadiusOk() (*int32, bool)`

GetInnerRadiusOk returns a tuple with the InnerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerRadius

`func (o *GeographicArea) SetInnerRadius(v int32)`

SetInnerRadius sets InnerRadius field to given value.


### GetUncertaintyRadius

`func (o *GeographicArea) GetUncertaintyRadius() float32`

GetUncertaintyRadius returns the UncertaintyRadius field if non-nil, zero value otherwise.

### GetUncertaintyRadiusOk

`func (o *GeographicArea) GetUncertaintyRadiusOk() (*float32, bool)`

GetUncertaintyRadiusOk returns a tuple with the UncertaintyRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyRadius

`func (o *GeographicArea) SetUncertaintyRadius(v float32)`

SetUncertaintyRadius sets UncertaintyRadius field to given value.


### GetOffsetAngle

`func (o *GeographicArea) GetOffsetAngle() int32`

GetOffsetAngle returns the OffsetAngle field if non-nil, zero value otherwise.

### GetOffsetAngleOk

`func (o *GeographicArea) GetOffsetAngleOk() (*int32, bool)`

GetOffsetAngleOk returns a tuple with the OffsetAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetAngle

`func (o *GeographicArea) SetOffsetAngle(v int32)`

SetOffsetAngle sets OffsetAngle field to given value.


### GetIncludedAngle

`func (o *GeographicArea) GetIncludedAngle() int32`

GetIncludedAngle returns the IncludedAngle field if non-nil, zero value otherwise.

### GetIncludedAngleOk

`func (o *GeographicArea) GetIncludedAngleOk() (*int32, bool)`

GetIncludedAngleOk returns a tuple with the IncludedAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedAngle

`func (o *GeographicArea) SetIncludedAngle(v int32)`

SetIncludedAngle sets IncludedAngle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


