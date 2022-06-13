# RedundantTransmissionExpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpatialValidCon** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \&quot;Label1.Label2.Label3\&quot;). | [optional] 
**RedTransExps** | [**[]RedundantTransmissionExpPerTS**](RedundantTransmissionExpPerTS.md) |  | 

## Methods

### NewRedundantTransmissionExpInfo

`func NewRedundantTransmissionExpInfo(redTransExps []RedundantTransmissionExpPerTS, ) *RedundantTransmissionExpInfo`

NewRedundantTransmissionExpInfo instantiates a new RedundantTransmissionExpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedundantTransmissionExpInfoWithDefaults

`func NewRedundantTransmissionExpInfoWithDefaults() *RedundantTransmissionExpInfo`

NewRedundantTransmissionExpInfoWithDefaults instantiates a new RedundantTransmissionExpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpatialValidCon

`func (o *RedundantTransmissionExpInfo) GetSpatialValidCon() NetworkAreaInfo`

GetSpatialValidCon returns the SpatialValidCon field if non-nil, zero value otherwise.

### GetSpatialValidConOk

`func (o *RedundantTransmissionExpInfo) GetSpatialValidConOk() (*NetworkAreaInfo, bool)`

GetSpatialValidConOk returns a tuple with the SpatialValidCon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpatialValidCon

`func (o *RedundantTransmissionExpInfo) SetSpatialValidCon(v NetworkAreaInfo)`

SetSpatialValidCon sets SpatialValidCon field to given value.

### HasSpatialValidCon

`func (o *RedundantTransmissionExpInfo) HasSpatialValidCon() bool`

HasSpatialValidCon returns a boolean if a field has been set.

### GetDnn

`func (o *RedundantTransmissionExpInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *RedundantTransmissionExpInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *RedundantTransmissionExpInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *RedundantTransmissionExpInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetRedTransExps

`func (o *RedundantTransmissionExpInfo) GetRedTransExps() []RedundantTransmissionExpPerTS`

GetRedTransExps returns the RedTransExps field if non-nil, zero value otherwise.

### GetRedTransExpsOk

`func (o *RedundantTransmissionExpInfo) GetRedTransExpsOk() (*[]RedundantTransmissionExpPerTS, bool)`

GetRedTransExpsOk returns a tuple with the RedTransExps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedTransExps

`func (o *RedundantTransmissionExpInfo) SetRedTransExps(v []RedundantTransmissionExpPerTS)`

SetRedTransExps sets RedTransExps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


