# HandlersRunnerMachineLogsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int32** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**Lines** | Pointer to [**[]HandlersRunnerMachineLogLine**](HandlersRunnerMachineLogLine.md) |  | [optional] 
**RunReference** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerMachineLogsRequest

`func NewHandlersRunnerMachineLogsRequest() *HandlersRunnerMachineLogsRequest`

NewHandlersRunnerMachineLogsRequest instantiates a new HandlersRunnerMachineLogsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerMachineLogsRequestWithDefaults

`func NewHandlersRunnerMachineLogsRequestWithDefaults() *HandlersRunnerMachineLogsRequest`

NewHandlersRunnerMachineLogsRequestWithDefaults instantiates a new HandlersRunnerMachineLogsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *HandlersRunnerMachineLogsRequest) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *HandlersRunnerMachineLogsRequest) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *HandlersRunnerMachineLogsRequest) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *HandlersRunnerMachineLogsRequest) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetJobId

`func (o *HandlersRunnerMachineLogsRequest) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *HandlersRunnerMachineLogsRequest) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *HandlersRunnerMachineLogsRequest) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *HandlersRunnerMachineLogsRequest) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetLines

`func (o *HandlersRunnerMachineLogsRequest) GetLines() []HandlersRunnerMachineLogLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *HandlersRunnerMachineLogsRequest) GetLinesOk() (*[]HandlersRunnerMachineLogLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *HandlersRunnerMachineLogsRequest) SetLines(v []HandlersRunnerMachineLogLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *HandlersRunnerMachineLogsRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetRunReference

`func (o *HandlersRunnerMachineLogsRequest) GetRunReference() string`

GetRunReference returns the RunReference field if non-nil, zero value otherwise.

### GetRunReferenceOk

`func (o *HandlersRunnerMachineLogsRequest) GetRunReferenceOk() (*string, bool)`

GetRunReferenceOk returns a tuple with the RunReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunReference

`func (o *HandlersRunnerMachineLogsRequest) SetRunReference(v string)`

SetRunReference sets RunReference field to given value.

### HasRunReference

`func (o *HandlersRunnerMachineLogsRequest) HasRunReference() bool`

HasRunReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


