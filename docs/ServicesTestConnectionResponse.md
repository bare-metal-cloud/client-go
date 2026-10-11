# ServicesTestConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketCount** | Pointer to **int32** |  | [optional] 
**Buckets** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewServicesTestConnectionResponse

`func NewServicesTestConnectionResponse() *ServicesTestConnectionResponse`

NewServicesTestConnectionResponse instantiates a new ServicesTestConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesTestConnectionResponseWithDefaults

`func NewServicesTestConnectionResponseWithDefaults() *ServicesTestConnectionResponse`

NewServicesTestConnectionResponseWithDefaults instantiates a new ServicesTestConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketCount

`func (o *ServicesTestConnectionResponse) GetBucketCount() int32`

GetBucketCount returns the BucketCount field if non-nil, zero value otherwise.

### GetBucketCountOk

`func (o *ServicesTestConnectionResponse) GetBucketCountOk() (*int32, bool)`

GetBucketCountOk returns a tuple with the BucketCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketCount

`func (o *ServicesTestConnectionResponse) SetBucketCount(v int32)`

SetBucketCount sets BucketCount field to given value.

### HasBucketCount

`func (o *ServicesTestConnectionResponse) HasBucketCount() bool`

HasBucketCount returns a boolean if a field has been set.

### GetBuckets

`func (o *ServicesTestConnectionResponse) GetBuckets() []string`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ServicesTestConnectionResponse) GetBucketsOk() (*[]string, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ServicesTestConnectionResponse) SetBuckets(v []string)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *ServicesTestConnectionResponse) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetMessage

`func (o *ServicesTestConnectionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServicesTestConnectionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServicesTestConnectionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServicesTestConnectionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *ServicesTestConnectionResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ServicesTestConnectionResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ServicesTestConnectionResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ServicesTestConnectionResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


