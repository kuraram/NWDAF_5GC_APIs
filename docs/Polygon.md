# Polygon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PointList** | [**[]GeographicalCoordinates**](GeographicalCoordinates.md) | List of points. | 

## Methods

### NewPolygon

`func NewPolygon(pointList []GeographicalCoordinates, ) *Polygon`

NewPolygon instantiates a new Polygon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolygonWithDefaults

`func NewPolygonWithDefaults() *Polygon`

NewPolygonWithDefaults instantiates a new Polygon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPointList

`func (o *Polygon) GetPointList() []GeographicalCoordinates`

GetPointList returns the PointList field if non-nil, zero value otherwise.

### GetPointListOk

`func (o *Polygon) GetPointListOk() (*[]GeographicalCoordinates, bool)`

GetPointListOk returns a tuple with the PointList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointList

`func (o *Polygon) SetPointList(v []GeographicalCoordinates)`

SetPointList sets PointList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


