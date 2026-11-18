# \RunnersInsightsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersInsights**](RunnersInsightsAPI.md#GetOrgsOrgIdRunnersInsights) | **Get** /orgs/{org_id}/runners/insights | The org&#39;s runners insights — the concurrency series, the pinned percentiles, and the open window



## GetOrgsOrgIdRunnersInsights

> HandlersRunnerInsightsResponse GetOrgsOrgIdRunnersInsights(ctx, orgId).Profile(profile).Window(window).Execute()

The org's runners insights — the concurrency series, the pinned percentiles, and the open window



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
	profile := "profile_example" // string | Placement profile filter (empty = org-wide) (optional)
	window := "window_example" // string | RFC3339 hour — scopes the pinned windows to that hour (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersInsightsAPI.GetOrgsOrgIdRunnersInsights(context.Background(), orgId).Profile(profile).Window(window).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersInsightsAPI.GetOrgsOrgIdRunnersInsights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersInsights`: HandlersRunnerInsightsResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersInsightsAPI.GetOrgsOrgIdRunnersInsights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profile** | **string** | Placement profile filter (empty &#x3D; org-wide) | 
 **window** | **string** | RFC3339 hour — scopes the pinned windows to that hour | 

### Return type

[**HandlersRunnerInsightsResponse**](HandlersRunnerInsightsResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

