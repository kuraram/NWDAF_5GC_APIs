# WlanPerformanceReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsIds** | Pointer to **[]string** |  | [optional] 
**BssIds** | Pointer to **[]string** |  | [optional] 
**WlanOrderCriter** | Pointer to [**WlanOrderingCriterion**](WlanOrderingCriterion.md) |  | [optional] 
**Order** | Pointer to [**MatchingDirection**](MatchingDirection.md) |  | [optional] 

## Methods

### NewWlanPerformanceReq

`func NewWlanPerformanceReq() *WlanPerformanceReq`

NewWlanPerformanceReq instantiates a new WlanPerformanceReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanPerformanceReqWithDefaults

`func NewWlanPerformanceReqWithDefaults() *WlanPerformanceReq`

NewWlanPerformanceReqWithDefaults instantiates a new WlanPerformanceReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsIds

`func (o *WlanPerformanceReq) GetSsIds() []string`

GetSsIds returns the SsIds field if non-nil, zero value otherwise.

### GetSsIdsOk

`func (o *WlanPerformanceReq) GetSsIdsOk() (*[]string, bool)`

GetSsIdsOk returns a tuple with the SsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsIds

`func (o *WlanPerformanceReq) SetSsIds(v []string)`

SetSsIds sets SsIds field to given value.

### HasSsIds

`func (o *WlanPerformanceReq) HasSsIds() bool`

HasSsIds returns a boolean if a field has been set.

### GetBssIds

`func (o *WlanPerformanceReq) GetBssIds() []string`

GetBssIds returns the BssIds field if non-nil, zero value otherwise.

### GetBssIdsOk

`func (o *WlanPerformanceReq) GetBssIdsOk() (*[]string, bool)`

GetBssIdsOk returns a tuple with the BssIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssIds

`func (o *WlanPerformanceReq) SetBssIds(v []string)`

SetBssIds sets BssIds field to given value.

### HasBssIds

`func (o *WlanPerformanceReq) HasBssIds() bool`

HasBssIds returns a boolean if a field has been set.

### GetWlanOrderCriter

`func (o *WlanPerformanceReq) GetWlanOrderCriter() WlanOrderingCriterion`

GetWlanOrderCriter returns the WlanOrderCriter field if non-nil, zero value otherwise.

### GetWlanOrderCriterOk

`func (o *WlanPerformanceReq) GetWlanOrderCriterOk() (*WlanOrderingCriterion, bool)`

GetWlanOrderCriterOk returns a tuple with the WlanOrderCriter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanOrderCriter

`func (o *WlanPerformanceReq) SetWlanOrderCriter(v WlanOrderingCriterion)`

SetWlanOrderCriter sets WlanOrderCriter field to given value.

### HasWlanOrderCriter

`func (o *WlanPerformanceReq) HasWlanOrderCriter() bool`

HasWlanOrderCriter returns a boolean if a field has been set.

### GetOrder

`func (o *WlanPerformanceReq) GetOrder() MatchingDirection`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WlanPerformanceReq) GetOrderOk() (*MatchingDirection, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WlanPerformanceReq) SetOrder(v MatchingDirection)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WlanPerformanceReq) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


