# \GitWebhooksAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostHooksGithub**](GitWebhooksAPI.md#PostHooksGithub) | **Post** /hooks/github | GitHub push webhook (external-git trigger)
[**PostHooksGitlab**](GitWebhooksAPI.md#PostHooksGitlab) | **Post** /hooks/gitlab | GitLab push webhook (external-git trigger)



## PostHooksGithub

> map[string]interface{} PostHooksGithub(ctx).XGitHubDelivery(xGitHubDelivery).XGitHubEvent(xGitHubEvent).XHubSignature256(xHubSignature256).Execute()

GitHub push webhook (external-git trigger)



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
	xGitHubDelivery := "xGitHubDelivery_example" // string | The delivery id
	xGitHubEvent := "xGitHubEvent_example" // string | The event type (push)
	xHubSignature256 := "xHubSignature256_example" // string | The HMAC-SHA256 signature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitWebhooksAPI.PostHooksGithub(context.Background()).XGitHubDelivery(xGitHubDelivery).XGitHubEvent(xGitHubEvent).XHubSignature256(xHubSignature256).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitWebhooksAPI.PostHooksGithub``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostHooksGithub`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitWebhooksAPI.PostHooksGithub`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostHooksGithubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGitHubDelivery** | **string** | The delivery id | 
 **xGitHubEvent** | **string** | The event type (push) | 
 **xHubSignature256** | **string** | The HMAC-SHA256 signature | 

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


## PostHooksGitlab

> map[string]interface{} PostHooksGitlab(ctx).XGitlabEventUUID(xGitlabEventUUID).XGitlabEvent(xGitlabEvent).XGitlabToken(xGitlabToken).Execute()

GitLab push webhook (external-git trigger)



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
	xGitlabEventUUID := "xGitlabEventUUID_example" // string | The delivery id
	xGitlabEvent := "xGitlabEvent_example" // string | The event type (Push Hook)
	xGitlabToken := "xGitlabToken_example" // string | The secret token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitWebhooksAPI.PostHooksGitlab(context.Background()).XGitlabEventUUID(xGitlabEventUUID).XGitlabEvent(xGitlabEvent).XGitlabToken(xGitlabToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitWebhooksAPI.PostHooksGitlab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostHooksGitlab`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitWebhooksAPI.PostHooksGitlab`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostHooksGitlabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGitlabEventUUID** | **string** | The delivery id | 
 **xGitlabEvent** | **string** | The event type (Push Hook) | 
 **xGitlabToken** | **string** | The secret token | 

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

