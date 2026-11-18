# \AgentUpdateAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdAgentUpdatePreferences**](AgentUpdateAPI.md#GetOrgsOrgIdAgentUpdatePreferences) | **Get** /orgs/{org_id}/agent-update-preferences | Get the org&#39;s agent update preference
[**PostAgentsIdUpdate**](AgentUpdateAPI.md#PostAgentsIdUpdate) | **Post** /agents/{id}/update | Update an agent to a target version (org-scoped manual path)
[**PostClustersIdAgentUpdate**](AgentUpdateAPI.md#PostClustersIdAgentUpdate) | **Post** /clusters/{id}/agent-update | Update a cluster&#39;s bmc-agent-k8s agent to a pinned image tag
[**PutOrgsOrgIdAgentUpdatePreferences**](AgentUpdateAPI.md#PutOrgsOrgIdAgentUpdatePreferences) | **Put** /orgs/{org_id}/agent-update-preferences | Update the org&#39;s agent update preference



## GetOrgsOrgIdAgentUpdatePreferences

> ModelsAgentUpdatePreference GetOrgsOrgIdAgentUpdatePreferences(ctx, orgId).Execute()

Get the org's agent update preference



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
	orgId := "orgId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentUpdateAPI.GetOrgsOrgIdAgentUpdatePreferences(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentUpdateAPI.GetOrgsOrgIdAgentUpdatePreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdAgentUpdatePreferences`: ModelsAgentUpdatePreference
	fmt.Fprintf(os.Stdout, "Response from `AgentUpdateAPI.GetOrgsOrgIdAgentUpdatePreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdAgentUpdatePreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsAgentUpdatePreference**](ModelsAgentUpdatePreference.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentsIdUpdate

> map[string]interface{} PostAgentsIdUpdate(ctx, id).Request(request).Execute()

Update an agent to a target version (org-scoped manual path)



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
	id := int32(56) // int32 | Agent DB id
	request := *openapiclient.NewHandlersAgentUpdateRequest() // HandlersAgentUpdateRequest | The optional target version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentUpdateAPI.PostAgentsIdUpdate(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentUpdateAPI.PostAgentsIdUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentsIdUpdate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentUpdateAPI.PostAgentsIdUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Agent DB id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentsIdUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersAgentUpdateRequest**](HandlersAgentUpdateRequest.md) | The optional target version | 

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


## PostClustersIdAgentUpdate

> map[string]interface{} PostClustersIdAgentUpdate(ctx, id).Request(request).Execute()

Update a cluster's bmc-agent-k8s agent to a pinned image tag



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
	id := int32(56) // int32 | Cluster DB id
	request := *openapiclient.NewHandlersClusterAgentUpdateRequest() // HandlersClusterAgentUpdateRequest | The optional explicit image tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentUpdateAPI.PostClustersIdAgentUpdate(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentUpdateAPI.PostClustersIdAgentUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAgentUpdate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentUpdateAPI.PostClustersIdAgentUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster DB id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAgentUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersClusterAgentUpdateRequest**](HandlersClusterAgentUpdateRequest.md) | The optional explicit image tag | 

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


## PutOrgsOrgIdAgentUpdatePreferences

> ModelsAgentUpdatePreference PutOrgsOrgIdAgentUpdatePreferences(ctx, orgId).Request(request).Execute()

Update the org's agent update preference



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
	orgId := "orgId_example" // string | Organization ID
	request := *openapiclient.NewHandlersUpdateAgentUpdatePreferenceRequest() // HandlersUpdateAgentUpdatePreferenceRequest | The preference delta

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentUpdateAPI.PutOrgsOrgIdAgentUpdatePreferences(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentUpdateAPI.PutOrgsOrgIdAgentUpdatePreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdAgentUpdatePreferences`: ModelsAgentUpdatePreference
	fmt.Fprintf(os.Stdout, "Response from `AgentUpdateAPI.PutOrgsOrgIdAgentUpdatePreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdAgentUpdatePreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersUpdateAgentUpdatePreferenceRequest**](HandlersUpdateAgentUpdatePreferenceRequest.md) | The preference delta | 

### Return type

[**ModelsAgentUpdatePreference**](ModelsAgentUpdatePreference.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

