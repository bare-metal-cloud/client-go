# \EdgeGitAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGitEdgeOrgIdProjectIdRepoRefPacksPushRef**](EdgeGitAPI.md#GetGitEdgeOrgIdProjectIdRepoRefPacksPushRef) | **Get** /git/edge/{org_id}/{project_id}/{repo_ref}/packs/{push_ref} | Fetch a stored pack (fetch-token redemption)
[**GetGitEdgeOrgIdProjectIdSignals**](EdgeGitAPI.md#GetGitEdgeOrgIdProjectIdSignals) | **Get** /git/edge/{org_id}/{project_id}/signals | Poll the next fetch signal (the addon&#39;s mirror sidecar)
[**PostGitEdgeOrgIdProjectIdRepoRefGitReceivePack**](EdgeGitAPI.md#PostGitEdgeOrgIdProjectIdRepoRefGitReceivePack) | **Post** /git/edge/{org_id}/{project_id}/{repo_ref}/git-receive-pack | Edge git receive-pack (BMC-hosted git push)



## GetGitEdgeOrgIdProjectIdRepoRefPacksPushRef

> *os.File GetGitEdgeOrgIdProjectIdRepoRefPacksPushRef(ctx, orgId, projectId, repoRef, pushRef).Authorization(authorization).Execute()

Fetch a stored pack (fetch-token redemption)



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
	projectId := int32(56) // int32 | The project id
	repoRef := "repoRef_example" // string | The repository reference
	pushRef := "pushRef_example" // string | The push reference
	authorization := "authorization_example" // string | The fetch token (Bearer)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeGitAPI.GetGitEdgeOrgIdProjectIdRepoRefPacksPushRef(context.Background(), orgId, projectId, repoRef, pushRef).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeGitAPI.GetGitEdgeOrgIdProjectIdRepoRefPacksPushRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitEdgeOrgIdProjectIdRepoRefPacksPushRef`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `EdgeGitAPI.GetGitEdgeOrgIdProjectIdRepoRefPacksPushRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 
**projectId** | **int32** | The project id | 
**repoRef** | **string** | The repository reference | 
**pushRef** | **string** | The push reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitEdgeOrgIdProjectIdRepoRefPacksPushRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | The fetch token (Bearer) | 

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


## GetGitEdgeOrgIdProjectIdSignals

> EdgegitFetchSignalPayload GetGitEdgeOrgIdProjectIdSignals(ctx, orgId, projectId).Execute()

Poll the next fetch signal (the addon's mirror sidecar)



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
	projectId := int32(56) // int32 | The project id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeGitAPI.GetGitEdgeOrgIdProjectIdSignals(context.Background(), orgId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeGitAPI.GetGitEdgeOrgIdProjectIdSignals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitEdgeOrgIdProjectIdSignals`: EdgegitFetchSignalPayload
	fmt.Fprintf(os.Stdout, "Response from `EdgeGitAPI.GetGitEdgeOrgIdProjectIdSignals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 
**projectId** | **int32** | The project id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitEdgeOrgIdProjectIdSignalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EdgegitFetchSignalPayload**](EdgegitFetchSignalPayload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitEdgeOrgIdProjectIdRepoRefGitReceivePack

> map[string]interface{} PostGitEdgeOrgIdProjectIdRepoRefGitReceivePack(ctx, orgId, projectId, repoRef).XBMCBranch(xBMCBranch).XBMCCommit(xBMCCommit).XBMCPusher(xBMCPusher).Execute()

Edge git receive-pack (BMC-hosted git push)



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
	projectId := int32(56) // int32 | The project id
	repoRef := "repoRef_example" // string | The repository reference
	xBMCBranch := "xBMCBranch_example" // string | The pushed branch (the v1 metadata grammar; clamped to 255) (optional)
	xBMCCommit := "xBMCCommit_example" // string | The pushed commit SHA (the v1 metadata grammar; clamped to 64) (optional)
	xBMCPusher := "xBMCPusher_example" // string | The pushing identity (the v1 metadata grammar; clamped to 64) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeGitAPI.PostGitEdgeOrgIdProjectIdRepoRefGitReceivePack(context.Background(), orgId, projectId, repoRef).XBMCBranch(xBMCBranch).XBMCCommit(xBMCCommit).XBMCPusher(xBMCPusher).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeGitAPI.PostGitEdgeOrgIdProjectIdRepoRefGitReceivePack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitEdgeOrgIdProjectIdRepoRefGitReceivePack`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EdgeGitAPI.PostGitEdgeOrgIdProjectIdRepoRefGitReceivePack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | The org id | 
**projectId** | **int32** | The project id | 
**repoRef** | **string** | The repository reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitEdgeOrgIdProjectIdRepoRefGitReceivePackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xBMCBranch** | **string** | The pushed branch (the v1 metadata grammar; clamped to 255) | 
 **xBMCCommit** | **string** | The pushed commit SHA (the v1 metadata grammar; clamped to 64) | 
 **xBMCPusher** | **string** | The pushing identity (the v1 metadata grammar; clamped to 64) | 

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

