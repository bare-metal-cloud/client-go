# StoragebucketProjectionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | Pointer to [**StoragebucketConfidence**](StoragebucketConfidence.md) |  | [optional] 
**DaysToFull** | Pointer to **float32** |  | [optional] 
**Points** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**StoragebucketStatus**](StoragebucketStatus.md) |  | [optional] 

## Methods

### NewStoragebucketProjectionView

`func NewStoragebucketProjectionView() *StoragebucketProjectionView`

NewStoragebucketProjectionView instantiates a new StoragebucketProjectionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragebucketProjectionViewWithDefaults

`func NewStoragebucketProjectionViewWithDefaults() *StoragebucketProjectionView`

NewStoragebucketProjectionViewWithDefaults instantiates a new StoragebucketProjectionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *StoragebucketProjectionView) GetConfidence() StoragebucketConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *StoragebucketProjectionView) GetConfidenceOk() (*StoragebucketConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *StoragebucketProjectionView) SetConfidence(v StoragebucketConfidence)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *StoragebucketProjectionView) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetDaysToFull

`func (o *StoragebucketProjectionView) GetDaysToFull() float32`

GetDaysToFull returns the DaysToFull field if non-nil, zero value otherwise.

### GetDaysToFullOk

`func (o *StoragebucketProjectionView) GetDaysToFullOk() (*float32, bool)`

GetDaysToFullOk returns a tuple with the DaysToFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToFull

`func (o *StoragebucketProjectionView) SetDaysToFull(v float32)`

SetDaysToFull sets DaysToFull field to given value.

### HasDaysToFull

`func (o *StoragebucketProjectionView) HasDaysToFull() bool`

HasDaysToFull returns a boolean if a field has been set.

### GetPoints

`func (o *StoragebucketProjectionView) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *StoragebucketProjectionView) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *StoragebucketProjectionView) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *StoragebucketProjectionView) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetStatus

`func (o *StoragebucketProjectionView) GetStatus() StoragebucketStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StoragebucketProjectionView) GetStatusOk() (*StoragebucketStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StoragebucketProjectionView) SetStatus(v StoragebucketStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StoragebucketProjectionView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


