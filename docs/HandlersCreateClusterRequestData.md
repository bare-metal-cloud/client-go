# HandlersCreateClusterRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdoptExisting** | Pointer to **bool** | AdoptExisting is the EXPLICIT ADOPT opt-in (story 124.3, constraint 4): accept pre-existing, RUNNING, same-version k3s on the members instead of installing. Without it, a member with running k3s is 409 existing_k3s_running. The server verifies EVERY member via the contamination preflight before honoring it (verified adoption records cluster.adopted&#x3D;true; an unverified opt-in is not honored). | [optional] 
**AgentDeployMethod** | Pointer to **string** | AgentDeployMethod selects how the k8s-agent is deployed inside the cluster after provisioning: \&quot;kubectl\&quot; (default, applies manifest via client-go) or \&quot;helm\&quot; (installs via Helm SDK from the chart repo). | [optional] 
**FailurePolicy** | Pointer to **string** | FailurePolicy is the E5-7 LOCKED failure policy for multi-master creates (story 125.3): \&quot;system_managed\&quot; (per-master retries to the deadline, then the quorum-preconditioned auto-return ladder) or \&quot;user_managed\&quot; (hold + action_required on first failure, zero autonomous actions). Locked for the operation; persisted with the operation record (job payload + the D6 hold) — never a cluster column. Multi-master only; single-master ignores it. | [optional] 
**Ha** | Pointer to **bool** | High availability mode (cloud providers) | [optional] 
**InstallConfig** | Pointer to **map[string]string** | InstallConfig overrides distribution install defaults. Keys recognized by the distribution (e.g. \&quot;snapshotter\&quot;) set the corresponding flag; unknown keys are passed through as --key&#x3D;value. Empty &#x3D; use distribution defaults. | [optional] 
**MasterCount** | Pointer to **int32** | Number of master nodes (cloud providers) | [optional] 
**NetworkId** | Pointer to **int32** | NetworkID is the explicit network contract (story 124.17, D12): the network every member VM must belong to, selected in wizard step 1. Nullable — null &#x3D; the acknowledged \&quot;No network — advanced\&quot; choice (every member must be networkless). Validated server-side against every member; violations surface as structured 409s (vm_not_in_network / vm_in_cluster). | [optional] 
**Stack** | **string** |  | 
**Version** | **string** |  | 
**VmIds** | Pointer to **[]int32** | VMIDs is the array of VM IDs for BYOV clusters. The first VM is assigned the master role; all others are workers. For explicit per-VM role assignment, use VMRoles instead. | [optional] 
**VmRoles** | Pointer to [**[]HandlersVMRole**](HandlersVMRole.md) | VMRoles allows explicit per-VM role assignment. When provided, it overrides the implicit first-VM-is-master convention. Each entry maps a VM ID to its role (\&quot;master\&quot; or \&quot;worker\&quot;). If omitted, VMIDs is used with the first element as master. | [optional] 
**WorkerCount** | Pointer to **int32** | Number of worker nodes (cloud providers) | [optional] 

## Methods

### NewHandlersCreateClusterRequestData

`func NewHandlersCreateClusterRequestData(stack string, version string, ) *HandlersCreateClusterRequestData`

NewHandlersCreateClusterRequestData instantiates a new HandlersCreateClusterRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateClusterRequestDataWithDefaults

`func NewHandlersCreateClusterRequestDataWithDefaults() *HandlersCreateClusterRequestData`

NewHandlersCreateClusterRequestDataWithDefaults instantiates a new HandlersCreateClusterRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdoptExisting

`func (o *HandlersCreateClusterRequestData) GetAdoptExisting() bool`

GetAdoptExisting returns the AdoptExisting field if non-nil, zero value otherwise.

### GetAdoptExistingOk

`func (o *HandlersCreateClusterRequestData) GetAdoptExistingOk() (*bool, bool)`

GetAdoptExistingOk returns a tuple with the AdoptExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptExisting

`func (o *HandlersCreateClusterRequestData) SetAdoptExisting(v bool)`

SetAdoptExisting sets AdoptExisting field to given value.

### HasAdoptExisting

`func (o *HandlersCreateClusterRequestData) HasAdoptExisting() bool`

HasAdoptExisting returns a boolean if a field has been set.

### GetAgentDeployMethod

`func (o *HandlersCreateClusterRequestData) GetAgentDeployMethod() string`

GetAgentDeployMethod returns the AgentDeployMethod field if non-nil, zero value otherwise.

### GetAgentDeployMethodOk

`func (o *HandlersCreateClusterRequestData) GetAgentDeployMethodOk() (*string, bool)`

GetAgentDeployMethodOk returns a tuple with the AgentDeployMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentDeployMethod

`func (o *HandlersCreateClusterRequestData) SetAgentDeployMethod(v string)`

SetAgentDeployMethod sets AgentDeployMethod field to given value.

### HasAgentDeployMethod

`func (o *HandlersCreateClusterRequestData) HasAgentDeployMethod() bool`

