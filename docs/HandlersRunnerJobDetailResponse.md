# HandlersRunnerJobDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **string** |  | [optional] 
**Conclusion** | Pointer to **string** | Conclusion is GitHub&#39;s workflow_job.conclusion (github rows only; empty until a completed delivery lands). | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is the execution&#39;s arrival. | [optional] 
**DurationSeconds** | Pointer to **int32** |  | [optional] 
**GithubLogUrl** | Pointer to **string** |  | [optional] 
**Hosted** | Pointer to **bool** |  | [optional] 
**JobRef** | Pointer to **string** | JobRef is the execution&#39;s reference; (source, jobRef) resolves the detail. | [optional] 
**LogBlobDigest** | Pointer to **string** | LogBlobDigest is the sha256 of the captured gzip log in the org&#39;s bucket — the log read&#39;s handle. Empty &#x3D; never captured. | [optional] 
**LogSize** | Pointer to **int32** |  | [optional] 
**LogUploadedAt** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**QueueWaitSeconds** | Pointer to **int32** | QueueWaitSeconds / DurationSeconds are the view&#39;s NULL-armed computations; null &#x3D; unknown, never zero (D2). | [optional] 
**Repo** | Pointer to **string** | Repo / Workflow / Profile / Runner / State / Result / GithubLogURL ride the ledger row verbatim (the list&#39;s contract). | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**RunId** | Pointer to **int32** | RunID is GitHub&#39;s workflow run id (0 for the BMC-run sources); LogSize the capture&#39;s COMPRESSED byte size (0 until captured); Hosted true only for github rows. | [optional] 
**Runner** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** | SourceID + Source are the row&#39;s identity ((source, sourceId)); Source is \&quot;pool\&quot; | \&quot;builder\&quot; | \&quot;github\&quot;. | [optional] 
**StartedAt** | Pointer to **string** | StartedAt / CompletedAt / LogUploadedAt: the first two are GITHUB&#39;S own stamps (the mirror&#39;s source of truth; null when GitHub never sent one), the third the capture&#39;s commit moment (null until the first capture). | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Workflow** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerJobDetailResponse

`func NewHandlersRunnerJobDetailResponse() *HandlersRunnerJobDetailResponse`

NewHandlersRunnerJobDetailResponse instantiates a new HandlersRunnerJobDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerJobDetailResponseWithDefaults

`func NewHandlersRunnerJobDetailResponseWithDefaults() *HandlersRunnerJobDetailResponse`

NewHandlersRunnerJobDetailResponseWithDefaults instantiates a new HandlersRunnerJobDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *HandlersRunnerJobDetailResponse) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *HandlersRunnerJobDetailResponse) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *HandlersRunnerJobDetailResponse) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *HandlersRunnerJobDetailResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetConclusion

