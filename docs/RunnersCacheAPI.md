# \RunnersCacheAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersCache**](RunnersCacheAPI.md#GetOrgsOrgIdRunnersCache) | **Get** /orgs/{org_id}/runners/cache | The org&#39;s cache scopes
[**GetOrgsOrgIdRunnersCacheEvents**](RunnersCacheAPI.md#GetOrgsOrgIdRunnersCacheEvents) | **Get** /orgs/{org_id}/runners/cache/events | The org&#39;s cache events ledger



## GetOrgsOrgIdRunnersCache

> HandlersRunnerCacheResponse GetOrgsOrgIdRunnersCache(ctx, orgId).Execute()

The org's cache scopes



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
	resp, r, err := apiClient.RunnersCacheAPI.GetOrgsOrgIdRunnersCache(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersCacheAPI.GetOrgsOrgIdRunnersCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersCache`: HandlersRunnerCacheResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersCacheAPI.GetOrgsOrgIdRunnersCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerCacheResponse**](HandlersRunnerCacheResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersCacheEvents

> HandlersRunnerCacheEventsResponse GetOrgsOrgIdRunnersCacheEvents(ctx, orgId).Execute()

The org's cache events ledger



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
	resp, r, err := apiClient.RunnersCacheAPI.GetOrgsOrgIdRunnersCacheEvents(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersCacheAPI.GetOrgsOrgIdRunnersCacheEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersCacheEvents`: HandlersRunnerCacheEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersCacheAPI.GetOrgsOrgIdRunnersCacheEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersCacheEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerCacheEventsResponse**](HandlersRunnerCacheEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

