# ModelsManagedSSHConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** | Optional description | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IpAddress** | Pointer to **string** | Connection details | [optional] 
**IsBastion** | Pointer to **bool** | Mark as bastion host | [optional] 
**IsVerified** | Pointer to **bool** | Connection tested successfully | [optional] 
**LastTestedAt** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Connection metadata | [optional] 
**OrgId** | Pointer to **int32** | Tenant scope (ADR-062) | [optional] 
**ParentConnection** | Pointer to [**ModelsManagedSSHConnection**](ModelsManagedSSHConnection.md) |  | [optional] 
**ParentConnectionId** | Pointer to **int32** | Bastion/parent connection (nullable) If set, this connection requires jumping through the parent bastion host | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**PrivateKeyEncryptedDataId** | Pointer to **int32** | Encrypted SSH Key (Story 18-14: Refactored to use polymorphic encrypted_data table) SECURITY FIX: Previously stored plaintext - now encrypted! References the EncryptedData table where the SSH private key is stored encrypted | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UseAgentForwarding** | Pointer to **bool** | SSH Agent forwarding (when using bastion) If true, uses bastion&#39;s SSH agent instead of a separate key | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsManagedSSHConnection

`func NewModelsManagedSSHConnection() *ModelsManagedSSHConnection`

NewModelsManagedSSHConnection instantiates a new ModelsManagedSSHConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsManagedSSHConnectionWithDefaults

`func NewModelsManagedSSHConnectionWithDefaults() *ModelsManagedSSHConnection`

NewModelsManagedSSHConnectionWithDefaults instantiates a new ModelsManagedSSHConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ModelsManagedSSHConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsManagedSSHConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsManagedSSHConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsManagedSSHConnection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ModelsManagedSSHConnection) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ModelsManagedSSHConnection) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ModelsManagedSSHConnection) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ModelsManagedSSHConnection) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDescription

`func (o *ModelsManagedSSHConnection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsManagedSSHConnection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsManagedSSHConnection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsManagedSSHConnection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ModelsManagedSSHConnection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsManagedSSHConnection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsManagedSSHConnection) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsManagedSSHConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *ModelsManagedSSHConnection) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ModelsManagedSSHConnection) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ModelsManagedSSHConnection) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ModelsManagedSSHConnection) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsBastion

`func (o *ModelsManagedSSHConnection) GetIsBastion() bool`

GetIsBastion returns the IsBastion field if non-nil, zero value otherwise.

### GetIsBastionOk

`func (o *ModelsManagedSSHConnection) GetIsBastionOk() (*bool, bool)`

GetIsBastionOk returns a tuple with the IsBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBastion

`func (o *ModelsManagedSSHConnection) SetIsBastion(v bool)`

SetIsBastion sets IsBastion field to given value.

### HasIsBastion

`func (o *ModelsManagedSSHConnection) HasIsBastion() bool`

HasIsBastion returns a boolean if a field has been set.

### GetIsVerified

`func (o *ModelsManagedSSHConnection) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *ModelsManagedSSHConnection) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *ModelsManagedSSHConnection) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *ModelsManagedSSHConnection) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetLastTestedAt

`func (o *ModelsManagedSSHConnection) GetLastTestedAt() string`

GetLastTestedAt returns the LastTestedAt field if non-nil, zero value otherwise.

### GetLastTestedAtOk

`func (o *ModelsManagedSSHConnection) GetLastTestedAtOk() (*string, bool)`

GetLastTestedAtOk returns a tuple with the LastTestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestedAt

`func (o *ModelsManagedSSHConnection) SetLastTestedAt(v string)`

SetLastTestedAt sets LastTestedAt field to given value.

### HasLastTestedAt

`func (o *ModelsManagedSSHConnection) HasLastTestedAt() bool`

HasLastTestedAt returns a boolean if a field has been set.

### GetName

