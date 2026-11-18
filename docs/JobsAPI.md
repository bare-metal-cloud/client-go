# \JobsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJobs**](JobsAPI.md#GetJobs) | **Get** /jobs | List jobs in the active project
[**GetJobsJobId**](JobsAPI.md#GetJobsJobId) | **Get** /jobs/{job_id} | Get a job with its step timeline
[**GetMultipassHostsJobsId**](JobsAPI.md#GetMultipassHostsJobsId) | **Get** /multipass-hosts/jobs/{id} | Get the status of a multipass host job



## GetJobs

> HandlersJobListResponse GetJobs(ctx).Page(page).Limit(limit).Status(status).JobType(jobType).UserId(userId).Execute()

List jobs in the active project

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
	page := int32(56) // int32 | Page (default 1) (optional)
	limit := int32(56) // int32 | Page size (optional)
	status := "status_example" // string | Filter by status (optional)
	jobType := "jobType_example" // string | Filter by job type (optional)
	userId := int32(56) // int32 | Only jobs started by this user (admins may pass any; others only their own) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.GetJobs(context.Background()).Page(page).Limit(limit).Status(status).JobType(jobType).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobs`: HandlersJobListResponse
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page (default 1) | 
 **limit** | **int32** | Page size | 
 **status** | **string** | Filter by status | 
 **jobType** | **string** | Filter by job type | 
 **userId** | **int32** | Only jobs started by this user (admins may pass any; others only their own) | 

### Return type

[**HandlersJobListResponse**](HandlersJobListResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobsJobId

> HandlersJobDetailResponse GetJobsJobId(ctx, jobId).Order(order).Execute()

Get a job with its step timeline

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
	jobId := "jobId_example" // string | Job ID
	order := "order_example" // string | asc or desc (default desc) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.GetJobsJobId(context.Background(), jobId).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.GetJobsJobId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobsJobId`: HandlersJobDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.GetJobsJobId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsJobIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | **string** | asc or desc (default desc) | 

### Return type

[**HandlersJobDetailResponse**](HandlersJobDetailResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipassHostsJobsId

> ModelsJobMetadata GetMultipassHostsJobsId(ctx, id).Execute()

Get the status of a multipass host job

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
	id := "id_example" // string | Job ID (UUID)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.GetMultipassHostsJobsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.GetMultipassHostsJobsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipassHostsJobsId`: ModelsJobMetadata
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.GetMultipassHostsJobsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipassHostsJobsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsJobMetadata**](ModelsJobMetadata.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

