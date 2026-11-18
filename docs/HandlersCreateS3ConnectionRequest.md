# HandlersCreateS3ConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** | azblob (Azure Blob) fields — issue #3576. Account is the storage account (the endpoint synthesizes from it), Container the blob container, and AccessKey the storage key (key mode).  LinkedCredentialID is link mode for EVERY provider (issue #3874): the connection borrows an org cloud credential and must carry no raw keys. Accepted pairs: aws/minio/wasabi/other link an aws credential, gcs links a gcp credential, azblob links an azure credential.  ServiceAccountJSON is the gcs native credential (issue #3879): paste the JSON key Google Cloud issues for a service account. It replaces the HMAC pair on new GCS connections; region still records the location and an endpoint may be omitted. | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**DefaultBucket** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**LinkedCredentialId** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Provider** | **string** |  | 
**Region** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**SecretAccessKey** | Pointer to **string** |  | [optional] 
**ServiceAccountJson** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersCreateS3ConnectionRequest

`func NewHandlersCreateS3ConnectionRequest(name string, provider string, ) *HandlersCreateS3ConnectionRequest`

NewHandlersCreateS3ConnectionRequest instantiates a new HandlersCreateS3ConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateS3ConnectionRequestWithDefaults

`func NewHandlersCreateS3ConnectionRequestWithDefaults() *HandlersCreateS3ConnectionRequest`

NewHandlersCreateS3ConnectionRequestWithDefaults instantiates a new HandlersCreateS3ConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *HandlersCreateS3ConnectionRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *HandlersCreateS3ConnectionRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *HandlersCreateS3ConnectionRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *HandlersCreateS3ConnectionRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *HandlersCreateS3ConnectionRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *HandlersCreateS3ConnectionRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *HandlersCreateS3ConnectionRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *HandlersCreateS3ConnectionRequest) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAccount

`func (o *HandlersCreateS3ConnectionRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HandlersCreateS3ConnectionRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HandlersCreateS3ConnectionRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *HandlersCreateS3ConnectionRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetContainer

`func (o *HandlersCreateS3ConnectionRequest) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *HandlersCreateS3ConnectionRequest) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *HandlersCreateS3ConnectionRequest) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *HandlersCreateS3ConnectionRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetDefaultBucket

`func (o *HandlersCreateS3ConnectionRequest) GetDefaultBucket() string`

GetDefaultBucket returns the DefaultBucket field if non-nil, zero value otherwise.

### GetDefaultBucketOk

`func (o *HandlersCreateS3ConnectionRequest) GetDefaultBucketOk() (*string, bool)`

GetDefaultBucketOk returns a tuple with the DefaultBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBucket

`func (o *HandlersCreateS3ConnectionRequest) SetDefaultBucket(v string)`

SetDefaultBucket sets DefaultBucket field to given value.

### HasDefaultBucket

`func (o *HandlersCreateS3ConnectionRequest) HasDefaultBucket() bool`

HasDefaultBucket returns a boolean if a field has been set.

### GetEndpoint

`func (o *HandlersCreateS3ConnectionRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HandlersCreateS3ConnectionRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HandlersCreateS3ConnectionRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HandlersCreateS3ConnectionRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetLinkedCredentialId

`func (o *HandlersCreateS3ConnectionRequest) GetLinkedCredentialId() int32`

GetLinkedCredentialId returns the LinkedCredentialId field if non-nil, zero value otherwise.

### GetLinkedCredentialIdOk

`func (o *HandlersCreateS3ConnectionRequest) GetLinkedCredentialIdOk() (*int32, bool)`

GetLinkedCredentialIdOk returns a tuple with the LinkedCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCredentialId

`func (o *HandlersCreateS3ConnectionRequest) SetLinkedCredentialId(v int32)`

SetLinkedCredentialId sets LinkedCredentialId field to given value.

### HasLinkedCredentialId

`func (o *HandlersCreateS3ConnectionRequest) HasLinkedCredentialId() bool`

HasLinkedCredentialId returns a boolean if a field has been set.

### GetName

`func (o *HandlersCreateS3ConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersCreateS3ConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersCreateS3ConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *HandlersCreateS3ConnectionRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersCreateS3ConnectionRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersCreateS3ConnectionRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRegion

`func (o *HandlersCreateS3ConnectionRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersCreateS3ConnectionRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersCreateS3ConnectionRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HandlersCreateS3ConnectionRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetScope

`func (o *HandlersCreateS3ConnectionRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersCreateS3ConnectionRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersCreateS3ConnectionRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersCreateS3ConnectionRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *HandlersCreateS3ConnectionRequest) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *HandlersCreateS3ConnectionRequest) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *HandlersCreateS3ConnectionRequest) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *HandlersCreateS3ConnectionRequest) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetServiceAccountJson

`func (o *HandlersCreateS3ConnectionRequest) GetServiceAccountJson() string`

GetServiceAccountJson returns the ServiceAccountJson field if non-nil, zero value otherwise.

### GetServiceAccountJsonOk

`func (o *HandlersCreateS3ConnectionRequest) GetServiceAccountJsonOk() (*string, bool)`

GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountJson

`func (o *HandlersCreateS3ConnectionRequest) SetServiceAccountJson(v string)`

SetServiceAccountJson sets ServiceAccountJson field to given value.

### HasServiceAccountJson

`func (o *HandlersCreateS3ConnectionRequest) HasServiceAccountJson() bool`

HasServiceAccountJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


