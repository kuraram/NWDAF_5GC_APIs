# ConsumerNfInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 

## Methods

### NewConsumerNfInformation

`func NewConsumerNfInformation() *ConsumerNfInformation`

NewConsumerNfInformation instantiates a new ConsumerNfInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerNfInformationWithDefaults

`func NewConsumerNfInformationWithDefaults() *ConsumerNfInformation`

NewConsumerNfInformationWithDefaults instantiates a new ConsumerNfInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfId

`func (o *ConsumerNfInformation) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *ConsumerNfInformation) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *ConsumerNfInformation) SetNfId(v string)`

SetNfId sets NfId field to given value.

### HasNfId

`func (o *ConsumerNfInformation) HasNfId() bool`

HasNfId returns a boolean if a field has been set.

### GetTaiList

`func (o *ConsumerNfInformation) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *ConsumerNfInformation) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *ConsumerNfInformation) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *ConsumerNfInformation) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


