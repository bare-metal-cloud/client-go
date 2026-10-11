# HandlersSSHConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** | Optional description | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IpAddress** | Pointer to **string** | Connection details | [optional] 
**IsBastion** | Pointer to **bool** | Mark as bastion host | [optional] 
**IsMultipassHost** | Pointer to **bool** |  | [optional] 
**IsVerified** | Pointer to **bool** | Connection tested successfully | [optional] 
**LastTestedAt** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Connection metadata | [optional] 
**OrgId** | Pointer to **int32** | Tenant scope (ADR-062) | [optional] 
**ParentConnection** | Pointer to [**ModelsManagedSSHConnection**](ModelsManagedSSHConnection.md) |  | [optional] 
**ParentConnectionId** | Pointer to **int32** | Bastion/parent connection (nullable) If set, this connection requires jumping through the parent bastion host | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**PrivateKeyEncryptedDataId** | Pointer to **int32** | Encrypted SSH Key (Story 18-14: Refactored to use polymorphic encrypted_data table) SECURITY FIX: Previously stored plaintext - now encrypted! References the EncryptedData table where the SSH private key is stored encrypted | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UseAgentForwarding** | Pointer to **bool** | SSH Agent forwarding (when using bastion) If true, uses bastion&#39;s SSH agent instead of a separate key | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersSSHConnectionResponse

`func NewHandlersSSHConnectionResponse() *HandlersSSHConnectionResponse`

NewHandlersSSHConnectionResponse instantiates a new HandlersSSHConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersSSHConnectionResponseWithDefaults

`func NewHandlersSSHConnectionResponseWithDefaults() *HandlersSSHConnectionResponse`

NewHandlersSSHConnectionResponseWithDefaults instantiates a new HandlersSSHConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HandlersSSHConnectionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersSSHConnectionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersSSHConnectionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersSSHConnectionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *HandlersSSHConnectionResponse) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *HandlersSSHConnectionResponse) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *HandlersSSHConnectionResponse) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *HandlersSSHConnectionResponse) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDescription

`func (o *HandlersSSHConnectionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersSSHConnectionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersSSHConnectionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersSSHConnectionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *HandlersSSHConnectionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersSSHConnectionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersSSHConnectionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersSSHConnectionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *HandlersSSHConnectionResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *HandlersSSHConnectionResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *HandlersSSHConnectionResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *HandlersSSHConnectionResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIsBastion

`func (o *HandlersSSHConnectionResponse) GetIsBastion() bool`

GetIsBastion returns the IsBastion field if non-nil, zero value otherwise.

### GetIsBastionOk

`func (o *HandlersSSHConnectionResponse) GetIsBastionOk() (*bool, bool)`

GetIsBastionOk returns a tuple with the IsBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBastion

`func (o *HandlersSSHConnectionResponse) SetIsBastion(v bool)`

SetIsBastion sets IsBastion field to given value.

### HasIsBastion

`func (o *HandlersSSHConnectionResponse) HasIsBastion() bool`

HasIsBastion returns a boolean if a field has been set.

### GetIsMultipassHost

`func (o *HandlersSSHConnectionResponse) GetIsMultipassHost() bool`

GetIsMultipassHost returns the IsMultipassHost field if non-nil, zero value otherwise.

### GetIsMultipassHostOk

`func (o *HandlersSSHConnectionResponse) GetIsMultipassHostOk() (*bool, bool)`

GetIsMultipassHostOk returns a tuple with the IsMultipassHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultipassHost

`func (o *HandlersSSHConnectionResponse) SetIsMultipassHost(v bool)`

SetIsMultipassHost sets IsMultipassHost field to given value.

### HasIsMultipassHost

`func (o *HandlersSSHConnectionResponse) HasIsMultipassHost() bool`

HasIsMultipassHost returns a boolean if a field has been set.

### GetIsVerified

`func (o *HandlersSSHConnectionResponse) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *HandlersSSHConnectionResponse) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *HandlersSSHConnectionResponse) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *HandlersSSHConnectionResponse) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetLastTestedAt

`func (o *HandlersSSHConnectionResponse) GetLastTestedAt() string`

GetLastTestedAt returns the LastTestedAt field if non-nil, zero value otherwise.

### GetLastTestedAtOk

`func (o *HandlersSSHConnectionResponse) GetLastTestedAtOk() (*string, bool)`

GetLastTestedAtOk returns a tuple with the LastTestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTestedAt

`func (o *HandlersSSHConnectionResponse) SetLastTestedAt(v string)`

SetLastTestedAt sets LastTestedAt field to given value.

