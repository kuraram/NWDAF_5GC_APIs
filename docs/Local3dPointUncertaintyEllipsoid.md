# Local3dPointUncertaintyEllipsoid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalOrigin** | [**LocalOrigin**](LocalOrigin.md) |  | 
**Point** | [**RelativeCartesianLocation**](RelativeCartesianLocation.md) |  | 
**UncertaintyEllipsoid** | [**UncertaintyEllipsoid**](UncertaintyEllipsoid.md) |  | 
**Confidence** | **int32** | Indicates value of confidence. | 

## Methods

### NewLocal3dPointUncertaintyEllipsoid

`func NewLocal3dPointUncertaintyEllipsoid(localOrigin LocalOrigin, point RelativeCartesianLocation, uncertaintyEllipsoid UncertaintyEllipsoid, confidence int32, ) *Local3dPointUncertaintyEllipsoid`

NewLocal3dPointUncertaintyEllipsoid instantiates a new Local3dPointUncertaintyEllipsoid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocal3dPointUncertaintyEllipsoidWithDefaults

`func NewLocal3dPointUncertaintyEllipsoidWithDefaults() *Local3dPointUncertaintyEllipsoid`

NewLocal3dPointUncertaintyEllipsoidWithDefaults instantiates a new Local3dPointUncertaintyEllipsoid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalOrigin

`func (o *Local3dPointUncertaintyEllipsoid) GetLocalOrigin() LocalOrigin`

GetLocalOrigin returns the LocalOrigin field if non-nil, zero value otherwise.

### GetLocalOriginOk

`func (o *Local3dPointUncertaintyEllipsoid) GetLocalOriginOk() (*LocalOrigin, bool)`

GetLocalOriginOk returns a tuple with the LocalOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalOrigin

`func (o *Local3dPointUncertaintyEllipsoid) SetLocalOrigin(v LocalOrigin)`

SetLocalOrigin sets LocalOrigin field to given value.


### GetPoint

`func (o *Local3dPointUncertaintyEllipsoid) GetPoint() RelativeCartesianLocation`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *Local3dPointUncertaintyEllipsoid) GetPointOk() (*RelativeCartesianLocation, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *Local3dPointUncertaintyEllipsoid) SetPoint(v RelativeCartesianLocation)`

SetPoint sets Point field to given value.


### GetUncertaintyEllipsoid

`func (o *Local3dPointUncertaintyEllipsoid) GetUncertaintyEllipsoid() UncertaintyEllipsoid`

GetUncertaintyEllipsoid returns the UncertaintyEllipsoid field if non-nil, zero value otherwise.

### GetUncertaintyEllipsoidOk

`func (o *Local3dPointUncertaintyEllipsoid) GetUncertaintyEllipsoidOk() (*UncertaintyEllipsoid, bool)`

GetUncertaintyEllipsoidOk returns a tuple with the UncertaintyEllipsoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyEllipsoid

`func (o *Local3dPointUncertaintyEllipsoid) SetUncertaintyEllipsoid(v UncertaintyEllipsoid)`

SetUncertaintyEllipsoid sets UncertaintyEllipsoid field to given value.


### GetConfidence

`func (o *Local3dPointUncertaintyEllipsoid) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *Local3dPointUncertaintyEllipsoid) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *Local3dPointUncertaintyEllipsoid) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


