# AgentopsCheckK3sResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiHealthy** | Pointer to **bool** | APIHealthy is the read-only API-liveness proof (every non-agent-role RUNNING install — review M1): &#x60;k3s kubectl get --raw&#x3D;/readyz&#x60; with a short in-kubectl timeout. false &#x3D; not proven (agent role has no local API to ask, or the API did not answer). | [optional] 
**Binary** | Pointer to **bool** |  | [optional] 
**DataDir** | Pointer to **bool** |  | [optional] 
**DiskFreeBytes** | Pointer to **int32** |  | [optional] 
**NodeTokenHash** | Pointer to **string** | NodeTokenHash is the SHA256 hex of the k3s server join token (/var/lib/rancher/k3s/server/node-token) — the cluster-identity fingerprint the adopt decision&#39;s one-cluster rule compares. The HASH only: the raw token never crosses the wire. \&quot;\&quot; when the file is absent. A token that exists but cannot be fingerprinted (unreadable, sha256sum missing) FAILS the whole probe (non-zero exit → the caller&#39;s fail-closed probe_unavailable path) — it must never read as an empty fingerprint that vacuously \&quot;agrees\&quot; with another member&#39;s. | [optional] 
**Role** | Pointer to **string** | Role classifies the install: RoleServer or RoleAgent; \&quot;\&quot; when no data dir was seen (the unit + pgrep evidence still speak). | [optional] 
**Running** | Pointer to **bool** |  | [optional] 
**Unit** | Pointer to **string** | Unit is the systemd unit found ACTIVE: \&quot;k3s.service\&quot; (server) or \&quot;k3s-agent.service\&quot; (worker). \&quot;\&quot; when nothing is active or the guest is systemd-less (the pgrep fallback cannot name a unit). | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentopsCheckK3sResult

`func NewAgentopsCheckK3sResult() *AgentopsCheckK3sResult`

NewAgentopsCheckK3sResult instantiates a new AgentopsCheckK3sResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentopsCheckK3sResultWithDefaults

`func NewAgentopsCheckK3sResultWithDefaults() *AgentopsCheckK3sResult`

NewAgentopsCheckK3sResultWithDefaults instantiates a new AgentopsCheckK3sResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiHealthy

`func (o *AgentopsCheckK3sResult) GetApiHealthy() bool`

GetApiHealthy returns the ApiHealthy field if non-nil, zero value otherwise.

### GetApiHealthyOk

`func (o *AgentopsCheckK3sResult) GetApiHealthyOk() (*bool, bool)`

GetApiHealthyOk returns a tuple with the ApiHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiHealthy

`func (o *AgentopsCheckK3sResult) SetApiHealthy(v bool)`

SetApiHealthy sets ApiHealthy field to given value.

### HasApiHealthy

`func (o *AgentopsCheckK3sResult) HasApiHealthy() bool`

HasApiHealthy returns a boolean if a field has been set.

### GetBinary

`func (o *AgentopsCheckK3sResult) GetBinary() bool`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *AgentopsCheckK3sResult) GetBinaryOk() (*bool, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *AgentopsCheckK3sResult) SetBinary(v bool)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *AgentopsCheckK3sResult) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetDataDir

`func (o *AgentopsCheckK3sResult) GetDataDir() bool`

GetDataDir returns the DataDir field if non-nil, zero value otherwise.

### GetDataDirOk

`func (o *AgentopsCheckK3sResult) GetDataDirOk() (*bool, bool)`

GetDataDirOk returns a tuple with the DataDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDir

`func (o *AgentopsCheckK3sResult) SetDataDir(v bool)`

SetDataDir sets DataDir field to given value.

### HasDataDir

`func (o *AgentopsCheckK3sResult) HasDataDir() bool`

HasDataDir returns a boolean if a field has been set.

### GetDiskFreeBytes

`func (o *AgentopsCheckK3sResult) GetDiskFreeBytes() int32`

GetDiskFreeBytes returns the DiskFreeBytes field if non-nil, zero value otherwise.

### GetDiskFreeBytesOk

`func (o *AgentopsCheckK3sResult) GetDiskFreeBytesOk() (*int32, bool)`

GetDiskFreeBytesOk returns a tuple with the DiskFreeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFreeBytes

`func (o *AgentopsCheckK3sResult) SetDiskFreeBytes(v int32)`

SetDiskFreeBytes sets DiskFreeBytes field to given value.

### HasDiskFreeBytes

`func (o *AgentopsCheckK3sResult) HasDiskFreeBytes() bool`

HasDiskFreeBytes returns a boolean if a field has been set.

### GetNodeTokenHash

`func (o *AgentopsCheckK3sResult) GetNodeTokenHash() string`

GetNodeTokenHash returns the NodeTokenHash field if non-nil, zero value otherwise.

### GetNodeTokenHashOk

`func (o *AgentopsCheckK3sResult) GetNodeTokenHashOk() (*string, bool)`

GetNodeTokenHashOk returns a tuple with the NodeTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTokenHash

`func (o *AgentopsCheckK3sResult) SetNodeTokenHash(v string)`

SetNodeTokenHash sets NodeTokenHash field to given value.

### HasNodeTokenHash

`func (o *AgentopsCheckK3sResult) HasNodeTokenHash() bool`

HasNodeTokenHash returns a boolean if a field has been set.

### GetRole

`func (o *AgentopsCheckK3sResult) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AgentopsCheckK3sResult) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AgentopsCheckK3sResult) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AgentopsCheckK3sResult) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRunning

`func (o *AgentopsCheckK3sResult) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *AgentopsCheckK3sResult) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *AgentopsCheckK3sResult) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *AgentopsCheckK3sResult) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetUnit

`func (o *AgentopsCheckK3sResult) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AgentopsCheckK3sResult) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AgentopsCheckK3sResult) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AgentopsCheckK3sResult) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetVersion

`func (o *AgentopsCheckK3sResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgentopsCheckK3sResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgentopsCheckK3sResult) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AgentopsCheckK3sResult) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


