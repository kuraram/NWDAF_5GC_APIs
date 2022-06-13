# GNbId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BitLength** | **int32** | Unsigned integer representing the bit length of the gNB ID as defined in clause 9.3.1.6 of 3GPP TS 38.413 [11], within the range 22 to 32. | 
**GNBValue** | **string** | This represents the identifier of the gNB. The value of the gNB ID shall be encoded in hexadecimal representation. Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The padding 0 shall be added to make multiple nibbles,  the most significant character representing the padding 0 if required together with the 4 most significant bits of the gNB ID  shall appear first in the string, and the character representing the 4 least significant bit of the gNB ID shall appear last in the string.  | 

## Methods

### NewGNbId

`func NewGNbId(bitLength int32, gNBValue string, ) *GNbId`

NewGNbId instantiates a new GNbId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGNbIdWithDefaults

`func NewGNbIdWithDefaults() *GNbId`

NewGNbIdWithDefaults instantiates a new GNbId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitLength

`func (o *GNbId) GetBitLength() int32`

GetBitLength returns the BitLength field if non-nil, zero value otherwise.

### GetBitLengthOk

`func (o *GNbId) GetBitLengthOk() (*int32, bool)`

GetBitLengthOk returns a tuple with the BitLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitLength

`func (o *GNbId) SetBitLength(v int32)`

SetBitLength sets BitLength field to given value.


### GetGNBValue

`func (o *GNbId) GetGNBValue() string`

GetGNBValue returns the GNBValue field if non-nil, zero value otherwise.

### GetGNBValueOk

`func (o *GNbId) GetGNBValueOk() (*string, bool)`

GetGNBValueOk returns a tuple with the GNBValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGNBValue

`func (o *GNbId) SetGNBValue(v string)`

SetGNBValue sets GNBValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


