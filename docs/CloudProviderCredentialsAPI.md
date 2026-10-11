# \CloudProviderCredentialsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCloudProviderCredentialsId**](CloudProviderCredentialsAPI.md#DeleteCloudProviderCredentialsId) | **Delete** /cloud-provider-credentials/{id} | Delete cloud provider credentials
[**GetCloudProviderCredentials**](CloudProviderCredentialsAPI.md#GetCloudProviderCredentials) | **Get** /cloud-provider-credentials | List all cloud provider credentials
[**PostCloudProviderCredentials**](CloudProviderCredentialsAPI.md#PostCloudProviderCredentials) | **Post** /cloud-provider-credentials | Create cloud provider credentials
[**PostCloudProviderCredentialsIdValidate**](CloudProviderCredentialsAPI.md#PostCloudProviderCredentialsIdValidate) | **Post** /cloud-provider-credentials/{id}/validate | Validate cloud provider credentials
[**PostCloudProviderCredentialsValidate**](CloudProviderCredentialsAPI.md#PostCloudProviderCredentialsValidate) | **Post** /cloud-provider-credentials/validate | Validate credentials without saving
[**PutCloudProviderCredentialsId**](CloudProviderCredentialsAPI.md#PutCloudProviderCredentialsId) | **Put** /cloud-provider-credentials/{id} | Update cloud provider credentials



## DeleteCloudProviderCredentialsId

> DeleteCloudProviderCredentialsId(ctx, id).Execute()

Delete cloud provider credentials



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
	id := "id_example" // string | Credential ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudProviderCredentialsAPI.DeleteCloudProviderCredentialsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.DeleteCloudProviderCredentialsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credential ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudProviderCredentialsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetCloudProviderCredentials

> []map[string]interface{} GetCloudProviderCredentials(ctx).Execute()

List all cloud provider credentials



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
	resp, r, err := apiClient.CloudProviderCredentialsAPI.GetCloudProviderCredentials(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.GetCloudProviderCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProviderCredentials`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.GetCloudProviderCredentials`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProviderCredentialsRequest struct via the builder pattern


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudProviderCredentials

> map[string]interface{} PostCloudProviderCredentials(ctx).Request(request).Execute()

Create cloud provider credentials



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
	request := *openapiclient.NewHandlersCreateCloudProviderCredentialRequest(map[string]interface{}{"key": interface{}(123)}, "Label_example", "Provider_example") // HandlersCreateCloudProviderCredentialRequest | Credential details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.PostCloudProviderCredentials(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.PostCloudProviderCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudProviderCredentials`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.PostCloudProviderCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudProviderCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersCreateCloudProviderCredentialRequest**](HandlersCreateCloudProviderCredentialRequest.md) | Credential details | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [EncryptionKeyAuth](../README.md#EncryptionKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudProviderCredentialsIdValidate

> map[string]interface{} PostCloudProviderCredentialsIdValidate(ctx, id).Execute()

Validate cloud provider credentials



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
	id := "id_example" // string | Credential ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.PostCloudProviderCredentialsIdValidate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.PostCloudProviderCredentialsIdValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudProviderCredentialsIdValidate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.PostCloudProviderCredentialsIdValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credential ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudProviderCredentialsIdValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [EncryptionKeyAuth](../README.md#EncryptionKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudProviderCredentialsValidate

> map[string]interface{} PostCloudProviderCredentialsValidate(ctx).Request(request).Execute()

Validate credentials without saving



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
	request := *openapiclient.NewHandlersValidateCredentialsRequest(map[string]interface{}{"key": interface{}(123)}, "Provider_example") // HandlersValidateCredentialsRequest | Credentials to validate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.PostCloudProviderCredentialsValidate(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.PostCloudProviderCredentialsValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudProviderCredentialsValidate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.PostCloudProviderCredentialsValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudProviderCredentialsValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersValidateCredentialsRequest**](HandlersValidateCredentialsRequest.md) | Credentials to validate | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [EncryptionKeyAuth](../README.md#EncryptionKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCloudProviderCredentialsId

> map[string]interface{} PutCloudProviderCredentialsId(ctx, id).Request(request).Execute()

Update cloud provider credentials



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
	id := "id_example" // string | Credential ID
	request := *openapiclient.NewHandlersUpdateCloudProviderCredentialRequest(map[string]interface{}{"key": interface{}(123)}, "Label_example") // HandlersUpdateCloudProviderCredentialRequest | Updated credential details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProviderCredentialsAPI.PutCloudProviderCredentialsId(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProviderCredentialsAPI.PutCloudProviderCredentialsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCloudProviderCredentialsId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProviderCredentialsAPI.PutCloudProviderCredentialsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credential ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCloudProviderCredentialsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersUpdateCloudProviderCredentialRequest**](HandlersUpdateCloudProviderCredentialRequest.md) | Updated credential details | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [EncryptionKeyAuth](../README.md#EncryptionKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

