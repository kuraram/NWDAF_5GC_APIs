# QosRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5qi** | Pointer to **int32** | Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255 | [optional] 
**GfbrUl** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**GfbrDl** | Pointer to **string** | String representing a bit rate prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;. | [optional] 
**ResType** | Pointer to [**QosResourceType**](QosResourceType.md) |  | [optional] 
**Pdb** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. | [optional] 
**Per** | Pointer to **string** | String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \&quot;scalar x 10-k\&quot; where the scalar and the exponent k are each encoded as one decimal digit.  | [optional] 

## Methods

### NewQosRequirement

`func NewQosRequirement() *QosRequirement`

NewQosRequirement instantiates a new QosRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosRequirementWithDefaults

`func NewQosRequirementWithDefaults() *QosRequirement`

NewQosRequirementWithDefaults instantiates a new QosRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar5qi

`func (o *QosRequirement) GetVar5qi() int32`

GetVar5qi returns the Var5qi field if non-nil, zero value otherwise.

### GetVar5qiOk

`func (o *QosRequirement) GetVar5qiOk() (*int32, bool)`

GetVar5qiOk returns a tuple with the Var5qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qi

`func (o *QosRequirement) SetVar5qi(v int32)`

SetVar5qi sets Var5qi field to given value.

### HasVar5qi

`func (o *QosRequirement) HasVar5qi() bool`

HasVar5qi returns a boolean if a field has been set.

### GetGfbrUl

`func (o *QosRequirement) GetGfbrUl() string`

GetGfbrUl returns the GfbrUl field if non-nil, zero value otherwise.

### GetGfbrUlOk

`func (o *QosRequirement) GetGfbrUlOk() (*string, bool)`

GetGfbrUlOk returns a tuple with the GfbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGfbrUl

`func (o *QosRequirement) SetGfbrUl(v string)`

SetGfbrUl sets GfbrUl field to given value.

### HasGfbrUl

`func (o *QosRequirement) HasGfbrUl() bool`

HasGfbrUl returns a boolean if a field has been set.

### GetGfbrDl

`func (o *QosRequirement) GetGfbrDl() string`

GetGfbrDl returns the GfbrDl field if non-nil, zero value otherwise.

### GetGfbrDlOk

`func (o *QosRequirement) GetGfbrDlOk() (*string, bool)`

GetGfbrDlOk returns a tuple with the GfbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGfbrDl

`func (o *QosRequirement) SetGfbrDl(v string)`

SetGfbrDl sets GfbrDl field to given value.

### HasGfbrDl

`func (o *QosRequirement) HasGfbrDl() bool`

HasGfbrDl returns a boolean if a field has been set.

### GetResType

`func (o *QosRequirement) GetResType() QosResourceType`

GetResType returns the ResType field if non-nil, zero value otherwise.

### GetResTypeOk

`func (o *QosRequirement) GetResTypeOk() (*QosResourceType, bool)`

GetResTypeOk returns a tuple with the ResType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResType

`func (o *QosRequirement) SetResType(v QosResourceType)`

SetResType sets ResType field to given value.

### HasResType

`func (o *QosRequirement) HasResType() bool`

HasResType returns a boolean if a field has been set.

### GetPdb

`func (o *QosRequirement) GetPdb() int32`

GetPdb returns the Pdb field if non-nil, zero value otherwise.

### GetPdbOk

`func (o *QosRequirement) GetPdbOk() (*int32, bool)`

GetPdbOk returns a tuple with the Pdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdb

`func (o *QosRequirement) SetPdb(v int32)`

SetPdb sets Pdb field to given value.

### HasPdb

`func (o *QosRequirement) HasPdb() bool`

HasPdb returns a boolean if a field has been set.

### GetPer

`func (o *QosRequirement) GetPer() string`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *QosRequirement) GetPerOk() (*string, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *QosRequirement) SetPer(v string)`

SetPer sets Per field to given value.

### HasPer

`func (o *QosRequirement) HasPer() bool`

HasPer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


