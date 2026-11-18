# HandlersRunnerCloudKeySaveBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** | AccessKeyID is the key&#39;s public identifier half. | [optional] 
**Name** | Pointer to **string** | Name is the org&#39;s display label for the key. | [optional] 
**Provider** | Pointer to **string** | Provider is the cloud the key serves (aws today — the only provider with a live verifier). | [optional] 
**Region** | Pointer to **string** | Region is the cloud region the key serves (required for aws). | [optional] 
**SecretAccessKey** | Pointer to **string** | SecretAccessKey is the key&#39;s secret half. TLS-only: the live check consumes it, the seal stores it encrypted, nothing logs it. | [optional] 

## Methods

### NewHandlersRunnerCloudKeySaveBody

`func NewHandlersRunnerCloudKeySaveBody() *HandlersRunnerCloudKeySaveBody`

NewHandlersRunnerCloudKeySaveBody instantiates a new HandlersRunnerCloudKeySaveBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerCloudKeySaveBodyWithDefaults

`func NewHandlersRunnerCloudKeySaveBodyWithDefaults() *HandlersRunnerCloudKeySaveBody`

NewHandlersRunnerCloudKeySaveBodyWithDefaults instantiates a new HandlersRunnerCloudKeySaveBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *HandlersRunnerCloudKeySaveBody) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *HandlersRunnerCloudKeySaveBody) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *HandlersRunnerCloudKeySaveBody) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *HandlersRunnerCloudKeySaveBody) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetName

`func (o *HandlersRunnerCloudKeySaveBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersRunnerCloudKeySaveBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersRunnerCloudKeySaveBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersRunnerCloudKeySaveBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *HandlersRunnerCloudKeySaveBody) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersRunnerCloudKeySaveBody) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersRunnerCloudKeySaveBody) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *HandlersRunnerCloudKeySaveBody) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *HandlersRunnerCloudKeySaveBody) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersRunnerCloudKeySaveBody) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersRunnerCloudKeySaveBody) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HandlersRunnerCloudKeySaveBody) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *HandlersRunnerCloudKeySaveBody) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *HandlersRunnerCloudKeySaveBody) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *HandlersRunnerCloudKeySaveBody) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *HandlersRunnerCloudKeySaveBody) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


