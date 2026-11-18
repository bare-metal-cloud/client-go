# HandlersTestS3ConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** | azblob fields (issue #3576). AccessKey is the storage key; Account the storage account. | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**DefaultBucket** | Pointer to **string** | DefaultBucket lets the verify step check a named bucket. For gcs this is the native probe&#39;s Bucket.Attrs target (issue #3879): a credential scoped to one bucket verifies without project-wide list permission. | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**LinkedCredentialId** | Pointer to **int32** | LinkedCredentialID (issue #3874) verifies a link BEFORE create: the handler resolves the org-scoped credential through the vault (the vault gate lives in the handler) and probes with its secret. The wizard&#39;s verify step calls this before create, so a bad link fails there, not at create. | [optional] 
**Provider** | **string** |  | 
**Region** | Pointer to **string** |  | [optional] 
**SecretAccessKey** | Pointer to **string** |  | [optional] 
**ServiceAccountJson** | Pointer to **string** | ServiceAccountJSON is the gcs native credential (issue #3879). When set, the verify runs over the Cloud Storage API; when empty, a gcs request keeps the SigV4 path with the HMAC pair (legacy rows). | [optional] 

## Methods

### NewHandlersTestS3ConnectionRequest

`func NewHandlersTestS3ConnectionRequest(provider string, ) *HandlersTestS3ConnectionRequest`

NewHandlersTestS3ConnectionRequest instantiates a new HandlersTestS3ConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersTestS3ConnectionRequestWithDefaults

`func NewHandlersTestS3ConnectionRequestWithDefaults() *HandlersTestS3ConnectionRequest`

NewHandlersTestS3ConnectionRequestWithDefaults instantiates a new HandlersTestS3ConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *HandlersTestS3ConnectionRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *HandlersTestS3ConnectionRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *HandlersTestS3ConnectionRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *HandlersTestS3ConnectionRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *HandlersTestS3ConnectionRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *HandlersTestS3ConnectionRequest) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *HandlersTestS3ConnectionRequest) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *HandlersTestS3ConnectionRequest) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAccount

`func (o *HandlersTestS3ConnectionRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HandlersTestS3ConnectionRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HandlersTestS3ConnectionRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *HandlersTestS3ConnectionRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetContainer

`func (o *HandlersTestS3ConnectionRequest) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *HandlersTestS3ConnectionRequest) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *HandlersTestS3ConnectionRequest) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *HandlersTestS3ConnectionRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetDefaultBucket

`func (o *HandlersTestS3ConnectionRequest) GetDefaultBucket() string`

GetDefaultBucket returns the DefaultBucket field if non-nil, zero value otherwise.

### GetDefaultBucketOk

`func (o *HandlersTestS3ConnectionRequest) GetDefaultBucketOk() (*string, bool)`

GetDefaultBucketOk returns a tuple with the DefaultBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBucket

`func (o *HandlersTestS3ConnectionRequest) SetDefaultBucket(v string)`

SetDefaultBucket sets DefaultBucket field to given value.

### HasDefaultBucket

`func (o *HandlersTestS3ConnectionRequest) HasDefaultBucket() bool`

HasDefaultBucket returns a boolean if a field has been set.

### GetEndpoint

`func (o *HandlersTestS3ConnectionRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HandlersTestS3ConnectionRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HandlersTestS3ConnectionRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HandlersTestS3ConnectionRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetLinkedCredentialId

`func (o *HandlersTestS3ConnectionRequest) GetLinkedCredentialId() int32`

GetLinkedCredentialId returns the LinkedCredentialId field if non-nil, zero value otherwise.

### GetLinkedCredentialIdOk

`func (o *HandlersTestS3ConnectionRequest) GetLinkedCredentialIdOk() (*int32, bool)`

GetLinkedCredentialIdOk returns a tuple with the LinkedCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCredentialId

`func (o *HandlersTestS3ConnectionRequest) SetLinkedCredentialId(v int32)`

SetLinkedCredentialId sets LinkedCredentialId field to given value.

### HasLinkedCredentialId

`func (o *HandlersTestS3ConnectionRequest) HasLinkedCredentialId() bool`

HasLinkedCredentialId returns a boolean if a field has been set.

### GetProvider

`func (o *HandlersTestS3ConnectionRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersTestS3ConnectionRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersTestS3ConnectionRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRegion

`func (o *HandlersTestS3ConnectionRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersTestS3ConnectionRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersTestS3ConnectionRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HandlersTestS3ConnectionRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *HandlersTestS3ConnectionRequest) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *HandlersTestS3ConnectionRequest) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *HandlersTestS3ConnectionRequest) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *HandlersTestS3ConnectionRequest) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetServiceAccountJson

`func (o *HandlersTestS3ConnectionRequest) GetServiceAccountJson() string`

GetServiceAccountJson returns the ServiceAccountJson field if non-nil, zero value otherwise.

### GetServiceAccountJsonOk

`func (o *HandlersTestS3ConnectionRequest) GetServiceAccountJsonOk() (*string, bool)`

GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountJson

`func (o *HandlersTestS3ConnectionRequest) SetServiceAccountJson(v string)`

SetServiceAccountJson sets ServiceAccountJson field to given value.

### HasServiceAccountJson

`func (o *HandlersTestS3ConnectionRequest) HasServiceAccountJson() bool`

HasServiceAccountJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


