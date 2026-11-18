# HandlersGitBuildTriggerBindRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** |  | [optional] 
**Build** | Pointer to **string** |  | [optional] 
**BuilderImage** | Pointer to **string** |  | [optional] 
**EgressAllowlist** | Pointer to **[]string** | EgressAllowlist is OPTIONAL (the 2026-10-08 drill composition): the hostname-suffix list the dispatched job&#39;s egress guard allows. Absent/empty composes the deny-all default (fail closed); the entry grammar is the contracts battery&#39;s egress_suffix law, enforced at dispatch — the one law, never a bind-time second copy. | [optional] 
**Install** | Pointer to **string** |  | [optional] 
**MachineBox** | Pointer to **[]int32** | MachineBox decodes base64 automatically (encoding/json&#39;s []byte rule). | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**PushRepository** | Pointer to **string** |  | [optional] 
**RepoFull** | Pointer to **string** |  | [optional] 
**Runtime** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**VaultEnvelope** | Pointer to [**GitopsVaultSealedEnvelope**](GitopsVaultSealedEnvelope.md) | The sealed-delivery bind (issue #4328, design SD5; refined AC 6): the browser generated the HMAC secret locally and sealed BOTH blobs before upload — the vault row under the org vault key (the #4025 envelope, base64 over [nonce][ciphertext+tag] per blob) and the machine box (the anonymous sealed box, base64 over the wire framing) addressed to the verifying machine&#39;s delivery public key. The server stores both verbatim; plaintext never rides this wire in either direction. | [optional] 
**VerifyingAgentId** | Pointer to **int32** | VerifyingAgentID is the connection&#39;s ONE designated verifying machine (SD5&#39;s v1 designation rule). | [optional] 

## Methods

### NewHandlersGitBuildTriggerBindRequest

`func NewHandlersGitBuildTriggerBindRequest() *HandlersGitBuildTriggerBindRequest`

NewHandlersGitBuildTriggerBindRequest instantiates a new HandlersGitBuildTriggerBindRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersGitBuildTriggerBindRequestWithDefaults

`func NewHandlersGitBuildTriggerBindRequestWithDefaults() *HandlersGitBuildTriggerBindRequest`

NewHandlersGitBuildTriggerBindRequestWithDefaults instantiates a new HandlersGitBuildTriggerBindRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *HandlersGitBuildTriggerBindRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *HandlersGitBuildTriggerBindRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *HandlersGitBuildTriggerBindRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *HandlersGitBuildTriggerBindRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBuild

`func (o *HandlersGitBuildTriggerBindRequest) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *HandlersGitBuildTriggerBindRequest) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *HandlersGitBuildTriggerBindRequest) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *HandlersGitBuildTriggerBindRequest) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetBuilderImage

`func (o *HandlersGitBuildTriggerBindRequest) GetBuilderImage() string`

GetBuilderImage returns the BuilderImage field if non-nil, zero value otherwise.

### GetBuilderImageOk

`func (o *HandlersGitBuildTriggerBindRequest) GetBuilderImageOk() (*string, bool)`

GetBuilderImageOk returns a tuple with the BuilderImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderImage

`func (o *HandlersGitBuildTriggerBindRequest) SetBuilderImage(v string)`

SetBuilderImage sets BuilderImage field to given value.

### HasBuilderImage

`func (o *HandlersGitBuildTriggerBindRequest) HasBuilderImage() bool`

HasBuilderImage returns a boolean if a field has been set.

### GetEgressAllowlist

`func (o *HandlersGitBuildTriggerBindRequest) GetEgressAllowlist() []string`

GetEgressAllowlist returns the EgressAllowlist field if non-nil, zero value otherwise.

### GetEgressAllowlistOk

`func (o *HandlersGitBuildTriggerBindRequest) GetEgressAllowlistOk() (*[]string, bool)`

GetEgressAllowlistOk returns a tuple with the EgressAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressAllowlist

`func (o *HandlersGitBuildTriggerBindRequest) SetEgressAllowlist(v []string)`

SetEgressAllowlist sets EgressAllowlist field to given value.

### HasEgressAllowlist

`func (o *HandlersGitBuildTriggerBindRequest) HasEgressAllowlist() bool`

HasEgressAllowlist returns a boolean if a field has been set.

### GetInstall

`func (o *HandlersGitBuildTriggerBindRequest) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *HandlersGitBuildTriggerBindRequest) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *HandlersGitBuildTriggerBindRequest) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *HandlersGitBuildTriggerBindRequest) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetMachineBox

`func (o *HandlersGitBuildTriggerBindRequest) GetMachineBox() []int32`

GetMachineBox returns the MachineBox field if non-nil, zero value otherwise.

### GetMachineBoxOk

