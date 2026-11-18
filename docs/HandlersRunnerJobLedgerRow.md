# HandlersRunnerJobLedgerRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is the execution&#39;s arrival. | [optional] 
**DurationSeconds** | Pointer to **int32** | DurationSeconds is the server-computed run duration; null when live or the source records no terminal timestamp. | [optional] 
**GithubLogUrl** | Pointer to **string** | GithubLogURL is the GitHub job&#39;s log page (P1.11&#39;s link-out); empty for the other sources. | [optional] 
**Hosted** | Pointer to **bool** | Hosted is true only for github rows: a GitHub-hosted job BMC never ran. | [optional] 
**JobRef** | Pointer to **string** | JobRef is the execution&#39;s reference: the pool job id, the builder&#39;s RunReference#attempt, GitHub&#39;s job id. | [optional] 
**Profile** | Pointer to **string** | Profile stays empty in v1 (the placement note lives on the refusals, B2.5). | [optional] 
**QueueWaitSeconds** | Pointer to **int32** | QueueWaitSeconds is the server-computed queue wait; null &#x3D; unknown, never zero (D2). | [optional] 
**Repo** | Pointer to **string** | Repo is the repository (the github source&#39;s repository name; empty for the BMC-run sources in v1). | [optional] 
**Result** | Pointer to **string** | Result is the terminal reason / conclusion; empty while live. | [optional] 
**RunId** | Pointer to **int32** | RunID is GitHub&#39;s workflow run id (0 for the BMC-run sources). | [optional] 
**Runner** | Pointer to **string** | Runner is the machine the execution ran on (the pool VM&#39;s provider id, the builder&#39;s machine agent id; empty for github rows — GitHub-hosted, BMC never ran them). | [optional] 
**Source** | Pointer to **string** | Source is \&quot;pool\&quot; | \&quot;builder\&quot; | \&quot;github\&quot;. | [optional] 
**SourceId** | Pointer to **string** | SourceID is the source row&#39;s id as text; (source, sourceId) is the row&#39;s identity. | [optional] 
**State** | Pointer to **string** | State is the source&#39;s state verbatim. | [optional] 
**Workflow** | Pointer to **string** | Workflow stays empty in v1 — GitHub&#39;s workflow name has no delivering story yet (the events feed carries actions only); surfaced when a story owns it. | [optional] 

## Methods

### NewHandlersRunnerJobLedgerRow

`func NewHandlersRunnerJobLedgerRow() *HandlersRunnerJobLedgerRow`

NewHandlersRunnerJobLedgerRow instantiates a new HandlersRunnerJobLedgerRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerJobLedgerRowWithDefaults

`func NewHandlersRunnerJobLedgerRowWithDefaults() *HandlersRunnerJobLedgerRow`

NewHandlersRunnerJobLedgerRowWithDefaults instantiates a new HandlersRunnerJobLedgerRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HandlersRunnerJobLedgerRow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersRunnerJobLedgerRow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersRunnerJobLedgerRow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersRunnerJobLedgerRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *HandlersRunnerJobLedgerRow) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *HandlersRunnerJobLedgerRow) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *HandlersRunnerJobLedgerRow) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *HandlersRunnerJobLedgerRow) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetGithubLogUrl

`func (o *HandlersRunnerJobLedgerRow) GetGithubLogUrl() string`

GetGithubLogUrl returns the GithubLogUrl field if non-nil, zero value otherwise.

### GetGithubLogUrlOk

`func (o *HandlersRunnerJobLedgerRow) GetGithubLogUrlOk() (*string, bool)`

GetGithubLogUrlOk returns a tuple with the GithubLogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLogUrl

`func (o *HandlersRunnerJobLedgerRow) SetGithubLogUrl(v string)`

SetGithubLogUrl sets GithubLogUrl field to given value.

### HasGithubLogUrl

`func (o *HandlersRunnerJobLedgerRow) HasGithubLogUrl() bool`

HasGithubLogUrl returns a boolean if a field has been set.

### GetHosted

`func (o *HandlersRunnerJobLedgerRow) GetHosted() bool`

GetHosted returns the Hosted field if non-nil, zero value otherwise.

### GetHostedOk

`func (o *HandlersRunnerJobLedgerRow) GetHostedOk() (*bool, bool)`

