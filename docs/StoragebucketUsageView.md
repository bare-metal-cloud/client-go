# StoragebucketUsageView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** |  | [optional] 
**Latest** | Pointer to [**StoragebucketLatestView**](StoragebucketLatestView.md) |  | [optional] 
**Projection** | Pointer to [**StoragebucketProjectionView**](StoragebucketProjectionView.md) |  | [optional] 

## Methods

### NewStoragebucketUsageView

`func NewStoragebucketUsageView() *StoragebucketUsageView`

NewStoragebucketUsageView instantiates a new StoragebucketUsageView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragebucketUsageViewWithDefaults

`func NewStoragebucketUsageViewWithDefaults() *StoragebucketUsageView`

NewStoragebucketUsageViewWithDefaults instantiates a new StoragebucketUsageView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *StoragebucketUsageView) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *StoragebucketUsageView) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *StoragebucketUsageView) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *StoragebucketUsageView) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetLatest

`func (o *StoragebucketUsageView) GetLatest() StoragebucketLatestView`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *StoragebucketUsageView) GetLatestOk() (*StoragebucketLatestView, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *StoragebucketUsageView) SetLatest(v StoragebucketLatestView)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *StoragebucketUsageView) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetProjection

`func (o *StoragebucketUsageView) GetProjection() StoragebucketProjectionView`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *StoragebucketUsageView) GetProjectionOk() (*StoragebucketProjectionView, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *StoragebucketUsageView) SetProjection(v StoragebucketProjectionView)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *StoragebucketUsageView) HasProjection() bool`

HasProjection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


