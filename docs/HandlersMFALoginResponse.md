# HandlersMFALoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | CSRF token for the session | [optional] 
**MfaVerified** | Pointer to **bool** | Always true on success | [optional] 
**RefreshToken** | Pointer to **string** | Opaque refresh token | [optional] 
**Token** | Pointer to **string** | JWT access token | [optional] 

## Methods

### NewHandlersMFALoginResponse

`func NewHandlersMFALoginResponse() *HandlersMFALoginResponse`

NewHandlersMFALoginResponse instantiates a new HandlersMFALoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMFALoginResponseWithDefaults

`func NewHandlersMFALoginResponseWithDefaults() *HandlersMFALoginResponse`

NewHandlersMFALoginResponseWithDefaults instantiates a new HandlersMFALoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *HandlersMFALoginResponse) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *HandlersMFALoginResponse) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *HandlersMFALoginResponse) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *HandlersMFALoginResponse) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMfaVerified

`func (o *HandlersMFALoginResponse) GetMfaVerified() bool`

GetMfaVerified returns the MfaVerified field if non-nil, zero value otherwise.

### GetMfaVerifiedOk

`func (o *HandlersMFALoginResponse) GetMfaVerifiedOk() (*bool, bool)`

GetMfaVerifiedOk returns a tuple with the MfaVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerified

`func (o *HandlersMFALoginResponse) SetMfaVerified(v bool)`

SetMfaVerified sets MfaVerified field to given value.

### HasMfaVerified

`func (o *HandlersMFALoginResponse) HasMfaVerified() bool`

HasMfaVerified returns a boolean if a field has been set.

### GetRefreshToken

`func (o *HandlersMFALoginResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *HandlersMFALoginResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *HandlersMFALoginResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *HandlersMFALoginResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetToken

`func (o *HandlersMFALoginResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *HandlersMFALoginResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *HandlersMFALoginResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *HandlersMFALoginResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


