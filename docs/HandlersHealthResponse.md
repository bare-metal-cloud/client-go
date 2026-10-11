# HandlersHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to **map[string]string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersHealthResponse

`func NewHandlersHealthResponse() *HandlersHealthResponse`

NewHandlersHealthResponse instantiates a new HandlersHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersHealthResponseWithDefaults

`func NewHandlersHealthResponseWithDefaults() *HandlersHealthResponse`

NewHandlersHealthResponseWithDefaults instantiates a new HandlersHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *HandlersHealthResponse) GetChecks() map[string]string`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *HandlersHealthResponse) GetChecksOk() (*map[string]string, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *HandlersHealthResponse) SetChecks(v map[string]string)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *HandlersHealthResponse) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetService

`func (o *HandlersHealthResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *HandlersHealthResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *HandlersHealthResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *HandlersHealthResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *HandlersHealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlersHealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlersHealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlersHealthResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *HandlersHealthResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HandlersHealthResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HandlersHealthResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *HandlersHealthResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


