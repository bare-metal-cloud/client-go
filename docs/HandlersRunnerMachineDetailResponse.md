# HandlersRunnerMachineDetailResponse

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
**LabelProbe** | Pointer to **string** | LabelProbe is the label-probe verdict: \&quot;pass\&quot; only when verified runner capability evidence is on record (AAR.2&#39;s evidence law — never a default pass), else \&quot;unknown\&quot;. | [optional] 
**LabelProbeVerifiedAt** | Pointer to **string** | LabelProbeVerifiedAt is the AGE of that evidence: the oldest verified_at across the required labels on a pass, the retained last-good timestamp on a downgraded probe, omitted when no probe ever ran (B1.5). | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**LastJobId** | Pointer to **string** | LastJobID is the machine&#39;s most recent job (live first, then the latest terminal); empty when it never ran one. | [optional] 
**Lifecycle** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeedsAttention** | Pointer to **bool** | NeedsAttention is the quarantine surface&#39;s state flag (AAR.5 item 4): true while the machine holds quarantined — the panel renders \&quot;needs attention / re-enlist\&quot; from it. Re-enlist is the return action (the existing lifecycle service); there is no silent untag path (release refuses from quarantined — the transition table). | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** | Owner is the owning org&#39;s name. | [optional] 
**Ownership** | Pointer to **string** |  | [optional] 
**ProfilesServed** | Pointer to **[]string** | ProfilesServed names the profiles this machine can serve; empty in B1.1 (B2.2&#39;s profile read lands the vocabulary). | [optional] 
**ProviderId** | Pointer to **string** | ProviderID is the pool&#39;s join key (GH #4141): the machine&#39;s virtual_machines row id as a string — the id space the job ledger&#39;s runner column (runner_jobs.vm_provider_id, minted by the floor as the VM row id) and the connections detail&#39;s runnerIds carry. The frontend joins machines to ledger rows and connection runners through it (in live mode the agent-row id matches neither). Empty (omitted) on an agent row with no VM behind it. | [optional] 
**Reachability** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** | Scope / Ownership surface the machine&#39;s live binding&#39;s pool scope and ownership class (AAR.5 item 1, GH #3206): \&quot;org\&quot;/\&quot;customer_owned\&quot; on every self-hosted machine in this org&#39;s list — the org surface only ever creates org-scope, customer-owned bindings. Empty on an unbound row (there is no binding to classify). | [optional] 
**Shape** | Pointer to [**HandlersRunnerMachineShape**](HandlersRunnerMachineShape.md) |  | [optional] 
**TodayMinutes** | Pointer to **float32** |  | [optional] 
**TodayRuns** | Pointer to **int32** | TodayRuns / TodayMinutes are today&#39;s (UTC) completed runs and metered minutes for THIS machine — the phase-1 pages&#39; per-row facts (P1.3&#39;s list, P1.6&#39;s detail). | [optional] 

## Methods

### NewHandlersRunnerMachineDetailResponse

`func NewHandlersRunnerMachineDetailResponse() *HandlersRunnerMachineDetailResponse`

NewHandlersRunnerMachineDetailResponse instantiates a new HandlersRunnerMachineDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerMachineDetailResponseWithDefaults

`func NewHandlersRunnerMachineDetailResponseWithDefaults() *HandlersRunnerMachineDetailResponse`

NewHandlersRunnerMachineDetailResponseWithDefaults instantiates a new HandlersRunnerMachineDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentVersion

`func (o *HandlersRunnerMachineDetailResponse) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *HandlersRunnerMachineDetailResponse) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *HandlersRunnerMachineDetailResponse) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *HandlersRunnerMachineDetailResponse) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetArch

`func (o *HandlersRunnerMachineDetailResponse) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *HandlersRunnerMachineDetailResponse) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *HandlersRunnerMachineDetailResponse) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *HandlersRunnerMachineDetailResponse) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetBound

`func (o *HandlersRunnerMachineDetailResponse) GetBound() bool`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *HandlersRunnerMachineDetailResponse) GetBoundOk() (*bool, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *HandlersRunnerMachineDetailResponse) SetBound(v bool)`

SetBound sets Bound field to given value.

### HasBound

`func (o *HandlersRunnerMachineDetailResponse) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetCleanupDiskPercent

