# HandlersDelegationMandateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedOps** | Pointer to **[]string** |  | [optional] 
**CertSignature** | Pointer to **string** |  | [optional] 
**EnvelopeSignature** | Pointer to **string** |  | [optional] 
**ExpiryUnix** | Pointer to **int32** |  | [optional] 
**HarnessFingerprint** | Pointer to **string** |  | [optional] 
**HarnessPubkey** | Pointer to **string** |  | [optional] 
**MachineId** | Pointer to **string** |  | [optional] 
**MandateId** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**RateLimitN** | Pointer to **int32** |  | [optional] 
**RateWindowSeconds** | Pointer to **int32** |  | [optional] 
**RevokeRef** | Pointer to **string** |  | [optional] 
**SignerFingerprint** | Pointer to **string** |  | [optional] 
**SignerUserId** | Pointer to **int32** |  | [optional] 
**SpendLimitCents** | Pointer to **int32** |  | [optional] 
**SupersedeRevoke** | Pointer to [**HandlersRevokeDelegationMandateRequest**](HandlersRevokeDelegationMandateRequest.md) | SupersedeRevoke is the browser-minted revoke token for the machine&#39;s PREVIOUS mandate (review should-fix 3): a re-mint flips the backend row to superseded, but the machine keeps the predecessor certificate INSTALLED until a revoke push names it — without the token, a stale installed mandate plus a replayed unconsumed action signature could resurrect narrowed-away authority. Required whenever this mint supersedes a live predecessor; the nonce must name THAT mandate. | [optional] 

## Methods

### NewHandlersDelegationMandateRequest

`func NewHandlersDelegationMandateRequest() *HandlersDelegationMandateRequest`

NewHandlersDelegationMandateRequest instantiates a new HandlersDelegationMandateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDelegationMandateRequestWithDefaults

`func NewHandlersDelegationMandateRequestWithDefaults() *HandlersDelegationMandateRequest`

NewHandlersDelegationMandateRequestWithDefaults instantiates a new HandlersDelegationMandateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedOps

`func (o *HandlersDelegationMandateRequest) GetAllowedOps() []string`

GetAllowedOps returns the AllowedOps field if non-nil, zero value otherwise.

### GetAllowedOpsOk

`func (o *HandlersDelegationMandateRequest) GetAllowedOpsOk() (*[]string, bool)`

GetAllowedOpsOk returns a tuple with the AllowedOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOps

`func (o *HandlersDelegationMandateRequest) SetAllowedOps(v []string)`

SetAllowedOps sets AllowedOps field to given value.

### HasAllowedOps

`func (o *HandlersDelegationMandateRequest) HasAllowedOps() bool`

HasAllowedOps returns a boolean if a field has been set.

### GetCertSignature

`func (o *HandlersDelegationMandateRequest) GetCertSignature() string`

GetCertSignature returns the CertSignature field if non-nil, zero value otherwise.

### GetCertSignatureOk

`func (o *HandlersDelegationMandateRequest) GetCertSignatureOk() (*string, bool)`

GetCertSignatureOk returns a tuple with the CertSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertSignature

`func (o *HandlersDelegationMandateRequest) SetCertSignature(v string)`

SetCertSignature sets CertSignature field to given value.

### HasCertSignature

`func (o *HandlersDelegationMandateRequest) HasCertSignature() bool`

HasCertSignature returns a boolean if a field has been set.

### GetEnvelopeSignature

`func (o *HandlersDelegationMandateRequest) GetEnvelopeSignature() string`

GetEnvelopeSignature returns the EnvelopeSignature field if non-nil, zero value otherwise.

### GetEnvelopeSignatureOk

`func (o *HandlersDelegationMandateRequest) GetEnvelopeSignatureOk() (*string, bool)`

GetEnvelopeSignatureOk returns a tuple with the EnvelopeSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvelopeSignature

`func (o *HandlersDelegationMandateRequest) SetEnvelopeSignature(v string)`

SetEnvelopeSignature sets EnvelopeSignature field to given value.

### HasEnvelopeSignature

`func (o *HandlersDelegationMandateRequest) HasEnvelopeSignature() bool`

HasEnvelopeSignature returns a boolean if a field has been set.

