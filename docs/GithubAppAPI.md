# \GithubAppAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrganizationsOrgIdGithubAppInstallationsId**](GithubAppAPI.md#DeleteOrganizationsOrgIdGithubAppInstallationsId) | **Delete** /organizations/{org_id}/github-app/installations/{id} | Disconnect one of the org&#39;s GitHub App installations
[**GetGithubAppInstallCallback**](GithubAppAPI.md#GetGithubAppInstallCallback) | **Get** /github-app/install/callback | GitHub App install callback
[**GetOrganizationsOrgIdGithubAppInstall**](GithubAppAPI.md#GetOrganizationsOrgIdGithubAppInstall) | **Get** /organizations/{org_id}/github-app/install | Start the org&#39;s GitHub App installation flow
[**GetOrganizationsOrgIdGithubAppInstallations**](GithubAppAPI.md#GetOrganizationsOrgIdGithubAppInstallations) | **Get** /organizations/{org_id}/github-app/installations | List the org&#39;s GitHub App installations
[**PostHooksGithubApp**](GithubAppAPI.md#PostHooksGithubApp) | **Post** /hooks/github-app | GitHub App webhook (lifecycle + workflow_job)



## DeleteOrganizationsOrgIdGithubAppInstallationsId

> string DeleteOrganizationsOrgIdGithubAppInstallationsId(ctx, orgId, id).Execute()

Disconnect one of the org's GitHub App installations



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
	id := int32(56) // int32 | Installation row id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GithubAppAPI.DeleteOrganizationsOrgIdGithubAppInstallationsId(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GithubAppAPI.DeleteOrganizationsOrgIdGithubAppInstallationsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganizationsOrgIdGithubAppInstallationsId`: string
	fmt.Fprintf(os.Stdout, "Response from `GithubAppAPI.DeleteOrganizationsOrgIdGithubAppInstallationsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Installation row id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationsOrgIdGithubAppInstallationsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubAppInstallCallback

> GetGithubAppInstallCallback(ctx).InstallationId(installationId).SetupAction(setupAction).State(state).Code(code).Error_(error_).Execute()

GitHub App install callback



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
	installationId := int32(56) // int32 | GitHub's installation id (present on both redirect legs)
	setupAction := "setupAction_example" // string | install | update | request
	state := "state_example" // string | The signed state token from the start endpoint (both hops)
	code := "code_example" // string | GitHub's OAuth code (hop 2 only) (optional)
	error_ := "error__example" // string | GitHub's OAuth error code (the user clicked Cancel on the authorize page) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GithubAppAPI.GetGithubAppInstallCallback(context.Background()).InstallationId(installationId).SetupAction(setupAction).State(state).Code(code).Error_(error_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GithubAppAPI.GetGithubAppInstallCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubAppInstallCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **installationId** | **int32** | GitHub&#39;s installation id (present on both redirect legs) | 
 **setupAction** | **string** | install | update | request | 
 **state** | **string** | The signed state token from the start endpoint (both hops) | 
 **code** | **string** | GitHub&#39;s OAuth code (hop 2 only) | 
 **error_** | **string** | GitHub&#39;s OAuth error code (the user clicked Cancel on the authorize page) | 

### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsOrgIdGithubAppInstall

> map[string]interface{} GetOrganizationsOrgIdGithubAppInstall(ctx, orgId).Execute()

Start the org's GitHub App installation flow



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
	resp, r, err := apiClient.GithubAppAPI.GetOrganizationsOrgIdGithubAppInstall(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GithubAppAPI.GetOrganizationsOrgIdGithubAppInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsOrgIdGithubAppInstall`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GithubAppAPI.GetOrganizationsOrgIdGithubAppInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsOrgIdGithubAppInstallRequest struct via the builder pattern


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


## GetOrganizationsOrgIdGithubAppInstallations

> map[string]interface{} GetOrganizationsOrgIdGithubAppInstallations(ctx, orgId).Execute()

List the org's GitHub App installations



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
	resp, r, err := apiClient.GithubAppAPI.GetOrganizationsOrgIdGithubAppInstallations(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GithubAppAPI.GetOrganizationsOrgIdGithubAppInstallations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsOrgIdGithubAppInstallations`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GithubAppAPI.GetOrganizationsOrgIdGithubAppInstallations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsOrgIdGithubAppInstallationsRequest struct via the builder pattern


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


## PostHooksGithubApp

> map[string]interface{} PostHooksGithubApp(ctx).XGitHubEvent(xGitHubEvent).XHubSignature256(xHubSignature256).XGitHubDelivery(xGitHubDelivery).Execute()

GitHub App webhook (lifecycle + workflow_job)



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
	xGitHubEvent := "xGitHubEvent_example" // string | installation | installation_target | workflow_job | ping
	xHubSignature256 := "xHubSignature256_example" // string | The HMAC-SHA256 signature
	xGitHubDelivery := "xGitHubDelivery_example" // string | The delivery id (the workflow_job dedupe key) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GithubAppAPI.PostHooksGithubApp(context.Background()).XGitHubEvent(xGitHubEvent).XHubSignature256(xHubSignature256).XGitHubDelivery(xGitHubDelivery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GithubAppAPI.PostHooksGithubApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostHooksGithubApp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GithubAppAPI.PostHooksGithubApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostHooksGithubAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGitHubEvent** | **string** | installation | installation_target | workflow_job | ping | 
 **xHubSignature256** | **string** | The HMAC-SHA256 signature | 
 **xGitHubDelivery** | **string** | The delivery id (the workflow_job dedupe key) | 

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

