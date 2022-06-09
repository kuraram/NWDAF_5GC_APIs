# LocationAreaId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**Lac** | **string** |  | 

## Methods

### NewLocationAreaId

`func NewLocationAreaId(plmnId PlmnId, lac string, ) *LocationAreaId`

NewLocationAreaId instantiates a new LocationAreaId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAreaIdWithDefaults

`func NewLocationAreaIdWithDefaults() *LocationAreaId`

NewLocationAreaIdWithDefaults instantiates a new LocationAreaId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *LocationAreaId) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *LocationAreaId) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *LocationAreaId) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetLac

`func (o *LocationAreaId) GetLac() string`

GetLac returns the Lac field if non-nil, zero value otherwise.

### GetLacOk

`func (o *LocationAreaId) GetLacOk() (*string, bool)`

GetLacOk returns a tuple with the Lac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLac

`func (o *LocationAreaId) SetLac(v string)`

SetLac sets Lac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


