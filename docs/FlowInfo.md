# FlowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **int32** | Indicates the IP flow. | 
**FlowDescriptions** | Pointer to **[]string** | Indicates the packet filters of the IP flow. Refer to subclause 5.3.8 of 3GPP TS 29.214 for encoding. It shall contain UL and/or DL IP flow description. | [optional] 

## Methods

### NewFlowInfo

`func NewFlowInfo(flowId int32, ) *FlowInfo`

NewFlowInfo instantiates a new FlowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowInfoWithDefaults

`func NewFlowInfoWithDefaults() *FlowInfo`

NewFlowInfoWithDefaults instantiates a new FlowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *FlowInfo) GetFlowId() int32`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *FlowInfo) GetFlowIdOk() (*int32, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *FlowInfo) SetFlowId(v int32)`

SetFlowId sets FlowId field to given value.


### GetFlowDescriptions

`func (o *FlowInfo) GetFlowDescriptions() []string`

GetFlowDescriptions returns the FlowDescriptions field if non-nil, zero value otherwise.

### GetFlowDescriptionsOk

`func (o *FlowInfo) GetFlowDescriptionsOk() (*[]string, bool)`

GetFlowDescriptionsOk returns a tuple with the FlowDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDescriptions

`func (o *FlowInfo) SetFlowDescriptions(v []string)`

SetFlowDescriptions sets FlowDescriptions field to given value.

### HasFlowDescriptions

`func (o *FlowInfo) HasFlowDescriptions() bool`

HasFlowDescriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


