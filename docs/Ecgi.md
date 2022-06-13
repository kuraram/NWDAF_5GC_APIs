# Ecgi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**EutraCellId** | **string** | 28-bit string identifying an E-UTRA Cell Id as specified in clause 9.3.1.9 of  3GPP TS 38.413, in hexadecimal representation. Each character in the string shall take a  value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most  significant character representing the 4 most significant bits of the Cell Id shall appear  first in the string, and the character representing the 4 least significant bit of the  Cell Id shall appear last in the string.   | 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1). | [optional] 

## Methods

### NewEcgi

`func NewEcgi(plmnId PlmnId, eutraCellId string, ) *Ecgi`

NewEcgi instantiates a new Ecgi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcgiWithDefaults

`func NewEcgiWithDefaults() *Ecgi`

NewEcgiWithDefaults instantiates a new Ecgi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *Ecgi) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *Ecgi) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *Ecgi) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetEutraCellId

`func (o *Ecgi) GetEutraCellId() string`

GetEutraCellId returns the EutraCellId field if non-nil, zero value otherwise.

### GetEutraCellIdOk

`func (o *Ecgi) GetEutraCellIdOk() (*string, bool)`

GetEutraCellIdOk returns a tuple with the EutraCellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEutraCellId

`func (o *Ecgi) SetEutraCellId(v string)`

SetEutraCellId sets EutraCellId field to given value.


### GetNid

`func (o *Ecgi) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *Ecgi) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *Ecgi) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *Ecgi) HasNid() bool`

HasNid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


