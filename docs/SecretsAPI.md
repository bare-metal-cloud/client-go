# \SecretsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecretsTypeIdUsage**](SecretsAPI.md#GetSecretsTypeIdUsage) | **Get** /secrets/{type}/{id}/usage | List active usages for a secret
[**GetSecretsTypeIdVersions**](SecretsAPI.md#GetSecretsTypeIdVersions) | **Get** /secrets/{type}/{id}/versions | List secret ciphertext versions
[**PostSecretsTypeIdVersionsVersionNumberRollback**](SecretsAPI.md#PostSecretsTypeIdVersionsVersionNumberRollback) | **Post** /secrets/{type}/{id}/versions/{version_number}/rollback | Roll back secret to a prior version (copy-forward)



## GetSecretsTypeIdUsage

> map[string]interface{} GetSecretsTypeIdUsage(ctx, type_, id).Execute()

List active usages for a secret



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
	type_ := "type__example" // string | Secret type (ssh_key, cloud_credential, s3_connection, image_registry, ssh_connection)
	id := int32(56) // int32 | Numeric secret id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetSecretsTypeIdUsage(context.Background(), type_, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecretsTypeIdUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretsTypeIdUsage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecretsTypeIdUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Secret type (ssh_key, cloud_credential, s3_connection, image_registry, ssh_connection) | 
**id** | **int32** | Numeric secret id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsTypeIdUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetSecretsTypeIdVersions

> map[string]interface{} GetSecretsTypeIdVersions(ctx, type_, id).FieldName(fieldName).Cursor(cursor).Limit(limit).Execute()

List secret ciphertext versions



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
	type_ := "type__example" // string | Secret type (ssh_key, cloud_credential, s3_connection, image_registry, ssh_connection)
	id := int32(56) // int32 | Numeric secret id
	fieldName := "fieldName_example" // string | Encrypted field name (defaults per secret type) (optional)
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetSecretsTypeIdVersions(context.Background(), type_, id).FieldName(fieldName).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecretsTypeIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretsTypeIdVersions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecretsTypeIdVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Secret type (ssh_key, cloud_credential, s3_connection, image_registry, ssh_connection) | 
**id** | **int32** | Numeric secret id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsTypeIdVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fieldName** | **string** | Encrypted field name (defaults per secret type) | 
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 

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


## PostSecretsTypeIdVersionsVersionNumberRollback

> map[string]interface{} PostSecretsTypeIdVersionsVersionNumberRollback(ctx, type_, id, versionNumber).Execute()

Roll back secret to a prior version (copy-forward)



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
	type_ := "type__example" // string | Secret type (ssh_key, cloud_credential, s3_connection, image_registry, ssh_connection)
	id := int32(56) // int32 | Numeric secret id
	versionNumber := int32(56) // int32 | Target version_number to restore (copy-forward)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.PostSecretsTypeIdVersionsVersionNumberRollback(context.Background(), type_, id, versionNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.PostSecretsTypeIdVersionsVersionNumberRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSecretsTypeIdVersionsVersionNumberRollback`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.PostSecretsTypeIdVersionsVersionNumberRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Secret type (ssh_key, cloud_credential, s3_connection, image_registry, ssh_connection) | 
**id** | **int32** | Numeric secret id | 
**versionNumber** | **int32** | Target version_number to restore (copy-forward) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecretsTypeIdVersionsVersionNumberRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

