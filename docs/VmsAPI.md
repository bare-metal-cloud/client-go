# \VmsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostVmsVerifySsh**](VmsAPI.md#PostVmsVerifySsh) | **Post** /vms/verify-ssh | Verify SSH connection



## PostVmsVerifySsh

> HandlersVerifySSHResponse PostVmsVerifySsh(ctx).Request(request).Execute()

Verify SSH connection



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
	request := *openapiclient.NewHandlersVerifySSHRequest("Host_example", int32(123), "Username_example") // HandlersVerifySSHRequest | SSH verification request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VmsAPI.PostVmsVerifySsh(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VmsAPI.PostVmsVerifySsh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVmsVerifySsh`: HandlersVerifySSHResponse
	fmt.Fprintf(os.Stdout, "Response from `VmsAPI.PostVmsVerifySsh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVmsVerifySshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersVerifySSHRequest**](HandlersVerifySSHRequest.md) | SSH verification request | 

### Return type

[**HandlersVerifySSHResponse**](HandlersVerifySSHResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

