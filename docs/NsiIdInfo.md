# NsiIdInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | [**Snssai**](Snssai.md) |  | 
**NsiIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNsiIdInfo

`func NewNsiIdInfo(snssai Snssai, ) *NsiIdInfo`

NewNsiIdInfo instantiates a new NsiIdInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsiIdInfoWithDefaults

`func NewNsiIdInfoWithDefaults() *NsiIdInfo`

NewNsiIdInfoWithDefaults instantiates a new NsiIdInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *NsiIdInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *NsiIdInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *NsiIdInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetNsiIds

`func (o *NsiIdInfo) GetNsiIds() []string`

GetNsiIds returns the NsiIds field if non-nil, zero value otherwise.

### GetNsiIdsOk

`func (o *NsiIdInfo) GetNsiIdsOk() (*[]string, bool)`

GetNsiIdsOk returns a tuple with the NsiIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiIds

`func (o *NsiIdInfo) SetNsiIds(v []string)`

SetNsiIds sets NsiIds field to given value.

### HasNsiIds

`func (o *NsiIdInfo) HasNsiIds() bool`

HasNsiIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