`func (o *HandlersGitBuildTriggerBindRequest) GetMachineBoxOk() (*[]int32, bool)`

GetMachineBoxOk returns a tuple with the MachineBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineBox

`func (o *HandlersGitBuildTriggerBindRequest) SetMachineBox(v []int32)`

SetMachineBox sets MachineBox field to given value.

### HasMachineBox

`func (o *HandlersGitBuildTriggerBindRequest) HasMachineBox() bool`

HasMachineBox returns a boolean if a field has been set.

### GetPort

`func (o *HandlersGitBuildTriggerBindRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlersGitBuildTriggerBindRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlersGitBuildTriggerBindRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HandlersGitBuildTriggerBindRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProjectId

`func (o *HandlersGitBuildTriggerBindRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HandlersGitBuildTriggerBindRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HandlersGitBuildTriggerBindRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *HandlersGitBuildTriggerBindRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPushRepository

`func (o *HandlersGitBuildTriggerBindRequest) GetPushRepository() string`

GetPushRepository returns the PushRepository field if non-nil, zero value otherwise.

### GetPushRepositoryOk

`func (o *HandlersGitBuildTriggerBindRequest) GetPushRepositoryOk() (*string, bool)`

GetPushRepositoryOk returns a tuple with the PushRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushRepository

`func (o *HandlersGitBuildTriggerBindRequest) SetPushRepository(v string)`

SetPushRepository sets PushRepository field to given value.

### HasPushRepository

`func (o *HandlersGitBuildTriggerBindRequest) HasPushRepository() bool`

HasPushRepository returns a boolean if a field has been set.

### GetRepoFull

`func (o *HandlersGitBuildTriggerBindRequest) GetRepoFull() string`

GetRepoFull returns the RepoFull field if non-nil, zero value otherwise.

### GetRepoFullOk

`func (o *HandlersGitBuildTriggerBindRequest) GetRepoFullOk() (*string, bool)`

GetRepoFullOk returns a tuple with the RepoFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoFull

`func (o *HandlersGitBuildTriggerBindRequest) SetRepoFull(v string)`

SetRepoFull sets RepoFull field to given value.

### HasRepoFull

`func (o *HandlersGitBuildTriggerBindRequest) HasRepoFull() bool`

HasRepoFull returns a boolean if a field has been set.

### GetRuntime

`func (o *HandlersGitBuildTriggerBindRequest) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *HandlersGitBuildTriggerBindRequest) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *HandlersGitBuildTriggerBindRequest) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *HandlersGitBuildTriggerBindRequest) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetStart

`func (o *HandlersGitBuildTriggerBindRequest) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HandlersGitBuildTriggerBindRequest) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HandlersGitBuildTriggerBindRequest) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *HandlersGitBuildTriggerBindRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStrategy

`func (o *HandlersGitBuildTriggerBindRequest) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *HandlersGitBuildTriggerBindRequest) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *HandlersGitBuildTriggerBindRequest) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *HandlersGitBuildTriggerBindRequest) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetVaultEnvelope

`func (o *HandlersGitBuildTriggerBindRequest) GetVaultEnvelope() GitopsVaultSealedEnvelope`

GetVaultEnvelope returns the VaultEnvelope field if non-nil, zero value otherwise.

### GetVaultEnvelopeOk

`func (o *HandlersGitBuildTriggerBindRequest) GetVaultEnvelopeOk() (*GitopsVaultSealedEnvelope, bool)`

GetVaultEnvelopeOk returns a tuple with the VaultEnvelope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultEnvelope

`func (o *HandlersGitBuildTriggerBindRequest) SetVaultEnvelope(v GitopsVaultSealedEnvelope)`

SetVaultEnvelope sets VaultEnvelope field to given value.

### HasVaultEnvelope

`func (o *HandlersGitBuildTriggerBindRequest) HasVaultEnvelope() bool`

HasVaultEnvelope returns a boolean if a field has been set.

### GetVerifyingAgentId

`func (o *HandlersGitBuildTriggerBindRequest) GetVerifyingAgentId() int32`

GetVerifyingAgentId returns the VerifyingAgentId field if non-nil, zero value otherwise.

### GetVerifyingAgentIdOk

`func (o *HandlersGitBuildTriggerBindRequest) GetVerifyingAgentIdOk() (*int32, bool)`

GetVerifyingAgentIdOk returns a tuple with the VerifyingAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyingAgentId

`func (o *HandlersGitBuildTriggerBindRequest) SetVerifyingAgentId(v int32)`

SetVerifyingAgentId sets VerifyingAgentId field to given value.

### HasVerifyingAgentId

`func (o *HandlersGitBuildTriggerBindRequest) HasVerifyingAgentId() bool`

HasVerifyingAgentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


