# HandlersBlueprintUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyRefusal** | Pointer to [**HandlersApplyRefusal**](HandlersApplyRefusal.md) |  | [optional] 
**Definitions** | Pointer to [**HandlersDeploykitDefinitionStatus**](HandlersDeploykitDefinitionStatus.md) | Definitions is the definition record&#39;s status block (entry 9&#39;s shape, entry 6&#39;s surfacing): the latest revision&#39;s pin and the source lifecycle&#39;s named state, on the surface the user stares at. | [optional] 
**PreviewNote** | Pointer to **string** | PreviewNote annotates a DEFERRED delta preview (the mid-roll honesty gate, mirrored from the route&#39;s own gate): a stack transition is in flight, so the as-deployed capture is about to move and a delta previewed now would not be the delta the sweeper runs. The preview resumes once the transition lands (the route refuses dispatching until then); the compiled-artifact diff still rides. | [optional] 
**Review** | Pointer to [**BlueprintdriftUpdateReview**](BlueprintdriftUpdateReview.md) |  | [optional] 
**SourceReconnect** | Pointer to [**HandlersBlueprintSourceReconnect**](HandlersBlueprintSourceReconnect.md) | SourceReconnect names the D4 reconnect-review path when the source sits orphaned: a fresh review&#39;s confirmed apply mints a new definition revision and journals the break. | [optional] 
**StackReview** | Pointer to [**HandlersStackReviewPayload**](HandlersStackReviewPayload.md) | StackReview is the SHARED stack transition review&#39;s delta (epic stack-evolution entry 6): a stack parent&#39;s lawful set change renders the entry-2 review — who joins, what they wire to, what dies, the data fate, the wave plan — instead of the dead-end refusal. Absent for single-component installs and refused transitions. | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersBlueprintUpdateResponse

`func NewHandlersBlueprintUpdateResponse() *HandlersBlueprintUpdateResponse`

NewHandlersBlueprintUpdateResponse instantiates a new HandlersBlueprintUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersBlueprintUpdateResponseWithDefaults

`func NewHandlersBlueprintUpdateResponseWithDefaults() *HandlersBlueprintUpdateResponse`

NewHandlersBlueprintUpdateResponseWithDefaults instantiates a new HandlersBlueprintUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyRefusal

`func (o *HandlersBlueprintUpdateResponse) GetApplyRefusal() HandlersApplyRefusal`

GetApplyRefusal returns the ApplyRefusal field if non-nil, zero value otherwise.

### GetApplyRefusalOk

`func (o *HandlersBlueprintUpdateResponse) GetApplyRefusalOk() (*HandlersApplyRefusal, bool)`

GetApplyRefusalOk returns a tuple with the ApplyRefusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyRefusal

`func (o *HandlersBlueprintUpdateResponse) SetApplyRefusal(v HandlersApplyRefusal)`

SetApplyRefusal sets ApplyRefusal field to given value.

### HasApplyRefusal

`func (o *HandlersBlueprintUpdateResponse) HasApplyRefusal() bool`

HasApplyRefusal returns a boolean if a field has been set.

### GetDefinitions

`func (o *HandlersBlueprintUpdateResponse) GetDefinitions() HandlersDeploykitDefinitionStatus`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *HandlersBlueprintUpdateResponse) GetDefinitionsOk() (*HandlersDeploykitDefinitionStatus, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *HandlersBlueprintUpdateResponse) SetDefinitions(v HandlersDeploykitDefinitionStatus)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *HandlersBlueprintUpdateResponse) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetPreviewNote

`func (o *HandlersBlueprintUpdateResponse) GetPreviewNote() string`

GetPreviewNote returns the PreviewNote field if non-nil, zero value otherwise.

### GetPreviewNoteOk

`func (o *HandlersBlueprintUpdateResponse) GetPreviewNoteOk() (*string, bool)`

GetPreviewNoteOk returns a tuple with the PreviewNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewNote

`func (o *HandlersBlueprintUpdateResponse) SetPreviewNote(v string)`

SetPreviewNote sets PreviewNote field to given value.

### HasPreviewNote

`func (o *HandlersBlueprintUpdateResponse) HasPreviewNote() bool`

HasPreviewNote returns a boolean if a field has been set.

### GetReview

`func (o *HandlersBlueprintUpdateResponse) GetReview() BlueprintdriftUpdateReview`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *HandlersBlueprintUpdateResponse) GetReviewOk() (*BlueprintdriftUpdateReview, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *HandlersBlueprintUpdateResponse) SetReview(v BlueprintdriftUpdateReview)`

SetReview sets Review field to given value.

### HasReview

`func (o *HandlersBlueprintUpdateResponse) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetSourceReconnect

`func (o *HandlersBlueprintUpdateResponse) GetSourceReconnect() HandlersBlueprintSourceReconnect`

GetSourceReconnect returns the SourceReconnect field if non-nil, zero value otherwise.

### GetSourceReconnectOk

`func (o *HandlersBlueprintUpdateResponse) GetSourceReconnectOk() (*HandlersBlueprintSourceReconnect, bool)`

GetSourceReconnectOk returns a tuple with the SourceReconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceReconnect

`func (o *HandlersBlueprintUpdateResponse) SetSourceReconnect(v HandlersBlueprintSourceReconnect)`

SetSourceReconnect sets SourceReconnect field to given value.

### HasSourceReconnect

`func (o *HandlersBlueprintUpdateResponse) HasSourceReconnect() bool`

HasSourceReconnect returns a boolean if a field has been set.

### GetStackReview

`func (o *HandlersBlueprintUpdateResponse) GetStackReview() HandlersStackReviewPayload`

GetStackReview returns the StackReview field if non-nil, zero value otherwise.

### GetStackReviewOk

`func (o *HandlersBlueprintUpdateResponse) GetStackReviewOk() (*HandlersStackReviewPayload, bool)`

GetStackReviewOk returns a tuple with the StackReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackReview

`func (o *HandlersBlueprintUpdateResponse) SetStackReview(v HandlersStackReviewPayload)`

SetStackReview sets StackReview field to given value.

### HasStackReview

`func (o *HandlersBlueprintUpdateResponse) HasStackReview() bool`

HasStackReview returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersBlueprintUpdateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersBlueprintUpdateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersBlueprintUpdateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersBlueprintUpdateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


