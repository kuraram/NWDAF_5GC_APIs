# ResourceUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUsage** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MemoryUsage** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**StorageUsage** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewResourceUsage

`func NewResourceUsage() *ResourceUsage`

NewResourceUsage instantiates a new ResourceUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUsageWithDefaults

`func NewResourceUsageWithDefaults() *ResourceUsage`

NewResourceUsageWithDefaults instantiates a new ResourceUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUsage

`func (o *ResourceUsage) GetCpuUsage() int32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ResourceUsage) GetCpuUsageOk() (*int32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ResourceUsage) SetCpuUsage(v int32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *ResourceUsage) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *ResourceUsage) GetMemoryUsage() int32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *ResourceUsage) GetMemoryUsageOk() (*int32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *ResourceUsage) SetMemoryUsage(v int32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *ResourceUsage) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetStorageUsage

`func (o *ResourceUsage) GetStorageUsage() int32`

GetStorageUsage returns the StorageUsage field if non-nil, zero value otherwise.

### GetStorageUsageOk

`func (o *ResourceUsage) GetStorageUsageOk() (*int32, bool)`

GetStorageUsageOk returns a tuple with the StorageUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsage

`func (o *ResourceUsage) SetStorageUsage(v int32)`

SetStorageUsage sets StorageUsage field to given value.

### HasStorageUsage

`func (o *ResourceUsage) HasStorageUsage() bool`

HasStorageUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