`func (o *HandlersRunnerMachineDetailResponse) GetCleanupDiskPercent() int32`

GetCleanupDiskPercent returns the CleanupDiskPercent field if non-nil, zero value otherwise.

### GetCleanupDiskPercentOk

`func (o *HandlersRunnerMachineDetailResponse) GetCleanupDiskPercentOk() (*int32, bool)`

GetCleanupDiskPercentOk returns a tuple with the CleanupDiskPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupDiskPercent

`func (o *HandlersRunnerMachineDetailResponse) SetCleanupDiskPercent(v int32)`

SetCleanupDiskPercent sets CleanupDiskPercent field to given value.

### HasCleanupDiskPercent

`func (o *HandlersRunnerMachineDetailResponse) HasCleanupDiskPercent() bool`

HasCleanupDiskPercent returns a boolean if a field has been set.

### GetCredentialState

`func (o *HandlersRunnerMachineDetailResponse) GetCredentialState() string`

GetCredentialState returns the CredentialState field if non-nil, zero value otherwise.

### GetCredentialStateOk

`func (o *HandlersRunnerMachineDetailResponse) GetCredentialStateOk() (*string, bool)`

GetCredentialStateOk returns a tuple with the CredentialState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialState

`func (o *HandlersRunnerMachineDetailResponse) SetCredentialState(v string)`

SetCredentialState sets CredentialState field to given value.

### HasCredentialState

`func (o *HandlersRunnerMachineDetailResponse) HasCredentialState() bool`

HasCredentialState returns a boolean if a field has been set.

### GetCurrentJobId

`func (o *HandlersRunnerMachineDetailResponse) GetCurrentJobId() string`

GetCurrentJobId returns the CurrentJobId field if non-nil, zero value otherwise.

### GetCurrentJobIdOk

`func (o *HandlersRunnerMachineDetailResponse) GetCurrentJobIdOk() (*string, bool)`

GetCurrentJobIdOk returns a tuple with the CurrentJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentJobId

`func (o *HandlersRunnerMachineDetailResponse) SetCurrentJobId(v string)`

SetCurrentJobId sets CurrentJobId field to given value.

### HasCurrentJobId

`func (o *HandlersRunnerMachineDetailResponse) HasCurrentJobId() bool`

HasCurrentJobId returns a boolean if a field has been set.

### GetDrain

`func (o *HandlersRunnerMachineDetailResponse) GetDrain() HandlersRunnerMachineDrainInfo`

GetDrain returns the Drain field if non-nil, zero value otherwise.

### GetDrainOk

`func (o *HandlersRunnerMachineDetailResponse) GetDrainOk() (*HandlersRunnerMachineDrainInfo, bool)`

GetDrainOk returns a tuple with the Drain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrain

`func (o *HandlersRunnerMachineDetailResponse) SetDrain(v HandlersRunnerMachineDrainInfo)`

SetDrain sets Drain field to given value.

### HasDrain

`func (o *HandlersRunnerMachineDetailResponse) HasDrain() bool`

HasDrain returns a boolean if a field has been set.

### GetEnrolledAt

`func (o *HandlersRunnerMachineDetailResponse) GetEnrolledAt() string`

GetEnrolledAt returns the EnrolledAt field if non-nil, zero value otherwise.

### GetEnrolledAtOk

`func (o *HandlersRunnerMachineDetailResponse) GetEnrolledAtOk() (*string, bool)`

GetEnrolledAtOk returns a tuple with the EnrolledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolledAt

`func (o *HandlersRunnerMachineDetailResponse) SetEnrolledAt(v string)`

SetEnrolledAt sets EnrolledAt field to given value.

### HasEnrolledAt

`func (o *HandlersRunnerMachineDetailResponse) HasEnrolledAt() bool`

HasEnrolledAt returns a boolean if a field has been set.

### GetHeartbeatAt

`func (o *HandlersRunnerMachineDetailResponse) GetHeartbeatAt() string`

GetHeartbeatAt returns the HeartbeatAt field if non-nil, zero value otherwise.

### GetHeartbeatAtOk

`func (o *HandlersRunnerMachineDetailResponse) GetHeartbeatAtOk() (*string, bool)`

