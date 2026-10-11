# ModelsClusterNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int32** |  | [optional] 
**ClusterId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**JoinedAt** | Pointer to **string** |  | [optional] 
**LastAttemptAt** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** | Phase is the install-lifecycle checkpoint (resume source of truth). Primary path (issue #1438): pending→setting_up→joining→ready, with \&quot;failed\&quot; as a resettable terminal (stuck nodes converge to it via the per-node timeout). | [optional] 
**PhaseError** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** | \&quot;master\&quot;, \&quot;worker\&quot; | [optional] 
**Status** | Pointer to **string** | legacy: \&quot;joining\&quot;, \&quot;ready\&quot;, \&quot;not-ready\&quot;, \&quot;removed\&quot; | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VirtualMachine** | Pointer to [**ModelsVirtualMachine**](ModelsVirtualMachine.md) |  | [optional] 
**VirtualMachineId** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsClusterNode

`func NewModelsClusterNode() *ModelsClusterNode`

NewModelsClusterNode instantiates a new ModelsClusterNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsClusterNodeWithDefaults

`func NewModelsClusterNodeWithDefaults() *ModelsClusterNode`

NewModelsClusterNodeWithDefaults instantiates a new ModelsClusterNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *ModelsClusterNode) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *ModelsClusterNode) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *ModelsClusterNode) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *ModelsClusterNode) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetClusterId

`func (o *ModelsClusterNode) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ModelsClusterNode) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ModelsClusterNode) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ModelsClusterNode) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ModelsClusterNode) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelsClusterNode) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelsClusterNode) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelsClusterNode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ModelsClusterNode) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsClusterNode) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsClusterNode) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsClusterNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJoinedAt

`func (o *ModelsClusterNode) GetJoinedAt() string`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *ModelsClusterNode) GetJoinedAtOk() (*string, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *ModelsClusterNode) SetJoinedAt(v string)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *ModelsClusterNode) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetLastAttemptAt

`func (o *ModelsClusterNode) GetLastAttemptAt() string`

GetLastAttemptAt returns the LastAttemptAt field if non-nil, zero value otherwise.

### GetLastAttemptAtOk

`func (o *ModelsClusterNode) GetLastAttemptAtOk() (*string, bool)`

GetLastAttemptAtOk returns a tuple with the LastAttemptAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttemptAt

`func (o *ModelsClusterNode) SetLastAttemptAt(v string)`

SetLastAttemptAt sets LastAttemptAt field to given value.

### HasLastAttemptAt

`func (o *ModelsClusterNode) HasLastAttemptAt() bool`

HasLastAttemptAt returns a boolean if a field has been set.

### GetPhase

`func (o *ModelsClusterNode) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ModelsClusterNode) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ModelsClusterNode) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ModelsClusterNode) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPhaseError

`func (o *ModelsClusterNode) GetPhaseError() string`

GetPhaseError returns the PhaseError field if non-nil, zero value otherwise.

### GetPhaseErrorOk

`func (o *ModelsClusterNode) GetPhaseErrorOk() (*string, bool)`

GetPhaseErrorOk returns a tuple with the PhaseError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseError

`func (o *ModelsClusterNode) SetPhaseError(v string)`

SetPhaseError sets PhaseError field to given value.

### HasPhaseError

`func (o *ModelsClusterNode) HasPhaseError() bool`

HasPhaseError returns a boolean if a field has been set.

### GetRole

`func (o *ModelsClusterNode) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ModelsClusterNode) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ModelsClusterNode) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ModelsClusterNode) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *ModelsClusterNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelsClusterNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelsClusterNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelsClusterNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelsClusterNode) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelsClusterNode) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelsClusterNode) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelsClusterNode) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *ModelsClusterNode) GetVirtualMachine() ModelsVirtualMachine`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *ModelsClusterNode) GetVirtualMachineOk() (*ModelsVirtualMachine, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *ModelsClusterNode) SetVirtualMachine(v ModelsVirtualMachine)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *ModelsClusterNode) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### GetVirtualMachineId

`func (o *ModelsClusterNode) GetVirtualMachineId() int32`

GetVirtualMachineId returns the VirtualMachineId field if non-nil, zero value otherwise.

### GetVirtualMachineIdOk

`func (o *ModelsClusterNode) GetVirtualMachineIdOk() (*int32, bool)`

GetVirtualMachineIdOk returns a tuple with the VirtualMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineId

`func (o *ModelsClusterNode) SetVirtualMachineId(v int32)`

SetVirtualMachineId sets VirtualMachineId field to given value.

### HasVirtualMachineId

`func (o *ModelsClusterNode) HasVirtualMachineId() bool`

HasVirtualMachineId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


