# HandlersBuildLedgerView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollbackAvailable** | Pointer to **bool** |  | [optional] 
**Services** | Pointer to [**[]HandlersBuildLedgerServiceGroup**](HandlersBuildLedgerServiceGroup.md) |  | [optional] 

## Methods

### NewHandlersBuildLedgerView

`func NewHandlersBuildLedgerView() *HandlersBuildLedgerView`

NewHandlersBuildLedgerView instantiates a new HandlersBuildLedgerView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersBuildLedgerViewWithDefaults

`func NewHandlersBuildLedgerViewWithDefaults() *HandlersBuildLedgerView`

NewHandlersBuildLedgerViewWithDefaults instantiates a new HandlersBuildLedgerView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollbackAvailable

`func (o *HandlersBuildLedgerView) GetRollbackAvailable() bool`

GetRollbackAvailable returns the RollbackAvailable field if non-nil, zero value otherwise.

### GetRollbackAvailableOk

`func (o *HandlersBuildLedgerView) GetRollbackAvailableOk() (*bool, bool)`

GetRollbackAvailableOk returns a tuple with the RollbackAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackAvailable

`func (o *HandlersBuildLedgerView) SetRollbackAvailable(v bool)`

SetRollbackAvailable sets RollbackAvailable field to given value.

### HasRollbackAvailable

`func (o *HandlersBuildLedgerView) HasRollbackAvailable() bool`

HasRollbackAvailable returns a boolean if a field has been set.

### GetServices

`func (o *HandlersBuildLedgerView) GetServices() []HandlersBuildLedgerServiceGroup`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *HandlersBuildLedgerView) GetServicesOk() (*[]HandlersBuildLedgerServiceGroup, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *HandlersBuildLedgerView) SetServices(v []HandlersBuildLedgerServiceGroup)`

SetServices sets Services field to given value.

### HasServices

`func (o *HandlersBuildLedgerView) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


