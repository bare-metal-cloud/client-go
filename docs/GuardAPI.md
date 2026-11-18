# \GuardAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVirtualMachinesIdReapprovals**](GuardAPI.md#GetVirtualMachinesIdReapprovals) | **Get** /virtual-machines/{id}/reapprovals | List the VM&#39;s pending re-approval asks
[**PostVirtualMachinesIdGuardMandatesOpIdRevoke**](GuardAPI.md#PostVirtualMachinesIdGuardMandatesOpIdRevoke) | **Post** /virtual-machines/{id}/guard-mandates/{opId}/revoke | Revoke a standing guard mandate
[**PostVirtualMachinesIdGuardTokens**](GuardAPI.md#PostVirtualMachinesIdGuardTokens) | **Post** /virtual-machines/{id}/guard-tokens | Store browser-minted guard op tokens
[**PostVirtualMachinesIdReapprovalsRidDismiss**](GuardAPI.md#PostVirtualMachinesIdReapprovalsRidDismiss) | **Post** /virtual-machines/{id}/reapprovals/{rid}/dismiss | Dismiss a re-approval ask



## GetVirtualMachinesIdReapprovals

> HandlersPendingReapprovalsResponse GetVirtualMachinesIdReapprovals(ctx, id).Execute()

List the VM's pending re-approval asks



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
	resp, r, err := apiClient.GuardAPI.GetVirtualMachinesIdReapprovals(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuardAPI.GetVirtualMachinesIdReapprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdReapprovals`: HandlersPendingReapprovalsResponse
	fmt.Fprintf(os.Stdout, "Response from `GuardAPI.GetVirtualMachinesIdReapprovals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdReapprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersPendingReapprovalsResponse**](HandlersPendingReapprovalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVirtualMachinesIdGuardMandatesOpIdRevoke

> map[string]string PostVirtualMachinesIdGuardMandatesOpIdRevoke(ctx, id, opId).Execute()

Revoke a standing guard mandate



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
	opId := "opId_example" // string | The mandate's op_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuardAPI.PostVirtualMachinesIdGuardMandatesOpIdRevoke(context.Background(), id, opId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuardAPI.PostVirtualMachinesIdGuardMandatesOpIdRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdGuardMandatesOpIdRevoke`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `GuardAPI.PostVirtualMachinesIdGuardMandatesOpIdRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 
**opId** | **string** | The mandate&#39;s op_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdGuardMandatesOpIdRevokeRequest struct via the builder pattern


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


## PostVirtualMachinesIdGuardTokens

> HandlersStoreGuardTokensResponse PostVirtualMachinesIdGuardTokens(ctx, id).Body(body).Execute()

Store browser-minted guard op tokens



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
	body := *openapiclient.NewHandlersStoreGuardTokensRequest() // HandlersStoreGuardTokensRequest | The minted tokens

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuardAPI.PostVirtualMachinesIdGuardTokens(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuardAPI.PostVirtualMachinesIdGuardTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdGuardTokens`: HandlersStoreGuardTokensResponse
	fmt.Fprintf(os.Stdout, "Response from `GuardAPI.PostVirtualMachinesIdGuardTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdGuardTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersStoreGuardTokensRequest**](HandlersStoreGuardTokensRequest.md) | The minted tokens | 

### Return type

[**HandlersStoreGuardTokensResponse**](HandlersStoreGuardTokensResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVirtualMachinesIdReapprovalsRidDismiss

> map[string]string PostVirtualMachinesIdReapprovalsRidDismiss(ctx, id, rid).Execute()

Dismiss a re-approval ask



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
	rid := "rid_example" // string | Re-approval ask id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuardAPI.PostVirtualMachinesIdReapprovalsRidDismiss(context.Background(), id, rid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuardAPI.PostVirtualMachinesIdReapprovalsRidDismiss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdReapprovalsRidDismiss`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `GuardAPI.PostVirtualMachinesIdReapprovalsRidDismiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 
**rid** | **string** | Re-approval ask id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdReapprovalsRidDismissRequest struct via the builder pattern


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

