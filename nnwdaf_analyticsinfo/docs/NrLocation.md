# NrLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tai** | [**Tai**](Tai.md) |  | 
**Ncgi** | [**Ncgi**](Ncgi.md) |  | 
**IgnoreNcgi** | Pointer to **bool** |  | [optional] [default to false]
**AgeOfLocationInformation** | Pointer to **int32** |  | [optional] 
**UeLocationTimestamp** | Pointer to **time.Time** |  | [optional] 
**GeographicalInformation** | Pointer to **string** |  | [optional] 
**GeodeticInformation** | Pointer to **string** |  | [optional] 
**GlobalGnbId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 

## Methods

### NewNrLocation

`func NewNrLocation(tai Tai, ncgi Ncgi, ) *NrLocation`

NewNrLocation instantiates a new NrLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrLocationWithDefaults

`func NewNrLocationWithDefaults() *NrLocation`

NewNrLocationWithDefaults instantiates a new NrLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTai

`func (o *NrLocation) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *NrLocation) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *NrLocation) SetTai(v Tai)`

SetTai sets Tai field to given value.


### GetNcgi

`func (o *NrLocation) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *NrLocation) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *NrLocation) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.


### GetIgnoreNcgi

`func (o *NrLocation) GetIgnoreNcgi() bool`

GetIgnoreNcgi returns the IgnoreNcgi field if non-nil, zero value otherwise.

### GetIgnoreNcgiOk

`func (o *NrLocation) GetIgnoreNcgiOk() (*bool, bool)`

GetIgnoreNcgiOk returns a tuple with the IgnoreNcgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreNcgi

`func (o *NrLocation) SetIgnoreNcgi(v bool)`

SetIgnoreNcgi sets IgnoreNcgi field to given value.

### HasIgnoreNcgi

`func (o *NrLocation) HasIgnoreNcgi() bool`

HasIgnoreNcgi returns a boolean if a field has been set.

### GetAgeOfLocationInformation

`func (o *NrLocation) GetAgeOfLocationInformation() int32`

GetAgeOfLocationInformation returns the AgeOfLocationInformation field if non-nil, zero value otherwise.

### GetAgeOfLocationInformationOk

`func (o *NrLocation) GetAgeOfLocationInformationOk() (*int32, bool)`

GetAgeOfLocationInformationOk returns a tuple with the AgeOfLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationInformation

`func (o *NrLocation) SetAgeOfLocationInformation(v int32)`

SetAgeOfLocationInformation sets AgeOfLocationInformation field to given value.

### HasAgeOfLocationInformation

`func (o *NrLocation) HasAgeOfLocationInformation() bool`

HasAgeOfLocationInformation returns a boolean if a field has been set.

### GetUeLocationTimestamp

`func (o *NrLocation) GetUeLocationTimestamp() time.Time`

GetUeLocationTimestamp returns the UeLocationTimestamp field if non-nil, zero value otherwise.

### GetUeLocationTimestampOk

`func (o *NrLocation) GetUeLocationTimestampOk() (*time.Time, bool)`

GetUeLocationTimestampOk returns a tuple with the UeLocationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocationTimestamp

`func (o *NrLocation) SetUeLocationTimestamp(v time.Time)`

SetUeLocationTimestamp sets UeLocationTimestamp field to given value.

### HasUeLocationTimestamp

`func (o *NrLocation) HasUeLocationTimestamp() bool`

HasUeLocationTimestamp returns a boolean if a field has been set.

### GetGeographicalInformation

`func (o *NrLocation) GetGeographicalInformation() string`

GetGeographicalInformation returns the GeographicalInformation field if non-nil, zero value otherwise.

### GetGeographicalInformationOk

`func (o *NrLocation) GetGeographicalInformationOk() (*string, bool)`

GetGeographicalInformationOk returns a tuple with the GeographicalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicalInformation

`func (o *NrLocation) SetGeographicalInformation(v string)`

SetGeographicalInformation sets GeographicalInformation field to given value.

### HasGeographicalInformation

`func (o *NrLocation) HasGeographicalInformation() bool`

HasGeographicalInformation returns a boolean if a field has been set.

### GetGeodeticInformation

`func (o *NrLocation) GetGeodeticInformation() string`

GetGeodeticInformation returns the GeodeticInformation field if non-nil, zero value otherwise.

### GetGeodeticInformationOk

`func (o *NrLocation) GetGeodeticInformationOk() (*string, bool)`

GetGeodeticInformationOk returns a tuple with the GeodeticInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeodeticInformation

`func (o *NrLocation) SetGeodeticInformation(v string)`

SetGeodeticInformation sets GeodeticInformation field to given value.

### HasGeodeticInformation

`func (o *NrLocation) HasGeodeticInformation() bool`

HasGeodeticInformation returns a boolean if a field has been set.

### GetGlobalGnbId

`func (o *NrLocation) GetGlobalGnbId() GlobalRanNodeId`

GetGlobalGnbId returns the GlobalGnbId field if non-nil, zero value otherwise.

### GetGlobalGnbIdOk

`func (o *NrLocation) GetGlobalGnbIdOk() (*GlobalRanNodeId, bool)`

GetGlobalGnbIdOk returns a tuple with the GlobalGnbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalGnbId

`func (o *NrLocation) SetGlobalGnbId(v GlobalRanNodeId)`

SetGlobalGnbId sets GlobalGnbId field to given value.

### HasGlobalGnbId

`func (o *NrLocation) HasGlobalGnbId() bool`

HasGlobalGnbId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


