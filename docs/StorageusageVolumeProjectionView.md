# StorageusageVolumeProjectionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | Pointer to [**StorageusageProjectionConfidence**](StorageusageProjectionConfidence.md) |  | [optional] 
**DaysToFull** | Pointer to **float32** |  | [optional] 
**Points** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**StorageusageProjectionStatus**](StorageusageProjectionStatus.md) |  | [optional] 

## Methods

### NewStorageusageVolumeProjectionView

`func NewStorageusageVolumeProjectionView() *StorageusageVolumeProjectionView`

NewStorageusageVolumeProjectionView instantiates a new StorageusageVolumeProjectionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageusageVolumeProjectionViewWithDefaults

`func NewStorageusageVolumeProjectionViewWithDefaults() *StorageusageVolumeProjectionView`

NewStorageusageVolumeProjectionViewWithDefaults instantiates a new StorageusageVolumeProjectionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *StorageusageVolumeProjectionView) GetConfidence() StorageusageProjectionConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *StorageusageVolumeProjectionView) GetConfidenceOk() (*StorageusageProjectionConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *StorageusageVolumeProjectionView) SetConfidence(v StorageusageProjectionConfidence)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *StorageusageVolumeProjectionView) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetDaysToFull

`func (o *StorageusageVolumeProjectionView) GetDaysToFull() float32`

GetDaysToFull returns the DaysToFull field if non-nil, zero value otherwise.

### GetDaysToFullOk

`func (o *StorageusageVolumeProjectionView) GetDaysToFullOk() (*float32, bool)`

GetDaysToFullOk returns a tuple with the DaysToFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToFull

`func (o *StorageusageVolumeProjectionView) SetDaysToFull(v float32)`

SetDaysToFull sets DaysToFull field to given value.

### HasDaysToFull

`func (o *StorageusageVolumeProjectionView) HasDaysToFull() bool`

HasDaysToFull returns a boolean if a field has been set.

### GetPoints

`func (o *StorageusageVolumeProjectionView) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *StorageusageVolumeProjectionView) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *StorageusageVolumeProjectionView) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *StorageusageVolumeProjectionView) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetStatus

`func (o *StorageusageVolumeProjectionView) GetStatus() StorageusageProjectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageusageVolumeProjectionView) GetStatusOk() (*StorageusageProjectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageusageVolumeProjectionView) SetStatus(v StorageusageProjectionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageusageVolumeProjectionView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


