# ModelsRunnerPackingHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | Arch is the host architecture (the packs&#39; classes must match it). | [optional] 
**Base** | Pointer to **string** | Base is the GENERIC OS base (agent only — nothing runner-shaped in it, even on our own hosts). | [optional] 
**Cache** | Pointer to [**ModelsRunnerPackingHostCache**](ModelsRunnerPackingHostCache.md) | Cache is the pool spec&#39;s cache block; a disabled block means no cache volumes. | [optional] 
**CpuCores** | Pointer to **int32** | CPUCores and MemoryGB are the host&#39;s metal facts (MemoryGB is the DECLARED unit — the floor&#39;s projection converts GB to the MB the packing arithmetic uses). | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled gates the floor feed: a disabled declaration stays readable but feeds no floor. NO gorm default tag — a defaulted bool would omit a false value on struct Create and silently enable the declaration (the migration&#39;s SQL default covers raw inserts only). | [optional] 
**HostId** | Pointer to **string** | HostID names the real machine — globally unique (see the header). | [optional] 
**Id** | Pointer to **int32** | ID trails the composites (fieldalignment). | [optional] 
**Labels** | Pointer to **[]string** | Labels are the capability labels every slice on the host carries. | [optional] 
**MemoryGb** | Pointer to **int32** |  | [optional] 
**MinFreeGb** | Pointer to **int32** | MinFreeGB is the disk admission floor (§5.3 check 4). | [optional] 
**Name** | Pointer to **string** | Name is the org-facing label (what the pools page renders). | [optional] 
**OrgId** | Pointer to **int32** | OrgID is the declaring org — every read and write scopes by it. | [optional] 
**Packs** | Pointer to [**[]ModelsRunnerPackingHostPack**](ModelsRunnerPackingHostPack.md) | Packs are the per-flavor packs (budgets + the concurrency lid + the warm floor). | [optional] 
**RunnerForge** | Pointer to **string** | RunnerForge + RunnerVersionPin are the pool&#39;s download-at-enlist pin (the slice&#39;s runner downloads fresh at enlist, against this pin — never a pre-baked image). | [optional] 
**RunnerVersionPin** | Pointer to **string** |  | [optional] 
**StoragePool** | Pointer to **string** | StoragePool names the host storage pool slices draw from. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Virtualization** | Pointer to **string** | Virtualization selects the backend class: incus-vm | incus-container | docker. | [optional] 

## Methods

### NewModelsRunnerPackingHost

`func NewModelsRunnerPackingHost() *ModelsRunnerPackingHost`

NewModelsRunnerPackingHost instantiates a new ModelsRunnerPackingHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsRunnerPackingHostWithDefaults

`func NewModelsRunnerPackingHostWithDefaults() *ModelsRunnerPackingHost`

NewModelsRunnerPackingHostWithDefaults instantiates a new ModelsRunnerPackingHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *ModelsRunnerPackingHost) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ModelsRunnerPackingHost) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ModelsRunnerPackingHost) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ModelsRunnerPackingHost) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetBase

`func (o *ModelsRunnerPackingHost) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *ModelsRunnerPackingHost) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *ModelsRunnerPackingHost) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *ModelsRunnerPackingHost) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetCache

`func (o *ModelsRunnerPackingHost) GetCache() ModelsRunnerPackingHostCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *ModelsRunnerPackingHost) GetCacheOk() (*ModelsRunnerPackingHostCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *ModelsRunnerPackingHost) SetCache(v ModelsRunnerPackingHostCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *ModelsRunnerPackingHost) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCpuCores

`func (o *ModelsRunnerPackingHost) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *ModelsRunnerPackingHost) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *ModelsRunnerPackingHost) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *ModelsRunnerPackingHost) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsRunnerPackingHost) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsRunnerPackingHost) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsRunnerPackingHost) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsRunnerPackingHost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelsRunnerPackingHost) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelsRunnerPackingHost) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelsRunnerPackingHost) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelsRunnerPackingHost) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHostId

`func (o *ModelsRunnerPackingHost) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ModelsRunnerPackingHost) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ModelsRunnerPackingHost) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ModelsRunnerPackingHost) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetId

`func (o *ModelsRunnerPackingHost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsRunnerPackingHost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsRunnerPackingHost) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsRunnerPackingHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *ModelsRunnerPackingHost) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ModelsRunnerPackingHost) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ModelsRunnerPackingHost) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ModelsRunnerPackingHost) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMemoryGb

`func (o *ModelsRunnerPackingHost) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *ModelsRunnerPackingHost) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *ModelsRunnerPackingHost) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.

### HasMemoryGb

