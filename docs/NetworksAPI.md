# \NetworksAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetworksId**](NetworksAPI.md#DeleteNetworksId) | **Delete** /networks/{id} | Delete a network
[**DeleteNetworksIdVmsVmId**](NetworksAPI.md#DeleteNetworksIdVmsVmId) | **Delete** /networks/{id}/vms/{vm_id} | Remove VM from network
[**GetNetworks**](NetworksAPI.md#GetNetworks) | **Get** /networks | List networks
[**GetNetworksId**](NetworksAPI.md#GetNetworksId) | **Get** /networks/{id} | Get network details
[**GetNetworksIdCandidates**](NetworksAPI.md#GetNetworksIdCandidates) | **Get** /networks/{id}/candidates | List candidate VMs for a network
[**GetNetworksIdJoinStatus**](NetworksAPI.md#GetNetworksIdJoinStatus) | **Get** /networks/{id}/join-status | Get network join status
[**GetNetworksIdNodes**](NetworksAPI.md#GetNetworksIdNodes) | **Get** /networks/{id}/nodes | List network nodes
[**PostNetworks**](NetworksAPI.md#PostNetworks) | **Post** /networks | Create a new network
[**PostNetworksIdDraftIntentDismiss**](NetworksAPI.md#PostNetworksIdDraftIntentDismiss) | **Post** /networks/{id}/draft-intent/dismiss | Dismiss a network&#39;s cluster-create draft intent
[**PostNetworksIdVms**](NetworksAPI.md#PostNetworksIdVms) | **Post** /networks/{id}/vms | Add VM to network
[**PostNetworksIdVmsVmIdCancel**](NetworksAPI.md#PostNetworksIdVmsVmIdCancel) | **Post** /networks/{id}/vms/{vm_id}/cancel | Cancel in-flight VM-to-network join
[**PostNetworksIdVmsVmIdRetry**](NetworksAPI.md#PostNetworksIdVmsVmIdRetry) | **Post** /networks/{id}/vms/{vm_id}/retry | Retry failed VM-to-network join



## DeleteNetworksId

> map[string]string DeleteNetworksId(ctx, id).Execute()

Delete a network



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
	id := "id_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworksId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworksId`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworksIdVmsVmId

> map[string]string DeleteNetworksIdVmsVmId(ctx, id, vmId).Execute()

Remove VM from network



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
	id := "id_example" // string | Network ID
	vmId := "vmId_example" // string | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworksIdVmsVmId(context.Background(), id, vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworksIdVmsVmId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworksIdVmsVmId`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworksIdVmsVmId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network ID | 
**vmId** | **string** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworksIdVmsVmIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworks

> []NetworkbackendNetwork GetNetworks(ctx).Execute()

List networks



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
	resp, r, err := apiClient.NetworksAPI.GetNetworks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworks`: []NetworkbackendNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksRequest struct via the builder pattern


### Return type

[**[]NetworkbackendNetwork**](NetworkbackendNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworksId

> NetworkbackendNetwork GetNetworksId(ctx, id).Execute()

Get network details



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
	id := "id_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworksId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworksId`: NetworkbackendNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksIdRequest struct via the builder pattern


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


## GetNetworksIdCandidates

> NetworkbackendCandidatesResult GetNetworksIdCandidates(ctx, id).Context(context).Execute()

List candidate VMs for a network



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
	id := "id_example" // string | Network ID
	context := "context_example" // string | Caller context (cluster-create | standalone | add-node) (optional) (default to "cluster-create")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworksIdCandidates(context.Background(), id).Context(context).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworksIdCandidates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworksIdCandidates`: NetworkbackendCandidatesResult
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworksIdCandidates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksIdCandidatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **context** | **string** | Caller context (cluster-create | standalone | add-node) | [default to &quot;cluster-create&quot;]

### Return type

[**NetworkbackendCandidatesResult**](NetworkbackendCandidatesResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworksIdJoinStatus

> NetworkbackendNetworkJoinStatusResult GetNetworksIdJoinStatus(ctx, id).Execute()

Get network join status



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
	id := "id_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworksIdJoinStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworksIdJoinStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworksIdJoinStatus`: NetworkbackendNetworkJoinStatusResult
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworksIdJoinStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksIdJoinStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkbackendNetworkJoinStatusResult**](NetworkbackendNetworkJoinStatusResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworksIdNodes

> map[string]interface{} GetNetworksIdNodes(ctx, id).Execute()

List network nodes



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
	id := "id_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworksIdNodes(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworksIdNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworksIdNodes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworksIdNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksIdNodesRequest struct via the builder pattern


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


## PostNetworks

> NetworkbackendNetwork PostNetworks(ctx).Request(request).Execute()

Create a new network



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
	request := *openapiclient.NewHandlersCreateNetworkRequest() // HandlersCreateNetworkRequest | Network creation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.PostNetworks(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.PostNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNetworks`: NetworkbackendNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.PostNetworks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersCreateNetworkRequest**](HandlersCreateNetworkRequest.md) | Network creation request | 

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


## PostNetworksIdDraftIntentDismiss

> map[string]string PostNetworksIdDraftIntentDismiss(ctx, id).Execute()

Dismiss a network's cluster-create draft intent



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
	id := "id_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.PostNetworksIdDraftIntentDismiss(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.PostNetworksIdDraftIntentDismiss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNetworksIdDraftIntentDismiss`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.PostNetworksIdDraftIntentDismiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNetworksIdDraftIntentDismissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNetworksIdVms

> map[string]string PostNetworksIdVms(ctx, id).Request(request).Execute()

Add VM to network



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
	id := "id_example" // string | Network ID
	request := *openapiclient.NewHandlersAddVMToNetworkRequest() // HandlersAddVMToNetworkRequest | VM add request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.PostNetworksIdVms(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.PostNetworksIdVms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNetworksIdVms`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.PostNetworksIdVms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNetworksIdVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersAddVMToNetworkRequest**](HandlersAddVMToNetworkRequest.md) | VM add request | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNetworksIdVmsVmIdCancel

> map[string]string PostNetworksIdVmsVmIdCancel(ctx, id, vmId).Execute()

Cancel in-flight VM-to-network join



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
	id := "id_example" // string | Network ID
	vmId := "vmId_example" // string | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.PostNetworksIdVmsVmIdCancel(context.Background(), id, vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.PostNetworksIdVmsVmIdCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNetworksIdVmsVmIdCancel`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.PostNetworksIdVmsVmIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network ID | 
**vmId** | **string** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNetworksIdVmsVmIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNetworksIdVmsVmIdRetry

> map[string]string PostNetworksIdVmsVmIdRetry(ctx, id, vmId).Execute()

Retry failed VM-to-network join



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
	id := "id_example" // string | Network ID
	vmId := "vmId_example" // string | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.PostNetworksIdVmsVmIdRetry(context.Background(), id, vmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.PostNetworksIdVmsVmIdRetry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNetworksIdVmsVmIdRetry`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.PostNetworksIdVmsVmIdRetry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Network ID | 
**vmId** | **string** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNetworksIdVmsVmIdRetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

