# \S3ConnectionsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteS3ConnectionsId**](S3ConnectionsAPI.md#DeleteS3ConnectionsId) | **Delete** /s3-connections/{id} | Delete an S3 connection
[**DeleteS3ConnectionsIdBucketsBucketObjectsKey**](S3ConnectionsAPI.md#DeleteS3ConnectionsIdBucketsBucketObjectsKey) | **Delete** /s3-connections/{id}/buckets/{bucket}/objects/{key} | Delete an S3 object (nested path)
[**DeleteS3ConnectionsIdObjects**](S3ConnectionsAPI.md#DeleteS3ConnectionsIdObjects) | **Delete** /s3-connections/{id}/objects | Delete an S3 object
[**GetS3Connections**](S3ConnectionsAPI.md#GetS3Connections) | **Get** /s3-connections | List all S3 connections
[**GetS3ConnectionsIdBuckets**](S3ConnectionsAPI.md#GetS3ConnectionsIdBuckets) | **Get** /s3-connections/{id}/buckets | List buckets for an S3 connection
[**GetS3ConnectionsIdObjects**](S3ConnectionsAPI.md#GetS3ConnectionsIdObjects) | **Get** /s3-connections/{id}/objects | List objects in a bucket
[**GetS3ConnectionsIdObjectsDownload**](S3ConnectionsAPI.md#GetS3ConnectionsIdObjectsDownload) | **Get** /s3-connections/{id}/objects/download | Download an S3 object
[**GetS3ConnectionsProviderInfo**](S3ConnectionsAPI.md#GetS3ConnectionsProviderInfo) | **Get** /s3-connections/provider-info | Get S3 provider information
[**PostS3Connections**](S3ConnectionsAPI.md#PostS3Connections) | **Post** /s3-connections | Create a new S3 connection
[**PostS3ConnectionsTest**](S3ConnectionsAPI.md#PostS3ConnectionsTest) | **Post** /s3-connections/test | Test an S3 connection before saving
[**PutS3ConnectionsIdBucketsBucketObjectsKey**](S3ConnectionsAPI.md#PutS3ConnectionsIdBucketsBucketObjectsKey) | **Put** /s3-connections/{id}/buckets/{bucket}/objects/{key} | Upload an S3 object (nested path)



## DeleteS3ConnectionsId

> DeleteS3ConnectionsId(ctx, id).Execute()

Delete an S3 connection

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
	id := "id_example" // string | S3 Connection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3ConnectionsAPI.DeleteS3ConnectionsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.DeleteS3ConnectionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteS3ConnectionsIdRequest struct via the builder pattern


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


## DeleteS3ConnectionsIdBucketsBucketObjectsKey

> DeleteS3ConnectionsIdBucketsBucketObjectsKey(ctx, id, bucket, key).Execute()

Delete an S3 object (nested path)



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
	id := int32(56) // int32 | S3 Connection ID
	bucket := "bucket_example" // string | Bucket name
	key := "key_example" // string | Object key (may contain '/')

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3ConnectionsAPI.DeleteS3ConnectionsIdBucketsBucketObjectsKey(context.Background(), id, bucket, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.DeleteS3ConnectionsIdBucketsBucketObjectsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | S3 Connection ID | 
**bucket** | **string** | Bucket name | 
**key** | **string** | Object key (may contain &#39;/&#39;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteS3ConnectionsIdBucketsBucketObjectsKeyRequest struct via the builder pattern


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


## DeleteS3ConnectionsIdObjects

> DeleteS3ConnectionsIdObjects(ctx, id).Bucket(bucket).Key(key).Execute()

Delete an S3 object



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
	id := int32(56) // int32 | S3 Connection ID
	bucket := "bucket_example" // string | Bucket name
	key := "key_example" // string | Object key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3ConnectionsAPI.DeleteS3ConnectionsIdObjects(context.Background(), id).Bucket(bucket).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.DeleteS3ConnectionsIdObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteS3ConnectionsIdObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucket** | **string** | Bucket name | 
 **key** | **string** | Object key | 

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


## GetS3Connections

> GetS3Connections(ctx).Execute()

List all S3 connections

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
	r, err := apiClient.S3ConnectionsAPI.GetS3Connections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.GetS3Connections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3ConnectionsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3ConnectionsIdBuckets

> ServicesS3BucketListResult GetS3ConnectionsIdBuckets(ctx, id).Q(q).Execute()

List buckets for an S3 connection



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
	id := int32(56) // int32 | S3 Connection ID
	q := "q_example" // string | Substring filter on bucket name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3ConnectionsAPI.GetS3ConnectionsIdBuckets(context.Background(), id).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.GetS3ConnectionsIdBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3ConnectionsIdBuckets`: ServicesS3BucketListResult
	fmt.Fprintf(os.Stdout, "Response from `S3ConnectionsAPI.GetS3ConnectionsIdBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3ConnectionsIdBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Substring filter on bucket name | 

### Return type

[**ServicesS3BucketListResult**](ServicesS3BucketListResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3ConnectionsIdObjects

> ServicesS3ObjectListResult GetS3ConnectionsIdObjects(ctx, id).Bucket(bucket).Prefix(prefix).ContinuationToken(continuationToken).Q(q).N(n).Delimiter(delimiter).Execute()

List objects in a bucket



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
	id := int32(56) // int32 | S3 Connection ID
	bucket := "bucket_example" // string | Bucket name
	prefix := "prefix_example" // string | Key prefix (optional)
	continuationToken := "continuationToken_example" // string | ListObjectsV2 continuation token (optional)
	q := "q_example" // string | Substring filter on object key (current page) (optional)
	n := int32(56) // int32 | Page size (default 50, max 100) (optional)
	delimiter := "delimiter_example" // string | Delimiter (e.g. / for folder-style prefixes) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3ConnectionsAPI.GetS3ConnectionsIdObjects(context.Background(), id).Bucket(bucket).Prefix(prefix).ContinuationToken(continuationToken).Q(q).N(n).Delimiter(delimiter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.GetS3ConnectionsIdObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3ConnectionsIdObjects`: ServicesS3ObjectListResult
	fmt.Fprintf(os.Stdout, "Response from `S3ConnectionsAPI.GetS3ConnectionsIdObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3ConnectionsIdObjectsRequest struct via the builder pattern


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

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3ConnectionsIdObjectsDownload

> *os.File GetS3ConnectionsIdObjectsDownload(ctx, id).Bucket(bucket).Key(key).Execute()

Download an S3 object



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
	id := int32(56) // int32 | S3 Connection ID
	bucket := "bucket_example" // string | Bucket name
	key := "key_example" // string | Object key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3ConnectionsAPI.GetS3ConnectionsIdObjectsDownload(context.Background(), id).Bucket(bucket).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.GetS3ConnectionsIdObjectsDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3ConnectionsIdObjectsDownload`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `S3ConnectionsAPI.GetS3ConnectionsIdObjectsDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | S3 Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3ConnectionsIdObjectsDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bucket** | **string** | Bucket name | 
 **key** | **string** | Object key | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3ConnectionsProviderInfo

> GetS3ConnectionsProviderInfo(ctx).Provider(provider).Execute()

Get S3 provider information

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
	provider := "provider_example" // string | Provider type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.S3ConnectionsAPI.GetS3ConnectionsProviderInfo(context.Background()).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.GetS3ConnectionsProviderInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetS3ConnectionsProviderInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Provider type | 

### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostS3Connections

> ModelsS3Connection PostS3Connections(ctx).Request(request).Execute()

Create a new S3 connection

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
	request := *openapiclient.NewHandlersCreateS3ConnectionRequest("AccessKeyId_example", "Endpoint_example", "Name_example", "Provider_example", "SecretAccessKey_example") // HandlersCreateS3ConnectionRequest | S3 connection details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3ConnectionsAPI.PostS3Connections(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.PostS3Connections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostS3Connections`: ModelsS3Connection
	fmt.Fprintf(os.Stdout, "Response from `S3ConnectionsAPI.PostS3Connections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostS3ConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersCreateS3ConnectionRequest**](HandlersCreateS3ConnectionRequest.md) | S3 connection details | 

### Return type

[**ModelsS3Connection**](ModelsS3Connection.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [EncryptionKeyAuth](../README.md#EncryptionKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostS3ConnectionsTest

> ServicesTestConnectionResponse PostS3ConnectionsTest(ctx).Request(request).Execute()

Test an S3 connection before saving

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
	request := *openapiclient.NewHandlersTestS3ConnectionRequest("AccessKeyId_example", "Endpoint_example", "Provider_example", "SecretAccessKey_example") // HandlersTestS3ConnectionRequest | S3 connection test details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3ConnectionsAPI.PostS3ConnectionsTest(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.PostS3ConnectionsTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostS3ConnectionsTest`: ServicesTestConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `S3ConnectionsAPI.PostS3ConnectionsTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostS3ConnectionsTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersTestS3ConnectionRequest**](HandlersTestS3ConnectionRequest.md) | S3 connection test details | 

### Return type

[**ServicesTestConnectionResponse**](ServicesTestConnectionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutS3ConnectionsIdBucketsBucketObjectsKey

> ServicesS3UploadResult PutS3ConnectionsIdBucketsBucketObjectsKey(ctx, id, bucket, key).Body(body).Execute()

Upload an S3 object (nested path)



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
	id := int32(56) // int32 | S3 Connection ID
	bucket := "bucket_example" // string | Bucket name
	key := "key_example" // string | Object key (may contain '/')
	body := []int32{int32(123)} // []int32 | Object bytes (≤32 MiB)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3ConnectionsAPI.PutS3ConnectionsIdBucketsBucketObjectsKey(context.Background(), id, bucket, key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3ConnectionsAPI.PutS3ConnectionsIdBucketsBucketObjectsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutS3ConnectionsIdBucketsBucketObjectsKey`: ServicesS3UploadResult
	fmt.Fprintf(os.Stdout, "Response from `S3ConnectionsAPI.PutS3ConnectionsIdBucketsBucketObjectsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | S3 Connection ID | 
**bucket** | **string** | Bucket name | 
**key** | **string** | Object key (may contain &#39;/&#39;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutS3ConnectionsIdBucketsBucketObjectsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **[]int32** | Object bytes (≤32 MiB) | 

### Return type

[**ServicesS3UploadResult**](ServicesS3UploadResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

