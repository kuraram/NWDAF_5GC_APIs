# NsiLoadLevelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadLevelInformation** | **int32** | Load level information of the network slice and the optionally associated network slice instance. | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**NsiId** | Pointer to **string** |  | [optional] 

## Methods

### NewNsiLoadLevelInfo

`func NewNsiLoadLevelInfo(loadLevelInformation int32, snssai Snssai, ) *NsiLoadLevelInfo`

NewNsiLoadLevelInfo instantiates a new NsiLoadLevelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsiLoadLevelInfoWithDefaults

`func NewNsiLoadLevelInfoWithDefaults() *NsiLoadLevelInfo`

NewNsiLoadLevelInfoWithDefaults instantiates a new NsiLoadLevelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadLevelInformation

`func (o *NsiLoadLevelInfo) GetLoadLevelInformation() int32`

GetLoadLevelInformation returns the LoadLevelInformation field if non-nil, zero value otherwise.

### GetLoadLevelInformationOk

`func (o *NsiLoadLevelInfo) GetLoadLevelInformationOk() (*int32, bool)`

GetLoadLevelInformationOk returns a tuple with the LoadLevelInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadLevelInformation

`func (o *NsiLoadLevelInfo) SetLoadLevelInformation(v int32)`

SetLoadLevelInformation sets LoadLevelInformation field to given value.


### GetSnssai

`func (o *NsiLoadLevelInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *NsiLoadLevelInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *NsiLoadLevelInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetNsiId

`func (o *NsiLoadLevelInfo) GetNsiId() string`

GetNsiId returns the NsiId field if non-nil, zero value otherwise.

### GetNsiIdOk

`func (o *NsiLoadLevelInfo) GetNsiIdOk() (*string, bool)`

GetNsiIdOk returns a tuple with the NsiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiId

`func (o *NsiLoadLevelInfo) SetNsiId(v string)`

SetNsiId sets NsiId field to given value.

### HasNsiId

`func (o *NsiLoadLevelInfo) HasNsiId() bool`

HasNsiId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


