# HandlersMetricHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Granularity** | Pointer to **string** |  | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**Series** | Pointer to [**[]HandlersMetricHistorySeries**](HandlersMetricHistorySeries.md) |  | [optional] 
**VmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersMetricHistoryResponse

`func NewHandlersMetricHistoryResponse() *HandlersMetricHistoryResponse`

NewHandlersMetricHistoryResponse instantiates a new HandlersMetricHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMetricHistoryResponseWithDefaults

`func NewHandlersMetricHistoryResponseWithDefaults() *HandlersMetricHistoryResponse`

NewHandlersMetricHistoryResponseWithDefaults instantiates a new HandlersMetricHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGranularity

`func (o *HandlersMetricHistoryResponse) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *HandlersMetricHistoryResponse) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *HandlersMetricHistoryResponse) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *HandlersMetricHistoryResponse) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetRange

`func (o *HandlersMetricHistoryResponse) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *HandlersMetricHistoryResponse) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *HandlersMetricHistoryResponse) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *HandlersMetricHistoryResponse) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSeries

`func (o *HandlersMetricHistoryResponse) GetSeries() []HandlersMetricHistorySeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *HandlersMetricHistoryResponse) GetSeriesOk() (*[]HandlersMetricHistorySeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *HandlersMetricHistoryResponse) SetSeries(v []HandlersMetricHistorySeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *HandlersMetricHistoryResponse) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetVmId

`func (o *HandlersMetricHistoryResponse) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *HandlersMetricHistoryResponse) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *HandlersMetricHistoryResponse) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *HandlersMetricHistoryResponse) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


