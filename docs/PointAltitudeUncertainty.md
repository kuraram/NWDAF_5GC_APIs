# PointAltitudeUncertainty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Point** | [**GeographicalCoordinates**](GeographicalCoordinates.md) |  | 
**Altitude** | **float64** | Indicates value of altitude. | 
**UncertaintyEllipse** | [**UncertaintyEllipse**](UncertaintyEllipse.md) |  | 
**UncertaintyAltitude** | **float32** | Indicates value of uncertainty. | 
**Confidence** | **int32** | Indicates value of confidence. | 

## Methods

### NewPointAltitudeUncertainty

`func NewPointAltitudeUncertainty(point GeographicalCoordinates, altitude float64, uncertaintyEllipse UncertaintyEllipse, uncertaintyAltitude float32, confidence int32, ) *PointAltitudeUncertainty`

NewPointAltitudeUncertainty instantiates a new PointAltitudeUncertainty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointAltitudeUncertaintyWithDefaults

`func NewPointAltitudeUncertaintyWithDefaults() *PointAltitudeUncertainty`

NewPointAltitudeUncertaintyWithDefaults instantiates a new PointAltitudeUncertainty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoint

`func (o *PointAltitudeUncertainty) GetPoint() GeographicalCoordinates`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *PointAltitudeUncertainty) GetPointOk() (*GeographicalCoordinates, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *PointAltitudeUncertainty) SetPoint(v GeographicalCoordinates)`

SetPoint sets Point field to given value.


### GetAltitude

`func (o *PointAltitudeUncertainty) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *PointAltitudeUncertainty) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *PointAltitudeUncertainty) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.


### GetUncertaintyEllipse

`func (o *PointAltitudeUncertainty) GetUncertaintyEllipse() UncertaintyEllipse`

GetUncertaintyEllipse returns the UncertaintyEllipse field if non-nil, zero value otherwise.

### GetUncertaintyEllipseOk

`func (o *PointAltitudeUncertainty) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool)`

GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyEllipse

`func (o *PointAltitudeUncertainty) SetUncertaintyEllipse(v UncertaintyEllipse)`

SetUncertaintyEllipse sets UncertaintyEllipse field to given value.


### GetUncertaintyAltitude

`func (o *PointAltitudeUncertainty) GetUncertaintyAltitude() float32`

GetUncertaintyAltitude returns the UncertaintyAltitude field if non-nil, zero value otherwise.

### GetUncertaintyAltitudeOk

`func (o *PointAltitudeUncertainty) GetUncertaintyAltitudeOk() (*float32, bool)`

GetUncertaintyAltitudeOk returns a tuple with the UncertaintyAltitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyAltitude

`func (o *PointAltitudeUncertainty) SetUncertaintyAltitude(v float32)`

SetUncertaintyAltitude sets UncertaintyAltitude field to given value.


### GetConfidence

`func (o *PointAltitudeUncertainty) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *PointAltitudeUncertainty) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *PointAltitudeUncertainty) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


