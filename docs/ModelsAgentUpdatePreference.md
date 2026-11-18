# ModelsAgentUpdatePreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoWindowHours** | Pointer to **int32** | AutoWindowHours: the rollout window the auto path&#39;s candidate rule reads (design §3.5: connected-or-window agents). Default 24. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt / UpdatedAt bookkeeping (gorm auto-touch). | [optional] 
**Id** | Pointer to **int32** | ID is the surrogate key (BIGSERIAL in migration 000170). | [optional] 
**Mode** | Pointer to **string** | Mode: manual (default) | auto. The story-5 worker enrolls only auto orgs; manual orgs update by explicit click only. | [optional] 
**OrgId** | Pointer to **int32** | OrgID: the owning organization; UNIQUE — the one-row-per-org law. | [optional] 
**PinnedVersion** | Pointer to **string** | PinnedVersion: the org&#39;s version pin (NULL &#x3D; unpinned). Rail-verified at write time; auto rollouts skip pinned orgs, manual clicks target the pin. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsAgentUpdatePreference

`func NewModelsAgentUpdatePreference() *ModelsAgentUpdatePreference`

NewModelsAgentUpdatePreference instantiates a new ModelsAgentUpdatePreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsAgentUpdatePreferenceWithDefaults

`func NewModelsAgentUpdatePreferenceWithDefaults() *ModelsAgentUpdatePreference`

NewModelsAgentUpdatePreferenceWithDefaults instantiates a new ModelsAgentUpdatePreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoWindowHours

`func (o *ModelsAgentUpdatePreference) GetAutoWindowHours() int32`

GetAutoWindowHours returns the AutoWindowHours field if non-nil, zero value otherwise.

### GetAutoWindowHoursOk

`func (o *ModelsAgentUpdatePreference) GetAutoWindowHoursOk() (*int32, bool)`

GetAutoWindowHoursOk returns a tuple with the AutoWindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWindowHours

`func (o *ModelsAgentUpdatePreference) SetAutoWindowHours(v int32)`

SetAutoWindowHours sets AutoWindowHours field to given value.

### HasAutoWindowHours

`func (o *ModelsAgentUpdatePreference) HasAutoWindowHours() bool`

HasAutoWindowHours returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsAgentUpdatePreference) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsAgentUpdatePreference) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsAgentUpdatePreference) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsAgentUpdatePreference) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ModelsAgentUpdatePreference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsAgentUpdatePreference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsAgentUpdatePreference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsAgentUpdatePreference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *ModelsAgentUpdatePreference) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ModelsAgentUpdatePreference) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ModelsAgentUpdatePreference) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ModelsAgentUpdatePreference) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsAgentUpdatePreference) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsAgentUpdatePreference) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsAgentUpdatePreference) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsAgentUpdatePreference) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPinnedVersion

`func (o *ModelsAgentUpdatePreference) GetPinnedVersion() string`

GetPinnedVersion returns the PinnedVersion field if non-nil, zero value otherwise.

### GetPinnedVersionOk

`func (o *ModelsAgentUpdatePreference) GetPinnedVersionOk() (*string, bool)`

GetPinnedVersionOk returns a tuple with the PinnedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedVersion

`func (o *ModelsAgentUpdatePreference) SetPinnedVersion(v string)`

SetPinnedVersion sets PinnedVersion field to given value.

### HasPinnedVersion

`func (o *ModelsAgentUpdatePreference) HasPinnedVersion() bool`

HasPinnedVersion returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsAgentUpdatePreference) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsAgentUpdatePreference) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsAgentUpdatePreference) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsAgentUpdatePreference) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


