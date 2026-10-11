# \SshKeysAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSshKeysId**](SshKeysAPI.md#DeleteSshKeysId) | **Delete** /ssh-keys/{id} | Delete an SSH key
[**GetSshKeys**](SshKeysAPI.md#GetSshKeys) | **Get** /ssh-keys | List all SSH keys
[**GetSshKeysId**](SshKeysAPI.md#GetSshKeysId) | **Get** /ssh-keys/{id} | Get an SSH key by ID
[**PostSshKeys**](SshKeysAPI.md#PostSshKeys) | **Post** /ssh-keys | Create a new SSH key



## DeleteSshKeysId

> DeleteSshKeysId(ctx, id).Execute()

Delete an SSH key



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
	id := "id_example" // string | SSH Key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SshKeysAPI.DeleteSshKeysId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshKeysAPI.DeleteSshKeysId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SSH Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSshKeysIdRequest struct via the builder pattern


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


## GetSshKeys

> []HandlersSshKeyListItem GetSshKeys(ctx).Execute()

List all SSH keys



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
	resp, r, err := apiClient.SshKeysAPI.GetSshKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshKeysAPI.GetSshKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshKeys`: []HandlersSshKeyListItem
	fmt.Fprintf(os.Stdout, "Response from `SshKeysAPI.GetSshKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshKeysRequest struct via the builder pattern


### Return type

[**[]HandlersSshKeyListItem**](HandlersSshKeyListItem.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshKeysId

> ModelsSSHKey GetSshKeysId(ctx, id).Execute()

Get an SSH key by ID



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
	id := "id_example" // string | SSH Key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SshKeysAPI.GetSshKeysId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshKeysAPI.GetSshKeysId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshKeysId`: ModelsSSHKey
	fmt.Fprintf(os.Stdout, "Response from `SshKeysAPI.GetSshKeysId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SSH Key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshKeysIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsSSHKey**](ModelsSSHKey.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [EncryptionKeyAuth](../README.md#EncryptionKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSshKeys

> ModelsSSHKey PostSshKeys(ctx).Request(request).Execute()

Create a new SSH key



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
	request := *openapiclient.NewHandlersCreateSSHKeyRequest("Key_example", "Title_example") // HandlersCreateSSHKeyRequest | SSH key details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SshKeysAPI.PostSshKeys(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SshKeysAPI.PostSshKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSshKeys`: ModelsSSHKey
	fmt.Fprintf(os.Stdout, "Response from `SshKeysAPI.PostSshKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSshKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersCreateSSHKeyRequest**](HandlersCreateSSHKeyRequest.md) | SSH key details | 

### Return type

[**ModelsSSHKey**](ModelsSSHKey.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [EncryptionKeyAuth](../README.md#EncryptionKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

