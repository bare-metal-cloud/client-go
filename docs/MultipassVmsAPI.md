# \MultipassVmsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1MultipassVmsId**](MultipassVmsAPI.md#GetApiV1MultipassVmsId) | **Get** /api/v1/multipass-vms/{id} | Get Multipass VM
[**GetApiV1MultipassVmsIdHost**](MultipassVmsAPI.md#GetApiV1MultipassVmsIdHost) | **Get** /api/v1/multipass-vms/{id}/host | Get Multipass VM Host
[**GetApiV1VirtualMachinesHostVmIdMultipassVms**](MultipassVmsAPI.md#GetApiV1VirtualMachinesHostVmIdMultipassVms) | **Get** /api/v1/virtual-machines/{host_vm_id}/multipass-vms | List Multipass VMs
[**PostApiV1MultipassVmsDelete**](MultipassVmsAPI.md#PostApiV1MultipassVmsDelete) | **Post** /api/v1/multipass-vms/delete | Handle Multipass Delete Request
[**PostApiV1MultipassVmsSync**](MultipassVmsAPI.md#PostApiV1MultipassVmsSync) | **Post** /api/v1/multipass-vms/sync | Handle Multipass Sync Message



## GetApiV1MultipassVmsId

> ModelsMultipassVM GetApiV1MultipassVmsId(ctx, id).Execute()

Get Multipass VM



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
	resp, r, err := apiClient.MultipassVmsAPI.GetApiV1MultipassVmsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassVmsAPI.GetApiV1MultipassVmsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1MultipassVmsId`: ModelsMultipassVM
	fmt.Fprintf(os.Stdout, "Response from `MultipassVmsAPI.GetApiV1MultipassVmsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MultipassVmsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsMultipassVM**](ModelsMultipassVM.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1MultipassVmsIdHost

> ModelsVirtualMachine GetApiV1MultipassVmsIdHost(ctx, id).Execute()

Get Multipass VM Host



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
	resp, r, err := apiClient.MultipassVmsAPI.GetApiV1MultipassVmsIdHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassVmsAPI.GetApiV1MultipassVmsIdHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1MultipassVmsIdHost`: ModelsVirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `MultipassVmsAPI.GetApiV1MultipassVmsIdHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1MultipassVmsIdHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsVirtualMachine**](ModelsVirtualMachine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1VirtualMachinesHostVmIdMultipassVms

> []ModelsMultipassVM GetApiV1VirtualMachinesHostVmIdMultipassVms(ctx, hostVmId).Execute()

List Multipass VMs



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
	hostVmId := int32(56) // int32 | Host VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipassVmsAPI.GetApiV1VirtualMachinesHostVmIdMultipassVms(context.Background(), hostVmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassVmsAPI.GetApiV1VirtualMachinesHostVmIdMultipassVms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1VirtualMachinesHostVmIdMultipassVms`: []ModelsMultipassVM
	fmt.Fprintf(os.Stdout, "Response from `MultipassVmsAPI.GetApiV1VirtualMachinesHostVmIdMultipassVms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostVmId** | **int32** | Host VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1VirtualMachinesHostVmIdMultipassVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelsMultipassVM**](ModelsMultipassVM.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1MultipassVmsDelete

> map[string]string PostApiV1MultipassVmsDelete(ctx).Request(request).Execute()

Handle Multipass Delete Request



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
	request := *openapiclient.NewHandlersMultipassDeleteRequest() // HandlersMultipassDeleteRequest | Delete request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipassVmsAPI.PostApiV1MultipassVmsDelete(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassVmsAPI.PostApiV1MultipassVmsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1MultipassVmsDelete`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `MultipassVmsAPI.PostApiV1MultipassVmsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1MultipassVmsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersMultipassDeleteRequest**](HandlersMultipassDeleteRequest.md) | Delete request | 

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


## PostApiV1MultipassVmsSync

> map[string]string PostApiV1MultipassVmsSync(ctx).Request(request).Execute()

Handle Multipass Sync Message



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
	request := *openapiclient.NewHandlersMultipassSyncRequest() // HandlersMultipassSyncRequest | Sync message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipassVmsAPI.PostApiV1MultipassVmsSync(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassVmsAPI.PostApiV1MultipassVmsSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1MultipassVmsSync`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `MultipassVmsAPI.PostApiV1MultipassVmsSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1MultipassVmsSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersMultipassSyncRequest**](HandlersMultipassSyncRequest.md) | Sync message | 

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

