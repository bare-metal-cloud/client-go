# HandlersDatabaseExtensionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**Extensions** | Pointer to [**[]EngineportsPGExtension**](EngineportsPGExtension.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersDatabaseExtensionsResponse

`func NewHandlersDatabaseExtensionsResponse() *HandlersDatabaseExtensionsResponse`

NewHandlersDatabaseExtensionsResponse instantiates a new HandlersDatabaseExtensionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDatabaseExtensionsResponseWithDefaults

`func NewHandlersDatabaseExtensionsResponseWithDefaults() *HandlersDatabaseExtensionsResponse`

NewHandlersDatabaseExtensionsResponseWithDefaults instantiates a new HandlersDatabaseExtensionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *HandlersDatabaseExtensionsResponse) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *HandlersDatabaseExtensionsResponse) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *HandlersDatabaseExtensionsResponse) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *HandlersDatabaseExtensionsResponse) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetExtensions

`func (o *HandlersDatabaseExtensionsResponse) GetExtensions() []EngineportsPGExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *HandlersDatabaseExtensionsResponse) GetExtensionsOk() (*[]EngineportsPGExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *HandlersDatabaseExtensionsResponse) SetExtensions(v []EngineportsPGExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *HandlersDatabaseExtensionsResponse) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetReason

`func (o *HandlersDatabaseExtensionsResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *HandlersDatabaseExtensionsResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *HandlersDatabaseExtensionsResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *HandlersDatabaseExtensionsResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