`func (o *ModelsManagedSSHConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsManagedSSHConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsManagedSSHConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsManagedSSHConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *ModelsManagedSSHConnection) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ModelsManagedSSHConnection) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ModelsManagedSSHConnection) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ModelsManagedSSHConnection) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetParentConnection

`func (o *ModelsManagedSSHConnection) GetParentConnection() ModelsManagedSSHConnection`

GetParentConnection returns the ParentConnection field if non-nil, zero value otherwise.

### GetParentConnectionOk

`func (o *ModelsManagedSSHConnection) GetParentConnectionOk() (*ModelsManagedSSHConnection, bool)`

GetParentConnectionOk returns a tuple with the ParentConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnection

`func (o *ModelsManagedSSHConnection) SetParentConnection(v ModelsManagedSSHConnection)`

SetParentConnection sets ParentConnection field to given value.

### HasParentConnection

`func (o *ModelsManagedSSHConnection) HasParentConnection() bool`

HasParentConnection returns a boolean if a field has been set.

### GetParentConnectionId

`func (o *ModelsManagedSSHConnection) GetParentConnectionId() int32`

GetParentConnectionId returns the ParentConnectionId field if non-nil, zero value otherwise.

### GetParentConnectionIdOk

`func (o *ModelsManagedSSHConnection) GetParentConnectionIdOk() (*int32, bool)`

GetParentConnectionIdOk returns a tuple with the ParentConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectionId

`func (o *ModelsManagedSSHConnection) SetParentConnectionId(v int32)`

SetParentConnectionId sets ParentConnectionId field to given value.

### HasParentConnectionId

`func (o *ModelsManagedSSHConnection) HasParentConnectionId() bool`

HasParentConnectionId returns a boolean if a field has been set.

### GetPort

`func (o *ModelsManagedSSHConnection) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ModelsManagedSSHConnection) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ModelsManagedSSHConnection) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ModelsManagedSSHConnection) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateKeyEncryptedDataId

`func (o *ModelsManagedSSHConnection) GetPrivateKeyEncryptedDataId() int32`

GetPrivateKeyEncryptedDataId returns the PrivateKeyEncryptedDataId field if non-nil, zero value otherwise.

### GetPrivateKeyEncryptedDataIdOk

`func (o *ModelsManagedSSHConnection) GetPrivateKeyEncryptedDataIdOk() (*int32, bool)`

GetPrivateKeyEncryptedDataIdOk returns a tuple with the PrivateKeyEncryptedDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEncryptedDataId

`func (o *ModelsManagedSSHConnection) SetPrivateKeyEncryptedDataId(v int32)`

SetPrivateKeyEncryptedDataId sets PrivateKeyEncryptedDataId field to given value.

### HasPrivateKeyEncryptedDataId

`func (o *ModelsManagedSSHConnection) HasPrivateKeyEncryptedDataId() bool`

HasPrivateKeyEncryptedDataId returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelsManagedSSHConnection) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelsManagedSSHConnection) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelsManagedSSHConnection) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelsManagedSSHConnection) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsManagedSSHConnection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsManagedSSHConnection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsManagedSSHConnection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsManagedSSHConnection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseAgentForwarding

`func (o *ModelsManagedSSHConnection) GetUseAgentForwarding() bool`

GetUseAgentForwarding returns the UseAgentForwarding field if non-nil, zero value otherwise.

### GetUseAgentForwardingOk

`func (o *ModelsManagedSSHConnection) GetUseAgentForwardingOk() (*bool, bool)`

GetUseAgentForwardingOk returns a tuple with the UseAgentForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgentForwarding

`func (o *ModelsManagedSSHConnection) SetUseAgentForwarding(v bool)`

SetUseAgentForwarding sets UseAgentForwarding field to given value.

### HasUseAgentForwarding

`func (o *ModelsManagedSSHConnection) HasUseAgentForwarding() bool`

HasUseAgentForwarding returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsManagedSSHConnection) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsManagedSSHConnection) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsManagedSSHConnection) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsManagedSSHConnection) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


