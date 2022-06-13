# RetainabilityThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelFlowNum** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**RelTimeUnit** | Pointer to [**TimeUnit**](TimeUnit.md) |  | [optional] 
**RelFlowRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | [optional] 

## Methods

### NewRetainabilityThreshold

`func NewRetainabilityThreshold() *RetainabilityThreshold`

NewRetainabilityThreshold instantiates a new RetainabilityThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetainabilityThresholdWithDefaults

`func NewRetainabilityThresholdWithDefaults() *RetainabilityThreshold`

NewRetainabilityThresholdWithDefaults instantiates a new RetainabilityThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelFlowNum

`func (o *RetainabilityThreshold) GetRelFlowNum() int32`

GetRelFlowNum returns the RelFlowNum field if non-nil, zero value otherwise.

### GetRelFlowNumOk

`func (o *RetainabilityThreshold) GetRelFlowNumOk() (*int32, bool)`

GetRelFlowNumOk returns a tuple with the RelFlowNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelFlowNum

`func (o *RetainabilityThreshold) SetRelFlowNum(v int32)`

SetRelFlowNum sets RelFlowNum field to given value.

### HasRelFlowNum

`func (o *RetainabilityThreshold) HasRelFlowNum() bool`

HasRelFlowNum returns a boolean if a field has been set.

### GetRelTimeUnit

`func (o *RetainabilityThreshold) GetRelTimeUnit() TimeUnit`

GetRelTimeUnit returns the RelTimeUnit field if non-nil, zero value otherwise.

### GetRelTimeUnitOk

`func (o *RetainabilityThreshold) GetRelTimeUnitOk() (*TimeUnit, bool)`

GetRelTimeUnitOk returns a tuple with the RelTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelTimeUnit

`func (o *RetainabilityThreshold) SetRelTimeUnit(v TimeUnit)`

SetRelTimeUnit sets RelTimeUnit field to given value.

### HasRelTimeUnit

`func (o *RetainabilityThreshold) HasRelTimeUnit() bool`

HasRelTimeUnit returns a boolean if a field has been set.

### GetRelFlowRatio

`func (o *RetainabilityThreshold) GetRelFlowRatio() int32`

GetRelFlowRatio returns the RelFlowRatio field if non-nil, zero value otherwise.

### GetRelFlowRatioOk

`func (o *RetainabilityThreshold) GetRelFlowRatioOk() (*int32, bool)`

GetRelFlowRatioOk returns a tuple with the RelFlowRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelFlowRatio

`func (o *RetainabilityThreshold) SetRelFlowRatio(v int32)`

SetRelFlowRatio sets RelFlowRatio field to given value.

### HasRelFlowRatio

`func (o *RetainabilityThreshold) HasRelFlowRatio() bool`

HasRelFlowRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


