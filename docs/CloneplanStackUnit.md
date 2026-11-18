# CloneplanStackUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]CloneplanStackComponent**](CloneplanStackComponent.md) | Components are the stack&#39;s components in manifest (deploy) order — each with its own install facts (the component installs are the stack&#39;s durable truth; the parent record stays revision-less). | [optional] 
**Stack** | Pointer to **string** | Stack is the stack manifest&#39;s name (the capture&#39;s). | [optional] 
**Wiring** | Pointer to [**[]CloneplanWiringRef**](CloneplanWiringRef.md) | Wiring is the captured wiring edges as reference items: consumer, slot, provider, kind, and the wiring secret&#39;s NAME. Keys carries the edge&#39;s reference values (host, port) — never a secret value (the capture&#39;s own law). | [optional] 

## Methods

### NewCloneplanStackUnit

`func NewCloneplanStackUnit() *CloneplanStackUnit`

NewCloneplanStackUnit instantiates a new CloneplanStackUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanStackUnitWithDefaults

`func NewCloneplanStackUnitWithDefaults() *CloneplanStackUnit`

NewCloneplanStackUnitWithDefaults instantiates a new CloneplanStackUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *CloneplanStackUnit) GetComponents() []CloneplanStackComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *CloneplanStackUnit) GetComponentsOk() (*[]CloneplanStackComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *CloneplanStackUnit) SetComponents(v []CloneplanStackComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *CloneplanStackUnit) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetStack

`func (o *CloneplanStackUnit) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *CloneplanStackUnit) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *CloneplanStackUnit) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *CloneplanStackUnit) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetWiring

`func (o *CloneplanStackUnit) GetWiring() []CloneplanWiringRef`

GetWiring returns the Wiring field if non-nil, zero value otherwise.

### GetWiringOk

`func (o *CloneplanStackUnit) GetWiringOk() (*[]CloneplanWiringRef, bool)`

GetWiringOk returns a tuple with the Wiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiring

`func (o *CloneplanStackUnit) SetWiring(v []CloneplanWiringRef)`

SetWiring sets Wiring field to given value.

### HasWiring

`func (o *CloneplanStackUnit) HasWiring() bool`

HasWiring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


