# AddrFqdn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**Fqdn** | Pointer to **string** | Indicates an FQDN. | [optional] 

## Methods

### NewAddrFqdn

`func NewAddrFqdn() *AddrFqdn`

NewAddrFqdn instantiates a new AddrFqdn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddrFqdnWithDefaults

`func NewAddrFqdnWithDefaults() *AddrFqdn`

NewAddrFqdnWithDefaults instantiates a new AddrFqdn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddr

`func (o *AddrFqdn) GetIpAddr() IpAddr`

GetIpAddr returns the IpAddr field if non-nil, zero value otherwise.

### GetIpAddrOk

`func (o *AddrFqdn) GetIpAddrOk() (*IpAddr, bool)`

GetIpAddrOk returns a tuple with the IpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddr

`func (o *AddrFqdn) SetIpAddr(v IpAddr)`

SetIpAddr sets IpAddr field to given value.

### HasIpAddr

`func (o *AddrFqdn) HasIpAddr() bool`

HasIpAddr returns a boolean if a field has been set.

### GetFqdn

`func (o *AddrFqdn) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *AddrFqdn) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *AddrFqdn) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *AddrFqdn) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


