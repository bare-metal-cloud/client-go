# HandlersStoreGuardTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuedVia** | Pointer to **string** | IssuedVia records the minting surface (audit provenance: \&quot;destructive-click\&quot;, \&quot;re-approval\&quot;, \&quot;setup-render\&quot;, \&quot;runner-install\&quot;). | [optional] 
**ReapprovalId** | Pointer to **string** | ReapprovalID closes the stranded ask this mint answers (empty for a destructive click that is not answering anything). | [optional] 
**Tokens** | Pointer to [**[]HandlersGuardTokenRequest**](HandlersGuardTokenRequest.md) |  | [optional] 

## Methods

### NewHandlersStoreGuardTokensRequest

`func NewHandlersStoreGuardTokensRequest() *HandlersStoreGuardTokensRequest`

NewHandlersStoreGuardTokensRequest instantiates a new HandlersStoreGuardTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStoreGuardTokensRequestWithDefaults

`func NewHandlersStoreGuardTokensRequestWithDefaults() *HandlersStoreGuardTokensRequest`

NewHandlersStoreGuardTokensRequestWithDefaults instantiates a new HandlersStoreGuardTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuedVia

`func (o *HandlersStoreGuardTokensRequest) GetIssuedVia() string`

GetIssuedVia returns the IssuedVia field if non-nil, zero value otherwise.

### GetIssuedViaOk

`func (o *HandlersStoreGuardTokensRequest) GetIssuedViaOk() (*string, bool)`

GetIssuedViaOk returns a tuple with the IssuedVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedVia

`func (o *HandlersStoreGuardTokensRequest) SetIssuedVia(v string)`

SetIssuedVia sets IssuedVia field to given value.

### HasIssuedVia

`func (o *HandlersStoreGuardTokensRequest) HasIssuedVia() bool`

HasIssuedVia returns a boolean if a field has been set.

### GetReapprovalId

`func (o *HandlersStoreGuardTokensRequest) GetReapprovalId() string`

GetReapprovalId returns the ReapprovalId field if non-nil, zero value otherwise.

### GetReapprovalIdOk

`func (o *HandlersStoreGuardTokensRequest) GetReapprovalIdOk() (*string, bool)`

GetReapprovalIdOk returns a tuple with the ReapprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReapprovalId

`func (o *HandlersStoreGuardTokensRequest) SetReapprovalId(v string)`

SetReapprovalId sets ReapprovalId field to given value.

### HasReapprovalId

`func (o *HandlersStoreGuardTokensRequest) HasReapprovalId() bool`

HasReapprovalId returns a boolean if a field has been set.

### GetTokens

`func (o *HandlersStoreGuardTokensRequest) GetTokens() []HandlersGuardTokenRequest`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *HandlersStoreGuardTokensRequest) GetTokensOk() (*[]HandlersGuardTokenRequest, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *HandlersStoreGuardTokensRequest) SetTokens(v []HandlersGuardTokenRequest)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *HandlersStoreGuardTokensRequest) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


