# EventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnySlice** | Pointer to **bool** | FALSE represents not applicable for all slices. TRUE represents applicable for all slices. | [optional] 
**AppIds** | Pointer to **[]string** | Identification(s) of application to which the subscription applies. | [optional] 
**Dnns** | Pointer to **[]string** | Identification(s) of DNN to which the subscription applies. | [optional] 
**Dnais** | Pointer to **[]string** |  | [optional] 
**Event** | [**NwdafEvent**](NwdafEvent.md) |  | 
**ExtraReportReq** | Pointer to [**EventReportingRequirement**](EventReportingRequirement.md) |  | [optional] 
**LadnDnns** | Pointer to **[]string** | Identification(s) of LADN DNN to indicate the LADN service area as the AOI. | [optional] 
**LoadLevelThreshold** | Pointer to **int32** | Indicates that the NWDAF shall report the corresponding network slice load level to the NF service consumer where the load level of the network slice identified by snssais is reached. | [optional] 
**NotificationMethod** | Pointer to [**NotificationMethod**](NotificationMethod.md) |  | [optional] 
**MatchingDir** | Pointer to [**MatchingDirection**](MatchingDirection.md) |  | [optional] 
**NfLoadLvlThds** | Pointer to [**[]ThresholdLevel**](ThresholdLevel.md) | Shall be supplied in order to start reporting when an average load level is reached. | [optional] 
**NfInstanceIds** | Pointer to **[]string** |  | [optional] 
**NfSetIds** | Pointer to **[]string** |  | [optional] 
**NfTypes** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**NetworkArea** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**VisitedAreas** | Pointer to [**[]NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**MaxTopAppUlNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MaxTopAppDlNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**NsiIdInfos** | Pointer to [**[]NsiIdInfo**](NsiIdInfo.md) |  | [optional] 
**NsiLevelThrds** | Pointer to **[]int32** |  | [optional] 
**QosRequ** | Pointer to [**QosRequirement**](QosRequirement.md) |  | [optional] 
**QosFlowRetThds** | Pointer to [**[]RetainabilityThreshold**](RetainabilityThreshold.md) |  | [optional] 
**RanUeThrouThds** | Pointer to **[]string** |  | [optional] 
**RepetitionPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**Snssaia** | Pointer to [**[]Snssai**](Snssai.md) | Identification(s) of network slice to which the subscription applies. It corresponds to snssais in the data model definition of 3GPP TS 29.520. | [optional] 
**TgtUe** | Pointer to [**TargetUeInformation**](TargetUeInformation.md) |  | [optional] 
**CongThresholds** | Pointer to [**[]ThresholdLevel**](ThresholdLevel.md) |  | [optional] 
**NwPerfRequs** | Pointer to [**[]NetworkPerfRequirement**](NetworkPerfRequirement.md) |  | [optional] 
**BwRequs** | Pointer to [**[]BwRequirement**](BwRequirement.md) |  | [optional] 
**ExcepRequs** | Pointer to [**[]Exception**](Exception.md) |  | [optional] 
**ExptAnaType** | Pointer to [**ExpectedAnalyticsType**](ExpectedAnalyticsType.md) |  | [optional] 
**ExptUeBehav** | Pointer to [**ExpectedUeBehaviourData**](ExpectedUeBehaviourData.md) |  | [optional] 
**RatTypes** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**Freqs** | Pointer to **[]int32** |  | [optional] 
**ListOfAnaSubsets** | Pointer to [**[]AnalyticsSubset**](AnalyticsSubset.md) |  | [optional] 
**DisperReqs** | Pointer to [**[]DispersionRequirement**](DispersionRequirement.md) |  | [optional] 
**RedTransReqs** | Pointer to [**[]RedundantTransmissionExpReq**](RedundantTransmissionExpReq.md) |  | [optional] 
**WlanReqs** | Pointer to [**[]WlanPerformanceReq**](WlanPerformanceReq.md) |  | [optional] 
**UpfId** | Pointer to **string** | Identifies the UPF. | [optional] 
**AppServerAddrs** | Pointer to [**[]AddrFqdn**](AddrFqdn.md) |  | [optional] 
**DnPerfReqs** | Pointer to [**[]DnPerformanceReq**](DnPerformanceReq.md) |  | [optional] 

## Methods

### NewEventSubscription

`func NewEventSubscription(event NwdafEvent, ) *EventSubscription`

NewEventSubscription instantiates a new EventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubscriptionWithDefaults

`func NewEventSubscriptionWithDefaults() *EventSubscription`

NewEventSubscriptionWithDefaults instantiates a new EventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnySlice

`func (o *EventSubscription) GetAnySlice() bool`

GetAnySlice returns the AnySlice field if non-nil, zero value otherwise.

### GetAnySliceOk

`func (o *EventSubscription) GetAnySliceOk() (*bool, bool)`

GetAnySliceOk returns a tuple with the AnySlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnySlice

`func (o *EventSubscription) SetAnySlice(v bool)`

SetAnySlice sets AnySlice field to given value.

### HasAnySlice

`func (o *EventSubscription) HasAnySlice() bool`

HasAnySlice returns a boolean if a field has been set.

### GetAppIds

`func (o *EventSubscription) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *EventSubscription) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *EventSubscription) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *EventSubscription) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetDnns

`func (o *EventSubscription) GetDnns() []string`

GetDnns returns the Dnns field if non-nil, zero value otherwise.

### GetDnnsOk

`func (o *EventSubscription) GetDnnsOk() (*[]string, bool)`

GetDnnsOk returns a tuple with the Dnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnns

`func (o *EventSubscription) SetDnns(v []string)`

SetDnns sets Dnns field to given value.

### HasDnns

`func (o *EventSubscription) HasDnns() bool`

HasDnns returns a boolean if a field has been set.

### GetDnais

`func (o *EventSubscription) GetDnais() []string`

GetDnais returns the Dnais field if non-nil, zero value otherwise.

### GetDnaisOk

`func (o *EventSubscription) GetDnaisOk() (*[]string, bool)`

GetDnaisOk returns a tuple with the Dnais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnais

`func (o *EventSubscription) SetDnais(v []string)`

SetDnais sets Dnais field to given value.

### HasDnais

`func (o *EventSubscription) HasDnais() bool`

HasDnais returns a boolean if a field has been set.

### GetEvent

`func (o *EventSubscription) GetEvent() NwdafEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventSubscription) GetEventOk() (*NwdafEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventSubscription) SetEvent(v NwdafEvent)`

SetEvent sets Event field to given value.


### GetExtraReportReq

`func (o *EventSubscription) GetExtraReportReq() EventReportingRequirement`

GetExtraReportReq returns the ExtraReportReq field if non-nil, zero value otherwise.

### GetExtraReportReqOk

`func (o *EventSubscription) GetExtraReportReqOk() (*EventReportingRequirement, bool)`

GetExtraReportReqOk returns a tuple with the ExtraReportReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraReportReq

`func (o *EventSubscription) SetExtraReportReq(v EventReportingRequirement)`

SetExtraReportReq sets ExtraReportReq field to given value.

### HasExtraReportReq

`func (o *EventSubscription) HasExtraReportReq() bool`

HasExtraReportReq returns a boolean if a field has been set.

### GetLadnDnns

`func (o *EventSubscription) GetLadnDnns() []string`

GetLadnDnns returns the LadnDnns field if non-nil, zero value otherwise.

### GetLadnDnnsOk

`func (o *EventSubscription) GetLadnDnnsOk() (*[]string, bool)`

GetLadnDnnsOk returns a tuple with the LadnDnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLadnDnns

`func (o *EventSubscription) SetLadnDnns(v []string)`

SetLadnDnns sets LadnDnns field to given value.

### HasLadnDnns

`func (o *EventSubscription) HasLadnDnns() bool`

HasLadnDnns returns a boolean if a field has been set.

### GetLoadLevelThreshold

`func (o *EventSubscription) GetLoadLevelThreshold() int32`

GetLoadLevelThreshold returns the LoadLevelThreshold field if non-nil, zero value otherwise.

### GetLoadLevelThresholdOk

`func (o *EventSubscription) GetLoadLevelThresholdOk() (*int32, bool)`

GetLoadLevelThresholdOk returns a tuple with the LoadLevelThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadLevelThreshold

`func (o *EventSubscription) SetLoadLevelThreshold(v int32)`

SetLoadLevelThreshold sets LoadLevelThreshold field to given value.

### HasLoadLevelThreshold

`func (o *EventSubscription) HasLoadLevelThreshold() bool`

HasLoadLevelThreshold returns a boolean if a field has been set.

### GetNotificationMethod

`func (o *EventSubscription) GetNotificationMethod() NotificationMethod`

GetNotificationMethod returns the NotificationMethod field if non-nil, zero value otherwise.

### GetNotificationMethodOk

`func (o *EventSubscription) GetNotificationMethodOk() (*NotificationMethod, bool)`

GetNotificationMethodOk returns a tuple with the NotificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMethod

`func (o *EventSubscription) SetNotificationMethod(v NotificationMethod)`

SetNotificationMethod sets NotificationMethod field to given value.

### HasNotificationMethod

`func (o *EventSubscription) HasNotificationMethod() bool`

HasNotificationMethod returns a boolean if a field has been set.

### GetMatchingDir

`func (o *EventSubscription) GetMatchingDir() MatchingDirection`

GetMatchingDir returns the MatchingDir field if non-nil, zero value otherwise.

### GetMatchingDirOk

`func (o *EventSubscription) GetMatchingDirOk() (*MatchingDirection, bool)`

GetMatchingDirOk returns a tuple with the MatchingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingDir

`func (o *EventSubscription) SetMatchingDir(v MatchingDirection)`

SetMatchingDir sets MatchingDir field to given value.

### HasMatchingDir

`func (o *EventSubscription) HasMatchingDir() bool`

HasMatchingDir returns a boolean if a field has been set.

### GetNfLoadLvlThds

`func (o *EventSubscription) GetNfLoadLvlThds() []ThresholdLevel`

GetNfLoadLvlThds returns the NfLoadLvlThds field if non-nil, zero value otherwise.

### GetNfLoadLvlThdsOk

`func (o *EventSubscription) GetNfLoadLvlThdsOk() (*[]ThresholdLevel, bool)`

GetNfLoadLvlThdsOk returns a tuple with the NfLoadLvlThds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfLoadLvlThds

`func (o *EventSubscription) SetNfLoadLvlThds(v []ThresholdLevel)`

SetNfLoadLvlThds sets NfLoadLvlThds field to given value.

### HasNfLoadLvlThds

`func (o *EventSubscription) HasNfLoadLvlThds() bool`

HasNfLoadLvlThds returns a boolean if a field has been set.

### GetNfInstanceIds

`func (o *EventSubscription) GetNfInstanceIds() []string`

GetNfInstanceIds returns the NfInstanceIds field if non-nil, zero value otherwise.

### GetNfInstanceIdsOk

`func (o *EventSubscription) GetNfInstanceIdsOk() (*[]string, bool)`

GetNfInstanceIdsOk returns a tuple with the NfInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceIds

`func (o *EventSubscription) SetNfInstanceIds(v []string)`

SetNfInstanceIds sets NfInstanceIds field to given value.

### HasNfInstanceIds

`func (o *EventSubscription) HasNfInstanceIds() bool`

HasNfInstanceIds returns a boolean if a field has been set.

### GetNfSetIds

`func (o *EventSubscription) GetNfSetIds() []string`

GetNfSetIds returns the NfSetIds field if non-nil, zero value otherwise.

### GetNfSetIdsOk

`func (o *EventSubscription) GetNfSetIdsOk() (*[]string, bool)`

GetNfSetIdsOk returns a tuple with the NfSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetIds

`func (o *EventSubscription) SetNfSetIds(v []string)`

SetNfSetIds sets NfSetIds field to given value.

### HasNfSetIds

`func (o *EventSubscription) HasNfSetIds() bool`

HasNfSetIds returns a boolean if a field has been set.

### GetNfTypes

`func (o *EventSubscription) GetNfTypes() []NFType`

GetNfTypes returns the NfTypes field if non-nil, zero value otherwise.

### GetNfTypesOk

`func (o *EventSubscription) GetNfTypesOk() (*[]NFType, bool)`

GetNfTypesOk returns a tuple with the NfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfTypes

`func (o *EventSubscription) SetNfTypes(v []NFType)`

SetNfTypes sets NfTypes field to given value.

### HasNfTypes

`func (o *EventSubscription) HasNfTypes() bool`

HasNfTypes returns a boolean if a field has been set.

### GetNetworkArea

`func (o *EventSubscription) GetNetworkArea() NetworkAreaInfo`

GetNetworkArea returns the NetworkArea field if non-nil, zero value otherwise.

### GetNetworkAreaOk

`func (o *EventSubscription) GetNetworkAreaOk() (*NetworkAreaInfo, bool)`

GetNetworkAreaOk returns a tuple with the NetworkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkArea

`func (o *EventSubscription) SetNetworkArea(v NetworkAreaInfo)`

SetNetworkArea sets NetworkArea field to given value.

### HasNetworkArea

`func (o *EventSubscription) HasNetworkArea() bool`

HasNetworkArea returns a boolean if a field has been set.

### GetVisitedAreas

`func (o *EventSubscription) GetVisitedAreas() []NetworkAreaInfo`

GetVisitedAreas returns the VisitedAreas field if non-nil, zero value otherwise.

### GetVisitedAreasOk

`func (o *EventSubscription) GetVisitedAreasOk() (*[]NetworkAreaInfo, bool)`

GetVisitedAreasOk returns a tuple with the VisitedAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitedAreas

`func (o *EventSubscription) SetVisitedAreas(v []NetworkAreaInfo)`

SetVisitedAreas sets VisitedAreas field to given value.

### HasVisitedAreas

`func (o *EventSubscription) HasVisitedAreas() bool`

HasVisitedAreas returns a boolean if a field has been set.

### GetMaxTopAppUlNbr

`func (o *EventSubscription) GetMaxTopAppUlNbr() int32`

GetMaxTopAppUlNbr returns the MaxTopAppUlNbr field if non-nil, zero value otherwise.

### GetMaxTopAppUlNbrOk

`func (o *EventSubscription) GetMaxTopAppUlNbrOk() (*int32, bool)`

GetMaxTopAppUlNbrOk returns a tuple with the MaxTopAppUlNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTopAppUlNbr

`func (o *EventSubscription) SetMaxTopAppUlNbr(v int32)`

SetMaxTopAppUlNbr sets MaxTopAppUlNbr field to given value.

### HasMaxTopAppUlNbr

`func (o *EventSubscription) HasMaxTopAppUlNbr() bool`

HasMaxTopAppUlNbr returns a boolean if a field has been set.

### GetMaxTopAppDlNbr

`func (o *EventSubscription) GetMaxTopAppDlNbr() int32`

GetMaxTopAppDlNbr returns the MaxTopAppDlNbr field if non-nil, zero value otherwise.

### GetMaxTopAppDlNbrOk

`func (o *EventSubscription) GetMaxTopAppDlNbrOk() (*int32, bool)`

GetMaxTopAppDlNbrOk returns a tuple with the MaxTopAppDlNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTopAppDlNbr

`func (o *EventSubscription) SetMaxTopAppDlNbr(v int32)`

SetMaxTopAppDlNbr sets MaxTopAppDlNbr field to given value.

### HasMaxTopAppDlNbr

`func (o *EventSubscription) HasMaxTopAppDlNbr() bool`

HasMaxTopAppDlNbr returns a boolean if a field has been set.

### GetNsiIdInfos

`func (o *EventSubscription) GetNsiIdInfos() []NsiIdInfo`

GetNsiIdInfos returns the NsiIdInfos field if non-nil, zero value otherwise.

### GetNsiIdInfosOk

`func (o *EventSubscription) GetNsiIdInfosOk() (*[]NsiIdInfo, bool)`

GetNsiIdInfosOk returns a tuple with the NsiIdInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiIdInfos

`func (o *EventSubscription) SetNsiIdInfos(v []NsiIdInfo)`

SetNsiIdInfos sets NsiIdInfos field to given value.

### HasNsiIdInfos

`func (o *EventSubscription) HasNsiIdInfos() bool`

HasNsiIdInfos returns a boolean if a field has been set.

### GetNsiLevelThrds

`func (o *EventSubscription) GetNsiLevelThrds() []int32`

GetNsiLevelThrds returns the NsiLevelThrds field if non-nil, zero value otherwise.

### GetNsiLevelThrdsOk

`func (o *EventSubscription) GetNsiLevelThrdsOk() (*[]int32, bool)`

GetNsiLevelThrdsOk returns a tuple with the NsiLevelThrds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiLevelThrds

`func (o *EventSubscription) SetNsiLevelThrds(v []int32)`

SetNsiLevelThrds sets NsiLevelThrds field to given value.

### HasNsiLevelThrds

`func (o *EventSubscription) HasNsiLevelThrds() bool`

HasNsiLevelThrds returns a boolean if a field has been set.

### GetQosRequ

`func (o *EventSubscription) GetQosRequ() QosRequirement`

GetQosRequ returns the QosRequ field if non-nil, zero value otherwise.

### GetQosRequOk

`func (o *EventSubscription) GetQosRequOk() (*QosRequirement, bool)`

GetQosRequOk returns a tuple with the QosRequ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosRequ

`func (o *EventSubscription) SetQosRequ(v QosRequirement)`

SetQosRequ sets QosRequ field to given value.

### HasQosRequ

`func (o *EventSubscription) HasQosRequ() bool`

HasQosRequ returns a boolean if a field has been set.

### GetQosFlowRetThds

`func (o *EventSubscription) GetQosFlowRetThds() []RetainabilityThreshold`

GetQosFlowRetThds returns the QosFlowRetThds field if non-nil, zero value otherwise.

### GetQosFlowRetThdsOk

`func (o *EventSubscription) GetQosFlowRetThdsOk() (*[]RetainabilityThreshold, bool)`

GetQosFlowRetThdsOk returns a tuple with the QosFlowRetThds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowRetThds

`func (o *EventSubscription) SetQosFlowRetThds(v []RetainabilityThreshold)`

SetQosFlowRetThds sets QosFlowRetThds field to given value.

### HasQosFlowRetThds

`func (o *EventSubscription) HasQosFlowRetThds() bool`

HasQosFlowRetThds returns a boolean if a field has been set.

### GetRanUeThrouThds

`func (o *EventSubscription) GetRanUeThrouThds() []string`

GetRanUeThrouThds returns the RanUeThrouThds field if non-nil, zero value otherwise.

### GetRanUeThrouThdsOk

`func (o *EventSubscription) GetRanUeThrouThdsOk() (*[]string, bool)`

GetRanUeThrouThdsOk returns a tuple with the RanUeThrouThds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanUeThrouThds

`func (o *EventSubscription) SetRanUeThrouThds(v []string)`

SetRanUeThrouThds sets RanUeThrouThds field to given value.

### HasRanUeThrouThds

`func (o *EventSubscription) HasRanUeThrouThds() bool`

HasRanUeThrouThds returns a boolean if a field has been set.

### GetRepetitionPeriod

`func (o *EventSubscription) GetRepetitionPeriod() int32`

GetRepetitionPeriod returns the RepetitionPeriod field if non-nil, zero value otherwise.

### GetRepetitionPeriodOk

`func (o *EventSubscription) GetRepetitionPeriodOk() (*int32, bool)`

GetRepetitionPeriodOk returns a tuple with the RepetitionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionPeriod

`func (o *EventSubscription) SetRepetitionPeriod(v int32)`

SetRepetitionPeriod sets RepetitionPeriod field to given value.

### HasRepetitionPeriod

`func (o *EventSubscription) HasRepetitionPeriod() bool`

HasRepetitionPeriod returns a boolean if a field has been set.

### GetSnssaia

`func (o *EventSubscription) GetSnssaia() []Snssai`

GetSnssaia returns the Snssaia field if non-nil, zero value otherwise.

### GetSnssaiaOk

`func (o *EventSubscription) GetSnssaiaOk() (*[]Snssai, bool)`

GetSnssaiaOk returns a tuple with the Snssaia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaia

`func (o *EventSubscription) SetSnssaia(v []Snssai)`

SetSnssaia sets Snssaia field to given value.

### HasSnssaia

`func (o *EventSubscription) HasSnssaia() bool`

HasSnssaia returns a boolean if a field has been set.

### GetTgtUe

`func (o *EventSubscription) GetTgtUe() TargetUeInformation`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *EventSubscription) GetTgtUeOk() (*TargetUeInformation, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *EventSubscription) SetTgtUe(v TargetUeInformation)`

