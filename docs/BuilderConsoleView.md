# BuilderConsoleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int32** |  | [optional] 
**AttemptId** | Pointer to **int32** |  | [optional] 
**DispatchedAt** | Pointer to **string** |  | [optional] 
**Failure** | Pointer to [**BuilderBuildFailureView**](BuilderBuildFailureView.md) |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**LogCursor** | Pointer to **int32** |  | [optional] 
**LogStatus** | Pointer to **string** | LogStatus is \&quot;starting\&quot; until the first line lands (A28), then \&quot;streaming\&quot; while the attempt is live; a TERMINAL attempt says so per the ledger fact — \&quot;complete\&quot; (succeeded) or \&quot;failed\&quot; — a finished build&#39;s pane never reads as still-streaming. LogCursor is the tail page&#39;s last seq — the cursor the console&#39;s next Since poll passes (the page contract, not a MAX read). | [optional] 
**LogTail** | Pointer to [**[]BuilderConsoleLogLine**](BuilderConsoleLogLine.md) |  | [optional] 
**MachineAgentId** | Pointer to **int32** |  | [optional] 
**PullFailure** | Pointer to [**BuilderPullFailureView**](BuilderPullFailureView.md) |  | [optional] 
**Revision** | Pointer to [**BuilderRevisionSummary**](BuilderRevisionSummary.md) |  | [optional] 
**RunReference** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | State is the ledger&#39;s own state for the attempt (submitting / dispatched / lost / succeeded / failed) — proven, never mapped. | [optional] 
**Steps** | Pointer to [**[]BuilderConsoleStep**](BuilderConsoleStep.md) |  | [optional] 

## Methods

### NewBuilderConsoleView

`func NewBuilderConsoleView() *BuilderConsoleView`

NewBuilderConsoleView instantiates a new BuilderConsoleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuilderConsoleViewWithDefaults

`func NewBuilderConsoleViewWithDefaults() *BuilderConsoleView`

NewBuilderConsoleViewWithDefaults instantiates a new BuilderConsoleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *BuilderConsoleView) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *BuilderConsoleView) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *BuilderConsoleView) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *BuilderConsoleView) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetAttemptId

`func (o *BuilderConsoleView) GetAttemptId() int32`

GetAttemptId returns the AttemptId field if non-nil, zero value otherwise.

### GetAttemptIdOk

`func (o *BuilderConsoleView) GetAttemptIdOk() (*int32, bool)`

GetAttemptIdOk returns a tuple with the AttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptId

`func (o *BuilderConsoleView) SetAttemptId(v int32)`

SetAttemptId sets AttemptId field to given value.

### HasAttemptId

`func (o *BuilderConsoleView) HasAttemptId() bool`

HasAttemptId returns a boolean if a field has been set.

### GetDispatchedAt

`func (o *BuilderConsoleView) GetDispatchedAt() string`

GetDispatchedAt returns the DispatchedAt field if non-nil, zero value otherwise.

### GetDispatchedAtOk

`func (o *BuilderConsoleView) GetDispatchedAtOk() (*string, bool)`

GetDispatchedAtOk returns a tuple with the DispatchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchedAt

`func (o *BuilderConsoleView) SetDispatchedAt(v string)`

SetDispatchedAt sets DispatchedAt field to given value.

### HasDispatchedAt

`func (o *BuilderConsoleView) HasDispatchedAt() bool`

HasDispatchedAt returns a boolean if a field has been set.

### GetFailure

`func (o *BuilderConsoleView) GetFailure() BuilderBuildFailureView`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *BuilderConsoleView) GetFailureOk() (*BuilderBuildFailureView, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *BuilderConsoleView) SetFailure(v BuilderBuildFailureView)`

SetFailure sets Failure field to given value.

### HasFailure

`func (o *BuilderConsoleView) HasFailure() bool`

HasFailure returns a boolean if a field has been set.

### GetJobId

`func (o *BuilderConsoleView) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BuilderConsoleView) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BuilderConsoleView) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BuilderConsoleView) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetLogCursor

`func (o *BuilderConsoleView) GetLogCursor() int32`

GetLogCursor returns the LogCursor field if non-nil, zero value otherwise.

### GetLogCursorOk

`func (o *BuilderConsoleView) GetLogCursorOk() (*int32, bool)`

