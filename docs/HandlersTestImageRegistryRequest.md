# HandlersTestImageRegistryRequest

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
**Type** | **string** |  | 
**Url** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersTestImageRegistryRequest

`func NewHandlersTestImageRegistryRequest(type_ string, url string, ) *HandlersTestImageRegistryRequest`

NewHandlersTestImageRegistryRequest instantiates a new HandlersTestImageRegistryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersTestImageRegistryRequestWithDefaults

`func NewHandlersTestImageRegistryRequestWithDefaults() *HandlersTestImageRegistryRequest`

NewHandlersTestImageRegistryRequestWithDefaults instantiates a new HandlersTestImageRegistryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *HandlersTestImageRegistryRequest) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *HandlersTestImageRegistryRequest) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *HandlersTestImageRegistryRequest) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *HandlersTestImageRegistryRequest) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAwsCredentialId

`func (o *HandlersTestImageRegistryRequest) GetAwsCredentialId() int32`

GetAwsCredentialId returns the AwsCredentialId field if non-nil, zero value otherwise.

### GetAwsCredentialIdOk

`func (o *HandlersTestImageRegistryRequest) GetAwsCredentialIdOk() (*int32, bool)`

GetAwsCredentialIdOk returns a tuple with the AwsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCredentialId

`func (o *HandlersTestImageRegistryRequest) SetAwsCredentialId(v int32)`

SetAwsCredentialId sets AwsCredentialId field to given value.

### HasAwsCredentialId

`func (o *HandlersTestImageRegistryRequest) HasAwsCredentialId() bool`

HasAwsCredentialId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *HandlersTestImageRegistryRequest) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *HandlersTestImageRegistryRequest) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *HandlersTestImageRegistryRequest) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *HandlersTestImageRegistryRequest) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAzureCredentialId

`func (o *HandlersTestImageRegistryRequest) GetAzureCredentialId() int32`

GetAzureCredentialId returns the AzureCredentialId field if non-nil, zero value otherwise.

### GetAzureCredentialIdOk

`func (o *HandlersTestImageRegistryRequest) GetAzureCredentialIdOk() (*int32, bool)`

GetAzureCredentialIdOk returns a tuple with the AzureCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCredentialId

`func (o *HandlersTestImageRegistryRequest) SetAzureCredentialId(v int32)`

SetAzureCredentialId sets AzureCredentialId field to given value.

### HasAzureCredentialId

`func (o *HandlersTestImageRegistryRequest) HasAzureCredentialId() bool`

HasAzureCredentialId returns a boolean if a field has been set.

### GetCaCertificate

`func (o *HandlersTestImageRegistryRequest) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *HandlersTestImageRegistryRequest) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *HandlersTestImageRegistryRequest) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *HandlersTestImageRegistryRequest) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetEcrRegion

`func (o *HandlersTestImageRegistryRequest) GetEcrRegion() string`

GetEcrRegion returns the EcrRegion field if non-nil, zero value otherwise.

### GetEcrRegionOk

`func (o *HandlersTestImageRegistryRequest) GetEcrRegionOk() (*string, bool)`

GetEcrRegionOk returns a tuple with the EcrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcrRegion

`func (o *HandlersTestImageRegistryRequest) SetEcrRegion(v string)`

SetEcrRegion sets EcrRegion field to given value.

### HasEcrRegion

`func (o *HandlersTestImageRegistryRequest) HasEcrRegion() bool`

HasEcrRegion returns a boolean if a field has been set.

### GetGcpCredentialId

`func (o *HandlersTestImageRegistryRequest) GetGcpCredentialId() int32`

GetGcpCredentialId returns the GcpCredentialId field if non-nil, zero value otherwise.

### GetGcpCredentialIdOk

`func (o *HandlersTestImageRegistryRequest) GetGcpCredentialIdOk() (*int32, bool)`

GetGcpCredentialIdOk returns a tuple with the GcpCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredentialId

`func (o *HandlersTestImageRegistryRequest) SetGcpCredentialId(v int32)`

SetGcpCredentialId sets GcpCredentialId field to given value.

### HasGcpCredentialId

`func (o *HandlersTestImageRegistryRequest) HasGcpCredentialId() bool`

HasGcpCredentialId returns a boolean if a field has been set.

### GetLoginServer

`func (o *HandlersTestImageRegistryRequest) GetLoginServer() string`

GetLoginServer returns the LoginServer field if non-nil, zero value otherwise.

### GetLoginServerOk

`func (o *HandlersTestImageRegistryRequest) GetLoginServerOk() (*string, bool)`

GetLoginServerOk returns a tuple with the LoginServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginServer

`func (o *HandlersTestImageRegistryRequest) SetLoginServer(v string)`

SetLoginServer sets LoginServer field to given value.

### HasLoginServer

`func (o *HandlersTestImageRegistryRequest) HasLoginServer() bool`

HasLoginServer returns a boolean if a field has been set.

### GetPassword

`func (o *HandlersTestImageRegistryRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HandlersTestImageRegistryRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HandlersTestImageRegistryRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HandlersTestImageRegistryRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetType

`func (o *HandlersTestImageRegistryRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HandlersTestImageRegistryRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HandlersTestImageRegistryRequest) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *HandlersTestImageRegistryRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HandlersTestImageRegistryRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HandlersTestImageRegistryRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *HandlersTestImageRegistryRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HandlersTestImageRegistryRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HandlersTestImageRegistryRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HandlersTestImageRegistryRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


