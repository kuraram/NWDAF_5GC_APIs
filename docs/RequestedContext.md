# RequestedContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contexts** | [**[]ContextType**](ContextType.md) | List of analytics context types. | 

## Methods

### NewRequestedContext

`func NewRequestedContext(contexts []ContextType, ) *RequestedContext`

NewRequestedContext instantiates a new RequestedContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedContextWithDefaults

`func NewRequestedContextWithDefaults() *RequestedContext`

NewRequestedContextWithDefaults instantiates a new RequestedContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContexts

`func (o *RequestedContext) GetContexts() []ContextType`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *RequestedContext) GetContextsOk() (*[]ContextType, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *RequestedContext) SetContexts(v []ContextType)`

SetContexts sets Contexts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


