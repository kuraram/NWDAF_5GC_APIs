# AnalyticsMetadataIndication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataWindow** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**DataStatProps** | Pointer to [**[]DatasetStatisticalProperty**](DatasetStatisticalProperty.md) |  | [optional] 
**Strategy** | Pointer to [**OutputStrategy**](OutputStrategy.md) |  | [optional] 
**AggrNwdafIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAnalyticsMetadataIndication

`func NewAnalyticsMetadataIndication() *AnalyticsMetadataIndication`

NewAnalyticsMetadataIndication instantiates a new AnalyticsMetadataIndication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsMetadataIndicationWithDefaults

`func NewAnalyticsMetadataIndicationWithDefaults() *AnalyticsMetadataIndication`

NewAnalyticsMetadataIndicationWithDefaults instantiates a new AnalyticsMetadataIndication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataWindow

`func (o *AnalyticsMetadataIndication) GetDataWindow() TimeWindow`

GetDataWindow returns the DataWindow field if non-nil, zero value otherwise.

### GetDataWindowOk

`func (o *AnalyticsMetadataIndication) GetDataWindowOk() (*TimeWindow, bool)`

GetDataWindowOk returns a tuple with the DataWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWindow

`func (o *AnalyticsMetadataIndication) SetDataWindow(v TimeWindow)`

SetDataWindow sets DataWindow field to given value.

### HasDataWindow

`func (o *AnalyticsMetadataIndication) HasDataWindow() bool`

HasDataWindow returns a boolean if a field has been set.

### GetDataStatProps

`func (o *AnalyticsMetadataIndication) GetDataStatProps() []DatasetStatisticalProperty`

GetDataStatProps returns the DataStatProps field if non-nil, zero value otherwise.

### GetDataStatPropsOk

`func (o *AnalyticsMetadataIndication) GetDataStatPropsOk() (*[]DatasetStatisticalProperty, bool)`

GetDataStatPropsOk returns a tuple with the DataStatProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStatProps

`func (o *AnalyticsMetadataIndication) SetDataStatProps(v []DatasetStatisticalProperty)`

SetDataStatProps sets DataStatProps field to given value.

### HasDataStatProps

`func (o *AnalyticsMetadataIndication) HasDataStatProps() bool`

HasDataStatProps returns a boolean if a field has been set.

### GetStrategy

`func (o *AnalyticsMetadataIndication) GetStrategy() OutputStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *AnalyticsMetadataIndication) GetStrategyOk() (*OutputStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *AnalyticsMetadataIndication) SetStrategy(v OutputStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *AnalyticsMetadataIndication) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetAggrNwdafIds

`func (o *AnalyticsMetadataIndication) GetAggrNwdafIds() []string`

GetAggrNwdafIds returns the AggrNwdafIds field if non-nil, zero value otherwise.

### GetAggrNwdafIdsOk

`func (o *AnalyticsMetadataIndication) GetAggrNwdafIdsOk() (*[]string, bool)`

GetAggrNwdafIdsOk returns a tuple with the AggrNwdafIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggrNwdafIds

`func (o *AnalyticsMetadataIndication) SetAggrNwdafIds(v []string)`

SetAggrNwdafIds sets AggrNwdafIds field to given value.

### HasAggrNwdafIds

`func (o *AnalyticsMetadataIndication) HasAggrNwdafIds() bool`

HasAggrNwdafIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


