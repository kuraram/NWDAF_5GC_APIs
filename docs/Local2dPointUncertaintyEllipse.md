# Local2dPointUncertaintyEllipse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalOrigin** | [**LocalOrigin**](LocalOrigin.md) |  | 
**Point** | [**RelativeCartesianLocation**](RelativeCartesianLocation.md) |  | 
**UncertaintyEllipse** | [**UncertaintyEllipse**](UncertaintyEllipse.md) |  | 
**Confidence** | **int32** | Indicates value of confidence. | 

## Methods

### NewLocal2dPointUncertaintyEllipse

`func NewLocal2dPointUncertaintyEllipse(localOrigin LocalOrigin, point RelativeCartesianLocation, uncertaintyEllipse UncertaintyEllipse, confidence int32, ) *Local2dPointUncertaintyEllipse`

NewLocal2dPointUncertaintyEllipse instantiates a new Local2dPointUncertaintyEllipse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocal2dPointUncertaintyEllipseWithDefaults

`func NewLocal2dPointUncertaintyEllipseWithDefaults() *Local2dPointUncertaintyEllipse`

NewLocal2dPointUncertaintyEllipseWithDefaults instantiates a new Local2dPointUncertaintyEllipse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalOrigin

`func (o *Local2dPointUncertaintyEllipse) GetLocalOrigin() LocalOrigin`

GetLocalOrigin returns the LocalOrigin field if non-nil, zero value otherwise.

### GetLocalOriginOk

`func (o *Local2dPointUncertaintyEllipse) GetLocalOriginOk() (*LocalOrigin, bool)`

GetLocalOriginOk returns a tuple with the LocalOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalOrigin

`func (o *Local2dPointUncertaintyEllipse) SetLocalOrigin(v LocalOrigin)`

SetLocalOrigin sets LocalOrigin field to given value.


### GetPoint

`func (o *Local2dPointUncertaintyEllipse) GetPoint() RelativeCartesianLocation`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *Local2dPointUncertaintyEllipse) GetPointOk() (*RelativeCartesianLocation, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *Local2dPointUncertaintyEllipse) SetPoint(v RelativeCartesianLocation)`

SetPoint sets Point field to given value.


### GetUncertaintyEllipse

`func (o *Local2dPointUncertaintyEllipse) GetUncertaintyEllipse() UncertaintyEllipse`

GetUncertaintyEllipse returns the UncertaintyEllipse field if non-nil, zero value otherwise.

### GetUncertaintyEllipseOk

`func (o *Local2dPointUncertaintyEllipse) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool)`

GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncertaintyEllipse

`func (o *Local2dPointUncertaintyEllipse) SetUncertaintyEllipse(v UncertaintyEllipse)`

SetUncertaintyEllipse sets UncertaintyEllipse field to given value.


### GetConfidence

`func (o *Local2dPointUncertaintyEllipse) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *Local2dPointUncertaintyEllipse) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *Local2dPointUncertaintyEllipse) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


