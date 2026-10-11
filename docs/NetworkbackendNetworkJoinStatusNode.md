# NetworkbackendNetworkJoinStatusNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Job** | Pointer to [**NetworkbackendNetworkJoinJobState**](NetworkbackendNetworkJoinJobState.md) |  | [optional] 
**JoinedAt** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewNetworkbackendNetworkJoinStatusNode

`func NewNetworkbackendNetworkJoinStatusNode() *NetworkbackendNetworkJoinStatusNode`

NewNetworkbackendNetworkJoinStatusNode instantiates a new NetworkbackendNetworkJoinStatusNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkbackendNetworkJoinStatusNodeWithDefaults

`func NewNetworkbackendNetworkJoinStatusNodeWithDefaults() *NetworkbackendNetworkJoinStatusNode`

NewNetworkbackendNetworkJoinStatusNodeWithDefaults instantiates a new NetworkbackendNetworkJoinStatusNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *NetworkbackendNetworkJoinStatusNode) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *NetworkbackendNetworkJoinStatusNode) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *NetworkbackendNetworkJoinStatusNode) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *NetworkbackendNetworkJoinStatusNode) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetError

`func (o *NetworkbackendNetworkJoinStatusNode) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NetworkbackendNetworkJoinStatusNode) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NetworkbackendNetworkJoinStatusNode) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *NetworkbackendNetworkJoinStatusNode) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *NetworkbackendNetworkJoinStatusNode) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkbackendNetworkJoinStatusNode) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkbackendNetworkJoinStatusNode) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkbackendNetworkJoinStatusNode) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetIpAddress

`func (o *NetworkbackendNetworkJoinStatusNode) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkbackendNetworkJoinStatusNode) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkbackendNetworkJoinStatusNode) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkbackendNetworkJoinStatusNode) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetJob

`func (o *NetworkbackendNetworkJoinStatusNode) GetJob() NetworkbackendNetworkJoinJobState`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *NetworkbackendNetworkJoinStatusNode) GetJobOk() (*NetworkbackendNetworkJoinJobState, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *NetworkbackendNetworkJoinStatusNode) SetJob(v NetworkbackendNetworkJoinJobState)`

SetJob sets Job field to given value.

### HasJob

`func (o *NetworkbackendNetworkJoinStatusNode) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJoinedAt

`func (o *NetworkbackendNetworkJoinStatusNode) GetJoinedAt() string`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *NetworkbackendNetworkJoinStatusNode) GetJoinedAtOk() (*string, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *NetworkbackendNetworkJoinStatusNode) SetJoinedAt(v string)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *NetworkbackendNetworkJoinStatusNode) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetName

`func (o *NetworkbackendNetworkJoinStatusNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkbackendNetworkJoinStatusNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkbackendNetworkJoinStatusNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkbackendNetworkJoinStatusNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhase

`func (o *NetworkbackendNetworkJoinStatusNode) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkbackendNetworkJoinStatusNode) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkbackendNetworkJoinStatusNode) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *NetworkbackendNetworkJoinStatusNode) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NetworkbackendNetworkJoinStatusNode) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NetworkbackendNetworkJoinStatusNode) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NetworkbackendNetworkJoinStatusNode) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NetworkbackendNetworkJoinStatusNode) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVmId

`func (o *NetworkbackendNetworkJoinStatusNode) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *NetworkbackendNetworkJoinStatusNode) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *NetworkbackendNetworkJoinStatusNode) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *NetworkbackendNetworkJoinStatusNode) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


