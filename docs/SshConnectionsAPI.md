# \SshConnectionsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSshConnectionsId**](SshConnectionsAPI.md#DeleteSshConnectionsId) | **Delete** /ssh-connections/{id} | Delete an SSH connection
[**GetSshConnections**](SshConnectionsAPI.md#GetSshConnections) | **Get** /ssh-connections | List all SSH connections
[**PostSshConnections**](SshConnectionsAPI.md#PostSshConnections) | **Post** /ssh-connections | Create a new SSH connection (disabled)



## DeleteSshConnectionsId

> DeleteSshConnectionsId(ctx, id).Execute()

Delete an SSH connection



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
	id := "id_example" // string | Connection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SshConnectionsAPI.DeleteSshConnectionsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshConnectionsAPI.DeleteSshConnectionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSshConnectionsIdRequest struct via the builder pattern


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


## GetSshConnections

> []HandlersSSHConnectionResponse GetSshConnections(ctx).BastionOnly(bastionOnly).Execute()

List all SSH connections



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
	bastionOnly := true // bool | Filter bastion hosts only (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SshConnectionsAPI.GetSshConnections(context.Background()).BastionOnly(bastionOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshConnectionsAPI.GetSshConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshConnections`: []HandlersSSHConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `SshConnectionsAPI.GetSshConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSshConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bastionOnly** | **bool** | Filter bastion hosts only | 

### Return type

[**[]HandlersSSHConnectionResponse**](HandlersSSHConnectionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshConnections

> ModelsManagedSSHConnection PostSshConnections(ctx).Request(request).Execute()

Create a new SSH connection (disabled)



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
	request := *openapiclient.NewHandlersCreateSSHConnectionRequest("IpAddress_example", "Name_example", "PrivateSshKey_example") // HandlersCreateSSHConnectionRequest | SSH connection details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SshConnectionsAPI.PostSshConnections(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshConnectionsAPI.PostSshConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSshConnections`: ModelsManagedSSHConnection
	fmt.Fprintf(os.Stdout, "Response from `SshConnectionsAPI.PostSshConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSshConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersCreateSSHConnectionRequest**](HandlersCreateSSHConnectionRequest.md) | SSH connection details | 

### Return type

[**ModelsManagedSSHConnection**](ModelsManagedSSHConnection.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

