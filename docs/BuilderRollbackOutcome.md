# BuilderRollbackOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to **bool** |  | [optional] 
**Confirmation** | Pointer to [**BuilderRollbackConfirmation**](BuilderRollbackConfirmation.md) | Confirmation is the FR48 payload. | [optional] 
**DurationMs** | Pointer to **int32** | DurationMs is the measured wall-clock of the whole rollback (re-point + probe(s)) on the Gate&#39;s clock; WithinBudget is the ≤30 s budget&#39;s honest comparison. | [optional] 
**Gate** | Pointer to [**ModelsBuildHealthGate**](ModelsBuildHealthGate.md) | Gate is the run&#39;s terminal row (proven / blocked / failed) — the timeline&#39;s health fact for the re-entered probe gate. | [optional] 
**Routed** | Pointer to **bool** | Routed / Blocked mirror the gate verdict (traffic routes on N-1, or the probe failed and nothing routes). | [optional] 
**Target** | Pointer to [**ModelsBuildRevision**](ModelsBuildRevision.md) | Target is the revision the service re-pointed at. | [optional] 
**WithinBudget** | Pointer to **bool** |  | [optional] 

## Methods

### NewBuilderRollbackOutcome

`func NewBuilderRollbackOutcome() *BuilderRollbackOutcome`

NewBuilderRollbackOutcome instantiates a new BuilderRollbackOutcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuilderRollbackOutcomeWithDefaults

`func NewBuilderRollbackOutcomeWithDefaults() *BuilderRollbackOutcome`

NewBuilderRollbackOutcomeWithDefaults instantiates a new BuilderRollbackOutcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *BuilderRollbackOutcome) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *BuilderRollbackOutcome) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *BuilderRollbackOutcome) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *BuilderRollbackOutcome) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetConfirmation

`func (o *BuilderRollbackOutcome) GetConfirmation() BuilderRollbackConfirmation`

GetConfirmation returns the Confirmation field if non-nil, zero value otherwise.

### GetConfirmationOk

`func (o *BuilderRollbackOutcome) GetConfirmationOk() (*BuilderRollbackConfirmation, bool)`

GetConfirmationOk returns a tuple with the Confirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmation

`func (o *BuilderRollbackOutcome) SetConfirmation(v BuilderRollbackConfirmation)`

SetConfirmation sets Confirmation field to given value.

### HasConfirmation

`func (o *BuilderRollbackOutcome) HasConfirmation() bool`

HasConfirmation returns a boolean if a field has been set.

### GetDurationMs

`func (o *BuilderRollbackOutcome) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *BuilderRollbackOutcome) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *BuilderRollbackOutcome) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *BuilderRollbackOutcome) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetGate

`func (o *BuilderRollbackOutcome) GetGate() ModelsBuildHealthGate`

GetGate returns the Gate field if non-nil, zero value otherwise.

### GetGateOk

`func (o *BuilderRollbackOutcome) GetGateOk() (*ModelsBuildHealthGate, bool)`

GetGateOk returns a tuple with the Gate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGate

`func (o *BuilderRollbackOutcome) SetGate(v ModelsBuildHealthGate)`

SetGate sets Gate field to given value.

### HasGate

`func (o *BuilderRollbackOutcome) HasGate() bool`

HasGate returns a boolean if a field has been set.

### GetRouted

`func (o *BuilderRollbackOutcome) GetRouted() bool`

GetRouted returns the Routed field if non-nil, zero value otherwise.

### GetRoutedOk

`func (o *BuilderRollbackOutcome) GetRoutedOk() (*bool, bool)`

GetRoutedOk returns a tuple with the Routed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouted

`func (o *BuilderRollbackOutcome) SetRouted(v bool)`

SetRouted sets Routed field to given value.

### HasRouted

`func (o *BuilderRollbackOutcome) HasRouted() bool`

HasRouted returns a boolean if a field has been set.

### GetTarget

`func (o *BuilderRollbackOutcome) GetTarget() ModelsBuildRevision`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BuilderRollbackOutcome) GetTargetOk() (*ModelsBuildRevision, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BuilderRollbackOutcome) SetTarget(v ModelsBuildRevision)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *BuilderRollbackOutcome) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetWithinBudget

`func (o *BuilderRollbackOutcome) GetWithinBudget() bool`

GetWithinBudget returns the WithinBudget field if non-nil, zero value otherwise.

### GetWithinBudgetOk

`func (o *BuilderRollbackOutcome) GetWithinBudgetOk() (*bool, bool)`

GetWithinBudgetOk returns a tuple with the WithinBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithinBudget

`func (o *BuilderRollbackOutcome) SetWithinBudget(v bool)`

SetWithinBudget sets WithinBudget field to given value.

### HasWithinBudget

`func (o *BuilderRollbackOutcome) HasWithinBudget() bool`

HasWithinBudget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


