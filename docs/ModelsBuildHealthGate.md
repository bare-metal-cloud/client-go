# ModelsBuildHealthGate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorUserId** | Pointer to **int32** | ActorUserID is the pipeline actor&#39;s user (the click&#39;s user on a rollback; the deploy trigger&#39;s user on a deploy leg). 0 when the actor is the pipeline itself. | [optional] 
**AttemptRef** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **string** | The deploy identity: the InstallRef triple&#39;s cluster + stack plus the service inside it — the service whose traffic this gate owns for the run&#39;s duration (the one-operation-at-a-time scope). | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DurationMs** | Pointer to **int32** |  | [optional] 
**Evidence** | Pointer to **[]int32** | Evidence is the probe&#39;s output, VERBATIM ([]string as JSON) — the pass lines on proven, the failure lines on blocked, the leg&#39;s named reason on failed. Never a paraphrase, never an invented checkmark. | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** | Kind is deploy or rollback (the vocabulary above). | [optional] 
**OrgId** | Pointer to **int32** | Tenant scope (ADR-062): the run lives inside one org and one project. Every read binds both halves. | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**RevisionId** | Pointer to **int32** | The revision this gate ran for — the build_revisions row the leg deploys (the target on a rollback). RevisionID is the row id the timeline resolves its health fact through; Revision is the human-facing monotonic number; AttemptRef is the revision&#39;s attempt reference (provenance); ImageDigest is the pinned digest the probe&#39;s workload runs. | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** | StartedAt / FinishedAt bound the run; DurationMs is the measured wall-clock of the whole gate run (the rollback&#39;s ≤30 s budget is recorded here — measured, not assumed). | [optional] 
**State** | Pointer to **string** | State is the run&#39;s truthful state: running → exactly one of proven / blocked / failed (the vocabulary above). | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsBuildHealthGate

`func NewModelsBuildHealthGate() *ModelsBuildHealthGate`

NewModelsBuildHealthGate instantiates a new ModelsBuildHealthGate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsBuildHealthGateWithDefaults

`func NewModelsBuildHealthGateWithDefaults() *ModelsBuildHealthGate`

NewModelsBuildHealthGateWithDefaults instantiates a new ModelsBuildHealthGate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorUserId

`func (o *ModelsBuildHealthGate) GetActorUserId() int32`

GetActorUserId returns the ActorUserId field if non-nil, zero value otherwise.

### GetActorUserIdOk

`func (o *ModelsBuildHealthGate) GetActorUserIdOk() (*int32, bool)`

GetActorUserIdOk returns a tuple with the ActorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUserId

`func (o *ModelsBuildHealthGate) SetActorUserId(v int32)`

SetActorUserId sets ActorUserId field to given value.

### HasActorUserId

`func (o *ModelsBuildHealthGate) HasActorUserId() bool`

HasActorUserId returns a boolean if a field has been set.

### GetAttemptRef

`func (o *ModelsBuildHealthGate) GetAttemptRef() string`

GetAttemptRef returns the AttemptRef field if non-nil, zero value otherwise.

### GetAttemptRefOk

`func (o *ModelsBuildHealthGate) GetAttemptRefOk() (*string, bool)`

GetAttemptRefOk returns a tuple with the AttemptRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptRef

`func (o *ModelsBuildHealthGate) SetAttemptRef(v string)`

SetAttemptRef sets AttemptRef field to given value.

### HasAttemptRef

`func (o *ModelsBuildHealthGate) HasAttemptRef() bool`

HasAttemptRef returns a boolean if a field has been set.

### GetClusterId

`func (o *ModelsBuildHealthGate) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ModelsBuildHealthGate) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ModelsBuildHealthGate) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ModelsBuildHealthGate) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsBuildHealthGate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsBuildHealthGate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsBuildHealthGate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsBuildHealthGate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDurationMs

`func (o *ModelsBuildHealthGate) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ModelsBuildHealthGate) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ModelsBuildHealthGate) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ModelsBuildHealthGate) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetEvidence

`func (o *ModelsBuildHealthGate) GetEvidence() []int32`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *ModelsBuildHealthGate) GetEvidenceOk() (*[]int32, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *ModelsBuildHealthGate) SetEvidence(v []int32)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *ModelsBuildHealthGate) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ModelsBuildHealthGate) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ModelsBuildHealthGate) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ModelsBuildHealthGate) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ModelsBuildHealthGate) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *ModelsBuildHealthGate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsBuildHealthGate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsBuildHealthGate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsBuildHealthGate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageDigest

`func (o *ModelsBuildHealthGate) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ModelsBuildHealthGate) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ModelsBuildHealthGate) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ModelsBuildHealthGate) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetKind

`func (o *ModelsBuildHealthGate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ModelsBuildHealthGate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ModelsBuildHealthGate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ModelsBuildHealthGate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsBuildHealthGate) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsBuildHealthGate) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsBuildHealthGate) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsBuildHealthGate) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsBuildHealthGate) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsBuildHealthGate) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsBuildHealthGate) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsBuildHealthGate) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRevision

`func (o *ModelsBuildHealthGate) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ModelsBuildHealthGate) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ModelsBuildHealthGate) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ModelsBuildHealthGate) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisionId

`func (o *ModelsBuildHealthGate) GetRevisionId() int32`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ModelsBuildHealthGate) GetRevisionIdOk() (*int32, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ModelsBuildHealthGate) SetRevisionId(v int32)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *ModelsBuildHealthGate) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetService

`func (o *ModelsBuildHealthGate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ModelsBuildHealthGate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ModelsBuildHealthGate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ModelsBuildHealthGate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStack

`func (o *ModelsBuildHealthGate) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *ModelsBuildHealthGate) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *ModelsBuildHealthGate) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *ModelsBuildHealthGate) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetStartedAt

`func (o *ModelsBuildHealthGate) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ModelsBuildHealthGate) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ModelsBuildHealthGate) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ModelsBuildHealthGate) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetState

`func (o *ModelsBuildHealthGate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelsBuildHealthGate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelsBuildHealthGate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ModelsBuildHealthGate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsBuildHealthGate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsBuildHealthGate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsBuildHealthGate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsBuildHealthGate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


