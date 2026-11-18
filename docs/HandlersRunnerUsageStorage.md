# HandlersRunnerUsageStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Series** | Pointer to [**[]HandlersRunnerUsageStoragePoint**](HandlersRunnerUsageStoragePoint.md) |  | [optional] 

## Methods

### NewHandlersRunnerUsageStorage

`func NewHandlersRunnerUsageStorage() *HandlersRunnerUsageStorage`

NewHandlersRunnerUsageStorage instantiates a new HandlersRunnerUsageStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerUsageStorageWithDefaults

`func NewHandlersRunnerUsageStorageWithDefaults() *HandlersRunnerUsageStorage`

NewHandlersRunnerUsageStorageWithDefaults instantiates a new HandlersRunnerUsageStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *HandlersRunnerUsageStorage) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HandlersRunnerUsageStorage) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HandlersRunnerUsageStorage) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HandlersRunnerUsageStorage) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSeries

`func (o *HandlersRunnerUsageStorage) GetSeries() []HandlersRunnerUsageStoragePoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *HandlersRunnerUsageStorage) GetSeriesOk() (*[]HandlersRunnerUsageStoragePoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *HandlersRunnerUsageStorage) SetSeries(v []HandlersRunnerUsageStoragePoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *HandlersRunnerUsageStorage) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


