# HandlersStoreDelegationMandatesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuedVia** | Pointer to **string** |  | [optional] 
**Mandates** | Pointer to [**[]HandlersDelegationMandateRequest**](HandlersDelegationMandateRequest.md) |  | [optional] 

## Methods

### NewHandlersStoreDelegationMandatesRequest

`func NewHandlersStoreDelegationMandatesRequest() *HandlersStoreDelegationMandatesRequest`

NewHandlersStoreDelegationMandatesRequest instantiates a new HandlersStoreDelegationMandatesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStoreDelegationMandatesRequestWithDefaults

`func NewHandlersStoreDelegationMandatesRequestWithDefaults() *HandlersStoreDelegationMandatesRequest`

NewHandlersStoreDelegationMandatesRequestWithDefaults instantiates a new HandlersStoreDelegationMandatesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuedVia

`func (o *HandlersStoreDelegationMandatesRequest) GetIssuedVia() string`

GetIssuedVia returns the IssuedVia field if non-nil, zero value otherwise.

### GetIssuedViaOk

`func (o *HandlersStoreDelegationMandatesRequest) GetIssuedViaOk() (*string, bool)`

GetIssuedViaOk returns a tuple with the IssuedVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedVia

`func (o *HandlersStoreDelegationMandatesRequest) SetIssuedVia(v string)`

SetIssuedVia sets IssuedVia field to given value.

### HasIssuedVia

`func (o *HandlersStoreDelegationMandatesRequest) HasIssuedVia() bool`

HasIssuedVia returns a boolean if a field has been set.

### GetMandates

`func (o *HandlersStoreDelegationMandatesRequest) GetMandates() []HandlersDelegationMandateRequest`

GetMandates returns the Mandates field if non-nil, zero value otherwise.

### GetMandatesOk

`func (o *HandlersStoreDelegationMandatesRequest) GetMandatesOk() (*[]HandlersDelegationMandateRequest, bool)`

GetMandatesOk returns a tuple with the Mandates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandates

`func (o *HandlersStoreDelegationMandatesRequest) SetMandates(v []HandlersDelegationMandateRequest)`

SetMandates sets Mandates field to given value.

### HasMandates

`func (o *HandlersStoreDelegationMandatesRequest) HasMandates() bool`

HasMandates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


