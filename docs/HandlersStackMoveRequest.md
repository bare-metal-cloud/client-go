# HandlersStackMoveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproveCollisions** | Pointer to **bool** | ApproveCollisions is the staged-import review&#39;s grant. | [optional] 
**Confirms** | Pointer to **map[string]string** | Confirms echoes the dispatch tokens the dialog minted — the per-member tokens AND the window token. | [optional] 
**ExportKey** | Pointer to **string** | ExportKey pins the export (empty &#x3D; the stack&#39;s latest pickable set). | [optional] 
**Facts** | Pointer to [**EngineContextFacts**](EngineContextFacts.md) | Facts are the TARGET cluster&#39;s context facts (the import leg compiles against them). | [optional] 
**Flavor** | Pointer to **string** | Flavor is the move path: \&quot;cold\&quot; (the fence-first default) or \&quot;replication\&quot; (the fast path — refused unless the recorded set&#39;s every stateful-authoritative member declared &#x60;move: replication&#x60;; an incapable set is NEVER silently downgraded, it is refused with the cold path named). Empty means cold. | [optional] 
**Retire** | Pointer to **string** | Retire is the source-retirement choice — explicit, never a default: \&quot;keep-data\&quot; (the retired-original tombstone as standing cancellation) or \&quot;purge\&quot; (destroy; the move&#39;s target probe gate is the verification). | [optional] 
**SourceClusterId** | Pointer to **int32** | SourceClusterID names the cluster the export ran on (the set&#39;s recorded scope; same project — the SAME-ORG gate runs against the source cluster&#39;s recorded org). | [optional] 
**Stack** | Pointer to **string** | Stack is the stack manifest&#39;s name the export recorded. | [optional] 

## Methods

### NewHandlersStackMoveRequest

`func NewHandlersStackMoveRequest() *HandlersStackMoveRequest`

NewHandlersStackMoveRequest instantiates a new HandlersStackMoveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStackMoveRequestWithDefaults

`func NewHandlersStackMoveRequestWithDefaults() *HandlersStackMoveRequest`

NewHandlersStackMoveRequestWithDefaults instantiates a new HandlersStackMoveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproveCollisions

`func (o *HandlersStackMoveRequest) GetApproveCollisions() bool`

GetApproveCollisions returns the ApproveCollisions field if non-nil, zero value otherwise.

### GetApproveCollisionsOk

`func (o *HandlersStackMoveRequest) GetApproveCollisionsOk() (*bool, bool)`

GetApproveCollisionsOk returns a tuple with the ApproveCollisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveCollisions

`func (o *HandlersStackMoveRequest) SetApproveCollisions(v bool)`

SetApproveCollisions sets ApproveCollisions field to given value.

### HasApproveCollisions

`func (o *HandlersStackMoveRequest) HasApproveCollisions() bool`

HasApproveCollisions returns a boolean if a field has been set.

### GetConfirms

`func (o *HandlersStackMoveRequest) GetConfirms() map[string]string`

GetConfirms returns the Confirms field if non-nil, zero value otherwise.

### GetConfirmsOk

`func (o *HandlersStackMoveRequest) GetConfirmsOk() (*map[string]string, bool)`

GetConfirmsOk returns a tuple with the Confirms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirms

`func (o *HandlersStackMoveRequest) SetConfirms(v map[string]string)`

SetConfirms sets Confirms field to given value.

### HasConfirms

`func (o *HandlersStackMoveRequest) HasConfirms() bool`

HasConfirms returns a boolean if a field has been set.

### GetExportKey

`func (o *HandlersStackMoveRequest) GetExportKey() string`

GetExportKey returns the ExportKey field if non-nil, zero value otherwise.

### GetExportKeyOk

`func (o *HandlersStackMoveRequest) GetExportKeyOk() (*string, bool)`

GetExportKeyOk returns a tuple with the ExportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportKey

`func (o *HandlersStackMoveRequest) SetExportKey(v string)`

SetExportKey sets ExportKey field to given value.

### HasExportKey

`func (o *HandlersStackMoveRequest) HasExportKey() bool`

HasExportKey returns a boolean if a field has been set.

### GetFacts

`func (o *HandlersStackMoveRequest) GetFacts() EngineContextFacts`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *HandlersStackMoveRequest) GetFactsOk() (*EngineContextFacts, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *HandlersStackMoveRequest) SetFacts(v EngineContextFacts)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *HandlersStackMoveRequest) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetFlavor

`func (o *HandlersStackMoveRequest) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *HandlersStackMoveRequest) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *HandlersStackMoveRequest) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *HandlersStackMoveRequest) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetRetire

`func (o *HandlersStackMoveRequest) GetRetire() string`

GetRetire returns the Retire field if non-nil, zero value otherwise.

### GetRetireOk

`func (o *HandlersStackMoveRequest) GetRetireOk() (*string, bool)`

GetRetireOk returns a tuple with the Retire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetire

`func (o *HandlersStackMoveRequest) SetRetire(v string)`

SetRetire sets Retire field to given value.

### HasRetire

`func (o *HandlersStackMoveRequest) HasRetire() bool`

HasRetire returns a boolean if a field has been set.

### GetSourceClusterId

`func (o *HandlersStackMoveRequest) GetSourceClusterId() int32`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### GetSourceClusterIdOk

`func (o *HandlersStackMoveRequest) GetSourceClusterIdOk() (*int32, bool)`

GetSourceClusterIdOk returns a tuple with the SourceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterId

`func (o *HandlersStackMoveRequest) SetSourceClusterId(v int32)`

SetSourceClusterId sets SourceClusterId field to given value.

### HasSourceClusterId

`func (o *HandlersStackMoveRequest) HasSourceClusterId() bool`

HasSourceClusterId returns a boolean if a field has been set.

### GetStack

`func (o *HandlersStackMoveRequest) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *HandlersStackMoveRequest) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *HandlersStackMoveRequest) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *HandlersStackMoveRequest) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


