# HandlersVMCheckK3sResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckedAt** | Pointer to **string** |  | [optional] 
**Decision** | Pointer to [**VmpreflightRejection**](VmpreflightRejection.md) |  | [optional] 
**K3sState** | Pointer to **string** |  | [optional] 
**Probe** | Pointer to [**AgentopsCheckK3sResult**](AgentopsCheckK3sResult.md) |  | [optional] 
**VmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersVMCheckK3sResponse

`func NewHandlersVMCheckK3sResponse() *HandlersVMCheckK3sResponse`

NewHandlersVMCheckK3sResponse instantiates a new HandlersVMCheckK3sResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersVMCheckK3sResponseWithDefaults

`func NewHandlersVMCheckK3sResponseWithDefaults() *HandlersVMCheckK3sResponse`

NewHandlersVMCheckK3sResponseWithDefaults instantiates a new HandlersVMCheckK3sResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckedAt

`func (o *HandlersVMCheckK3sResponse) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *HandlersVMCheckK3sResponse) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *HandlersVMCheckK3sResponse) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *HandlersVMCheckK3sResponse) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### GetDecision

`func (o *HandlersVMCheckK3sResponse) GetDecision() VmpreflightRejection`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *HandlersVMCheckK3sResponse) GetDecisionOk() (*VmpreflightRejection, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *HandlersVMCheckK3sResponse) SetDecision(v VmpreflightRejection)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *HandlersVMCheckK3sResponse) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetK3sState

`func (o *HandlersVMCheckK3sResponse) GetK3sState() string`

GetK3sState returns the K3sState field if non-nil, zero value otherwise.

### GetK3sStateOk

`func (o *HandlersVMCheckK3sResponse) GetK3sStateOk() (*string, bool)`

GetK3sStateOk returns a tuple with the K3sState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK3sState

`func (o *HandlersVMCheckK3sResponse) SetK3sState(v string)`

SetK3sState sets K3sState field to given value.

### HasK3sState

`func (o *HandlersVMCheckK3sResponse) HasK3sState() bool`

HasK3sState returns a boolean if a field has been set.

### GetProbe

`func (o *HandlersVMCheckK3sResponse) GetProbe() AgentopsCheckK3sResult`

GetProbe returns the Probe field if non-nil, zero value otherwise.

### GetProbeOk

`func (o *HandlersVMCheckK3sResponse) GetProbeOk() (*AgentopsCheckK3sResult, bool)`

GetProbeOk returns a tuple with the Probe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbe

`func (o *HandlersVMCheckK3sResponse) SetProbe(v AgentopsCheckK3sResult)`

SetProbe sets Probe field to given value.

### HasProbe

`func (o *HandlersVMCheckK3sResponse) HasProbe() bool`

HasProbe returns a boolean if a field has been set.

### GetVmId

`func (o *HandlersVMCheckK3sResponse) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *HandlersVMCheckK3sResponse) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *HandlersVMCheckK3sResponse) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *HandlersVMCheckK3sResponse) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


