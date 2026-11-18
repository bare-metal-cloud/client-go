# ClonereplayReportUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** | Component names the stack component (kind stack only). | [optional] 
**DataClass** | Pointer to **string** | DataClass is the unit&#39;s declared data class (the plan&#39;s), shown so the report states how the data classes decided the replay. | [optional] 
**Err** | Pointer to **string** | Err carries the failing unit&#39;s refusal/failure text (failed lines). | [optional] 
**Install** | Pointer to **string** | Install is the target install&#39;s name. | [optional] 
**Kind** | Pointer to **string** | Kind is the unit kind: stack component or loose install. | [optional] 
**Outcome** | Pointer to **string** | Outcome is deployed | already-replayed | failed. | [optional] 
**Provenance** | Pointer to **string** | Provenance is the stamped provenance: cloned for every deployed unit (the 134.3 law). | [optional] 
**ReplayEmpty** | Pointer to **bool** | ReplayEmpty marks the stateful-authoritative-with-no-seed-hook case: the unit replayed EMPTY, and the report says so — never a guess. | [optional] 
**Revision** | Pointer to **int32** | Revision is the revision the unit&#39;s deploy minted (deployed lines). | [optional] 
**SlotNotes** | Pointer to **[]string** | SlotNotes carries the interim disclosure lines the derive rendered for the owning unit&#39;s managed slots (managed-xyz story A, issue #4275&#39;s wire delta: backward-compatible JSON). The settled report states the working-but-empty posture beside the outcome. Story B retracts the line when the seed lands; story C adds the durability line. | [optional] 
**Stack** | Pointer to **string** | Stack names the owning stack (kind stack only). | [optional] 

## Methods

### NewClonereplayReportUnit

`func NewClonereplayReportUnit() *ClonereplayReportUnit`

NewClonereplayReportUnit instantiates a new ClonereplayReportUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClonereplayReportUnitWithDefaults

`func NewClonereplayReportUnitWithDefaults() *ClonereplayReportUnit`

NewClonereplayReportUnitWithDefaults instantiates a new ClonereplayReportUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ClonereplayReportUnit) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ClonereplayReportUnit) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ClonereplayReportUnit) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ClonereplayReportUnit) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetDataClass

`func (o *ClonereplayReportUnit) GetDataClass() string`

GetDataClass returns the DataClass field if non-nil, zero value otherwise.

### GetDataClassOk

`func (o *ClonereplayReportUnit) GetDataClassOk() (*string, bool)`

GetDataClassOk returns a tuple with the DataClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataClass

`func (o *ClonereplayReportUnit) SetDataClass(v string)`

SetDataClass sets DataClass field to given value.

### HasDataClass

`func (o *ClonereplayReportUnit) HasDataClass() bool`

HasDataClass returns a boolean if a field has been set.

### GetErr

`func (o *ClonereplayReportUnit) GetErr() string`

GetErr returns the Err field if non-nil, zero value otherwise.

### GetErrOk

`func (o *ClonereplayReportUnit) GetErrOk() (*string, bool)`

GetErrOk returns a tuple with the Err field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErr

`func (o *ClonereplayReportUnit) SetErr(v string)`

SetErr sets Err field to given value.

### HasErr

`func (o *ClonereplayReportUnit) HasErr() bool`

HasErr returns a boolean if a field has been set.

### GetInstall

`func (o *ClonereplayReportUnit) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *ClonereplayReportUnit) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *ClonereplayReportUnit) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *ClonereplayReportUnit) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetKind

`func (o *ClonereplayReportUnit) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClonereplayReportUnit) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClonereplayReportUnit) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClonereplayReportUnit) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOutcome

`func (o *ClonereplayReportUnit) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ClonereplayReportUnit) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ClonereplayReportUnit) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *ClonereplayReportUnit) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetProvenance

`func (o *ClonereplayReportUnit) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *ClonereplayReportUnit) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *ClonereplayReportUnit) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *ClonereplayReportUnit) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetReplayEmpty

`func (o *ClonereplayReportUnit) GetReplayEmpty() bool`

GetReplayEmpty returns the ReplayEmpty field if non-nil, zero value otherwise.

### GetReplayEmptyOk

`func (o *ClonereplayReportUnit) GetReplayEmptyOk() (*bool, bool)`

GetReplayEmptyOk returns a tuple with the ReplayEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayEmpty

`func (o *ClonereplayReportUnit) SetReplayEmpty(v bool)`

SetReplayEmpty sets ReplayEmpty field to given value.

### HasReplayEmpty

`func (o *ClonereplayReportUnit) HasReplayEmpty() bool`

HasReplayEmpty returns a boolean if a field has been set.

### GetRevision

`func (o *ClonereplayReportUnit) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ClonereplayReportUnit) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ClonereplayReportUnit) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ClonereplayReportUnit) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSlotNotes

`func (o *ClonereplayReportUnit) GetSlotNotes() []string`

GetSlotNotes returns the SlotNotes field if non-nil, zero value otherwise.

### GetSlotNotesOk

`func (o *ClonereplayReportUnit) GetSlotNotesOk() (*[]string, bool)`

GetSlotNotesOk returns a tuple with the SlotNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNotes

`func (o *ClonereplayReportUnit) SetSlotNotes(v []string)`

SetSlotNotes sets SlotNotes field to given value.

### HasSlotNotes

`func (o *ClonereplayReportUnit) HasSlotNotes() bool`

HasSlotNotes returns a boolean if a field has been set.

### GetStack

`func (o *ClonereplayReportUnit) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *ClonereplayReportUnit) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *ClonereplayReportUnit) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *ClonereplayReportUnit) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


