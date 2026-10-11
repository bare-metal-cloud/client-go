# HandlersOrgVaultUnlockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** |  | [optional] 
**MfaToken** | Pointer to **string** | MFA fields are forward-compatible (BE-X will add org MFA — see vault_unlock_mfa.go for the user-side analog). BE-2 ignores them. | [optional] 
**OrgVaultPassword** | Pointer to **string** |  | [optional] 
**TrustedDevice** | Pointer to **bool** |  | [optional] 

## Methods

### NewHandlersOrgVaultUnlockRequest

`func NewHandlersOrgVaultUnlockRequest() *HandlersOrgVaultUnlockRequest`

NewHandlersOrgVaultUnlockRequest instantiates a new HandlersOrgVaultUnlockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersOrgVaultUnlockRequestWithDefaults

`func NewHandlersOrgVaultUnlockRequestWithDefaults() *HandlersOrgVaultUnlockRequest`

NewHandlersOrgVaultUnlockRequestWithDefaults instantiates a new HandlersOrgVaultUnlockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *HandlersOrgVaultUnlockRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *HandlersOrgVaultUnlockRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *HandlersOrgVaultUnlockRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *HandlersOrgVaultUnlockRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetMfaToken

`func (o *HandlersOrgVaultUnlockRequest) GetMfaToken() string`

GetMfaToken returns the MfaToken field if non-nil, zero value otherwise.

### GetMfaTokenOk

`func (o *HandlersOrgVaultUnlockRequest) GetMfaTokenOk() (*string, bool)`

GetMfaTokenOk returns a tuple with the MfaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaToken

`func (o *HandlersOrgVaultUnlockRequest) SetMfaToken(v string)`

SetMfaToken sets MfaToken field to given value.

### HasMfaToken

`func (o *HandlersOrgVaultUnlockRequest) HasMfaToken() bool`

HasMfaToken returns a boolean if a field has been set.

### GetOrgVaultPassword

`func (o *HandlersOrgVaultUnlockRequest) GetOrgVaultPassword() string`

GetOrgVaultPassword returns the OrgVaultPassword field if non-nil, zero value otherwise.

### GetOrgVaultPasswordOk

`func (o *HandlersOrgVaultUnlockRequest) GetOrgVaultPasswordOk() (*string, bool)`

GetOrgVaultPasswordOk returns a tuple with the OrgVaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgVaultPassword

`func (o *HandlersOrgVaultUnlockRequest) SetOrgVaultPassword(v string)`

SetOrgVaultPassword sets OrgVaultPassword field to given value.

### HasOrgVaultPassword

`func (o *HandlersOrgVaultUnlockRequest) HasOrgVaultPassword() bool`

HasOrgVaultPassword returns a boolean if a field has been set.

### GetTrustedDevice

`func (o *HandlersOrgVaultUnlockRequest) GetTrustedDevice() bool`

GetTrustedDevice returns the TrustedDevice field if non-nil, zero value otherwise.

### GetTrustedDeviceOk

`func (o *HandlersOrgVaultUnlockRequest) GetTrustedDeviceOk() (*bool, bool)`

GetTrustedDeviceOk returns a tuple with the TrustedDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedDevice

`func (o *HandlersOrgVaultUnlockRequest) SetTrustedDevice(v bool)`

SetTrustedDevice sets TrustedDevice field to given value.

### HasTrustedDevice

`func (o *HandlersOrgVaultUnlockRequest) HasTrustedDevice() bool`

HasTrustedDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


