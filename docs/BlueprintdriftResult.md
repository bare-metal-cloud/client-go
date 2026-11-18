# BlueprintdriftResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprint** | Pointer to **bool** | Blueprint is the per-stack capability flag: this install is a blueprint deployment, so the stack view may offer Update. | [optional] 
**CheckedAt** | Pointer to **string** | CheckedAt is when the served result was checked. | [optional] 
**CommitsAhead** | Pointer to **int32** | CommitsAhead is the compare count (best-effort; 0 &#x3D; absent). | [optional] 
**CompareUrl** | Pointer to **string** | CompareURL is the upstream commit-compare link (context). | [optional] 
**Degraded** | Pointer to **bool** | Degraded marks a check that could not run (quota spent, provider error): NO hint — never an error to the view. | [optional] 
**DegradedReason** | Pointer to [**BlueprintdriftDegradedReason**](BlueprintdriftDegradedReason.md) | DegradedReason names WHY a degraded check could not run (issue #3775): credential_unresolved (the source-pull credential cannot mint — a revoked/suspended GitHub App installation), probe_failed (a transient provider/probe failure), or quota (the shared budget refused). Empty when not degraded. ADDITIVE: consumers that read state/degraded only are unchanged. | [optional] 
**DigestDrift** | Pointer to [**[]BlueprintdriftDigestDriftRow**](BlueprintdriftDigestDriftRow.md) | DigestDrift carries the tag-level movements (the digest-moved state&#39;s rows; attached to ref-moved too). | [optional] 
**Dismissed** | Pointer to **bool** | Dismissed marks THIS instance as dismissed by its owner (the FE renders nothing; the state survives until a new upstream SHA range). | [optional] 
**InstalledSha** | Pointer to **string** | InstalledSHA / UpstreamSHA are the movement&#39;s two ends (UpstreamSHA empty on a ref that no longer resolves). | [optional] 
**Key** | Pointer to **string** | Key is THIS instance&#39;s dismissal key (the upstream SHA range, or the state key) — the FE&#39;s dismiss call echoes it back. | [optional] 
**SlotDrift** | Pointer to [**[]BlueprintdriftSlotDriftRow**](BlueprintdriftSlotDriftRow.md) | SlotDrift carries the managed-slot declaration movements (the managed-xyz story 7 AC1 half): the repo file&#39;s slot declaration is the baseline, and a changed floor or engine names its rows here. A bound-version move that satisfies the floor never lands here — R4: the floor is a minimum, satisfying it is not drift. | [optional] 
**State** | Pointer to [**BlueprintdriftHintState**](BlueprintdriftHintState.md) | State is the named state. \&quot;none\&quot; (and \&quot;none\&quot; + Degraded) renders NO hint — the FE renders nothing either way. | [optional] 
**UpstreamSha** | Pointer to **string** |  | [optional] 

## Methods

### NewBlueprintdriftResult

`func NewBlueprintdriftResult() *BlueprintdriftResult`

NewBlueprintdriftResult instantiates a new BlueprintdriftResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintdriftResultWithDefaults

`func NewBlueprintdriftResultWithDefaults() *BlueprintdriftResult`

NewBlueprintdriftResultWithDefaults instantiates a new BlueprintdriftResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprint

`func (o *BlueprintdriftResult) GetBlueprint() bool`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *BlueprintdriftResult) GetBlueprintOk() (*bool, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *BlueprintdriftResult) SetBlueprint(v bool)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *BlueprintdriftResult) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetCheckedAt

`func (o *BlueprintdriftResult) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *BlueprintdriftResult) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *BlueprintdriftResult) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *BlueprintdriftResult) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### GetCommitsAhead

`func (o *BlueprintdriftResult) GetCommitsAhead() int32`

GetCommitsAhead returns the CommitsAhead field if non-nil, zero value otherwise.

### GetCommitsAheadOk

`func (o *BlueprintdriftResult) GetCommitsAheadOk() (*int32, bool)`

GetCommitsAheadOk returns a tuple with the CommitsAhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsAhead

`func (o *BlueprintdriftResult) SetCommitsAhead(v int32)`

SetCommitsAhead sets CommitsAhead field to given value.

### HasCommitsAhead

`func (o *BlueprintdriftResult) HasCommitsAhead() bool`

HasCommitsAhead returns a boolean if a field has been set.

### GetCompareUrl

`func (o *BlueprintdriftResult) GetCompareUrl() string`

GetCompareUrl returns the CompareUrl field if non-nil, zero value otherwise.

### GetCompareUrlOk

`func (o *BlueprintdriftResult) GetCompareUrlOk() (*string, bool)`

GetCompareUrlOk returns a tuple with the CompareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareUrl

`func (o *BlueprintdriftResult) SetCompareUrl(v string)`

SetCompareUrl sets CompareUrl field to given value.

### HasCompareUrl

`func (o *BlueprintdriftResult) HasCompareUrl() bool`

HasCompareUrl returns a boolean if a field has been set.

### GetDegraded

`func (o *BlueprintdriftResult) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *BlueprintdriftResult) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *BlueprintdriftResult) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *BlueprintdriftResult) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetDegradedReason

