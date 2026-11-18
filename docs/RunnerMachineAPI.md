# \RunnerMachineAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1RunnerMachineVerifyPoll**](RunnerMachineAPI.md#GetApiV1RunnerMachineVerifyPoll) | **Get** /api/v1/runner/machine/verify/poll | Verify-relay long-poll (machine chain)
[**GetRunnerMachineCancels**](RunnerMachineAPI.md#GetRunnerMachineCancels) | **Get** /runner/machine/cancels | Poll a build machine&#39;s pending cancel directives (machine credential)
[**GetRunnerMachineDispatch**](RunnerMachineAPI.md#GetRunnerMachineDispatch) | **Get** /runner/machine/dispatch | Poll a build machine&#39;s dispatched jobs (machine credential)
[**PostApiV1RunnerMachineVerifyVerdict**](RunnerMachineAPI.md#PostApiV1RunnerMachineVerifyVerdict) | **Post** /api/v1/runner/machine/verify/verdict | Verify-relay verdict (machine chain)
[**PostRunnerMachineBuildsCompletion**](RunnerMachineAPI.md#PostRunnerMachineBuildsCompletion) | **Post** /runner/machine/builds/completion | Deliver a build completion (machine credential)
[**PostRunnerMachineBuildsFailure**](RunnerMachineAPI.md#PostRunnerMachineBuildsFailure) | **Post** /runner/machine/builds/failure | Deliver a build failure report (machine credential)
[**PostRunnerMachineFacts**](RunnerMachineAPI.md#PostRunnerMachineFacts) | **Post** /runner/machine/facts | Deliver a start attempt&#39;s probed runtime facts (machine credential)
[**PostRunnerMachineJobsComplete**](RunnerMachineAPI.md#PostRunnerMachineJobsComplete) | **Post** /runner/machine/jobs/complete | Deliver a runner-start job&#39;s terminal (machine credential)
[**PostRunnerMachineJobsJobIdHeartbeat**](RunnerMachineAPI.md#PostRunnerMachineJobsJobIdHeartbeat) | **Post** /runner/machine/jobs/{job_id}/heartbeat | Report a build job heartbeat (machine credential)
[**PostRunnerMachineLogs**](RunnerMachineAPI.md#PostRunnerMachineLogs) | **Post** /runner/machine/logs | Deliver build log lines (machine credential)
[**PostRunnerMachineMetrics**](RunnerMachineAPI.md#PostRunnerMachineMetrics) | **Post** /runner/machine/metrics | Deliver sampled metrics (machine credential)



## GetApiV1RunnerMachineVerifyPoll

> VerifyrelayRequest GetApiV1RunnerMachineVerifyPoll(ctx).Execute()

Verify-relay long-poll (machine chain)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.GetApiV1RunnerMachineVerifyPoll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.GetApiV1RunnerMachineVerifyPoll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1RunnerMachineVerifyPoll`: VerifyrelayRequest
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.GetApiV1RunnerMachineVerifyPoll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1RunnerMachineVerifyPollRequest struct via the builder pattern


### Return type

[**VerifyrelayRequest**](VerifyrelayRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunnerMachineCancels

> BuilderMachineCancelResponse GetRunnerMachineCancels(ctx).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()

Poll a build machine's pending cancel directives (machine credential)



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
	xBMCMachineSig := "xBMCMachineSig_example" // string | The machine's Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE=enforce) (optional)
	xBMCMachineTs := "xBMCMachineTs_example" // string | The signed request timestamp, decimal unix seconds (60s skew window) (optional)
	xBMCMachineNonce := "xBMCMachineNonce_example" // string | The request's one-time nonce (replayed nonces are refused) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.GetRunnerMachineCancels(context.Background()).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.GetRunnerMachineCancels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerMachineCancels`: BuilderMachineCancelResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.GetRunnerMachineCancels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerMachineCancelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBMCMachineSig** | **string** | The machine&#39;s Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE&#x3D;enforce) | 
 **xBMCMachineTs** | **string** | The signed request timestamp, decimal unix seconds (60s skew window) | 
 **xBMCMachineNonce** | **string** | The request&#39;s one-time nonce (replayed nonces are refused) | 

### Return type

[**BuilderMachineCancelResponse**](BuilderMachineCancelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunnerMachineDispatch

> RunnerpoolMachinePollResponse GetRunnerMachineDispatch(ctx).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()

Poll a build machine's dispatched jobs (machine credential)



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
	xBMCMachineSig := "xBMCMachineSig_example" // string | The machine's Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE=enforce) (optional)
	xBMCMachineTs := "xBMCMachineTs_example" // string | The signed request timestamp, decimal unix seconds (60s skew window) (optional)
	xBMCMachineNonce := "xBMCMachineNonce_example" // string | The request's one-time nonce (replayed nonces are refused) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.GetRunnerMachineDispatch(context.Background()).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.GetRunnerMachineDispatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerMachineDispatch`: RunnerpoolMachinePollResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.GetRunnerMachineDispatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerMachineDispatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBMCMachineSig** | **string** | The machine&#39;s Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE&#x3D;enforce) | 
 **xBMCMachineTs** | **string** | The signed request timestamp, decimal unix seconds (60s skew window) | 
 **xBMCMachineNonce** | **string** | The request&#39;s one-time nonce (replayed nonces are refused) | 

### Return type

[**RunnerpoolMachinePollResponse**](RunnerpoolMachinePollResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1RunnerMachineVerifyVerdict

> map[string]bool PostApiV1RunnerMachineVerifyVerdict(ctx).Body(body).Execute()

Verify-relay verdict (machine chain)



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
	body := *openapiclient.NewHandlersVerifyVerdictRequest() // HandlersVerifyVerdictRequest | The verdict

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.PostApiV1RunnerMachineVerifyVerdict(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.PostApiV1RunnerMachineVerifyVerdict``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1RunnerMachineVerifyVerdict`: map[string]bool
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.PostApiV1RunnerMachineVerifyVerdict`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1RunnerMachineVerifyVerdictRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersVerifyVerdictRequest**](HandlersVerifyVerdictRequest.md) | The verdict | 

### Return type

**map[string]bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerMachineBuildsCompletion

> HandlersRunnerBuildCompletionResponse PostRunnerMachineBuildsCompletion(ctx).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()

Deliver a build completion (machine credential)



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
	body := *openapiclient.NewHandlersRunnerBuildCompletionRequest() // HandlersRunnerBuildCompletionRequest | The signed completion document
	xBMCMachineSig := "xBMCMachineSig_example" // string | The machine's Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE=enforce) (optional)
	xBMCMachineTs := "xBMCMachineTs_example" // string | The signed request timestamp, decimal unix seconds (60s skew window) (optional)
	xBMCMachineNonce := "xBMCMachineNonce_example" // string | The request's one-time nonce (replayed nonces are refused) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.PostRunnerMachineBuildsCompletion(context.Background()).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.PostRunnerMachineBuildsCompletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerMachineBuildsCompletion`: HandlersRunnerBuildCompletionResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.PostRunnerMachineBuildsCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerMachineBuildsCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersRunnerBuildCompletionRequest**](HandlersRunnerBuildCompletionRequest.md) | The signed completion document | 
 **xBMCMachineSig** | **string** | The machine&#39;s Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE&#x3D;enforce) | 
 **xBMCMachineTs** | **string** | The signed request timestamp, decimal unix seconds (60s skew window) | 
 **xBMCMachineNonce** | **string** | The request&#39;s one-time nonce (replayed nonces are refused) | 

### Return type

[**HandlersRunnerBuildCompletionResponse**](HandlersRunnerBuildCompletionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerMachineBuildsFailure

> map[string]string PostRunnerMachineBuildsFailure(ctx).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()

Deliver a build failure report (machine credential)



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
	body := *openapiclient.NewHandlersRunnerBuildCompletionRequest() // HandlersRunnerBuildCompletionRequest | The signed failure report
	xBMCMachineSig := "xBMCMachineSig_example" // string | The machine's Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE=enforce) (optional)
	xBMCMachineTs := "xBMCMachineTs_example" // string | The signed request timestamp, decimal unix seconds (60s skew window) (optional)
	xBMCMachineNonce := "xBMCMachineNonce_example" // string | The request's one-time nonce (replayed nonces are refused) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.PostRunnerMachineBuildsFailure(context.Background()).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.PostRunnerMachineBuildsFailure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerMachineBuildsFailure`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.PostRunnerMachineBuildsFailure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerMachineBuildsFailureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersRunnerBuildCompletionRequest**](HandlersRunnerBuildCompletionRequest.md) | The signed failure report | 
 **xBMCMachineSig** | **string** | The machine&#39;s Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE&#x3D;enforce) | 
 **xBMCMachineTs** | **string** | The signed request timestamp, decimal unix seconds (60s skew window) | 
 **xBMCMachineNonce** | **string** | The request&#39;s one-time nonce (replayed nonces are refused) | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerMachineFacts

> map[string]string PostRunnerMachineFacts(ctx).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()

Deliver a start attempt's probed runtime facts (machine credential)



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
	body := *openapiclient.NewHandlersRunnerMachineFactsRequest() // HandlersRunnerMachineFactsRequest | The probed facts report
	xBMCMachineSig := "xBMCMachineSig_example" // string | The machine's Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE=enforce) (optional)
	xBMCMachineTs := "xBMCMachineTs_example" // string | The signed request timestamp, decimal unix seconds (60s skew window) (optional)
	xBMCMachineNonce := "xBMCMachineNonce_example" // string | The request's one-time nonce (replayed nonces are refused) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.PostRunnerMachineFacts(context.Background()).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.PostRunnerMachineFacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerMachineFacts`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.PostRunnerMachineFacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerMachineFactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersRunnerMachineFactsRequest**](HandlersRunnerMachineFactsRequest.md) | The probed facts report | 
 **xBMCMachineSig** | **string** | The machine&#39;s Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE&#x3D;enforce) | 
 **xBMCMachineTs** | **string** | The signed request timestamp, decimal unix seconds (60s skew window) | 
 **xBMCMachineNonce** | **string** | The request&#39;s one-time nonce (replayed nonces are refused) | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerMachineJobsComplete

> map[string]string PostRunnerMachineJobsComplete(ctx).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()

Deliver a runner-start job's terminal (machine credential)



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
	body := *openapiclient.NewHandlersRunnerBuildCompletionRequest() // HandlersRunnerBuildCompletionRequest | The signed start-terminal document
	xBMCMachineSig := "xBMCMachineSig_example" // string | The machine's Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE=enforce) (optional)
	xBMCMachineTs := "xBMCMachineTs_example" // string | The signed request timestamp, decimal unix seconds (60s skew window) (optional)
	xBMCMachineNonce := "xBMCMachineNonce_example" // string | The request's one-time nonce (replayed nonces are refused) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.PostRunnerMachineJobsComplete(context.Background()).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.PostRunnerMachineJobsComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerMachineJobsComplete`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.PostRunnerMachineJobsComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerMachineJobsCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersRunnerBuildCompletionRequest**](HandlersRunnerBuildCompletionRequest.md) | The signed start-terminal document | 
 **xBMCMachineSig** | **string** | The machine&#39;s Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE&#x3D;enforce) | 
 **xBMCMachineTs** | **string** | The signed request timestamp, decimal unix seconds (60s skew window) | 
 **xBMCMachineNonce** | **string** | The request&#39;s one-time nonce (replayed nonces are refused) | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerMachineJobsJobIdHeartbeat

> map[string]string PostRunnerMachineJobsJobIdHeartbeat(ctx, jobId).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()

Report a build job heartbeat (machine credential)



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
	jobId := "jobId_example" // string | The build job id
	xBMCMachineSig := "xBMCMachineSig_example" // string | The machine's Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE=enforce) (optional)
	xBMCMachineTs := "xBMCMachineTs_example" // string | The signed request timestamp, decimal unix seconds (60s skew window) (optional)
	xBMCMachineNonce := "xBMCMachineNonce_example" // string | The request's one-time nonce (replayed nonces are refused) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.PostRunnerMachineJobsJobIdHeartbeat(context.Background(), jobId).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.PostRunnerMachineJobsJobIdHeartbeat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerMachineJobsJobIdHeartbeat`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.PostRunnerMachineJobsJobIdHeartbeat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The build job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerMachineJobsJobIdHeartbeatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBMCMachineSig** | **string** | The machine&#39;s Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE&#x3D;enforce) | 
 **xBMCMachineTs** | **string** | The signed request timestamp, decimal unix seconds (60s skew window) | 
 **xBMCMachineNonce** | **string** | The request&#39;s one-time nonce (replayed nonces are refused) | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerMachineLogs

> map[string]int32 PostRunnerMachineLogs(ctx).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()

Deliver build log lines (machine credential)



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
	body := *openapiclient.NewHandlersRunnerMachineLogsRequest() // HandlersRunnerMachineLogsRequest | The log batch
	xBMCMachineSig := "xBMCMachineSig_example" // string | The machine's Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE=enforce) (optional)
	xBMCMachineTs := "xBMCMachineTs_example" // string | The signed request timestamp, decimal unix seconds (60s skew window) (optional)
	xBMCMachineNonce := "xBMCMachineNonce_example" // string | The request's one-time nonce (replayed nonces are refused) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.PostRunnerMachineLogs(context.Background()).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.PostRunnerMachineLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerMachineLogs`: map[string]int32
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.PostRunnerMachineLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerMachineLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersRunnerMachineLogsRequest**](HandlersRunnerMachineLogsRequest.md) | The log batch | 
 **xBMCMachineSig** | **string** | The machine&#39;s Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE&#x3D;enforce) | 
 **xBMCMachineTs** | **string** | The signed request timestamp, decimal unix seconds (60s skew window) | 
 **xBMCMachineNonce** | **string** | The request&#39;s one-time nonce (replayed nonces are refused) | 

### Return type

**map[string]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerMachineMetrics

> map[string]int32 PostRunnerMachineMetrics(ctx).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()

Deliver sampled metrics (machine credential)



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
	body := *openapiclient.NewContractsRunnerMetrics() // ContractsRunnerMetrics | The sampled-metrics batch
	xBMCMachineSig := "xBMCMachineSig_example" // string | The machine's Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE=enforce) (optional)
	xBMCMachineTs := "xBMCMachineTs_example" // string | The signed request timestamp, decimal unix seconds (60s skew window) (optional)
	xBMCMachineNonce := "xBMCMachineNonce_example" // string | The request's one-time nonce (replayed nonces are refused) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerMachineAPI.PostRunnerMachineMetrics(context.Background()).Body(body).XBMCMachineSig(xBMCMachineSig).XBMCMachineTs(xBMCMachineTs).XBMCMachineNonce(xBMCMachineNonce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerMachineAPI.PostRunnerMachineMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerMachineMetrics`: map[string]int32
	fmt.Fprintf(os.Stdout, "Response from `RunnerMachineAPI.PostRunnerMachineMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerMachineMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ContractsRunnerMetrics**](ContractsRunnerMetrics.md) | The sampled-metrics batch | 
 **xBMCMachineSig** | **string** | The machine&#39;s Ed25519 signature over the request payload (AAR.1 slice C; required once BMC_MACHINE_REQUEST_SIGNATURE&#x3D;enforce) | 
 **xBMCMachineTs** | **string** | The signed request timestamp, decimal unix seconds (60s skew window) | 
 **xBMCMachineNonce** | **string** | The request&#39;s one-time nonce (replayed nonces are refused) | 

### Return type

**map[string]int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