SetTgtUe sets TgtUe field to given value.

### HasTgtUe

`func (o *EventSubscription) HasTgtUe() bool`

HasTgtUe returns a boolean if a field has been set.

### GetCongThresholds

`func (o *EventSubscription) GetCongThresholds() []ThresholdLevel`

GetCongThresholds returns the CongThresholds field if non-nil, zero value otherwise.

### GetCongThresholdsOk

`func (o *EventSubscription) GetCongThresholdsOk() (*[]ThresholdLevel, bool)`

GetCongThresholdsOk returns a tuple with the CongThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongThresholds

`func (o *EventSubscription) SetCongThresholds(v []ThresholdLevel)`

SetCongThresholds sets CongThresholds field to given value.

### HasCongThresholds

`func (o *EventSubscription) HasCongThresholds() bool`

HasCongThresholds returns a boolean if a field has been set.

### GetNwPerfRequs

`func (o *EventSubscription) GetNwPerfRequs() []NetworkPerfRequirement`

GetNwPerfRequs returns the NwPerfRequs field if non-nil, zero value otherwise.

### GetNwPerfRequsOk

`func (o *EventSubscription) GetNwPerfRequsOk() (*[]NetworkPerfRequirement, bool)`

GetNwPerfRequsOk returns a tuple with the NwPerfRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfRequs

