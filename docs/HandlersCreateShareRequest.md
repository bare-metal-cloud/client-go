# HandlersCreateShareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAtUnix** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LinkedRelease** | Pointer to [**HandlersShareReleaseBody**](HandlersShareReleaseBody.md) | LinkedRelease is the release of a linked image registry&#39;s cloud credential, required for a registry in cloud_credential mode and refused otherwise (BAR-1675, ADR-104 D2). | [optional] 
**RecipientUserId** | Pointer to **int32** |  | [optional] 
**Release** | Pointer to [**HandlersShareReleaseBody**](HandlersShareReleaseBody.md) | Release is the sharer&#39;s one-time release of the secret&#39;s own value. It is required when the secret has one, and refused otherwise (ADR-104 D2). | [optional] 
**SecretId** | Pointer to **int32** |  | [optional] 
**SecretType** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersCreateShareRequest

`func NewHandlersCreateShareRequest() *HandlersCreateShareRequest`

NewHandlersCreateShareRequest instantiates a new HandlersCreateShareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateShareRequestWithDefaults

`func NewHandlersCreateShareRequestWithDefaults() *HandlersCreateShareRequest`

NewHandlersCreateShareRequestWithDefaults instantiates a new HandlersCreateShareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAtUnix

`func (o *HandlersCreateShareRequest) GetExpiresAtUnix() int32`

GetExpiresAtUnix returns the ExpiresAtUnix field if non-nil, zero value otherwise.

### GetExpiresAtUnixOk

`func (o *HandlersCreateShareRequest) GetExpiresAtUnixOk() (*int32, bool)`

GetExpiresAtUnixOk returns a tuple with the ExpiresAtUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAtUnix

`func (o *HandlersCreateShareRequest) SetExpiresAtUnix(v int32)`

SetExpiresAtUnix sets ExpiresAtUnix field to given value.

### HasExpiresAtUnix

`func (o *HandlersCreateShareRequest) HasExpiresAtUnix() bool`

HasExpiresAtUnix returns a boolean if a field has been set.

### GetId

`func (o *HandlersCreateShareRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersCreateShareRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersCreateShareRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersCreateShareRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinkedRelease

`func (o *HandlersCreateShareRequest) GetLinkedRelease() HandlersShareReleaseBody`

GetLinkedRelease returns the LinkedRelease field if non-nil, zero value otherwise.

### GetLinkedReleaseOk

`func (o *HandlersCreateShareRequest) GetLinkedReleaseOk() (*HandlersShareReleaseBody, bool)`

GetLinkedReleaseOk returns a tuple with the LinkedRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedRelease

`func (o *HandlersCreateShareRequest) SetLinkedRelease(v HandlersShareReleaseBody)`

SetLinkedRelease sets LinkedRelease field to given value.

### HasLinkedRelease

`func (o *HandlersCreateShareRequest) HasLinkedRelease() bool`

HasLinkedRelease returns a boolean if a field has been set.

### GetRecipientUserId

`func (o *HandlersCreateShareRequest) GetRecipientUserId() int32`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *HandlersCreateShareRequest) GetRecipientUserIdOk() (*int32, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *HandlersCreateShareRequest) SetRecipientUserId(v int32)`

SetRecipientUserId sets RecipientUserId field to given value.

### HasRecipientUserId

`func (o *HandlersCreateShareRequest) HasRecipientUserId() bool`

HasRecipientUserId returns a boolean if a field has been set.

### GetRelease

`func (o *HandlersCreateShareRequest) GetRelease() HandlersShareReleaseBody`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *HandlersCreateShareRequest) GetReleaseOk() (*HandlersShareReleaseBody, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *HandlersCreateShareRequest) SetRelease(v HandlersShareReleaseBody)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *HandlersCreateShareRequest) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetSecretId

`func (o *HandlersCreateShareRequest) GetSecretId() int32`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *HandlersCreateShareRequest) GetSecretIdOk() (*int32, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *HandlersCreateShareRequest) SetSecretId(v int32)`

SetSecretId sets SecretId field to given value.

### HasSecretId

`func (o *HandlersCreateShareRequest) HasSecretId() bool`

HasSecretId returns a boolean if a field has been set.

### GetSecretType

`func (o *HandlersCreateShareRequest) GetSecretType() string`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *HandlersCreateShareRequest) GetSecretTypeOk() (*string, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *HandlersCreateShareRequest) SetSecretType(v string)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *HandlersCreateShareRequest) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


