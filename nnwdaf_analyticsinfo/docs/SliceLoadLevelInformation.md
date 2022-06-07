# SliceLoadLevelInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadLevelInformation** | **int32** | Load level information of the network slice and the optionally associated network slice instance. | 
**Snssais** | [**[]Snssai**](Snssai.md) | Identification(s) of network slice to which the subscription applies. | 

## Methods

### NewSliceLoadLevelInformation

`func NewSliceLoadLevelInformation(loadLevelInformation int32, snssais []Snssai, ) *SliceLoadLevelInformation`

NewSliceLoadLevelInformation instantiates a new SliceLoadLevelInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceLoadLevelInformationWithDefaults

`func NewSliceLoadLevelInformationWithDefaults() *SliceLoadLevelInformation`

NewSliceLoadLevelInformationWithDefaults instantiates a new SliceLoadLevelInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadLevelInformation

`func (o *SliceLoadLevelInformation) GetLoadLevelInformation() int32`

GetLoadLevelInformation returns the LoadLevelInformation field if non-nil, zero value otherwise.

### GetLoadLevelInformationOk

`func (o *SliceLoadLevelInformation) GetLoadLevelInformationOk() (*int32, bool)`

GetLoadLevelInformationOk returns a tuple with the LoadLevelInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadLevelInformation

`func (o *SliceLoadLevelInformation) SetLoadLevelInformation(v int32)`

SetLoadLevelInformation sets LoadLevelInformation field to given value.


### GetSnssais

`func (o *SliceLoadLevelInformation) GetSnssais() []Snssai`

GetSnssais returns the Snssais field if non-nil, zero value otherwise.

### GetSnssaisOk

`func (o *SliceLoadLevelInformation) GetSnssaisOk() (*[]Snssai, bool)`

GetSnssaisOk returns a tuple with the Snssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssais

`func (o *SliceLoadLevelInformation) SetSnssais(v []Snssai)`

SetSnssais sets Snssais field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


