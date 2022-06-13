# ApplicationVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | String providing an application identifier. | 
**AppVolume** | **int64** | Unsigned integer identifying a volume in units of bytes. | 

## Methods

### NewApplicationVolume

`func NewApplicationVolume(appId string, appVolume int64, ) *ApplicationVolume`

NewApplicationVolume instantiates a new ApplicationVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationVolumeWithDefaults

`func NewApplicationVolumeWithDefaults() *ApplicationVolume`

NewApplicationVolumeWithDefaults instantiates a new ApplicationVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ApplicationVolume) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ApplicationVolume) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ApplicationVolume) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetAppVolume

`func (o *ApplicationVolume) GetAppVolume() int64`

GetAppVolume returns the AppVolume field if non-nil, zero value otherwise.

### GetAppVolumeOk

`func (o *ApplicationVolume) GetAppVolumeOk() (*int64, bool)`

GetAppVolumeOk returns a tuple with the AppVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVolume

`func (o *ApplicationVolume) SetAppVolume(v int64)`

SetAppVolume sets AppVolume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


