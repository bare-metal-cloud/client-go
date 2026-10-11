# \VirtualMachinesAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVirtualMachinesId**](VirtualMachinesAPI.md#DeleteVirtualMachinesId) | **Delete** /virtual-machines/{id} | Delete a VM
[**GetVirtualMachinesJobs**](VirtualMachinesAPI.md#GetVirtualMachinesJobs) | **Get** /virtual-machines/jobs | List VM operations
[**GetVirtualMachinesJobsId**](VirtualMachinesAPI.md#GetVirtualMachinesJobsId) | **Get** /virtual-machines/jobs/{id} | Get VM operation
[**PostCloudProvidersAzureProvision**](VirtualMachinesAPI.md#PostCloudProvidersAzureProvision) | **Post** /cloud-providers/azure/provision | Provision an Azure VM
[**PostVirtualMachinesProvisionAws**](VirtualMachinesAPI.md#PostVirtualMachinesProvisionAws) | **Post** /virtual-machines/provision/aws | Provision an AWS EC2 instance
[**PostVirtualMachinesProvisionAwsWorker**](VirtualMachinesAPI.md#PostVirtualMachinesProvisionAwsWorker) | **Post** /virtual-machines/provision/aws/worker | Provision an AWS EC2 instance via background worker
[**PostVirtualMachinesProvisionGcp**](VirtualMachinesAPI.md#PostVirtualMachinesProvisionGcp) | **Post** /virtual-machines/provision/gcp | Provision a GCP Compute Engine instance
[**PostVmsIdCheckK3s**](VirtualMachinesAPI.md#PostVmsIdCheckK3s) | **Post** /vms/{id}/check-k3s | Check a VM for pre-existing k3s



## DeleteVirtualMachinesId

> DeleteVirtualMachinesId200Response DeleteVirtualMachinesId(ctx, id).Force(force).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.DeleteVirtualMachinesId(context.Background(), id).Force(force).Execute()
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


## PostVmsIdCheckK3s

> HandlersVMCheckK3sResponse PostVmsIdCheckK3s(ctx, id).Body(body).Execute()

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
	resp, r, err := apiClient.VirtualMachinesAPI.PostVmsIdCheckK3s(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PostVmsIdCheckK3s``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVmsIdCheckK3s`: HandlersVMCheckK3sResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PostVmsIdCheckK3s`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVmsIdCheckK3sRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersVMCheckK3sRequest**](HandlersVMCheckK3sRequest.md) | Optional decision params. In the BODY: &#39;+&#39; in k3s tags decodes as space in a query string. | 

### Return type

[**HandlersVMCheckK3sResponse**](HandlersVMCheckK3sResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

