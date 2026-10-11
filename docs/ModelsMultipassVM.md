# ModelsMultipassVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** | Agent managing this VM | [optional] 
**Cpus** | Pointer to **int32** | Resource allocation (optional, from agent sync) | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to [**GormDeletedAt**](GormDeletedAt.md) |  | [optional] 
**Disk** | Pointer to **string** |  | [optional] 
**HostVm** | Pointer to [**ModelsVirtualMachine**](ModelsVirtualMachine.md) |  | [optional] 
**HostVmId** | Pointer to **int32** | Host VM relationship | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Ipv4** | Pointer to **string** |  | [optional] 
**LastSeenAt** | Pointer to **string** | Last seen tracks when the agent last reported this VM&#39;s state | [optional] 
**Memory** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | VM identity (from Multipass CLI) | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsMultipassVM

`func NewModelsMultipassVM() *ModelsMultipassVM`

NewModelsMultipassVM instantiates a new ModelsMultipassVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsMultipassVMWithDefaults

`func NewModelsMultipassVMWithDefaults() *ModelsMultipassVM`

NewModelsMultipassVMWithDefaults instantiates a new ModelsMultipassVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *ModelsMultipassVM) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ModelsMultipassVM) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ModelsMultipassVM) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ModelsMultipassVM) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetCpus

`func (o *ModelsMultipassVM) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *ModelsMultipassVM) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *ModelsMultipassVM) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *ModelsMultipassVM) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsMultipassVM) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsMultipassVM) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsMultipassVM) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsMultipassVM) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsMultipassVM) GetDeletedAt() GormDeletedAt`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsMultipassVM) GetDeletedAtOk() (*GormDeletedAt, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsMultipassVM) SetDeletedAt(v GormDeletedAt)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsMultipassVM) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDisk

`func (o *ModelsMultipassVM) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ModelsMultipassVM) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ModelsMultipassVM) SetDisk(v string)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ModelsMultipassVM) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetHostVm

`func (o *ModelsMultipassVM) GetHostVm() ModelsVirtualMachine`

GetHostVm returns the HostVm field if non-nil, zero value otherwise.

### GetHostVmOk

`func (o *ModelsMultipassVM) GetHostVmOk() (*ModelsVirtualMachine, bool)`

GetHostVmOk returns a tuple with the HostVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVm

`func (o *ModelsMultipassVM) SetHostVm(v ModelsVirtualMachine)`

SetHostVm sets HostVm field to given value.

### HasHostVm

`func (o *ModelsMultipassVM) HasHostVm() bool`

HasHostVm returns a boolean if a field has been set.

### GetHostVmId

`func (o *ModelsMultipassVM) GetHostVmId() int32`

GetHostVmId returns the HostVmId field if non-nil, zero value otherwise.

### GetHostVmIdOk

`func (o *ModelsMultipassVM) GetHostVmIdOk() (*int32, bool)`

GetHostVmIdOk returns a tuple with the HostVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVmId

`func (o *ModelsMultipassVM) SetHostVmId(v int32)`

SetHostVmId sets HostVmId field to given value.

### HasHostVmId

`func (o *ModelsMultipassVM) HasHostVmId() bool`

HasHostVmId returns a boolean if a field has been set.

### GetId

`func (o *ModelsMultipassVM) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsMultipassVM) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsMultipassVM) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsMultipassVM) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ModelsMultipassVM) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ModelsMultipassVM) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ModelsMultipassVM) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ModelsMultipassVM) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIpv4

`func (o *ModelsMultipassVM) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *ModelsMultipassVM) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *ModelsMultipassVM) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *ModelsMultipassVM) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *ModelsMultipassVM) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ModelsMultipassVM) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ModelsMultipassVM) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *ModelsMultipassVM) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetMemory

`func (o *ModelsMultipassVM) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ModelsMultipassVM) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ModelsMultipassVM) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ModelsMultipassVM) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetName

`func (o *ModelsMultipassVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsMultipassVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsMultipassVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsMultipassVM) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *ModelsMultipassVM) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelsMultipassVM) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelsMultipassVM) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ModelsMultipassVM) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsMultipassVM) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsMultipassVM) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsMultipassVM) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsMultipassVM) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUuid

`func (o *ModelsMultipassVM) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ModelsMultipassVM) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ModelsMultipassVM) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ModelsMultipassVM) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


