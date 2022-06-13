# EventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnySlice** | Pointer to **bool** | FALSE represents not applicable for all slices. TRUE represents applicable for all slices. | [optional] 
**Snssais** | Pointer to [**[]Snssai**](Snssai.md) | Identification(s) of network slice to which the subscription belongs. | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**Dnns** | Pointer to **[]string** |  | [optional] 
**Dnais** | Pointer to **[]string** |  | [optional] 
**LadnDnns** | Pointer to **[]string** | Identification(s) of LADN DNN to indicate the LADN service area as the AOI. | [optional] 
**NetworkArea** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**VisitedAreas** | Pointer to [**[]NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**MaxTopAppUlNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MaxTopAppDlNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**NfInstanceIds** | Pointer to **[]string** |  | [optional] 
**NfSetIds** | Pointer to **[]string** |  | [optional] 
**NfTypes** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**NsiIdInfos** | Pointer to [**[]NsiIdInfo**](NsiIdInfo.md) |  | [optional] 
**QosRequ** | Pointer to [**QosRequirement**](QosRequirement.md) |  | [optional] 
**NwPerfTypes** | Pointer to [**[]NetworkPerfType**](NetworkPerfType.md) |  | [optional] 
**BwRequs** | Pointer to [**[]BwRequirement**](BwRequirement.md) |  | [optional] 
**ExcepIds** | Pointer to [**[]ExceptionId**](ExceptionId.md) |  | [optional] 
**ExptAnaType** | Pointer to [**ExpectedAnalyticsType**](ExpectedAnalyticsType.md) |  | [optional] 
**ExptUeBehav** | Pointer to [**ExpectedUeBehaviourData**](ExpectedUeBehaviourData.md) |  | [optional] 
**RatTypes** | Pointer to [**[]RatType**](RatType.md) |  | [optional] 
**Freqs** | Pointer to **[]int32** |  | [optional] 
**DisperReqs** | Pointer to [**[]DispersionRequirement**](DispersionRequirement.md) |  | [optional] 
**RedTransReqs** | Pointer to [**[]RedundantTransmissionExpReq**](RedundantTransmissionExpReq.md) |  | [optional] 
**WlanReqs** | Pointer to [**[]WlanPerformanceReq**](WlanPerformanceReq.md) |  | [optional] 
**ListOfAnaSubsets** | Pointer to [**[]AnalyticsSubset**](AnalyticsSubset.md) |  | [optional] 
**UpfId** | Pointer to **string** | Identifies the UPF. | [optional] 
**AppServerAddrs** | Pointer to [**[]AddrFqdn**](AddrFqdn.md) |  | [optional] 
**DnPerfReqs** | Pointer to [**[]DnPerformanceReq**](DnPerformanceReq.md) |  | [optional] 

## Methods

### NewEventFilter

`func NewEventFilter() *EventFilter`

NewEventFilter instantiates a new EventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventFilterWithDefaults

`func NewEventFilterWithDefaults() *EventFilter`

NewEventFilterWithDefaults instantiates a new EventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnySlice

`func (o *EventFilter) GetAnySlice() bool`

GetAnySlice returns the AnySlice field if non-nil, zero value otherwise.

### GetAnySliceOk

`func (o *EventFilter) GetAnySliceOk() (*bool, bool)`

GetAnySliceOk returns a tuple with the AnySlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnySlice

`func (o *EventFilter) SetAnySlice(v bool)`

SetAnySlice sets AnySlice field to given value.

### HasAnySlice

`func (o *EventFilter) HasAnySlice() bool`

HasAnySlice returns a boolean if a field has been set.

### GetSnssais

`func (o *EventFilter) GetSnssais() []Snssai`

GetSnssais returns the Snssais field if non-nil, zero value otherwise.

### GetSnssaisOk

`func (o *EventFilter) GetSnssaisOk() (*[]Snssai, bool)`

GetSnssaisOk returns a tuple with the Snssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssais

`func (o *EventFilter) SetSnssais(v []Snssai)`

SetSnssais sets Snssais field to given value.

### HasSnssais

`func (o *EventFilter) HasSnssais() bool`

HasSnssais returns a boolean if a field has been set.

### GetAppIds

`func (o *EventFilter) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *EventFilter) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *EventFilter) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *EventFilter) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetDnns

`func (o *EventFilter) GetDnns() []string`

GetDnns returns the Dnns field if non-nil, zero value otherwise.

### GetDnnsOk

`func (o *EventFilter) GetDnnsOk() (*[]string, bool)`

GetDnnsOk returns a tuple with the Dnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnns

`func (o *EventFilter) SetDnns(v []string)`

SetDnns sets Dnns field to given value.

### HasDnns

`func (o *EventFilter) HasDnns() bool`

HasDnns returns a boolean if a field has been set.

### GetDnais

`func (o *EventFilter) GetDnais() []string`

GetDnais returns the Dnais field if non-nil, zero value otherwise.

### GetDnaisOk

`func (o *EventFilter) GetDnaisOk() (*[]string, bool)`

GetDnaisOk returns a tuple with the Dnais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnais

`func (o *EventFilter) SetDnais(v []string)`

SetDnais sets Dnais field to given value.

### HasDnais

`func (o *EventFilter) HasDnais() bool`

HasDnais returns a boolean if a field has been set.

### GetLadnDnns

`func (o *EventFilter) GetLadnDnns() []string`

GetLadnDnns returns the LadnDnns field if non-nil, zero value otherwise.

### GetLadnDnnsOk

`func (o *EventFilter) GetLadnDnnsOk() (*[]string, bool)`

GetLadnDnnsOk returns a tuple with the LadnDnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLadnDnns

`func (o *EventFilter) SetLadnDnns(v []string)`

SetLadnDnns sets LadnDnns field to given value.

### HasLadnDnns

`func (o *EventFilter) HasLadnDnns() bool`

HasLadnDnns returns a boolean if a field has been set.

### GetNetworkArea

`func (o *EventFilter) GetNetworkArea() NetworkAreaInfo`

GetNetworkArea returns the NetworkArea field if non-nil, zero value otherwise.

### GetNetworkAreaOk

`func (o *EventFilter) GetNetworkAreaOk() (*NetworkAreaInfo, bool)`

GetNetworkAreaOk returns a tuple with the NetworkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkArea

`func (o *EventFilter) SetNetworkArea(v NetworkAreaInfo)`

SetNetworkArea sets NetworkArea field to given value.

### HasNetworkArea

`func (o *EventFilter) HasNetworkArea() bool`

HasNetworkArea returns a boolean if a field has been set.

### GetVisitedAreas

`func (o *EventFilter) GetVisitedAreas() []NetworkAreaInfo`

GetVisitedAreas returns the VisitedAreas field if non-nil, zero value otherwise.

### GetVisitedAreasOk

`func (o *EventFilter) GetVisitedAreasOk() (*[]NetworkAreaInfo, bool)`

GetVisitedAreasOk returns a tuple with the VisitedAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitedAreas

`func (o *EventFilter) SetVisitedAreas(v []NetworkAreaInfo)`

SetVisitedAreas sets VisitedAreas field to given value.

### HasVisitedAreas

`func (o *EventFilter) HasVisitedAreas() bool`

HasVisitedAreas returns a boolean if a field has been set.

### GetMaxTopAppUlNbr

`func (o *EventFilter) GetMaxTopAppUlNbr() int32`

GetMaxTopAppUlNbr returns the MaxTopAppUlNbr field if non-nil, zero value otherwise.

### GetMaxTopAppUlNbrOk

`func (o *EventFilter) GetMaxTopAppUlNbrOk() (*int32, bool)`

GetMaxTopAppUlNbrOk returns a tuple with the MaxTopAppUlNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTopAppUlNbr

`func (o *EventFilter) SetMaxTopAppUlNbr(v int32)`

SetMaxTopAppUlNbr sets MaxTopAppUlNbr field to given value.

### HasMaxTopAppUlNbr

`func (o *EventFilter) HasMaxTopAppUlNbr() bool`

HasMaxTopAppUlNbr returns a boolean if a field has been set.

### GetMaxTopAppDlNbr

`func (o *EventFilter) GetMaxTopAppDlNbr() int32`

GetMaxTopAppDlNbr returns the MaxTopAppDlNbr field if non-nil, zero value otherwise.

### GetMaxTopAppDlNbrOk

`func (o *EventFilter) GetMaxTopAppDlNbrOk() (*int32, bool)`

GetMaxTopAppDlNbrOk returns a tuple with the MaxTopAppDlNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTopAppDlNbr

`func (o *EventFilter) SetMaxTopAppDlNbr(v int32)`

SetMaxTopAppDlNbr sets MaxTopAppDlNbr field to given value.

### HasMaxTopAppDlNbr

`func (o *EventFilter) HasMaxTopAppDlNbr() bool`

HasMaxTopAppDlNbr returns a boolean if a field has been set.

### GetNfInstanceIds

`func (o *EventFilter) GetNfInstanceIds() []string`

GetNfInstanceIds returns the NfInstanceIds field if non-nil, zero value otherwise.

### GetNfInstanceIdsOk

`func (o *EventFilter) GetNfInstanceIdsOk() (*[]string, bool)`

GetNfInstanceIdsOk returns a tuple with the NfInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceIds

`func (o *EventFilter) SetNfInstanceIds(v []string)`

SetNfInstanceIds sets NfInstanceIds field to given value.

### HasNfInstanceIds

`func (o *EventFilter) HasNfInstanceIds() bool`

HasNfInstanceIds returns a boolean if a field has been set.

### GetNfSetIds

`func (o *EventFilter) GetNfSetIds() []string`

GetNfSetIds returns the NfSetIds field if non-nil, zero value otherwise.

### GetNfSetIdsOk

`func (o *EventFilter) GetNfSetIdsOk() (*[]string, bool)`

GetNfSetIdsOk returns a tuple with the NfSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetIds

`func (o *EventFilter) SetNfSetIds(v []string)`

SetNfSetIds sets NfSetIds field to given value.

### HasNfSetIds

`func (o *EventFilter) HasNfSetIds() bool`

HasNfSetIds returns a boolean if a field has been set.

### GetNfTypes

`func (o *EventFilter) GetNfTypes() []NFType`

GetNfTypes returns the NfTypes field if non-nil, zero value otherwise.

### GetNfTypesOk

`func (o *EventFilter) GetNfTypesOk() (*[]NFType, bool)`

GetNfTypesOk returns a tuple with the NfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfTypes

`func (o *EventFilter) SetNfTypes(v []NFType)`

SetNfTypes sets NfTypes field to given value.

### HasNfTypes

`func (o *EventFilter) HasNfTypes() bool`

HasNfTypes returns a boolean if a field has been set.

### GetNsiIdInfos

`func (o *EventFilter) GetNsiIdInfos() []NsiIdInfo`

GetNsiIdInfos returns the NsiIdInfos field if non-nil, zero value otherwise.

### GetNsiIdInfosOk

`func (o *EventFilter) GetNsiIdInfosOk() (*[]NsiIdInfo, bool)`

GetNsiIdInfosOk returns a tuple with the NsiIdInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiIdInfos

`func (o *EventFilter) SetNsiIdInfos(v []NsiIdInfo)`

SetNsiIdInfos sets NsiIdInfos field to given value.

### HasNsiIdInfos

`func (o *EventFilter) HasNsiIdInfos() bool`

HasNsiIdInfos returns a boolean if a field has been set.

### GetQosRequ

`func (o *EventFilter) GetQosRequ() QosRequirement`

GetQosRequ returns the QosRequ field if non-nil, zero value otherwise.

### GetQosRequOk

`func (o *EventFilter) GetQosRequOk() (*QosRequirement, bool)`

GetQosRequOk returns a tuple with the QosRequ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosRequ

`func (o *EventFilter) SetQosRequ(v QosRequirement)`

SetQosRequ sets QosRequ field to given value.

### HasQosRequ

`func (o *EventFilter) HasQosRequ() bool`

HasQosRequ returns a boolean if a field has been set.

### GetNwPerfTypes

`func (o *EventFilter) GetNwPerfTypes() []NetworkPerfType`

GetNwPerfTypes returns the NwPerfTypes field if non-nil, zero value otherwise.

### GetNwPerfTypesOk

`func (o *EventFilter) GetNwPerfTypesOk() (*[]NetworkPerfType, bool)`

GetNwPerfTypesOk returns a tuple with the NwPerfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwPerfTypes

`func (o *EventFilter) SetNwPerfTypes(v []NetworkPerfType)`

SetNwPerfTypes sets NwPerfTypes field to given value.

### HasNwPerfTypes

`func (o *EventFilter) HasNwPerfTypes() bool`

HasNwPerfTypes returns a boolean if a field has been set.

### GetBwRequs

`func (o *EventFilter) GetBwRequs() []BwRequirement`

GetBwRequs returns the BwRequs field if non-nil, zero value otherwise.

### GetBwRequsOk

`func (o *EventFilter) GetBwRequsOk() (*[]BwRequirement, bool)`

GetBwRequsOk returns a tuple with the BwRequs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwRequs

`func (o *EventFilter) SetBwRequs(v []BwRequirement)`

SetBwRequs sets BwRequs field to given value.

### HasBwRequs

`func (o *EventFilter) HasBwRequs() bool`

HasBwRequs returns a boolean if a field has been set.

### GetExcepIds

`func (o *EventFilter) GetExcepIds() []ExceptionId`

GetExcepIds returns the ExcepIds field if non-nil, zero value otherwise.

### GetExcepIdsOk

`func (o *EventFilter) GetExcepIdsOk() (*[]ExceptionId, bool)`

GetExcepIdsOk returns a tuple with the ExcepIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcepIds

`func (o *EventFilter) SetExcepIds(v []ExceptionId)`

SetExcepIds sets ExcepIds field to given value.

### HasExcepIds

`func (o *EventFilter) HasExcepIds() bool`

HasExcepIds returns a boolean if a field has been set.

### GetExptAnaType

`func (o *EventFilter) GetExptAnaType() ExpectedAnalyticsType`

GetExptAnaType returns the ExptAnaType field if non-nil, zero value otherwise.

### GetExptAnaTypeOk

`func (o *EventFilter) GetExptAnaTypeOk() (*ExpectedAnalyticsType, bool)`

GetExptAnaTypeOk returns a tuple with the ExptAnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptAnaType

`func (o *EventFilter) SetExptAnaType(v ExpectedAnalyticsType)`

SetExptAnaType sets ExptAnaType field to given value.

### HasExptAnaType

`func (o *EventFilter) HasExptAnaType() bool`

HasExptAnaType returns a boolean if a field has been set.

### GetExptUeBehav

`func (o *EventFilter) GetExptUeBehav() ExpectedUeBehaviourData`

GetExptUeBehav returns the ExptUeBehav field if non-nil, zero value otherwise.

### GetExptUeBehavOk

`func (o *EventFilter) GetExptUeBehavOk() (*ExpectedUeBehaviourData, bool)`

GetExptUeBehavOk returns a tuple with the ExptUeBehav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExptUeBehav

`func (o *EventFilter) SetExptUeBehav(v ExpectedUeBehaviourData)`

SetExptUeBehav sets ExptUeBehav field to given value.

### HasExptUeBehav

`func (o *EventFilter) HasExptUeBehav() bool`

HasExptUeBehav returns a boolean if a field has been set.

### GetRatTypes

`func (o *EventFilter) GetRatTypes() []RatType`

GetRatTypes returns the RatTypes field if non-nil, zero value otherwise.

### GetRatTypesOk

`func (o *EventFilter) GetRatTypesOk() (*[]RatType, bool)`

GetRatTypesOk returns a tuple with the RatTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatTypes

`func (o *EventFilter) SetRatTypes(v []RatType)`

SetRatTypes sets RatTypes field to given value.

### HasRatTypes

`func (o *EventFilter) HasRatTypes() bool`

HasRatTypes returns a boolean if a field has been set.

### GetFreqs

`func (o *EventFilter) GetFreqs() []int32`

GetFreqs returns the Freqs field if non-nil, zero value otherwise.

### GetFreqsOk

`func (o *EventFilter) GetFreqsOk() (*[]int32, bool)`

GetFreqsOk returns a tuple with the Freqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreqs

`func (o *EventFilter) SetFreqs(v []int32)`

SetFreqs sets Freqs field to given value.

### HasFreqs

`func (o *EventFilter) HasFreqs() bool`

HasFreqs returns a boolean if a field has been set.

### GetDisperReqs

`func (o *EventFilter) GetDisperReqs() []DispersionRequirement`

GetDisperReqs returns the DisperReqs field if non-nil, zero value otherwise.

### GetDisperReqsOk

`func (o *EventFilter) GetDisperReqsOk() (*[]DispersionRequirement, bool)`

GetDisperReqsOk returns a tuple with the DisperReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisperReqs

`func (o *EventFilter) SetDisperReqs(v []DispersionRequirement)`

SetDisperReqs sets DisperReqs field to given value.

### HasDisperReqs

`func (o *EventFilter) HasDisperReqs() bool`

HasDisperReqs returns a boolean if a field has been set.

### GetRedTransReqs

`func (o *EventFilter) GetRedTransReqs() []RedundantTransmissionExpReq`

GetRedTransReqs returns the RedTransReqs field if non-nil, zero value otherwise.

### GetRedTransReqsOk

`func (o *EventFilter) GetRedTransReqsOk() (*[]RedundantTransmissionExpReq, bool)`

GetRedTransReqsOk returns a tuple with the RedTransReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedTransReqs

`func (o *EventFilter) SetRedTransReqs(v []RedundantTransmissionExpReq)`

SetRedTransReqs sets RedTransReqs field to given value.

### HasRedTransReqs

`func (o *EventFilter) HasRedTransReqs() bool`

HasRedTransReqs returns a boolean if a field has been set.

### GetWlanReqs

`func (o *EventFilter) GetWlanReqs() []WlanPerformanceReq`

GetWlanReqs returns the WlanReqs field if non-nil, zero value otherwise.

### GetWlanReqsOk

`func (o *EventFilter) GetWlanReqsOk() (*[]WlanPerformanceReq, bool)`

GetWlanReqsOk returns a tuple with the WlanReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanReqs

`func (o *EventFilter) SetWlanReqs(v []WlanPerformanceReq)`

SetWlanReqs sets WlanReqs field to given value.

### HasWlanReqs

`func (o *EventFilter) HasWlanReqs() bool`

HasWlanReqs returns a boolean if a field has been set.

### GetListOfAnaSubsets

`func (o *EventFilter) GetListOfAnaSubsets() []AnalyticsSubset`

GetListOfAnaSubsets returns the ListOfAnaSubsets field if non-nil, zero value otherwise.

### GetListOfAnaSubsetsOk

`func (o *EventFilter) GetListOfAnaSubsetsOk() (*[]AnalyticsSubset, bool)`

GetListOfAnaSubsetsOk returns a tuple with the ListOfAnaSubsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfAnaSubsets

`func (o *EventFilter) SetListOfAnaSubsets(v []AnalyticsSubset)`

SetListOfAnaSubsets sets ListOfAnaSubsets field to given value.

### HasListOfAnaSubsets

`func (o *EventFilter) HasListOfAnaSubsets() bool`

HasListOfAnaSubsets returns a boolean if a field has been set.

### GetUpfId

`func (o *EventFilter) GetUpfId() string`

GetUpfId returns the UpfId field if non-nil, zero value otherwise.

### GetUpfIdOk

`func (o *EventFilter) GetUpfIdOk() (*string, bool)`

GetUpfIdOk returns a tuple with the UpfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfId

`func (o *EventFilter) SetUpfId(v string)`

SetUpfId sets UpfId field to given value.

### HasUpfId

`func (o *EventFilter) HasUpfId() bool`

HasUpfId returns a boolean if a field has been set.

### GetAppServerAddrs

`func (o *EventFilter) GetAppServerAddrs() []AddrFqdn`

GetAppServerAddrs returns the AppServerAddrs field if non-nil, zero value otherwise.

### GetAppServerAddrsOk

`func (o *EventFilter) GetAppServerAddrsOk() (*[]AddrFqdn, bool)`

GetAppServerAddrsOk returns a tuple with the AppServerAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerAddrs

`func (o *EventFilter) SetAppServerAddrs(v []AddrFqdn)`

SetAppServerAddrs sets AppServerAddrs field to given value.

### HasAppServerAddrs

`func (o *EventFilter) HasAppServerAddrs() bool`

HasAppServerAddrs returns a boolean if a field has been set.

### GetDnPerfReqs

`func (o *EventFilter) GetDnPerfReqs() []DnPerformanceReq`

GetDnPerfReqs returns the DnPerfReqs field if non-nil, zero value otherwise.

### GetDnPerfReqsOk

`func (o *EventFilter) GetDnPerfReqsOk() (*[]DnPerformanceReq, bool)`

GetDnPerfReqsOk returns a tuple with the DnPerfReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPerfReqs

`func (o *EventFilter) SetDnPerfReqs(v []DnPerformanceReq)`

SetDnPerfReqs sets DnPerfReqs field to given value.

### HasDnPerfReqs

`func (o *EventFilter) HasDnPerfReqs() bool`

HasDnPerfReqs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


