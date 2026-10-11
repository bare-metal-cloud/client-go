# ModelsVirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | Pointer to **string** | AccessMode (Epic 89 — BYOV VM Console) determines interactive-shell policy.   byov     — user-owned VM (own_managed, user cloud key, multipass guest).              End-user Console / SSH allowed (role-gated).   platform — provisioned with a BMC system cloud credential              (CloudProviderCredential.IsSystemOwned). End-user shell              forbidden; future admin panel is a separate surface.  Default \&quot;byov\&quot; preserves existing UX for rows that predate this field (GORM backfills the column default on AutoMigrate — dev-stage policy). | [optional] 
**AgentConnectedAt** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **string** | Unique agent identifier | [optional] 
**AgentInstalled** | Pointer to **bool** | Agent installation tracking (Story 18-vm-04) | [optional] 
**AgentInstalledAt** | Pointer to **string** |  | [optional] 
**Architecture** | Pointer to **string** | x86_64, aarch64, etc. | [optional] 
**Cluster** | Pointer to [**ModelsCluster**](ModelsCluster.md) |  | [optional] 
**ClusterId** | Pointer to **int32** | Cluster association (optional - VM can be standalone or part of a cluster) Note: Cascade deletion is handled at application level based on ClusterOrigin - OriginCreated: VMs are deleted when cluster is deleted (if user opts in) - OriginImported: VMs are never deleted automatically; only k8s is uninstalled | [optional] 
**ConnectionId** | Pointer to **int32** |  | [optional] 
**ConnectionType** | Pointer to **string** | Polymorphic association for connection connection_type examples: managed_ssh_connections, aws_key_pair, gcp_service_account | [optional] 
**CpuCores** | Pointer to **int32** | Resource information (optional) | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**DeletionReasonCode** | Pointer to **string** | DeletionReasonCode is the machine-contract reason a VM removal landed in cleanup_incomplete (issue #1454, AC3): one of the VMRemovalReason* constants. NULL on every row not in cleanup_incomplete. The FE maps codes → copy; unknown codes render status_reason (sanitized human fallback detail), so the exact strings are a wire contract. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DiskSize** | Pointer to **int32** | In GB | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**K3sState** | Pointer to **string** | K3SState is the guest-truth k3s lifecycle state of this VM&#39;s host (story 124-1 AC6, epic-124): &#39;none&#39; &#x3D; no k3s installed on the guest; &#39;installed&#39; &#x3D; a successful SetupNode/uninstall-symmetric moment completed; &#39;unknown&#39; &#x3D; never observed (the default — rows predating the column and guests outside cluster flows backfill nothing). Writers land ONLY at the guest-state moments that exist today, both via WriteNodePhase (the ADR-060 single-writer choke point): phase ready → installed, phase removed → none. The 124.3 probe writer and the detach / delete-keep writers arrive with their owning stories. | [optional] 
**KernelVersion** | Pointer to **string** |  | [optional] 
**LastSeenAt** | Pointer to **string** |  | [optional] 
**MetricsLastCollectedAt** | Pointer to **string** | Last successful metrics collection | [optional] 
**MetricsRetentionHours** | Pointer to **int32** | Metrics collection (Story 20-12: On-Demand Metrics) | [optional] 
**Name** | Pointer to **string** | VM identification | [optional] 
**Network** | Pointer to [**ModelsNetwork**](ModelsNetwork.md) |  | [optional] 
**NetworkId** | Pointer to **int32** | Network association (optional - VM can be part of a network) | [optional] 
**OperatingSystem** | Pointer to **string** | System information (auto-detected after connection) | [optional] 
**OrgId** | Pointer to **int32** | Tenant scope (ADR-062): replaces the former owner-only UserID. | [optional] 
**Origin** | Pointer to [**ModelsResourceOrigin**](ModelsResourceOrigin.md) |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Provider** | Pointer to **string** | Provider information | [optional] 
**ProviderVmId** | Pointer to **string** | Provider-specific VM identifier (e.g., EC2 instance ID, GCP instance name, Azure VM name) | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**RamSize** | Pointer to **int32** | In MB | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** | \&quot;bastion\&quot;, \&quot;worker\&quot;, \&quot;\&quot; (standalone) | [optional] 
**SpecsStatus** | Pointer to **string** | SpecsStatus is the EVIDENCE status of the CPU/RAM/disk spec columns (story 124-2, D4): &#39;reported&#39; &#x3D; an agent metrics report carrying cpu_count/memory_total_bytes/disk_total_bytes has been seen for this VM; &#39;unknown&#39; &#x3D; never observed (the default — no report, no inference). The only writer after creation is the D4 projection in the metrics collector (services.SpecProjector): create-time declared specs and the SSH display migration are not evidence — specs_status stays &#39;unknown&#39; for them. The status is NEVER derived from the column values (never \&quot;0 means unknown\&quot;). Consumers: cluster-create preflight (124.3) and the VM pages&#39; specs pill. | [optional] 
**SshKeyEncryptedDataId** | Pointer to **int32** | Encrypted SSH Key Storage (Polymorphic encrypted_data table) References the EncryptedData table where the SSH private key is stored encrypted | [optional] 
**SshPublicKey** | Pointer to **string** | SSH Access (for VMs where we provisioned SSH keys) | [optional] 
**SshUsername** | Pointer to **string** | SSH username (default: ubuntu) | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**StatusReason** | Pointer to **string** |  | [optional] 
**StoppedAt** | Pointer to **string** | Story 18-vm-07: Graceful shutdown timestamp | [optional] 
**Tags** | Pointer to **string** | Tags (JSON field) | [optional] 
**Tier** | Pointer to [**ModelsResourceTier**](ModelsResourceTier.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsVirtualMachine

`func NewModelsVirtualMachine() *ModelsVirtualMachine`

NewModelsVirtualMachine instantiates a new ModelsVirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsVirtualMachineWithDefaults

`func NewModelsVirtualMachineWithDefaults() *ModelsVirtualMachine`

NewModelsVirtualMachineWithDefaults instantiates a new ModelsVirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *ModelsVirtualMachine) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *ModelsVirtualMachine) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *ModelsVirtualMachine) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *ModelsVirtualMachine) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetAgentConnectedAt

