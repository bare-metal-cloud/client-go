# HandlersRunnerMachineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentVersion** | Pointer to **string** |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
**Bound** | Pointer to **bool** | Bound: true &#x3D; a live runner binding (runner_pool_enlistments row, released_at IS NULL); false &#x3D; unbound, available to Add. | [optional] 
**CleanupDiskPercent** | Pointer to **int32** | CleanupDiskPercent is the machine&#39;s auto disk-cleanup threshold (AAR.5 item 3): 0 &#x3D; not configured, else the stored 10-90 band value. The setting&#39;s own read/write rides the cleanup route. | [optional] 
**CredentialState** | Pointer to **string** | CredentialState is B1.7&#39;s visible rotation state; empty (omitted) on healthy machines — B1.1 always omits it. | [optional] 
**CurrentJobId** | Pointer to **string** | CurrentJobID is the running job&#39;s id while the lifecycle is busy. | [optional] 
**Drain** | Pointer to [**HandlersRunnerMachineDrainInfo**](HandlersRunnerMachineDrainInfo.md) | Drain is the machine&#39;s live transition bookkeeping (B1.4): present while the machine is draining, quarantined-with-a-pending-drain or returning. The overdue flag is the drain-timeout REPORT — the operator&#39;s signal that the force option is open. | [optional] 
**EnrolledAt** | Pointer to **string** |  | [optional] 
**HeartbeatAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Lifecycle** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeedsAttention** | Pointer to **bool** | NeedsAttention is the quarantine surface&#39;s state flag (AAR.5 item 4): true while the machine holds quarantined — the panel renders \&quot;needs attention / re-enlist\&quot; from it. Re-enlist is the return action (the existing lifecycle service); there is no silent untag path (release refuses from quarantined — the transition table). | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Ownership** | Pointer to **string** |  | [optional] 
**ProfilesServed** | Pointer to **[]string** | ProfilesServed names the profiles this machine can serve; empty in B1.1 (B2.2&#39;s profile read lands the vocabulary). | [optional] 
**ProviderId** | Pointer to **string** | ProviderID is the pool&#39;s join key (GH #4141): the machine&#39;s virtual_machines row id as a string — the id space the job ledger&#39;s runner column (runner_jobs.vm_provider_id, minted by the floor as the VM row id) and the connections detail&#39;s runnerIds carry. The frontend joins machines to ledger rows and connection runners through it (in live mode the agent-row id matches neither). Empty (omitted) on an agent row with no VM behind it. | [optional] 
**Reachability** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** | Scope / Ownership surface the machine&#39;s live binding&#39;s pool scope and ownership class (AAR.5 item 1, GH #3206): \&quot;org\&quot;/\&quot;customer_owned\&quot; on every self-hosted machine in this org&#39;s list — the org surface only ever creates org-scope, customer-owned bindings. Empty on an unbound row (there is no binding to classify). | [optional] 
**Shape** | Pointer to [**HandlersRunnerMachineShape**](HandlersRunnerMachineShape.md) |  | [optional] 
**TodayMinutes** | Pointer to **float32** |  | [optional] 
**TodayRuns** | Pointer to **int32** | TodayRuns / TodayMinutes are today&#39;s (UTC) completed runs and metered minutes for THIS machine — the phase-1 pages&#39; per-row facts (P1.3&#39;s list, P1.6&#39;s detail). | [optional] 

## Methods

### NewHandlersRunnerMachineResponse

`func NewHandlersRunnerMachineResponse() *HandlersRunnerMachineResponse`

NewHandlersRunnerMachineResponse instantiates a new HandlersRunnerMachineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerMachineResponseWithDefaults

`func NewHandlersRunnerMachineResponseWithDefaults() *HandlersRunnerMachineResponse`

NewHandlersRunnerMachineResponseWithDefaults instantiates a new HandlersRunnerMachineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentVersion

`func (o *HandlersRunnerMachineResponse) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *HandlersRunnerMachineResponse) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *HandlersRunnerMachineResponse) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *HandlersRunnerMachineResponse) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetArch

`func (o *HandlersRunnerMachineResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *HandlersRunnerMachineResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *HandlersRunnerMachineResponse) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *HandlersRunnerMachineResponse) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetBound

`func (o *HandlersRunnerMachineResponse) GetBound() bool`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *HandlersRunnerMachineResponse) GetBoundOk() (*bool, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *HandlersRunnerMachineResponse) SetBound(v bool)`

SetBound sets Bound field to given value.

### HasBound

`func (o *HandlersRunnerMachineResponse) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetCleanupDiskPercent

`func (o *HandlersRunnerMachineResponse) GetCleanupDiskPercent() int32`

