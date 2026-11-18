# \RunnersArtifactsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgsOrgIdArtifactsArtifactId**](RunnersArtifactsAPI.md#DeleteOrgsOrgIdArtifactsArtifactId) | **Delete** /orgs/{org_id}/artifacts/{artifact_id} | Delete a workflow artifact (admin)
[**GetOrgsOrgIdArtifacts**](RunnersArtifactsAPI.md#GetOrgsOrgIdArtifacts) | **Get** /orgs/{org_id}/artifacts | List the org&#39;s workflow artifacts
[**GetOrgsOrgIdArtifactsArtifactIdDownload**](RunnersArtifactsAPI.md#GetOrgsOrgIdArtifactsArtifactIdDownload) | **Get** /orgs/{org_id}/artifacts/{artifact_id}/download | Download a workflow artifact&#39;s bytes
[**GetOrgsOrgIdArtifactsSettings**](RunnersArtifactsAPI.md#GetOrgsOrgIdArtifactsSettings) | **Get** /orgs/{org_id}/artifacts/settings | Get the org&#39;s artifact retention settings
[**GetOrgsOrgIdArtifactsStorage**](RunnersArtifactsAPI.md#GetOrgsOrgIdArtifactsStorage) | **Get** /orgs/{org_id}/artifacts/storage | Get the org&#39;s artifact storage view
[**PutOrgsOrgIdArtifactsArtifactIdRetention**](RunnersArtifactsAPI.md#PutOrgsOrgIdArtifactsArtifactIdRetention) | **Put** /orgs/{org_id}/artifacts/{artifact_id}/retention | Set an artifact&#39;s retention override (null clears)
[**PutOrgsOrgIdArtifactsSettings**](RunnersArtifactsAPI.md#PutOrgsOrgIdArtifactsSettings) | **Put** /orgs/{org_id}/artifacts/settings | Set the org&#39;s artifact retention settings (admin)



## DeleteOrgsOrgIdArtifactsArtifactId

> map[string]interface{} DeleteOrgsOrgIdArtifactsArtifactId(ctx, orgId, artifactId).Execute()

Delete a workflow artifact (admin)



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
	artifactId := int32(56) // int32 | The artifact row id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersArtifactsAPI.DeleteOrgsOrgIdArtifactsArtifactId(context.Background(), orgId, artifactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersArtifactsAPI.DeleteOrgsOrgIdArtifactsArtifactId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdArtifactsArtifactId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RunnersArtifactsAPI.DeleteOrgsOrgIdArtifactsArtifactId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**artifactId** | **int32** | The artifact row id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdArtifactsArtifactIdRequest struct via the builder pattern


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


## GetOrgsOrgIdArtifacts

> map[string]interface{} GetOrgsOrgIdArtifacts(ctx, orgId).Repository(repository).Run(run).Limit(limit).Execute()

List the org's workflow artifacts



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
	repository := "repository_example" // string | Filter by repository bare name (exact) (optional)
	run := int32(56) // int32 | Filter by workflow run id (optional)
	limit := int32(56) // int32 | Page size (default 50, at most 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersArtifactsAPI.GetOrgsOrgIdArtifacts(context.Background(), orgId).Repository(repository).Run(run).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersArtifactsAPI.GetOrgsOrgIdArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdArtifacts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RunnersArtifactsAPI.GetOrgsOrgIdArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | **string** | Filter by repository bare name (exact) | 
 **run** | **int32** | Filter by workflow run id | 
 **limit** | **int32** | Page size (default 50, at most 200) | 

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


## GetOrgsOrgIdArtifactsArtifactIdDownload

> *os.File GetOrgsOrgIdArtifactsArtifactIdDownload(ctx, orgId, artifactId).Execute()

Download a workflow artifact's bytes



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
	artifactId := int32(56) // int32 | The artifact row id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersArtifactsAPI.GetOrgsOrgIdArtifactsArtifactIdDownload(context.Background(), orgId, artifactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersArtifactsAPI.GetOrgsOrgIdArtifactsArtifactIdDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdArtifactsArtifactIdDownload`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RunnersArtifactsAPI.GetOrgsOrgIdArtifactsArtifactIdDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**artifactId** | **int32** | The artifact row id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdArtifactsArtifactIdDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetOrgsOrgIdArtifactsSettings

> HandlersRunnerArtifactSettingsRow GetOrgsOrgIdArtifactsSettings(ctx, orgId).Execute()

Get the org's artifact retention settings



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
	resp, r, err := apiClient.RunnersArtifactsAPI.GetOrgsOrgIdArtifactsSettings(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersArtifactsAPI.GetOrgsOrgIdArtifactsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdArtifactsSettings`: HandlersRunnerArtifactSettingsRow
	fmt.Fprintf(os.Stdout, "Response from `RunnersArtifactsAPI.GetOrgsOrgIdArtifactsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdArtifactsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerArtifactSettingsRow**](HandlersRunnerArtifactSettingsRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdArtifactsStorage

> HandlersRunnerArtifactStorageRow GetOrgsOrgIdArtifactsStorage(ctx, orgId).Execute()

Get the org's artifact storage view



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
	resp, r, err := apiClient.RunnersArtifactsAPI.GetOrgsOrgIdArtifactsStorage(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersArtifactsAPI.GetOrgsOrgIdArtifactsStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdArtifactsStorage`: HandlersRunnerArtifactStorageRow
	fmt.Fprintf(os.Stdout, "Response from `RunnersArtifactsAPI.GetOrgsOrgIdArtifactsStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdArtifactsStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerArtifactStorageRow**](HandlersRunnerArtifactStorageRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdArtifactsArtifactIdRetention

> HandlersRunnerArtifactUserRow PutOrgsOrgIdArtifactsArtifactIdRetention(ctx, orgId, artifactId).Body(body).Execute()

Set an artifact's retention override (null clears)



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
	artifactId := int32(56) // int32 | The artifact row id
	body := *openapiclient.NewHandlersRunnerArtifactRetentionBody() // HandlersRunnerArtifactRetentionBody | days is 1 to 365, or null to clear the override

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersArtifactsAPI.PutOrgsOrgIdArtifactsArtifactIdRetention(context.Background(), orgId, artifactId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersArtifactsAPI.PutOrgsOrgIdArtifactsArtifactIdRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdArtifactsArtifactIdRetention`: HandlersRunnerArtifactUserRow
	fmt.Fprintf(os.Stdout, "Response from `RunnersArtifactsAPI.PutOrgsOrgIdArtifactsArtifactIdRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**artifactId** | **int32** | The artifact row id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdArtifactsArtifactIdRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersRunnerArtifactRetentionBody**](HandlersRunnerArtifactRetentionBody.md) | days is 1 to 365, or null to clear the override | 

### Return type

[**HandlersRunnerArtifactUserRow**](HandlersRunnerArtifactUserRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdArtifactsSettings

> HandlersRunnerArtifactSettingsRow PutOrgsOrgIdArtifactsSettings(ctx, orgId).Body(body).Execute()

Set the org's artifact retention settings (admin)



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
	body := *openapiclient.NewHandlersRunnerArtifactSettingsBody() // HandlersRunnerArtifactSettingsBody | deleteWhenRunEnds, or days 1-90

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersArtifactsAPI.PutOrgsOrgIdArtifactsSettings(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersArtifactsAPI.PutOrgsOrgIdArtifactsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdArtifactsSettings`: HandlersRunnerArtifactSettingsRow
	fmt.Fprintf(os.Stdout, "Response from `RunnersArtifactsAPI.PutOrgsOrgIdArtifactsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdArtifactsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRunnerArtifactSettingsBody**](HandlersRunnerArtifactSettingsBody.md) | deleteWhenRunEnds, or days 1-90 | 

### Return type

[**HandlersRunnerArtifactSettingsRow**](HandlersRunnerArtifactSettingsRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

