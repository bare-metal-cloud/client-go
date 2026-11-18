# HandlersRunnerCloudKeyRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyIdMasked** | Pointer to **string** | AccessKeyIDMasked is the stored access key id with its middle replaced (the page renders the identifier, never a secret, never the raw id). | [optional] 
**Attached** | Pointer to **bool** | Attached marks the org&#39;s resolved key for its provider — the one the machine-create custody chain unseals into the spot adapter&#39;s per-exec config. | [optional] 
**AttachedPools** | Pointer to **[]string** | AttachedPools names the pool specs referencing this key (empty until the packing-declaration store lands — the attach is org-level until then). | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is the key&#39;s store stamp (unix seconds). | [optional] 
**KeyId** | Pointer to **int32** | KeyID is the stored key&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is the org&#39;s display label for the key. | [optional] 
**Provider** | Pointer to **string** | Provider is the cloud the key serves (aws today — the spot arm&#39;s provider vocabulary; gcp/azure are the design&#39;s accepted set but have no verifier, so they refuse at the verify-and-store). | [optional] 
**Region** | Pointer to **string** | Region is the cloud region the key serves (the adapter&#39;s config needs it for every machine identity). | [optional] 

## Methods

### NewHandlersRunnerCloudKeyRow

`func NewHandlersRunnerCloudKeyRow() *HandlersRunnerCloudKeyRow`

NewHandlersRunnerCloudKeyRow instantiates a new HandlersRunnerCloudKeyRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerCloudKeyRowWithDefaults

`func NewHandlersRunnerCloudKeyRowWithDefaults() *HandlersRunnerCloudKeyRow`

NewHandlersRunnerCloudKeyRowWithDefaults instantiates a new HandlersRunnerCloudKeyRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyIdMasked

`func (o *HandlersRunnerCloudKeyRow) GetAccessKeyIdMasked() string`

GetAccessKeyIdMasked returns the AccessKeyIdMasked field if non-nil, zero value otherwise.

### GetAccessKeyIdMaskedOk

`func (o *HandlersRunnerCloudKeyRow) GetAccessKeyIdMaskedOk() (*string, bool)`

GetAccessKeyIdMaskedOk returns a tuple with the AccessKeyIdMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyIdMasked

`func (o *HandlersRunnerCloudKeyRow) SetAccessKeyIdMasked(v string)`

SetAccessKeyIdMasked sets AccessKeyIdMasked field to given value.

### HasAccessKeyIdMasked

`func (o *HandlersRunnerCloudKeyRow) HasAccessKeyIdMasked() bool`

HasAccessKeyIdMasked returns a boolean if a field has been set.

### GetAttached

`func (o *HandlersRunnerCloudKeyRow) GetAttached() bool`

GetAttached returns the Attached field if non-nil, zero value otherwise.

### GetAttachedOk

`func (o *HandlersRunnerCloudKeyRow) GetAttachedOk() (*bool, bool)`

GetAttachedOk returns a tuple with the Attached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttached

`func (o *HandlersRunnerCloudKeyRow) SetAttached(v bool)`

SetAttached sets Attached field to given value.

### HasAttached

`func (o *HandlersRunnerCloudKeyRow) HasAttached() bool`

HasAttached returns a boolean if a field has been set.

### GetAttachedPools

`func (o *HandlersRunnerCloudKeyRow) GetAttachedPools() []string`

GetAttachedPools returns the AttachedPools field if non-nil, zero value otherwise.

### GetAttachedPoolsOk

`func (o *HandlersRunnerCloudKeyRow) GetAttachedPoolsOk() (*[]string, bool)`

GetAttachedPoolsOk returns a tuple with the AttachedPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedPools

`func (o *HandlersRunnerCloudKeyRow) SetAttachedPools(v []string)`

SetAttachedPools sets AttachedPools field to given value.

### HasAttachedPools

`func (o *HandlersRunnerCloudKeyRow) HasAttachedPools() bool`

HasAttachedPools returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersRunnerCloudKeyRow) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersRunnerCloudKeyRow) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersRunnerCloudKeyRow) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersRunnerCloudKeyRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetKeyId

`func (o *HandlersRunnerCloudKeyRow) GetKeyId() int32`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *HandlersRunnerCloudKeyRow) GetKeyIdOk() (*int32, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *HandlersRunnerCloudKeyRow) SetKeyId(v int32)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *HandlersRunnerCloudKeyRow) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetName

`func (o *HandlersRunnerCloudKeyRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersRunnerCloudKeyRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersRunnerCloudKeyRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersRunnerCloudKeyRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *HandlersRunnerCloudKeyRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersRunnerCloudKeyRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersRunnerCloudKeyRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *HandlersRunnerCloudKeyRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *HandlersRunnerCloudKeyRow) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersRunnerCloudKeyRow) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersRunnerCloudKeyRow) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HandlersRunnerCloudKeyRow) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


