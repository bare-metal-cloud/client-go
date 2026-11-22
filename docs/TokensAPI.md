# \TokensAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMeTokensId**](TokensAPI.md#DeleteMeTokensId) | **Delete** /me/tokens/{id} | Revoke a personal access token
[**GetMeTokens**](TokensAPI.md#GetMeTokens) | **Get** /me/tokens | List the caller&#39;s personal access tokens
[**PostMeTokens**](TokensAPI.md#PostMeTokens) | **Post** /me/tokens | Create a personal access token



## DeleteMeTokensId

> HandlersPatMessageResponse DeleteMeTokensId(ctx, id).Execute()

Revoke a personal access token



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
	id := int32(56) // int32 | Token ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.DeleteMeTokensId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.DeleteMeTokensId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMeTokensId`: HandlersPatMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.DeleteMeTokensId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMeTokensIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersPatMessageResponse**](HandlersPatMessageResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeTokens

> HandlersPatListResponse GetMeTokens(ctx).Execute()

List the caller's personal access tokens



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
	resp, r, err := apiClient.TokensAPI.GetMeTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetMeTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeTokens`: HandlersPatListResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetMeTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeTokensRequest struct via the builder pattern


### Return type

[**HandlersPatListResponse**](HandlersPatListResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMeTokens

> HandlersPatCreatedResponse PostMeTokens(ctx).Body(body).Execute()

Create a personal access token



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
	body := *openapiclient.NewHandlersPatRequest() // HandlersPatRequest | Token name, expiry in days (0 = no expiry, absent = 90) and scopes (absent = full trust)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.PostMeTokens(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.PostMeTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMeTokens`: HandlersPatCreatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.PostMeTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMeTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersPatRequest**](HandlersPatRequest.md) | Token name, expiry in days (0 &#x3D; no expiry, absent &#x3D; 90) and scopes (absent &#x3D; full trust) | 

### Return type

[**HandlersPatCreatedResponse**](HandlersPatCreatedResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

