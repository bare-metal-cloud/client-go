# HandlersGuardTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinarySha256** | Pointer to **string** |  | [optional] 
**ExpiryUnix** | Pointer to **int32** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**MachineId** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**OpId** | Pointer to **string** |  | [optional] 
**OpName** | Pointer to **string** |  | [optional] 
**RotateNewPubkey** | Pointer to **string** | The RotateGuardKey payload (rotation class only): the new guard PUBLIC key (base64-std) and the dual-key window&#39;s end (unix seconds). Stored verbatim; the agent verifies and applies. | [optional] 
**RotatePrevValidUntil** | Pointer to **int32** |  | [optional] 
**Signature** | Pointer to **string** | Signature is base64-std of the raw Ed25519 signature over the canonical token encoding (the browser&#39;s wire form; stored verbatim). | [optional] 
**SignerFingerprint** | Pointer to **string** |  | [optional] 
**SignerUserId** | Pointer to **int32** | Signer attribution (#4036, per-user minting keys): the minting user&#39;s id and guard-key fingerprint (base64url SHA-256 of the pubkey). The backend CANNOT verify the claim (it holds no guard key) — it records what the authenticated session minted, and the AGENT&#39;s roster verifies the signature itself. | [optional] 

## Methods

### NewHandlersGuardTokenRequest

`func NewHandlersGuardTokenRequest() *HandlersGuardTokenRequest`

NewHandlersGuardTokenRequest instantiates a new HandlersGuardTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersGuardTokenRequestWithDefaults

`func NewHandlersGuardTokenRequestWithDefaults() *HandlersGuardTokenRequest`

NewHandlersGuardTokenRequestWithDefaults instantiates a new HandlersGuardTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinarySha256

`func (o *HandlersGuardTokenRequest) GetBinarySha256() string`

GetBinarySha256 returns the BinarySha256 field if non-nil, zero value otherwise.

### GetBinarySha256Ok

`func (o *HandlersGuardTokenRequest) GetBinarySha256Ok() (*string, bool)`

GetBinarySha256Ok returns a tuple with the BinarySha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinarySha256

`func (o *HandlersGuardTokenRequest) SetBinarySha256(v string)`

SetBinarySha256 sets BinarySha256 field to given value.

### HasBinarySha256

`func (o *HandlersGuardTokenRequest) HasBinarySha256() bool`

HasBinarySha256 returns a boolean if a field has been set.

### GetExpiryUnix

`func (o *HandlersGuardTokenRequest) GetExpiryUnix() int32`

GetExpiryUnix returns the ExpiryUnix field if non-nil, zero value otherwise.

### GetExpiryUnixOk

`func (o *HandlersGuardTokenRequest) GetExpiryUnixOk() (*int32, bool)`

GetExpiryUnixOk returns a tuple with the ExpiryUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryUnix

`func (o *HandlersGuardTokenRequest) SetExpiryUnix(v int32)`

SetExpiryUnix sets ExpiryUnix field to given value.

### HasExpiryUnix

`func (o *HandlersGuardTokenRequest) HasExpiryUnix() bool`

HasExpiryUnix returns a boolean if a field has been set.

### GetKind

`func (o *HandlersGuardTokenRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HandlersGuardTokenRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HandlersGuardTokenRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HandlersGuardTokenRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMachineId

`func (o *HandlersGuardTokenRequest) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *HandlersGuardTokenRequest) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *HandlersGuardTokenRequest) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *HandlersGuardTokenRequest) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetNonce

