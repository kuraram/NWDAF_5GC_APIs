# ClassCriterion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisperClass** | [**DispersionClass**](DispersionClass.md) |  | 
**ClassThreshold** | **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | 
**ThresMatch** | [**MatchingDirection**](MatchingDirection.md) |  | 

## Methods

### NewClassCriterion

`func NewClassCriterion(disperClass DispersionClass, classThreshold int32, thresMatch MatchingDirection, ) *ClassCriterion`

NewClassCriterion instantiates a new ClassCriterion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassCriterionWithDefaults

`func NewClassCriterionWithDefaults() *ClassCriterion`

NewClassCriterionWithDefaults instantiates a new ClassCriterion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisperClass

`func (o *ClassCriterion) GetDisperClass() DispersionClass`

GetDisperClass returns the DisperClass field if non-nil, zero value otherwise.

### GetDisperClassOk

`func (o *ClassCriterion) GetDisperClassOk() (*DispersionClass, bool)`

GetDisperClassOk returns a tuple with the DisperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperClass

`func (o *ClassCriterion) SetDisperClass(v DispersionClass)`

SetDisperClass sets DisperClass field to given value.


### GetClassThreshold

`func (o *ClassCriterion) GetClassThreshold() int32`

GetClassThreshold returns the ClassThreshold field if non-nil, zero value otherwise.

### GetClassThresholdOk

`func (o *ClassCriterion) GetClassThresholdOk() (*int32, bool)`

GetClassThresholdOk returns a tuple with the ClassThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassThreshold

`func (o *ClassCriterion) SetClassThreshold(v int32)`

SetClassThreshold sets ClassThreshold field to given value.


### GetThresMatch

`func (o *ClassCriterion) GetThresMatch() MatchingDirection`

GetThresMatch returns the ThresMatch field if non-nil, zero value otherwise.

### GetThresMatchOk

`func (o *ClassCriterion) GetThresMatchOk() (*MatchingDirection, bool)`

GetThresMatchOk returns a tuple with the ThresMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresMatch

`func (o *ClassCriterion) SetThresMatch(v MatchingDirection)`

SetThresMatch sets ThresMatch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


