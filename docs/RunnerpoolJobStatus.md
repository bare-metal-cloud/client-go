# RunnerpoolJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdmittedAt** | Pointer to **string** |  | [optional] 
**DispatchedAt** | Pointer to **string** |  | [optional] 
**DurationSeconds** | Pointer to **float32** |  | [optional] 
**ExitCode** | Pointer to **int32** |  | [optional] 
**Hosted** | Pointer to **bool** | Hosted is the hosted badge: every job this route reports is a hosted-chain job. | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProviderVmId** | Pointer to **string** | ProviderVMID is the floor VM&#39;s provider id while the job binds one (the Epic 145.2 swap-scope note on the CP-API handler type covers its exposure). | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ReasonClass** | Pointer to **string** |  | [optional] 
**RunReference** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to [**RunnerpoolStage**](RunnerpoolStage.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TerminalAt** | Pointer to **string** |  | [optional] 
**TerminalKind** | Pointer to **string** |  | [optional] 
**VmSeconds** | Pointer to **float32** | VMSeconds is the control-plane-computed metered amount; DurationSeconds is the admitted-to-now-or-terminal wall clock. | [optional] 

## Methods

### NewRunnerpoolJobStatus

`func NewRunnerpoolJobStatus() *RunnerpoolJobStatus`

NewRunnerpoolJobStatus instantiates a new RunnerpoolJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerpoolJobStatusWithDefaults

`func NewRunnerpoolJobStatusWithDefaults() *RunnerpoolJobStatus`

NewRunnerpoolJobStatusWithDefaults instantiates a new RunnerpoolJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmittedAt

`func (o *RunnerpoolJobStatus) GetAdmittedAt() string`

GetAdmittedAt returns the AdmittedAt field if non-nil, zero value otherwise.

### GetAdmittedAtOk

`func (o *RunnerpoolJobStatus) GetAdmittedAtOk() (*string, bool)`

GetAdmittedAtOk returns a tuple with the AdmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmittedAt

`func (o *RunnerpoolJobStatus) SetAdmittedAt(v string)`

SetAdmittedAt sets AdmittedAt field to given value.

### HasAdmittedAt

`func (o *RunnerpoolJobStatus) HasAdmittedAt() bool`

HasAdmittedAt returns a boolean if a field has been set.

### GetDispatchedAt

`func (o *RunnerpoolJobStatus) GetDispatchedAt() string`

GetDispatchedAt returns the DispatchedAt field if non-nil, zero value otherwise.

### GetDispatchedAtOk

`func (o *RunnerpoolJobStatus) GetDispatchedAtOk() (*string, bool)`

GetDispatchedAtOk returns a tuple with the DispatchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchedAt

`func (o *RunnerpoolJobStatus) SetDispatchedAt(v string)`

SetDispatchedAt sets DispatchedAt field to given value.

### HasDispatchedAt

`func (o *RunnerpoolJobStatus) HasDispatchedAt() bool`

HasDispatchedAt returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *RunnerpoolJobStatus) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *RunnerpoolJobStatus) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *RunnerpoolJobStatus) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *RunnerpoolJobStatus) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetExitCode

`func (o *RunnerpoolJobStatus) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *RunnerpoolJobStatus) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *RunnerpoolJobStatus) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *RunnerpoolJobStatus) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetHosted

`func (o *RunnerpoolJobStatus) GetHosted() bool`

GetHosted returns the Hosted field if non-nil, zero value otherwise.

### GetHostedOk

`func (o *RunnerpoolJobStatus) GetHostedOk() (*bool, bool)`

GetHostedOk returns a tuple with the Hosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosted

`func (o *RunnerpoolJobStatus) SetHosted(v bool)`

SetHosted sets Hosted field to given value.

### HasHosted

`func (o *RunnerpoolJobStatus) HasHosted() bool`

HasHosted returns a boolean if a field has been set.

### GetJobId

`func (o *RunnerpoolJobStatus) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *RunnerpoolJobStatus) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *RunnerpoolJobStatus) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *RunnerpoolJobStatus) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetOrgId

