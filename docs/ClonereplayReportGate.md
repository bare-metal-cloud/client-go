# ClonereplayReportGate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentHeartbeat** | Pointer to **string** | AgentHeartbeat: the BMC agent&#39;s heartbeat is fresh (or the failure). | [optional] 
**DeploymentsAvailable** | Pointer to **string** | DeploymentsAvailable: every deployment across the cloned installs&#39; namespaces carries the Available condition (or the failure naming the unavailable ones). | [optional] 
**NoPendingPvcs** | Pointer to **string** | NoPendingPVCs: no PersistentVolumeClaim sits in Pending (or the failure naming them). | [optional] 
**NodesReady** | Pointer to **string** | NodesReady: every cluster node reports Ready (\&quot;ok (3/3 nodes Ready)\&quot; or the failure naming the NotReady nodes). | [optional] 
**PvNodeAffinity** | Pointer to **string** | PVNodeAffinity is N/A for a without-data clone — no PVs travel. The check is marked, never omitted. | [optional] 
**WebhooksResolve** | Pointer to **string** | WebhooksResolve is N/A for a without-data clone — webhooks re-install behind their own probe gates. Marked, never omitted. | [optional] 

## Methods

### NewClonereplayReportGate

`func NewClonereplayReportGate() *ClonereplayReportGate`

NewClonereplayReportGate instantiates a new ClonereplayReportGate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClonereplayReportGateWithDefaults

`func NewClonereplayReportGateWithDefaults() *ClonereplayReportGate`

NewClonereplayReportGateWithDefaults instantiates a new ClonereplayReportGate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentHeartbeat

`func (o *ClonereplayReportGate) GetAgentHeartbeat() string`

GetAgentHeartbeat returns the AgentHeartbeat field if non-nil, zero value otherwise.

### GetAgentHeartbeatOk

`func (o *ClonereplayReportGate) GetAgentHeartbeatOk() (*string, bool)`

GetAgentHeartbeatOk returns a tuple with the AgentHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentHeartbeat

`func (o *ClonereplayReportGate) SetAgentHeartbeat(v string)`

SetAgentHeartbeat sets AgentHeartbeat field to given value.

### HasAgentHeartbeat

`func (o *ClonereplayReportGate) HasAgentHeartbeat() bool`

HasAgentHeartbeat returns a boolean if a field has been set.

### GetDeploymentsAvailable

`func (o *ClonereplayReportGate) GetDeploymentsAvailable() string`

GetDeploymentsAvailable returns the DeploymentsAvailable field if non-nil, zero value otherwise.

### GetDeploymentsAvailableOk

`func (o *ClonereplayReportGate) GetDeploymentsAvailableOk() (*string, bool)`

GetDeploymentsAvailableOk returns a tuple with the DeploymentsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsAvailable

`func (o *ClonereplayReportGate) SetDeploymentsAvailable(v string)`

SetDeploymentsAvailable sets DeploymentsAvailable field to given value.

### HasDeploymentsAvailable

`func (o *ClonereplayReportGate) HasDeploymentsAvailable() bool`

HasDeploymentsAvailable returns a boolean if a field has been set.

### GetNoPendingPvcs

`func (o *ClonereplayReportGate) GetNoPendingPvcs() string`

GetNoPendingPvcs returns the NoPendingPvcs field if non-nil, zero value otherwise.

### GetNoPendingPvcsOk

`func (o *ClonereplayReportGate) GetNoPendingPvcsOk() (*string, bool)`

GetNoPendingPvcsOk returns a tuple with the NoPendingPvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoPendingPvcs

`func (o *ClonereplayReportGate) SetNoPendingPvcs(v string)`

SetNoPendingPvcs sets NoPendingPvcs field to given value.

### HasNoPendingPvcs

`func (o *ClonereplayReportGate) HasNoPendingPvcs() bool`

HasNoPendingPvcs returns a boolean if a field has been set.

### GetNodesReady

`func (o *ClonereplayReportGate) GetNodesReady() string`

GetNodesReady returns the NodesReady field if non-nil, zero value otherwise.

### GetNodesReadyOk

`func (o *ClonereplayReportGate) GetNodesReadyOk() (*string, bool)`

GetNodesReadyOk returns a tuple with the NodesReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesReady

`func (o *ClonereplayReportGate) SetNodesReady(v string)`

SetNodesReady sets NodesReady field to given value.

### HasNodesReady

`func (o *ClonereplayReportGate) HasNodesReady() bool`

HasNodesReady returns a boolean if a field has been set.

### GetPvNodeAffinity

`func (o *ClonereplayReportGate) GetPvNodeAffinity() string`

GetPvNodeAffinity returns the PvNodeAffinity field if non-nil, zero value otherwise.

### GetPvNodeAffinityOk

`func (o *ClonereplayReportGate) GetPvNodeAffinityOk() (*string, bool)`

GetPvNodeAffinityOk returns a tuple with the PvNodeAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvNodeAffinity

`func (o *ClonereplayReportGate) SetPvNodeAffinity(v string)`

SetPvNodeAffinity sets PvNodeAffinity field to given value.

### HasPvNodeAffinity

`func (o *ClonereplayReportGate) HasPvNodeAffinity() bool`

HasPvNodeAffinity returns a boolean if a field has been set.

### GetWebhooksResolve

`func (o *ClonereplayReportGate) GetWebhooksResolve() string`

GetWebhooksResolve returns the WebhooksResolve field if non-nil, zero value otherwise.

### GetWebhooksResolveOk

`func (o *ClonereplayReportGate) GetWebhooksResolveOk() (*string, bool)`

GetWebhooksResolveOk returns a tuple with the WebhooksResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksResolve

`func (o *ClonereplayReportGate) SetWebhooksResolve(v string)`

SetWebhooksResolve sets WebhooksResolve field to given value.

### HasWebhooksResolve

`func (o *ClonereplayReportGate) HasWebhooksResolve() bool`

HasWebhooksResolve returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


