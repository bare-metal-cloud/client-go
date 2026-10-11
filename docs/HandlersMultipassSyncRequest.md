# HandlersMultipassSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** |  | [optional] 
**HostVmId** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 
**Vms** | Pointer to [**[]HandlersMultipassVMEntryRequest**](HandlersMultipassVMEntryRequest.md) |  | [optional] 

## Methods

### NewHandlersMultipassSyncRequest

`func NewHandlersMultipassSyncRequest() *HandlersMultipassSyncRequest`

NewHandlersMultipassSyncRequest instantiates a new HandlersMultipassSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMultipassSyncRequestWithDefaults

`func NewHandlersMultipassSyncRequestWithDefaults() *HandlersMultipassSyncRequest`

NewHandlersMultipassSyncRequestWithDefaults instantiates a new HandlersMultipassSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *HandlersMultipassSyncRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *HandlersMultipassSyncRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *HandlersMultipassSyncRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *HandlersMultipassSyncRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetHostVmId

`func (o *HandlersMultipassSyncRequest) GetHostVmId() int32`

GetHostVmId returns the HostVmId field if non-nil, zero value otherwise.

### GetHostVmIdOk

`func (o *HandlersMultipassSyncRequest) GetHostVmIdOk() (*int32, bool)`

GetHostVmIdOk returns a tuple with the HostVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVmId

`func (o *HandlersMultipassSyncRequest) SetHostVmId(v int32)`

SetHostVmId sets HostVmId field to given value.

### HasHostVmId

`func (o *HandlersMultipassSyncRequest) HasHostVmId() bool`

HasHostVmId returns a boolean if a field has been set.

### GetTimestamp

`func (o *HandlersMultipassSyncRequest) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HandlersMultipassSyncRequest) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HandlersMultipassSyncRequest) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HandlersMultipassSyncRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVms

`func (o *HandlersMultipassSyncRequest) GetVms() []HandlersMultipassVMEntryRequest`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *HandlersMultipassSyncRequest) GetVmsOk() (*[]HandlersMultipassVMEntryRequest, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *HandlersMultipassSyncRequest) SetVms(v []HandlersMultipassVMEntryRequest)`

SetVms sets Vms field to given value.

### HasVms

`func (o *HandlersMultipassSyncRequest) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


