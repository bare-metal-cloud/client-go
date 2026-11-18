# HandlersStoreGuardTokensResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolved** | Pointer to **string** | the ask closed by this intake (\&quot;\&quot; when none) | [optional] 
**Stored** | Pointer to **[]string** | the op_ids persisted | [optional] 

## Methods

### NewHandlersStoreGuardTokensResponse

`func NewHandlersStoreGuardTokensResponse() *HandlersStoreGuardTokensResponse`

NewHandlersStoreGuardTokensResponse instantiates a new HandlersStoreGuardTokensResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStoreGuardTokensResponseWithDefaults

`func NewHandlersStoreGuardTokensResponseWithDefaults() *HandlersStoreGuardTokensResponse`

NewHandlersStoreGuardTokensResponseWithDefaults instantiates a new HandlersStoreGuardTokensResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolved

`func (o *HandlersStoreGuardTokensResponse) GetResolved() string`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *HandlersStoreGuardTokensResponse) GetResolvedOk() (*string, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *HandlersStoreGuardTokensResponse) SetResolved(v string)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *HandlersStoreGuardTokensResponse) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetStored

`func (o *HandlersStoreGuardTokensResponse) GetStored() []string`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *HandlersStoreGuardTokensResponse) GetStoredOk() (*[]string, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *HandlersStoreGuardTokensResponse) SetStored(v []string)`

SetStored sets Stored field to given value.

### HasStored

`func (o *HandlersStoreGuardTokensResponse) HasStored() bool`

HasStored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


