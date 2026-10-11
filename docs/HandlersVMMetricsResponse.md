# HandlersVMMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to **map[string]float32** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**VmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersVMMetricsResponse

`func NewHandlersVMMetricsResponse() *HandlersVMMetricsResponse`

NewHandlersVMMetricsResponse instantiates a new HandlersVMMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersVMMetricsResponseWithDefaults

`func NewHandlersVMMetricsResponseWithDefaults() *HandlersVMMetricsResponse`

NewHandlersVMMetricsResponseWithDefaults instantiates a new HandlersVMMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *HandlersVMMetricsResponse) GetMetrics() map[string]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *HandlersVMMetricsResponse) GetMetricsOk() (*map[string]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *HandlersVMMetricsResponse) SetMetrics(v map[string]float32)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *HandlersVMMetricsResponse) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetTimestamp

`func (o *HandlersVMMetricsResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HandlersVMMetricsResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HandlersVMMetricsResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HandlersVMMetricsResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVmId

`func (o *HandlersVMMetricsResponse) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *HandlersVMMetricsResponse) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *HandlersVMMetricsResponse) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *HandlersVMMetricsResponse) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


