# HandlersRunnerProjectSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosted** | Pointer to **bool** |  | [optional] 
**Jobs** | Pointer to [**[]RunnerpoolJobStatus**](RunnerpoolJobStatus.md) |  | [optional] 
**MinutesCeiling** | Pointer to **float32** |  | [optional] 
**MinutesMetered** | Pointer to **float32** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**Quota** | Pointer to [**RunnerpoolQuotaStatus**](RunnerpoolQuotaStatus.md) |  | [optional] 
**Reservation** | Pointer to [**RunnerpoolReservationState**](RunnerpoolReservationState.md) |  | [optional] 

## Methods

### NewHandlersRunnerProjectSection

`func NewHandlersRunnerProjectSection() *HandlersRunnerProjectSection`

NewHandlersRunnerProjectSection instantiates a new HandlersRunnerProjectSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerProjectSectionWithDefaults

`func NewHandlersRunnerProjectSectionWithDefaults() *HandlersRunnerProjectSection`

NewHandlersRunnerProjectSectionWithDefaults instantiates a new HandlersRunnerProjectSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosted

`func (o *HandlersRunnerProjectSection) GetHosted() bool`

GetHosted returns the Hosted field if non-nil, zero value otherwise.

### GetHostedOk

`func (o *HandlersRunnerProjectSection) GetHostedOk() (*bool, bool)`

GetHostedOk returns a tuple with the Hosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosted

`func (o *HandlersRunnerProjectSection) SetHosted(v bool)`

SetHosted sets Hosted field to given value.

### HasHosted

`func (o *HandlersRunnerProjectSection) HasHosted() bool`

HasHosted returns a boolean if a field has been set.

### GetJobs

`func (o *HandlersRunnerProjectSection) GetJobs() []RunnerpoolJobStatus`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *HandlersRunnerProjectSection) GetJobsOk() (*[]RunnerpoolJobStatus, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *HandlersRunnerProjectSection) SetJobs(v []RunnerpoolJobStatus)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *HandlersRunnerProjectSection) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetMinutesCeiling

`func (o *HandlersRunnerProjectSection) GetMinutesCeiling() float32`

GetMinutesCeiling returns the MinutesCeiling field if non-nil, zero value otherwise.

### GetMinutesCeilingOk

`func (o *HandlersRunnerProjectSection) GetMinutesCeilingOk() (*float32, bool)`

GetMinutesCeilingOk returns a tuple with the MinutesCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesCeiling

`func (o *HandlersRunnerProjectSection) SetMinutesCeiling(v float32)`

SetMinutesCeiling sets MinutesCeiling field to given value.

### HasMinutesCeiling

`func (o *HandlersRunnerProjectSection) HasMinutesCeiling() bool`

HasMinutesCeiling returns a boolean if a field has been set.

### GetMinutesMetered

`func (o *HandlersRunnerProjectSection) GetMinutesMetered() float32`

GetMinutesMetered returns the MinutesMetered field if non-nil, zero value otherwise.

### GetMinutesMeteredOk

`func (o *HandlersRunnerProjectSection) GetMinutesMeteredOk() (*float32, bool)`

GetMinutesMeteredOk returns a tuple with the MinutesMetered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesMetered

`func (o *HandlersRunnerProjectSection) SetMinutesMetered(v float32)`

SetMinutesMetered sets MinutesMetered field to given value.

### HasMinutesMetered

`func (o *HandlersRunnerProjectSection) HasMinutesMetered() bool`

HasMinutesMetered returns a boolean if a field has been set.

### GetOrgId

`func (o *HandlersRunnerProjectSection) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *HandlersRunnerProjectSection) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *HandlersRunnerProjectSection) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *HandlersRunnerProjectSection) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetProjectId

`func (o *HandlersRunnerProjectSection) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HandlersRunnerProjectSection) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HandlersRunnerProjectSection) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *HandlersRunnerProjectSection) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetQuota

`func (o *HandlersRunnerProjectSection) GetQuota() RunnerpoolQuotaStatus`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *HandlersRunnerProjectSection) GetQuotaOk() (*RunnerpoolQuotaStatus, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *HandlersRunnerProjectSection) SetQuota(v RunnerpoolQuotaStatus)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *HandlersRunnerProjectSection) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetReservation

`func (o *HandlersRunnerProjectSection) GetReservation() RunnerpoolReservationState`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *HandlersRunnerProjectSection) GetReservationOk() (*RunnerpoolReservationState, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *HandlersRunnerProjectSection) SetReservation(v RunnerpoolReservationState)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *HandlersRunnerProjectSection) HasReservation() bool`

HasReservation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


