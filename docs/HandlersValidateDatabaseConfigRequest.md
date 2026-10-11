# HandlersValidateDatabaseConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupBucketName** | Pointer to **string** |  | [optional] 
**BackupEnabled** | Pointer to **bool** |  | [optional] 
**BackupSchedule** | Pointer to **string** |  | [optional] 
**ClusterId** | Pointer to **int32** |  | [optional] 
**CpuLimit** | Pointer to **string** |  | [optional] 
**CpuRequest** | Pointer to **string** |  | [optional] 
**HaEnabled** | Pointer to **bool** |  | [optional] 
**MemoryLimit** | Pointer to **string** |  | [optional] 
**MemoryRequest** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**ReplicaCount** | Pointer to **int32** |  | [optional] 
**S3ConnectionId** | Pointer to **int32** |  | [optional] 
**StorageClass** | Pointer to **string** |  | [optional] 
**StorageSize** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersValidateDatabaseConfigRequest

`func NewHandlersValidateDatabaseConfigRequest() *HandlersValidateDatabaseConfigRequest`

NewHandlersValidateDatabaseConfigRequest instantiates a new HandlersValidateDatabaseConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersValidateDatabaseConfigRequestWithDefaults

`func NewHandlersValidateDatabaseConfigRequestWithDefaults() *HandlersValidateDatabaseConfigRequest`

NewHandlersValidateDatabaseConfigRequestWithDefaults instantiates a new HandlersValidateDatabaseConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupBucketName

`func (o *HandlersValidateDatabaseConfigRequest) GetBackupBucketName() string`

GetBackupBucketName returns the BackupBucketName field if non-nil, zero value otherwise.

### GetBackupBucketNameOk

`func (o *HandlersValidateDatabaseConfigRequest) GetBackupBucketNameOk() (*string, bool)`

GetBackupBucketNameOk returns a tuple with the BackupBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupBucketName

`func (o *HandlersValidateDatabaseConfigRequest) SetBackupBucketName(v string)`

SetBackupBucketName sets BackupBucketName field to given value.

### HasBackupBucketName

`func (o *HandlersValidateDatabaseConfigRequest) HasBackupBucketName() bool`

HasBackupBucketName returns a boolean if a field has been set.

### GetBackupEnabled

`func (o *HandlersValidateDatabaseConfigRequest) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *HandlersValidateDatabaseConfigRequest) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *HandlersValidateDatabaseConfigRequest) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *HandlersValidateDatabaseConfigRequest) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.

### GetBackupSchedule

`func (o *HandlersValidateDatabaseConfigRequest) GetBackupSchedule() string`

GetBackupSchedule returns the BackupSchedule field if non-nil, zero value otherwise.

### GetBackupScheduleOk

`func (o *HandlersValidateDatabaseConfigRequest) GetBackupScheduleOk() (*string, bool)`

GetBackupScheduleOk returns a tuple with the BackupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedule

`func (o *HandlersValidateDatabaseConfigRequest) SetBackupSchedule(v string)`

SetBackupSchedule sets BackupSchedule field to given value.

### HasBackupSchedule

`func (o *HandlersValidateDatabaseConfigRequest) HasBackupSchedule() bool`

HasBackupSchedule returns a boolean if a field has been set.

### GetClusterId

`func (o *HandlersValidateDatabaseConfigRequest) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HandlersValidateDatabaseConfigRequest) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HandlersValidateDatabaseConfigRequest) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HandlersValidateDatabaseConfigRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCpuLimit

`func (o *HandlersValidateDatabaseConfigRequest) GetCpuLimit() string`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *HandlersValidateDatabaseConfigRequest) GetCpuLimitOk() (*string, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *HandlersValidateDatabaseConfigRequest) SetCpuLimit(v string)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *HandlersValidateDatabaseConfigRequest) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetCpuRequest

`func (o *HandlersValidateDatabaseConfigRequest) GetCpuRequest() string`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *HandlersValidateDatabaseConfigRequest) GetCpuRequestOk() (*string, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *HandlersValidateDatabaseConfigRequest) SetCpuRequest(v string)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *HandlersValidateDatabaseConfigRequest) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### GetHaEnabled

`func (o *HandlersValidateDatabaseConfigRequest) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *HandlersValidateDatabaseConfigRequest) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *HandlersValidateDatabaseConfigRequest) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *HandlersValidateDatabaseConfigRequest) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *HandlersValidateDatabaseConfigRequest) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *HandlersValidateDatabaseConfigRequest) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *HandlersValidateDatabaseConfigRequest) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *HandlersValidateDatabaseConfigRequest) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMemoryRequest

`func (o *HandlersValidateDatabaseConfigRequest) GetMemoryRequest() string`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *HandlersValidateDatabaseConfigRequest) GetMemoryRequestOk() (*string, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *HandlersValidateDatabaseConfigRequest) SetMemoryRequest(v string)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *HandlersValidateDatabaseConfigRequest) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.

### GetName

`func (o *HandlersValidateDatabaseConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersValidateDatabaseConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersValidateDatabaseConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersValidateDatabaseConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *HandlersValidateDatabaseConfigRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HandlersValidateDatabaseConfigRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HandlersValidateDatabaseConfigRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HandlersValidateDatabaseConfigRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetReplicaCount

`func (o *HandlersValidateDatabaseConfigRequest) GetReplicaCount() int32`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *HandlersValidateDatabaseConfigRequest) GetReplicaCountOk() (*int32, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *HandlersValidateDatabaseConfigRequest) SetReplicaCount(v int32)`

SetReplicaCount sets ReplicaCount field to given value.

### HasReplicaCount

`func (o *HandlersValidateDatabaseConfigRequest) HasReplicaCount() bool`

HasReplicaCount returns a boolean if a field has been set.

### GetS3ConnectionId

`func (o *HandlersValidateDatabaseConfigRequest) GetS3ConnectionId() int32`

GetS3ConnectionId returns the S3ConnectionId field if non-nil, zero value otherwise.

### GetS3ConnectionIdOk

`func (o *HandlersValidateDatabaseConfigRequest) GetS3ConnectionIdOk() (*int32, bool)`

GetS3ConnectionIdOk returns a tuple with the S3ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ConnectionId

`func (o *HandlersValidateDatabaseConfigRequest) SetS3ConnectionId(v int32)`

SetS3ConnectionId sets S3ConnectionId field to given value.

### HasS3ConnectionId

`func (o *HandlersValidateDatabaseConfigRequest) HasS3ConnectionId() bool`

HasS3ConnectionId returns a boolean if a field has been set.

### GetStorageClass

`func (o *HandlersValidateDatabaseConfigRequest) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *HandlersValidateDatabaseConfigRequest) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *HandlersValidateDatabaseConfigRequest) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *HandlersValidateDatabaseConfigRequest) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetStorageSize

`func (o *HandlersValidateDatabaseConfigRequest) GetStorageSize() string`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *HandlersValidateDatabaseConfigRequest) GetStorageSizeOk() (*string, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *HandlersValidateDatabaseConfigRequest) SetStorageSize(v string)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *HandlersValidateDatabaseConfigRequest) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetType

`func (o *HandlersValidateDatabaseConfigRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HandlersValidateDatabaseConfigRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HandlersValidateDatabaseConfigRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HandlersValidateDatabaseConfigRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *HandlersValidateDatabaseConfigRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HandlersValidateDatabaseConfigRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HandlersValidateDatabaseConfigRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HandlersValidateDatabaseConfigRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


