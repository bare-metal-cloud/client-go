# HandlersTimelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Services** | Pointer to [**[]BuilderTimelineRow**](BuilderTimelineRow.md) |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersTimelineRequest

`func NewHandlersTimelineRequest() *HandlersTimelineRequest`

NewHandlersTimelineRequest instantiates a new HandlersTimelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersTimelineRequestWithDefaults

`func NewHandlersTimelineRequestWithDefaults() *HandlersTimelineRequest`

NewHandlersTimelineRequestWithDefaults instantiates a new HandlersTimelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *HandlersTimelineRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HandlersTimelineRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HandlersTimelineRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *HandlersTimelineRequest) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetServices

`func (o *HandlersTimelineRequest) GetServices() []BuilderTimelineRow`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *HandlersTimelineRequest) GetServicesOk() (*[]BuilderTimelineRow, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *HandlersTimelineRequest) SetServices(v []BuilderTimelineRow)`

SetServices sets Services field to given value.

### HasServices

`func (o *HandlersTimelineRequest) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetStack

`func (o *HandlersTimelineRequest) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *HandlersTimelineRequest) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *HandlersTimelineRequest) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *HandlersTimelineRequest) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


