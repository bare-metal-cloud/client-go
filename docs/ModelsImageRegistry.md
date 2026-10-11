# ModelsImageRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **string** | AuthMode is the explicit auth contract (Story 79.1). Empty on legacy rows; ResolveAuthMode() derives from Type (+ FK) until backfilled by writes. | [optional] 
**AwsCredentialId** | Pointer to **int32** | For ECR: AWS credential reference and region | [optional] 
**AwsRegion** | Pointer to **string** | Renamed from ECRRegion for consistency | [optional] 
**AzureCredentialId** | Pointer to **int32** | For ACR: Azure credential reference and login server (for username/password auth) | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**GcpCredentialId** | Pointer to **int32** | For GCR: GCP credential reference | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastTestedAt** | Pointer to **string** |  | [optional] 
**LoginServer** | Pointer to **string** | For ACR username/password auth | [optional] 
**Name** | Pointer to **string** | Connection identification | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | Metadata | [optional] 
**Type** | Pointer to **string** | Registry type constant | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | Registry endpoint (auto-filled for known types, editable for private) | [optional] 
**Username** | Pointer to **string** | Username for basic auth (Docker Hub, GHCR, private registries) Password is stored encrypted using polymorphic encrypted_data table | [optional] 

## Methods

### NewModelsImageRegistry

`func NewModelsImageRegistry() *ModelsImageRegistry`

NewModelsImageRegistry instantiates a new ModelsImageRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsImageRegistryWithDefaults

`func NewModelsImageRegistryWithDefaults() *ModelsImageRegistry`

NewModelsImageRegistryWithDefaults instantiates a new ModelsImageRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *ModelsImageRegistry) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *ModelsImageRegistry) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *ModelsImageRegistry) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *ModelsImageRegistry) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAwsCredentialId

`func (o *ModelsImageRegistry) GetAwsCredentialId() int32`

GetAwsCredentialId returns the AwsCredentialId field if non-nil, zero value otherwise.

### GetAwsCredentialIdOk

`func (o *ModelsImageRegistry) GetAwsCredentialIdOk() (*int32, bool)`

GetAwsCredentialIdOk returns a tuple with the AwsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCredentialId

`func (o *ModelsImageRegistry) SetAwsCredentialId(v int32)`

SetAwsCredentialId sets AwsCredentialId field to given value.

### HasAwsCredentialId

`func (o *ModelsImageRegistry) HasAwsCredentialId() bool`

HasAwsCredentialId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *ModelsImageRegistry) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *ModelsImageRegistry) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *ModelsImageRegistry) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *ModelsImageRegistry) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAzureCredentialId

`func (o *ModelsImageRegistry) GetAzureCredentialId() int32`

GetAzureCredentialId returns the AzureCredentialId field if non-nil, zero value otherwise.

### GetAzureCredentialIdOk

`func (o *ModelsImageRegistry) GetAzureCredentialIdOk() (*int32, bool)`

GetAzureCredentialIdOk returns a tuple with the AzureCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCredentialId

`func (o *ModelsImageRegistry) SetAzureCredentialId(v int32)`

SetAzureCredentialId sets AzureCredentialId field to given value.

### HasAzureCredentialId

`func (o *ModelsImageRegistry) HasAzureCredentialId() bool`

HasAzureCredentialId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsImageRegistry) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsImageRegistry) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsImageRegistry) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsImageRegistry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsImageRegistry) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsImageRegistry) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsImageRegistry) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsImageRegistry) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetGcpCredentialId

`func (o *ModelsImageRegistry) GetGcpCredentialId() int32`

GetGcpCredentialId returns the GcpCredentialId field if non-nil, zero value otherwise.

### GetGcpCredentialIdOk

`func (o *ModelsImageRegistry) GetGcpCredentialIdOk() (*int32, bool)`

GetGcpCredentialIdOk returns a tuple with the GcpCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredentialId

`func (o *ModelsImageRegistry) SetGcpCredentialId(v int32)`

SetGcpCredentialId sets GcpCredentialId field to given value.

### HasGcpCredentialId

`func (o *ModelsImageRegistry) HasGcpCredentialId() bool`

HasGcpCredentialId returns a boolean if a field has been set.

### GetId

`func (o *ModelsImageRegistry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsImageRegistry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsImageRegistry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsImageRegistry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastTestedAt

`func (o *ModelsImageRegistry) GetLastTestedAt() string`

GetLastTestedAt returns the LastTestedAt field if non-nil, zero value otherwise.

### GetLastTestedAtOk

`func (o *ModelsImageRegistry) GetLastTestedAtOk() (*string, bool)`

GetLastTestedAtOk returns a tuple with the LastTestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestedAt

`func (o *ModelsImageRegistry) SetLastTestedAt(v string)`

SetLastTestedAt sets LastTestedAt field to given value.

### HasLastTestedAt

`func (o *ModelsImageRegistry) HasLastTestedAt() bool`

HasLastTestedAt returns a boolean if a field has been set.

### GetLoginServer

`func (o *ModelsImageRegistry) GetLoginServer() string`

GetLoginServer returns the LoginServer field if non-nil, zero value otherwise.

### GetLoginServerOk

`func (o *ModelsImageRegistry) GetLoginServerOk() (*string, bool)`

GetLoginServerOk returns a tuple with the LoginServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginServer

`func (o *ModelsImageRegistry) SetLoginServer(v string)`

SetLoginServer sets LoginServer field to given value.

### HasLoginServer

`func (o *ModelsImageRegistry) HasLoginServer() bool`

HasLoginServer returns a boolean if a field has been set.

### GetName

`func (o *ModelsImageRegistry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsImageRegistry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsImageRegistry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsImageRegistry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsImageRegistry) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsImageRegistry) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsImageRegistry) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsImageRegistry) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsImageRegistry) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsImageRegistry) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsImageRegistry) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsImageRegistry) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsImageRegistry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsImageRegistry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsImageRegistry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsImageRegistry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ModelsImageRegistry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelsImageRegistry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelsImageRegistry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelsImageRegistry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsImageRegistry) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsImageRegistry) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsImageRegistry) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsImageRegistry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *ModelsImageRegistry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelsImageRegistry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelsImageRegistry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ModelsImageRegistry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsImageRegistry) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsImageRegistry) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsImageRegistry) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsImageRegistry) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


