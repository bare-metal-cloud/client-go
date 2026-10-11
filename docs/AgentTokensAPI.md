# \AgentTokensAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAgentTokensValidate**](AgentTokensAPI.md#PostAgentTokensValidate) | **Post** /agent/tokens/validate | Validate token



## PostAgentTokensValidate

> HandlersValidateTokenResponse PostAgentTokensValidate(ctx).Request(request).Execute()

Validate token



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
	request := *openapiclient.NewHandlersValidateTokenRequest() // HandlersValidateTokenRequest | Token validation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTokensAPI.PostAgentTokensValidate(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTokensAPI.PostAgentTokensValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentTokensValidate`: HandlersValidateTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentTokensAPI.PostAgentTokensValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentTokensValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersValidateTokenRequest**](HandlersValidateTokenRequest.md) | Token validation request | 

### Return type

[**HandlersValidateTokenResponse**](HandlersValidateTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

