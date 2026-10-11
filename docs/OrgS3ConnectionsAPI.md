# \OrgS3ConnectionsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgsOrgIdS3ConnectionsId**](OrgS3ConnectionsAPI.md#DeleteOrgsOrgIdS3ConnectionsId) | **Delete** /orgs/{org_id}/s3-connections/{id} | Delete an org-scoped S3 connection
[**GetOrgsOrgIdS3Connections**](OrgS3ConnectionsAPI.md#GetOrgsOrgIdS3Connections) | **Get** /orgs/{org_id}/s3-connections | List org-scoped S3 connections
[**PostOrgsOrgIdS3Connections**](OrgS3ConnectionsAPI.md#PostOrgsOrgIdS3Connections) | **Post** /orgs/{org_id}/s3-connections | Create an org-scoped S3 connection
[**PostOrgsOrgIdS3ConnectionsIdRotateCredentials**](OrgS3ConnectionsAPI.md#PostOrgsOrgIdS3ConnectionsIdRotateCredentials) | **Post** /orgs/{org_id}/s3-connections/{id}/rotate-credentials | Rotate credentials for an org-scoped S3 connection
[**PutOrgsOrgIdS3ConnectionsId**](OrgS3ConnectionsAPI.md#PutOrgsOrgIdS3ConnectionsId) | **Put** /orgs/{org_id}/s3-connections/{id} | Update an org-scoped S3 connection (non-credential fields only)



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
	request := *openapiclient.NewHandlersCreateOrgS3ConnectionRequest("AccessKeyId_example", "Endpoint_example", "Name_example", "Provider_example", "SecretAccessKey_example") // HandlersCreateOrgS3ConnectionRequest | S3 connection details

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
	request := *openapiclient.NewHandlersRotateOrgS3ConnectionCredentialsRequest("AccessKeyId_example", "SecretAccessKey_example") // HandlersRotateOrgS3ConnectionCredentialsRequest | New credentials

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

