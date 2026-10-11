# ModelsJobMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**CurrentStep** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**JobType** | Pointer to **string** | vm:install, cluster:install, database:deploy, … | [optional] 
**MaxRetries** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **[]int32** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **int32** | 1&#x3D;highest, 5&#x3D;lowest | [optional] 
**Progress** | Pointer to **int32** | Fine-grained progress roll-up (0–100) and current step label. Updated by the step writer in the same tx that writes job_steps. List queries read these columns directly — no N+1 on job_steps. | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ResourceId** | Pointer to **int32** |  | [optional] 
**ResourceType** | Pointer to **string** | Generic resource link (ADR-060 V6 — replaces cluster_id/vm_id). \&quot;cluster\&quot; | \&quot;vm\&quot; | \&quot;database\&quot; | \&quot;network\&quot; | \&quot;multipass\&quot; | … | [optional] 
**RetryCount** | Pointer to **int32** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ModelsJobMetadataStatus**](ModelsJobMetadataStatus.md) | Coarse status enum (pending → running → completed|failed|cancelled). Status transitions are driven by a generic mechanism (ADR-060 Option A), never hand-set by business-logic error handling. | [optional] 
**TimeoutSeconds** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**WorkerId** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** | Temporal domain key (e.g. \&quot;provision-cluster-42\&quot;). Set when the job triggers a Temporal workflow; nil for Asynq-only jobs. | [optional] 

## Methods

### NewModelsJobMetadata

`func NewModelsJobMetadata() *ModelsJobMetadata`

NewModelsJobMetadata instantiates a new ModelsJobMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsJobMetadataWithDefaults

`func NewModelsJobMetadataWithDefaults() *ModelsJobMetadata`

NewModelsJobMetadataWithDefaults instantiates a new ModelsJobMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *ModelsJobMetadata) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ModelsJobMetadata) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ModelsJobMetadata) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ModelsJobMetadata) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsJobMetadata) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsJobMetadata) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsJobMetadata) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsJobMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsJobMetadata) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsJobMetadata) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsJobMetadata) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsJobMetadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCurrentStep

`func (o *ModelsJobMetadata) GetCurrentStep() string`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *ModelsJobMetadata) GetCurrentStepOk() (*string, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *ModelsJobMetadata) SetCurrentStep(v string)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *ModelsJobMetadata) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ModelsJobMetadata) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ModelsJobMetadata) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ModelsJobMetadata) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ModelsJobMetadata) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetJobId

`func (o *ModelsJobMetadata) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ModelsJobMetadata) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ModelsJobMetadata) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ModelsJobMetadata) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobType

`func (o *ModelsJobMetadata) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ModelsJobMetadata) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ModelsJobMetadata) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ModelsJobMetadata) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetMaxRetries

`func (o *ModelsJobMetadata) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *ModelsJobMetadata) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *ModelsJobMetadata) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *ModelsJobMetadata) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsJobMetadata) GetMetadata() []int32`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsJobMetadata) GetMetadataOk() (*[]int32, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsJobMetadata) SetMetadata(v []int32)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsJobMetadata) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsJobMetadata) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsJobMetadata) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsJobMetadata) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsJobMetadata) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPriority

`func (o *ModelsJobMetadata) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModelsJobMetadata) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModelsJobMetadata) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ModelsJobMetadata) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProgress

`func (o *ModelsJobMetadata) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ModelsJobMetadata) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ModelsJobMetadata) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ModelsJobMetadata) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsJobMetadata) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsJobMetadata) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsJobMetadata) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsJobMetadata) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetResourceId

`func (o *ModelsJobMetadata) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ModelsJobMetadata) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ModelsJobMetadata) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ModelsJobMetadata) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *ModelsJobMetadata) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ModelsJobMetadata) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ModelsJobMetadata) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ModelsJobMetadata) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetRetryCount

`func (o *ModelsJobMetadata) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ModelsJobMetadata) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ModelsJobMetadata) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ModelsJobMetadata) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetStartedAt

`func (o *ModelsJobMetadata) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ModelsJobMetadata) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ModelsJobMetadata) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ModelsJobMetadata) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsJobMetadata) GetStatus() ModelsJobMetadataStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsJobMetadata) GetStatusOk() (*ModelsJobMetadataStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsJobMetadata) SetStatus(v ModelsJobMetadataStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsJobMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *ModelsJobMetadata) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ModelsJobMetadata) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ModelsJobMetadata) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *ModelsJobMetadata) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsJobMetadata) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsJobMetadata) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsJobMetadata) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsJobMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkerId

`func (o *ModelsJobMetadata) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ModelsJobMetadata) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ModelsJobMetadata) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *ModelsJobMetadata) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *ModelsJobMetadata) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ModelsJobMetadata) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ModelsJobMetadata) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ModelsJobMetadata) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


