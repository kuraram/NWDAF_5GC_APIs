# CellGlobalId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**Lac** | **string** |  | 
**CellId** | **string** |  | 

## Methods

### NewCellGlobalId

`func NewCellGlobalId(plmnId PlmnId, lac string, cellId string, ) *CellGlobalId`

NewCellGlobalId instantiates a new CellGlobalId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellGlobalIdWithDefaults

`func NewCellGlobalIdWithDefaults() *CellGlobalId`

NewCellGlobalIdWithDefaults instantiates a new CellGlobalId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *CellGlobalId) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *CellGlobalId) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *CellGlobalId) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetLac

`func (o *CellGlobalId) GetLac() string`

GetLac returns the Lac field if non-nil, zero value otherwise.

### GetLacOk

`func (o *CellGlobalId) GetLacOk() (*string, bool)`

GetLacOk returns a tuple with the Lac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLac

`func (o *CellGlobalId) SetLac(v string)`

SetLac sets Lac field to given value.


### GetCellId

`func (o *CellGlobalId) GetCellId() string`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *CellGlobalId) GetCellIdOk() (*string, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *CellGlobalId) SetCellId(v string)`

SetCellId sets CellId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


