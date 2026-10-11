# \HelmOperationsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClustersIdHelmReleasesName**](HelmOperationsAPI.md#DeleteClustersIdHelmReleasesName) | **Delete** /clusters/{id}/helm/releases/{name} | Uninstall Helm Release
[**GetClustersIdHelmReleases**](HelmOperationsAPI.md#GetClustersIdHelmReleases) | **Get** /clusters/{id}/helm/releases | List Helm Releases
[**GetClustersIdHelmReleasesName**](HelmOperationsAPI.md#GetClustersIdHelmReleasesName) | **Get** /clusters/{id}/helm/releases/{name} | Get Helm Release Status
[**PostClustersIdHelmInstall**](HelmOperationsAPI.md#PostClustersIdHelmInstall) | **Post** /clusters/{id}/helm/install | Install Helm Chart



## DeleteClustersIdHelmReleasesName

> map[string]interface{} DeleteClustersIdHelmReleasesName(ctx, id, name).Namespace(namespace).Execute()

Uninstall Helm Release



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
	id := "id_example" // string | Cluster ID
	name := "name_example" // string | Release Name
	namespace := "namespace_example" // string | Filter by namespace (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmOperationsAPI.DeleteClustersIdHelmReleasesName(context.Background(), id, name).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmOperationsAPI.DeleteClustersIdHelmReleasesName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClustersIdHelmReleasesName`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `HelmOperationsAPI.DeleteClustersIdHelmReleasesName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**name** | **string** | Release Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClustersIdHelmReleasesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespace** | **string** | Filter by namespace | 

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


## GetClustersIdHelmReleases

> map[string]interface{} GetClustersIdHelmReleases(ctx, id).Namespace(namespace).Execute()

List Helm Releases



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
	id := "id_example" // string | Cluster ID
	namespace := "namespace_example" // string | Filter by namespace (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmOperationsAPI.GetClustersIdHelmReleases(context.Background(), id).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmOperationsAPI.GetClustersIdHelmReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdHelmReleases`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `HelmOperationsAPI.GetClustersIdHelmReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdHelmReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | Filter by namespace | 

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


## GetClustersIdHelmReleasesName

> map[string]interface{} GetClustersIdHelmReleasesName(ctx, id, name).Namespace(namespace).Execute()

Get Helm Release Status



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
	id := "id_example" // string | Cluster ID
	name := "name_example" // string | Release Name
	namespace := "namespace_example" // string | Filter by namespace (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmOperationsAPI.GetClustersIdHelmReleasesName(context.Background(), id, name).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmOperationsAPI.GetClustersIdHelmReleasesName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdHelmReleasesName`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `HelmOperationsAPI.GetClustersIdHelmReleasesName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**name** | **string** | Release Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdHelmReleasesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespace** | **string** | Filter by namespace | 

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


## PostClustersIdHelmInstall

> map[string]interface{} PostClustersIdHelmInstall(ctx, id).Request(request).Execute()

Install Helm Chart



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
	id := "id_example" // string | Cluster ID
	request := *openapiclient.NewHandlersInstallChartRequest() // HandlersInstallChartRequest | Installation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelmOperationsAPI.PostClustersIdHelmInstall(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelmOperationsAPI.PostClustersIdHelmInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdHelmInstall`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `HelmOperationsAPI.PostClustersIdHelmInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdHelmInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersInstallChartRequest**](HandlersInstallChartRequest.md) | Installation request | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

