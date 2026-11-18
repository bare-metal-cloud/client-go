# HandlersStackRollbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**TargetRevisionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersStackRollbackRequest

`func NewHandlersStackRollbackRequest() *HandlersStackRollbackRequest`

NewHandlersStackRollbackRequest instantiates a new HandlersStackRollbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStackRollbackRequestWithDefaults

`func NewHandlersStackRollbackRequestWithDefaults() *HandlersStackRollbackRequest`

NewHandlersStackRollbackRequestWithDefaults instantiates a new HandlersStackRollbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *HandlersStackRollbackRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HandlersStackRollbackRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HandlersStackRollbackRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HandlersStackRollbackRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetService

`func (o *HandlersStackRollbackRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *HandlersStackRollbackRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *HandlersStackRollbackRequest) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *HandlersStackRollbackRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStack

`func (o *HandlersStackRollbackRequest) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *HandlersStackRollbackRequest) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *HandlersStackRollbackRequest) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *HandlersStackRollbackRequest) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetTargetRevisionId

`func (o *HandlersStackRollbackRequest) GetTargetRevisionId() int32`

GetTargetRevisionId returns the TargetRevisionId field if non-nil, zero value otherwise.

### GetTargetRevisionIdOk

`func (o *HandlersStackRollbackRequest) GetTargetRevisionIdOk() (*int32, bool)`

GetTargetRevisionIdOk returns a tuple with the TargetRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRevisionId

`func (o *HandlersStackRollbackRequest) SetTargetRevisionId(v int32)`

SetTargetRevisionId sets TargetRevisionId field to given value.

### HasTargetRevisionId

`func (o *HandlersStackRollbackRequest) HasTargetRevisionId() bool`

HasTargetRevisionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


