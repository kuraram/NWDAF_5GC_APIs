# SpecificAnalyticsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** |  | [optional] 
**ProducerId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**ProducerSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;  set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or \&quot;set&lt;SetID&gt;.  &lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with &lt;MCC&gt; encoded as defined in clause 5.4.2  (\&quot;Mcc\&quot; data type definition) &lt;MNC&gt; encoded as defined in clause 5.4.2 (\&quot;Mnc\&quot; data type  definition) &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but  with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic  characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end with either an  alphabetic character or a digit.   | [optional] 
**NwdafEvSub** | Pointer to [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | [optional] 

## Methods

### NewSpecificAnalyticsSubscription

`func NewSpecificAnalyticsSubscription() *SpecificAnalyticsSubscription`

NewSpecificAnalyticsSubscription instantiates a new SpecificAnalyticsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecificAnalyticsSubscriptionWithDefaults

`func NewSpecificAnalyticsSubscriptionWithDefaults() *SpecificAnalyticsSubscription`

NewSpecificAnalyticsSubscriptionWithDefaults instantiates a new SpecificAnalyticsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *SpecificAnalyticsSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SpecificAnalyticsSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SpecificAnalyticsSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SpecificAnalyticsSubscription) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetProducerId

`func (o *SpecificAnalyticsSubscription) GetProducerId() string`

GetProducerId returns the ProducerId field if non-nil, zero value otherwise.

### GetProducerIdOk

`func (o *SpecificAnalyticsSubscription) GetProducerIdOk() (*string, bool)`

GetProducerIdOk returns a tuple with the ProducerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerId

`func (o *SpecificAnalyticsSubscription) SetProducerId(v string)`

SetProducerId sets ProducerId field to given value.

### HasProducerId

`func (o *SpecificAnalyticsSubscription) HasProducerId() bool`

HasProducerId returns a boolean if a field has been set.

### GetProducerSetId

`func (o *SpecificAnalyticsSubscription) GetProducerSetId() string`

GetProducerSetId returns the ProducerSetId field if non-nil, zero value otherwise.

### GetProducerSetIdOk

`func (o *SpecificAnalyticsSubscription) GetProducerSetIdOk() (*string, bool)`

GetProducerSetIdOk returns a tuple with the ProducerSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerSetId

`func (o *SpecificAnalyticsSubscription) SetProducerSetId(v string)`

SetProducerSetId sets ProducerSetId field to given value.

### HasProducerSetId

`func (o *SpecificAnalyticsSubscription) HasProducerSetId() bool`

HasProducerSetId returns a boolean if a field has been set.

### GetNwdafEvSub

`func (o *SpecificAnalyticsSubscription) GetNwdafEvSub() NnwdafEventsSubscription`

GetNwdafEvSub returns the NwdafEvSub field if non-nil, zero value otherwise.

### GetNwdafEvSubOk

`func (o *SpecificAnalyticsSubscription) GetNwdafEvSubOk() (*NnwdafEventsSubscription, bool)`

GetNwdafEvSubOk returns a tuple with the NwdafEvSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafEvSub

`func (o *SpecificAnalyticsSubscription) SetNwdafEvSub(v NnwdafEventsSubscription)`

SetNwdafEvSub sets NwdafEvSub field to given value.

### HasNwdafEvSub

`func (o *SpecificAnalyticsSubscription) HasNwdafEvSub() bool`

HasNwdafEvSub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


