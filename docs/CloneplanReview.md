# CloneplanReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecidedAt** | Pointer to **string** |  | [optional] 
**DecidedBy** | Pointer to **int32** |  | [optional] 
**Durability** | Pointer to [**[]CloneplanDurabilityOverride**](CloneplanDurabilityOverride.md) | Durability are the per-install durability overrides (story C, issue #4277): the clone proposes off; an override replaces the proposal, and an override onto a destination another install writes refuses (CheckDurabilityOverrides — the two-writer guard). | [optional] 
**Endpoints** | Pointer to **[]string** | Endpoints are the linked-external re-resolution approvals: the binding-used endpoint names the operator approved for re-resolution against the target&#39;s registrations. | [optional] 
**EnvLocal** | Pointer to [**[]CloneplanEnvLocalApproval**](CloneplanEnvLocalApproval.md) | EnvLocal are the env_local approvals (install + param key → the approved target value). | [optional] 
**Excluded** | Pointer to **[]string** | Excluded lists the plan units (by install name — a stack unit&#39;s name is its PARENT install) the review excludes from the clone. Every unit not named is accepted. | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloneplanReview

`func NewCloneplanReview() *CloneplanReview`

NewCloneplanReview instantiates a new CloneplanReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanReviewWithDefaults

`func NewCloneplanReviewWithDefaults() *CloneplanReview`

NewCloneplanReviewWithDefaults instantiates a new CloneplanReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecidedAt

`func (o *CloneplanReview) GetDecidedAt() string`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *CloneplanReview) GetDecidedAtOk() (*string, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *CloneplanReview) SetDecidedAt(v string)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *CloneplanReview) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetDecidedBy

`func (o *CloneplanReview) GetDecidedBy() int32`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *CloneplanReview) GetDecidedByOk() (*int32, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *CloneplanReview) SetDecidedBy(v int32)`

SetDecidedBy sets DecidedBy field to given value.

### HasDecidedBy

`func (o *CloneplanReview) HasDecidedBy() bool`

HasDecidedBy returns a boolean if a field has been set.

### GetDurability

`func (o *CloneplanReview) GetDurability() []CloneplanDurabilityOverride`

GetDurability returns the Durability field if non-nil, zero value otherwise.

### GetDurabilityOk

`func (o *CloneplanReview) GetDurabilityOk() (*[]CloneplanDurabilityOverride, bool)`

GetDurabilityOk returns a tuple with the Durability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurability

`func (o *CloneplanReview) SetDurability(v []CloneplanDurabilityOverride)`

SetDurability sets Durability field to given value.

### HasDurability

`func (o *CloneplanReview) HasDurability() bool`

HasDurability returns a boolean if a field has been set.

### GetEndpoints

`func (o *CloneplanReview) GetEndpoints() []string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *CloneplanReview) GetEndpointsOk() (*[]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *CloneplanReview) SetEndpoints(v []string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *CloneplanReview) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetEnvLocal

`func (o *CloneplanReview) GetEnvLocal() []CloneplanEnvLocalApproval`

GetEnvLocal returns the EnvLocal field if non-nil, zero value otherwise.

### GetEnvLocalOk

`func (o *CloneplanReview) GetEnvLocalOk() (*[]CloneplanEnvLocalApproval, bool)`

GetEnvLocalOk returns a tuple with the EnvLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvLocal

`func (o *CloneplanReview) SetEnvLocal(v []CloneplanEnvLocalApproval)`

SetEnvLocal sets EnvLocal field to given value.

### HasEnvLocal

`func (o *CloneplanReview) HasEnvLocal() bool`

HasEnvLocal returns a boolean if a field has been set.

### GetExcluded

`func (o *CloneplanReview) GetExcluded() []string`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *CloneplanReview) GetExcludedOk() (*[]string, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *CloneplanReview) SetExcluded(v []string)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *CloneplanReview) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.

### GetPlanId

`func (o *CloneplanReview) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CloneplanReview) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CloneplanReview) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *CloneplanReview) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *CloneplanReview) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CloneplanReview) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CloneplanReview) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CloneplanReview) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


