# CloneplanInstallFacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bundle** | Pointer to [**BundleRef**](BundleRef.md) | Bundle is the chosen revision&#39;s digest-pinned bundle ref (AC 2) — the replay Builds from THIS ref, never a re-resolved catalog entry. Nil when the unit has no replay revision. | [optional] 
**DataClass** | Pointer to **string** | DataClass is the install component&#39;s declared data class (AC 2) — what decides the replay&#39;s data behavior (stateless redeploys, stateful-derivable deploys fresh, stateful-authoritative bootstraps empty; OPERATIONS.md:255). DataClassUndeclared when the frozen state carries no declaration. A loose unit carries its single component&#39;s class; a stack unit&#39;s parent record carries none (the components carry theirs). | [optional] 
**DisplayName** | Pointer to **string** | DisplayName / PhysicalName are the record&#39;s display/physical-name facts, carried AS STORED (AC 2). A restore-diverged source&#39;s PhysicalName diverges from its InstallName permanently (the alias law) — the plan shows the divergence; story 134.3&#39;s identity law gives the CLONE a fresh PhysicalName and never copies this one. | [optional] 
**Durability** | Pointer to [**CloneplanDurabilityPlan**](CloneplanDurabilityPlan.md) | Durability is the durability plan (AC 6, RULED 2026-09-27): durability does NOT travel — the proposal is the platform default (off) and the source&#39;s policy rides as reference; the review may override per install. | [optional] 
**HooksDeclared** | Pointer to **[]string** | HooksDeclared are the install&#39;s DECLARED seed/sanitize hooks (AC 2; the bundle format&#39;s hook surface — bundle-declared per OPERATIONS.md:257). Empty today: the frozen bundle format declares no seed/sanitize hooks yet, and the plan says so rather than guessing one. | [optional] 
**Params** | Pointer to [**[]CloneplanParam**](CloneplanParam.md) | Params are the chosen revision&#39;s non-sensitive resolved params (AC 2), each with its portability classification (the Param vocabulary). Credential-class params are listed by key with their values stripped (AC 3). | [optional] 
**PhysicalName** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to [**EngineInstallRef**](EngineInstallRef.md) | Ref is the source install&#39;s identity (project/cluster/install). | [optional] 
**Replay** | Pointer to [**CloneplanReplaySource**](CloneplanReplaySource.md) | Replay records how the revision was chosen (ReplaySource). | [optional] 
**ReplayEmpty** | Pointer to **bool** | ReplayEmpty is the review-facing fact 134.2 replays by: a stateful-authoritative install with NO declared seed hook replays EMPTY, and the plan review says so — never a guess. | [optional] 
**ReplayRevision** | Pointer to **int32** | ReplayRevision is the proposed replayed revision number: the record&#39;s StableRevision, falling back to Current when no stable pointer exists (RULED 2026-09-27). Zero with Replay none. | [optional] 

## Methods

### NewCloneplanInstallFacts

`func NewCloneplanInstallFacts() *CloneplanInstallFacts`

NewCloneplanInstallFacts instantiates a new CloneplanInstallFacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanInstallFactsWithDefaults

`func NewCloneplanInstallFactsWithDefaults() *CloneplanInstallFacts`

NewCloneplanInstallFactsWithDefaults instantiates a new CloneplanInstallFacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundle

`func (o *CloneplanInstallFacts) GetBundle() BundleRef`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *CloneplanInstallFacts) GetBundleOk() (*BundleRef, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *CloneplanInstallFacts) SetBundle(v BundleRef)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *CloneplanInstallFacts) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetDataClass

`func (o *CloneplanInstallFacts) GetDataClass() string`

GetDataClass returns the DataClass field if non-nil, zero value otherwise.

### GetDataClassOk

`func (o *CloneplanInstallFacts) GetDataClassOk() (*string, bool)`

GetDataClassOk returns a tuple with the DataClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClass

`func (o *CloneplanInstallFacts) SetDataClass(v string)`

SetDataClass sets DataClass field to given value.

### HasDataClass

`func (o *CloneplanInstallFacts) HasDataClass() bool`

HasDataClass returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloneplanInstallFacts) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloneplanInstallFacts) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloneplanInstallFacts) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloneplanInstallFacts) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDurability

`func (o *CloneplanInstallFacts) GetDurability() CloneplanDurabilityPlan`

GetDurability returns the Durability field if non-nil, zero value otherwise.

### GetDurabilityOk

