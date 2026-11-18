# HandlersRunnerPackingHostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**Base** | Pointer to **string** |  | [optional] 
**Cache** | Pointer to [**ModelsRunnerPackingHostCache**](ModelsRunnerPackingHostCache.md) |  | [optional] 
**CpuCores** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HostId** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**MemoryGb** | Pointer to **int32** |  | [optional] 
**MinFreeGb** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Packs** | Pointer to [**[]ModelsRunnerPackingHostPack**](ModelsRunnerPackingHostPack.md) |  | [optional] 
**RunnerForge** | Pointer to **string** |  | [optional] 
**RunnerVersionPin** | Pointer to **string** |  | [optional] 
**StoragePool** | Pointer to **string** |  | [optional] 
**Virtualization** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerPackingHostBody

`func NewHandlersRunnerPackingHostBody() *HandlersRunnerPackingHostBody`

NewHandlersRunnerPackingHostBody instantiates a new HandlersRunnerPackingHostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerPackingHostBodyWithDefaults

`func NewHandlersRunnerPackingHostBodyWithDefaults() *HandlersRunnerPackingHostBody`

NewHandlersRunnerPackingHostBodyWithDefaults instantiates a new HandlersRunnerPackingHostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *HandlersRunnerPackingHostBody) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *HandlersRunnerPackingHostBody) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *HandlersRunnerPackingHostBody) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *HandlersRunnerPackingHostBody) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetBase

`func (o *HandlersRunnerPackingHostBody) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *HandlersRunnerPackingHostBody) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *HandlersRunnerPackingHostBody) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *HandlersRunnerPackingHostBody) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetCache

`func (o *HandlersRunnerPackingHostBody) GetCache() ModelsRunnerPackingHostCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *HandlersRunnerPackingHostBody) GetCacheOk() (*ModelsRunnerPackingHostCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *HandlersRunnerPackingHostBody) SetCache(v ModelsRunnerPackingHostCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *HandlersRunnerPackingHostBody) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCpuCores

`func (o *HandlersRunnerPackingHostBody) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *HandlersRunnerPackingHostBody) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *HandlersRunnerPackingHostBody) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *HandlersRunnerPackingHostBody) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetEnabled

`func (o *HandlersRunnerPackingHostBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HandlersRunnerPackingHostBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HandlersRunnerPackingHostBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HandlersRunnerPackingHostBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHostId

`func (o *HandlersRunnerPackingHostBody) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *HandlersRunnerPackingHostBody) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *HandlersRunnerPackingHostBody) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *HandlersRunnerPackingHostBody) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetLabels

`func (o *HandlersRunnerPackingHostBody) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HandlersRunnerPackingHostBody) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HandlersRunnerPackingHostBody) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HandlersRunnerPackingHostBody) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMemoryGb

`func (o *HandlersRunnerPackingHostBody) GetMemoryGb() int32`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *HandlersRunnerPackingHostBody) GetMemoryGbOk() (*int32, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *HandlersRunnerPackingHostBody) SetMemoryGb(v int32)`

SetMemoryGb sets MemoryGb field to given value.

### HasMemoryGb

`func (o *HandlersRunnerPackingHostBody) HasMemoryGb() bool`

HasMemoryGb returns a boolean if a field has been set.

### GetMinFreeGb

`func (o *HandlersRunnerPackingHostBody) GetMinFreeGb() int32`

GetMinFreeGb returns the MinFreeGb field if non-nil, zero value otherwise.

### GetMinFreeGbOk

`func (o *HandlersRunnerPackingHostBody) GetMinFreeGbOk() (*int32, bool)`

GetMinFreeGbOk returns a tuple with the MinFreeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFreeGb

`func (o *HandlersRunnerPackingHostBody) SetMinFreeGb(v int32)`

SetMinFreeGb sets MinFreeGb field to given value.

### HasMinFreeGb

`func (o *HandlersRunnerPackingHostBody) HasMinFreeGb() bool`

HasMinFreeGb returns a boolean if a field has been set.

### GetName

`func (o *HandlersRunnerPackingHostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersRunnerPackingHostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersRunnerPackingHostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersRunnerPackingHostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPacks

`func (o *HandlersRunnerPackingHostBody) GetPacks() []ModelsRunnerPackingHostPack`

GetPacks returns the Packs field if non-nil, zero value otherwise.

### GetPacksOk

`func (o *HandlersRunnerPackingHostBody) GetPacksOk() (*[]ModelsRunnerPackingHostPack, bool)`

GetPacksOk returns a tuple with the Packs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacks

`func (o *HandlersRunnerPackingHostBody) SetPacks(v []ModelsRunnerPackingHostPack)`

SetPacks sets Packs field to given value.

### HasPacks

`func (o *HandlersRunnerPackingHostBody) HasPacks() bool`

HasPacks returns a boolean if a field has been set.

### GetRunnerForge

`func (o *HandlersRunnerPackingHostBody) GetRunnerForge() string`

GetRunnerForge returns the RunnerForge field if non-nil, zero value otherwise.

### GetRunnerForgeOk

`func (o *HandlersRunnerPackingHostBody) GetRunnerForgeOk() (*string, bool)`

GetRunnerForgeOk returns a tuple with the RunnerForge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerForge

`func (o *HandlersRunnerPackingHostBody) SetRunnerForge(v string)`

SetRunnerForge sets RunnerForge field to given value.

### HasRunnerForge

`func (o *HandlersRunnerPackingHostBody) HasRunnerForge() bool`

HasRunnerForge returns a boolean if a field has been set.

### GetRunnerVersionPin

`func (o *HandlersRunnerPackingHostBody) GetRunnerVersionPin() string`

GetRunnerVersionPin returns the RunnerVersionPin field if non-nil, zero value otherwise.

### GetRunnerVersionPinOk

`func (o *HandlersRunnerPackingHostBody) GetRunnerVersionPinOk() (*string, bool)`

GetRunnerVersionPinOk returns a tuple with the RunnerVersionPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerVersionPin

`func (o *HandlersRunnerPackingHostBody) SetRunnerVersionPin(v string)`

SetRunnerVersionPin sets RunnerVersionPin field to given value.

### HasRunnerVersionPin

`func (o *HandlersRunnerPackingHostBody) HasRunnerVersionPin() bool`

HasRunnerVersionPin returns a boolean if a field has been set.

### GetStoragePool

`func (o *HandlersRunnerPackingHostBody) GetStoragePool() string`

GetStoragePool returns the StoragePool field if non-nil, zero value otherwise.

### GetStoragePoolOk

`func (o *HandlersRunnerPackingHostBody) GetStoragePoolOk() (*string, bool)`

GetStoragePoolOk returns a tuple with the StoragePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePool

`func (o *HandlersRunnerPackingHostBody) SetStoragePool(v string)`

SetStoragePool sets StoragePool field to given value.

### HasStoragePool

`func (o *HandlersRunnerPackingHostBody) HasStoragePool() bool`

HasStoragePool returns a boolean if a field has been set.

### GetVirtualization

`func (o *HandlersRunnerPackingHostBody) GetVirtualization() string`

GetVirtualization returns the Virtualization field if non-nil, zero value otherwise.

### GetVirtualizationOk

`func (o *HandlersRunnerPackingHostBody) GetVirtualizationOk() (*string, bool)`

GetVirtualizationOk returns a tuple with the Virtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualization

`func (o *HandlersRunnerPackingHostBody) SetVirtualization(v string)`

SetVirtualization sets Virtualization field to given value.

### HasVirtualization

`func (o *HandlersRunnerPackingHostBody) HasVirtualization() bool`

HasVirtualization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


