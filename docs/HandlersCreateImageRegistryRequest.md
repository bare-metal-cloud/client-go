# HandlersCreateImageRegistryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **string** |  | [optional] 
**AwsCredentialId** | Pointer to **int32** |  | [optional] 
**AwsRegion** | Pointer to **string** |  | [optional] 
**AzureCredentialId** | Pointer to **int32** |  | [optional] 
**CaCertificate** | Pointer to **string** | For private registries | [optional] 
**EcrRegion** | Pointer to **string** | legacy alias | [optional] 
**GcpCredentialId** | Pointer to **int32** |  | [optional] 
**LoginServer** | Pointer to **string** | For ACR username/password auth | [optional] 
**Name** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Url** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersCreateImageRegistryRequest

`func NewHandlersCreateImageRegistryRequest(name string, type_ string, url string, ) *HandlersCreateImageRegistryRequest`

NewHandlersCreateImageRegistryRequest instantiates a new HandlersCreateImageRegistryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateImageRegistryRequestWithDefaults

`func NewHandlersCreateImageRegistryRequestWithDefaults() *HandlersCreateImageRegistryRequest`

NewHandlersCreateImageRegistryRequestWithDefaults instantiates a new HandlersCreateImageRegistryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *HandlersCreateImageRegistryRequest) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *HandlersCreateImageRegistryRequest) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *HandlersCreateImageRegistryRequest) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *HandlersCreateImageRegistryRequest) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAwsCredentialId

`func (o *HandlersCreateImageRegistryRequest) GetAwsCredentialId() int32`

GetAwsCredentialId returns the AwsCredentialId field if non-nil, zero value otherwise.

### GetAwsCredentialIdOk

`func (o *HandlersCreateImageRegistryRequest) GetAwsCredentialIdOk() (*int32, bool)`

GetAwsCredentialIdOk returns a tuple with the AwsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCredentialId

`func (o *HandlersCreateImageRegistryRequest) SetAwsCredentialId(v int32)`

SetAwsCredentialId sets AwsCredentialId field to given value.

### HasAwsCredentialId

`func (o *HandlersCreateImageRegistryRequest) HasAwsCredentialId() bool`

HasAwsCredentialId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *HandlersCreateImageRegistryRequest) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *HandlersCreateImageRegistryRequest) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *HandlersCreateImageRegistryRequest) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *HandlersCreateImageRegistryRequest) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAzureCredentialId

`func (o *HandlersCreateImageRegistryRequest) GetAzureCredentialId() int32`

GetAzureCredentialId returns the AzureCredentialId field if non-nil, zero value otherwise.

### GetAzureCredentialIdOk

`func (o *HandlersCreateImageRegistryRequest) GetAzureCredentialIdOk() (*int32, bool)`

GetAzureCredentialIdOk returns a tuple with the AzureCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCredentialId

`func (o *HandlersCreateImageRegistryRequest) SetAzureCredentialId(v int32)`

SetAzureCredentialId sets AzureCredentialId field to given value.

### HasAzureCredentialId

`func (o *HandlersCreateImageRegistryRequest) HasAzureCredentialId() bool`

HasAzureCredentialId returns a boolean if a field has been set.

### GetCaCertificate

`func (o *HandlersCreateImageRegistryRequest) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *HandlersCreateImageRegistryRequest) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *HandlersCreateImageRegistryRequest) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *HandlersCreateImageRegistryRequest) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetEcrRegion

`func (o *HandlersCreateImageRegistryRequest) GetEcrRegion() string`

GetEcrRegion returns the EcrRegion field if non-nil, zero value otherwise.

### GetEcrRegionOk

`func (o *HandlersCreateImageRegistryRequest) GetEcrRegionOk() (*string, bool)`

GetEcrRegionOk returns a tuple with the EcrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcrRegion

`func (o *HandlersCreateImageRegistryRequest) SetEcrRegion(v string)`

SetEcrRegion sets EcrRegion field to given value.

### HasEcrRegion

`func (o *HandlersCreateImageRegistryRequest) HasEcrRegion() bool`

HasEcrRegion returns a boolean if a field has been set.

### GetGcpCredentialId

`func (o *HandlersCreateImageRegistryRequest) GetGcpCredentialId() int32`

GetGcpCredentialId returns the GcpCredentialId field if non-nil, zero value otherwise.

### GetGcpCredentialIdOk

`func (o *HandlersCreateImageRegistryRequest) GetGcpCredentialIdOk() (*int32, bool)`

GetGcpCredentialIdOk returns a tuple with the GcpCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredentialId

`func (o *HandlersCreateImageRegistryRequest) SetGcpCredentialId(v int32)`

SetGcpCredentialId sets GcpCredentialId field to given value.

### HasGcpCredentialId

`func (o *HandlersCreateImageRegistryRequest) HasGcpCredentialId() bool`

HasGcpCredentialId returns a boolean if a field has been set.

### GetLoginServer

`func (o *HandlersCreateImageRegistryRequest) GetLoginServer() string`

GetLoginServer returns the LoginServer field if non-nil, zero value otherwise.

### GetLoginServerOk

`func (o *HandlersCreateImageRegistryRequest) GetLoginServerOk() (*string, bool)`

GetLoginServerOk returns a tuple with the LoginServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginServer

`func (o *HandlersCreateImageRegistryRequest) SetLoginServer(v string)`

SetLoginServer sets LoginServer field to given value.

### HasLoginServer

`func (o *HandlersCreateImageRegistryRequest) HasLoginServer() bool`

HasLoginServer returns a boolean if a field has been set.

### GetName

`func (o *HandlersCreateImageRegistryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersCreateImageRegistryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersCreateImageRegistryRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *HandlersCreateImageRegistryRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HandlersCreateImageRegistryRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HandlersCreateImageRegistryRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HandlersCreateImageRegistryRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetScope

`func (o *HandlersCreateImageRegistryRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersCreateImageRegistryRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersCreateImageRegistryRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersCreateImageRegistryRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetType

`func (o *HandlersCreateImageRegistryRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HandlersCreateImageRegistryRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HandlersCreateImageRegistryRequest) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *HandlersCreateImageRegistryRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HandlersCreateImageRegistryRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HandlersCreateImageRegistryRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *HandlersCreateImageRegistryRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HandlersCreateImageRegistryRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HandlersCreateImageRegistryRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HandlersCreateImageRegistryRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


