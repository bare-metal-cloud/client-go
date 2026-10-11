# HandlersDeploykitResolveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bundle** | Pointer to [**BundleRef**](BundleRef.md) |  | [optional] 
**Facts** | Pointer to [**EngineContextFacts**](EngineContextFacts.md) |  | [optional] 
**Input** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


