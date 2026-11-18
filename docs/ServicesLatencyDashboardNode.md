# ServicesLatencyDashboardNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NodeStatus** | Pointer to **string** | NodeStatus is the honesty state: measured | off_not_measured | preparing_not_measured | address_diverged_not_measured | not_measured. | [optional] 
**PowerState** | Pointer to **string** | PowerState is the virtual_machines status (the power-state truth). | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**VmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewServicesLatencyDashboardNode

`func NewServicesLatencyDashboardNode() *ServicesLatencyDashboardNode`

NewServicesLatencyDashboardNode instantiates a new ServicesLatencyDashboardNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesLatencyDashboardNodeWithDefaults

`func NewServicesLatencyDashboardNodeWithDefaults() *ServicesLatencyDashboardNode`

NewServicesLatencyDashboardNodeWithDefaults instantiates a new ServicesLatencyDashboardNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServicesLatencyDashboardNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicesLatencyDashboardNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicesLatencyDashboardNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicesLatencyDashboardNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeStatus

`func (o *ServicesLatencyDashboardNode) GetNodeStatus() string`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *ServicesLatencyDashboardNode) GetNodeStatusOk() (*string, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *ServicesLatencyDashboardNode) SetNodeStatus(v string)`

SetNodeStatus sets NodeStatus field to given value.

### HasNodeStatus

`func (o *ServicesLatencyDashboardNode) HasNodeStatus() bool`

HasNodeStatus returns a boolean if a field has been set.

### GetPowerState

`func (o *ServicesLatencyDashboardNode) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *ServicesLatencyDashboardNode) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *ServicesLatencyDashboardNode) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *ServicesLatencyDashboardNode) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetRole

`func (o *ServicesLatencyDashboardNode) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServicesLatencyDashboardNode) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServicesLatencyDashboardNode) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ServicesLatencyDashboardNode) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetVmId

`func (o *ServicesLatencyDashboardNode) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *ServicesLatencyDashboardNode) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *ServicesLatencyDashboardNode) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *ServicesLatencyDashboardNode) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


