# HandlersAWSVMWorkerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedSshIps** | Pointer to **[]string** | Security options | [optional] 
**AmiId** | Pointer to **string** | Advanced options (for explicit user control) If provided, these bypass the VM planner&#39;s decision-making | [optional] 
**BastionEnabled** | Pointer to **bool** | Story 18-7-8: SSH Security fields (Context-Based Access Control) | [optional] 
**BastionHost** | Pointer to **string** | Optional bastion hostname/IP override | [optional] 
**Cpus** | **int32** |  | 
**CredentialId** | **int32** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DiskGb** | **int32** |  | 
**Image** | Pointer to **string** | Optional, defaults to ubuntu-22.04 | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**KeyPairName** | Pointer to **string** |  | [optional] 
**MemoryMb** | **int32** |  | 
**Name** | **string** |  | 
**Region** | **string** | Required specifications | 
**SecurityGroupId** | Pointer to **string** |  | [optional] 
**SshPublicKey** | Pointer to **string** | SSH configuration | [optional] 
**SubnetId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags | [optional] 
**UserPublicIps** | Pointer to **[]string** | Manual SSH access IPs (appended to allow list) | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersAWSVMWorkerRequest

`func NewHandlersAWSVMWorkerRequest(cpus int32, credentialId int32, diskGb int32, memoryMb int32, name string, region string, ) *HandlersAWSVMWorkerRequest`

NewHandlersAWSVMWorkerRequest instantiates a new HandlersAWSVMWorkerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersAWSVMWorkerRequestWithDefaults

`func NewHandlersAWSVMWorkerRequestWithDefaults() *HandlersAWSVMWorkerRequest`

NewHandlersAWSVMWorkerRequestWithDefaults instantiates a new HandlersAWSVMWorkerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedSshIps

`func (o *HandlersAWSVMWorkerRequest) GetAllowedSshIps() []string`

GetAllowedSshIps returns the AllowedSshIps field if non-nil, zero value otherwise.

### GetAllowedSshIpsOk

`func (o *HandlersAWSVMWorkerRequest) GetAllowedSshIpsOk() (*[]string, bool)`

GetAllowedSshIpsOk returns a tuple with the AllowedSshIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSshIps

`func (o *HandlersAWSVMWorkerRequest) SetAllowedSshIps(v []string)`

SetAllowedSshIps sets AllowedSshIps field to given value.

### HasAllowedSshIps

`func (o *HandlersAWSVMWorkerRequest) HasAllowedSshIps() bool`

HasAllowedSshIps returns a boolean if a field has been set.

### GetAmiId

`func (o *HandlersAWSVMWorkerRequest) GetAmiId() string`

GetAmiId returns the AmiId field if non-nil, zero value otherwise.

### GetAmiIdOk

`func (o *HandlersAWSVMWorkerRequest) GetAmiIdOk() (*string, bool)`

GetAmiIdOk returns a tuple with the AmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiId

`func (o *HandlersAWSVMWorkerRequest) SetAmiId(v string)`

SetAmiId sets AmiId field to given value.

### HasAmiId

`func (o *HandlersAWSVMWorkerRequest) HasAmiId() bool`

HasAmiId returns a boolean if a field has been set.

### GetBastionEnabled

`func (o *HandlersAWSVMWorkerRequest) GetBastionEnabled() bool`

GetBastionEnabled returns the BastionEnabled field if non-nil, zero value otherwise.

### GetBastionEnabledOk

`func (o *HandlersAWSVMWorkerRequest) GetBastionEnabledOk() (*bool, bool)`

GetBastionEnabledOk returns a tuple with the BastionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionEnabled

`func (o *HandlersAWSVMWorkerRequest) SetBastionEnabled(v bool)`

SetBastionEnabled sets BastionEnabled field to given value.

### HasBastionEnabled

`func (o *HandlersAWSVMWorkerRequest) HasBastionEnabled() bool`

HasBastionEnabled returns a boolean if a field has been set.

### GetBastionHost

`func (o *HandlersAWSVMWorkerRequest) GetBastionHost() string`

GetBastionHost returns the BastionHost field if non-nil, zero value otherwise.

### GetBastionHostOk

`func (o *HandlersAWSVMWorkerRequest) GetBastionHostOk() (*string, bool)`

GetBastionHostOk returns a tuple with the BastionHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastionHost

`func (o *HandlersAWSVMWorkerRequest) SetBastionHost(v string)`

SetBastionHost sets BastionHost field to given value.

### HasBastionHost

`func (o *HandlersAWSVMWorkerRequest) HasBastionHost() bool`

