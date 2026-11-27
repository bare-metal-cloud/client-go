# ModelsGitConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **string** | APIBase is the connection&#39;s optional API base — the self-hosted Forgejo instance ROOT (issue #3644; the provider client appends /api/v1). A plain URL, never a secret: it rides the row, NOT the vault. Empty &#x3D; the provider default (Codeberg for forgejo). Only a forgejo row may carry a non-empty base; the service refuses one on the pinned-base providers (github/gitlab) and the plain-git provider — at write time and again at every use. | [optional] 
**AuthKind** | Pointer to **string** | AuthKind is the credential SHAPE (story 145.2): &#39;token&#39; (the 143.1 default — every pre-145.2 row reads back here) or &#39;ssh_key&#39; (SSH private-key material for the plain-git engine path). Material lives ONLY in the vault either way. | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastTestedAt** | Pointer to **string** |  | [optional] 
**Login** | Pointer to **string** | Login is the derived view of a GitHub App installation connection (BAR-5959): the GitHub account login the picker surfaces (\&quot;installed for acme\&quot;). Not a column — app rows carry it in Name behind the \&quot;github-app:\&quot; prefix; hand rows carry none. Populated by the gitops service&#39;s list/get paths. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Status mirrors ImageRegistry&#39;s connection lifecycle. LastTestedAt is the tested-at stamp Law 7&#39;s reconnect state reads — never guessed. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VerifyingAgentId** | Pointer to **int32** | VerifyingAgentID is the connection&#39;s ONE designated verifying machine (issue #4328; design zero-trust-sealed-delivery-2026-11-10 SD5, the v1 designation rule: one verifying machine pinned per git connection, chosen at bind from the org&#39;s trusted connected machines). The webhook edge relays every raw push to this machine for HMAC verification — the machine is the only holder of the opened secret (its sealed-delivery row). Nil &#x3D; no verifying machine pinned (a connection bound before the sealed-delivery enrollment, or whose pin was never set): the edge answers 503 — fail closed — exactly as it does when the pinned machine is untrusted or unreachable. Moving the pin is a re-seal (the browser re-seals to the new machine&#39;s delivery key); revoking the pinned machine refuses every delivery (SD4). | [optional] 

## Methods

### NewModelsGitConnection

`func NewModelsGitConnection() *ModelsGitConnection`

NewModelsGitConnection instantiates a new ModelsGitConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsGitConnectionWithDefaults

`func NewModelsGitConnectionWithDefaults() *ModelsGitConnection`

NewModelsGitConnectionWithDefaults instantiates a new ModelsGitConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *ModelsGitConnection) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *ModelsGitConnection) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *ModelsGitConnection) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *ModelsGitConnection) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### GetAuthKind

`func (o *ModelsGitConnection) GetAuthKind() string`

GetAuthKind returns the AuthKind field if non-nil, zero value otherwise.

### GetAuthKindOk

`func (o *ModelsGitConnection) GetAuthKindOk() (*string, bool)`

GetAuthKindOk returns a tuple with the AuthKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKind

`func (o *ModelsGitConnection) SetAuthKind(v string)`

SetAuthKind sets AuthKind field to given value.

### HasAuthKind

`func (o *ModelsGitConnection) HasAuthKind() bool`

HasAuthKind returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsGitConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsGitConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsGitConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsGitConnection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsGitConnection) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsGitConnection) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsGitConnection) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsGitConnection) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetId

`func (o *ModelsGitConnection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsGitConnection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsGitConnection) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsGitConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastTestedAt

`func (o *ModelsGitConnection) GetLastTestedAt() string`

GetLastTestedAt returns the LastTestedAt field if non-nil, zero value otherwise.

### GetLastTestedAtOk

`func (o *ModelsGitConnection) GetLastTestedAtOk() (*string, bool)`

GetLastTestedAtOk returns a tuple with the LastTestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestedAt

`func (o *ModelsGitConnection) SetLastTestedAt(v string)`

SetLastTestedAt sets LastTestedAt field to given value.

### HasLastTestedAt

`func (o *ModelsGitConnection) HasLastTestedAt() bool`

HasLastTestedAt returns a boolean if a field has been set.

### GetLogin

`func (o *ModelsGitConnection) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *ModelsGitConnection) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *ModelsGitConnection) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *ModelsGitConnection) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetName

`func (o *ModelsGitConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsGitConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsGitConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsGitConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsGitConnection) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsGitConnection) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsGitConnection) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsGitConnection) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProvider

`func (o *ModelsGitConnection) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelsGitConnection) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelsGitConnection) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelsGitConnection) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsGitConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsGitConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsGitConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsGitConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsGitConnection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsGitConnection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsGitConnection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsGitConnection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVerifyingAgentId

`func (o *ModelsGitConnection) GetVerifyingAgentId() int32`

GetVerifyingAgentId returns the VerifyingAgentId field if non-nil, zero value otherwise.

### GetVerifyingAgentIdOk

`func (o *ModelsGitConnection) GetVerifyingAgentIdOk() (*int32, bool)`

GetVerifyingAgentIdOk returns a tuple with the VerifyingAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyingAgentId

`func (o *ModelsGitConnection) SetVerifyingAgentId(v int32)`

SetVerifyingAgentId sets VerifyingAgentId field to given value.

### HasVerifyingAgentId

`func (o *ModelsGitConnection) HasVerifyingAgentId() bool`

HasVerifyingAgentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


