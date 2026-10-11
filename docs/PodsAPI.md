# \PodsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClustersIdPods**](PodsAPI.md#GetClustersIdPods) | **Get** /clusters/{id}/pods | List pods in cluster
[**GetClustersIdPodsNamespacePod**](PodsAPI.md#GetClustersIdPodsNamespacePod) | **Get** /clusters/{id}/pods/{namespace}/{pod} | Get pod details
[**GetClustersIdPodsNamespacePodExec**](PodsAPI.md#GetClustersIdPodsNamespacePodExec) | **Get** /clusters/{id}/pods/{namespace}/{pod}/exec | Execute shell in pod via WebSocket
[**GetClustersIdPodsNamespacePodLogs**](PodsAPI.md#GetClustersIdPodsNamespacePodLogs) | **Get** /clusters/{id}/pods/{namespace}/{pod}/logs | Stream pod logs via WebSocket



## GetClustersIdPods

> map[string]interface{} GetClustersIdPods(ctx, id).Namespace(namespace).Execute()

List pods in cluster



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
	id := int32(56) // int32 | Cluster ID
	namespace := "namespace_example" // string | Namespace (default: all namespaces) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PodsAPI.GetClustersIdPods(context.Background(), id).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PodsAPI.GetClustersIdPods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdPods`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PodsAPI.GetClustersIdPods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Namespace (default: all namespaces) | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdPodsNamespacePod

> map[string]interface{} GetClustersIdPodsNamespacePod(ctx, id, namespace, pod).Execute()

Get pod details



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
	id := int32(56) // int32 | Cluster ID
	namespace := "namespace_example" // string | Pod namespace
	pod := "pod_example" // string | Pod name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PodsAPI.GetClustersIdPodsNamespacePod(context.Background(), id, namespace, pod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PodsAPI.GetClustersIdPodsNamespacePod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdPodsNamespacePod`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PodsAPI.GetClustersIdPodsNamespacePod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**namespace** | **string** | Pod namespace | 
**pod** | **string** | Pod name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdPodsNamespacePodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdPodsNamespacePodExec

> GetClustersIdPodsNamespacePodExec(ctx, id, namespace, pod).Container(container).Command(command).Execute()

Execute shell in pod via WebSocket



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
	id := int32(56) // int32 | Cluster ID
	namespace := "namespace_example" // string | Pod namespace
	pod := "pod_example" // string | Pod name
	container := "container_example" // string | Container name (optional, defaults to first container) (optional)
	command := "command_example" // string | Command to execute (default: /bin/sh) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PodsAPI.GetClustersIdPodsNamespacePodExec(context.Background(), id, namespace, pod).Container(container).Command(command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PodsAPI.GetClustersIdPodsNamespacePodExec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**namespace** | **string** | Pod namespace | 
**pod** | **string** | Pod name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdPodsNamespacePodExecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **container** | **string** | Container name (optional, defaults to first container) | 
 **command** | **string** | Command to execute (default: /bin/sh) | 

### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdPodsNamespacePodLogs

> GetClustersIdPodsNamespacePodLogs(ctx, id, namespace, pod).Container(container).Follow(follow).Tail(tail).Timestamps(timestamps).Previous(previous).Execute()

Stream pod logs via WebSocket



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
	id := int32(56) // int32 | Cluster ID
	namespace := "namespace_example" // string | Pod namespace
	pod := "pod_example" // string | Pod name
	container := "container_example" // string | Container name (optional, defaults to first container) (optional)
	follow := true // bool | Follow logs (stream continuously) (optional)
	tail := int32(56) // int32 | Number of lines from end (default: 100) (optional)
	timestamps := true // bool | Include timestamps (optional)
	previous := true // bool | Get logs from previous container instance (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PodsAPI.GetClustersIdPodsNamespacePodLogs(context.Background(), id, namespace, pod).Container(container).Follow(follow).Tail(tail).Timestamps(timestamps).Previous(previous).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PodsAPI.GetClustersIdPodsNamespacePodLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**namespace** | **string** | Pod namespace | 
**pod** | **string** | Pod name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdPodsNamespacePodLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **container** | **string** | Container name (optional, defaults to first container) | 
 **follow** | **bool** | Follow logs (stream continuously) | 
 **tail** | **int32** | Number of lines from end (default: 100) | 
 **timestamps** | **bool** | Include timestamps | 
 **previous** | **bool** | Get logs from previous container instance | 

### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

