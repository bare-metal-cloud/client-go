# \MetricsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVirtualMachinesIdMetrics**](MetricsAPI.md#GetVirtualMachinesIdMetrics) | **Get** /virtual-machines/{id}/metrics | Get VM metrics
[**GetVirtualMachinesIdMetricsHistory**](MetricsAPI.md#GetVirtualMachinesIdMetricsHistory) | **Get** /virtual-machines/{id}/metrics/history | Get VM metrics history



## GetVirtualMachinesIdMetrics

> HandlersVMMetricsResponse GetVirtualMachinesIdMetrics(ctx, id).Execute()

Get VM metrics



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
	id := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetVirtualMachinesIdMetrics(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetVirtualMachinesIdMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdMetrics`: HandlersVMMetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetVirtualMachinesIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersVMMetricsResponse**](HandlersVMMetricsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachinesIdMetricsHistory

> HandlersMetricHistoryResponse GetVirtualMachinesIdMetricsHistory(ctx, id).Range_(range_).Granularity(granularity).Metric(metric).Execute()

Get VM metrics history



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
	id := int32(56) // int32 | VM ID
	range_ := "range__example" // string | Time range (1h, 24h, 7d) (optional) (default to "1h")
	granularity := "granularity_example" // string | Granularity (raw, hourly, daily) (optional) (default to "raw")
	metric := "metric_example" // string | Specific metric name (default: all) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetVirtualMachinesIdMetricsHistory(context.Background(), id).Range_(range_).Granularity(granularity).Metric(metric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetVirtualMachinesIdMetricsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdMetricsHistory`: HandlersMetricHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetVirtualMachinesIdMetricsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdMetricsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **string** | Time range (1h, 24h, 7d) | [default to &quot;1h&quot;]
 **granularity** | **string** | Granularity (raw, hourly, daily) | [default to &quot;raw&quot;]
 **metric** | **string** | Specific metric name (default: all) | 

### Return type

[**HandlersMetricHistoryResponse**](HandlersMetricHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

