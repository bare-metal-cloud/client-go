# \RunnersConnectionsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersConnections**](RunnersConnectionsAPI.md#GetOrgsOrgIdRunnersConnections) | **Get** /orgs/{org_id}/runners/connections | List the org&#39;s GitHub connections — every installation with its derived state and repository count
[**GetOrgsOrgIdRunnersConnectionsConnectionId**](RunnersConnectionsAPI.md#GetOrgsOrgIdRunnersConnectionsConnectionId) | **Get** /orgs/{org_id}/runners/connections/{connection_id} | One connection — its repositories with mapping honesty flags, the org&#39;s runner ids, the profiles in use
[**PutOrgsOrgIdRunnersConnectionsConnectionIdMapping**](RunnersConnectionsAPI.md#PutOrgsOrgIdRunnersConnectionsConnectionIdMapping) | **Put** /orgs/{org_id}/runners/connections/{connection_id}/mapping | Save the connection&#39;s repository-to-profile mapping (replaces the whole set)



## GetOrgsOrgIdRunnersConnections

> HandlersRunnerConnectionsResponse GetOrgsOrgIdRunnersConnections(ctx, orgId).Execute()

List the org's GitHub connections — every installation with its derived state and repository count



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
	resp, r, err := apiClient.RunnersConnectionsAPI.GetOrgsOrgIdRunnersConnections(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersConnectionsAPI.GetOrgsOrgIdRunnersConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersConnections`: HandlersRunnerConnectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersConnectionsAPI.GetOrgsOrgIdRunnersConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerConnectionsResponse**](HandlersRunnerConnectionsResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersConnectionsConnectionId

> HandlersRunnerConnectionDetailResponse GetOrgsOrgIdRunnersConnectionsConnectionId(ctx, orgId, connectionId).Execute()

One connection — its repositories with mapping honesty flags, the org's runner ids, the profiles in use



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
	connectionId := int32(56) // int32 | The connection id (the installation row id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersConnectionsAPI.GetOrgsOrgIdRunnersConnectionsConnectionId(context.Background(), orgId, connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersConnectionsAPI.GetOrgsOrgIdRunnersConnectionsConnectionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersConnectionsConnectionId`: HandlersRunnerConnectionDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersConnectionsAPI.GetOrgsOrgIdRunnersConnectionsConnectionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**connectionId** | **int32** | The connection id (the installation row id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersConnectionsConnectionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerConnectionDetailResponse**](HandlersRunnerConnectionDetailResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdRunnersConnectionsConnectionIdMapping

> HandlersRunnerConnectionDetailResponse PutOrgsOrgIdRunnersConnectionsConnectionIdMapping(ctx, orgId, connectionId).Body(body).Execute()

Save the connection's repository-to-profile mapping (replaces the whole set)



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
	connectionId := int32(56) // int32 | The connection id (the installation row id)
	body := *openapiclient.NewHandlersSaveConnectionMappingBody() // HandlersSaveConnectionMappingBody | The full repository list with its profile answers

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersConnectionsAPI.PutOrgsOrgIdRunnersConnectionsConnectionIdMapping(context.Background(), orgId, connectionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersConnectionsAPI.PutOrgsOrgIdRunnersConnectionsConnectionIdMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdRunnersConnectionsConnectionIdMapping`: HandlersRunnerConnectionDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersConnectionsAPI.PutOrgsOrgIdRunnersConnectionsConnectionIdMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**connectionId** | **int32** | The connection id (the installation row id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdRunnersConnectionsConnectionIdMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersSaveConnectionMappingBody**](HandlersSaveConnectionMappingBody.md) | The full repository list with its profile answers | 

### Return type

[**HandlersRunnerConnectionDetailResponse**](HandlersRunnerConnectionDetailResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

