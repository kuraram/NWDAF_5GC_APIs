# RoutingAreaId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**Lac** | **string** | Location Area Code | 
**Rac** | **string** | Routing Area Code | 

## Methods

### NewRoutingAreaId

`func NewRoutingAreaId(plmnId PlmnId, lac string, rac string, ) *RoutingAreaId`

NewRoutingAreaId instantiates a new RoutingAreaId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingAreaIdWithDefaults

`func NewRoutingAreaIdWithDefaults() *RoutingAreaId`

NewRoutingAreaIdWithDefaults instantiates a new RoutingAreaId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *RoutingAreaId) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *RoutingAreaId) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *RoutingAreaId) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetLac

`func (o *RoutingAreaId) GetLac() string`

GetLac returns the Lac field if non-nil, zero value otherwise.

### GetLacOk

`func (o *RoutingAreaId) GetLacOk() (*string, bool)`

GetLacOk returns a tuple with the Lac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLac

`func (o *RoutingAreaId) SetLac(v string)`

SetLac sets Lac field to given value.


### GetRac

`func (o *RoutingAreaId) GetRac() string`

GetRac returns the Rac field if non-nil, zero value otherwise.

### GetRacOk

`func (o *RoutingAreaId) GetRacOk() (*string, bool)`

GetRacOk returns a tuple with the Rac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRac

`func (o *RoutingAreaId) SetRac(v string)`

SetRac sets Rac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


