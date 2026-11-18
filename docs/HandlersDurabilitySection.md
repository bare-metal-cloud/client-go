# HandlersDurabilitySection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badge** | Pointer to [**HandlersDurabilityBadge**](HandlersDurabilityBadge.md) |  | [optional] 
**Components** | Pointer to [**[]HandlersDurabilitySection**](HandlersDurabilitySection.md) |  | [optional] 
**Install** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Points** | Pointer to [**[]HandlersDurabilityPoint**](HandlersDurabilityPoint.md) |  | [optional] 
**RetentionPreset** | Pointer to **string** |  | [optional] 
**Runs** | Pointer to [**[]HandlersDurabilityRun**](HandlersDurabilityRun.md) |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Streaks** | Pointer to [**HandlersDurabilityStreaks**](HandlersDurabilityStreaks.md) |  | [optional] 
**TargetDestination** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersDurabilitySection

`func NewHandlersDurabilitySection() *HandlersDurabilitySection`

NewHandlersDurabilitySection instantiates a new HandlersDurabilitySection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDurabilitySectionWithDefaults

`func NewHandlersDurabilitySectionWithDefaults() *HandlersDurabilitySection`

NewHandlersDurabilitySectionWithDefaults instantiates a new HandlersDurabilitySection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadge

`func (o *HandlersDurabilitySection) GetBadge() HandlersDurabilityBadge`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *HandlersDurabilitySection) GetBadgeOk() (*HandlersDurabilityBadge, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *HandlersDurabilitySection) SetBadge(v HandlersDurabilityBadge)`

SetBadge sets Badge field to given value.

### HasBadge

`func (o *HandlersDurabilitySection) HasBadge() bool`

HasBadge returns a boolean if a field has been set.

### GetComponents

`func (o *HandlersDurabilitySection) GetComponents() []HandlersDurabilitySection`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *HandlersDurabilitySection) GetComponentsOk() (*[]HandlersDurabilitySection, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *HandlersDurabilitySection) SetComponents(v []HandlersDurabilitySection)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *HandlersDurabilitySection) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetInstall

`func (o *HandlersDurabilitySection) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *HandlersDurabilitySection) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *HandlersDurabilitySection) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *HandlersDurabilitySection) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetMode

`func (o *HandlersDurabilitySection) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HandlersDurabilitySection) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HandlersDurabilitySection) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HandlersDurabilitySection) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPoints

`func (o *HandlersDurabilitySection) GetPoints() []HandlersDurabilityPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *HandlersDurabilitySection) GetPointsOk() (*[]HandlersDurabilityPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *HandlersDurabilitySection) SetPoints(v []HandlersDurabilityPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *HandlersDurabilitySection) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetRetentionPreset

`func (o *HandlersDurabilitySection) GetRetentionPreset() string`

GetRetentionPreset returns the RetentionPreset field if non-nil, zero value otherwise.

### GetRetentionPresetOk

`func (o *HandlersDurabilitySection) GetRetentionPresetOk() (*string, bool)`

GetRetentionPresetOk returns a tuple with the RetentionPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPreset

`func (o *HandlersDurabilitySection) SetRetentionPreset(v string)`

SetRetentionPreset sets RetentionPreset field to given value.

### HasRetentionPreset

`func (o *HandlersDurabilitySection) HasRetentionPreset() bool`

HasRetentionPreset returns a boolean if a field has been set.

### GetRuns

`func (o *HandlersDurabilitySection) GetRuns() []HandlersDurabilityRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *HandlersDurabilitySection) GetRunsOk() (*[]HandlersDurabilityRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *HandlersDurabilitySection) SetRuns(v []HandlersDurabilityRun)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *HandlersDurabilitySection) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetSchedule

`func (o *HandlersDurabilitySection) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *HandlersDurabilitySection) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *HandlersDurabilitySection) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *HandlersDurabilitySection) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersDurabilitySection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersDurabilitySection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersDurabilitySection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersDurabilitySection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStreaks

`func (o *HandlersDurabilitySection) GetStreaks() HandlersDurabilityStreaks`

GetStreaks returns the Streaks field if non-nil, zero value otherwise.

### GetStreaksOk

`func (o *HandlersDurabilitySection) GetStreaksOk() (*HandlersDurabilityStreaks, bool)`

GetStreaksOk returns a tuple with the Streaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreaks

`func (o *HandlersDurabilitySection) SetStreaks(v HandlersDurabilityStreaks)`

SetStreaks sets Streaks field to given value.

### HasStreaks

`func (o *HandlersDurabilitySection) HasStreaks() bool`

HasStreaks returns a boolean if a field has been set.

### GetTargetDestination

`func (o *HandlersDurabilitySection) GetTargetDestination() string`

GetTargetDestination returns the TargetDestination field if non-nil, zero value otherwise.

### GetTargetDestinationOk

`func (o *HandlersDurabilitySection) GetTargetDestinationOk() (*string, bool)`

GetTargetDestinationOk returns a tuple with the TargetDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDestination

`func (o *HandlersDurabilitySection) SetTargetDestination(v string)`

SetTargetDestination sets TargetDestination field to given value.

### HasTargetDestination

`func (o *HandlersDurabilitySection) HasTargetDestination() bool`

HasTargetDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


