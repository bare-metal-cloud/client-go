# CloneplanDurabilityPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Proposed** | Pointer to [**EngineDurabilityPolicy**](EngineDurabilityPolicy.md) | Proposed is what the replay will set unless the review overrides — the platform default: mode off. | [optional] 
**Source** | Pointer to [**EngineDurabilityPolicy**](EngineDurabilityPolicy.md) | Source is the source install&#39;s durability policy, as stored — reference only, never copied into the proposal. | [optional] 

## Methods

### NewCloneplanDurabilityPlan

`func NewCloneplanDurabilityPlan() *CloneplanDurabilityPlan`

NewCloneplanDurabilityPlan instantiates a new CloneplanDurabilityPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanDurabilityPlanWithDefaults

`func NewCloneplanDurabilityPlanWithDefaults() *CloneplanDurabilityPlan`

NewCloneplanDurabilityPlanWithDefaults instantiates a new CloneplanDurabilityPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProposed

`func (o *CloneplanDurabilityPlan) GetProposed() EngineDurabilityPolicy`

GetProposed returns the Proposed field if non-nil, zero value otherwise.

### GetProposedOk

`func (o *CloneplanDurabilityPlan) GetProposedOk() (*EngineDurabilityPolicy, bool)`

GetProposedOk returns a tuple with the Proposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposed

`func (o *CloneplanDurabilityPlan) SetProposed(v EngineDurabilityPolicy)`

SetProposed sets Proposed field to given value.

### HasProposed

`func (o *CloneplanDurabilityPlan) HasProposed() bool`

HasProposed returns a boolean if a field has been set.

### GetSource

`func (o *CloneplanDurabilityPlan) GetSource() EngineDurabilityPolicy`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloneplanDurabilityPlan) GetSourceOk() (*EngineDurabilityPolicy, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloneplanDurabilityPlan) SetSource(v EngineDurabilityPolicy)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloneplanDurabilityPlan) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


