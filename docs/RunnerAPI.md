# \RunnerAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRunnerProjectsProjectIdAssignmentsAgentId**](RunnerAPI.md#DeleteRunnerProjectsProjectIdAssignmentsAgentId) | **Delete** /runner/projects/{project_id}/assignments/{agent_id} | Remove a build machine assignment (panel)
[**DeleteRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts**](RunnerAPI.md#DeleteRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts) | **Delete** /runner-runtime/_apis/pipelines/workflows/{runId}/artifacts | Delete an artifact by run and name
[**GetRunnerJobsJobId**](RunnerAPI.md#GetRunnerJobsJobId) | **Get** /runner/jobs/{job_id} | Poll a hosted job&#39;s status (engine identity; v1 BMC-admin)
[**GetRunnerPanelProjectsProjectIdSection**](RunnerAPI.md#GetRunnerPanelProjectsProjectIdSection) | **Get** /runner/panel/projects/{project_id}/section | The per-project hosted-runner section (panel)
[**GetRunnerProjectsProjectIdAssignments**](RunnerAPI.md#GetRunnerProjectsProjectIdAssignments) | **Get** /runner/projects/{project_id}/assignments | List the project&#39;s assigned build machines (panel)
[**GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts**](RunnerAPI.md#GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts) | **Get** /runner-runtime/_apis/pipelines/workflows/{runId}/artifacts | List a run&#39;s finalized artifacts (the download entry)
[**GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsName**](RunnerAPI.md#GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsName) | **Get** /runner-runtime/_apis/pipelines/workflows/{runId}/artifacts/{name} | Get one artifact&#39;s detail + a fresh container URL
[**GetRunnerRuntimeRunnerBlob**](RunnerAPI.md#GetRunnerRuntimeRunnerBlob) | **Get** /runner-runtime/runner-blob | Download the artifact&#39;s bytes through the signed blob URL
[**GetRunnerRuntimeRunnerFiles**](RunnerAPI.md#GetRunnerRuntimeRunnerFiles) | **Get** /runner-runtime/runner-files | List the files the artifact holds (the item listing)
[**PatchRunnerRuntimeRunnerBlob**](RunnerAPI.md#PatchRunnerRuntimeRunnerBlob) | **Patch** /runner-runtime/runner-blob | Upload one artifact chunk through the signed blob URL
[**PostRunnerBuildsCompletion**](RunnerAPI.md#PostRunnerBuildsCompletion) | **Post** /runner/builds/completion | Report a build completion (builder channel; v1 BMC-admin)
[**PostRunnerCompletions**](RunnerAPI.md#PostRunnerCompletions) | **Post** /runner/completions | Report a job completion (agent channel; v1 BMC-admin)
[**PostRunnerJobs**](RunnerAPI.md#PostRunnerJobs) | **Post** /runner/jobs | Submit a hosted-runner job (engine identity; v1 BMC-admin)
[**PostRunnerJobsJobIdHeartbeat**](RunnerAPI.md#PostRunnerJobsJobIdHeartbeat) | **Post** /runner/jobs/{job_id}/heartbeat | Report a job heartbeat (agent channel; v1 BMC-admin)
[**PostRunnerJobsJobIdKill**](RunnerAPI.md#PostRunnerJobsJobIdKill) | **Post** /runner/jobs/{job_id}/kill | Kill a hosted-runner job (org owner or BMC-admin)
[**PostRunnerProjectsProjectIdAssignments**](RunnerAPI.md#PostRunnerProjectsProjectIdAssignments) | **Post** /runner/projects/{project_id}/assignments | Assign a build machine to the project (panel)
[**PostRunnerProjectsProjectIdBuildsDispatch**](RunnerAPI.md#PostRunnerProjectsProjectIdBuildsDispatch) | **Post** /runner/projects/{project_id}/builds/dispatch | Dispatch a build (pipeline entry point; entry 2.5 wires the caller)
[**PostRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts**](RunnerAPI.md#PostRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts) | **Post** /runner-runtime/_apis/pipelines/workflows/{runId}/artifacts | Create an artifact (the job&#39;s identity)
[**PutRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts**](RunnerAPI.md#PutRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts) | **Put** /runner-runtime/_apis/pipelines/workflows/{runId}/artifacts | Finalize an artifact (commit the bytes to the store)



## DeleteRunnerProjectsProjectIdAssignmentsAgentId

> map[string]string DeleteRunnerProjectsProjectIdAssignmentsAgentId(ctx, projectId, agentId).Execute()

Remove a build machine assignment (panel)



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
	projectId := int32(56) // int32 | The project id
	agentId := int32(56) // int32 | The agent id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.DeleteRunnerProjectsProjectIdAssignmentsAgentId(context.Background(), projectId, agentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.DeleteRunnerProjectsProjectIdAssignmentsAgentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRunnerProjectsProjectIdAssignmentsAgentId`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.DeleteRunnerProjectsProjectIdAssignmentsAgentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The project id | 
**agentId** | **int32** | The agent id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRunnerProjectsProjectIdAssignmentsAgentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DeleteRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts

> map[string]string DeleteRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts(ctx, runId).ArtifactName(artifactName).Execute()

Delete an artifact by run and name



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
	runId := int32(56) // int32 | The workflow run id
	artifactName := "artifactName_example" // string | The artifact's name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.DeleteRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts(context.Background(), runId).ArtifactName(artifactName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.DeleteRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.DeleteRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **int32** | The workflow run id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artifactName** | **string** | The artifact&#39;s name | 

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


## GetRunnerJobsJobId

> RunnerpoolJobStatus GetRunnerJobsJobId(ctx, jobId).Execute()

Poll a hosted job's status (engine identity; v1 BMC-admin)



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
	jobId := "jobId_example" // string | The job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.GetRunnerJobsJobId(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.GetRunnerJobsJobId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerJobsJobId`: RunnerpoolJobStatus
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.GetRunnerJobsJobId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerJobsJobIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RunnerpoolJobStatus**](RunnerpoolJobStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunnerPanelProjectsProjectIdSection

> HandlersRunnerProjectSection GetRunnerPanelProjectsProjectIdSection(ctx, projectId).Execute()

The per-project hosted-runner section (panel)



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
	projectId := int32(56) // int32 | The project id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.GetRunnerPanelProjectsProjectIdSection(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.GetRunnerPanelProjectsProjectIdSection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerPanelProjectsProjectIdSection`: HandlersRunnerProjectSection
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.GetRunnerPanelProjectsProjectIdSection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerPanelProjectsProjectIdSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerProjectSection**](HandlersRunnerProjectSection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunnerProjectsProjectIdAssignments

> map[string][]HandlersRunnerAssignmentRow GetRunnerProjectsProjectIdAssignments(ctx, projectId).Execute()

List the project's assigned build machines (panel)



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
	projectId := int32(56) // int32 | The project id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.GetRunnerProjectsProjectIdAssignments(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.GetRunnerProjectsProjectIdAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerProjectsProjectIdAssignments`: map[string][]HandlersRunnerAssignmentRow
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.GetRunnerProjectsProjectIdAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerProjectsProjectIdAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string][]HandlersRunnerAssignmentRow**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts

> map[string]interface{} GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts(ctx, runId).Execute()

List a run's finalized artifacts (the download entry)



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
	runId := int32(56) // int32 | The workflow run id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **int32** | The workflow run id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsName

> map[string]string GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsName(ctx, runId, name).Execute()

Get one artifact's detail + a fresh container URL



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
	runId := int32(56) // int32 | The workflow run id
	name := "name_example" // string | The artifact's name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsName(context.Background(), runId, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsName`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.GetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **int32** | The workflow run id | 
**name** | **string** | The artifact&#39;s name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetRunnerRuntimeRunnerBlob

> *os.File GetRunnerRuntimeRunnerBlob(ctx).Sig(sig).Execute()

Download the artifact's bytes through the signed blob URL



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
	sig := "sig_example" // string | The signed blob URL's signature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.GetRunnerRuntimeRunnerBlob(context.Background()).Sig(sig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.GetRunnerRuntimeRunnerBlob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerRuntimeRunnerBlob`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.GetRunnerRuntimeRunnerBlob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerRuntimeRunnerBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sig** | **string** | The signed blob URL&#39;s signature | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunnerRuntimeRunnerFiles

> map[string]interface{} GetRunnerRuntimeRunnerFiles(ctx).Sig(sig).Execute()

List the files the artifact holds (the item listing)



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
	sig := "sig_example" // string | The signed container URL's signature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.GetRunnerRuntimeRunnerFiles(context.Background()).Sig(sig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.GetRunnerRuntimeRunnerFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunnerRuntimeRunnerFiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.GetRunnerRuntimeRunnerFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunnerRuntimeRunnerFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sig** | **string** | The signed container URL&#39;s signature | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRunnerRuntimeRunnerBlob

> map[string]string PatchRunnerRuntimeRunnerBlob(ctx).Sig(sig).Chunk(chunk).Execute()

Upload one artifact chunk through the signed blob URL



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
	sig := "sig_example" // string | The signed blob URL's signature
	chunk := int32(56) // int32 | The chunk's sequential index (0-based, strictly increasing)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PatchRunnerRuntimeRunnerBlob(context.Background()).Sig(sig).Chunk(chunk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PatchRunnerRuntimeRunnerBlob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRunnerRuntimeRunnerBlob`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PatchRunnerRuntimeRunnerBlob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchRunnerRuntimeRunnerBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sig** | **string** | The signed blob URL&#39;s signature | 
 **chunk** | **int32** | The chunk&#39;s sequential index (0-based, strictly increasing) | 

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


## PostRunnerBuildsCompletion

> HandlersRunnerBuildCompletionResponse PostRunnerBuildsCompletion(ctx).Body(body).Execute()

Report a build completion (builder channel; v1 BMC-admin)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PostRunnerBuildsCompletion(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PostRunnerBuildsCompletion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerBuildsCompletion`: HandlersRunnerBuildCompletionResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PostRunnerBuildsCompletion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerBuildsCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersRunnerBuildCompletionRequest**](HandlersRunnerBuildCompletionRequest.md) | The signed completion document | 

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


## PostRunnerCompletions

> HandlersRunnerCompletionResponse PostRunnerCompletions(ctx).Body(body).Execute()

Report a job completion (agent channel; v1 BMC-admin)



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
	body := *openapiclient.NewHandlersRunnerCompletionRequest() // HandlersRunnerCompletionRequest | The completion report

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PostRunnerCompletions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PostRunnerCompletions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerCompletions`: HandlersRunnerCompletionResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PostRunnerCompletions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerCompletionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersRunnerCompletionRequest**](HandlersRunnerCompletionRequest.md) | The completion report | 

### Return type

[**HandlersRunnerCompletionResponse**](HandlersRunnerCompletionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerJobs

> HandlersRunnerJobAdmittedResponse PostRunnerJobs(ctx).Body(body).Execute()

Submit a hosted-runner job (engine identity; v1 BMC-admin)



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
	body := *openapiclient.NewHandlersRunnerJobAdmissionRequest() // HandlersRunnerJobAdmissionRequest | The signed job-spec submission

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PostRunnerJobs(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PostRunnerJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerJobs`: HandlersRunnerJobAdmittedResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PostRunnerJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersRunnerJobAdmissionRequest**](HandlersRunnerJobAdmissionRequest.md) | The signed job-spec submission | 

### Return type

[**HandlersRunnerJobAdmittedResponse**](HandlersRunnerJobAdmittedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerJobsJobIdHeartbeat

> map[string]string PostRunnerJobsJobIdHeartbeat(ctx, jobId).Body(body).Execute()

Report a job heartbeat (agent channel; v1 BMC-admin)



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
	jobId := "jobId_example" // string | The job id
	body := *openapiclient.NewHandlersRunnerHeartbeatRequest() // HandlersRunnerHeartbeatRequest | An optional advisory heartbeat timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PostRunnerJobsJobIdHeartbeat(context.Background(), jobId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PostRunnerJobsJobIdHeartbeat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerJobsJobIdHeartbeat`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PostRunnerJobsJobIdHeartbeat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerJobsJobIdHeartbeatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRunnerHeartbeatRequest**](HandlersRunnerHeartbeatRequest.md) | An optional advisory heartbeat timestamp | 

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


## PostRunnerJobsJobIdKill

> HandlersRunnerKillResponse PostRunnerJobsJobIdKill(ctx, jobId).Body(body).Execute()

Kill a hosted-runner job (org owner or BMC-admin)



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
	jobId := "jobId_example" // string | The job id
	body := *openapiclient.NewHandlersRunnerKillRequest() // HandlersRunnerKillRequest | An optional kill reason (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PostRunnerJobsJobIdKill(context.Background(), jobId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PostRunnerJobsJobIdKill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerJobsJobIdKill`: HandlersRunnerKillResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PostRunnerJobsJobIdKill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerJobsJobIdKillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRunnerKillRequest**](HandlersRunnerKillRequest.md) | An optional kill reason | 

### Return type

[**HandlersRunnerKillResponse**](HandlersRunnerKillResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerProjectsProjectIdAssignments

> HandlersRunnerAssignmentRow PostRunnerProjectsProjectIdAssignments(ctx, projectId).Body(body).Execute()

Assign a build machine to the project (panel)



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
	projectId := int32(56) // int32 | The project id
	body := map[string]int32{"key": int32(123)} // map[string]int32 | agent_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PostRunnerProjectsProjectIdAssignments(context.Background(), projectId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PostRunnerProjectsProjectIdAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerProjectsProjectIdAssignments`: HandlersRunnerAssignmentRow
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PostRunnerProjectsProjectIdAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerProjectsProjectIdAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]int32** | agent_id | 

### Return type

[**HandlersRunnerAssignmentRow**](HandlersRunnerAssignmentRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerProjectsProjectIdBuildsDispatch

> map[string]interface{} PostRunnerProjectsProjectIdBuildsDispatch(ctx, projectId).Body(body).Execute()

Dispatch a build (pipeline entry point; entry 2.5 wires the caller)



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
	projectId := int32(56) // int32 | The project id
	body := *openapiclient.NewHandlersRunnerBuildDispatchRequest() // HandlersRunnerBuildDispatchRequest | The resolved build

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PostRunnerProjectsProjectIdBuildsDispatch(context.Background(), projectId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PostRunnerProjectsProjectIdBuildsDispatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerProjectsProjectIdBuildsDispatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PostRunnerProjectsProjectIdBuildsDispatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** | The project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerProjectsProjectIdBuildsDispatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRunnerBuildDispatchRequest**](HandlersRunnerBuildDispatchRequest.md) | The resolved build | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts

> map[string]string PostRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts(ctx, runId).Body(body).Execute()

Create an artifact (the job's identity)



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
	runId := int32(56) // int32 | The workflow run id (must be the token job's run)
	body := map[string]string{"key": "Inner_example"} // map[string]string | The create grammar (Name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PostRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts(context.Background(), runId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PostRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PostRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **int32** | The workflow run id (must be the token job&#39;s run) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]string** | The create grammar (Name) | 

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


## PutRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts

> map[string]string PutRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts(ctx, runId).ArtifactName(artifactName).Execute()

Finalize an artifact (commit the bytes to the store)



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
	runId := int32(56) // int32 | The workflow run id
	artifactName := "artifactName_example" // string | The artifact's name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerAPI.PutRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts(context.Background(), runId).ArtifactName(artifactName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerAPI.PutRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnerAPI.PutRunnerRuntimeApisPipelinesWorkflowsRunIdArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **int32** | The workflow run id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRunnerRuntimeApisPipelinesWorkflowsRunIdArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artifactName** | **string** | The artifact&#39;s name | 

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

