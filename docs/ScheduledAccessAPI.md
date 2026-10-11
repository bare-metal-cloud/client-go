# \ScheduledAccessAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApiV1ServiceAccountsSaIdSealSecret**](ScheduledAccessAPI.md#PostApiV1ServiceAccountsSaIdSealSecret) | **Post** /api/v1/service-accounts/{sa_id}/seal-secret | Seal a service-account secret for worker consumption



## PostApiV1ServiceAccountsSaIdSealSecret

> map[string]interface{} PostApiV1ServiceAccountsSaIdSealSecret(ctx, saId).Body(body).Execute()

Seal a service-account secret for worker consumption



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduledAccessAPI.PostApiV1ServiceAccountsSaIdSealSecret(context.Background(), saId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledAccessAPI.PostApiV1ServiceAccountsSaIdSealSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1ServiceAccountsSaIdSealSecret`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ScheduledAccessAPI.PostApiV1ServiceAccountsSaIdSealSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saId** | **int32** | Service Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1ServiceAccountsSaIdSealSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersSealSecretRequest**](HandlersSealSecretRequest.md) | Seal request | 

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

