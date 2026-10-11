# \AgentInstallationAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAgentInstallCommands**](AgentInstallationAPI.md#PostAgentInstallCommands) | **Post** /agent/install/commands | Get unified install command



## PostAgentInstallCommands

> HandlersUnifiedInstallCommandResponse PostAgentInstallCommands(ctx).Request(request).Execute()

Get unified install command



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
	request := *openapiclient.NewHandlersInstallCommandPostRequest("ClusterType_example") // HandlersInstallCommandPostRequest | Install command request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentInstallationAPI.PostAgentInstallCommands(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentInstallationAPI.PostAgentInstallCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentInstallCommands`: HandlersUnifiedInstallCommandResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentInstallationAPI.PostAgentInstallCommands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentInstallCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersInstallCommandPostRequest**](HandlersInstallCommandPostRequest.md) | Install command request | 

### Return type

[**HandlersUnifiedInstallCommandResponse**](HandlersUnifiedInstallCommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

