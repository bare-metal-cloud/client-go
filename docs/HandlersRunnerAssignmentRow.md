# HandlersRunnerAssignmentRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **int32** |  | [optional] 
**AgentStatus** | Pointer to **string** |  | [optional] 
**AssignedBy** | Pointer to **int32** |  | [optional] 
**Capabilities** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**MachinePollExpiresAt** | Pointer to **string** | MachinePollExpiresAt is the credential&#39;s expiry (the 24h machine-poll TTL), rendered when the token is. | [optional] 
**MachinePollToken** | Pointer to **string** | MachinePollToken is the machine chain&#39;s credential PLAINTEXT, returned exactly once on the assignment that minted it (epic 164 entry 12): the machine presents it as &#x60;Authorization: Bearer &lt;token&gt;&#x60; on /api/v1/runner/machine/_*. Only the hash rests server-side. Empty on every other read — the plaintext is never shown twice. | [optional] 
**Reachable** | Pointer to **bool** |  | [optional] 

## Methods

### NewHandlersRunnerAssignmentRow

`func NewHandlersRunnerAssignmentRow() *HandlersRunnerAssignmentRow`

NewHandlersRunnerAssignmentRow instantiates a new HandlersRunnerAssignmentRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerAssignmentRowWithDefaults

`func NewHandlersRunnerAssignmentRowWithDefaults() *HandlersRunnerAssignmentRow`

NewHandlersRunnerAssignmentRowWithDefaults instantiates a new HandlersRunnerAssignmentRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *HandlersRunnerAssignmentRow) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *HandlersRunnerAssignmentRow) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *HandlersRunnerAssignmentRow) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *HandlersRunnerAssignmentRow) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentStatus

`func (o *HandlersRunnerAssignmentRow) GetAgentStatus() string`

GetAgentStatus returns the AgentStatus field if non-nil, zero value otherwise.

### GetAgentStatusOk

`func (o *HandlersRunnerAssignmentRow) GetAgentStatusOk() (*string, bool)`

GetAgentStatusOk returns a tuple with the AgentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStatus

`func (o *HandlersRunnerAssignmentRow) SetAgentStatus(v string)`

SetAgentStatus sets AgentStatus field to given value.

### HasAgentStatus

`func (o *HandlersRunnerAssignmentRow) HasAgentStatus() bool`

HasAgentStatus returns a boolean if a field has been set.

### GetAssignedBy

`func (o *HandlersRunnerAssignmentRow) GetAssignedBy() int32`

GetAssignedBy returns the AssignedBy field if non-nil, zero value otherwise.

### GetAssignedByOk

`func (o *HandlersRunnerAssignmentRow) GetAssignedByOk() (*int32, bool)`

GetAssignedByOk returns a tuple with the AssignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBy

`func (o *HandlersRunnerAssignmentRow) SetAssignedBy(v int32)`

SetAssignedBy sets AssignedBy field to given value.

### HasAssignedBy

`func (o *HandlersRunnerAssignmentRow) HasAssignedBy() bool`

HasAssignedBy returns a boolean if a field has been set.

### GetCapabilities

`func (o *HandlersRunnerAssignmentRow) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *HandlersRunnerAssignmentRow) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *HandlersRunnerAssignmentRow) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *HandlersRunnerAssignmentRow) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersRunnerAssignmentRow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersRunnerAssignmentRow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersRunnerAssignmentRow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersRunnerAssignmentRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMachinePollExpiresAt

`func (o *HandlersRunnerAssignmentRow) GetMachinePollExpiresAt() string`

GetMachinePollExpiresAt returns the MachinePollExpiresAt field if non-nil, zero value otherwise.

### GetMachinePollExpiresAtOk

`func (o *HandlersRunnerAssignmentRow) GetMachinePollExpiresAtOk() (*string, bool)`

GetMachinePollExpiresAtOk returns a tuple with the MachinePollExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinePollExpiresAt

`func (o *HandlersRunnerAssignmentRow) SetMachinePollExpiresAt(v string)`

SetMachinePollExpiresAt sets MachinePollExpiresAt field to given value.

### HasMachinePollExpiresAt

`func (o *HandlersRunnerAssignmentRow) HasMachinePollExpiresAt() bool`

HasMachinePollExpiresAt returns a boolean if a field has been set.

### GetMachinePollToken

`func (o *HandlersRunnerAssignmentRow) GetMachinePollToken() string`

GetMachinePollToken returns the MachinePollToken field if non-nil, zero value otherwise.

### GetMachinePollTokenOk

`func (o *HandlersRunnerAssignmentRow) GetMachinePollTokenOk() (*string, bool)`

GetMachinePollTokenOk returns a tuple with the MachinePollToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachinePollToken

`func (o *HandlersRunnerAssignmentRow) SetMachinePollToken(v string)`

SetMachinePollToken sets MachinePollToken field to given value.

### HasMachinePollToken

`func (o *HandlersRunnerAssignmentRow) HasMachinePollToken() bool`

HasMachinePollToken returns a boolean if a field has been set.

### GetReachable

`func (o *HandlersRunnerAssignmentRow) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *HandlersRunnerAssignmentRow) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *HandlersRunnerAssignmentRow) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *HandlersRunnerAssignmentRow) HasReachable() bool`

HasReachable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


