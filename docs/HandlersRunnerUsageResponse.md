# HandlersRunnerUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis states the counting rules (D1 — the rules ride the answer). | [optional] 
**CurrentPeriod** | Pointer to [**HandlersRunnerUsageCurrentPeriod**](HandlersRunnerUsageCurrentPeriod.md) | CurrentPeriod is the open month&#39;s projection block; nil when the requested period is a closed month (a projection of the past would invent a future). | [optional] 
**Daily** | Pointer to [**[]HandlersRunnerUsageDailyRow**](HandlersRunnerUsageDailyRow.md) | Daily is the requested period&#39;s UTC day × platform rows, day then platform ascending. The CSV export serves these rows verbatim. | [optional] 
**Period** | Pointer to **string** | Period echoes the answer&#39;s YYYY-MM period (the default filled in). | [optional] 
**Platform** | Pointer to [**[]HandlersRunnerUsagePlatformMinutes**](HandlersRunnerUsagePlatformMinutes.md) | Platform is the period&#39;s per-platform totals (the daily rows summed — exactly, no second query). | [optional] 
**PreviousPeriod** | Pointer to [**HandlersRunnerUsagePreviousPeriod**](HandlersRunnerUsagePreviousPeriod.md) | PreviousPeriod is the calendar month before the requested period. | [optional] 
**Storage** | Pointer to [**HandlersRunnerUsageStorage**](HandlersRunnerUsageStorage.md) | Storage is the recorded v1 boundary: an empty series with the reason (artifact metadata arrives with the artifact store stories). | [optional] 
**Total** | Pointer to **float32** | Total is the period&#39;s total minutes — the daily rows&#39; exact sum of raw floats (never rounded; the screen may format). | [optional] 

## Methods

### NewHandlersRunnerUsageResponse

`func NewHandlersRunnerUsageResponse() *HandlersRunnerUsageResponse`

NewHandlersRunnerUsageResponse instantiates a new HandlersRunnerUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerUsageResponseWithDefaults

`func NewHandlersRunnerUsageResponseWithDefaults() *HandlersRunnerUsageResponse`

NewHandlersRunnerUsageResponseWithDefaults instantiates a new HandlersRunnerUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *HandlersRunnerUsageResponse) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *HandlersRunnerUsageResponse) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *HandlersRunnerUsageResponse) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *HandlersRunnerUsageResponse) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetCurrentPeriod

`func (o *HandlersRunnerUsageResponse) GetCurrentPeriod() HandlersRunnerUsageCurrentPeriod`

GetCurrentPeriod returns the CurrentPeriod field if non-nil, zero value otherwise.

### GetCurrentPeriodOk

`func (o *HandlersRunnerUsageResponse) GetCurrentPeriodOk() (*HandlersRunnerUsageCurrentPeriod, bool)`

GetCurrentPeriodOk returns a tuple with the CurrentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriod

`func (o *HandlersRunnerUsageResponse) SetCurrentPeriod(v HandlersRunnerUsageCurrentPeriod)`

SetCurrentPeriod sets CurrentPeriod field to given value.

### HasCurrentPeriod

`func (o *HandlersRunnerUsageResponse) HasCurrentPeriod() bool`

HasCurrentPeriod returns a boolean if a field has been set.

### GetDaily

`func (o *HandlersRunnerUsageResponse) GetDaily() []HandlersRunnerUsageDailyRow`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *HandlersRunnerUsageResponse) GetDailyOk() (*[]HandlersRunnerUsageDailyRow, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *HandlersRunnerUsageResponse) SetDaily(v []HandlersRunnerUsageDailyRow)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *HandlersRunnerUsageResponse) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetPeriod

`func (o *HandlersRunnerUsageResponse) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *HandlersRunnerUsageResponse) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *HandlersRunnerUsageResponse) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *HandlersRunnerUsageResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPlatform

`func (o *HandlersRunnerUsageResponse) GetPlatform() []HandlersRunnerUsagePlatformMinutes`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *HandlersRunnerUsageResponse) GetPlatformOk() (*[]HandlersRunnerUsagePlatformMinutes, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *HandlersRunnerUsageResponse) SetPlatform(v []HandlersRunnerUsagePlatformMinutes)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *HandlersRunnerUsageResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPreviousPeriod

`func (o *HandlersRunnerUsageResponse) GetPreviousPeriod() HandlersRunnerUsagePreviousPeriod`

GetPreviousPeriod returns the PreviousPeriod field if non-nil, zero value otherwise.

### GetPreviousPeriodOk

`func (o *HandlersRunnerUsageResponse) GetPreviousPeriodOk() (*HandlersRunnerUsagePreviousPeriod, bool)`

GetPreviousPeriodOk returns a tuple with the PreviousPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPeriod

`func (o *HandlersRunnerUsageResponse) SetPreviousPeriod(v HandlersRunnerUsagePreviousPeriod)`

SetPreviousPeriod sets PreviousPeriod field to given value.

### HasPreviousPeriod

`func (o *HandlersRunnerUsageResponse) HasPreviousPeriod() bool`

HasPreviousPeriod returns a boolean if a field has been set.

### GetStorage

`func (o *HandlersRunnerUsageResponse) GetStorage() HandlersRunnerUsageStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *HandlersRunnerUsageResponse) GetStorageOk() (*HandlersRunnerUsageStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *HandlersRunnerUsageResponse) SetStorage(v HandlersRunnerUsageStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *HandlersRunnerUsageResponse) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetTotal

`func (o *HandlersRunnerUsageResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HandlersRunnerUsageResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HandlersRunnerUsageResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HandlersRunnerUsageResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


