# \MultipassHostsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMultipassHostsId**](MultipassHostsAPI.md#DeleteMultipassHostsId) | **Delete** /multipass-hosts/{id} | Delete a Multipass host
[**GetMultipassHosts**](MultipassHostsAPI.md#GetMultipassHosts) | **Get** /multipass-hosts | List all Multipass hosts
[**GetMultipassHostsId**](MultipassHostsAPI.md#GetMultipassHostsId) | **Get** /multipass-hosts/{id} | Get a Multipass host by ID
[**PostMultipassHosts**](MultipassHostsAPI.md#PostMultipassHosts) | **Post** /multipass-hosts | Create a new Multipass host



## DeleteMultipassHostsId

> DeleteMultipassHostsId(ctx, id).Execute()

Delete a Multipass host



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
	id := "id_example" // string | Multipass Host ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultipassHostsAPI.DeleteMultipassHostsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassHostsAPI.DeleteMultipassHostsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Multipass Host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMultipassHostsIdRequest struct via the builder pattern


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


## GetMultipassHosts

> []HandlersMultipassHostResponse GetMultipassHosts(ctx).Execute()

List all Multipass hosts



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
	resp, r, err := apiClient.MultipassHostsAPI.GetMultipassHosts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassHostsAPI.GetMultipassHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipassHosts`: []HandlersMultipassHostResponse
	fmt.Fprintf(os.Stdout, "Response from `MultipassHostsAPI.GetMultipassHosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipassHostsRequest struct via the builder pattern


### Return type

[**[]HandlersMultipassHostResponse**](HandlersMultipassHostResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipassHostsId

> ModelsMultipassHost GetMultipassHostsId(ctx, id).Execute()

Get a Multipass host by ID



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
	id := "id_example" // string | Multipass Host ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipassHostsAPI.GetMultipassHostsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassHostsAPI.GetMultipassHostsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipassHostsId`: ModelsMultipassHost
	fmt.Fprintf(os.Stdout, "Response from `MultipassHostsAPI.GetMultipassHostsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Multipass Host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipassHostsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsMultipassHost**](ModelsMultipassHost.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMultipassHosts

> ModelsMultipassHost PostMultipassHosts(ctx).Request(request).Execute()

Create a new Multipass host



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
	request := *openapiclient.NewHandlersCreateMultipassHostRequest("Name_example") // HandlersCreateMultipassHostRequest | Multipass host creation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipassHostsAPI.PostMultipassHosts(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipassHostsAPI.PostMultipassHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMultipassHosts`: ModelsMultipassHost
	fmt.Fprintf(os.Stdout, "Response from `MultipassHostsAPI.PostMultipassHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMultipassHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersCreateMultipassHostRequest**](HandlersCreateMultipassHostRequest.md) | Multipass host creation request | 

### Return type

[**ModelsMultipassHost**](ModelsMultipassHost.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

