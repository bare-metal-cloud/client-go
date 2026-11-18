# \MultipassVmsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMultipassVmsId**](MultipassVmsAPI.md#GetMultipassVmsId) | **Get** /multipass-vms/{id} | Get Multipass VM
[**GetMultipassVmsIdHost**](MultipassVmsAPI.md#GetMultipassVmsIdHost) | **Get** /multipass-vms/{id}/host | Get Multipass VM Host
[**GetVirtualMachinesHostVmIdMultipassVms**](MultipassVmsAPI.md#GetVirtualMachinesHostVmIdMultipassVms) | **Get** /virtual-machines/{host_vm_id}/multipass-vms | List Multipass VMs
[**PostApiV1MultipassVmsDelete**](MultipassVmsAPI.md#PostApiV1MultipassVmsDelete) | **Post** /api/v1/multipass-vms/delete | Handle Multipass Delete Request
[**PostApiV1MultipassVmsSync**](MultipassVmsAPI.md#PostApiV1MultipassVmsSync) | **Post** /api/v1/multipass-vms/sync | Handle Multipass Sync Message



## GetMultipassVmsId

> ModelsMultipassVM GetMultipassVmsId(ctx, id).Execute()

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
	resp, r, err := apiClient.MultipassVmsAPI.GetMultipassVmsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassVmsAPI.GetMultipassVmsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipassVmsId`: ModelsMultipassVM
	fmt.Fprintf(os.Stdout, "Response from `MultipassVmsAPI.GetMultipassVmsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipassVmsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsMultipassVM**](ModelsMultipassVM.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipassVmsIdHost

> ModelsVirtualMachine GetMultipassVmsIdHost(ctx, id).Execute()

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
	resp, r, err := apiClient.MultipassVmsAPI.GetMultipassVmsIdHost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassVmsAPI.GetMultipassVmsIdHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipassVmsIdHost`: ModelsVirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `MultipassVmsAPI.GetMultipassVmsIdHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipassVmsIdHostRequest struct via the builder pattern


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


## GetVirtualMachinesHostVmIdMultipassVms

> []ModelsMultipassVM GetVirtualMachinesHostVmIdMultipassVms(ctx, hostVmId).Execute()

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
	resp, r, err := apiClient.MultipassVmsAPI.GetVirtualMachinesHostVmIdMultipassVms(context.Background(), hostVmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassVmsAPI.GetVirtualMachinesHostVmIdMultipassVms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesHostVmIdMultipassVms`: []ModelsMultipassVM
	fmt.Fprintf(os.Stdout, "Response from `MultipassVmsAPI.GetVirtualMachinesHostVmIdMultipassVms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostVmId** | **int32** | Host VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesHostVmIdMultipassVmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelsMultipassVM**](ModelsMultipassVM.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

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

