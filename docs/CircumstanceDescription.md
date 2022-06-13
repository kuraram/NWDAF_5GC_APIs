# CircumstanceDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Freq** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 
**Tm** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**LocArea** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**Vol** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 

## Methods

### NewCircumstanceDescription

`func NewCircumstanceDescription() *CircumstanceDescription`

NewCircumstanceDescription instantiates a new CircumstanceDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircumstanceDescriptionWithDefaults

`func NewCircumstanceDescriptionWithDefaults() *CircumstanceDescription`

NewCircumstanceDescriptionWithDefaults instantiates a new CircumstanceDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreq

`func (o *CircumstanceDescription) GetFreq() float32`

GetFreq returns the Freq field if non-nil, zero value otherwise.

### GetFreqOk

`func (o *CircumstanceDescription) GetFreqOk() (*float32, bool)`

GetFreqOk returns a tuple with the Freq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreq

`func (o *CircumstanceDescription) SetFreq(v float32)`

SetFreq sets Freq field to given value.

### HasFreq

`func (o *CircumstanceDescription) HasFreq() bool`

HasFreq returns a boolean if a field has been set.

### GetTm

`func (o *CircumstanceDescription) GetTm() time.Time`

GetTm returns the Tm field if non-nil, zero value otherwise.

### GetTmOk

`func (o *CircumstanceDescription) GetTmOk() (*time.Time, bool)`

GetTmOk returns a tuple with the Tm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTm

`func (o *CircumstanceDescription) SetTm(v time.Time)`

SetTm sets Tm field to given value.

### HasTm

`func (o *CircumstanceDescription) HasTm() bool`

HasTm returns a boolean if a field has been set.

### GetLocArea

`func (o *CircumstanceDescription) GetLocArea() NetworkAreaInfo`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *CircumstanceDescription) GetLocAreaOk() (*NetworkAreaInfo, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *CircumstanceDescription) SetLocArea(v NetworkAreaInfo)`

SetLocArea sets LocArea field to given value.

### HasLocArea

`func (o *CircumstanceDescription) HasLocArea() bool`

HasLocArea returns a boolean if a field has been set.

### GetVol

`func (o *CircumstanceDescription) GetVol() int64`

GetVol returns the Vol field if non-nil, zero value otherwise.

### GetVolOk

`func (o *CircumstanceDescription) GetVolOk() (*int64, bool)`

GetVolOk returns a tuple with the Vol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol

`func (o *CircumstanceDescription) SetVol(v int64)`

SetVol sets Vol field to given value.

### HasVol

`func (o *CircumstanceDescription) HasVol() bool`

HasVol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


