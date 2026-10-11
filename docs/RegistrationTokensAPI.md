# \RegistrationTokensAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAgentRegistrationToken**](RegistrationTokensAPI.md#PostAgentRegistrationToken) | **Post** /agent/registration-token | Generate GitHub-style registration token



## PostAgentRegistrationToken

> HandlersGenerateRegistrationTokenResponse PostAgentRegistrationToken(ctx).Request(request).Execute()

Generate GitHub-style registration token



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
	request := *openapiclient.NewHandlersGenerateRegistrationTokenRequest("ClusterName_example", "ClusterType_example") // HandlersGenerateRegistrationTokenRequest | Token generation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistrationTokensAPI.PostAgentRegistrationToken(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrationTokensAPI.PostAgentRegistrationToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentRegistrationToken`: HandlersGenerateRegistrationTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `RegistrationTokensAPI.PostAgentRegistrationToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentRegistrationTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersGenerateRegistrationTokenRequest**](HandlersGenerateRegistrationTokenRequest.md) | Token generation request | 

### Return type

[**HandlersGenerateRegistrationTokenResponse**](HandlersGenerateRegistrationTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