### HasLastTestedAt

`func (o *HandlersSSHConnectionResponse) HasLastTestedAt() bool`

HasLastTestedAt returns a boolean if a field has been set.

### GetName

`func (o *HandlersSSHConnectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersSSHConnectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersSSHConnectionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersSSHConnectionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *HandlersSSHConnectionResponse) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *HandlersSSHConnectionResponse) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *HandlersSSHConnectionResponse) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *HandlersSSHConnectionResponse) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetParentConnection

`func (o *HandlersSSHConnectionResponse) GetParentConnection() ModelsManagedSSHConnection`

GetParentConnection returns the ParentConnection field if non-nil, zero value otherwise.

### GetParentConnectionOk

`func (o *HandlersSSHConnectionResponse) GetParentConnectionOk() (*ModelsManagedSSHConnection, bool)`

GetParentConnectionOk returns a tuple with the ParentConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnection

`func (o *HandlersSSHConnectionResponse) SetParentConnection(v ModelsManagedSSHConnection)`

SetParentConnection sets ParentConnection field to given value.

### HasParentConnection

`func (o *HandlersSSHConnectionResponse) HasParentConnection() bool`

HasParentConnection returns a boolean if a field has been set.

### GetParentConnectionId

`func (o *HandlersSSHConnectionResponse) GetParentConnectionId() int32`

GetParentConnectionId returns the ParentConnectionId field if non-nil, zero value otherwise.

### GetParentConnectionIdOk

`func (o *HandlersSSHConnectionResponse) GetParentConnectionIdOk() (*int32, bool)`

GetParentConnectionIdOk returns a tuple with the ParentConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectionId

`func (o *HandlersSSHConnectionResponse) SetParentConnectionId(v int32)`

SetParentConnectionId sets ParentConnectionId field to given value.

### HasParentConnectionId

`func (o *HandlersSSHConnectionResponse) HasParentConnectionId() bool`

HasParentConnectionId returns a boolean if a field has been set.

### GetPort

`func (o *HandlersSSHConnectionResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlersSSHConnectionResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlersSSHConnectionResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HandlersSSHConnectionResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateKeyEncryptedDataId

`func (o *HandlersSSHConnectionResponse) GetPrivateKeyEncryptedDataId() int32`

GetPrivateKeyEncryptedDataId returns the PrivateKeyEncryptedDataId field if non-nil, zero value otherwise.

### GetPrivateKeyEncryptedDataIdOk

`func (o *HandlersSSHConnectionResponse) GetPrivateKeyEncryptedDataIdOk() (*int32, bool)`

GetPrivateKeyEncryptedDataIdOk returns a tuple with the PrivateKeyEncryptedDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEncryptedDataId

`func (o *HandlersSSHConnectionResponse) SetPrivateKeyEncryptedDataId(v int32)`

SetPrivateKeyEncryptedDataId sets PrivateKeyEncryptedDataId field to given value.

### HasPrivateKeyEncryptedDataId

`func (o *HandlersSSHConnectionResponse) HasPrivateKeyEncryptedDataId() bool`

HasPrivateKeyEncryptedDataId returns a boolean if a field has been set.

### GetProjectId

`func (o *HandlersSSHConnectionResponse) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HandlersSSHConnectionResponse) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HandlersSSHConnectionResponse) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *HandlersSSHConnectionResponse) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetScope

`func (o *HandlersSSHConnectionResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersSSHConnectionResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersSSHConnectionResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersSSHConnectionResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HandlersSSHConnectionResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HandlersSSHConnectionResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HandlersSSHConnectionResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HandlersSSHConnectionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseAgentForwarding

`func (o *HandlersSSHConnectionResponse) GetUseAgentForwarding() bool`

GetUseAgentForwarding returns the UseAgentForwarding field if non-nil, zero value otherwise.

### GetUseAgentForwardingOk

`func (o *HandlersSSHConnectionResponse) GetUseAgentForwardingOk() (*bool, bool)`

GetUseAgentForwardingOk returns a tuple with the UseAgentForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgentForwarding

`func (o *HandlersSSHConnectionResponse) SetUseAgentForwarding(v bool)`

SetUseAgentForwarding sets UseAgentForwarding field to given value.

### HasUseAgentForwarding

`func (o *HandlersSSHConnectionResponse) HasUseAgentForwarding() bool`

HasUseAgentForwarding returns a boolean if a field has been set.

### GetUsername

`func (o *HandlersSSHConnectionResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HandlersSSHConnectionResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HandlersSSHConnectionResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HandlersSSHConnectionResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