GetHeartbeatAtOk returns a tuple with the HeartbeatAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatAt

`func (o *HandlersRunnerMachineDetailResponse) SetHeartbeatAt(v string)`

SetHeartbeatAt sets HeartbeatAt field to given value.

### HasHeartbeatAt

`func (o *HandlersRunnerMachineDetailResponse) HasHeartbeatAt() bool`

HasHeartbeatAt returns a boolean if a field has been set.

### GetId

`func (o *HandlersRunnerMachineDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersRunnerMachineDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersRunnerMachineDetailResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersRunnerMachineDetailResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabelProbe

`func (o *HandlersRunnerMachineDetailResponse) GetLabelProbe() string`

GetLabelProbe returns the LabelProbe field if non-nil, zero value otherwise.

### GetLabelProbeOk

`func (o *HandlersRunnerMachineDetailResponse) GetLabelProbeOk() (*string, bool)`

GetLabelProbeOk returns a tuple with the LabelProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelProbe

`func (o *HandlersRunnerMachineDetailResponse) SetLabelProbe(v string)`

SetLabelProbe sets LabelProbe field to given value.

### HasLabelProbe

`func (o *HandlersRunnerMachineDetailResponse) HasLabelProbe() bool`

HasLabelProbe returns a boolean if a field has been set.

### GetLabelProbeVerifiedAt

`func (o *HandlersRunnerMachineDetailResponse) GetLabelProbeVerifiedAt() string`

GetLabelProbeVerifiedAt returns the LabelProbeVerifiedAt field if non-nil, zero value otherwise.

### GetLabelProbeVerifiedAtOk

`func (o *HandlersRunnerMachineDetailResponse) GetLabelProbeVerifiedAtOk() (*string, bool)`

GetLabelProbeVerifiedAtOk returns a tuple with the LabelProbeVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelProbeVerifiedAt

`func (o *HandlersRunnerMachineDetailResponse) SetLabelProbeVerifiedAt(v string)`

SetLabelProbeVerifiedAt sets LabelProbeVerifiedAt field to given value.

### HasLabelProbeVerifiedAt

`func (o *HandlersRunnerMachineDetailResponse) HasLabelProbeVerifiedAt() bool`

HasLabelProbeVerifiedAt returns a boolean if a field has been set.

### GetLabels

`func (o *HandlersRunnerMachineDetailResponse) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HandlersRunnerMachineDetailResponse) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HandlersRunnerMachineDetailResponse) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HandlersRunnerMachineDetailResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastJobId

`func (o *HandlersRunnerMachineDetailResponse) GetLastJobId() string`

GetLastJobId returns the LastJobId field if non-nil, zero value otherwise.

### GetLastJobIdOk

`func (o *HandlersRunnerMachineDetailResponse) GetLastJobIdOk() (*string, bool)`

GetLastJobIdOk returns a tuple with the LastJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastJobId

`func (o *HandlersRunnerMachineDetailResponse) SetLastJobId(v string)`

SetLastJobId sets LastJobId field to given value.

### HasLastJobId

`func (o *HandlersRunnerMachineDetailResponse) HasLastJobId() bool`

HasLastJobId returns a boolean if a field has been set.

### GetLifecycle

`func (o *HandlersRunnerMachineDetailResponse) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *HandlersRunnerMachineDetailResponse) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *HandlersRunnerMachineDetailResponse) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *HandlersRunnerMachineDetailResponse) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetName

`func (o *HandlersRunnerMachineDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersRunnerMachineDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersRunnerMachineDetailResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersRunnerMachineDetailResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedsAttention

`func (o *HandlersRunnerMachineDetailResponse) GetNeedsAttention() bool`

GetNeedsAttention returns the NeedsAttention field if non-nil, zero value otherwise.

### GetNeedsAttentionOk

`func (o *HandlersRunnerMachineDetailResponse) GetNeedsAttentionOk() (*bool, bool)`

GetNeedsAttentionOk returns a tuple with the NeedsAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsAttention

`func (o *HandlersRunnerMachineDetailResponse) SetNeedsAttention(v bool)`

SetNeedsAttention sets NeedsAttention field to given value.

### HasNeedsAttention

`func (o *HandlersRunnerMachineDetailResponse) HasNeedsAttention() bool`

HasNeedsAttention returns a boolean if a field has been set.

### GetOs

`func (o *HandlersRunnerMachineDetailResponse) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *HandlersRunnerMachineDetailResponse) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *HandlersRunnerMachineDetailResponse) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *HandlersRunnerMachineDetailResponse) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOwner

`func (o *HandlersRunnerMachineDetailResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *HandlersRunnerMachineDetailResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *HandlersRunnerMachineDetailResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *HandlersRunnerMachineDetailResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnership

`func (o *HandlersRunnerMachineDetailResponse) GetOwnership() string`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *HandlersRunnerMachineDetailResponse) GetOwnershipOk() (*string, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *HandlersRunnerMachineDetailResponse) SetOwnership(v string)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *HandlersRunnerMachineDetailResponse) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.

