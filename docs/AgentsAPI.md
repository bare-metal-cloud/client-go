# \AgentsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentsClusterIdHealth**](AgentsAPI.md#GetAgentsClusterIdHealth) | **Get** /agents/{cluster_id}/health | Get agent health
[**GetAgentsClusterIdInfo**](AgentsAPI.md#GetAgentsClusterIdInfo) | **Get** /agents/{cluster_id}/info | Get agent node info
[**GetAgentsClusterIdQuery**](AgentsAPI.md#GetAgentsClusterIdQuery) | **Get** /agents/{cluster_id}/query | Query custom agent endpoint
[**GetAgentsClusterIdStats**](AgentsAPI.md#GetAgentsClusterIdStats) | **Get** /agents/{cluster_id}/stats | Get agent node stats



## GetAgentsClusterIdHealth

> map[string]interface{} GetAgentsClusterIdHealth(ctx, clusterId).Execute()

Get agent health



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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.GetAgentsClusterIdHealth(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsClusterIdHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsClusterIdHealth`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsClusterIdHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsClusterIdHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAgentsClusterIdInfo

> map[string]interface{} GetAgentsClusterIdInfo(ctx, clusterId).Execute()

Get agent node info



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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.GetAgentsClusterIdInfo(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsClusterIdInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsClusterIdInfo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsClusterIdInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsClusterIdInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAgentsClusterIdQuery

> map[string]interface{} GetAgentsClusterIdQuery(ctx, clusterId).Endpoint(endpoint).Execute()

Query custom agent endpoint



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
	clusterId := "clusterId_example" // string | Cluster ID
	endpoint := "endpoint_example" // string | Endpoint path (e.g., /query/custom)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.GetAgentsClusterIdQuery(context.Background(), clusterId).Endpoint(endpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsClusterIdQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsClusterIdQuery`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsClusterIdQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsClusterIdQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpoint** | **string** | Endpoint path (e.g., /query/custom) | 

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


## GetAgentsClusterIdStats

> map[string]interface{} GetAgentsClusterIdStats(ctx, clusterId).Execute()

Get agent node stats



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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.GetAgentsClusterIdStats(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.GetAgentsClusterIdStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsClusterIdStats`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.GetAgentsClusterIdStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsClusterIdStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

