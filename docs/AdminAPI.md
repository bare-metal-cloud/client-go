# \AdminAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV1AdminClustersIdInstallsInstall**](AdminAPI.md#DeleteApiV1AdminClustersIdInstallsInstall) | **Delete** /api/v1/admin/clusters/{id}/installs/{install} | Delete one install on a cluster (platform admin)
[**GetAdminExplorerRequests**](AdminAPI.md#GetAdminExplorerRequests) | **Get** /admin/explorer/requests | List explorer access requests (BMC-admin queue)
[**GetAdminExplorerViewas**](AdminAPI.md#GetAdminExplorerViewas) | **Get** /admin/explorer/viewas | View-as session status (BMC-admin)
[**GetApiV1AdminAgents**](AdminAPI.md#GetApiV1AdminAgents) | **Get** /api/v1/admin/agents | List the agent fleet (platform admin)
[**GetApiV1AdminAgentsId**](AdminAPI.md#GetApiV1AdminAgentsId) | **Get** /api/v1/admin/agents/{id} | One agent&#39;s fleet detail (platform admin)
[**GetApiV1AdminAuditEvents**](AdminAPI.md#GetApiV1AdminAuditEvents) | **Get** /api/v1/admin/audit-events | List platform audit events (platform admin)
[**GetApiV1AdminBackupOrgsOrgIdEntitlement**](AdminAPI.md#GetApiV1AdminBackupOrgsOrgIdEntitlement) | **Get** /api/v1/admin/backup/orgs/{org_id}/entitlement | Get an org&#39;s backup entitlement (admin)
[**GetApiV1AdminClusters**](AdminAPI.md#GetApiV1AdminClusters) | **Get** /api/v1/admin/clusters | List clusters across orgs (platform admin)
[**GetApiV1AdminClustersIdInstalls**](AdminAPI.md#GetApiV1AdminClustersIdInstalls) | **Get** /api/v1/admin/clusters/{id}/installs | List one cluster&#39;s installs across orgs (platform admin)
[**GetApiV1AdminDeployTriage**](AdminAPI.md#GetApiV1AdminDeployTriage) | **Get** /api/v1/admin/deploy-triage | Cross-org failed-deploy triage list (platform admin)
[**GetApiV1AdminDeployTriageJobsJobId**](AdminAPI.md#GetApiV1AdminDeployTriageJobsJobId) | **Get** /api/v1/admin/deploy-triage/jobs/{job_id} | Job debugger by id across orgs (platform admin)
[**GetApiV1AdminMonitoringOrgsOrgIdEntitlement**](AdminAPI.md#GetApiV1AdminMonitoringOrgsOrgIdEntitlement) | **Get** /api/v1/admin/monitoring/orgs/{org_id}/entitlement | Get an org&#39;s effective monitoring entitlement (admin)
[**GetApiV1AdminOrgs**](AdminAPI.md#GetApiV1AdminOrgs) | **Get** /api/v1/admin/orgs | List organizations across orgs (platform admin)
[**GetApiV1AdminOverviewFleet**](AdminAPI.md#GetApiV1AdminOverviewFleet) | **Get** /api/v1/admin/overview/fleet | Admin Overview fleet counters (platform admin)
[**GetApiV1AdminOverviewHealth**](AdminAPI.md#GetApiV1AdminOverviewHealth) | **Get** /api/v1/admin/overview/health | Admin Overview health strip (platform admin)
[**GetApiV1AdminOverviewSensors**](AdminAPI.md#GetApiV1AdminOverviewSensors) | **Get** /api/v1/admin/overview/sensors | Admin Overview sensors row (platform admin)
[**GetApiV1AdminOverviewWork**](AdminAPI.md#GetApiV1AdminOverviewWork) | **Get** /api/v1/admin/overview/work | Admin Overview work-in-flight board (platform admin)
[**GetApiV1AdminPlatformRunnerImages**](AdminAPI.md#GetApiV1AdminPlatformRunnerImages) | **Get** /api/v1/admin/platform/runner-images | List the base image catalogue (platform admin)
[**GetApiV1AdminUsers**](AdminAPI.md#GetApiV1AdminUsers) | **Get** /api/v1/admin/users | List users across orgs (platform admin)
[**PostAdminExplorerGrantsIdApprove**](AdminAPI.md#PostAdminExplorerGrantsIdApprove) | **Post** /admin/explorer/grants/{id}/approve | Approve an explorer access request (BMC-admin)
[**PostAdminExplorerGrantsIdDeny**](AdminAPI.md#PostAdminExplorerGrantsIdDeny) | **Post** /admin/explorer/grants/{id}/deny | Deny an explorer access request (BMC-admin)
[**PostAdminExplorerGrantsIdRevoke**](AdminAPI.md#PostAdminExplorerGrantsIdRevoke) | **Post** /admin/explorer/grants/{id}/revoke | Revoke an active explorer grant (BMC-admin)
[**PostAdminExplorerViewas**](AdminAPI.md#PostAdminExplorerViewas) | **Post** /admin/explorer/viewas | Start a view-as-client session (BMC-admin)
[**PostAdminExplorerViewasStop**](AdminAPI.md#PostAdminExplorerViewasStop) | **Post** /admin/explorer/viewas/stop | Stop the view-as-client session (BMC-admin)
[**PostAdminFeatureGates**](AdminAPI.md#PostAdminFeatureGates) | **Post** /admin/feature-gates | Flip a feature gate (BMC-admin)
[**PostAdminRunnerReservations**](AdminAPI.md#PostAdminRunnerReservations) | **Post** /admin/runner/reservations | Grant a hosted-runner reservation (BMC-admin)
[**PostAdminRunnerReservationsOrgIdRevoke**](AdminAPI.md#PostAdminRunnerReservationsOrgIdRevoke) | **Post** /admin/runner/reservations/{org_id}/revoke | Revoke a hosted-runner reservation (BMC-admin)
[**PostApiV1AdminAgentsIdForceDelete**](AdminAPI.md#PostApiV1AdminAgentsIdForceDelete) | **Post** /api/v1/admin/agents/{id}/force-delete | Force-delete an agent (platform admin) — revoke + archive
[**PostApiV1AdminUsersIdMfaDisable**](AdminAPI.md#PostApiV1AdminUsersIdMfaDisable) | **Post** /api/v1/admin/users/{id}/mfa/disable | Admin force-disable user MFA
[**PutApiV1AdminBackupOrgsOrgIdEntitlement**](AdminAPI.md#PutApiV1AdminBackupOrgsOrgIdEntitlement) | **Put** /api/v1/admin/backup/orgs/{org_id}/entitlement | Set or clear an org&#39;s backup entitlement (admin)
[**PutApiV1AdminMonitoringOrgsOrgIdExemption**](AdminAPI.md#PutApiV1AdminMonitoringOrgsOrgIdExemption) | **Put** /api/v1/admin/monitoring/orgs/{org_id}/exemption | Set or clear an org&#39;s enforcement exemption (admin)
[**PutApiV1AdminMonitoringOrgsOrgIdLimits**](AdminAPI.md#PutApiV1AdminMonitoringOrgsOrgIdLimits) | **Put** /api/v1/admin/monitoring/orgs/{org_id}/limits | Set an org&#39;s monitoring limits (admin set-path)



## DeleteApiV1AdminClustersIdInstallsInstall

> map[string]interface{} DeleteApiV1AdminClustersIdInstallsInstall(ctx, id, install).Body(body).Execute()

Delete one install on a cluster (platform admin)



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersAdminInstallDeleteBody() // HandlersAdminInstallDeleteBody | Delete reason

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.DeleteApiV1AdminClustersIdInstallsInstall(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.DeleteApiV1AdminClustersIdInstallsInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApiV1AdminClustersIdInstallsInstall`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.DeleteApiV1AdminClustersIdInstallsInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1AdminClustersIdInstallsInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersAdminInstallDeleteBody**](HandlersAdminInstallDeleteBody.md) | Delete reason | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminExplorerRequests

> []HandlersExplorerGrantView GetAdminExplorerRequests(ctx).Status(status).Execute()

List explorer access requests (BMC-admin queue)



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
	status := "status_example" // string | pending (default) | requested | active | expired | revoked | denied | all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetAdminExplorerRequests(context.Background()).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminExplorerRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminExplorerRequests`: []HandlersExplorerGrantView
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminExplorerRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminExplorerRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | pending (default) | requested | active | expired | revoked | denied | all | 

### Return type

[**[]HandlersExplorerGrantView**](HandlersExplorerGrantView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminExplorerViewas

> map[string]interface{} GetAdminExplorerViewas(ctx).Execute()

View-as session status (BMC-admin)



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
	resp, r, err := apiClient.AdminAPI.GetAdminExplorerViewas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminExplorerViewas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminExplorerViewas`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminExplorerViewas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminExplorerViewasRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminAgents

> map[string]interface{} GetApiV1AdminAgents(ctx).Page(page).Limit(limit).Status(status).Search(search).Execute()

List the agent fleet (platform admin)



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
	page := int32(56) // int32 | 1-based page (default 1) (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	status := "status_example" // string | Filter by one lifecycle state (optional)
	search := "search_example" // string | Match org name, VM name, routing key, or agent id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminAgents(context.Background()).Page(page).Limit(limit).Status(status).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminAgents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | 1-based page (default 1) | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **status** | **string** | Filter by one lifecycle state | 
 **search** | **string** | Match org name, VM name, routing key, or agent id | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminAgentsId

> map[string]interface{} GetApiV1AdminAgentsId(ctx, id).Execute()

One agent's fleet detail (platform admin)



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
	id := int32(56) // int32 | Agent id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminAgentsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminAgentsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminAgentsId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminAgentsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Agent id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminAgentsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminAuditEvents

> map[string]interface{} GetApiV1AdminAuditEvents(ctx).ActorUserId(actorUserId).Action(action).ResourceType(resourceType).Outcome(outcome).SubjectOrgId(subjectOrgId).SubjectUserId(subjectUserId).Reason(reason).From(from).To(to).Cursor(cursor).Limit(limit).Execute()

List platform audit events (platform admin)



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
	actorUserId := int32(56) // int32 | Filter by actor user id (both queries) (optional)
	action := "action_example" // string | Filter by action (optional)
	resourceType := "resourceType_example" // string | Filter by resource_type (optional)
	outcome := "outcome_example" // string | Filter by outcome (success|error|denied|unconfirmed) (optional)
	subjectOrgId := int32(56) // int32 | Filter by subject org (platform rows) (optional)
	subjectUserId := int32(56) // int32 | Filter by impersonated subject (impersonated rows) (optional)
	reason := "reason_example" // string | Substring match on the impersonated reason (optional)
	from := "from_example" // string | Inclusive lower bound (RFC3339) (optional)
	to := "to_example" // string | Inclusive upper bound (RFC3339) (optional)
	cursor := "cursor_example" // string | Opaque pagination cursor (platform query) (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminAuditEvents(context.Background()).ActorUserId(actorUserId).Action(action).ResourceType(resourceType).Outcome(outcome).SubjectOrgId(subjectOrgId).SubjectUserId(subjectUserId).Reason(reason).From(from).To(to).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminAuditEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminAuditEvents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminAuditEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminAuditEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actorUserId** | **int32** | Filter by actor user id (both queries) | 
 **action** | **string** | Filter by action | 
 **resourceType** | **string** | Filter by resource_type | 
 **outcome** | **string** | Filter by outcome (success|error|denied|unconfirmed) | 
 **subjectOrgId** | **int32** | Filter by subject org (platform rows) | 
 **subjectUserId** | **int32** | Filter by impersonated subject (impersonated rows) | 
 **reason** | **string** | Substring match on the impersonated reason | 
 **from** | **string** | Inclusive lower bound (RFC3339) | 
 **to** | **string** | Inclusive upper bound (RFC3339) | 
 **cursor** | **string** | Opaque pagination cursor (platform query) | 
 **limit** | **int32** | Page size (default 50, max 200) | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminBackupOrgsOrgIdEntitlement

> map[string]interface{} GetApiV1AdminBackupOrgsOrgIdEntitlement(ctx, orgId).Execute()

Get an org's backup entitlement (admin)



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
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminBackupOrgsOrgIdEntitlement(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminBackupOrgsOrgIdEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminBackupOrgsOrgIdEntitlement`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminBackupOrgsOrgIdEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminBackupOrgsOrgIdEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminClusters

> map[string]interface{} GetApiV1AdminClusters(ctx).Execute()

List clusters across orgs (platform admin)



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
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminClusters`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminClustersRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminClustersIdInstalls

> map[string]interface{} GetApiV1AdminClustersIdInstalls(ctx, id).Execute()

List one cluster's installs across orgs (platform admin)



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
	id := "id_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminClustersIdInstalls(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminClustersIdInstalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminClustersIdInstalls`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminClustersIdInstalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminClustersIdInstallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminDeployTriage

> map[string]interface{} GetApiV1AdminDeployTriage(ctx).Limit(limit).Execute()

Cross-org failed-deploy triage list (platform admin)



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
	limit := int32(56) // int32 | Page size (default 100, max 500) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminDeployTriage(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminDeployTriage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminDeployTriage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminDeployTriage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminDeployTriageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Page size (default 100, max 500) | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminDeployTriageJobsJobId

> map[string]interface{} GetApiV1AdminDeployTriageJobsJobId(ctx, jobId).Execute()

Job debugger by id across orgs (platform admin)



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
	jobId := "jobId_example" // string | Job id (UUID)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminDeployTriageJobsJobId(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminDeployTriageJobsJobId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminDeployTriageJobsJobId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminDeployTriageJobsJobId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminDeployTriageJobsJobIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminMonitoringOrgsOrgIdEntitlement

> map[string]interface{} GetApiV1AdminMonitoringOrgsOrgIdEntitlement(ctx, orgId).Execute()

Get an org's effective monitoring entitlement (admin)



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
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminMonitoringOrgsOrgIdEntitlement(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminMonitoringOrgsOrgIdEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminMonitoringOrgsOrgIdEntitlement`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminMonitoringOrgsOrgIdEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminMonitoringOrgsOrgIdEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminOrgs

> map[string]interface{} GetApiV1AdminOrgs(ctx).Search(search).Cursor(cursor).Limit(limit).Execute()

List organizations across orgs (platform admin)



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
	search := "search_example" // string | Substring search (org name, org slug, member email) (optional)
	cursor := "cursor_example" // string | Opaque keyset cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminOrgs(context.Background()).Search(search).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminOrgs`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminOrgs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminOrgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Substring search (org name, org slug, member email) | 
 **cursor** | **string** | Opaque keyset cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminOverviewFleet

> GetApiV1AdminOverviewFleet200Response GetApiV1AdminOverviewFleet(ctx).Execute()

Admin Overview fleet counters (platform admin)



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
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminOverviewFleet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminOverviewFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminOverviewFleet`: GetApiV1AdminOverviewFleet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminOverviewFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminOverviewFleetRequest struct via the builder pattern


### Return type

[**GetApiV1AdminOverviewFleet200Response**](GetApiV1AdminOverviewFleet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminOverviewHealth

> GetApiV1AdminOverviewHealth200Response GetApiV1AdminOverviewHealth(ctx).Execute()

Admin Overview health strip (platform admin)



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
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminOverviewHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminOverviewHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminOverviewHealth`: GetApiV1AdminOverviewHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminOverviewHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminOverviewHealthRequest struct via the builder pattern


### Return type

[**GetApiV1AdminOverviewHealth200Response**](GetApiV1AdminOverviewHealth200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminOverviewSensors

> GetApiV1AdminOverviewSensors200Response GetApiV1AdminOverviewSensors(ctx).Execute()

Admin Overview sensors row (platform admin)



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
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminOverviewSensors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminOverviewSensors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminOverviewSensors`: GetApiV1AdminOverviewSensors200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminOverviewSensors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminOverviewSensorsRequest struct via the builder pattern


### Return type

[**GetApiV1AdminOverviewSensors200Response**](GetApiV1AdminOverviewSensors200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminOverviewWork

> GetApiV1AdminOverviewWork200Response GetApiV1AdminOverviewWork(ctx).Execute()

Admin Overview work-in-flight board (platform admin)



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
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminOverviewWork(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminOverviewWork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminOverviewWork`: GetApiV1AdminOverviewWork200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminOverviewWork`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminOverviewWorkRequest struct via the builder pattern


### Return type

[**GetApiV1AdminOverviewWork200Response**](GetApiV1AdminOverviewWork200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminPlatformRunnerImages

> HandlersRunnerImagesResponse GetApiV1AdminPlatformRunnerImages(ctx).Execute()

List the base image catalogue (platform admin)



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
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminPlatformRunnerImages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminPlatformRunnerImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminPlatformRunnerImages`: HandlersRunnerImagesResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminPlatformRunnerImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminPlatformRunnerImagesRequest struct via the builder pattern


### Return type

[**HandlersRunnerImagesResponse**](HandlersRunnerImagesResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1AdminUsers

> map[string]interface{} GetApiV1AdminUsers(ctx).Search(search).Cursor(cursor).Limit(limit).Execute()

List users across orgs (platform admin)



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
	search := "search_example" // string | Substring search on user email (optional)
	cursor := "cursor_example" // string | Opaque keyset cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.GetApiV1AdminUsers(context.Background()).Search(search).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetApiV1AdminUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1AdminUsers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetApiV1AdminUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1AdminUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Substring search on user email | 
 **cursor** | **string** | Opaque keyset cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminExplorerGrantsIdApprove

> HandlersExplorerGrantView PostAdminExplorerGrantsIdApprove(ctx, id).Body(body).Execute()

Approve an explorer access request (BMC-admin)



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
	id := int32(56) // int32 | Grant ID
	body := map[string]interface{}{ ... } // map[string]interface{} | {\\ (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminExplorerGrantsIdApprove(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminExplorerGrantsIdApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminExplorerGrantsIdApprove`: HandlersExplorerGrantView
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminExplorerGrantsIdApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminExplorerGrantsIdApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | {\\ | 

### Return type

[**HandlersExplorerGrantView**](HandlersExplorerGrantView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminExplorerGrantsIdDeny

> HandlersExplorerGrantView PostAdminExplorerGrantsIdDeny(ctx, id).Body(body).Execute()

Deny an explorer access request (BMC-admin)



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
	id := int32(56) // int32 | Grant ID
	body := map[string]interface{}{ ... } // map[string]interface{} | {\\ (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminExplorerGrantsIdDeny(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminExplorerGrantsIdDeny``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminExplorerGrantsIdDeny`: HandlersExplorerGrantView
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminExplorerGrantsIdDeny`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminExplorerGrantsIdDenyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | {\\ | 

### Return type

[**HandlersExplorerGrantView**](HandlersExplorerGrantView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminExplorerGrantsIdRevoke

> HandlersExplorerGrantView PostAdminExplorerGrantsIdRevoke(ctx, id).Body(body).Execute()

Revoke an active explorer grant (BMC-admin)



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
	id := int32(56) // int32 | Grant ID
	body := map[string]interface{}{ ... } // map[string]interface{} | {\\ (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminExplorerGrantsIdRevoke(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminExplorerGrantsIdRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminExplorerGrantsIdRevoke`: HandlersExplorerGrantView
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminExplorerGrantsIdRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminExplorerGrantsIdRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | {\\ | 

### Return type

[**HandlersExplorerGrantView**](HandlersExplorerGrantView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminExplorerViewas

> map[string]interface{} PostAdminExplorerViewas(ctx).Body(body).Execute()

Start a view-as-client session (BMC-admin)



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
	body := map[string]interface{}{ ... } // map[string]interface{} | {\\

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminExplorerViewas(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminExplorerViewas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminExplorerViewas`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminExplorerViewas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminExplorerViewasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | {\\ | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminExplorerViewasStop

> map[string]interface{} PostAdminExplorerViewasStop(ctx).Execute()

Stop the view-as-client session (BMC-admin)



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
	resp, r, err := apiClient.AdminAPI.PostAdminExplorerViewasStop(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminExplorerViewasStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminExplorerViewasStop`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminExplorerViewasStop`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminExplorerViewasStopRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminFeatureGates

> HandlersFeatureGateFlipResponse PostAdminFeatureGates(ctx).Body(body).Execute()

Flip a feature gate (BMC-admin)



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
	body := *openapiclient.NewHandlersFeatureGateFlipRequest() // HandlersFeatureGateFlipRequest | The flip

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminFeatureGates(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminFeatureGates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminFeatureGates`: HandlersFeatureGateFlipResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminFeatureGates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminFeatureGatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersFeatureGateFlipRequest**](HandlersFeatureGateFlipRequest.md) | The flip | 

### Return type

[**HandlersFeatureGateFlipResponse**](HandlersFeatureGateFlipResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminRunnerReservations

> HandlersRunnerReservationResponse PostAdminRunnerReservations(ctx).Body(body).Execute()

Grant a hosted-runner reservation (BMC-admin)



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
	body := *openapiclient.NewHandlersRunnerReservationGrantRequest() // HandlersRunnerReservationGrantRequest | The grant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminRunnerReservations(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminRunnerReservations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminRunnerReservations`: HandlersRunnerReservationResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminRunnerReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminRunnerReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlersRunnerReservationGrantRequest**](HandlersRunnerReservationGrantRequest.md) | The grant | 

### Return type

[**HandlersRunnerReservationResponse**](HandlersRunnerReservationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdminRunnerReservationsOrgIdRevoke

> HandlersRunnerReservationResponse PostAdminRunnerReservationsOrgIdRevoke(ctx, orgId).Body(body).Execute()

Revoke a hosted-runner reservation (BMC-admin)



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
	body := *openapiclient.NewHandlersRunnerReservationRevokeRequest() // HandlersRunnerReservationRevokeRequest | The revoke

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostAdminRunnerReservationsOrgIdRevoke(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostAdminRunnerReservationsOrgIdRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminRunnerReservationsOrgIdRevoke`: HandlersRunnerReservationResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostAdminRunnerReservationsOrgIdRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminRunnerReservationsOrgIdRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRunnerReservationRevokeRequest**](HandlersRunnerReservationRevokeRequest.md) | The revoke | 

### Return type

[**HandlersRunnerReservationResponse**](HandlersRunnerReservationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1AdminAgentsIdForceDelete

> map[string]interface{} PostApiV1AdminAgentsIdForceDelete(ctx, id).Body(body).Execute()

Force-delete an agent (platform admin) — revoke + archive



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
	id := int32(56) // int32 | Agent DB id
	body := map[string]interface{}{ ... } // map[string]interface{} | { \\

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostApiV1AdminAgentsIdForceDelete(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostApiV1AdminAgentsIdForceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1AdminAgentsIdForceDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostApiV1AdminAgentsIdForceDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Agent DB id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1AdminAgentsIdForceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | { \\ | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1AdminUsersIdMfaDisable

> HandlersMFADisableResponse PostApiV1AdminUsersIdMfaDisable(ctx, id).Execute()

Admin force-disable user MFA



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
	id := int32(56) // int32 | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PostApiV1AdminUsersIdMfaDisable(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PostApiV1AdminUsersIdMfaDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1AdminUsersIdMfaDisable`: HandlersMFADisableResponse
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PostApiV1AdminUsersIdMfaDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1AdminUsersIdMfaDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersMFADisableResponse**](HandlersMFADisableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiV1AdminBackupOrgsOrgIdEntitlement

> map[string]interface{} PutApiV1AdminBackupOrgsOrgIdEntitlement(ctx, orgId).Request(request).Execute()

Set or clear an org's backup entitlement (admin)



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
	request := *openapiclient.NewHandlersSetBackupEntitlementRequest() // HandlersSetBackupEntitlementRequest | Entitlement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PutApiV1AdminBackupOrgsOrgIdEntitlement(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PutApiV1AdminBackupOrgsOrgIdEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApiV1AdminBackupOrgsOrgIdEntitlement`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PutApiV1AdminBackupOrgsOrgIdEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1AdminBackupOrgsOrgIdEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersSetBackupEntitlementRequest**](HandlersSetBackupEntitlementRequest.md) | Entitlement | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiV1AdminMonitoringOrgsOrgIdExemption

> map[string]interface{} PutApiV1AdminMonitoringOrgsOrgIdExemption(ctx, orgId).Request(request).Execute()

Set or clear an org's enforcement exemption (admin)



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
	request := *openapiclient.NewHandlersSetExemptionRequest() // HandlersSetExemptionRequest | Exemption

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PutApiV1AdminMonitoringOrgsOrgIdExemption(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PutApiV1AdminMonitoringOrgsOrgIdExemption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApiV1AdminMonitoringOrgsOrgIdExemption`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PutApiV1AdminMonitoringOrgsOrgIdExemption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1AdminMonitoringOrgsOrgIdExemptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersSetExemptionRequest**](HandlersSetExemptionRequest.md) | Exemption | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiV1AdminMonitoringOrgsOrgIdLimits

> map[string]interface{} PutApiV1AdminMonitoringOrgsOrgIdLimits(ctx, orgId).Request(request).Execute()

Set an org's monitoring limits (admin set-path)



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
	request := *openapiclient.NewHandlersSetMonitoringLimitsRequest() // HandlersSetMonitoringLimitsRequest | Limits

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PutApiV1AdminMonitoringOrgsOrgIdLimits(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PutApiV1AdminMonitoringOrgsOrgIdLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApiV1AdminMonitoringOrgsOrgIdLimits`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PutApiV1AdminMonitoringOrgsOrgIdLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1AdminMonitoringOrgsOrgIdLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersSetMonitoringLimitsRequest**](HandlersSetMonitoringLimitsRequest.md) | Limits | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

