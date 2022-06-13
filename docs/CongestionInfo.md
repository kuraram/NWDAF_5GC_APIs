# CongestionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CongType** | [**CongestionType**](CongestionType.md) |  | 
**TimeIntev** | [**TimeWindow**](TimeWindow.md) |  | 
**Nsi** | [**ThresholdLevel**](ThresholdLevel.md) |  | 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TopAppListUl** | Pointer to [**[]TopApplication**](TopApplication.md) |  | [optional] 
**TopAppListDl** | Pointer to [**[]TopApplication**](TopApplication.md) |  | [optional] 

## Methods

### NewCongestionInfo

`func NewCongestionInfo(congType CongestionType, timeIntev TimeWindow, nsi ThresholdLevel, ) *CongestionInfo`

NewCongestionInfo instantiates a new CongestionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCongestionInfoWithDefaults

`func NewCongestionInfoWithDefaults() *CongestionInfo`

NewCongestionInfoWithDefaults instantiates a new CongestionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCongType

`func (o *CongestionInfo) GetCongType() CongestionType`

GetCongType returns the CongType field if non-nil, zero value otherwise.

### GetCongTypeOk

`func (o *CongestionInfo) GetCongTypeOk() (*CongestionType, bool)`

GetCongTypeOk returns a tuple with the CongType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongType

`func (o *CongestionInfo) SetCongType(v CongestionType)`

SetCongType sets CongType field to given value.


### GetTimeIntev

`func (o *CongestionInfo) GetTimeIntev() TimeWindow`

GetTimeIntev returns the TimeIntev field if non-nil, zero value otherwise.

### GetTimeIntevOk

`func (o *CongestionInfo) GetTimeIntevOk() (*TimeWindow, bool)`

GetTimeIntevOk returns a tuple with the TimeIntev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntev

`func (o *CongestionInfo) SetTimeIntev(v TimeWindow)`

SetTimeIntev sets TimeIntev field to given value.


### GetNsi

`func (o *CongestionInfo) GetNsi() ThresholdLevel`

GetNsi returns the Nsi field if non-nil, zero value otherwise.

### GetNsiOk

`func (o *CongestionInfo) GetNsiOk() (*ThresholdLevel, bool)`

GetNsiOk returns a tuple with the Nsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsi

`func (o *CongestionInfo) SetNsi(v ThresholdLevel)`

SetNsi sets Nsi field to given value.


### GetConfidence

`func (o *CongestionInfo) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CongestionInfo) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CongestionInfo) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *CongestionInfo) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetTopAppListUl

`func (o *CongestionInfo) GetTopAppListUl() []TopApplication`

GetTopAppListUl returns the TopAppListUl field if non-nil, zero value otherwise.

### GetTopAppListUlOk

`func (o *CongestionInfo) GetTopAppListUlOk() (*[]TopApplication, bool)`

GetTopAppListUlOk returns a tuple with the TopAppListUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopAppListUl

`func (o *CongestionInfo) SetTopAppListUl(v []TopApplication)`

SetTopAppListUl sets TopAppListUl field to given value.

### HasTopAppListUl

`func (o *CongestionInfo) HasTopAppListUl() bool`

HasTopAppListUl returns a boolean if a field has been set.

### GetTopAppListDl

`func (o *CongestionInfo) GetTopAppListDl() []TopApplication`

GetTopAppListDl returns the TopAppListDl field if non-nil, zero value otherwise.

### GetTopAppListDlOk

`func (o *CongestionInfo) GetTopAppListDlOk() (*[]TopApplication, bool)`

GetTopAppListDlOk returns a tuple with the TopAppListDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopAppListDl

`func (o *CongestionInfo) SetTopAppListDl(v []TopApplication)`

SetTopAppListDl sets TopAppListDl field to given value.

### HasTopAppListDl

`func (o *CongestionInfo) HasTopAppListDl() bool`

HasTopAppListDl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


