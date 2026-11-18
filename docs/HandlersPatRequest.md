# HandlersPatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresInDays** | Pointer to **int32** | pointer: nil &#x3D; absent, 0 &#x3D; explicit no-expiry | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHandlersPatRequest

`func NewHandlersPatRequest() *HandlersPatRequest`

NewHandlersPatRequest instantiates a new HandlersPatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersPatRequestWithDefaults

`func NewHandlersPatRequestWithDefaults() *HandlersPatRequest`

NewHandlersPatRequestWithDefaults instantiates a new HandlersPatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresInDays

`func (o *HandlersPatRequest) GetExpiresInDays() int32`

GetExpiresInDays returns the ExpiresInDays field if non-nil, zero value otherwise.

### GetExpiresInDaysOk

`func (o *HandlersPatRequest) GetExpiresInDaysOk() (*int32, bool)`

GetExpiresInDaysOk returns a tuple with the ExpiresInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInDays

`func (o *HandlersPatRequest) SetExpiresInDays(v int32)`

SetExpiresInDays sets ExpiresInDays field to given value.

### HasExpiresInDays

`func (o *HandlersPatRequest) HasExpiresInDays() bool`

HasExpiresInDays returns a boolean if a field has been set.

### GetName

`func (o *HandlersPatRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersPatRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersPatRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HandlersPatRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScopes

`func (o *HandlersPatRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *HandlersPatRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *HandlersPatRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *HandlersPatRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


