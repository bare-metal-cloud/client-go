# \StacksAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostClustersIdAppsInstallStackArchive**](StacksAPI.md#PostClustersIdAppsInstallStackArchive) | **Post** /clusters/{id}/apps/{install}/stack-archive | Export a stack install as one self-contained stack archive
[**PostClustersIdAppsInstallStackArchiveImport**](StacksAPI.md#PostClustersIdAppsInstallStackArchiveImport) | **Post** /clusters/{id}/apps/{install}/stack-archive-import | Import a stack from an exported stack archive set (restore-run dispatch)
[**PostClustersIdAppsInstallStackCopy**](StacksAPI.md#PostClustersIdAppsInstallStackCopy) | **Post** /clusters/{id}/apps/{install}/stack-copy | Copy a stack install to a new install on this cluster (one composed run)
[**PostClustersIdAppsInstallStackMove**](StacksAPI.md#PostClustersIdAppsInstallStackMove) | **Post** /clusters/{id}/apps/{install}/stack-move | Move a stack install to a new install on this cluster (fence-first, one composed run)



## PostClustersIdAppsInstallStackArchive

> map[string]interface{} PostClustersIdAppsInstallStackArchive(ctx, id, install).Execute()

Export a stack install as one self-contained stack archive



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
	install := "install_example" // string | The stack install's name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.PostClustersIdAppsInstallStackArchive(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.PostClustersIdAppsInstallStackArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallStackArchive`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.PostClustersIdAppsInstallStackArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | The stack install&#39;s name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallStackArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdAppsInstallStackArchiveImport

> map[string]interface{} PostClustersIdAppsInstallStackArchiveImport(ctx, id, install).Body(body).Execute()

Import a stack from an exported stack archive set (restore-run dispatch)



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
	id := int32(56) // int32 | Cluster ID (the import target)
	install := "install_example" // string | The target stack install's name
	body := *openapiclient.NewHandlersStackArchiveImportRequest() // HandlersStackArchiveImportRequest | The source set and the review decisions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.PostClustersIdAppsInstallStackArchiveImport(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.PostClustersIdAppsInstallStackArchiveImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallStackArchiveImport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.PostClustersIdAppsInstallStackArchiveImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID (the import target) | 
**install** | **string** | The target stack install&#39;s name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallStackArchiveImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersStackArchiveImportRequest**](HandlersStackArchiveImportRequest.md) | The source set and the review decisions | 

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


## PostClustersIdAppsInstallStackCopy

> map[string]interface{} PostClustersIdAppsInstallStackCopy(ctx, id, install).Body(body).Execute()

Copy a stack install to a new install on this cluster (one composed run)



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
	id := int32(56) // int32 | Cluster ID (the copy target)
	install := "install_example" // string | The target stack install's name
	body := *openapiclient.NewHandlersStackCopyRequest() // HandlersStackCopyRequest | The source set, the fence decision, and the review decisions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.PostClustersIdAppsInstallStackCopy(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.PostClustersIdAppsInstallStackCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallStackCopy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.PostClustersIdAppsInstallStackCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID (the copy target) | 
**install** | **string** | The target stack install&#39;s name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallStackCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersStackCopyRequest**](HandlersStackCopyRequest.md) | The source set, the fence decision, and the review decisions | 

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


## PostClustersIdAppsInstallStackMove

> map[string]interface{} PostClustersIdAppsInstallStackMove(ctx, id, install).Body(body).Execute()

Move a stack install to a new install on this cluster (fence-first, one composed run)



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
	id := int32(56) // int32 | Cluster ID (the move target)
	install := "install_example" // string | The target stack install's name
	body := *openapiclient.NewHandlersStackMoveRequest() // HandlersStackMoveRequest | The source set, the retirement choice, and the review decisions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.PostClustersIdAppsInstallStackMove(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.PostClustersIdAppsInstallStackMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallStackMove`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.PostClustersIdAppsInstallStackMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID (the move target) | 
**install** | **string** | The target stack install&#39;s name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallStackMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersStackMoveRequest**](HandlersStackMoveRequest.md) | The source set, the retirement choice, and the review decisions | 

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