`func (o *EventSubscription) SetNwPerfRequs(v []NetworkPerfRequirement)`

SetNwPerfRequs sets NwPerfRequs field to given value.

### HasNwPerfRequs

`func (o *EventSubscription) HasNwPerfRequs() bool`

HasNwPerfRequs returns a boolean if a field has been set.

### GetBwRequs

`func (o *EventSubscription) GetBwRequs() []BwRequirement`

GetBwRequs returns the BwRequs field if non-nil, zero value otherwise.

### GetBwRequsOk

`func (o *EventSubscription) GetBwRequsOk() (*[]BwRequirement, bool)`

GetBwRequsOk returns a tuple with the BwRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwRequs

`func (o *EventSubscription) SetBwRequs(v []BwRequirement)`

SetBwRequs sets BwRequs field to given value.

### HasBwRequs

`func (o *EventSubscription) HasBwRequs() bool`

HasBwRequs returns a boolean if a field has been set.

### GetExcepRequs

`func (o *EventSubscription) GetExcepRequs() []Exception`

GetExcepRequs returns the ExcepRequs field if non-nil, zero value otherwise.

### GetExcepRequsOk

`func (o *EventSubscription) GetExcepRequsOk() (*[]Exception, bool)`

GetExcepRequsOk returns a tuple with the ExcepRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepRequs

