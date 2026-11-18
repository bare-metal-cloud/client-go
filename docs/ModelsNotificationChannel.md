# ModelsNotificationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled gates delivery: the sink resolves ENABLED channels only. Disabling is the operator&#39;s mute (a disabled channel gets nothing and hears nothing — no dead-channel notices under the fire-once law).  gorm bool default trap: with &#x60;default:true&#x60;, Create SKIPS the zero value, so a struct-literal insert with Enabled:false silently lands true. Flips go through explicit Updates only — the store creates channels enabled and never inserts a disabled row (test seeders must likewise update the column explicitly). | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Kind** | Pointer to **string** | Kind is the channel vocabulary: webhook | email (validated at the service layer; no DB CHECK constraint — the column is data-safe). | [optional] 
**Label** | Pointer to **string** | Label is the channel&#39;s user-chosen, NON-secret display name (epic fcp entry 6, decision D21): required at create, unique per org — the service layer validates before every write (a duplicate answers 409 naming the holder) and migration 000152&#39;s (org_id, label) unique index is the at-rest backstop. The label is the rule-form picker&#39;s primary text; the sealed mask stays secondary.  The default:&#39;&#39; is the explorer_grant precedent, NOT a real default: AutoMigrate runs BEFORE goose at boot, and a plain &#x60;not null&#x60; ADD COLUMN against a populated notification_channels table lands \&quot;column contains null values\&quot; — boot aborts before 000152 can backfill. The &#39;&#39; default lets the pre-backfill rows exist for the seconds between the two; the service layer refuses blank labels on every write, and 000152 fills the &#39;&#39; cells by created order. | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Target** | Pointer to **string** | Target is the delivery target: for webhook channels the SEALED URL blob (custody opens it only on the delivery path), for email channels the address. Never the webhook plaintext. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsNotificationChannel

`func NewModelsNotificationChannel() *ModelsNotificationChannel`

NewModelsNotificationChannel instantiates a new ModelsNotificationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsNotificationChannelWithDefaults

`func NewModelsNotificationChannelWithDefaults() *ModelsNotificationChannel`

NewModelsNotificationChannelWithDefaults instantiates a new ModelsNotificationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ModelsNotificationChannel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsNotificationChannel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsNotificationChannel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsNotificationChannel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelsNotificationChannel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelsNotificationChannel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelsNotificationChannel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelsNotificationChannel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *ModelsNotificationChannel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsNotificationChannel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsNotificationChannel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsNotificationChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ModelsNotificationChannel) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ModelsNotificationChannel) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ModelsNotificationChannel) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ModelsNotificationChannel) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *ModelsNotificationChannel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModelsNotificationChannel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModelsNotificationChannel) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ModelsNotificationChannel) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsNotificationChannel) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsNotificationChannel) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsNotificationChannel) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsNotificationChannel) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetTarget

`func (o *ModelsNotificationChannel) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ModelsNotificationChannel) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ModelsNotificationChannel) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ModelsNotificationChannel) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsNotificationChannel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsNotificationChannel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsNotificationChannel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsNotificationChannel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