`func (o *CloneplanInstallFacts) GetDurabilityOk() (*CloneplanDurabilityPlan, bool)`

GetDurabilityOk returns a tuple with the Durability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurability

`func (o *CloneplanInstallFacts) SetDurability(v CloneplanDurabilityPlan)`

SetDurability sets Durability field to given value.

### HasDurability

`func (o *CloneplanInstallFacts) HasDurability() bool`

HasDurability returns a boolean if a field has been set.

### GetHooksDeclared

`func (o *CloneplanInstallFacts) GetHooksDeclared() []string`

GetHooksDeclared returns the HooksDeclared field if non-nil, zero value otherwise.

### GetHooksDeclaredOk

`func (o *CloneplanInstallFacts) GetHooksDeclaredOk() (*[]string, bool)`

GetHooksDeclaredOk returns a tuple with the HooksDeclared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooksDeclared

`func (o *CloneplanInstallFacts) SetHooksDeclared(v []string)`

SetHooksDeclared sets HooksDeclared field to given value.

### HasHooksDeclared

`func (o *CloneplanInstallFacts) HasHooksDeclared() bool`

HasHooksDeclared returns a boolean if a field has been set.

### GetParams

`func (o *CloneplanInstallFacts) GetParams() []CloneplanParam`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CloneplanInstallFacts) GetParamsOk() (*[]CloneplanParam, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CloneplanInstallFacts) SetParams(v []CloneplanParam)`

SetParams sets Params field to given value.

### HasParams

`func (o *CloneplanInstallFacts) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetPhysicalName

`func (o *CloneplanInstallFacts) GetPhysicalName() string`

GetPhysicalName returns the PhysicalName field if non-nil, zero value otherwise.

### GetPhysicalNameOk

`func (o *CloneplanInstallFacts) GetPhysicalNameOk() (*string, bool)`

GetPhysicalNameOk returns a tuple with the PhysicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalName

`func (o *CloneplanInstallFacts) SetPhysicalName(v string)`

SetPhysicalName sets PhysicalName field to given value.

### HasPhysicalName

`func (o *CloneplanInstallFacts) HasPhysicalName() bool`

HasPhysicalName returns a boolean if a field has been set.

### GetRef

`func (o *CloneplanInstallFacts) GetRef() EngineInstallRef`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CloneplanInstallFacts) GetRefOk() (*EngineInstallRef, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CloneplanInstallFacts) SetRef(v EngineInstallRef)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CloneplanInstallFacts) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetReplay

`func (o *CloneplanInstallFacts) GetReplay() CloneplanReplaySource`

GetReplay returns the Replay field if non-nil, zero value otherwise.

### GetReplayOk

`func (o *CloneplanInstallFacts) GetReplayOk() (*CloneplanReplaySource, bool)`

GetReplayOk returns a tuple with the Replay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplay

`func (o *CloneplanInstallFacts) SetReplay(v CloneplanReplaySource)`

SetReplay sets Replay field to given value.

### HasReplay

`func (o *CloneplanInstallFacts) HasReplay() bool`

HasReplay returns a boolean if a field has been set.

### GetReplayEmpty

`func (o *CloneplanInstallFacts) GetReplayEmpty() bool`

GetReplayEmpty returns the ReplayEmpty field if non-nil, zero value otherwise.

### GetReplayEmptyOk

`func (o *CloneplanInstallFacts) GetReplayEmptyOk() (*bool, bool)`

GetReplayEmptyOk returns a tuple with the ReplayEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayEmpty

`func (o *CloneplanInstallFacts) SetReplayEmpty(v bool)`

SetReplayEmpty sets ReplayEmpty field to given value.

### HasReplayEmpty

`func (o *CloneplanInstallFacts) HasReplayEmpty() bool`

HasReplayEmpty returns a boolean if a field has been set.

### GetReplayRevision

`func (o *CloneplanInstallFacts) GetReplayRevision() int32`

GetReplayRevision returns the ReplayRevision field if non-nil, zero value otherwise.

### GetReplayRevisionOk

`func (o *CloneplanInstallFacts) GetReplayRevisionOk() (*int32, bool)`

GetReplayRevisionOk returns a tuple with the ReplayRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayRevision

`func (o *CloneplanInstallFacts) SetReplayRevision(v int32)`

SetReplayRevision sets ReplayRevision field to given value.

### HasReplayRevision

`func (o *CloneplanInstallFacts) HasReplayRevision() bool`

HasReplayRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


