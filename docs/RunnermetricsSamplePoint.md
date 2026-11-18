# RunnermetricsSamplePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapturedAt** | Pointer to **string** |  | [optional] 
**CpuMillicores** | Pointer to **int32** |  | [optional] 
**DiskBytes** | Pointer to **int32** |  | [optional] 
**MemoryBytes** | Pointer to **int32** |  | [optional] 
**Runner** | Pointer to **string** |  | [optional] 
**Slice** | Pointer to **string** |  | [optional] 
**WindowFromUnixMs** | Pointer to **int32** |  | [optional] 
**WindowToUnixMs** | Pointer to **int32** |  | [optional] 

## Methods

### NewRunnermetricsSamplePoint

`func NewRunnermetricsSamplePoint() *RunnermetricsSamplePoint`

NewRunnermetricsSamplePoint instantiates a new RunnermetricsSamplePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnermetricsSamplePointWithDefaults

`func NewRunnermetricsSamplePointWithDefaults() *RunnermetricsSamplePoint`

NewRunnermetricsSamplePointWithDefaults instantiates a new RunnermetricsSamplePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapturedAt

`func (o *RunnermetricsSamplePoint) GetCapturedAt() string`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *RunnermetricsSamplePoint) GetCapturedAtOk() (*string, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *RunnermetricsSamplePoint) SetCapturedAt(v string)`

SetCapturedAt sets CapturedAt field to given value.

### HasCapturedAt

`func (o *RunnermetricsSamplePoint) HasCapturedAt() bool`

HasCapturedAt returns a boolean if a field has been set.

### GetCpuMillicores

`func (o *RunnermetricsSamplePoint) GetCpuMillicores() int32`

GetCpuMillicores returns the CpuMillicores field if non-nil, zero value otherwise.

### GetCpuMillicoresOk

`func (o *RunnermetricsSamplePoint) GetCpuMillicoresOk() (*int32, bool)`

GetCpuMillicoresOk returns a tuple with the CpuMillicores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMillicores

`func (o *RunnermetricsSamplePoint) SetCpuMillicores(v int32)`

SetCpuMillicores sets CpuMillicores field to given value.

### HasCpuMillicores

`func (o *RunnermetricsSamplePoint) HasCpuMillicores() bool`

HasCpuMillicores returns a boolean if a field has been set.

### GetDiskBytes

`func (o *RunnermetricsSamplePoint) GetDiskBytes() int32`

GetDiskBytes returns the DiskBytes field if non-nil, zero value otherwise.

### GetDiskBytesOk

`func (o *RunnermetricsSamplePoint) GetDiskBytesOk() (*int32, bool)`

GetDiskBytesOk returns a tuple with the DiskBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskBytes

`func (o *RunnermetricsSamplePoint) SetDiskBytes(v int32)`

SetDiskBytes sets DiskBytes field to given value.

### HasDiskBytes

`func (o *RunnermetricsSamplePoint) HasDiskBytes() bool`

HasDiskBytes returns a boolean if a field has been set.

### GetMemoryBytes

`func (o *RunnermetricsSamplePoint) GetMemoryBytes() int32`

GetMemoryBytes returns the MemoryBytes field if non-nil, zero value otherwise.

### GetMemoryBytesOk

`func (o *RunnermetricsSamplePoint) GetMemoryBytesOk() (*int32, bool)`

GetMemoryBytesOk returns a tuple with the MemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryBytes

`func (o *RunnermetricsSamplePoint) SetMemoryBytes(v int32)`

SetMemoryBytes sets MemoryBytes field to given value.

### HasMemoryBytes

`func (o *RunnermetricsSamplePoint) HasMemoryBytes() bool`

HasMemoryBytes returns a boolean if a field has been set.

### GetRunner

`func (o *RunnermetricsSamplePoint) GetRunner() string`

GetRunner returns the Runner field if non-nil, zero value otherwise.

### GetRunnerOk

`func (o *RunnermetricsSamplePoint) GetRunnerOk() (*string, bool)`

GetRunnerOk returns a tuple with the Runner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunner

`func (o *RunnermetricsSamplePoint) SetRunner(v string)`

SetRunner sets Runner field to given value.

### HasRunner

`func (o *RunnermetricsSamplePoint) HasRunner() bool`

HasRunner returns a boolean if a field has been set.

### GetSlice

`func (o *RunnermetricsSamplePoint) GetSlice() string`

GetSlice returns the Slice field if non-nil, zero value otherwise.

### GetSliceOk

`func (o *RunnermetricsSamplePoint) GetSliceOk() (*string, bool)`

GetSliceOk returns a tuple with the Slice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlice

`func (o *RunnermetricsSamplePoint) SetSlice(v string)`

SetSlice sets Slice field to given value.

### HasSlice

`func (o *RunnermetricsSamplePoint) HasSlice() bool`

HasSlice returns a boolean if a field has been set.

### GetWindowFromUnixMs

`func (o *RunnermetricsSamplePoint) GetWindowFromUnixMs() int32`

GetWindowFromUnixMs returns the WindowFromUnixMs field if non-nil, zero value otherwise.

### GetWindowFromUnixMsOk

`func (o *RunnermetricsSamplePoint) GetWindowFromUnixMsOk() (*int32, bool)`

GetWindowFromUnixMsOk returns a tuple with the WindowFromUnixMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowFromUnixMs

`func (o *RunnermetricsSamplePoint) SetWindowFromUnixMs(v int32)`

SetWindowFromUnixMs sets WindowFromUnixMs field to given value.

### HasWindowFromUnixMs

`func (o *RunnermetricsSamplePoint) HasWindowFromUnixMs() bool`

HasWindowFromUnixMs returns a boolean if a field has been set.

### GetWindowToUnixMs

`func (o *RunnermetricsSamplePoint) GetWindowToUnixMs() int32`

GetWindowToUnixMs returns the WindowToUnixMs field if non-nil, zero value otherwise.

### GetWindowToUnixMsOk

`func (o *RunnermetricsSamplePoint) GetWindowToUnixMsOk() (*int32, bool)`

GetWindowToUnixMsOk returns a tuple with the WindowToUnixMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowToUnixMs

`func (o *RunnermetricsSamplePoint) SetWindowToUnixMs(v int32)`

SetWindowToUnixMs sets WindowToUnixMs field to given value.

### HasWindowToUnixMs

`func (o *RunnermetricsSamplePoint) HasWindowToUnixMs() bool`

HasWindowToUnixMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