`func (o *EventSubscription) SetExcepRequs(v []Exception)`

SetExcepRequs sets ExcepRequs field to given value.

### HasExcepRequs

`func (o *EventSubscription) HasExcepRequs() bool`

HasExcepRequs returns a boolean if a field has been set.

### GetExptAnaType

`func (o *EventSubscription) GetExptAnaType() ExpectedAnalyticsType`

GetExptAnaType returns the ExptAnaType field if non-nil, zero value otherwise.

### GetExptAnaTypeOk

`func (o *EventSubscription) GetExptAnaTypeOk() (*ExpectedAnalyticsType, bool)`

GetExptAnaTypeOk returns a tuple with the ExptAnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptAnaType

`func (o *EventSubscription) SetExptAnaType(v ExpectedAnalyticsType)`

SetExptAnaType sets ExptAnaType field to given value.

### HasExptAnaType

`func (o *EventSubscription) HasExptAnaType() bool`

HasExptAnaType returns a boolean if a field has been set.

### GetExptUeBehav

`func (o *EventSubscription) GetExptUeBehav() ExpectedUeBehaviourData`

GetExptUeBehav returns the ExptUeBehav field if non-nil, zero value otherwise.

### GetExptUeBehavOk

`func (o *EventSubscription) GetExptUeBehavOk() (*ExpectedUeBehaviourData, bool)`

