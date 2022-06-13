# UeCommunication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommDur** | **int32** | indicating a time in seconds. | 
**CommDurVariance** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 
**PerioTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**PerioTimeVariance** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 
**Ts** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TsVariance** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 
**RecurringTime** | Pointer to [**ScheduledCommunicationTime**](ScheduledCommunicationTime.md) |  | [optional] 
**TrafChar** | [**TrafficCharacterization**](TrafficCharacterization.md) |  | 
**Ratio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | [optional] 
**PerioCommInd** | Pointer to **bool** |  | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**AnaOfAppList** | Pointer to [**AppListForUeComm**](AppListForUeComm.md) |  | [optional] 
**SessInactTimer** | Pointer to [**SessInactTimerForUeComm**](SessInactTimerForUeComm.md) |  | [optional] 

## Methods

### NewUeCommunication

`func NewUeCommunication(commDur int32, trafChar TrafficCharacterization, ) *UeCommunication`

NewUeCommunication instantiates a new UeCommunication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeCommunicationWithDefaults

`func NewUeCommunicationWithDefaults() *UeCommunication`

NewUeCommunicationWithDefaults instantiates a new UeCommunication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommDur

`func (o *UeCommunication) GetCommDur() int32`

GetCommDur returns the CommDur field if non-nil, zero value otherwise.

### GetCommDurOk

`func (o *UeCommunication) GetCommDurOk() (*int32, bool)`

GetCommDurOk returns a tuple with the CommDur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommDur

`func (o *UeCommunication) SetCommDur(v int32)`

SetCommDur sets CommDur field to given value.


### GetCommDurVariance

`func (o *UeCommunication) GetCommDurVariance() float32`

GetCommDurVariance returns the CommDurVariance field if non-nil, zero value otherwise.

### GetCommDurVarianceOk

`func (o *UeCommunication) GetCommDurVarianceOk() (*float32, bool)`

GetCommDurVarianceOk returns a tuple with the CommDurVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommDurVariance

`func (o *UeCommunication) SetCommDurVariance(v float32)`

SetCommDurVariance sets CommDurVariance field to given value.

### HasCommDurVariance

`func (o *UeCommunication) HasCommDurVariance() bool`

HasCommDurVariance returns a boolean if a field has been set.

### GetPerioTime

`func (o *UeCommunication) GetPerioTime() int32`

GetPerioTime returns the PerioTime field if non-nil, zero value otherwise.

### GetPerioTimeOk

`func (o *UeCommunication) GetPerioTimeOk() (*int32, bool)`

GetPerioTimeOk returns a tuple with the PerioTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerioTime

`func (o *UeCommunication) SetPerioTime(v int32)`

SetPerioTime sets PerioTime field to given value.

### HasPerioTime

`func (o *UeCommunication) HasPerioTime() bool`

HasPerioTime returns a boolean if a field has been set.

### GetPerioTimeVariance

`func (o *UeCommunication) GetPerioTimeVariance() float32`

GetPerioTimeVariance returns the PerioTimeVariance field if non-nil, zero value otherwise.

### GetPerioTimeVarianceOk

`func (o *UeCommunication) GetPerioTimeVarianceOk() (*float32, bool)`

GetPerioTimeVarianceOk returns a tuple with the PerioTimeVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerioTimeVariance

`func (o *UeCommunication) SetPerioTimeVariance(v float32)`

SetPerioTimeVariance sets PerioTimeVariance field to given value.

### HasPerioTimeVariance

`func (o *UeCommunication) HasPerioTimeVariance() bool`

HasPerioTimeVariance returns a boolean if a field has been set.

### GetTs