GetCleanupDiskPercent returns the CleanupDiskPercent field if non-nil, zero value otherwise.

### GetCleanupDiskPercentOk

`func (o *HandlersRunnerMachineResponse) GetCleanupDiskPercentOk() (*int32, bool)`

GetCleanupDiskPercentOk returns a tuple with the CleanupDiskPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupDiskPercent

`func (o *HandlersRunnerMachineResponse) SetCleanupDiskPercent(v int32)`

SetCleanupDiskPercent sets CleanupDiskPercent field to given value.

### HasCleanupDiskPercent

`func (o *HandlersRunnerMachineResponse) HasCleanupDiskPercent() bool`

HasCleanupDiskPercent returns a boolean if a field has been set.

### GetCredentialState

`func (o *HandlersRunnerMachineResponse) GetCredentialState() string`

GetCredentialState returns the CredentialState field if non-nil, zero value otherwise.

### GetCredentialStateOk

`func (o *HandlersRunnerMachineResponse) GetCredentialStateOk() (*string, bool)`

GetCredentialStateOk returns a tuple with the CredentialState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialState

`func (o *HandlersRunnerMachineResponse) SetCredentialState(v string)`

SetCredentialState sets CredentialState field to given value.

### HasCredentialState

`func (o *HandlersRunnerMachineResponse) HasCredentialState() bool`

HasCredentialState returns a boolean if a field has been set.

### GetCurrentJobId

`func (o *HandlersRunnerMachineResponse) GetCurrentJobId() string`

GetCurrentJobId returns the CurrentJobId field if non-nil, zero value otherwise.

### GetCurrentJobIdOk

`func (o *HandlersRunnerMachineResponse) GetCurrentJobIdOk() (*string, bool)`

GetCurrentJobIdOk returns a tuple with the CurrentJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentJobId

`func (o *HandlersRunnerMachineResponse) SetCurrentJobId(v string)`

SetCurrentJobId sets CurrentJobId field to given value.

### HasCurrentJobId

`func (o *HandlersRunnerMachineResponse) HasCurrentJobId() bool`

HasCurrentJobId returns a boolean if a field has been set.

### GetDrain

`func (o *HandlersRunnerMachineResponse) GetDrain() HandlersRunnerMachineDrainInfo`

GetDrain returns the Drain field if non-nil, zero value otherwise.

### GetDrainOk

`func (o *HandlersRunnerMachineResponse) GetDrainOk() (*HandlersRunnerMachineDrainInfo, bool)`

GetDrainOk returns a tuple with the Drain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrain

`func (o *HandlersRunnerMachineResponse) SetDrain(v HandlersRunnerMachineDrainInfo)`

SetDrain sets Drain field to given value.

### HasDrain

`func (o *HandlersRunnerMachineResponse) HasDrain() bool`

HasDrain returns a boolean if a field has been set.

### GetEnrolledAt

`func (o *HandlersRunnerMachineResponse) GetEnrolledAt() string`

GetEnrolledAt returns the EnrolledAt field if non-nil, zero value otherwise.

### GetEnrolledAtOk

`func (o *HandlersRunnerMachineResponse) GetEnrolledAtOk() (*string, bool)`

GetEnrolledAtOk returns a tuple with the EnrolledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledAt

`func (o *HandlersRunnerMachineResponse) SetEnrolledAt(v string)`

SetEnrolledAt sets EnrolledAt field to given value.

### HasEnrolledAt

`func (o *HandlersRunnerMachineResponse) HasEnrolledAt() bool`

HasEnrolledAt returns a boolean if a field has been set.

### GetHeartbeatAt

`func (o *HandlersRunnerMachineResponse) GetHeartbeatAt() string`

GetHeartbeatAt returns the HeartbeatAt field if non-nil, zero value otherwise.

### GetHeartbeatAtOk

`func (o *HandlersRunnerMachineResponse) GetHeartbeatAtOk() (*string, bool)`

GetHeartbeatAtOk returns a tuple with the HeartbeatAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatAt

`func (o *HandlersRunnerMachineResponse) SetHeartbeatAt(v string)`

SetHeartbeatAt sets HeartbeatAt field to given value.

### HasHeartbeatAt

`func (o *HandlersRunnerMachineResponse) HasHeartbeatAt() bool`

HasHeartbeatAt returns a boolean if a field has been set.

### GetId

`func (o *HandlersRunnerMachineResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersRunnerMachineResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersRunnerMachineResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersRunnerMachineResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *HandlersRunnerMachineResponse) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HandlersRunnerMachineResponse) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HandlersRunnerMachineResponse) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HandlersRunnerMachineResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLifecycle

`func (o *HandlersRunnerMachineResponse) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *HandlersRunnerMachineResponse) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *HandlersRunnerMachineResponse) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *HandlersRunnerMachineResponse) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetName

