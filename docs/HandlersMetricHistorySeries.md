# HandlersMetricHistorySeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** |  | [optional] 
**Points** | Pointer to [**[]HandlersMetricHistoryPoint**](HandlersMetricHistoryPoint.md) |  | [optional] 

## Methods

### NewHandlersMetricHistorySeries

`func NewHandlersMetricHistorySeries() *HandlersMetricHistorySeries`

NewHandlersMetricHistorySeries instantiates a new HandlersMetricHistorySeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMetricHistorySeriesWithDefaults

`func NewHandlersMetricHistorySeriesWithDefaults() *HandlersMetricHistorySeries`

NewHandlersMetricHistorySeriesWithDefaults instantiates a new HandlersMetricHistorySeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *HandlersMetricHistorySeries) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *HandlersMetricHistorySeries) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *HandlersMetricHistorySeries) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *HandlersMetricHistorySeries) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPoints

`func (o *HandlersMetricHistorySeries) GetPoints() []HandlersMetricHistoryPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *HandlersMetricHistorySeries) GetPointsOk() (*[]HandlersMetricHistoryPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *HandlersMetricHistorySeries) SetPoints(v []HandlersMetricHistoryPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *HandlersMetricHistorySeries) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