`func (o *ModelsVirtualMachine) GetAgentConnectedAt() string`

GetAgentConnectedAt returns the AgentConnectedAt field if non-nil, zero value otherwise.

### GetAgentConnectedAtOk

`func (o *ModelsVirtualMachine) GetAgentConnectedAtOk() (*string, bool)`

GetAgentConnectedAtOk returns a tuple with the AgentConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConnectedAt

`func (o *ModelsVirtualMachine) SetAgentConnectedAt(v string)`

SetAgentConnectedAt sets AgentConnectedAt field to given value.

### HasAgentConnectedAt

`func (o *ModelsVirtualMachine) HasAgentConnectedAt() bool`

HasAgentConnectedAt returns a boolean if a field has been set.

### GetAgentId

`func (o *ModelsVirtualMachine) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ModelsVirtualMachine) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ModelsVirtualMachine) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ModelsVirtualMachine) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *ModelsVirtualMachine) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *ModelsVirtualMachine) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *ModelsVirtualMachine) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *ModelsVirtualMachine) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetAgentInstalledAt

`func (o *ModelsVirtualMachine) GetAgentInstalledAt() string`

GetAgentInstalledAt returns the AgentInstalledAt field if non-nil, zero value otherwise.

### GetAgentInstalledAtOk

`func (o *ModelsVirtualMachine) GetAgentInstalledAtOk() (*string, bool)`

GetAgentInstalledAtOk returns a tuple with the AgentInstalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalledAt

`func (o *ModelsVirtualMachine) SetAgentInstalledAt(v string)`

SetAgentInstalledAt sets AgentInstalledAt field to given value.

### HasAgentInstalledAt

`func (o *ModelsVirtualMachine) HasAgentInstalledAt() bool`

HasAgentInstalledAt returns a boolean if a field has been set.

### GetArchitecture

`func (o *ModelsVirtualMachine) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ModelsVirtualMachine) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ModelsVirtualMachine) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ModelsVirtualMachine) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetCluster

