# \NotificationDigestSettingsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationsOrgIdNotificationDigestSettings**](NotificationDigestSettingsAPI.md#GetOrganizationsOrgIdNotificationDigestSettings) | **Get** /organizations/{org_id}/notification-digest-settings | Read the org&#39;s notification digest settings
[**PutOrganizationsOrgIdNotificationDigestSettings**](NotificationDigestSettingsAPI.md#PutOrganizationsOrgIdNotificationDigestSettings) | **Put** /organizations/{org_id}/notification-digest-settings | Update the org&#39;s notification digest settings



## GetOrganizationsOrgIdNotificationDigestSettings

> ModelsOrgDigestSetting GetOrganizationsOrgIdNotificationDigestSettings(ctx, orgId).Execute()

Read the org's notification digest settings



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
	resp, r, err := apiClient.NotificationDigestSettingsAPI.GetOrganizationsOrgIdNotificationDigestSettings(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationDigestSettingsAPI.GetOrganizationsOrgIdNotificationDigestSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsOrgIdNotificationDigestSettings`: ModelsOrgDigestSetting
	fmt.Fprintf(os.Stdout, "Response from `NotificationDigestSettingsAPI.GetOrganizationsOrgIdNotificationDigestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsOrgIdNotificationDigestSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsOrgDigestSetting**](ModelsOrgDigestSetting.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrganizationsOrgIdNotificationDigestSettings

> ModelsOrgDigestSetting PutOrganizationsOrgIdNotificationDigestSettings(ctx, orgId).Body(body).Execute()

Update the org's notification digest settings



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
	body := map[string]interface{}{ ... } // map[string]interface{} | Settings payload: cadence=weekly|daily, include_empty=bool (all fields optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationDigestSettingsAPI.PutOrganizationsOrgIdNotificationDigestSettings(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationDigestSettingsAPI.PutOrganizationsOrgIdNotificationDigestSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrganizationsOrgIdNotificationDigestSettings`: ModelsOrgDigestSetting
	fmt.Fprintf(os.Stdout, "Response from `NotificationDigestSettingsAPI.PutOrganizationsOrgIdNotificationDigestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrganizationsOrgIdNotificationDigestSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Settings payload: cadence&#x3D;weekly|daily, include_empty&#x3D;bool (all fields optional) | 

### Return type

[**ModelsOrgDigestSetting**](ModelsOrgDigestSetting.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

