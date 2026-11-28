# \GitopsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClustersIdAppsInstallGitops**](GitopsAPI.md#DeleteClustersIdAppsInstallGitops) | **Delete** /clusters/{id}/apps/{install}/gitops | Disable CI deploys for an app install
[**DeleteClustersIdAppsInstallGitopsRegistration**](GitopsAPI.md#DeleteClustersIdAppsInstallGitopsRegistration) | **Delete** /clusters/{id}/apps/{install}/gitops/registration | Cancel a pending CI registration
[**DeleteOrgsOrgIdGitopsBuildTriggersTriggerId**](GitopsAPI.md#DeleteOrgsOrgIdGitopsBuildTriggersTriggerId) | **Delete** /orgs/{org_id}/gitops/build-triggers/{trigger_id} | Unbind a repo from its project build
[**DeleteOrgsOrgIdGitopsConnectionsId**](GitopsAPI.md#DeleteOrgsOrgIdGitopsConnectionsId) | **Delete** /orgs/{org_id}/gitops/connections/{id} | Revoke a git connection and its deploy tokens
[**GetOrgsOrgIdGitopsBuildTriggers**](GitopsAPI.md#GetOrgsOrgIdGitopsBuildTriggers) | **Get** /orgs/{org_id}/gitops/build-triggers | List the org&#39;s external-git build bindings
[**GetOrgsOrgIdGitopsConnections**](GitopsAPI.md#GetOrgsOrgIdGitopsConnections) | **Get** /orgs/{org_id}/gitops/connections | List the org&#39;s git connections
[**GetOrgsOrgIdGitopsConnectionsIdBranches**](GitopsAPI.md#GetOrgsOrgIdGitopsConnectionsIdBranches) | **Get** /orgs/{org_id}/gitops/connections/{id}/branches | List branches of a repository on a git connection
[**GetOrgsOrgIdGitopsConnectionsIdRepos**](GitopsAPI.md#GetOrgsOrgIdGitopsConnectionsIdRepos) | **Get** /orgs/{org_id}/gitops/connections/{id}/repos | List repositories on a git connection
[**GetOrgsOrgIdGitopsVerifyingMachines**](GitopsAPI.md#GetOrgsOrgIdGitopsVerifyingMachines) | **Get** /orgs/{org_id}/gitops/verifying-machines | List the org&#39;s verifying machines (public delivery keys)
[**PostClustersIdAppsInstallGitopsEnable**](GitopsAPI.md#PostClustersIdAppsInstallGitopsEnable) | **Post** /clusters/{id}/apps/{install}/gitops/enable | Enable CI deploys for an app install
[**PostClustersIdAppsInstallGitopsRegistration**](GitopsAPI.md#PostClustersIdAppsInstallGitopsRegistration) | **Post** /clusters/{id}/apps/{install}/gitops/registration | Create the deploy setup for a repository
[**PostClustersIdAppsInstallGitopsRegistrationRotate**](GitopsAPI.md#PostClustersIdAppsInstallGitopsRegistrationRotate) | **Post** /clusters/{id}/apps/{install}/gitops/registration/rotate | Rotate the registration-class deploy token
[**PostClustersIdGitopsDetect**](GitopsAPI.md#PostClustersIdGitopsDetect) | **Post** /clusters/{id}/gitops/detect | Inspect a repository for the create-from-repo shape
[**PostOrgsOrgIdGitopsConnections**](GitopsAPI.md#PostOrgsOrgIdGitopsConnections) | **Post** /orgs/{org_id}/gitops/connections | Connect a git provider to the org
[**PostOrgsOrgIdGitopsConnectionsIdBuildTrigger**](GitopsAPI.md#PostOrgsOrgIdGitopsConnectionsIdBuildTrigger) | **Post** /orgs/{org_id}/gitops/connections/{id}/build-trigger | Bind a repo to a project&#39;s build (external-git trigger)
[**PostOrgsOrgIdGitopsConnectionsIdCredentials**](GitopsAPI.md#PostOrgsOrgIdGitopsConnectionsIdCredentials) | **Post** /orgs/{org_id}/gitops/connections/{id}/credentials | Replace a git connection&#39;s credentials (reconnect)
[**PostOrgsOrgIdGitopsConnectionsIdTest**](GitopsAPI.md#PostOrgsOrgIdGitopsConnectionsIdTest) | **Post** /orgs/{org_id}/gitops/connections/{id}/test | Re-validate a git connection&#39;s stored credentials



## DeleteClustersIdAppsInstallGitops

> map[string]string DeleteClustersIdAppsInstallGitops(ctx, id, install).Execute()

Disable CI deploys for an app install

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
	id := int32(56) // int32 | Cluster ID
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.DeleteClustersIdAppsInstallGitops(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.DeleteClustersIdAppsInstallGitops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClustersIdAppsInstallGitops`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.DeleteClustersIdAppsInstallGitops`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClustersIdAppsInstallGitopsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClustersIdAppsInstallGitopsRegistration

> map[string]string DeleteClustersIdAppsInstallGitopsRegistration(ctx, id, install).Execute()

Cancel a pending CI registration

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
	id := int32(56) // int32 | Cluster ID
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.DeleteClustersIdAppsInstallGitopsRegistration(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.DeleteClustersIdAppsInstallGitopsRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClustersIdAppsInstallGitopsRegistration`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.DeleteClustersIdAppsInstallGitopsRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClustersIdAppsInstallGitopsRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgsOrgIdGitopsBuildTriggersTriggerId

> map[string]bool DeleteOrgsOrgIdGitopsBuildTriggersTriggerId(ctx, orgId, triggerId).Execute()

Unbind a repo from its project build



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
	triggerId := int32(56) // int32 | The build trigger id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.DeleteOrgsOrgIdGitopsBuildTriggersTriggerId(context.Background(), orgId, triggerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.DeleteOrgsOrgIdGitopsBuildTriggersTriggerId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdGitopsBuildTriggersTriggerId`: map[string]bool
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.DeleteOrgsOrgIdGitopsBuildTriggersTriggerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 
**triggerId** | **int32** | The build trigger id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdGitopsBuildTriggersTriggerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]bool**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgsOrgIdGitopsConnectionsId

> map[string]string DeleteOrgsOrgIdGitopsConnectionsId(ctx, orgId, id).Execute()

Revoke a git connection and its deploy tokens

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
	id := int32(56) // int32 | Connection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.DeleteOrgsOrgIdGitopsConnectionsId(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.DeleteOrgsOrgIdGitopsConnectionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdGitopsConnectionsId`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.DeleteOrgsOrgIdGitopsConnectionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdGitopsConnectionsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdGitopsBuildTriggers

> []HandlersGitBuildTriggerRow GetOrgsOrgIdGitopsBuildTriggers(ctx, orgId).Execute()

List the org's external-git build bindings



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.GetOrgsOrgIdGitopsBuildTriggers(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.GetOrgsOrgIdGitopsBuildTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdGitopsBuildTriggers`: []HandlersGitBuildTriggerRow
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.GetOrgsOrgIdGitopsBuildTriggers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdGitopsBuildTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HandlersGitBuildTriggerRow**](HandlersGitBuildTriggerRow.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdGitopsConnections

> []ModelsGitConnection GetOrgsOrgIdGitopsConnections(ctx, orgId).Execute()

List the org's git connections

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
	resp, r, err := apiClient.GitopsAPI.GetOrgsOrgIdGitopsConnections(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.GetOrgsOrgIdGitopsConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdGitopsConnections`: []ModelsGitConnection
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.GetOrgsOrgIdGitopsConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdGitopsConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelsGitConnection**](ModelsGitConnection.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdGitopsConnectionsIdBranches

> []string GetOrgsOrgIdGitopsConnectionsIdBranches(ctx, orgId, id).Repo(repo).Execute()

List branches of a repository on a git connection

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
	id := int32(56) // int32 | Connection ID
	repo := "repo_example" // string | Repository (owner/name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.GetOrgsOrgIdGitopsConnectionsIdBranches(context.Background(), orgId, id).Repo(repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.GetOrgsOrgIdGitopsConnectionsIdBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdGitopsConnectionsIdBranches`: []string
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.GetOrgsOrgIdGitopsConnectionsIdBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdGitopsConnectionsIdBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **repo** | **string** | Repository (owner/name) | 

### Return type

**[]string**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdGitopsConnectionsIdRepos

> map[string]interface{} GetOrgsOrgIdGitopsConnectionsIdRepos(ctx, orgId, id).Page(page).PerPage(perPage).Execute()

List repositories on a git connection

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
	id := int32(56) // int32 | Connection ID
	page := int32(56) // int32 | Page (optional)
	perPage := int32(56) // int32 | Page size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.GetOrgsOrgIdGitopsConnectionsIdRepos(context.Background(), orgId, id).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.GetOrgsOrgIdGitopsConnectionsIdRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdGitopsConnectionsIdRepos`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.GetOrgsOrgIdGitopsConnectionsIdRepos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdGitopsConnectionsIdReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page | 
 **perPage** | **int32** | Page size | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdGitopsVerifyingMachines

> []HandlersVerifyingMachineRow GetOrgsOrgIdGitopsVerifyingMachines(ctx, orgId).ConnectionId(connectionId).Execute()

List the org's verifying machines (public delivery keys)



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
	connectionId := int32(56) // int32 | The git connection id (optional; scopes the list to its pinned verifier) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.GetOrgsOrgIdGitopsVerifyingMachines(context.Background(), orgId).ConnectionId(connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.GetOrgsOrgIdGitopsVerifyingMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdGitopsVerifyingMachines`: []HandlersVerifyingMachineRow
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.GetOrgsOrgIdGitopsVerifyingMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdGitopsVerifyingMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionId** | **int32** | The git connection id (optional; scopes the list to its pinned verifier) | 

### Return type

[**[]HandlersVerifyingMachineRow**](HandlersVerifyingMachineRow.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdAppsInstallGitopsEnable

> map[string]interface{} PostClustersIdAppsInstallGitopsEnable(ctx, id, install).Request(request).Execute()

Enable CI deploys for an app install

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
	id := int32(56) // int32 | Cluster ID
	install := "install_example" // string | Install name
	request := *openapiclient.NewHandlersEnableInstallRequest() // HandlersEnableInstallRequest | Repo and branch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.PostClustersIdAppsInstallGitopsEnable(context.Background(), id, install).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.PostClustersIdAppsInstallGitopsEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallGitopsEnable`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.PostClustersIdAppsInstallGitopsEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallGitopsEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**HandlersEnableInstallRequest**](HandlersEnableInstallRequest.md) | Repo and branch | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdAppsInstallGitopsRegistration

> map[string]interface{} PostClustersIdAppsInstallGitopsRegistration(ctx, id, install).Body(body).Execute()

Create the deploy setup for a repository



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
	id := int32(56) // int32 | Cluster ID
	install := "install_example" // string | Install name (the confirmed, editable derivation)
	body := *openapiclient.NewHandlersCreateRegistrationRequest() // HandlersCreateRegistrationRequest | The confirmed shape

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.PostClustersIdAppsInstallGitopsRegistration(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.PostClustersIdAppsInstallGitopsRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallGitopsRegistration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.PostClustersIdAppsInstallGitopsRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name (the confirmed, editable derivation) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallGitopsRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersCreateRegistrationRequest**](HandlersCreateRegistrationRequest.md) | The confirmed shape | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdAppsInstallGitopsRegistrationRotate

> map[string]interface{} PostClustersIdAppsInstallGitopsRegistrationRotate(ctx, id, install).Execute()

Rotate the registration-class deploy token



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
	id := int32(56) // int32 | Cluster ID
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.PostClustersIdAppsInstallGitopsRegistrationRotate(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.PostClustersIdAppsInstallGitopsRegistrationRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallGitopsRegistrationRotate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.PostClustersIdAppsInstallGitopsRegistrationRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallGitopsRegistrationRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdGitopsDetect

> GitopsCreateDetection PostClustersIdGitopsDetect(ctx, id).Body(body).Execute()

Inspect a repository for the create-from-repo shape



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
	id := int32(56) // int32 | Cluster ID
	body := *openapiclient.NewHandlersDetectCreateRequest() // HandlersDetectCreateRequest | The picker's selection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.PostClustersIdGitopsDetect(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.PostClustersIdGitopsDetect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdGitopsDetect`: GitopsCreateDetection
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.PostClustersIdGitopsDetect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdGitopsDetectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersDetectCreateRequest**](HandlersDetectCreateRequest.md) | The picker&#39;s selection | 

### Return type

[**GitopsCreateDetection**](GitopsCreateDetection.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdGitopsConnections

> map[string]interface{} PostOrgsOrgIdGitopsConnections(ctx, orgId).Request(request).Execute()

Connect a git provider to the org

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
	request := *openapiclient.NewHandlersCreateGitopsConnectionRequest() // HandlersCreateGitopsConnectionRequest | Provider and credentials (sealed in the vault)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.PostOrgsOrgIdGitopsConnections(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.PostOrgsOrgIdGitopsConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdGitopsConnections`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.PostOrgsOrgIdGitopsConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdGitopsConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersCreateGitopsConnectionRequest**](HandlersCreateGitopsConnectionRequest.md) | Provider and credentials (sealed in the vault) | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdGitopsConnectionsIdBuildTrigger

> HandlersGitBuildTriggerBindResponse PostOrgsOrgIdGitopsConnectionsIdBuildTrigger(ctx, orgId, id).Body(body).Execute()

Bind a repo to a project's build (external-git trigger)



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
	id := int32(56) // int32 | The git connection id
	body := *openapiclient.NewHandlersGitBuildTriggerBindRequest() // HandlersGitBuildTriggerBindRequest | The binding

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.PostOrgsOrgIdGitopsConnectionsIdBuildTrigger(context.Background(), orgId, id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.PostOrgsOrgIdGitopsConnectionsIdBuildTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdGitopsConnectionsIdBuildTrigger`: HandlersGitBuildTriggerBindResponse
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.PostOrgsOrgIdGitopsConnectionsIdBuildTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 
**id** | **int32** | The git connection id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdGitopsConnectionsIdBuildTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersGitBuildTriggerBindRequest**](HandlersGitBuildTriggerBindRequest.md) | The binding | 

### Return type

[**HandlersGitBuildTriggerBindResponse**](HandlersGitBuildTriggerBindResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdGitopsConnectionsIdCredentials

> map[string]interface{} PostOrgsOrgIdGitopsConnectionsIdCredentials(ctx, orgId, id).Request(request).Execute()

Replace a git connection's credentials (reconnect)

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
	id := int32(56) // int32 | Connection ID
	request := *openapiclient.NewHandlersCreateGitopsConnectionRequest() // HandlersCreateGitopsConnectionRequest | Fresh credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.PostOrgsOrgIdGitopsConnectionsIdCredentials(context.Background(), orgId, id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.PostOrgsOrgIdGitopsConnectionsIdCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdGitopsConnectionsIdCredentials`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.PostOrgsOrgIdGitopsConnectionsIdCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdGitopsConnectionsIdCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**HandlersCreateGitopsConnectionRequest**](HandlersCreateGitopsConnectionRequest.md) | Fresh credentials | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdGitopsConnectionsIdTest

> map[string]interface{} PostOrgsOrgIdGitopsConnectionsIdTest(ctx, orgId, id).Execute()

Re-validate a git connection's stored credentials

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
	id := int32(56) // int32 | Connection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitopsAPI.PostOrgsOrgIdGitopsConnectionsIdTest(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitopsAPI.PostOrgsOrgIdGitopsConnectionsIdTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdGitopsConnectionsIdTest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitopsAPI.PostOrgsOrgIdGitopsConnectionsIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdGitopsConnectionsIdTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

