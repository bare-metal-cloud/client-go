# HandlersDeploykitInstallStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildBinding** | Pointer to [**HandlersDeploykitBuildBinding**](HandlersDeploykitBuildBinding.md) | BuildBinding is the project&#39;s git build (BAR-601, additive): the trigger, repo and branch the stack deploys from, or null. Always served, like the list entry. | [optional] 
**Bundle** | Pointer to [**BundleRef**](BundleRef.md) | Bundle is the deployed revision&#39;s bundle ref (story 142.5, additive and omitempty): the installed identity the upgrade affordance needs — the capability-filtered versions listing is queried with its name and the running version is named beside the affordance and cut from the targets. Absent on a record with no deployed revision. | [optional] 
**Components** | Pointer to [**[]HandlersDeploykitComponentStatus**](HandlersDeploykitComponentStatus.md) |  | [optional] 
**CurrentRevision** | Pointer to **int32** |  | [optional] 
**Definitions** | Pointer to [**HandlersDeploykitDefinitionStatus**](HandlersDeploykitDefinitionStatus.md) | Definitions is the stack&#39;s platform-held definition record (epic stack-evolution entry 9, additive and omitempty): the latest revision&#39;s pin and the source lifecycle&#39;s NAMED state — \&quot;orphaned\&quot; when the import repository no longer answers (D4). The stack keeps running and evolving from stored revisions while orphaned; the named state is what the panel renders. Absent for installs with no definition revisions (single-component installs and every pre-entry-9 stack). | [optional] 
**Detail** | Pointer to [**HandlersDeploykitInstallStatusDetail**](HandlersDeploykitInstallStatusDetail.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]EngineportsInstallProgressEvent**](EngineportsInstallProgressEvent.md) |  | [optional] 
**GatePause** | Pointer to [**HandlersDeploykitGatePauseStatus**](HandlersDeploykitGatePauseStatus.md) | GatePause is the rule-10 gate pause block (story 137.2, issue #2962, additive and omitempty): when the newest gate progress event is an OPEN pause, the panel renders the Resume/Abort affordance off it. Absent on every install whose events carry no open pause. | [optional] 
**Generation** | Pointer to **int32** |  | [optional] 
**Gitops** | Pointer to [**GitopsInstallStateSection**](GitopsInstallStateSection.md) | Gitops is the install&#39;s CI-deploy state section (story 143.2, additive and omitempty): enabled/reconnect/pending/installed/ abandoned/disabled with the connection, repo coordinates, and the last CI deploy. Absent when the install carries no gitops surface. | [optional] 
**Install** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to [**[]EngineLabel**](EngineLabel.md) | Labels is the record&#39;s movable label pointers (story 149.6, additive and omitempty): name → revision (\&quot;next\&quot;, \&quot;staging\&quot; — the deploykit label law), served verbatim so the deploy timeline&#39;s revision cards can join the labels by revision number. A record with no labels serves no key at all (the pre-149.6 shape). | [optional] 
**Revisions** | Pointer to [**[]HandlersDeploykitRevisionStatus**](HandlersDeploykitRevisionStatus.md) | Revisions is the per-revision provenance surface (story 143.3, additive and omitempty): every revision&#39;s bundle identity and provenance, with the CI build subject beside the revisions that carry one. A nil Build is the MANUAL subject — never inherited (Law 12b): a panel deploy after a CI deploy renders the manual subject on its own revision while the CI revision keeps its own. The panel&#39;s rollback affordance reads the CI-stamped numbers off this list. | [optional] 
**Slots** | Pointer to [**[]HandlersDeploykitSlotStatus**](HandlersDeploykitSlotStatus.md) | Slots is the stack view&#39;s slot panel (epic managed-xyz story 5b, issue #4085, additive and omitempty): the stack&#39;s managed slots with the SERVER-AUTHORED action list — the rotate action renders from here, never invented client-side. Absent on a stack that declares no slots. | [optional] 
**StableRevision** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersDeploykitInstallStatusResponse

`func NewHandlersDeploykitInstallStatusResponse() *HandlersDeploykitInstallStatusResponse`

NewHandlersDeploykitInstallStatusResponse instantiates a new HandlersDeploykitInstallStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDeploykitInstallStatusResponseWithDefaults

`func NewHandlersDeploykitInstallStatusResponseWithDefaults() *HandlersDeploykitInstallStatusResponse`

NewHandlersDeploykitInstallStatusResponseWithDefaults instantiates a new HandlersDeploykitInstallStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildBinding

`func (o *HandlersDeploykitInstallStatusResponse) GetBuildBinding() HandlersDeploykitBuildBinding`

GetBuildBinding returns the BuildBinding field if non-nil, zero value otherwise.

### GetBuildBindingOk

`func (o *HandlersDeploykitInstallStatusResponse) GetBuildBindingOk() (*HandlersDeploykitBuildBinding, bool)`

GetBuildBindingOk returns a tuple with the BuildBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildBinding

`func (o *HandlersDeploykitInstallStatusResponse) SetBuildBinding(v HandlersDeploykitBuildBinding)`

SetBuildBinding sets BuildBinding field to given value.

### HasBuildBinding

`func (o *HandlersDeploykitInstallStatusResponse) HasBuildBinding() bool`

HasBuildBinding returns a boolean if a field has been set.

### GetBundle

`func (o *HandlersDeploykitInstallStatusResponse) GetBundle() BundleRef`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *HandlersDeploykitInstallStatusResponse) GetBundleOk() (*BundleRef, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *HandlersDeploykitInstallStatusResponse) SetBundle(v BundleRef)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *HandlersDeploykitInstallStatusResponse) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetComponents

