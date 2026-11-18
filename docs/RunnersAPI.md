# \RunnersAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersMachinesMachineIdMetrics**](RunnersAPI.md#GetOrgsOrgIdRunnersMachinesMachineIdMetrics) | **Get** /orgs/{org_id}/runners/machines/{machine_id}/metrics | A machine&#39;s sampled metrics series
[**GetOrgsOrgIdRunnersMetrics**](RunnersAPI.md#GetOrgsOrgIdRunnersMetrics) | **Get** /orgs/{org_id}/runners/metrics | The org&#39;s sampled metrics series



## GetOrgsOrgIdRunnersMachinesMachineIdMetrics

> HandlersRunnerMetricsSeriesResponse GetOrgsOrgIdRunnersMachinesMachineIdMetrics(ctx, orgId, machineId).From(from).To(to).Limit(limit).Execute()

A machine's sampled metrics series



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
	orgId := int32(56) // int32 | The org id (must match the caller's active org)
	machineId := int32(56) // int32 | The machine id
	from := "from_example" // string | The window's start (RFC3339; default 24 hours ago) (optional)
	to := "to_example" // string | The window's end (RFC3339; default now) (optional)
	limit := int32(56) // int32 | The row cap (default 500, at most 2000) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersAPI.GetOrgsOrgIdRunnersMachinesMachineIdMetrics(context.Background(), orgId, machineId).From(from).To(to).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersAPI.GetOrgsOrgIdRunnersMachinesMachineIdMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersMachinesMachineIdMetrics`: HandlersRunnerMetricsSeriesResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersAPI.GetOrgsOrgIdRunnersMachinesMachineIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id (must match the caller&#39;s active org) | 
**machineId** | **int32** | The machine id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersMachinesMachineIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **string** | The window&#39;s start (RFC3339; default 24 hours ago) | 
 **to** | **string** | The window&#39;s end (RFC3339; default now) | 
 **limit** | **int32** | The row cap (default 500, at most 2000) | 

### Return type

[**HandlersRunnerMetricsSeriesResponse**](HandlersRunnerMetricsSeriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersMetrics

> HandlersRunnerMetricsSeriesResponse GetOrgsOrgIdRunnersMetrics(ctx, orgId).MachineId(machineId).From(from).To(to).Limit(limit).Execute()

The org's sampled metrics series



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
	orgId := int32(56) // int32 | The org id (must match the caller's active org)
	machineId := int32(56) // int32 | Narrow to one machine (optional)
	from := "from_example" // string | The window's start (RFC3339; default 24 hours ago) (optional)
	to := "to_example" // string | The window's end (RFC3339; default now) (optional)
	limit := int32(56) // int32 | The row cap (default 500, at most 2000) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersAPI.GetOrgsOrgIdRunnersMetrics(context.Background(), orgId).MachineId(machineId).From(from).To(to).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersAPI.GetOrgsOrgIdRunnersMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersMetrics`: HandlersRunnerMetricsSeriesResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersAPI.GetOrgsOrgIdRunnersMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id (must match the caller&#39;s active org) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **machineId** | **int32** | Narrow to one machine | 
 **from** | **string** | The window&#39;s start (RFC3339; default 24 hours ago) | 
 **to** | **string** | The window&#39;s end (RFC3339; default now) | 
 **limit** | **int32** | The row cap (default 500, at most 2000) | 

### Return type

[**HandlersRunnerMetricsSeriesResponse**](HandlersRunnerMetricsSeriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

