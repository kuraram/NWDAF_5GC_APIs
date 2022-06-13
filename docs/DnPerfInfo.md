# DnPerfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003; it shall contain either a DNN Network Identifier, or a full DNN with both the Network Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots (e.g. \&quot;Label1.Label2.Label3\&quot;). | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**DnPerf** | [**[]DnPerf**](DnPerf.md) |  | 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewDnPerfInfo

`func NewDnPerfInfo(dnPerf []DnPerf, ) *DnPerfInfo`

NewDnPerfInfo instantiates a new DnPerfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnPerfInfoWithDefaults

`func NewDnPerfInfoWithDefaults() *DnPerfInfo`

NewDnPerfInfoWithDefaults instantiates a new DnPerfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *DnPerfInfo) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *DnPerfInfo) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *DnPerfInfo) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *DnPerfInfo) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDnn

`func (o *DnPerfInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *DnPerfInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *DnPerfInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *DnPerfInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *DnPerfInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *DnPerfInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *DnPerfInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *DnPerfInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetDnPerf

`func (o *DnPerfInfo) GetDnPerf() []DnPerf`

GetDnPerf returns the DnPerf field if non-nil, zero value otherwise.

### GetDnPerfOk

`func (o *DnPerfInfo) GetDnPerfOk() (*[]DnPerf, bool)`

GetDnPerfOk returns a tuple with the DnPerf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnPerf

`func (o *DnPerfInfo) SetDnPerf(v []DnPerf)`

SetDnPerf sets DnPerf field to given value.


### GetConfidence

`func (o *DnPerfInfo) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *DnPerfInfo) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *DnPerfInfo) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *DnPerfInfo) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


