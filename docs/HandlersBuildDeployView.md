# HandlersBuildDeployView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int32** |  | [optional] 
**AttemptId** | Pointer to **int32** | AttemptID is the build-attempt ledger row id: the key the console and logs reads take (/build-attempts/:attempt_id/console, /logs). Zero, and omitted, when held (no row was written). | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**QueuePosition** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to **string** | held only — the ruled gate text | [optional] 
**ReasonCode** | Pointer to **string** | held only | [optional] 
**RunReference** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | dispatched | queued | held | [optional] 
**Trigger** | Pointer to **string** | always \&quot;cli\&quot; | [optional] 

## Methods

### NewHandlersBuildDeployView

`func NewHandlersBuildDeployView() *HandlersBuildDeployView`

NewHandlersBuildDeployView instantiates a new HandlersBuildDeployView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersBuildDeployViewWithDefaults

`func NewHandlersBuildDeployViewWithDefaults() *HandlersBuildDeployView`

NewHandlersBuildDeployViewWithDefaults instantiates a new HandlersBuildDeployView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *HandlersBuildDeployView) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *HandlersBuildDeployView) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *HandlersBuildDeployView) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *HandlersBuildDeployView) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetAttemptId

`func (o *HandlersBuildDeployView) GetAttemptId() int32`

GetAttemptId returns the AttemptId field if non-nil, zero value otherwise.

### GetAttemptIdOk

`func (o *HandlersBuildDeployView) GetAttemptIdOk() (*int32, bool)`

GetAttemptIdOk returns a tuple with the AttemptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptId

`func (o *HandlersBuildDeployView) SetAttemptId(v int32)`

SetAttemptId sets AttemptId field to given value.

### HasAttemptId

`func (o *HandlersBuildDeployView) HasAttemptId() bool`

HasAttemptId returns a boolean if a field has been set.

### GetJobId

`func (o *HandlersBuildDeployView) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *HandlersBuildDeployView) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *HandlersBuildDeployView) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *HandlersBuildDeployView) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetQueuePosition

`func (o *HandlersBuildDeployView) GetQueuePosition() int32`

GetQueuePosition returns the QueuePosition field if non-nil, zero value otherwise.

### GetQueuePositionOk

`func (o *HandlersBuildDeployView) GetQueuePositionOk() (*int32, bool)`

GetQueuePositionOk returns a tuple with the QueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePosition

`func (o *HandlersBuildDeployView) SetQueuePosition(v int32)`

SetQueuePosition sets QueuePosition field to given value.

### HasQueuePosition

`func (o *HandlersBuildDeployView) HasQueuePosition() bool`

HasQueuePosition returns a boolean if a field has been set.

### GetReason

`func (o *HandlersBuildDeployView) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HandlersBuildDeployView) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HandlersBuildDeployView) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HandlersBuildDeployView) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *HandlersBuildDeployView) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *HandlersBuildDeployView) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *HandlersBuildDeployView) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *HandlersBuildDeployView) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetRunReference

`func (o *HandlersBuildDeployView) GetRunReference() string`

GetRunReference returns the RunReference field if non-nil, zero value otherwise.

### GetRunReferenceOk

`func (o *HandlersBuildDeployView) GetRunReferenceOk() (*string, bool)`

GetRunReferenceOk returns a tuple with the RunReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunReference

`func (o *HandlersBuildDeployView) SetRunReference(v string)`

SetRunReference sets RunReference field to given value.

### HasRunReference

`func (o *HandlersBuildDeployView) HasRunReference() bool`

HasRunReference returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersBuildDeployView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersBuildDeployView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersBuildDeployView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersBuildDeployView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrigger

`func (o *HandlersBuildDeployView) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *HandlersBuildDeployView) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *HandlersBuildDeployView) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *HandlersBuildDeployView) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


