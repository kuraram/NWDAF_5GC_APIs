# GeraLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationNumber** | Pointer to **string** |  | [optional] 
**Cgi** | Pointer to [**CellGlobalId**](CellGlobalId.md) |  | [optional] 
**Rai** | Pointer to [**RoutingAreaId**](RoutingAreaId.md) |  | [optional] 
**Sai** | Pointer to [**ServiceAreaId**](ServiceAreaId.md) |  | [optional] 
**Lai** | Pointer to [**LocationAreaId**](LocationAreaId.md) |  | [optional] 
**VlrNumber** | Pointer to **string** |  | [optional] 
**MscNumber** | Pointer to **string** |  | [optional] 
**AgeOfLocationInformation** | Pointer to **int32** |  | [optional] 
**UeLocationTimestamp** | Pointer to **time.Time** |  | [optional] 
**GeographicalInformation** | Pointer to **string** |  | [optional] 
**GeodeticInformation** | Pointer to **string** |  | [optional] 

## Methods

### NewGeraLocation

`func NewGeraLocation() *GeraLocation`

NewGeraLocation instantiates a new GeraLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeraLocationWithDefaults

`func NewGeraLocationWithDefaults() *GeraLocation`

NewGeraLocationWithDefaults instantiates a new GeraLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationNumber

`func (o *GeraLocation) GetLocationNumber() string`

GetLocationNumber returns the LocationNumber field if non-nil, zero value otherwise.

### GetLocationNumberOk

`func (o *GeraLocation) GetLocationNumberOk() (*string, bool)`

GetLocationNumberOk returns a tuple with the LocationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationNumber

`func (o *GeraLocation) SetLocationNumber(v string)`

SetLocationNumber sets LocationNumber field to given value.

### HasLocationNumber

`func (o *GeraLocation) HasLocationNumber() bool`

HasLocationNumber returns a boolean if a field has been set.

### GetCgi

`func (o *GeraLocation) GetCgi() CellGlobalId`

GetCgi returns the Cgi field if non-nil, zero value otherwise.

### GetCgiOk

`func (o *GeraLocation) GetCgiOk() (*CellGlobalId, bool)`

GetCgiOk returns a tuple with the Cgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgi

`func (o *GeraLocation) SetCgi(v CellGlobalId)`

SetCgi sets Cgi field to given value.

### HasCgi

`func (o *GeraLocation) HasCgi() bool`

HasCgi returns a boolean if a field has been set.

### GetRai

`func (o *GeraLocation) GetRai() RoutingAreaId`

GetRai returns the Rai field if non-nil, zero value otherwise.

### GetRaiOk

`func (o *GeraLocation) GetRaiOk() (*RoutingAreaId, bool)`

GetRaiOk returns a tuple with the Rai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRai

`func (o *GeraLocation) SetRai(v RoutingAreaId)`

SetRai sets Rai field to given value.

### HasRai

`func (o *GeraLocation) HasRai() bool`

HasRai returns a boolean if a field has been set.

### GetSai

`func (o *GeraLocation) GetSai() ServiceAreaId`

GetSai returns the Sai field if non-nil, zero value otherwise.

### GetSaiOk

`func (o *GeraLocation) GetSaiOk() (*ServiceAreaId, bool)`

GetSaiOk returns a tuple with the Sai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSai

`func (o *GeraLocation) SetSai(v ServiceAreaId)`

SetSai sets Sai field to given value.

### HasSai

`func (o *GeraLocation) HasSai() bool`

HasSai returns a boolean if a field has been set.

### GetLai

`func (o *GeraLocation) GetLai() LocationAreaId`

GetLai returns the Lai field if non-nil, zero value otherwise.

### GetLaiOk

`func (o *GeraLocation) GetLaiOk() (*LocationAreaId, bool)`

GetLaiOk returns a tuple with the Lai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLai

`func (o *GeraLocation) SetLai(v LocationAreaId)`

SetLai sets Lai field to given value.

### HasLai

`func (o *GeraLocation) HasLai() bool`

HasLai returns a boolean if a field has been set.

### GetVlrNumber

