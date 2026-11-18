# ContractsMetricSample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtUnixMs** | Pointer to **int32** |  | [optional] 
**CpuMillicores** | Pointer to **int32** | integer millicores — no floats on the canonical wire | [optional] 
**DiskBytes** | Pointer to **int32** |  | [optional] 
**MemoryBytes** | Pointer to **int32** |  | [optional] 
**Runner** | Pointer to **string** | the RunnerHandle id | [optional] 
**Slice** | Pointer to **string** | the slice identity on the machine | [optional] 

## Methods

### NewContractsMetricSample

`func NewContractsMetricSample() *ContractsMetricSample`

NewContractsMetricSample instantiates a new ContractsMetricSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractsMetricSampleWithDefaults

`func NewContractsMetricSampleWithDefaults() *ContractsMetricSample`

NewContractsMetricSampleWithDefaults instantiates a new ContractsMetricSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtUnixMs

`func (o *ContractsMetricSample) GetAtUnixMs() int32`

GetAtUnixMs returns the AtUnixMs field if non-nil, zero value otherwise.

### GetAtUnixMsOk

`func (o *ContractsMetricSample) GetAtUnixMsOk() (*int32, bool)`

GetAtUnixMsOk returns a tuple with the AtUnixMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtUnixMs

`func (o *ContractsMetricSample) SetAtUnixMs(v int32)`

SetAtUnixMs sets AtUnixMs field to given value.

### HasAtUnixMs

`func (o *ContractsMetricSample) HasAtUnixMs() bool`

HasAtUnixMs returns a boolean if a field has been set.

### GetCpuMillicores

`func (o *ContractsMetricSample) GetCpuMillicores() int32`

GetCpuMillicores returns the CpuMillicores field if non-nil, zero value otherwise.

### GetCpuMillicoresOk

`func (o *ContractsMetricSample) GetCpuMillicoresOk() (*int32, bool)`

GetCpuMillicoresOk returns a tuple with the CpuMillicores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMillicores

`func (o *ContractsMetricSample) SetCpuMillicores(v int32)`

SetCpuMillicores sets CpuMillicores field to given value.

### HasCpuMillicores

`func (o *ContractsMetricSample) HasCpuMillicores() bool`

HasCpuMillicores returns a boolean if a field has been set.

### GetDiskBytes

`func (o *ContractsMetricSample) GetDiskBytes() int32`

GetDiskBytes returns the DiskBytes field if non-nil, zero value otherwise.

### GetDiskBytesOk

`func (o *ContractsMetricSample) GetDiskBytesOk() (*int32, bool)`

GetDiskBytesOk returns a tuple with the DiskBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskBytes

`func (o *ContractsMetricSample) SetDiskBytes(v int32)`

SetDiskBytes sets DiskBytes field to given value.

### HasDiskBytes

`func (o *ContractsMetricSample) HasDiskBytes() bool`

HasDiskBytes returns a boolean if a field has been set.

### GetMemoryBytes

`func (o *ContractsMetricSample) GetMemoryBytes() int32`

GetMemoryBytes returns the MemoryBytes field if non-nil, zero value otherwise.

### GetMemoryBytesOk

`func (o *ContractsMetricSample) GetMemoryBytesOk() (*int32, bool)`

GetMemoryBytesOk returns a tuple with the MemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBytes

`func (o *ContractsMetricSample) SetMemoryBytes(v int32)`

SetMemoryBytes sets MemoryBytes field to given value.

### HasMemoryBytes

`func (o *ContractsMetricSample) HasMemoryBytes() bool`

HasMemoryBytes returns a boolean if a field has been set.

### GetRunner

`func (o *ContractsMetricSample) GetRunner() string`

GetRunner returns the Runner field if non-nil, zero value otherwise.

### GetRunnerOk

`func (o *ContractsMetricSample) GetRunnerOk() (*string, bool)`

GetRunnerOk returns a tuple with the Runner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunner

`func (o *ContractsMetricSample) SetRunner(v string)`

SetRunner sets Runner field to given value.

### HasRunner

`func (o *ContractsMetricSample) HasRunner() bool`

HasRunner returns a boolean if a field has been set.

### GetSlice

`func (o *ContractsMetricSample) GetSlice() string`

GetSlice returns the Slice field if non-nil, zero value otherwise.

### GetSliceOk

`func (o *ContractsMetricSample) GetSliceOk() (*string, bool)`

GetSliceOk returns a tuple with the Slice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlice

`func (o *ContractsMetricSample) SetSlice(v string)`

SetSlice sets Slice field to given value.

### HasSlice

`func (o *ContractsMetricSample) HasSlice() bool`

HasSlice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


