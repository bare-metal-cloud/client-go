# \ConsoleAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostVmsIdConsoleDevMode**](ConsoleAPI.md#PostVmsIdConsoleDevMode) | **Post** /vms/{id}/console/dev-mode | Record VM Console developer-mode toggle (audit only)



## PostVmsIdConsoleDevMode

> PostVmsIdConsoleDevMode(ctx, id).Body(body).Execute()

Record VM Console developer-mode toggle (audit only)



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
	id := int32(56) // int32 | Virtual Machine id
	body := *openapiclient.NewHandlersVMConsoleDevModeRequest() // HandlersVMConsoleDevModeRequest | Dev mode payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConsoleAPI.PostVmsIdConsoleDevMode(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleAPI.PostVmsIdConsoleDevMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Virtual Machine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVmsIdConsoleDevModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersVMConsoleDevModeRequest**](HandlersVMConsoleDevModeRequest.md) | Dev mode payload | 

### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