`func (o *RunnerpoolJobStatus) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RunnerpoolJobStatus) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RunnerpoolJobStatus) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *RunnerpoolJobStatus) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *RunnerpoolJobStatus) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RunnerpoolJobStatus) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RunnerpoolJobStatus) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RunnerpoolJobStatus) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProviderVmId

`func (o *RunnerpoolJobStatus) GetProviderVmId() string`

GetProviderVmId returns the ProviderVmId field if non-nil, zero value otherwise.

### GetProviderVmIdOk

`func (o *RunnerpoolJobStatus) GetProviderVmIdOk() (*string, bool)`

GetProviderVmIdOk returns a tuple with the ProviderVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVmId

`func (o *RunnerpoolJobStatus) SetProviderVmId(v string)`

SetProviderVmId sets ProviderVmId field to given value.

### HasProviderVmId

`func (o *RunnerpoolJobStatus) HasProviderVmId() bool`

HasProviderVmId returns a boolean if a field has been set.

### GetReason

`func (o *RunnerpoolJobStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RunnerpoolJobStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RunnerpoolJobStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RunnerpoolJobStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonClass

`func (o *RunnerpoolJobStatus) GetReasonClass() string`

GetReasonClass returns the ReasonClass field if non-nil, zero value otherwise.

### GetReasonClassOk

`func (o *RunnerpoolJobStatus) GetReasonClassOk() (*string, bool)`

GetReasonClassOk returns a tuple with the ReasonClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonClass

`func (o *RunnerpoolJobStatus) SetReasonClass(v string)`

SetReasonClass sets ReasonClass field to given value.

### HasReasonClass

`func (o *RunnerpoolJobStatus) HasReasonClass() bool`

HasReasonClass returns a boolean if a field has been set.

### GetRunReference

`func (o *RunnerpoolJobStatus) GetRunReference() string`

GetRunReference returns the RunReference field if non-nil, zero value otherwise.

### GetRunReferenceOk

`func (o *RunnerpoolJobStatus) GetRunReferenceOk() (*string, bool)`

GetRunReferenceOk returns a tuple with the RunReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunReference

`func (o *RunnerpoolJobStatus) SetRunReference(v string)`

SetRunReference sets RunReference field to given value.

### HasRunReference

`func (o *RunnerpoolJobStatus) HasRunReference() bool`

HasRunReference returns a boolean if a field has been set.

### GetStage

`func (o *RunnerpoolJobStatus) GetStage() RunnerpoolStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *RunnerpoolJobStatus) GetStageOk() (*RunnerpoolStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *RunnerpoolJobStatus) SetStage(v RunnerpoolStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *RunnerpoolJobStatus) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetState

`func (o *RunnerpoolJobStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RunnerpoolJobStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RunnerpoolJobStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RunnerpoolJobStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTerminalAt

`func (o *RunnerpoolJobStatus) GetTerminalAt() string`

GetTerminalAt returns the TerminalAt field if non-nil, zero value otherwise.

### GetTerminalAtOk

`func (o *RunnerpoolJobStatus) GetTerminalAtOk() (*string, bool)`

GetTerminalAtOk returns a tuple with the TerminalAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalAt

`func (o *RunnerpoolJobStatus) SetTerminalAt(v string)`

SetTerminalAt sets TerminalAt field to given value.

### HasTerminalAt

`func (o *RunnerpoolJobStatus) HasTerminalAt() bool`

HasTerminalAt returns a boolean if a field has been set.

### GetTerminalKind

`func (o *RunnerpoolJobStatus) GetTerminalKind() string`

GetTerminalKind returns the TerminalKind field if non-nil, zero value otherwise.

### GetTerminalKindOk

`func (o *RunnerpoolJobStatus) GetTerminalKindOk() (*string, bool)`

GetTerminalKindOk returns a tuple with the TerminalKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalKind

`func (o *RunnerpoolJobStatus) SetTerminalKind(v string)`

SetTerminalKind sets TerminalKind field to given value.

### HasTerminalKind

`func (o *RunnerpoolJobStatus) HasTerminalKind() bool`

HasTerminalKind returns a boolean if a field has been set.

### GetVmSeconds

`func (o *RunnerpoolJobStatus) GetVmSeconds() float32`

GetVmSeconds returns the VmSeconds field if non-nil, zero value otherwise.

### GetVmSecondsOk

`func (o *RunnerpoolJobStatus) GetVmSecondsOk() (*float32, bool)`

GetVmSecondsOk returns a tuple with the VmSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSeconds

`func (o *RunnerpoolJobStatus) SetVmSeconds(v float32)`

SetVmSeconds sets VmSeconds field to given value.

### HasVmSeconds

`func (o *RunnerpoolJobStatus) HasVmSeconds() bool`

HasVmSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


