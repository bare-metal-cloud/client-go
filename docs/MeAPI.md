# \MeAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApiV1MeSudo**](MeAPI.md#PostApiV1MeSudo) | **Post** /api/v1/me/sudo | Refresh MFA sudo window for privileged actions



## PostApiV1MeSudo

> map[string]interface{} PostApiV1MeSudo(ctx).Request(request).Execute()

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
	resp, r, err := apiClient.MeAPI.PostApiV1MeSudo(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.PostApiV1MeSudo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1MeSudo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.PostApiV1MeSudo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1MeSudoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersSudoChallengeRequest**](HandlersSudoChallengeRequest.md) | TOTP code | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

