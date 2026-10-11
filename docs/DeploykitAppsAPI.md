# \DeploykitAppsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClustersIdAppsInstallForm**](DeploykitAppsAPI.md#GetClustersIdAppsInstallForm) | **Get** /clusters/{id}/apps/{install}/form | Render the deploy form (FormSpec) for an install&#39;s bundle
[**GetClustersIdAppsVersions**](DeploykitAppsAPI.md#GetClustersIdAppsVersions) | **Get** /clusters/{id}/apps/versions | List the catalog versions of one bundle this engine can run
[**PostClustersIdAppsInstallResolve**](DeploykitAppsAPI.md#PostClustersIdAppsInstallResolve) | **Post** /clusters/{id}/apps/{install}/resolve | Resolve the form against user input and platform context facts



## GetClustersIdAppsInstallForm

> GetClustersIdAppsInstallForm(ctx, id, install).Execute()

Render the deploy form (FormSpec) for an install's bundle

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
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallForm(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsVersions

> GetClustersIdAppsVersions(ctx, id).Bundle(bundle).Execute()

List the catalog versions of one bundle this engine can run

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
	bundle := "bundle_example" // string | Bundle name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsVersions(context.Background(), id).Bundle(bundle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundle** | **string** | Bundle name | 

### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdAppsInstallResolve

> PostClustersIdAppsInstallResolve(ctx, id, install).Body(body).Execute()

Resolve the form against user input and platform context facts

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
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersDeploykitResolveRequest() // HandlersDeploykitResolveRequest | Resolve request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallResolve(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersDeploykitResolveRequest**](HandlersDeploykitResolveRequest.md) | Resolve request | 

### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

