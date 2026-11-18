# HandlersImportVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SshConnection** | Pointer to [**HandlersSSHConnectionRequest**](HandlersSSHConnectionRequest.md) |  | [optional] 

## Methods

### NewHandlersImportVMRequest

`func NewHandlersImportVMRequest() *HandlersImportVMRequest`

NewHandlersImportVMRequest instantiates a new HandlersImportVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersImportVMRequestWithDefaults

`func NewHandlersImportVMRequestWithDefaults() *HandlersImportVMRequest`

NewHandlersImportVMRequestWithDefaults instantiates a new HandlersImportVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HandlersImportVMRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HandlersImportVMRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HandlersImportVMRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HandlersImportVMRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *HandlersImportVMRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersImportVMRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersImportVMRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersImportVMRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSshConnection

`func (o *HandlersImportVMRequest) GetSshConnection() HandlersSSHConnectionRequest`

GetSshConnection returns the SshConnection field if non-nil, zero value otherwise.

### GetSshConnectionOk

`func (o *HandlersImportVMRequest) GetSshConnectionOk() (*HandlersSSHConnectionRequest, bool)`

GetSshConnectionOk returns a tuple with the SshConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConnection

`func (o *HandlersImportVMRequest) SetSshConnection(v HandlersSSHConnectionRequest)`

SetSshConnection sets SshConnection field to given value.

### HasSshConnection

`func (o *HandlersImportVMRequest) HasSshConnection() bool`

HasSshConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


