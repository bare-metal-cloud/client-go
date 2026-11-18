# CloneplanPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt / UpdatedAt are the derivation stamps (store-owned on persistence). | [optional] 
**Endpoints** | Pointer to [**[]CloneplanEndpointRef**](CloneplanEndpointRef.md) | Endpoints are the source scope&#39;s linked-external endpoint registrations as target re-resolution items: name, type, host, and port ONLY (AC 4). The registry&#39;s sealed credentials have no field to ride — the store&#39;s reference read never selects them. | [optional] 
**Id** | Pointer to **string** | ID is the plan&#39;s stable id (names-free; minted at derivation when the request did not pin one). | [optional] 
**SchemaVersion** | Pointer to **int32** | SchemaVersion is the document&#39;s schema version (PlanSchemaV1). | [optional] 
**Source** | Pointer to [**EngineInstallRef**](EngineInstallRef.md) | Source / Target name the clone&#39;s two scopes. The InstallRef&#39;s ProjectID + ClusterID are the scope key (the engine identity); InstallName is empty on both — the plan covers the whole scope. | [optional] 
**SourceStateDigest** | Pointer to **string** | SourceStateDigest pins the derivation&#39;s inputs (the doc comment&#39;s one sentence on the package): the source records&#39; generations, the StackDAGStore captures, and the external-endpoint registrations. CheckFresh refuses execution against a source whose digest changed. | [optional] 
**State** | Pointer to [**CloneplanPlanState**](CloneplanPlanState.md) | State is the derivation outcome (PlanState). | [optional] 
**Target** | Pointer to [**EngineInstallRef**](EngineInstallRef.md) |  | [optional] 
**Units** | Pointer to [**[]CloneplanUnit**](CloneplanUnit.md) | Units are the plan&#39;s install units in source install-name order — one per loose install, one per stack-shaped install (whose components fold INTO the stack unit and never appear as loose units, AC 9). | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCloneplanPlan

`func NewCloneplanPlan() *CloneplanPlan`

NewCloneplanPlan instantiates a new CloneplanPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanPlanWithDefaults

`func NewCloneplanPlanWithDefaults() *CloneplanPlan`

NewCloneplanPlanWithDefaults instantiates a new CloneplanPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloneplanPlan) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloneplanPlan) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloneplanPlan) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloneplanPlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEndpoints

`func (o *CloneplanPlan) GetEndpoints() []CloneplanEndpointRef`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *CloneplanPlan) GetEndpointsOk() (*[]CloneplanEndpointRef, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *CloneplanPlan) SetEndpoints(v []CloneplanEndpointRef)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *CloneplanPlan) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetId

`func (o *CloneplanPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloneplanPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloneplanPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloneplanPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *CloneplanPlan) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CloneplanPlan) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CloneplanPlan) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CloneplanPlan) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSource

`func (o *CloneplanPlan) GetSource() EngineInstallRef`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloneplanPlan) GetSourceOk() (*EngineInstallRef, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloneplanPlan) SetSource(v EngineInstallRef)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloneplanPlan) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceStateDigest

`func (o *CloneplanPlan) GetSourceStateDigest() string`

GetSourceStateDigest returns the SourceStateDigest field if non-nil, zero value otherwise.

### GetSourceStateDigestOk

`func (o *CloneplanPlan) GetSourceStateDigestOk() (*string, bool)`

GetSourceStateDigestOk returns a tuple with the SourceStateDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStateDigest

`func (o *CloneplanPlan) SetSourceStateDigest(v string)`

SetSourceStateDigest sets SourceStateDigest field to given value.

### HasSourceStateDigest

`func (o *CloneplanPlan) HasSourceStateDigest() bool`

HasSourceStateDigest returns a boolean if a field has been set.

### GetState

`func (o *CloneplanPlan) GetState() CloneplanPlanState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloneplanPlan) GetStateOk() (*CloneplanPlanState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloneplanPlan) SetState(v CloneplanPlanState)`

SetState sets State field to given value.

### HasState

`func (o *CloneplanPlan) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTarget

`func (o *CloneplanPlan) GetTarget() EngineInstallRef`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloneplanPlan) GetTargetOk() (*EngineInstallRef, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloneplanPlan) SetTarget(v EngineInstallRef)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloneplanPlan) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUnits

`func (o *CloneplanPlan) GetUnits() []CloneplanUnit`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CloneplanPlan) GetUnitsOk() (*[]CloneplanUnit, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CloneplanPlan) SetUnits(v []CloneplanUnit)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *CloneplanPlan) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloneplanPlan) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloneplanPlan) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloneplanPlan) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloneplanPlan) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


