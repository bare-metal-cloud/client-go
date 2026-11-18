# \PersonGrantsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgsOrgIdPersonGrantsId**](PersonGrantsAPI.md#DeleteOrgsOrgIdPersonGrantsId) | **Delete** /orgs/{org_id}/person-grants/{id} | Delete a person grant
[**GetOrgsOrgIdPersonGrants**](PersonGrantsAPI.md#GetOrgsOrgIdPersonGrants) | **Get** /orgs/{org_id}/person-grants | List the org&#39;s person grants
[**PostOrgsOrgIdPersonGrants**](PersonGrantsAPI.md#PostOrgsOrgIdPersonGrants) | **Post** /orgs/{org_id}/person-grants | Grant a person access to one database
[**PostOrgsOrgIdPersonGrantsIdRevoke**](PersonGrantsAPI.md#PostOrgsOrgIdPersonGrantsIdRevoke) | **Post** /orgs/{org_id}/person-grants/{id}/revoke | Revoke a person grant



## DeleteOrgsOrgIdPersonGrantsId

> DeleteOrgsOrgIdPersonGrantsId(ctx, orgId, id).Force(force).Execute()

Delete a person grant



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
	orgId := "orgId_example" // string | Organization ID
	id := "id_example" // string | Person Grant ID
	force := true // bool | Force the delete past the Used-by guard (admin/owner only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonGrantsAPI.DeleteOrgsOrgIdPersonGrantsId(context.Background(), orgId, id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonGrantsAPI.DeleteOrgsOrgIdPersonGrantsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**id** | **string** | Person Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdPersonGrantsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | Force the delete past the Used-by guard (admin/owner only) | 

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


## GetOrgsOrgIdPersonGrants

> []ModelsPersonGrant GetOrgsOrgIdPersonGrants(ctx, orgId).Status(status).Execute()

List the org's person grants



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
	orgId := "orgId_example" // string | Organization ID
	status := "status_example" // string | Filter: active, suspended, revoked, or all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonGrantsAPI.GetOrgsOrgIdPersonGrants(context.Background(), orgId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonGrantsAPI.GetOrgsOrgIdPersonGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdPersonGrants`: []ModelsPersonGrant
	fmt.Fprintf(os.Stdout, "Response from `PersonGrantsAPI.GetOrgsOrgIdPersonGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdPersonGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter: active, suspended, revoked, or all | 

### Return type

[**[]ModelsPersonGrant**](ModelsPersonGrant.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdPersonGrants

> ModelsPersonGrant PostOrgsOrgIdPersonGrants(ctx, orgId).Request(request).Execute()

Grant a person access to one database



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
	orgId := "orgId_example" // string | Organization ID
	request := *openapiclient.NewHandlersCreatePersonGrantRequest() // HandlersCreatePersonGrantRequest | The grant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonGrantsAPI.PostOrgsOrgIdPersonGrants(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonGrantsAPI.PostOrgsOrgIdPersonGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdPersonGrants`: ModelsPersonGrant
	fmt.Fprintf(os.Stdout, "Response from `PersonGrantsAPI.PostOrgsOrgIdPersonGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdPersonGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersCreatePersonGrantRequest**](HandlersCreatePersonGrantRequest.md) | The grant | 

### Return type

[**ModelsPersonGrant**](ModelsPersonGrant.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdPersonGrantsIdRevoke

> ModelsPersonGrant PostOrgsOrgIdPersonGrantsIdRevoke(ctx, orgId, id).Execute()

Revoke a person grant



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
	orgId := "orgId_example" // string | Organization ID
	id := "id_example" // string | Person Grant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonGrantsAPI.PostOrgsOrgIdPersonGrantsIdRevoke(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonGrantsAPI.PostOrgsOrgIdPersonGrantsIdRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdPersonGrantsIdRevoke`: ModelsPersonGrant
	fmt.Fprintf(os.Stdout, "Response from `PersonGrantsAPI.PostOrgsOrgIdPersonGrantsIdRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organization ID | 
**id** | **string** | Person Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdPersonGrantsIdRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsPersonGrant**](ModelsPersonGrant.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

