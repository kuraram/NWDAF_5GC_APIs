# DnPerformanceReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnPerfOrderCriter** | Pointer to [**DnPerfOrderingCriterion**](DnPerfOrderingCriterion.md) |  | [optional] 
**Order** | Pointer to [**MatchingDirection**](MatchingDirection.md) |  | [optional] 
**ReportThresholds** | Pointer to [**[]ThresholdLevel**](ThresholdLevel.md) |  | [optional] 

## Methods

### NewDnPerformanceReq

`func NewDnPerformanceReq() *DnPerformanceReq`

NewDnPerformanceReq instantiates a new DnPerformanceReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnPerformanceReqWithDefaults

`func NewDnPerformanceReqWithDefaults() *DnPerformanceReq`

NewDnPerformanceReqWithDefaults instantiates a new DnPerformanceReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnPerfOrderCriter

`func (o *DnPerformanceReq) GetDnPerfOrderCriter() DnPerfOrderingCriterion`

GetDnPerfOrderCriter returns the DnPerfOrderCriter field if non-nil, zero value otherwise.

### GetDnPerfOrderCriterOk

`func (o *DnPerformanceReq) GetDnPerfOrderCriterOk() (*DnPerfOrderingCriterion, bool)`

GetDnPerfOrderCriterOk returns a tuple with the DnPerfOrderCriter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPerfOrderCriter

`func (o *DnPerformanceReq) SetDnPerfOrderCriter(v DnPerfOrderingCriterion)`

SetDnPerfOrderCriter sets DnPerfOrderCriter field to given value.

### HasDnPerfOrderCriter

`func (o *DnPerformanceReq) HasDnPerfOrderCriter() bool`

HasDnPerfOrderCriter returns a boolean if a field has been set.

### GetOrder

`func (o *DnPerformanceReq) GetOrder() MatchingDirection`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DnPerformanceReq) GetOrderOk() (*MatchingDirection, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DnPerformanceReq) SetOrder(v MatchingDirection)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *DnPerformanceReq) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReportThresholds

`func (o *DnPerformanceReq) GetReportThresholds() []ThresholdLevel`

GetReportThresholds returns the ReportThresholds field if non-nil, zero value otherwise.

### GetReportThresholdsOk

`func (o *DnPerformanceReq) GetReportThresholdsOk() (*[]ThresholdLevel, bool)`

GetReportThresholdsOk returns a tuple with the ReportThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportThresholds

`func (o *DnPerformanceReq) SetReportThresholds(v []ThresholdLevel)`

SetReportThresholds sets ReportThresholds field to given value.

### HasReportThresholds

`func (o *DnPerformanceReq) HasReportThresholds() bool`

HasReportThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


