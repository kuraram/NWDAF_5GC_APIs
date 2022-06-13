# AnalyticsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TimeStampGen** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**AnaMetaInfo** | Pointer to [**AnalyticsMetadataInfo**](AnalyticsMetadataInfo.md) |  | [optional] 
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
**SmccExps** | Pointer to [**[]SmcceInfo**](SmcceInfo.md) |  | [optional] 
**DisperInfos** | Pointer to [**[]DispersionInfo**](DispersionInfo.md) |  | [optional] 
**RedTransInfos** | Pointer to [**[]RedundantTransmissionExpInfo**](RedundantTransmissionExpInfo.md) |  | [optional] 
**WlanInfos** | Pointer to [**[]WlanPerformanceInfo**](WlanPerformanceInfo.md) |  | [optional] 
**DnPerfInfos** | Pointer to [**[]DnPerfInfo**](DnPerfInfo.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

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

### GetStart

`func (o *AnalyticsData) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AnalyticsData) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AnalyticsData) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *AnalyticsData) HasStart() bool`

HasStart returns a boolean if a field has been set.

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

### GetAnaMetaInfo

`func (o *AnalyticsData) GetAnaMetaInfo() AnalyticsMetadataInfo`

GetAnaMetaInfo returns the AnaMetaInfo field if non-nil, zero value otherwise.

### GetAnaMetaInfoOk

`func (o *AnalyticsData) GetAnaMetaInfoOk() (*AnalyticsMetadataInfo, bool)`

GetAnaMetaInfoOk returns a tuple with the AnaMetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaMetaInfo

`func (o *AnalyticsData) SetAnaMetaInfo(v AnalyticsMetadataInfo)`

SetAnaMetaInfo sets AnaMetaInfo field to given value.

### HasAnaMetaInfo

`func (o *AnalyticsData) HasAnaMetaInfo() bool`

HasAnaMetaInfo returns a boolean if a field has been set.

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

### GetSmccExps

`func (o *AnalyticsData) GetSmccExps() []SmcceInfo`

GetSmccExps returns the SmccExps field if non-nil, zero value otherwise.

### GetSmccExpsOk

`func (o *AnalyticsData) GetSmccExpsOk() (*[]SmcceInfo, bool)`

GetSmccExpsOk returns a tuple with the SmccExps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmccExps

`func (o *AnalyticsData) SetSmccExps(v []SmcceInfo)`

SetSmccExps sets SmccExps field to given value.

### HasSmccExps

`func (o *AnalyticsData) HasSmccExps() bool`

HasSmccExps returns a boolean if a field has been set.

### GetDisperInfos

`func (o *AnalyticsData) GetDisperInfos() []DispersionInfo`

GetDisperInfos returns the DisperInfos field if non-nil, zero value otherwise.

### GetDisperInfosOk

`func (o *AnalyticsData) GetDisperInfosOk() (*[]DispersionInfo, bool)`

GetDisperInfosOk returns a tuple with the DisperInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperInfos

`func (o *AnalyticsData) SetDisperInfos(v []DispersionInfo)`

SetDisperInfos sets DisperInfos field to given value.

### HasDisperInfos

`func (o *AnalyticsData) HasDisperInfos() bool`

HasDisperInfos returns a boolean if a field has been set.

### GetRedTransInfos

`func (o *AnalyticsData) GetRedTransInfos() []RedundantTransmissionExpInfo`

GetRedTransInfos returns the RedTransInfos field if non-nil, zero value otherwise.

### GetRedTransInfosOk

`func (o *AnalyticsData) GetRedTransInfosOk() (*[]RedundantTransmissionExpInfo, bool)`

GetRedTransInfosOk returns a tuple with the RedTransInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedTransInfos

`func (o *AnalyticsData) SetRedTransInfos(v []RedundantTransmissionExpInfo)`

SetRedTransInfos sets RedTransInfos field to given value.

### HasRedTransInfos

`func (o *AnalyticsData) HasRedTransInfos() bool`

HasRedTransInfos returns a boolean if a field has been set.

### GetWlanInfos

`func (o *AnalyticsData) GetWlanInfos() []WlanPerformanceInfo`

GetWlanInfos returns the WlanInfos field if non-nil, zero value otherwise.

### GetWlanInfosOk

`func (o *AnalyticsData) GetWlanInfosOk() (*[]WlanPerformanceInfo, bool)`

GetWlanInfosOk returns a tuple with the WlanInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanInfos

`func (o *AnalyticsData) SetWlanInfos(v []WlanPerformanceInfo)`

SetWlanInfos sets WlanInfos field to given value.

### HasWlanInfos

`func (o *AnalyticsData) HasWlanInfos() bool`

HasWlanInfos returns a boolean if a field has been set.

### GetDnPerfInfos

`func (o *AnalyticsData) GetDnPerfInfos() []DnPerfInfo`

GetDnPerfInfos returns the DnPerfInfos field if non-nil, zero value otherwise.

### GetDnPerfInfosOk

`func (o *AnalyticsData) GetDnPerfInfosOk() (*[]DnPerfInfo, bool)`

GetDnPerfInfosOk returns a tuple with the DnPerfInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPerfInfos

`func (o *AnalyticsData) SetDnPerfInfos(v []DnPerfInfo)`

SetDnPerfInfos sets DnPerfInfos field to given value.

### HasDnPerfInfos

`func (o *AnalyticsData) HasDnPerfInfos() bool`

HasDnPerfInfos returns a boolean if a field has been set.

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


