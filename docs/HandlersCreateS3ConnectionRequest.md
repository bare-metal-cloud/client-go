# HandlersCreateS3ConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** |  | 
**DefaultBucket** | Pointer to **string** |  | [optional] 
**Endpoint** | **string** |  | 
**Name** | **string** |  | 
**Provider** | **string** |  | 
**Region** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**SecretAccessKey** | **string** |  | 

## Methods

### NewHandlersCreateS3ConnectionRequest

`func NewHandlersCreateS3ConnectionRequest(accessKeyId string, endpoint string, name string, provider string, secretAccessKey string, ) *HandlersCreateS3ConnectionRequest`

NewHandlersCreateS3ConnectionRequest instantiates a new HandlersCreateS3ConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateS3ConnectionRequestWithDefaults

`func NewHandlersCreateS3ConnectionRequestWithDefaults() *HandlersCreateS3ConnectionRequest`

NewHandlersCreateS3ConnectionRequestWithDefaults instantiates a new HandlersCreateS3ConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


