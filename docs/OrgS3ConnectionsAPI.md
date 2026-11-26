# \OrgS3ConnectionsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgsOrgIdS3ConnectionsId**](OrgS3ConnectionsAPI.md#DeleteOrgsOrgIdS3ConnectionsId) | **Delete** /orgs/{org_id}/s3-connections/{id} | Delete an org-scoped S3 connection
[**DeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoal**](OrgS3ConnectionsAPI.md#DeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoal) | **Delete** /orgs/{org_id}/s3-connections/{id}/buckets/{bucket}/goal | Clear the fill goal for a bucket
[**GetOrgsOrgIdS3Connections**](OrgS3ConnectionsAPI.md#GetOrgsOrgIdS3Connections) | **Get** /orgs/{org_id}/s3-connections | List org-scoped S3 connections
[**GetOrgsOrgIdS3ConnectionsIdBucketsBucketUsage**](OrgS3ConnectionsAPI.md#GetOrgsOrgIdS3ConnectionsIdBucketsBucketUsage) | **Get** /orgs/{org_id}/s3-connections/{id}/buckets/{bucket}/usage | Get bucket usage for an org-scoped S3 connection
[**PostOrgsOrgIdS3Connections**](OrgS3ConnectionsAPI.md#PostOrgsOrgIdS3Connections) | **Post** /orgs/{org_id}/s3-connections | Create an org-scoped S3 connection
[**PostOrgsOrgIdS3ConnectionsIdRotateCredentials**](OrgS3ConnectionsAPI.md#PostOrgsOrgIdS3ConnectionsIdRotateCredentials) | **Post** /orgs/{org_id}/s3-connections/{id}/rotate-credentials | Rotate credentials for an org-scoped S3 connection
[**PutOrgsOrgIdS3ConnectionsId**](OrgS3ConnectionsAPI.md#PutOrgsOrgIdS3ConnectionsId) | **Put** /orgs/{org_id}/s3-connections/{id} | Update an org-scoped S3 connection (non-credential fields only)
[**PutOrgsOrgIdS3ConnectionsIdBucketsBucketGoal**](OrgS3ConnectionsAPI.md#PutOrgsOrgIdS3ConnectionsIdBucketsBucketGoal) | **Put** /orgs/{org_id}/s3-connections/{id}/buckets/{bucket}/goal | Set the fill goal for a bucket



## DeleteOrgsOrgIdS3ConnectionsId

> string DeleteOrgsOrgIdS3ConnectionsId(ctx, orgId, id).Execute()

Delete an org-scoped S3 connection

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
	id := int32(56) // int32 | S3 Connection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.DeleteOrgsOrgIdS3ConnectionsId(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.DeleteOrgsOrgIdS3ConnectionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdS3ConnectionsId`: string
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.DeleteOrgsOrgIdS3ConnectionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdS3ConnectionsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoal

> map[string]interface{} DeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoal(ctx, orgId, id, bucket).Execute()

Clear the fill goal for a bucket



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
	id := int32(56) // int32 | S3 Connection ID
	bucket := "bucket_example" // string | Bucket name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.DeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoal(context.Background(), orgId, id, bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.DeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoal`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.DeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | S3 Connection ID | 
**bucket** | **string** | Bucket name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdS3Connections

> []HandlersOrgS3ConnectionResponse GetOrgsOrgIdS3Connections(ctx, orgId).Execute()

List org-scoped S3 connections

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
	resp, r, err := apiClient.OrgS3ConnectionsAPI.GetOrgsOrgIdS3Connections(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.GetOrgsOrgIdS3Connections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdS3Connections`: []HandlersOrgS3ConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.GetOrgsOrgIdS3Connections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdS3ConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HandlersOrgS3ConnectionResponse**](HandlersOrgS3ConnectionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdS3ConnectionsIdBucketsBucketUsage

> StoragebucketUsageView GetOrgsOrgIdS3ConnectionsIdBucketsBucketUsage(ctx, orgId, id, bucket).Execute()

Get bucket usage for an org-scoped S3 connection



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
	id := int32(56) // int32 | S3 Connection ID
	bucket := "bucket_example" // string | Bucket name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdBucketsBucketUsage(context.Background(), orgId, id, bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdBucketsBucketUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdS3ConnectionsIdBucketsBucketUsage`: StoragebucketUsageView
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdBucketsBucketUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | S3 Connection ID | 
**bucket** | **string** | Bucket name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdS3ConnectionsIdBucketsBucketUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StoragebucketUsageView**](StoragebucketUsageView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdS3Connections

> HandlersOrgS3ConnectionResponse PostOrgsOrgIdS3Connections(ctx, orgId).Request(request).Execute()

Create an org-scoped S3 connection

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
	request := *openapiclient.NewHandlersCreateOrgS3ConnectionRequest("Name_example", "Provider_example") // HandlersCreateOrgS3ConnectionRequest | S3 connection details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.PostOrgsOrgIdS3Connections(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.PostOrgsOrgIdS3Connections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdS3Connections`: HandlersOrgS3ConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.PostOrgsOrgIdS3Connections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdS3ConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersCreateOrgS3ConnectionRequest**](HandlersCreateOrgS3ConnectionRequest.md) | S3 connection details | 

### Return type

[**HandlersOrgS3ConnectionResponse**](HandlersOrgS3ConnectionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdS3ConnectionsIdRotateCredentials

> HandlersOrgS3ConnectionResponse PostOrgsOrgIdS3ConnectionsIdRotateCredentials(ctx, orgId, id).Request(request).Execute()

Rotate credentials for an org-scoped S3 connection

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
	id := int32(56) // int32 | S3 Connection ID
	request := *openapiclient.NewHandlersRotateOrgS3ConnectionCredentialsRequest() // HandlersRotateOrgS3ConnectionCredentialsRequest | New credentials

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.PostOrgsOrgIdS3ConnectionsIdRotateCredentials(context.Background(), orgId, id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.PostOrgsOrgIdS3ConnectionsIdRotateCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdS3ConnectionsIdRotateCredentials`: HandlersOrgS3ConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.PostOrgsOrgIdS3ConnectionsIdRotateCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdS3ConnectionsIdRotateCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**HandlersRotateOrgS3ConnectionCredentialsRequest**](HandlersRotateOrgS3ConnectionCredentialsRequest.md) | New credentials | 

### Return type

[**HandlersOrgS3ConnectionResponse**](HandlersOrgS3ConnectionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdS3ConnectionsId

> HandlersOrgS3ConnectionResponse PutOrgsOrgIdS3ConnectionsId(ctx, orgId, id).Request(request).Execute()

Update an org-scoped S3 connection (non-credential fields only)

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
	id := int32(56) // int32 | S3 Connection ID
	request := *openapiclient.NewHandlersUpdateOrgS3ConnectionRequest() // HandlersUpdateOrgS3ConnectionRequest | Fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.PutOrgsOrgIdS3ConnectionsId(context.Background(), orgId, id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.PutOrgsOrgIdS3ConnectionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdS3ConnectionsId`: HandlersOrgS3ConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.PutOrgsOrgIdS3ConnectionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdS3ConnectionsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**HandlersUpdateOrgS3ConnectionRequest**](HandlersUpdateOrgS3ConnectionRequest.md) | Fields to update | 

### Return type

[**HandlersOrgS3ConnectionResponse**](HandlersOrgS3ConnectionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdS3ConnectionsIdBucketsBucketGoal

> StoragebucketBucketGoalView PutOrgsOrgIdS3ConnectionsIdBucketsBucketGoal(ctx, orgId, id, bucket).Body(body).Execute()

Set the fill goal for a bucket



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
	id := int32(56) // int32 | S3 Connection ID
	bucket := "bucket_example" // string | Bucket name
	body := *openapiclient.NewHandlersBucketGoalRequest() // HandlersBucketGoalRequest | The goal in bytes (must be positive)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.PutOrgsOrgIdS3ConnectionsIdBucketsBucketGoal(context.Background(), orgId, id, bucket).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.PutOrgsOrgIdS3ConnectionsIdBucketsBucketGoal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdS3ConnectionsIdBucketsBucketGoal`: StoragebucketBucketGoalView
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.PutOrgsOrgIdS3ConnectionsIdBucketsBucketGoal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | S3 Connection ID | 
**bucket** | **string** | Bucket name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdS3ConnectionsIdBucketsBucketGoalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**HandlersBucketGoalRequest**](HandlersBucketGoalRequest.md) | The goal in bytes (must be positive) | 

### Return type

[**StoragebucketBucketGoalView**](StoragebucketBucketGoalView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

