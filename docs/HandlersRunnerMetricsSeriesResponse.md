# HandlersRunnerMetricsSeriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Samples** | Pointer to [**[]RunnermetricsSamplePoint**](RunnermetricsSamplePoint.md) |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 

## Methods

### NewHandlersRunnerMetricsSeriesResponse

`func NewHandlersRunnerMetricsSeriesResponse() *HandlersRunnerMetricsSeriesResponse`

NewHandlersRunnerMetricsSeriesResponse instantiates a new HandlersRunnerMetricsSeriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerMetricsSeriesResponseWithDefaults

`func NewHandlersRunnerMetricsSeriesResponseWithDefaults() *HandlersRunnerMetricsSeriesResponse`

NewHandlersRunnerMetricsSeriesResponseWithDefaults instantiates a new HandlersRunnerMetricsSeriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *HandlersRunnerMetricsSeriesResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HandlersRunnerMetricsSeriesResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HandlersRunnerMetricsSeriesResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HandlersRunnerMetricsSeriesResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSamples

`func (o *HandlersRunnerMetricsSeriesResponse) GetSamples() []RunnermetricsSamplePoint`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *HandlersRunnerMetricsSeriesResponse) GetSamplesOk() (*[]RunnermetricsSamplePoint, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *HandlersRunnerMetricsSeriesResponse) SetSamples(v []RunnermetricsSamplePoint)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *HandlersRunnerMetricsSeriesResponse) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetTruncated

`func (o *HandlersRunnerMetricsSeriesResponse) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *HandlersRunnerMetricsSeriesResponse) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *HandlersRunnerMetricsSeriesResponse) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *HandlersRunnerMetricsSeriesResponse) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


