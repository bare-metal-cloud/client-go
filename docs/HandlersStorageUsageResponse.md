# HandlersStorageUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**map[string]HandlersStorageConnectionUsage**](HandlersStorageConnectionUsage.md) |  | [optional] 

## Methods

### NewHandlersStorageUsageResponse

`func NewHandlersStorageUsageResponse() *HandlersStorageUsageResponse`

NewHandlersStorageUsageResponse instantiates a new HandlersStorageUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStorageUsageResponseWithDefaults

`func NewHandlersStorageUsageResponseWithDefaults() *HandlersStorageUsageResponse`

NewHandlersStorageUsageResponseWithDefaults instantiates a new HandlersStorageUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *HandlersStorageUsageResponse) GetConnections() map[string]HandlersStorageConnectionUsage`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *HandlersStorageUsageResponse) GetConnectionsOk() (*map[string]HandlersStorageConnectionUsage, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *HandlersStorageUsageResponse) SetConnections(v map[string]HandlersStorageConnectionUsage)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *HandlersStorageUsageResponse) HasConnections() bool`

HasConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


