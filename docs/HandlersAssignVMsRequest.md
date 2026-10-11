# HandlersAssignVMsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to **map[string]string** | vm_id -&gt; role mapping (optional, auto-assign if omitted) | [optional] 
**VmIds** | **[]int32** |  | 

## Methods

### NewHandlersAssignVMsRequest

`func NewHandlersAssignVMsRequest(vmIds []int32, ) *HandlersAssignVMsRequest`

NewHandlersAssignVMsRequest instantiates a new HandlersAssignVMsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersAssignVMsRequestWithDefaults

`func NewHandlersAssignVMsRequestWithDefaults() *HandlersAssignVMsRequest`

NewHandlersAssignVMsRequestWithDefaults instantiates a new HandlersAssignVMsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *HandlersAssignVMsRequest) GetRoles() map[string]string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *HandlersAssignVMsRequest) GetRolesOk() (*map[string]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *HandlersAssignVMsRequest) SetRoles(v map[string]string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *HandlersAssignVMsRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetVmIds

`func (o *HandlersAssignVMsRequest) GetVmIds() []int32`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *HandlersAssignVMsRequest) GetVmIdsOk() (*[]int32, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *HandlersAssignVMsRequest) SetVmIds(v []int32)`

SetVmIds sets VmIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


