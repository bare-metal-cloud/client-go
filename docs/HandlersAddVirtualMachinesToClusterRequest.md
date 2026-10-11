# HandlersAddVirtualMachinesToClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** |  | 
**VmIds** | **[]int32** |  | 

## Methods

### NewHandlersAddVirtualMachinesToClusterRequest

`func NewHandlersAddVirtualMachinesToClusterRequest(role string, vmIds []int32, ) *HandlersAddVirtualMachinesToClusterRequest`

NewHandlersAddVirtualMachinesToClusterRequest instantiates a new HandlersAddVirtualMachinesToClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersAddVirtualMachinesToClusterRequestWithDefaults

`func NewHandlersAddVirtualMachinesToClusterRequestWithDefaults() *HandlersAddVirtualMachinesToClusterRequest`

NewHandlersAddVirtualMachinesToClusterRequestWithDefaults instantiates a new HandlersAddVirtualMachinesToClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *HandlersAddVirtualMachinesToClusterRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *HandlersAddVirtualMachinesToClusterRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *HandlersAddVirtualMachinesToClusterRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetVmIds

`func (o *HandlersAddVirtualMachinesToClusterRequest) GetVmIds() []int32`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *HandlersAddVirtualMachinesToClusterRequest) GetVmIdsOk() (*[]int32, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *HandlersAddVirtualMachinesToClusterRequest) SetVmIds(v []int32)`

SetVmIds sets VmIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


