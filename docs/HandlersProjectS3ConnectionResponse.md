# HandlersProjectS3ConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Issue #3576 — azblob fields (empty for the S3-family providers). | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to **map[string]string** | Credentials is ALWAYS a masked view — never decrypted on this path. | [optional] 
**DefaultBucket** | Pointer to **map[string]interface{}** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**HasVaultCredentials** | Pointer to **bool** | HasVaultCredentials tells the FE whether a \&quot;rotate\&quot; affordance applies. | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastUsedAt** | Pointer to **map[string]interface{}** |  | [optional] 
**LastVerifiedAt** | Pointer to **string** | LastVerifiedAt is RFC3339 (UTC) or \&quot;\&quot; when the connection has never passed a check (BAR-5678). Populated by the LIST path only — the single-row mutations answer \&quot;\&quot;, the same law as the org responses. | [optional] 
**LinkedCredentialId** | Pointer to **int32** | LinkedCredentialID is set on link-mode rows of every provider. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** | Region is the S3-family region (issue #3874); empty for azblob rows. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersProjectS3ConnectionResponse

`func NewHandlersProjectS3ConnectionResponse() *HandlersProjectS3ConnectionResponse`

NewHandlersProjectS3ConnectionResponse instantiates a new HandlersProjectS3ConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersProjectS3ConnectionResponseWithDefaults

`func NewHandlersProjectS3ConnectionResponseWithDefaults() *HandlersProjectS3ConnectionResponse`

NewHandlersProjectS3ConnectionResponseWithDefaults instantiates a new HandlersProjectS3ConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *HandlersProjectS3ConnectionResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HandlersProjectS3ConnectionResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HandlersProjectS3ConnectionResponse) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *HandlersProjectS3ConnectionResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetContainer

`func (o *HandlersProjectS3ConnectionResponse) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *HandlersProjectS3ConnectionResponse) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *HandlersProjectS3ConnectionResponse) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *HandlersProjectS3ConnectionResponse) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersProjectS3ConnectionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersProjectS3ConnectionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersProjectS3ConnectionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersProjectS3ConnectionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentials

`func (o *HandlersProjectS3ConnectionResponse) GetCredentials() map[string]string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *HandlersProjectS3ConnectionResponse) GetCredentialsOk() (*map[string]string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *HandlersProjectS3ConnectionResponse) SetCredentials(v map[string]string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *HandlersProjectS3ConnectionResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDefaultBucket

`func (o *HandlersProjectS3ConnectionResponse) GetDefaultBucket() map[string]interface{}`

GetDefaultBucket returns the DefaultBucket field if non-nil, zero value otherwise.

### GetDefaultBucketOk

`func (o *HandlersProjectS3ConnectionResponse) GetDefaultBucketOk() (*map[string]interface{}, bool)`

GetDefaultBucketOk returns a tuple with the DefaultBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBucket

`func (o *HandlersProjectS3ConnectionResponse) SetDefaultBucket(v map[string]interface{})`

SetDefaultBucket sets DefaultBucket field to given value.

### HasDefaultBucket

`func (o *HandlersProjectS3ConnectionResponse) HasDefaultBucket() bool`

HasDefaultBucket returns a boolean if a field has been set.

### GetEndpoint

`func (o *HandlersProjectS3ConnectionResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HandlersProjectS3ConnectionResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HandlersProjectS3ConnectionResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HandlersProjectS3ConnectionResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHasVaultCredentials

`func (o *HandlersProjectS3ConnectionResponse) GetHasVaultCredentials() bool`

GetHasVaultCredentials returns the HasVaultCredentials field if non-nil, zero value otherwise.

### GetHasVaultCredentialsOk

`func (o *HandlersProjectS3ConnectionResponse) GetHasVaultCredentialsOk() (*bool, bool)`

GetHasVaultCredentialsOk returns a tuple with the HasVaultCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVaultCredentials

`func (o *HandlersProjectS3ConnectionResponse) SetHasVaultCredentials(v bool)`

SetHasVaultCredentials sets HasVaultCredentials field to given value.

### HasHasVaultCredentials

`func (o *HandlersProjectS3ConnectionResponse) HasHasVaultCredentials() bool`

HasHasVaultCredentials returns a boolean if a field has been set.

### GetId

`func (o *HandlersProjectS3ConnectionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersProjectS3ConnectionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersProjectS3ConnectionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersProjectS3ConnectionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *HandlersProjectS3ConnectionResponse) GetLastUsedAt() map[string]interface{}`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *HandlersProjectS3ConnectionResponse) GetLastUsedAtOk() (*map[string]interface{}, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *HandlersProjectS3ConnectionResponse) SetLastUsedAt(v map[string]interface{})`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *HandlersProjectS3ConnectionResponse) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetLastVerifiedAt

`func (o *HandlersProjectS3ConnectionResponse) GetLastVerifiedAt() string`

GetLastVerifiedAt returns the LastVerifiedAt field if non-nil, zero value otherwise.

### GetLastVerifiedAtOk

`func (o *HandlersProjectS3ConnectionResponse) GetLastVerifiedAtOk() (*string, bool)`

GetLastVerifiedAtOk returns a tuple with the LastVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifiedAt

`func (o *HandlersProjectS3ConnectionResponse) SetLastVerifiedAt(v string)`

SetLastVerifiedAt sets LastVerifiedAt field to given value.

### HasLastVerifiedAt

`func (o *HandlersProjectS3ConnectionResponse) HasLastVerifiedAt() bool`

HasLastVerifiedAt returns a boolean if a field has been set.

### GetLinkedCredentialId

`func (o *HandlersProjectS3ConnectionResponse) GetLinkedCredentialId() int32`

GetLinkedCredentialId returns the LinkedCredentialId field if non-nil, zero value otherwise.

### GetLinkedCredentialIdOk

`func (o *HandlersProjectS3ConnectionResponse) GetLinkedCredentialIdOk() (*int32, bool)`

GetLinkedCredentialIdOk returns a tuple with the LinkedCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCredentialId

`func (o *HandlersProjectS3ConnectionResponse) SetLinkedCredentialId(v int32)`

SetLinkedCredentialId sets LinkedCredentialId field to given value.

### HasLinkedCredentialId

`func (o *HandlersProjectS3ConnectionResponse) HasLinkedCredentialId() bool`

HasLinkedCredentialId returns a boolean if a field has been set.

### GetName

`func (o *HandlersProjectS3ConnectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersProjectS3ConnectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersProjectS3ConnectionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersProjectS3ConnectionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *HandlersProjectS3ConnectionResponse) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HandlersProjectS3ConnectionResponse) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HandlersProjectS3ConnectionResponse) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *HandlersProjectS3ConnectionResponse) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProvider

`func (o *HandlersProjectS3ConnectionResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersProjectS3ConnectionResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersProjectS3ConnectionResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *HandlersProjectS3ConnectionResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *HandlersProjectS3ConnectionResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersProjectS3ConnectionResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersProjectS3ConnectionResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HandlersProjectS3ConnectionResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HandlersProjectS3ConnectionResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HandlersProjectS3ConnectionResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HandlersProjectS3ConnectionResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HandlersProjectS3ConnectionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


