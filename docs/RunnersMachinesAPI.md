# \RunnersMachinesAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersEvents**](RunnersMachinesAPI.md#GetOrgsOrgIdRunnersEvents) | **Get** /orgs/{org_id}/runners/events | The org&#39;s events feed — every machine&#39;s transitions beside the org&#39;s workflow transitions
[**GetOrgsOrgIdRunnersMachines**](RunnersMachinesAPI.md#GetOrgsOrgIdRunnersMachines) | **Get** /orgs/{org_id}/runners/machines | List the org&#39;s runner machines — bound runners AND unbound VMs
[**GetOrgsOrgIdRunnersMachinesMachineId**](RunnersMachinesAPI.md#GetOrgsOrgIdRunnersMachinesMachineId) | **Get** /orgs/{org_id}/runners/machines/{machine_id} | One machine&#39;s detail (identity and health reads ride it)
[**GetOrgsOrgIdRunnersMachinesMachineIdCleanup**](RunnersMachinesAPI.md#GetOrgsOrgIdRunnersMachinesMachineIdCleanup) | **Get** /orgs/{org_id}/runners/machines/{machine_id}/cleanup | One machine&#39;s auto disk-cleanup setting
[**GetOrgsOrgIdRunnersMachinesMachineIdEvents**](RunnersMachinesAPI.md#GetOrgsOrgIdRunnersMachinesMachineIdEvents) | **Get** /orgs/{org_id}/runners/machines/{machine_id}/events | One machine&#39;s events feed — the workflow transitions and the lifecycle transitions, unified
[**GetOrgsOrgIdRunnersMachinesMachineIdProfileEligibility**](RunnersMachinesAPI.md#GetOrgsOrgIdRunnersMachinesMachineIdProfileEligibility) | **Get** /orgs/{org_id}/runners/machines/{machine_id}/profile-eligibility | One machine&#39;s eligibility for every runner profile — the engine&#39;s verdicts with reasons
[**PostOrgsOrgIdRunnersMachines**](RunnersMachinesAPI.md#PostOrgsOrgIdRunnersMachines) | **Post** /orgs/{org_id}/runners/machines | Raw machine creation — typed refusal (the install path is the only join)
[**PostOrgsOrgIdRunnersMachinesMachineIdAdd**](RunnersMachinesAPI.md#PostOrgsOrgIdRunnersMachinesMachineIdAdd) | **Post** /orgs/{org_id}/runners/machines/{machine_id}/add | Add a detected machine as a runner (create the binding)
[**PostOrgsOrgIdRunnersMachinesMachineIdQuarantine**](RunnersMachinesAPI.md#PostOrgsOrgIdRunnersMachinesMachineIdQuarantine) | **Post** /orgs/{org_id}/runners/machines/{machine_id}/quarantine | Quarantine a runner — operator-only (ruling 13, test-marked until AAR.4)
[**PostOrgsOrgIdRunnersMachinesMachineIdRegenerateMachineId**](RunnersMachinesAPI.md#PostOrgsOrgIdRunnersMachinesMachineIdRegenerateMachineId) | **Post** /orgs/{org_id}/runners/machines/{machine_id}/regenerate-machine-id | Regenerate a machine&#39;s machine-id (operator-only, typed confirmation)
[**PostOrgsOrgIdRunnersMachinesMachineIdRelease**](RunnersMachinesAPI.md#PostOrgsOrgIdRunnersMachinesMachineIdRelease) | **Post** /orgs/{org_id}/runners/machines/{machine_id}/release | Release a runner (untag the binding — immediately, or through a drain)
[**PostOrgsOrgIdRunnersMachinesMachineIdReturn**](RunnersMachinesAPI.md#PostOrgsOrgIdRunnersMachinesMachineIdReturn) | **Post** /orgs/{org_id}/runners/machines/{machine_id}/return | Return a quarantined runner to service — operator-only
[**PostOrgsOrgIdRunnersMachinesMachineIdRotateCredential**](RunnersMachinesAPI.md#PostOrgsOrgIdRunnersMachinesMachineIdRotateCredential) | **Post** /orgs/{org_id}/runners/machines/{machine_id}/rotate-credential | Rotate a runner&#39;s machine credential (operator-only, show-once)
[**PutOrgsOrgIdRunnersMachinesMachineIdCleanup**](RunnersMachinesAPI.md#PutOrgsOrgIdRunnersMachinesMachineIdCleanup) | **Put** /orgs/{org_id}/runners/machines/{machine_id}/cleanup | Set a machine&#39;s auto disk-cleanup threshold (operator-only)



## GetOrgsOrgIdRunnersEvents

> HandlersRunnerMachineEventsResponse GetOrgsOrgIdRunnersEvents(ctx, orgId).Repo(repo).Run(run).Action(action).Limit(limit).Execute()

The org's events feed — every machine's transitions beside the org's workflow transitions



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
	repo := "repo_example" // string | Exact repository-name filter (workflow rows) (optional)
	run := int32(56) // int32 | Exact workflow-run-id filter (workflow rows) (optional)
	action := "action_example" // string | Exact action filter (the row's action verb) (optional)
	limit := int32(56) // int32 | Maximum rows (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.GetOrgsOrgIdRunnersEvents(context.Background(), orgId).Repo(repo).Run(run).Action(action).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.GetOrgsOrgIdRunnersEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersEvents`: HandlersRunnerMachineEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.GetOrgsOrgIdRunnersEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repo** | **string** | Exact repository-name filter (workflow rows) | 
 **run** | **int32** | Exact workflow-run-id filter (workflow rows) | 
 **action** | **string** | Exact action filter (the row&#39;s action verb) | 
 **limit** | **int32** | Maximum rows (default 50, max 200) | 

### Return type

[**HandlersRunnerMachineEventsResponse**](HandlersRunnerMachineEventsResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersMachines

> HandlersRunnerMachinesResponse GetOrgsOrgIdRunnersMachines(ctx, orgId).Execute()

List the org's runner machines — bound runners AND unbound VMs

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
	resp, r, err := apiClient.RunnersMachinesAPI.GetOrgsOrgIdRunnersMachines(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.GetOrgsOrgIdRunnersMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersMachines`: HandlersRunnerMachinesResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.GetOrgsOrgIdRunnersMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerMachinesResponse**](HandlersRunnerMachinesResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersMachinesMachineId

> HandlersRunnerMachineDetailResponse GetOrgsOrgIdRunnersMachinesMachineId(ctx, orgId, machineId).Execute()

One machine's detail (identity and health reads ride it)

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
	machineId := int32(56) // int32 | Machine ID (the agent row's id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineId(context.Background(), orgId, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersMachinesMachineId`: HandlersRunnerMachineDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**machineId** | **int32** | Machine ID (the agent row&#39;s id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersMachinesMachineIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerMachineDetailResponse**](HandlersRunnerMachineDetailResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersMachinesMachineIdCleanup

> HandlersRunnerMachineCleanupResponse GetOrgsOrgIdRunnersMachinesMachineIdCleanup(ctx, orgId, machineId).Execute()

One machine's auto disk-cleanup setting

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
	machineId := int32(56) // int32 | Machine ID (the agent row's id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineIdCleanup(context.Background(), orgId, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineIdCleanup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersMachinesMachineIdCleanup`: HandlersRunnerMachineCleanupResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineIdCleanup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**machineId** | **int32** | Machine ID (the agent row&#39;s id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersMachinesMachineIdCleanupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerMachineCleanupResponse**](HandlersRunnerMachineCleanupResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersMachinesMachineIdEvents

> HandlersRunnerMachineEventsResponse GetOrgsOrgIdRunnersMachinesMachineIdEvents(ctx, orgId, machineId).Repo(repo).Run(run).Action(action).Limit(limit).Execute()

One machine's events feed — the workflow transitions and the lifecycle transitions, unified



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
	machineId := int32(56) // int32 | Machine ID (the agent row's id)
	repo := "repo_example" // string | Exact repository-name filter (workflow rows) (optional)
	run := int32(56) // int32 | Exact workflow-run-id filter (workflow rows) (optional)
	action := "action_example" // string | Exact action filter (the row's action verb) (optional)
	limit := int32(56) // int32 | Maximum rows (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineIdEvents(context.Background(), orgId, machineId).Repo(repo).Run(run).Action(action).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersMachinesMachineIdEvents`: HandlersRunnerMachineEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**machineId** | **int32** | Machine ID (the agent row&#39;s id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersMachinesMachineIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **repo** | **string** | Exact repository-name filter (workflow rows) | 
 **run** | **int32** | Exact workflow-run-id filter (workflow rows) | 
 **action** | **string** | Exact action filter (the row&#39;s action verb) | 
 **limit** | **int32** | Maximum rows (default 50, max 200) | 

### Return type

[**HandlersRunnerMachineEventsResponse**](HandlersRunnerMachineEventsResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersMachinesMachineIdProfileEligibility

> HandlersRunnerProfileEligibilityResponse GetOrgsOrgIdRunnersMachinesMachineIdProfileEligibility(ctx, orgId, machineId).Execute()

One machine's eligibility for every runner profile — the engine's verdicts with reasons



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
	machineId := int32(56) // int32 | Machine ID (the agent row's id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineIdProfileEligibility(context.Background(), orgId, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineIdProfileEligibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersMachinesMachineIdProfileEligibility`: HandlersRunnerProfileEligibilityResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.GetOrgsOrgIdRunnersMachinesMachineIdProfileEligibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**machineId** | **int32** | Machine ID (the agent row&#39;s id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersMachinesMachineIdProfileEligibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerProfileEligibilityResponse**](HandlersRunnerProfileEligibilityResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersMachines

> PostOrgsOrgIdRunnersMachines(ctx, orgId).Execute()

Raw machine creation — typed refusal (the install path is the only join)



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
	r, err := apiClient.RunnersMachinesAPI.PostOrgsOrgIdRunnersMachines(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersMachinesMachineIdAdd

> HandlersRunnerMachineActionResponse PostOrgsOrgIdRunnersMachinesMachineIdAdd(ctx, orgId, machineId).Execute()

Add a detected machine as a runner (create the binding)



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
	machineId := int32(56) // int32 | Machine ID (the agent row's id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdAdd(context.Background(), orgId, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersMachinesMachineIdAdd`: HandlersRunnerMachineActionResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**machineId** | **int32** | Machine ID (the agent row&#39;s id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersMachinesMachineIdAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerMachineActionResponse**](HandlersRunnerMachineActionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersMachinesMachineIdQuarantine

> HandlersRunnerMachineActionResponse PostOrgsOrgIdRunnersMachinesMachineIdQuarantine(ctx, orgId, machineId).Execute()

Quarantine a runner — operator-only (ruling 13, test-marked until AAR.4)



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
	machineId := int32(56) // int32 | Machine ID (the agent row's id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdQuarantine(context.Background(), orgId, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdQuarantine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersMachinesMachineIdQuarantine`: HandlersRunnerMachineActionResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdQuarantine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**machineId** | **int32** | Machine ID (the agent row&#39;s id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersMachinesMachineIdQuarantineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerMachineActionResponse**](HandlersRunnerMachineActionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersMachinesMachineIdRegenerateMachineId

> HandlersRunnerMachineIdentityResponse PostOrgsOrgIdRunnersMachinesMachineIdRegenerateMachineId(ctx, orgId, machineId).Body(body).Execute()

Regenerate a machine's machine-id (operator-only, typed confirmation)



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
	machineId := int32(56) // int32 | Machine ID (the agent row's id)
	body := map[string]interface{}{ ... } // map[string]interface{} | Typed confirmation object: one field, confirm, whose value is regenerate_machine_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdRegenerateMachineId(context.Background(), orgId, machineId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdRegenerateMachineId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersMachinesMachineIdRegenerateMachineId`: HandlersRunnerMachineIdentityResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdRegenerateMachineId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**machineId** | **int32** | Machine ID (the agent row&#39;s id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersMachinesMachineIdRegenerateMachineIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | Typed confirmation object: one field, confirm, whose value is regenerate_machine_id | 

### Return type

[**HandlersRunnerMachineIdentityResponse**](HandlersRunnerMachineIdentityResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersMachinesMachineIdRelease

> HandlersRunnerMachineActionResponse PostOrgsOrgIdRunnersMachinesMachineIdRelease(ctx, orgId, machineId).Force(force).Execute()

Release a runner (untag the binding — immediately, or through a drain)



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
	machineId := int32(56) // int32 | Machine ID (the agent row's id)
	force := map[string]interface{}{ ... } // map[string]interface{} | Optional {\\ (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdRelease(context.Background(), orgId, machineId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersMachinesMachineIdRelease`: HandlersRunnerMachineActionResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**machineId** | **int32** | Machine ID (the agent row&#39;s id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersMachinesMachineIdReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **map[string]interface{}** | Optional {\\ | 

### Return type

[**HandlersRunnerMachineActionResponse**](HandlersRunnerMachineActionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersMachinesMachineIdReturn

> HandlersRunnerMachineActionResponse PostOrgsOrgIdRunnersMachinesMachineIdReturn(ctx, orgId, machineId).Execute()

Return a quarantined runner to service — operator-only



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
	machineId := int32(56) // int32 | Machine ID (the agent row's id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdReturn(context.Background(), orgId, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdReturn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersMachinesMachineIdReturn`: HandlersRunnerMachineActionResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**machineId** | **int32** | Machine ID (the agent row&#39;s id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersMachinesMachineIdReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerMachineActionResponse**](HandlersRunnerMachineActionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersMachinesMachineIdRotateCredential

> HandlersRunnerMachineCredentialRotationResponse PostOrgsOrgIdRunnersMachinesMachineIdRotateCredential(ctx, orgId, machineId).Execute()

Rotate a runner's machine credential (operator-only, show-once)



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
	orgId := int32(56) // int32 | Organization id
	machineId := int32(56) // int32 | The machine id (the agent id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdRotateCredential(context.Background(), orgId, machineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdRotateCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersMachinesMachineIdRotateCredential`: HandlersRunnerMachineCredentialRotationResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.PostOrgsOrgIdRunnersMachinesMachineIdRotateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization id | 
**machineId** | **int32** | The machine id (the agent id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersMachinesMachineIdRotateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerMachineCredentialRotationResponse**](HandlersRunnerMachineCredentialRotationResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdRunnersMachinesMachineIdCleanup

> HandlersRunnerMachineCleanupResponse PutOrgsOrgIdRunnersMachinesMachineIdCleanup(ctx, orgId, machineId).Body(body).Execute()

Set a machine's auto disk-cleanup threshold (operator-only)



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
	machineId := int32(56) // int32 | Machine ID (the agent row's id)
	body := map[string]interface{}{ ... } // map[string]interface{} | {\\

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinesAPI.PutOrgsOrgIdRunnersMachinesMachineIdCleanup(context.Background(), orgId, machineId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinesAPI.PutOrgsOrgIdRunnersMachinesMachineIdCleanup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdRunnersMachinesMachineIdCleanup`: HandlersRunnerMachineCleanupResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinesAPI.PutOrgsOrgIdRunnersMachinesMachineIdCleanup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**machineId** | **int32** | Machine ID (the agent row&#39;s id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdRunnersMachinesMachineIdCleanupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | {\\ | 

### Return type

[**HandlersRunnerMachineCleanupResponse**](HandlersRunnerMachineCleanupResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

