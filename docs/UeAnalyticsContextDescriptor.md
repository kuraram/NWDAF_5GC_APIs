# UeAnalyticsContextDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause 2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2 of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2 of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of 3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention defined in 3GPP TS 29.501.  | 
**AnaTypes** | [**[]NwdafEvent**](NwdafEvent.md) | List of analytics types for which UE related analytics contexts can be retrieved. | 

## Methods

### NewUeAnalyticsContextDescriptor

`func NewUeAnalyticsContextDescriptor(supi string, anaTypes []NwdafEvent, ) *UeAnalyticsContextDescriptor`

NewUeAnalyticsContextDescriptor instantiates a new UeAnalyticsContextDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeAnalyticsContextDescriptorWithDefaults

`func NewUeAnalyticsContextDescriptorWithDefaults() *UeAnalyticsContextDescriptor`

NewUeAnalyticsContextDescriptorWithDefaults instantiates a new UeAnalyticsContextDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UeAnalyticsContextDescriptor) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeAnalyticsContextDescriptor) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeAnalyticsContextDescriptor) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetAnaTypes

`func (o *UeAnalyticsContextDescriptor) GetAnaTypes() []NwdafEvent`

GetAnaTypes returns the AnaTypes field if non-nil, zero value otherwise.

### GetAnaTypesOk

`func (o *UeAnalyticsContextDescriptor) GetAnaTypesOk() (*[]NwdafEvent, bool)`

GetAnaTypesOk returns a tuple with the AnaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaTypes

`func (o *UeAnalyticsContextDescriptor) SetAnaTypes(v []NwdafEvent)`

SetAnaTypes sets AnaTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


