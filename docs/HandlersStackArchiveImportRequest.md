# HandlersStackArchiveImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproveCollisions** | Pointer to **bool** | ApproveCollisions is the staged-import review&#39;s grant. | [optional] 
**Confirms** | Pointer to **map[string]string** | Confirms echoes the dispatch tokens the dialog minted, per archived component. | [optional] 
**ExportKey** | Pointer to **string** | ExportKey pins the export (empty &#x3D; the stack&#39;s latest pickable set). | [optional] 
**Facts** | Pointer to [**EngineContextFacts**](EngineContextFacts.md) | Facts are the TARGET cluster&#39;s context facts (the same source the install route&#39;s resolved payload carries — the panel&#39;s cluster context). The import compiles against them; a dispatch without them refuses at the worker. | [optional] 
**SourceClusterId** | Pointer to **int32** | SourceClusterID names the cluster the export ran on (the set&#39;s recorded scope; same project — the copy is same-org, the move threat model is not this story&#39;s door to open). | [optional] 
**Stack** | Pointer to **string** | Stack is the stack manifest&#39;s name the export recorded. | [optional] 

## Methods

### NewHandlersStackArchiveImportRequest

`func NewHandlersStackArchiveImportRequest() *HandlersStackArchiveImportRequest`

NewHandlersStackArchiveImportRequest instantiates a new HandlersStackArchiveImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStackArchiveImportRequestWithDefaults

`func NewHandlersStackArchiveImportRequestWithDefaults() *HandlersStackArchiveImportRequest`

NewHandlersStackArchiveImportRequestWithDefaults instantiates a new HandlersStackArchiveImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproveCollisions

`func (o *HandlersStackArchiveImportRequest) GetApproveCollisions() bool`

GetApproveCollisions returns the ApproveCollisions field if non-nil, zero value otherwise.

### GetApproveCollisionsOk

`func (o *HandlersStackArchiveImportRequest) GetApproveCollisionsOk() (*bool, bool)`

GetApproveCollisionsOk returns a tuple with the ApproveCollisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveCollisions

`func (o *HandlersStackArchiveImportRequest) SetApproveCollisions(v bool)`

SetApproveCollisions sets ApproveCollisions field to given value.

### HasApproveCollisions

`func (o *HandlersStackArchiveImportRequest) HasApproveCollisions() bool`

HasApproveCollisions returns a boolean if a field has been set.

### GetConfirms

`func (o *HandlersStackArchiveImportRequest) GetConfirms() map[string]string`

GetConfirms returns the Confirms field if non-nil, zero value otherwise.

### GetConfirmsOk

`func (o *HandlersStackArchiveImportRequest) GetConfirmsOk() (*map[string]string, bool)`

GetConfirmsOk returns a tuple with the Confirms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirms

`func (o *HandlersStackArchiveImportRequest) SetConfirms(v map[string]string)`

SetConfirms sets Confirms field to given value.

### HasConfirms

`func (o *HandlersStackArchiveImportRequest) HasConfirms() bool`

HasConfirms returns a boolean if a field has been set.

### GetExportKey

`func (o *HandlersStackArchiveImportRequest) GetExportKey() string`

GetExportKey returns the ExportKey field if non-nil, zero value otherwise.

### GetExportKeyOk

`func (o *HandlersStackArchiveImportRequest) GetExportKeyOk() (*string, bool)`

GetExportKeyOk returns a tuple with the ExportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportKey

`func (o *HandlersStackArchiveImportRequest) SetExportKey(v string)`

SetExportKey sets ExportKey field to given value.

### HasExportKey

`func (o *HandlersStackArchiveImportRequest) HasExportKey() bool`

HasExportKey returns a boolean if a field has been set.

### GetFacts

`func (o *HandlersStackArchiveImportRequest) GetFacts() EngineContextFacts`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *HandlersStackArchiveImportRequest) GetFactsOk() (*EngineContextFacts, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *HandlersStackArchiveImportRequest) SetFacts(v EngineContextFacts)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *HandlersStackArchiveImportRequest) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetSourceClusterId

`func (o *HandlersStackArchiveImportRequest) GetSourceClusterId() int32`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### GetSourceClusterIdOk

`func (o *HandlersStackArchiveImportRequest) GetSourceClusterIdOk() (*int32, bool)`

GetSourceClusterIdOk returns a tuple with the SourceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterId

`func (o *HandlersStackArchiveImportRequest) SetSourceClusterId(v int32)`

SetSourceClusterId sets SourceClusterId field to given value.

### HasSourceClusterId

`func (o *HandlersStackArchiveImportRequest) HasSourceClusterId() bool`

HasSourceClusterId returns a boolean if a field has been set.

### GetStack

`func (o *HandlersStackArchiveImportRequest) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *HandlersStackArchiveImportRequest) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *HandlersStackArchiveImportRequest) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *HandlersStackArchiveImportRequest) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


