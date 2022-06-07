# AnalyticsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **time.Time** |  | [optional] 
**TimeStampGen** | Pointer to **time.Time** |  | [optional] 
**SliceLoadLevelInfos** | Pointer to [**[]SliceLoadLevelInformation**](SliceLoadLevelInformation.md) | The slices and their load level information. | [optional] 
**NsiLoadLevelInfos** | Pointer to [**[]NsiLoadLevelInfo**](NsiLoadLevelInfo.md) |  | [optional] 
**NfLoadLevelInfos** | Pointer to [**[]NfLoadLevelInformation**](NfLoadLevelInformation.md) |  | [optional] 
**NwPerfs** | Pointer to [**[]NetworkPerfInfo**](NetworkPerfInfo.md) |  | [optional] 
**SvcExps** | Pointer to [**[]ServiceExperienceInfo**](ServiceExperienceInfo.md) |  | [optional] 
**QosSustainInfos** | Pointer to [**[]QosSustainabilityInfo**](QosSustainabilityInfo.md) |  | [optional] 
**UeMobs** | Pointer to [**[]UeMobility**](UeMobility.md) |  | [optional] 
**UeComms** | Pointer to [**[]UeCommunication**](UeCommunication.md) |  | [optional] 
**UserDataCongInfos** | Pointer to [**[]UserDataCongestionInfo**](UserDataCongestionInfo.md) |  | [optional] 
**AbnorBehavrs** | Pointer to [**[]AbnormalBehaviour**](AbnormalBehaviour.md) |  | [optional] 
**SuppFeat** | Pointer to **string** |  | [optional] 

## Methods

### NewAnalyticsData

`func NewAnalyticsData() *AnalyticsData`

NewAnalyticsData instantiates a new AnalyticsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsDataWithDefaults

`func NewAnalyticsDataWithDefaults() *AnalyticsData`

NewAnalyticsDataWithDefaults instantiates a new AnalyticsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *AnalyticsData) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AnalyticsData) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AnalyticsData) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AnalyticsData) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetTimeStampGen

`func (o *AnalyticsData) GetTimeStampGen() time.Time`

GetTimeStampGen returns the TimeStampGen field if non-nil, zero value otherwise.

### GetTimeStampGenOk

`func (o *AnalyticsData) GetTimeStampGenOk() (*time.Time, bool)`

GetTimeStampGenOk returns a tuple with the TimeStampGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStampGen

`func (o *AnalyticsData) SetTimeStampGen(v time.Time)`

SetTimeStampGen sets TimeStampGen field to given value.

### HasTimeStampGen

`func (o *AnalyticsData) HasTimeStampGen() bool`

HasTimeStampGen returns a boolean if a field has been set.

### GetSliceLoadLevelInfos

`func (o *AnalyticsData) GetSliceLoadLevelInfos() []SliceLoadLevelInformation`

GetSliceLoadLevelInfos returns the SliceLoadLevelInfos field if non-nil, zero value otherwise.

### GetSliceLoadLevelInfosOk

`func (o *AnalyticsData) GetSliceLoadLevelInfosOk() (*[]SliceLoadLevelInformation, bool)`

GetSliceLoadLevelInfosOk returns a tuple with the SliceLoadLevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceLoadLevelInfos

`func (o *AnalyticsData) SetSliceLoadLevelInfos(v []SliceLoadLevelInformation)`

SetSliceLoadLevelInfos sets SliceLoadLevelInfos field to given value.

### HasSliceLoadLevelInfos

`func (o *AnalyticsData) HasSliceLoadLevelInfos() bool`

HasSliceLoadLevelInfos returns a boolean if a field has been set.

### GetNsiLoadLevelInfos

`func (o *AnalyticsData) GetNsiLoadLevelInfos() []NsiLoadLevelInfo`

GetNsiLoadLevelInfos returns the NsiLoadLevelInfos field if non-nil, zero value otherwise.

### GetNsiLoadLevelInfosOk

`func (o *AnalyticsData) GetNsiLoadLevelInfosOk() (*[]NsiLoadLevelInfo, bool)`

GetNsiLoadLevelInfosOk returns a tuple with the NsiLoadLevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiLoadLevelInfos

`func (o *AnalyticsData) SetNsiLoadLevelInfos(v []NsiLoadLevelInfo)`

SetNsiLoadLevelInfos sets NsiLoadLevelInfos field to given value.

### HasNsiLoadLevelInfos

`func (o *AnalyticsData) HasNsiLoadLevelInfos() bool`

HasNsiLoadLevelInfos returns a boolean if a field has been set.

### GetNfLoadLevelInfos

`func (o *AnalyticsData) GetNfLoadLevelInfos() []NfLoadLevelInformation`

GetNfLoadLevelInfos returns the NfLoadLevelInfos field if non-nil, zero value otherwise.

### GetNfLoadLevelInfosOk

`func (o *AnalyticsData) GetNfLoadLevelInfosOk() (*[]NfLoadLevelInformation, bool)`

GetNfLoadLevelInfosOk returns a tuple with the NfLoadLevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfLoadLevelInfos

`func (o *AnalyticsData) SetNfLoadLevelInfos(v []NfLoadLevelInformation)`

SetNfLoadLevelInfos sets NfLoadLevelInfos field to given value.

