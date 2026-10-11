# ModelsDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupBucketName** | Pointer to **string** |  | [optional] 
**BackupEnabled** | Pointer to **bool** | Backups | [optional] 
**BackupSchedule** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **int32** |  | [optional] 
**CpuLimit** | Pointer to **string** | e.g. 1000m | [optional] 
**CpuRequest** | Pointer to **string** | Resource configuration | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**HaEnabled** | Pointer to **bool** | High availability | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastUsedAt** | Pointer to **string** |  | [optional] 
**MemoryLimit** | Pointer to **string** | e.g. 1Gi | [optional] 
**MemoryRequest** | Pointer to **string** | e.g. 128Mi | [optional] 
**Name** | Pointer to **string** | Basic configuration | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProvisionedAt** | Pointer to **string** |  | [optional] 
**ReplicaCount** | Pointer to **int32** |  | [optional] 
**S3ConnectionId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**ModelsDatabaseStatus**](ModelsDatabaseStatus.md) | Status | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**StorageClass** | Pointer to **string** |  | [optional] 
**StorageSize** | Pointer to **string** | Storage configuration | [optional] 
**Type** | Pointer to [**ModelsDatabaseType**](ModelsDatabaseType.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsDatabase

`func NewModelsDatabase() *ModelsDatabase`

NewModelsDatabase instantiates a new ModelsDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsDatabaseWithDefaults

`func NewModelsDatabaseWithDefaults() *ModelsDatabase`

NewModelsDatabaseWithDefaults instantiates a new ModelsDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupBucketName

`func (o *ModelsDatabase) GetBackupBucketName() string`

GetBackupBucketName returns the BackupBucketName field if non-nil, zero value otherwise.

### GetBackupBucketNameOk

`func (o *ModelsDatabase) GetBackupBucketNameOk() (*string, bool)`

GetBackupBucketNameOk returns a tuple with the BackupBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupBucketName

`func (o *ModelsDatabase) SetBackupBucketName(v string)`

SetBackupBucketName sets BackupBucketName field to given value.

### HasBackupBucketName

`func (o *ModelsDatabase) HasBackupBucketName() bool`

HasBackupBucketName returns a boolean if a field has been set.

### GetBackupEnabled

`func (o *ModelsDatabase) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *ModelsDatabase) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *ModelsDatabase) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *ModelsDatabase) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.

### GetBackupSchedule

`func (o *ModelsDatabase) GetBackupSchedule() string`

GetBackupSchedule returns the BackupSchedule field if non-nil, zero value otherwise.

### GetBackupScheduleOk

`func (o *ModelsDatabase) GetBackupScheduleOk() (*string, bool)`

GetBackupScheduleOk returns a tuple with the BackupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedule

`func (o *ModelsDatabase) SetBackupSchedule(v string)`

SetBackupSchedule sets BackupSchedule field to given value.

### HasBackupSchedule

`func (o *ModelsDatabase) HasBackupSchedule() bool`

HasBackupSchedule returns a boolean if a field has been set.

### GetClusterId

`func (o *ModelsDatabase) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ModelsDatabase) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ModelsDatabase) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ModelsDatabase) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCpuLimit

`func (o *ModelsDatabase) GetCpuLimit() string`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *ModelsDatabase) GetCpuLimitOk() (*string, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *ModelsDatabase) SetCpuLimit(v string)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *ModelsDatabase) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetCpuRequest

`func (o *ModelsDatabase) GetCpuRequest() string`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *ModelsDatabase) GetCpuRequestOk() (*string, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *ModelsDatabase) SetCpuRequest(v string)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *ModelsDatabase) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsDatabase) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsDatabase) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsDatabase) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsDatabase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsDatabase) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsDatabase) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsDatabase) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsDatabase) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetHaEnabled

`func (o *ModelsDatabase) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *ModelsDatabase) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *ModelsDatabase) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *ModelsDatabase) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetId

`func (o *ModelsDatabase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsDatabase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsDatabase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsDatabase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ModelsDatabase) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ModelsDatabase) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ModelsDatabase) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ModelsDatabase) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *ModelsDatabase) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *ModelsDatabase) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *ModelsDatabase) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *ModelsDatabase) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMemoryRequest

`func (o *ModelsDatabase) GetMemoryRequest() string`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *ModelsDatabase) GetMemoryRequestOk() (*string, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *ModelsDatabase) SetMemoryRequest(v string)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *ModelsDatabase) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.

### GetName

`func (o *ModelsDatabase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsDatabase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsDatabase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsDatabase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ModelsDatabase) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ModelsDatabase) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ModelsDatabase) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ModelsDatabase) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsDatabase) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsDatabase) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsDatabase) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsDatabase) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsDatabase) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsDatabase) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsDatabase) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsDatabase) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProvisionedAt

`func (o *ModelsDatabase) GetProvisionedAt() string`

GetProvisionedAt returns the ProvisionedAt field if non-nil, zero value otherwise.

### GetProvisionedAtOk

`func (o *ModelsDatabase) GetProvisionedAtOk() (*string, bool)`

GetProvisionedAtOk returns a tuple with the ProvisionedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedAt

`func (o *ModelsDatabase) SetProvisionedAt(v string)`

SetProvisionedAt sets ProvisionedAt field to given value.

### HasProvisionedAt

`func (o *ModelsDatabase) HasProvisionedAt() bool`

HasProvisionedAt returns a boolean if a field has been set.

### GetReplicaCount

`func (o *ModelsDatabase) GetReplicaCount() int32`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *ModelsDatabase) GetReplicaCountOk() (*int32, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *ModelsDatabase) SetReplicaCount(v int32)`

SetReplicaCount sets ReplicaCount field to given value.

### HasReplicaCount

`func (o *ModelsDatabase) HasReplicaCount() bool`

HasReplicaCount returns a boolean if a field has been set.

### GetS3ConnectionId

`func (o *ModelsDatabase) GetS3ConnectionId() int32`

GetS3ConnectionId returns the S3ConnectionId field if non-nil, zero value otherwise.

### GetS3ConnectionIdOk

`func (o *ModelsDatabase) GetS3ConnectionIdOk() (*int32, bool)`

GetS3ConnectionIdOk returns a tuple with the S3ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ConnectionId

`func (o *ModelsDatabase) SetS3ConnectionId(v int32)`

SetS3ConnectionId sets S3ConnectionId field to given value.

### HasS3ConnectionId

`func (o *ModelsDatabase) HasS3ConnectionId() bool`

HasS3ConnectionId returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsDatabase) GetStatus() ModelsDatabaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsDatabase) GetStatusOk() (*ModelsDatabaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsDatabase) SetStatus(v ModelsDatabaseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsDatabase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ModelsDatabase) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ModelsDatabase) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ModelsDatabase) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ModelsDatabase) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStorageClass

`func (o *ModelsDatabase) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *ModelsDatabase) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *ModelsDatabase) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *ModelsDatabase) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetStorageSize

`func (o *ModelsDatabase) GetStorageSize() string`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *ModelsDatabase) GetStorageSizeOk() (*string, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *ModelsDatabase) SetStorageSize(v string)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *ModelsDatabase) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetType

`func (o *ModelsDatabase) GetType() ModelsDatabaseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelsDatabase) GetTypeOk() (*ModelsDatabaseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelsDatabase) SetType(v ModelsDatabaseType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelsDatabase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsDatabase) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsDatabase) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsDatabase) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsDatabase) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *ModelsDatabase) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelsDatabase) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelsDatabase) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelsDatabase) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


