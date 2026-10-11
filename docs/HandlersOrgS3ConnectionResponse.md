# HandlersOrgS3ConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to **map[string]string** | Credentials is ALWAYS a masked view. The list path does not even decrypt (AC5); the create/edit paths return this same minimal placeholder. | [optional] 
**DefaultBucket** | Pointer to **map[string]interface{}** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**HasVaultCredentials** | Pointer to **bool** | HasVaultCredentials tells the FE whether a \&quot;rotate\&quot; affordance applies (104.4). It does NOT leak the secret. | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastUsedAt** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersOrgS3ConnectionResponse

`func NewHandlersOrgS3ConnectionResponse() *HandlersOrgS3ConnectionResponse`

NewHandlersOrgS3ConnectionResponse instantiates a new HandlersOrgS3ConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersOrgS3ConnectionResponseWithDefaults

`func NewHandlersOrgS3ConnectionResponseWithDefaults() *HandlersOrgS3ConnectionResponse`

NewHandlersOrgS3ConnectionResponseWithDefaults instantiates a new HandlersOrgS3ConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HandlersOrgS3ConnectionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersOrgS3ConnectionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersOrgS3ConnectionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersOrgS3ConnectionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentials

`func (o *HandlersOrgS3ConnectionResponse) GetCredentials() map[string]string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *HandlersOrgS3ConnectionResponse) GetCredentialsOk() (*map[string]string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *HandlersOrgS3ConnectionResponse) SetCredentials(v map[string]string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *HandlersOrgS3ConnectionResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDefaultBucket

`func (o *HandlersOrgS3ConnectionResponse) GetDefaultBucket() map[string]interface{}`

GetDefaultBucket returns the DefaultBucket field if non-nil, zero value otherwise.

### GetDefaultBucketOk

`func (o *HandlersOrgS3ConnectionResponse) GetDefaultBucketOk() (*map[string]interface{}, bool)`

GetDefaultBucketOk returns a tuple with the DefaultBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBucket

`func (o *HandlersOrgS3ConnectionResponse) SetDefaultBucket(v map[string]interface{})`

SetDefaultBucket sets DefaultBucket field to given value.

### HasDefaultBucket

`func (o *HandlersOrgS3ConnectionResponse) HasDefaultBucket() bool`

HasDefaultBucket returns a boolean if a field has been set.

### GetEndpoint

`func (o *HandlersOrgS3ConnectionResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *HandlersOrgS3ConnectionResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *HandlersOrgS3ConnectionResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *HandlersOrgS3ConnectionResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetHasVaultCredentials

`func (o *HandlersOrgS3ConnectionResponse) GetHasVaultCredentials() bool`

GetHasVaultCredentials returns the HasVaultCredentials field if non-nil, zero value otherwise.

### GetHasVaultCredentialsOk

`func (o *HandlersOrgS3ConnectionResponse) GetHasVaultCredentialsOk() (*bool, bool)`

GetHasVaultCredentialsOk returns a tuple with the HasVaultCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVaultCredentials

`func (o *HandlersOrgS3ConnectionResponse) SetHasVaultCredentials(v bool)`

SetHasVaultCredentials sets HasVaultCredentials field to given value.

### HasHasVaultCredentials

`func (o *HandlersOrgS3ConnectionResponse) HasHasVaultCredentials() bool`

HasHasVaultCredentials returns a boolean if a field has been set.

### GetId

`func (o *HandlersOrgS3ConnectionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlersOrgS3ConnectionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlersOrgS3ConnectionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HandlersOrgS3ConnectionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *HandlersOrgS3ConnectionResponse) GetLastUsedAt() map[string]interface{}`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *HandlersOrgS3ConnectionResponse) GetLastUsedAtOk() (*map[string]interface{}, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *HandlersOrgS3ConnectionResponse) SetLastUsedAt(v map[string]interface{})`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *HandlersOrgS3ConnectionResponse) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetName

`func (o *HandlersOrgS3ConnectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersOrgS3ConnectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersOrgS3ConnectionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersOrgS3ConnectionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *HandlersOrgS3ConnectionResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersOrgS3ConnectionResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersOrgS3ConnectionResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *HandlersOrgS3ConnectionResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HandlersOrgS3ConnectionResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HandlersOrgS3ConnectionResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HandlersOrgS3ConnectionResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HandlersOrgS3ConnectionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


