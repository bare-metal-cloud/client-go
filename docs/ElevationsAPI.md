# \ElevationsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApiV1OrgsOrgIdElevationApproversUserId**](ElevationsAPI.md#DeleteApiV1OrgsOrgIdElevationApproversUserId) | **Delete** /api/v1/orgs/{org_id}/elevation-approvers/{user_id} | Remove an elevation approver
[**GetApiV1OrgsOrgIdElevationApprovers**](ElevationsAPI.md#GetApiV1OrgsOrgIdElevationApprovers) | **Get** /api/v1/orgs/{org_id}/elevation-approvers | List elevation approvers
[**GetApiV1OrgsOrgIdElevations**](ElevationsAPI.md#GetApiV1OrgsOrgIdElevations) | **Get** /api/v1/orgs/{org_id}/elevations | List admin elevations
[**GetApiV1OrgsOrgIdElevationsCurrent**](ElevationsAPI.md#GetApiV1OrgsOrgIdElevationsCurrent) | **Get** /api/v1/orgs/{org_id}/elevations/current | Current admin elevation
[**PostApiV1OrgsOrgIdElevations**](ElevationsAPI.md#PostApiV1OrgsOrgIdElevations) | **Post** /api/v1/orgs/{org_id}/elevations | Request admin elevation
[**PostApiV1OrgsOrgIdElevationsIdApprove**](ElevationsAPI.md#PostApiV1OrgsOrgIdElevationsIdApprove) | **Post** /api/v1/orgs/{org_id}/elevations/{id}/approve | Approve admin elevation
[**PostApiV1OrgsOrgIdElevationsIdDeny**](ElevationsAPI.md#PostApiV1OrgsOrgIdElevationsIdDeny) | **Post** /api/v1/orgs/{org_id}/elevations/{id}/deny | Deny admin elevation
[**PostApiV1OrgsOrgIdElevationsIdEnd**](ElevationsAPI.md#PostApiV1OrgsOrgIdElevationsIdEnd) | **Post** /api/v1/orgs/{org_id}/elevations/{id}/end | End or cancel admin elevation
[**PutApiV1OrgsOrgIdElevationApproversUserId**](ElevationsAPI.md#PutApiV1OrgsOrgIdElevationApproversUserId) | **Put** /api/v1/orgs/{org_id}/elevation-approvers/{user_id} | Add an elevation approver



## DeleteApiV1OrgsOrgIdElevationApproversUserId

> DeleteApiV1OrgsOrgIdElevationApproversUserId(ctx, orgId, userId).Execute()

Remove an elevation approver

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
	userId := int32(56) // int32 | Admin user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ElevationsAPI.DeleteApiV1OrgsOrgIdElevationApproversUserId(context.Background(), orgId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElevationsAPI.DeleteApiV1OrgsOrgIdElevationApproversUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**userId** | **int32** | Admin user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiV1OrgsOrgIdElevationApproversUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1OrgsOrgIdElevationApprovers

> []HandlersApproverView GetApiV1OrgsOrgIdElevationApprovers(ctx, orgId).Execute()

List elevation approvers

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
	resp, r, err := apiClient.ElevationsAPI.GetApiV1OrgsOrgIdElevationApprovers(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElevationsAPI.GetApiV1OrgsOrgIdElevationApprovers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1OrgsOrgIdElevationApprovers`: []HandlersApproverView
	fmt.Fprintf(os.Stdout, "Response from `ElevationsAPI.GetApiV1OrgsOrgIdElevationApprovers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1OrgsOrgIdElevationApproversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HandlersApproverView**](HandlersApproverView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1OrgsOrgIdElevations

> []HandlersElevationView GetApiV1OrgsOrgIdElevations(ctx, orgId).Status(status).UserId(userId).Execute()

List admin elevations

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
	status := "status_example" // string | Filter by status (optional)
	userId := int32(56) // int32 | Filter by user (Owner and Admins only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElevationsAPI.GetApiV1OrgsOrgIdElevations(context.Background(), orgId).Status(status).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElevationsAPI.GetApiV1OrgsOrgIdElevations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1OrgsOrgIdElevations`: []HandlersElevationView
	fmt.Fprintf(os.Stdout, "Response from `ElevationsAPI.GetApiV1OrgsOrgIdElevations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1OrgsOrgIdElevationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter by status | 
 **userId** | **int32** | Filter by user (Owner and Admins only) | 

### Return type

[**[]HandlersElevationView**](HandlersElevationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1OrgsOrgIdElevationsCurrent

> HandlersElevationView GetApiV1OrgsOrgIdElevationsCurrent(ctx, orgId).Execute()

Current admin elevation

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
	resp, r, err := apiClient.ElevationsAPI.GetApiV1OrgsOrgIdElevationsCurrent(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElevationsAPI.GetApiV1OrgsOrgIdElevationsCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1OrgsOrgIdElevationsCurrent`: HandlersElevationView
	fmt.Fprintf(os.Stdout, "Response from `ElevationsAPI.GetApiV1OrgsOrgIdElevationsCurrent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1OrgsOrgIdElevationsCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersElevationView**](HandlersElevationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1OrgsOrgIdElevations

> HandlersElevationView PostApiV1OrgsOrgIdElevations(ctx, orgId).Body(body).Execute()

Request admin elevation



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
	body := map[string]interface{}{ ... } // map[string]interface{} | reason (10-500 chars) and duration_minutes (5-60, default 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElevationsAPI.PostApiV1OrgsOrgIdElevations(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElevationsAPI.PostApiV1OrgsOrgIdElevations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1OrgsOrgIdElevations`: HandlersElevationView
	fmt.Fprintf(os.Stdout, "Response from `ElevationsAPI.PostApiV1OrgsOrgIdElevations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1OrgsOrgIdElevationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | reason (10-500 chars) and duration_minutes (5-60, default 30) | 

### Return type

[**HandlersElevationView**](HandlersElevationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1OrgsOrgIdElevationsIdApprove

> HandlersElevationView PostApiV1OrgsOrgIdElevationsIdApprove(ctx, orgId, id).Execute()

Approve admin elevation

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
	id := int32(56) // int32 | Elevation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElevationsAPI.PostApiV1OrgsOrgIdElevationsIdApprove(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElevationsAPI.PostApiV1OrgsOrgIdElevationsIdApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1OrgsOrgIdElevationsIdApprove`: HandlersElevationView
	fmt.Fprintf(os.Stdout, "Response from `ElevationsAPI.PostApiV1OrgsOrgIdElevationsIdApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Elevation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1OrgsOrgIdElevationsIdApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersElevationView**](HandlersElevationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1OrgsOrgIdElevationsIdDeny

> HandlersElevationView PostApiV1OrgsOrgIdElevationsIdDeny(ctx, orgId, id).Body(body).Execute()

Deny admin elevation

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
	id := int32(56) // int32 | Elevation ID
	body := map[string]interface{}{ ... } // map[string]interface{} | reason (optional) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElevationsAPI.PostApiV1OrgsOrgIdElevationsIdDeny(context.Background(), orgId, id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElevationsAPI.PostApiV1OrgsOrgIdElevationsIdDeny``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1OrgsOrgIdElevationsIdDeny`: HandlersElevationView
	fmt.Fprintf(os.Stdout, "Response from `ElevationsAPI.PostApiV1OrgsOrgIdElevationsIdDeny`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Elevation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1OrgsOrgIdElevationsIdDenyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | reason (optional) | 

### Return type

[**HandlersElevationView**](HandlersElevationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiV1OrgsOrgIdElevationsIdEnd

> HandlersElevationView PostApiV1OrgsOrgIdElevationsIdEnd(ctx, orgId, id).Execute()

End or cancel admin elevation

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
	id := int32(56) // int32 | Elevation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElevationsAPI.PostApiV1OrgsOrgIdElevationsIdEnd(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElevationsAPI.PostApiV1OrgsOrgIdElevationsIdEnd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1OrgsOrgIdElevationsIdEnd`: HandlersElevationView
	fmt.Fprintf(os.Stdout, "Response from `ElevationsAPI.PostApiV1OrgsOrgIdElevationsIdEnd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Elevation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1OrgsOrgIdElevationsIdEndRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersElevationView**](HandlersElevationView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApiV1OrgsOrgIdElevationApproversUserId

> map[string]interface{} PutApiV1OrgsOrgIdElevationApproversUserId(ctx, orgId, userId).Execute()

Add an elevation approver

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
	userId := int32(56) // int32 | Admin user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElevationsAPI.PutApiV1OrgsOrgIdElevationApproversUserId(context.Background(), orgId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElevationsAPI.PutApiV1OrgsOrgIdElevationApproversUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApiV1OrgsOrgIdElevationApproversUserId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ElevationsAPI.PutApiV1OrgsOrgIdElevationApproversUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**userId** | **int32** | Admin user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1OrgsOrgIdElevationApproversUserIdRequest struct via the builder pattern


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

