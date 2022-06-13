# ContextElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextId** | [**AnalyticsContextIdentifier**](AnalyticsContextIdentifier.md) |  | 
**PendAnalytics** | Pointer to [**[]EventNotification**](EventNotification.md) | Output analytics for the analytics subscription which have not yet been sent to the analytics consumer. | [optional] 
**HistAnalytics** | Pointer to [**[]EventNotification**](EventNotification.md) | Historical output analytics. | [optional] 
**LastOutputTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**AggrSubs** | Pointer to [**[]SpecificAnalyticsSubscription**](SpecificAnalyticsSubscription.md) | Information about analytics subscriptions that the NWDAF has with other NWDAFs to perform aggregation. | [optional] 
**HistData** | Pointer to [**[]HistoricalData**](HistoricalData.md) | Historical data related to the analytics subscription. | [optional] 
**AdrfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**AdrfDataTypes** | Pointer to [**[]AdrfDataType**](AdrfDataType.md) | Type(s) of data stored in the ADRF by the NWDAF. | [optional] 
**AggrNwdafIds** | Pointer to **[]string** | NWDAF identifiers of NWDAF instances used by the NWDAF service consumer when aggregating multiple analytics subscriptions. | [optional] 
**ModelProvIds** | Pointer to **[]string** | Identifiers of NWDAFs that provide ML models used by the NF service consumer. | [optional] 

## Methods

### NewContextElement

`func NewContextElement(contextId AnalyticsContextIdentifier, ) *ContextElement`

NewContextElement instantiates a new ContextElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextElementWithDefaults

`func NewContextElementWithDefaults() *ContextElement`

NewContextElementWithDefaults instantiates a new ContextElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextId

`func (o *ContextElement) GetContextId() AnalyticsContextIdentifier`

GetContextId returns the ContextId field if non-nil, zero value otherwise.

### GetContextIdOk

`func (o *ContextElement) GetContextIdOk() (*AnalyticsContextIdentifier, bool)`

GetContextIdOk returns a tuple with the ContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextId

`func (o *ContextElement) SetContextId(v AnalyticsContextIdentifier)`

SetContextId sets ContextId field to given value.


### GetPendAnalytics

`func (o *ContextElement) GetPendAnalytics() []EventNotification`

GetPendAnalytics returns the PendAnalytics field if non-nil, zero value otherwise.

### GetPendAnalyticsOk

`func (o *ContextElement) GetPendAnalyticsOk() (*[]EventNotification, bool)`

GetPendAnalyticsOk returns a tuple with the PendAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendAnalytics

`func (o *ContextElement) SetPendAnalytics(v []EventNotification)`

SetPendAnalytics sets PendAnalytics field to given value.

### HasPendAnalytics

`func (o *ContextElement) HasPendAnalytics() bool`

HasPendAnalytics returns a boolean if a field has been set.

### GetHistAnalytics

`func (o *ContextElement) GetHistAnalytics() []EventNotification`

GetHistAnalytics returns the HistAnalytics field if non-nil, zero value otherwise.

### GetHistAnalyticsOk

`func (o *ContextElement) GetHistAnalyticsOk() (*[]EventNotification, bool)`

GetHistAnalyticsOk returns a tuple with the HistAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistAnalytics

`func (o *ContextElement) SetHistAnalytics(v []EventNotification)`

SetHistAnalytics sets HistAnalytics field to given value.

### HasHistAnalytics

`func (o *ContextElement) HasHistAnalytics() bool`

HasHistAnalytics returns a boolean if a field has been set.

### GetLastOutputTime

`func (o *ContextElement) GetLastOutputTime() time.Time`

GetLastOutputTime returns the LastOutputTime field if non-nil, zero value otherwise.

### GetLastOutputTimeOk

`func (o *ContextElement) GetLastOutputTimeOk() (*time.Time, bool)`

GetLastOutputTimeOk returns a tuple with the LastOutputTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOutputTime

`func (o *ContextElement) SetLastOutputTime(v time.Time)`

SetLastOutputTime sets LastOutputTime field to given value.

### HasLastOutputTime