`func (o *ModelsVirtualMachine) GetCluster() ModelsCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ModelsVirtualMachine) GetClusterOk() (*ModelsCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ModelsVirtualMachine) SetCluster(v ModelsCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ModelsVirtualMachine) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetClusterId

`func (o *ModelsVirtualMachine) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ModelsVirtualMachine) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ModelsVirtualMachine) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ModelsVirtualMachine) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetConnectionId

`func (o *ModelsVirtualMachine) GetConnectionId() int32`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ModelsVirtualMachine) GetConnectionIdOk() (*int32, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ModelsVirtualMachine) SetConnectionId(v int32)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ModelsVirtualMachine) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionType

`func (o *ModelsVirtualMachine) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ModelsVirtualMachine) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ModelsVirtualMachine) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *ModelsVirtualMachine) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetCpuCores

`func (o *ModelsVirtualMachine) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *ModelsVirtualMachine) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *ModelsVirtualMachine) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *ModelsVirtualMachine) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsVirtualMachine) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsVirtualMachine) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsVirtualMachine) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsVirtualMachine) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsVirtualMachine) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsVirtualMachine) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsVirtualMachine) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsVirtualMachine) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDeletionReasonCode

`func (o *ModelsVirtualMachine) GetDeletionReasonCode() string`

GetDeletionReasonCode returns the DeletionReasonCode field if non-nil, zero value otherwise.

### GetDeletionReasonCodeOk

`func (o *ModelsVirtualMachine) GetDeletionReasonCodeOk() (*string, bool)`

GetDeletionReasonCodeOk returns a tuple with the DeletionReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionReasonCode

`func (o *ModelsVirtualMachine) SetDeletionReasonCode(v string)`

SetDeletionReasonCode sets DeletionReasonCode field to given value.

### HasDeletionReasonCode

`func (o *ModelsVirtualMachine) HasDeletionReasonCode() bool`

HasDeletionReasonCode returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsVirtualMachine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsVirtualMachine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsVirtualMachine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsVirtualMachine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskSize

`func (o *ModelsVirtualMachine) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ModelsVirtualMachine) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ModelsVirtualMachine) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *ModelsVirtualMachine) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetId

`func (o *ModelsVirtualMachine) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsVirtualMachine) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsVirtualMachine) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsVirtualMachine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetK3sState

`func (o *ModelsVirtualMachine) GetK3sState() string`

GetK3sState returns the K3sState field if non-nil, zero value otherwise.

### GetK3sStateOk

`func (o *ModelsVirtualMachine) GetK3sStateOk() (*string, bool)`

GetK3sStateOk returns a tuple with the K3sState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK3sState

`func (o *ModelsVirtualMachine) SetK3sState(v string)`

SetK3sState sets K3sState field to given value.

### HasK3sState

`func (o *ModelsVirtualMachine) HasK3sState() bool`

HasK3sState returns a boolean if a field has been set.

### GetKernelVersion

