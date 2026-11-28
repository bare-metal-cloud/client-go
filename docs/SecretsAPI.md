# \SecretsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdS3ConnectionsIdVersions**](SecretsAPI.md#GetOrgsOrgIdS3ConnectionsIdVersions) | **Get** /orgs/{org_id}/s3-connections/{id}/versions | List an org-scoped S3 connection&#39;s credential versions
[**GetSecrets**](SecretsAPI.md#GetSecrets) | **Get** /secrets | List all secrets in the unified secrets hub
[**GetSecretsTypeId**](SecretsAPI.md#GetSecretsTypeId) | **Get** /secrets/{type}/{id} | Get one keyring secret&#39;s detail bundle
[**GetSecretsTypeIdUsage**](SecretsAPI.md#GetSecretsTypeIdUsage) | **Get** /secrets/{type}/{id}/usage | List active usages for a secret
[**GetSecretsTypeIdValue**](SecretsAPI.md#GetSecretsTypeIdValue) | **Get** /secrets/{type}/{id}/value | Reveal one keyring secret&#39;s value material (zero-peek exception)
[**GetSecretsTypeIdVersions**](SecretsAPI.md#GetSecretsTypeIdVersions) | **Get** /secrets/{type}/{id}/versions | List secret ciphertext versions
[**PostSecretsTypeIdTestConnection**](SecretsAPI.md#PostSecretsTypeIdTestConnection) | **Post** /secrets/{type}/{id}/test-connection | Test a secret&#39;s connection
[**PostSecretsTypeIdVersionsVersionNumberRollback**](SecretsAPI.md#PostSecretsTypeIdVersionsVersionNumberRollback) | **Post** /secrets/{type}/{id}/versions/{version_number}/rollback | Roll back secret to a prior version (copy-forward)



## GetOrgsOrgIdS3ConnectionsIdVersions

> map[string]interface{} GetOrgsOrgIdS3ConnectionsIdVersions(ctx, orgId, id).FieldName(fieldName).Cursor(cursor).Limit(limit).Execute()

List an org-scoped S3 connection's credential versions



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
	orgId := int32(56) // int32 | Numeric org id (the page org)
	id := int32(56) // int32 | Numeric S3 connection id
	fieldName := "fieldName_example" // string | Encrypted field name (defaults to credentials) (optional)
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetOrgsOrgIdS3ConnectionsIdVersions(context.Background(), orgId, id).FieldName(fieldName).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetOrgsOrgIdS3ConnectionsIdVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdS3ConnectionsIdVersions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetOrgsOrgIdS3ConnectionsIdVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Numeric org id (the page org) | 
**id** | **int32** | Numeric S3 connection id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdS3ConnectionsIdVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fieldName** | **string** | Encrypted field name (defaults to credentials) | 
 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 

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


## GetSecrets

> map[string]interface{} GetSecrets(ctx).XProjectId(xProjectId).ProjectId(projectId).Limit(limit).Cursor(cursor).Scope(scope).Execute()

List all secrets in the unified secrets hub



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
	xProjectId := "xProjectId_example" // string | Project selector: a project id (single-project mode) or the value all (every accessible project in the active org); defaults to the session's resolved project (optional)
	projectId := "projectId_example" // string | Query-string mirror of the X-Project-Id header (optional)
	limit := int32(56) // int32 | Page size (default 50, capped at 200) (optional)
	cursor := "cursor_example" // string | Opaque keyset cursor from a previous page (optional)
	scope := "scope_example" // string | Legacy scope filter; accepted but no longer applied (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetSecrets(context.Background()).XProjectId(xProjectId).ProjectId(projectId).Limit(limit).Cursor(cursor).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecrets`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xProjectId** | **string** | Project selector: a project id (single-project mode) or the value all (every accessible project in the active org); defaults to the session&#39;s resolved project | 
 **projectId** | **string** | Query-string mirror of the X-Project-Id header | 
 **limit** | **int32** | Page size (default 50, capped at 200) | 
 **cursor** | **string** | Opaque keyset cursor from a previous page | 
 **scope** | **string** | Legacy scope filter; accepted but no longer applied | 

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


## GetSecretsTypeId

> HandlersKeyringDetailResponse GetSecretsTypeId(ctx, type_, id).XProjectId(xProjectId).Execute()

Get one keyring secret's detail bundle



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
	xProjectId := "xProjectId_example" // string | Project selector (the hub list's scope) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetSecretsTypeId(context.Background(), type_, id).XProjectId(xProjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecretsTypeId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretsTypeId`: HandlersKeyringDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecretsTypeId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Secret type (ssh_key, cloud_credential, s3_connection, image_registry, ssh_connection) | 
**id** | **int32** | Numeric secret id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsTypeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xProjectId** | **string** | Project selector (the hub list&#39;s scope) | 

### Return type

[**HandlersKeyringDetailResponse**](HandlersKeyringDetailResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetSecretsTypeIdValue

> HandlersSecretValueEnvelope GetSecretsTypeIdValue(ctx, type_, id).Version(version).Execute()

Reveal one keyring secret's value material (zero-peek exception)



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
	version := int32(56) // int32 | A specific version_number to serve (default: the server-resolved current version) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.GetSecretsTypeIdValue(context.Background(), type_, id).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.GetSecretsTypeIdValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretsTypeIdValue`: HandlersSecretValueEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.GetSecretsTypeIdValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Secret type (ssh_key, cloud_credential, s3_connection, image_registry, ssh_connection) | 
**id** | **int32** | Numeric secret id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsTypeIdValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **int32** | A specific version_number to serve (default: the server-resolved current version) | 

### Return type

[**HandlersSecretValueEnvelope**](HandlersSecretValueEnvelope.md)

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


## PostSecretsTypeIdTestConnection

> HandlersSecretTestConnectionResponse PostSecretsTypeIdTestConnection(ctx, type_, id).Execute()

Test a secret's connection



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
	type_ := "type__example" // string | Secret type
	id := int32(56) // int32 | Secret ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.PostSecretsTypeIdTestConnection(context.Background(), type_, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.PostSecretsTypeIdTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSecretsTypeIdTestConnection`: HandlersSecretTestConnectionResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.PostSecretsTypeIdTestConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Secret type | 
**id** | **int32** | Secret ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecretsTypeIdTestConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersSecretTestConnectionResponse**](HandlersSecretTestConnectionResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

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

