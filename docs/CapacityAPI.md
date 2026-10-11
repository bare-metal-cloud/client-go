# \CapacityAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudProvidersProviderOptions**](CapacityAPI.md#GetCloudProvidersProviderOptions) | **Get** /cloud-providers/{provider}/options | Get provider-specific options (regions, instance types)
[**GetClustersIdAvailableVms**](CapacityAPI.md#GetClustersIdAvailableVms) | **Get** /clusters/{id}/available-vms | Get available VMs for cluster
[**GetClustersIdCapacitySources**](CapacityAPI.md#GetClustersIdCapacitySources) | **Get** /clusters/{id}/capacity-sources | Get available capacity sources for a cluster
[**GetVirtualMachinesUnassigned**](CapacityAPI.md#GetVirtualMachinesUnassigned) | **Get** /virtual-machines/unassigned | Get VMs not assigned to any cluster
[**PostClustersIdVms**](CapacityAPI.md#PostClustersIdVms) | **Post** /clusters/{id}/vms | Assign VMs to a cluster



## GetCloudProvidersProviderOptions

> map[string]interface{} GetCloudProvidersProviderOptions(ctx, provider).Execute()

Get provider-specific options (regions, instance types)



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
	provider := "provider_example" // string | Provider (aws, gcp, azure, digitalocean)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CapacityAPI.GetCloudProvidersProviderOptions(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapacityAPI.GetCloudProvidersProviderOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersProviderOptions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CapacityAPI.GetCloudProvidersProviderOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider (aws, gcp, azure, digitalocean) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersProviderOptionsRequest struct via the builder pattern


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


## GetClustersIdAvailableVms

> []ModelsVirtualMachine GetClustersIdAvailableVms(ctx, id).Execute()

Get available VMs for cluster



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
	resp, r, err := apiClient.CapacityAPI.GetClustersIdAvailableVms(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapacityAPI.GetClustersIdAvailableVms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAvailableVms`: []ModelsVirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `CapacityAPI.GetClustersIdAvailableVms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAvailableVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelsVirtualMachine**](ModelsVirtualMachine.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdCapacitySources

> map[string]interface{} GetClustersIdCapacitySources(ctx, id).Execute()

Get available capacity sources for a cluster



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
	resp, r, err := apiClient.CapacityAPI.GetClustersIdCapacitySources(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapacityAPI.GetClustersIdCapacitySources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdCapacitySources`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CapacityAPI.GetClustersIdCapacitySources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdCapacitySourcesRequest struct via the builder pattern


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


## GetVirtualMachinesUnassigned

> map[string]interface{} GetVirtualMachinesUnassigned(ctx).Page(page).Limit(limit).Execute()

Get VMs not assigned to any cluster



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
	page := int32(56) // int32 | Page number (default 1) (optional)
	limit := int32(56) // int32 | Items per page (default 50, max 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CapacityAPI.GetVirtualMachinesUnassigned(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapacityAPI.GetVirtualMachinesUnassigned``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesUnassigned`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CapacityAPI.GetVirtualMachinesUnassigned`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesUnassignedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (default 1) | 
 **limit** | **int32** | Items per page (default 50, max 100) | 

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


## PostClustersIdVms

> []ModelsClusterNode PostClustersIdVms(ctx, id).Request(request).Execute()

Assign VMs to a cluster



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
	request := *openapiclient.NewHandlersAssignVMsRequest([]int32{int32(123)}) // HandlersAssignVMsRequest | VMs and roles

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CapacityAPI.PostClustersIdVms(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapacityAPI.PostClustersIdVms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdVms`: []ModelsClusterNode
	fmt.Fprintf(os.Stdout, "Response from `CapacityAPI.PostClustersIdVms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersAssignVMsRequest**](HandlersAssignVMsRequest.md) | VMs and roles | 

### Return type

[**[]ModelsClusterNode**](ModelsClusterNode.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