`func (o *HandlersRunnerMachineResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersRunnerMachineResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersRunnerMachineResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersRunnerMachineResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedsAttention

`func (o *HandlersRunnerMachineResponse) GetNeedsAttention() bool`

GetNeedsAttention returns the NeedsAttention field if non-nil, zero value otherwise.

### GetNeedsAttentionOk

`func (o *HandlersRunnerMachineResponse) GetNeedsAttentionOk() (*bool, bool)`

GetNeedsAttentionOk returns a tuple with the NeedsAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsAttention

`func (o *HandlersRunnerMachineResponse) SetNeedsAttention(v bool)`

SetNeedsAttention sets NeedsAttention field to given value.

### HasNeedsAttention

`func (o *HandlersRunnerMachineResponse) HasNeedsAttention() bool`

HasNeedsAttention returns a boolean if a field has been set.

### GetOs

`func (o *HandlersRunnerMachineResponse) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *HandlersRunnerMachineResponse) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *HandlersRunnerMachineResponse) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *HandlersRunnerMachineResponse) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOwnership

`func (o *HandlersRunnerMachineResponse) GetOwnership() string`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *HandlersRunnerMachineResponse) GetOwnershipOk() (*string, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *HandlersRunnerMachineResponse) SetOwnership(v string)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *HandlersRunnerMachineResponse) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.

### GetProfilesServed

`func (o *HandlersRunnerMachineResponse) GetProfilesServed() []string`

GetProfilesServed returns the ProfilesServed field if non-nil, zero value otherwise.

### GetProfilesServedOk

`func (o *HandlersRunnerMachineResponse) GetProfilesServedOk() (*[]string, bool)`

GetProfilesServedOk returns a tuple with the ProfilesServed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesServed

`func (o *HandlersRunnerMachineResponse) SetProfilesServed(v []string)`

SetProfilesServed sets ProfilesServed field to given value.

### HasProfilesServed

`func (o *HandlersRunnerMachineResponse) HasProfilesServed() bool`

HasProfilesServed returns a boolean if a field has been set.

### GetProviderId

`func (o *HandlersRunnerMachineResponse) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *HandlersRunnerMachineResponse) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *HandlersRunnerMachineResponse) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *HandlersRunnerMachineResponse) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetReachability

`func (o *HandlersRunnerMachineResponse) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *HandlersRunnerMachineResponse) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *HandlersRunnerMachineResponse) SetReachability(v string)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *HandlersRunnerMachineResponse) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetScope

`func (o *HandlersRunnerMachineResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersRunnerMachineResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersRunnerMachineResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersRunnerMachineResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetShape

`func (o *HandlersRunnerMachineResponse) GetShape() HandlersRunnerMachineShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *HandlersRunnerMachineResponse) GetShapeOk() (*HandlersRunnerMachineShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *HandlersRunnerMachineResponse) SetShape(v HandlersRunnerMachineShape)`

SetShape sets Shape field to given value.

### HasShape

`func (o *HandlersRunnerMachineResponse) HasShape() bool`

HasShape returns a boolean if a field has been set.

### GetTodayMinutes

`func (o *HandlersRunnerMachineResponse) GetTodayMinutes() float32`

GetTodayMinutes returns the TodayMinutes field if non-nil, zero value otherwise.

### GetTodayMinutesOk

`func (o *HandlersRunnerMachineResponse) GetTodayMinutesOk() (*float32, bool)`

GetTodayMinutesOk returns a tuple with the TodayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayMinutes

`func (o *HandlersRunnerMachineResponse) SetTodayMinutes(v float32)`

SetTodayMinutes sets TodayMinutes field to given value.

### HasTodayMinutes

`func (o *HandlersRunnerMachineResponse) HasTodayMinutes() bool`

HasTodayMinutes returns a boolean if a field has been set.

### GetTodayRuns

`func (o *HandlersRunnerMachineResponse) GetTodayRuns() int32`

GetTodayRuns returns the TodayRuns field if non-nil, zero value otherwise.

### GetTodayRunsOk

`func (o *HandlersRunnerMachineResponse) GetTodayRunsOk() (*int32, bool)`

GetTodayRunsOk returns a tuple with the TodayRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayRuns

`func (o *HandlersRunnerMachineResponse) SetTodayRuns(v int32)`

SetTodayRuns sets TodayRuns field to given value.

### HasTodayRuns

`func (o *HandlersRunnerMachineResponse) HasTodayRuns() bool`

HasTodayRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