`func (o *GeraLocation) GetVlrNumber() string`

GetVlrNumber returns the VlrNumber field if non-nil, zero value otherwise.

### GetVlrNumberOk

`func (o *GeraLocation) GetVlrNumberOk() (*string, bool)`

GetVlrNumberOk returns a tuple with the VlrNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlrNumber

`func (o *GeraLocation) SetVlrNumber(v string)`

SetVlrNumber sets VlrNumber field to given value.

### HasVlrNumber

`func (o *GeraLocation) HasVlrNumber() bool`

HasVlrNumber returns a boolean if a field has been set.

### GetMscNumber

`func (o *GeraLocation) GetMscNumber() string`

GetMscNumber returns the MscNumber field if non-nil, zero value otherwise.

### GetMscNumberOk

`func (o *GeraLocation) GetMscNumberOk() (*string, bool)`

GetMscNumberOk returns a tuple with the MscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMscNumber

`func (o *GeraLocation) SetMscNumber(v string)`

SetMscNumber sets MscNumber field to given value.

### HasMscNumber

`func (o *GeraLocation) HasMscNumber() bool`

HasMscNumber returns a boolean if a field has been set.

### GetAgeOfLocationInformation

`func (o *GeraLocation) GetAgeOfLocationInformation() int32`

GetAgeOfLocationInformation returns the AgeOfLocationInformation field if non-nil, zero value otherwise.

### GetAgeOfLocationInformationOk

`func (o *GeraLocation) GetAgeOfLocationInformationOk() (*int32, bool)`

GetAgeOfLocationInformationOk returns a tuple with the AgeOfLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationInformation

`func (o *GeraLocation) SetAgeOfLocationInformation(v int32)`

SetAgeOfLocationInformation sets AgeOfLocationInformation field to given value.

### HasAgeOfLocationInformation

`func (o *GeraLocation) HasAgeOfLocationInformation() bool`

HasAgeOfLocationInformation returns a boolean if a field has been set.

### GetUeLocationTimestamp

`func (o *GeraLocation) GetUeLocationTimestamp() time.Time`

GetUeLocationTimestamp returns the UeLocationTimestamp field if non-nil, zero value otherwise.

### GetUeLocationTimestampOk

`func (o *GeraLocation) GetUeLocationTimestampOk() (*time.Time, bool)`

GetUeLocationTimestampOk returns a tuple with the UeLocationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocationTimestamp

`func (o *GeraLocation) SetUeLocationTimestamp(v time.Time)`

SetUeLocationTimestamp sets UeLocationTimestamp field to given value.

### HasUeLocationTimestamp

`func (o *GeraLocation) HasUeLocationTimestamp() bool`

HasUeLocationTimestamp returns a boolean if a field has been set.

### GetGeographicalInformation

`func (o *GeraLocation) GetGeographicalInformation() string`

GetGeographicalInformation returns the GeographicalInformation field if non-nil, zero value otherwise.

### GetGeographicalInformationOk

`func (o *GeraLocation) GetGeographicalInformationOk() (*string, bool)`

GetGeographicalInformationOk returns a tuple with the GeographicalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicalInformation

`func (o *GeraLocation) SetGeographicalInformation(v string)`

SetGeographicalInformation sets GeographicalInformation field to given value.

### HasGeographicalInformation

`func (o *GeraLocation) HasGeographicalInformation() bool`

HasGeographicalInformation returns a boolean if a field has been set.

### GetGeodeticInformation

`func (o *GeraLocation) GetGeodeticInformation() string`

GetGeodeticInformation returns the GeodeticInformation field if non-nil, zero value otherwise.

### GetGeodeticInformationOk

`func (o *GeraLocation) GetGeodeticInformationOk() (*string, bool)`

GetGeodeticInformationOk returns a tuple with the GeodeticInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeodeticInformation

`func (o *GeraLocation) SetGeodeticInformation(v string)`

SetGeodeticInformation sets GeodeticInformation field to given value.

### HasGeodeticInformation

`func (o *GeraLocation) HasGeodeticInformation() bool`

HasGeodeticInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


