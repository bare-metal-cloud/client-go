# \VirtualMachinesAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVirtualMachinesId**](VirtualMachinesAPI.md#DeleteVirtualMachinesId) | **Delete** /virtual-machines/{id} | Delete a VM
[**GetVirtualMachines**](VirtualMachinesAPI.md#GetVirtualMachines) | **Get** /virtual-machines | List my virtual machines
[**GetVirtualMachinesId**](VirtualMachinesAPI.md#GetVirtualMachinesId) | **Get** /virtual-machines/{id} | Get one virtual machine
[**GetVirtualMachinesIdAddressChangeImpact**](VirtualMachinesAPI.md#GetVirtualMachinesIdAddressChangeImpact) | **Get** /virtual-machines/{id}/address-change-impact | Preflight the machine&#39;s detected public-address change
[**GetVirtualMachinesIdAddressChangeStatus**](VirtualMachinesAPI.md#GetVirtualMachinesIdAddressChangeStatus) | **Get** /virtual-machines/{id}/address-change/status | Re-check the machine&#39;s address change
[**GetVirtualMachinesIdConnections**](VirtualMachinesAPI.md#GetVirtualMachinesIdConnections) | **Get** /virtual-machines/{id}/connections | Get a VM&#39;s agent connection event history
[**GetVirtualMachinesIdContextTransitions**](VirtualMachinesAPI.md#GetVirtualMachinesIdContextTransitions) | **Get** /virtual-machines/{id}/context-transitions | Get a VM&#39;s workload-context transition history
[**GetVirtualMachinesIdContexts**](VirtualMachinesAPI.md#GetVirtualMachinesIdContexts) | **Get** /virtual-machines/{id}/contexts | Get a VM&#39;s current workload contexts
[**GetVirtualMachinesIdDeleteImpact**](VirtualMachinesAPI.md#GetVirtualMachinesIdDeleteImpact) | **Get** /virtual-machines/{id}/delete-impact | Preflight the consequences of deleting a VM
[**GetVirtualMachinesIdTransitions**](VirtualMachinesAPI.md#GetVirtualMachinesIdTransitions) | **Get** /virtual-machines/{id}/transitions | Get a VM&#39;s liveness transition history
[**GetVirtualMachinesIdUptime**](VirtualMachinesAPI.md#GetVirtualMachinesIdUptime) | **Get** /virtual-machines/{id}/uptime | Get a VM&#39;s computed availability over a window
[**GetVirtualMachinesJobs**](VirtualMachinesAPI.md#GetVirtualMachinesJobs) | **Get** /virtual-machines/jobs | List VM operations
[**GetVirtualMachinesJobsId**](VirtualMachinesAPI.md#GetVirtualMachinesJobsId) | **Get** /virtual-machines/jobs/{id} | Get VM operation
[**GetVmsIdInstallStatus**](VirtualMachinesAPI.md#GetVmsIdInstallStatus) | **Get** /vms/{id}/install-status | Get the agent install status of a VM
[**PostCloudProvidersAzureProvision**](VirtualMachinesAPI.md#PostCloudProvidersAzureProvision) | **Post** /cloud-providers/azure/provision | Provision an Azure VM
[**PostVirtualMachines**](VirtualMachinesAPI.md#PostVirtualMachines) | **Post** /virtual-machines | Create a virtual machine
[**PostVirtualMachinesIdAddressChangeExecute**](VirtualMachinesAPI.md#PostVirtualMachinesIdAddressChangeExecute) | **Post** /virtual-machines/{id}/address-change/execute | Confirm the machine&#39;s address change and act per custody
[**PostVirtualMachinesIdCheckK3s**](VirtualMachinesAPI.md#PostVirtualMachinesIdCheckK3s) | **Post** /virtual-machines/{id}/check-k3s | Check a VM for pre-existing k3s
[**PostVirtualMachinesIdRestore**](VirtualMachinesAPI.md#PostVirtualMachinesIdRestore) | **Post** /virtual-machines/{id}/restore | Restore a falsely terminated VM record
[**PostVirtualMachinesIdTransfer**](VirtualMachinesAPI.md#PostVirtualMachinesIdTransfer) | **Post** /virtual-machines/{id}/transfer | Transfer a VM to another project of the same org
[**PostVirtualMachinesImport**](VirtualMachinesAPI.md#PostVirtualMachinesImport) | **Post** /virtual-machines/import | Import an existing machine
[**PostVirtualMachinesImportManual**](VirtualMachinesAPI.md#PostVirtualMachinesImportManual) | **Post** /virtual-machines/import/manual | Get the manual install command for a machine
[**PostVirtualMachinesImportSsh**](VirtualMachinesAPI.md#PostVirtualMachinesImportSsh) | **Post** /virtual-machines/import/ssh | Import a machine with SSH auto-push
[**PostVirtualMachinesProvisionAws**](VirtualMachinesAPI.md#PostVirtualMachinesProvisionAws) | **Post** /virtual-machines/provision/aws | Provision an AWS EC2 instance
[**PostVirtualMachinesProvisionAwsWorker**](VirtualMachinesAPI.md#PostVirtualMachinesProvisionAwsWorker) | **Post** /virtual-machines/provision/aws/worker | Provision an AWS EC2 instance via background worker
[**PostVirtualMachinesProvisionGcp**](VirtualMachinesAPI.md#PostVirtualMachinesProvisionGcp) | **Post** /virtual-machines/provision/gcp | Provision a GCP Compute Engine instance



## DeleteVirtualMachinesId

> DeleteVirtualMachinesId200Response DeleteVirtualMachinesId(ctx, id).Force(force).ImpactFingerprint(impactFingerprint).Execute()

Delete a VM



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
	id := int32(56) // int32 | VM ID
	force := true // bool | Force-remove immediately (requires admin or owner; only meaningful for deleting / cleanup_incomplete / deletion_failed VMs) (optional)
	impactFingerprint := "impactFingerprint_example" // string | Epic #3964: the fingerprint GET /virtual-machines/{id}/delete-impact returned. When provided and the consequences moved since that preflight, the delete refuses with 409 consequence_changed carrying the fresh impact. Absent = legacy behavior. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.DeleteVirtualMachinesId(context.Background(), id).Force(force).ImpactFingerprint(impactFingerprint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.DeleteVirtualMachinesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVirtualMachinesId`: DeleteVirtualMachinesId200Response
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.DeleteVirtualMachinesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualMachinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force-remove immediately (requires admin or owner; only meaningful for deleting / cleanup_incomplete / deletion_failed VMs) | 
 **impactFingerprint** | **string** | Epic #3964: the fingerprint GET /virtual-machines/{id}/delete-impact returned. When provided and the consequences moved since that preflight, the delete refuses with 409 consequence_changed carrying the fresh impact. Absent &#x3D; legacy behavior. | 

### Return type

[**DeleteVirtualMachinesId200Response**](DeleteVirtualMachinesId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachines

> []map[string]interface{} GetVirtualMachines(ctx).Execute()

List my virtual machines

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
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachines`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesRequest struct via the builder pattern


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachinesId

> ModelsVirtualMachine GetVirtualMachinesId(ctx, id).Execute()

Get one virtual machine

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
	id := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesId`: ModelsVirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsVirtualMachine**](ModelsVirtualMachine.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachinesIdAddressChangeImpact

> AddresschangeImpact GetVirtualMachinesIdAddressChangeImpact(ctx, id).Execute()

Preflight the machine's detected public-address change



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
	id := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesIdAddressChangeImpact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesIdAddressChangeImpact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdAddressChangeImpact`: AddresschangeImpact
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesIdAddressChangeImpact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdAddressChangeImpactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddresschangeImpact**](AddresschangeImpact.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachinesIdAddressChangeStatus

> AddresschangeStatus GetVirtualMachinesIdAddressChangeStatus(ctx, id).Execute()

Re-check the machine's address change



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
	id := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesIdAddressChangeStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesIdAddressChangeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdAddressChangeStatus`: AddresschangeStatus
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesIdAddressChangeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdAddressChangeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddresschangeStatus**](AddresschangeStatus.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachinesIdConnections

> map[string]interface{} GetVirtualMachinesIdConnections(ctx, id).Cursor(cursor).Limit(limit).Execute()

Get a VM's agent connection event history



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
	id := int32(56) // int32 | VM ID
	cursor := "cursor_example" // string | Exclusive lower bound: return events with id > cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesIdConnections(context.Background(), id).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesIdConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdConnections`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesIdConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Exclusive lower bound: return events with id &gt; cursor | 
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


## GetVirtualMachinesIdContextTransitions

> map[string]interface{} GetVirtualMachinesIdContextTransitions(ctx, id).Limit(limit).Execute()

Get a VM's workload-context transition history



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
	id := int32(56) // int32 | VM ID
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesIdContextTransitions(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesIdContextTransitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdContextTransitions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesIdContextTransitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdContextTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetVirtualMachinesIdContexts

> map[string]interface{} GetVirtualMachinesIdContexts(ctx, id).Execute()

Get a VM's current workload contexts



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
	id := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesIdContexts(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesIdContexts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdContexts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesIdContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdContextsRequest struct via the builder pattern


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


## GetVirtualMachinesIdDeleteImpact

> ServicesImpactResponse GetVirtualMachinesIdDeleteImpact(ctx, id).Execute()

Preflight the consequences of deleting a VM



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
	id := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesIdDeleteImpact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesIdDeleteImpact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdDeleteImpact`: ServicesImpactResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesIdDeleteImpact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdDeleteImpactRequest struct via the builder pattern


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


## GetVirtualMachinesIdTransitions

> map[string]interface{} GetVirtualMachinesIdTransitions(ctx, id).Cursor(cursor).Limit(limit).Execute()

Get a VM's liveness transition history



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
	id := int32(56) // int32 | VM ID
	cursor := "cursor_example" // string | Exclusive lower bound: return transitions with id > cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesIdTransitions(context.Background(), id).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesIdTransitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdTransitions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesIdTransitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdTransitionsRequest struct via the builder pattern


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


## GetVirtualMachinesIdUptime

> map[string]interface{} GetVirtualMachinesIdUptime(ctx, id).Window(window).Execute()

Get a VM's computed availability over a window



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
	id := int32(56) // int32 | VM ID
	window := "window_example" // string | Look-back window (1h, 24h, 7d, 30d; default 7d) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesIdUptime(context.Background(), id).Window(window).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesIdUptime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdUptime`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesIdUptime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdUptimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **window** | **string** | Look-back window (1h, 24h, 7d, 30d; default 7d) | 

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


## GetVirtualMachinesJobs

> []ModelsJobMetadata GetVirtualMachinesJobs(ctx).Execute()

List VM operations



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
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesJobs`: []ModelsJobMetadata
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesJobsRequest struct via the builder pattern


### Return type

[**[]ModelsJobMetadata**](ModelsJobMetadata.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachinesJobsId

> ModelsJobMetadata GetVirtualMachinesJobsId(ctx, id).Execute()

Get VM operation



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
	id := "id_example" // string | Job ID (UUID)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVirtualMachinesJobsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVirtualMachinesJobsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesJobsId`: ModelsJobMetadata
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVirtualMachinesJobsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesJobsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsJobMetadata**](ModelsJobMetadata.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVmsIdInstallStatus

> map[string]interface{} GetVmsIdInstallStatus(ctx, id).Execute()

Get the agent install status of a VM

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
	id := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.GetVmsIdInstallStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.GetVmsIdInstallStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVmsIdInstallStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.GetVmsIdInstallStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVmsIdInstallStatusRequest struct via the builder pattern


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


## PostCloudProvidersAzureProvision

> ModelsVirtualMachine PostCloudProvidersAzureProvision(ctx).Request(request).Execute()

Provision an Azure VM



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
	request := *openapiclient.NewHandlersProvisionAzureVMRequest(int32(123), "Name_example", "Region_example", "VmSize_example") // HandlersProvisionAzureVMRequest | Instance configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostCloudProvidersAzureProvision(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostCloudProvidersAzureProvision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudProvidersAzureProvision`: ModelsVirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostCloudProvidersAzureProvision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudProvidersAzureProvisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersProvisionAzureVMRequest**](HandlersProvisionAzureVMRequest.md) | Instance configuration | 

### Return type

[**ModelsVirtualMachine**](ModelsVirtualMachine.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVirtualMachines

> map[string]interface{} PostVirtualMachines(ctx).Request(request).Execute()

Create a virtual machine

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
	request := *openapiclient.NewHandlersCreateVirtualMachineRequest("Name_example", "Provider_example") // HandlersCreateVirtualMachineRequest | VM to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachines(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachines`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersCreateVirtualMachineRequest**](HandlersCreateVirtualMachineRequest.md) | VM to create | 

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


## PostVirtualMachinesIdAddressChangeExecute

> AddresschangeExecuteResult PostVirtualMachinesIdAddressChangeExecute(ctx, id).Body(body).Execute()

Confirm the machine's address change and act per custody



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
	id := int32(56) // int32 | VM ID
	body := *openapiclient.NewHandlersAddressChangeExecuteRequest() // HandlersAddressChangeExecuteRequest | The confirmed impact's fingerprint

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachinesIdAddressChangeExecute(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachinesIdAddressChangeExecute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdAddressChangeExecute`: AddresschangeExecuteResult
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachinesIdAddressChangeExecute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdAddressChangeExecuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersAddressChangeExecuteRequest**](HandlersAddressChangeExecuteRequest.md) | The confirmed impact&#39;s fingerprint | 

### Return type

[**AddresschangeExecuteResult**](AddresschangeExecuteResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVirtualMachinesIdCheckK3s

> HandlersVMCheckK3sResponse PostVirtualMachinesIdCheckK3s(ctx, id).Body(body).Execute()

Check a VM for pre-existing k3s



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
	id := int32(56) // int32 | VM ID
	body := *openapiclient.NewHandlersVMCheckK3sRequest() // HandlersVMCheckK3sRequest | Optional decision params. In the BODY: '+' in k3s tags decodes as space in a query string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachinesIdCheckK3s(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachinesIdCheckK3s``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdCheckK3s`: HandlersVMCheckK3sResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachinesIdCheckK3s`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdCheckK3sRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersVMCheckK3sRequest**](HandlersVMCheckK3sRequest.md) | Optional decision params. In the BODY: &#39;+&#39; in k3s tags decodes as space in a query string. | 

### Return type

[**HandlersVMCheckK3sResponse**](HandlersVMCheckK3sResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVirtualMachinesIdRestore

> map[string]interface{} PostVirtualMachinesIdRestore(ctx, id).Execute()

Restore a falsely terminated VM record

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
	id := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachinesIdRestore(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachinesIdRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdRestore`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachinesIdRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdRestoreRequest struct via the builder pattern


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


## PostVirtualMachinesIdTransfer

> PostVirtualMachinesIdTransfer200Response PostVirtualMachinesIdTransfer(ctx, id).Body(body).Execute()

Transfer a VM to another project of the same org



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
	id := int32(56) // int32 | VM ID
	body := *openapiclient.NewPostVirtualMachinesIdTransferRequest() // PostVirtualMachinesIdTransferRequest | Destination project id (same org)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachinesIdTransfer(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachinesIdTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdTransfer`: PostVirtualMachinesIdTransfer200Response
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachinesIdTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PostVirtualMachinesIdTransferRequest**](PostVirtualMachinesIdTransferRequest.md) | Destination project id (same org) | 

### Return type

[**PostVirtualMachinesIdTransfer200Response**](PostVirtualMachinesIdTransfer200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVirtualMachinesImport

> map[string]interface{} PostVirtualMachinesImport(ctx).Request(request).Execute()

Import an existing machine

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
	request := *openapiclient.NewHandlersImportVMRequest() // HandlersImportVMRequest | Machine to import (SSH creds optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachinesImport(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachinesImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesImport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachinesImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersImportVMRequest**](HandlersImportVMRequest.md) | Machine to import (SSH creds optional) | 

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


## PostVirtualMachinesImportManual

> HandlersUnifiedInstallCommandResponse PostVirtualMachinesImportManual(ctx).Request(request).Execute()

Get the manual install command for a machine

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
	request := *openapiclient.NewHandlersInstallCommandPostRequest("ClusterType_example") // HandlersInstallCommandPostRequest | Machine details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachinesImportManual(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachinesImportManual``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesImportManual`: HandlersUnifiedInstallCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachinesImportManual`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesImportManualRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersInstallCommandPostRequest**](HandlersInstallCommandPostRequest.md) | Machine details | 

### Return type

[**HandlersUnifiedInstallCommandResponse**](HandlersUnifiedInstallCommandResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVirtualMachinesImportSsh

> map[string]interface{} PostVirtualMachinesImportSsh(ctx).Request(request).Execute()

Import a machine with SSH auto-push

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
	request := *openapiclient.NewHandlersImportVMRequest() // HandlersImportVMRequest | Machine plus SSH connection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachinesImportSsh(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachinesImportSsh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesImportSsh`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachinesImportSsh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesImportSshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersImportVMRequest**](HandlersImportVMRequest.md) | Machine plus SSH connection | 

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


## PostVirtualMachinesProvisionAws

> ModelsVirtualMachine PostVirtualMachinesProvisionAws(ctx).Request(request).Execute()

Provision an AWS EC2 instance



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
	request := *openapiclient.NewHandlersProvisionAWSVMRequest(int32(123), "InstanceType_example", "Name_example", "Region_example") // HandlersProvisionAWSVMRequest | Instance configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachinesProvisionAws(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachinesProvisionAws``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesProvisionAws`: ModelsVirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachinesProvisionAws`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesProvisionAwsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersProvisionAWSVMRequest**](HandlersProvisionAWSVMRequest.md) | Instance configuration | 

### Return type

[**ModelsVirtualMachine**](ModelsVirtualMachine.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVirtualMachinesProvisionAwsWorker

> HandlersAWSVMWorkerResponse PostVirtualMachinesProvisionAwsWorker(ctx).Request(request).Execute()

Provision an AWS EC2 instance via background worker



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
	request := *openapiclient.NewHandlersAWSVMWorkerRequest(int32(123), int32(123), int32(123), int32(123), "Name_example", "Region_example") // HandlersAWSVMWorkerRequest | Instance configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachinesProvisionAwsWorker(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachinesProvisionAwsWorker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesProvisionAwsWorker`: HandlersAWSVMWorkerResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachinesProvisionAwsWorker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesProvisionAwsWorkerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersAWSVMWorkerRequest**](HandlersAWSVMWorkerRequest.md) | Instance configuration | 

### Return type

[**HandlersAWSVMWorkerResponse**](HandlersAWSVMWorkerResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVirtualMachinesProvisionGcp

> ModelsVirtualMachine PostVirtualMachinesProvisionGcp(ctx).Request(request).Execute()

Provision a GCP Compute Engine instance



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
	request := *openapiclient.NewHandlersProvisionGCPVMRequest(int32(123), "MachineType_example", "Name_example", "Zone_example") // HandlersProvisionGCPVMRequest | Instance configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.PostVirtualMachinesProvisionGcp(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVirtualMachinesProvisionGcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesProvisionGcp`: ModelsVirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVirtualMachinesProvisionGcp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesProvisionGcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersProvisionGCPVMRequest**](HandlersProvisionGCPVMRequest.md) | Instance configuration | 

### Return type

[**ModelsVirtualMachine**](ModelsVirtualMachine.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