`func (o *HandlersDeploykitInstallStatusResponse) GetComponents() []HandlersDeploykitComponentStatus`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *HandlersDeploykitInstallStatusResponse) GetComponentsOk() (*[]HandlersDeploykitComponentStatus, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *HandlersDeploykitInstallStatusResponse) SetComponents(v []HandlersDeploykitComponentStatus)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *HandlersDeploykitInstallStatusResponse) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCurrentRevision

`func (o *HandlersDeploykitInstallStatusResponse) GetCurrentRevision() int32`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *HandlersDeploykitInstallStatusResponse) GetCurrentRevisionOk() (*int32, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *HandlersDeploykitInstallStatusResponse) SetCurrentRevision(v int32)`

SetCurrentRevision sets CurrentRevision field to given value.

### HasCurrentRevision

`func (o *HandlersDeploykitInstallStatusResponse) HasCurrentRevision() bool`

HasCurrentRevision returns a boolean if a field has been set.

### GetDefinitions

`func (o *HandlersDeploykitInstallStatusResponse) GetDefinitions() HandlersDeploykitDefinitionStatus`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *HandlersDeploykitInstallStatusResponse) GetDefinitionsOk() (*HandlersDeploykitDefinitionStatus, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *HandlersDeploykitInstallStatusResponse) SetDefinitions(v HandlersDeploykitDefinitionStatus)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *HandlersDeploykitInstallStatusResponse) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetDetail

`func (o *HandlersDeploykitInstallStatusResponse) GetDetail() HandlersDeploykitInstallStatusDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *HandlersDeploykitInstallStatusResponse) GetDetailOk() (*HandlersDeploykitInstallStatusDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *HandlersDeploykitInstallStatusResponse) SetDetail(v HandlersDeploykitInstallStatusDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *HandlersDeploykitInstallStatusResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDisplayName

`func (o *HandlersDeploykitInstallStatusResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HandlersDeploykitInstallStatusResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HandlersDeploykitInstallStatusResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HandlersDeploykitInstallStatusResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEvents

`func (o *HandlersDeploykitInstallStatusResponse) GetEvents() []EngineportsInstallProgressEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *HandlersDeploykitInstallStatusResponse) GetEventsOk() (*[]EngineportsInstallProgressEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *HandlersDeploykitInstallStatusResponse) SetEvents(v []EngineportsInstallProgressEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *HandlersDeploykitInstallStatusResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetGatePause

`func (o *HandlersDeploykitInstallStatusResponse) GetGatePause() HandlersDeploykitGatePauseStatus`

GetGatePause returns the GatePause field if non-nil, zero value otherwise.

### GetGatePauseOk

`func (o *HandlersDeploykitInstallStatusResponse) GetGatePauseOk() (*HandlersDeploykitGatePauseStatus, bool)`

GetGatePauseOk returns a tuple with the GatePause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatePause

`func (o *HandlersDeploykitInstallStatusResponse) SetGatePause(v HandlersDeploykitGatePauseStatus)`

SetGatePause sets GatePause field to given value.

### HasGatePause

`func (o *HandlersDeploykitInstallStatusResponse) HasGatePause() bool`

HasGatePause returns a boolean if a field has been set.

### GetGeneration

`func (o *HandlersDeploykitInstallStatusResponse) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *HandlersDeploykitInstallStatusResponse) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *HandlersDeploykitInstallStatusResponse) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *HandlersDeploykitInstallStatusResponse) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetGitops

`func (o *HandlersDeploykitInstallStatusResponse) GetGitops() GitopsInstallStateSection`

GetGitops returns the Gitops field if non-nil, zero value otherwise.

### GetGitopsOk

`func (o *HandlersDeploykitInstallStatusResponse) GetGitopsOk() (*GitopsInstallStateSection, bool)`

GetGitopsOk returns a tuple with the Gitops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitops

`func (o *HandlersDeploykitInstallStatusResponse) SetGitops(v GitopsInstallStateSection)`

SetGitops sets Gitops field to given value.

### HasGitops

`func (o *HandlersDeploykitInstallStatusResponse) HasGitops() bool`

HasGitops returns a boolean if a field has been set.

### GetInstall

`func (o *HandlersDeploykitInstallStatusResponse) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *HandlersDeploykitInstallStatusResponse) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *HandlersDeploykitInstallStatusResponse) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *HandlersDeploykitInstallStatusResponse) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetLabels

