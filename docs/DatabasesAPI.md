# \DatabasesAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServicesDatabasesId**](DatabasesAPI.md#DeleteServicesDatabasesId) | **Delete** /services/databases/{id} | Delete a database
[**GetServicesDatabases**](DatabasesAPI.md#GetServicesDatabases) | **Get** /services/databases | List all databases
[**GetServicesDatabasesConfigSchema**](DatabasesAPI.md#GetServicesDatabasesConfigSchema) | **Get** /services/databases/config/schema | Get database configuration schema
[**GetServicesDatabasesId**](DatabasesAPI.md#GetServicesDatabasesId) | **Get** /services/databases/{id} | Get database details
[**GetServicesDatabasesIdConnection**](DatabasesAPI.md#GetServicesDatabasesIdConnection) | **Get** /services/databases/{id}/connection | Get database connection info
[**PostServicesDatabases**](DatabasesAPI.md#PostServicesDatabases) | **Post** /services/databases | Create database configuration
[**PostServicesDatabasesConfigValidate**](DatabasesAPI.md#PostServicesDatabasesConfigValidate) | **Post** /services/databases/config/validate | Validate database configuration
[**PostServicesDatabasesIdConnectionResetPassword**](DatabasesAPI.md#PostServicesDatabasesIdConnectionResetPassword) | **Post** /services/databases/{id}/connection/reset-password | Reset database password
[**PostServicesDatabasesIdRestart**](DatabasesAPI.md#PostServicesDatabasesIdRestart) | **Post** /services/databases/{id}/restart | Restart a database (graceful rolling restart)
[**PostServicesDatabasesIdScaleResources**](DatabasesAPI.md#PostServicesDatabasesIdScaleResources) | **Post** /services/databases/{id}/scale-resources | Scale database CPU/memory resources
[**PostServicesDatabasesIdScaleStorage**](DatabasesAPI.md#PostServicesDatabasesIdScaleStorage) | **Post** /services/databases/{id}/scale-storage | Scale database storage (increase only)
[**PostServicesDatabasesIdStart**](DatabasesAPI.md#PostServicesDatabasesIdStart) | **Post** /services/databases/{id}/start | Start a stopped database (restore replicas)
[**PostServicesDatabasesIdStop**](DatabasesAPI.md#PostServicesDatabasesIdStop) | **Post** /services/databases/{id}/stop | Stop a database (scale to 0 replicas)



## DeleteServicesDatabasesId

> map[string]interface{} DeleteServicesDatabasesId(ctx, id).Request(request).Execute()

Delete a database

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
	id := "id_example" // string | Database ID
	request := *openapiclient.NewHandlersDeleteDatabaseRequest("ConfirmName_example") // HandlersDeleteDatabaseRequest | Confirmation with database name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.DeleteServicesDatabasesId(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteServicesDatabasesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServicesDatabasesId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteServicesDatabasesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServicesDatabasesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersDeleteDatabaseRequest**](HandlersDeleteDatabaseRequest.md) | Confirmation with database name | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesDatabases

> GetServicesDatabases(ctx).Execute()

List all databases

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
	r, err := apiClient.DatabasesAPI.GetServicesDatabases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesRequest struct via the builder pattern


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


## GetServicesDatabasesConfigSchema

> GetServicesDatabasesConfigSchema(ctx).Type_(type_).ClusterId(clusterId).Execute()

Get database configuration schema

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
	type_ := "type__example" // string | Database type (postgresql, mongodb, mariadb)
	clusterId := "clusterId_example" // string | Cluster ID for storage classes

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatabasesAPI.GetServicesDatabasesConfigSchema(context.Background()).Type_(type_).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesConfigSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesConfigSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Database type (postgresql, mongodb, mariadb) | 
 **clusterId** | **string** | Cluster ID for storage classes | 

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


## GetServicesDatabasesId

> GetServicesDatabasesId(ctx, id).Execute()

Get database details

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
	id := "id_example" // string | Database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DatabasesAPI.GetServicesDatabasesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetServicesDatabasesIdConnection

> ServicesDatabaseConnectionInfo GetServicesDatabasesIdConnection(ctx, id).Execute()

Get database connection info



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
	id := int32(56) // int32 | Database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdConnection`: ServicesDatabaseConnectionInfo
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServicesDatabaseConnectionInfo**](ServicesDatabaseConnectionInfo.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServicesDatabases

> ModelsDatabase PostServicesDatabases(ctx).Request(request).Execute()

Create database configuration

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
	request := *openapiclient.NewHandlersValidateDatabaseConfigRequest() // HandlersValidateDatabaseConfigRequest | Database configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabases(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabases`: ModelsDatabase
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersValidateDatabaseConfigRequest**](HandlersValidateDatabaseConfigRequest.md) | Database configuration | 

### Return type

[**ModelsDatabase**](ModelsDatabase.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServicesDatabasesConfigValidate

> map[string]interface{} PostServicesDatabasesConfigValidate(ctx).Request(request).Execute()

Validate database configuration

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
	request := *openapiclient.NewHandlersValidateDatabaseConfigRequest() // HandlersValidateDatabaseConfigRequest | Database configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesConfigValidate(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesConfigValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesConfigValidate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesConfigValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesConfigValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersValidateDatabaseConfigRequest**](HandlersValidateDatabaseConfigRequest.md) | Database configuration | 

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


## PostServicesDatabasesIdConnectionResetPassword

> map[string]interface{} PostServicesDatabasesIdConnectionResetPassword(ctx, id).Execute()

Reset database password



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
	id := int32(56) // int32 | Database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdConnectionResetPassword(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdConnectionResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdConnectionResetPassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdConnectionResetPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdConnectionResetPasswordRequest struct via the builder pattern


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


## PostServicesDatabasesIdRestart

> map[string]interface{} PostServicesDatabasesIdRestart(ctx, id).Execute()

Restart a database (graceful rolling restart)

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
	id := "id_example" // string | Database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdRestart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdRestart`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdRestart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServicesDatabasesIdScaleResources

> map[string]interface{} PostServicesDatabasesIdScaleResources(ctx, id).Request(request).Execute()

Scale database CPU/memory resources

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
	id := "id_example" // string | Database ID
	request := *openapiclient.NewHandlersScaleResourcesRequest() // HandlersScaleResourcesRequest | New resource limits

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdScaleResources(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdScaleResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdScaleResources`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdScaleResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdScaleResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersScaleResourcesRequest**](HandlersScaleResourcesRequest.md) | New resource limits | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServicesDatabasesIdScaleStorage

> map[string]interface{} PostServicesDatabasesIdScaleStorage(ctx, id).Request(request).Execute()

Scale database storage (increase only)

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
	id := "id_example" // string | Database ID
	request := *openapiclient.NewHandlersScaleStorageRequest("NewSize_example") // HandlersScaleStorageRequest | New storage size

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdScaleStorage(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdScaleStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdScaleStorage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdScaleStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdScaleStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersScaleStorageRequest**](HandlersScaleStorageRequest.md) | New storage size | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServicesDatabasesIdStart

> map[string]interface{} PostServicesDatabasesIdStart(ctx, id).Execute()

Start a stopped database (restore replicas)

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
	id := "id_example" // string | Database ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdStart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdStart`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServicesDatabasesIdStop

> map[string]interface{} PostServicesDatabasesIdStop(ctx, id).Request(request).Execute()

Stop a database (scale to 0 replicas)

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
	id := "id_example" // string | Database ID
	request := *openapiclient.NewHandlersStopDatabaseRequest("ConfirmName_example") // HandlersStopDatabaseRequest | Confirmation with database name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdStop(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdStop`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersStopDatabaseRequest**](HandlersStopDatabaseRequest.md) | Confirmation with database name | 

### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

