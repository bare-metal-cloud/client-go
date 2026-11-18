# CloneplanStackComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** | Component is the stack manifest&#39;s component name. | [optional] 
**Facts** | Pointer to [**CloneplanInstallFacts**](CloneplanInstallFacts.md) | Facts are the component install&#39;s facts (its own record, revision, bundle, params, data class, durability). | [optional] 
**Position** | Pointer to **int32** | Position is the component&#39;s index in the manifest&#39;s canonical order (the capture&#39;s deploy-order tiebreak). | [optional] 

## Methods

### NewCloneplanStackComponent

`func NewCloneplanStackComponent() *CloneplanStackComponent`

NewCloneplanStackComponent instantiates a new CloneplanStackComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanStackComponentWithDefaults

`func NewCloneplanStackComponentWithDefaults() *CloneplanStackComponent`

NewCloneplanStackComponentWithDefaults instantiates a new CloneplanStackComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *CloneplanStackComponent) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *CloneplanStackComponent) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *CloneplanStackComponent) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *CloneplanStackComponent) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetFacts

`func (o *CloneplanStackComponent) GetFacts() CloneplanInstallFacts`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *CloneplanStackComponent) GetFactsOk() (*CloneplanInstallFacts, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *CloneplanStackComponent) SetFacts(v CloneplanInstallFacts)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *CloneplanStackComponent) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetPosition

`func (o *CloneplanStackComponent) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CloneplanStackComponent) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CloneplanStackComponent) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CloneplanStackComponent) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


