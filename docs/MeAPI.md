# \MeAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMeSudo**](MeAPI.md#PostMeSudo) | **Post** /me/sudo | Refresh MFA sudo window for privileged actions



## PostMeSudo

> map[string]interface{} PostMeSudo(ctx).Request(request).Execute()

Refresh MFA sudo window for privileged actions



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
	request := *openapiclient.NewHandlersSudoChallengeRequest() // HandlersSudoChallengeRequest | TOTP code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.PostMeSudo(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.PostMeSudo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMeSudo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.PostMeSudo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMeSudoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersSudoChallengeRequest**](HandlersSudoChallengeRequest.md) | TOTP code | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

