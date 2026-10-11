# ModelsMultipassHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** | Optional description | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsLocal** | Pointer to **bool** | true for local machine, false for remote | [optional] 
**IsVerified** | Pointer to **bool** | Status and verification | [optional] 
**LastTestedAt** | Pointer to **string** |  | [optional] 
**MultipassVersion** | Pointer to **string** | Multipass version info | [optional] 
**Name** | Pointer to **string** | Host identification | [optional] 
**OrgId** | Pointer to **int32** | Tenant scope (ADR-062) | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**SshConnection** | Pointer to [**ModelsManagedSSHConnection**](ModelsManagedSSHConnection.md) |  | [optional] 
**SshConnectionId** | Pointer to **int32** | Reference to ManagedSSHConnection (the jump host entry) | [optional] 
**Status** | Pointer to **string** | pending, active, error | [optional] 
**StatusReason** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsMultipassHost

`func NewModelsMultipassHost() *ModelsMultipassHost`

NewModelsMultipassHost instantiates a new ModelsMultipassHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsMultipassHostWithDefaults

`func NewModelsMultipassHostWithDefaults() *ModelsMultipassHost`

NewModelsMultipassHostWithDefaults instantiates a new ModelsMultipassHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ModelsMultipassHost) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsMultipassHost) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsMultipassHost) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsMultipassHost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsMultipassHost) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsMultipassHost) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsMultipassHost) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsMultipassHost) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsMultipassHost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsMultipassHost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsMultipassHost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsMultipassHost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ModelsMultipassHost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsMultipassHost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsMultipassHost) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsMultipassHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsLocal

`func (o *ModelsMultipassHost) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *ModelsMultipassHost) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *ModelsMultipassHost) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *ModelsMultipassHost) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetIsVerified

`func (o *ModelsMultipassHost) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *ModelsMultipassHost) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *ModelsMultipassHost) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *ModelsMultipassHost) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetLastTestedAt

`func (o *ModelsMultipassHost) GetLastTestedAt() string`

GetLastTestedAt returns the LastTestedAt field if non-nil, zero value otherwise.

### GetLastTestedAtOk

`func (o *ModelsMultipassHost) GetLastTestedAtOk() (*string, bool)`

GetLastTestedAtOk returns a tuple with the LastTestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestedAt

`func (o *ModelsMultipassHost) SetLastTestedAt(v string)`

SetLastTestedAt sets LastTestedAt field to given value.

### HasLastTestedAt

`func (o *ModelsMultipassHost) HasLastTestedAt() bool`

HasLastTestedAt returns a boolean if a field has been set.

### GetMultipassVersion

`func (o *ModelsMultipassHost) GetMultipassVersion() string`

GetMultipassVersion returns the MultipassVersion field if non-nil, zero value otherwise.

### GetMultipassVersionOk

`func (o *ModelsMultipassHost) GetMultipassVersionOk() (*string, bool)`

GetMultipassVersionOk returns a tuple with the MultipassVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipassVersion

`func (o *ModelsMultipassHost) SetMultipassVersion(v string)`

SetMultipassVersion sets MultipassVersion field to given value.

### HasMultipassVersion

`func (o *ModelsMultipassHost) HasMultipassVersion() bool`

HasMultipassVersion returns a boolean if a field has been set.

### GetName

`func (o *ModelsMultipassHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsMultipassHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsMultipassHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsMultipassHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsMultipassHost) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsMultipassHost) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsMultipassHost) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsMultipassHost) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsMultipassHost) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsMultipassHost) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsMultipassHost) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsMultipassHost) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSshConnection

`func (o *ModelsMultipassHost) GetSshConnection() ModelsManagedSSHConnection`

GetSshConnection returns the SshConnection field if non-nil, zero value otherwise.

### GetSshConnectionOk

`func (o *ModelsMultipassHost) GetSshConnectionOk() (*ModelsManagedSSHConnection, bool)`

GetSshConnectionOk returns a tuple with the SshConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConnection

`func (o *ModelsMultipassHost) SetSshConnection(v ModelsManagedSSHConnection)`

SetSshConnection sets SshConnection field to given value.

### HasSshConnection

`func (o *ModelsMultipassHost) HasSshConnection() bool`

HasSshConnection returns a boolean if a field has been set.

### GetSshConnectionId

`func (o *ModelsMultipassHost) GetSshConnectionId() int32`

GetSshConnectionId returns the SshConnectionId field if non-nil, zero value otherwise.

### GetSshConnectionIdOk

`func (o *ModelsMultipassHost) GetSshConnectionIdOk() (*int32, bool)`

GetSshConnectionIdOk returns a tuple with the SshConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConnectionId

`func (o *ModelsMultipassHost) SetSshConnectionId(v int32)`

SetSshConnectionId sets SshConnectionId field to given value.

### HasSshConnectionId

`func (o *ModelsMultipassHost) HasSshConnectionId() bool`

HasSshConnectionId returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsMultipassHost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsMultipassHost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsMultipassHost) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsMultipassHost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *ModelsMultipassHost) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *ModelsMultipassHost) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *ModelsMultipassHost) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *ModelsMultipassHost) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsMultipassHost) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsMultipassHost) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsMultipassHost) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsMultipassHost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


