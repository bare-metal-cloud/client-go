# CloneplanUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facts** | Pointer to [**CloneplanInstallFacts**](CloneplanInstallFacts.md) | Facts are the unit&#39;s install facts (identity + replay plan). | [optional] 
**SlotNotes** | Pointer to **[]string** | SlotNotes carries the disclosure lines the derive rendered for this unit&#39;s managed slots (managed-xyz story A, issue #4275&#39;s wire delta: backward-compatible JSON — omitempty, no schema bump). The interim line states the working-but-empty posture (\&quot;database provisions fresh at the target; starts empty — data copy follows in the seed story\&quot;) on the database class and the blast-radius line on the cache class (\&quot;sessions, counters, queues are not copied — they start empty\&quot;). Story B retracts the interim line when the seed lands; story C adds the durability line. | [optional] 
**Stack** | Pointer to [**CloneplanStackUnit**](CloneplanStackUnit.md) | Stack is non-nil exactly when the source install is stack-shaped (a StackDAGStore capture exists for it). | [optional] 

## Methods

### NewCloneplanUnit

`func NewCloneplanUnit() *CloneplanUnit`

NewCloneplanUnit instantiates a new CloneplanUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanUnitWithDefaults

`func NewCloneplanUnitWithDefaults() *CloneplanUnit`

NewCloneplanUnitWithDefaults instantiates a new CloneplanUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacts

`func (o *CloneplanUnit) GetFacts() CloneplanInstallFacts`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *CloneplanUnit) GetFactsOk() (*CloneplanInstallFacts, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *CloneplanUnit) SetFacts(v CloneplanInstallFacts)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *CloneplanUnit) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetSlotNotes

`func (o *CloneplanUnit) GetSlotNotes() []string`

GetSlotNotes returns the SlotNotes field if non-nil, zero value otherwise.

### GetSlotNotesOk

`func (o *CloneplanUnit) GetSlotNotesOk() (*[]string, bool)`

GetSlotNotesOk returns a tuple with the SlotNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNotes

`func (o *CloneplanUnit) SetSlotNotes(v []string)`

SetSlotNotes sets SlotNotes field to given value.

### HasSlotNotes

`func (o *CloneplanUnit) HasSlotNotes() bool`

HasSlotNotes returns a boolean if a field has been set.

### GetStack

`func (o *CloneplanUnit) GetStack() CloneplanStackUnit`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *CloneplanUnit) GetStackOk() (*CloneplanStackUnit, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *CloneplanUnit) SetStack(v CloneplanStackUnit)`

SetStack sets Stack field to given value.

### HasStack

`func (o *CloneplanUnit) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


