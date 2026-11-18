# \RunnersPackingHostsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgsOrgIdRunnersPackingHostsHostId**](RunnersPackingHostsAPI.md#DeleteOrgsOrgIdRunnersPackingHostsHostId) | **Delete** /orgs/{org_id}/runners/packing-hosts/{host_id} | Withdraw a packing host declaration (admin)
[**GetOrgsOrgIdRunnersPackingHosts**](RunnersPackingHostsAPI.md#GetOrgsOrgIdRunnersPackingHosts) | **Get** /orgs/{org_id}/runners/packing-hosts | List the org&#39;s packing host declarations
[**GetOrgsOrgIdRunnersPackingHostsHostId**](RunnersPackingHostsAPI.md#GetOrgsOrgIdRunnersPackingHostsHostId) | **Get** /orgs/{org_id}/runners/packing-hosts/{host_id} | Get one packing host declaration
[**PostOrgsOrgIdRunnersPackingHosts**](RunnersPackingHostsAPI.md#PostOrgsOrgIdRunnersPackingHosts) | **Post** /orgs/{org_id}/runners/packing-hosts | Declare a packing host (admin)
[**PutOrgsOrgIdRunnersPackingHostsHostId**](RunnersPackingHostsAPI.md#PutOrgsOrgIdRunnersPackingHostsHostId) | **Put** /orgs/{org_id}/runners/packing-hosts/{host_id} | Rewrite a packing host declaration (admin)



## DeleteOrgsOrgIdRunnersPackingHostsHostId

> map[string]bool DeleteOrgsOrgIdRunnersPackingHostsHostId(ctx, orgId, hostId).Execute()

Withdraw a packing host declaration (admin)



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
	hostId := "hostId_example" // string | Packing host id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersPackingHostsAPI.DeleteOrgsOrgIdRunnersPackingHostsHostId(context.Background(), orgId, hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPackingHostsAPI.DeleteOrgsOrgIdRunnersPackingHostsHostId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdRunnersPackingHostsHostId`: map[string]bool
	fmt.Fprintf(os.Stdout, "Response from `RunnersPackingHostsAPI.DeleteOrgsOrgIdRunnersPackingHostsHostId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**hostId** | **string** | Packing host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdRunnersPackingHostsHostIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersPackingHosts

> map[string]interface{} GetOrgsOrgIdRunnersPackingHosts(ctx, orgId).Execute()

List the org's packing host declarations



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
	resp, r, err := apiClient.RunnersPackingHostsAPI.GetOrgsOrgIdRunnersPackingHosts(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPackingHostsAPI.GetOrgsOrgIdRunnersPackingHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersPackingHosts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RunnersPackingHostsAPI.GetOrgsOrgIdRunnersPackingHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersPackingHostsRequest struct via the builder pattern


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


## GetOrgsOrgIdRunnersPackingHostsHostId

> ModelsRunnerPackingHost GetOrgsOrgIdRunnersPackingHostsHostId(ctx, orgId, hostId).Execute()

Get one packing host declaration



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
	hostId := "hostId_example" // string | Packing host id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersPackingHostsAPI.GetOrgsOrgIdRunnersPackingHostsHostId(context.Background(), orgId, hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPackingHostsAPI.GetOrgsOrgIdRunnersPackingHostsHostId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersPackingHostsHostId`: ModelsRunnerPackingHost
	fmt.Fprintf(os.Stdout, "Response from `RunnersPackingHostsAPI.GetOrgsOrgIdRunnersPackingHostsHostId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**hostId** | **string** | Packing host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersPackingHostsHostIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsRunnerPackingHost**](ModelsRunnerPackingHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersPackingHosts

> ModelsRunnerPackingHost PostOrgsOrgIdRunnersPackingHosts(ctx, orgId).Body(body).Execute()

Declare a packing host (admin)



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
	body := *openapiclient.NewHandlersRunnerPackingHostBody() // HandlersRunnerPackingHostBody | the declaration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersPackingHostsAPI.PostOrgsOrgIdRunnersPackingHosts(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPackingHostsAPI.PostOrgsOrgIdRunnersPackingHosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersPackingHosts`: ModelsRunnerPackingHost
	fmt.Fprintf(os.Stdout, "Response from `RunnersPackingHostsAPI.PostOrgsOrgIdRunnersPackingHosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersPackingHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRunnerPackingHostBody**](HandlersRunnerPackingHostBody.md) | the declaration | 

### Return type

[**ModelsRunnerPackingHost**](ModelsRunnerPackingHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdRunnersPackingHostsHostId

> ModelsRunnerPackingHost PutOrgsOrgIdRunnersPackingHostsHostId(ctx, orgId, hostId).Body(body).Execute()

Rewrite a packing host declaration (admin)



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
	hostId := "hostId_example" // string | Packing host id
	body := *openapiclient.NewHandlersRunnerPackingHostBody() // HandlersRunnerPackingHostBody | the declaration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersPackingHostsAPI.PutOrgsOrgIdRunnersPackingHostsHostId(context.Background(), orgId, hostId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPackingHostsAPI.PutOrgsOrgIdRunnersPackingHostsHostId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdRunnersPackingHostsHostId`: ModelsRunnerPackingHost
	fmt.Fprintf(os.Stdout, "Response from `RunnersPackingHostsAPI.PutOrgsOrgIdRunnersPackingHostsHostId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**hostId** | **string** | Packing host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdRunnersPackingHostsHostIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersRunnerPackingHostBody**](HandlersRunnerPackingHostBody.md) | the declaration | 

### Return type

[**ModelsRunnerPackingHost**](ModelsRunnerPackingHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

