# NetworkAreaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ecgis** | Pointer to [**[]Ecgi**](Ecgi.md) | Contains a list of E-UTRA cell identities. | [optional] 
**Ncgis** | Pointer to [**[]Ncgi**](Ncgi.md) | Contains a list of NR cell identities. | [optional] 
**GRanNodeIds** | Pointer to [**[]GlobalRanNodeId**](GlobalRanNodeId.md) | Contains a list of NG RAN nodes. | [optional] 
**Tais** | Pointer to [**[]Tai**](Tai.md) | Contains a list of tracking area identities. | [optional] 

## Methods

### NewNetworkAreaInfo

`func NewNetworkAreaInfo() *NetworkAreaInfo`

NewNetworkAreaInfo instantiates a new NetworkAreaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkAreaInfoWithDefaults

`func NewNetworkAreaInfoWithDefaults() *NetworkAreaInfo`

NewNetworkAreaInfoWithDefaults instantiates a new NetworkAreaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcgis

`func (o *NetworkAreaInfo) GetEcgis() []Ecgi`

GetEcgis returns the Ecgis field if non-nil, zero value otherwise.

### GetEcgisOk

`func (o *NetworkAreaInfo) GetEcgisOk() (*[]Ecgi, bool)`

GetEcgisOk returns a tuple with the Ecgis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgis

`func (o *NetworkAreaInfo) SetEcgis(v []Ecgi)`

SetEcgis sets Ecgis field to given value.

### HasEcgis

`func (o *NetworkAreaInfo) HasEcgis() bool`

HasEcgis returns a boolean if a field has been set.

### GetNcgis

`func (o *NetworkAreaInfo) GetNcgis() []Ncgi`

GetNcgis returns the Ncgis field if non-nil, zero value otherwise.

### GetNcgisOk

`func (o *NetworkAreaInfo) GetNcgisOk() (*[]Ncgi, bool)`

GetNcgisOk returns a tuple with the Ncgis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgis

`func (o *NetworkAreaInfo) SetNcgis(v []Ncgi)`

SetNcgis sets Ncgis field to given value.

### HasNcgis

`func (o *NetworkAreaInfo) HasNcgis() bool`

HasNcgis returns a boolean if a field has been set.

### GetGRanNodeIds

`func (o *NetworkAreaInfo) GetGRanNodeIds() []GlobalRanNodeId`

GetGRanNodeIds returns the GRanNodeIds field if non-nil, zero value otherwise.

### GetGRanNodeIdsOk

`func (o *NetworkAreaInfo) GetGRanNodeIdsOk() (*[]GlobalRanNodeId, bool)`

GetGRanNodeIdsOk returns a tuple with the GRanNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGRanNodeIds

`func (o *NetworkAreaInfo) SetGRanNodeIds(v []GlobalRanNodeId)`

SetGRanNodeIds sets GRanNodeIds field to given value.

### HasGRanNodeIds

`func (o *NetworkAreaInfo) HasGRanNodeIds() bool`

HasGRanNodeIds returns a boolean if a field has been set.

### GetTais

`func (o *NetworkAreaInfo) GetTais() []Tai`

GetTais returns the Tais field if non-nil, zero value otherwise.

### GetTaisOk

`func (o *NetworkAreaInfo) GetTaisOk() (*[]Tai, bool)`

GetTaisOk returns a tuple with the Tais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTais

`func (o *NetworkAreaInfo) SetTais(v []Tai)`

SetTais sets Tais field to given value.

### HasTais

`func (o *NetworkAreaInfo) HasTais() bool`

HasTais returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