GetExptUeBehavOk returns a tuple with the ExptUeBehav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptUeBehav

`func (o *EventSubscription) SetExptUeBehav(v ExpectedUeBehaviourData)`

SetExptUeBehav sets ExptUeBehav field to given value.

### HasExptUeBehav

`func (o *EventSubscription) HasExptUeBehav() bool`

HasExptUeBehav returns a boolean if a field has been set.

### GetRatTypes

`func (o *EventSubscription) GetRatTypes() []RatType`

GetRatTypes returns the RatTypes field if non-nil, zero value otherwise.

### GetRatTypesOk

`func (o *EventSubscription) GetRatTypesOk() (*[]RatType, bool)`

GetRatTypesOk returns a tuple with the RatTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatTypes

`func (o *EventSubscription) SetRatTypes(v []RatType)`

SetRatTypes sets RatTypes field to given value.

### HasRatTypes

`func (o *EventSubscription) HasRatTypes() bool`

HasRatTypes returns a boolean if a field has been set.

### GetFreqs

`func (o *EventSubscription) GetFreqs() []int32`

GetFreqs returns the Freqs field if non-nil, zero value otherwise.

### GetFreqsOk

`func (o *EventSubscription) GetFreqsOk() (*[]int32, bool)`

GetFreqsOk returns a tuple with the Freqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreqs

