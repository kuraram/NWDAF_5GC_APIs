# IpEthFlowDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpTrafficFilter** | Pointer to **string** | Defines a packet filter of an IP flow. | [optional] 
**EthTrafficFilter** | Pointer to [**EthFlowDescription**](EthFlowDescription.md) |  | [optional] 

## Methods

### NewIpEthFlowDescription

`func NewIpEthFlowDescription() *IpEthFlowDescription`

NewIpEthFlowDescription instantiates a new IpEthFlowDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpEthFlowDescriptionWithDefaults

`func NewIpEthFlowDescriptionWithDefaults() *IpEthFlowDescription`

NewIpEthFlowDescriptionWithDefaults instantiates a new IpEthFlowDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpTrafficFilter

`func (o *IpEthFlowDescription) GetIpTrafficFilter() string`

GetIpTrafficFilter returns the IpTrafficFilter field if non-nil, zero value otherwise.

### GetIpTrafficFilterOk

`func (o *IpEthFlowDescription) GetIpTrafficFilterOk() (*string, bool)`

GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpTrafficFilter

`func (o *IpEthFlowDescription) SetIpTrafficFilter(v string)`

SetIpTrafficFilter sets IpTrafficFilter field to given value.

### HasIpTrafficFilter

`func (o *IpEthFlowDescription) HasIpTrafficFilter() bool`

HasIpTrafficFilter returns a boolean if a field has been set.

### GetEthTrafficFilter

`func (o *IpEthFlowDescription) GetEthTrafficFilter() EthFlowDescription`

GetEthTrafficFilter returns the EthTrafficFilter field if non-nil, zero value otherwise.

### GetEthTrafficFilterOk

`func (o *IpEthFlowDescription) GetEthTrafficFilterOk() (*EthFlowDescription, bool)`

GetEthTrafficFilterOk returns a tuple with the EthTrafficFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthTrafficFilter

`func (o *IpEthFlowDescription) SetEthTrafficFilter(v EthFlowDescription)`

SetEthTrafficFilter sets EthTrafficFilter field to given value.

### HasEthTrafficFilter

`func (o *IpEthFlowDescription) HasEthTrafficFilter() bool`

HasEthTrafficFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


