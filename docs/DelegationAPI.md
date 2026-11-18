# \DelegationAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAiAgentsMandates**](DelegationAPI.md#GetAiAgentsMandates) | **Get** /ai-agents/mandates | List the org&#39;s AI-delegation mandates
[**GetAiAgentsToolCatalog**](DelegationAPI.md#GetAiAgentsToolCatalog) | **Get** /ai-agents/tool-catalog | List the AI agent&#39;s tool catalog
[**PostAiAgentsMandates**](DelegationAPI.md#PostAiAgentsMandates) | **Post** /ai-agents/mandates | Store browser-minted AI-delegation mandates
[**PostAiAgentsMandatesMidRevoke**](DelegationAPI.md#PostAiAgentsMandatesMidRevoke) | **Post** /ai-agents/mandates/{mid}/revoke | Revoke an AI-delegation mandate
[**PostVirtualMachinesIdDelegationActions**](DelegationAPI.md#PostVirtualMachinesIdDelegationActions) | **Post** /virtual-machines/{id}/delegation-actions | Store a delegated action
[**PostVirtualMachinesIdDelegationProposals**](DelegationAPI.md#PostVirtualMachinesIdDelegationProposals) | **Post** /virtual-machines/{id}/delegation-proposals | Propose a destructive operation for human co-sign



## GetAiAgentsMandates

> HandlersDelegationMandatesResponse GetAiAgentsMandates(ctx).Execute()

List the org's AI-delegation mandates



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
	resp, r, err := apiClient.DelegationAPI.GetAiAgentsMandates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.GetAiAgentsMandates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiAgentsMandates`: HandlersDelegationMandatesResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.GetAiAgentsMandates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiAgentsMandatesRequest struct via the builder pattern


### Return type

[**HandlersDelegationMandatesResponse**](HandlersDelegationMandatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiAgentsToolCatalog

> HandlersAgentToolCatalogResponse GetAiAgentsToolCatalog(ctx).Execute()

List the AI agent's tool catalog



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
	resp, r, err := apiClient.DelegationAPI.GetAiAgentsToolCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.GetAiAgentsToolCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiAgentsToolCatalog`: HandlersAgentToolCatalogResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.GetAiAgentsToolCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiAgentsToolCatalogRequest struct via the builder pattern


### Return type

[**HandlersAgentToolCatalogResponse**](HandlersAgentToolCatalogResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiAgentsMandates

> HandlersStoreDelegationMandatesResponse PostAiAgentsMandates(ctx).Body(body).Execute()

Store browser-minted AI-delegation mandates



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
	body := *openapiclient.NewHandlersStoreDelegationMandatesRequest() // HandlersStoreDelegationMandatesRequest | The minted mandates

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegationAPI.PostAiAgentsMandates(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.PostAiAgentsMandates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiAgentsMandates`: HandlersStoreDelegationMandatesResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.PostAiAgentsMandates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAiAgentsMandatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersStoreDelegationMandatesRequest**](HandlersStoreDelegationMandatesRequest.md) | The minted mandates | 

### Return type

[**HandlersStoreDelegationMandatesResponse**](HandlersStoreDelegationMandatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAiAgentsMandatesMidRevoke

> map[string]string PostAiAgentsMandatesMidRevoke(ctx, mid).Body(body).Execute()

Revoke an AI-delegation mandate



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
	mid := "mid_example" // string | Mandate id
	body := *openapiclient.NewHandlersRevokeDelegationMandateRequest() // HandlersRevokeDelegationMandateRequest | The minted revoke token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegationAPI.PostAiAgentsMandatesMidRevoke(context.Background(), mid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.PostAiAgentsMandatesMidRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAiAgentsMandatesMidRevoke`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.PostAiAgentsMandatesMidRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mid** | **string** | Mandate id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAiAgentsMandatesMidRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRevokeDelegationMandateRequest**](HandlersRevokeDelegationMandateRequest.md) | The minted revoke token | 

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


## PostVirtualMachinesIdDelegationActions

> map[string]string PostVirtualMachinesIdDelegationActions(ctx, id).Body(body).Execute()

Store a delegated action



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
	body := *openapiclient.NewHandlersDelegatedActionRequest() // HandlersDelegatedActionRequest | The signed action

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegationAPI.PostVirtualMachinesIdDelegationActions(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.PostVirtualMachinesIdDelegationActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdDelegationActions`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.PostVirtualMachinesIdDelegationActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdDelegationActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersDelegatedActionRequest**](HandlersDelegatedActionRequest.md) | The signed action | 

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


## PostVirtualMachinesIdDelegationProposals

> map[string]string PostVirtualMachinesIdDelegationProposals(ctx, id).Body(body).Execute()

Propose a destructive operation for human co-sign



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
	body := *openapiclient.NewHandlersDelegationProposalRequest() // HandlersDelegationProposalRequest | The proposal

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegationAPI.PostVirtualMachinesIdDelegationProposals(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegationAPI.PostVirtualMachinesIdDelegationProposals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdDelegationProposals`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DelegationAPI.PostVirtualMachinesIdDelegationProposals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdDelegationProposalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersDelegationProposalRequest**](HandlersDelegationProposalRequest.md) | The proposal | 

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

