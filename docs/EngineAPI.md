# \EngineAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClustersIdEngine**](EngineAPI.md#GetClustersIdEngine) | **Get** /clusters/{id}/engine | Engine surfaces and status for a cluster
[**GetClustersIdEngineInterrupted**](EngineAPI.md#GetClustersIdEngineInterrupted) | **Get** /clusters/{id}/engine/interrupted | Interrupted-work report
[**GetClustersIdEngineMigration**](EngineAPI.md#GetClustersIdEngineMigration) | **Get** /clusters/{id}/engine/migration | Migration wizard state
[**GetClustersIdEnginePipelineRuns**](EngineAPI.md#GetClustersIdEnginePipelineRuns) | **Get** /clusters/{id}/engine/pipeline-runs | Pipeline-run history (the customer plane)
[**PostClustersIdEngineDispatch**](EngineAPI.md#PostClustersIdEngineDispatch) | **Post** /clusters/{id}/engine/dispatch | Dispatch a pipeline run manually
[**PostClustersIdEngineMigration**](EngineAPI.md#PostClustersIdEngineMigration) | **Post** /clusters/{id}/engine/migration | Start the migration wizard
[**PostClustersIdEngineMigrationAbort**](EngineAPI.md#PostClustersIdEngineMigrationAbort) | **Post** /clusters/{id}/engine/migration/abort | Abort the migration before the flip
[**PostClustersIdEngineMigrationCutover**](EngineAPI.md#PostClustersIdEngineMigrationCutover) | **Post** /clusters/{id}/engine/migration/cutover | Execute the migration cutover
[**PostClustersIdEngineNamespaces**](EngineAPI.md#PostClustersIdEngineNamespaces) | **Post** /clusters/{id}/engine/namespaces | Update the declared namespace set
[**PostClustersIdEngineRevert**](EngineAPI.md#PostClustersIdEngineRevert) | **Post** /clusters/{id}/engine/revert | Revert to the BMC engine



## GetClustersIdEngine

> map[string]interface{} GetClustersIdEngine(ctx, id).Execute()

Engine surfaces and status for a cluster



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.GetClustersIdEngine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.GetClustersIdEngine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdEngine`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.GetClustersIdEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdEngineRequest struct via the builder pattern


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


## GetClustersIdEngineInterrupted

> map[string]interface{} GetClustersIdEngineInterrupted(ctx, id).Execute()

Interrupted-work report



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.GetClustersIdEngineInterrupted(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.GetClustersIdEngineInterrupted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdEngineInterrupted`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.GetClustersIdEngineInterrupted`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdEngineInterruptedRequest struct via the builder pattern


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


## GetClustersIdEngineMigration

> map[string]interface{} GetClustersIdEngineMigration(ctx, id).Execute()

Migration wizard state



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.GetClustersIdEngineMigration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.GetClustersIdEngineMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdEngineMigration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.GetClustersIdEngineMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdEngineMigrationRequest struct via the builder pattern


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


## GetClustersIdEnginePipelineRuns

> map[string]interface{} GetClustersIdEnginePipelineRuns(ctx, id).Limit(limit).Execute()

Pipeline-run history (the customer plane)



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
	limit := int32(56) // int32 | Maximum rows (default 50, capped at 500) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.GetClustersIdEnginePipelineRuns(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.GetClustersIdEnginePipelineRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdEnginePipelineRuns`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.GetClustersIdEnginePipelineRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdEnginePipelineRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum rows (default 50, capped at 500) | 

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


## PostClustersIdEngineDispatch

> map[string]interface{} PostClustersIdEngineDispatch(ctx, id).Body(body).Execute()

Dispatch a pipeline run manually



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
	body := map[string]interface{}{ ... } // map[string]interface{} | repo, branch?, ref

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.PostClustersIdEngineDispatch(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.PostClustersIdEngineDispatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdEngineDispatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.PostClustersIdEngineDispatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdEngineDispatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | repo, branch?, ref | 

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


## PostClustersIdEngineMigration

> map[string]interface{} PostClustersIdEngineMigration(ctx, id).Body(body).Execute()

Start the migration wizard



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
	body := map[string]interface{}{ ... } // map[string]interface{} | repo, git_connection_id, namespaces, the three consent booleans

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.PostClustersIdEngineMigration(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.PostClustersIdEngineMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdEngineMigration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.PostClustersIdEngineMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdEngineMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | repo, git_connection_id, namespaces, the three consent booleans | 

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


## PostClustersIdEngineMigrationAbort

> map[string]interface{} PostClustersIdEngineMigrationAbort(ctx, id).Execute()

Abort the migration before the flip



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.PostClustersIdEngineMigrationAbort(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.PostClustersIdEngineMigrationAbort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdEngineMigrationAbort`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.PostClustersIdEngineMigrationAbort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdEngineMigrationAbortRequest struct via the builder pattern


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


## PostClustersIdEngineMigrationCutover

> map[string]interface{} PostClustersIdEngineMigrationCutover(ctx, id).Execute()

Execute the migration cutover



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.PostClustersIdEngineMigrationCutover(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.PostClustersIdEngineMigrationCutover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdEngineMigrationCutover`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.PostClustersIdEngineMigrationCutover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdEngineMigrationCutoverRequest struct via the builder pattern


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


## PostClustersIdEngineNamespaces

> map[string]interface{} PostClustersIdEngineNamespaces(ctx, id).Body(body).Execute()

Update the declared namespace set



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
	body := map[string]interface{}{ ... } // map[string]interface{} | namespaces: string array

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.PostClustersIdEngineNamespaces(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.PostClustersIdEngineNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdEngineNamespaces`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.PostClustersIdEngineNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdEngineNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | namespaces: string array | 

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


## PostClustersIdEngineRevert

> map[string]interface{} PostClustersIdEngineRevert(ctx, id).Execute()

Revert to the BMC engine



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.PostClustersIdEngineRevert(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.PostClustersIdEngineRevert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdEngineRevert`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.PostClustersIdEngineRevert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdEngineRevertRequest struct via the builder pattern


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

