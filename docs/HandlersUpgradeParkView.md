# HandlersUpgradeParkView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **int32** |  | [optional] 
**CommitSha** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**ImageDigest** | Pointer to **string** |  | [optional] 
**Install** | Pointer to **string** |  | [optional] 
**MigrationSummary** | Pointer to **string** |  | [optional] 
**PanelUrl** | Pointer to **string** |  | [optional] 
**ParkId** | Pointer to **int32** |  | [optional] 
**PlanFingerprint** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**RunUrl** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Verdicts** | Pointer to [**[]GitopsVerdictComponent**](GitopsVerdictComponent.md) |  | [optional] 
**VolumeDelta** | Pointer to [**GitopsVolumeDelta**](GitopsVolumeDelta.md) |  | [optional] 

## Methods

### NewHandlersUpgradeParkView

`func NewHandlersUpgradeParkView() *HandlersUpgradeParkView`

NewHandlersUpgradeParkView instantiates a new HandlersUpgradeParkView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersUpgradeParkViewWithDefaults

`func NewHandlersUpgradeParkViewWithDefaults() *HandlersUpgradeParkView`

NewHandlersUpgradeParkViewWithDefaults instantiates a new HandlersUpgradeParkView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *HandlersUpgradeParkView) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HandlersUpgradeParkView) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HandlersUpgradeParkView) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HandlersUpgradeParkView) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCommitSha

`func (o *HandlersUpgradeParkView) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *HandlersUpgradeParkView) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *HandlersUpgradeParkView) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *HandlersUpgradeParkView) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HandlersUpgradeParkView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlersUpgradeParkView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlersUpgradeParkView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HandlersUpgradeParkView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetImageDigest

`func (o *HandlersUpgradeParkView) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *HandlersUpgradeParkView) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *HandlersUpgradeParkView) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *HandlersUpgradeParkView) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.

### GetInstall

`func (o *HandlersUpgradeParkView) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *HandlersUpgradeParkView) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *HandlersUpgradeParkView) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *HandlersUpgradeParkView) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetMigrationSummary

`func (o *HandlersUpgradeParkView) GetMigrationSummary() string`

GetMigrationSummary returns the MigrationSummary field if non-nil, zero value otherwise.

### GetMigrationSummaryOk

`func (o *HandlersUpgradeParkView) GetMigrationSummaryOk() (*string, bool)`

GetMigrationSummaryOk returns a tuple with the MigrationSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationSummary

`func (o *HandlersUpgradeParkView) SetMigrationSummary(v string)`

SetMigrationSummary sets MigrationSummary field to given value.

### HasMigrationSummary

`func (o *HandlersUpgradeParkView) HasMigrationSummary() bool`

HasMigrationSummary returns a boolean if a field has been set.

### GetPanelUrl

`func (o *HandlersUpgradeParkView) GetPanelUrl() string`

GetPanelUrl returns the PanelUrl field if non-nil, zero value otherwise.

### GetPanelUrlOk

`func (o *HandlersUpgradeParkView) GetPanelUrlOk() (*string, bool)`

GetPanelUrlOk returns a tuple with the PanelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelUrl

`func (o *HandlersUpgradeParkView) SetPanelUrl(v string)`

SetPanelUrl sets PanelUrl field to given value.

### HasPanelUrl

`func (o *HandlersUpgradeParkView) HasPanelUrl() bool`

HasPanelUrl returns a boolean if a field has been set.

### GetParkId

`func (o *HandlersUpgradeParkView) GetParkId() int32`

GetParkId returns the ParkId field if non-nil, zero value otherwise.

### GetParkIdOk

`func (o *HandlersUpgradeParkView) GetParkIdOk() (*int32, bool)`

GetParkIdOk returns a tuple with the ParkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkId

`func (o *HandlersUpgradeParkView) SetParkId(v int32)`

SetParkId sets ParkId field to given value.

### HasParkId

`func (o *HandlersUpgradeParkView) HasParkId() bool`

HasParkId returns a boolean if a field has been set.

### GetPlanFingerprint

`func (o *HandlersUpgradeParkView) GetPlanFingerprint() string`

GetPlanFingerprint returns the PlanFingerprint field if non-nil, zero value otherwise.

### GetPlanFingerprintOk

`func (o *HandlersUpgradeParkView) GetPlanFingerprintOk() (*string, bool)`

GetPlanFingerprintOk returns a tuple with the PlanFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanFingerprint

`func (o *HandlersUpgradeParkView) SetPlanFingerprint(v string)`

SetPlanFingerprint sets PlanFingerprint field to given value.

### HasPlanFingerprint

`func (o *HandlersUpgradeParkView) HasPlanFingerprint() bool`

HasPlanFingerprint returns a boolean if a field has been set.

### GetRepo

`func (o *HandlersUpgradeParkView) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *HandlersUpgradeParkView) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *HandlersUpgradeParkView) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *HandlersUpgradeParkView) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRunUrl

`func (o *HandlersUpgradeParkView) GetRunUrl() string`

GetRunUrl returns the RunUrl field if non-nil, zero value otherwise.

### GetRunUrlOk

`func (o *HandlersUpgradeParkView) GetRunUrlOk() (*string, bool)`

GetRunUrlOk returns a tuple with the RunUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunUrl

`func (o *HandlersUpgradeParkView) SetRunUrl(v string)`

SetRunUrl sets RunUrl field to given value.

### HasRunUrl

`func (o *HandlersUpgradeParkView) HasRunUrl() bool`

HasRunUrl returns a boolean if a field has been set.

### GetState

`func (o *HandlersUpgradeParkView) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersUpgradeParkView) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersUpgradeParkView) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersUpgradeParkView) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVerdicts

`func (o *HandlersUpgradeParkView) GetVerdicts() []GitopsVerdictComponent`

GetVerdicts returns the Verdicts field if non-nil, zero value otherwise.

### GetVerdictsOk

`func (o *HandlersUpgradeParkView) GetVerdictsOk() (*[]GitopsVerdictComponent, bool)`

GetVerdictsOk returns a tuple with the Verdicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdicts

`func (o *HandlersUpgradeParkView) SetVerdicts(v []GitopsVerdictComponent)`

SetVerdicts sets Verdicts field to given value.

### HasVerdicts

`func (o *HandlersUpgradeParkView) HasVerdicts() bool`

HasVerdicts returns a boolean if a field has been set.

### GetVolumeDelta

`func (o *HandlersUpgradeParkView) GetVolumeDelta() GitopsVolumeDelta`

GetVolumeDelta returns the VolumeDelta field if non-nil, zero value otherwise.

### GetVolumeDeltaOk

`func (o *HandlersUpgradeParkView) GetVolumeDeltaOk() (*GitopsVolumeDelta, bool)`

GetVolumeDeltaOk returns a tuple with the VolumeDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDelta

`func (o *HandlersUpgradeParkView) SetVolumeDelta(v GitopsVolumeDelta)`

SetVolumeDelta sets VolumeDelta field to given value.

### HasVolumeDelta

`func (o *HandlersUpgradeParkView) HasVolumeDelta() bool`

HasVolumeDelta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


