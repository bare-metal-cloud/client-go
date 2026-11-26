# HandlersAccountAlertEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryGaps** | Pointer to [**[]AlertpicksGap**](AlertpicksGap.md) |  | [optional] 
**Events** | Pointer to [**[]ModelsAlertEvent**](ModelsAlertEvent.md) |  | [optional] 

## Methods

### NewHandlersAccountAlertEventsResponse

`func NewHandlersAccountAlertEventsResponse() *HandlersAccountAlertEventsResponse`

NewHandlersAccountAlertEventsResponse instantiates a new HandlersAccountAlertEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersAccountAlertEventsResponseWithDefaults

`func NewHandlersAccountAlertEventsResponseWithDefaults() *HandlersAccountAlertEventsResponse`

NewHandlersAccountAlertEventsResponseWithDefaults instantiates a new HandlersAccountAlertEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryGaps

`func (o *HandlersAccountAlertEventsResponse) GetDeliveryGaps() []AlertpicksGap`

GetDeliveryGaps returns the DeliveryGaps field if non-nil, zero value otherwise.

### GetDeliveryGapsOk

`func (o *HandlersAccountAlertEventsResponse) GetDeliveryGapsOk() (*[]AlertpicksGap, bool)`

GetDeliveryGapsOk returns a tuple with the DeliveryGaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryGaps

`func (o *HandlersAccountAlertEventsResponse) SetDeliveryGaps(v []AlertpicksGap)`

SetDeliveryGaps sets DeliveryGaps field to given value.

### HasDeliveryGaps

`func (o *HandlersAccountAlertEventsResponse) HasDeliveryGaps() bool`

HasDeliveryGaps returns a boolean if a field has been set.

### GetEvents

`func (o *HandlersAccountAlertEventsResponse) GetEvents() []ModelsAlertEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *HandlersAccountAlertEventsResponse) GetEventsOk() (*[]ModelsAlertEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *HandlersAccountAlertEventsResponse) SetEvents(v []ModelsAlertEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *HandlersAccountAlertEventsResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


