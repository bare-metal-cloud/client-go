# HandlersRunnerCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int32** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**ExitCode** | Pointer to **int32** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**RunReference** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerCompletionRequest

`func NewHandlersRunnerCompletionRequest() *HandlersRunnerCompletionRequest`

NewHandlersRunnerCompletionRequest instantiates a new HandlersRunnerCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerCompletionRequestWithDefaults

`func NewHandlersRunnerCompletionRequestWithDefaults() *HandlersRunnerCompletionRequest`

NewHandlersRunnerCompletionRequestWithDefaults instantiates a new HandlersRunnerCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *HandlersRunnerCompletionRequest) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *HandlersRunnerCompletionRequest) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *HandlersRunnerCompletionRequest) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *HandlersRunnerCompletionRequest) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetEndedAt

`func (o *HandlersRunnerCompletionRequest) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *HandlersRunnerCompletionRequest) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *HandlersRunnerCompletionRequest) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *HandlersRunnerCompletionRequest) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetExitCode

`func (o *HandlersRunnerCompletionRequest) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *HandlersRunnerCompletionRequest) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *HandlersRunnerCompletionRequest) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *HandlersRunnerCompletionRequest) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetJobId

`func (o *HandlersRunnerCompletionRequest) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *HandlersRunnerCompletionRequest) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *HandlersRunnerCompletionRequest) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *HandlersRunnerCompletionRequest) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetRunReference

`func (o *HandlersRunnerCompletionRequest) GetRunReference() string`

GetRunReference returns the RunReference field if non-nil, zero value otherwise.

### GetRunReferenceOk

`func (o *HandlersRunnerCompletionRequest) GetRunReferenceOk() (*string, bool)`

GetRunReferenceOk returns a tuple with the RunReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunReference

`func (o *HandlersRunnerCompletionRequest) SetRunReference(v string)`

SetRunReference sets RunReference field to given value.

### HasRunReference

`func (o *HandlersRunnerCompletionRequest) HasRunReference() bool`

HasRunReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


