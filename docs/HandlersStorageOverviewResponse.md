# HandlersStorageOverviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveStorageAlerts** | Pointer to **int32** | ActiveStorageAlerts counts the org&#39;s unresolved storage_volume and storage_bucket alert events — the same active predicate both per-resource status reads use. | [optional] 
**BucketBytes** | Pointer to **int32** | BucketBytes sums every bucket&#39;s latest snapshot across the org&#39;s S3 connections. Null until a bucket snapshot carries bytes (LB-1 not landed), and never counting a partial snapshot (a lower bound). | [optional] 
**BucketsMeasured** | Pointer to **int32** |  | [optional] 
**VolumeProvisionedBytes** | Pointer to **int32** |  | [optional] 
**VolumeUsedBytes** | Pointer to **int32** | VolumeUsedBytes and VolumeProvisionedBytes sum the org&#39;s latest volume readings over one reading set: the volumes that reported BOTH figures. Null when no volume did. VolumesMeasured counts exactly that set, so the count and the two sums always name the same volumes. | [optional] 
**VolumesFullWithin14d** | Pointer to **int32** | VolumesFullWithin14d counts volumes whose fill projection lands within 14 days. A volume with no real projection (not enough history, no growth, no capacity) is never counted. | [optional] 
**VolumesMeasured** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersStorageOverviewResponse

`func NewHandlersStorageOverviewResponse() *HandlersStorageOverviewResponse`

NewHandlersStorageOverviewResponse instantiates a new HandlersStorageOverviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStorageOverviewResponseWithDefaults

`func NewHandlersStorageOverviewResponseWithDefaults() *HandlersStorageOverviewResponse`

NewHandlersStorageOverviewResponseWithDefaults instantiates a new HandlersStorageOverviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveStorageAlerts

`func (o *HandlersStorageOverviewResponse) GetActiveStorageAlerts() int32`

GetActiveStorageAlerts returns the ActiveStorageAlerts field if non-nil, zero value otherwise.

### GetActiveStorageAlertsOk

`func (o *HandlersStorageOverviewResponse) GetActiveStorageAlertsOk() (*int32, bool)`

GetActiveStorageAlertsOk returns a tuple with the ActiveStorageAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStorageAlerts

`func (o *HandlersStorageOverviewResponse) SetActiveStorageAlerts(v int32)`

SetActiveStorageAlerts sets ActiveStorageAlerts field to given value.

### HasActiveStorageAlerts

`func (o *HandlersStorageOverviewResponse) HasActiveStorageAlerts() bool`

HasActiveStorageAlerts returns a boolean if a field has been set.

### GetBucketBytes

`func (o *HandlersStorageOverviewResponse) GetBucketBytes() int32`

GetBucketBytes returns the BucketBytes field if non-nil, zero value otherwise.

### GetBucketBytesOk

`func (o *HandlersStorageOverviewResponse) GetBucketBytesOk() (*int32, bool)`

GetBucketBytesOk returns a tuple with the BucketBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketBytes

`func (o *HandlersStorageOverviewResponse) SetBucketBytes(v int32)`

SetBucketBytes sets BucketBytes field to given value.

### HasBucketBytes

`func (o *HandlersStorageOverviewResponse) HasBucketBytes() bool`

HasBucketBytes returns a boolean if a field has been set.

### GetBucketsMeasured

`func (o *HandlersStorageOverviewResponse) GetBucketsMeasured() int32`

GetBucketsMeasured returns the BucketsMeasured field if non-nil, zero value otherwise.

### GetBucketsMeasuredOk

`func (o *HandlersStorageOverviewResponse) GetBucketsMeasuredOk() (*int32, bool)`

GetBucketsMeasuredOk returns a tuple with the BucketsMeasured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketsMeasured

`func (o *HandlersStorageOverviewResponse) SetBucketsMeasured(v int32)`

SetBucketsMeasured sets BucketsMeasured field to given value.

### HasBucketsMeasured

`func (o *HandlersStorageOverviewResponse) HasBucketsMeasured() bool`

HasBucketsMeasured returns a boolean if a field has been set.

### GetVolumeProvisionedBytes

`func (o *HandlersStorageOverviewResponse) GetVolumeProvisionedBytes() int32`

GetVolumeProvisionedBytes returns the VolumeProvisionedBytes field if non-nil, zero value otherwise.

### GetVolumeProvisionedBytesOk

`func (o *HandlersStorageOverviewResponse) GetVolumeProvisionedBytesOk() (*int32, bool)`

GetVolumeProvisionedBytesOk returns a tuple with the VolumeProvisionedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeProvisionedBytes

`func (o *HandlersStorageOverviewResponse) SetVolumeProvisionedBytes(v int32)`

SetVolumeProvisionedBytes sets VolumeProvisionedBytes field to given value.

### HasVolumeProvisionedBytes

`func (o *HandlersStorageOverviewResponse) HasVolumeProvisionedBytes() bool`

HasVolumeProvisionedBytes returns a boolean if a field has been set.

### GetVolumeUsedBytes

`func (o *HandlersStorageOverviewResponse) GetVolumeUsedBytes() int32`

GetVolumeUsedBytes returns the VolumeUsedBytes field if non-nil, zero value otherwise.

### GetVolumeUsedBytesOk

`func (o *HandlersStorageOverviewResponse) GetVolumeUsedBytesOk() (*int32, bool)`

GetVolumeUsedBytesOk returns a tuple with the VolumeUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUsedBytes

`func (o *HandlersStorageOverviewResponse) SetVolumeUsedBytes(v int32)`

SetVolumeUsedBytes sets VolumeUsedBytes field to given value.

### HasVolumeUsedBytes

`func (o *HandlersStorageOverviewResponse) HasVolumeUsedBytes() bool`

HasVolumeUsedBytes returns a boolean if a field has been set.

### GetVolumesFullWithin14d

`func (o *HandlersStorageOverviewResponse) GetVolumesFullWithin14d() int32`

GetVolumesFullWithin14d returns the VolumesFullWithin14d field if non-nil, zero value otherwise.

### GetVolumesFullWithin14dOk

`func (o *HandlersStorageOverviewResponse) GetVolumesFullWithin14dOk() (*int32, bool)`

GetVolumesFullWithin14dOk returns a tuple with the VolumesFullWithin14d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesFullWithin14d

`func (o *HandlersStorageOverviewResponse) SetVolumesFullWithin14d(v int32)`

SetVolumesFullWithin14d sets VolumesFullWithin14d field to given value.

### HasVolumesFullWithin14d

`func (o *HandlersStorageOverviewResponse) HasVolumesFullWithin14d() bool`

HasVolumesFullWithin14d returns a boolean if a field has been set.

### GetVolumesMeasured

`func (o *HandlersStorageOverviewResponse) GetVolumesMeasured() int32`

GetVolumesMeasured returns the VolumesMeasured field if non-nil, zero value otherwise.

### GetVolumesMeasuredOk

`func (o *HandlersStorageOverviewResponse) GetVolumesMeasuredOk() (*int32, bool)`

GetVolumesMeasuredOk returns a tuple with the VolumesMeasured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesMeasured

`func (o *HandlersStorageOverviewResponse) SetVolumesMeasured(v int32)`

SetVolumesMeasured sets VolumesMeasured field to given value.

### HasVolumesMeasured

`func (o *HandlersStorageOverviewResponse) HasVolumesMeasured() bool`

HasVolumesMeasured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


