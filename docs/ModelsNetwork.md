# ModelsNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ModelsNetworkConfig**](ModelsNetworkConfig.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUser** | Pointer to [**ModelsUser**](ModelsUser.md) |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**DraftName** | Pointer to **string** | DraftName is the cluster name the wizard was heading toward when the network was created (paired with Purpose&#x3D;cluster-create). | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastReconcileAt** | Pointer to **string** | LastReconcileAt is the timestamp of the most recent reconciler touch. Used for observability + to detect reconciler stalls. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Purpose** | Pointer to [**ModelsNetworkPurpose**](ModelsNetworkPurpose.md) | Purpose is the durable cluster-create draft intent (124.17): \&quot;\&quot; for ordinary networks, \&quot;cluster-create\&quot; for wizard-created drafts. See the NetworkPurpose docs. GORM AutoMigrate (dev-stage policy) adds the column on next boot; 000055 keeps the goose path convergent. | [optional] 
**ReconcileAttempts** | Pointer to **int32** | ReconcileAttempts is incremented by the Story 101-12 Temporal reconciler every time it retries a stuck row. When it reaches MaxReconcileAttempts (10) the row flips to NetworkStatusErrorPermanent. GORM AutoMigrate (dev-stage policy) adds the column on next boot. | [optional] 
**Status** | Pointer to [**ModelsNetworkStatus**](ModelsNetworkStatus.md) |  | [optional] 
**Type** | Pointer to [**ModelsNetworkType**](ModelsNetworkType.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Vms** | Pointer to [**[]ModelsVirtualMachine**](ModelsVirtualMachine.md) | Relations | [optional] 

## Methods

### NewModelsNetwork

`func NewModelsNetwork() *ModelsNetwork`

NewModelsNetwork instantiates a new ModelsNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsNetworkWithDefaults

`func NewModelsNetworkWithDefaults() *ModelsNetwork`

NewModelsNetworkWithDefaults instantiates a new ModelsNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ModelsNetwork) GetConfig() ModelsNetworkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ModelsNetwork) GetConfigOk() (*ModelsNetworkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ModelsNetwork) SetConfig(v ModelsNetworkConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ModelsNetwork) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsNetwork) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsNetwork) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsNetwork) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsNetwork) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *ModelsNetwork) GetCreatedByUser() ModelsUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *ModelsNetwork) GetCreatedByUserOk() (*ModelsUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *ModelsNetwork) SetCreatedByUser(v ModelsUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *ModelsNetwork) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsNetwork) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsNetwork) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsNetwork) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsNetwork) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ModelsNetwork) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ModelsNetwork) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ModelsNetwork) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ModelsNetwork) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDraftName

`func (o *ModelsNetwork) GetDraftName() string`

GetDraftName returns the DraftName field if non-nil, zero value otherwise.

### GetDraftNameOk

`func (o *ModelsNetwork) GetDraftNameOk() (*string, bool)`

GetDraftNameOk returns a tuple with the DraftName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftName

`func (o *ModelsNetwork) SetDraftName(v string)`

SetDraftName sets DraftName field to given value.

### HasDraftName

`func (o *ModelsNetwork) HasDraftName() bool`

HasDraftName returns a boolean if a field has been set.

### GetId

`func (o *ModelsNetwork) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsNetwork) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsNetwork) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastReconcileAt

`func (o *ModelsNetwork) GetLastReconcileAt() string`

GetLastReconcileAt returns the LastReconcileAt field if non-nil, zero value otherwise.

### GetLastReconcileAtOk

`func (o *ModelsNetwork) GetLastReconcileAtOk() (*string, bool)`

GetLastReconcileAtOk returns a tuple with the LastReconcileAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReconcileAt

`func (o *ModelsNetwork) SetLastReconcileAt(v string)`

SetLastReconcileAt sets LastReconcileAt field to given value.

### HasLastReconcileAt

`func (o *ModelsNetwork) HasLastReconcileAt() bool`

HasLastReconcileAt returns a boolean if a field has been set.

### GetName

`func (o *ModelsNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsNetwork) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsNetwork) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsNetwork) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsNetwork) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsNetwork) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsNetwork) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsNetwork) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsNetwork) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPurpose

`func (o *ModelsNetwork) GetPurpose() ModelsNetworkPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ModelsNetwork) GetPurposeOk() (*ModelsNetworkPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ModelsNetwork) SetPurpose(v ModelsNetworkPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *ModelsNetwork) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetReconcileAttempts

`func (o *ModelsNetwork) GetReconcileAttempts() int32`

GetReconcileAttempts returns the ReconcileAttempts field if non-nil, zero value otherwise.

### GetReconcileAttemptsOk

`func (o *ModelsNetwork) GetReconcileAttemptsOk() (*int32, bool)`

GetReconcileAttemptsOk returns a tuple with the ReconcileAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconcileAttempts

`func (o *ModelsNetwork) SetReconcileAttempts(v int32)`

SetReconcileAttempts sets ReconcileAttempts field to given value.

### HasReconcileAttempts

`func (o *ModelsNetwork) HasReconcileAttempts() bool`

HasReconcileAttempts returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsNetwork) GetStatus() ModelsNetworkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsNetwork) GetStatusOk() (*ModelsNetworkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsNetwork) SetStatus(v ModelsNetworkStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsNetwork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ModelsNetwork) GetType() ModelsNetworkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelsNetwork) GetTypeOk() (*ModelsNetworkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelsNetwork) SetType(v ModelsNetworkType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelsNetwork) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsNetwork) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsNetwork) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsNetwork) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsNetwork) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVms

`func (o *ModelsNetwork) GetVms() []ModelsVirtualMachine`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *ModelsNetwork) GetVmsOk() (*[]ModelsVirtualMachine, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *ModelsNetwork) SetVms(v []ModelsVirtualMachine)`

SetVms sets Vms field to given value.

### HasVms

`func (o *ModelsNetwork) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


