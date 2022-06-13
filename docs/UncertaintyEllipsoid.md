# UncertaintyEllipsoid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SemiMajor** | **float32** | Indicates value of uncertainty. | 
**SemiMinor** | **float32** | Indicates value of uncertainty. | 
**Vertical** | **float32** | Indicates value of uncertainty. | 
**OrientationMajor** | **int32** | Indicates value of orientation angle. | 

## Methods

### NewUncertaintyEllipsoid

`func NewUncertaintyEllipsoid(semiMajor float32, semiMinor float32, vertical float32, orientationMajor int32, ) *UncertaintyEllipsoid`

NewUncertaintyEllipsoid instantiates a new UncertaintyEllipsoid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUncertaintyEllipsoidWithDefaults

`func NewUncertaintyEllipsoidWithDefaults() *UncertaintyEllipsoid`

NewUncertaintyEllipsoidWithDefaults instantiates a new UncertaintyEllipsoid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSemiMajor

`func (o *UncertaintyEllipsoid) GetSemiMajor() float32`

GetSemiMajor returns the SemiMajor field if non-nil, zero value otherwise.

### GetSemiMajorOk

`func (o *UncertaintyEllipsoid) GetSemiMajorOk() (*float32, bool)`

GetSemiMajorOk returns a tuple with the SemiMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemiMajor

`func (o *UncertaintyEllipsoid) SetSemiMajor(v float32)`

SetSemiMajor sets SemiMajor field to given value.


### GetSemiMinor

`func (o *UncertaintyEllipsoid) GetSemiMinor() float32`

GetSemiMinor returns the SemiMinor field if non-nil, zero value otherwise.

### GetSemiMinorOk

`func (o *UncertaintyEllipsoid) GetSemiMinorOk() (*float32, bool)`

GetSemiMinorOk returns a tuple with the SemiMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemiMinor

`func (o *UncertaintyEllipsoid) SetSemiMinor(v float32)`

SetSemiMinor sets SemiMinor field to given value.


### GetVertical

`func (o *UncertaintyEllipsoid) GetVertical() float32`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *UncertaintyEllipsoid) GetVerticalOk() (*float32, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *UncertaintyEllipsoid) SetVertical(v float32)`

SetVertical sets Vertical field to given value.


### GetOrientationMajor

`func (o *UncertaintyEllipsoid) GetOrientationMajor() int32`

GetOrientationMajor returns the OrientationMajor field if non-nil, zero value otherwise.

### GetOrientationMajorOk

`func (o *UncertaintyEllipsoid) GetOrientationMajorOk() (*int32, bool)`

GetOrientationMajorOk returns a tuple with the OrientationMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientationMajor

`func (o *UncertaintyEllipsoid) SetOrientationMajor(v int32)`

SetOrientationMajor sets OrientationMajor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


