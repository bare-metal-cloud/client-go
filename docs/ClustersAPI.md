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
[**GetClustersIdBackupMigration**](ClustersAPI.md#GetClustersIdBackupMigration) | **Get** /clusters/{id}/backup-migration | Get the cluster&#39;s backup destination migration status
[**GetClustersIdBackupRuns**](ClustersAPI.md#GetClustersIdBackupRuns) | **Get** /clusters/{id}/backup-runs | List the cluster backup plan&#39;s runs
[**GetClustersIdBackupStatus**](ClustersAPI.md#GetClustersIdBackupStatus) | **Get** /clusters/{id}/backup-status | Get the cluster backup plan&#39;s status
[**GetClustersIdDeleteImpact**](ClustersAPI.md#GetClustersIdDeleteImpact) | **Get** /clusters/{id}/delete-impact | Preflight the consequences of deleting a platform
[**GetClustersIdEdgeHostname**](ClustersAPI.md#GetClustersIdEdgeHostname) | **Get** /clusters/{id}/edge-hostname | The cluster&#39;s platform-minted edge hostname
[**GetClustersIdEvents**](ClustersAPI.md#GetClustersIdEvents) | **Get** /clusters/{id}/events | Get cluster status transition history
[**GetClustersIdExplorerInfo**](ClustersAPI.md#GetClustersIdExplorerInfo) | **Get** /clusters/{id}/explorer/info | Get K8s dashboard access information
[**GetClustersIdK8sViewAccess**](ClustersAPI.md#GetClustersIdK8sViewAccess) | **Get** /clusters/{id}/k8s-view/access | K8s View access state (the teaching boundary shape)
[**GetClustersIdK8sWatchProxyPath**](ClustersAPI.md#GetClustersIdK8sWatchProxyPath) | **Get** /clusters/{id}/k8s-watch/{proxyPath} | Watch K8s resources via the explorer watch protocol
[**GetClustersIdLatency**](ClustersAPI.md#GetClustersIdLatency) | **Get** /clusters/{id}/latency | Real-time latency between the cluster&#39;s nodes
[**GetClustersIdMonitorStatus**](ClustersAPI.md#GetClustersIdMonitorStatus) | **Get** /clusters/{id}/monitor-status | Get a cluster&#39;s monitor-context status and composition
[**GetClustersIdNetwork**](ClustersAPI.md#GetClustersIdNetwork) | **Get** /clusters/{id}/network | Get cluster network
[**GetClustersIdTimeline**](ClustersAPI.md#GetClustersIdTimeline) | **Get** /clusters/{id}/timeline | Get cluster resource timeline
[**GetClustersIdVirtualMachinesNew**](ClustersAPI.md#GetClustersIdVirtualMachinesNew) | **Get** /clusters/{id}/virtual-machines/new | Get available VMs to add to cluster
[**GetClustersIdVirtualMachinesVmIdRemoveImpact**](ClustersAPI.md#GetClustersIdVirtualMachinesVmIdRemoveImpact) | **Get** /clusters/{id}/virtual-machines/{vm_id}/remove-impact | Preflight the consequences of removing a VM from its platform
[**GetClustersNew**](ClustersAPI.md#GetClustersNew) | **Get** /clusters/new | Get cluster creation options
[**PatchClustersIdBackupPlan**](ClustersAPI.md#PatchClustersIdBackupPlan) | **Patch** /clusters/{id}/backup-plan | Update the cluster&#39;s backup plan
[**PostClusters**](ClustersAPI.md#PostClusters) | **Post** /clusters | Create a new cluster
[**PostClustersIdBackupMigrationResume**](ClustersAPI.md#PostClustersIdBackupMigrationResume) | **Post** /clusters/{id}/backup-migration/resume | Resume a failed backup destination migration
[**PostClustersIdBackupRunsRunIdRestore**](ClustersAPI.md#PostClustersIdBackupRunsRunIdRestore) | **Post** /clusters/{id}/backup-runs/{run_id}/restore | Restore an install&#39;s data from a recorded backup run
[**PostClustersIdExplorerSetup**](ClustersAPI.md#PostClustersIdExplorerSetup) | **Post** /clusters/{id}/explorer/setup | Setup K8s cluster explorer tunnel
[**PostClustersIdK8sViewAccessRequest**](ClustersAPI.md#PostClustersIdK8sViewAccessRequest) | **Post** /clusters/{id}/k8s-view/access/request | Request K8s View access (disabled at launch)
[**PostClustersIdNetwork**](ClustersAPI.md#PostClustersIdNetwork) | **Post** /clusters/{id}/network | Set cluster network
[**PostClustersIdVirtualMachines**](ClustersAPI.md#PostClustersIdVirtualMachines) | **Post** /clusters/{id}/virtual-machines | Add VMs to cluster
[**PostClustersImport**](ClustersAPI.md#PostClustersImport) | **Post** /clusters/import | Import an existing cluster (unified)
[**PostClustersPeerProbe**](ClustersAPI.md#PostClustersPeerProbe) | **Post** /clusters/peer-probe | Briefly test connections between the selected machines
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


## GetClustersIdBackupMigration

> map[string]interface{} GetClustersIdBackupMigration(ctx, id).Execute()

Get the cluster's backup destination migration status



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
	resp, r, err := apiClient.ClustersAPI.GetClustersIdBackupMigration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdBackupMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdBackupMigration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdBackupMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdBackupMigrationRequest struct via the builder pattern


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


## GetClustersIdBackupRuns

> map[string]interface{} GetClustersIdBackupRuns(ctx, id).Limit(limit).Execute()

List the cluster backup plan's runs



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
	limit := int32(56) // int32 | Page size (default 100, max 500) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClustersIdBackupRuns(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdBackupRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdBackupRuns`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdBackupRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdBackupRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Page size (default 100, max 500) | 

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


## GetClustersIdBackupStatus

> map[string]interface{} GetClustersIdBackupStatus(ctx, id).Execute()

Get the cluster backup plan's status



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
	resp, r, err := apiClient.ClustersAPI.GetClustersIdBackupStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdBackupStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdBackupStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdBackupStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdBackupStatusRequest struct via the builder pattern


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


## GetClustersIdDeleteImpact

> ServicesImpactResponse GetClustersIdDeleteImpact(ctx, id).Execute()

Preflight the consequences of deleting a platform



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
	resp, r, err := apiClient.ClustersAPI.GetClustersIdDeleteImpact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdDeleteImpact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdDeleteImpact`: ServicesImpactResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdDeleteImpact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdDeleteImpactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicesImpactResponse**](ServicesImpactResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdEdgeHostname

> HandlersClusterEdgeHostnameView GetClustersIdEdgeHostname(ctx, id).Execute()

The cluster's platform-minted edge hostname



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
	resp, r, err := apiClient.ClustersAPI.GetClustersIdEdgeHostname(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdEdgeHostname``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdEdgeHostname`: HandlersClusterEdgeHostnameView
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdEdgeHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdEdgeHostnameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersClusterEdgeHostnameView**](HandlersClusterEdgeHostnameView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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


## GetClustersIdK8sViewAccess

> HandlersExplorerAccessView GetClustersIdK8sViewAccess(ctx, id).Execute()

K8s View access state (the teaching boundary shape)



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
	resp, r, err := apiClient.ClustersAPI.GetClustersIdK8sViewAccess(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdK8sViewAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdK8sViewAccess`: HandlersExplorerAccessView
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdK8sViewAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdK8sViewAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersExplorerAccessView**](HandlersExplorerAccessView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdK8sWatchProxyPath

> GetClustersIdK8sWatchProxyPath(ctx, id, proxyPath).Execute()

Watch K8s resources via the explorer watch protocol



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
	proxyPath := "proxyPath_example" // string | Path segment (backfilled by normalize-spec)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClustersAPI.GetClustersIdK8sWatchProxyPath(context.Background(), id, proxyPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdK8sWatchProxyPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**proxyPath** | **string** | Path segment (backfilled by normalize-spec) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdK8sWatchProxyPathRequest struct via the builder pattern


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


## GetClustersIdLatency

> ServicesClusterLatencyDashboard GetClustersIdLatency(ctx, id).Execute()

Real-time latency between the cluster's nodes



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
	resp, r, err := apiClient.ClustersAPI.GetClustersIdLatency(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdLatency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdLatency`: ServicesClusterLatencyDashboard
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdLatency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdLatencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicesClusterLatencyDashboard**](ServicesClusterLatencyDashboard.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdMonitorStatus

> map[string]interface{} GetClustersIdMonitorStatus(ctx, id).Execute()

Get a cluster's monitor-context status and composition



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
	resp, r, err := apiClient.ClustersAPI.GetClustersIdMonitorStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdMonitorStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdMonitorStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdMonitorStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdMonitorStatusRequest struct via the builder pattern


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


## GetClustersIdTimeline

> map[string]interface{} GetClustersIdTimeline(ctx, id).Cursor(cursor).Limit(limit).ResourceType(resourceType).ResourceRef(resourceRef).Action(action).ActorType(actorType).From(from).To(to).Execute()

Get cluster resource timeline



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
	cursor := "cursor_example" // string | Opaque keyset cursor (minted by this route) (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	resourceType := "resourceType_example" // string | Filter by resource_type (exact) (optional)
	resourceRef := "resourceRef_example" // string | Filter by resource_ref (exact) (optional)
	action := "action_example" // string | Filter by action (exact) (optional)
	actorType := "actorType_example" // string | Filter by actor_type: user or system (optional)
	from := "from_example" // string | Inclusive lower bound on created_at (RFC3339) (optional)
	to := "to_example" // string | Inclusive upper bound on created_at (RFC3339) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClustersIdTimeline(context.Background(), id).Cursor(cursor).Limit(limit).ResourceType(resourceType).ResourceRef(resourceRef).Action(action).ActorType(actorType).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdTimeline`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque keyset cursor (minted by this route) | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **resourceType** | **string** | Filter by resource_type (exact) | 
 **resourceRef** | **string** | Filter by resource_ref (exact) | 
 **action** | **string** | Filter by action (exact) | 
 **actorType** | **string** | Filter by actor_type: user or system | 
 **from** | **string** | Inclusive lower bound on created_at (RFC3339) | 
 **to** | **string** | Inclusive upper bound on created_at (RFC3339) | 

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


## GetClustersIdVirtualMachinesVmIdRemoveImpact

> ServicesImpactResponse GetClustersIdVirtualMachinesVmIdRemoveImpact(ctx, id, vmId).Execute()

Preflight the consequences of removing a VM from its platform



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
	vmId := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.GetClustersIdVirtualMachinesVmIdRemoveImpact(context.Background(), id, vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.GetClustersIdVirtualMachinesVmIdRemoveImpact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdVirtualMachinesVmIdRemoveImpact`: ServicesImpactResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.GetClustersIdVirtualMachinesVmIdRemoveImpact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**vmId** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdVirtualMachinesVmIdRemoveImpactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServicesImpactResponse**](ServicesImpactResponse.md)

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


## PatchClustersIdBackupPlan

> map[string]interface{} PatchClustersIdBackupPlan(ctx, id).Body(body).Execute()

Update the cluster's backup plan



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
	body := *openapiclient.NewHandlersUpdateBackupPlanBody() // HandlersUpdateBackupPlanBody | Plan fields to change (all optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.PatchClustersIdBackupPlan(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PatchClustersIdBackupPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchClustersIdBackupPlan`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.PatchClustersIdBackupPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchClustersIdBackupPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersUpdateBackupPlanBody**](HandlersUpdateBackupPlanBody.md) | Plan fields to change (all optional) | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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


## PostClustersIdBackupMigrationResume

> map[string]interface{} PostClustersIdBackupMigrationResume(ctx, id).Execute()

Resume a failed backup destination migration



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
	resp, r, err := apiClient.ClustersAPI.PostClustersIdBackupMigrationResume(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PostClustersIdBackupMigrationResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdBackupMigrationResume`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.PostClustersIdBackupMigrationResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdBackupMigrationResumeRequest struct via the builder pattern


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


## PostClustersIdBackupRunsRunIdRestore

> map[string]interface{} PostClustersIdBackupRunsRunIdRestore(ctx, id, runId).Body(body).Execute()

Restore an install's data from a recorded backup run



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
	runId := int32(56) // int32 | Backup run ID (the restore point)
	body := *openapiclient.NewHandlersBackupRestoreRequest() // HandlersBackupRestoreRequest | The fresh confirm (empty on the first POST — the dialog answer)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.PostClustersIdBackupRunsRunIdRestore(context.Background(), id, runId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PostClustersIdBackupRunsRunIdRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdBackupRunsRunIdRestore`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.PostClustersIdBackupRunsRunIdRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**runId** | **int32** | Backup run ID (the restore point) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdBackupRunsRunIdRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersBackupRestoreRequest**](HandlersBackupRestoreRequest.md) | The fresh confirm (empty on the first POST — the dialog answer) | 

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


## PostClustersIdK8sViewAccessRequest

> PostClustersIdK8sViewAccessRequest(ctx, id).Body(body).Execute()

Request K8s View access (disabled at launch)



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
	body := map[string]interface{}{ ... } // map[string]interface{} | Ignored (nothing is filed) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClustersAPI.PostClustersIdK8sViewAccessRequest(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PostClustersIdK8sViewAccessRequest``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostClustersIdK8sViewAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Ignored (nothing is filed) | 

### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
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

> HandlersImportClusterResponse PostClustersImport(ctx).Request(request).XVaultRentalID(xVaultRentalID).Execute()

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
	xVaultRentalID := "xVaultRentalID_example" // string | Rental id (kubeconfig path only; the cluster.import.seal rental minted for this import) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.PostClustersImport(context.Background()).Request(request).XVaultRentalID(xVaultRentalID).Execute()
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
 **xVaultRentalID** | **string** | Rental id (kubeconfig path only; the cluster.import.seal rental minted for this import) | 

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


## PostClustersPeerProbe

> ServicesPeerProbeResponse PostClustersPeerProbe(ctx).Request(request).Execute()

Briefly test connections between the selected machines



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
	request := *openapiclient.NewHandlersPeerProbePayload([]int32{int32(123)}) // HandlersPeerProbePayload | Selection, probe plane and network

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClustersAPI.PostClustersPeerProbe(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClustersAPI.PostClustersPeerProbe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersPeerProbe`: ServicesPeerProbeResponse
	fmt.Fprintf(os.Stdout, "Response from `ClustersAPI.PostClustersPeerProbe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersPeerProbeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersPeerProbePayload**](HandlersPeerProbePayload.md) | Selection, probe plane and network | 

### Return type

[**ServicesPeerProbeResponse**](ServicesPeerProbeResponse.md)

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