HasAgentDeployMethod returns a boolean if a field has been set.

### GetFailurePolicy

`func (o *HandlersCreateClusterRequestData) GetFailurePolicy() string`

GetFailurePolicy returns the FailurePolicy field if non-nil, zero value otherwise.

### GetFailurePolicyOk

`func (o *HandlersCreateClusterRequestData) GetFailurePolicyOk() (*string, bool)`

GetFailurePolicyOk returns a tuple with the FailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailurePolicy

`func (o *HandlersCreateClusterRequestData) SetFailurePolicy(v string)`

SetFailurePolicy sets FailurePolicy field to given value.

### HasFailurePolicy

`func (o *HandlersCreateClusterRequestData) HasFailurePolicy() bool`

HasFailurePolicy returns a boolean if a field has been set.

### GetHa

`func (o *HandlersCreateClusterRequestData) GetHa() bool`

GetHa returns the Ha field if non-nil, zero value otherwise.

### GetHaOk

`func (o *HandlersCreateClusterRequestData) GetHaOk() (*bool, bool)`

GetHaOk returns a tuple with the Ha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHa

`func (o *HandlersCreateClusterRequestData) SetHa(v bool)`

SetHa sets Ha field to given value.

### HasHa

`func (o *HandlersCreateClusterRequestData) HasHa() bool`

HasHa returns a boolean if a field has been set.

### GetInstallConfig

`func (o *HandlersCreateClusterRequestData) GetInstallConfig() map[string]string`

GetInstallConfig returns the InstallConfig field if non-nil, zero value otherwise.

### GetInstallConfigOk

`func (o *HandlersCreateClusterRequestData) GetInstallConfigOk() (*map[string]string, bool)`

GetInstallConfigOk returns a tuple with the InstallConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallConfig

`func (o *HandlersCreateClusterRequestData) SetInstallConfig(v map[string]string)`

SetInstallConfig sets InstallConfig field to given value.

### HasInstallConfig

`func (o *HandlersCreateClusterRequestData) HasInstallConfig() bool`

HasInstallConfig returns a boolean if a field has been set.

### GetMasterCount

`func (o *HandlersCreateClusterRequestData) GetMasterCount() int32`

GetMasterCount returns the MasterCount field if non-nil, zero value otherwise.

### GetMasterCountOk

`func (o *HandlersCreateClusterRequestData) GetMasterCountOk() (*int32, bool)`

GetMasterCountOk returns a tuple with the MasterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterCount

`func (o *HandlersCreateClusterRequestData) SetMasterCount(v int32)`

SetMasterCount sets MasterCount field to given value.

### HasMasterCount

`func (o *HandlersCreateClusterRequestData) HasMasterCount() bool`

HasMasterCount returns a boolean if a field has been set.

### GetNetworkId

`func (o *HandlersCreateClusterRequestData) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *HandlersCreateClusterRequestData) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *HandlersCreateClusterRequestData) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *HandlersCreateClusterRequestData) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetStack

`func (o *HandlersCreateClusterRequestData) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *HandlersCreateClusterRequestData) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *HandlersCreateClusterRequestData) SetStack(v string)`

SetStack sets Stack field to given value.


### GetVersion

`func (o *HandlersCreateClusterRequestData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HandlersCreateClusterRequestData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HandlersCreateClusterRequestData) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVmIds

`func (o *HandlersCreateClusterRequestData) GetVmIds() []int32`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *HandlersCreateClusterRequestData) GetVmIdsOk() (*[]int32, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *HandlersCreateClusterRequestData) SetVmIds(v []int32)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *HandlersCreateClusterRequestData) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### GetVmRoles

`func (o *HandlersCreateClusterRequestData) GetVmRoles() []HandlersVMRole`

GetVmRoles returns the VmRoles field if non-nil, zero value otherwise.

### GetVmRolesOk

`func (o *HandlersCreateClusterRequestData) GetVmRolesOk() (*[]HandlersVMRole, bool)`

GetVmRolesOk returns a tuple with the VmRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRoles

`func (o *HandlersCreateClusterRequestData) SetVmRoles(v []HandlersVMRole)`

SetVmRoles sets VmRoles field to given value.

### HasVmRoles

`func (o *HandlersCreateClusterRequestData) HasVmRoles() bool`

HasVmRoles returns a boolean if a field has been set.

### GetWorkerCount

`func (o *HandlersCreateClusterRequestData) GetWorkerCount() int32`

GetWorkerCount returns the WorkerCount field if non-nil, zero value otherwise.

### GetWorkerCountOk

`func (o *HandlersCreateClusterRequestData) GetWorkerCountOk() (*int32, bool)`

GetWorkerCountOk returns a tuple with the WorkerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerCount

`func (o *HandlersCreateClusterRequestData) SetWorkerCount(v int32)`

SetWorkerCount sets WorkerCount field to given value.

### HasWorkerCount

`func (o *HandlersCreateClusterRequestData) HasWorkerCount() bool`

HasWorkerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