### HasNfLoadLevelInfos

`func (o *AnalyticsData) HasNfLoadLevelInfos() bool`

HasNfLoadLevelInfos returns a boolean if a field has been set.

### GetNwPerfs

`func (o *AnalyticsData) GetNwPerfs() []NetworkPerfInfo`

GetNwPerfs returns the NwPerfs field if non-nil, zero value otherwise.

### GetNwPerfsOk

`func (o *AnalyticsData) GetNwPerfsOk() (*[]NetworkPerfInfo, bool)`

GetNwPerfsOk returns a tuple with the NwPerfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfs

`func (o *AnalyticsData) SetNwPerfs(v []NetworkPerfInfo)`

SetNwPerfs sets NwPerfs field to given value.

### HasNwPerfs

`func (o *AnalyticsData) HasNwPerfs() bool`

HasNwPerfs returns a boolean if a field has been set.

### GetSvcExps

`func (o *AnalyticsData) GetSvcExps() []ServiceExperienceInfo`

GetSvcExps returns the SvcExps field if non-nil, zero value otherwise.

### GetSvcExpsOk

`func (o *AnalyticsData) GetSvcExpsOk() (*[]ServiceExperienceInfo, bool)`

GetSvcExpsOk returns a tuple with the SvcExps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcExps

`func (o *AnalyticsData) SetSvcExps(v []ServiceExperienceInfo)`

SetSvcExps sets SvcExps field to given value.

### HasSvcExps

`func (o *AnalyticsData) HasSvcExps() bool`

HasSvcExps returns a boolean if a field has been set.

### GetQosSustainInfos

`func (o *AnalyticsData) GetQosSustainInfos() []QosSustainabilityInfo`

GetQosSustainInfos returns the QosSustainInfos field if non-nil, zero value otherwise.

### GetQosSustainInfosOk

`func (o *AnalyticsData) GetQosSustainInfosOk() (*[]QosSustainabilityInfo, bool)`

GetQosSustainInfosOk returns a tuple with the QosSustainInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSustainInfos

`func (o *AnalyticsData) SetQosSustainInfos(v []QosSustainabilityInfo)`

SetQosSustainInfos sets QosSustainInfos field to given value.

### HasQosSustainInfos

`func (o *AnalyticsData) HasQosSustainInfos() bool`

HasQosSustainInfos returns a boolean if a field has been set.

### GetUeMobs

`func (o *AnalyticsData) GetUeMobs() []UeMobility`

GetUeMobs returns the UeMobs field if non-nil, zero value otherwise.

### GetUeMobsOk

`func (o *AnalyticsData) GetUeMobsOk() (*[]UeMobility, bool)`

GetUeMobsOk returns a tuple with the UeMobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMobs

`func (o *AnalyticsData) SetUeMobs(v []UeMobility)`

SetUeMobs sets UeMobs field to given value.

### HasUeMobs

`func (o *AnalyticsData) HasUeMobs() bool`

HasUeMobs returns a boolean if a field has been set.

### GetUeComms

`func (o *AnalyticsData) GetUeComms() []UeCommunication`

GetUeComms returns the UeComms field if non-nil, zero value otherwise.

### GetUeCommsOk

`func (o *AnalyticsData) GetUeCommsOk() (*[]UeCommunication, bool)`

GetUeCommsOk returns a tuple with the UeComms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeComms

`func (o *AnalyticsData) SetUeComms(v []UeCommunication)`

SetUeComms sets UeComms field to given value.

### HasUeComms

`func (o *AnalyticsData) HasUeComms() bool`

HasUeComms returns a boolean if a field has been set.

### GetUserDataCongInfos

`func (o *AnalyticsData) GetUserDataCongInfos() []UserDataCongestionInfo`

GetUserDataCongInfos returns the UserDataCongInfos field if non-nil, zero value otherwise.

### GetUserDataCongInfosOk

`func (o *AnalyticsData) GetUserDataCongInfosOk() (*[]UserDataCongestionInfo, bool)`

GetUserDataCongInfosOk returns a tuple with the UserDataCongInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataCongInfos

`func (o *AnalyticsData) SetUserDataCongInfos(v []UserDataCongestionInfo)`

SetUserDataCongInfos sets UserDataCongInfos field to given value.

### HasUserDataCongInfos

`func (o *AnalyticsData) HasUserDataCongInfos() bool`

HasUserDataCongInfos returns a boolean if a field has been set.

### GetAbnorBehavrs

`func (o *AnalyticsData) GetAbnorBehavrs() []AbnormalBehaviour`

GetAbnorBehavrs returns the AbnorBehavrs field if non-nil, zero value otherwise.

### GetAbnorBehavrsOk

`func (o *AnalyticsData) GetAbnorBehavrsOk() (*[]AbnormalBehaviour, bool)`

GetAbnorBehavrsOk returns a tuple with the AbnorBehavrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnorBehavrs

`func (o *AnalyticsData) SetAbnorBehavrs(v []AbnormalBehaviour)`

SetAbnorBehavrs sets AbnorBehavrs field to given value.

### HasAbnorBehavrs

`func (o *AnalyticsData) HasAbnorBehavrs() bool`

HasAbnorBehavrs returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AnalyticsData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AnalyticsData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AnalyticsData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AnalyticsData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


