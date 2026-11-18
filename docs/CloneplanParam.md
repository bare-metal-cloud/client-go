# CloneplanParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the param&#39;s canonical key (the frozen ClassifyParam grammar). | [optional] 
**Portability** | Pointer to **string** | Portability is the plan&#39;s classification (the Param* vocabulary). | [optional] 
**ProposedValue** | Pointer to **map[string]interface{}** | ProposedValue is the env_local param&#39;s proposed target value. Nil at derivation: nothing has resolved target facts yet — the review (or the target resolution it approves) fills it, and the replay renders the approved value. | [optional] 
**SourceValue** | Pointer to **map[string]interface{}** | SourceValue is the env_local param&#39;s source-side resolved value — shown beside the proposed target value (AC 5). | [optional] 
**Undeclared** | Pointer to **bool** | Undeclared flags the \&quot;no declaration exists\&quot; case (AC 5): the param replays verbatim and is flagged as such — the records-predating-the- portability-classes law. Nothing is guessed: an omitted class is a fact, not an implied portable. | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value is the param&#39;s resolved value — NON-sensitive by the revision-params law. Secret-class params NEVER carry one (the field stays empty; the value re-escrows on the target). | [optional] 

## Methods

### NewCloneplanParam

`func NewCloneplanParam() *CloneplanParam`

NewCloneplanParam instantiates a new CloneplanParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanParamWithDefaults

`func NewCloneplanParamWithDefaults() *CloneplanParam`

NewCloneplanParamWithDefaults instantiates a new CloneplanParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CloneplanParam) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloneplanParam) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloneplanParam) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloneplanParam) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPortability

`func (o *CloneplanParam) GetPortability() string`

GetPortability returns the Portability field if non-nil, zero value otherwise.

### GetPortabilityOk

`func (o *CloneplanParam) GetPortabilityOk() (*string, bool)`

GetPortabilityOk returns a tuple with the Portability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortability

`func (o *CloneplanParam) SetPortability(v string)`

SetPortability sets Portability field to given value.

### HasPortability

`func (o *CloneplanParam) HasPortability() bool`

HasPortability returns a boolean if a field has been set.

### GetProposedValue

`func (o *CloneplanParam) GetProposedValue() map[string]interface{}`

GetProposedValue returns the ProposedValue field if non-nil, zero value otherwise.

### GetProposedValueOk

`func (o *CloneplanParam) GetProposedValueOk() (*map[string]interface{}, bool)`

GetProposedValueOk returns a tuple with the ProposedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedValue

`func (o *CloneplanParam) SetProposedValue(v map[string]interface{})`

SetProposedValue sets ProposedValue field to given value.

### HasProposedValue

`func (o *CloneplanParam) HasProposedValue() bool`

HasProposedValue returns a boolean if a field has been set.

### GetSourceValue

`func (o *CloneplanParam) GetSourceValue() map[string]interface{}`

GetSourceValue returns the SourceValue field if non-nil, zero value otherwise.

### GetSourceValueOk

`func (o *CloneplanParam) GetSourceValueOk() (*map[string]interface{}, bool)`

GetSourceValueOk returns a tuple with the SourceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceValue

`func (o *CloneplanParam) SetSourceValue(v map[string]interface{})`

SetSourceValue sets SourceValue field to given value.

### HasSourceValue

`func (o *CloneplanParam) HasSourceValue() bool`

HasSourceValue returns a boolean if a field has been set.

### GetUndeclared

`func (o *CloneplanParam) GetUndeclared() bool`

GetUndeclared returns the Undeclared field if non-nil, zero value otherwise.

### GetUndeclaredOk

`func (o *CloneplanParam) GetUndeclaredOk() (*bool, bool)`

GetUndeclaredOk returns a tuple with the Undeclared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndeclared

`func (o *CloneplanParam) SetUndeclared(v bool)`

SetUndeclared sets Undeclared field to given value.

### HasUndeclared

`func (o *CloneplanParam) HasUndeclared() bool`

HasUndeclared returns a boolean if a field has been set.

### GetValue

`func (o *CloneplanParam) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloneplanParam) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloneplanParam) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CloneplanParam) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


