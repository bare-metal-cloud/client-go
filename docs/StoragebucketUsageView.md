# StoragebucketUsageView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** |  | [optional] 
**CapacityBytes** | Pointer to **int32** |  | [optional] 
**Latest** | Pointer to [**StoragebucketLatestView**](StoragebucketLatestView.md) |  | [optional] 
**Projection** | Pointer to [**StoragebucketProjectionView**](StoragebucketProjectionView.md) |  | [optional] 
**Trend** | Pointer to [**[]StoragebucketPoint**](StoragebucketPoint.md) |  | [optional] 

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

### GetCapacityBytes

`func (o *StoragebucketUsageView) GetCapacityBytes() int32`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *StoragebucketUsageView) GetCapacityBytesOk() (*int32, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *StoragebucketUsageView) SetCapacityBytes(v int32)`

SetCapacityBytes sets CapacityBytes field to given value.

### HasCapacityBytes

`func (o *StoragebucketUsageView) HasCapacityBytes() bool`

HasCapacityBytes returns a boolean if a field has been set.

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

### GetTrend

`func (o *StoragebucketUsageView) GetTrend() []StoragebucketPoint`

GetTrend returns the Trend field if non-nil, zero value otherwise.

### GetTrendOk

`func (o *StoragebucketUsageView) GetTrendOk() (*[]StoragebucketPoint, bool)`

GetTrendOk returns a tuple with the Trend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrend

`func (o *StoragebucketUsageView) SetTrend(v []StoragebucketPoint)`

SetTrend sets Trend field to given value.

### HasTrend

`func (o *StoragebucketUsageView) HasTrend() bool`

HasTrend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


