# \ScheduledAccessAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostServiceAccountsSaIdSealSecret**](ScheduledAccessAPI.md#PostServiceAccountsSaIdSealSecret) | **Post** /service-accounts/{sa_id}/seal-secret | Seal a service-account secret for consumption



## PostServiceAccountsSaIdSealSecret

> map[string]interface{} PostServiceAccountsSaIdSealSecret(ctx, saId).Body(body).XVaultRentalID(xVaultRentalID).Execute()

Seal a service-account secret for consumption



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
	saId := int32(56) // int32 | Service Account ID
	body := *openapiclient.NewHandlersSealSecretRequest() // HandlersSealSecretRequest | Seal request
	xVaultRentalID := "xVaultRentalID_example" // string | Rental id (human path; the sa.secret.seal rental for this secret row) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduledAccessAPI.PostServiceAccountsSaIdSealSecret(context.Background(), saId).Body(body).XVaultRentalID(xVaultRentalID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledAccessAPI.PostServiceAccountsSaIdSealSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceAccountsSaIdSealSecret`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ScheduledAccessAPI.PostServiceAccountsSaIdSealSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saId** | **int32** | Service Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceAccountsSaIdSealSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersSealSecretRequest**](HandlersSealSecretRequest.md) | Seal request | 
 **xVaultRentalID** | **string** | Rental id (human path; the sa.secret.seal rental for this secret row) | 

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

