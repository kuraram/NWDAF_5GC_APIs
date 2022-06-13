# NsiLoadLevelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadLevelInformation** | **int32** | Load level information of the network slice and the optionally associated network slice instance. | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**NsiId** | Pointer to **string** | Contains the Identifier of the selected Network Slice instance | [optional] 
**ResUsage** | Pointer to [**ResourceUsage**](ResourceUsage.md) |  | [optional] 
**NumOfExceedLoadLevelThr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**ExceedLoadLevelThrInd** | Pointer to **bool** |  | [optional] 
**NetworkArea** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**TimePeriod** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**NumOfUes** | Pointer to [**NumberAverage**](NumberAverage.md) |  | [optional] 
**NumOfPduSess** | Pointer to [**NumberAverage**](NumberAverage.md) |  | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

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

### GetResUsage

`func (o *NsiLoadLevelInfo) GetResUsage() ResourceUsage`

GetResUsage returns the ResUsage field if non-nil, zero value otherwise.

### GetResUsageOk

`func (o *NsiLoadLevelInfo) GetResUsageOk() (*ResourceUsage, bool)`

GetResUsageOk returns a tuple with the ResUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResUsage

`func (o *NsiLoadLevelInfo) SetResUsage(v ResourceUsage)`

SetResUsage sets ResUsage field to given value.

### HasResUsage

`func (o *NsiLoadLevelInfo) HasResUsage() bool`

HasResUsage returns a boolean if a field has been set.

### GetNumOfExceedLoadLevelThr

`func (o *NsiLoadLevelInfo) GetNumOfExceedLoadLevelThr() int32`

GetNumOfExceedLoadLevelThr returns the NumOfExceedLoadLevelThr field if non-nil, zero value otherwise.

### GetNumOfExceedLoadLevelThrOk

`func (o *NsiLoadLevelInfo) GetNumOfExceedLoadLevelThrOk() (*int32, bool)`

GetNumOfExceedLoadLevelThrOk returns a tuple with the NumOfExceedLoadLevelThr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfExceedLoadLevelThr

`func (o *NsiLoadLevelInfo) SetNumOfExceedLoadLevelThr(v int32)`

SetNumOfExceedLoadLevelThr sets NumOfExceedLoadLevelThr field to given value.

### HasNumOfExceedLoadLevelThr

`func (o *NsiLoadLevelInfo) HasNumOfExceedLoadLevelThr() bool`

HasNumOfExceedLoadLevelThr returns a boolean if a field has been set.

### GetExceedLoadLevelThrInd

`func (o *NsiLoadLevelInfo) GetExceedLoadLevelThrInd() bool`

GetExceedLoadLevelThrInd returns the ExceedLoadLevelThrInd field if non-nil, zero value otherwise.

### GetExceedLoadLevelThrIndOk

`func (o *NsiLoadLevelInfo) GetExceedLoadLevelThrIndOk() (*bool, bool)`

GetExceedLoadLevelThrIndOk returns a tuple with the ExceedLoadLevelThrInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceedLoadLevelThrInd

`func (o *NsiLoadLevelInfo) SetExceedLoadLevelThrInd(v bool)`

SetExceedLoadLevelThrInd sets ExceedLoadLevelThrInd field to given value.

### HasExceedLoadLevelThrInd

`func (o *NsiLoadLevelInfo) HasExceedLoadLevelThrInd() bool`

HasExceedLoadLevelThrInd returns a boolean if a field has been set.

### GetNetworkArea

`func (o *NsiLoadLevelInfo) GetNetworkArea() NetworkAreaInfo`

GetNetworkArea returns the NetworkArea field if non-nil, zero value otherwise.

### GetNetworkAreaOk

`func (o *NsiLoadLevelInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool)`

GetNetworkAreaOk returns a tuple with the NetworkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkArea

`func (o *NsiLoadLevelInfo) SetNetworkArea(v NetworkAreaInfo)`

SetNetworkArea sets NetworkArea field to given value.

### HasNetworkArea

`func (o *NsiLoadLevelInfo) HasNetworkArea() bool`

HasNetworkArea returns a boolean if a field has been set.

### GetTimePeriod

`func (o *NsiLoadLevelInfo) GetTimePeriod() TimeWindow`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *NsiLoadLevelInfo) GetTimePeriodOk() (*TimeWindow, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *NsiLoadLevelInfo) SetTimePeriod(v TimeWindow)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *NsiLoadLevelInfo) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetNumOfUes

`func (o *NsiLoadLevelInfo) GetNumOfUes() NumberAverage`

GetNumOfUes returns the NumOfUes field if non-nil, zero value otherwise.

### GetNumOfUesOk

`func (o *NsiLoadLevelInfo) GetNumOfUesOk() (*NumberAverage, bool)`

GetNumOfUesOk returns a tuple with the NumOfUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfUes

`func (o *NsiLoadLevelInfo) SetNumOfUes(v NumberAverage)`

SetNumOfUes sets NumOfUes field to given value.

### HasNumOfUes

`func (o *NsiLoadLevelInfo) HasNumOfUes() bool`

HasNumOfUes returns a boolean if a field has been set.

### GetNumOfPduSess

`func (o *NsiLoadLevelInfo) GetNumOfPduSess() NumberAverage`

GetNumOfPduSess returns the NumOfPduSess field if non-nil, zero value otherwise.

### GetNumOfPduSessOk

`func (o *NsiLoadLevelInfo) GetNumOfPduSessOk() (*NumberAverage, bool)`

GetNumOfPduSessOk returns a tuple with the NumOfPduSess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfPduSess

`func (o *NsiLoadLevelInfo) SetNumOfPduSess(v NumberAverage)`

SetNumOfPduSess sets NumOfPduSess field to given value.

### HasNumOfPduSess

`func (o *NsiLoadLevelInfo) HasNumOfPduSess() bool`

HasNumOfPduSess returns a boolean if a field has been set.

### GetConfidence

`func (o *NsiLoadLevelInfo) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *NsiLoadLevelInfo) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *NsiLoadLevelInfo) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *NsiLoadLevelInfo) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


