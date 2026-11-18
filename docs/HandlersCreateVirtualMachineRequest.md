# HandlersCreateVirtualMachineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BastionEnabled** | Pointer to **bool** | Story 18-7-8: SSH Security fields (Context-Based Access Control) | [optional] 
**BastionHost** | Pointer to **string** | Optional bastion hostname/IP override | [optional] 
**ConnectionId** | Pointer to **int32** |  | [optional] 
**ConnectionType** | Pointer to **string** | For provider-managed VMs (multipass, aws, etc) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MultipassConfig** | Pointer to [**HandlersMultipassVMConfig**](HandlersMultipassVMConfig.md) | Multipass VM configuration | [optional] 
**Name** | **string** |  | 
**Provider** | **string** | own_managed, multipass, aws, gcp, etc. | 
**SshConnection** | Pointer to [**HandlersCreateSSHConnectionRequest**](HandlersCreateSSHConnectionRequest.md) | SSH Connection details (for own_managed provider) | [optional] 
**UserPublicIps** | Pointer to **[]string** | Manual SSH access IPs (appended to allow list) | [optional] 

## Methods

### NewHandlersCreateVirtualMachineRequest

`func NewHandlersCreateVirtualMachineRequest(name string, provider string, ) *HandlersCreateVirtualMachineRequest`

NewHandlersCreateVirtualMachineRequest instantiates a new HandlersCreateVirtualMachineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateVirtualMachineRequestWithDefaults

`func NewHandlersCreateVirtualMachineRequestWithDefaults() *HandlersCreateVirtualMachineRequest`

NewHandlersCreateVirtualMachineRequestWithDefaults instantiates a new HandlersCreateVirtualMachineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBastionEnabled

`func (o *HandlersCreateVirtualMachineRequest) GetBastionEnabled() bool`

GetBastionEnabled returns the BastionEnabled field if non-nil, zero value otherwise.

### GetBastionEnabledOk

`func (o *HandlersCreateVirtualMachineRequest) GetBastionEnabledOk() (*bool, bool)`

GetBastionEnabledOk returns a tuple with the BastionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionEnabled

`func (o *HandlersCreateVirtualMachineRequest) SetBastionEnabled(v bool)`

SetBastionEnabled sets BastionEnabled field to given value.

### HasBastionEnabled

`func (o *HandlersCreateVirtualMachineRequest) HasBastionEnabled() bool`

HasBastionEnabled returns a boolean if a field has been set.

### GetBastionHost

`func (o *HandlersCreateVirtualMachineRequest) GetBastionHost() string`

GetBastionHost returns the BastionHost field if non-nil, zero value otherwise.

### GetBastionHostOk

`func (o *HandlersCreateVirtualMachineRequest) GetBastionHostOk() (*string, bool)`

GetBastionHostOk returns a tuple with the BastionHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionHost

`func (o *HandlersCreateVirtualMachineRequest) SetBastionHost(v string)`

SetBastionHost sets BastionHost field to given value.

### HasBastionHost

`func (o *HandlersCreateVirtualMachineRequest) HasBastionHost() bool`

HasBastionHost returns a boolean if a field has been set.

### GetConnectionId

`func (o *HandlersCreateVirtualMachineRequest) GetConnectionId() int32`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *HandlersCreateVirtualMachineRequest) GetConnectionIdOk() (*int32, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *HandlersCreateVirtualMachineRequest) SetConnectionId(v int32)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *HandlersCreateVirtualMachineRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionType

`func (o *HandlersCreateVirtualMachineRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HandlersCreateVirtualMachineRequest) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HandlersCreateVirtualMachineRequest) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *HandlersCreateVirtualMachineRequest) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetDescription

`func (o *HandlersCreateVirtualMachineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersCreateVirtualMachineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersCreateVirtualMachineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersCreateVirtualMachineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultipassConfig

`func (o *HandlersCreateVirtualMachineRequest) GetMultipassConfig() HandlersMultipassVMConfig`

GetMultipassConfig returns the MultipassConfig field if non-nil, zero value otherwise.

### GetMultipassConfigOk

`func (o *HandlersCreateVirtualMachineRequest) GetMultipassConfigOk() (*HandlersMultipassVMConfig, bool)`

GetMultipassConfigOk returns a tuple with the MultipassConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipassConfig

`func (o *HandlersCreateVirtualMachineRequest) SetMultipassConfig(v HandlersMultipassVMConfig)`

SetMultipassConfig sets MultipassConfig field to given value.

### HasMultipassConfig

`func (o *HandlersCreateVirtualMachineRequest) HasMultipassConfig() bool`

HasMultipassConfig returns a boolean if a field has been set.

### GetName

`func (o *HandlersCreateVirtualMachineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersCreateVirtualMachineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersCreateVirtualMachineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *HandlersCreateVirtualMachineRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersCreateVirtualMachineRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersCreateVirtualMachineRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSshConnection

`func (o *HandlersCreateVirtualMachineRequest) GetSshConnection() HandlersCreateSSHConnectionRequest`

GetSshConnection returns the SshConnection field if non-nil, zero value otherwise.

### GetSshConnectionOk

`func (o *HandlersCreateVirtualMachineRequest) GetSshConnectionOk() (*HandlersCreateSSHConnectionRequest, bool)`

GetSshConnectionOk returns a tuple with the SshConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConnection

`func (o *HandlersCreateVirtualMachineRequest) SetSshConnection(v HandlersCreateSSHConnectionRequest)`

SetSshConnection sets SshConnection field to given value.

### HasSshConnection

`func (o *HandlersCreateVirtualMachineRequest) HasSshConnection() bool`

HasSshConnection returns a boolean if a field has been set.

### GetUserPublicIps

`func (o *HandlersCreateVirtualMachineRequest) GetUserPublicIps() []string`

GetUserPublicIps returns the UserPublicIps field if non-nil, zero value otherwise.

### GetUserPublicIpsOk

`func (o *HandlersCreateVirtualMachineRequest) GetUserPublicIpsOk() (*[]string, bool)`

GetUserPublicIpsOk returns a tuple with the UserPublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPublicIps

`func (o *HandlersCreateVirtualMachineRequest) SetUserPublicIps(v []string)`

SetUserPublicIps sets UserPublicIps field to given value.

### HasUserPublicIps

`func (o *HandlersCreateVirtualMachineRequest) HasUserPublicIps() bool`

HasUserPublicIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


