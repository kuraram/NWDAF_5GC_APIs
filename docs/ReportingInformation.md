# ReportingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImmRep** | Pointer to **bool** |  | [optional] 
**NotifMethod** | Pointer to [**NotificationMethod1**](NotificationMethod1.md) |  | [optional] 
**MaxReportNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MonDur** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SampRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | [optional] 
**PartitionCriteria** | Pointer to [**[]PartitioningCriteria**](PartitioningCriteria.md) | Criteria for partitioning the UEs before applying the sampling ratio. | [optional] 
**GrpRepTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NotifFlag** | Pointer to [**NotificationFlag**](NotificationFlag.md) |  | [optional] 

## Methods

### NewReportingInformation

`func NewReportingInformation() *ReportingInformation`

NewReportingInformation instantiates a new ReportingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingInformationWithDefaults

`func NewReportingInformationWithDefaults() *ReportingInformation`

NewReportingInformationWithDefaults instantiates a new ReportingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImmRep

`func (o *ReportingInformation) GetImmRep() bool`

GetImmRep returns the ImmRep field if non-nil, zero value otherwise.

### GetImmRepOk

`func (o *ReportingInformation) GetImmRepOk() (*bool, bool)`

GetImmRepOk returns a tuple with the ImmRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmRep

`func (o *ReportingInformation) SetImmRep(v bool)`

SetImmRep sets ImmRep field to given value.

### HasImmRep

`func (o *ReportingInformation) HasImmRep() bool`

HasImmRep returns a boolean if a field has been set.

### GetNotifMethod

`func (o *ReportingInformation) GetNotifMethod() NotificationMethod1`

GetNotifMethod returns the NotifMethod field if non-nil, zero value otherwise.

### GetNotifMethodOk

`func (o *ReportingInformation) GetNotifMethodOk() (*NotificationMethod1, bool)`

GetNotifMethodOk returns a tuple with the NotifMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifMethod

`func (o *ReportingInformation) SetNotifMethod(v NotificationMethod1)`

SetNotifMethod sets NotifMethod field to given value.

### HasNotifMethod

`func (o *ReportingInformation) HasNotifMethod() bool`

HasNotifMethod returns a boolean if a field has been set.

### GetMaxReportNbr

`func (o *ReportingInformation) GetMaxReportNbr() int32`

GetMaxReportNbr returns the MaxReportNbr field if non-nil, zero value otherwise.

### GetMaxReportNbrOk

`func (o *ReportingInformation) GetMaxReportNbrOk() (*int32, bool)`

GetMaxReportNbrOk returns a tuple with the MaxReportNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReportNbr

`func (o *ReportingInformation) SetMaxReportNbr(v int32)`

SetMaxReportNbr sets MaxReportNbr field to given value.

### HasMaxReportNbr

`func (o *ReportingInformation) HasMaxReportNbr() bool`

HasMaxReportNbr returns a boolean if a field has been set.

### GetMonDur

`func (o *ReportingInformation) GetMonDur() time.Time`

GetMonDur returns the MonDur field if non-nil, zero value otherwise.

### GetMonDurOk

`func (o *ReportingInformation) GetMonDurOk() (*time.Time, bool)`

GetMonDurOk returns a tuple with the MonDur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonDur

`func (o *ReportingInformation) SetMonDur(v time.Time)`

SetMonDur sets MonDur field to given value.

### HasMonDur

`func (o *ReportingInformation) HasMonDur() bool`

HasMonDur returns a boolean if a field has been set.

### GetRepPeriod

`func (o *ReportingInformation) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *ReportingInformation) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *ReportingInformation) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *ReportingInformation) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetSampRatio

`func (o *ReportingInformation) GetSampRatio() int32`

GetSampRatio returns the SampRatio field if non-nil, zero value otherwise.

### GetSampRatioOk

`func (o *ReportingInformation) GetSampRatioOk() (*int32, bool)`

GetSampRatioOk returns a tuple with the SampRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampRatio

`func (o *ReportingInformation) SetSampRatio(v int32)`

SetSampRatio sets SampRatio field to given value.

### HasSampRatio

`func (o *ReportingInformation) HasSampRatio() bool`

HasSampRatio returns a boolean if a field has been set.

### GetPartitionCriteria

`func (o *ReportingInformation) GetPartitionCriteria() []PartitioningCriteria`

GetPartitionCriteria returns the PartitionCriteria field if non-nil, zero value otherwise.

### GetPartitionCriteriaOk

`func (o *ReportingInformation) GetPartitionCriteriaOk() (*[]PartitioningCriteria, bool)`

GetPartitionCriteriaOk returns a tuple with the PartitionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCriteria

`func (o *ReportingInformation) SetPartitionCriteria(v []PartitioningCriteria)`

SetPartitionCriteria sets PartitionCriteria field to given value.

### HasPartitionCriteria

`func (o *ReportingInformation) HasPartitionCriteria() bool`

HasPartitionCriteria returns a boolean if a field has been set.

### GetGrpRepTime

`func (o *ReportingInformation) GetGrpRepTime() int32`

GetGrpRepTime returns the GrpRepTime field if non-nil, zero value otherwise.

### GetGrpRepTimeOk

`func (o *ReportingInformation) GetGrpRepTimeOk() (*int32, bool)`

GetGrpRepTimeOk returns a tuple with the GrpRepTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpRepTime

`func (o *ReportingInformation) SetGrpRepTime(v int32)`

SetGrpRepTime sets GrpRepTime field to given value.

### HasGrpRepTime

`func (o *ReportingInformation) HasGrpRepTime() bool`

HasGrpRepTime returns a boolean if a field has been set.

### GetNotifFlag

`func (o *ReportingInformation) GetNotifFlag() NotificationFlag`

GetNotifFlag returns the NotifFlag field if non-nil, zero value otherwise.

### GetNotifFlagOk

`func (o *ReportingInformation) GetNotifFlagOk() (*NotificationFlag, bool)`

GetNotifFlagOk returns a tuple with the NotifFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifFlag

`func (o *ReportingInformation) SetNotifFlag(v NotificationFlag)`

SetNotifFlag sets NotifFlag field to given value.

### HasNotifFlag

`func (o *ReportingInformation) HasNotifFlag() bool`

HasNotifFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


