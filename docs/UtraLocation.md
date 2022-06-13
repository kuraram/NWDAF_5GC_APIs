# UtraLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cgi** | Pointer to [**CellGlobalId**](CellGlobalId.md) |  | [optional] 
**Sai** | Pointer to [**ServiceAreaId**](ServiceAreaId.md) |  | [optional] 
**Lai** | Pointer to [**LocationAreaId**](LocationAreaId.md) |  | [optional] 
**Rai** | Pointer to [**RoutingAreaId**](RoutingAreaId.md) |  | [optional] 
**AgeOfLocationInformation** | Pointer to **int32** | The value represents the elapsed time in minutes since the last network contact of the mobile station.  Value \&quot;0\&quot; indicates that the location information was obtained after a successful paging procedure for  Active Location Retrieval when the UE is in idle mode or after a successful location reporting procedure  the UE is in connected mode.Any other value than \&quot;0\&quot; indicates that the location information is the last known one.  See 3GPP TS 29.002 clause 17.7.8.  | [optional] 
**UeLocationTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**GeographicalInformation** | Pointer to **string** | Refer to geographical Information.See 3GPP TS 23.032 clause 7.3.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used. | [optional] 
**GeodeticInformation** | Pointer to **string** | Refers to Calling Geodetic Location. See ITU-T Recommendation Q.763 (1999) clause 3.88.2.  Only the description of an ellipsoid point with uncertainty circle is allowed to be used.  | [optional] 

## Methods

### NewUtraLocation

`func NewUtraLocation() *UtraLocation`

NewUtraLocation instantiates a new UtraLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtraLocationWithDefaults

`func NewUtraLocationWithDefaults() *UtraLocation`

NewUtraLocationWithDefaults instantiates a new UtraLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCgi

`func (o *UtraLocation) GetCgi() CellGlobalId`

GetCgi returns the Cgi field if non-nil, zero value otherwise.

### GetCgiOk

`func (o *UtraLocation) GetCgiOk() (*CellGlobalId, bool)`

GetCgiOk returns a tuple with the Cgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgi

`func (o *UtraLocation) SetCgi(v CellGlobalId)`

SetCgi sets Cgi field to given value.

### HasCgi

`func (o *UtraLocation) HasCgi() bool`

HasCgi returns a boolean if a field has been set.

### GetSai

`func (o *UtraLocation) GetSai() ServiceAreaId`

GetSai returns the Sai field if non-nil, zero value otherwise.

### GetSaiOk

`func (o *UtraLocation) GetSaiOk() (*ServiceAreaId, bool)`

GetSaiOk returns a tuple with the Sai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSai

`func (o *UtraLocation) SetSai(v ServiceAreaId)`

SetSai sets Sai field to given value.

### HasSai

`func (o *UtraLocation) HasSai() bool`

HasSai returns a boolean if a field has been set.

### GetLai

`func (o *UtraLocation) GetLai() LocationAreaId`

GetLai returns the Lai field if non-nil, zero value otherwise.

### GetLaiOk

`func (o *UtraLocation) GetLaiOk() (*LocationAreaId, bool)`

GetLaiOk returns a tuple with the Lai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLai

`func (o *UtraLocation) SetLai(v LocationAreaId)`

SetLai sets Lai field to given value.

### HasLai

`func (o *UtraLocation) HasLai() bool`

HasLai returns a boolean if a field has been set.

### GetRai

`func (o *UtraLocation) GetRai() RoutingAreaId`

GetRai returns the Rai field if non-nil, zero value otherwise.

### GetRaiOk

`func (o *UtraLocation) GetRaiOk() (*RoutingAreaId, bool)`

GetRaiOk returns a tuple with the Rai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRai

`func (o *UtraLocation) SetRai(v RoutingAreaId)`

SetRai sets Rai field to given value.

### HasRai

`func (o *UtraLocation) HasRai() bool`

HasRai returns a boolean if a field has been set.

### GetAgeOfLocationInformation

`func (o *UtraLocation) GetAgeOfLocationInformation() int32`

GetAgeOfLocationInformation returns the AgeOfLocationInformation field if non-nil, zero value otherwise.

### GetAgeOfLocationInformationOk

`func (o *UtraLocation) GetAgeOfLocationInformationOk() (*int32, bool)`

GetAgeOfLocationInformationOk returns a tuple with the AgeOfLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationInformation

`func (o *UtraLocation) SetAgeOfLocationInformation(v int32)`

SetAgeOfLocationInformation sets AgeOfLocationInformation field to given value.

### HasAgeOfLocationInformation

`func (o *UtraLocation) HasAgeOfLocationInformation() bool`

HasAgeOfLocationInformation returns a boolean if a field has been set.

### GetUeLocationTimestamp

`func (o *UtraLocation) GetUeLocationTimestamp() time.Time`

GetUeLocationTimestamp returns the UeLocationTimestamp field if non-nil, zero value otherwise.

### GetUeLocationTimestampOk

`func (o *UtraLocation) GetUeLocationTimestampOk() (*time.Time, bool)`

GetUeLocationTimestampOk returns a tuple with the UeLocationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocationTimestamp

`func (o *UtraLocation) SetUeLocationTimestamp(v time.Time)`

SetUeLocationTimestamp sets UeLocationTimestamp field to given value.

### HasUeLocationTimestamp

`func (o *UtraLocation) HasUeLocationTimestamp() bool`

HasUeLocationTimestamp returns a boolean if a field has been set.

### GetGeographicalInformation

`func (o *UtraLocation) GetGeographicalInformation() string`

GetGeographicalInformation returns the GeographicalInformation field if non-nil, zero value otherwise.

### GetGeographicalInformationOk

`func (o *UtraLocation) GetGeographicalInformationOk() (*string, bool)`

GetGeographicalInformationOk returns a tuple with the GeographicalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicalInformation

`func (o *UtraLocation) SetGeographicalInformation(v string)`

SetGeographicalInformation sets GeographicalInformation field to given value.

### HasGeographicalInformation

`func (o *UtraLocation) HasGeographicalInformation() bool`

HasGeographicalInformation returns a boolean if a field has been set.

### GetGeodeticInformation

`func (o *UtraLocation) GetGeodeticInformation() string`

GetGeodeticInformation returns the GeodeticInformation field if non-nil, zero value otherwise.

### GetGeodeticInformationOk

`func (o *UtraLocation) GetGeodeticInformationOk() (*string, bool)`

GetGeodeticInformationOk returns a tuple with the GeodeticInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeodeticInformation

`func (o *UtraLocation) SetGeodeticInformation(v string)`

SetGeodeticInformation sets GeodeticInformation field to given value.

### HasGeodeticInformation

`func (o *UtraLocation) HasGeodeticInformation() bool`

HasGeodeticInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


