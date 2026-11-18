# HandlersDeploykitInstallSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildBinding** | Pointer to [**HandlersDeploykitBuildBinding**](HandlersDeploykitBuildBinding.md) | BuildBinding is the project&#39;s git build (BAR-601): the trigger, repo and branch the stack deploys from, or null when the project has no single binding. Always served (no omitempty) so the console reads an explicit null rather than an absent key. | [optional] 
**ClusterId** | Pointer to **int32** |  | [optional] 
**CurrentRevision** | Pointer to **int32** |  | [optional] 
**Detail** | Pointer to [**HandlersDeploykitInstallStatusDetail**](HandlersDeploykitInstallStatusDetail.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Generation** | Pointer to **int32** |  | [optional] 
**Install** | Pointer to **string** |  | [optional] 
**ManagedByStack** | Pointer to **string** | ManagedByStack is the managed-slot marker (epic managed-xyz story 4, AC2 — additive and omitempty): the OWNING STACK&#39;s install name when the install is a managed slot&#39;s provisioning (a ManagedSlotBinding claims it). The typed-apps law is \&quot;unified list, per-type surfaces\&quot; — slot installs are SHOWN, as managed rows the FE badges from this field; the install-scoped verbs are fenced by AC5&#39;s refusal (owned by stack; use the stack surface). | [optional] 
**Namespace** | Pointer to **string** | Namespace is the record&#39;s PhysicalName (story 149.4, additive and omitempty): the literal k8s namespace the engine renders everything into — the stacks surface&#39;s pods bridge scopes the K8s View to it. It equals InstallName at fresh install and diverges permanently after a restore cutover (the alias law), so the client never derives it from the install name. | [optional] 
**Park** | Pointer to [**GitopsUpgradeParkBadge**](GitopsUpgradeParkBadge.md) | Park is the honest-state badge (issue #2709 §2.7, additive and omitempty): the install&#39;s active parked upgrade — a parked app reads as parked on the list surface too, never only on its detail. | [optional] 
**StableRevision** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersDeploykitInstallSummary

`func NewHandlersDeploykitInstallSummary() *HandlersDeploykitInstallSummary`

NewHandlersDeploykitInstallSummary instantiates a new HandlersDeploykitInstallSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDeploykitInstallSummaryWithDefaults

`func NewHandlersDeploykitInstallSummaryWithDefaults() *HandlersDeploykitInstallSummary`

NewHandlersDeploykitInstallSummaryWithDefaults instantiates a new HandlersDeploykitInstallSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildBinding

`func (o *HandlersDeploykitInstallSummary) GetBuildBinding() HandlersDeploykitBuildBinding`

GetBuildBinding returns the BuildBinding field if non-nil, zero value otherwise.

### GetBuildBindingOk

`func (o *HandlersDeploykitInstallSummary) GetBuildBindingOk() (*HandlersDeploykitBuildBinding, bool)`

GetBuildBindingOk returns a tuple with the BuildBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildBinding

`func (o *HandlersDeploykitInstallSummary) SetBuildBinding(v HandlersDeploykitBuildBinding)`

SetBuildBinding sets BuildBinding field to given value.

### HasBuildBinding

`func (o *HandlersDeploykitInstallSummary) HasBuildBinding() bool`

HasBuildBinding returns a boolean if a field has been set.

### GetClusterId

`func (o *HandlersDeploykitInstallSummary) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HandlersDeploykitInstallSummary) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HandlersDeploykitInstallSummary) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HandlersDeploykitInstallSummary) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCurrentRevision

`func (o *HandlersDeploykitInstallSummary) GetCurrentRevision() int32`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *HandlersDeploykitInstallSummary) GetCurrentRevisionOk() (*int32, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *HandlersDeploykitInstallSummary) SetCurrentRevision(v int32)`

SetCurrentRevision sets CurrentRevision field to given value.

### HasCurrentRevision

`func (o *HandlersDeploykitInstallSummary) HasCurrentRevision() bool`

HasCurrentRevision returns a boolean if a field has been set.

### GetDetail

`func (o *HandlersDeploykitInstallSummary) GetDetail() HandlersDeploykitInstallStatusDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *HandlersDeploykitInstallSummary) GetDetailOk() (*HandlersDeploykitInstallStatusDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *HandlersDeploykitInstallSummary) SetDetail(v HandlersDeploykitInstallStatusDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *HandlersDeploykitInstallSummary) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDisplayName

`func (o *HandlersDeploykitInstallSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HandlersDeploykitInstallSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HandlersDeploykitInstallSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HandlersDeploykitInstallSummary) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGeneration

`func (o *HandlersDeploykitInstallSummary) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *HandlersDeploykitInstallSummary) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *HandlersDeploykitInstallSummary) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *HandlersDeploykitInstallSummary) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetInstall

`func (o *HandlersDeploykitInstallSummary) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *HandlersDeploykitInstallSummary) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *HandlersDeploykitInstallSummary) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *HandlersDeploykitInstallSummary) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetManagedByStack

`func (o *HandlersDeploykitInstallSummary) GetManagedByStack() string`

GetManagedByStack returns the ManagedByStack field if non-nil, zero value otherwise.

### GetManagedByStackOk

`func (o *HandlersDeploykitInstallSummary) GetManagedByStackOk() (*string, bool)`

GetManagedByStackOk returns a tuple with the ManagedByStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedByStack

`func (o *HandlersDeploykitInstallSummary) SetManagedByStack(v string)`

SetManagedByStack sets ManagedByStack field to given value.

### HasManagedByStack

`func (o *HandlersDeploykitInstallSummary) HasManagedByStack() bool`

HasManagedByStack returns a boolean if a field has been set.

### GetNamespace

`func (o *HandlersDeploykitInstallSummary) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *HandlersDeploykitInstallSummary) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *HandlersDeploykitInstallSummary) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *HandlersDeploykitInstallSummary) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPark

`func (o *HandlersDeploykitInstallSummary) GetPark() GitopsUpgradeParkBadge`

GetPark returns the Park field if non-nil, zero value otherwise.

### GetParkOk

`func (o *HandlersDeploykitInstallSummary) GetParkOk() (*GitopsUpgradeParkBadge, bool)`

GetParkOk returns a tuple with the Park field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPark

`func (o *HandlersDeploykitInstallSummary) SetPark(v GitopsUpgradeParkBadge)`

SetPark sets Park field to given value.

### HasPark

`func (o *HandlersDeploykitInstallSummary) HasPark() bool`

HasPark returns a boolean if a field has been set.

### GetStableRevision

`func (o *HandlersDeploykitInstallSummary) GetStableRevision() int32`

GetStableRevision returns the StableRevision field if non-nil, zero value otherwise.

### GetStableRevisionOk

`func (o *HandlersDeploykitInstallSummary) GetStableRevisionOk() (*int32, bool)`

GetStableRevisionOk returns a tuple with the StableRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableRevision

`func (o *HandlersDeploykitInstallSummary) SetStableRevision(v int32)`

SetStableRevision sets StableRevision field to given value.

### HasStableRevision

`func (o *HandlersDeploykitInstallSummary) HasStableRevision() bool`

HasStableRevision returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersDeploykitInstallSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersDeploykitInstallSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersDeploykitInstallSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersDeploykitInstallSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