GetHostedOk returns a tuple with the Hosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosted

`func (o *HandlersRunnerJobLedgerRow) SetHosted(v bool)`

SetHosted sets Hosted field to given value.

### HasHosted

`func (o *HandlersRunnerJobLedgerRow) HasHosted() bool`

HasHosted returns a boolean if a field has been set.

### GetJobRef

`func (o *HandlersRunnerJobLedgerRow) GetJobRef() string`

GetJobRef returns the JobRef field if non-nil, zero value otherwise.

### GetJobRefOk

`func (o *HandlersRunnerJobLedgerRow) GetJobRefOk() (*string, bool)`

GetJobRefOk returns a tuple with the JobRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRef

`func (o *HandlersRunnerJobLedgerRow) SetJobRef(v string)`

SetJobRef sets JobRef field to given value.

### HasJobRef

`func (o *HandlersRunnerJobLedgerRow) HasJobRef() bool`

HasJobRef returns a boolean if a field has been set.

### GetProfile

`func (o *HandlersRunnerJobLedgerRow) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *HandlersRunnerJobLedgerRow) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *HandlersRunnerJobLedgerRow) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *HandlersRunnerJobLedgerRow) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetQueueWaitSeconds

`func (o *HandlersRunnerJobLedgerRow) GetQueueWaitSeconds() int32`

GetQueueWaitSeconds returns the QueueWaitSeconds field if non-nil, zero value otherwise.

### GetQueueWaitSecondsOk

`func (o *HandlersRunnerJobLedgerRow) GetQueueWaitSecondsOk() (*int32, bool)`

GetQueueWaitSecondsOk returns a tuple with the QueueWaitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueWaitSeconds

`func (o *HandlersRunnerJobLedgerRow) SetQueueWaitSeconds(v int32)`

SetQueueWaitSeconds sets QueueWaitSeconds field to given value.

### HasQueueWaitSeconds

`func (o *HandlersRunnerJobLedgerRow) HasQueueWaitSeconds() bool`

HasQueueWaitSeconds returns a boolean if a field has been set.

### GetRepo

`func (o *HandlersRunnerJobLedgerRow) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *HandlersRunnerJobLedgerRow) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *HandlersRunnerJobLedgerRow) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *HandlersRunnerJobLedgerRow) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetResult

`func (o *HandlersRunnerJobLedgerRow) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *HandlersRunnerJobLedgerRow) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *HandlersRunnerJobLedgerRow) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *HandlersRunnerJobLedgerRow) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRunId

`func (o *HandlersRunnerJobLedgerRow) GetRunId() int32`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *HandlersRunnerJobLedgerRow) GetRunIdOk() (*int32, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *HandlersRunnerJobLedgerRow) SetRunId(v int32)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *HandlersRunnerJobLedgerRow) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetRunner

`func (o *HandlersRunnerJobLedgerRow) GetRunner() string`

GetRunner returns the Runner field if non-nil, zero value otherwise.

### GetRunnerOk

`func (o *HandlersRunnerJobLedgerRow) GetRunnerOk() (*string, bool)`

GetRunnerOk returns a tuple with the Runner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunner

`func (o *HandlersRunnerJobLedgerRow) SetRunner(v string)`

SetRunner sets Runner field to given value.

### HasRunner

`func (o *HandlersRunnerJobLedgerRow) HasRunner() bool`

HasRunner returns a boolean if a field has been set.

### GetSource

`func (o *HandlersRunnerJobLedgerRow) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HandlersRunnerJobLedgerRow) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HandlersRunnerJobLedgerRow) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *HandlersRunnerJobLedgerRow) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceId

`func (o *HandlersRunnerJobLedgerRow) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HandlersRunnerJobLedgerRow) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HandlersRunnerJobLedgerRow) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *HandlersRunnerJobLedgerRow) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetState

`func (o *HandlersRunnerJobLedgerRow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersRunnerJobLedgerRow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersRunnerJobLedgerRow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersRunnerJobLedgerRow) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkflow

`func (o *HandlersRunnerJobLedgerRow) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *HandlersRunnerJobLedgerRow) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *HandlersRunnerJobLedgerRow) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *HandlersRunnerJobLedgerRow) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


