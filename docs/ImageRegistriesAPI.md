# \ImageRegistriesAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteImageRegistriesId**](ImageRegistriesAPI.md#DeleteImageRegistriesId) | **Delete** /image-registries/{id} | Delete an image registry
[**DeleteImageRegistriesIdRepositoriesNameManifestsRef**](ImageRegistriesAPI.md#DeleteImageRegistriesIdRepositoriesNameManifestsRef) | **Delete** /image-registries/{id}/repositories/{name}/manifests/{ref} | Delete a manifest by tag or digest
[**DeleteImageRegistriesIdRepositoriesNameTagsTag**](ImageRegistriesAPI.md#DeleteImageRegistriesIdRepositoriesNameTagsTag) | **Delete** /image-registries/{id}/repositories/{name}/tags/{tag} | Delete a repository tag
[**GetImageRegistries**](ImageRegistriesAPI.md#GetImageRegistries) | **Get** /image-registries | List all image registries
[**GetImageRegistriesId**](ImageRegistriesAPI.md#GetImageRegistriesId) | **Get** /image-registries/{id} | Get an image registry by ID
[**GetImageRegistriesIdRepositories**](ImageRegistriesAPI.md#GetImageRegistriesIdRepositories) | **Get** /image-registries/{id}/repositories | List repositories in an image registry
[**GetImageRegistriesIdRepositoriesNameManifestsRef**](ImageRegistriesAPI.md#GetImageRegistriesIdRepositoriesNameManifestsRef) | **Get** /image-registries/{id}/repositories/{name}/manifests/{ref} | Get manifest metadata for a tag or digest
[**GetImageRegistriesIdRepositoriesNameTags**](ImageRegistriesAPI.md#GetImageRegistriesIdRepositoriesNameTags) | **Get** /image-registries/{id}/repositories/{name}/tags | List tags for a repository
[**GetImageRegistriesRegistryInfo**](ImageRegistriesAPI.md#GetImageRegistriesRegistryInfo) | **Get** /image-registries/registry-info | Get registry type information
[**PatchImageRegistriesId**](ImageRegistriesAPI.md#PatchImageRegistriesId) | **Patch** /image-registries/{id} | Patch image registry credentials / cloud FK
[**PostImageRegistries**](ImageRegistriesAPI.md#PostImageRegistries) | **Post** /image-registries | Create a new image registry
[**PostImageRegistriesTest**](ImageRegistriesAPI.md#PostImageRegistriesTest) | **Post** /image-registries/test | Test an image registry connection



## DeleteImageRegistriesId

> DeleteImageRegistriesId(ctx, id).Execute()

Delete an image registry



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
	id := "id_example" // string | Registry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageRegistriesAPI.DeleteImageRegistriesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.DeleteImageRegistriesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRegistriesIdRequest struct via the builder pattern


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


## DeleteImageRegistriesIdRepositoriesNameManifestsRef

> DeleteImageRegistriesIdRepositoriesNameManifestsRef(ctx, id, name, ref).Execute()

Delete a manifest by tag or digest



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
	id := int32(56) // int32 | Registry ID
	name := "name_example" // string | Repository name (URL-encoded)
	ref := "ref_example" // string | Tag or digest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageRegistriesAPI.DeleteImageRegistriesIdRepositoriesNameManifestsRef(context.Background(), id, name, ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.DeleteImageRegistriesIdRepositoriesNameManifestsRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Registry ID | 
**name** | **string** | Repository name (URL-encoded) | 
**ref** | **string** | Tag or digest | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRegistriesIdRepositoriesNameManifestsRefRequest struct via the builder pattern


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


## DeleteImageRegistriesIdRepositoriesNameTagsTag

> DeleteImageRegistriesIdRepositoriesNameTagsTag(ctx, id, name, tag).Execute()

Delete a repository tag



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
	id := int32(56) // int32 | Registry ID
	name := "name_example" // string | Repository name (URL-encoded)
	tag := "tag_example" // string | Tag name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageRegistriesAPI.DeleteImageRegistriesIdRepositoriesNameTagsTag(context.Background(), id, name, tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.DeleteImageRegistriesIdRepositoriesNameTagsTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Registry ID | 
**name** | **string** | Repository name (URL-encoded) | 
**tag** | **string** | Tag name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRegistriesIdRepositoriesNameTagsTagRequest struct via the builder pattern


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


## GetImageRegistries

> []ModelsImageRegistry GetImageRegistries(ctx).Execute()

List all image registries



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
	resp, r, err := apiClient.ImageRegistriesAPI.GetImageRegistries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.GetImageRegistries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageRegistries`: []ModelsImageRegistry
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistriesAPI.GetImageRegistries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRegistriesRequest struct via the builder pattern


### Return type

[**[]ModelsImageRegistry**](ModelsImageRegistry.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageRegistriesId

> ModelsImageRegistry GetImageRegistriesId(ctx, id).Execute()

Get an image registry by ID



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
	id := "id_example" // string | Registry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageRegistriesAPI.GetImageRegistriesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.GetImageRegistriesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageRegistriesId`: ModelsImageRegistry
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistriesAPI.GetImageRegistriesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRegistriesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsImageRegistry**](ModelsImageRegistry.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageRegistriesIdRepositories

> ServicesCatalogListResult GetImageRegistriesIdRepositories(ctx, id).Q(q).N(n).Last(last).Execute()

List repositories in an image registry



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
	id := int32(56) // int32 | Registry ID
	q := "q_example" // string | Substring filter on repository name (optional)
	n := int32(56) // int32 | Page size (default 50, max 100) (optional)
	last := "last_example" // string | Pagination cursor (OCI last= or ECR NextToken) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageRegistriesAPI.GetImageRegistriesIdRepositories(context.Background(), id).Q(q).N(n).Last(last).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.GetImageRegistriesIdRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageRegistriesIdRepositories`: ServicesCatalogListResult
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistriesAPI.GetImageRegistriesIdRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRegistriesIdRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Substring filter on repository name | 
 **n** | **int32** | Page size (default 50, max 100) | 
 **last** | **string** | Pagination cursor (OCI last&#x3D; or ECR NextToken) | 

### Return type

[**ServicesCatalogListResult**](ServicesCatalogListResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageRegistriesIdRepositoriesNameManifestsRef

> ServicesManifestMetadata GetImageRegistriesIdRepositoriesNameManifestsRef(ctx, id, name, ref).Execute()

Get manifest metadata for a tag or digest



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
	id := int32(56) // int32 | Registry ID
	name := "name_example" // string | Repository name (URL-encoded)
	ref := "ref_example" // string | Tag or digest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageRegistriesAPI.GetImageRegistriesIdRepositoriesNameManifestsRef(context.Background(), id, name, ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.GetImageRegistriesIdRepositoriesNameManifestsRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageRegistriesIdRepositoriesNameManifestsRef`: ServicesManifestMetadata
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistriesAPI.GetImageRegistriesIdRepositoriesNameManifestsRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Registry ID | 
**name** | **string** | Repository name (URL-encoded) | 
**ref** | **string** | Tag or digest | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRegistriesIdRepositoriesNameManifestsRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServicesManifestMetadata**](ServicesManifestMetadata.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageRegistriesIdRepositoriesNameTags

> ServicesTagsListResult GetImageRegistriesIdRepositoriesNameTags(ctx, id, name).N(n).Last(last).Execute()

List tags for a repository



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
	id := int32(56) // int32 | Registry ID
	name := "name_example" // string | Repository name (URL-encoded)
	n := int32(56) // int32 | Page size (default 50, max 100) (optional)
	last := "last_example" // string | Tags list cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageRegistriesAPI.GetImageRegistriesIdRepositoriesNameTags(context.Background(), id, name).N(n).Last(last).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.GetImageRegistriesIdRepositoriesNameTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageRegistriesIdRepositoriesNameTags`: ServicesTagsListResult
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistriesAPI.GetImageRegistriesIdRepositoriesNameTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Registry ID | 
**name** | **string** | Repository name (URL-encoded) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRegistriesIdRepositoriesNameTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **n** | **int32** | Page size (default 50, max 100) | 
 **last** | **string** | Tags list cursor | 

### Return type

[**ServicesTagsListResult**](ServicesTagsListResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageRegistriesRegistryInfo

> map[string]interface{} GetImageRegistriesRegistryInfo(ctx).Type_(type_).Execute()

Get registry type information



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
	type_ := "type__example" // string | Registry type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageRegistriesAPI.GetImageRegistriesRegistryInfo(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.GetImageRegistriesRegistryInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageRegistriesRegistryInfo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistriesAPI.GetImageRegistriesRegistryInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRegistriesRegistryInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Registry type | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchImageRegistriesId

> map[string]interface{} PatchImageRegistriesId(ctx, id).Request(request).Execute()

Patch image registry credentials / cloud FK



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
	id := int32(56) // int32 | Registry ID
	request := *openapiclient.NewHandlersPatchImageRegistryRequest() // HandlersPatchImageRegistryRequest | Credential / FK updates

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageRegistriesAPI.PatchImageRegistriesId(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.PatchImageRegistriesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchImageRegistriesId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistriesAPI.PatchImageRegistriesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Registry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchImageRegistriesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersPatchImageRegistryRequest**](HandlersPatchImageRegistryRequest.md) | Credential / FK updates | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImageRegistries

> ModelsImageRegistry PostImageRegistries(ctx).Request(request).Execute()

Create a new image registry



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
	request := *openapiclient.NewHandlersCreateImageRegistryRequest("Name_example", "Type_example", "Url_example") // HandlersCreateImageRegistryRequest | Registry details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageRegistriesAPI.PostImageRegistries(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.PostImageRegistries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImageRegistries`: ModelsImageRegistry
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistriesAPI.PostImageRegistries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImageRegistriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersCreateImageRegistryRequest**](HandlersCreateImageRegistryRequest.md) | Registry details | 

### Return type

[**ModelsImageRegistry**](ModelsImageRegistry.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [EncryptionKeyAuth](../README.md#EncryptionKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImageRegistriesTest

> HandlersTestImageRegistryResponse PostImageRegistriesTest(ctx).Request(request).Execute()

Test an image registry connection



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
	request := *openapiclient.NewHandlersTestImageRegistryRequest("Type_example", "Url_example") // HandlersTestImageRegistryRequest | Registry connection details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageRegistriesAPI.PostImageRegistriesTest(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageRegistriesAPI.PostImageRegistriesTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImageRegistriesTest`: HandlersTestImageRegistryResponse
	fmt.Fprintf(os.Stdout, "Response from `ImageRegistriesAPI.PostImageRegistriesTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImageRegistriesTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersTestImageRegistryRequest**](HandlersTestImageRegistryRequest.md) | Registry connection details | 

### Return type

[**HandlersTestImageRegistryResponse**](HandlersTestImageRegistryResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

