# HandlersStorageConnectionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumers** | Pointer to [**[]HandlersS3ConnectionConsumer**](HandlersS3ConnectionConsumer.md) | Consumers is every live resource storing its backups or artifacts in this connection (clusters with backup plans, managed databases, and the runner store / artifact store chips when stored rows exist) — the same shape the delete guard&#39;s 409 body carries. | [optional] 
**Used** | Pointer to **bool** | Used is Consumers non-empty; false means genuinely unused, and the console may then allow every verb (FR-Z-2). A read the handler could not make is a 500, never a false — unknown is never \&quot;unused\&quot; (FR-Z-13). | [optional] 

## Methods

### NewHandlersStorageConnectionUsage

`func NewHandlersStorageConnectionUsage() *HandlersStorageConnectionUsage`

NewHandlersStorageConnectionUsage instantiates a new HandlersStorageConnectionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStorageConnectionUsageWithDefaults

`func NewHandlersStorageConnectionUsageWithDefaults() *HandlersStorageConnectionUsage`

NewHandlersStorageConnectionUsageWithDefaults instantiates a new HandlersStorageConnectionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumers

`func (o *HandlersStorageConnectionUsage) GetConsumers() []HandlersS3ConnectionConsumer`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *HandlersStorageConnectionUsage) GetConsumersOk() (*[]HandlersS3ConnectionConsumer, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *HandlersStorageConnectionUsage) SetConsumers(v []HandlersS3ConnectionConsumer)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *HandlersStorageConnectionUsage) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetUsed

`func (o *HandlersStorageConnectionUsage) GetUsed() bool`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *HandlersStorageConnectionUsage) GetUsedOk() (*bool, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *HandlersStorageConnectionUsage) SetUsed(v bool)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *HandlersStorageConnectionUsage) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


