# HandlersRunnerCacheEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]HandlersRunnerCacheEventRow**](HandlersRunnerCacheEventRow.md) |  | [optional] 
**Total** | Pointer to **int32** | Total is the pre-cap ledger size (the events surface&#39;s shape). | [optional] 

## Methods

### NewHandlersRunnerCacheEventsResponse

`func NewHandlersRunnerCacheEventsResponse() *HandlersRunnerCacheEventsResponse`

NewHandlersRunnerCacheEventsResponse instantiates a new HandlersRunnerCacheEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerCacheEventsResponseWithDefaults

`func NewHandlersRunnerCacheEventsResponseWithDefaults() *HandlersRunnerCacheEventsResponse`

NewHandlersRunnerCacheEventsResponseWithDefaults instantiates a new HandlersRunnerCacheEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *HandlersRunnerCacheEventsResponse) GetEvents() []HandlersRunnerCacheEventRow`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *HandlersRunnerCacheEventsResponse) GetEventsOk() (*[]HandlersRunnerCacheEventRow, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *HandlersRunnerCacheEventsResponse) SetEvents(v []HandlersRunnerCacheEventRow)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *HandlersRunnerCacheEventsResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTotal

`func (o *HandlersRunnerCacheEventsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HandlersRunnerCacheEventsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HandlersRunnerCacheEventsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HandlersRunnerCacheEventsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


