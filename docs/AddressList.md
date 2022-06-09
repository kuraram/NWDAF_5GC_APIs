# AddressList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Addrs** | Pointer to **[]string** |  | [optional] 
**Ipv6Addrs** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 

## Methods

### NewAddressList

`func NewAddressList() *AddressList`

NewAddressList instantiates a new AddressList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressListWithDefaults

`func NewAddressListWithDefaults() *AddressList`

NewAddressListWithDefaults instantiates a new AddressList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Addrs

`func (o *AddressList) GetIpv4Addrs() []string`

GetIpv4Addrs returns the Ipv4Addrs field if non-nil, zero value otherwise.

### GetIpv4AddrsOk

`func (o *AddressList) GetIpv4AddrsOk() (*[]string, bool)`

GetIpv4AddrsOk returns a tuple with the Ipv4Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addrs

`func (o *AddressList) SetIpv4Addrs(v []string)`

SetIpv4Addrs sets Ipv4Addrs field to given value.

### HasIpv4Addrs

`func (o *AddressList) HasIpv4Addrs() bool`

HasIpv4Addrs returns a boolean if a field has been set.

### GetIpv6Addrs

`func (o *AddressList) GetIpv6Addrs() []Ipv6Addr`

GetIpv6Addrs returns the Ipv6Addrs field if non-nil, zero value otherwise.

### GetIpv6AddrsOk

`func (o *AddressList) GetIpv6AddrsOk() (*[]Ipv6Addr, bool)`

GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addrs

`func (o *AddressList) SetIpv6Addrs(v []Ipv6Addr)`

SetIpv6Addrs sets Ipv6Addrs field to given value.

### HasIpv6Addrs

`func (o *AddressList) HasIpv6Addrs() bool`

HasIpv6Addrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


