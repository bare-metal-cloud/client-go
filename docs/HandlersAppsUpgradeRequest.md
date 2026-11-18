# HandlersAppsUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyzeWindowSeconds** | Pointer to **int32** |  | [optional] 
**Confirm** | Pointer to **string** | Confirm is the user&#39;s fresh consent for a migrating upgrade (the dialog&#39;s confirm). Empty on the first POST — the dialog answer. | [optional] 
**Gates** | Pointer to [**[]EngineMetricGate**](EngineMetricGate.md) | Story 137.2 (issue #2962) — the rule-10 gate set. Gates is the upgrade&#39;s gate list (v1 evaluates the gateway provider; the reserved phase-2 provider names refuse at the plan compile below with the pointed error). Manual-kind gates compile into the plan&#39;s confirmations — the dispatch then demands a fresh confirm token like a migrating upgrade. Soak is the post-promotion soak window in seconds (0 &#x3D; none); RetainOldRevision pins the pre-upgrade revision across the soak (must cover it — the plan refuses otherwise); AnalyzeWindow is the pre-switch evaluation window (0 &#x3D; the 10m default). Durations ride seconds on the wire, nanoseconds on the queue payload. | [optional] 
**RetainOldRevisionSeconds** | Pointer to **int32** |  | [optional] 
**SoakSeconds** | Pointer to **int32** |  | [optional] 
**TargetDigest** | Pointer to **string** | TargetDigest is the blueprint target&#39;s compile pin (story 165.5, additive and omitempty): required beside a 40-hex TargetVersion, refused otherwise — a digestless SHA is never dispatchable. | [optional] 
**TargetVersion** | Pointer to **string** | TargetVersion is the catalog version the install upgrades to. The capability filter has already hidden what this engine cannot run. For a BLUEPRINT stack&#39;s update (story 165.5) it is the re-resolved commit SHA (the 40-hex shape — IsRef&#39;s discriminator) and TargetDigest carries the new compile&#39;s pin. | [optional] 

## Methods

### NewHandlersAppsUpgradeRequest

`func NewHandlersAppsUpgradeRequest() *HandlersAppsUpgradeRequest`

NewHandlersAppsUpgradeRequest instantiates a new HandlersAppsUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersAppsUpgradeRequestWithDefaults

`func NewHandlersAppsUpgradeRequestWithDefaults() *HandlersAppsUpgradeRequest`

NewHandlersAppsUpgradeRequestWithDefaults instantiates a new HandlersAppsUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzeWindowSeconds

`func (o *HandlersAppsUpgradeRequest) GetAnalyzeWindowSeconds() int32`

GetAnalyzeWindowSeconds returns the AnalyzeWindowSeconds field if non-nil, zero value otherwise.

### GetAnalyzeWindowSecondsOk

`func (o *HandlersAppsUpgradeRequest) GetAnalyzeWindowSecondsOk() (*int32, bool)`

GetAnalyzeWindowSecondsOk returns a tuple with the AnalyzeWindowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeWindowSeconds

`func (o *HandlersAppsUpgradeRequest) SetAnalyzeWindowSeconds(v int32)`

SetAnalyzeWindowSeconds sets AnalyzeWindowSeconds field to given value.

### HasAnalyzeWindowSeconds

`func (o *HandlersAppsUpgradeRequest) HasAnalyzeWindowSeconds() bool`

HasAnalyzeWindowSeconds returns a boolean if a field has been set.

### GetConfirm

`func (o *HandlersAppsUpgradeRequest) GetConfirm() string`

GetConfirm returns the Confirm field if non-nil, zero value otherwise.

### GetConfirmOk

`func (o *HandlersAppsUpgradeRequest) GetConfirmOk() (*string, bool)`

GetConfirmOk returns a tuple with the Confirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirm

`func (o *HandlersAppsUpgradeRequest) SetConfirm(v string)`

SetConfirm sets Confirm field to given value.

### HasConfirm

`func (o *HandlersAppsUpgradeRequest) HasConfirm() bool`

HasConfirm returns a boolean if a field has been set.

### GetGates

`func (o *HandlersAppsUpgradeRequest) GetGates() []EngineMetricGate`

GetGates returns the Gates field if non-nil, zero value otherwise.

### GetGatesOk

`func (o *HandlersAppsUpgradeRequest) GetGatesOk() (*[]EngineMetricGate, bool)`

GetGatesOk returns a tuple with the Gates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGates

`func (o *HandlersAppsUpgradeRequest) SetGates(v []EngineMetricGate)`

SetGates sets Gates field to given value.

### HasGates

`func (o *HandlersAppsUpgradeRequest) HasGates() bool`

HasGates returns a boolean if a field has been set.

### GetRetainOldRevisionSeconds

`func (o *HandlersAppsUpgradeRequest) GetRetainOldRevisionSeconds() int32`

GetRetainOldRevisionSeconds returns the RetainOldRevisionSeconds field if non-nil, zero value otherwise.

### GetRetainOldRevisionSecondsOk

`func (o *HandlersAppsUpgradeRequest) GetRetainOldRevisionSecondsOk() (*int32, bool)`

GetRetainOldRevisionSecondsOk returns a tuple with the RetainOldRevisionSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainOldRevisionSeconds

`func (o *HandlersAppsUpgradeRequest) SetRetainOldRevisionSeconds(v int32)`

SetRetainOldRevisionSeconds sets RetainOldRevisionSeconds field to given value.

### HasRetainOldRevisionSeconds

`func (o *HandlersAppsUpgradeRequest) HasRetainOldRevisionSeconds() bool`

HasRetainOldRevisionSeconds returns a boolean if a field has been set.

### GetSoakSeconds

`func (o *HandlersAppsUpgradeRequest) GetSoakSeconds() int32`

GetSoakSeconds returns the SoakSeconds field if non-nil, zero value otherwise.

### GetSoakSecondsOk

`func (o *HandlersAppsUpgradeRequest) GetSoakSecondsOk() (*int32, bool)`

GetSoakSecondsOk returns a tuple with the SoakSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoakSeconds

`func (o *HandlersAppsUpgradeRequest) SetSoakSeconds(v int32)`

SetSoakSeconds sets SoakSeconds field to given value.

### HasSoakSeconds

`func (o *HandlersAppsUpgradeRequest) HasSoakSeconds() bool`

HasSoakSeconds returns a boolean if a field has been set.

### GetTargetDigest

`func (o *HandlersAppsUpgradeRequest) GetTargetDigest() string`

GetTargetDigest returns the TargetDigest field if non-nil, zero value otherwise.

### GetTargetDigestOk

`func (o *HandlersAppsUpgradeRequest) GetTargetDigestOk() (*string, bool)`

GetTargetDigestOk returns a tuple with the TargetDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDigest

`func (o *HandlersAppsUpgradeRequest) SetTargetDigest(v string)`

SetTargetDigest sets TargetDigest field to given value.

### HasTargetDigest

`func (o *HandlersAppsUpgradeRequest) HasTargetDigest() bool`

HasTargetDigest returns a boolean if a field has been set.

### GetTargetVersion

`func (o *HandlersAppsUpgradeRequest) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *HandlersAppsUpgradeRequest) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *HandlersAppsUpgradeRequest) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *HandlersAppsUpgradeRequest) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


