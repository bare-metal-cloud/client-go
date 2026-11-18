# HandlersBlueprintDriftResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drift** | Pointer to [**BlueprintdriftResult**](BlueprintdriftResult.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersBlueprintDriftResponse

`func NewHandlersBlueprintDriftResponse() *HandlersBlueprintDriftResponse`

NewHandlersBlueprintDriftResponse instantiates a new HandlersBlueprintDriftResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersBlueprintDriftResponseWithDefaults

`func NewHandlersBlueprintDriftResponseWithDefaults() *HandlersBlueprintDriftResponse`

NewHandlersBlueprintDriftResponseWithDefaults instantiates a new HandlersBlueprintDriftResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrift

`func (o *HandlersBlueprintDriftResponse) GetDrift() BlueprintdriftResult`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *HandlersBlueprintDriftResponse) GetDriftOk() (*BlueprintdriftResult, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *HandlersBlueprintDriftResponse) SetDrift(v BlueprintdriftResult)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *HandlersBlueprintDriftResponse) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersBlueprintDriftResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersBlueprintDriftResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersBlueprintDriftResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersBlueprintDriftResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


