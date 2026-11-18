# HandlersRunnerUsageCurrentPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysElapsed** | Pointer to **int32** | UTC calendar days including today (the DECIDED calendar-so-far law) | [optional] 
**DaysInPeriod** | Pointer to **int32** |  | [optional] 
**IncludedMinutes** | Pointer to **float32** |  | [optional] 
**IncludedSource** | Pointer to **string** | the named constant&#39;s provenance | [optional] 
**MinutesSoFar** | Pointer to **float32** |  | [optional] 
**Projection** | Pointer to **float32** | whole minutes: minutesSoFar / daysElapsed × daysInPeriod, rounded | [optional] 

## Methods

### NewHandlersRunnerUsageCurrentPeriod

`func NewHandlersRunnerUsageCurrentPeriod() *HandlersRunnerUsageCurrentPeriod`

NewHandlersRunnerUsageCurrentPeriod instantiates a new HandlersRunnerUsageCurrentPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerUsageCurrentPeriodWithDefaults

`func NewHandlersRunnerUsageCurrentPeriodWithDefaults() *HandlersRunnerUsageCurrentPeriod`

NewHandlersRunnerUsageCurrentPeriodWithDefaults instantiates a new HandlersRunnerUsageCurrentPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysElapsed

`func (o *HandlersRunnerUsageCurrentPeriod) GetDaysElapsed() int32`

GetDaysElapsed returns the DaysElapsed field if non-nil, zero value otherwise.

### GetDaysElapsedOk

`func (o *HandlersRunnerUsageCurrentPeriod) GetDaysElapsedOk() (*int32, bool)`

GetDaysElapsedOk returns a tuple with the DaysElapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysElapsed

`func (o *HandlersRunnerUsageCurrentPeriod) SetDaysElapsed(v int32)`

SetDaysElapsed sets DaysElapsed field to given value.

### HasDaysElapsed

`func (o *HandlersRunnerUsageCurrentPeriod) HasDaysElapsed() bool`

HasDaysElapsed returns a boolean if a field has been set.

### GetDaysInPeriod

`func (o *HandlersRunnerUsageCurrentPeriod) GetDaysInPeriod() int32`

GetDaysInPeriod returns the DaysInPeriod field if non-nil, zero value otherwise.

### GetDaysInPeriodOk

`func (o *HandlersRunnerUsageCurrentPeriod) GetDaysInPeriodOk() (*int32, bool)`

GetDaysInPeriodOk returns a tuple with the DaysInPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInPeriod

`func (o *HandlersRunnerUsageCurrentPeriod) SetDaysInPeriod(v int32)`

SetDaysInPeriod sets DaysInPeriod field to given value.

### HasDaysInPeriod

`func (o *HandlersRunnerUsageCurrentPeriod) HasDaysInPeriod() bool`

HasDaysInPeriod returns a boolean if a field has been set.

### GetIncludedMinutes

`func (o *HandlersRunnerUsageCurrentPeriod) GetIncludedMinutes() float32`

GetIncludedMinutes returns the IncludedMinutes field if non-nil, zero value otherwise.

### GetIncludedMinutesOk

`func (o *HandlersRunnerUsageCurrentPeriod) GetIncludedMinutesOk() (*float32, bool)`

GetIncludedMinutesOk returns a tuple with the IncludedMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedMinutes

`func (o *HandlersRunnerUsageCurrentPeriod) SetIncludedMinutes(v float32)`

SetIncludedMinutes sets IncludedMinutes field to given value.

### HasIncludedMinutes

`func (o *HandlersRunnerUsageCurrentPeriod) HasIncludedMinutes() bool`

HasIncludedMinutes returns a boolean if a field has been set.

### GetIncludedSource

`func (o *HandlersRunnerUsageCurrentPeriod) GetIncludedSource() string`

GetIncludedSource returns the IncludedSource field if non-nil, zero value otherwise.

### GetIncludedSourceOk

`func (o *HandlersRunnerUsageCurrentPeriod) GetIncludedSourceOk() (*string, bool)`

GetIncludedSourceOk returns a tuple with the IncludedSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSource

`func (o *HandlersRunnerUsageCurrentPeriod) SetIncludedSource(v string)`

SetIncludedSource sets IncludedSource field to given value.

### HasIncludedSource

`func (o *HandlersRunnerUsageCurrentPeriod) HasIncludedSource() bool`

HasIncludedSource returns a boolean if a field has been set.

### GetMinutesSoFar

`func (o *HandlersRunnerUsageCurrentPeriod) GetMinutesSoFar() float32`

GetMinutesSoFar returns the MinutesSoFar field if non-nil, zero value otherwise.

### GetMinutesSoFarOk

`func (o *HandlersRunnerUsageCurrentPeriod) GetMinutesSoFarOk() (*float32, bool)`

GetMinutesSoFarOk returns a tuple with the MinutesSoFar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesSoFar

`func (o *HandlersRunnerUsageCurrentPeriod) SetMinutesSoFar(v float32)`

SetMinutesSoFar sets MinutesSoFar field to given value.

### HasMinutesSoFar

`func (o *HandlersRunnerUsageCurrentPeriod) HasMinutesSoFar() bool`

HasMinutesSoFar returns a boolean if a field has been set.

### GetProjection

`func (o *HandlersRunnerUsageCurrentPeriod) GetProjection() float32`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *HandlersRunnerUsageCurrentPeriod) GetProjectionOk() (*float32, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *HandlersRunnerUsageCurrentPeriod) SetProjection(v float32)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *HandlersRunnerUsageCurrentPeriod) HasProjection() bool`

HasProjection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