`func (o *ContextElement) HasLastOutputTime() bool`

HasLastOutputTime returns a boolean if a field has been set.

### GetAggrSubs

`func (o *ContextElement) GetAggrSubs() []SpecificAnalyticsSubscription`

GetAggrSubs returns the AggrSubs field if non-nil, zero value otherwise.

### GetAggrSubsOk

`func (o *ContextElement) GetAggrSubsOk() (*[]SpecificAnalyticsSubscription, bool)`

GetAggrSubsOk returns a tuple with the AggrSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggrSubs

`func (o *ContextElement) SetAggrSubs(v []SpecificAnalyticsSubscription)`

SetAggrSubs sets AggrSubs field to given value.

### HasAggrSubs

`func (o *ContextElement) HasAggrSubs() bool`

HasAggrSubs returns a boolean if a field has been set.

### GetHistData

`func (o *ContextElement) GetHistData() []HistoricalData`

GetHistData returns the HistData field if non-nil, zero value otherwise.

### GetHistDataOk

`func (o *ContextElement) GetHistDataOk() (*[]HistoricalData, bool)`

GetHistDataOk returns a tuple with the HistData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistData

`func (o *ContextElement) SetHistData(v []HistoricalData)`

SetHistData sets HistData field to given value.

### HasHistData

`func (o *ContextElement) HasHistData() bool`

HasHistData returns a boolean if a field has been set.

### GetAdrfId

`func (o *ContextElement) GetAdrfId() string`

GetAdrfId returns the AdrfId field if non-nil, zero value otherwise.

### GetAdrfIdOk

`func (o *ContextElement) GetAdrfIdOk() (*string, bool)`

GetAdrfIdOk returns a tuple with the AdrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrfId

`func (o *ContextElement) SetAdrfId(v string)`

SetAdrfId sets AdrfId field to given value.

### HasAdrfId

`func (o *ContextElement) HasAdrfId() bool`

HasAdrfId returns a boolean if a field has been set.

### GetAdrfDataTypes

`func (o *ContextElement) GetAdrfDataTypes() []AdrfDataType`

GetAdrfDataTypes returns the AdrfDataTypes field if non-nil, zero value otherwise.

### GetAdrfDataTypesOk

`func (o *ContextElement) GetAdrfDataTypesOk() (*[]AdrfDataType, bool)`

GetAdrfDataTypesOk returns a tuple with the AdrfDataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrfDataTypes

`func (o *ContextElement) SetAdrfDataTypes(v []AdrfDataType)`

SetAdrfDataTypes sets AdrfDataTypes field to given value.

### HasAdrfDataTypes

`func (o *ContextElement) HasAdrfDataTypes() bool`

HasAdrfDataTypes returns a boolean if a field has been set.

### GetAggrNwdafIds

`func (o *ContextElement) GetAggrNwdafIds() []string`

GetAggrNwdafIds returns the AggrNwdafIds field if non-nil, zero value otherwise.

### GetAggrNwdafIdsOk

`func (o *ContextElement) GetAggrNwdafIdsOk() (*[]string, bool)`

GetAggrNwdafIdsOk returns a tuple with the AggrNwdafIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggrNwdafIds

`func (o *ContextElement) SetAggrNwdafIds(v []string)`

SetAggrNwdafIds sets AggrNwdafIds field to given value.

### HasAggrNwdafIds

`func (o *ContextElement) HasAggrNwdafIds() bool`

HasAggrNwdafIds returns a boolean if a field has been set.

### GetModelProvIds

`func (o *ContextElement) GetModelProvIds() []string`

GetModelProvIds returns the ModelProvIds field if non-nil, zero value otherwise.

### GetModelProvIdsOk

`func (o *ContextElement) GetModelProvIdsOk() (*[]string, bool)`

GetModelProvIdsOk returns a tuple with the ModelProvIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvIds

`func (o *ContextElement) SetModelProvIds(v []string)`

SetModelProvIds sets ModelProvIds field to given value.

### HasModelProvIds

`func (o *ContextElement) HasModelProvIds() bool`

HasModelProvIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


