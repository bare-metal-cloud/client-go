# HandlersErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersErrorResponse

`func NewHandlersErrorResponse() *HandlersErrorResponse`

NewHandlersErrorResponse instantiates a new HandlersErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersErrorResponseWithDefaults

`func NewHandlersErrorResponseWithDefaults() *HandlersErrorResponse`

NewHandlersErrorResponseWithDefaults instantiates a new HandlersErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *HandlersErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HandlersErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HandlersErrorResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *HandlersErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *HandlersErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HandlersErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HandlersErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HandlersErrorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


