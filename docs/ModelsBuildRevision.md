# ModelsBuildRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptRef** | Pointer to **string** | AttemptRef is THE idempotency key: \&quot;&lt;run_reference&gt;#&lt;attempt&gt;\&quot; (BuildRevisionAttemptRef). UNIQUE — two completions of the same attempt converge on this one row. | [optional] 
**BuilderImageDigest** | Pointer to **string** | BuilderImageDigest is the digest-pinned builder image the build ran on (the D4 pinning law — provenance survives registry transit). | [optional] 
**ClusterId** | Pointer to **string** | The deploykit InstallRef triple (D8) plus the service inside the stack: (cluster_id, stack, service) names the deploy unit the revision was built for. The (stack, service, revision) half is UNIQUE inside the tenancy scope — idx_build_revision_identity. | [optional] 
**Commit** | Pointer to **string** | Commit is the source commit the build ran against. | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CredentialRef** | Pointer to **string** | CredentialRef is the minted credential&#39;s reference (D8&#39;s payload): the attempt&#39;s KindBuild token LOOKUP HASH (the #1511 keyed-hash form, 64 hex chars — a reference, never the plaintext secret, and hash-at-rest is the storage law). Copied from the attempt ledger row at acceptance. | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ImageDigest** | Pointer to **string** | ImageDigest is the produced image&#39;s digest (\&quot;sha256:&lt;64 hex&gt;\&quot; — the publish shape discipline the revision registry adopts from gitops publish.go). RegistryRef is the registry repository reference the digest was verified in (host/repository). Together they render the pinned reference: RegistryRef + \&quot;@\&quot; + ImageDigest. | [optional] 
**OrgId** | Pointer to **int32** | Tenant scope (ADR-062): the revision lives inside one org and one project. Every read binds org+project — a cross-org read is a miss, never a wildcard. Both halves ride the identity unique index (mirroring migration 000122&#39;s six-column form). | [optional] 
**Outcome** | Pointer to **string** | Outcome is the verified completion&#39;s outcome. A revision row exists only for a verified publish (publish-before-pin, D7), so today the only value is \&quot;succeeded\&quot;; the column is the D8 payload&#39;s outcome field, carried for the schema&#39;s completeness. | [optional] 
**Plan** | Pointer to **[]int32** | Plan is the resolved build block (the Session-32 grammar as compiled) — the JSON the plan was built from, stored verbatim so a pruned image can be re-built from the plan + commit (design §12 question 2&#39;s answer). | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**RegistryRef** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **int32** | Revision is the monotonic per-identity number (1, 2, 3, ... — assigned by the repository layer, never a caller&#39;s claim). | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**TriggerSource** | Pointer to **string** | TriggerSource is what asked for the deploy (the trigger_source grammar above): written once at registration, never updated — the append-only law covers it like every other column. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsBuildRevision

`func NewModelsBuildRevision() *ModelsBuildRevision`

NewModelsBuildRevision instantiates a new ModelsBuildRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsBuildRevisionWithDefaults

`func NewModelsBuildRevisionWithDefaults() *ModelsBuildRevision`

NewModelsBuildRevisionWithDefaults instantiates a new ModelsBuildRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptRef

`func (o *ModelsBuildRevision) GetAttemptRef() string`

GetAttemptRef returns the AttemptRef field if non-nil, zero value otherwise.

### GetAttemptRefOk

`func (o *ModelsBuildRevision) GetAttemptRefOk() (*string, bool)`

GetAttemptRefOk returns a tuple with the AttemptRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptRef

`func (o *ModelsBuildRevision) SetAttemptRef(v string)`

SetAttemptRef sets AttemptRef field to given value.

### HasAttemptRef

`func (o *ModelsBuildRevision) HasAttemptRef() bool`

HasAttemptRef returns a boolean if a field has been set.

### GetBuilderImageDigest

`func (o *ModelsBuildRevision) GetBuilderImageDigest() string`

GetBuilderImageDigest returns the BuilderImageDigest field if non-nil, zero value otherwise.

### GetBuilderImageDigestOk

`func (o *ModelsBuildRevision) GetBuilderImageDigestOk() (*string, bool)`

GetBuilderImageDigestOk returns a tuple with the BuilderImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderImageDigest

`func (o *ModelsBuildRevision) SetBuilderImageDigest(v string)`

SetBuilderImageDigest sets BuilderImageDigest field to given value.

### HasBuilderImageDigest

`func (o *ModelsBuildRevision) HasBuilderImageDigest() bool`

HasBuilderImageDigest returns a boolean if a field has been set.

### GetClusterId

`func (o *ModelsBuildRevision) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ModelsBuildRevision) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ModelsBuildRevision) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ModelsBuildRevision) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCommit

`func (o *ModelsBuildRevision) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ModelsBuildRevision) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ModelsBuildRevision) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ModelsBuildRevision) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsBuildRevision) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsBuildRevision) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsBuildRevision) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsBuildRevision) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialRef

