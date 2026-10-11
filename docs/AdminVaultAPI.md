# \AdminVaultAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1AdminVaultRotationHealth**](AdminVaultAPI.md#GetApiV1AdminVaultRotationHealth) | **Get** /api/v1/admin/vault/rotation-health | SMK rotation health check
[**GetApiV1AdminVaultRotationStatus**](AdminVaultAPI.md#GetApiV1AdminVaultRotationStatus) | **Get** /api/v1/admin/vault/rotation-status | SMK rotation status



## GetApiV1AdminVaultRotationHealth

> HandlersVaultRotationHealthResponse GetApiV1AdminVaultRotationHealth(ctx).Execute()

SMK rotation health check



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
	resp, r, err := apiClient.AdminVaultAPI.GetApiV1AdminVaultRotationHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminVaultAPI.GetApiV1AdminVaultRotationHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminVaultRotationHealth`: HandlersVaultRotationHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminVaultAPI.GetApiV1AdminVaultRotationHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminVaultRotationHealthRequest struct via the builder pattern


### Return type

[**HandlersVaultRotationHealthResponse**](HandlersVaultRotationHealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminVaultRotationStatus

> HandlersVaultRotationAdminStatusResponse GetApiV1AdminVaultRotationStatus(ctx).Execute()

SMK rotation status



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
	resp, r, err := apiClient.AdminVaultAPI.GetApiV1AdminVaultRotationStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminVaultAPI.GetApiV1AdminVaultRotationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminVaultRotationStatus`: HandlersVaultRotationAdminStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminVaultAPI.GetApiV1AdminVaultRotationStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminVaultRotationStatusRequest struct via the builder pattern


### Return type

[**HandlersVaultRotationAdminStatusResponse**](HandlersVaultRotationAdminStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