`func (o *UeCommunication) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *UeCommunication) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *UeCommunication) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *UeCommunication) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTsVariance

`func (o *UeCommunication) GetTsVariance() float32`

GetTsVariance returns the TsVariance field if non-nil, zero value otherwise.

### GetTsVarianceOk

`func (o *UeCommunication) GetTsVarianceOk() (*float32, bool)`

GetTsVarianceOk returns a tuple with the TsVariance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsVariance

`func (o *UeCommunication) SetTsVariance(v float32)`

SetTsVariance sets TsVariance field to given value.

### HasTsVariance

`func (o *UeCommunication) HasTsVariance() bool`

HasTsVariance returns a boolean if a field has been set.

### GetRecurringTime

`func (o *UeCommunication) GetRecurringTime() ScheduledCommunicationTime`

GetRecurringTime returns the RecurringTime field if non-nil, zero value otherwise.

### GetRecurringTimeOk

`func (o *UeCommunication) GetRecurringTimeOk() (*ScheduledCommunicationTime, bool)`

GetRecurringTimeOk returns a tuple with the RecurringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTime

`func (o *UeCommunication) SetRecurringTime(v ScheduledCommunicationTime)`

SetRecurringTime sets RecurringTime field to given value.

### HasRecurringTime

`func (o *UeCommunication) HasRecurringTime() bool`

HasRecurringTime returns a boolean if a field has been set.

### GetTrafChar

`func (o *UeCommunication) GetTrafChar() TrafficCharacterization`

GetTrafChar returns the TrafChar field if non-nil, zero value otherwise.

### GetTrafCharOk

`func (o *UeCommunication) GetTrafCharOk() (*TrafficCharacterization, bool)`

GetTrafCharOk returns a tuple with the TrafChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafChar

`func (o *UeCommunication) SetTrafChar(v TrafficCharacterization)`

SetTrafChar sets TrafChar field to given value.


### GetRatio

`func (o *UeCommunication) GetRatio() int32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *UeCommunication) GetRatioOk() (*int32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *UeCommunication) SetRatio(v int32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *UeCommunication) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetPerioCommInd

`func (o *UeCommunication) GetPerioCommInd() bool`

GetPerioCommInd returns the PerioCommInd field if non-nil, zero value otherwise.

### GetPerioCommIndOk

`func (o *UeCommunication) GetPerioCommIndOk() (*bool, bool)`

GetPerioCommIndOk returns a tuple with the PerioCommInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerioCommInd

`func (o *UeCommunication) SetPerioCommInd(v bool)`

SetPerioCommInd sets PerioCommInd field to given value.

### HasPerioCommInd

`func (o *UeCommunication) HasPerioCommInd() bool`

HasPerioCommInd returns a boolean if a field has been set.

### GetConfidence

`func (o *UeCommunication) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *UeCommunication) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *UeCommunication) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *UeCommunication) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetAnaOfAppList

`func (o *UeCommunication) GetAnaOfAppList() AppListForUeComm`

GetAnaOfAppList returns the AnaOfAppList field if non-nil, zero value otherwise.

### GetAnaOfAppListOk

`func (o *UeCommunication) GetAnaOfAppListOk() (*AppListForUeComm, bool)`

GetAnaOfAppListOk returns a tuple with the AnaOfAppList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaOfAppList

`func (o *UeCommunication) SetAnaOfAppList(v AppListForUeComm)`

SetAnaOfAppList sets AnaOfAppList field to given value.

### HasAnaOfAppList

`func (o *UeCommunication) HasAnaOfAppList() bool`

HasAnaOfAppList returns a boolean if a field has been set.

### GetSessInactTimer

`func (o *UeCommunication) GetSessInactTimer() SessInactTimerForUeComm`

GetSessInactTimer returns the SessInactTimer field if non-nil, zero value otherwise.

### GetSessInactTimerOk

`func (o *UeCommunication) GetSessInactTimerOk() (*SessInactTimerForUeComm, bool)`

GetSessInactTimerOk returns a tuple with the SessInactTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessInactTimer

`func (o *UeCommunication) SetSessInactTimer(v SessInactTimerForUeComm)`

SetSessInactTimer sets SessInactTimer field to given value.

### HasSessInactTimer

`func (o *UeCommunication) HasSessInactTimer() bool`

HasSessInactTimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


