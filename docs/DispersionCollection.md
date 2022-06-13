# DispersionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeLoc** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Supis** | Pointer to **[]string** |  | [optional] 
**Gpsis** | Pointer to **[]string** |  | [optional] 
**AppVolumes** | Pointer to [**[]ApplicationVolume**](ApplicationVolume.md) |  | [optional] 
**DisperAmount** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**DisperClass** | Pointer to [**DispersionClass**](DispersionClass.md) |  | [optional] 
**UsageRank** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**PercentileRank** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | [optional] 
**UeRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewDispersionCollection

`func NewDispersionCollection() *DispersionCollection`

NewDispersionCollection instantiates a new DispersionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDispersionCollectionWithDefaults

`func NewDispersionCollectionWithDefaults() *DispersionCollection`

NewDispersionCollectionWithDefaults instantiates a new DispersionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeLoc

`func (o *DispersionCollection) GetUeLoc() UserLocation`

GetUeLoc returns the UeLoc field if non-nil, zero value otherwise.

### GetUeLocOk

`func (o *DispersionCollection) GetUeLocOk() (*UserLocation, bool)`

GetUeLocOk returns a tuple with the UeLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLoc

`func (o *DispersionCollection) SetUeLoc(v UserLocation)`

SetUeLoc sets UeLoc field to given value.

### HasUeLoc

`func (o *DispersionCollection) HasUeLoc() bool`

HasUeLoc returns a boolean if a field has been set.

### GetSnssai

`func (o *DispersionCollection) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *DispersionCollection) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *DispersionCollection) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *DispersionCollection) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetSupis

`func (o *DispersionCollection) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *DispersionCollection) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *DispersionCollection) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *DispersionCollection) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetGpsis

`func (o *DispersionCollection) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *DispersionCollection) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *DispersionCollection) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *DispersionCollection) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetAppVolumes

`func (o *DispersionCollection) GetAppVolumes() []ApplicationVolume`

GetAppVolumes returns the AppVolumes field if non-nil, zero value otherwise.

### GetAppVolumesOk

`func (o *DispersionCollection) GetAppVolumesOk() (*[]ApplicationVolume, bool)`

GetAppVolumesOk returns a tuple with the AppVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVolumes

`func (o *DispersionCollection) SetAppVolumes(v []ApplicationVolume)`

SetAppVolumes sets AppVolumes field to given value.

### HasAppVolumes

`func (o *DispersionCollection) HasAppVolumes() bool`

HasAppVolumes returns a boolean if a field has been set.

### GetDisperAmount

`func (o *DispersionCollection) GetDisperAmount() int32`

GetDisperAmount returns the DisperAmount field if non-nil, zero value otherwise.

### GetDisperAmountOk

`func (o *DispersionCollection) GetDisperAmountOk() (*int32, bool)`

GetDisperAmountOk returns a tuple with the DisperAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperAmount

`func (o *DispersionCollection) SetDisperAmount(v int32)`

SetDisperAmount sets DisperAmount field to given value.

### HasDisperAmount

`func (o *DispersionCollection) HasDisperAmount() bool`

HasDisperAmount returns a boolean if a field has been set.

### GetDisperClass

`func (o *DispersionCollection) GetDisperClass() DispersionClass`

GetDisperClass returns the DisperClass field if non-nil, zero value otherwise.

### GetDisperClassOk

`func (o *DispersionCollection) GetDisperClassOk() (*DispersionClass, bool)`

GetDisperClassOk returns a tuple with the DisperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperClass

`func (o *DispersionCollection) SetDisperClass(v DispersionClass)`

SetDisperClass sets DisperClass field to given value.

### HasDisperClass

`func (o *DispersionCollection) HasDisperClass() bool`

HasDisperClass returns a boolean if a field has been set.

### GetUsageRank

`func (o *DispersionCollection) GetUsageRank() int32`

GetUsageRank returns the UsageRank field if non-nil, zero value otherwise.

### GetUsageRankOk

`func (o *DispersionCollection) GetUsageRankOk() (*int32, bool)`

GetUsageRankOk returns a tuple with the UsageRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRank

`func (o *DispersionCollection) SetUsageRank(v int32)`

SetUsageRank sets UsageRank field to given value.

### HasUsageRank

`func (o *DispersionCollection) HasUsageRank() bool`

HasUsageRank returns a boolean if a field has been set.

### GetPercentileRank

`func (o *DispersionCollection) GetPercentileRank() int32`

GetPercentileRank returns the PercentileRank field if non-nil, zero value otherwise.

### GetPercentileRankOk

`func (o *DispersionCollection) GetPercentileRankOk() (*int32, bool)`

GetPercentileRankOk returns a tuple with the PercentileRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileRank

`func (o *DispersionCollection) SetPercentileRank(v int32)`

SetPercentileRank sets PercentileRank field to given value.

### HasPercentileRank

`func (o *DispersionCollection) HasPercentileRank() bool`

HasPercentileRank returns a boolean if a field has been set.

### GetUeRatio

`func (o *DispersionCollection) GetUeRatio() int32`

GetUeRatio returns the UeRatio field if non-nil, zero value otherwise.

### GetUeRatioOk

`func (o *DispersionCollection) GetUeRatioOk() (*int32, bool)`

GetUeRatioOk returns a tuple with the UeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeRatio

`func (o *DispersionCollection) SetUeRatio(v int32)`

SetUeRatio sets UeRatio field to given value.

### HasUeRatio

`func (o *DispersionCollection) HasUeRatio() bool`

HasUeRatio returns a boolean if a field has been set.

### GetConfidence

`func (o *DispersionCollection) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *DispersionCollection) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *DispersionCollection) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *DispersionCollection) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