HasBastionHost returns a boolean if a field has been set.

### GetCpus

`func (o *HandlersAWSVMWorkerRequest) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *HandlersAWSVMWorkerRequest) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *HandlersAWSVMWorkerRequest) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### GetCredentialId

`func (o *HandlersAWSVMWorkerRequest) GetCredentialId() int32`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *HandlersAWSVMWorkerRequest) GetCredentialIdOk() (*int32, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *HandlersAWSVMWorkerRequest) SetCredentialId(v int32)`

SetCredentialId sets CredentialId field to given value.


### GetDescription

`func (o *HandlersAWSVMWorkerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersAWSVMWorkerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersAWSVMWorkerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersAWSVMWorkerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskGb

`func (o *HandlersAWSVMWorkerRequest) GetDiskGb() int32`

GetDiskGb returns the DiskGb field if non-nil, zero value otherwise.

### GetDiskGbOk

`func (o *HandlersAWSVMWorkerRequest) GetDiskGbOk() (*int32, bool)`

GetDiskGbOk returns a tuple with the DiskGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGb

`func (o *HandlersAWSVMWorkerRequest) SetDiskGb(v int32)`

SetDiskGb sets DiskGb field to given value.


### GetImage

`func (o *HandlersAWSVMWorkerRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *HandlersAWSVMWorkerRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *HandlersAWSVMWorkerRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *HandlersAWSVMWorkerRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInstanceType

`func (o *HandlersAWSVMWorkerRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *HandlersAWSVMWorkerRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *HandlersAWSVMWorkerRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *HandlersAWSVMWorkerRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetKeyPairName

`func (o *HandlersAWSVMWorkerRequest) GetKeyPairName() string`

GetKeyPairName returns the KeyPairName field if non-nil, zero value otherwise.

### GetKeyPairNameOk

`func (o *HandlersAWSVMWorkerRequest) GetKeyPairNameOk() (*string, bool)`

GetKeyPairNameOk returns a tuple with the KeyPairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairName

`func (o *HandlersAWSVMWorkerRequest) SetKeyPairName(v string)`

SetKeyPairName sets KeyPairName field to given value.

### HasKeyPairName

`func (o *HandlersAWSVMWorkerRequest) HasKeyPairName() bool`

HasKeyPairName returns a boolean if a field has been set.

### GetMemoryMb

`func (o *HandlersAWSVMWorkerRequest) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *HandlersAWSVMWorkerRequest) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *HandlersAWSVMWorkerRequest) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.


### GetName

`func (o *HandlersAWSVMWorkerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersAWSVMWorkerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersAWSVMWorkerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *HandlersAWSVMWorkerRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersAWSVMWorkerRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersAWSVMWorkerRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSecurityGroupId

`func (o *HandlersAWSVMWorkerRequest) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *HandlersAWSVMWorkerRequest) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *HandlersAWSVMWorkerRequest) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *HandlersAWSVMWorkerRequest) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### GetSshPublicKey

`func (o *HandlersAWSVMWorkerRequest) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *HandlersAWSVMWorkerRequest) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *HandlersAWSVMWorkerRequest) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *HandlersAWSVMWorkerRequest) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### GetSubnetId

`func (o *HandlersAWSVMWorkerRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *HandlersAWSVMWorkerRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *HandlersAWSVMWorkerRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *HandlersAWSVMWorkerRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetTags

`func (o *HandlersAWSVMWorkerRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HandlersAWSVMWorkerRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HandlersAWSVMWorkerRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HandlersAWSVMWorkerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUserPublicIps

`func (o *HandlersAWSVMWorkerRequest) GetUserPublicIps() []string`

GetUserPublicIps returns the UserPublicIps field if non-nil, zero value otherwise.

### GetUserPublicIpsOk

`func (o *HandlersAWSVMWorkerRequest) GetUserPublicIpsOk() (*[]string, bool)`

GetUserPublicIpsOk returns a tuple with the UserPublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPublicIps

`func (o *HandlersAWSVMWorkerRequest) SetUserPublicIps(v []string)`

SetUserPublicIps sets UserPublicIps field to given value.

### HasUserPublicIps

`func (o *HandlersAWSVMWorkerRequest) HasUserPublicIps() bool`

HasUserPublicIps returns a boolean if a field has been set.

### GetVpcId

`func (o *HandlersAWSVMWorkerRequest) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *HandlersAWSVMWorkerRequest) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *HandlersAWSVMWorkerRequest) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *HandlersAWSVMWorkerRequest) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


