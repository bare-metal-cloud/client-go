# \RunnersPackingPoolsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersPackingPools**](RunnersPackingPoolsAPI.md#GetOrgsOrgIdRunnersPackingPools) | **Get** /orgs/{org_id}/runners/packing-pools | List the org&#39;s packing pools
[**GetOrgsOrgIdRunnersPackingPoolsPoolIdSettings**](RunnersPackingPoolsAPI.md#GetOrgsOrgIdRunnersPackingPoolsPoolIdSettings) | **Get** /orgs/{org_id}/runners/packing-pools/{pool_id}/settings | Get a packing pool&#39;s per-pool settings
[**GetOrgsOrgIdRunnersPackingPoolsPoolIdSlices**](RunnersPackingPoolsAPI.md#GetOrgsOrgIdRunnersPackingPoolsPoolIdSlices) | **Get** /orgs/{org_id}/runners/packing-pools/{pool_id}/slices | One packing pool&#39;s slices
[**PutOrgsOrgIdRunnersPackingPoolsPoolIdSettings**](RunnersPackingPoolsAPI.md#PutOrgsOrgIdRunnersPackingPoolsPoolIdSettings) | **Put** /orgs/{org_id}/runners/packing-pools/{pool_id}/settings | Set a packing pool&#39;s per-pool settings (admin)



## GetOrgsOrgIdRunnersPackingPools

> HandlersRunnerPackingPoolsResponse GetOrgsOrgIdRunnersPackingPools(ctx, orgId).Execute()

List the org's packing pools



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
	resp, r, err := apiClient.RunnersPackingPoolsAPI.GetOrgsOrgIdRunnersPackingPools(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPackingPoolsAPI.GetOrgsOrgIdRunnersPackingPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersPackingPools`: HandlersRunnerPackingPoolsResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersPackingPoolsAPI.GetOrgsOrgIdRunnersPackingPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersPackingPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerPackingPoolsResponse**](HandlersRunnerPackingPoolsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersPackingPoolsPoolIdSettings

> HandlersRunnerPoolHostSettingsBody GetOrgsOrgIdRunnersPackingPoolsPoolIdSettings(ctx, orgId, poolId).Execute()

Get a packing pool's per-pool settings



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
	poolId := "poolId_example" // string | Packing host id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersPackingPoolsAPI.GetOrgsOrgIdRunnersPackingPoolsPoolIdSettings(context.Background(), orgId, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPackingPoolsAPI.GetOrgsOrgIdRunnersPackingPoolsPoolIdSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersPackingPoolsPoolIdSettings`: HandlersRunnerPoolHostSettingsBody
	fmt.Fprintf(os.Stdout, "Response from `RunnersPackingPoolsAPI.GetOrgsOrgIdRunnersPackingPoolsPoolIdSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**poolId** | **string** | Packing host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersPackingPoolsPoolIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerPoolHostSettingsBody**](HandlersRunnerPoolHostSettingsBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersPackingPoolsPoolIdSlices

> HandlersRunnerPackingSlicesResponse GetOrgsOrgIdRunnersPackingPoolsPoolIdSlices(ctx, orgId, poolId).Execute()

One packing pool's slices



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
	poolId := "poolId_example" // string | Packing host ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersPackingPoolsAPI.GetOrgsOrgIdRunnersPackingPoolsPoolIdSlices(context.Background(), orgId, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPackingPoolsAPI.GetOrgsOrgIdRunnersPackingPoolsPoolIdSlices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersPackingPoolsPoolIdSlices`: HandlersRunnerPackingSlicesResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersPackingPoolsAPI.GetOrgsOrgIdRunnersPackingPoolsPoolIdSlices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**poolId** | **string** | Packing host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersPackingPoolsPoolIdSlicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerPackingSlicesResponse**](HandlersRunnerPackingSlicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdRunnersPackingPoolsPoolIdSettings

> HandlersRunnerPoolHostSettingsBody PutOrgsOrgIdRunnersPackingPoolsPoolIdSettings(ctx, orgId, poolId).Body(body).Execute()

Set a packing pool's per-pool settings (admin)



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
	poolId := "poolId_example" // string | Packing host id
	body := *openapiclient.NewHandlersRunnerPoolHostSettingsBody() // HandlersRunnerPoolHostSettingsBody | the per-pool settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersPackingPoolsAPI.PutOrgsOrgIdRunnersPackingPoolsPoolIdSettings(context.Background(), orgId, poolId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersPackingPoolsAPI.PutOrgsOrgIdRunnersPackingPoolsPoolIdSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdRunnersPackingPoolsPoolIdSettings`: HandlersRunnerPoolHostSettingsBody
	fmt.Fprintf(os.Stdout, "Response from `RunnersPackingPoolsAPI.PutOrgsOrgIdRunnersPackingPoolsPoolIdSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**poolId** | **string** | Packing host id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdRunnersPackingPoolsPoolIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersRunnerPoolHostSettingsBody**](HandlersRunnerPoolHostSettingsBody.md) | the per-pool settings | 

### Return type

[**HandlersRunnerPoolHostSettingsBody**](HandlersRunnerPoolHostSettingsBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