`func (o *EventSubscription) SetFreqs(v []int32)`

SetFreqs sets Freqs field to given value.

### HasFreqs

`func (o *EventSubscription) HasFreqs() bool`

HasFreqs returns a boolean if a field has been set.

### GetListOfAnaSubsets

`func (o *EventSubscription) GetListOfAnaSubsets() []AnalyticsSubset`

GetListOfAnaSubsets returns the ListOfAnaSubsets field if non-nil, zero value otherwise.

### GetListOfAnaSubsetsOk

`func (o *EventSubscription) GetListOfAnaSubsetsOk() (*[]AnalyticsSubset, bool)`

GetListOfAnaSubsetsOk returns a tuple with the ListOfAnaSubsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfAnaSubsets

`func (o *EventSubscription) SetListOfAnaSubsets(v []AnalyticsSubset)`

SetListOfAnaSubsets sets ListOfAnaSubsets field to given value.

### HasListOfAnaSubsets

`func (o *EventSubscription) HasListOfAnaSubsets() bool`

HasListOfAnaSubsets returns a boolean if a field has been set.

### GetDisperReqs

`func (o *EventSubscription) GetDisperReqs() []DispersionRequirement`

GetDisperReqs returns the DisperReqs field if non-nil, zero value otherwise.

### GetDisperReqsOk

