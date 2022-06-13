# AppListForUeComm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**StartTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**AppDur** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**OccurRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent. | [optional] 
**SpatialValidity** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewAppListForUeComm

`func NewAppListForUeComm() *AppListForUeComm`

NewAppListForUeComm instantiates a new AppListForUeComm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppListForUeCommWithDefaults

`func NewAppListForUeCommWithDefaults() *AppListForUeComm`

NewAppListForUeCommWithDefaults instantiates a new AppListForUeComm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AppListForUeComm) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppListForUeComm) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppListForUeComm) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppListForUeComm) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetStartTime

`func (o *AppListForUeComm) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AppListForUeComm) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AppListForUeComm) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AppListForUeComm) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetAppDur

`func (o *AppListForUeComm) GetAppDur() int32`

GetAppDur returns the AppDur field if non-nil, zero value otherwise.

### GetAppDurOk

`func (o *AppListForUeComm) GetAppDurOk() (*int32, bool)`

GetAppDurOk returns a tuple with the AppDur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDur

`func (o *AppListForUeComm) SetAppDur(v int32)`

SetAppDur sets AppDur field to given value.

### HasAppDur

`func (o *AppListForUeComm) HasAppDur() bool`

HasAppDur returns a boolean if a field has been set.

### GetOccurRatio

`func (o *AppListForUeComm) GetOccurRatio() int32`

GetOccurRatio returns the OccurRatio field if non-nil, zero value otherwise.

### GetOccurRatioOk

`func (o *AppListForUeComm) GetOccurRatioOk() (*int32, bool)`

GetOccurRatioOk returns a tuple with the OccurRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurRatio

`func (o *AppListForUeComm) SetOccurRatio(v int32)`

SetOccurRatio sets OccurRatio field to given value.

### HasOccurRatio

`func (o *AppListForUeComm) HasOccurRatio() bool`

HasOccurRatio returns a boolean if a field has been set.

### GetSpatialValidity

`func (o *AppListForUeComm) GetSpatialValidity() NetworkAreaInfo`

GetSpatialValidity returns the SpatialValidity field if non-nil, zero value otherwise.

### GetSpatialValidityOk

`func (o *AppListForUeComm) GetSpatialValidityOk() (*NetworkAreaInfo, bool)`

GetSpatialValidityOk returns a tuple with the SpatialValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpatialValidity

`func (o *AppListForUeComm) SetSpatialValidity(v NetworkAreaInfo)`

SetSpatialValidity sets SpatialValidity field to given value.

### HasSpatialValidity

`func (o *AppListForUeComm) HasSpatialValidity() bool`

HasSpatialValidity returns a boolean if a field has been set.

### GetConfidence

`func (o *AppListForUeComm) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *AppListForUeComm) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *AppListForUeComm) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *AppListForUeComm) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


