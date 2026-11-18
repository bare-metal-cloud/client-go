# ClonereplayReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointResolutions** | Pointer to **[]string** | EndpointResolutions are the linked-external resolutions applied: the binding&#39;s endpoint name resolved against the target&#39;s registration. | [optional] 
**EnvLocalResolutions** | Pointer to **[]string** | EnvLocalResolutions are the env_local resolutions the review approved, as rendered (install.key &#x3D; value). | [optional] 
**Excluded** | Pointer to **[]string** | Excluded names the plan units the review excluded (story 134.4) — what the review left out is part of the run&#39;s truth. | [optional] 
**Failures** | Pointer to **[]string** | Failures name every failing install/component — a failed run&#39;s report is the operator&#39;s first read. | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 
**Gate** | Pointer to [**ClonereplayReportGate**](ClonereplayReportGate.md) | Gate is the post-clone gate&#39;s record (nil until the gate ran). | [optional] 
**OrderingBasis** | Pointer to **string** | OrderingBasis states the replay&#39;s ordering (the honest version — see OrderingBasis). | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Runbook** | Pointer to **string** | Runbook is the recovery runbook (RunbookText) — both recovery paths written into the report&#39;s doc. | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 
**SourceClusterId** | Pointer to **string** |  | [optional] 
**SourceProjectId** | Pointer to **string** | Source / Target name the clone&#39;s two scopes. | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | State is the report&#39;s state (the ReportState vocabulary). | [optional] 
**TargetClusterId** | Pointer to **string** |  | [optional] 
**TargetProjectId** | Pointer to **string** |  | [optional] 
**Units** | Pointer to [**[]ClonereplayReportUnit**](ClonereplayReportUnit.md) | Units are the per-install outcomes (stack components ride their own lines — the component records ARE the stack&#39;s durable truth). | [optional] 

## Methods

### NewClonereplayReport

`func NewClonereplayReport() *ClonereplayReport`

NewClonereplayReport instantiates a new ClonereplayReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClonereplayReportWithDefaults

`func NewClonereplayReportWithDefaults() *ClonereplayReport`

NewClonereplayReportWithDefaults instantiates a new ClonereplayReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointResolutions

`func (o *ClonereplayReport) GetEndpointResolutions() []string`

GetEndpointResolutions returns the EndpointResolutions field if non-nil, zero value otherwise.

### GetEndpointResolutionsOk

`func (o *ClonereplayReport) GetEndpointResolutionsOk() (*[]string, bool)`

GetEndpointResolutionsOk returns a tuple with the EndpointResolutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointResolutions

`func (o *ClonereplayReport) SetEndpointResolutions(v []string)`

SetEndpointResolutions sets EndpointResolutions field to given value.

### HasEndpointResolutions

`func (o *ClonereplayReport) HasEndpointResolutions() bool`

HasEndpointResolutions returns a boolean if a field has been set.

### GetEnvLocalResolutions

`func (o *ClonereplayReport) GetEnvLocalResolutions() []string`

GetEnvLocalResolutions returns the EnvLocalResolutions field if non-nil, zero value otherwise.

### GetEnvLocalResolutionsOk

`func (o *ClonereplayReport) GetEnvLocalResolutionsOk() (*[]string, bool)`

GetEnvLocalResolutionsOk returns a tuple with the EnvLocalResolutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvLocalResolutions

`func (o *ClonereplayReport) SetEnvLocalResolutions(v []string)`

SetEnvLocalResolutions sets EnvLocalResolutions field to given value.

### HasEnvLocalResolutions

`func (o *ClonereplayReport) HasEnvLocalResolutions() bool`

HasEnvLocalResolutions returns a boolean if a field has been set.

### GetExcluded

`func (o *ClonereplayReport) GetExcluded() []string`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *ClonereplayReport) GetExcludedOk() (*[]string, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *ClonereplayReport) SetExcluded(v []string)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *ClonereplayReport) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.

### GetFailures

`func (o *ClonereplayReport) GetFailures() []string`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *ClonereplayReport) GetFailuresOk() (*[]string, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *ClonereplayReport) SetFailures(v []string)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *ClonereplayReport) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ClonereplayReport) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ClonereplayReport) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ClonereplayReport) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ClonereplayReport) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetGate

`func (o *ClonereplayReport) GetGate() ClonereplayReportGate`

GetGate returns the Gate field if non-nil, zero value otherwise.

### GetGateOk

`func (o *ClonereplayReport) GetGateOk() (*ClonereplayReportGate, bool)`

GetGateOk returns a tuple with the Gate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGate

`func (o *ClonereplayReport) SetGate(v ClonereplayReportGate)`

SetGate sets Gate field to given value.

