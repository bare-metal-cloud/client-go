# \OrgS3ConnectionsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgsOrgIdS3ConnectionsId**](OrgS3ConnectionsAPI.md#DeleteOrgsOrgIdS3ConnectionsId) | **Delete** /orgs/{org_id}/s3-connections/{id} | Delete an org-scoped S3 connection
[**DeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoal**](OrgS3ConnectionsAPI.md#DeleteOrgsOrgIdS3ConnectionsIdBucketsBucketGoal) | **Delete** /orgs/{org_id}/s3-connections/{id}/buckets/{bucket}/goal | Clear the fill goal for a bucket
[**GetOrgsOrgIdS3Connections**](OrgS3ConnectionsAPI.md#GetOrgsOrgIdS3Connections) | **Get** /orgs/{org_id}/s3-connections | List org-scoped S3 connections
[**GetOrgsOrgIdS3ConnectionsIdBuckets**](OrgS3ConnectionsAPI.md#GetOrgsOrgIdS3ConnectionsIdBuckets) | **Get** /orgs/{org_id}/s3-connections/{id}/buckets | List buckets for an org-scoped S3 connection
[**GetOrgsOrgIdS3ConnectionsIdBucketsBucketUsage**](OrgS3ConnectionsAPI.md#GetOrgsOrgIdS3ConnectionsIdBucketsBucketUsage) | **Get** /orgs/{org_id}/s3-connections/{id}/buckets/{bucket}/usage | Get bucket usage for an org-scoped S3 connection
[**GetOrgsOrgIdS3ConnectionsIdObjects**](OrgS3ConnectionsAPI.md#GetOrgsOrgIdS3ConnectionsIdObjects) | **Get** /orgs/{org_id}/s3-connections/{id}/objects | List objects in a bucket of an org-scoped S3 connection
[**GetOrgsOrgIdS3ConnectionsIdObjectsDownload**](OrgS3ConnectionsAPI.md#GetOrgsOrgIdS3ConnectionsIdObjectsDownload) | **Get** /orgs/{org_id}/s3-connections/{id}/objects/download | Download an object from an org-scoped S3 connection
[**PostOrgsOrgIdS3Connections**](OrgS3ConnectionsAPI.md#PostOrgsOrgIdS3Connections) | **Post** /orgs/{org_id}/s3-connections | Create an org-scoped S3 connection
[**PostOrgsOrgIdS3ConnectionsIdRotateCredentials**](OrgS3ConnectionsAPI.md#PostOrgsOrgIdS3ConnectionsIdRotateCredentials) | **Post** /orgs/{org_id}/s3-connections/{id}/rotate-credentials | Rotate credentials for an org-scoped S3 connection
[**PutOrgsOrgIdS3ConnectionsId**](OrgS3ConnectionsAPI.md#PutOrgsOrgIdS3ConnectionsId) | **Put** /orgs/{org_id}/s3-connections/{id} | Update an org-scoped S3 connection (non-credential fields only)
[**PutOrgsOrgIdS3ConnectionsIdBucketsBucketGoal**](OrgS3ConnectionsAPI.md#PutOrgsOrgIdS3ConnectionsIdBucketsBucketGoal) | **Put** /orgs/{org_id}/s3-connections/{id}/buckets/{bucket}/goal | Set the fill goal for a bucket



## DeleteOrgsOrgIdS3ConnectionsId

> string DeleteOrgsOrgIdS3ConnectionsId(ctx, orgId, id).Force(force).Execute()

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
	force := true // bool | Force the delete past the secret-reference guard (admin/owner only); a backup-plan or managed-database consumer is never forceable (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.DeleteOrgsOrgIdS3ConnectionsId(context.Background(), orgId, id).Force(force).Execute()
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


 **force** | **bool** | Force the delete past the secret-reference guard (admin/owner only); a backup-plan or managed-database consumer is never forceable | 

### Return type

**string**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

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

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdS3ConnectionsIdBuckets

> ServicesS3BucketListResult GetOrgsOrgIdS3ConnectionsIdBuckets(ctx, orgId, id).Q(q).Execute()

List buckets for an org-scoped S3 connection



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
	q := "q_example" // string | Substring filter on bucket name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdBuckets(context.Background(), orgId, id).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdS3ConnectionsIdBuckets`: ServicesS3BucketListResult
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdS3ConnectionsIdBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** | Substring filter on bucket name | 

### Return type

[**ServicesS3BucketListResult**](ServicesS3BucketListResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

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


## GetOrgsOrgIdS3ConnectionsIdObjects

> ServicesS3ObjectListResult GetOrgsOrgIdS3ConnectionsIdObjects(ctx, orgId, id).Bucket(bucket).Prefix(prefix).ContinuationToken(continuationToken).Q(q).N(n).Delimiter(delimiter).Execute()

List objects in a bucket of an org-scoped S3 connection



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
	prefix := "prefix_example" // string | Key prefix (optional)
	continuationToken := "continuationToken_example" // string | ListObjectsV2 continuation token (optional)
	q := "q_example" // string | Substring filter on object key (current page) (optional)
	n := int32(56) // int32 | Page size (default 50, max 100) (optional)
	delimiter := "delimiter_example" // string | Delimiter (e.g. / for folder-style prefixes) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdObjects(context.Background(), orgId, id).Bucket(bucket).Prefix(prefix).ContinuationToken(continuationToken).Q(q).N(n).Delimiter(delimiter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdS3ConnectionsIdObjects`: ServicesS3ObjectListResult
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdS3ConnectionsIdObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bucket** | **string** | Bucket name | 
 **prefix** | **string** | Key prefix | 
 **continuationToken** | **string** | ListObjectsV2 continuation token | 
 **q** | **string** | Substring filter on object key (current page) | 
 **n** | **int32** | Page size (default 50, max 100) | 
 **delimiter** | **string** | Delimiter (e.g. / for folder-style prefixes) | 

### Return type

[**ServicesS3ObjectListResult**](ServicesS3ObjectListResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdS3ConnectionsIdObjectsDownload

> *os.File GetOrgsOrgIdS3ConnectionsIdObjectsDownload(ctx, orgId, id).Bucket(bucket).Key(key).Execute()

Download an object from an org-scoped S3 connection



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
	key := "key_example" // string | Object key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdObjectsDownload(context.Background(), orgId, id).Bucket(bucket).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdObjectsDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdS3ConnectionsIdObjectsDownload`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrgS3ConnectionsAPI.GetOrgsOrgIdS3ConnectionsIdObjectsDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdS3ConnectionsIdObjectsDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bucket** | **string** | Bucket name | 
 **key** | **string** | Object key | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

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

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

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

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

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

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

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

