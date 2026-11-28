# \BuildsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdBuildAttemptsAttemptIdConsole**](BuildsAPI.md#GetOrgsOrgIdBuildAttemptsAttemptIdConsole) | **Get** /orgs/{org_id}/build-attempts/{attempt_id}/console | Read a build attempt&#39;s console (plan steps, evidence, log tail)
[**GetOrgsOrgIdBuildAttemptsAttemptIdLogs**](BuildsAPI.md#GetOrgsOrgIdBuildAttemptsAttemptIdLogs) | **Get** /orgs/{org_id}/build-attempts/{attempt_id}/logs | Download a build attempt&#39;s full log
[**GetOrgsOrgIdStacksRevisions**](BuildsAPI.md#GetOrgsOrgIdStacksRevisions) | **Get** /orgs/{org_id}/stacks/revisions | Read a stack&#39;s deploy revision ledger (per-service provenance)
[**PostOrgsOrgIdBuildAttemptsAttemptIdCancel**](BuildsAPI.md#PostOrgsOrgIdBuildAttemptsAttemptIdCancel) | **Post** /orgs/{org_id}/build-attempts/{attempt_id}/cancel | Cancel a build attempt (the deploys ledger&#39;s cancel verb)
[**PostOrgsOrgIdStacksBuildTimeline**](BuildsAPI.md#PostOrgsOrgIdStacksBuildTimeline) | **Post** /orgs/{org_id}/stacks/build-timeline | Read a stack deploy&#39;s timeline (proven states, dependency order)
[**PostOrgsOrgIdStacksDeploy**](BuildsAPI.md#PostOrgsOrgIdStacksDeploy) | **Post** /orgs/{org_id}/stacks/deploy | Trigger a bound build manually (the CLI deploy verb)
[**PostOrgsOrgIdStacksRollback**](BuildsAPI.md#PostOrgsOrgIdStacksRollback) | **Post** /orgs/{org_id}/stacks/rollback | Roll a stack service back to an older registered revision



## GetOrgsOrgIdBuildAttemptsAttemptIdConsole

> BuilderConsoleView GetOrgsOrgIdBuildAttemptsAttemptIdConsole(ctx, orgId, attemptId).Execute()

Read a build attempt's console (plan steps, evidence, log tail)



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
	orgId := int32(56) // int32 | The org id
	attemptId := int32(56) // int32 | The build attempt's ledger row id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.GetOrgsOrgIdBuildAttemptsAttemptIdConsole(context.Background(), orgId, attemptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.GetOrgsOrgIdBuildAttemptsAttemptIdConsole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdBuildAttemptsAttemptIdConsole`: BuilderConsoleView
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.GetOrgsOrgIdBuildAttemptsAttemptIdConsole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 
**attemptId** | **int32** | The build attempt&#39;s ledger row id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdBuildAttemptsAttemptIdConsoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BuilderConsoleView**](BuilderConsoleView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdBuildAttemptsAttemptIdLogs

> string GetOrgsOrgIdBuildAttemptsAttemptIdLogs(ctx, orgId, attemptId).Execute()

Download a build attempt's full log



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
	orgId := int32(56) // int32 | The org id
	attemptId := int32(56) // int32 | The build attempt's ledger row id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.GetOrgsOrgIdBuildAttemptsAttemptIdLogs(context.Background(), orgId, attemptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.GetOrgsOrgIdBuildAttemptsAttemptIdLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdBuildAttemptsAttemptIdLogs`: string
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.GetOrgsOrgIdBuildAttemptsAttemptIdLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 
**attemptId** | **int32** | The build attempt&#39;s ledger row id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdBuildAttemptsAttemptIdLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdStacksRevisions

> HandlersBuildLedgerView GetOrgsOrgIdStacksRevisions(ctx, orgId).ClusterId(clusterId).Stack(stack).Execute()

Read a stack's deploy revision ledger (per-service provenance)



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
	orgId := int32(56) // int32 | The org id
	clusterId := "clusterId_example" // string | The cluster id
	stack := "stack_example" // string | The stack name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.GetOrgsOrgIdStacksRevisions(context.Background(), orgId).ClusterId(clusterId).Stack(stack).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.GetOrgsOrgIdStacksRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdStacksRevisions`: HandlersBuildLedgerView
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.GetOrgsOrgIdStacksRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdStacksRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterId** | **string** | The cluster id | 
 **stack** | **string** | The stack name | 

### Return type

[**HandlersBuildLedgerView**](HandlersBuildLedgerView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdBuildAttemptsAttemptIdCancel

> HandlersBuildCancelView PostOrgsOrgIdBuildAttemptsAttemptIdCancel(ctx, orgId, attemptId).Execute()

Cancel a build attempt (the deploys ledger's cancel verb)



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
	orgId := int32(56) // int32 | The org id
	attemptId := int32(56) // int32 | The build attempt's ledger row id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.PostOrgsOrgIdBuildAttemptsAttemptIdCancel(context.Background(), orgId, attemptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.PostOrgsOrgIdBuildAttemptsAttemptIdCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdBuildAttemptsAttemptIdCancel`: HandlersBuildCancelView
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.PostOrgsOrgIdBuildAttemptsAttemptIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 
**attemptId** | **int32** | The build attempt&#39;s ledger row id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdBuildAttemptsAttemptIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersBuildCancelView**](HandlersBuildCancelView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdStacksBuildTimeline

> BuilderTimelineView PostOrgsOrgIdStacksBuildTimeline(ctx, orgId).Body(body).Execute()

Read a stack deploy's timeline (proven states, dependency order)



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
	orgId := int32(56) // int32 | The org id
	body := *openapiclient.NewHandlersTimelineRequest() // HandlersTimelineRequest | The deploy identity and the service manifest in dependency order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.PostOrgsOrgIdStacksBuildTimeline(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.PostOrgsOrgIdStacksBuildTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdStacksBuildTimeline`: BuilderTimelineView
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.PostOrgsOrgIdStacksBuildTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdStacksBuildTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersTimelineRequest**](HandlersTimelineRequest.md) | The deploy identity and the service manifest in dependency order | 

### Return type

[**BuilderTimelineView**](BuilderTimelineView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdStacksDeploy

> HandlersBuildDeployView PostOrgsOrgIdStacksDeploy(ctx, orgId).Body(body).Execute()

Trigger a bound build manually (the CLI deploy verb)



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
	orgId := int32(56) // int32 | The org id
	body := *openapiclient.NewHandlersBuildDeployRequest() // HandlersBuildDeployRequest | The binding identity (ref optional: branch, tag, or SHA)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.PostOrgsOrgIdStacksDeploy(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.PostOrgsOrgIdStacksDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdStacksDeploy`: HandlersBuildDeployView
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.PostOrgsOrgIdStacksDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdStacksDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersBuildDeployRequest**](HandlersBuildDeployRequest.md) | The binding identity (ref optional: branch, tag, or SHA) | 

### Return type

[**HandlersBuildDeployView**](HandlersBuildDeployView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdStacksRollback

> BuilderRollbackOutcome PostOrgsOrgIdStacksRollback(ctx, orgId).Body(body).Execute()

Roll a stack service back to an older registered revision



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
	orgId := int32(56) // int32 | The org id
	body := *openapiclient.NewHandlersStackRollbackRequest() // HandlersStackRollbackRequest | The deploy identity and the target revision id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.PostOrgsOrgIdStacksRollback(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.PostOrgsOrgIdStacksRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdStacksRollback`: BuilderRollbackOutcome
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.PostOrgsOrgIdStacksRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdStacksRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersStackRollbackRequest**](HandlersStackRollbackRequest.md) | The deploy identity and the target revision id | 

### Return type

[**BuilderRollbackOutcome**](BuilderRollbackOutcome.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