`func (o *ModelsRunnerPackingHost) HasMemoryGb() bool`

HasMemoryGb returns a boolean if a field has been set.

### GetMinFreeGb

`func (o *ModelsRunnerPackingHost) GetMinFreeGb() int32`

GetMinFreeGb returns the MinFreeGb field if non-nil, zero value otherwise.

### GetMinFreeGbOk

`func (o *ModelsRunnerPackingHost) GetMinFreeGbOk() (*int32, bool)`

GetMinFreeGbOk returns a tuple with the MinFreeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFreeGb

`func (o *ModelsRunnerPackingHost) SetMinFreeGb(v int32)`

SetMinFreeGb sets MinFreeGb field to given value.

### HasMinFreeGb

`func (o *ModelsRunnerPackingHost) HasMinFreeGb() bool`

HasMinFreeGb returns a boolean if a field has been set.

### GetName

`func (o *ModelsRunnerPackingHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsRunnerPackingHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsRunnerPackingHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsRunnerPackingHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsRunnerPackingHost) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsRunnerPackingHost) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsRunnerPackingHost) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsRunnerPackingHost) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPacks

`func (o *ModelsRunnerPackingHost) GetPacks() []ModelsRunnerPackingHostPack`

GetPacks returns the Packs field if non-nil, zero value otherwise.

### GetPacksOk

`func (o *ModelsRunnerPackingHost) GetPacksOk() (*[]ModelsRunnerPackingHostPack, bool)`

GetPacksOk returns a tuple with the Packs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacks

`func (o *ModelsRunnerPackingHost) SetPacks(v []ModelsRunnerPackingHostPack)`

SetPacks sets Packs field to given value.

### HasPacks

`func (o *ModelsRunnerPackingHost) HasPacks() bool`

HasPacks returns a boolean if a field has been set.

### GetRunnerForge

`func (o *ModelsRunnerPackingHost) GetRunnerForge() string`

GetRunnerForge returns the RunnerForge field if non-nil, zero value otherwise.

### GetRunnerForgeOk

`func (o *ModelsRunnerPackingHost) GetRunnerForgeOk() (*string, bool)`

GetRunnerForgeOk returns a tuple with the RunnerForge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerForge

`func (o *ModelsRunnerPackingHost) SetRunnerForge(v string)`

SetRunnerForge sets RunnerForge field to given value.

### HasRunnerForge

`func (o *ModelsRunnerPackingHost) HasRunnerForge() bool`

HasRunnerForge returns a boolean if a field has been set.

### GetRunnerVersionPin

`func (o *ModelsRunnerPackingHost) GetRunnerVersionPin() string`

GetRunnerVersionPin returns the RunnerVersionPin field if non-nil, zero value otherwise.

### GetRunnerVersionPinOk

`func (o *ModelsRunnerPackingHost) GetRunnerVersionPinOk() (*string, bool)`

GetRunnerVersionPinOk returns a tuple with the RunnerVersionPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerVersionPin

`func (o *ModelsRunnerPackingHost) SetRunnerVersionPin(v string)`

SetRunnerVersionPin sets RunnerVersionPin field to given value.

### HasRunnerVersionPin

`func (o *ModelsRunnerPackingHost) HasRunnerVersionPin() bool`

HasRunnerVersionPin returns a boolean if a field has been set.

### GetStoragePool

`func (o *ModelsRunnerPackingHost) GetStoragePool() string`

GetStoragePool returns the StoragePool field if non-nil, zero value otherwise.

### GetStoragePoolOk

`func (o *ModelsRunnerPackingHost) GetStoragePoolOk() (*string, bool)`

GetStoragePoolOk returns a tuple with the StoragePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePool

`func (o *ModelsRunnerPackingHost) SetStoragePool(v string)`

SetStoragePool sets StoragePool field to given value.

### HasStoragePool

`func (o *ModelsRunnerPackingHost) HasStoragePool() bool`

HasStoragePool returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsRunnerPackingHost) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsRunnerPackingHost) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsRunnerPackingHost) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsRunnerPackingHost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVirtualization

`func (o *ModelsRunnerPackingHost) GetVirtualization() string`

GetVirtualization returns the Virtualization field if non-nil, zero value otherwise.

### GetVirtualizationOk

`func (o *ModelsRunnerPackingHost) GetVirtualizationOk() (*string, bool)`

GetVirtualizationOk returns a tuple with the Virtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualization

`func (o *ModelsRunnerPackingHost) SetVirtualization(v string)`

SetVirtualization sets Virtualization field to given value.

### HasVirtualization

`func (o *ModelsRunnerPackingHost) HasVirtualization() bool`

HasVirtualization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


