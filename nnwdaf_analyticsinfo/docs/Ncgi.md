# Ncgi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**NrCellId** | **string** |  | 
**Nid** | Pointer to **string** |  | [optional] 

## Methods

### NewNcgi

`func NewNcgi(plmnId PlmnId, nrCellId string, ) *Ncgi`

NewNcgi instantiates a new Ncgi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNcgiWithDefaults

`func NewNcgiWithDefaults() *Ncgi`

NewNcgiWithDefaults instantiates a new Ncgi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *Ncgi) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *Ncgi) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *Ncgi) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetNrCellId

`func (o *Ncgi) GetNrCellId() string`

GetNrCellId returns the NrCellId field if non-nil, zero value otherwise.

### GetNrCellIdOk

`func (o *Ncgi) GetNrCellIdOk() (*string, bool)`

GetNrCellIdOk returns a tuple with the NrCellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrCellId

`func (o *Ncgi) SetNrCellId(v string)`

SetNrCellId sets NrCellId field to given value.


### GetNid

`func (o *Ncgi) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *Ncgi) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *Ncgi) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *Ncgi) HasNid() bool`

HasNid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


