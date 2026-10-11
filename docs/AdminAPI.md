# \AdminAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApiV1AdminUsersIdMfaDisable**](AdminAPI.md#PostApiV1AdminUsersIdMfaDisable) | **Post** /api/v1/admin/users/{id}/mfa/disable | Admin force-disable user MFA



## PostApiV1AdminUsersIdMfaDisable

> HandlersMFADisableResponse PostApiV1AdminUsersIdMfaDisable(ctx, id).Execute()

Admin force-disable user MFA



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
	id := int32(56) // int32 | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostApiV1AdminUsersIdMfaDisable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostApiV1AdminUsersIdMfaDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1AdminUsersIdMfaDisable`: HandlersMFADisableResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostApiV1AdminUsersIdMfaDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1AdminUsersIdMfaDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersMFADisableResponse**](HandlersMFADisableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