`func (o *ModelsVirtualMachine) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *ModelsVirtualMachine) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *ModelsVirtualMachine) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.

### HasKernelVersion

`func (o *ModelsVirtualMachine) HasKernelVersion() bool`

HasKernelVersion returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *ModelsVirtualMachine) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ModelsVirtualMachine) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ModelsVirtualMachine) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *ModelsVirtualMachine) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetMetricsLastCollectedAt

`func (o *ModelsVirtualMachine) GetMetricsLastCollectedAt() string`

GetMetricsLastCollectedAt returns the MetricsLastCollectedAt field if non-nil, zero value otherwise.

### GetMetricsLastCollectedAtOk

`func (o *ModelsVirtualMachine) GetMetricsLastCollectedAtOk() (*string, bool)`

GetMetricsLastCollectedAtOk returns a tuple with the MetricsLastCollectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsLastCollectedAt

`func (o *ModelsVirtualMachine) SetMetricsLastCollectedAt(v string)`

SetMetricsLastCollectedAt sets MetricsLastCollectedAt field to given value.

### HasMetricsLastCollectedAt

`func (o *ModelsVirtualMachine) HasMetricsLastCollectedAt() bool`

HasMetricsLastCollectedAt returns a boolean if a field has been set.

### GetMetricsRetentionHours

`func (o *ModelsVirtualMachine) GetMetricsRetentionHours() int32`

GetMetricsRetentionHours returns the MetricsRetentionHours field if non-nil, zero value otherwise.

### GetMetricsRetentionHoursOk

`func (o *ModelsVirtualMachine) GetMetricsRetentionHoursOk() (*int32, bool)`

GetMetricsRetentionHoursOk returns a tuple with the MetricsRetentionHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionHours

`func (o *ModelsVirtualMachine) SetMetricsRetentionHours(v int32)`

SetMetricsRetentionHours sets MetricsRetentionHours field to given value.

### HasMetricsRetentionHours

`func (o *ModelsVirtualMachine) HasMetricsRetentionHours() bool`

HasMetricsRetentionHours returns a boolean if a field has been set.

### GetName

`func (o *ModelsVirtualMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsVirtualMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsVirtualMachine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsVirtualMachine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *ModelsVirtualMachine) GetNetwork() ModelsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ModelsVirtualMachine) GetNetworkOk() (*ModelsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ModelsVirtualMachine) SetNetwork(v ModelsNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ModelsVirtualMachine) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkId

`func (o *ModelsVirtualMachine) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ModelsVirtualMachine) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ModelsVirtualMachine) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ModelsVirtualMachine) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *ModelsVirtualMachine) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ModelsVirtualMachine) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ModelsVirtualMachine) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ModelsVirtualMachine) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsVirtualMachine) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsVirtualMachine) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsVirtualMachine) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsVirtualMachine) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrigin

`func (o *ModelsVirtualMachine) GetOrigin() ModelsResourceOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ModelsVirtualMachine) GetOriginOk() (*ModelsResourceOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ModelsVirtualMachine) SetOrigin(v ModelsResourceOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ModelsVirtualMachine) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetOsVersion

`func (o *ModelsVirtualMachine) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ModelsVirtualMachine) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ModelsVirtualMachine) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *ModelsVirtualMachine) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPrivateIp

`func (o *ModelsVirtualMachine) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *ModelsVirtualMachine) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *ModelsVirtualMachine) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *ModelsVirtualMachine) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsVirtualMachine) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsVirtualMachine) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsVirtualMachine) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsVirtualMachine) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProvider

`func (o *ModelsVirtualMachine) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelsVirtualMachine) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelsVirtualMachine) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelsVirtualMachine) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderVmId

`func (o *ModelsVirtualMachine) GetProviderVmId() string`

GetProviderVmId returns the ProviderVmId field if non-nil, zero value otherwise.

### GetProviderVmIdOk

`func (o *ModelsVirtualMachine) GetProviderVmIdOk() (*string, bool)`

GetProviderVmIdOk returns a tuple with the ProviderVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVmId

`func (o *ModelsVirtualMachine) SetProviderVmId(v string)`

SetProviderVmId sets ProviderVmId field to given value.

### HasProviderVmId

`func (o *ModelsVirtualMachine) HasProviderVmId() bool`

HasProviderVmId returns a boolean if a field has been set.

### GetPublicIp

`func (o *ModelsVirtualMachine) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ModelsVirtualMachine) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ModelsVirtualMachine) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ModelsVirtualMachine) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRamSize

`func (o *ModelsVirtualMachine) GetRamSize() int32`

GetRamSize returns the RamSize field if non-nil, zero value otherwise.

### GetRamSizeOk

