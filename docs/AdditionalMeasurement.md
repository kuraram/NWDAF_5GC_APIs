# AdditionalMeasurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnexpLoc** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**UnexpFlowTeps** | Pointer to [**[]IpEthFlowDescription**](IpEthFlowDescription.md) |  | [optional] 
**UnexpWakes** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**DdosAttack** | Pointer to [**AddressList**](AddressList.md) |  | [optional] 
**WrgDest** | Pointer to [**AddressList**](AddressList.md) |  | [optional] 
**Circums** | Pointer to [**[]CircumstanceDescription**](CircumstanceDescription.md) |  | [optional] 

## Methods

### NewAdditionalMeasurement

`func NewAdditionalMeasurement() *AdditionalMeasurement`

NewAdditionalMeasurement instantiates a new AdditionalMeasurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalMeasurementWithDefaults

`func NewAdditionalMeasurementWithDefaults() *AdditionalMeasurement`

NewAdditionalMeasurementWithDefaults instantiates a new AdditionalMeasurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnexpLoc

`func (o *AdditionalMeasurement) GetUnexpLoc() NetworkAreaInfo`

GetUnexpLoc returns the UnexpLoc field if non-nil, zero value otherwise.

### GetUnexpLocOk

`func (o *AdditionalMeasurement) GetUnexpLocOk() (*NetworkAreaInfo, bool)`

GetUnexpLocOk returns a tuple with the UnexpLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexpLoc

`func (o *AdditionalMeasurement) SetUnexpLoc(v NetworkAreaInfo)`

SetUnexpLoc sets UnexpLoc field to given value.

### HasUnexpLoc

`func (o *AdditionalMeasurement) HasUnexpLoc() bool`

HasUnexpLoc returns a boolean if a field has been set.

### GetUnexpFlowTeps

`func (o *AdditionalMeasurement) GetUnexpFlowTeps() []IpEthFlowDescription`

GetUnexpFlowTeps returns the UnexpFlowTeps field if non-nil, zero value otherwise.

### GetUnexpFlowTepsOk

`func (o *AdditionalMeasurement) GetUnexpFlowTepsOk() (*[]IpEthFlowDescription, bool)`

GetUnexpFlowTepsOk returns a tuple with the UnexpFlowTeps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexpFlowTeps

`func (o *AdditionalMeasurement) SetUnexpFlowTeps(v []IpEthFlowDescription)`

SetUnexpFlowTeps sets UnexpFlowTeps field to given value.

### HasUnexpFlowTeps

`func (o *AdditionalMeasurement) HasUnexpFlowTeps() bool`

HasUnexpFlowTeps returns a boolean if a field has been set.

### GetUnexpWakes

`func (o *AdditionalMeasurement) GetUnexpWakes() []time.Time`

GetUnexpWakes returns the UnexpWakes field if non-nil, zero value otherwise.

### GetUnexpWakesOk

`func (o *AdditionalMeasurement) GetUnexpWakesOk() (*[]time.Time, bool)`

GetUnexpWakesOk returns a tuple with the UnexpWakes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexpWakes

`func (o *AdditionalMeasurement) SetUnexpWakes(v []time.Time)`

SetUnexpWakes sets UnexpWakes field to given value.

### HasUnexpWakes

`func (o *AdditionalMeasurement) HasUnexpWakes() bool`

HasUnexpWakes returns a boolean if a field has been set.

### GetDdosAttack

`func (o *AdditionalMeasurement) GetDdosAttack() AddressList`

GetDdosAttack returns the DdosAttack field if non-nil, zero value otherwise.

### GetDdosAttackOk

`func (o *AdditionalMeasurement) GetDdosAttackOk() (*AddressList, bool)`

GetDdosAttackOk returns a tuple with the DdosAttack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosAttack

`func (o *AdditionalMeasurement) SetDdosAttack(v AddressList)`

SetDdosAttack sets DdosAttack field to given value.

### HasDdosAttack

`func (o *AdditionalMeasurement) HasDdosAttack() bool`

HasDdosAttack returns a boolean if a field has been set.

### GetWrgDest

`func (o *AdditionalMeasurement) GetWrgDest() AddressList`

GetWrgDest returns the WrgDest field if non-nil, zero value otherwise.

### GetWrgDestOk

`func (o *AdditionalMeasurement) GetWrgDestOk() (*AddressList, bool)`

GetWrgDestOk returns a tuple with the WrgDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrgDest

`func (o *AdditionalMeasurement) SetWrgDest(v AddressList)`

SetWrgDest sets WrgDest field to given value.

### HasWrgDest

`func (o *AdditionalMeasurement) HasWrgDest() bool`

HasWrgDest returns a boolean if a field has been set.

### GetCircums

`func (o *AdditionalMeasurement) GetCircums() []CircumstanceDescription`

GetCircums returns the Circums field if non-nil, zero value otherwise.

### GetCircumsOk

`func (o *AdditionalMeasurement) GetCircumsOk() (*[]CircumstanceDescription, bool)`

GetCircumsOk returns a tuple with the Circums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircums

`func (o *AdditionalMeasurement) SetCircums(v []CircumstanceDescription)`

SetCircums sets Circums field to given value.

### HasCircums

`func (o *AdditionalMeasurement) HasCircums() bool`

HasCircums returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


