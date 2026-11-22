# \StorageAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdStorageVolumes**](StorageAPI.md#GetOrgsOrgIdStorageVolumes) | **Get** /orgs/{org_id}/storage/volumes | List PVC usage for an organization



## GetOrgsOrgIdStorageVolumes

> HandlersStorageVolumesResponse GetOrgsOrgIdStorageVolumes(ctx, orgId).Execute()

List PVC usage for an organization



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
	orgId := int32(56) // int32 | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetOrgsOrgIdStorageVolumes(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetOrgsOrgIdStorageVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdStorageVolumes`: HandlersStorageVolumesResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetOrgsOrgIdStorageVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdStorageVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersStorageVolumesResponse**](HandlersStorageVolumesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

