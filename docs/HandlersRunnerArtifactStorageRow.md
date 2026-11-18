# HandlersRunnerArtifactStorageRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the ruling-14 default bucket&#39;s NAME — null when the org has no connected bucket (never a fabricated name). | [optional] 
**Retention** | Pointer to [**HandlersRunnerArtifactSettingsRow**](HandlersRunnerArtifactSettingsRow.md) | Retention is the org&#39;s effective setting in the settings row&#39;s shape (the usage card renders it beside the bytes). | [optional] 
**StoredBytes** | Pointer to **int32** | StoredBytes is SUM(size) over the org&#39;s state&#x3D;&#39;finalized&#39; rows only (expired rows&#39; bytes are released; created rows are transient chunks). | [optional] 

## Methods

### NewHandlersRunnerArtifactStorageRow

`func NewHandlersRunnerArtifactStorageRow() *HandlersRunnerArtifactStorageRow`

NewHandlersRunnerArtifactStorageRow instantiates a new HandlersRunnerArtifactStorageRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerArtifactStorageRowWithDefaults

`func NewHandlersRunnerArtifactStorageRowWithDefaults() *HandlersRunnerArtifactStorageRow`

NewHandlersRunnerArtifactStorageRowWithDefaults instantiates a new HandlersRunnerArtifactStorageRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *HandlersRunnerArtifactStorageRow) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *HandlersRunnerArtifactStorageRow) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *HandlersRunnerArtifactStorageRow) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *HandlersRunnerArtifactStorageRow) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetRetention

`func (o *HandlersRunnerArtifactStorageRow) GetRetention() HandlersRunnerArtifactSettingsRow`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *HandlersRunnerArtifactStorageRow) GetRetentionOk() (*HandlersRunnerArtifactSettingsRow, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *HandlersRunnerArtifactStorageRow) SetRetention(v HandlersRunnerArtifactSettingsRow)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *HandlersRunnerArtifactStorageRow) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetStoredBytes

`func (o *HandlersRunnerArtifactStorageRow) GetStoredBytes() int32`

GetStoredBytes returns the StoredBytes field if non-nil, zero value otherwise.

### GetStoredBytesOk

`func (o *HandlersRunnerArtifactStorageRow) GetStoredBytesOk() (*int32, bool)`

GetStoredBytesOk returns a tuple with the StoredBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredBytes

`func (o *HandlersRunnerArtifactStorageRow) SetStoredBytes(v int32)`

SetStoredBytes sets StoredBytes field to given value.

### HasStoredBytes

`func (o *HandlersRunnerArtifactStorageRow) HasStoredBytes() bool`

HasStoredBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


