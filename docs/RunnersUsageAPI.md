# \RunnersUsageAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersUsage**](RunnersUsageAPI.md#GetOrgsOrgIdRunnersUsage) | **Get** /orgs/{org_id}/runners/usage | The org&#39;s runners usage — unit-minutes per day by platform, the projection, and the previous period
[**GetOrgsOrgIdRunnersUsageCsv**](RunnersUsageAPI.md#GetOrgsOrgIdRunnersUsageCsv) | **Get** /orgs/{org_id}/runners/usage.csv | The org&#39;s runners usage as CSV — the screen&#39;s daily rows verbatim
[**GetOrgsOrgIdRunnersUsageNow**](RunnersUsageAPI.md#GetOrgsOrgIdRunnersUsageNow) | **Get** /orgs/{org_id}/runners/usage/now | The org&#39;s runners usage \&quot;now\&quot; figures — the Overview&#39;s live snapshot



## GetOrgsOrgIdRunnersUsage

> HandlersRunnerUsageResponse GetOrgsOrgIdRunnersUsage(ctx, orgId).Period(period).Execute()

The org's runners usage — unit-minutes per day by platform, the projection, and the previous period



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
	period := "period_example" // string | The month, YYYY-MM (default: the current UTC month) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersUsageAPI.GetOrgsOrgIdRunnersUsage(context.Background(), orgId).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersUsageAPI.GetOrgsOrgIdRunnersUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersUsage`: HandlersRunnerUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersUsageAPI.GetOrgsOrgIdRunnersUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | The month, YYYY-MM (default: the current UTC month) | 

### Return type

[**HandlersRunnerUsageResponse**](HandlersRunnerUsageResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersUsageCsv

> string GetOrgsOrgIdRunnersUsageCsv(ctx, orgId).Period(period).Execute()

The org's runners usage as CSV — the screen's daily rows verbatim



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
	period := "period_example" // string | The month, YYYY-MM (default: the current UTC month) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersUsageAPI.GetOrgsOrgIdRunnersUsageCsv(context.Background(), orgId).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersUsageAPI.GetOrgsOrgIdRunnersUsageCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersUsageCsv`: string
	fmt.Fprintf(os.Stdout, "Response from `RunnersUsageAPI.GetOrgsOrgIdRunnersUsageCsv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersUsageCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | The month, YYYY-MM (default: the current UTC month) | 

### Return type

**string**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersUsageNow

> HandlersRunnerUsageNowResponse GetOrgsOrgIdRunnersUsageNow(ctx, orgId).Execute()

The org's runners usage \"now\" figures — the Overview's live snapshot



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
	resp, r, err := apiClient.RunnersUsageAPI.GetOrgsOrgIdRunnersUsageNow(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersUsageAPI.GetOrgsOrgIdRunnersUsageNow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersUsageNow`: HandlersRunnerUsageNowResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersUsageAPI.GetOrgsOrgIdRunnersUsageNow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersUsageNowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerUsageNowResponse**](HandlersRunnerUsageNowResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

