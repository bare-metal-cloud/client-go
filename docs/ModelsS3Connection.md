# ModelsS3Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Issue #3576 — Azure Blob (azblob) fields. Account is the storage account name (the endpoint synthesizes from it) and Container is the blob container; both are empty for every other provider. Bucket/region semantics do not apply to azblob rows. | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**DefaultBucket** | Pointer to **string** | Optional: Default bucket for this connection | [optional] 
**Endpoint** | Pointer to **string** | S3 endpoint (auto-filled based on provider, editable for \&quot;other\&quot;) AWS: https://s3.&lt;region&gt;.amazonaws.com MinIO: https://&lt;your-minio-host&gt;:9000 Wasabi: https://s3.wasabisys.com GCS: https://storage.googleapis.com azblob: https://&lt;account&gt;.blob.core.windows.net (synthesized at create when the request omits it — the column is not null, so the synthesized value is stored). | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastUsedAt** | Pointer to **string** | Metadata | [optional] 
**LinkedCredentialId** | Pointer to **int32** | Issues #3576 (azblob) and #3874 (the S3 family) — link mode: when set, the connection carries NO vault credentials of its own (CredentialsEncryptedDataID stays nil) and authenticates at use time through the linked CloudProviderCredential. azblob links an azure credential&#39;s service principal (the ACR image-registry shape); the S3 family links an aws-shaped access-key pair (gcs links a gcp credential only when it carries HMAC keys). Accepted pairs: aws/minio/wasabi/other link an aws credential, gcs links a gcp credential, azblob links an azure credential. | [optional] 
**Name** | Pointer to **string** | Connection identification | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** | Story 104.3: ProjectID is now nullable (*uint) so a row may be org-scoped (ProjectID &#x3D; nil) instead of project-scoped. The org-vs- project distinction is one branch: &#x60;project_id IS NULL&#x60; &#x3D; org-scoped. AutoMigrate applies the NOT NULL → nullable change (dev-stage policy). | [optional] 
**Provider** | Pointer to **string** | aws, minio, wasabi, gcs, other, azblob | [optional] 
**Region** | Pointer to **string** | Issue #3874 — the S3-family region. Key-mode rows also carry it inside the sealed credential blob (that is where the probe read it before), but link-mode rows store NO blob of their own, so the region the create validated has nowhere else to live. Explore, the unified test-connection check and the catalog capture read it here. Empty for azblob rows. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsS3Connection

`func NewModelsS3Connection() *ModelsS3Connection`

NewModelsS3Connection instantiates a new ModelsS3Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsS3ConnectionWithDefaults

`func NewModelsS3ConnectionWithDefaults() *ModelsS3Connection`

NewModelsS3ConnectionWithDefaults instantiates a new ModelsS3Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ModelsS3Connection) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ModelsS3Connection) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ModelsS3Connection) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ModelsS3Connection) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetContainer

`func (o *ModelsS3Connection) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelsS3Connection) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelsS3Connection) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelsS3Connection) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsS3Connection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsS3Connection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsS3Connection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsS3Connection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsS3Connection) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsS3Connection) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsS3Connection) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsS3Connection) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDefaultBucket

`func (o *ModelsS3Connection) GetDefaultBucket() string`

GetDefaultBucket returns the DefaultBucket field if non-nil, zero value otherwise.

### GetDefaultBucketOk

`func (o *ModelsS3Connection) GetDefaultBucketOk() (*string, bool)`

GetDefaultBucketOk returns a tuple with the DefaultBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBucket

`func (o *ModelsS3Connection) SetDefaultBucket(v string)`

SetDefaultBucket sets DefaultBucket field to given value.

### HasDefaultBucket

`func (o *ModelsS3Connection) HasDefaultBucket() bool`

HasDefaultBucket returns a boolean if a field has been set.

### GetEndpoint

`func (o *ModelsS3Connection) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ModelsS3Connection) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ModelsS3Connection) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ModelsS3Connection) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetId

`func (o *ModelsS3Connection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsS3Connection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsS3Connection) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsS3Connection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ModelsS3Connection) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ModelsS3Connection) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ModelsS3Connection) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ModelsS3Connection) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetLinkedCredentialId

`func (o *ModelsS3Connection) GetLinkedCredentialId() int32`

GetLinkedCredentialId returns the LinkedCredentialId field if non-nil, zero value otherwise.

### GetLinkedCredentialIdOk

`func (o *ModelsS3Connection) GetLinkedCredentialIdOk() (*int32, bool)`

GetLinkedCredentialIdOk returns a tuple with the LinkedCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCredentialId

`func (o *ModelsS3Connection) SetLinkedCredentialId(v int32)`

SetLinkedCredentialId sets LinkedCredentialId field to given value.

### HasLinkedCredentialId

`func (o *ModelsS3Connection) HasLinkedCredentialId() bool`

HasLinkedCredentialId returns a boolean if a field has been set.

### GetName

`func (o *ModelsS3Connection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsS3Connection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsS3Connection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsS3Connection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsS3Connection) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsS3Connection) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsS3Connection) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsS3Connection) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsS3Connection) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsS3Connection) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsS3Connection) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsS3Connection) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProvider

`func (o *ModelsS3Connection) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelsS3Connection) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelsS3Connection) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelsS3Connection) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *ModelsS3Connection) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ModelsS3Connection) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ModelsS3Connection) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ModelsS3Connection) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsS3Connection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsS3Connection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsS3Connection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsS3Connection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