### GetExpiryUnix

`func (o *HandlersDelegationMandateRequest) GetExpiryUnix() int32`

GetExpiryUnix returns the ExpiryUnix field if non-nil, zero value otherwise.

### GetExpiryUnixOk

`func (o *HandlersDelegationMandateRequest) GetExpiryUnixOk() (*int32, bool)`

GetExpiryUnixOk returns a tuple with the ExpiryUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryUnix

`func (o *HandlersDelegationMandateRequest) SetExpiryUnix(v int32)`

SetExpiryUnix sets ExpiryUnix field to given value.

### HasExpiryUnix

`func (o *HandlersDelegationMandateRequest) HasExpiryUnix() bool`

HasExpiryUnix returns a boolean if a field has been set.

### GetHarnessFingerprint

`func (o *HandlersDelegationMandateRequest) GetHarnessFingerprint() string`

GetHarnessFingerprint returns the HarnessFingerprint field if non-nil, zero value otherwise.

### GetHarnessFingerprintOk

`func (o *HandlersDelegationMandateRequest) GetHarnessFingerprintOk() (*string, bool)`

GetHarnessFingerprintOk returns a tuple with the HarnessFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarnessFingerprint

`func (o *HandlersDelegationMandateRequest) SetHarnessFingerprint(v string)`

SetHarnessFingerprint sets HarnessFingerprint field to given value.

### HasHarnessFingerprint

`func (o *HandlersDelegationMandateRequest) HasHarnessFingerprint() bool`

HasHarnessFingerprint returns a boolean if a field has been set.

### GetHarnessPubkey

`func (o *HandlersDelegationMandateRequest) GetHarnessPubkey() string`

GetHarnessPubkey returns the HarnessPubkey field if non-nil, zero value otherwise.

### GetHarnessPubkeyOk

`func (o *HandlersDelegationMandateRequest) GetHarnessPubkeyOk() (*string, bool)`

GetHarnessPubkeyOk returns a tuple with the HarnessPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarnessPubkey

`func (o *HandlersDelegationMandateRequest) SetHarnessPubkey(v string)`

SetHarnessPubkey sets HarnessPubkey field to given value.

### HasHarnessPubkey

`func (o *HandlersDelegationMandateRequest) HasHarnessPubkey() bool`

HasHarnessPubkey returns a boolean if a field has been set.

### GetMachineId

`func (o *HandlersDelegationMandateRequest) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *HandlersDelegationMandateRequest) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *HandlersDelegationMandateRequest) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *HandlersDelegationMandateRequest) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetMandateId

`func (o *HandlersDelegationMandateRequest) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *HandlersDelegationMandateRequest) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *HandlersDelegationMandateRequest) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.

### HasMandateId

`func (o *HandlersDelegationMandateRequest) HasMandateId() bool`

HasMandateId returns a boolean if a field has been set.

### GetNonce