`func (o *ModelsVirtualMachine) GetRamSizeOk() (*int32, bool)`

GetRamSizeOk returns a tuple with the RamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamSize

`func (o *ModelsVirtualMachine) SetRamSize(v int32)`

SetRamSize sets RamSize field to given value.

### HasRamSize

`func (o *ModelsVirtualMachine) HasRamSize() bool`

HasRamSize returns a boolean if a field has been set.

### GetRegion

`func (o *ModelsVirtualMachine) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ModelsVirtualMachine) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ModelsVirtualMachine) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ModelsVirtualMachine) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRole

`func (o *ModelsVirtualMachine) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ModelsVirtualMachine) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ModelsVirtualMachine) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ModelsVirtualMachine) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSpecsStatus

`func (o *ModelsVirtualMachine) GetSpecsStatus() string`

GetSpecsStatus returns the SpecsStatus field if non-nil, zero value otherwise.

### GetSpecsStatusOk

`func (o *ModelsVirtualMachine) GetSpecsStatusOk() (*string, bool)`

GetSpecsStatusOk returns a tuple with the SpecsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecsStatus

`func (o *ModelsVirtualMachine) SetSpecsStatus(v string)`

SetSpecsStatus sets SpecsStatus field to given value.

### HasSpecsStatus

`func (o *ModelsVirtualMachine) HasSpecsStatus() bool`

HasSpecsStatus returns a boolean if a field has been set.

### GetSshKeyEncryptedDataId

`func (o *ModelsVirtualMachine) GetSshKeyEncryptedDataId() int32`

GetSshKeyEncryptedDataId returns the SshKeyEncryptedDataId field if non-nil, zero value otherwise.

### GetSshKeyEncryptedDataIdOk

`func (o *ModelsVirtualMachine) GetSshKeyEncryptedDataIdOk() (*int32, bool)`

GetSshKeyEncryptedDataIdOk returns a tuple with the SshKeyEncryptedDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyEncryptedDataId

`func (o *ModelsVirtualMachine) SetSshKeyEncryptedDataId(v int32)`

SetSshKeyEncryptedDataId sets SshKeyEncryptedDataId field to given value.

### HasSshKeyEncryptedDataId

`func (o *ModelsVirtualMachine) HasSshKeyEncryptedDataId() bool`

HasSshKeyEncryptedDataId returns a boolean if a field has been set.

### GetSshPublicKey

`func (o *ModelsVirtualMachine) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *ModelsVirtualMachine) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *ModelsVirtualMachine) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *ModelsVirtualMachine) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### GetSshUsername

`func (o *ModelsVirtualMachine) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ModelsVirtualMachine) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ModelsVirtualMachine) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ModelsVirtualMachine) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsVirtualMachine) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsVirtualMachine) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsVirtualMachine) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsVirtualMachine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *ModelsVirtualMachine) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *ModelsVirtualMachine) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *ModelsVirtualMachine) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *ModelsVirtualMachine) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetStoppedAt

`func (o *ModelsVirtualMachine) GetStoppedAt() string`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *ModelsVirtualMachine) GetStoppedAtOk() (*string, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *ModelsVirtualMachine) SetStoppedAt(v string)`

SetStoppedAt sets StoppedAt field to given value.

### HasStoppedAt

`func (o *ModelsVirtualMachine) HasStoppedAt() bool`

HasStoppedAt returns a boolean if a field has been set.

### GetTags

`func (o *ModelsVirtualMachine) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelsVirtualMachine) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelsVirtualMachine) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelsVirtualMachine) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTier

`func (o *ModelsVirtualMachine) GetTier() ModelsResourceTier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *ModelsVirtualMachine) GetTierOk() (*ModelsResourceTier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *ModelsVirtualMachine) SetTier(v ModelsResourceTier)`

SetTier sets Tier field to given value.

### HasTier

`func (o *ModelsVirtualMachine) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsVirtualMachine) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsVirtualMachine) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsVirtualMachine) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsVirtualMachine) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


