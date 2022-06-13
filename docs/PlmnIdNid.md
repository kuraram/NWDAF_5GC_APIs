# PlmnIdNid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mcc** | **string** | Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413. | 
**Mnc** | **string** | Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413. | 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1). | [optional] 

## Methods

### NewPlmnIdNid

`func NewPlmnIdNid(mcc string, mnc string, ) *PlmnIdNid`

NewPlmnIdNid instantiates a new PlmnIdNid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlmnIdNidWithDefaults

`func NewPlmnIdNidWithDefaults() *PlmnIdNid`

NewPlmnIdNidWithDefaults instantiates a new PlmnIdNid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcc

`func (o *PlmnIdNid) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *PlmnIdNid) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *PlmnIdNid) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetMnc

`func (o *PlmnIdNid) GetMnc() string`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *PlmnIdNid) GetMncOk() (*string, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *PlmnIdNid) SetMnc(v string)`

SetMnc sets Mnc field to given value.


### GetNid

`func (o *PlmnIdNid) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *PlmnIdNid) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *PlmnIdNid) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *PlmnIdNid) HasNid() bool`

HasNid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


