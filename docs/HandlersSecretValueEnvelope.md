# HandlersSecretValueEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | Pointer to **string** |  | [optional] 
**EncryptedDataId** | Pointer to **int32** | EncryptedDataID is the value row the envelope was read from. The browser sends it back as the release marker when it shares the secret (BAR-1675, ADR-104 D3). It identifies a row; it carries no key or plaintext. | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**WrappedDek** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersSecretValueEnvelope

`func NewHandlersSecretValueEnvelope() *HandlersSecretValueEnvelope`

NewHandlersSecretValueEnvelope instantiates a new HandlersSecretValueEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersSecretValueEnvelopeWithDefaults

`func NewHandlersSecretValueEnvelopeWithDefaults() *HandlersSecretValueEnvelope`

NewHandlersSecretValueEnvelopeWithDefaults instantiates a new HandlersSecretValueEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *HandlersSecretValueEnvelope) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *HandlersSecretValueEnvelope) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *HandlersSecretValueEnvelope) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.

### HasCiphertext

`func (o *HandlersSecretValueEnvelope) HasCiphertext() bool`

HasCiphertext returns a boolean if a field has been set.

### GetEncryptedDataId

`func (o *HandlersSecretValueEnvelope) GetEncryptedDataId() int32`

GetEncryptedDataId returns the EncryptedDataId field if non-nil, zero value otherwise.

### GetEncryptedDataIdOk

`func (o *HandlersSecretValueEnvelope) GetEncryptedDataIdOk() (*int32, bool)`

GetEncryptedDataIdOk returns a tuple with the EncryptedDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedDataId

`func (o *HandlersSecretValueEnvelope) SetEncryptedDataId(v int32)`

SetEncryptedDataId sets EncryptedDataId field to given value.

### HasEncryptedDataId

`func (o *HandlersSecretValueEnvelope) HasEncryptedDataId() bool`

HasEncryptedDataId returns a boolean if a field has been set.

### GetFormat

`func (o *HandlersSecretValueEnvelope) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *HandlersSecretValueEnvelope) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *HandlersSecretValueEnvelope) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *HandlersSecretValueEnvelope) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetVersion

`func (o *HandlersSecretValueEnvelope) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HandlersSecretValueEnvelope) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HandlersSecretValueEnvelope) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HandlersSecretValueEnvelope) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWrappedDek

`func (o *HandlersSecretValueEnvelope) GetWrappedDek() string`

GetWrappedDek returns the WrappedDek field if non-nil, zero value otherwise.

### GetWrappedDekOk

`func (o *HandlersSecretValueEnvelope) GetWrappedDekOk() (*string, bool)`

GetWrappedDekOk returns a tuple with the WrappedDek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrappedDek

`func (o *HandlersSecretValueEnvelope) SetWrappedDek(v string)`

SetWrappedDek sets WrappedDek field to given value.

### HasWrappedDek

`func (o *HandlersSecretValueEnvelope) HasWrappedDek() bool`

HasWrappedDek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


