# HandlersMetricHistoryPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **float32** | unix timestamp (ms) for charting | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewHandlersMetricHistoryPoint

`func NewHandlersMetricHistoryPoint() *HandlersMetricHistoryPoint`

NewHandlersMetricHistoryPoint instantiates a new HandlersMetricHistoryPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMetricHistoryPointWithDefaults

`func NewHandlersMetricHistoryPointWithDefaults() *HandlersMetricHistoryPoint`

NewHandlersMetricHistoryPointWithDefaults instantiates a new HandlersMetricHistoryPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *HandlersMetricHistoryPoint) GetTime() float32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *HandlersMetricHistoryPoint) GetTimeOk() (*float32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *HandlersMetricHistoryPoint) SetTime(v float32)`

SetTime sets Time field to given value.

### HasTime

`func (o *HandlersMetricHistoryPoint) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetValue

`func (o *HandlersMetricHistoryPoint) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HandlersMetricHistoryPoint) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HandlersMetricHistoryPoint) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *HandlersMetricHistoryPoint) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


