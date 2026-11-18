# ModelsRunnerPackingHostPack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | Pointer to **string** | ClassName names the VM class this flavor serves (the class the catalog carries and the class jobs fit). | [optional] 
**CpuCores** | Pointer to **int32** |  | [optional] 
**DiskGb** | Pointer to **int32** |  | [optional] 
**MaxConcurrent** | Pointer to **int32** | MaxConcurrent is the flavor lid: live slices of the flavor on the host never exceed it. | [optional] 
**MinIdle** | Pointer to **int32** | MinIdle is the warm floor per flavor — a budget promise the host-budget arithmetic reserves. | [optional] 
**RamMb** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsRunnerPackingHostPack

`func NewModelsRunnerPackingHostPack() *ModelsRunnerPackingHostPack`

NewModelsRunnerPackingHostPack instantiates a new ModelsRunnerPackingHostPack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsRunnerPackingHostPackWithDefaults

`func NewModelsRunnerPackingHostPackWithDefaults() *ModelsRunnerPackingHostPack`

NewModelsRunnerPackingHostPackWithDefaults instantiates a new ModelsRunnerPackingHostPack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *ModelsRunnerPackingHostPack) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *ModelsRunnerPackingHostPack) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *ModelsRunnerPackingHostPack) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *ModelsRunnerPackingHostPack) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetCpuCores

`func (o *ModelsRunnerPackingHostPack) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *ModelsRunnerPackingHostPack) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *ModelsRunnerPackingHostPack) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *ModelsRunnerPackingHostPack) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetDiskGb

`func (o *ModelsRunnerPackingHostPack) GetDiskGb() int32`

GetDiskGb returns the DiskGb field if non-nil, zero value otherwise.

### GetDiskGbOk

`func (o *ModelsRunnerPackingHostPack) GetDiskGbOk() (*int32, bool)`

GetDiskGbOk returns a tuple with the DiskGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGb

`func (o *ModelsRunnerPackingHostPack) SetDiskGb(v int32)`

SetDiskGb sets DiskGb field to given value.

### HasDiskGb

`func (o *ModelsRunnerPackingHostPack) HasDiskGb() bool`

HasDiskGb returns a boolean if a field has been set.

### GetMaxConcurrent

`func (o *ModelsRunnerPackingHostPack) GetMaxConcurrent() int32`

GetMaxConcurrent returns the MaxConcurrent field if non-nil, zero value otherwise.

### GetMaxConcurrentOk

`func (o *ModelsRunnerPackingHostPack) GetMaxConcurrentOk() (*int32, bool)`

GetMaxConcurrentOk returns a tuple with the MaxConcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrent

`func (o *ModelsRunnerPackingHostPack) SetMaxConcurrent(v int32)`

SetMaxConcurrent sets MaxConcurrent field to given value.

### HasMaxConcurrent

`func (o *ModelsRunnerPackingHostPack) HasMaxConcurrent() bool`

HasMaxConcurrent returns a boolean if a field has been set.

### GetMinIdle

`func (o *ModelsRunnerPackingHostPack) GetMinIdle() int32`

GetMinIdle returns the MinIdle field if non-nil, zero value otherwise.

### GetMinIdleOk

`func (o *ModelsRunnerPackingHostPack) GetMinIdleOk() (*int32, bool)`

GetMinIdleOk returns a tuple with the MinIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIdle

`func (o *ModelsRunnerPackingHostPack) SetMinIdle(v int32)`

SetMinIdle sets MinIdle field to given value.

### HasMinIdle

`func (o *ModelsRunnerPackingHostPack) HasMinIdle() bool`

HasMinIdle returns a boolean if a field has been set.

### GetRamMb

`func (o *ModelsRunnerPackingHostPack) GetRamMb() int32`

GetRamMb returns the RamMb field if non-nil, zero value otherwise.

### GetRamMbOk

`func (o *ModelsRunnerPackingHostPack) GetRamMbOk() (*int32, bool)`

GetRamMbOk returns a tuple with the RamMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamMb

`func (o *ModelsRunnerPackingHostPack) SetRamMb(v int32)`

SetRamMb sets RamMb field to given value.

### HasRamMb

`func (o *ModelsRunnerPackingHostPack) HasRamMb() bool`

HasRamMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


