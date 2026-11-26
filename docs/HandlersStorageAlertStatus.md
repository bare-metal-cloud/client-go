# HandlersStorageAlertStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveCount** | Pointer to **int32** |  | [optional] 
**Events** | Pointer to [**[]ModelsAlertEvent**](ModelsAlertEvent.md) |  | [optional] 
**RuleCount** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersStorageAlertStatus

`func NewHandlersStorageAlertStatus() *HandlersStorageAlertStatus`

NewHandlersStorageAlertStatus instantiates a new HandlersStorageAlertStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStorageAlertStatusWithDefaults

`func NewHandlersStorageAlertStatusWithDefaults() *HandlersStorageAlertStatus`

NewHandlersStorageAlertStatusWithDefaults instantiates a new HandlersStorageAlertStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveCount

`func (o *HandlersStorageAlertStatus) GetActiveCount() int32`

GetActiveCount returns the ActiveCount field if non-nil, zero value otherwise.

### GetActiveCountOk

`func (o *HandlersStorageAlertStatus) GetActiveCountOk() (*int32, bool)`

GetActiveCountOk returns a tuple with the ActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCount

`func (o *HandlersStorageAlertStatus) SetActiveCount(v int32)`

SetActiveCount sets ActiveCount field to given value.

### HasActiveCount

`func (o *HandlersStorageAlertStatus) HasActiveCount() bool`

HasActiveCount returns a boolean if a field has been set.

### GetEvents

`func (o *HandlersStorageAlertStatus) GetEvents() []ModelsAlertEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *HandlersStorageAlertStatus) GetEventsOk() (*[]ModelsAlertEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *HandlersStorageAlertStatus) SetEvents(v []ModelsAlertEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *HandlersStorageAlertStatus) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetRuleCount

`func (o *HandlersStorageAlertStatus) GetRuleCount() int32`

GetRuleCount returns the RuleCount field if non-nil, zero value otherwise.

### GetRuleCountOk

`func (o *HandlersStorageAlertStatus) GetRuleCountOk() (*int32, bool)`

GetRuleCountOk returns a tuple with the RuleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleCount

`func (o *HandlersStorageAlertStatus) SetRuleCount(v int32)`

SetRuleCount sets RuleCount field to given value.

### HasRuleCount

`func (o *HandlersStorageAlertStatus) HasRuleCount() bool`

HasRuleCount returns a boolean if a field has been set.

### GetState

`func (o *HandlersStorageAlertStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HandlersStorageAlertStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HandlersStorageAlertStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HandlersStorageAlertStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


