# ServicesS3ObjectListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** |  | [optional] 
**CommonPrefixes** | Pointer to **[]string** |  | [optional] 
**IsTruncated** | Pointer to **bool** |  | [optional] 
**NextContinuationToken** | Pointer to **string** |  | [optional] 
**Objects** | Pointer to [**[]ServicesS3ObjectInfo**](ServicesS3ObjectInfo.md) |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 

## Methods

### NewServicesS3ObjectListResult

`func NewServicesS3ObjectListResult() *ServicesS3ObjectListResult`

NewServicesS3ObjectListResult instantiates a new ServicesS3ObjectListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesS3ObjectListResultWithDefaults

`func NewServicesS3ObjectListResultWithDefaults() *ServicesS3ObjectListResult`

NewServicesS3ObjectListResultWithDefaults instantiates a new ServicesS3ObjectListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *ServicesS3ObjectListResult) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ServicesS3ObjectListResult) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ServicesS3ObjectListResult) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ServicesS3ObjectListResult) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCommonPrefixes

`func (o *ServicesS3ObjectListResult) GetCommonPrefixes() []string`

GetCommonPrefixes returns the CommonPrefixes field if non-nil, zero value otherwise.

### GetCommonPrefixesOk

`func (o *ServicesS3ObjectListResult) GetCommonPrefixesOk() (*[]string, bool)`

GetCommonPrefixesOk returns a tuple with the CommonPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonPrefixes

`func (o *ServicesS3ObjectListResult) SetCommonPrefixes(v []string)`

SetCommonPrefixes sets CommonPrefixes field to given value.

### HasCommonPrefixes

`func (o *ServicesS3ObjectListResult) HasCommonPrefixes() bool`

HasCommonPrefixes returns a boolean if a field has been set.

### GetIsTruncated

`func (o *ServicesS3ObjectListResult) GetIsTruncated() bool`

GetIsTruncated returns the IsTruncated field if non-nil, zero value otherwise.

### GetIsTruncatedOk

`func (o *ServicesS3ObjectListResult) GetIsTruncatedOk() (*bool, bool)`

GetIsTruncatedOk returns a tuple with the IsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTruncated

`func (o *ServicesS3ObjectListResult) SetIsTruncated(v bool)`

SetIsTruncated sets IsTruncated field to given value.

### HasIsTruncated

`func (o *ServicesS3ObjectListResult) HasIsTruncated() bool`

HasIsTruncated returns a boolean if a field has been set.

### GetNextContinuationToken

`func (o *ServicesS3ObjectListResult) GetNextContinuationToken() string`

GetNextContinuationToken returns the NextContinuationToken field if non-nil, zero value otherwise.

### GetNextContinuationTokenOk

`func (o *ServicesS3ObjectListResult) GetNextContinuationTokenOk() (*string, bool)`

GetNextContinuationTokenOk returns a tuple with the NextContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextContinuationToken

`func (o *ServicesS3ObjectListResult) SetNextContinuationToken(v string)`

SetNextContinuationToken sets NextContinuationToken field to given value.

### HasNextContinuationToken

`func (o *ServicesS3ObjectListResult) HasNextContinuationToken() bool`

HasNextContinuationToken returns a boolean if a field has been set.

### GetObjects

`func (o *ServicesS3ObjectListResult) GetObjects() []ServicesS3ObjectInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ServicesS3ObjectListResult) GetObjectsOk() (*[]ServicesS3ObjectInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ServicesS3ObjectListResult) SetObjects(v []ServicesS3ObjectInfo)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ServicesS3ObjectListResult) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetPrefix

`func (o *ServicesS3ObjectListResult) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ServicesS3ObjectListResult) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ServicesS3ObjectListResult) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ServicesS3ObjectListResult) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


