# \CloneAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClustersIdClonePlansPlanId**](CloneAPI.md#GetClustersIdClonePlansPlanId) | **Get** /clusters/{id}/clone/plans/{plan_id} | Fetch one clone plan with its review and aggregate
[**GetClustersIdClonePlansPlanIdStatus**](CloneAPI.md#GetClustersIdClonePlansPlanIdStatus) | **Get** /clusters/{id}/clone/plans/{plan_id}/status | The clone&#39;s track surface: report plus live events
[**PostClustersIdClonePlans**](CloneAPI.md#PostClustersIdClonePlans) | **Post** /clusters/{id}/clone/plans | Derive a clone plan from a source cluster into a target
[**PostClustersIdClonePlansPlanIdExecute**](CloneAPI.md#PostClustersIdClonePlansPlanIdExecute) | **Post** /clusters/{id}/clone/plans/{plan_id}/execute | Submit the reviewed plan for execution
[**PutClustersIdClonePlansPlanIdReview**](CloneAPI.md#PutClustersIdClonePlansPlanIdReview) | **Put** /clusters/{id}/clone/plans/{plan_id}/review | Submit the plan review (exclusions, env_local and endpoint



## GetClustersIdClonePlansPlanId

> HandlersClonePlanResponse GetClustersIdClonePlansPlanId(ctx, id, planId).Execute()

Fetch one clone plan with its review and aggregate

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
	id := "id_example" // string | Source cluster ID
	planId := "planId_example" // string | Clone plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloneAPI.GetClustersIdClonePlansPlanId(context.Background(), id, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloneAPI.GetClustersIdClonePlansPlanId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdClonePlansPlanId`: HandlersClonePlanResponse
	fmt.Fprintf(os.Stdout, "Response from `CloneAPI.GetClustersIdClonePlansPlanId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source cluster ID | 
**planId** | **string** | Clone plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdClonePlansPlanIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersClonePlanResponse**](HandlersClonePlanResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdClonePlansPlanIdStatus

> HandlersCloneStatusResponse GetClustersIdClonePlansPlanIdStatus(ctx, id, planId).Execute()

The clone's track surface: report plus live events



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
	id := "id_example" // string | Source cluster ID
	planId := "planId_example" // string | Clone plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloneAPI.GetClustersIdClonePlansPlanIdStatus(context.Background(), id, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloneAPI.GetClustersIdClonePlansPlanIdStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdClonePlansPlanIdStatus`: HandlersCloneStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `CloneAPI.GetClustersIdClonePlansPlanIdStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source cluster ID | 
**planId** | **string** | Clone plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdClonePlansPlanIdStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersCloneStatusResponse**](HandlersCloneStatusResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdClonePlans

> HandlersClonePlanResponse PostClustersIdClonePlans(ctx, id).Body(body).Execute()

Derive a clone plan from a source cluster into a target



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
	id := "id_example" // string | Source cluster ID
	body := *openapiclient.NewHandlersCloneDeriveRequest() // HandlersCloneDeriveRequest | Target cluster

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloneAPI.PostClustersIdClonePlans(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloneAPI.PostClustersIdClonePlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdClonePlans`: HandlersClonePlanResponse
	fmt.Fprintf(os.Stdout, "Response from `CloneAPI.PostClustersIdClonePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdClonePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersCloneDeriveRequest**](HandlersCloneDeriveRequest.md) | Target cluster | 

### Return type

[**HandlersClonePlanResponse**](HandlersClonePlanResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdClonePlansPlanIdExecute

> HandlersCloneExecuteResponse PostClustersIdClonePlansPlanIdExecute(ctx, id, planId).Execute()

Submit the reviewed plan for execution



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
	id := "id_example" // string | Source cluster ID
	planId := "planId_example" // string | Clone plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloneAPI.PostClustersIdClonePlansPlanIdExecute(context.Background(), id, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloneAPI.PostClustersIdClonePlansPlanIdExecute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdClonePlansPlanIdExecute`: HandlersCloneExecuteResponse
	fmt.Fprintf(os.Stdout, "Response from `CloneAPI.PostClustersIdClonePlansPlanIdExecute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source cluster ID | 
**planId** | **string** | Clone plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdClonePlansPlanIdExecuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersCloneExecuteResponse**](HandlersCloneExecuteResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutClustersIdClonePlansPlanIdReview

> HandlersClonePlanResponse PutClustersIdClonePlansPlanIdReview(ctx, id, planId).Body(body).Execute()

Submit the plan review (exclusions, env_local and endpoint



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
	id := "id_example" // string | Source cluster ID
	planId := "planId_example" // string | Clone plan ID
	body := *openapiclient.NewHandlersCloneReviewRequest() // HandlersCloneReviewRequest | Review decisions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloneAPI.PutClustersIdClonePlansPlanIdReview(context.Background(), id, planId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloneAPI.PutClustersIdClonePlansPlanIdReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClustersIdClonePlansPlanIdReview`: HandlersClonePlanResponse
	fmt.Fprintf(os.Stdout, "Response from `CloneAPI.PutClustersIdClonePlansPlanIdReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source cluster ID | 
**planId** | **string** | Clone plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClustersIdClonePlansPlanIdReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersCloneReviewRequest**](HandlersCloneReviewRequest.md) | Review decisions | 

### Return type

[**HandlersClonePlanResponse**](HandlersClonePlanResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

