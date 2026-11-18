# EdgegitFetchSignalPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch + CommitSHA name the ref the push carried (the mover&#39;s update-ref input — a bare packfile carries no refs). | [optional] 
**CommitSha** | Pointer to **string** |  | [optional] 
**FetchPath** | Pointer to **string** | FetchPath is the pack&#39;s fetch path relative to the edge base URL the sidecar&#39;s env carries (the Bearer redemption URL). | [optional] 
**FetchToken** | Pointer to **string** | FetchToken is the 12-char Base32 PLAINTEXT fetch credential (the registration-token mint&#39;s #1511 law: shown once, hash-only at rest). Single-use, 15-minute TTL, bound to (org, project, repo, push). | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**PackDigest** | Pointer to **string** | PackDigest is the stored pack&#39;s sha256 (hex) — the sidecar verifies what it fetched against it before landing the pack in Forgejo. | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**PushRef** | Pointer to **string** |  | [optional] 
**RepoRef** | Pointer to **string** |  | [optional] 
**SizeBytes** | Pointer to **int32** |  | [optional] 

## Methods

### NewEdgegitFetchSignalPayload

`func NewEdgegitFetchSignalPayload() *EdgegitFetchSignalPayload`

NewEdgegitFetchSignalPayload instantiates a new EdgegitFetchSignalPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgegitFetchSignalPayloadWithDefaults

`func NewEdgegitFetchSignalPayloadWithDefaults() *EdgegitFetchSignalPayload`

NewEdgegitFetchSignalPayloadWithDefaults instantiates a new EdgegitFetchSignalPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *EdgegitFetchSignalPayload) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *EdgegitFetchSignalPayload) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *EdgegitFetchSignalPayload) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *EdgegitFetchSignalPayload) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCommitSha

`func (o *EdgegitFetchSignalPayload) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *EdgegitFetchSignalPayload) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *EdgegitFetchSignalPayload) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *EdgegitFetchSignalPayload) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetFetchPath

`func (o *EdgegitFetchSignalPayload) GetFetchPath() string`

GetFetchPath returns the FetchPath field if non-nil, zero value otherwise.

### GetFetchPathOk

`func (o *EdgegitFetchSignalPayload) GetFetchPathOk() (*string, bool)`

GetFetchPathOk returns a tuple with the FetchPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchPath

`func (o *EdgegitFetchSignalPayload) SetFetchPath(v string)`

SetFetchPath sets FetchPath field to given value.

### HasFetchPath

`func (o *EdgegitFetchSignalPayload) HasFetchPath() bool`

HasFetchPath returns a boolean if a field has been set.

### GetFetchToken

`func (o *EdgegitFetchSignalPayload) GetFetchToken() string`

GetFetchToken returns the FetchToken field if non-nil, zero value otherwise.

### GetFetchTokenOk

`func (o *EdgegitFetchSignalPayload) GetFetchTokenOk() (*string, bool)`

GetFetchTokenOk returns a tuple with the FetchToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchToken

`func (o *EdgegitFetchSignalPayload) SetFetchToken(v string)`

SetFetchToken sets FetchToken field to given value.

### HasFetchToken

`func (o *EdgegitFetchSignalPayload) HasFetchToken() bool`

HasFetchToken returns a boolean if a field has been set.

### GetOrgId

`func (o *EdgegitFetchSignalPayload) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EdgegitFetchSignalPayload) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EdgegitFetchSignalPayload) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *EdgegitFetchSignalPayload) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPackDigest

`func (o *EdgegitFetchSignalPayload) GetPackDigest() string`

GetPackDigest returns the PackDigest field if non-nil, zero value otherwise.

### GetPackDigestOk

`func (o *EdgegitFetchSignalPayload) GetPackDigestOk() (*string, bool)`

GetPackDigestOk returns a tuple with the PackDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackDigest

`func (o *EdgegitFetchSignalPayload) SetPackDigest(v string)`

SetPackDigest sets PackDigest field to given value.

### HasPackDigest

`func (o *EdgegitFetchSignalPayload) HasPackDigest() bool`

HasPackDigest returns a boolean if a field has been set.

### GetProjectId

`func (o *EdgegitFetchSignalPayload) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EdgegitFetchSignalPayload) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EdgegitFetchSignalPayload) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *EdgegitFetchSignalPayload) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPushRef

`func (o *EdgegitFetchSignalPayload) GetPushRef() string`

GetPushRef returns the PushRef field if non-nil, zero value otherwise.

### GetPushRefOk

`func (o *EdgegitFetchSignalPayload) GetPushRefOk() (*string, bool)`

GetPushRefOk returns a tuple with the PushRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushRef

`func (o *EdgegitFetchSignalPayload) SetPushRef(v string)`

SetPushRef sets PushRef field to given value.

### HasPushRef

`func (o *EdgegitFetchSignalPayload) HasPushRef() bool`

HasPushRef returns a boolean if a field has been set.

### GetRepoRef

`func (o *EdgegitFetchSignalPayload) GetRepoRef() string`

GetRepoRef returns the RepoRef field if non-nil, zero value otherwise.

### GetRepoRefOk

`func (o *EdgegitFetchSignalPayload) GetRepoRefOk() (*string, bool)`

GetRepoRefOk returns a tuple with the RepoRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoRef

`func (o *EdgegitFetchSignalPayload) SetRepoRef(v string)`

SetRepoRef sets RepoRef field to given value.

### HasRepoRef

`func (o *EdgegitFetchSignalPayload) HasRepoRef() bool`

HasRepoRef returns a boolean if a field has been set.

### GetSizeBytes

`func (o *EdgegitFetchSignalPayload) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *EdgegitFetchSignalPayload) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *EdgegitFetchSignalPayload) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *EdgegitFetchSignalPayload) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