GetLogCursorOk returns a tuple with the LogCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCursor

`func (o *BuilderConsoleView) SetLogCursor(v int32)`

SetLogCursor sets LogCursor field to given value.

### HasLogCursor

`func (o *BuilderConsoleView) HasLogCursor() bool`

HasLogCursor returns a boolean if a field has been set.

### GetLogStatus

`func (o *BuilderConsoleView) GetLogStatus() string`

GetLogStatus returns the LogStatus field if non-nil, zero value otherwise.

### GetLogStatusOk

`func (o *BuilderConsoleView) GetLogStatusOk() (*string, bool)`

GetLogStatusOk returns a tuple with the LogStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStatus

`func (o *BuilderConsoleView) SetLogStatus(v string)`

SetLogStatus sets LogStatus field to given value.

### HasLogStatus

`func (o *BuilderConsoleView) HasLogStatus() bool`

HasLogStatus returns a boolean if a field has been set.

### GetLogTail

`func (o *BuilderConsoleView) GetLogTail() []BuilderConsoleLogLine`

GetLogTail returns the LogTail field if non-nil, zero value otherwise.

### GetLogTailOk

`func (o *BuilderConsoleView) GetLogTailOk() (*[]BuilderConsoleLogLine, bool)`

GetLogTailOk returns a tuple with the LogTail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTail

`func (o *BuilderConsoleView) SetLogTail(v []BuilderConsoleLogLine)`

SetLogTail sets LogTail field to given value.

### HasLogTail

`func (o *BuilderConsoleView) HasLogTail() bool`

HasLogTail returns a boolean if a field has been set.

### GetMachineAgentId

`func (o *BuilderConsoleView) GetMachineAgentId() int32`

GetMachineAgentId returns the MachineAgentId field if non-nil, zero value otherwise.

### GetMachineAgentIdOk

`func (o *BuilderConsoleView) GetMachineAgentIdOk() (*int32, bool)`

GetMachineAgentIdOk returns a tuple with the MachineAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineAgentId

`func (o *BuilderConsoleView) SetMachineAgentId(v int32)`

SetMachineAgentId sets MachineAgentId field to given value.

### HasMachineAgentId

`func (o *BuilderConsoleView) HasMachineAgentId() bool`

HasMachineAgentId returns a boolean if a field has been set.

### GetPullFailure

`func (o *BuilderConsoleView) GetPullFailure() BuilderPullFailureView`

GetPullFailure returns the PullFailure field if non-nil, zero value otherwise.

### GetPullFailureOk

`func (o *BuilderConsoleView) GetPullFailureOk() (*BuilderPullFailureView, bool)`

GetPullFailureOk returns a tuple with the PullFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullFailure

`func (o *BuilderConsoleView) SetPullFailure(v BuilderPullFailureView)`

SetPullFailure sets PullFailure field to given value.

### HasPullFailure

`func (o *BuilderConsoleView) HasPullFailure() bool`

HasPullFailure returns a boolean if a field has been set.

### GetRevision

`func (o *BuilderConsoleView) GetRevision() BuilderRevisionSummary`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BuilderConsoleView) GetRevisionOk() (*BuilderRevisionSummary, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BuilderConsoleView) SetRevision(v BuilderRevisionSummary)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BuilderConsoleView) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRunReference

`func (o *BuilderConsoleView) GetRunReference() string`

GetRunReference returns the RunReference field if non-nil, zero value otherwise.

### GetRunReferenceOk

`func (o *BuilderConsoleView) GetRunReferenceOk() (*string, bool)`

GetRunReferenceOk returns a tuple with the RunReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunReference

`func (o *BuilderConsoleView) SetRunReference(v string)`

SetRunReference sets RunReference field to given value.

### HasRunReference

`func (o *BuilderConsoleView) HasRunReference() bool`

HasRunReference returns a boolean if a field has been set.

### GetState

`func (o *BuilderConsoleView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BuilderConsoleView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BuilderConsoleView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BuilderConsoleView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSteps

`func (o *BuilderConsoleView) GetSteps() []BuilderConsoleStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *BuilderConsoleView) GetStepsOk() (*[]BuilderConsoleStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *BuilderConsoleView) SetSteps(v []BuilderConsoleStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *BuilderConsoleView) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


