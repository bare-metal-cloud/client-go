# HandlersCreateCloudProviderCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **map[string]interface{}** |  | 
**Label** | **string** |  | 
**Provider** | **string** |  | 
**Scope** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersCreateCloudProviderCredentialRequest

`func NewHandlersCreateCloudProviderCredentialRequest(credentials map[string]interface{}, label string, provider string, ) *HandlersCreateCloudProviderCredentialRequest`

NewHandlersCreateCloudProviderCredentialRequest instantiates a new HandlersCreateCloudProviderCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateCloudProviderCredentialRequestWithDefaults

`func NewHandlersCreateCloudProviderCredentialRequestWithDefaults() *HandlersCreateCloudProviderCredentialRequest`

NewHandlersCreateCloudProviderCredentialRequestWithDefaults instantiates a new HandlersCreateCloudProviderCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *HandlersCreateCloudProviderCredentialRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *HandlersCreateCloudProviderCredentialRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *HandlersCreateCloudProviderCredentialRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.


### GetLabel

`func (o *HandlersCreateCloudProviderCredentialRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *HandlersCreateCloudProviderCredentialRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *HandlersCreateCloudProviderCredentialRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetProvider

`func (o *HandlersCreateCloudProviderCredentialRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersCreateCloudProviderCredentialRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersCreateCloudProviderCredentialRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetScope

`func (o *HandlersCreateCloudProviderCredentialRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersCreateCloudProviderCredentialRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersCreateCloudProviderCredentialRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersCreateCloudProviderCredentialRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


