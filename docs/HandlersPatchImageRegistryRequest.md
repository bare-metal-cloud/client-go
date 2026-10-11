# HandlersPatchImageRegistryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to **string** |  | [optional] 
**AwsCredentialId** | Pointer to **int32** |  | [optional] 
**AwsRegion** | Pointer to **string** |  | [optional] 
**AzureCredentialId** | Pointer to **int32** |  | [optional] 
**CaCertificate** | Pointer to **string** |  | [optional] 
**EcrRegion** | Pointer to **string** |  | [optional] 
**GcpCredentialId** | Pointer to **int32** |  | [optional] 
**LoginServer** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersPatchImageRegistryRequest

`func NewHandlersPatchImageRegistryRequest() *HandlersPatchImageRegistryRequest`

NewHandlersPatchImageRegistryRequest instantiates a new HandlersPatchImageRegistryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersPatchImageRegistryRequestWithDefaults

`func NewHandlersPatchImageRegistryRequestWithDefaults() *HandlersPatchImageRegistryRequest`

NewHandlersPatchImageRegistryRequestWithDefaults instantiates a new HandlersPatchImageRegistryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *HandlersPatchImageRegistryRequest) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *HandlersPatchImageRegistryRequest) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *HandlersPatchImageRegistryRequest) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *HandlersPatchImageRegistryRequest) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAwsCredentialId

`func (o *HandlersPatchImageRegistryRequest) GetAwsCredentialId() int32`

GetAwsCredentialId returns the AwsCredentialId field if non-nil, zero value otherwise.

### GetAwsCredentialIdOk

`func (o *HandlersPatchImageRegistryRequest) GetAwsCredentialIdOk() (*int32, bool)`

GetAwsCredentialIdOk returns a tuple with the AwsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCredentialId

`func (o *HandlersPatchImageRegistryRequest) SetAwsCredentialId(v int32)`

SetAwsCredentialId sets AwsCredentialId field to given value.

### HasAwsCredentialId

`func (o *HandlersPatchImageRegistryRequest) HasAwsCredentialId() bool`

HasAwsCredentialId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *HandlersPatchImageRegistryRequest) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *HandlersPatchImageRegistryRequest) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *HandlersPatchImageRegistryRequest) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *HandlersPatchImageRegistryRequest) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAzureCredentialId

`func (o *HandlersPatchImageRegistryRequest) GetAzureCredentialId() int32`

GetAzureCredentialId returns the AzureCredentialId field if non-nil, zero value otherwise.

### GetAzureCredentialIdOk

`func (o *HandlersPatchImageRegistryRequest) GetAzureCredentialIdOk() (*int32, bool)`

GetAzureCredentialIdOk returns a tuple with the AzureCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCredentialId

`func (o *HandlersPatchImageRegistryRequest) SetAzureCredentialId(v int32)`

SetAzureCredentialId sets AzureCredentialId field to given value.

### HasAzureCredentialId

`func (o *HandlersPatchImageRegistryRequest) HasAzureCredentialId() bool`

HasAzureCredentialId returns a boolean if a field has been set.

### GetCaCertificate

`func (o *HandlersPatchImageRegistryRequest) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *HandlersPatchImageRegistryRequest) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *HandlersPatchImageRegistryRequest) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *HandlersPatchImageRegistryRequest) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetEcrRegion

`func (o *HandlersPatchImageRegistryRequest) GetEcrRegion() string`

GetEcrRegion returns the EcrRegion field if non-nil, zero value otherwise.

### GetEcrRegionOk

`func (o *HandlersPatchImageRegistryRequest) GetEcrRegionOk() (*string, bool)`

GetEcrRegionOk returns a tuple with the EcrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcrRegion

`func (o *HandlersPatchImageRegistryRequest) SetEcrRegion(v string)`

SetEcrRegion sets EcrRegion field to given value.

### HasEcrRegion

`func (o *HandlersPatchImageRegistryRequest) HasEcrRegion() bool`

HasEcrRegion returns a boolean if a field has been set.

### GetGcpCredentialId

`func (o *HandlersPatchImageRegistryRequest) GetGcpCredentialId() int32`

GetGcpCredentialId returns the GcpCredentialId field if non-nil, zero value otherwise.

### GetGcpCredentialIdOk

`func (o *HandlersPatchImageRegistryRequest) GetGcpCredentialIdOk() (*int32, bool)`

GetGcpCredentialIdOk returns a tuple with the GcpCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredentialId

`func (o *HandlersPatchImageRegistryRequest) SetGcpCredentialId(v int32)`

SetGcpCredentialId sets GcpCredentialId field to given value.

### HasGcpCredentialId

`func (o *HandlersPatchImageRegistryRequest) HasGcpCredentialId() bool`

HasGcpCredentialId returns a boolean if a field has been set.

### GetLoginServer

`func (o *HandlersPatchImageRegistryRequest) GetLoginServer() string`

GetLoginServer returns the LoginServer field if non-nil, zero value otherwise.

### GetLoginServerOk

`func (o *HandlersPatchImageRegistryRequest) GetLoginServerOk() (*string, bool)`

GetLoginServerOk returns a tuple with the LoginServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginServer

`func (o *HandlersPatchImageRegistryRequest) SetLoginServer(v string)`

SetLoginServer sets LoginServer field to given value.

### HasLoginServer

`func (o *HandlersPatchImageRegistryRequest) HasLoginServer() bool`

HasLoginServer returns a boolean if a field has been set.

### GetPassword

`func (o *HandlersPatchImageRegistryRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HandlersPatchImageRegistryRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HandlersPatchImageRegistryRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HandlersPatchImageRegistryRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetScope

`func (o *HandlersPatchImageRegistryRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersPatchImageRegistryRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersPatchImageRegistryRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersPatchImageRegistryRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUsername

`func (o *HandlersPatchImageRegistryRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HandlersPatchImageRegistryRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HandlersPatchImageRegistryRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HandlersPatchImageRegistryRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


