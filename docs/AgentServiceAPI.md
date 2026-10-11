# \AgentServiceAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentEnrollmentStatus**](AgentServiceAPI.md#GetAgentEnrollmentStatus) | **Get** /agent/enrollment/status | Recover enrollment identity for a consumed token
[**GetAgentServiceHealth**](AgentServiceAPI.md#GetAgentServiceHealth) | **Get** /agent-service/health | Agent service health check
[**PostAgentExchangeToken**](AgentServiceAPI.md#PostAgentExchangeToken) | **Post** /agent/exchange-token | Exchange registration token for JWT
[**PostAgentShutdown**](AgentServiceAPI.md#PostAgentShutdown) | **Post** /agent/shutdown | Receive agent shutdown notification



## GetAgentEnrollmentStatus

> map[string]string GetAgentEnrollmentStatus(ctx).Token(token).Execute()

Recover enrollment identity for a consumed token



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
	token := "token_example" // string | Registration token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentServiceAPI.GetAgentEnrollmentStatus(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.GetAgentEnrollmentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentEnrollmentStatus`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.GetAgentEnrollmentStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentEnrollmentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Registration token | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentServiceHealth

> AgentHealthResponse GetAgentServiceHealth(ctx).Execute()

Agent service health check



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
	resp, r, err := apiClient.AgentServiceAPI.GetAgentServiceHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.GetAgentServiceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentServiceHealth`: AgentHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.GetAgentServiceHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentServiceHealthRequest struct via the builder pattern


### Return type

[**AgentHealthResponse**](AgentHealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentExchangeToken

> map[string]string PostAgentExchangeToken(ctx).Request(request).Execute()

Exchange registration token for JWT



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
	request := map[string]interface{}{ ... } // map[string]interface{} | Exchange request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentServiceAPI.PostAgentExchangeToken(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.PostAgentExchangeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentExchangeToken`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.PostAgentExchangeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentExchangeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | **map[string]interface{}** | Exchange request | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentShutdown

> map[string]string PostAgentShutdown(ctx).Request(request).Execute()

Receive agent shutdown notification



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
	request := *openapiclient.NewAgentShutdownNotificationRequest() // AgentShutdownNotificationRequest | Shutdown notification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentServiceAPI.PostAgentShutdown(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.PostAgentShutdown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentShutdown`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.PostAgentShutdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentShutdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**AgentShutdownNotificationRequest**](AgentShutdownNotificationRequest.md) | Shutdown notification | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

