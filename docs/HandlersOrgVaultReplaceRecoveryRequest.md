# HandlersOrgVaultReplaceRecoveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledged** | Pointer to **bool** | Acknowledged is the A9 gate: the replace dialog&#39;s \&quot;I stored this key\&quot; checkbox. The endpoint refuses to invalidate the old key until the browser says the user stored the new one. | [optional] 
**RecoveryKeyFingerprint** | Pointer to **string** | RecoveryKeyFingerprint is base64url(SHA-256(new recovery key)) — computed in the browser; the server records it, never the key. | [optional] 
**RecoverySalt** | Pointer to **string** | RecoverySalt is the fresh 32-byte Argon2id salt for the new recovery derivation (browser-generated, public by construction). | [optional] 
**RecoverySlot** | Pointer to **string** | RecoverySlot is the new recovery-wrapped vault key (opaque, components/vault/slots format, wrapped under Argon2id(new recovery key, RecoverySalt) in the browser). | [optional] 

## Methods

### NewHandlersOrgVaultReplaceRecoveryRequest

`func NewHandlersOrgVaultReplaceRecoveryRequest() *HandlersOrgVaultReplaceRecoveryRequest`

NewHandlersOrgVaultReplaceRecoveryRequest instantiates a new HandlersOrgVaultReplaceRecoveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersOrgVaultReplaceRecoveryRequestWithDefaults

`func NewHandlersOrgVaultReplaceRecoveryRequestWithDefaults() *HandlersOrgVaultReplaceRecoveryRequest`

NewHandlersOrgVaultReplaceRecoveryRequestWithDefaults instantiates a new HandlersOrgVaultReplaceRecoveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledged

`func (o *HandlersOrgVaultReplaceRecoveryRequest) GetAcknowledged() bool`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *HandlersOrgVaultReplaceRecoveryRequest) GetAcknowledgedOk() (*bool, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *HandlersOrgVaultReplaceRecoveryRequest) SetAcknowledged(v bool)`

SetAcknowledged sets Acknowledged field to given value.

### HasAcknowledged

`func (o *HandlersOrgVaultReplaceRecoveryRequest) HasAcknowledged() bool`

HasAcknowledged returns a boolean if a field has been set.

### GetRecoveryKeyFingerprint

`func (o *HandlersOrgVaultReplaceRecoveryRequest) GetRecoveryKeyFingerprint() string`

GetRecoveryKeyFingerprint returns the RecoveryKeyFingerprint field if non-nil, zero value otherwise.

### GetRecoveryKeyFingerprintOk

`func (o *HandlersOrgVaultReplaceRecoveryRequest) GetRecoveryKeyFingerprintOk() (*string, bool)`

GetRecoveryKeyFingerprintOk returns a tuple with the RecoveryKeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeyFingerprint

`func (o *HandlersOrgVaultReplaceRecoveryRequest) SetRecoveryKeyFingerprint(v string)`

SetRecoveryKeyFingerprint sets RecoveryKeyFingerprint field to given value.

### HasRecoveryKeyFingerprint

`func (o *HandlersOrgVaultReplaceRecoveryRequest) HasRecoveryKeyFingerprint() bool`

HasRecoveryKeyFingerprint returns a boolean if a field has been set.

### GetRecoverySalt

`func (o *HandlersOrgVaultReplaceRecoveryRequest) GetRecoverySalt() string`

GetRecoverySalt returns the RecoverySalt field if non-nil, zero value otherwise.

### GetRecoverySaltOk

`func (o *HandlersOrgVaultReplaceRecoveryRequest) GetRecoverySaltOk() (*string, bool)`

GetRecoverySaltOk returns a tuple with the RecoverySalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverySalt

`func (o *HandlersOrgVaultReplaceRecoveryRequest) SetRecoverySalt(v string)`

SetRecoverySalt sets RecoverySalt field to given value.

### HasRecoverySalt

`func (o *HandlersOrgVaultReplaceRecoveryRequest) HasRecoverySalt() bool`

HasRecoverySalt returns a boolean if a field has been set.

### GetRecoverySlot

`func (o *HandlersOrgVaultReplaceRecoveryRequest) GetRecoverySlot() string`

GetRecoverySlot returns the RecoverySlot field if non-nil, zero value otherwise.

### GetRecoverySlotOk

`func (o *HandlersOrgVaultReplaceRecoveryRequest) GetRecoverySlotOk() (*string, bool)`

GetRecoverySlotOk returns a tuple with the RecoverySlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverySlot

`func (o *HandlersOrgVaultReplaceRecoveryRequest) SetRecoverySlot(v string)`

SetRecoverySlot sets RecoverySlot field to given value.

### HasRecoverySlot

`func (o *HandlersOrgVaultReplaceRecoveryRequest) HasRecoverySlot() bool`

HasRecoverySlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


