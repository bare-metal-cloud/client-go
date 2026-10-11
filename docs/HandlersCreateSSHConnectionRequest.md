# HandlersCreateSSHConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**IpAddress** | **string** |  | 
**IsBastion** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**ParentConnectionId** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**PrivateSshKey** | **string** |  | 
**Scope** | Pointer to **string** | \&quot;user\&quot; (default) or \&quot;org\&quot; (BE-4) | [optional] 
**UseAgentForwarding** | Pointer to **bool** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersCreateSSHConnectionRequest

`func NewHandlersCreateSSHConnectionRequest(ipAddress string, name string, privateSshKey string, ) *HandlersCreateSSHConnectionRequest`

NewHandlersCreateSSHConnectionRequest instantiates a new HandlersCreateSSHConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateSSHConnectionRequestWithDefaults

`func NewHandlersCreateSSHConnectionRequestWithDefaults() *HandlersCreateSSHConnectionRequest`

NewHandlersCreateSSHConnectionRequestWithDefaults instantiates a new HandlersCreateSSHConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HandlersCreateSSHConnectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersCreateSSHConnectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersCreateSSHConnectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersCreateSSHConnectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddress

`func (o *HandlersCreateSSHConnectionRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *HandlersCreateSSHConnectionRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *HandlersCreateSSHConnectionRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetIsBastion

`func (o *HandlersCreateSSHConnectionRequest) GetIsBastion() bool`

GetIsBastion returns the IsBastion field if non-nil, zero value otherwise.

### GetIsBastionOk

`func (o *HandlersCreateSSHConnectionRequest) GetIsBastionOk() (*bool, bool)`

GetIsBastionOk returns a tuple with the IsBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBastion

`func (o *HandlersCreateSSHConnectionRequest) SetIsBastion(v bool)`

SetIsBastion sets IsBastion field to given value.

### HasIsBastion

`func (o *HandlersCreateSSHConnectionRequest) HasIsBastion() bool`

HasIsBastion returns a boolean if a field has been set.

### GetName

`func (o *HandlersCreateSSHConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersCreateSSHConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersCreateSSHConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentConnectionId

`func (o *HandlersCreateSSHConnectionRequest) GetParentConnectionId() int32`

GetParentConnectionId returns the ParentConnectionId field if non-nil, zero value otherwise.

### GetParentConnectionIdOk

`func (o *HandlersCreateSSHConnectionRequest) GetParentConnectionIdOk() (*int32, bool)`

GetParentConnectionIdOk returns a tuple with the ParentConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectionId

`func (o *HandlersCreateSSHConnectionRequest) SetParentConnectionId(v int32)`

SetParentConnectionId sets ParentConnectionId field to given value.

### HasParentConnectionId

`func (o *HandlersCreateSSHConnectionRequest) HasParentConnectionId() bool`

HasParentConnectionId returns a boolean if a field has been set.

### GetPort

`func (o *HandlersCreateSSHConnectionRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlersCreateSSHConnectionRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlersCreateSSHConnectionRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HandlersCreateSSHConnectionRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateSshKey

`func (o *HandlersCreateSSHConnectionRequest) GetPrivateSshKey() string`

GetPrivateSshKey returns the PrivateSshKey field if non-nil, zero value otherwise.

### GetPrivateSshKeyOk

`func (o *HandlersCreateSSHConnectionRequest) GetPrivateSshKeyOk() (*string, bool)`

GetPrivateSshKeyOk returns a tuple with the PrivateSshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSshKey

`func (o *HandlersCreateSSHConnectionRequest) SetPrivateSshKey(v string)`

SetPrivateSshKey sets PrivateSshKey field to given value.


### GetScope

`func (o *HandlersCreateSSHConnectionRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersCreateSSHConnectionRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersCreateSSHConnectionRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersCreateSSHConnectionRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUseAgentForwarding

`func (o *HandlersCreateSSHConnectionRequest) GetUseAgentForwarding() bool`

GetUseAgentForwarding returns the UseAgentForwarding field if non-nil, zero value otherwise.

### GetUseAgentForwardingOk

`func (o *HandlersCreateSSHConnectionRequest) GetUseAgentForwardingOk() (*bool, bool)`

GetUseAgentForwardingOk returns a tuple with the UseAgentForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgentForwarding

`func (o *HandlersCreateSSHConnectionRequest) SetUseAgentForwarding(v bool)`

SetUseAgentForwarding sets UseAgentForwarding field to given value.

### HasUseAgentForwarding

`func (o *HandlersCreateSSHConnectionRequest) HasUseAgentForwarding() bool`

HasUseAgentForwarding returns a boolean if a field has been set.

### GetUsername

`func (o *HandlersCreateSSHConnectionRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HandlersCreateSSHConnectionRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HandlersCreateSSHConnectionRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HandlersCreateSSHConnectionRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


