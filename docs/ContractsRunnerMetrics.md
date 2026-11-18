# ContractsRunnerMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | Pointer to [**[]ContractsMetricSample**](ContractsMetricSample.md) |  | [optional] 
**Window** | Pointer to [**ContractsMetricsWindow**](ContractsMetricsWindow.md) |  | [optional] 

## Methods

### NewContractsRunnerMetrics

`func NewContractsRunnerMetrics() *ContractsRunnerMetrics`

NewContractsRunnerMetrics instantiates a new ContractsRunnerMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractsRunnerMetricsWithDefaults

`func NewContractsRunnerMetricsWithDefaults() *ContractsRunnerMetrics`

NewContractsRunnerMetricsWithDefaults instantiates a new ContractsRunnerMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *ContractsRunnerMetrics) GetSamples() []ContractsMetricSample`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *ContractsRunnerMetrics) GetSamplesOk() (*[]ContractsMetricSample, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *ContractsRunnerMetrics) SetSamples(v []ContractsMetricSample)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *ContractsRunnerMetrics) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetWindow

`func (o *ContractsRunnerMetrics) GetWindow() ContractsMetricsWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ContractsRunnerMetrics) GetWindowOk() (*ContractsMetricsWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ContractsRunnerMetrics) SetWindow(v ContractsMetricsWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *ContractsRunnerMetrics) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


