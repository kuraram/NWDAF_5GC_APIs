# NnwdafEventsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSubscriptions** | [**[]EventSubscription**](EventSubscription.md) | Subscribed events | 
**EvtReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**NotificationURI** | Pointer to **string** | String providing an URI formatted according to RFC 3986 | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**EventNotifications** | Pointer to [**[]EventNotification**](EventNotification.md) |  | [optional] 
**FailEventReports** | Pointer to [**[]FailureEventInfo**](FailureEventInfo.md) |  | [optional] 
**PrevSub** | Pointer to [**SpecificAnalyticsSubscription1**](SpecificAnalyticsSubscription1.md) |  | [optional] 
**ConsNfInfo** | Pointer to [**ConsumerNfInformation**](ConsumerNfInformation.md) |  | [optional] 

## Methods

### NewNnwdafEventsSubscription

`func NewNnwdafEventsSubscription(eventSubscriptions []EventSubscription, ) *NnwdafEventsSubscription`

NewNnwdafEventsSubscription instantiates a new NnwdafEventsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNnwdafEventsSubscriptionWithDefaults

`func NewNnwdafEventsSubscriptionWithDefaults() *NnwdafEventsSubscription`

NewNnwdafEventsSubscriptionWithDefaults instantiates a new NnwdafEventsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSubscriptions

`func (o *NnwdafEventsSubscription) GetEventSubscriptions() []EventSubscription`

GetEventSubscriptions returns the EventSubscriptions field if non-nil, zero value otherwise.

### GetEventSubscriptionsOk

`func (o *NnwdafEventsSubscription) GetEventSubscriptionsOk() (*[]EventSubscription, bool)`

GetEventSubscriptionsOk returns a tuple with the EventSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubscriptions

`func (o *NnwdafEventsSubscription) SetEventSubscriptions(v []EventSubscription)`

SetEventSubscriptions sets EventSubscriptions field to given value.


### GetEvtReq

`func (o *NnwdafEventsSubscription) GetEvtReq() ReportingInformation`

GetEvtReq returns the EvtReq field if non-nil, zero value otherwise.

### GetEvtReqOk

`func (o *NnwdafEventsSubscription) GetEvtReqOk() (*ReportingInformation, bool)`

GetEvtReqOk returns a tuple with the EvtReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvtReq

`func (o *NnwdafEventsSubscription) SetEvtReq(v ReportingInformation)`

SetEvtReq sets EvtReq field to given value.

### HasEvtReq

`func (o *NnwdafEventsSubscription) HasEvtReq() bool`

HasEvtReq returns a boolean if a field has been set.

### GetNotificationURI

`func (o *NnwdafEventsSubscription) GetNotificationURI() string`

GetNotificationURI returns the NotificationURI field if non-nil, zero value otherwise.

### GetNotificationURIOk

`func (o *NnwdafEventsSubscription) GetNotificationURIOk() (*string, bool)`

GetNotificationURIOk returns a tuple with the NotificationURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationURI

`func (o *NnwdafEventsSubscription) SetNotificationURI(v string)`

SetNotificationURI sets NotificationURI field to given value.

### HasNotificationURI

`func (o *NnwdafEventsSubscription) HasNotificationURI() bool`

HasNotificationURI returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NnwdafEventsSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NnwdafEventsSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NnwdafEventsSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NnwdafEventsSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetEventNotifications

`func (o *NnwdafEventsSubscription) GetEventNotifications() []EventNotification`

GetEventNotifications returns the EventNotifications field if non-nil, zero value otherwise.

### GetEventNotificationsOk

`func (o *NnwdafEventsSubscription) GetEventNotificationsOk() (*[]EventNotification, bool)`

GetEventNotificationsOk returns a tuple with the EventNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifications

`func (o *NnwdafEventsSubscription) SetEventNotifications(v []EventNotification)`

SetEventNotifications sets EventNotifications field to given value.

### HasEventNotifications

`func (o *NnwdafEventsSubscription) HasEventNotifications() bool`

HasEventNotifications returns a boolean if a field has been set.

### GetFailEventReports

`func (o *NnwdafEventsSubscription) GetFailEventReports() []FailureEventInfo`

GetFailEventReports returns the FailEventReports field if non-nil, zero value otherwise.

### GetFailEventReportsOk

`func (o *NnwdafEventsSubscription) GetFailEventReportsOk() (*[]FailureEventInfo, bool)`

GetFailEventReportsOk returns a tuple with the FailEventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailEventReports

`func (o *NnwdafEventsSubscription) SetFailEventReports(v []FailureEventInfo)`

SetFailEventReports sets FailEventReports field to given value.

### HasFailEventReports

`func (o *NnwdafEventsSubscription) HasFailEventReports() bool`

HasFailEventReports returns a boolean if a field has been set.

### GetPrevSub

`func (o *NnwdafEventsSubscription) GetPrevSub() SpecificAnalyticsSubscription1`

GetPrevSub returns the PrevSub field if non-nil, zero value otherwise.

### GetPrevSubOk

`func (o *NnwdafEventsSubscription) GetPrevSubOk() (*SpecificAnalyticsSubscription1, bool)`

GetPrevSubOk returns a tuple with the PrevSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevSub

`func (o *NnwdafEventsSubscription) SetPrevSub(v SpecificAnalyticsSubscription1)`

SetPrevSub sets PrevSub field to given value.

### HasPrevSub

`func (o *NnwdafEventsSubscription) HasPrevSub() bool`

HasPrevSub returns a boolean if a field has been set.

### GetConsNfInfo

`func (o *NnwdafEventsSubscription) GetConsNfInfo() ConsumerNfInformation`

GetConsNfInfo returns the ConsNfInfo field if non-nil, zero value otherwise.

### GetConsNfInfoOk

`func (o *NnwdafEventsSubscription) GetConsNfInfoOk() (*ConsumerNfInformation, bool)`

GetConsNfInfoOk returns a tuple with the ConsNfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsNfInfo

`func (o *NnwdafEventsSubscription) SetConsNfInfo(v ConsumerNfInformation)`

SetConsNfInfo sets ConsNfInfo field to given value.

### HasConsNfInfo

`func (o *NnwdafEventsSubscription) HasConsNfInfo() bool`

HasConsNfInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


