# EngineportsStrandRetargetPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candidates** | Pointer to [**[]EngineportsStrandRecordDiff**](EngineportsStrandRecordDiff.md) | Candidates are the attributable records (outcome retarget). | [optional] 
**ClusterId** | Pointer to **string** | ClusterID is the triple&#39;s cluster ref (the engine&#39;s numeric cluster row id in string form). | [optional] 
**DeadTargets** | Pointer to [**[]EngineportsStrandDeadTarget**](EngineportsStrandDeadTarget.md) | DeadTargets lists every remembered dead target (the attribution set). | [optional] 
**Instance** | Pointer to **string** | Instance is the arrival addon instance id. | [optional] 
**LiveTarget** | Pointer to **string** | LiveTarget is the live tunnel&#39;s full CNAME — the re-target destination every candidate would move to. | [optional] 
**Untouched** | Pointer to [**[]EngineportsStrandRecordDiff**](EngineportsStrandRecordDiff.md) | Untouched are the records the preview names but will never write. | [optional] 
**Zones** | Pointer to [**[]EngineportsRecordCleanupOutcome**](EngineportsRecordCleanupOutcome.md) | Zones carries the zone-level outcomes (no-custody / failed) for zones whose records could not be enumerated. | [optional] 

## Methods

### NewEngineportsStrandRetargetPreview

`func NewEngineportsStrandRetargetPreview() *EngineportsStrandRetargetPreview`

NewEngineportsStrandRetargetPreview instantiates a new EngineportsStrandRetargetPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineportsStrandRetargetPreviewWithDefaults

`func NewEngineportsStrandRetargetPreviewWithDefaults() *EngineportsStrandRetargetPreview`

NewEngineportsStrandRetargetPreviewWithDefaults instantiates a new EngineportsStrandRetargetPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidates

`func (o *EngineportsStrandRetargetPreview) GetCandidates() []EngineportsStrandRecordDiff`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *EngineportsStrandRetargetPreview) GetCandidatesOk() (*[]EngineportsStrandRecordDiff, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *EngineportsStrandRetargetPreview) SetCandidates(v []EngineportsStrandRecordDiff)`

SetCandidates sets Candidates field to given value.

### HasCandidates

`func (o *EngineportsStrandRetargetPreview) HasCandidates() bool`

HasCandidates returns a boolean if a field has been set.

### GetClusterId

`func (o *EngineportsStrandRetargetPreview) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EngineportsStrandRetargetPreview) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EngineportsStrandRetargetPreview) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EngineportsStrandRetargetPreview) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetDeadTargets

`func (o *EngineportsStrandRetargetPreview) GetDeadTargets() []EngineportsStrandDeadTarget`

GetDeadTargets returns the DeadTargets field if non-nil, zero value otherwise.

### GetDeadTargetsOk

`func (o *EngineportsStrandRetargetPreview) GetDeadTargetsOk() (*[]EngineportsStrandDeadTarget, bool)`

GetDeadTargetsOk returns a tuple with the DeadTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadTargets

`func (o *EngineportsStrandRetargetPreview) SetDeadTargets(v []EngineportsStrandDeadTarget)`

SetDeadTargets sets DeadTargets field to given value.

### HasDeadTargets

`func (o *EngineportsStrandRetargetPreview) HasDeadTargets() bool`

HasDeadTargets returns a boolean if a field has been set.

### GetInstance

`func (o *EngineportsStrandRetargetPreview) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *EngineportsStrandRetargetPreview) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *EngineportsStrandRetargetPreview) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *EngineportsStrandRetargetPreview) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetLiveTarget

`func (o *EngineportsStrandRetargetPreview) GetLiveTarget() string`

GetLiveTarget returns the LiveTarget field if non-nil, zero value otherwise.

### GetLiveTargetOk

`func (o *EngineportsStrandRetargetPreview) GetLiveTargetOk() (*string, bool)`

GetLiveTargetOk returns a tuple with the LiveTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveTarget

`func (o *EngineportsStrandRetargetPreview) SetLiveTarget(v string)`

SetLiveTarget sets LiveTarget field to given value.

### HasLiveTarget

`func (o *EngineportsStrandRetargetPreview) HasLiveTarget() bool`

HasLiveTarget returns a boolean if a field has been set.

### GetUntouched

`func (o *EngineportsStrandRetargetPreview) GetUntouched() []EngineportsStrandRecordDiff`

GetUntouched returns the Untouched field if non-nil, zero value otherwise.

### GetUntouchedOk

`func (o *EngineportsStrandRetargetPreview) GetUntouchedOk() (*[]EngineportsStrandRecordDiff, bool)`

GetUntouchedOk returns a tuple with the Untouched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntouched

`func (o *EngineportsStrandRetargetPreview) SetUntouched(v []EngineportsStrandRecordDiff)`

SetUntouched sets Untouched field to given value.

### HasUntouched

`func (o *EngineportsStrandRetargetPreview) HasUntouched() bool`

HasUntouched returns a boolean if a field has been set.

### GetZones

`func (o *EngineportsStrandRetargetPreview) GetZones() []EngineportsRecordCleanupOutcome`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *EngineportsStrandRetargetPreview) GetZonesOk() (*[]EngineportsRecordCleanupOutcome, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *EngineportsStrandRetargetPreview) SetZones(v []EngineportsRecordCleanupOutcome)`

SetZones sets Zones field to given value.

### HasZones

`func (o *EngineportsStrandRetargetPreview) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