`func (o *EventSubscription) GetDisperReqsOk() (*[]DispersionRequirement, bool)`

GetDisperReqsOk returns a tuple with the DisperReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperReqs

`func (o *EventSubscription) SetDisperReqs(v []DispersionRequirement)`

SetDisperReqs sets DisperReqs field to given value.

### HasDisperReqs

`func (o *EventSubscription) HasDisperReqs() bool`

HasDisperReqs returns a boolean if a field has been set.

### GetRedTransReqs

`func (o *EventSubscription) GetRedTransReqs() []RedundantTransmissionExpReq`

GetRedTransReqs returns the RedTransReqs field if non-nil, zero value otherwise.

### GetRedTransReqsOk

`func (o *EventSubscription) GetRedTransReqsOk() (*[]RedundantTransmissionExpReq, bool)`

GetRedTransReqsOk returns a tuple with the RedTransReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedTransReqs

`func (o *EventSubscription) SetRedTransReqs(v []RedundantTransmissionExpReq)`

SetRedTransReqs sets RedTransReqs field to given value.

### HasRedTransReqs

`func (o *EventSubscription) HasRedTransReqs() bool`

HasRedTransReqs returns a boolean if a field has been set.

### GetWlanReqs

`func (o *EventSubscription) GetWlanReqs() []WlanPerformanceReq`

