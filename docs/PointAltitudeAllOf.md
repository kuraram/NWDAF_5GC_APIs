# PointAltitudeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Point** | [**GeographicalCoordinates**](GeographicalCoordinates.md) |  | 
**Altitude** | **float64** | Indicates value of altitude. | 

## Methods

### NewPointAltitudeAllOf

`func NewPointAltitudeAllOf(point GeographicalCoordinates, altitude float64, ) *PointAltitudeAllOf`

NewPointAltitudeAllOf instantiates a new PointAltitudeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointAltitudeAllOfWithDefaults

`func NewPointAltitudeAllOfWithDefaults() *PointAltitudeAllOf`

NewPointAltitudeAllOfWithDefaults instantiates a new PointAltitudeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoint

`func (o *PointAltitudeAllOf) GetPoint() GeographicalCoordinates`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *PointAltitudeAllOf) GetPointOk() (*GeographicalCoordinates, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *PointAltitudeAllOf) SetPoint(v GeographicalCoordinates)`

SetPoint sets Point field to given value.


### GetAltitude

`func (o *PointAltitudeAllOf) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *PointAltitudeAllOf) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *PointAltitudeAllOf) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


