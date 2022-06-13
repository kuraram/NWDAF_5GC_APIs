# PointAltitudeUncertaintyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Point** | [**GeographicalCoordinates**](GeographicalCoordinates.md) |  | 
**Altitude** | **float64** | Indicates value of altitude. | 
**UncertaintyEllipse** | [**UncertaintyEllipse**](UncertaintyEllipse.md) |  | 
**UncertaintyAltitude** | **float32** | Indicates value of uncertainty. | 
**Confidence** | **int32** | Indicates value of confidence. | 

## Methods

### NewPointAltitudeUncertaintyAllOf

`func NewPointAltitudeUncertaintyAllOf(point GeographicalCoordinates, altitude float64, uncertaintyEllipse UncertaintyEllipse, uncertaintyAltitude float32, confidence int32, ) *PointAltitudeUncertaintyAllOf`

NewPointAltitudeUncertaintyAllOf instantiates a new PointAltitudeUncertaintyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointAltitudeUncertaintyAllOfWithDefaults

`func NewPointAltitudeUncertaintyAllOfWithDefaults() *PointAltitudeUncertaintyAllOf`

NewPointAltitudeUncertaintyAllOfWithDefaults instantiates a new PointAltitudeUncertaintyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoint

`func (o *PointAltitudeUncertaintyAllOf) GetPoint() GeographicalCoordinates`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *PointAltitudeUncertaintyAllOf) GetPointOk() (*GeographicalCoordinates, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *PointAltitudeUncertaintyAllOf) SetPoint(v GeographicalCoordinates)`

SetPoint sets Point field to given value.


### GetAltitude

`func (o *PointAltitudeUncertaintyAllOf) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *PointAltitudeUncertaintyAllOf) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *PointAltitudeUncertaintyAllOf) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.


### GetUncertaintyEllipse

`func (o *PointAltitudeUncertaintyAllOf) GetUncertaintyEllipse() UncertaintyEllipse`

GetUncertaintyEllipse returns the UncertaintyEllipse field if non-nil, zero value otherwise.

### GetUncertaintyEllipseOk

`func (o *PointAltitudeUncertaintyAllOf) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool)`

GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyEllipse

`func (o *PointAltitudeUncertaintyAllOf) SetUncertaintyEllipse(v UncertaintyEllipse)`

SetUncertaintyEllipse sets UncertaintyEllipse field to given value.


### GetUncertaintyAltitude

`func (o *PointAltitudeUncertaintyAllOf) GetUncertaintyAltitude() float32`

GetUncertaintyAltitude returns the UncertaintyAltitude field if non-nil, zero value otherwise.

### GetUncertaintyAltitudeOk

`func (o *PointAltitudeUncertaintyAllOf) GetUncertaintyAltitudeOk() (*float32, bool)`

GetUncertaintyAltitudeOk returns a tuple with the UncertaintyAltitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyAltitude

`func (o *PointAltitudeUncertaintyAllOf) SetUncertaintyAltitude(v float32)`

SetUncertaintyAltitude sets UncertaintyAltitude field to given value.


### GetConfidence

`func (o *PointAltitudeUncertaintyAllOf) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *PointAltitudeUncertaintyAllOf) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *PointAltitudeUncertaintyAllOf) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


