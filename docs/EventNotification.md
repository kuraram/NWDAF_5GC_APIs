# EventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**NwdafEvent**](NwdafEvent.md) |  | 
**Start** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TimeStampGen** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**FailNotifyCode** | Pointer to [**NwdafFailureCode**](NwdafFailureCode.md) |  | [optional] 
**RvWaitTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**AnaMetaInfo** | Pointer to [**AnalyticsMetadataInfo**](AnalyticsMetadataInfo.md) |  | [optional] 
**NfLoadLevelInfos** | Pointer to [**[]NfLoadLevelInformation**](NfLoadLevelInformation.md) |  | [optional] 
**NsiLoadLevelInfos** | Pointer to [**[]NsiLoadLevelInfo**](NsiLoadLevelInfo.md) |  | [optional] 
**SliceLoadLevelInfo** | Pointer to [**SliceLoadLevelInformation**](SliceLoadLevelInformation.md) |  | [optional] 
**SvcExps** | Pointer to [**[]ServiceExperienceInfo**](ServiceExperienceInfo.md) |  | [optional] 
**QosSustainInfos** | Pointer to [**[]QosSustainabilityInfo**](QosSustainabilityInfo.md) |  | [optional] 
**UeComms** | Pointer to [**[]UeCommunication**](UeCommunication.md) |  | [optional] 
**UeMobs** | Pointer to [**[]UeMobility**](UeMobility.md) |  | [optional] 
**UserDataCongInfos** | Pointer to [**[]UserDataCongestionInfo**](UserDataCongestionInfo.md) |  | [optional] 
**AbnorBehavrs** | Pointer to [**[]AbnormalBehaviour**](AbnormalBehaviour.md) |  | [optional] 
**NwPerfs** | Pointer to [**[]NetworkPerfInfo**](NetworkPerfInfo.md) |  | [optional] 
**DnPerfInfos** | Pointer to [**[]DnPerfInfo**](DnPerfInfo.md) |  | [optional] 
**DisperInfos** | Pointer to [**[]DispersionInfo**](DispersionInfo.md) |  | [optional] 
**RedTransInfos** | Pointer to [**[]RedundantTransmissionExpInfo**](RedundantTransmissionExpInfo.md) |  | [optional] 
**WlanInfos** | Pointer to [**[]WlanPerformanceInfo**](WlanPerformanceInfo.md) |  | [optional] 

## Methods

### NewEventNotification

`func NewEventNotification(event NwdafEvent, ) *EventNotification`

NewEventNotification instantiates a new EventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationWithDefaults

`func NewEventNotificationWithDefaults() *EventNotification`

NewEventNotificationWithDefaults instantiates a new EventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventNotification) GetEvent() NwdafEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventNotification) GetEventOk() (*NwdafEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventNotification) SetEvent(v NwdafEvent)`

SetEvent sets Event field to given value.


### GetStart

`func (o *EventNotification) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EventNotification) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EventNotification) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *EventNotification) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetExpiry

