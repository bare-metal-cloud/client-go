# HandlersRunnerBuildCompletionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptRef** | Pointer to **string** |  | [optional] 
**Duplicate** | Pointer to **bool** | Duplicate is true when this completion is the idempotent replay of an already-registered attempt — the SAME revision came back and no second row exists. | [optional] 
**ImageRef** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerBuildCompletionResponse

`func NewHandlersRunnerBuildCompletionResponse() *HandlersRunnerBuildCompletionResponse`

NewHandlersRunnerBuildCompletionResponse instantiates a new HandlersRunnerBuildCompletionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerBuildCompletionResponseWithDefaults

`func NewHandlersRunnerBuildCompletionResponseWithDefaults() *HandlersRunnerBuildCompletionResponse`

NewHandlersRunnerBuildCompletionResponseWithDefaults instantiates a new HandlersRunnerBuildCompletionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptRef

`func (o *HandlersRunnerBuildCompletionResponse) GetAttemptRef() string`

GetAttemptRef returns the AttemptRef field if non-nil, zero value otherwise.

### GetAttemptRefOk

`func (o *HandlersRunnerBuildCompletionResponse) GetAttemptRefOk() (*string, bool)`

GetAttemptRefOk returns a tuple with the AttemptRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptRef

`func (o *HandlersRunnerBuildCompletionResponse) SetAttemptRef(v string)`

SetAttemptRef sets AttemptRef field to given value.

### HasAttemptRef

`func (o *HandlersRunnerBuildCompletionResponse) HasAttemptRef() bool`

HasAttemptRef returns a boolean if a field has been set.

### GetDuplicate

`func (o *HandlersRunnerBuildCompletionResponse) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *HandlersRunnerBuildCompletionResponse) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *HandlersRunnerBuildCompletionResponse) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *HandlersRunnerBuildCompletionResponse) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetImageRef

`func (o *HandlersRunnerBuildCompletionResponse) GetImageRef() string`

GetImageRef returns the ImageRef field if non-nil, zero value otherwise.

### GetImageRefOk

`func (o *HandlersRunnerBuildCompletionResponse) GetImageRefOk() (*string, bool)`

GetImageRefOk returns a tuple with the ImageRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRef

`func (o *HandlersRunnerBuildCompletionResponse) SetImageRef(v string)`

SetImageRef sets ImageRef field to given value.

### HasImageRef

`func (o *HandlersRunnerBuildCompletionResponse) HasImageRef() bool`

HasImageRef returns a boolean if a field has been set.

### GetRevision

`func (o *HandlersRunnerBuildCompletionResponse) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *HandlersRunnerBuildCompletionResponse) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *HandlersRunnerBuildCompletionResponse) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *HandlersRunnerBuildCompletionResponse) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetService

`func (o *HandlersRunnerBuildCompletionResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *HandlersRunnerBuildCompletionResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *HandlersRunnerBuildCompletionResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *HandlersRunnerBuildCompletionResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStack

`func (o *HandlersRunnerBuildCompletionResponse) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *HandlersRunnerBuildCompletionResponse) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *HandlersRunnerBuildCompletionResponse) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *HandlersRunnerBuildCompletionResponse) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


