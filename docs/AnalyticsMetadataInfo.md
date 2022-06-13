# AnalyticsMetadataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumSamples** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**DataWindow** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**DataStatProps** | Pointer to [**[]DatasetStatisticalProperty**](DatasetStatisticalProperty.md) |  | [optional] 
**Strategy** | Pointer to [**OutputStrategy**](OutputStrategy.md) |  | [optional] 
**Accuracy** | Pointer to [**Accuracy**](Accuracy.md) |  | [optional] 

## Methods

### NewAnalyticsMetadataInfo

`func NewAnalyticsMetadataInfo() *AnalyticsMetadataInfo`

NewAnalyticsMetadataInfo instantiates a new AnalyticsMetadataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsMetadataInfoWithDefaults

`func NewAnalyticsMetadataInfoWithDefaults() *AnalyticsMetadataInfo`

NewAnalyticsMetadataInfoWithDefaults instantiates a new AnalyticsMetadataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumSamples

`func (o *AnalyticsMetadataInfo) GetNumSamples() int32`

GetNumSamples returns the NumSamples field if non-nil, zero value otherwise.

### GetNumSamplesOk

`func (o *AnalyticsMetadataInfo) GetNumSamplesOk() (*int32, bool)`

GetNumSamplesOk returns a tuple with the NumSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSamples

`func (o *AnalyticsMetadataInfo) SetNumSamples(v int32)`

SetNumSamples sets NumSamples field to given value.

### HasNumSamples

`func (o *AnalyticsMetadataInfo) HasNumSamples() bool`

HasNumSamples returns a boolean if a field has been set.

### GetDataWindow

`func (o *AnalyticsMetadataInfo) GetDataWindow() TimeWindow`

GetDataWindow returns the DataWindow field if non-nil, zero value otherwise.

### GetDataWindowOk

`func (o *AnalyticsMetadataInfo) GetDataWindowOk() (*TimeWindow, bool)`

GetDataWindowOk returns a tuple with the DataWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWindow

`func (o *AnalyticsMetadataInfo) SetDataWindow(v TimeWindow)`

SetDataWindow sets DataWindow field to given value.

### HasDataWindow

`func (o *AnalyticsMetadataInfo) HasDataWindow() bool`

HasDataWindow returns a boolean if a field has been set.

### GetDataStatProps

`func (o *AnalyticsMetadataInfo) GetDataStatProps() []DatasetStatisticalProperty`

GetDataStatProps returns the DataStatProps field if non-nil, zero value otherwise.

### GetDataStatPropsOk

`func (o *AnalyticsMetadataInfo) GetDataStatPropsOk() (*[]DatasetStatisticalProperty, bool)`

GetDataStatPropsOk returns a tuple with the DataStatProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStatProps

`func (o *AnalyticsMetadataInfo) SetDataStatProps(v []DatasetStatisticalProperty)`

SetDataStatProps sets DataStatProps field to given value.

### HasDataStatProps

`func (o *AnalyticsMetadataInfo) HasDataStatProps() bool`

HasDataStatProps returns a boolean if a field has been set.

### GetStrategy

`func (o *AnalyticsMetadataInfo) GetStrategy() OutputStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *AnalyticsMetadataInfo) GetStrategyOk() (*OutputStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *AnalyticsMetadataInfo) SetStrategy(v OutputStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *AnalyticsMetadataInfo) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetAccuracy

`func (o *AnalyticsMetadataInfo) GetAccuracy() Accuracy`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *AnalyticsMetadataInfo) GetAccuracyOk() (*Accuracy, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *AnalyticsMetadataInfo) SetAccuracy(v Accuracy)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *AnalyticsMetadataInfo) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


