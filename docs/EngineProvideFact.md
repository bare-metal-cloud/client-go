# EngineProvideFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityFree** | Pointer to **map[string]string** | CapacityFree maps the capacity keys to free figures (k8s resource quantities, the ProvidesContext shape) the platform can answer for. A declared requires capacity key this map cannot answer refuses the binding at resolve with the named key — the facts are the offer side&#39;s whole truth now, so an unanswered declared fact is a gap, never a guess. | [optional] 
**Name** | Pointer to **string** | Name is the offer&#39;s identity — the offering install&#39;s name or the registered endpoint&#39;s name: the reference a re-link recommendation names. | [optional] 
**Type** | Pointer to **string** | Type is the provider type the offer serves (the BindingSlot.Type vocabulary: postgres, redis, ...). | [optional] 
**Version** | Pointer to **string** | Version is the offered software version when the platform can answer for it. Empty means the offer reports no version — a slot whose requires declares a version constraint cannot verify against this offer, and a linked binding refuses at resolve naming the gap. | [optional] 

## Methods

### NewEngineProvideFact

`func NewEngineProvideFact() *EngineProvideFact`

NewEngineProvideFact instantiates a new EngineProvideFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineProvideFactWithDefaults

`func NewEngineProvideFactWithDefaults() *EngineProvideFact`

NewEngineProvideFactWithDefaults instantiates a new EngineProvideFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityFree

`func (o *EngineProvideFact) GetCapacityFree() map[string]string`

GetCapacityFree returns the CapacityFree field if non-nil, zero value otherwise.

### GetCapacityFreeOk

`func (o *EngineProvideFact) GetCapacityFreeOk() (*map[string]string, bool)`

GetCapacityFreeOk returns a tuple with the CapacityFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityFree

`func (o *EngineProvideFact) SetCapacityFree(v map[string]string)`

SetCapacityFree sets CapacityFree field to given value.

### HasCapacityFree

`func (o *EngineProvideFact) HasCapacityFree() bool`

HasCapacityFree returns a boolean if a field has been set.

### GetName

`func (o *EngineProvideFact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineProvideFact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineProvideFact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EngineProvideFact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EngineProvideFact) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EngineProvideFact) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EngineProvideFact) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EngineProvideFact) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *EngineProvideFact) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EngineProvideFact) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EngineProvideFact) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EngineProvideFact) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