### HasGate

`func (o *ClonereplayReport) HasGate() bool`

HasGate returns a boolean if a field has been set.

### GetOrderingBasis

`func (o *ClonereplayReport) GetOrderingBasis() string`

GetOrderingBasis returns the OrderingBasis field if non-nil, zero value otherwise.

### GetOrderingBasisOk

`func (o *ClonereplayReport) GetOrderingBasisOk() (*string, bool)`

GetOrderingBasisOk returns a tuple with the OrderingBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderingBasis

`func (o *ClonereplayReport) SetOrderingBasis(v string)`

SetOrderingBasis sets OrderingBasis field to given value.

### HasOrderingBasis

`func (o *ClonereplayReport) HasOrderingBasis() bool`

HasOrderingBasis returns a boolean if a field has been set.

### GetPlanId

`func (o *ClonereplayReport) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ClonereplayReport) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ClonereplayReport) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *ClonereplayReport) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetRunbook

`func (o *ClonereplayReport) GetRunbook() string`

GetRunbook returns the Runbook field if non-nil, zero value otherwise.

### GetRunbookOk

`func (o *ClonereplayReport) GetRunbookOk() (*string, bool)`

GetRunbookOk returns a tuple with the Runbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunbook

`func (o *ClonereplayReport) SetRunbook(v string)`

SetRunbook sets Runbook field to given value.

### HasRunbook

`func (o *ClonereplayReport) HasRunbook() bool`

HasRunbook returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ClonereplayReport) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ClonereplayReport) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ClonereplayReport) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ClonereplayReport) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSourceClusterId

`func (o *ClonereplayReport) GetSourceClusterId() string`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### GetSourceClusterIdOk

`func (o *ClonereplayReport) GetSourceClusterIdOk() (*string, bool)`

GetSourceClusterIdOk returns a tuple with the SourceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterId

`func (o *ClonereplayReport) SetSourceClusterId(v string)`

SetSourceClusterId sets SourceClusterId field to given value.

### HasSourceClusterId

`func (o *ClonereplayReport) HasSourceClusterId() bool`

HasSourceClusterId returns a boolean if a field has been set.

### GetSourceProjectId

`func (o *ClonereplayReport) GetSourceProjectId() string`

GetSourceProjectId returns the SourceProjectId field if non-nil, zero value otherwise.

### GetSourceProjectIdOk

`func (o *ClonereplayReport) GetSourceProjectIdOk() (*string, bool)`

GetSourceProjectIdOk returns a tuple with the SourceProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProjectId

`func (o *ClonereplayReport) SetSourceProjectId(v string)`

SetSourceProjectId sets SourceProjectId field to given value.

### HasSourceProjectId

`func (o *ClonereplayReport) HasSourceProjectId() bool`

HasSourceProjectId returns a boolean if a field has been set.

### GetStartedAt

`func (o *ClonereplayReport) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ClonereplayReport) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ClonereplayReport) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ClonereplayReport) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetState

`func (o *ClonereplayReport) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClonereplayReport) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClonereplayReport) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ClonereplayReport) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTargetClusterId

`func (o *ClonereplayReport) GetTargetClusterId() string`

GetTargetClusterId returns the TargetClusterId field if non-nil, zero value otherwise.

### GetTargetClusterIdOk

`func (o *ClonereplayReport) GetTargetClusterIdOk() (*string, bool)`

GetTargetClusterIdOk returns a tuple with the TargetClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterId

`func (o *ClonereplayReport) SetTargetClusterId(v string)`

SetTargetClusterId sets TargetClusterId field to given value.

### HasTargetClusterId

`func (o *ClonereplayReport) HasTargetClusterId() bool`

HasTargetClusterId returns a boolean if a field has been set.

### GetTargetProjectId

`func (o *ClonereplayReport) GetTargetProjectId() string`

GetTargetProjectId returns the TargetProjectId field if non-nil, zero value otherwise.

### GetTargetProjectIdOk

`func (o *ClonereplayReport) GetTargetProjectIdOk() (*string, bool)`

GetTargetProjectIdOk returns a tuple with the TargetProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProjectId

`func (o *ClonereplayReport) SetTargetProjectId(v string)`

SetTargetProjectId sets TargetProjectId field to given value.

### HasTargetProjectId

`func (o *ClonereplayReport) HasTargetProjectId() bool`

HasTargetProjectId returns a boolean if a field has been set.

### GetUnits

`func (o *ClonereplayReport) GetUnits() []ClonereplayReportUnit`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ClonereplayReport) GetUnitsOk() (*[]ClonereplayReportUnit, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ClonereplayReport) SetUnits(v []ClonereplayReportUnit)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ClonereplayReport) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


