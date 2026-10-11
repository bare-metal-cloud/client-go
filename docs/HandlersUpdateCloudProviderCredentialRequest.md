# HandlersUpdateCloudProviderCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **map[string]interface{}** |  | 
**Label** | **string** |  | 
**Scope** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersUpdateCloudProviderCredentialRequest

`func NewHandlersUpdateCloudProviderCredentialRequest(credentials map[string]interface{}, label string, ) *HandlersUpdateCloudProviderCredentialRequest`

NewHandlersUpdateCloudProviderCredentialRequest instantiates a new HandlersUpdateCloudProviderCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersUpdateCloudProviderCredentialRequestWithDefaults

`func NewHandlersUpdateCloudProviderCredentialRequestWithDefaults() *HandlersUpdateCloudProviderCredentialRequest`

NewHandlersUpdateCloudProviderCredentialRequestWithDefaults instantiates a new HandlersUpdateCloudProviderCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *HandlersUpdateCloudProviderCredentialRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *HandlersUpdateCloudProviderCredentialRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *HandlersUpdateCloudProviderCredentialRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.


### GetLabel

`func (o *HandlersUpdateCloudProviderCredentialRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HandlersUpdateCloudProviderCredentialRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HandlersUpdateCloudProviderCredentialRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetScope

`func (o *HandlersUpdateCloudProviderCredentialRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersUpdateCloudProviderCredentialRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersUpdateCloudProviderCredentialRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersUpdateCloudProviderCredentialRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


