# \RunnersInstallsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersInstallsInstallId**](RunnersInstallsAPI.md#GetOrgsOrgIdRunnersInstallsInstallId) | **Get** /orgs/{org_id}/runners/installs/{install_id} | The install record&#39;s state — the Add page&#39;s poll
[**PostOrgsOrgIdRunnersInstallsInstallIdCancel**](RunnersInstallsAPI.md#PostOrgsOrgIdRunnersInstallsInstallIdCancel) | **Post** /orgs/{org_id}/runners/installs/{install_id}/cancel | Cancel an install from any pre-add state
[**PostOrgsOrgIdRunnersMachinesInstallCommand**](RunnersInstallsAPI.md#PostOrgsOrgIdRunnersMachinesInstallCommand) | **Post** /orgs/{org_id}/runners/machines/install-command | Issue the runner install command for a chosen machine
[**PostRunnersInstallAck**](RunnersInstallsAPI.md#PostRunnersInstallAck) | **Post** /runners/install-ack | The install command&#39;s first machine contact (one-time token)



## GetOrgsOrgIdRunnersInstallsInstallId

> HandlersRunnerInstallStatusResponse GetOrgsOrgIdRunnersInstallsInstallId(ctx, orgId, installId).Execute()

The install record's state — the Add page's poll



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
	orgId := int32(56) // int32 | Organization ID
	installId := "installId_example" // string | Install ID (the issue answer's installId)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersInstallsAPI.GetOrgsOrgIdRunnersInstallsInstallId(context.Background(), orgId, installId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersInstallsAPI.GetOrgsOrgIdRunnersInstallsInstallId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersInstallsInstallId`: HandlersRunnerInstallStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersInstallsAPI.GetOrgsOrgIdRunnersInstallsInstallId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**installId** | **string** | Install ID (the issue answer&#39;s installId) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersInstallsInstallIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerInstallStatusResponse**](HandlersRunnerInstallStatusResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersInstallsInstallIdCancel

> HandlersRunnerInstallStatusResponse PostOrgsOrgIdRunnersInstallsInstallIdCancel(ctx, orgId, installId).Execute()

Cancel an install from any pre-add state



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
	orgId := int32(56) // int32 | Organization ID
	installId := "installId_example" // string | Install ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersInstallsAPI.PostOrgsOrgIdRunnersInstallsInstallIdCancel(context.Background(), orgId, installId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersInstallsAPI.PostOrgsOrgIdRunnersInstallsInstallIdCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersInstallsInstallIdCancel`: HandlersRunnerInstallStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersInstallsAPI.PostOrgsOrgIdRunnersInstallsInstallIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**installId** | **string** | Install ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersInstallsInstallIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerInstallStatusResponse**](HandlersRunnerInstallStatusResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersMachinesInstallCommand

> HandlersRunnerInstallCommandResponse PostOrgsOrgIdRunnersMachinesInstallCommand(ctx, orgId).Body(body).Execute()

Issue the runner install command for a chosen machine



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
	orgId := int32(56) // int32 | Organization ID
	body := *openapiclient.NewHandlersRunnerInstallIssueRequest() // HandlersRunnerInstallIssueRequest | The chosen machine (the agent id, the picker's row id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersInstallsAPI.PostOrgsOrgIdRunnersMachinesInstallCommand(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersInstallsAPI.PostOrgsOrgIdRunnersMachinesInstallCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersMachinesInstallCommand`: HandlersRunnerInstallCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersInstallsAPI.PostOrgsOrgIdRunnersMachinesInstallCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersMachinesInstallCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRunnerInstallIssueRequest**](HandlersRunnerInstallIssueRequest.md) | The chosen machine (the agent id, the picker&#39;s row id) | 

### Return type

[**HandlersRunnerInstallCommandResponse**](HandlersRunnerInstallCommandResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnersInstallAck

> HandlersRunnerInstallStatusResponse PostRunnersInstallAck(ctx).Body(body).Execute()

The install command's first machine contact (one-time token)



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
	body := *openapiclient.NewHandlersRunnerInstallAckRequest() // HandlersRunnerInstallAckRequest | The one-time token, the machine fingerprint, the preflight verdicts

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersInstallsAPI.PostRunnersInstallAck(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersInstallsAPI.PostRunnersInstallAck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnersInstallAck`: HandlersRunnerInstallStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersInstallsAPI.PostRunnersInstallAck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnersInstallAckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersRunnerInstallAckRequest**](HandlersRunnerInstallAckRequest.md) | The one-time token, the machine fingerprint, the preflight verdicts | 

### Return type

[**HandlersRunnerInstallStatusResponse**](HandlersRunnerInstallStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

