# HandlersCreateOrgS3ConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** | azblob (Azure Blob) fields — issue #3576. | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**DefaultBucket** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**LinkedCredentialId** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Provider** | **string** |  | 
**Region** | Pointer to **string** |  | [optional] 
**SecretAccessKey** | Pointer to **string** |  | [optional] 
**ServiceAccountJson** | Pointer to **string** | ServiceAccountJSON is the gcs native credential (issue #3879): the plaintext JSON key for a service account, replacing the HMAC pair. | [optional] 

## Methods

### NewHandlersCreateOrgS3ConnectionRequest

`func NewHandlersCreateOrgS3ConnectionRequest(name string, provider string, ) *HandlersCreateOrgS3ConnectionRequest`

NewHandlersCreateOrgS3ConnectionRequest instantiates a new HandlersCreateOrgS3ConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateOrgS3ConnectionRequestWithDefaults

`func NewHandlersCreateOrgS3ConnectionRequestWithDefaults() *HandlersCreateOrgS3ConnectionRequest`

NewHandlersCreateOrgS3ConnectionRequestWithDefaults instantiates a new HandlersCreateOrgS3ConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *HandlersCreateOrgS3ConnectionRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *HandlersCreateOrgS3ConnectionRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *HandlersCreateOrgS3ConnectionRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *HandlersCreateOrgS3ConnectionRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *HandlersCreateOrgS3ConnectionRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *HandlersCreateOrgS3ConnectionRequest) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAccount

`func (o *HandlersCreateOrgS3ConnectionRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HandlersCreateOrgS3ConnectionRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *HandlersCreateOrgS3ConnectionRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetContainer

`func (o *HandlersCreateOrgS3ConnectionRequest) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *HandlersCreateOrgS3ConnectionRequest) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *HandlersCreateOrgS3ConnectionRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetDefaultBucket

`func (o *HandlersCreateOrgS3ConnectionRequest) GetDefaultBucket() string`

GetDefaultBucket returns the DefaultBucket field if non-nil, zero value otherwise.

### GetDefaultBucketOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetDefaultBucketOk() (*string, bool)`

GetDefaultBucketOk returns a tuple with the DefaultBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBucket

`func (o *HandlersCreateOrgS3ConnectionRequest) SetDefaultBucket(v string)`

SetDefaultBucket sets DefaultBucket field to given value.

### HasDefaultBucket

`func (o *HandlersCreateOrgS3ConnectionRequest) HasDefaultBucket() bool`

HasDefaultBucket returns a boolean if a field has been set.

### GetEndpoint

`func (o *HandlersCreateOrgS3ConnectionRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HandlersCreateOrgS3ConnectionRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HandlersCreateOrgS3ConnectionRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetLinkedCredentialId

`func (o *HandlersCreateOrgS3ConnectionRequest) GetLinkedCredentialId() int32`

GetLinkedCredentialId returns the LinkedCredentialId field if non-nil, zero value otherwise.

### GetLinkedCredentialIdOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetLinkedCredentialIdOk() (*int32, bool)`

GetLinkedCredentialIdOk returns a tuple with the LinkedCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCredentialId

`func (o *HandlersCreateOrgS3ConnectionRequest) SetLinkedCredentialId(v int32)`

SetLinkedCredentialId sets LinkedCredentialId field to given value.

### HasLinkedCredentialId

`func (o *HandlersCreateOrgS3ConnectionRequest) HasLinkedCredentialId() bool`

HasLinkedCredentialId returns a boolean if a field has been set.

### GetName

`func (o *HandlersCreateOrgS3ConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersCreateOrgS3ConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *HandlersCreateOrgS3ConnectionRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersCreateOrgS3ConnectionRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRegion

`func (o *HandlersCreateOrgS3ConnectionRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersCreateOrgS3ConnectionRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HandlersCreateOrgS3ConnectionRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *HandlersCreateOrgS3ConnectionRequest) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *HandlersCreateOrgS3ConnectionRequest) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *HandlersCreateOrgS3ConnectionRequest) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetServiceAccountJson

`func (o *HandlersCreateOrgS3ConnectionRequest) GetServiceAccountJson() string`

GetServiceAccountJson returns the ServiceAccountJson field if non-nil, zero value otherwise.

### GetServiceAccountJsonOk

`func (o *HandlersCreateOrgS3ConnectionRequest) GetServiceAccountJsonOk() (*string, bool)`

GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountJson

`func (o *HandlersCreateOrgS3ConnectionRequest) SetServiceAccountJson(v string)`

SetServiceAccountJson sets ServiceAccountJson field to given value.

### HasServiceAccountJson

`func (o *HandlersCreateOrgS3ConnectionRequest) HasServiceAccountJson() bool`

HasServiceAccountJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


