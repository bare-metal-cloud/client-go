# ServicesJobListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CurrentStep** | Pointer to **string** | ADR-065 roll-up | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**ResourceId** | Pointer to **int32** | ADR-065 | [optional] 
**ResourceType** | Pointer to **string** | ADR-065 | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ModelsJobMetadataStatus**](ModelsJobMetadataStatus.md) |  | [optional] 
**StepCount** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewServicesJobListItem

`func NewServicesJobListItem() *ServicesJobListItem`

NewServicesJobListItem instantiates a new ServicesJobListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesJobListItemWithDefaults

`func NewServicesJobListItemWithDefaults() *ServicesJobListItem`

NewServicesJobListItemWithDefaults instantiates a new ServicesJobListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *ServicesJobListItem) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ServicesJobListItem) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ServicesJobListItem) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ServicesJobListItem) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServicesJobListItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServicesJobListItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServicesJobListItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServicesJobListItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentStep

`func (o *ServicesJobListItem) GetCurrentStep() string`

GetCurrentStep returns the CurrentStep field if non-nil, zero value otherwise.

### GetCurrentStepOk

`func (o *ServicesJobListItem) GetCurrentStepOk() (*string, bool)`

GetCurrentStepOk returns a tuple with the CurrentStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStep

`func (o *ServicesJobListItem) SetCurrentStep(v string)`

SetCurrentStep sets CurrentStep field to given value.

### HasCurrentStep

`func (o *ServicesJobListItem) HasCurrentStep() bool`

HasCurrentStep returns a boolean if a field has been set.

### GetJobId

`func (o *ServicesJobListItem) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ServicesJobListItem) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ServicesJobListItem) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ServicesJobListItem) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobType

`func (o *ServicesJobListItem) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ServicesJobListItem) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ServicesJobListItem) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ServicesJobListItem) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetProgress

`func (o *ServicesJobListItem) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ServicesJobListItem) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ServicesJobListItem) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ServicesJobListItem) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetResourceId

`func (o *ServicesJobListItem) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ServicesJobListItem) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ServicesJobListItem) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ServicesJobListItem) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *ServicesJobListItem) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ServicesJobListItem) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ServicesJobListItem) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ServicesJobListItem) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStartedAt

`func (o *ServicesJobListItem) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ServicesJobListItem) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ServicesJobListItem) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ServicesJobListItem) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ServicesJobListItem) GetStatus() ModelsJobMetadataStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServicesJobListItem) GetStatusOk() (*ModelsJobMetadataStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServicesJobListItem) SetStatus(v ModelsJobMetadataStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServicesJobListItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStepCount

`func (o *ServicesJobListItem) GetStepCount() int32`

GetStepCount returns the StepCount field if non-nil, zero value otherwise.

### GetStepCountOk

`func (o *ServicesJobListItem) GetStepCountOk() (*int32, bool)`

GetStepCountOk returns a tuple with the StepCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepCount

`func (o *ServicesJobListItem) SetStepCount(v int32)`

SetStepCount sets StepCount field to given value.

### HasStepCount

`func (o *ServicesJobListItem) HasStepCount() bool`

HasStepCount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServicesJobListItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServicesJobListItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServicesJobListItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServicesJobListItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *ServicesJobListItem) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ServicesJobListItem) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ServicesJobListItem) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ServicesJobListItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