GetWlanReqs returns the WlanReqs field if non-nil, zero value otherwise.

### GetWlanReqsOk

`func (o *EventSubscription) GetWlanReqsOk() (*[]WlanPerformanceReq, bool)`

GetWlanReqsOk returns a tuple with the WlanReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanReqs

`func (o *EventSubscription) SetWlanReqs(v []WlanPerformanceReq)`

SetWlanReqs sets WlanReqs field to given value.

### HasWlanReqs

`func (o *EventSubscription) HasWlanReqs() bool`

HasWlanReqs returns a boolean if a field has been set.

### GetUpfId

`func (o *EventSubscription) GetUpfId() string`

GetUpfId returns the UpfId field if non-nil, zero value otherwise.

### GetUpfIdOk

`func (o *EventSubscription) GetUpfIdOk() (*string, bool)`

GetUpfIdOk returns a tuple with the UpfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfId

`func (o *EventSubscription) SetUpfId(v string)`

SetUpfId sets UpfId field to given value.

### HasUpfId

`func (o *EventSubscription) HasUpfId() bool`

HasUpfId returns a boolean if a field has been set.

### GetAppServerAddrs

`func (o *EventSubscription) GetAppServerAddrs() []AddrFqdn`

GetAppServerAddrs returns the AppServerAddrs field if non-nil, zero value otherwise.

### GetAppServerAddrsOk

`func (o *EventSubscription) GetAppServerAddrsOk() (*[]AddrFqdn, bool)`

GetAppServerAddrsOk returns a tuple with the AppServerAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerAddrs

`func (o *EventSubscription) SetAppServerAddrs(v []AddrFqdn)`

SetAppServerAddrs sets AppServerAddrs field to given value.

### HasAppServerAddrs

`func (o *EventSubscription) HasAppServerAddrs() bool`

HasAppServerAddrs returns a boolean if a field has been set.

### GetDnPerfReqs

`func (o *EventSubscription) GetDnPerfReqs() []DnPerformanceReq`

GetDnPerfReqs returns the DnPerfReqs field if non-nil, zero value otherwise.

### GetDnPerfReqsOk

`func (o *EventSubscription) GetDnPerfReqsOk() (*[]DnPerformanceReq, bool)`

GetDnPerfReqsOk returns a tuple with the DnPerfReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPerfReqs

`func (o *EventSubscription) SetDnPerfReqs(v []DnPerformanceReq)`

SetDnPerfReqs sets DnPerfReqs field to given value.

### HasDnPerfReqs

`func (o *EventSubscription) HasDnPerfReqs() bool`

HasDnPerfReqs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


