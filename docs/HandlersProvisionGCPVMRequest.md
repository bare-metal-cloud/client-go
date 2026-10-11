# HandlersProvisionGCPVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialId** | **int32** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DiskSizeGb** | Pointer to **int32** |  | [optional] 
**DiskType** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**MachineType** | **string** |  | 
**Name** | **string** |  | 
**Network** | Pointer to **string** |  | [optional] 
**SimpleMode** | Pointer to **bool** | Simple/Advanced mode | [optional] 
**SshPublicKey** | Pointer to **string** | SSH key options | [optional] 
**Subnetwork** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags (labels in GCP) | [optional] 
**Zone** | **string** |  | 

## Methods

### NewHandlersProvisionGCPVMRequest

`func NewHandlersProvisionGCPVMRequest(credentialId int32, machineType string, name string, zone string, ) *HandlersProvisionGCPVMRequest`

NewHandlersProvisionGCPVMRequest instantiates a new HandlersProvisionGCPVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersProvisionGCPVMRequestWithDefaults

`func NewHandlersProvisionGCPVMRequestWithDefaults() *HandlersProvisionGCPVMRequest`

NewHandlersProvisionGCPVMRequestWithDefaults instantiates a new HandlersProvisionGCPVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialId

`func (o *HandlersProvisionGCPVMRequest) GetCredentialId() int32`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *HandlersProvisionGCPVMRequest) GetCredentialIdOk() (*int32, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *HandlersProvisionGCPVMRequest) SetCredentialId(v int32)`

SetCredentialId sets CredentialId field to given value.


### GetDescription

`func (o *HandlersProvisionGCPVMRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersProvisionGCPVMRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersProvisionGCPVMRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersProvisionGCPVMRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskSizeGb

`func (o *HandlersProvisionGCPVMRequest) GetDiskSizeGb() int32`

GetDiskSizeGb returns the DiskSizeGb field if non-nil, zero value otherwise.

### GetDiskSizeGbOk

`func (o *HandlersProvisionGCPVMRequest) GetDiskSizeGbOk() (*int32, bool)`

GetDiskSizeGbOk returns a tuple with the DiskSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGb

`func (o *HandlersProvisionGCPVMRequest) SetDiskSizeGb(v int32)`

SetDiskSizeGb sets DiskSizeGb field to given value.

### HasDiskSizeGb

`func (o *HandlersProvisionGCPVMRequest) HasDiskSizeGb() bool`

HasDiskSizeGb returns a boolean if a field has been set.

### GetDiskType

`func (o *HandlersProvisionGCPVMRequest) GetDiskType() string`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *HandlersProvisionGCPVMRequest) GetDiskTypeOk() (*string, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *HandlersProvisionGCPVMRequest) SetDiskType(v string)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *HandlersProvisionGCPVMRequest) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetImage

`func (o *HandlersProvisionGCPVMRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *HandlersProvisionGCPVMRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *HandlersProvisionGCPVMRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *HandlersProvisionGCPVMRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMachineType

`func (o *HandlersProvisionGCPVMRequest) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *HandlersProvisionGCPVMRequest) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *HandlersProvisionGCPVMRequest) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.


### GetName

`func (o *HandlersProvisionGCPVMRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersProvisionGCPVMRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersProvisionGCPVMRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNetwork

`func (o *HandlersProvisionGCPVMRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *HandlersProvisionGCPVMRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *HandlersProvisionGCPVMRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *HandlersProvisionGCPVMRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSimpleMode

`func (o *HandlersProvisionGCPVMRequest) GetSimpleMode() bool`

GetSimpleMode returns the SimpleMode field if non-nil, zero value otherwise.

### GetSimpleModeOk

`func (o *HandlersProvisionGCPVMRequest) GetSimpleModeOk() (*bool, bool)`

GetSimpleModeOk returns a tuple with the SimpleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleMode

`func (o *HandlersProvisionGCPVMRequest) SetSimpleMode(v bool)`

SetSimpleMode sets SimpleMode field to given value.

### HasSimpleMode

`func (o *HandlersProvisionGCPVMRequest) HasSimpleMode() bool`

HasSimpleMode returns a boolean if a field has been set.

### GetSshPublicKey

`func (o *HandlersProvisionGCPVMRequest) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *HandlersProvisionGCPVMRequest) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *HandlersProvisionGCPVMRequest) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *HandlersProvisionGCPVMRequest) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### GetSubnetwork

`func (o *HandlersProvisionGCPVMRequest) GetSubnetwork() string`

GetSubnetwork returns the Subnetwork field if non-nil, zero value otherwise.

### GetSubnetworkOk

`func (o *HandlersProvisionGCPVMRequest) GetSubnetworkOk() (*string, bool)`

GetSubnetworkOk returns a tuple with the Subnetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetwork

`func (o *HandlersProvisionGCPVMRequest) SetSubnetwork(v string)`

SetSubnetwork sets Subnetwork field to given value.

### HasSubnetwork

`func (o *HandlersProvisionGCPVMRequest) HasSubnetwork() bool`

HasSubnetwork returns a boolean if a field has been set.

### GetTags

`func (o *HandlersProvisionGCPVMRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HandlersProvisionGCPVMRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HandlersProvisionGCPVMRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HandlersProvisionGCPVMRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetZone

`func (o *HandlersProvisionGCPVMRequest) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HandlersProvisionGCPVMRequest) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HandlersProvisionGCPVMRequest) SetZone(v string)`

SetZone sets Zone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