`func (o *HandlersGuardTokenRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *HandlersGuardTokenRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *HandlersGuardTokenRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *HandlersGuardTokenRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetOpId

`func (o *HandlersGuardTokenRequest) GetOpId() string`

GetOpId returns the OpId field if non-nil, zero value otherwise.

### GetOpIdOk

`func (o *HandlersGuardTokenRequest) GetOpIdOk() (*string, bool)`

GetOpIdOk returns a tuple with the OpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpId

`func (o *HandlersGuardTokenRequest) SetOpId(v string)`

SetOpId sets OpId field to given value.

### HasOpId

`func (o *HandlersGuardTokenRequest) HasOpId() bool`

HasOpId returns a boolean if a field has been set.

### GetOpName

`func (o *HandlersGuardTokenRequest) GetOpName() string`

GetOpName returns the OpName field if non-nil, zero value otherwise.

### GetOpNameOk

`func (o *HandlersGuardTokenRequest) GetOpNameOk() (*string, bool)`

GetOpNameOk returns a tuple with the OpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpName

`func (o *HandlersGuardTokenRequest) SetOpName(v string)`

SetOpName sets OpName field to given value.

### HasOpName

`func (o *HandlersGuardTokenRequest) HasOpName() bool`

HasOpName returns a boolean if a field has been set.

### GetRotateNewPubkey

`func (o *HandlersGuardTokenRequest) GetRotateNewPubkey() string`

GetRotateNewPubkey returns the RotateNewPubkey field if non-nil, zero value otherwise.

### GetRotateNewPubkeyOk

`func (o *HandlersGuardTokenRequest) GetRotateNewPubkeyOk() (*string, bool)`

GetRotateNewPubkeyOk returns a tuple with the RotateNewPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateNewPubkey

`func (o *HandlersGuardTokenRequest) SetRotateNewPubkey(v string)`

SetRotateNewPubkey sets RotateNewPubkey field to given value.

### HasRotateNewPubkey

`func (o *HandlersGuardTokenRequest) HasRotateNewPubkey() bool`

HasRotateNewPubkey returns a boolean if a field has been set.

### GetRotatePrevValidUntil

`func (o *HandlersGuardTokenRequest) GetRotatePrevValidUntil() int32`

GetRotatePrevValidUntil returns the RotatePrevValidUntil field if non-nil, zero value otherwise.

### GetRotatePrevValidUntilOk

`func (o *HandlersGuardTokenRequest) GetRotatePrevValidUntilOk() (*int32, bool)`

GetRotatePrevValidUntilOk returns a tuple with the RotatePrevValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatePrevValidUntil

`func (o *HandlersGuardTokenRequest) SetRotatePrevValidUntil(v int32)`

SetRotatePrevValidUntil sets RotatePrevValidUntil field to given value.

### HasRotatePrevValidUntil

`func (o *HandlersGuardTokenRequest) HasRotatePrevValidUntil() bool`

HasRotatePrevValidUntil returns a boolean if a field has been set.

### GetSignature

`func (o *HandlersGuardTokenRequest) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *HandlersGuardTokenRequest) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *HandlersGuardTokenRequest) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *HandlersGuardTokenRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSignerFingerprint

`func (o *HandlersGuardTokenRequest) GetSignerFingerprint() string`

GetSignerFingerprint returns the SignerFingerprint field if non-nil, zero value otherwise.

### GetSignerFingerprintOk

`func (o *HandlersGuardTokenRequest) GetSignerFingerprintOk() (*string, bool)`

GetSignerFingerprintOk returns a tuple with the SignerFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerFingerprint

`func (o *HandlersGuardTokenRequest) SetSignerFingerprint(v string)`

SetSignerFingerprint sets SignerFingerprint field to given value.

### HasSignerFingerprint

`func (o *HandlersGuardTokenRequest) HasSignerFingerprint() bool`

HasSignerFingerprint returns a boolean if a field has been set.

### GetSignerUserId

`func (o *HandlersGuardTokenRequest) GetSignerUserId() int32`

GetSignerUserId returns the SignerUserId field if non-nil, zero value otherwise.

### GetSignerUserIdOk

`func (o *HandlersGuardTokenRequest) GetSignerUserIdOk() (*int32, bool)`

GetSignerUserIdOk returns a tuple with the SignerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUserId

`func (o *HandlersGuardTokenRequest) SetSignerUserId(v int32)`

SetSignerUserId sets SignerUserId field to given value.

### HasSignerUserId

`func (o *HandlersGuardTokenRequest) HasSignerUserId() bool`

HasSignerUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


