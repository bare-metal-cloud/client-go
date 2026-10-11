# AgentShutdownNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** | Agent unique identifier | [optional] 
**Reason** | Pointer to **string** | Shutdown reason (e.g., \&quot;graceful_shutdown\&quot;) | [optional] 
**Signal** | Pointer to **string** | Signal received (e.g., \&quot;terminated\&quot;, \&quot;interrupt\&quot;) | [optional] 
**Timestamp** | Pointer to **string** | Shutdown timestamp | [optional] 

## Methods

### NewAgentShutdownNotificationRequest

`func NewAgentShutdownNotificationRequest() *AgentShutdownNotificationRequest`

NewAgentShutdownNotificationRequest instantiates a new AgentShutdownNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentShutdownNotificationRequestWithDefaults

`func NewAgentShutdownNotificationRequestWithDefaults() *AgentShutdownNotificationRequest`

NewAgentShutdownNotificationRequestWithDefaults instantiates a new AgentShutdownNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *AgentShutdownNotificationRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentShutdownNotificationRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentShutdownNotificationRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AgentShutdownNotificationRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetReason

`func (o *AgentShutdownNotificationRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AgentShutdownNotificationRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AgentShutdownNotificationRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AgentShutdownNotificationRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSignal

`func (o *AgentShutdownNotificationRequest) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *AgentShutdownNotificationRequest) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *AgentShutdownNotificationRequest) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *AgentShutdownNotificationRequest) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetTimestamp

`func (o *AgentShutdownNotificationRequest) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AgentShutdownNotificationRequest) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AgentShutdownNotificationRequest) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AgentShutdownNotificationRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