`func (o *BlueprintdriftResult) GetDegradedReason() BlueprintdriftDegradedReason`

GetDegradedReason returns the DegradedReason field if non-nil, zero value otherwise.

### GetDegradedReasonOk

`func (o *BlueprintdriftResult) GetDegradedReasonOk() (*BlueprintdriftDegradedReason, bool)`

GetDegradedReasonOk returns a tuple with the DegradedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedReason

`func (o *BlueprintdriftResult) SetDegradedReason(v BlueprintdriftDegradedReason)`

SetDegradedReason sets DegradedReason field to given value.

### HasDegradedReason

`func (o *BlueprintdriftResult) HasDegradedReason() bool`

HasDegradedReason returns a boolean if a field has been set.

### GetDigestDrift

`func (o *BlueprintdriftResult) GetDigestDrift() []BlueprintdriftDigestDriftRow`

GetDigestDrift returns the DigestDrift field if non-nil, zero value otherwise.

### GetDigestDriftOk

`func (o *BlueprintdriftResult) GetDigestDriftOk() (*[]BlueprintdriftDigestDriftRow, bool)`

GetDigestDriftOk returns a tuple with the DigestDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestDrift

`func (o *BlueprintdriftResult) SetDigestDrift(v []BlueprintdriftDigestDriftRow)`

SetDigestDrift sets DigestDrift field to given value.

### HasDigestDrift

`func (o *BlueprintdriftResult) HasDigestDrift() bool`

HasDigestDrift returns a boolean if a field has been set.

### GetDismissed

`func (o *BlueprintdriftResult) GetDismissed() bool`

GetDismissed returns the Dismissed field if non-nil, zero value otherwise.

### GetDismissedOk

`func (o *BlueprintdriftResult) GetDismissedOk() (*bool, bool)`

GetDismissedOk returns a tuple with the Dismissed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissed

`func (o *BlueprintdriftResult) SetDismissed(v bool)`

SetDismissed sets Dismissed field to given value.

### HasDismissed

`func (o *BlueprintdriftResult) HasDismissed() bool`

HasDismissed returns a boolean if a field has been set.

### GetInstalledSha

`func (o *BlueprintdriftResult) GetInstalledSha() string`

GetInstalledSha returns the InstalledSha field if non-nil, zero value otherwise.

### GetInstalledShaOk

`func (o *BlueprintdriftResult) GetInstalledShaOk() (*string, bool)`

GetInstalledShaOk returns a tuple with the InstalledSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledSha

`func (o *BlueprintdriftResult) SetInstalledSha(v string)`

SetInstalledSha sets InstalledSha field to given value.

### HasInstalledSha

`func (o *BlueprintdriftResult) HasInstalledSha() bool`

HasInstalledSha returns a boolean if a field has been set.

### GetKey

`func (o *BlueprintdriftResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BlueprintdriftResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BlueprintdriftResult) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BlueprintdriftResult) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSlotDrift

`func (o *BlueprintdriftResult) GetSlotDrift() []BlueprintdriftSlotDriftRow`

GetSlotDrift returns the SlotDrift field if non-nil, zero value otherwise.

### GetSlotDriftOk

`func (o *BlueprintdriftResult) GetSlotDriftOk() (*[]BlueprintdriftSlotDriftRow, bool)`

GetSlotDriftOk returns a tuple with the SlotDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotDrift

`func (o *BlueprintdriftResult) SetSlotDrift(v []BlueprintdriftSlotDriftRow)`

SetSlotDrift sets SlotDrift field to given value.

### HasSlotDrift

`func (o *BlueprintdriftResult) HasSlotDrift() bool`

HasSlotDrift returns a boolean if a field has been set.

### GetState

`func (o *BlueprintdriftResult) GetState() BlueprintdriftHintState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BlueprintdriftResult) GetStateOk() (*BlueprintdriftHintState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BlueprintdriftResult) SetState(v BlueprintdriftHintState)`

SetState sets State field to given value.

### HasState

`func (o *BlueprintdriftResult) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpstreamSha

`func (o *BlueprintdriftResult) GetUpstreamSha() string`

GetUpstreamSha returns the UpstreamSha field if non-nil, zero value otherwise.

### GetUpstreamShaOk

`func (o *BlueprintdriftResult) GetUpstreamShaOk() (*string, bool)`

GetUpstreamShaOk returns a tuple with the UpstreamSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamSha

`func (o *BlueprintdriftResult) SetUpstreamSha(v string)`

SetUpstreamSha sets UpstreamSha field to given value.

### HasUpstreamSha

`func (o *BlueprintdriftResult) HasUpstreamSha() bool`

HasUpstreamSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


