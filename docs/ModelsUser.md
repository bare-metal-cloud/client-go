# ModelsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to [**[]ModelsNetwork**](ModelsNetwork.md) |  | [optional] 
**OauthAccounts** | Pointer to [**[]ModelsOAuthAccount**](ModelsOAuthAccount.md) |  | [optional] 
**UiMode** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsUser

`func NewModelsUser() *ModelsUser`

NewModelsUser instantiates a new ModelsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsUserWithDefaults

`func NewModelsUserWithDefaults() *ModelsUser`

NewModelsUserWithDefaults instantiates a new ModelsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *ModelsUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ModelsUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ModelsUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ModelsUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *ModelsUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModelsUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *ModelsUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelsUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *ModelsUser) GetNetworks() []ModelsNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ModelsUser) GetNetworksOk() (*[]ModelsNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ModelsUser) SetNetworks(v []ModelsNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ModelsUser) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOauthAccounts

`func (o *ModelsUser) GetOauthAccounts() []ModelsOAuthAccount`

GetOauthAccounts returns the OauthAccounts field if non-nil, zero value otherwise.

### GetOauthAccountsOk

`func (o *ModelsUser) GetOauthAccountsOk() (*[]ModelsOAuthAccount, bool)`

GetOauthAccountsOk returns a tuple with the OauthAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAccounts

`func (o *ModelsUser) SetOauthAccounts(v []ModelsOAuthAccount)`

SetOauthAccounts sets OauthAccounts field to given value.

### HasOauthAccounts

`func (o *ModelsUser) HasOauthAccounts() bool`

HasOauthAccounts returns a boolean if a field has been set.

### GetUiMode

`func (o *ModelsUser) GetUiMode() string`

GetUiMode returns the UiMode field if non-nil, zero value otherwise.

### GetUiModeOk

`func (o *ModelsUser) GetUiModeOk() (*string, bool)`

GetUiModeOk returns a tuple with the UiMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiMode

`func (o *ModelsUser) SetUiMode(v string)`

SetUiMode sets UiMode field to given value.

### HasUiMode

`func (o *ModelsUser) HasUiMode() bool`

HasUiMode returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVerified

`func (o *ModelsUser) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ModelsUser) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ModelsUser) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ModelsUser) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


