# HandlersMFAEmailLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MfaToken** | **string** | Pending MFA challenge token | 

## Methods

### NewHandlersMFAEmailLinkRequest

`func NewHandlersMFAEmailLinkRequest(mfaToken string, ) *HandlersMFAEmailLinkRequest`

NewHandlersMFAEmailLinkRequest instantiates a new HandlersMFAEmailLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMFAEmailLinkRequestWithDefaults

`func NewHandlersMFAEmailLinkRequestWithDefaults() *HandlersMFAEmailLinkRequest`

NewHandlersMFAEmailLinkRequestWithDefaults instantiates a new HandlersMFAEmailLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMfaToken

`func (o *HandlersMFAEmailLinkRequest) GetMfaToken() string`

GetMfaToken returns the MfaToken field if non-nil, zero value otherwise.

### GetMfaTokenOk

`func (o *HandlersMFAEmailLinkRequest) GetMfaTokenOk() (*string, bool)`

GetMfaTokenOk returns a tuple with the MfaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaToken

`func (o *HandlersMFAEmailLinkRequest) SetMfaToken(v string)`

SetMfaToken sets MfaToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


