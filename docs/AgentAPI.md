# \AgentAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentInstallSh**](AgentAPI.md#GetAgentInstallSh) | **Get** /agent/install.sh | Get agent installation script
[**GetAgentManifests**](AgentAPI.md#GetAgentManifests) | **Get** /agent/manifests | Get agent installation manifest



## GetAgentInstallSh

> string GetAgentInstallSh(ctx).Token(token).Execute()

Get agent installation script



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
	token := "token_example" // string | Agent registration token (short token)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentInstallSh(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentInstallSh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentInstallSh`: string
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentInstallSh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInstallShRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Agent registration token (short token) | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentManifests

> string GetAgentManifests(ctx).Token(token).Name(name).Namespace(namespace).ImageTag(imageTag).Execute()

Get agent installation manifest



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
	token := "token_example" // string | Agent registration token (short token)
	name := "name_example" // string | Cluster name (optional)
	namespace := "namespace_example" // string | Namespace (default: bare-metal-agent) (optional)
	imageTag := "imageTag_example" // string | Agent image tag (default: latest) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.GetAgentManifests(context.Background()).Token(token).Name(name).Namespace(namespace).ImageTag(imageTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.GetAgentManifests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentManifests`: string
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.GetAgentManifests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentManifestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Agent registration token (short token) | 
 **name** | **string** | Cluster name | 
 **namespace** | **string** | Namespace (default: bare-metal-agent) | 
 **imageTag** | **string** | Agent image tag (default: latest) | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

