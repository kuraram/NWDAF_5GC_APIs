# EventReportingRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accuracy** | Pointer to [**Accuracy**](Accuracy.md) |  | [optional] 
**AccPerSubset** | Pointer to [**[]Accuracy**](Accuracy.md) | Each element indicates the preferred accuracy level per analytics subset. It may be present if the \&quot;listOfAnaSubsets\&quot; attribute is present in the subscription request when the subscription event is NF_LOAD, UE_COMM, DISPERSION, NETWORK_PERFORMANCE, WLAN_PERFORMANCE, DN_PERFORMANCE or SERVICE_EXPERIENCE.  | [optional] 
**StartTs** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**EndTs** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**OffsetPeriod** | Pointer to **int32** | Offset period in units of seconds to the reporting time, if the value is negative means statistics in the past offset period, otherwise a positive value means prediction in the future offset period. May be present if the \&quot;repPeriod\&quot; attribute is included within the \&quot;evtReq\&quot; attribute. | [optional] 
**SampRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | [optional] 
**MaxObjectNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MaxSupiNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TimeAnaNeeded** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**AnaMeta** | Pointer to [**[]AnalyticsMetadata**](AnalyticsMetadata.md) |  | [optional] 
**AnaMetaInd** | Pointer to [**AnalyticsMetadataIndication**](AnalyticsMetadataIndication.md) |  | [optional] 

## Methods

### NewEventReportingRequirement

`func NewEventReportingRequirement() *EventReportingRequirement`

NewEventReportingRequirement instantiates a new EventReportingRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventReportingRequirementWithDefaults

`func NewEventReportingRequirementWithDefaults() *EventReportingRequirement`

NewEventReportingRequirementWithDefaults instantiates a new EventReportingRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracy

`func (o *EventReportingRequirement) GetAccuracy() Accuracy`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *EventReportingRequirement) GetAccuracyOk() (*Accuracy, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *EventReportingRequirement) SetAccuracy(v Accuracy)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *EventReportingRequirement) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetAccPerSubset

`func (o *EventReportingRequirement) GetAccPerSubset() []Accuracy`

GetAccPerSubset returns the AccPerSubset field if non-nil, zero value otherwise.

### GetAccPerSubsetOk

`func (o *EventReportingRequirement) GetAccPerSubsetOk() (*[]Accuracy, bool)`

GetAccPerSubsetOk returns a tuple with the AccPerSubset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccPerSubset

`func (o *EventReportingRequirement) SetAccPerSubset(v []Accuracy)`

SetAccPerSubset sets AccPerSubset field to given value.

### HasAccPerSubset

`func (o *EventReportingRequirement) HasAccPerSubset() bool`

HasAccPerSubset returns a boolean if a field has been set.

### GetStartTs

