# HandlersRunnerUsageNowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinutesToday** | Pointer to **float32** | today&#39;s UTC minutes, the daily rows&#39; own population | [optional] 
**RunningJobs** | Pointer to **int32** | runner_jobs in dispatched/running for the org | [optional] 
**StorageGB** | Pointer to **float32** | the recorded v1 boundary: 0 until the B5.2+ metadata | [optional] 

## Methods

### NewHandlersRunnerUsageNowResponse

`func NewHandlersRunnerUsageNowResponse() *HandlersRunnerUsageNowResponse`

NewHandlersRunnerUsageNowResponse instantiates a new HandlersRunnerUsageNowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerUsageNowResponseWithDefaults

`func NewHandlersRunnerUsageNowResponseWithDefaults() *HandlersRunnerUsageNowResponse`

NewHandlersRunnerUsageNowResponseWithDefaults instantiates a new HandlersRunnerUsageNowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinutesToday

`func (o *HandlersRunnerUsageNowResponse) GetMinutesToday() float32`

GetMinutesToday returns the MinutesToday field if non-nil, zero value otherwise.

### GetMinutesTodayOk

`func (o *HandlersRunnerUsageNowResponse) GetMinutesTodayOk() (*float32, bool)`

GetMinutesTodayOk returns a tuple with the MinutesToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesToday

`func (o *HandlersRunnerUsageNowResponse) SetMinutesToday(v float32)`

SetMinutesToday sets MinutesToday field to given value.

### HasMinutesToday

`func (o *HandlersRunnerUsageNowResponse) HasMinutesToday() bool`

HasMinutesToday returns a boolean if a field has been set.

### GetRunningJobs

`func (o *HandlersRunnerUsageNowResponse) GetRunningJobs() int32`

GetRunningJobs returns the RunningJobs field if non-nil, zero value otherwise.

### GetRunningJobsOk

`func (o *HandlersRunnerUsageNowResponse) GetRunningJobsOk() (*int32, bool)`

GetRunningJobsOk returns a tuple with the RunningJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningJobs

`func (o *HandlersRunnerUsageNowResponse) SetRunningJobs(v int32)`

SetRunningJobs sets RunningJobs field to given value.

### HasRunningJobs

`func (o *HandlersRunnerUsageNowResponse) HasRunningJobs() bool`

HasRunningJobs returns a boolean if a field has been set.

### GetStorageGB

`func (o *HandlersRunnerUsageNowResponse) GetStorageGB() float32`

GetStorageGB returns the StorageGB field if non-nil, zero value otherwise.

### GetStorageGBOk

`func (o *HandlersRunnerUsageNowResponse) GetStorageGBOk() (*float32, bool)`

GetStorageGBOk returns a tuple with the StorageGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGB

`func (o *HandlersRunnerUsageNowResponse) SetStorageGB(v float32)`

SetStorageGB sets StorageGB field to given value.

### HasStorageGB

`func (o *HandlersRunnerUsageNowResponse) HasStorageGB() bool`

HasStorageGB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


