# \RunnersJobsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersJobs**](RunnersJobsAPI.md#GetOrgsOrgIdRunnersJobs) | **Get** /orgs/{org_id}/runners/jobs | List the org&#39;s unified job ledger — every execution across pool, builder, and GitHub
[**GetOrgsOrgIdRunnersJobsJobRef**](RunnersJobsAPI.md#GetOrgsOrgIdRunnersJobsJobRef) | **Get** /orgs/{org_id}/runners/jobs/{job_ref} | Get one job&#39;s detail — the ledger row expanded with the source record
[**GetOrgsOrgIdRunnersJobsJobRefLog**](RunnersJobsAPI.md#GetOrgsOrgIdRunnersJobsJobRefLog) | **Get** /orgs/{org_id}/runners/jobs/{job_ref}/log | Read a job&#39;s captured log
[**GetOrgsOrgIdRunnersJobsJobRefLogLines**](RunnersJobsAPI.md#GetOrgsOrgIdRunnersJobsJobRefLogLines) | **Get** /orgs/{org_id}/runners/jobs/{job_ref}/log/lines | Backfill a job&#39;s live log stream by sequence number



## GetOrgsOrgIdRunnersJobs

> HandlersRunnerJobsLedgerResponse GetOrgsOrgIdRunnersJobs(ctx, orgId).Limit(limit).State(state).Execute()

List the org's unified job ledger — every execution across pool, builder, and GitHub



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
	limit := int32(56) // int32 | Maximum rows (default 50, max 200) (optional)
	state := "state_example" // string | Exact-match state filter (the source's state verbatim — the three sources use THREE vocabularies: pool pending|provisioned|dispatched|running|terminal; builder submitting|dispatched|queued|lost|succeeded|failed|cancelled; github queued|in_progress|completed — filter per source or omit) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersJobsAPI.GetOrgsOrgIdRunnersJobs(context.Background(), orgId).Limit(limit).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersJobsAPI.GetOrgsOrgIdRunnersJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersJobs`: HandlersRunnerJobsLedgerResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersJobsAPI.GetOrgsOrgIdRunnersJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum rows (default 50, max 200) | 
 **state** | **string** | Exact-match state filter (the source&#39;s state verbatim — the three sources use THREE vocabularies: pool pending|provisioned|dispatched|running|terminal; builder submitting|dispatched|queued|lost|succeeded|failed|cancelled; github queued|in_progress|completed — filter per source or omit) | 

### Return type

[**HandlersRunnerJobsLedgerResponse**](HandlersRunnerJobsLedgerResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersJobsJobRef

> HandlersRunnerJobDetailResponse GetOrgsOrgIdRunnersJobsJobRef(ctx, orgId, jobRef).Source(source).Execute()

Get one job's detail — the ledger row expanded with the source record



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
	jobRef := "jobRef_example" // string | The job's reference (the ledger row's jobRef)
	source := "source_example" // string | Narrow the reference to one source (pool | builder | github) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersJobsAPI.GetOrgsOrgIdRunnersJobsJobRef(context.Background(), orgId, jobRef).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersJobsAPI.GetOrgsOrgIdRunnersJobsJobRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersJobsJobRef`: HandlersRunnerJobDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersJobsAPI.GetOrgsOrgIdRunnersJobsJobRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**jobRef** | **string** | The job&#39;s reference (the ledger row&#39;s jobRef) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersJobsJobRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **string** | Narrow the reference to one source (pool | builder | github) | 

### Return type

[**HandlersRunnerJobDetailResponse**](HandlersRunnerJobDetailResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersJobsJobRefLog

> string GetOrgsOrgIdRunnersJobsJobRefLog(ctx, orgId, jobRef).Source(source).Execute()

Read a job's captured log



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
	jobRef := "jobRef_example" // string | The job's reference (the ledger row's jobRef)
	source := "source_example" // string | Narrow the reference to one source (pool | builder | github) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersJobsAPI.GetOrgsOrgIdRunnersJobsJobRefLog(context.Background(), orgId, jobRef).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersJobsAPI.GetOrgsOrgIdRunnersJobsJobRefLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersJobsJobRefLog`: string
	fmt.Fprintf(os.Stdout, "Response from `RunnersJobsAPI.GetOrgsOrgIdRunnersJobsJobRefLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**jobRef** | **string** | The job&#39;s reference (the ledger row&#39;s jobRef) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersJobsJobRefLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **string** | Narrow the reference to one source (pool | builder | github) | 

### Return type

**string**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersJobsJobRefLogLines

> HandlersRunnerJobLogLinesResponse GetOrgsOrgIdRunnersJobsJobRefLogLines(ctx, orgId, jobRef).Source(source).AfterSeq(afterSeq).Execute()

Backfill a job's live log stream by sequence number



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
	jobRef := "jobRef_example" // string | The job's reference (the ledger row's jobRef)
	source := "source_example" // string | Narrow the reference to one source (pool | builder | github) (optional)
	afterSeq := int32(56) // int32 | Return only lines with a larger sequence (default 0 — the whole buffered window) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersJobsAPI.GetOrgsOrgIdRunnersJobsJobRefLogLines(context.Background(), orgId, jobRef).Source(source).AfterSeq(afterSeq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersJobsAPI.GetOrgsOrgIdRunnersJobsJobRefLogLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersJobsJobRefLogLines`: HandlersRunnerJobLogLinesResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersJobsAPI.GetOrgsOrgIdRunnersJobsJobRefLogLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**jobRef** | **string** | The job&#39;s reference (the ledger row&#39;s jobRef) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersJobsJobRefLogLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **string** | Narrow the reference to one source (pool | builder | github) | 
 **afterSeq** | **int32** | Return only lines with a larger sequence (default 0 — the whole buffered window) | 

### Return type

[**HandlersRunnerJobLogLinesResponse**](HandlersRunnerJobLogLinesResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

