# \ClustersAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClustersId**](ClustersAPI.md#DeleteClustersId) | **Delete** /clusters/{id} | Delete a cluster
[**GetApiClustersValidateNameName**](ClustersAPI.md#GetApiClustersValidateNameName) | **Get** /api/clusters/validate-name/{name} | Validate cluster name
[**GetClusters**](ClustersAPI.md#GetClusters) | **Get** /clusters | List all clusters
[**GetClustersId**](ClustersAPI.md#GetClustersId) | **Get** /clusters/{id} | Get cluster details
[**GetClustersIdAgentDeployment**](ClustersAPI.md#GetClustersIdAgentDeployment) | **Get** /clusters/{id}/agent-deployment | Get cluster agent deployment progress
[**GetClustersIdAgentDeploymentWs**](ClustersAPI.md#GetClustersIdAgentDeploymentWs) | **Get** /clusters/{id}/agent-deployment-ws | WebSocket for cluster agent deployment progress
[**GetClustersIdEvents**](ClustersAPI.md#GetClustersIdEvents) | **Get** /clusters/{id}/events | Get cluster status transition history
[**GetClustersIdExplorerInfo**](ClustersAPI.md#GetClustersIdExplorerInfo) | **Get** /clusters/{id}/explorer/info | Get K8s dashboard access information
[**GetClustersIdNetwork**](ClustersAPI.md#GetClustersIdNetwork) | **Get** /clusters/{id}/network | Get cluster network
[**GetClustersIdVirtualMachinesNew**](ClustersAPI.md#GetClustersIdVirtualMachinesNew) | **Get** /clusters/{id}/virtual-machines/new | Get available VMs to add to cluster
[**GetClustersNew**](ClustersAPI.md#GetClustersNew) | **Get** /clusters/new | Get cluster creation options
[**PostClusters**](ClustersAPI.md#PostClusters) | **Post** /clusters | Create a new cluster
[**PostClustersIdExplorerSetup**](ClustersAPI.md#PostClustersIdExplorerSetup) | **Post** /clusters/{id}/explorer/setup | Setup K8s cluster explorer tunnel
[**PostClustersIdNetwork**](ClustersAPI.md#PostClustersIdNetwork) | **Post** /clusters/{id}/network | Set cluster network
[**PostClustersIdVirtualMachines**](ClustersAPI.md#PostClustersIdVirtualMachines) | **Post** /clusters/{id}/virtual-machines | Add VMs to cluster
[**PostClustersImport**](ClustersAPI.md#PostClustersImport) | **Post** /clusters/import | Import an existing cluster (unified)
[**PostClustersRoleRecommendations**](ClustersAPI.md#PostClustersRoleRecommendations) | **Post** /clusters/role-recommendations | Recommend cluster roles for a VM selection



## DeleteClustersId

> map[string]interface{} DeleteClustersId(ctx, id).Body(body).Execute()

Delete a cluster



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
	body := *openapiclient.NewHandlersDeleteClusterRequest() // HandlersDeleteClusterRequest | Per-VM dispositions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.DeleteClustersId(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.DeleteClustersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClustersId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.DeleteClustersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClustersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersDeleteClusterRequest**](HandlersDeleteClusterRequest.md) | Per-VM dispositions | 

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


## GetApiClustersValidateNameName

> HandlersValidateClusterNameResponse GetApiClustersValidateNameName(ctx, name).Execute()

Validate cluster name



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
	name := "name_example" // string | Cluster name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetApiClustersValidateNameName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetApiClustersValidateNameName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiClustersValidateNameName`: HandlersValidateClusterNameResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetApiClustersValidateNameName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiClustersValidateNameNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersValidateClusterNameResponse**](HandlersValidateClusterNameResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusters

> []ModelsCluster GetClusters(ctx).Execute()

List all clusters



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
	resp, r, err := apiClient.ClustersAPI.GetClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClusters`: []ModelsCluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersRequest struct via the builder pattern


### Return type

[**[]ModelsCluster**](ModelsCluster.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersId

> map[string]interface{} GetClustersId(ctx, id).Execute()

Get cluster details



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClustersId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdRequest struct via the builder pattern


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


## GetClustersIdAgentDeployment

> map[string]interface{} GetClustersIdAgentDeployment(ctx, id).Execute()

Get cluster agent deployment progress



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
	resp, r, err := apiClient.ClustersAPI.GetClustersIdAgentDeployment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdAgentDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAgentDeployment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdAgentDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAgentDeploymentRequest struct via the builder pattern


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


## GetClustersIdAgentDeploymentWs

> GetClustersIdAgentDeploymentWs(ctx, id).Execute()

WebSocket for cluster agent deployment progress



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
	r, err := apiClient.ClustersAPI.GetClustersIdAgentDeploymentWs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdAgentDeploymentWs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAgentDeploymentWsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetClustersIdEvents

> map[string]interface{} GetClustersIdEvents(ctx, id).Cursor(cursor).Limit(limit).Execute()

Get cluster status transition history



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
	cursor := "cursor_example" // string | Exclusive lower bound: return transitions with id > cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClustersIdEvents(context.Background(), id).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdEvents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Exclusive lower bound: return transitions with id &gt; cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 

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


## GetClustersIdExplorerInfo

> map[string]interface{} GetClustersIdExplorerInfo(ctx, id).Execute()

Get K8s dashboard access information



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClustersIdExplorerInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdExplorerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdExplorerInfo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdExplorerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdExplorerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [EncryptionKeyAuth](../README.md#EncryptionKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdNetwork

> NetworkbackendNetwork GetClustersIdNetwork(ctx, id).Execute()

Get cluster network



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClustersIdNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdNetwork`: NetworkbackendNetwork
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkbackendNetwork**](NetworkbackendNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdVirtualMachinesNew

> map[string]interface{} GetClustersIdVirtualMachinesNew(ctx, id).Execute()

Get available VMs to add to cluster



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClustersIdVirtualMachinesNew(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdVirtualMachinesNew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdVirtualMachinesNew`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdVirtualMachinesNew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdVirtualMachinesNewRequest struct via the builder pattern


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


## GetClustersNew

> map[string]interface{} GetClustersNew(ctx).Provider(provider).Execute()

Get cluster creation options



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
	provider := "provider_example" // string | Provider type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClustersNew(context.Background()).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersNew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersNew`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersNew`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Provider type | 

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


## PostClusters

> ModelsCluster PostClusters(ctx).Request(request).Execute()

Create a new cluster



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
	request := *openapiclient.NewHandlersCreateClusterRequest(*openapiclient.NewHandlersCreateClusterRequestData("Stack_example", "Version_example"), "Name_example", "Provider_example") // HandlersCreateClusterRequest | Cluster creation details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.PostClusters(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClusters`: ModelsCluster
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersCreateClusterRequest**](HandlersCreateClusterRequest.md) | Cluster creation details | 

### Return type

[**ModelsCluster**](ModelsCluster.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdExplorerSetup

> map[string]interface{} PostClustersIdExplorerSetup(ctx, id).Execute()

Setup K8s cluster explorer tunnel



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.PostClustersIdExplorerSetup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PostClustersIdExplorerSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdExplorerSetup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.PostClustersIdExplorerSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdExplorerSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [EncryptionKeyAuth](../README.md#EncryptionKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdNetwork

> NetworkbackendNetwork PostClustersIdNetwork(ctx, id).Request(request).Execute()

Set cluster network



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
	request := *openapiclient.NewHandlersSetClusterNetworkRequest() // HandlersSetClusterNetworkRequest | Network set request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.PostClustersIdNetwork(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PostClustersIdNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdNetwork`: NetworkbackendNetwork
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.PostClustersIdNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersSetClusterNetworkRequest**](HandlersSetClusterNetworkRequest.md) | Network set request | 

### Return type

[**NetworkbackendNetwork**](NetworkbackendNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdVirtualMachines

> map[string]interface{} PostClustersIdVirtualMachines(ctx, id).Request(request).Execute()

Add VMs to cluster



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
	request := *openapiclient.NewHandlersAddVirtualMachinesToClusterRequest("Role_example", []int32{int32(123)}) // HandlersAddVirtualMachinesToClusterRequest | VMs to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.PostClustersIdVirtualMachines(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PostClustersIdVirtualMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdVirtualMachines`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.PostClustersIdVirtualMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersAddVirtualMachinesToClusterRequest**](HandlersAddVirtualMachinesToClusterRequest.md) | VMs to add | 

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


## PostClustersImport

> HandlersImportClusterResponse PostClustersImport(ctx).Request(request).Execute()

Import an existing cluster (unified)



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
	request := *openapiclient.NewHandlersImportClusterRequest("Name_example") // HandlersImportClusterRequest | Import cluster details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.PostClustersImport(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PostClustersImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersImport`: HandlersImportClusterResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.PostClustersImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersImportClusterRequest**](HandlersImportClusterRequest.md) | Import cluster details | 

### Return type

[**HandlersImportClusterResponse**](HandlersImportClusterResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersRoleRecommendations

> ServicesRoleRecommendationResponse PostClustersRoleRecommendations(ctx).Request(request).Execute()

Recommend cluster roles for a VM selection



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
	request := *openapiclient.NewHandlersRoleRecommendationPayload([]int32{int32(123)}) // HandlersRoleRecommendationPayload | Network choice and candidate VM ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.PostClustersRoleRecommendations(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PostClustersRoleRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersRoleRecommendations`: ServicesRoleRecommendationResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.PostClustersRoleRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersRoleRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersRoleRecommendationPayload**](HandlersRoleRecommendationPayload.md) | Network choice and candidate VM ids | 

### Return type

[**ServicesRoleRecommendationResponse**](ServicesRoleRecommendationResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

