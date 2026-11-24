# StorageusageVolumeUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityBytes** | Pointer to **int32** |  | [optional] 
**ClusterId** | Pointer to **int32** |  | [optional] 
**InstallName** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**PodName** | Pointer to **string** |  | [optional] 
**Projection** | Pointer to [**StorageusageVolumeProjectionView**](StorageusageVolumeProjectionView.md) |  | [optional] 
**PvcName** | Pointer to **string** |  | [optional] 
**SampledAt** | Pointer to **string** |  | [optional] 
**UsedBytes** | Pointer to **int32** |  | [optional] 
**WorkloadKind** | Pointer to **string** |  | [optional] 
**WorkloadName** | Pointer to **string** |  | [optional] 

## Methods

### NewStorageusageVolumeUsage

`func NewStorageusageVolumeUsage() *StorageusageVolumeUsage`

NewStorageusageVolumeUsage instantiates a new StorageusageVolumeUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageusageVolumeUsageWithDefaults

`func NewStorageusageVolumeUsageWithDefaults() *StorageusageVolumeUsage`

NewStorageusageVolumeUsageWithDefaults instantiates a new StorageusageVolumeUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityBytes

`func (o *StorageusageVolumeUsage) GetCapacityBytes() int32`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *StorageusageVolumeUsage) GetCapacityBytesOk() (*int32, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *StorageusageVolumeUsage) SetCapacityBytes(v int32)`

SetCapacityBytes sets CapacityBytes field to given value.

### HasCapacityBytes

`func (o *StorageusageVolumeUsage) HasCapacityBytes() bool`

HasCapacityBytes returns a boolean if a field has been set.

### GetClusterId

`func (o *StorageusageVolumeUsage) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StorageusageVolumeUsage) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StorageusageVolumeUsage) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *StorageusageVolumeUsage) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetInstallName

`func (o *StorageusageVolumeUsage) GetInstallName() string`

GetInstallName returns the InstallName field if non-nil, zero value otherwise.

### GetInstallNameOk

`func (o *StorageusageVolumeUsage) GetInstallNameOk() (*string, bool)`

GetInstallNameOk returns a tuple with the InstallName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallName

`func (o *StorageusageVolumeUsage) SetInstallName(v string)`

SetInstallName sets InstallName field to given value.

### HasInstallName

`func (o *StorageusageVolumeUsage) HasInstallName() bool`

HasInstallName returns a boolean if a field has been set.

### GetNamespace

`func (o *StorageusageVolumeUsage) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *StorageusageVolumeUsage) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *StorageusageVolumeUsage) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *StorageusageVolumeUsage) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPodName

`func (o *StorageusageVolumeUsage) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *StorageusageVolumeUsage) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *StorageusageVolumeUsage) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *StorageusageVolumeUsage) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetProjection

`func (o *StorageusageVolumeUsage) GetProjection() StorageusageVolumeProjectionView`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *StorageusageVolumeUsage) GetProjectionOk() (*StorageusageVolumeProjectionView, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *StorageusageVolumeUsage) SetProjection(v StorageusageVolumeProjectionView)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *StorageusageVolumeUsage) HasProjection() bool`

HasProjection returns a boolean if a field has been set.

### GetPvcName

`func (o *StorageusageVolumeUsage) GetPvcName() string`

GetPvcName returns the PvcName field if non-nil, zero value otherwise.

### GetPvcNameOk

`func (o *StorageusageVolumeUsage) GetPvcNameOk() (*string, bool)`

GetPvcNameOk returns a tuple with the PvcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcName

`func (o *StorageusageVolumeUsage) SetPvcName(v string)`

SetPvcName sets PvcName field to given value.

### HasPvcName

`func (o *StorageusageVolumeUsage) HasPvcName() bool`

HasPvcName returns a boolean if a field has been set.

### GetSampledAt

`func (o *StorageusageVolumeUsage) GetSampledAt() string`

GetSampledAt returns the SampledAt field if non-nil, zero value otherwise.

### GetSampledAtOk

`func (o *StorageusageVolumeUsage) GetSampledAtOk() (*string, bool)`

GetSampledAtOk returns a tuple with the SampledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampledAt

`func (o *StorageusageVolumeUsage) SetSampledAt(v string)`

SetSampledAt sets SampledAt field to given value.

### HasSampledAt

`func (o *StorageusageVolumeUsage) HasSampledAt() bool`

HasSampledAt returns a boolean if a field has been set.

### GetUsedBytes

`func (o *StorageusageVolumeUsage) GetUsedBytes() int32`

GetUsedBytes returns the UsedBytes field if non-nil, zero value otherwise.

### GetUsedBytesOk

`func (o *StorageusageVolumeUsage) GetUsedBytesOk() (*int32, bool)`

GetUsedBytesOk returns a tuple with the UsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBytes

`func (o *StorageusageVolumeUsage) SetUsedBytes(v int32)`

SetUsedBytes sets UsedBytes field to given value.

### HasUsedBytes

`func (o *StorageusageVolumeUsage) HasUsedBytes() bool`

HasUsedBytes returns a boolean if a field has been set.

### GetWorkloadKind

`func (o *StorageusageVolumeUsage) GetWorkloadKind() string`

GetWorkloadKind returns the WorkloadKind field if non-nil, zero value otherwise.

### GetWorkloadKindOk

`func (o *StorageusageVolumeUsage) GetWorkloadKindOk() (*string, bool)`

GetWorkloadKindOk returns a tuple with the WorkloadKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadKind

`func (o *StorageusageVolumeUsage) SetWorkloadKind(v string)`

SetWorkloadKind sets WorkloadKind field to given value.

### HasWorkloadKind

`func (o *StorageusageVolumeUsage) HasWorkloadKind() bool`

HasWorkloadKind returns a boolean if a field has been set.

### GetWorkloadName

`func (o *StorageusageVolumeUsage) GetWorkloadName() string`

GetWorkloadName returns the WorkloadName field if non-nil, zero value otherwise.

### GetWorkloadNameOk

`func (o *StorageusageVolumeUsage) GetWorkloadNameOk() (*string, bool)`

GetWorkloadNameOk returns a tuple with the WorkloadName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadName

`func (o *StorageusageVolumeUsage) SetWorkloadName(v string)`

SetWorkloadName sets WorkloadName field to given value.

### HasWorkloadName

`func (o *StorageusageVolumeUsage) HasWorkloadName() bool`

HasWorkloadName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


