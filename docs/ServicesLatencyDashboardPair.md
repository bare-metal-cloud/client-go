# ServicesLatencyDashboardPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Detail names the honest reason for a non-measured render. | [optional] 
**RttMs** | Pointer to **float32** | RTTms is the connect round trip in milliseconds (0 unless Status is measured or refused — the two classes a round trip backs). | [optional] 
**SourceVmId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | Status is the pair&#39;s HONESTY-ADJUSTED sample status: a stored sample whose endpoint is currently stopped or diverged degrades to LatMonPairNotMeasured at read time — RTTs are never served through a node whose plane or power state no longer backs them. | [optional] 
**TargetVmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewServicesLatencyDashboardPair

`func NewServicesLatencyDashboardPair() *ServicesLatencyDashboardPair`

NewServicesLatencyDashboardPair instantiates a new ServicesLatencyDashboardPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesLatencyDashboardPairWithDefaults

`func NewServicesLatencyDashboardPairWithDefaults() *ServicesLatencyDashboardPair`

NewServicesLatencyDashboardPairWithDefaults instantiates a new ServicesLatencyDashboardPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *ServicesLatencyDashboardPair) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ServicesLatencyDashboardPair) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ServicesLatencyDashboardPair) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ServicesLatencyDashboardPair) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetRttMs

`func (o *ServicesLatencyDashboardPair) GetRttMs() float32`

GetRttMs returns the RttMs field if non-nil, zero value otherwise.

### GetRttMsOk

`func (o *ServicesLatencyDashboardPair) GetRttMsOk() (*float32, bool)`

GetRttMsOk returns a tuple with the RttMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRttMs

`func (o *ServicesLatencyDashboardPair) SetRttMs(v float32)`

SetRttMs sets RttMs field to given value.

### HasRttMs

`func (o *ServicesLatencyDashboardPair) HasRttMs() bool`

HasRttMs returns a boolean if a field has been set.

### GetSourceVmId

`func (o *ServicesLatencyDashboardPair) GetSourceVmId() int32`

GetSourceVmId returns the SourceVmId field if non-nil, zero value otherwise.

### GetSourceVmIdOk

`func (o *ServicesLatencyDashboardPair) GetSourceVmIdOk() (*int32, bool)`

GetSourceVmIdOk returns a tuple with the SourceVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVmId

`func (o *ServicesLatencyDashboardPair) SetSourceVmId(v int32)`

SetSourceVmId sets SourceVmId field to given value.

### HasSourceVmId

`func (o *ServicesLatencyDashboardPair) HasSourceVmId() bool`

HasSourceVmId returns a boolean if a field has been set.

### GetStatus

`func (o *ServicesLatencyDashboardPair) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServicesLatencyDashboardPair) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServicesLatencyDashboardPair) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServicesLatencyDashboardPair) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetVmId

`func (o *ServicesLatencyDashboardPair) GetTargetVmId() int32`

GetTargetVmId returns the TargetVmId field if non-nil, zero value otherwise.

### GetTargetVmIdOk

`func (o *ServicesLatencyDashboardPair) GetTargetVmIdOk() (*int32, bool)`

GetTargetVmIdOk returns a tuple with the TargetVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmId

`func (o *ServicesLatencyDashboardPair) SetTargetVmId(v int32)`

SetTargetVmId sets TargetVmId field to given value.

### HasTargetVmId

`func (o *ServicesLatencyDashboardPair) HasTargetVmId() bool`

HasTargetVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


