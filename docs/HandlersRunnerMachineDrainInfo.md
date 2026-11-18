# HandlersRunnerMachineDrainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deadline** | Pointer to **string** | Deadline is the drain&#39;s report bound — Overdue goes true at it and the force option opens. | [optional] 
**JobId** | Pointer to **string** | JobID is the job a pending drain waits on. | [optional] 
**Overdue** | Pointer to **bool** | Overdue is the drain-timeout REPORT (AC4): true once the deadline passed with the drain still pending. | [optional] 
**PendingAction** | Pointer to **string** | PendingAction is what the standing drain resolves into (\&quot;release\&quot; or \&quot;quarantine\&quot;); \&quot;\&quot; on a held or returning machine. | [optional] 
**RequestedAt** | Pointer to **string** | RequestedAt is when the standing transition was issued. | [optional] 

## Methods

### NewHandlersRunnerMachineDrainInfo

`func NewHandlersRunnerMachineDrainInfo() *HandlersRunnerMachineDrainInfo`

NewHandlersRunnerMachineDrainInfo instantiates a new HandlersRunnerMachineDrainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerMachineDrainInfoWithDefaults

`func NewHandlersRunnerMachineDrainInfoWithDefaults() *HandlersRunnerMachineDrainInfo`

NewHandlersRunnerMachineDrainInfoWithDefaults instantiates a new HandlersRunnerMachineDrainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeadline

`func (o *HandlersRunnerMachineDrainInfo) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *HandlersRunnerMachineDrainInfo) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *HandlersRunnerMachineDrainInfo) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *HandlersRunnerMachineDrainInfo) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetJobId

`func (o *HandlersRunnerMachineDrainInfo) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *HandlersRunnerMachineDrainInfo) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *HandlersRunnerMachineDrainInfo) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *HandlersRunnerMachineDrainInfo) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOverdue

`func (o *HandlersRunnerMachineDrainInfo) GetOverdue() bool`

GetOverdue returns the Overdue field if non-nil, zero value otherwise.

### GetOverdueOk

`func (o *HandlersRunnerMachineDrainInfo) GetOverdueOk() (*bool, bool)`

GetOverdueOk returns a tuple with the Overdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdue

`func (o *HandlersRunnerMachineDrainInfo) SetOverdue(v bool)`

SetOverdue sets Overdue field to given value.

### HasOverdue

`func (o *HandlersRunnerMachineDrainInfo) HasOverdue() bool`

HasOverdue returns a boolean if a field has been set.

### GetPendingAction

`func (o *HandlersRunnerMachineDrainInfo) GetPendingAction() string`

GetPendingAction returns the PendingAction field if non-nil, zero value otherwise.

### GetPendingActionOk

`func (o *HandlersRunnerMachineDrainInfo) GetPendingActionOk() (*string, bool)`

GetPendingActionOk returns a tuple with the PendingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAction

`func (o *HandlersRunnerMachineDrainInfo) SetPendingAction(v string)`

SetPendingAction sets PendingAction field to given value.

### HasPendingAction

`func (o *HandlersRunnerMachineDrainInfo) HasPendingAction() bool`

HasPendingAction returns a boolean if a field has been set.

### GetRequestedAt

`func (o *HandlersRunnerMachineDrainInfo) GetRequestedAt() string`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *HandlersRunnerMachineDrainInfo) GetRequestedAtOk() (*string, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *HandlersRunnerMachineDrainInfo) SetRequestedAt(v string)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *HandlersRunnerMachineDrainInfo) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


