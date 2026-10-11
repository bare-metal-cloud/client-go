# \HealthAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealth**](HealthAPI.md#GetHealth) | **Get** /health | Health check endpoint
[**GetHealthDetailed**](HealthAPI.md#GetHealthDetailed) | **Get** /health/detailed | Detailed health check endpoint



## GetHealth

> HandlersHealthResponse GetHealth(ctx).Execute()

Health check endpoint



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
	resp, r, err := apiClient.HealthAPI.GetHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHealth`: HandlersHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthRequest struct via the builder pattern


### Return type

[**HandlersHealthResponse**](HandlersHealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealthDetailed

> HandlersDetailedHealthCheckResponse GetHealthDetailed(ctx).Execute()

Detailed health check endpoint



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
	resp, r, err := apiClient.HealthAPI.GetHealthDetailed(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GetHealthDetailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHealthDetailed`: HandlersDetailedHealthCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GetHealthDetailed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthDetailedRequest struct via the builder pattern


### Return type

[**HandlersDetailedHealthCheckResponse**](HandlersDetailedHealthCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

