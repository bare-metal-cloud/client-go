# \DurabilityAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostDurabilityBackupRunsAuthorize**](DurabilityAPI.md#PostDurabilityBackupRunsAuthorize) | **Post** /durability/backup-runs/authorize | Authorize one backup run (backend#4180 component 3)



## PostDurabilityBackupRunsAuthorize

> PostDurabilityBackupRunsAuthorize(ctx).Request(request).Execute()

Authorize one backup run (backend#4180 component 3)



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
	request := *openapiclient.NewHandlersBackupRunAuthorizeRequest() // HandlersBackupRunAuthorizeRequest | org_id, run_id, rental_id (the ticket)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DurabilityAPI.PostDurabilityBackupRunsAuthorize(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DurabilityAPI.PostDurabilityBackupRunsAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDurabilityBackupRunsAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersBackupRunAuthorizeRequest**](HandlersBackupRunAuthorizeRequest.md) | org_id, run_id, rental_id (the ticket) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

