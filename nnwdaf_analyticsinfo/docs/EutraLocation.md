# EutraLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tai** | [**Tai**](Tai.md) |  | 
**IgnoreTai** | Pointer to **bool** |  | [optional] [default to false]
**Ecgi** | [**Ecgi**](Ecgi.md) |  | 
**IgnoreEcgi** | Pointer to **bool** |  | [optional] [default to false]
**AgeOfLocationInformation** | Pointer to **int32** |  | [optional] 
**UeLocationTimestamp** | Pointer to **time.Time** |  | [optional] 
**GeographicalInformation** | Pointer to **string** |  | [optional] 
**GeodeticInformation** | Pointer to **string** |  | [optional] 
**GlobalNgenbId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**GlobalENbId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 

## Methods

### NewEutraLocation

`func NewEutraLocation(tai Tai, ecgi Ecgi, ) *EutraLocation`

NewEutraLocation instantiates a new EutraLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEutraLocationWithDefaults

`func NewEutraLocationWithDefaults() *EutraLocation`

NewEutraLocationWithDefaults instantiates a new EutraLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTai

`func (o *EutraLocation) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *EutraLocation) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *EutraLocation) SetTai(v Tai)`

SetTai sets Tai field to given value.


### GetIgnoreTai

`func (o *EutraLocation) GetIgnoreTai() bool`

GetIgnoreTai returns the IgnoreTai field if non-nil, zero value otherwise.

### GetIgnoreTaiOk

`func (o *EutraLocation) GetIgnoreTaiOk() (*bool, bool)`

GetIgnoreTaiOk returns a tuple with the IgnoreTai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTai

`func (o *EutraLocation) SetIgnoreTai(v bool)`

SetIgnoreTai sets IgnoreTai field to given value.

### HasIgnoreTai

`func (o *EutraLocation) HasIgnoreTai() bool`

HasIgnoreTai returns a boolean if a field has been set.

### GetEcgi

`func (o *EutraLocation) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *EutraLocation) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *EutraLocation) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.


### GetIgnoreEcgi

`func (o *EutraLocation) GetIgnoreEcgi() bool`

GetIgnoreEcgi returns the IgnoreEcgi field if non-nil, zero value otherwise.

### GetIgnoreEcgiOk

`func (o *EutraLocation) GetIgnoreEcgiOk() (*bool, bool)`

GetIgnoreEcgiOk returns a tuple with the IgnoreEcgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreEcgi

`func (o *EutraLocation) SetIgnoreEcgi(v bool)`

SetIgnoreEcgi sets IgnoreEcgi field to given value.

### HasIgnoreEcgi

`func (o *EutraLocation) HasIgnoreEcgi() bool`

HasIgnoreEcgi returns a boolean if a field has been set.

### GetAgeOfLocationInformation

`func (o *EutraLocation) GetAgeOfLocationInformation() int32`

GetAgeOfLocationInformation returns the AgeOfLocationInformation field if non-nil, zero value otherwise.

### GetAgeOfLocationInformationOk

`func (o *EutraLocation) GetAgeOfLocationInformationOk() (*int32, bool)`

GetAgeOfLocationInformationOk returns a tuple with the AgeOfLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationInformation

`func (o *EutraLocation) SetAgeOfLocationInformation(v int32)`

SetAgeOfLocationInformation sets AgeOfLocationInformation field to given value.

### HasAgeOfLocationInformation

`func (o *EutraLocation) HasAgeOfLocationInformation() bool`

HasAgeOfLocationInformation returns a boolean if a field has been set.

### GetUeLocationTimestamp

`func (o *EutraLocation) GetUeLocationTimestamp() time.Time`

GetUeLocationTimestamp returns the UeLocationTimestamp field if non-nil, zero value otherwise.

### GetUeLocationTimestampOk

`func (o *EutraLocation) GetUeLocationTimestampOk() (*time.Time, bool)`

GetUeLocationTimestampOk returns a tuple with the UeLocationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocationTimestamp

`func (o *EutraLocation) SetUeLocationTimestamp(v time.Time)`

SetUeLocationTimestamp sets UeLocationTimestamp field to given value.

### HasUeLocationTimestamp

`func (o *EutraLocation) HasUeLocationTimestamp() bool`

HasUeLocationTimestamp returns a boolean if a field has been set.

### GetGeographicalInformation

`func (o *EutraLocation) GetGeographicalInformation() string`

GetGeographicalInformation returns the GeographicalInformation field if non-nil, zero value otherwise.

### GetGeographicalInformationOk

`func (o *EutraLocation) GetGeographicalInformationOk() (*string, bool)`

GetGeographicalInformationOk returns a tuple with the GeographicalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicalInformation

`func (o *EutraLocation) SetGeographicalInformation(v string)`

SetGeographicalInformation sets GeographicalInformation field to given value.

### HasGeographicalInformation

`func (o *EutraLocation) HasGeographicalInformation() bool`

HasGeographicalInformation returns a boolean if a field has been set.

### GetGeodeticInformation

`func (o *EutraLocation) GetGeodeticInformation() string`

GetGeodeticInformation returns the GeodeticInformation field if non-nil, zero value otherwise.

### GetGeodeticInformationOk

`func (o *EutraLocation) GetGeodeticInformationOk() (*string, bool)`

GetGeodeticInformationOk returns a tuple with the GeodeticInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeodeticInformation

`func (o *EutraLocation) SetGeodeticInformation(v string)`

SetGeodeticInformation sets GeodeticInformation field to given value.

### HasGeodeticInformation

`func (o *EutraLocation) HasGeodeticInformation() bool`

HasGeodeticInformation returns a boolean if a field has been set.

### GetGlobalNgenbId

`func (o *EutraLocation) GetGlobalNgenbId() GlobalRanNodeId`

GetGlobalNgenbId returns the GlobalNgenbId field if non-nil, zero value otherwise.

### GetGlobalNgenbIdOk

`func (o *EutraLocation) GetGlobalNgenbIdOk() (*GlobalRanNodeId, bool)`

GetGlobalNgenbIdOk returns a tuple with the GlobalNgenbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalNgenbId

`func (o *EutraLocation) SetGlobalNgenbId(v GlobalRanNodeId)`

SetGlobalNgenbId sets GlobalNgenbId field to given value.

### HasGlobalNgenbId

`func (o *EutraLocation) HasGlobalNgenbId() bool`

HasGlobalNgenbId returns a boolean if a field has been set.

### GetGlobalENbId

`func (o *EutraLocation) GetGlobalENbId() GlobalRanNodeId`

GetGlobalENbId returns the GlobalENbId field if non-nil, zero value otherwise.

### GetGlobalENbIdOk

`func (o *EutraLocation) GetGlobalENbIdOk() (*GlobalRanNodeId, bool)`

GetGlobalENbIdOk returns a tuple with the GlobalENbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalENbId

`func (o *EutraLocation) SetGlobalENbId(v GlobalRanNodeId)`

SetGlobalENbId sets GlobalENbId field to given value.

### HasGlobalENbId

`func (o *EutraLocation) HasGlobalENbId() bool`

HasGlobalENbId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


