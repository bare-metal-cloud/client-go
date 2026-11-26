# \OrganizationsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutApiV1OrganizationsOrgIdUsageMeasurement**](OrganizationsAPI.md#PutApiV1OrganizationsOrgIdUsageMeasurement) | **Put** /api/v1/organizations/{org_id}/usage-measurement | Set the organization&#39;s usage-measurement toggle



## PutApiV1OrganizationsOrgIdUsageMeasurement

> map[string]interface{} PutApiV1OrganizationsOrgIdUsageMeasurement(ctx, orgId).Body(body).Execute()

Set the organization's usage-measurement toggle



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
	body := *openapiclient.NewHandlersUpdateUsageMeasurementRequest() // HandlersUpdateUsageMeasurementRequest | Toggle value

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PutApiV1OrganizationsOrgIdUsageMeasurement(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PutApiV1OrganizationsOrgIdUsageMeasurement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApiV1OrganizationsOrgIdUsageMeasurement`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PutApiV1OrganizationsOrgIdUsageMeasurement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApiV1OrganizationsOrgIdUsageMeasurementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersUpdateUsageMeasurementRequest**](HandlersUpdateUsageMeasurementRequest.md) | Toggle value | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

