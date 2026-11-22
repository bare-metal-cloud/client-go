# HandlersDatabaseStatusResponse

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
**MaxReplicaLagSeconds** | Pointer to **float32** |  | [optional] 
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

### NewHandlersDatabaseStatusResponse

`func NewHandlersDatabaseStatusResponse() *HandlersDatabaseStatusResponse`

NewHandlersDatabaseStatusResponse instantiates a new HandlersDatabaseStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDatabaseStatusResponseWithDefaults

`func NewHandlersDatabaseStatusResponseWithDefaults() *HandlersDatabaseStatusResponse`

NewHandlersDatabaseStatusResponseWithDefaults instantiates a new HandlersDatabaseStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupBucketName

`func (o *HandlersDatabaseStatusResponse) GetBackupBucketName() string`

GetBackupBucketName returns the BackupBucketName field if non-nil, zero value otherwise.

### GetBackupBucketNameOk

`func (o *HandlersDatabaseStatusResponse) GetBackupBucketNameOk() (*string, bool)`

GetBackupBucketNameOk returns a tuple with the BackupBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupBucketName

`func (o *HandlersDatabaseStatusResponse) SetBackupBucketName(v string)`

SetBackupBucketName sets BackupBucketName field to given value.

### HasBackupBucketName

`func (o *HandlersDatabaseStatusResponse) HasBackupBucketName() bool`

HasBackupBucketName returns a boolean if a field has been set.

### GetBackupEnabled

`func (o *HandlersDatabaseStatusResponse) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *HandlersDatabaseStatusResponse) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *HandlersDatabaseStatusResponse) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *HandlersDatabaseStatusResponse) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.

### GetBackupSchedule

`func (o *HandlersDatabaseStatusResponse) GetBackupSchedule() string`

GetBackupSchedule returns the BackupSchedule field if non-nil, zero value otherwise.

### GetBackupScheduleOk

`func (o *HandlersDatabaseStatusResponse) GetBackupScheduleOk() (*string, bool)`

GetBackupScheduleOk returns a tuple with the BackupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedule

`func (o *HandlersDatabaseStatusResponse) SetBackupSchedule(v string)`

SetBackupSchedule sets BackupSchedule field to given value.

### HasBackupSchedule

`func (o *HandlersDatabaseStatusResponse) HasBackupSchedule() bool`

HasBackupSchedule returns a boolean if a field has been set.

### GetClusterId

`func (o *HandlersDatabaseStatusResponse) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HandlersDatabaseStatusResponse) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HandlersDatabaseStatusResponse) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HandlersDatabaseStatusResponse) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCpuLimit

`func (o *HandlersDatabaseStatusResponse) GetCpuLimit() string`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *HandlersDatabaseStatusResponse) GetCpuLimitOk() (*string, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *HandlersDatabaseStatusResponse) SetCpuLimit(v string)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *HandlersDatabaseStatusResponse) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetCpuRequest

`func (o *HandlersDatabaseStatusResponse) GetCpuRequest() string`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *HandlersDatabaseStatusResponse) GetCpuRequestOk() (*string, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *HandlersDatabaseStatusResponse) SetCpuRequest(v string)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *HandlersDatabaseStatusResponse) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersDatabaseStatusResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersDatabaseStatusResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersDatabaseStatusResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersDatabaseStatusResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *HandlersDatabaseStatusResponse) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *HandlersDatabaseStatusResponse) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *HandlersDatabaseStatusResponse) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *HandlersDatabaseStatusResponse) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetHaEnabled

`func (o *HandlersDatabaseStatusResponse) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *HandlersDatabaseStatusResponse) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *HandlersDatabaseStatusResponse) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *HandlersDatabaseStatusResponse) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetId

`func (o *HandlersDatabaseStatusResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersDatabaseStatusResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersDatabaseStatusResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersDatabaseStatusResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *HandlersDatabaseStatusResponse) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *HandlersDatabaseStatusResponse) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *HandlersDatabaseStatusResponse) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *HandlersDatabaseStatusResponse) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetMaxReplicaLagSeconds

`func (o *HandlersDatabaseStatusResponse) GetMaxReplicaLagSeconds() float32`

GetMaxReplicaLagSeconds returns the MaxReplicaLagSeconds field if non-nil, zero value otherwise.

### GetMaxReplicaLagSecondsOk

`func (o *HandlersDatabaseStatusResponse) GetMaxReplicaLagSecondsOk() (*float32, bool)`

GetMaxReplicaLagSecondsOk returns a tuple with the MaxReplicaLagSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicaLagSeconds

`func (o *HandlersDatabaseStatusResponse) SetMaxReplicaLagSeconds(v float32)`

SetMaxReplicaLagSeconds sets MaxReplicaLagSeconds field to given value.

### HasMaxReplicaLagSeconds

