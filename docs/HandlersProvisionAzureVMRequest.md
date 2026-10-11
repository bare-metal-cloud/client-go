# HandlersProvisionAzureVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderCredentialId** | **int32** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DiskSizeGb** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**NetworkName** | Pointer to **string** |  | [optional] 
**Region** | **string** |  | 
**ResourceGroup** | Pointer to **string** |  | [optional] 
**SimpleMode** | Pointer to **bool** | Simple/Advanced mode | [optional] 
**SourceImage** | Pointer to **string** | Advanced options | [optional] 
**SshPublicKey** | Pointer to **string** | SSH key | [optional] 
**VmSize** | **string** |  | 

## Methods

### NewHandlersProvisionAzureVMRequest

`func NewHandlersProvisionAzureVMRequest(cloudProviderCredentialId int32, name string, region string, vmSize string, ) *HandlersProvisionAzureVMRequest`

NewHandlersProvisionAzureVMRequest instantiates a new HandlersProvisionAzureVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersProvisionAzureVMRequestWithDefaults

`func NewHandlersProvisionAzureVMRequestWithDefaults() *HandlersProvisionAzureVMRequest`

NewHandlersProvisionAzureVMRequestWithDefaults instantiates a new HandlersProvisionAzureVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderCredentialId

`func (o *HandlersProvisionAzureVMRequest) GetCloudProviderCredentialId() int32`

GetCloudProviderCredentialId returns the CloudProviderCredentialId field if non-nil, zero value otherwise.

### GetCloudProviderCredentialIdOk

`func (o *HandlersProvisionAzureVMRequest) GetCloudProviderCredentialIdOk() (*int32, bool)`

GetCloudProviderCredentialIdOk returns a tuple with the CloudProviderCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderCredentialId

`func (o *HandlersProvisionAzureVMRequest) SetCloudProviderCredentialId(v int32)`

SetCloudProviderCredentialId sets CloudProviderCredentialId field to given value.


### GetDescription

`func (o *HandlersProvisionAzureVMRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersProvisionAzureVMRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersProvisionAzureVMRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersProvisionAzureVMRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskSizeGb

`func (o *HandlersProvisionAzureVMRequest) GetDiskSizeGb() int32`

GetDiskSizeGb returns the DiskSizeGb field if non-nil, zero value otherwise.

### GetDiskSizeGbOk

`func (o *HandlersProvisionAzureVMRequest) GetDiskSizeGbOk() (*int32, bool)`

GetDiskSizeGbOk returns a tuple with the DiskSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGb

`func (o *HandlersProvisionAzureVMRequest) SetDiskSizeGb(v int32)`

SetDiskSizeGb sets DiskSizeGb field to given value.

### HasDiskSizeGb

`func (o *HandlersProvisionAzureVMRequest) HasDiskSizeGb() bool`

HasDiskSizeGb returns a boolean if a field has been set.

### GetName

`func (o *HandlersProvisionAzureVMRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersProvisionAzureVMRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersProvisionAzureVMRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkName

`func (o *HandlersProvisionAzureVMRequest) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *HandlersProvisionAzureVMRequest) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *HandlersProvisionAzureVMRequest) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *HandlersProvisionAzureVMRequest) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetRegion

`func (o *HandlersProvisionAzureVMRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersProvisionAzureVMRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersProvisionAzureVMRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetResourceGroup

`func (o *HandlersProvisionAzureVMRequest) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *HandlersProvisionAzureVMRequest) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *HandlersProvisionAzureVMRequest) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *HandlersProvisionAzureVMRequest) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetSimpleMode

`func (o *HandlersProvisionAzureVMRequest) GetSimpleMode() bool`

GetSimpleMode returns the SimpleMode field if non-nil, zero value otherwise.

### GetSimpleModeOk

`func (o *HandlersProvisionAzureVMRequest) GetSimpleModeOk() (*bool, bool)`

GetSimpleModeOk returns a tuple with the SimpleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleMode

`func (o *HandlersProvisionAzureVMRequest) SetSimpleMode(v bool)`

SetSimpleMode sets SimpleMode field to given value.

### HasSimpleMode

`func (o *HandlersProvisionAzureVMRequest) HasSimpleMode() bool`

HasSimpleMode returns a boolean if a field has been set.

### GetSourceImage

`func (o *HandlersProvisionAzureVMRequest) GetSourceImage() string`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *HandlersProvisionAzureVMRequest) GetSourceImageOk() (*string, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *HandlersProvisionAzureVMRequest) SetSourceImage(v string)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *HandlersProvisionAzureVMRequest) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetSshPublicKey

`func (o *HandlersProvisionAzureVMRequest) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *HandlersProvisionAzureVMRequest) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *HandlersProvisionAzureVMRequest) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *HandlersProvisionAzureVMRequest) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### GetVmSize

`func (o *HandlersProvisionAzureVMRequest) GetVmSize() string`

GetVmSize returns the VmSize field if non-nil, zero value otherwise.

### GetVmSizeOk

`func (o *HandlersProvisionAzureVMRequest) GetVmSizeOk() (*string, bool)`

GetVmSizeOk returns a tuple with the VmSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSize

`func (o *HandlersProvisionAzureVMRequest) SetVmSize(v string)`

SetVmSize sets VmSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


