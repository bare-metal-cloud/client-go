# HandlersDeploykitResolveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrivalInstance** | Pointer to **string** |  | [optional] 
**ArrivalKind** | Pointer to **string** | Arrival is the generic arrival pick (\&quot;public\&quot; | \&quot;tunnel\&quot;) with its optional addon instance sub-choice; ZoneID is the registry pick. | [optional] 
**Blueprint** | Pointer to [**HandlersInstallBlueprintSelection**](HandlersInstallBlueprintSelection.md) | Blueprint is the deploy wizard&#39;s ORIGINALLY SELECTED source (epic 165 story 165.5 — additive and omitempty): the repo + ref + file the card resolved, persisted at install time so the update flow re-derives from the SAME selection (the never-silently-switch- files law). Absent on every pre-165.5 request. | [optional] 
**Bundle** | Pointer to [**BundleRef**](BundleRef.md) |  | [optional] 
**Facts** | Pointer to [**EngineContextFacts**](EngineContextFacts.md) |  | [optional] 
**Input** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlersDeploykitResolveRequest

`func NewHandlersDeploykitResolveRequest() *HandlersDeploykitResolveRequest`

NewHandlersDeploykitResolveRequest instantiates a new HandlersDeploykitResolveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDeploykitResolveRequestWithDefaults

`func NewHandlersDeploykitResolveRequestWithDefaults() *HandlersDeploykitResolveRequest`

NewHandlersDeploykitResolveRequestWithDefaults instantiates a new HandlersDeploykitResolveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrivalInstance

`func (o *HandlersDeploykitResolveRequest) GetArrivalInstance() string`

GetArrivalInstance returns the ArrivalInstance field if non-nil, zero value otherwise.

### GetArrivalInstanceOk

`func (o *HandlersDeploykitResolveRequest) GetArrivalInstanceOk() (*string, bool)`

GetArrivalInstanceOk returns a tuple with the ArrivalInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalInstance

`func (o *HandlersDeploykitResolveRequest) SetArrivalInstance(v string)`

SetArrivalInstance sets ArrivalInstance field to given value.

### HasArrivalInstance

`func (o *HandlersDeploykitResolveRequest) HasArrivalInstance() bool`

HasArrivalInstance returns a boolean if a field has been set.

### GetArrivalKind

`func (o *HandlersDeploykitResolveRequest) GetArrivalKind() string`

GetArrivalKind returns the ArrivalKind field if non-nil, zero value otherwise.

### GetArrivalKindOk

`func (o *HandlersDeploykitResolveRequest) GetArrivalKindOk() (*string, bool)`

GetArrivalKindOk returns a tuple with the ArrivalKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalKind

`func (o *HandlersDeploykitResolveRequest) SetArrivalKind(v string)`

SetArrivalKind sets ArrivalKind field to given value.

### HasArrivalKind

`func (o *HandlersDeploykitResolveRequest) HasArrivalKind() bool`

HasArrivalKind returns a boolean if a field has been set.

### GetBlueprint

`func (o *HandlersDeploykitResolveRequest) GetBlueprint() HandlersInstallBlueprintSelection`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *HandlersDeploykitResolveRequest) GetBlueprintOk() (*HandlersInstallBlueprintSelection, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *HandlersDeploykitResolveRequest) SetBlueprint(v HandlersInstallBlueprintSelection)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *HandlersDeploykitResolveRequest) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetBundle

`func (o *HandlersDeploykitResolveRequest) GetBundle() BundleRef`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *HandlersDeploykitResolveRequest) GetBundleOk() (*BundleRef, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *HandlersDeploykitResolveRequest) SetBundle(v BundleRef)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *HandlersDeploykitResolveRequest) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetFacts

`func (o *HandlersDeploykitResolveRequest) GetFacts() EngineContextFacts`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *HandlersDeploykitResolveRequest) GetFactsOk() (*EngineContextFacts, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *HandlersDeploykitResolveRequest) SetFacts(v EngineContextFacts)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *HandlersDeploykitResolveRequest) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetInput

`func (o *HandlersDeploykitResolveRequest) GetInput() map[string]map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *HandlersDeploykitResolveRequest) GetInputOk() (*map[string]map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *HandlersDeploykitResolveRequest) SetInput(v map[string]map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *HandlersDeploykitResolveRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetTarget

`func (o *HandlersDeploykitResolveRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *HandlersDeploykitResolveRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *HandlersDeploykitResolveRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *HandlersDeploykitResolveRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetZoneId

`func (o *HandlersDeploykitResolveRequest) GetZoneId() int32`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *HandlersDeploykitResolveRequest) GetZoneIdOk() (*int32, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *HandlersDeploykitResolveRequest) SetZoneId(v int32)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *HandlersDeploykitResolveRequest) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


