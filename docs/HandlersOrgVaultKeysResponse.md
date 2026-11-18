# HandlersOrgVaultKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordSlot** | Pointer to **string** |  | [optional] 
**RecoveryKeyCreatedAt** | Pointer to **string** | RecoveryKeyCreatedAt is when the CURRENT recovery key was generated. | [optional] 
**RecoveryKeyFingerprint** | Pointer to **string** | RecoveryKeyFingerprint is base64url(SHA-256(recovery key)) — A9 metadata: proof a key exists, never the key. | [optional] 
**RecoveryKeyVersion** | Pointer to **int32** |  | [optional] 
**RecoverySalt** | Pointer to **string** |  | [optional] 
**RecoverySlot** | Pointer to **string** |  | [optional] 
**VaultSalt** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersOrgVaultKeysResponse

`func NewHandlersOrgVaultKeysResponse() *HandlersOrgVaultKeysResponse`

NewHandlersOrgVaultKeysResponse instantiates a new HandlersOrgVaultKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersOrgVaultKeysResponseWithDefaults

`func NewHandlersOrgVaultKeysResponseWithDefaults() *HandlersOrgVaultKeysResponse`

NewHandlersOrgVaultKeysResponseWithDefaults instantiates a new HandlersOrgVaultKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordSlot

`func (o *HandlersOrgVaultKeysResponse) GetPasswordSlot() string`

GetPasswordSlot returns the PasswordSlot field if non-nil, zero value otherwise.

### GetPasswordSlotOk

`func (o *HandlersOrgVaultKeysResponse) GetPasswordSlotOk() (*string, bool)`

GetPasswordSlotOk returns a tuple with the PasswordSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSlot

`func (o *HandlersOrgVaultKeysResponse) SetPasswordSlot(v string)`

SetPasswordSlot sets PasswordSlot field to given value.

### HasPasswordSlot

`func (o *HandlersOrgVaultKeysResponse) HasPasswordSlot() bool`

HasPasswordSlot returns a boolean if a field has been set.

### GetRecoveryKeyCreatedAt

`func (o *HandlersOrgVaultKeysResponse) GetRecoveryKeyCreatedAt() string`

GetRecoveryKeyCreatedAt returns the RecoveryKeyCreatedAt field if non-nil, zero value otherwise.

### GetRecoveryKeyCreatedAtOk

`func (o *HandlersOrgVaultKeysResponse) GetRecoveryKeyCreatedAtOk() (*string, bool)`

GetRecoveryKeyCreatedAtOk returns a tuple with the RecoveryKeyCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeyCreatedAt

`func (o *HandlersOrgVaultKeysResponse) SetRecoveryKeyCreatedAt(v string)`

SetRecoveryKeyCreatedAt sets RecoveryKeyCreatedAt field to given value.

### HasRecoveryKeyCreatedAt

`func (o *HandlersOrgVaultKeysResponse) HasRecoveryKeyCreatedAt() bool`

HasRecoveryKeyCreatedAt returns a boolean if a field has been set.

### GetRecoveryKeyFingerprint

`func (o *HandlersOrgVaultKeysResponse) GetRecoveryKeyFingerprint() string`

GetRecoveryKeyFingerprint returns the RecoveryKeyFingerprint field if non-nil, zero value otherwise.

### GetRecoveryKeyFingerprintOk

`func (o *HandlersOrgVaultKeysResponse) GetRecoveryKeyFingerprintOk() (*string, bool)`

GetRecoveryKeyFingerprintOk returns a tuple with the RecoveryKeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeyFingerprint

`func (o *HandlersOrgVaultKeysResponse) SetRecoveryKeyFingerprint(v string)`

SetRecoveryKeyFingerprint sets RecoveryKeyFingerprint field to given value.

### HasRecoveryKeyFingerprint

`func (o *HandlersOrgVaultKeysResponse) HasRecoveryKeyFingerprint() bool`

HasRecoveryKeyFingerprint returns a boolean if a field has been set.

### GetRecoveryKeyVersion

`func (o *HandlersOrgVaultKeysResponse) GetRecoveryKeyVersion() int32`

GetRecoveryKeyVersion returns the RecoveryKeyVersion field if non-nil, zero value otherwise.

### GetRecoveryKeyVersionOk

`func (o *HandlersOrgVaultKeysResponse) GetRecoveryKeyVersionOk() (*int32, bool)`

GetRecoveryKeyVersionOk returns a tuple with the RecoveryKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeyVersion

`func (o *HandlersOrgVaultKeysResponse) SetRecoveryKeyVersion(v int32)`

SetRecoveryKeyVersion sets RecoveryKeyVersion field to given value.

### HasRecoveryKeyVersion

`func (o *HandlersOrgVaultKeysResponse) HasRecoveryKeyVersion() bool`

HasRecoveryKeyVersion returns a boolean if a field has been set.

### GetRecoverySalt

`func (o *HandlersOrgVaultKeysResponse) GetRecoverySalt() string`

GetRecoverySalt returns the RecoverySalt field if non-nil, zero value otherwise.

### GetRecoverySaltOk

`func (o *HandlersOrgVaultKeysResponse) GetRecoverySaltOk() (*string, bool)`

GetRecoverySaltOk returns a tuple with the RecoverySalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverySalt

`func (o *HandlersOrgVaultKeysResponse) SetRecoverySalt(v string)`

SetRecoverySalt sets RecoverySalt field to given value.

### HasRecoverySalt

`func (o *HandlersOrgVaultKeysResponse) HasRecoverySalt() bool`

HasRecoverySalt returns a boolean if a field has been set.

### GetRecoverySlot

`func (o *HandlersOrgVaultKeysResponse) GetRecoverySlot() string`

GetRecoverySlot returns the RecoverySlot field if non-nil, zero value otherwise.

### GetRecoverySlotOk

`func (o *HandlersOrgVaultKeysResponse) GetRecoverySlotOk() (*string, bool)`

GetRecoverySlotOk returns a tuple with the RecoverySlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverySlot

`func (o *HandlersOrgVaultKeysResponse) SetRecoverySlot(v string)`

SetRecoverySlot sets RecoverySlot field to given value.

### HasRecoverySlot

`func (o *HandlersOrgVaultKeysResponse) HasRecoverySlot() bool`

HasRecoverySlot returns a boolean if a field has been set.

### GetVaultSalt

`func (o *HandlersOrgVaultKeysResponse) GetVaultSalt() string`

GetVaultSalt returns the VaultSalt field if non-nil, zero value otherwise.

### GetVaultSaltOk

`func (o *HandlersOrgVaultKeysResponse) GetVaultSaltOk() (*string, bool)`

GetVaultSaltOk returns a tuple with the VaultSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSalt

`func (o *HandlersOrgVaultKeysResponse) SetVaultSalt(v string)`

SetVaultSalt sets VaultSalt field to given value.

### HasVaultSalt

`func (o *HandlersOrgVaultKeysResponse) HasVaultSalt() bool`

HasVaultSalt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


