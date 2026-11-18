# ModelsJobStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **int32** | deprecated: use ResourceType+ResourceID (ADR-060 V6); 0 if not cluster-scoped | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**LastAttemptAt** | Pointer to **string** | Story 18.4.1: When we last tried this step | [optional] 
**Metadata** | Pointer to **[]int32** | Error messages, progress, etc. | [optional] 
**ResourceId** | Pointer to **int32** | ADR-060 V6: the row id of ResourceType | [optional] 
**ResourceType** | Pointer to **string** | ADR-060 V6: &#39;cluster&#39; | &#39;network_node&#39; | &#39;vm&#39; | ... | [optional] 
**RetryCount** | Pointer to **int32** | Story 18.4.1: Total retry attempts (0-indexed) | [optional] 
**Status** | Pointer to [**ModelsJobStepStatus**](ModelsJobStepStatus.md) |  | [optional] 
**StepName** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**WorkerId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelsJobStep

`func NewModelsJobStep() *ModelsJobStep`

NewModelsJobStep instantiates a new ModelsJobStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsJobStepWithDefaults

`func NewModelsJobStepWithDefaults() *ModelsJobStep`

NewModelsJobStepWithDefaults instantiates a new ModelsJobStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *ModelsJobStep) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ModelsJobStep) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ModelsJobStep) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ModelsJobStep) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsJobStep) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsJobStep) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsJobStep) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsJobStep) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetJobId

`func (o *ModelsJobStep) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ModelsJobStep) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ModelsJobStep) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ModelsJobStep) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetLastAttemptAt

`func (o *ModelsJobStep) GetLastAttemptAt() string`

GetLastAttemptAt returns the LastAttemptAt field if non-nil, zero value otherwise.

### GetLastAttemptAtOk

`func (o *ModelsJobStep) GetLastAttemptAtOk() (*string, bool)`

GetLastAttemptAtOk returns a tuple with the LastAttemptAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttemptAt

`func (o *ModelsJobStep) SetLastAttemptAt(v string)`

SetLastAttemptAt sets LastAttemptAt field to given value.

### HasLastAttemptAt

`func (o *ModelsJobStep) HasLastAttemptAt() bool`

HasLastAttemptAt returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsJobStep) GetMetadata() []int32`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsJobStep) GetMetadataOk() (*[]int32, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsJobStep) SetMetadata(v []int32)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsJobStep) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResourceId

`func (o *ModelsJobStep) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ModelsJobStep) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ModelsJobStep) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ModelsJobStep) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *ModelsJobStep) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ModelsJobStep) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ModelsJobStep) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ModelsJobStep) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetRetryCount

`func (o *ModelsJobStep) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ModelsJobStep) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ModelsJobStep) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ModelsJobStep) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsJobStep) GetStatus() ModelsJobStepStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsJobStep) GetStatusOk() (*ModelsJobStepStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsJobStep) SetStatus(v ModelsJobStepStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsJobStep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStepName

`func (o *ModelsJobStep) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *ModelsJobStep) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *ModelsJobStep) SetStepName(v string)`

SetStepName sets StepName field to given value.

### HasStepName

`func (o *ModelsJobStep) HasStepName() bool`

HasStepName returns a boolean if a field has been set.

### GetTime

`func (o *ModelsJobStep) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ModelsJobStep) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ModelsJobStep) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *ModelsJobStep) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetWorkerId

`func (o *ModelsJobStep) GetWorkerId() string`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *ModelsJobStep) GetWorkerIdOk() (*string, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *ModelsJobStep) SetWorkerId(v string)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *ModelsJobStep) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


