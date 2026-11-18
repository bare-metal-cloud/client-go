# \AgentServiceAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentEnrollmentStatus**](AgentServiceAPI.md#GetAgentEnrollmentStatus) | **Get** /agent/enrollment/status | Recover enrollment identity for a consumed token
[**GetAgentServiceHealth**](AgentServiceAPI.md#GetAgentServiceHealth) | **Get** /agent-service/health | Agent service health check
[**PostAgentExchangeToken**](AgentServiceAPI.md#PostAgentExchangeToken) | **Post** /agent/exchange-token | Exchange registration token for JWT
[**PostAgentShutdown**](AgentServiceAPI.md#PostAgentShutdown) | **Post** /agent/shutdown | Receive agent shutdown notification
[**PostAgentSupportBundles**](AgentServiceAPI.md#PostAgentSupportBundles) | **Post** /agent/support/bundles | Upload an agent support bundle
[**PostAgentUpdatesEvents**](AgentServiceAPI.md#PostAgentUpdatesEvents) | **Post** /agent/updates/events | Report an agent update restore event



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


## PostAgentSupportBundles

> AgentSupportBundleUploadResponse PostAgentSupportBundles(ctx).XBMCAgentID(xBMCAgentID).XBMCTs(xBMCTs).XBMCNonce(xBMCNonce).XBMCSig(xBMCSig).XBMCBundleNote(xBMCBundleNote).Execute()

Upload an agent support bundle



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
	xBMCAgentID := "xBMCAgentID_example" // string | Enrolled numeric agent id
	xBMCTs := "xBMCTs_example" // string | Unix timestamp (±60s skew window)
	xBMCNonce := "xBMCNonce_example" // string | Single-use hex nonce (replay window)
	xBMCSig := "xBMCSig_example" // string | Hex Ed25519 signature over agent_id|ts|nonce
	xBMCBundleNote := "xBMCBundleNote_example" // string | Optional operator note (masked upstream by the CLI) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentServiceAPI.PostAgentSupportBundles(context.Background()).XBMCAgentID(xBMCAgentID).XBMCTs(xBMCTs).XBMCNonce(xBMCNonce).XBMCSig(xBMCSig).XBMCBundleNote(xBMCBundleNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.PostAgentSupportBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentSupportBundles`: AgentSupportBundleUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.PostAgentSupportBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentSupportBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBMCAgentID** | **string** | Enrolled numeric agent id | 
 **xBMCTs** | **string** | Unix timestamp (±60s skew window) | 
 **xBMCNonce** | **string** | Single-use hex nonce (replay window) | 
 **xBMCSig** | **string** | Hex Ed25519 signature over agent_id|ts|nonce | 
 **xBMCBundleNote** | **string** | Optional operator note (masked upstream by the CLI) | 

### Return type

[**AgentSupportBundleUploadResponse**](AgentSupportBundleUploadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentUpdatesEvents

> AgentUpdateRestoreEventResponse PostAgentUpdatesEvents(ctx).XBMCAgentID(xBMCAgentID).XBMCTs(xBMCTs).XBMCNonce(xBMCNonce).XBMCSig(xBMCSig).Event(event).Execute()

Report an agent update restore event



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
	xBMCAgentID := "xBMCAgentID_example" // string | Enrolled numeric agent id
	xBMCTs := "xBMCTs_example" // string | Unix timestamp (±60s skew window)
	xBMCNonce := "xBMCNonce_example" // string | Single-use hex nonce (replay window)
	xBMCSig := "xBMCSig_example" // string | Hex Ed25519 signature over agent_id|ts|nonce
	event := *openapiclient.NewAgentUpdateRestoreEventRequest() // AgentUpdateRestoreEventRequest | The restore event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentServiceAPI.PostAgentUpdatesEvents(context.Background()).XBMCAgentID(xBMCAgentID).XBMCTs(xBMCTs).XBMCNonce(xBMCNonce).XBMCSig(xBMCSig).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentServiceAPI.PostAgentUpdatesEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentUpdatesEvents`: AgentUpdateRestoreEventResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentServiceAPI.PostAgentUpdatesEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentUpdatesEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBMCAgentID** | **string** | Enrolled numeric agent id | 
 **xBMCTs** | **string** | Unix timestamp (±60s skew window) | 
 **xBMCNonce** | **string** | Single-use hex nonce (replay window) | 
 **xBMCSig** | **string** | Hex Ed25519 signature over agent_id|ts|nonce | 
 **event** | [**AgentUpdateRestoreEventRequest**](AgentUpdateRestoreEventRequest.md) | The restore event | 

### Return type

[**AgentUpdateRestoreEventResponse**](AgentUpdateRestoreEventResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

