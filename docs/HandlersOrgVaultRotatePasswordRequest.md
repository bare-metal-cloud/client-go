# HandlersOrgVaultRotatePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentOrgVaultPassword** | Pointer to **string** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**NewOrgVaultPassword** | Pointer to **string** |  | [optional] 
**NewPasswordSlot** | Pointer to **string** |  | [optional] 
**NewVaultSalt** | Pointer to **string** | Client-wrapped new password slot (base64) + its fresh Argon2id salt (base64, 32 bytes) + the new password — the #4026 browser flow&#39;s artifacts. All-or-nothing: any subset is a 400. The slot is an opaque blob the server stores verbatim (the same shape components/vault/slots wraps server-side). | [optional] 
**TrustedDevice** | Pointer to **bool** |  | [optional] 

## Methods

### NewHandlersOrgVaultRotatePasswordRequest

`func NewHandlersOrgVaultRotatePasswordRequest() *HandlersOrgVaultRotatePasswordRequest`

NewHandlersOrgVaultRotatePasswordRequest instantiates a new HandlersOrgVaultRotatePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersOrgVaultRotatePasswordRequestWithDefaults

`func NewHandlersOrgVaultRotatePasswordRequestWithDefaults() *HandlersOrgVaultRotatePasswordRequest`

NewHandlersOrgVaultRotatePasswordRequestWithDefaults instantiates a new HandlersOrgVaultRotatePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentOrgVaultPassword

`func (o *HandlersOrgVaultRotatePasswordRequest) GetCurrentOrgVaultPassword() string`

GetCurrentOrgVaultPassword returns the CurrentOrgVaultPassword field if non-nil, zero value otherwise.

### GetCurrentOrgVaultPasswordOk

`func (o *HandlersOrgVaultRotatePasswordRequest) GetCurrentOrgVaultPasswordOk() (*string, bool)`

GetCurrentOrgVaultPasswordOk returns a tuple with the CurrentOrgVaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentOrgVaultPassword

`func (o *HandlersOrgVaultRotatePasswordRequest) SetCurrentOrgVaultPassword(v string)`

SetCurrentOrgVaultPassword sets CurrentOrgVaultPassword field to given value.

### HasCurrentOrgVaultPassword

`func (o *HandlersOrgVaultRotatePasswordRequest) HasCurrentOrgVaultPassword() bool`

HasCurrentOrgVaultPassword returns a boolean if a field has been set.

### GetDeviceId

`func (o *HandlersOrgVaultRotatePasswordRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HandlersOrgVaultRotatePasswordRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HandlersOrgVaultRotatePasswordRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HandlersOrgVaultRotatePasswordRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetNewOrgVaultPassword

`func (o *HandlersOrgVaultRotatePasswordRequest) GetNewOrgVaultPassword() string`

GetNewOrgVaultPassword returns the NewOrgVaultPassword field if non-nil, zero value otherwise.

### GetNewOrgVaultPasswordOk

`func (o *HandlersOrgVaultRotatePasswordRequest) GetNewOrgVaultPasswordOk() (*string, bool)`

GetNewOrgVaultPasswordOk returns a tuple with the NewOrgVaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrgVaultPassword

`func (o *HandlersOrgVaultRotatePasswordRequest) SetNewOrgVaultPassword(v string)`

SetNewOrgVaultPassword sets NewOrgVaultPassword field to given value.

### HasNewOrgVaultPassword

`func (o *HandlersOrgVaultRotatePasswordRequest) HasNewOrgVaultPassword() bool`

HasNewOrgVaultPassword returns a boolean if a field has been set.

### GetNewPasswordSlot

`func (o *HandlersOrgVaultRotatePasswordRequest) GetNewPasswordSlot() string`

GetNewPasswordSlot returns the NewPasswordSlot field if non-nil, zero value otherwise.

### GetNewPasswordSlotOk

`func (o *HandlersOrgVaultRotatePasswordRequest) GetNewPasswordSlotOk() (*string, bool)`

GetNewPasswordSlotOk returns a tuple with the NewPasswordSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPasswordSlot

`func (o *HandlersOrgVaultRotatePasswordRequest) SetNewPasswordSlot(v string)`

SetNewPasswordSlot sets NewPasswordSlot field to given value.

### HasNewPasswordSlot

`func (o *HandlersOrgVaultRotatePasswordRequest) HasNewPasswordSlot() bool`

HasNewPasswordSlot returns a boolean if a field has been set.

### GetNewVaultSalt

`func (o *HandlersOrgVaultRotatePasswordRequest) GetNewVaultSalt() string`

GetNewVaultSalt returns the NewVaultSalt field if non-nil, zero value otherwise.

### GetNewVaultSaltOk

`func (o *HandlersOrgVaultRotatePasswordRequest) GetNewVaultSaltOk() (*string, bool)`

GetNewVaultSaltOk returns a tuple with the NewVaultSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVaultSalt

`func (o *HandlersOrgVaultRotatePasswordRequest) SetNewVaultSalt(v string)`

SetNewVaultSalt sets NewVaultSalt field to given value.

### HasNewVaultSalt

`func (o *HandlersOrgVaultRotatePasswordRequest) HasNewVaultSalt() bool`

HasNewVaultSalt returns a boolean if a field has been set.

### GetTrustedDevice

`func (o *HandlersOrgVaultRotatePasswordRequest) GetTrustedDevice() bool`

GetTrustedDevice returns the TrustedDevice field if non-nil, zero value otherwise.

### GetTrustedDeviceOk

`func (o *HandlersOrgVaultRotatePasswordRequest) GetTrustedDeviceOk() (*bool, bool)`

GetTrustedDeviceOk returns a tuple with the TrustedDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDevice

`func (o *HandlersOrgVaultRotatePasswordRequest) SetTrustedDevice(v bool)`

SetTrustedDevice sets TrustedDevice field to given value.

### HasTrustedDevice

`func (o *HandlersOrgVaultRotatePasswordRequest) HasTrustedDevice() bool`

HasTrustedDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


