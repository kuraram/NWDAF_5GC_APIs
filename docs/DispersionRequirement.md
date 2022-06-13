# DispersionRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisperType** | [**DispersionType**](DispersionType.md) |  | 
**ClassCriters** | Pointer to [**[]ClassCriterion**](ClassCriterion.md) |  | [optional] 
**RankCriters** | Pointer to [**[]RankingCriterion**](RankingCriterion.md) |  | [optional] 
**DispOrderCriter** | Pointer to [**DispersionOrderingCriterion**](DispersionOrderingCriterion.md) |  | [optional] 
**Order** | Pointer to [**MatchingDirection**](MatchingDirection.md) |  | [optional] 

## Methods

### NewDispersionRequirement

`func NewDispersionRequirement(disperType DispersionType, ) *DispersionRequirement`

NewDispersionRequirement instantiates a new DispersionRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDispersionRequirementWithDefaults

`func NewDispersionRequirementWithDefaults() *DispersionRequirement`

NewDispersionRequirementWithDefaults instantiates a new DispersionRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisperType

`func (o *DispersionRequirement) GetDisperType() DispersionType`

GetDisperType returns the DisperType field if non-nil, zero value otherwise.

### GetDisperTypeOk

`func (o *DispersionRequirement) GetDisperTypeOk() (*DispersionType, bool)`

GetDisperTypeOk returns a tuple with the DisperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperType

`func (o *DispersionRequirement) SetDisperType(v DispersionType)`

SetDisperType sets DisperType field to given value.


### GetClassCriters

`func (o *DispersionRequirement) GetClassCriters() []ClassCriterion`

GetClassCriters returns the ClassCriters field if non-nil, zero value otherwise.

### GetClassCritersOk

`func (o *DispersionRequirement) GetClassCritersOk() (*[]ClassCriterion, bool)`

GetClassCritersOk returns a tuple with the ClassCriters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassCriters

`func (o *DispersionRequirement) SetClassCriters(v []ClassCriterion)`

SetClassCriters sets ClassCriters field to given value.

### HasClassCriters

`func (o *DispersionRequirement) HasClassCriters() bool`

HasClassCriters returns a boolean if a field has been set.

### GetRankCriters

`func (o *DispersionRequirement) GetRankCriters() []RankingCriterion`

GetRankCriters returns the RankCriters field if non-nil, zero value otherwise.

### GetRankCritersOk

`func (o *DispersionRequirement) GetRankCritersOk() (*[]RankingCriterion, bool)`

GetRankCritersOk returns a tuple with the RankCriters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankCriters

`func (o *DispersionRequirement) SetRankCriters(v []RankingCriterion)`

SetRankCriters sets RankCriters field to given value.

### HasRankCriters

`func (o *DispersionRequirement) HasRankCriters() bool`

HasRankCriters returns a boolean if a field has been set.

### GetDispOrderCriter

`func (o *DispersionRequirement) GetDispOrderCriter() DispersionOrderingCriterion`

GetDispOrderCriter returns the DispOrderCriter field if non-nil, zero value otherwise.

### GetDispOrderCriterOk

`func (o *DispersionRequirement) GetDispOrderCriterOk() (*DispersionOrderingCriterion, bool)`

GetDispOrderCriterOk returns a tuple with the DispOrderCriter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispOrderCriter

`func (o *DispersionRequirement) SetDispOrderCriter(v DispersionOrderingCriterion)`

SetDispOrderCriter sets DispOrderCriter field to given value.

### HasDispOrderCriter

`func (o *DispersionRequirement) HasDispOrderCriter() bool`

HasDispOrderCriter returns a boolean if a field has been set.

### GetOrder

`func (o *DispersionRequirement) GetOrder() MatchingDirection`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DispersionRequirement) GetOrderOk() (*MatchingDirection, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DispersionRequirement) SetOrder(v MatchingDirection)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *DispersionRequirement) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