`func (o *HandlersRunnerJobDetailResponse) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *HandlersRunnerJobDetailResponse) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *HandlersRunnerJobDetailResponse) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *HandlersRunnerJobDetailResponse) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersRunnerJobDetailResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersRunnerJobDetailResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersRunnerJobDetailResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersRunnerJobDetailResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *HandlersRunnerJobDetailResponse) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *HandlersRunnerJobDetailResponse) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *HandlersRunnerJobDetailResponse) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *HandlersRunnerJobDetailResponse) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetGithubLogUrl

`func (o *HandlersRunnerJobDetailResponse) GetGithubLogUrl() string`

GetGithubLogUrl returns the GithubLogUrl field if non-nil, zero value otherwise.

### GetGithubLogUrlOk

`func (o *HandlersRunnerJobDetailResponse) GetGithubLogUrlOk() (*string, bool)`

GetGithubLogUrlOk returns a tuple with the GithubLogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLogUrl

`func (o *HandlersRunnerJobDetailResponse) SetGithubLogUrl(v string)`

SetGithubLogUrl sets GithubLogUrl field to given value.

### HasGithubLogUrl

`func (o *HandlersRunnerJobDetailResponse) HasGithubLogUrl() bool`

HasGithubLogUrl returns a boolean if a field has been set.

### GetHosted

`func (o *HandlersRunnerJobDetailResponse) GetHosted() bool`

GetHosted returns the Hosted field if non-nil, zero value otherwise.

### GetHostedOk

`func (o *HandlersRunnerJobDetailResponse) GetHostedOk() (*bool, bool)`

GetHostedOk returns a tuple with the Hosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosted

`func (o *HandlersRunnerJobDetailResponse) SetHosted(v bool)`

SetHosted sets Hosted field to given value.

### HasHosted

`func (o *HandlersRunnerJobDetailResponse) HasHosted() bool`

HasHosted returns a boolean if a field has been set.

### GetJobRef

`func (o *HandlersRunnerJobDetailResponse) GetJobRef() string`

GetJobRef returns the JobRef field if non-nil, zero value otherwise.

### GetJobRefOk

`func (o *HandlersRunnerJobDetailResponse) GetJobRefOk() (*string, bool)`

GetJobRefOk returns a tuple with the JobRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRef

`func (o *HandlersRunnerJobDetailResponse) SetJobRef(v string)`

SetJobRef sets JobRef field to given value.

### HasJobRef

`func (o *HandlersRunnerJobDetailResponse) HasJobRef() bool`

HasJobRef returns a boolean if a field has been set.

### GetLogBlobDigest

`func (o *HandlersRunnerJobDetailResponse) GetLogBlobDigest() string`

GetLogBlobDigest returns the LogBlobDigest field if non-nil, zero value otherwise.

### GetLogBlobDigestOk

`func (o *HandlersRunnerJobDetailResponse) GetLogBlobDigestOk() (*string, bool)`

GetLogBlobDigestOk returns a tuple with the LogBlobDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBlobDigest

`func (o *HandlersRunnerJobDetailResponse) SetLogBlobDigest(v string)`

SetLogBlobDigest sets LogBlobDigest field to given value.

### HasLogBlobDigest

`func (o *HandlersRunnerJobDetailResponse) HasLogBlobDigest() bool`

HasLogBlobDigest returns a boolean if a field has been set.

### GetLogSize

`func (o *HandlersRunnerJobDetailResponse) GetLogSize() int32`

GetLogSize returns the LogSize field if non-nil, zero value otherwise.

### GetLogSizeOk

`func (o *HandlersRunnerJobDetailResponse) GetLogSizeOk() (*int32, bool)`

GetLogSizeOk returns a tuple with the LogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSize

`func (o *HandlersRunnerJobDetailResponse) SetLogSize(v int32)`

SetLogSize sets LogSize field to given value.

### HasLogSize

`func (o *HandlersRunnerJobDetailResponse) HasLogSize() bool`

HasLogSize returns a boolean if a field has been set.

### GetLogUploadedAt

`func (o *HandlersRunnerJobDetailResponse) GetLogUploadedAt() string`

GetLogUploadedAt returns the LogUploadedAt field if non-nil, zero value otherwise.

### GetLogUploadedAtOk

`func (o *HandlersRunnerJobDetailResponse) GetLogUploadedAtOk() (*string, bool)`

GetLogUploadedAtOk returns a tuple with the LogUploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUploadedAt

`func (o *HandlersRunnerJobDetailResponse) SetLogUploadedAt(v string)`

SetLogUploadedAt sets LogUploadedAt field to given value.

### HasLogUploadedAt

`func (o *HandlersRunnerJobDetailResponse) HasLogUploadedAt() bool`

HasLogUploadedAt returns a boolean if a field has been set.

### GetProfile

`func (o *HandlersRunnerJobDetailResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *HandlersRunnerJobDetailResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *HandlersRunnerJobDetailResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *HandlersRunnerJobDetailResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQueueWaitSeconds

`func (o *HandlersRunnerJobDetailResponse) GetQueueWaitSeconds() int32`

GetQueueWaitSeconds returns the QueueWaitSeconds field if non-nil, zero value otherwise.

### GetQueueWaitSecondsOk

`func (o *HandlersRunnerJobDetailResponse) GetQueueWaitSecondsOk() (*int32, bool)`

GetQueueWaitSecondsOk returns a tuple with the QueueWaitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueWaitSeconds

`func (o *HandlersRunnerJobDetailResponse) SetQueueWaitSeconds(v int32)`

SetQueueWaitSeconds sets QueueWaitSeconds field to given value.

### HasQueueWaitSeconds

`func (o *HandlersRunnerJobDetailResponse) HasQueueWaitSeconds() bool`

HasQueueWaitSeconds returns a boolean if a field has been set.

### GetRepo

`func (o *HandlersRunnerJobDetailResponse) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *HandlersRunnerJobDetailResponse) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *HandlersRunnerJobDetailResponse) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *HandlersRunnerJobDetailResponse) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetResult

`func (o *HandlersRunnerJobDetailResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *HandlersRunnerJobDetailResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *HandlersRunnerJobDetailResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *HandlersRunnerJobDetailResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRunId

`func (o *HandlersRunnerJobDetailResponse) GetRunId() int32`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *HandlersRunnerJobDetailResponse) GetRunIdOk() (*int32, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *HandlersRunnerJobDetailResponse) SetRunId(v int32)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *HandlersRunnerJobDetailResponse) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetRunner

`func (o *HandlersRunnerJobDetailResponse) GetRunner() string`

GetRunner returns the Runner field if non-nil, zero value otherwise.

### GetRunnerOk

`func (o *HandlersRunnerJobDetailResponse) GetRunnerOk() (*string, bool)`

GetRunnerOk returns a tuple with the Runner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunner

`func (o *HandlersRunnerJobDetailResponse) SetRunner(v string)`

SetRunner sets Runner field to given value.

### HasRunner

`func (o *HandlersRunnerJobDetailResponse) HasRunner() bool`

HasRunner returns a boolean if a field has been set.

### GetSource

`func (o *HandlersRunnerJobDetailResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HandlersRunnerJobDetailResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HandlersRunnerJobDetailResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *HandlersRunnerJobDetailResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceId

`func (o *HandlersRunnerJobDetailResponse) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HandlersRunnerJobDetailResponse) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HandlersRunnerJobDetailResponse) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *HandlersRunnerJobDetailResponse) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetStartedAt

`func (o *HandlersRunnerJobDetailResponse) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *HandlersRunnerJobDetailResponse) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *HandlersRunnerJobDetailResponse) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *HandlersRunnerJobDetailResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetState

`func (o *HandlersRunnerJobDetailResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersRunnerJobDetailResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersRunnerJobDetailResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersRunnerJobDetailResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkflow

`func (o *HandlersRunnerJobDetailResponse) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *HandlersRunnerJobDetailResponse) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *HandlersRunnerJobDetailResponse) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *HandlersRunnerJobDetailResponse) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


