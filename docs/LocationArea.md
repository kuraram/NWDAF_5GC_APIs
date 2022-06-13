# LocationArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeographicAreas** | Pointer to [**[]GeographicArea**](GeographicArea.md) | Identifies a list of geographic area of the user where the UE is located. | [optional] 
**CivicAddresses** | Pointer to [**[]CivicAddress**](CivicAddress.md) | Identifies a list of civic addresses of the user where the UE is located. | [optional] 
**NwAreaInfo** | Pointer to [**NetworkAreaInfo1**](NetworkAreaInfo1.md) |  | [optional] 
**UmtTime** | Pointer to [**UmtTime**](UmtTime.md) |  | [optional] 

## Methods

### NewLocationArea

`func NewLocationArea() *LocationArea`

NewLocationArea instantiates a new LocationArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationAreaWithDefaults

`func NewLocationAreaWithDefaults() *LocationArea`

NewLocationAreaWithDefaults instantiates a new LocationArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeographicAreas

`func (o *LocationArea) GetGeographicAreas() []GeographicArea`

GetGeographicAreas returns the GeographicAreas field if non-nil, zero value otherwise.

### GetGeographicAreasOk

`func (o *LocationArea) GetGeographicAreasOk() (*[]GeographicArea, bool)`

GetGeographicAreasOk returns a tuple with the GeographicAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicAreas

`func (o *LocationArea) SetGeographicAreas(v []GeographicArea)`

SetGeographicAreas sets GeographicAreas field to given value.

### HasGeographicAreas

`func (o *LocationArea) HasGeographicAreas() bool`

HasGeographicAreas returns a boolean if a field has been set.

### GetCivicAddresses

`func (o *LocationArea) GetCivicAddresses() []CivicAddress`

GetCivicAddresses returns the CivicAddresses field if non-nil, zero value otherwise.

### GetCivicAddressesOk

`func (o *LocationArea) GetCivicAddressesOk() (*[]CivicAddress, bool)`

GetCivicAddressesOk returns a tuple with the CivicAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddresses

`func (o *LocationArea) SetCivicAddresses(v []CivicAddress)`

SetCivicAddresses sets CivicAddresses field to given value.

### HasCivicAddresses

`func (o *LocationArea) HasCivicAddresses() bool`

HasCivicAddresses returns a boolean if a field has been set.

### GetNwAreaInfo

`func (o *LocationArea) GetNwAreaInfo() NetworkAreaInfo1`

GetNwAreaInfo returns the NwAreaInfo field if non-nil, zero value otherwise.

### GetNwAreaInfoOk

`func (o *LocationArea) GetNwAreaInfoOk() (*NetworkAreaInfo1, bool)`

GetNwAreaInfoOk returns a tuple with the NwAreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwAreaInfo

`func (o *LocationArea) SetNwAreaInfo(v NetworkAreaInfo1)`

SetNwAreaInfo sets NwAreaInfo field to given value.

### HasNwAreaInfo

`func (o *LocationArea) HasNwAreaInfo() bool`

HasNwAreaInfo returns a boolean if a field has been set.

### GetUmtTime

`func (o *LocationArea) GetUmtTime() UmtTime`

GetUmtTime returns the UmtTime field if non-nil, zero value otherwise.

### GetUmtTimeOk

`func (o *LocationArea) GetUmtTimeOk() (*UmtTime, bool)`

GetUmtTimeOk returns a tuple with the UmtTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmtTime

`func (o *LocationArea) SetUmtTime(v UmtTime)`

SetUmtTime sets UmtTime field to given value.

### HasUmtTime

`func (o *LocationArea) HasUmtTime() bool`

HasUmtTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


