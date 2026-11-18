# EngineDurabilityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveCadence** | Pointer to **string** | EffectiveCadence is the cadence as floored per node tier (STORAGE.md §3); recomputed on facts revision. The RPO badge displays THIS, never the requested one. | [optional] 
**Mode** | Pointer to **string** | off | scheduled | continuous | [optional] 
**PolicyGen** | Pointer to **int32** | PolicyGen is the separate CAS token for durability edits: inert edits bump THIS, never Generation, so a notify-grade change must not fence an in-flight upgrade into re-quiesce + second backup. | [optional] 
**RecoveryWindow** | Pointer to **string** | RecoveryWindow is the panel-configurable PITR window (story 132.3a, the epic-132 ruling): how far back point-in-time recovery promises, whole days (\&quot;7d\&quot;), default DefaultRecoveryWindow. Additive field with omitempty: policies written before 132.3a decode with it empty and re-marshal byte-identically. Inert at the record level — the effective bound is always the oldest surviving VERIFIED base, which retention (the preset) keeps alive. | [optional] 
**RetentionPreset** | Pointer to **string** | RetentionPreset is one of the named presets only — free-form per-object retention is illegal under the dependency-graph law (STORAGE.md §6). | [optional] 
**Schedule** | Pointer to **string** | cron (scheduled mode) | [optional] 
**SyncCadence** | Pointer to **string** | continuous file-sync cadence REQUESTED (e.g. \&quot;5m\&quot;) | [optional] 
**TargetDestination** | Pointer to **string** | S3 destination ref (user&#39;s bucket) | [optional] 

## Methods

### NewEngineDurabilityPolicy

`func NewEngineDurabilityPolicy() *EngineDurabilityPolicy`

NewEngineDurabilityPolicy instantiates a new EngineDurabilityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineDurabilityPolicyWithDefaults

`func NewEngineDurabilityPolicyWithDefaults() *EngineDurabilityPolicy`

NewEngineDurabilityPolicyWithDefaults instantiates a new EngineDurabilityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveCadence

`func (o *EngineDurabilityPolicy) GetEffectiveCadence() string`

GetEffectiveCadence returns the EffectiveCadence field if non-nil, zero value otherwise.

### GetEffectiveCadenceOk

`func (o *EngineDurabilityPolicy) GetEffectiveCadenceOk() (*string, bool)`

GetEffectiveCadenceOk returns a tuple with the EffectiveCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCadence

`func (o *EngineDurabilityPolicy) SetEffectiveCadence(v string)`

SetEffectiveCadence sets EffectiveCadence field to given value.

### HasEffectiveCadence

`func (o *EngineDurabilityPolicy) HasEffectiveCadence() bool`

HasEffectiveCadence returns a boolean if a field has been set.

### GetMode

`func (o *EngineDurabilityPolicy) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EngineDurabilityPolicy) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EngineDurabilityPolicy) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EngineDurabilityPolicy) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPolicyGen

`func (o *EngineDurabilityPolicy) GetPolicyGen() int32`

GetPolicyGen returns the PolicyGen field if non-nil, zero value otherwise.

### GetPolicyGenOk

`func (o *EngineDurabilityPolicy) GetPolicyGenOk() (*int32, bool)`

GetPolicyGenOk returns a tuple with the PolicyGen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGen

`func (o *EngineDurabilityPolicy) SetPolicyGen(v int32)`

SetPolicyGen sets PolicyGen field to given value.

### HasPolicyGen

`func (o *EngineDurabilityPolicy) HasPolicyGen() bool`

HasPolicyGen returns a boolean if a field has been set.

### GetRecoveryWindow

`func (o *EngineDurabilityPolicy) GetRecoveryWindow() string`

GetRecoveryWindow returns the RecoveryWindow field if non-nil, zero value otherwise.

### GetRecoveryWindowOk

`func (o *EngineDurabilityPolicy) GetRecoveryWindowOk() (*string, bool)`

GetRecoveryWindowOk returns a tuple with the RecoveryWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryWindow

`func (o *EngineDurabilityPolicy) SetRecoveryWindow(v string)`

SetRecoveryWindow sets RecoveryWindow field to given value.

### HasRecoveryWindow

`func (o *EngineDurabilityPolicy) HasRecoveryWindow() bool`

HasRecoveryWindow returns a boolean if a field has been set.

### GetRetentionPreset

`func (o *EngineDurabilityPolicy) GetRetentionPreset() string`

GetRetentionPreset returns the RetentionPreset field if non-nil, zero value otherwise.

### GetRetentionPresetOk

`func (o *EngineDurabilityPolicy) GetRetentionPresetOk() (*string, bool)`

GetRetentionPresetOk returns a tuple with the RetentionPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPreset

`func (o *EngineDurabilityPolicy) SetRetentionPreset(v string)`

SetRetentionPreset sets RetentionPreset field to given value.

### HasRetentionPreset

`func (o *EngineDurabilityPolicy) HasRetentionPreset() bool`

HasRetentionPreset returns a boolean if a field has been set.

### GetSchedule

`func (o *EngineDurabilityPolicy) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *EngineDurabilityPolicy) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *EngineDurabilityPolicy) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *EngineDurabilityPolicy) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSyncCadence

`func (o *EngineDurabilityPolicy) GetSyncCadence() string`

GetSyncCadence returns the SyncCadence field if non-nil, zero value otherwise.

### GetSyncCadenceOk

`func (o *EngineDurabilityPolicy) GetSyncCadenceOk() (*string, bool)`

GetSyncCadenceOk returns a tuple with the SyncCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCadence

`func (o *EngineDurabilityPolicy) SetSyncCadence(v string)`

SetSyncCadence sets SyncCadence field to given value.

### HasSyncCadence

`func (o *EngineDurabilityPolicy) HasSyncCadence() bool`

HasSyncCadence returns a boolean if a field has been set.

### GetTargetDestination

`func (o *EngineDurabilityPolicy) GetTargetDestination() string`

GetTargetDestination returns the TargetDestination field if non-nil, zero value otherwise.

### GetTargetDestinationOk

`func (o *EngineDurabilityPolicy) GetTargetDestinationOk() (*string, bool)`

GetTargetDestinationOk returns a tuple with the TargetDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDestination

`func (o *EngineDurabilityPolicy) SetTargetDestination(v string)`

SetTargetDestination sets TargetDestination field to given value.

### HasTargetDestination

`func (o *EngineDurabilityPolicy) HasTargetDestination() bool`

HasTargetDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


