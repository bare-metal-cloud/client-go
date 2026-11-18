# HandlersMultipassVMConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpus** | Pointer to **int32** | Number of CPUs (default: 2) | [optional] 
**Disk** | Pointer to **string** | Disk size (e.g., \&quot;10G\&quot;, default: \&quot;10G\&quot;) | [optional] 
**Image** | Pointer to **string** | Ubuntu image (e.g., \&quot;22.04\&quot;, default: \&quot;22.04\&quot;) | [optional] 
**Memory** | Pointer to **string** | Memory size (e.g., \&quot;2G\&quot;, default: \&quot;2G\&quot;) | [optional] 

## Methods

### NewHandlersMultipassVMConfig

`func NewHandlersMultipassVMConfig() *HandlersMultipassVMConfig`

NewHandlersMultipassVMConfig instantiates a new HandlersMultipassVMConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMultipassVMConfigWithDefaults

`func NewHandlersMultipassVMConfigWithDefaults() *HandlersMultipassVMConfig`

NewHandlersMultipassVMConfigWithDefaults instantiates a new HandlersMultipassVMConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpus

`func (o *HandlersMultipassVMConfig) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *HandlersMultipassVMConfig) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *HandlersMultipassVMConfig) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *HandlersMultipassVMConfig) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetDisk

`func (o *HandlersMultipassVMConfig) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *HandlersMultipassVMConfig) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *HandlersMultipassVMConfig) SetDisk(v string)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *HandlersMultipassVMConfig) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetImage

`func (o *HandlersMultipassVMConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *HandlersMultipassVMConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *HandlersMultipassVMConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *HandlersMultipassVMConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMemory

`func (o *HandlersMultipassVMConfig) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *HandlersMultipassVMConfig) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *HandlersMultipassVMConfig) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *HandlersMultipassVMConfig) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


