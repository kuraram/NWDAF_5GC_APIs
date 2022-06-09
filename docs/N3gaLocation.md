# N3gaLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**N3gppTai** | Pointer to [**Tai**](Tai.md) |  | [optional] 
**N3IwfId** | Pointer to **string** |  | [optional] 
**UeIpv4Addr** | Pointer to **string** |  | [optional] 
**UeIpv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**PortNumber** | Pointer to **int32** |  | [optional] 
**TnapId** | Pointer to [**TnapId**](TnapId.md) |  | [optional] 
**Protocol** | Pointer to [**TransportProtocol**](TransportProtocol.md) |  | [optional] 
**TwapId** | Pointer to [**TwapId**](TwapId.md) |  | [optional] 
**HfcNodeId** | Pointer to [**HfcNodeId**](HfcNodeId.md) |  | [optional] 
**Gli** | Pointer to **string** |  | [optional] 
**W5gbanLineType** | Pointer to [**LineType**](LineType.md) |  | [optional] 
**Gci** | Pointer to **string** |  | [optional] 

## Methods

### NewN3gaLocation

`func NewN3gaLocation() *N3gaLocation`

NewN3gaLocation instantiates a new N3gaLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN3gaLocationWithDefaults

`func NewN3gaLocationWithDefaults() *N3gaLocation`

NewN3gaLocationWithDefaults instantiates a new N3gaLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetN3gppTai

`func (o *N3gaLocation) GetN3gppTai() Tai`

GetN3gppTai returns the N3gppTai field if non-nil, zero value otherwise.

### GetN3gppTaiOk

`func (o *N3gaLocation) GetN3gppTaiOk() (*Tai, bool)`

GetN3gppTaiOk returns a tuple with the N3gppTai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3gppTai

`func (o *N3gaLocation) SetN3gppTai(v Tai)`

SetN3gppTai sets N3gppTai field to given value.

### HasN3gppTai

`func (o *N3gaLocation) HasN3gppTai() bool`

HasN3gppTai returns a boolean if a field has been set.

### GetN3IwfId

`func (o *N3gaLocation) GetN3IwfId() string`

GetN3IwfId returns the N3IwfId field if non-nil, zero value otherwise.

### GetN3IwfIdOk

`func (o *N3gaLocation) GetN3IwfIdOk() (*string, bool)`

GetN3IwfIdOk returns a tuple with the N3IwfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3IwfId

`func (o *N3gaLocation) SetN3IwfId(v string)`

SetN3IwfId sets N3IwfId field to given value.

### HasN3IwfId

`func (o *N3gaLocation) HasN3IwfId() bool`

HasN3IwfId returns a boolean if a field has been set.

### GetUeIpv4Addr

`func (o *N3gaLocation) GetUeIpv4Addr() string`

GetUeIpv4Addr returns the UeIpv4Addr field if non-nil, zero value otherwise.

### GetUeIpv4AddrOk

`func (o *N3gaLocation) GetUeIpv4AddrOk() (*string, bool)`

GetUeIpv4AddrOk returns a tuple with the UeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4Addr

`func (o *N3gaLocation) SetUeIpv4Addr(v string)`

SetUeIpv4Addr sets UeIpv4Addr field to given value.

### HasUeIpv4Addr

`func (o *N3gaLocation) HasUeIpv4Addr() bool`

HasUeIpv4Addr returns a boolean if a field has been set.

### GetUeIpv6Addr

`func (o *N3gaLocation) GetUeIpv6Addr() Ipv6Addr`

GetUeIpv6Addr returns the UeIpv6Addr field if non-nil, zero value otherwise.

### GetUeIpv6AddrOk

`func (o *N3gaLocation) GetUeIpv6AddrOk() (*Ipv6Addr, bool)`

GetUeIpv6AddrOk returns a tuple with the UeIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6Addr

`func (o *N3gaLocation) SetUeIpv6Addr(v Ipv6Addr)`

SetUeIpv6Addr sets UeIpv6Addr field to given value.

### HasUeIpv6Addr

`func (o *N3gaLocation) HasUeIpv6Addr() bool`

HasUeIpv6Addr returns a boolean if a field has been set.

### GetPortNumber

