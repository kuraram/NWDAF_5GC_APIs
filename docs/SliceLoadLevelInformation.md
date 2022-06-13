# SliceLoadLevelInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadLevelInformation** | **int32** | Load level information of the network slice and the optionally associated network slice instance. | 
**Snssais** | [**[]Snssai**](Snssai.md) | Identification(s) of network slice to which the subscription applies. | 
**NumOfUes** | Pointer to [**NumberAverage**](NumberAverage.md) |  | [optional] 
**NumOfPduSess** | Pointer to [**NumberAverage**](NumberAverage.md) |  | [optional] 
**ExceedLoadLevelThrInd** | Pointer to **bool** |  | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewSliceLoadLevelInformation

`func NewSliceLoadLevelInformation(loadLevelInformation int32, snssais []Snssai, ) *SliceLoadLevelInformation`

NewSliceLoadLevelInformation instantiates a new SliceLoadLevelInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceLoadLevelInformationWithDefaults

`func NewSliceLoadLevelInformationWithDefaults() *SliceLoadLevelInformation`

NewSliceLoadLevelInformationWithDefaults instantiates a new SliceLoadLevelInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadLevelInformation

`func (o *SliceLoadLevelInformation) GetLoadLevelInformation() int32`

GetLoadLevelInformation returns the LoadLevelInformation field if non-nil, zero value otherwise.

### GetLoadLevelInformationOk

`func (o *SliceLoadLevelInformation) GetLoadLevelInformationOk() (*int32, bool)`

GetLoadLevelInformationOk returns a tuple with the LoadLevelInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadLevelInformation

`func (o *SliceLoadLevelInformation) SetLoadLevelInformation(v int32)`

SetLoadLevelInformation sets LoadLevelInformation field to given value.


### GetSnssais

`func (o *SliceLoadLevelInformation) GetSnssais() []Snssai`

GetSnssais returns the Snssais field if non-nil, zero value otherwise.

### GetSnssaisOk

`func (o *SliceLoadLevelInformation) GetSnssaisOk() (*[]Snssai, bool)`

GetSnssaisOk returns a tuple with the Snssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssais

`func (o *SliceLoadLevelInformation) SetSnssais(v []Snssai)`

SetSnssais sets Snssais field to given value.


### GetNumOfUes

`func (o *SliceLoadLevelInformation) GetNumOfUes() NumberAverage`

GetNumOfUes returns the NumOfUes field if non-nil, zero value otherwise.

### GetNumOfUesOk

`func (o *SliceLoadLevelInformation) GetNumOfUesOk() (*NumberAverage, bool)`

GetNumOfUesOk returns a tuple with the NumOfUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfUes

`func (o *SliceLoadLevelInformation) SetNumOfUes(v NumberAverage)`

SetNumOfUes sets NumOfUes field to given value.

### HasNumOfUes

`func (o *SliceLoadLevelInformation) HasNumOfUes() bool`

HasNumOfUes returns a boolean if a field has been set.

### GetNumOfPduSess

`func (o *SliceLoadLevelInformation) GetNumOfPduSess() NumberAverage`

GetNumOfPduSess returns the NumOfPduSess field if non-nil, zero value otherwise.

### GetNumOfPduSessOk

`func (o *SliceLoadLevelInformation) GetNumOfPduSessOk() (*NumberAverage, bool)`

GetNumOfPduSessOk returns a tuple with the NumOfPduSess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfPduSess

`func (o *SliceLoadLevelInformation) SetNumOfPduSess(v NumberAverage)`

SetNumOfPduSess sets NumOfPduSess field to given value.

### HasNumOfPduSess

`func (o *SliceLoadLevelInformation) HasNumOfPduSess() bool`

HasNumOfPduSess returns a boolean if a field has been set.

### GetExceedLoadLevelThrInd

`func (o *SliceLoadLevelInformation) GetExceedLoadLevelThrInd() bool`

GetExceedLoadLevelThrInd returns the ExceedLoadLevelThrInd field if non-nil, zero value otherwise.

### GetExceedLoadLevelThrIndOk

`func (o *SliceLoadLevelInformation) GetExceedLoadLevelThrIndOk() (*bool, bool)`

GetExceedLoadLevelThrIndOk returns a tuple with the ExceedLoadLevelThrInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceedLoadLevelThrInd

`func (o *SliceLoadLevelInformation) SetExceedLoadLevelThrInd(v bool)`

SetExceedLoadLevelThrInd sets ExceedLoadLevelThrInd field to given value.

### HasExceedLoadLevelThrInd

`func (o *SliceLoadLevelInformation) HasExceedLoadLevelThrInd() bool`

HasExceedLoadLevelThrInd returns a boolean if a field has been set.

### GetConfidence

`func (o *SliceLoadLevelInformation) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SliceLoadLevelInformation) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SliceLoadLevelInformation) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *SliceLoadLevelInformation) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


