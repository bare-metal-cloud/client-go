# \RequirementsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1Requirements**](RequirementsAPI.md#GetApiV1Requirements) | **Get** /api/v1/requirements | List onboarding requirements
[**PostApiV1RequirementsKeyResolve**](RequirementsAPI.md#PostApiV1RequirementsKeyResolve) | **Post** /api/v1/requirements/{key}/resolve | Resolve an evidence requirement
[**PostApiV1RequirementsKeySkip**](RequirementsAPI.md#PostApiV1RequirementsKeySkip) | **Post** /api/v1/requirements/{key}/skip | Skip a requirement



## GetApiV1Requirements

> HandlersListRequirementsResponse GetApiV1Requirements(ctx).Execute()

List onboarding requirements



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
	resp, r, err := apiClient.RequirementsAPI.GetApiV1Requirements(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementsAPI.GetApiV1Requirements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1Requirements`: HandlersListRequirementsResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementsAPI.GetApiV1Requirements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1RequirementsRequest struct via the builder pattern


### Return type

[**HandlersListRequirementsResponse**](HandlersListRequirementsResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1RequirementsKeyResolve

> HandlersResolveResponse PostApiV1RequirementsKeyResolve(ctx, key).Body(body).Execute()

Resolve an evidence requirement



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
	key := "key_example" // string | Requirement key
	body := *openapiclient.NewHandlersResolveRequirementRequest() // HandlersResolveRequirementRequest | Resolution payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementsAPI.PostApiV1RequirementsKeyResolve(context.Background(), key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementsAPI.PostApiV1RequirementsKeyResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1RequirementsKeyResolve`: HandlersResolveResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementsAPI.PostApiV1RequirementsKeyResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Requirement key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1RequirementsKeyResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersResolveRequirementRequest**](HandlersResolveRequirementRequest.md) | Resolution payload | 

### Return type

[**HandlersResolveResponse**](HandlersResolveResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1RequirementsKeySkip

> HandlersResolveResponse PostApiV1RequirementsKeySkip(ctx, key).Execute()

Skip a requirement



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
	key := "key_example" // string | Requirement key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequirementsAPI.PostApiV1RequirementsKeySkip(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequirementsAPI.PostApiV1RequirementsKeySkip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1RequirementsKeySkip`: HandlersResolveResponse
	fmt.Fprintf(os.Stdout, "Response from `RequirementsAPI.PostApiV1RequirementsKeySkip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Requirement key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1RequirementsKeySkipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersResolveResponse**](HandlersResolveResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

