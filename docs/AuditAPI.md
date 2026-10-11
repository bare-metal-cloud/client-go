# \AuditAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdAuditEvents**](AuditAPI.md#GetOrgsOrgIdAuditEvents) | **Get** /orgs/{org_id}/audit-events | List org audit events



## GetOrgsOrgIdAuditEvents

> map[string]interface{} GetOrgsOrgIdAuditEvents(ctx, orgId).Cursor(cursor).Limit(limit).ResourceType(resourceType).ActorUserId(actorUserId).From(from).To(to).Execute()

List org audit events



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
	cursor := "cursor_example" // string | Opaque pagination cursor (optional)
	limit := int32(56) // int32 | Page size (default 50, max 200) (optional)
	resourceType := "resourceType_example" // string | Filter by resource_type (optional)
	actorUserId := int32(56) // int32 | Filter by actor user id (optional)
	from := "from_example" // string | Inclusive lower bound (RFC3339) (optional)
	to := "to_example" // string | Inclusive upper bound (RFC3339) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditAPI.GetOrgsOrgIdAuditEvents(context.Background(), orgId).Cursor(cursor).Limit(limit).ResourceType(resourceType).ActorUserId(actorUserId).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.GetOrgsOrgIdAuditEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdAuditEvents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuditAPI.GetOrgsOrgIdAuditEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdAuditEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | Opaque pagination cursor | 
 **limit** | **int32** | Page size (default 50, max 200) | 
 **resourceType** | **string** | Filter by resource_type | 
 **actorUserId** | **int32** | Filter by actor user id | 
 **from** | **string** | Inclusive lower bound (RFC3339) | 
 **to** | **string** | Inclusive upper bound (RFC3339) | 

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

