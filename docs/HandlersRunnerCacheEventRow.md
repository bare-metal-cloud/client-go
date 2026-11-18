# HandlersRunnerCacheEventRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int32** | At is the event&#39;s server stamp (unix seconds). | [optional] 
**Detail** | Pointer to **string** | Detail is the server&#39;s one-line record of the action. | [optional] 
**Kind** | Pointer to **string** | Kind is the action&#39;s word (promotion | eviction | trim | cold_fill | refusal). | [optional] 
**Scope** | Pointer to **string** | Scope is the event&#39;s cache scope. | [optional] 

## Methods

### NewHandlersRunnerCacheEventRow

`func NewHandlersRunnerCacheEventRow() *HandlersRunnerCacheEventRow`

NewHandlersRunnerCacheEventRow instantiates a new HandlersRunnerCacheEventRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerCacheEventRowWithDefaults

`func NewHandlersRunnerCacheEventRowWithDefaults() *HandlersRunnerCacheEventRow`

NewHandlersRunnerCacheEventRowWithDefaults instantiates a new HandlersRunnerCacheEventRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *HandlersRunnerCacheEventRow) GetAt() int32`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *HandlersRunnerCacheEventRow) GetAtOk() (*int32, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *HandlersRunnerCacheEventRow) SetAt(v int32)`

SetAt sets At field to given value.

### HasAt

`func (o *HandlersRunnerCacheEventRow) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetDetail

`func (o *HandlersRunnerCacheEventRow) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *HandlersRunnerCacheEventRow) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *HandlersRunnerCacheEventRow) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *HandlersRunnerCacheEventRow) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetKind

`func (o *HandlersRunnerCacheEventRow) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HandlersRunnerCacheEventRow) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HandlersRunnerCacheEventRow) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HandlersRunnerCacheEventRow) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetScope

`func (o *HandlersRunnerCacheEventRow) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersRunnerCacheEventRow) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersRunnerCacheEventRow) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersRunnerCacheEventRow) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


