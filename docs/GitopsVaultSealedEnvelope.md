# GitopsVaultSealedEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | Pointer to **[]int32** | base64-decoded bytes | [optional] 
**WrappedDek** | Pointer to **[]int32** | base64-decoded bytes | [optional] 

## Methods

### NewGitopsVaultSealedEnvelope

`func NewGitopsVaultSealedEnvelope() *GitopsVaultSealedEnvelope`

NewGitopsVaultSealedEnvelope instantiates a new GitopsVaultSealedEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitopsVaultSealedEnvelopeWithDefaults

`func NewGitopsVaultSealedEnvelopeWithDefaults() *GitopsVaultSealedEnvelope`

NewGitopsVaultSealedEnvelopeWithDefaults instantiates a new GitopsVaultSealedEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *GitopsVaultSealedEnvelope) GetCiphertext() []int32`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *GitopsVaultSealedEnvelope) GetCiphertextOk() (*[]int32, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *GitopsVaultSealedEnvelope) SetCiphertext(v []int32)`

SetCiphertext sets Ciphertext field to given value.

### HasCiphertext

`func (o *GitopsVaultSealedEnvelope) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.

### GetWrappedDek

`func (o *GitopsVaultSealedEnvelope) GetWrappedDek() []int32`

GetWrappedDek returns the WrappedDek field if non-nil, zero value otherwise.

### GetWrappedDekOk

`func (o *GitopsVaultSealedEnvelope) GetWrappedDekOk() (*[]int32, bool)`

GetWrappedDekOk returns a tuple with the WrappedDek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrappedDek

`func (o *GitopsVaultSealedEnvelope) SetWrappedDek(v []int32)`

SetWrappedDek sets WrappedDek field to given value.

### HasWrappedDek

`func (o *GitopsVaultSealedEnvelope) HasWrappedDek() bool`

HasWrappedDek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