`func (o *ModelsBuildRevision) GetCredentialRef() string`

GetCredentialRef returns the CredentialRef field if non-nil, zero value otherwise.

### GetCredentialRefOk

`func (o *ModelsBuildRevision) GetCredentialRefOk() (*string, bool)`

GetCredentialRefOk returns a tuple with the CredentialRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialRef

`func (o *ModelsBuildRevision) SetCredentialRef(v string)`

SetCredentialRef sets CredentialRef field to given value.

### HasCredentialRef

`func (o *ModelsBuildRevision) HasCredentialRef() bool`

HasCredentialRef returns a boolean if a field has been set.

### GetId

`func (o *ModelsBuildRevision) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsBuildRevision) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsBuildRevision) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsBuildRevision) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageDigest

`func (o *ModelsBuildRevision) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *ModelsBuildRevision) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *ModelsBuildRevision) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *ModelsBuildRevision) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsBuildRevision) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsBuildRevision) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsBuildRevision) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsBuildRevision) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOutcome

`func (o *ModelsBuildRevision) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ModelsBuildRevision) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ModelsBuildRevision) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *ModelsBuildRevision) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetPlan

`func (o *ModelsBuildRevision) GetPlan() []int32`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ModelsBuildRevision) GetPlanOk() (*[]int32, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ModelsBuildRevision) SetPlan(v []int32)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ModelsBuildRevision) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsBuildRevision) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsBuildRevision) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsBuildRevision) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsBuildRevision) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRegistryRef

`func (o *ModelsBuildRevision) GetRegistryRef() string`

GetRegistryRef returns the RegistryRef field if non-nil, zero value otherwise.

### GetRegistryRefOk

`func (o *ModelsBuildRevision) GetRegistryRefOk() (*string, bool)`

GetRegistryRefOk returns a tuple with the RegistryRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryRef

`func (o *ModelsBuildRevision) SetRegistryRef(v string)`

SetRegistryRef sets RegistryRef field to given value.

### HasRegistryRef

`func (o *ModelsBuildRevision) HasRegistryRef() bool`

HasRegistryRef returns a boolean if a field has been set.

### GetRevision

`func (o *ModelsBuildRevision) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ModelsBuildRevision) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ModelsBuildRevision) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ModelsBuildRevision) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetService

`func (o *ModelsBuildRevision) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ModelsBuildRevision) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ModelsBuildRevision) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ModelsBuildRevision) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStack

`func (o *ModelsBuildRevision) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *ModelsBuildRevision) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *ModelsBuildRevision) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *ModelsBuildRevision) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetTriggerSource

`func (o *ModelsBuildRevision) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *ModelsBuildRevision) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *ModelsBuildRevision) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.

### HasTriggerSource

`func (o *ModelsBuildRevision) HasTriggerSource() bool`

HasTriggerSource returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsBuildRevision) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsBuildRevision) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsBuildRevision) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsBuildRevision) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


