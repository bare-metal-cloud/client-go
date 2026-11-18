# EngineClassCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | Class is the storage class the fact describes. | [optional] 
**EligibleNodeCount** | Pointer to **string** | EligibleNodeCount is the platform&#39;s declared count of nodes eligible for the class, as a decimal string (\&quot;3\&quot;) so an UNPARSEABLE count classifies at the consume point. ABSENT (nil) is MISSING-FACT; negative or unparseable is MALFORMED-FACT; a per-node fact declaring 0 is the capacity-UNSUPPORTED refusal; a system fact carrying a NONZERO count is MALFORMED-FACT (a system fact declares 0). | [optional] 
**Nodes** | Pointer to [**[]EngineNodeCapacity**](EngineNodeCapacity.md) | Nodes are the per-node free figures; a system fact lawfully carries none. An empty or partial list (len(nodes) &lt; eligible_node_count) is MISSING-FACT; duplicated nodes or MORE entries than the declared eligible count are MALFORMED-FACT. omitempty: a system fact serializes without it. | [optional] 
**Scope** | Pointer to [**EngineCapacityScope**](EngineCapacityScope.md) | Scope is the fact&#39;s scope kind (the frozen CapacityScope vocabulary); a value outside it is MALFORMED-FACT. | [optional] 
**SystemFreeBytes** | Pointer to **string** | SystemFreeBytes is the system scope&#39;s single free figure (a quantity string); a per-node fact MUST carry no value — the exclusivity is MALFORMED-FACT. omitempty: a per-node fact serializes without it. | [optional] 

## Methods

### NewEngineClassCapacity

`func NewEngineClassCapacity() *EngineClassCapacity`

NewEngineClassCapacity instantiates a new EngineClassCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineClassCapacityWithDefaults

`func NewEngineClassCapacityWithDefaults() *EngineClassCapacity`

NewEngineClassCapacityWithDefaults instantiates a new EngineClassCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *EngineClassCapacity) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *EngineClassCapacity) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *EngineClassCapacity) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *EngineClassCapacity) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetEligibleNodeCount

`func (o *EngineClassCapacity) GetEligibleNodeCount() string`

GetEligibleNodeCount returns the EligibleNodeCount field if non-nil, zero value otherwise.

### GetEligibleNodeCountOk

`func (o *EngineClassCapacity) GetEligibleNodeCountOk() (*string, bool)`

GetEligibleNodeCountOk returns a tuple with the EligibleNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleNodeCount

`func (o *EngineClassCapacity) SetEligibleNodeCount(v string)`

SetEligibleNodeCount sets EligibleNodeCount field to given value.

### HasEligibleNodeCount

`func (o *EngineClassCapacity) HasEligibleNodeCount() bool`

HasEligibleNodeCount returns a boolean if a field has been set.

### GetNodes

`func (o *EngineClassCapacity) GetNodes() []EngineNodeCapacity`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *EngineClassCapacity) GetNodesOk() (*[]EngineNodeCapacity, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *EngineClassCapacity) SetNodes(v []EngineNodeCapacity)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *EngineClassCapacity) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetScope

`func (o *EngineClassCapacity) GetScope() EngineCapacityScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EngineClassCapacity) GetScopeOk() (*EngineCapacityScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EngineClassCapacity) SetScope(v EngineCapacityScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EngineClassCapacity) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSystemFreeBytes

`func (o *EngineClassCapacity) GetSystemFreeBytes() string`

GetSystemFreeBytes returns the SystemFreeBytes field if non-nil, zero value otherwise.

### GetSystemFreeBytesOk

`func (o *EngineClassCapacity) GetSystemFreeBytesOk() (*string, bool)`

GetSystemFreeBytesOk returns a tuple with the SystemFreeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFreeBytes

`func (o *EngineClassCapacity) SetSystemFreeBytes(v string)`

SetSystemFreeBytes sets SystemFreeBytes field to given value.

### HasSystemFreeBytes

`func (o *EngineClassCapacity) HasSystemFreeBytes() bool`

HasSystemFreeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