### GetProfilesServed

`func (o *HandlersRunnerMachineDetailResponse) GetProfilesServed() []string`

GetProfilesServed returns the ProfilesServed field if non-nil, zero value otherwise.

### GetProfilesServedOk

`func (o *HandlersRunnerMachineDetailResponse) GetProfilesServedOk() (*[]string, bool)`

GetProfilesServedOk returns a tuple with the ProfilesServed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesServed

`func (o *HandlersRunnerMachineDetailResponse) SetProfilesServed(v []string)`

SetProfilesServed sets ProfilesServed field to given value.

### HasProfilesServed

`func (o *HandlersRunnerMachineDetailResponse) HasProfilesServed() bool`

HasProfilesServed returns a boolean if a field has been set.

### GetProviderId

`func (o *HandlersRunnerMachineDetailResponse) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *HandlersRunnerMachineDetailResponse) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *HandlersRunnerMachineDetailResponse) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *HandlersRunnerMachineDetailResponse) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetReachability

`func (o *HandlersRunnerMachineDetailResponse) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *HandlersRunnerMachineDetailResponse) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *HandlersRunnerMachineDetailResponse) SetReachability(v string)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *HandlersRunnerMachineDetailResponse) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetScope

`func (o *HandlersRunnerMachineDetailResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersRunnerMachineDetailResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersRunnerMachineDetailResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersRunnerMachineDetailResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetShape

`func (o *HandlersRunnerMachineDetailResponse) GetShape() HandlersRunnerMachineShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *HandlersRunnerMachineDetailResponse) GetShapeOk() (*HandlersRunnerMachineShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *HandlersRunnerMachineDetailResponse) SetShape(v HandlersRunnerMachineShape)`

SetShape sets Shape field to given value.

### HasShape

`func (o *HandlersRunnerMachineDetailResponse) HasShape() bool`

HasShape returns a boolean if a field has been set.

### GetTodayMinutes

`func (o *HandlersRunnerMachineDetailResponse) GetTodayMinutes() float32`

GetTodayMinutes returns the TodayMinutes field if non-nil, zero value otherwise.

### GetTodayMinutesOk

`func (o *HandlersRunnerMachineDetailResponse) GetTodayMinutesOk() (*float32, bool)`

GetTodayMinutesOk returns a tuple with the TodayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayMinutes

`func (o *HandlersRunnerMachineDetailResponse) SetTodayMinutes(v float32)`

SetTodayMinutes sets TodayMinutes field to given value.

### HasTodayMinutes

`func (o *HandlersRunnerMachineDetailResponse) HasTodayMinutes() bool`

HasTodayMinutes returns a boolean if a field has been set.

### GetTodayRuns

`func (o *HandlersRunnerMachineDetailResponse) GetTodayRuns() int32`

GetTodayRuns returns the TodayRuns field if non-nil, zero value otherwise.

### GetTodayRunsOk

`func (o *HandlersRunnerMachineDetailResponse) GetTodayRunsOk() (*int32, bool)`

GetTodayRunsOk returns a tuple with the TodayRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayRuns

`func (o *HandlersRunnerMachineDetailResponse) SetTodayRuns(v int32)`

SetTodayRuns sets TodayRuns field to given value.

### HasTodayRuns

`func (o *HandlersRunnerMachineDetailResponse) HasTodayRuns() bool`

HasTodayRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


