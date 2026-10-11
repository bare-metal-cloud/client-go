# HandlersMFARecoverResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | CSRF token for the session | [optional] 
**MfaDisabled** | Pointer to **bool** | True if last code consumed (AC5) | [optional] 
**MfaVerified** | Pointer to **bool** | Always true on success | [optional] 
**RefreshToken** | Pointer to **string** | Opaque refresh token | [optional] 
**RemainingBackupCodes** | Pointer to **int32** | Number of unused backup codes left | [optional] 
**Token** | Pointer to **string** | JWT access token | [optional] 

## Methods

### NewHandlersMFARecoverResponse

`func NewHandlersMFARecoverResponse() *HandlersMFARecoverResponse`

NewHandlersMFARecoverResponse instantiates a new HandlersMFARecoverResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersMFARecoverResponseWithDefaults

`func NewHandlersMFARecoverResponseWithDefaults() *HandlersMFARecoverResponse`

NewHandlersMFARecoverResponseWithDefaults instantiates a new HandlersMFARecoverResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *HandlersMFARecoverResponse) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *HandlersMFARecoverResponse) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *HandlersMFARecoverResponse) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *HandlersMFARecoverResponse) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetMfaDisabled

`func (o *HandlersMFARecoverResponse) GetMfaDisabled() bool`

GetMfaDisabled returns the MfaDisabled field if non-nil, zero value otherwise.

### GetMfaDisabledOk

`func (o *HandlersMFARecoverResponse) GetMfaDisabledOk() (*bool, bool)`

GetMfaDisabledOk returns a tuple with the MfaDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaDisabled

`func (o *HandlersMFARecoverResponse) SetMfaDisabled(v bool)`

SetMfaDisabled sets MfaDisabled field to given value.

### HasMfaDisabled

`func (o *HandlersMFARecoverResponse) HasMfaDisabled() bool`

HasMfaDisabled returns a boolean if a field has been set.

### GetMfaVerified

`func (o *HandlersMFARecoverResponse) GetMfaVerified() bool`

GetMfaVerified returns the MfaVerified field if non-nil, zero value otherwise.

### GetMfaVerifiedOk

`func (o *HandlersMFARecoverResponse) GetMfaVerifiedOk() (*bool, bool)`

GetMfaVerifiedOk returns a tuple with the MfaVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaVerified

`func (o *HandlersMFARecoverResponse) SetMfaVerified(v bool)`

SetMfaVerified sets MfaVerified field to given value.

### HasMfaVerified

`func (o *HandlersMFARecoverResponse) HasMfaVerified() bool`

HasMfaVerified returns a boolean if a field has been set.

### GetRefreshToken

`func (o *HandlersMFARecoverResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *HandlersMFARecoverResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *HandlersMFARecoverResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *HandlersMFARecoverResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRemainingBackupCodes

`func (o *HandlersMFARecoverResponse) GetRemainingBackupCodes() int32`

GetRemainingBackupCodes returns the RemainingBackupCodes field if non-nil, zero value otherwise.

### GetRemainingBackupCodesOk

`func (o *HandlersMFARecoverResponse) GetRemainingBackupCodesOk() (*int32, bool)`

GetRemainingBackupCodesOk returns a tuple with the RemainingBackupCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBackupCodes

`func (o *HandlersMFARecoverResponse) SetRemainingBackupCodes(v int32)`

SetRemainingBackupCodes sets RemainingBackupCodes field to given value.

### HasRemainingBackupCodes

`func (o *HandlersMFARecoverResponse) HasRemainingBackupCodes() bool`

HasRemainingBackupCodes returns a boolean if a field has been set.

### GetToken

`func (o *HandlersMFARecoverResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *HandlersMFARecoverResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *HandlersMFARecoverResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *HandlersMFARecoverResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


