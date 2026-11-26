# \DatabasesAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServicesDatabasesId**](DatabasesAPI.md#DeleteServicesDatabasesId) | **Delete** /services/databases/{id} | Delete a database
[**DeleteServicesDatabasesIdAdminGrantsGrantId**](DatabasesAPI.md#DeleteServicesDatabasesIdAdminGrantsGrantId) | **Delete** /services/databases/{id}/admin/grants/{grantId} | Revoke a member&#39;s grant
[**DeleteServicesDatabasesIdAdminUsersName**](DatabasesAPI.md#DeleteServicesDatabasesIdAdminUsersName) | **Delete** /services/databases/{id}/admin/users/{name} | Delete a user
[**GetServicesDatabases**](DatabasesAPI.md#GetServicesDatabases) | **Get** /services/databases | List all databases
[**GetServicesDatabasesConfigSchema**](DatabasesAPI.md#GetServicesDatabasesConfigSchema) | **Get** /services/databases/config/schema | Get database configuration schema
[**GetServicesDatabasesId**](DatabasesAPI.md#GetServicesDatabasesId) | **Get** /services/databases/{id} | Get database details
[**GetServicesDatabasesIdAdminChanges**](DatabasesAPI.md#GetServicesDatabasesIdAdminChanges) | **Get** /services/databases/{id}/admin/changes | List the database service&#39;s admin change log
[**GetServicesDatabasesIdAdminDatabases**](DatabasesAPI.md#GetServicesDatabasesIdAdminDatabases) | **Get** /services/databases/{id}/admin/databases | List the databases on a database service
[**GetServicesDatabasesIdAdminGrants**](DatabasesAPI.md#GetServicesDatabasesIdAdminGrants) | **Get** /services/databases/{id}/admin/grants | See who can do what on a database service
[**GetServicesDatabasesIdAdminParameters**](DatabasesAPI.md#GetServicesDatabasesIdAdminParameters) | **Get** /services/databases/{id}/admin/parameters | List the engine parameters on a database service
[**GetServicesDatabasesIdAdminUsers**](DatabasesAPI.md#GetServicesDatabasesIdAdminUsers) | **Get** /services/databases/{id}/admin/users | List the users on a database service
[**GetServicesDatabasesIdConnection**](DatabasesAPI.md#GetServicesDatabasesIdConnection) | **Get** /services/databases/{id}/connection | Get database connection info
[**GetServicesDatabasesIdExportDumpJobDownload**](DatabasesAPI.md#GetServicesDatabasesIdExportDumpJobDownload) | **Get** /services/databases/{id}/export-dump/{job}/download | Download one completed database dump export
[**GetServicesDatabasesIdExtensions**](DatabasesAPI.md#GetServicesDatabasesIdExtensions) | **Get** /services/databases/{id}/extensions | List the Postgres extensions installed on a database
[**GetServicesDatabasesIdSchema**](DatabasesAPI.md#GetServicesDatabasesIdSchema) | **Get** /services/databases/{id}/schema | Browse a database&#39;s schema
[**GetServicesDatabasesIdSchemaSchemaTablesTableColumns**](DatabasesAPI.md#GetServicesDatabasesIdSchemaSchemaTablesTableColumns) | **Get** /services/databases/{id}/schema/{schema}/tables/{table}/columns | List one relation&#39;s columns
[**GetServicesDatabasesIdStatus**](DatabasesAPI.md#GetServicesDatabasesIdStatus) | **Get** /services/databases/{id}/status | Get database deployment status
[**PatchServicesDatabasesIdAdminGrantsGrantId**](DatabasesAPI.md#PatchServicesDatabasesIdAdminGrantsGrantId) | **Patch** /services/databases/{id}/admin/grants/{grantId} | Edit a member&#39;s grant
[**PatchServicesDatabasesIdAdminUsersName**](DatabasesAPI.md#PatchServicesDatabasesIdAdminUsersName) | **Patch** /services/databases/{id}/admin/users/{name} | Edit a user&#39;s role
[**PostServicesDatabases**](DatabasesAPI.md#PostServicesDatabases) | **Post** /services/databases | Create database configuration
[**PostServicesDatabasesConfigValidate**](DatabasesAPI.md#PostServicesDatabasesConfigValidate) | **Post** /services/databases/config/validate | Validate database configuration
[**PostServicesDatabasesIdAdminDatabases**](DatabasesAPI.md#PostServicesDatabasesIdAdminDatabases) | **Post** /services/databases/{id}/admin/databases | Create a database on a database service
[**PostServicesDatabasesIdAdminGrants**](DatabasesAPI.md#PostServicesDatabasesIdAdminGrants) | **Post** /services/databases/{id}/admin/grants | Grant a member access to a database service
[**PostServicesDatabasesIdAdminParametersApplyRestart**](DatabasesAPI.md#PostServicesDatabasesIdAdminParametersApplyRestart) | **Post** /services/databases/{id}/admin/parameters/apply-restart | Apply the staged parameter changes and restart
[**PostServicesDatabasesIdAdminUsers**](DatabasesAPI.md#PostServicesDatabasesIdAdminUsers) | **Post** /services/databases/{id}/admin/users | Create a user on a database service
[**PostServicesDatabasesIdAdminUsersNameRevoke**](DatabasesAPI.md#PostServicesDatabasesIdAdminUsersNameRevoke) | **Post** /services/databases/{id}/admin/users/{name}/revoke | Revoke a user&#39;s login
[**PostServicesDatabasesIdConnectionResetPassword**](DatabasesAPI.md#PostServicesDatabasesIdConnectionResetPassword) | **Post** /services/databases/{id}/connection/reset-password | Reset database password
[**PostServicesDatabasesIdDeploy**](DatabasesAPI.md#PostServicesDatabasesIdDeploy) | **Post** /services/databases/{id}/deploy | Deploy a database (first deploy or re-deploy after error/stop)
[**PostServicesDatabasesIdExportDump**](DatabasesAPI.md#PostServicesDatabasesIdExportDump) | **Post** /services/databases/{id}/export-dump | Dump (export) ONE managed database out for download
[**PostServicesDatabasesIdQuery**](DatabasesAPI.md#PostServicesDatabasesIdQuery) | **Post** /services/databases/{id}/query | Run one read-only query on a database
[**PostServicesDatabasesIdRecoverLane**](DatabasesAPI.md#PostServicesDatabasesIdRecoverLane) | **Post** /services/databases/{id}/recover-lane | Recover a stranded backup lane (admin)
[**PostServicesDatabasesIdRestart**](DatabasesAPI.md#PostServicesDatabasesIdRestart) | **Post** /services/databases/{id}/restart | Restart a database (graceful rolling restart)
[**PostServicesDatabasesIdRestore**](DatabasesAPI.md#PostServicesDatabasesIdRestore) | **Post** /services/databases/{id}/restore | Restore a database from a backup archive (deploykit engine)
[**PostServicesDatabasesIdRestoreDump**](DatabasesAPI.md#PostServicesDatabasesIdRestoreDump) | **Post** /services/databases/{id}/restore-dump | Restore ONE managed database from a raw dump file
[**PostServicesDatabasesIdScaleResources**](DatabasesAPI.md#PostServicesDatabasesIdScaleResources) | **Post** /services/databases/{id}/scale-resources | Scale database CPU/memory resources
[**PostServicesDatabasesIdScaleStorage**](DatabasesAPI.md#PostServicesDatabasesIdScaleStorage) | **Post** /services/databases/{id}/scale-storage | Scale database storage (increase only)
[**PostServicesDatabasesIdStart**](DatabasesAPI.md#PostServicesDatabasesIdStart) | **Post** /services/databases/{id}/start | Start a stopped database (restore replicas)
[**PostServicesDatabasesIdStop**](DatabasesAPI.md#PostServicesDatabasesIdStop) | **Post** /services/databases/{id}/stop | Stop a database (scale to 0 replicas)
[**PostServicesDatabasesIdUpgrade**](DatabasesAPI.md#PostServicesDatabasesIdUpgrade) | **Post** /services/databases/{id}/upgrade | Upgrade a database (deploykit engine)
[**PutServicesDatabasesIdAdminParametersName**](DatabasesAPI.md#PutServicesDatabasesIdAdminParametersName) | **Put** /services/databases/{id}/admin/parameters/{name} | Edit one engine parameter on a database service



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


## DeleteServicesDatabasesIdAdminGrantsGrantId

> map[string]interface{} DeleteServicesDatabasesIdAdminGrantsGrantId(ctx, id, grantId).Execute()

Revoke a member's grant



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
	grantId := int32(56) // int32 | Grant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.DeleteServicesDatabasesIdAdminGrantsGrantId(context.Background(), id, grantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteServicesDatabasesIdAdminGrantsGrantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServicesDatabasesIdAdminGrantsGrantId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteServicesDatabasesIdAdminGrantsGrantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 
**grantId** | **int32** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServicesDatabasesIdAdminGrantsGrantIdRequest struct via the builder pattern


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


## DeleteServicesDatabasesIdAdminUsersName

> map[string]interface{} DeleteServicesDatabasesIdAdminUsersName(ctx, id, name).Execute()

Delete a user



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
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.DeleteServicesDatabasesIdAdminUsersName(context.Background(), id, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteServicesDatabasesIdAdminUsersName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServicesDatabasesIdAdminUsersName`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteServicesDatabasesIdAdminUsersName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServicesDatabasesIdAdminUsersNameRequest struct via the builder pattern


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
	type_ := "type__example" // string | Database type (postgresql, mariadb, mysql)
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
 **type_** | **string** | Database type (postgresql, mariadb, mysql) | 
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


## GetServicesDatabasesIdAdminChanges

> map[string]interface{} GetServicesDatabasesIdAdminChanges(ctx, id).Execute()

List the database service's admin change log



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
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdAdminChanges(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdAdminChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdAdminChanges`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdAdminChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdAdminChangesRequest struct via the builder pattern


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


## GetServicesDatabasesIdAdminDatabases

> map[string]interface{} GetServicesDatabasesIdAdminDatabases(ctx, id).Execute()

List the databases on a database service



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
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdAdminDatabases(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdAdminDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdAdminDatabases`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdAdminDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdAdminDatabasesRequest struct via the builder pattern


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


## GetServicesDatabasesIdAdminGrants

> map[string]interface{} GetServicesDatabasesIdAdminGrants(ctx, id).Execute()

See who can do what on a database service



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
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdAdminGrants(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdAdminGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdAdminGrants`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdAdminGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdAdminGrantsRequest struct via the builder pattern


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


## GetServicesDatabasesIdAdminParameters

> map[string]interface{} GetServicesDatabasesIdAdminParameters(ctx, id).Execute()

List the engine parameters on a database service



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
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdAdminParameters(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdAdminParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdAdminParameters`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdAdminParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdAdminParametersRequest struct via the builder pattern


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


## GetServicesDatabasesIdAdminUsers

> map[string]interface{} GetServicesDatabasesIdAdminUsers(ctx, id).Execute()

List the users on a database service



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
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdAdminUsers(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdAdminUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdAdminUsers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdAdminUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdAdminUsersRequest struct via the builder pattern


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


## GetServicesDatabasesIdExportDumpJobDownload

> map[string]interface{} GetServicesDatabasesIdExportDumpJobDownload(ctx, id, job).Execute()

Download one completed database dump export



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
	job := "job_example" // string | The export job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdExportDumpJobDownload(context.Background(), id, job).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdExportDumpJobDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdExportDumpJobDownload`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdExportDumpJobDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 
**job** | **string** | The export job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdExportDumpJobDownloadRequest struct via the builder pattern


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


## GetServicesDatabasesIdExtensions

> HandlersDatabaseExtensionsResponse GetServicesDatabasesIdExtensions(ctx, id).Execute()

List the Postgres extensions installed on a database



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
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdExtensions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdExtensions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdExtensions`: HandlersDatabaseExtensionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdExtensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersDatabaseExtensionsResponse**](HandlersDatabaseExtensionsResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesDatabasesIdSchema

> map[string]interface{} GetServicesDatabasesIdSchema(ctx, id).Execute()

Browse a database's schema



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
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdSchema(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdSchemaRequest struct via the builder pattern


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


## GetServicesDatabasesIdSchemaSchemaTablesTableColumns

> map[string]interface{} GetServicesDatabasesIdSchemaSchemaTablesTableColumns(ctx, id, schema, table).Execute()

List one relation's columns



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
	schema := "schema_example" // string | Schema name
	table := "table_example" // string | Table name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdSchemaSchemaTablesTableColumns(context.Background(), id, schema, table).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdSchemaSchemaTablesTableColumns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdSchemaSchemaTablesTableColumns`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdSchemaSchemaTablesTableColumns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 
**schema** | **string** | Schema name | 
**table** | **string** | Table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdSchemaSchemaTablesTableColumnsRequest struct via the builder pattern


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


## GetServicesDatabasesIdStatus

> HandlersDatabaseStatusResponse GetServicesDatabasesIdStatus(ctx, id).Execute()

Get database deployment status



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
	resp, r, err := apiClient.DatabasesAPI.GetServicesDatabasesIdStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetServicesDatabasesIdStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicesDatabasesIdStatus`: HandlersDatabaseStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetServicesDatabasesIdStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesDatabasesIdStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersDatabaseStatusResponse**](HandlersDatabaseStatusResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchServicesDatabasesIdAdminGrantsGrantId

> map[string]interface{} PatchServicesDatabasesIdAdminGrantsGrantId(ctx, id, grantId).Request(request).Execute()

Edit a member's grant



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
	grantId := int32(56) // int32 | Grant ID
	request := *openapiclient.NewHandlersEditAdminGrantRequest() // HandlersEditAdminGrantRequest | The member id and the new grant axes

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PatchServicesDatabasesIdAdminGrantsGrantId(context.Background(), id, grantId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PatchServicesDatabasesIdAdminGrantsGrantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServicesDatabasesIdAdminGrantsGrantId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PatchServicesDatabasesIdAdminGrantsGrantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 
**grantId** | **int32** | Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServicesDatabasesIdAdminGrantsGrantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**HandlersEditAdminGrantRequest**](HandlersEditAdminGrantRequest.md) | The member id and the new grant axes | 

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


## PatchServicesDatabasesIdAdminUsersName

> map[string]interface{} PatchServicesDatabasesIdAdminUsersName(ctx, id, name).Request(request).Execute()

Edit a user's role



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
	name := "name_example" // string | User name
	request := *openapiclient.NewHandlersEditAdminUserRoleRequest() // HandlersEditAdminUserRoleRequest | The new role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PatchServicesDatabasesIdAdminUsersName(context.Background(), id, name).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PatchServicesDatabasesIdAdminUsersName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServicesDatabasesIdAdminUsersName`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PatchServicesDatabasesIdAdminUsersName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServicesDatabasesIdAdminUsersNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**HandlersEditAdminUserRoleRequest**](HandlersEditAdminUserRoleRequest.md) | The new role | 

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


## PostServicesDatabasesIdAdminDatabases

> map[string]interface{} PostServicesDatabasesIdAdminDatabases(ctx, id).Request(request).Execute()

Create a database on a database service



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
	request := *openapiclient.NewHandlersCreateAdminDatabaseRequest() // HandlersCreateAdminDatabaseRequest | The database name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdAdminDatabases(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdAdminDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdAdminDatabases`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdAdminDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdAdminDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersCreateAdminDatabaseRequest**](HandlersCreateAdminDatabaseRequest.md) | The database name | 

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


## PostServicesDatabasesIdAdminGrants

> map[string]interface{} PostServicesDatabasesIdAdminGrants(ctx, id).Request(request).Execute()

Grant a member access to a database service



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
	request := *openapiclient.NewHandlersCreateAdminGrantRequest() // HandlersCreateAdminGrantRequest | The member id and the grant axes

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdAdminGrants(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdAdminGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdAdminGrants`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdAdminGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdAdminGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersCreateAdminGrantRequest**](HandlersCreateAdminGrantRequest.md) | The member id and the grant axes | 

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


## PostServicesDatabasesIdAdminParametersApplyRestart

> map[string]interface{} PostServicesDatabasesIdAdminParametersApplyRestart(ctx, id).Execute()

Apply the staged parameter changes and restart



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
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdAdminParametersApplyRestart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdAdminParametersApplyRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdAdminParametersApplyRestart`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdAdminParametersApplyRestart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdAdminParametersApplyRestartRequest struct via the builder pattern


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


## PostServicesDatabasesIdAdminUsers

> map[string]interface{} PostServicesDatabasesIdAdminUsers(ctx, id).Request(request).Execute()

Create a user on a database service



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
	request := *openapiclient.NewHandlersCreateAdminUserRequest() // HandlersCreateAdminUserRequest | The user name, target database and role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdAdminUsers(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdAdminUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdAdminUsers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdAdminUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdAdminUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersCreateAdminUserRequest**](HandlersCreateAdminUserRequest.md) | The user name, target database and role | 

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


## PostServicesDatabasesIdAdminUsersNameRevoke

> map[string]interface{} PostServicesDatabasesIdAdminUsersNameRevoke(ctx, id, name).Execute()

Revoke a user's login



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
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdAdminUsersNameRevoke(context.Background(), id, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdAdminUsersNameRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdAdminUsersNameRevoke`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdAdminUsersNameRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdAdminUsersNameRevokeRequest struct via the builder pattern


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


## PostServicesDatabasesIdDeploy

> map[string]interface{} PostServicesDatabasesIdDeploy(ctx, id).Execute()

Deploy a database (first deploy or re-deploy after error/stop)

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
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdDeploy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdDeploy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdDeployRequest struct via the builder pattern


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


## PostServicesDatabasesIdExportDump

> map[string]interface{} PostServicesDatabasesIdExportDump(ctx, id).Body(body).Execute()

Dump (export) ONE managed database out for download



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
	body := *openapiclient.NewHandlersDumpExportRequest() // HandlersDumpExportRequest | destination_ref; object_key optional (defaults under managed-db-dumps/) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdExportDump(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdExportDump``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdExportDump`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdExportDump`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdExportDumpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersDumpExportRequest**](HandlersDumpExportRequest.md) | destination_ref; object_key optional (defaults under managed-db-dumps/) | 

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


## PostServicesDatabasesIdQuery

> DatabasequeryQueryResponse PostServicesDatabasesIdQuery(ctx, id).Execute()

Run one read-only query on a database



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
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdQuery(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdQuery`: DatabasequeryQueryResponse
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabasequeryQueryResponse**](DatabasequeryQueryResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServicesDatabasesIdRecoverLane

> map[string]interface{} PostServicesDatabasesIdRecoverLane(ctx, id).Request(request).Execute()

Recover a stranded backup lane (admin)



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
	request := *openapiclient.NewHandlersRecoverLaneRequest() // HandlersRecoverLaneRequest | Explicit operator acknowledgment (confirm must be true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdRecoverLane(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdRecoverLane``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdRecoverLane`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdRecoverLane`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdRecoverLaneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersRecoverLaneRequest**](HandlersRecoverLaneRequest.md) | Explicit operator acknowledgment (confirm must be true) | 

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


## PostServicesDatabasesIdRestore

> map[string]interface{} PostServicesDatabasesIdRestore(ctx, id).Request(request).Execute()

Restore a database from a backup archive (deploykit engine)



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
	request := *openapiclient.NewHandlersRestoreDatabaseRequest() // HandlersRestoreDatabaseRequest | Archive handle (destination_ref, object_key, digest) plus the explicit fresh confirm

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdRestore(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdRestore`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersRestoreDatabaseRequest**](HandlersRestoreDatabaseRequest.md) | Archive handle (destination_ref, object_key, digest) plus the explicit fresh confirm | 

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


## PostServicesDatabasesIdRestoreDump

> map[string]interface{} PostServicesDatabasesIdRestoreDump(ctx, id).Body(body).Execute()

Restore ONE managed database from a raw dump file



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
	body := *openapiclient.NewHandlersDumpRestoreRequest() // HandlersDumpRestoreRequest | destination_ref, object_key, optional digest; confirm on the re-post

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdRestoreDump(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdRestoreDump``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdRestoreDump`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdRestoreDump`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdRestoreDumpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersDumpRestoreRequest**](HandlersDumpRestoreRequest.md) | destination_ref, object_key, optional digest; confirm on the re-post | 

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


## PostServicesDatabasesIdUpgrade

> map[string]interface{} PostServicesDatabasesIdUpgrade(ctx, id).Request(request).Execute()

Upgrade a database (deploykit engine)



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
	request := *openapiclient.NewHandlersUpgradeDatabaseRequest() // HandlersUpgradeDatabaseRequest | Target version (empty upgrades to the catalog's latest) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostServicesDatabasesIdUpgrade(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostServicesDatabasesIdUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServicesDatabasesIdUpgrade`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostServicesDatabasesIdUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServicesDatabasesIdUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersUpgradeDatabaseRequest**](HandlersUpgradeDatabaseRequest.md) | Target version (empty upgrades to the catalog&#39;s latest) | 

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


## PutServicesDatabasesIdAdminParametersName

> map[string]interface{} PutServicesDatabasesIdAdminParametersName(ctx, id, name).Request(request).Execute()

Edit one engine parameter on a database service



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
	name := "name_example" // string | Parameter name (lowercase GUC name, optionally dotted)
	request := *openapiclient.NewHandlersSetAdminParameterRequest() // HandlersSetAdminParameterRequest | The parameter value

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PutServicesDatabasesIdAdminParametersName(context.Background(), id, name).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PutServicesDatabasesIdAdminParametersName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServicesDatabasesIdAdminParametersName`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PutServicesDatabasesIdAdminParametersName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID | 
**name** | **string** | Parameter name (lowercase GUC name, optionally dotted) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServicesDatabasesIdAdminParametersNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**HandlersSetAdminParameterRequest**](HandlersSetAdminParameterRequest.md) | The parameter value | 

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

