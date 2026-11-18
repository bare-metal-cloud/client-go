# HandlersOrgVaultRotateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **int32** |  | [optional] 
**Checklist** | Pointer to [**[]HandlersRotationItem**](HandlersRotationItem.md) |  | [optional] 
**Machine** | Pointer to **string** |  | [optional] 
**RecordedAt** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersOrgVaultRotateResponse

`func NewHandlersOrgVaultRotateResponse() *HandlersOrgVaultRotateResponse`

NewHandlersOrgVaultRotateResponse instantiates a new HandlersOrgVaultRotateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersOrgVaultRotateResponseWithDefaults

`func NewHandlersOrgVaultRotateResponseWithDefaults() *HandlersOrgVaultRotateResponse`

NewHandlersOrgVaultRotateResponseWithDefaults instantiates a new HandlersOrgVaultRotateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *HandlersOrgVaultRotateResponse) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *HandlersOrgVaultRotateResponse) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *HandlersOrgVaultRotateResponse) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *HandlersOrgVaultRotateResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetChecklist

`func (o *HandlersOrgVaultRotateResponse) GetChecklist() []HandlersRotationItem`

GetChecklist returns the Checklist field if non-nil, zero value otherwise.

### GetChecklistOk

`func (o *HandlersOrgVaultRotateResponse) GetChecklistOk() (*[]HandlersRotationItem, bool)`

GetChecklistOk returns a tuple with the Checklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklist

`func (o *HandlersOrgVaultRotateResponse) SetChecklist(v []HandlersRotationItem)`

SetChecklist sets Checklist field to given value.

### HasChecklist

`func (o *HandlersOrgVaultRotateResponse) HasChecklist() bool`

HasChecklist returns a boolean if a field has been set.

### GetMachine

`func (o *HandlersOrgVaultRotateResponse) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *HandlersOrgVaultRotateResponse) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *HandlersOrgVaultRotateResponse) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *HandlersOrgVaultRotateResponse) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetRecordedAt

`func (o *HandlersOrgVaultRotateResponse) GetRecordedAt() string`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *HandlersOrgVaultRotateResponse) GetRecordedAtOk() (*string, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *HandlersOrgVaultRotateResponse) SetRecordedAt(v string)`

SetRecordedAt sets RecordedAt field to given value.

### HasRecordedAt

`func (o *HandlersOrgVaultRotateResponse) HasRecordedAt() bool`

HasRecordedAt returns a boolean if a field has been set.

### GetState

`func (o *HandlersOrgVaultRotateResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersOrgVaultRotateResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersOrgVaultRotateResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersOrgVaultRotateResponse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


