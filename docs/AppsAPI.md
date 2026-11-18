# \AppsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClustersIdAppsInstallSlots**](AppsAPI.md#DeleteClustersIdAppsInstallSlots) | **Delete** /clusters/{id}/apps/{install}/slots | Delete failed stack slots of an app install
[**GetClustersIdAppsInstallEdgeMetrics**](AppsAPI.md#GetClustersIdAppsInstallEdgeMetrics) | **Get** /clusters/{id}/apps/{install}/edge-metrics | Read an app install&#39;s edge metrics
[**GetClustersIdAppsInstallStackStatus**](AppsAPI.md#GetClustersIdAppsInstallStackStatus) | **Get** /clusters/{id}/apps/{install}/stack-status | Read an app install&#39;s stack status
[**PostClustersIdAppsInstallUpgradeAbort**](AppsAPI.md#PostClustersIdAppsInstallUpgradeAbort) | **Post** /clusters/{id}/apps/{install}/upgrade/abort | Abort a paused upgrade
[**PostClustersIdAppsInstallUpgradeResume**](AppsAPI.md#PostClustersIdAppsInstallUpgradeResume) | **Post** /clusters/{id}/apps/{install}/upgrade/resume | Resume a paused upgrade



## DeleteClustersIdAppsInstallSlots

> map[string]interface{} DeleteClustersIdAppsInstallSlots(ctx, id, install).Execute()

Delete failed stack slots of an app install

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
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.DeleteClustersIdAppsInstallSlots(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.DeleteClustersIdAppsInstallSlots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClustersIdAppsInstallSlots`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.DeleteClustersIdAppsInstallSlots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClustersIdAppsInstallSlotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsInstallEdgeMetrics

> map[string]interface{} GetClustersIdAppsInstallEdgeMetrics(ctx, id, install).Window(window).Execute()

Read an app install's edge metrics

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
	install := "install_example" // string | Install name
	window := "window_example" // string | 1h by default (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetClustersIdAppsInstallEdgeMetrics(context.Background(), id, install).Window(window).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetClustersIdAppsInstallEdgeMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallEdgeMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetClustersIdAppsInstallEdgeMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallEdgeMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **window** | **string** | 1h by default | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsInstallStackStatus

> map[string]interface{} GetClustersIdAppsInstallStackStatus(ctx, id, install).Execute()

Read an app install's stack status

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
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.GetClustersIdAppsInstallStackStatus(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.GetClustersIdAppsInstallStackStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallStackStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.GetClustersIdAppsInstallStackStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallStackStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdAppsInstallUpgradeAbort

> map[string]interface{} PostClustersIdAppsInstallUpgradeAbort(ctx, id, install).Execute()

Abort a paused upgrade

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
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.PostClustersIdAppsInstallUpgradeAbort(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.PostClustersIdAppsInstallUpgradeAbort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallUpgradeAbort`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.PostClustersIdAppsInstallUpgradeAbort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallUpgradeAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdAppsInstallUpgradeResume

> map[string]interface{} PostClustersIdAppsInstallUpgradeResume(ctx, id, install).Execute()

Resume a paused upgrade



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
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.PostClustersIdAppsInstallUpgradeResume(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.PostClustersIdAppsInstallUpgradeResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallUpgradeResume`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.PostClustersIdAppsInstallUpgradeResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallUpgradeResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

