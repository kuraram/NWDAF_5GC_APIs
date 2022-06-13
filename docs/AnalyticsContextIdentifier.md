# AnalyticsContextIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | The identifier of a subscription. | [optional] 
**NfAnaCtxts** | Pointer to [**[]NwdafEvent**](NwdafEvent.md) | List of analytics types for which NF related analytics contexts can be retrieved. | [optional] 
**UeAnaCtxts** | Pointer to [**[]UeAnalyticsContextDescriptor**](UeAnalyticsContextDescriptor.md) | List of objects that indicate for which SUPI and analytics types combinations analytics context can be retrieved. | [optional] 

## Methods

### NewAnalyticsContextIdentifier

`func NewAnalyticsContextIdentifier() *AnalyticsContextIdentifier`

NewAnalyticsContextIdentifier instantiates a new AnalyticsContextIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsContextIdentifierWithDefaults

`func NewAnalyticsContextIdentifierWithDefaults() *AnalyticsContextIdentifier`

NewAnalyticsContextIdentifierWithDefaults instantiates a new AnalyticsContextIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AnalyticsContextIdentifier) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AnalyticsContextIdentifier) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AnalyticsContextIdentifier) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AnalyticsContextIdentifier) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetNfAnaCtxts

`func (o *AnalyticsContextIdentifier) GetNfAnaCtxts() []NwdafEvent`

GetNfAnaCtxts returns the NfAnaCtxts field if non-nil, zero value otherwise.

### GetNfAnaCtxtsOk

`func (o *AnalyticsContextIdentifier) GetNfAnaCtxtsOk() (*[]NwdafEvent, bool)`

GetNfAnaCtxtsOk returns a tuple with the NfAnaCtxts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfAnaCtxts

`func (o *AnalyticsContextIdentifier) SetNfAnaCtxts(v []NwdafEvent)`

SetNfAnaCtxts sets NfAnaCtxts field to given value.

### HasNfAnaCtxts

`func (o *AnalyticsContextIdentifier) HasNfAnaCtxts() bool`

HasNfAnaCtxts returns a boolean if a field has been set.

### GetUeAnaCtxts

`func (o *AnalyticsContextIdentifier) GetUeAnaCtxts() []UeAnalyticsContextDescriptor`

GetUeAnaCtxts returns the UeAnaCtxts field if non-nil, zero value otherwise.

### GetUeAnaCtxtsOk

`func (o *AnalyticsContextIdentifier) GetUeAnaCtxtsOk() (*[]UeAnalyticsContextDescriptor, bool)`

GetUeAnaCtxtsOk returns a tuple with the UeAnaCtxts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeAnaCtxts

`func (o *AnalyticsContextIdentifier) SetUeAnaCtxts(v []UeAnalyticsContextDescriptor)`

SetUeAnaCtxts sets UeAnaCtxts field to given value.

### HasUeAnaCtxts

`func (o *AnalyticsContextIdentifier) HasUeAnaCtxts() bool`

HasUeAnaCtxts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