`func (o *HandlersDelegationMandateRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *HandlersDelegationMandateRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *HandlersDelegationMandateRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *HandlersDelegationMandateRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetRateLimitN

`func (o *HandlersDelegationMandateRequest) GetRateLimitN() int32`

GetRateLimitN returns the RateLimitN field if non-nil, zero value otherwise.

### GetRateLimitNOk

`func (o *HandlersDelegationMandateRequest) GetRateLimitNOk() (*int32, bool)`

GetRateLimitNOk returns a tuple with the RateLimitN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitN

`func (o *HandlersDelegationMandateRequest) SetRateLimitN(v int32)`

SetRateLimitN sets RateLimitN field to given value.

### HasRateLimitN

`func (o *HandlersDelegationMandateRequest) HasRateLimitN() bool`

HasRateLimitN returns a boolean if a field has been set.

### GetRateWindowSeconds

`func (o *HandlersDelegationMandateRequest) GetRateWindowSeconds() int32`

GetRateWindowSeconds returns the RateWindowSeconds field if non-nil, zero value otherwise.

### GetRateWindowSecondsOk

`func (o *HandlersDelegationMandateRequest) GetRateWindowSecondsOk() (*int32, bool)`

GetRateWindowSecondsOk returns a tuple with the RateWindowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateWindowSeconds

`func (o *HandlersDelegationMandateRequest) SetRateWindowSeconds(v int32)`

SetRateWindowSeconds sets RateWindowSeconds field to given value.

### HasRateWindowSeconds

`func (o *HandlersDelegationMandateRequest) HasRateWindowSeconds() bool`

HasRateWindowSeconds returns a boolean if a field has been set.

### GetRevokeRef

`func (o *HandlersDelegationMandateRequest) GetRevokeRef() string`

GetRevokeRef returns the RevokeRef field if non-nil, zero value otherwise.

### GetRevokeRefOk

`func (o *HandlersDelegationMandateRequest) GetRevokeRefOk() (*string, bool)`

GetRevokeRefOk returns a tuple with the RevokeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeRef

`func (o *HandlersDelegationMandateRequest) SetRevokeRef(v string)`

SetRevokeRef sets RevokeRef field to given value.

### HasRevokeRef

`func (o *HandlersDelegationMandateRequest) HasRevokeRef() bool`

HasRevokeRef returns a boolean if a field has been set.

### GetSignerFingerprint

`func (o *HandlersDelegationMandateRequest) GetSignerFingerprint() string`

GetSignerFingerprint returns the SignerFingerprint field if non-nil, zero value otherwise.

### GetSignerFingerprintOk

`func (o *HandlersDelegationMandateRequest) GetSignerFingerprintOk() (*string, bool)`

GetSignerFingerprintOk returns a tuple with the SignerFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerFingerprint

`func (o *HandlersDelegationMandateRequest) SetSignerFingerprint(v string)`

SetSignerFingerprint sets SignerFingerprint field to given value.

### HasSignerFingerprint

`func (o *HandlersDelegationMandateRequest) HasSignerFingerprint() bool`

HasSignerFingerprint returns a boolean if a field has been set.

### GetSignerUserId

`func (o *HandlersDelegationMandateRequest) GetSignerUserId() int32`

GetSignerUserId returns the SignerUserId field if non-nil, zero value otherwise.

### GetSignerUserIdOk

`func (o *HandlersDelegationMandateRequest) GetSignerUserIdOk() (*int32, bool)`

GetSignerUserIdOk returns a tuple with the SignerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUserId

`func (o *HandlersDelegationMandateRequest) SetSignerUserId(v int32)`

SetSignerUserId sets SignerUserId field to given value.

### HasSignerUserId

`func (o *HandlersDelegationMandateRequest) HasSignerUserId() bool`

HasSignerUserId returns a boolean if a field has been set.

### GetSpendLimitCents

`func (o *HandlersDelegationMandateRequest) GetSpendLimitCents() int32`

GetSpendLimitCents returns the SpendLimitCents field if non-nil, zero value otherwise.

### GetSpendLimitCentsOk

`func (o *HandlersDelegationMandateRequest) GetSpendLimitCentsOk() (*int32, bool)`

GetSpendLimitCentsOk returns a tuple with the SpendLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimitCents

`func (o *HandlersDelegationMandateRequest) SetSpendLimitCents(v int32)`

SetSpendLimitCents sets SpendLimitCents field to given value.

### HasSpendLimitCents

`func (o *HandlersDelegationMandateRequest) HasSpendLimitCents() bool`

HasSpendLimitCents returns a boolean if a field has been set.

### GetSupersedeRevoke

`func (o *HandlersDelegationMandateRequest) GetSupersedeRevoke() HandlersRevokeDelegationMandateRequest`

GetSupersedeRevoke returns the SupersedeRevoke field if non-nil, zero value otherwise.

### GetSupersedeRevokeOk

`func (o *HandlersDelegationMandateRequest) GetSupersedeRevokeOk() (*HandlersRevokeDelegationMandateRequest, bool)`

GetSupersedeRevokeOk returns a tuple with the SupersedeRevoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupersedeRevoke

`func (o *HandlersDelegationMandateRequest) SetSupersedeRevoke(v HandlersRevokeDelegationMandateRequest)`

SetSupersedeRevoke sets SupersedeRevoke field to given value.

### HasSupersedeRevoke

`func (o *HandlersDelegationMandateRequest) HasSupersedeRevoke() bool`

HasSupersedeRevoke returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


