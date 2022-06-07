# TrafficCharacterization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | Pointer to **string** |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**FDescs** | Pointer to [**[]IpEthFlowDescription**](IpEthFlowDescription.md) |  | [optional] 
**UlVol** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**UlVolVariance** | Pointer to **float32** |  | [optional] 
**DlVol** | Pointer to **int64** | Unsigned integer identifying a volume in units of bytes. | [optional] 
**DlVolVariance** | Pointer to **float32** |  | [optional] 

## Methods

### NewTrafficCharacterization

`func NewTrafficCharacterization() *TrafficCharacterization`

NewTrafficCharacterization instantiates a new TrafficCharacterization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficCharacterizationWithDefaults

`func NewTrafficCharacterizationWithDefaults() *TrafficCharacterization`

NewTrafficCharacterizationWithDefaults instantiates a new TrafficCharacterization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *TrafficCharacterization) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *TrafficCharacterization) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *TrafficCharacterization) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *TrafficCharacterization) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *TrafficCharacterization) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *TrafficCharacterization) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *TrafficCharacterization) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *TrafficCharacterization) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetAppId

`func (o *TrafficCharacterization) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *TrafficCharacterization) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *TrafficCharacterization) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *TrafficCharacterization) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetFDescs

`func (o *TrafficCharacterization) GetFDescs() []IpEthFlowDescription`

GetFDescs returns the FDescs field if non-nil, zero value otherwise.

### GetFDescsOk

`func (o *TrafficCharacterization) GetFDescsOk() (*[]IpEthFlowDescription, bool)`

GetFDescsOk returns a tuple with the FDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFDescs

`func (o *TrafficCharacterization) SetFDescs(v []IpEthFlowDescription)`

SetFDescs sets FDescs field to given value.

### HasFDescs

`func (o *TrafficCharacterization) HasFDescs() bool`

HasFDescs returns a boolean if a field has been set.

### GetUlVol

`func (o *TrafficCharacterization) GetUlVol() int64`

GetUlVol returns the UlVol field if non-nil, zero value otherwise.

### GetUlVolOk

`func (o *TrafficCharacterization) GetUlVolOk() (*int64, bool)`

GetUlVolOk returns a tuple with the UlVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlVol

`func (o *TrafficCharacterization) SetUlVol(v int64)`

SetUlVol sets UlVol field to given value.

### HasUlVol

`func (o *TrafficCharacterization) HasUlVol() bool`

HasUlVol returns a boolean if a field has been set.

### GetUlVolVariance

`func (o *TrafficCharacterization) GetUlVolVariance() float32`

GetUlVolVariance returns the UlVolVariance field if non-nil, zero value otherwise.

### GetUlVolVarianceOk

`func (o *TrafficCharacterization) GetUlVolVarianceOk() (*float32, bool)`

GetUlVolVarianceOk returns a tuple with the UlVolVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUlVolVariance

`func (o *TrafficCharacterization) SetUlVolVariance(v float32)`

SetUlVolVariance sets UlVolVariance field to given value.

### HasUlVolVariance

`func (o *TrafficCharacterization) HasUlVolVariance() bool`

HasUlVolVariance returns a boolean if a field has been set.

### GetDlVol

`func (o *TrafficCharacterization) GetDlVol() int64`

GetDlVol returns the DlVol field if non-nil, zero value otherwise.

### GetDlVolOk

`func (o *TrafficCharacterization) GetDlVolOk() (*int64, bool)`

GetDlVolOk returns a tuple with the DlVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlVol

`func (o *TrafficCharacterization) SetDlVol(v int64)`

SetDlVol sets DlVol field to given value.

### HasDlVol

`func (o *TrafficCharacterization) HasDlVol() bool`

HasDlVol returns a boolean if a field has been set.

### GetDlVolVariance

`func (o *TrafficCharacterization) GetDlVolVariance() float32`

GetDlVolVariance returns the DlVolVariance field if non-nil, zero value otherwise.

### GetDlVolVarianceOk

`func (o *TrafficCharacterization) GetDlVolVarianceOk() (*float32, bool)`

GetDlVolVarianceOk returns a tuple with the DlVolVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlVolVariance

`func (o *TrafficCharacterization) SetDlVolVariance(v float32)`

SetDlVolVariance sets DlVolVariance field to given value.

### HasDlVolVariance

`func (o *TrafficCharacterization) HasDlVolVariance() bool`

HasDlVolVariance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


