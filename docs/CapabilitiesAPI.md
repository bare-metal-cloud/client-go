# \CapabilitiesAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCapabilityDefinitions**](CapabilitiesAPI.md#GetCapabilityDefinitions) | **Get** /capability-definitions | List capability definitions
[**GetVirtualMachinesIdCapabilities**](CapabilitiesAPI.md#GetVirtualMachinesIdCapabilities) | **Get** /virtual-machines/{id}/capabilities | Get a VM&#39;s agent capabilities
[**PostVirtualMachinesIdCapabilitiesRefresh**](CapabilitiesAPI.md#PostVirtualMachinesIdCapabilitiesRefresh) | **Post** /virtual-machines/{id}/capabilities/refresh | Refresh a VM&#39;s agent capabilities



## GetCapabilityDefinitions

> HandlersCapabilityDefinitionsResponse GetCapabilityDefinitions(ctx).Execute()

List capability definitions

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
	resp, r, err := apiClient.CapabilitiesAPI.GetCapabilityDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.GetCapabilityDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapabilityDefinitions`: HandlersCapabilityDefinitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CapabilitiesAPI.GetCapabilityDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCapabilityDefinitionsRequest struct via the builder pattern


### Return type

[**HandlersCapabilityDefinitionsResponse**](HandlersCapabilityDefinitionsResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachinesIdCapabilities

> map[string]interface{} GetVirtualMachinesIdCapabilities(ctx, id).Execute()

Get a VM's agent capabilities

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
	resp, r, err := apiClient.CapabilitiesAPI.GetVirtualMachinesIdCapabilities(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.GetVirtualMachinesIdCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdCapabilities`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CapabilitiesAPI.GetVirtualMachinesIdCapabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdCapabilitiesRequest struct via the builder pattern


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


## PostVirtualMachinesIdCapabilitiesRefresh

> map[string]interface{} PostVirtualMachinesIdCapabilitiesRefresh(ctx, id).Execute()

Refresh a VM's agent capabilities

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
	resp, r, err := apiClient.CapabilitiesAPI.PostVirtualMachinesIdCapabilitiesRefresh(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.PostVirtualMachinesIdCapabilitiesRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdCapabilitiesRefresh`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CapabilitiesAPI.PostVirtualMachinesIdCapabilitiesRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdCapabilitiesRefreshRequest struct via the builder pattern


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

