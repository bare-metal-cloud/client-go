# ModelsAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to **string** | Condition is the watched condition (IsValidAlertCondition&#39;s per-kind vocabulary). Empty for vm_threshold rules. | [optional] 
**CooldownSeconds** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ForDurationSeconds** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Scope** | Pointer to **string** | Scope says which resources a storage_volume rule covers (AlertScope*). Defaults to resource; other kinds ignore it. | [optional] 
**Severity** | Pointer to **string** | Severity is warning or critical (AlertSeverity*). It rides every fired event and the notification text. Defaults to warning. | [optional] 
**Status** | Pointer to **string** | Status is the evaluator-owned rule lifecycle surfaced in CRUD (AlertRuleStatus*). CRUD writes only Active; the sweep owns the target_missing flip (and never flips it back except through rule update, which re-arms the rule). | [optional] 
**TargetId** | Pointer to **int32** | TargetID is the target resource id for the non-VM kinds (today: the cluster id for cluster_health and stack_component). Zero for vm_threshold and platform_signal rules. | [optional] 
**TargetKind** | Pointer to **string** | TargetKind names what the rule watches (AlertTarget*). Defaults to vm_threshold so pre-158.3 rows (and payloads that omit it) keep their exact existing meaning. | [optional] 
**TargetName** | Pointer to **string** | TargetName names the sub-target: the stack component name, or the platform_sensors sensor key. Empty for vm_threshold and cluster_health. | [optional] 
**Threshold** | Pointer to **float32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsAlertRule

`func NewModelsAlertRule() *ModelsAlertRule`

NewModelsAlertRule instantiates a new ModelsAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAlertRuleWithDefaults

`func NewModelsAlertRuleWithDefaults() *ModelsAlertRule`

NewModelsAlertRuleWithDefaults instantiates a new ModelsAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *ModelsAlertRule) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *ModelsAlertRule) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *ModelsAlertRule) SetComparison(v string)`

SetComparison sets Comparison field to given value.

### HasComparison

`func (o *ModelsAlertRule) HasComparison() bool`

HasComparison returns a boolean if a field has been set.

### GetCondition

`func (o *ModelsAlertRule) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ModelsAlertRule) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ModelsAlertRule) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ModelsAlertRule) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetCooldownSeconds

`func (o *ModelsAlertRule) GetCooldownSeconds() int32`

GetCooldownSeconds returns the CooldownSeconds field if non-nil, zero value otherwise.

### GetCooldownSecondsOk

`func (o *ModelsAlertRule) GetCooldownSecondsOk() (*int32, bool)`

GetCooldownSecondsOk returns a tuple with the CooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownSeconds

`func (o *ModelsAlertRule) SetCooldownSeconds(v int32)`

SetCooldownSeconds sets CooldownSeconds field to given value.

### HasCooldownSeconds

`func (o *ModelsAlertRule) HasCooldownSeconds() bool`

HasCooldownSeconds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsAlertRule) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsAlertRule) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsAlertRule) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsAlertRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsAlertRule) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsAlertRule) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsAlertRule) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsAlertRule) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelsAlertRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelsAlertRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelsAlertRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelsAlertRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForDurationSeconds

`func (o *ModelsAlertRule) GetForDurationSeconds() int32`

GetForDurationSeconds returns the ForDurationSeconds field if non-nil, zero value otherwise.

### GetForDurationSecondsOk

`func (o *ModelsAlertRule) GetForDurationSecondsOk() (*int32, bool)`

GetForDurationSecondsOk returns a tuple with the ForDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForDurationSeconds

`func (o *ModelsAlertRule) SetForDurationSeconds(v int32)`

SetForDurationSeconds sets ForDurationSeconds field to given value.

### HasForDurationSeconds

`func (o *ModelsAlertRule) HasForDurationSeconds() bool`

HasForDurationSeconds returns a boolean if a field has been set.

### GetId

`func (o *ModelsAlertRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsAlertRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsAlertRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsAlertRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetric

`func (o *ModelsAlertRule) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ModelsAlertRule) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ModelsAlertRule) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ModelsAlertRule) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsAlertRule) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsAlertRule) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsAlertRule) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsAlertRule) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsAlertRule) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsAlertRule) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsAlertRule) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsAlertRule) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetScope

`func (o *ModelsAlertRule) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ModelsAlertRule) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ModelsAlertRule) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ModelsAlertRule) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSeverity

`func (o *ModelsAlertRule) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelsAlertRule) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelsAlertRule) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ModelsAlertRule) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsAlertRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsAlertRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsAlertRule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsAlertRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetId

`func (o *ModelsAlertRule) GetTargetId() int32`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ModelsAlertRule) GetTargetIdOk() (*int32, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ModelsAlertRule) SetTargetId(v int32)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ModelsAlertRule) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetKind

`func (o *ModelsAlertRule) GetTargetKind() string`

GetTargetKind returns the TargetKind field if non-nil, zero value otherwise.

### GetTargetKindOk

`func (o *ModelsAlertRule) GetTargetKindOk() (*string, bool)`

GetTargetKindOk returns a tuple with the TargetKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKind

`func (o *ModelsAlertRule) SetTargetKind(v string)`

SetTargetKind sets TargetKind field to given value.

### HasTargetKind

`func (o *ModelsAlertRule) HasTargetKind() bool`

HasTargetKind returns a boolean if a field has been set.

### GetTargetName

`func (o *ModelsAlertRule) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ModelsAlertRule) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ModelsAlertRule) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ModelsAlertRule) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetThreshold

`func (o *ModelsAlertRule) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ModelsAlertRule) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ModelsAlertRule) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ModelsAlertRule) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsAlertRule) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsAlertRule) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsAlertRule) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsAlertRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVmId

`func (o *ModelsAlertRule) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *ModelsAlertRule) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *ModelsAlertRule) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *ModelsAlertRule) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


