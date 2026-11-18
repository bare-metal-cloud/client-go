# HandlersStackCopyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproveCollisions** | Pointer to **bool** | ApproveCollisions is the staged-import review&#39;s grant. | [optional] 
**Confirms** | Pointer to **map[string]string** | Confirms echoes the dispatch tokens the dialog minted. | [optional] 
**ExportKey** | Pointer to **string** | ExportKey pins the export (empty &#x3D; the stack&#39;s latest pickable set). | [optional] 
**Facts** | Pointer to [**EngineContextFacts**](EngineContextFacts.md) | Facts are the TARGET cluster&#39;s context facts (the import leg compiles against them). | [optional] 
**Fence** | Pointer to **bool** | Fence is the copy&#39;s OPTIONAL source fence — DEFAULT OFF (the zero value: unfenced is the copy contract). | [optional] 
**SourceClusterId** | Pointer to **int32** | SourceClusterID names the cluster the export ran on (the set&#39;s recorded scope; same project — the SAME-ORG gate runs against the source cluster&#39;s recorded org). | [optional] 
**Stack** | Pointer to **string** | Stack is the stack manifest&#39;s name the export recorded. | [optional] 

## Methods

### NewHandlersStackCopyRequest

`func NewHandlersStackCopyRequest() *HandlersStackCopyRequest`

NewHandlersStackCopyRequest instantiates a new HandlersStackCopyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStackCopyRequestWithDefaults

`func NewHandlersStackCopyRequestWithDefaults() *HandlersStackCopyRequest`

NewHandlersStackCopyRequestWithDefaults instantiates a new HandlersStackCopyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproveCollisions

`func (o *HandlersStackCopyRequest) GetApproveCollisions() bool`

GetApproveCollisions returns the ApproveCollisions field if non-nil, zero value otherwise.

### GetApproveCollisionsOk

`func (o *HandlersStackCopyRequest) GetApproveCollisionsOk() (*bool, bool)`

GetApproveCollisionsOk returns a tuple with the ApproveCollisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveCollisions

`func (o *HandlersStackCopyRequest) SetApproveCollisions(v bool)`

SetApproveCollisions sets ApproveCollisions field to given value.

### HasApproveCollisions

`func (o *HandlersStackCopyRequest) HasApproveCollisions() bool`

HasApproveCollisions returns a boolean if a field has been set.

### GetConfirms

`func (o *HandlersStackCopyRequest) GetConfirms() map[string]string`

GetConfirms returns the Confirms field if non-nil, zero value otherwise.

### GetConfirmsOk

`func (o *HandlersStackCopyRequest) GetConfirmsOk() (*map[string]string, bool)`

GetConfirmsOk returns a tuple with the Confirms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirms

`func (o *HandlersStackCopyRequest) SetConfirms(v map[string]string)`

SetConfirms sets Confirms field to given value.

### HasConfirms

`func (o *HandlersStackCopyRequest) HasConfirms() bool`

HasConfirms returns a boolean if a field has been set.

### GetExportKey

`func (o *HandlersStackCopyRequest) GetExportKey() string`

GetExportKey returns the ExportKey field if non-nil, zero value otherwise.

### GetExportKeyOk

`func (o *HandlersStackCopyRequest) GetExportKeyOk() (*string, bool)`

GetExportKeyOk returns a tuple with the ExportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportKey

`func (o *HandlersStackCopyRequest) SetExportKey(v string)`

SetExportKey sets ExportKey field to given value.

### HasExportKey

`func (o *HandlersStackCopyRequest) HasExportKey() bool`

HasExportKey returns a boolean if a field has been set.

### GetFacts

`func (o *HandlersStackCopyRequest) GetFacts() EngineContextFacts`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *HandlersStackCopyRequest) GetFactsOk() (*EngineContextFacts, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *HandlersStackCopyRequest) SetFacts(v EngineContextFacts)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *HandlersStackCopyRequest) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetFence

`func (o *HandlersStackCopyRequest) GetFence() bool`

GetFence returns the Fence field if non-nil, zero value otherwise.

### GetFenceOk

`func (o *HandlersStackCopyRequest) GetFenceOk() (*bool, bool)`

GetFenceOk returns a tuple with the Fence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFence

`func (o *HandlersStackCopyRequest) SetFence(v bool)`

SetFence sets Fence field to given value.

### HasFence

`func (o *HandlersStackCopyRequest) HasFence() bool`

HasFence returns a boolean if a field has been set.

### GetSourceClusterId

`func (o *HandlersStackCopyRequest) GetSourceClusterId() int32`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### GetSourceClusterIdOk

`func (o *HandlersStackCopyRequest) GetSourceClusterIdOk() (*int32, bool)`

GetSourceClusterIdOk returns a tuple with the SourceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterId

`func (o *HandlersStackCopyRequest) SetSourceClusterId(v int32)`

SetSourceClusterId sets SourceClusterId field to given value.

### HasSourceClusterId

`func (o *HandlersStackCopyRequest) HasSourceClusterId() bool`

HasSourceClusterId returns a boolean if a field has been set.

### GetStack

`func (o *HandlersStackCopyRequest) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *HandlersStackCopyRequest) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *HandlersStackCopyRequest) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *HandlersStackCopyRequest) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


