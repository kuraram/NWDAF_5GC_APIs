# UncertaintyEllipse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SemiMajor** | **float32** | Indicates value of uncertainty. | 
**SemiMinor** | **float32** | Indicates value of uncertainty. | 
**OrientationMajor** | **int32** | Indicates value of orientation angle. | 

## Methods

### NewUncertaintyEllipse

`func NewUncertaintyEllipse(semiMajor float32, semiMinor float32, orientationMajor int32, ) *UncertaintyEllipse`

NewUncertaintyEllipse instantiates a new UncertaintyEllipse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUncertaintyEllipseWithDefaults

`func NewUncertaintyEllipseWithDefaults() *UncertaintyEllipse`

NewUncertaintyEllipseWithDefaults instantiates a new UncertaintyEllipse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSemiMajor

`func (o *UncertaintyEllipse) GetSemiMajor() float32`

GetSemiMajor returns the SemiMajor field if non-nil, zero value otherwise.

### GetSemiMajorOk

`func (o *UncertaintyEllipse) GetSemiMajorOk() (*float32, bool)`

GetSemiMajorOk returns a tuple with the SemiMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemiMajor

`func (o *UncertaintyEllipse) SetSemiMajor(v float32)`

SetSemiMajor sets SemiMajor field to given value.


### GetSemiMinor

`func (o *UncertaintyEllipse) GetSemiMinor() float32`

GetSemiMinor returns the SemiMinor field if non-nil, zero value otherwise.

### GetSemiMinorOk

`func (o *UncertaintyEllipse) GetSemiMinorOk() (*float32, bool)`

GetSemiMinorOk returns a tuple with the SemiMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemiMinor

`func (o *UncertaintyEllipse) SetSemiMinor(v float32)`

SetSemiMinor sets SemiMinor field to given value.


### GetOrientationMajor

`func (o *UncertaintyEllipse) GetOrientationMajor() int32`

GetOrientationMajor returns the OrientationMajor field if non-nil, zero value otherwise.

### GetOrientationMajorOk

`func (o *UncertaintyEllipse) GetOrientationMajorOk() (*int32, bool)`

GetOrientationMajorOk returns a tuple with the OrientationMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientationMajor

`func (o *UncertaintyEllipse) SetOrientationMajor(v int32)`

SetOrientationMajor sets OrientationMajor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


