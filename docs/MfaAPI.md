# \MfaAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1MfaStatus**](MfaAPI.md#GetApiV1MfaStatus) | **Get** /api/v1/mfa/status | Get MFA status
[**PostApiV1MfaDisable**](MfaAPI.md#PostApiV1MfaDisable) | **Post** /api/v1/mfa/disable | Disable own MFA
[**PostApiV1MfaSetup**](MfaAPI.md#PostApiV1MfaSetup) | **Post** /api/v1/mfa/setup | Setup MFA (generate TOTP secret)
[**PostApiV1MfaVerify**](MfaAPI.md#PostApiV1MfaVerify) | **Post** /api/v1/mfa/verify | Verify MFA (enable MFA + issue backup codes)



## GetApiV1MfaStatus

> HandlersMFAStatusResponse GetApiV1MfaStatus(ctx).Execute()

Get MFA status



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
	resp, r, err := apiClient.MfaAPI.GetApiV1MfaStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MfaAPI.GetApiV1MfaStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1MfaStatus`: HandlersMFAStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `MfaAPI.GetApiV1MfaStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MfaStatusRequest struct via the builder pattern


### Return type

[**HandlersMFAStatusResponse**](HandlersMFAStatusResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1MfaDisable

> HandlersMFADisableResponse PostApiV1MfaDisable(ctx).Request(request).Execute()

Disable own MFA



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
	request := *openapiclient.NewHandlersMFADisableRequest() // HandlersMFADisableRequest | Current TOTP code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MfaAPI.PostApiV1MfaDisable(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MfaAPI.PostApiV1MfaDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1MfaDisable`: HandlersMFADisableResponse
	fmt.Fprintf(os.Stdout, "Response from `MfaAPI.PostApiV1MfaDisable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1MfaDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersMFADisableRequest**](HandlersMFADisableRequest.md) | Current TOTP code | 

### Return type

[**HandlersMFADisableResponse**](HandlersMFADisableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1MfaSetup

> HandlersMFASetupResponse PostApiV1MfaSetup(ctx).Execute()

Setup MFA (generate TOTP secret)



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
	resp, r, err := apiClient.MfaAPI.PostApiV1MfaSetup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MfaAPI.PostApiV1MfaSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1MfaSetup`: HandlersMFASetupResponse
	fmt.Fprintf(os.Stdout, "Response from `MfaAPI.PostApiV1MfaSetup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1MfaSetupRequest struct via the builder pattern


### Return type

[**HandlersMFASetupResponse**](HandlersMFASetupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1MfaVerify

> HandlersMFAVerifyResponse PostApiV1MfaVerify(ctx).Request(request).Execute()

Verify MFA (enable MFA + issue backup codes)



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
	request := *openapiclient.NewHandlersMFAVerifyRequest() // HandlersMFAVerifyRequest | TOTP code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MfaAPI.PostApiV1MfaVerify(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MfaAPI.PostApiV1MfaVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1MfaVerify`: HandlersMFAVerifyResponse
	fmt.Fprintf(os.Stdout, "Response from `MfaAPI.PostApiV1MfaVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1MfaVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersMFAVerifyRequest**](HandlersMFAVerifyRequest.md) | TOTP code | 

### Return type

[**HandlersMFAVerifyResponse**](HandlersMFAVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

