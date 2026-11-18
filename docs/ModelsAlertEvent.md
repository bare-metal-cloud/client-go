# ModelsAlertEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledged** | Pointer to **bool** |  | [optional] 
**AcknowledgedAt** | Pointer to **string** | AcknowledgedAt/By and ResolvedAt/By record who acted and when (BAR-496). Acknowledge also sets the Acknowledged bool above, which stays the read the firing list filters on. | [optional] 
**AcknowledgedBy** | Pointer to **int32** |  | [optional] 
**BreachedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DedupKey** | Pointer to **string** | DedupKey names one alert occurrence. A second insert with the same (org, key) is a no-op, so an alert raised by a repeating sweep is stored and delivered once. Missed-backup alerts use \&quot;backup-missed:&lt;cluster&gt;:&lt;lane&gt;:&lt;target&gt;:&lt;window start&gt;\&quot; (BAR-496). NULL for every other alert, so the partial index ignores them. | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**ObservedValue** | Pointer to **float32** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ResolvedAt** | Pointer to **string** |  | [optional] 
**ResolvedBy** | Pointer to **int32** |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **int32** |  | [optional] 
**TargetKind** | Pointer to **string** | TargetKind/TargetID/TargetName mirror the firing rule&#39;s target (AlertTarget*). Pre-158.3 rows read as vm_threshold with target_id 0 — their vm_id is the subject, unchanged. | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**ThresholdValue** | Pointer to **float32** |  | [optional] 
**VmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsAlertEvent

`func NewModelsAlertEvent() *ModelsAlertEvent`

NewModelsAlertEvent instantiates a new ModelsAlertEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAlertEventWithDefaults

`func NewModelsAlertEventWithDefaults() *ModelsAlertEvent`

NewModelsAlertEventWithDefaults instantiates a new ModelsAlertEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledged

`func (o *ModelsAlertEvent) GetAcknowledged() bool`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *ModelsAlertEvent) GetAcknowledgedOk() (*bool, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *ModelsAlertEvent) SetAcknowledged(v bool)`

SetAcknowledged sets Acknowledged field to given value.

### HasAcknowledged

`func (o *ModelsAlertEvent) HasAcknowledged() bool`

HasAcknowledged returns a boolean if a field has been set.

### GetAcknowledgedAt

`func (o *ModelsAlertEvent) GetAcknowledgedAt() string`

GetAcknowledgedAt returns the AcknowledgedAt field if non-nil, zero value otherwise.

### GetAcknowledgedAtOk

`func (o *ModelsAlertEvent) GetAcknowledgedAtOk() (*string, bool)`

GetAcknowledgedAtOk returns a tuple with the AcknowledgedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedAt

`func (o *ModelsAlertEvent) SetAcknowledgedAt(v string)`

SetAcknowledgedAt sets AcknowledgedAt field to given value.

### HasAcknowledgedAt

`func (o *ModelsAlertEvent) HasAcknowledgedAt() bool`

HasAcknowledgedAt returns a boolean if a field has been set.

### GetAcknowledgedBy

`func (o *ModelsAlertEvent) GetAcknowledgedBy() int32`

GetAcknowledgedBy returns the AcknowledgedBy field if non-nil, zero value otherwise.

### GetAcknowledgedByOk

`func (o *ModelsAlertEvent) GetAcknowledgedByOk() (*int32, bool)`

GetAcknowledgedByOk returns a tuple with the AcknowledgedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedBy

`func (o *ModelsAlertEvent) SetAcknowledgedBy(v int32)`

SetAcknowledgedBy sets AcknowledgedBy field to given value.

### HasAcknowledgedBy

`func (o *ModelsAlertEvent) HasAcknowledgedBy() bool`

HasAcknowledgedBy returns a boolean if a field has been set.

### GetBreachedAt

`func (o *ModelsAlertEvent) GetBreachedAt() string`

GetBreachedAt returns the BreachedAt field if non-nil, zero value otherwise.

### GetBreachedAtOk

`func (o *ModelsAlertEvent) GetBreachedAtOk() (*string, bool)`

GetBreachedAtOk returns a tuple with the BreachedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreachedAt

`func (o *ModelsAlertEvent) SetBreachedAt(v string)`

SetBreachedAt sets BreachedAt field to given value.

### HasBreachedAt

`func (o *ModelsAlertEvent) HasBreachedAt() bool`

HasBreachedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsAlertEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsAlertEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsAlertEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsAlertEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDedupKey

`func (o *ModelsAlertEvent) GetDedupKey() string`

GetDedupKey returns the DedupKey field if non-nil, zero value otherwise.

### GetDedupKeyOk

`func (o *ModelsAlertEvent) GetDedupKeyOk() (*string, bool)`

GetDedupKeyOk returns a tuple with the DedupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupKey

`func (o *ModelsAlertEvent) SetDedupKey(v string)`

SetDedupKey sets DedupKey field to given value.

### HasDedupKey

`func (o *ModelsAlertEvent) HasDedupKey() bool`

HasDedupKey returns a boolean if a field has been set.

### GetId

`func (o *ModelsAlertEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsAlertEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsAlertEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsAlertEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ModelsAlertEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelsAlertEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelsAlertEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelsAlertEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetric

`func (o *ModelsAlertEvent) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ModelsAlertEvent) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ModelsAlertEvent) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ModelsAlertEvent) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetObservedValue

`func (o *ModelsAlertEvent) GetObservedValue() float32`

GetObservedValue returns the ObservedValue field if non-nil, zero value otherwise.

### GetObservedValueOk

`func (o *ModelsAlertEvent) GetObservedValueOk() (*float32, bool)`

GetObservedValueOk returns a tuple with the ObservedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedValue

`func (o *ModelsAlertEvent) SetObservedValue(v float32)`

SetObservedValue sets ObservedValue field to given value.

### HasObservedValue

`func (o *ModelsAlertEvent) HasObservedValue() bool`

HasObservedValue returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsAlertEvent) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsAlertEvent) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsAlertEvent) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsAlertEvent) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsAlertEvent) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsAlertEvent) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsAlertEvent) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsAlertEvent) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetResolvedAt

`func (o *ModelsAlertEvent) GetResolvedAt() string`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *ModelsAlertEvent) GetResolvedAtOk() (*string, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *ModelsAlertEvent) SetResolvedAt(v string)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *ModelsAlertEvent) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetResolvedBy

`func (o *ModelsAlertEvent) GetResolvedBy() int32`

GetResolvedBy returns the ResolvedBy field if non-nil, zero value otherwise.

### GetResolvedByOk

`func (o *ModelsAlertEvent) GetResolvedByOk() (*int32, bool)`

GetResolvedByOk returns a tuple with the ResolvedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedBy

`func (o *ModelsAlertEvent) SetResolvedBy(v int32)`

SetResolvedBy sets ResolvedBy field to given value.

### HasResolvedBy

`func (o *ModelsAlertEvent) HasResolvedBy() bool`

HasResolvedBy returns a boolean if a field has been set.

### GetRuleId

`func (o *ModelsAlertEvent) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ModelsAlertEvent) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ModelsAlertEvent) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ModelsAlertEvent) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetTargetId

`func (o *ModelsAlertEvent) GetTargetId() int32`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ModelsAlertEvent) GetTargetIdOk() (*int32, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ModelsAlertEvent) SetTargetId(v int32)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ModelsAlertEvent) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetKind

`func (o *ModelsAlertEvent) GetTargetKind() string`

GetTargetKind returns the TargetKind field if non-nil, zero value otherwise.

### GetTargetKindOk

`func (o *ModelsAlertEvent) GetTargetKindOk() (*string, bool)`

GetTargetKindOk returns a tuple with the TargetKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKind

`func (o *ModelsAlertEvent) SetTargetKind(v string)`

SetTargetKind sets TargetKind field to given value.

### HasTargetKind

`func (o *ModelsAlertEvent) HasTargetKind() bool`

HasTargetKind returns a boolean if a field has been set.

### GetTargetName

`func (o *ModelsAlertEvent) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ModelsAlertEvent) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ModelsAlertEvent) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ModelsAlertEvent) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetThresholdValue

`func (o *ModelsAlertEvent) GetThresholdValue() float32`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *ModelsAlertEvent) GetThresholdValueOk() (*float32, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *ModelsAlertEvent) SetThresholdValue(v float32)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *ModelsAlertEvent) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.

### GetVmId

`func (o *ModelsAlertEvent) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *ModelsAlertEvent) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *ModelsAlertEvent) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *ModelsAlertEvent) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