`func (o *EventReportingRequirement) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *EventReportingRequirement) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *EventReportingRequirement) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *EventReportingRequirement) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *EventReportingRequirement) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *EventReportingRequirement) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *EventReportingRequirement) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *EventReportingRequirement) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetOffsetPeriod

`func (o *EventReportingRequirement) GetOffsetPeriod() int32`

GetOffsetPeriod returns the OffsetPeriod field if non-nil, zero value otherwise.

### GetOffsetPeriodOk

`func (o *EventReportingRequirement) GetOffsetPeriodOk() (*int32, bool)`

GetOffsetPeriodOk returns a tuple with the OffsetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetPeriod

`func (o *EventReportingRequirement) SetOffsetPeriod(v int32)`

SetOffsetPeriod sets OffsetPeriod field to given value.

### HasOffsetPeriod

`func (o *EventReportingRequirement) HasOffsetPeriod() bool`

HasOffsetPeriod returns a boolean if a field has been set.

### GetSampRatio

`func (o *EventReportingRequirement) GetSampRatio() int32`

GetSampRatio returns the SampRatio field if non-nil, zero value otherwise.

### GetSampRatioOk

`func (o *EventReportingRequirement) GetSampRatioOk() (*int32, bool)`

GetSampRatioOk returns a tuple with the SampRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampRatio

`func (o *EventReportingRequirement) SetSampRatio(v int32)`

SetSampRatio sets SampRatio field to given value.

### HasSampRatio

`func (o *EventReportingRequirement) HasSampRatio() bool`

HasSampRatio returns a boolean if a field has been set.

### GetMaxObjectNbr

`func (o *EventReportingRequirement) GetMaxObjectNbr() int32`

GetMaxObjectNbr returns the MaxObjectNbr field if non-nil, zero value otherwise.

### GetMaxObjectNbrOk

`func (o *EventReportingRequirement) GetMaxObjectNbrOk() (*int32, bool)`

GetMaxObjectNbrOk returns a tuple with the MaxObjectNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxObjectNbr

`func (o *EventReportingRequirement) SetMaxObjectNbr(v int32)`

SetMaxObjectNbr sets MaxObjectNbr field to given value.

### HasMaxObjectNbr

`func (o *EventReportingRequirement) HasMaxObjectNbr() bool`

HasMaxObjectNbr returns a boolean if a field has been set.

### GetMaxSupiNbr

`func (o *EventReportingRequirement) GetMaxSupiNbr() int32`

GetMaxSupiNbr returns the MaxSupiNbr field if non-nil, zero value otherwise.

### GetMaxSupiNbrOk

`func (o *EventReportingRequirement) GetMaxSupiNbrOk() (*int32, bool)`

GetMaxSupiNbrOk returns a tuple with the MaxSupiNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSupiNbr

`func (o *EventReportingRequirement) SetMaxSupiNbr(v int32)`

SetMaxSupiNbr sets MaxSupiNbr field to given value.

### HasMaxSupiNbr

`func (o *EventReportingRequirement) HasMaxSupiNbr() bool`

HasMaxSupiNbr returns a boolean if a field has been set.

### GetTimeAnaNeeded

`func (o *EventReportingRequirement) GetTimeAnaNeeded() time.Time`

GetTimeAnaNeeded returns the TimeAnaNeeded field if non-nil, zero value otherwise.

### GetTimeAnaNeededOk

`func (o *EventReportingRequirement) GetTimeAnaNeededOk() (*time.Time, bool)`

GetTimeAnaNeededOk returns a tuple with the TimeAnaNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAnaNeeded

`func (o *EventReportingRequirement) SetTimeAnaNeeded(v time.Time)`

SetTimeAnaNeeded sets TimeAnaNeeded field to given value.

### HasTimeAnaNeeded

`func (o *EventReportingRequirement) HasTimeAnaNeeded() bool`

HasTimeAnaNeeded returns a boolean if a field has been set.

### GetAnaMeta

`func (o *EventReportingRequirement) GetAnaMeta() []AnalyticsMetadata`

GetAnaMeta returns the AnaMeta field if non-nil, zero value otherwise.

### GetAnaMetaOk

`func (o *EventReportingRequirement) GetAnaMetaOk() (*[]AnalyticsMetadata, bool)`

GetAnaMetaOk returns a tuple with the AnaMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaMeta

`func (o *EventReportingRequirement) SetAnaMeta(v []AnalyticsMetadata)`

SetAnaMeta sets AnaMeta field to given value.

### HasAnaMeta

`func (o *EventReportingRequirement) HasAnaMeta() bool`

HasAnaMeta returns a boolean if a field has been set.

### GetAnaMetaInd

`func (o *EventReportingRequirement) GetAnaMetaInd() AnalyticsMetadataIndication`

GetAnaMetaInd returns the AnaMetaInd field if non-nil, zero value otherwise.

### GetAnaMetaIndOk

`func (o *EventReportingRequirement) GetAnaMetaIndOk() (*AnalyticsMetadataIndication, bool)`

GetAnaMetaIndOk returns a tuple with the AnaMetaInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaMetaInd

`func (o *EventReportingRequirement) SetAnaMetaInd(v AnalyticsMetadataIndication)`

SetAnaMetaInd sets AnaMetaInd field to given value.

### HasAnaMetaInd

`func (o *EventReportingRequirement) HasAnaMetaInd() bool`

HasAnaMetaInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


