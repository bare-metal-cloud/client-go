# HandlersCreateSSHKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Scope** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewHandlersCreateSSHKeyRequest

`func NewHandlersCreateSSHKeyRequest(key string, title string, ) *HandlersCreateSSHKeyRequest`

NewHandlersCreateSSHKeyRequest instantiates a new HandlersCreateSSHKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersCreateSSHKeyRequestWithDefaults

`func NewHandlersCreateSSHKeyRequestWithDefaults() *HandlersCreateSSHKeyRequest`

NewHandlersCreateSSHKeyRequestWithDefaults instantiates a new HandlersCreateSSHKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *HandlersCreateSSHKeyRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *HandlersCreateSSHKeyRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *HandlersCreateSSHKeyRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetScope

`func (o *HandlersCreateSSHKeyRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *HandlersCreateSSHKeyRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *HandlersCreateSSHKeyRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *HandlersCreateSSHKeyRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTitle

`func (o *HandlersCreateSSHKeyRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HandlersCreateSSHKeyRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HandlersCreateSSHKeyRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