`func (o *HandlersDeploykitInstallStatusResponse) GetLabels() []EngineLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HandlersDeploykitInstallStatusResponse) GetLabelsOk() (*[]EngineLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HandlersDeploykitInstallStatusResponse) SetLabels(v []EngineLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HandlersDeploykitInstallStatusResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetRevisions

`func (o *HandlersDeploykitInstallStatusResponse) GetRevisions() []HandlersDeploykitRevisionStatus`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *HandlersDeploykitInstallStatusResponse) GetRevisionsOk() (*[]HandlersDeploykitRevisionStatus, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *HandlersDeploykitInstallStatusResponse) SetRevisions(v []HandlersDeploykitRevisionStatus)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *HandlersDeploykitInstallStatusResponse) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetSlots

`func (o *HandlersDeploykitInstallStatusResponse) GetSlots() []HandlersDeploykitSlotStatus`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *HandlersDeploykitInstallStatusResponse) GetSlotsOk() (*[]HandlersDeploykitSlotStatus, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *HandlersDeploykitInstallStatusResponse) SetSlots(v []HandlersDeploykitSlotStatus)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *HandlersDeploykitInstallStatusResponse) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetStableRevision

`func (o *HandlersDeploykitInstallStatusResponse) GetStableRevision() int32`

GetStableRevision returns the StableRevision field if non-nil, zero value otherwise.

### GetStableRevisionOk

`func (o *HandlersDeploykitInstallStatusResponse) GetStableRevisionOk() (*int32, bool)`

GetStableRevisionOk returns a tuple with the StableRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableRevision

`func (o *HandlersDeploykitInstallStatusResponse) SetStableRevision(v int32)`

SetStableRevision sets StableRevision field to given value.

### HasStableRevision

`func (o *HandlersDeploykitInstallStatusResponse) HasStableRevision() bool`

HasStableRevision returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersDeploykitInstallStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersDeploykitInstallStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersDeploykitInstallStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersDeploykitInstallStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


