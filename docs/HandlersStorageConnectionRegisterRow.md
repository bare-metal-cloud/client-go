# HandlersStorageConnectionRegisterRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to **map[string]string** | Always masked — never the plaintext credential material. | [optional] 
**DefaultBucket** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**HasVaultCredentials** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastUsedAt** | Pointer to **string** |  | [optional] 
**LastVerifiedAt** | Pointer to **string** | RFC3339 UTC of the last SUCCESSFUL check; \&quot;\&quot; when never verified (FR-Z-13: an unreported value is empty, never faked). | [optional] 
**LinkedCredentialId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** | FR-Z-3: the row&#39;s real scope — \&quot;org\&quot; (ProjectID nil) or \&quot;project\&quot;. Not the ADR-067 secrets-tier hard-coded value: the storage register reintroduces the real distinction per FR-Z-3. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersStorageConnectionRegisterRow

`func NewHandlersStorageConnectionRegisterRow() *HandlersStorageConnectionRegisterRow`

NewHandlersStorageConnectionRegisterRow instantiates a new HandlersStorageConnectionRegisterRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStorageConnectionRegisterRowWithDefaults

`func NewHandlersStorageConnectionRegisterRowWithDefaults() *HandlersStorageConnectionRegisterRow`

NewHandlersStorageConnectionRegisterRowWithDefaults instantiates a new HandlersStorageConnectionRegisterRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *HandlersStorageConnectionRegisterRow) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HandlersStorageConnectionRegisterRow) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HandlersStorageConnectionRegisterRow) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *HandlersStorageConnectionRegisterRow) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetContainer

`func (o *HandlersStorageConnectionRegisterRow) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *HandlersStorageConnectionRegisterRow) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *HandlersStorageConnectionRegisterRow) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *HandlersStorageConnectionRegisterRow) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersStorageConnectionRegisterRow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersStorageConnectionRegisterRow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersStorageConnectionRegisterRow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersStorageConnectionRegisterRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentials

`func (o *HandlersStorageConnectionRegisterRow) GetCredentials() map[string]string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *HandlersStorageConnectionRegisterRow) GetCredentialsOk() (*map[string]string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *HandlersStorageConnectionRegisterRow) SetCredentials(v map[string]string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *HandlersStorageConnectionRegisterRow) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDefaultBucket

`func (o *HandlersStorageConnectionRegisterRow) GetDefaultBucket() string`

GetDefaultBucket returns the DefaultBucket field if non-nil, zero value otherwise.

### GetDefaultBucketOk

`func (o *HandlersStorageConnectionRegisterRow) GetDefaultBucketOk() (*string, bool)`

GetDefaultBucketOk returns a tuple with the DefaultBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBucket

`func (o *HandlersStorageConnectionRegisterRow) SetDefaultBucket(v string)`

SetDefaultBucket sets DefaultBucket field to given value.

### HasDefaultBucket

`func (o *HandlersStorageConnectionRegisterRow) HasDefaultBucket() bool`

HasDefaultBucket returns a boolean if a field has been set.

### GetEndpoint

`func (o *HandlersStorageConnectionRegisterRow) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HandlersStorageConnectionRegisterRow) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HandlersStorageConnectionRegisterRow) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HandlersStorageConnectionRegisterRow) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHasVaultCredentials

`func (o *HandlersStorageConnectionRegisterRow) GetHasVaultCredentials() bool`

GetHasVaultCredentials returns the HasVaultCredentials field if non-nil, zero value otherwise.

### GetHasVaultCredentialsOk

`func (o *HandlersStorageConnectionRegisterRow) GetHasVaultCredentialsOk() (*bool, bool)`

GetHasVaultCredentialsOk returns a tuple with the HasVaultCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVaultCredentials

`func (o *HandlersStorageConnectionRegisterRow) SetHasVaultCredentials(v bool)`

SetHasVaultCredentials sets HasVaultCredentials field to given value.

### HasHasVaultCredentials

`func (o *HandlersStorageConnectionRegisterRow) HasHasVaultCredentials() bool`

HasHasVaultCredentials returns a boolean if a field has been set.

### GetId

`func (o *HandlersStorageConnectionRegisterRow) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersStorageConnectionRegisterRow) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersStorageConnectionRegisterRow) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersStorageConnectionRegisterRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *HandlersStorageConnectionRegisterRow) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *HandlersStorageConnectionRegisterRow) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *HandlersStorageConnectionRegisterRow) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *HandlersStorageConnectionRegisterRow) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetLastVerifiedAt

`func (o *HandlersStorageConnectionRegisterRow) GetLastVerifiedAt() string`

GetLastVerifiedAt returns the LastVerifiedAt field if non-nil, zero value otherwise.

### GetLastVerifiedAtOk

`func (o *HandlersStorageConnectionRegisterRow) GetLastVerifiedAtOk() (*string, bool)`

GetLastVerifiedAtOk returns a tuple with the LastVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifiedAt

`func (o *HandlersStorageConnectionRegisterRow) SetLastVerifiedAt(v string)`

SetLastVerifiedAt sets LastVerifiedAt field to given value.

### HasLastVerifiedAt

`func (o *HandlersStorageConnectionRegisterRow) HasLastVerifiedAt() bool`

HasLastVerifiedAt returns a boolean if a field has been set.

### GetLinkedCredentialId

`func (o *HandlersStorageConnectionRegisterRow) GetLinkedCredentialId() int32`

GetLinkedCredentialId returns the LinkedCredentialId field if non-nil, zero value otherwise.

### GetLinkedCredentialIdOk

`func (o *HandlersStorageConnectionRegisterRow) GetLinkedCredentialIdOk() (*int32, bool)`

GetLinkedCredentialIdOk returns a tuple with the LinkedCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCredentialId

`func (o *HandlersStorageConnectionRegisterRow) SetLinkedCredentialId(v int32)`

SetLinkedCredentialId sets LinkedCredentialId field to given value.

### HasLinkedCredentialId

`func (o *HandlersStorageConnectionRegisterRow) HasLinkedCredentialId() bool`

HasLinkedCredentialId returns a boolean if a field has been set.

### GetName

`func (o *HandlersStorageConnectionRegisterRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersStorageConnectionRegisterRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersStorageConnectionRegisterRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersStorageConnectionRegisterRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *HandlersStorageConnectionRegisterRow) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HandlersStorageConnectionRegisterRow) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HandlersStorageConnectionRegisterRow) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *HandlersStorageConnectionRegisterRow) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *HandlersStorageConnectionRegisterRow) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *HandlersStorageConnectionRegisterRow) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *HandlersStorageConnectionRegisterRow) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *HandlersStorageConnectionRegisterRow) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetProvider

`func (o *HandlersStorageConnectionRegisterRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersStorageConnectionRegisterRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersStorageConnectionRegisterRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *HandlersStorageConnectionRegisterRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *HandlersStorageConnectionRegisterRow) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersStorageConnectionRegisterRow) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersStorageConnectionRegisterRow) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HandlersStorageConnectionRegisterRow) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetScope

`func (o *HandlersStorageConnectionRegisterRow) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersStorageConnectionRegisterRow) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersStorageConnectionRegisterRow) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersStorageConnectionRegisterRow) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HandlersStorageConnectionRegisterRow) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HandlersStorageConnectionRegisterRow) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HandlersStorageConnectionRegisterRow) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HandlersStorageConnectionRegisterRow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


