# EventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnySlice** | Pointer to **bool** | FALSE represents not applicable for all slices. TRUE represents applicable for all slices. | [optional] 
**Snssais** | Pointer to [**[]Snssai**](Snssai.md) | Identification(s) of network slice to which the subscription belongs. | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**Dnns** | Pointer to **[]string** |  | [optional] 
**Dnais** | Pointer to **[]string** |  | [optional] 
**NetworkArea** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


