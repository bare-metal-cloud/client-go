# HandlersProvisionAWSVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ami** | Pointer to **string** |  | [optional] 
**CredentialId** | **int32** |  | 
**CustomPorts** | Pointer to [**[]HandlersCustomPortRule**](HandlersCustomPortRule.md) | Custom ports for security group | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DiskSize** | Pointer to **int32** |  | [optional] 
**InstanceType** | **string** |  | 
**KeyPairName** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Region** | **string** |  | 
**SecurityGroupIds** | Pointer to **[]string** | Security groups | [optional] 
**SimpleMode** | Pointer to **bool** | Simple/Advanced mode | [optional] 
**SshPublicKey** | Pointer to **string** |  | [optional] 
**SubnetId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags | [optional] 
**UseExistingKeyPair** | Pointer to **bool** | SSH key options | [optional] 

## Methods

### NewHandlersProvisionAWSVMRequest

`func NewHandlersProvisionAWSVMRequest(credentialId int32, instanceType string, name string, region string, ) *HandlersProvisionAWSVMRequest`

NewHandlersProvisionAWSVMRequest instantiates a new HandlersProvisionAWSVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersProvisionAWSVMRequestWithDefaults

`func NewHandlersProvisionAWSVMRequestWithDefaults() *HandlersProvisionAWSVMRequest`

NewHandlersProvisionAWSVMRequestWithDefaults instantiates a new HandlersProvisionAWSVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmi

`func (o *HandlersProvisionAWSVMRequest) GetAmi() string`

GetAmi returns the Ami field if non-nil, zero value otherwise.

### GetAmiOk

`func (o *HandlersProvisionAWSVMRequest) GetAmiOk() (*string, bool)`

GetAmiOk returns a tuple with the Ami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmi

`func (o *HandlersProvisionAWSVMRequest) SetAmi(v string)`

SetAmi sets Ami field to given value.

### HasAmi

`func (o *HandlersProvisionAWSVMRequest) HasAmi() bool`

HasAmi returns a boolean if a field has been set.

### GetCredentialId

`func (o *HandlersProvisionAWSVMRequest) GetCredentialId() int32`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *HandlersProvisionAWSVMRequest) GetCredentialIdOk() (*int32, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *HandlersProvisionAWSVMRequest) SetCredentialId(v int32)`

SetCredentialId sets CredentialId field to given value.


### GetCustomPorts

`func (o *HandlersProvisionAWSVMRequest) GetCustomPorts() []HandlersCustomPortRule`

GetCustomPorts returns the CustomPorts field if non-nil, zero value otherwise.

### GetCustomPortsOk

`func (o *HandlersProvisionAWSVMRequest) GetCustomPortsOk() (*[]HandlersCustomPortRule, bool)`

GetCustomPortsOk returns a tuple with the CustomPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPorts

`func (o *HandlersProvisionAWSVMRequest) SetCustomPorts(v []HandlersCustomPortRule)`

SetCustomPorts sets CustomPorts field to given value.

### HasCustomPorts

`func (o *HandlersProvisionAWSVMRequest) HasCustomPorts() bool`

HasCustomPorts returns a boolean if a field has been set.

### GetDescription

`func (o *HandlersProvisionAWSVMRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersProvisionAWSVMRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersProvisionAWSVMRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersProvisionAWSVMRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskSize

`func (o *HandlersProvisionAWSVMRequest) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *HandlersProvisionAWSVMRequest) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *HandlersProvisionAWSVMRequest) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *HandlersProvisionAWSVMRequest) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetInstanceType

`func (o *HandlersProvisionAWSVMRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *HandlersProvisionAWSVMRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *HandlersProvisionAWSVMRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetKeyPairName

`func (o *HandlersProvisionAWSVMRequest) GetKeyPairName() string`

GetKeyPairName returns the KeyPairName field if non-nil, zero value otherwise.

### GetKeyPairNameOk

`func (o *HandlersProvisionAWSVMRequest) GetKeyPairNameOk() (*string, bool)`

GetKeyPairNameOk returns a tuple with the KeyPairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairName

`func (o *HandlersProvisionAWSVMRequest) SetKeyPairName(v string)`

SetKeyPairName sets KeyPairName field to given value.

### HasKeyPairName

`func (o *HandlersProvisionAWSVMRequest) HasKeyPairName() bool`

HasKeyPairName returns a boolean if a field has been set.

### GetName

`func (o *HandlersProvisionAWSVMRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersProvisionAWSVMRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersProvisionAWSVMRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *HandlersProvisionAWSVMRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersProvisionAWSVMRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersProvisionAWSVMRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSecurityGroupIds

`func (o *HandlersProvisionAWSVMRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *HandlersProvisionAWSVMRequest) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *HandlersProvisionAWSVMRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *HandlersProvisionAWSVMRequest) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### GetSimpleMode

`func (o *HandlersProvisionAWSVMRequest) GetSimpleMode() bool`

GetSimpleMode returns the SimpleMode field if non-nil, zero value otherwise.

### GetSimpleModeOk

`func (o *HandlersProvisionAWSVMRequest) GetSimpleModeOk() (*bool, bool)`

GetSimpleModeOk returns a tuple with the SimpleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleMode

`func (o *HandlersProvisionAWSVMRequest) SetSimpleMode(v bool)`

SetSimpleMode sets SimpleMode field to given value.

### HasSimpleMode

`func (o *HandlersProvisionAWSVMRequest) HasSimpleMode() bool`

HasSimpleMode returns a boolean if a field has been set.

### GetSshPublicKey

`func (o *HandlersProvisionAWSVMRequest) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *HandlersProvisionAWSVMRequest) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *HandlersProvisionAWSVMRequest) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *HandlersProvisionAWSVMRequest) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### GetSubnetId

`func (o *HandlersProvisionAWSVMRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *HandlersProvisionAWSVMRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *HandlersProvisionAWSVMRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *HandlersProvisionAWSVMRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetTags

`func (o *HandlersProvisionAWSVMRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HandlersProvisionAWSVMRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HandlersProvisionAWSVMRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HandlersProvisionAWSVMRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUseExistingKeyPair

`func (o *HandlersProvisionAWSVMRequest) GetUseExistingKeyPair() bool`

GetUseExistingKeyPair returns the UseExistingKeyPair field if non-nil, zero value otherwise.

### GetUseExistingKeyPairOk

`func (o *HandlersProvisionAWSVMRequest) GetUseExistingKeyPairOk() (*bool, bool)`

GetUseExistingKeyPairOk returns a tuple with the UseExistingKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingKeyPair

`func (o *HandlersProvisionAWSVMRequest) SetUseExistingKeyPair(v bool)`

SetUseExistingKeyPair sets UseExistingKeyPair field to given value.

### HasUseExistingKeyPair

`func (o *HandlersProvisionAWSVMRequest) HasUseExistingKeyPair() bool`

HasUseExistingKeyPair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


