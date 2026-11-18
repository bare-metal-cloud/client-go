# ModelsPersonGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DatabaseId** | Pointer to **int32** |  | [optional] 
**GrantedAt** | Pointer to **string** |  | [optional] 
**GrantedBy** | Pointer to **int32** | Nullable stamps lead (fieldalignment). | [optional] 
**Id** | Pointer to **int32** | Identity columns: the grant binds the resource&#39;s STABLE ID (the databases uint PK) + org + project + the named person. | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to **string** | Reason is the admin&#39;s why (08.3&#39;s visibility: the grant carries its intent alongside its attribution). | [optional] 
**RevokedAt** | Pointer to **string** | revoke stamp | [optional] 
**Shape** | Pointer to **string** | The grant vocabulary. Status and Shape are CHECK-constrained at the DB — migration 000125 adds ck_person_grants_status and ck_person_grants_shape via idempotent DO-block ALTERs (NOT inline CREATE TABLE constraints: AutoMigrate runs before goose and creates this table from these tags, which cannot express a CHECK — the documented 000040 trap). | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SuspendedAt** | Pointer to **string** | membership-loss stamp | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsPersonGrant

`func NewModelsPersonGrant() *ModelsPersonGrant`

NewModelsPersonGrant instantiates a new ModelsPersonGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsPersonGrantWithDefaults

`func NewModelsPersonGrantWithDefaults() *ModelsPersonGrant`

NewModelsPersonGrantWithDefaults instantiates a new ModelsPersonGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ModelsPersonGrant) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsPersonGrant) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsPersonGrant) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsPersonGrant) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDatabaseId

`func (o *ModelsPersonGrant) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *ModelsPersonGrant) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *ModelsPersonGrant) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *ModelsPersonGrant) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### GetGrantedAt

`func (o *ModelsPersonGrant) GetGrantedAt() string`

GetGrantedAt returns the GrantedAt field if non-nil, zero value otherwise.

### GetGrantedAtOk

`func (o *ModelsPersonGrant) GetGrantedAtOk() (*string, bool)`

GetGrantedAtOk returns a tuple with the GrantedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedAt

`func (o *ModelsPersonGrant) SetGrantedAt(v string)`

SetGrantedAt sets GrantedAt field to given value.

### HasGrantedAt

`func (o *ModelsPersonGrant) HasGrantedAt() bool`

HasGrantedAt returns a boolean if a field has been set.

### GetGrantedBy

`func (o *ModelsPersonGrant) GetGrantedBy() int32`

GetGrantedBy returns the GrantedBy field if non-nil, zero value otherwise.

### GetGrantedByOk

`func (o *ModelsPersonGrant) GetGrantedByOk() (*int32, bool)`

GetGrantedByOk returns a tuple with the GrantedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedBy

`func (o *ModelsPersonGrant) SetGrantedBy(v int32)`

SetGrantedBy sets GrantedBy field to given value.

### HasGrantedBy

`func (o *ModelsPersonGrant) HasGrantedBy() bool`

HasGrantedBy returns a boolean if a field has been set.

### GetId

`func (o *ModelsPersonGrant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsPersonGrant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsPersonGrant) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsPersonGrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsPersonGrant) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsPersonGrant) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsPersonGrant) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsPersonGrant) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsPersonGrant) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsPersonGrant) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsPersonGrant) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsPersonGrant) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReason

`func (o *ModelsPersonGrant) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ModelsPersonGrant) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ModelsPersonGrant) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ModelsPersonGrant) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRevokedAt

`func (o *ModelsPersonGrant) GetRevokedAt() string`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *ModelsPersonGrant) GetRevokedAtOk() (*string, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *ModelsPersonGrant) SetRevokedAt(v string)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *ModelsPersonGrant) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### GetShape

`func (o *ModelsPersonGrant) GetShape() string`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *ModelsPersonGrant) GetShapeOk() (*string, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *ModelsPersonGrant) SetShape(v string)`

SetShape sets Shape field to given value.

### HasShape

`func (o *ModelsPersonGrant) HasShape() bool`

HasShape returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsPersonGrant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsPersonGrant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsPersonGrant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsPersonGrant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuspendedAt

`func (o *ModelsPersonGrant) GetSuspendedAt() string`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *ModelsPersonGrant) GetSuspendedAtOk() (*string, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *ModelsPersonGrant) SetSuspendedAt(v string)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *ModelsPersonGrant) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsPersonGrant) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsPersonGrant) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsPersonGrant) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsPersonGrant) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *ModelsPersonGrant) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelsPersonGrant) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelsPersonGrant) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelsPersonGrant) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


