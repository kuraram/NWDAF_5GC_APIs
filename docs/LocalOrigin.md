# LocalOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoordinateId** | Pointer to **string** |  | [optional] 
**Point** | Pointer to [**GeographicalCoordinates**](GeographicalCoordinates.md) |  | [optional] 

## Methods

### NewLocalOrigin

`func NewLocalOrigin() *LocalOrigin`

NewLocalOrigin instantiates a new LocalOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalOriginWithDefaults

`func NewLocalOriginWithDefaults() *LocalOrigin`

NewLocalOriginWithDefaults instantiates a new LocalOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinateId

`func (o *LocalOrigin) GetCoordinateId() string`

GetCoordinateId returns the CoordinateId field if non-nil, zero value otherwise.

### GetCoordinateIdOk

`func (o *LocalOrigin) GetCoordinateIdOk() (*string, bool)`

GetCoordinateIdOk returns a tuple with the CoordinateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinateId

`func (o *LocalOrigin) SetCoordinateId(v string)`

SetCoordinateId sets CoordinateId field to given value.

### HasCoordinateId

`func (o *LocalOrigin) HasCoordinateId() bool`

HasCoordinateId returns a boolean if a field has been set.

### GetPoint

`func (o *LocalOrigin) GetPoint() GeographicalCoordinates`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *LocalOrigin) GetPointOk() (*GeographicalCoordinates, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *LocalOrigin) SetPoint(v GeographicalCoordinates)`

SetPoint sets Point field to given value.

### HasPoint

`func (o *LocalOrigin) HasPoint() bool`

HasPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


