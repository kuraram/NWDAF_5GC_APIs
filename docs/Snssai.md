# Snssai

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sst** | **int32** | Unsigned integer, within the range 0 to 255, representing the Slice/Service Type.  It indicates the expected Network Slice behaviour in terms of features and services. Values 0 to 127 correspond to the standardized SST range. Values 128 to 255 correspond  to the Operator-specific range. See clause 28.4.2 of 3GPP TS 23.003. Standardized values are defined in clause 5.15.2.2 of 3GPP TS 23.501.   | 
**Sd** | Pointer to **string** | 3-octet string, representing the Slice Differentiator, in hexadecimal representation. Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits of the SD shall appear first in the string, and the character representing the 4 least significant bit of the SD shall appear last in the string. This is an optional parameter that complements the Slice/Service type(s) to allow to differentiate amongst multiple Network Slices of the same Slice/Service type. This IE shall be absent if no SD value is associated with the SST.  | [optional] 

## Methods

### NewSnssai

`func NewSnssai(sst int32, ) *Snssai`

NewSnssai instantiates a new Snssai object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnssaiWithDefaults

`func NewSnssaiWithDefaults() *Snssai`

NewSnssaiWithDefaults instantiates a new Snssai object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSst

`func (o *Snssai) GetSst() int32`

GetSst returns the Sst field if non-nil, zero value otherwise.

### GetSstOk

`func (o *Snssai) GetSstOk() (*int32, bool)`

GetSstOk returns a tuple with the Sst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSst

`func (o *Snssai) SetSst(v int32)`

SetSst sets Sst field to given value.


### GetSd

`func (o *Snssai) GetSd() string`

GetSd returns the Sd field if non-nil, zero value otherwise.

### GetSdOk

`func (o *Snssai) GetSdOk() (*string, bool)`

GetSdOk returns a tuple with the Sd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSd

`func (o *Snssai) SetSd(v string)`

SetSd sets Sd field to given value.

### HasSd

`func (o *Snssai) HasSd() bool`

HasSd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


