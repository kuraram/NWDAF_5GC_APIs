# ServiceExperienceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SvcExprc** | [**SvcExperience**](SvcExperience.md) |  | 
**SvcExprcVariance** | Pointer to **float32** |  | [optional] 
**Supis** | Pointer to **[]string** |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **int32** |  | [optional] 
**Dnn** | Pointer to **string** |  | [optional] 
**NetworkArea** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**NsiId** | Pointer to **string** |  | [optional] 
**Ratio** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceExperienceInfo

`func NewServiceExperienceInfo(svcExprc SvcExperience, ) *ServiceExperienceInfo`

NewServiceExperienceInfo instantiates a new ServiceExperienceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceExperienceInfoWithDefaults

`func NewServiceExperienceInfoWithDefaults() *ServiceExperienceInfo`

NewServiceExperienceInfoWithDefaults instantiates a new ServiceExperienceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSvcExprc

`func (o *ServiceExperienceInfo) GetSvcExprc() SvcExperience`

GetSvcExprc returns the SvcExprc field if non-nil, zero value otherwise.

### GetSvcExprcOk

`func (o *ServiceExperienceInfo) GetSvcExprcOk() (*SvcExperience, bool)`

GetSvcExprcOk returns a tuple with the SvcExprc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExprc

`func (o *ServiceExperienceInfo) SetSvcExprc(v SvcExperience)`

SetSvcExprc sets SvcExprc field to given value.


### GetSvcExprcVariance

`func (o *ServiceExperienceInfo) GetSvcExprcVariance() float32`

GetSvcExprcVariance returns the SvcExprcVariance field if non-nil, zero value otherwise.

### GetSvcExprcVarianceOk

`func (o *ServiceExperienceInfo) GetSvcExprcVarianceOk() (*float32, bool)`

GetSvcExprcVarianceOk returns a tuple with the SvcExprcVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExprcVariance

`func (o *ServiceExperienceInfo) SetSvcExprcVariance(v float32)`

SetSvcExprcVariance sets SvcExprcVariance field to given value.

### HasSvcExprcVariance

`func (o *ServiceExperienceInfo) HasSvcExprcVariance() bool`

HasSvcExprcVariance returns a boolean if a field has been set.

### GetSupis

`func (o *ServiceExperienceInfo) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *ServiceExperienceInfo) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *ServiceExperienceInfo) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *ServiceExperienceInfo) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetSnssai

`func (o *ServiceExperienceInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *ServiceExperienceInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *ServiceExperienceInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *ServiceExperienceInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetAppId

`func (o *ServiceExperienceInfo) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceExperienceInfo) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceExperienceInfo) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceExperienceInfo) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetConfidence

`func (o *ServiceExperienceInfo) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ServiceExperienceInfo) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ServiceExperienceInfo) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ServiceExperienceInfo) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetDnn

`func (o *ServiceExperienceInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *ServiceExperienceInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *ServiceExperienceInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *ServiceExperienceInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetNetworkArea

`func (o *ServiceExperienceInfo) GetNetworkArea() NetworkAreaInfo`

GetNetworkArea returns the NetworkArea field if non-nil, zero value otherwise.

### GetNetworkAreaOk

`func (o *ServiceExperienceInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool)`

GetNetworkAreaOk returns a tuple with the NetworkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkArea

`func (o *ServiceExperienceInfo) SetNetworkArea(v NetworkAreaInfo)`

SetNetworkArea sets NetworkArea field to given value.

### HasNetworkArea

`func (o *ServiceExperienceInfo) HasNetworkArea() bool`

HasNetworkArea returns a boolean if a field has been set.

### GetNsiId

`func (o *ServiceExperienceInfo) GetNsiId() string`

GetNsiId returns the NsiId field if non-nil, zero value otherwise.

### GetNsiIdOk

`func (o *ServiceExperienceInfo) GetNsiIdOk() (*string, bool)`

GetNsiIdOk returns a tuple with the NsiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiId

`func (o *ServiceExperienceInfo) SetNsiId(v string)`

SetNsiId sets NsiId field to given value.

### HasNsiId

`func (o *ServiceExperienceInfo) HasNsiId() bool`

HasNsiId returns a boolean if a field has been set.

### GetRatio

`func (o *ServiceExperienceInfo) GetRatio() int32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *ServiceExperienceInfo) GetRatioOk() (*int32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *ServiceExperienceInfo) SetRatio(v int32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *ServiceExperienceInfo) HasRatio() bool`

HasRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