`func (o *EventNotification) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *EventNotification) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *EventNotification) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *EventNotification) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetTimeStampGen

`func (o *EventNotification) GetTimeStampGen() time.Time`

GetTimeStampGen returns the TimeStampGen field if non-nil, zero value otherwise.

### GetTimeStampGenOk

`func (o *EventNotification) GetTimeStampGenOk() (*time.Time, bool)`

GetTimeStampGenOk returns a tuple with the TimeStampGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStampGen

`func (o *EventNotification) SetTimeStampGen(v time.Time)`

SetTimeStampGen sets TimeStampGen field to given value.

### HasTimeStampGen

`func (o *EventNotification) HasTimeStampGen() bool`

HasTimeStampGen returns a boolean if a field has been set.

### GetFailNotifyCode

`func (o *EventNotification) GetFailNotifyCode() NwdafFailureCode`

GetFailNotifyCode returns the FailNotifyCode field if non-nil, zero value otherwise.

### GetFailNotifyCodeOk

`func (o *EventNotification) GetFailNotifyCodeOk() (*NwdafFailureCode, bool)`

GetFailNotifyCodeOk returns a tuple with the FailNotifyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailNotifyCode

`func (o *EventNotification) SetFailNotifyCode(v NwdafFailureCode)`

SetFailNotifyCode sets FailNotifyCode field to given value.

### HasFailNotifyCode

`func (o *EventNotification) HasFailNotifyCode() bool`

HasFailNotifyCode returns a boolean if a field has been set.

### GetRvWaitTime

`func (o *EventNotification) GetRvWaitTime() int32`

GetRvWaitTime returns the RvWaitTime field if non-nil, zero value otherwise.

### GetRvWaitTimeOk

`func (o *EventNotification) GetRvWaitTimeOk() (*int32, bool)`

GetRvWaitTimeOk returns a tuple with the RvWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRvWaitTime

`func (o *EventNotification) SetRvWaitTime(v int32)`

SetRvWaitTime sets RvWaitTime field to given value.

### HasRvWaitTime

`func (o *EventNotification) HasRvWaitTime() bool`

HasRvWaitTime returns a boolean if a field has been set.

### GetAnaMetaInfo

`func (o *EventNotification) GetAnaMetaInfo() AnalyticsMetadataInfo`

GetAnaMetaInfo returns the AnaMetaInfo field if non-nil, zero value otherwise.

### GetAnaMetaInfoOk

`func (o *EventNotification) GetAnaMetaInfoOk() (*AnalyticsMetadataInfo, bool)`

GetAnaMetaInfoOk returns a tuple with the AnaMetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaMetaInfo

`func (o *EventNotification) SetAnaMetaInfo(v AnalyticsMetadataInfo)`

SetAnaMetaInfo sets AnaMetaInfo field to given value.

### HasAnaMetaInfo

`func (o *EventNotification) HasAnaMetaInfo() bool`

HasAnaMetaInfo returns a boolean if a field has been set.

### GetNfLoadLevelInfos

`func (o *EventNotification) GetNfLoadLevelInfos() []NfLoadLevelInformation`

GetNfLoadLevelInfos returns the NfLoadLevelInfos field if non-nil, zero value otherwise.

### GetNfLoadLevelInfosOk

`func (o *EventNotification) GetNfLoadLevelInfosOk() (*[]NfLoadLevelInformation, bool)`

GetNfLoadLevelInfosOk returns a tuple with the NfLoadLevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfLoadLevelInfos

`func (o *EventNotification) SetNfLoadLevelInfos(v []NfLoadLevelInformation)`

SetNfLoadLevelInfos sets NfLoadLevelInfos field to given value.

### HasNfLoadLevelInfos

`func (o *EventNotification) HasNfLoadLevelInfos() bool`

HasNfLoadLevelInfos returns a boolean if a field has been set.

### GetNsiLoadLevelInfos

`func (o *EventNotification) GetNsiLoadLevelInfos() []NsiLoadLevelInfo`

GetNsiLoadLevelInfos returns the NsiLoadLevelInfos field if non-nil, zero value otherwise.

### GetNsiLoadLevelInfosOk

`func (o *EventNotification) GetNsiLoadLevelInfosOk() (*[]NsiLoadLevelInfo, bool)`

GetNsiLoadLevelInfosOk returns a tuple with the NsiLoadLevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiLoadLevelInfos

`func (o *EventNotification) SetNsiLoadLevelInfos(v []NsiLoadLevelInfo)`

SetNsiLoadLevelInfos sets NsiLoadLevelInfos field to given value.

### HasNsiLoadLevelInfos

`func (o *EventNotification) HasNsiLoadLevelInfos() bool`

HasNsiLoadLevelInfos returns a boolean if a field has been set.

### GetSliceLoadLevelInfo

`func (o *EventNotification) GetSliceLoadLevelInfo() SliceLoadLevelInformation`

GetSliceLoadLevelInfo returns the SliceLoadLevelInfo field if non-nil, zero value otherwise.

### GetSliceLoadLevelInfoOk

`func (o *EventNotification) GetSliceLoadLevelInfoOk() (*SliceLoadLevelInformation, bool)`

GetSliceLoadLevelInfoOk returns a tuple with the SliceLoadLevelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceLoadLevelInfo

`func (o *EventNotification) SetSliceLoadLevelInfo(v SliceLoadLevelInformation)`

SetSliceLoadLevelInfo sets SliceLoadLevelInfo field to given value.

### HasSliceLoadLevelInfo

`func (o *EventNotification) HasSliceLoadLevelInfo() bool`

HasSliceLoadLevelInfo returns a boolean if a field has been set.

### GetSvcExps

`func (o *EventNotification) GetSvcExps() []ServiceExperienceInfo`

GetSvcExps returns the SvcExps field if non-nil, zero value otherwise.

### GetSvcExpsOk

`func (o *EventNotification) GetSvcExpsOk() (*[]ServiceExperienceInfo, bool)`

GetSvcExpsOk returns a tuple with the SvcExps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExps

`func (o *EventNotification) SetSvcExps(v []ServiceExperienceInfo)`

SetSvcExps sets SvcExps field to given value.

### HasSvcExps

`func (o *EventNotification) HasSvcExps() bool`

HasSvcExps returns a boolean if a field has been set.

### GetQosSustainInfos

`func (o *EventNotification) GetQosSustainInfos() []QosSustainabilityInfo`

GetQosSustainInfos returns the QosSustainInfos field if non-nil, zero value otherwise.

### GetQosSustainInfosOk

`func (o *EventNotification) GetQosSustainInfosOk() (*[]QosSustainabilityInfo, bool)`

GetQosSustainInfosOk returns a tuple with the QosSustainInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSustainInfos

`func (o *EventNotification) SetQosSustainInfos(v []QosSustainabilityInfo)`

SetQosSustainInfos sets QosSustainInfos field to given value.

### HasQosSustainInfos

`func (o *EventNotification) HasQosSustainInfos() bool`

HasQosSustainInfos returns a boolean if a field has been set.

### GetUeComms

`func (o *EventNotification) GetUeComms() []UeCommunication`

GetUeComms returns the UeComms field if non-nil, zero value otherwise.

### GetUeCommsOk

`func (o *EventNotification) GetUeCommsOk() (*[]UeCommunication, bool)`

GetUeCommsOk returns a tuple with the UeComms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeComms

`func (o *EventNotification) SetUeComms(v []UeCommunication)`

SetUeComms sets UeComms field to given value.

### HasUeComms

`func (o *EventNotification) HasUeComms() bool`

HasUeComms returns a boolean if a field has been set.

### GetUeMobs

`func (o *EventNotification) GetUeMobs() []UeMobility`

GetUeMobs returns the UeMobs field if non-nil, zero value otherwise.

### GetUeMobsOk

`func (o *EventNotification) GetUeMobsOk() (*[]UeMobility, bool)`

GetUeMobsOk returns a tuple with the UeMobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMobs

`func (o *EventNotification) SetUeMobs(v []UeMobility)`

SetUeMobs sets UeMobs field to given value.

### HasUeMobs

`func (o *EventNotification) HasUeMobs() bool`

HasUeMobs returns a boolean if a field has been set.

### GetUserDataCongInfos

`func (o *EventNotification) GetUserDataCongInfos() []UserDataCongestionInfo`

GetUserDataCongInfos returns the UserDataCongInfos field if non-nil, zero value otherwise.

### GetUserDataCongInfosOk

`func (o *EventNotification) GetUserDataCongInfosOk() (*[]UserDataCongestionInfo, bool)`

GetUserDataCongInfosOk returns a tuple with the UserDataCongInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataCongInfos

`func (o *EventNotification) SetUserDataCongInfos(v []UserDataCongestionInfo)`

SetUserDataCongInfos sets UserDataCongInfos field to given value.

### HasUserDataCongInfos

`func (o *EventNotification) HasUserDataCongInfos() bool`

HasUserDataCongInfos returns a boolean if a field has been set.

### GetAbnorBehavrs

`func (o *EventNotification) GetAbnorBehavrs() []AbnormalBehaviour`

GetAbnorBehavrs returns the AbnorBehavrs field if non-nil, zero value otherwise.

### GetAbnorBehavrsOk

`func (o *EventNotification) GetAbnorBehavrsOk() (*[]AbnormalBehaviour, bool)`

GetAbnorBehavrsOk returns a tuple with the AbnorBehavrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnorBehavrs

`func (o *EventNotification) SetAbnorBehavrs(v []AbnormalBehaviour)`

SetAbnorBehavrs sets AbnorBehavrs field to given value.

### HasAbnorBehavrs

`func (o *EventNotification) HasAbnorBehavrs() bool`

HasAbnorBehavrs returns a boolean if a field has been set.

### GetNwPerfs

`func (o *EventNotification) GetNwPerfs() []NetworkPerfInfo`

GetNwPerfs returns the NwPerfs field if non-nil, zero value otherwise.

### GetNwPerfsOk

`func (o *EventNotification) GetNwPerfsOk() (*[]NetworkPerfInfo, bool)`

GetNwPerfsOk returns a tuple with the NwPerfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfs

`func (o *EventNotification) SetNwPerfs(v []NetworkPerfInfo)`

SetNwPerfs sets NwPerfs field to given value.

### HasNwPerfs

`func (o *EventNotification) HasNwPerfs() bool`

HasNwPerfs returns a boolean if a field has been set.

### GetDnPerfInfos

`func (o *EventNotification) GetDnPerfInfos() []DnPerfInfo`

GetDnPerfInfos returns the DnPerfInfos field if non-nil, zero value otherwise.

### GetDnPerfInfosOk

`func (o *EventNotification) GetDnPerfInfosOk() (*[]DnPerfInfo, bool)`

GetDnPerfInfosOk returns a tuple with the DnPerfInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPerfInfos

`func (o *EventNotification) SetDnPerfInfos(v []DnPerfInfo)`

SetDnPerfInfos sets DnPerfInfos field to given value.

### HasDnPerfInfos

`func (o *EventNotification) HasDnPerfInfos() bool`

HasDnPerfInfos returns a boolean if a field has been set.

### GetDisperInfos

`func (o *EventNotification) GetDisperInfos() []DispersionInfo`

GetDisperInfos returns the DisperInfos field if non-nil, zero value otherwise.

### GetDisperInfosOk

`func (o *EventNotification) GetDisperInfosOk() (*[]DispersionInfo, bool)`

GetDisperInfosOk returns a tuple with the DisperInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperInfos

`func (o *EventNotification) SetDisperInfos(v []DispersionInfo)`

SetDisperInfos sets DisperInfos field to given value.

### HasDisperInfos

`func (o *EventNotification) HasDisperInfos() bool`

HasDisperInfos returns a boolean if a field has been set.

### GetRedTransInfos

`func (o *EventNotification) GetRedTransInfos() []RedundantTransmissionExpInfo`

GetRedTransInfos returns the RedTransInfos field if non-nil, zero value otherwise.

### GetRedTransInfosOk

`func (o *EventNotification) GetRedTransInfosOk() (*[]RedundantTransmissionExpInfo, bool)`

GetRedTransInfosOk returns a tuple with the RedTransInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedTransInfos

`func (o *EventNotification) SetRedTransInfos(v []RedundantTransmissionExpInfo)`

SetRedTransInfos sets RedTransInfos field to given value.

### HasRedTransInfos

`func (o *EventNotification) HasRedTransInfos() bool`

HasRedTransInfos returns a boolean if a field has been set.

### GetWlanInfos

`func (o *EventNotification) GetWlanInfos() []WlanPerformanceInfo`

GetWlanInfos returns the WlanInfos field if non-nil, zero value otherwise.

### GetWlanInfosOk

`func (o *EventNotification) GetWlanInfosOk() (*[]WlanPerformanceInfo, bool)`

GetWlanInfosOk returns a tuple with the WlanInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanInfos

`func (o *EventNotification) SetWlanInfos(v []WlanPerformanceInfo)`

SetWlanInfos sets WlanInfos field to given value.

### HasWlanInfos

`func (o *EventNotification) HasWlanInfos() bool`

HasWlanInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


