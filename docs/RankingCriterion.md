# RankingCriterion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HighBase** | **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | 
**LowBase** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | [optional] 

## Methods

### NewRankingCriterion

`func NewRankingCriterion(highBase int32, ) *RankingCriterion`

NewRankingCriterion instantiates a new RankingCriterion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRankingCriterionWithDefaults

`func NewRankingCriterionWithDefaults() *RankingCriterion`

NewRankingCriterionWithDefaults instantiates a new RankingCriterion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighBase

`func (o *RankingCriterion) GetHighBase() int32`

GetHighBase returns the HighBase field if non-nil, zero value otherwise.

### GetHighBaseOk

`func (o *RankingCriterion) GetHighBaseOk() (*int32, bool)`

GetHighBaseOk returns a tuple with the HighBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighBase

`func (o *RankingCriterion) SetHighBase(v int32)`

SetHighBase sets HighBase field to given value.


### GetLowBase

`func (o *RankingCriterion) GetLowBase() int32`

GetLowBase returns the LowBase field if non-nil, zero value otherwise.

### GetLowBaseOk

`func (o *RankingCriterion) GetLowBaseOk() (*int32, bool)`

GetLowBaseOk returns a tuple with the LowBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBase

`func (o *RankingCriterion) SetLowBase(v int32)`

SetLowBase sets LowBase field to given value.

### HasLowBase

`func (o *RankingCriterion) HasLowBase() bool`

HasLowBase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


