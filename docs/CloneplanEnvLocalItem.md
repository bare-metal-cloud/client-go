# CloneplanEnvLocalItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovedValue** | Pointer to **map[string]interface{}** | ApprovedValue is the review&#39;s approved target value (approved items only). | [optional] 
**Key** | Pointer to **string** | Key is \&quot;install.key\&quot; (a stack component&#39;s is its own install). | [optional] 
**SourceValue** | Pointer to **map[string]interface{}** | SourceValue is the source-side resolved value, as the plan carries it. | [optional] 

## Methods

### NewCloneplanEnvLocalItem

`func NewCloneplanEnvLocalItem() *CloneplanEnvLocalItem`

NewCloneplanEnvLocalItem instantiates a new CloneplanEnvLocalItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneplanEnvLocalItemWithDefaults

`func NewCloneplanEnvLocalItemWithDefaults() *CloneplanEnvLocalItem`

NewCloneplanEnvLocalItemWithDefaults instantiates a new CloneplanEnvLocalItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovedValue

`func (o *CloneplanEnvLocalItem) GetApprovedValue() map[string]interface{}`

GetApprovedValue returns the ApprovedValue field if non-nil, zero value otherwise.

### GetApprovedValueOk

`func (o *CloneplanEnvLocalItem) GetApprovedValueOk() (*map[string]interface{}, bool)`

GetApprovedValueOk returns a tuple with the ApprovedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedValue

`func (o *CloneplanEnvLocalItem) SetApprovedValue(v map[string]interface{})`

SetApprovedValue sets ApprovedValue field to given value.

### HasApprovedValue

`func (o *CloneplanEnvLocalItem) HasApprovedValue() bool`

HasApprovedValue returns a boolean if a field has been set.

### GetKey

`func (o *CloneplanEnvLocalItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloneplanEnvLocalItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloneplanEnvLocalItem) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloneplanEnvLocalItem) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSourceValue

`func (o *CloneplanEnvLocalItem) GetSourceValue() map[string]interface{}`

GetSourceValue returns the SourceValue field if non-nil, zero value otherwise.

### GetSourceValueOk

`func (o *CloneplanEnvLocalItem) GetSourceValueOk() (*map[string]interface{}, bool)`

GetSourceValueOk returns a tuple with the SourceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceValue

`func (o *CloneplanEnvLocalItem) SetSourceValue(v map[string]interface{})`

SetSourceValue sets SourceValue field to given value.

### HasSourceValue

`func (o *CloneplanEnvLocalItem) HasSourceValue() bool`

HasSourceValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


