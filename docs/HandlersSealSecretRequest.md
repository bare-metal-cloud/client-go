# HandlersSealSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | seal action (e.g. \&quot;seal_sa_secret\&quot;) — must be in seal policy AllowedActions | [optional] 
**EncryptedDataId** | Pointer to **int32** | the EncryptedData row to seal | [optional] 
**SecretType** | Pointer to **string** | broker secret-type label (e.g. \&quot;sa_credentials\&quot;) | [optional] 
**WorkerKey** | Pointer to **string** | hex-encoded target worker pubkey (allowlisted in seal policy) | [optional] 

## Methods

### NewHandlersSealSecretRequest

`func NewHandlersSealSecretRequest() *HandlersSealSecretRequest`

NewHandlersSealSecretRequest instantiates a new HandlersSealSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersSealSecretRequestWithDefaults

`func NewHandlersSealSecretRequestWithDefaults() *HandlersSealSecretRequest`

NewHandlersSealSecretRequestWithDefaults instantiates a new HandlersSealSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *HandlersSealSecretRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HandlersSealSecretRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HandlersSealSecretRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HandlersSealSecretRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetEncryptedDataId

`func (o *HandlersSealSecretRequest) GetEncryptedDataId() int32`

GetEncryptedDataId returns the EncryptedDataId field if non-nil, zero value otherwise.

### GetEncryptedDataIdOk

`func (o *HandlersSealSecretRequest) GetEncryptedDataIdOk() (*int32, bool)`

GetEncryptedDataIdOk returns a tuple with the EncryptedDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedDataId

`func (o *HandlersSealSecretRequest) SetEncryptedDataId(v int32)`

SetEncryptedDataId sets EncryptedDataId field to given value.

### HasEncryptedDataId

`func (o *HandlersSealSecretRequest) HasEncryptedDataId() bool`

HasEncryptedDataId returns a boolean if a field has been set.

### GetSecretType

`func (o *HandlersSealSecretRequest) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *HandlersSealSecretRequest) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *HandlersSealSecretRequest) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *HandlersSealSecretRequest) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetWorkerKey

`func (o *HandlersSealSecretRequest) GetWorkerKey() string`

GetWorkerKey returns the WorkerKey field if non-nil, zero value otherwise.

### GetWorkerKeyOk

`func (o *HandlersSealSecretRequest) GetWorkerKeyOk() (*string, bool)`

GetWorkerKeyOk returns a tuple with the WorkerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerKey

`func (o *HandlersSealSecretRequest) SetWorkerKey(v string)`

SetWorkerKey sets WorkerKey field to given value.

### HasWorkerKey

`func (o *HandlersSealSecretRequest) HasWorkerKey() bool`

HasWorkerKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


