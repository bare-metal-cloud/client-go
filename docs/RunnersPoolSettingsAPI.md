# \RunnersPoolSettingsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersPoolSettings**](RunnersPoolSettingsAPI.md#GetOrgsOrgIdRunnersPoolSettings) | **Get** /orgs/{org_id}/runners/pool-settings | Get the org&#39;s pool settings
[**PutOrgsOrgIdRunnersPoolSettings**](RunnersPoolSettingsAPI.md#PutOrgsOrgIdRunnersPoolSettings) | **Put** /orgs/{org_id}/runners/pool-settings | Set the org&#39;s pool settings (admin)



## GetOrgsOrgIdRunnersPoolSettings

> HandlersRunnerPoolSettingsBody GetOrgsOrgIdRunnersPoolSettings(ctx, orgId).Execute()

Get the org's pool settings



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
	resp, r, err := apiClient.RunnersPoolSettingsAPI.GetOrgsOrgIdRunnersPoolSettings(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPoolSettingsAPI.GetOrgsOrgIdRunnersPoolSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersPoolSettings`: HandlersRunnerPoolSettingsBody
	fmt.Fprintf(os.Stdout, "Response from `RunnersPoolSettingsAPI.GetOrgsOrgIdRunnersPoolSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersPoolSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerPoolSettingsBody**](HandlersRunnerPoolSettingsBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdRunnersPoolSettings

> HandlersRunnerPoolSettingsBody PutOrgsOrgIdRunnersPoolSettings(ctx, orgId).Body(body).Execute()

Set the org's pool settings (admin)



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
	body := *openapiclient.NewHandlersRunnerPoolSettingsBody() // HandlersRunnerPoolSettingsBody | the settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersPoolSettingsAPI.PutOrgsOrgIdRunnersPoolSettings(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPoolSettingsAPI.PutOrgsOrgIdRunnersPoolSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdRunnersPoolSettings`: HandlersRunnerPoolSettingsBody
	fmt.Fprintf(os.Stdout, "Response from `RunnersPoolSettingsAPI.PutOrgsOrgIdRunnersPoolSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdRunnersPoolSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRunnerPoolSettingsBody**](HandlersRunnerPoolSettingsBody.md) | the settings | 

### Return type

[**HandlersRunnerPoolSettingsBody**](HandlersRunnerPoolSettingsBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

