# \AuthAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthSessions**](AuthAPI.md#DeleteAuthSessions) | **Delete** /auth/sessions | Logout from the current session
[**GetAuthMe**](AuthAPI.md#GetAuthMe) | **Get** /auth/me | Get current authenticated user
[**GetAuthMeUiMode**](AuthAPI.md#GetAuthMeUiMode) | **Get** /auth/me/ui-mode | Get user&#39;s UI mode preference
[**GetAuthOauthDiagnostics**](AuthAPI.md#GetAuthOauthDiagnostics) | **Get** /auth/oauth-diagnostics | Get OAuth configuration diagnostics
[**PostApiV1AuthLoginMfa**](AuthAPI.md#PostApiV1AuthLoginMfa) | **Post** /api/v1/auth/login/mfa | MFA-Enforced Login Flow: complete MFA login
[**PostApiV1AuthMfaRecover**](AuthAPI.md#PostApiV1AuthMfaRecover) | **Post** /api/v1/auth/mfa/recover | MFA Backup Code Recovery
[**PostAuthSessions**](AuthAPI.md#PostAuthSessions) | **Post** /auth/sessions | Login to an existing account
[**PostAuthToken**](AuthAPI.md#PostAuthToken) | **Post** /auth/token | Generate a JWT token for the authenticated user
[**PostAuthTokenRefresh**](AuthAPI.md#PostAuthTokenRefresh) | **Post** /auth/token/refresh | Refresh access token
[**PostAuthUsers**](AuthAPI.md#PostAuthUsers) | **Post** /auth/users | Create a new user account
[**PostAuthVerify**](AuthAPI.md#PostAuthVerify) | **Post** /auth/verify | Verify email address
[**PutAuthMeUiMode**](AuthAPI.md#PutAuthMeUiMode) | **Put** /auth/me/ui-mode | Update user&#39;s UI mode preference



## DeleteAuthSessions

> map[string]bool DeleteAuthSessions(ctx).Execute()

Logout from the current session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.DeleteAuthSessions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.DeleteAuthSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAuthSessions`: map[string]bool
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.DeleteAuthSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthSessionsRequest struct via the builder pattern


### Return type

**map[string]bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthMe

> ModelsUser GetAuthMe(ctx).Execute()

Get current authenticated user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.GetAuthMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetAuthMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthMe`: ModelsUser
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetAuthMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthMeRequest struct via the builder pattern


### Return type

[**ModelsUser**](ModelsUser.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthMeUiMode

> map[string]string GetAuthMeUiMode(ctx).Execute()

Get user's UI mode preference



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.GetAuthMeUiMode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetAuthMeUiMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthMeUiMode`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetAuthMeUiMode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthMeUiModeRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOauthDiagnostics

> map[string]interface{} GetAuthOauthDiagnostics(ctx).Execute()

Get OAuth configuration diagnostics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.GetAuthOauthDiagnostics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.GetAuthOauthDiagnostics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthOauthDiagnostics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.GetAuthOauthDiagnostics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOauthDiagnosticsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1AuthLoginMfa

> HandlersMFALoginResponse PostApiV1AuthLoginMfa(ctx).Request(request).Execute()

MFA-Enforced Login Flow: complete MFA login



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	request := *openapiclient.NewHandlersMFALoginRequest("Code_example", "MfaToken_example") // HandlersMFALoginRequest | MFA token and TOTP code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.PostApiV1AuthLoginMfa(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiV1AuthLoginMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1AuthLoginMfa`: HandlersMFALoginResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.PostApiV1AuthLoginMfa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1AuthLoginMfaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersMFALoginRequest**](HandlersMFALoginRequest.md) | MFA token and TOTP code | 

### Return type

[**HandlersMFALoginResponse**](HandlersMFALoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1AuthMfaRecover

> HandlersMFARecoverResponse PostApiV1AuthMfaRecover(ctx).Request(request).Execute()

MFA Backup Code Recovery



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	request := *openapiclient.NewHandlersMFARecoverRequest("BackupCode_example", "MfaToken_example") // HandlersMFARecoverRequest | MFA challenge token and backup code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.PostApiV1AuthMfaRecover(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostApiV1AuthMfaRecover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1AuthMfaRecover`: HandlersMFARecoverResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.PostApiV1AuthMfaRecover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1AuthMfaRecoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersMFARecoverRequest**](HandlersMFARecoverRequest.md) | MFA challenge token and backup code | 

### Return type

[**HandlersMFARecoverResponse**](HandlersMFARecoverResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthSessions

> map[string]interface{} PostAuthSessions(ctx).Request(request).Execute()

Login to an existing account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	request := *openapiclient.NewHandlersLoginRequest("Email_example", "Password_example") // HandlersLoginRequest | Login credentials (email and password)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.PostAuthSessions(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostAuthSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthSessions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.PostAuthSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersLoginRequest**](HandlersLoginRequest.md) | Login credentials (email and password) | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthToken

> map[string]interface{} PostAuthToken(ctx).Execute()

Generate a JWT token for the authenticated user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.PostAuthToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostAuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthToken`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.PostAuthToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthTokenRefresh

> map[string]interface{} PostAuthTokenRefresh(ctx).Request(request).Execute()

Refresh access token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	request := *openapiclient.NewHandlersRefreshTokenRequest("RefreshToken_example") // HandlersRefreshTokenRequest | Refresh token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.PostAuthTokenRefresh(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostAuthTokenRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthTokenRefresh`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.PostAuthTokenRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthTokenRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersRefreshTokenRequest**](HandlersRefreshTokenRequest.md) | Refresh token | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthUsers

> map[string]interface{} PostAuthUsers(ctx).Request(request).Execute()

Create a new user account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	request := *openapiclient.NewHandlersSignupRequest("Email_example", "Name_example", "Password_example") // HandlersSignupRequest | Signup details (name, email, password)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.PostAuthUsers(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostAuthUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthUsers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.PostAuthUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersSignupRequest**](HandlersSignupRequest.md) | Signup details (name, email, password) | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthVerify

> map[string]interface{} PostAuthVerify(ctx).Request(request).Execute()

Verify email address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	request := *openapiclient.NewHandlersVerifyEmailRequest("Token_example") // HandlersVerifyEmailRequest | Verification token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.PostAuthVerify(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostAuthVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthVerify`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.PostAuthVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersVerifyEmailRequest**](HandlersVerifyEmailRequest.md) | Verification token | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAuthMeUiMode

> map[string]interface{} PutAuthMeUiMode(ctx).Request(request).Execute()

Update user's UI mode preference



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	request := *openapiclient.NewHandlersUIModeRequest() // HandlersUIModeRequest | UI mode (simple or advanced)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.PutAuthMeUiMode(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PutAuthMeUiMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAuthMeUiMode`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.PutAuthMeUiMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAuthMeUiModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersUIModeRequest**](HandlersUIModeRequest.md) | UI mode (simple or advanced) | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

