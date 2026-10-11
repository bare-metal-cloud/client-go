# AgentHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentsConnected** | Pointer to **int32** | Number of connected agents | [optional] 
**RedisConnected** | Pointer to **bool** | Redis connectivity status | [optional] 
**Status** | Pointer to **string** | \&quot;healthy\&quot; or \&quot;unhealthy\&quot; | [optional] 
**UptimeSeconds** | Pointer to **int32** | Service uptime | [optional] 
**Version** | Pointer to **string** | Service version | [optional] 

## Methods

### NewAgentHealthResponse

`func NewAgentHealthResponse() *AgentHealthResponse`

NewAgentHealthResponse instantiates a new AgentHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentHealthResponseWithDefaults

`func NewAgentHealthResponseWithDefaults() *AgentHealthResponse`

NewAgentHealthResponseWithDefaults instantiates a new AgentHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentsConnected

`func (o *AgentHealthResponse) GetAgentsConnected() int32`

GetAgentsConnected returns the AgentsConnected field if non-nil, zero value otherwise.

### GetAgentsConnectedOk

`func (o *AgentHealthResponse) GetAgentsConnectedOk() (*int32, bool)`

GetAgentsConnectedOk returns a tuple with the AgentsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentsConnected

`func (o *AgentHealthResponse) SetAgentsConnected(v int32)`

SetAgentsConnected sets AgentsConnected field to given value.

### HasAgentsConnected

`func (o *AgentHealthResponse) HasAgentsConnected() bool`

HasAgentsConnected returns a boolean if a field has been set.

### GetRedisConnected

`func (o *AgentHealthResponse) GetRedisConnected() bool`

GetRedisConnected returns the RedisConnected field if non-nil, zero value otherwise.

### GetRedisConnectedOk

`func (o *AgentHealthResponse) GetRedisConnectedOk() (*bool, bool)`

GetRedisConnectedOk returns a tuple with the RedisConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedisConnected

`func (o *AgentHealthResponse) SetRedisConnected(v bool)`

SetRedisConnected sets RedisConnected field to given value.

### HasRedisConnected

`func (o *AgentHealthResponse) HasRedisConnected() bool`

HasRedisConnected returns a boolean if a field has been set.

### GetStatus

`func (o *AgentHealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentHealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentHealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentHealthResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUptimeSeconds

`func (o *AgentHealthResponse) GetUptimeSeconds() int32`

GetUptimeSeconds returns the UptimeSeconds field if non-nil, zero value otherwise.

### GetUptimeSecondsOk

`func (o *AgentHealthResponse) GetUptimeSecondsOk() (*int32, bool)`

GetUptimeSecondsOk returns a tuple with the UptimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeSeconds

`func (o *AgentHealthResponse) SetUptimeSeconds(v int32)`

SetUptimeSeconds sets UptimeSeconds field to given value.

### HasUptimeSeconds

`func (o *AgentHealthResponse) HasUptimeSeconds() bool`

HasUptimeSeconds returns a boolean if a field has been set.

### GetVersion

`func (o *AgentHealthResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgentHealthResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgentHealthResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AgentHealthResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


