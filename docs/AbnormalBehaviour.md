# AbnormalBehaviour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supis** | Pointer to **[]string** |  | [optional] 
**Excep** | [**Exception**](Exception.md) |  | 
**Dnn** | Pointer to **string** |  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Ratio** | Pointer to **int32** |  | [optional] 
**Confidence** | Pointer to **int32** |  | [optional] 
**AddtMeasInfo** | Pointer to [**AdditionalMeasurement**](AdditionalMeasurement.md) |  | [optional] 

## Methods

### NewAbnormalBehaviour

`func NewAbnormalBehaviour(excep Exception, ) *AbnormalBehaviour`

NewAbnormalBehaviour instantiates a new AbnormalBehaviour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbnormalBehaviourWithDefaults

`func NewAbnormalBehaviourWithDefaults() *AbnormalBehaviour`

NewAbnormalBehaviourWithDefaults instantiates a new AbnormalBehaviour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupis

`func (o *AbnormalBehaviour) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *AbnormalBehaviour) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *AbnormalBehaviour) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *AbnormalBehaviour) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetExcep

`func (o *AbnormalBehaviour) GetExcep() Exception`

GetExcep returns the Excep field if non-nil, zero value otherwise.

### GetExcepOk

`func (o *AbnormalBehaviour) GetExcepOk() (*Exception, bool)`

GetExcepOk returns a tuple with the Excep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcep

`func (o *AbnormalBehaviour) SetExcep(v Exception)`

SetExcep sets Excep field to given value.


### GetDnn

`func (o *AbnormalBehaviour) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *AbnormalBehaviour) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *AbnormalBehaviour) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *AbnormalBehaviour) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *AbnormalBehaviour) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *AbnormalBehaviour) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *AbnormalBehaviour) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *AbnormalBehaviour) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetRatio

`func (o *AbnormalBehaviour) GetRatio() int32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *AbnormalBehaviour) GetRatioOk() (*int32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *AbnormalBehaviour) SetRatio(v int32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *AbnormalBehaviour) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetConfidence

`func (o *AbnormalBehaviour) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *AbnormalBehaviour) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *AbnormalBehaviour) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *AbnormalBehaviour) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetAddtMeasInfo

`func (o *AbnormalBehaviour) GetAddtMeasInfo() AdditionalMeasurement`

GetAddtMeasInfo returns the AddtMeasInfo field if non-nil, zero value otherwise.

### GetAddtMeasInfoOk

`func (o *AbnormalBehaviour) GetAddtMeasInfoOk() (*AdditionalMeasurement, bool)`

GetAddtMeasInfoOk returns a tuple with the AddtMeasInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddtMeasInfo

`func (o *AbnormalBehaviour) SetAddtMeasInfo(v AdditionalMeasurement)`

SetAddtMeasInfo sets AddtMeasInfo field to given value.

### HasAddtMeasInfo

`func (o *AbnormalBehaviour) HasAddtMeasInfo() bool`

HasAddtMeasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