`func (o *N3gaLocation) GetPortNumber() int32`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *N3gaLocation) GetPortNumberOk() (*int32, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *N3gaLocation) SetPortNumber(v int32)`

SetPortNumber sets PortNumber field to given value.

### HasPortNumber

`func (o *N3gaLocation) HasPortNumber() bool`

HasPortNumber returns a boolean if a field has been set.

### GetTnapId

`func (o *N3gaLocation) GetTnapId() TnapId`

GetTnapId returns the TnapId field if non-nil, zero value otherwise.

### GetTnapIdOk

`func (o *N3gaLocation) GetTnapIdOk() (*TnapId, bool)`

GetTnapIdOk returns a tuple with the TnapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTnapId

`func (o *N3gaLocation) SetTnapId(v TnapId)`

SetTnapId sets TnapId field to given value.

### HasTnapId

`func (o *N3gaLocation) HasTnapId() bool`

HasTnapId returns a boolean if a field has been set.

### GetProtocol

`func (o *N3gaLocation) GetProtocol() TransportProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *N3gaLocation) GetProtocolOk() (*TransportProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *N3gaLocation) SetProtocol(v TransportProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *N3gaLocation) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTwapId

`func (o *N3gaLocation) GetTwapId() TwapId`

GetTwapId returns the TwapId field if non-nil, zero value otherwise.

### GetTwapIdOk

`func (o *N3gaLocation) GetTwapIdOk() (*TwapId, bool)`

GetTwapIdOk returns a tuple with the TwapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwapId

`func (o *N3gaLocation) SetTwapId(v TwapId)`

SetTwapId sets TwapId field to given value.

### HasTwapId

`func (o *N3gaLocation) HasTwapId() bool`

HasTwapId returns a boolean if a field has been set.

### GetHfcNodeId

`func (o *N3gaLocation) GetHfcNodeId() HfcNodeId`

GetHfcNodeId returns the HfcNodeId field if non-nil, zero value otherwise.

### GetHfcNodeIdOk

`func (o *N3gaLocation) GetHfcNodeIdOk() (*HfcNodeId, bool)`

GetHfcNodeIdOk returns a tuple with the HfcNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHfcNodeId

`func (o *N3gaLocation) SetHfcNodeId(v HfcNodeId)`

SetHfcNodeId sets HfcNodeId field to given value.

### HasHfcNodeId

`func (o *N3gaLocation) HasHfcNodeId() bool`

HasHfcNodeId returns a boolean if a field has been set.

### GetGli

`func (o *N3gaLocation) GetGli() string`

GetGli returns the Gli field if non-nil, zero value otherwise.

### GetGliOk

`func (o *N3gaLocation) GetGliOk() (*string, bool)`

GetGliOk returns a tuple with the Gli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGli

`func (o *N3gaLocation) SetGli(v string)`

SetGli sets Gli field to given value.

### HasGli

`func (o *N3gaLocation) HasGli() bool`

HasGli returns a boolean if a field has been set.

### GetW5gbanLineType

`func (o *N3gaLocation) GetW5gbanLineType() LineType`

GetW5gbanLineType returns the W5gbanLineType field if non-nil, zero value otherwise.

### GetW5gbanLineTypeOk

`func (o *N3gaLocation) GetW5gbanLineTypeOk() (*LineType, bool)`

GetW5gbanLineTypeOk returns a tuple with the W5gbanLineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW5gbanLineType

`func (o *N3gaLocation) SetW5gbanLineType(v LineType)`

SetW5gbanLineType sets W5gbanLineType field to given value.

### HasW5gbanLineType

`func (o *N3gaLocation) HasW5gbanLineType() bool`

HasW5gbanLineType returns a boolean if a field has been set.

### GetGci

`func (o *N3gaLocation) GetGci() string`

GetGci returns the Gci field if non-nil, zero value otherwise.

### GetGciOk

`func (o *N3gaLocation) GetGciOk() (*string, bool)`

GetGciOk returns a tuple with the Gci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGci

`func (o *N3gaLocation) SetGci(v string)`

SetGci sets Gci field to given value.

### HasGci

`func (o *N3gaLocation) HasGci() bool`

HasGci returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