`func (o *HandlersDatabaseStatusResponse) HasMaxReplicaLagSeconds() bool`

HasMaxReplicaLagSeconds returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *HandlersDatabaseStatusResponse) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *HandlersDatabaseStatusResponse) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *HandlersDatabaseStatusResponse) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *HandlersDatabaseStatusResponse) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMemoryRequest

`func (o *HandlersDatabaseStatusResponse) GetMemoryRequest() string`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *HandlersDatabaseStatusResponse) GetMemoryRequestOk() (*string, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *HandlersDatabaseStatusResponse) SetMemoryRequest(v string)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *HandlersDatabaseStatusResponse) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.

### GetName

`func (o *HandlersDatabaseStatusResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersDatabaseStatusResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersDatabaseStatusResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersDatabaseStatusResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *HandlersDatabaseStatusResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HandlersDatabaseStatusResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HandlersDatabaseStatusResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HandlersDatabaseStatusResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrgId

`func (o *HandlersDatabaseStatusResponse) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *HandlersDatabaseStatusResponse) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *HandlersDatabaseStatusResponse) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *HandlersDatabaseStatusResponse) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *HandlersDatabaseStatusResponse) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HandlersDatabaseStatusResponse) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HandlersDatabaseStatusResponse) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *HandlersDatabaseStatusResponse) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProvisionedAt

`func (o *HandlersDatabaseStatusResponse) GetProvisionedAt() string`

GetProvisionedAt returns the ProvisionedAt field if non-nil, zero value otherwise.

### GetProvisionedAtOk

`func (o *HandlersDatabaseStatusResponse) GetProvisionedAtOk() (*string, bool)`

GetProvisionedAtOk returns a tuple with the ProvisionedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedAt

`func (o *HandlersDatabaseStatusResponse) SetProvisionedAt(v string)`

SetProvisionedAt sets ProvisionedAt field to given value.

### HasProvisionedAt

`func (o *HandlersDatabaseStatusResponse) HasProvisionedAt() bool`

HasProvisionedAt returns a boolean if a field has been set.

### GetReplicaCount

`func (o *HandlersDatabaseStatusResponse) GetReplicaCount() int32`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *HandlersDatabaseStatusResponse) GetReplicaCountOk() (*int32, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *HandlersDatabaseStatusResponse) SetReplicaCount(v int32)`

SetReplicaCount sets ReplicaCount field to given value.

### HasReplicaCount

`func (o *HandlersDatabaseStatusResponse) HasReplicaCount() bool`

HasReplicaCount returns a boolean if a field has been set.

### GetS3ConnectionId

`func (o *HandlersDatabaseStatusResponse) GetS3ConnectionId() int32`

GetS3ConnectionId returns the S3ConnectionId field if non-nil, zero value otherwise.

### GetS3ConnectionIdOk

`func (o *HandlersDatabaseStatusResponse) GetS3ConnectionIdOk() (*int32, bool)`

GetS3ConnectionIdOk returns a tuple with the S3ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ConnectionId

`func (o *HandlersDatabaseStatusResponse) SetS3ConnectionId(v int32)`

SetS3ConnectionId sets S3ConnectionId field to given value.

### HasS3ConnectionId

`func (o *HandlersDatabaseStatusResponse) HasS3ConnectionId() bool`

HasS3ConnectionId returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersDatabaseStatusResponse) GetStatus() ModelsDatabaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersDatabaseStatusResponse) GetStatusOk() (*ModelsDatabaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersDatabaseStatusResponse) SetStatus(v ModelsDatabaseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersDatabaseStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *HandlersDatabaseStatusResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *HandlersDatabaseStatusResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *HandlersDatabaseStatusResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *HandlersDatabaseStatusResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStorageClass

`func (o *HandlersDatabaseStatusResponse) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *HandlersDatabaseStatusResponse) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *HandlersDatabaseStatusResponse) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *HandlersDatabaseStatusResponse) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetStorageSize

`func (o *HandlersDatabaseStatusResponse) GetStorageSize() string`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *HandlersDatabaseStatusResponse) GetStorageSizeOk() (*string, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *HandlersDatabaseStatusResponse) SetStorageSize(v string)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *HandlersDatabaseStatusResponse) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetType

`func (o *HandlersDatabaseStatusResponse) GetType() ModelsDatabaseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HandlersDatabaseStatusResponse) GetTypeOk() (*ModelsDatabaseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HandlersDatabaseStatusResponse) SetType(v ModelsDatabaseType)`

SetType sets Type field to given value.

### HasType

`func (o *HandlersDatabaseStatusResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HandlersDatabaseStatusResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HandlersDatabaseStatusResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HandlersDatabaseStatusResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HandlersDatabaseStatusResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *HandlersDatabaseStatusResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HandlersDatabaseStatusResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HandlersDatabaseStatusResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HandlersDatabaseStatusResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


