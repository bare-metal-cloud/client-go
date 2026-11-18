# \NotificationChannelsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrganizationsOrgIdNotificationChannelsId**](NotificationChannelsAPI.md#DeleteOrganizationsOrgIdNotificationChannelsId) | **Delete** /organizations/{org_id}/notification-channels/{id} | Delete a notification channel
[**GetOrganizationsOrgIdNotificationChannels**](NotificationChannelsAPI.md#GetOrganizationsOrgIdNotificationChannels) | **Get** /organizations/{org_id}/notification-channels | List the org&#39;s notification channels
[**GetOrganizationsOrgIdNotificationChannelsId**](NotificationChannelsAPI.md#GetOrganizationsOrgIdNotificationChannelsId) | **Get** /organizations/{org_id}/notification-channels/{id} | Get one notification channel
[**PostOrganizationsOrgIdNotificationChannels**](NotificationChannelsAPI.md#PostOrganizationsOrgIdNotificationChannels) | **Post** /organizations/{org_id}/notification-channels | Register a notification channel
[**PostOrganizationsOrgIdNotificationChannelsIdTest**](NotificationChannelsAPI.md#PostOrganizationsOrgIdNotificationChannelsIdTest) | **Post** /organizations/{org_id}/notification-channels/{id}/test | Send a test delivery to the channel
[**PutOrganizationsOrgIdNotificationChannelsId**](NotificationChannelsAPI.md#PutOrganizationsOrgIdNotificationChannelsId) | **Put** /organizations/{org_id}/notification-channels/{id} | Update a notification channel



## DeleteOrganizationsOrgIdNotificationChannelsId

> map[string]interface{} DeleteOrganizationsOrgIdNotificationChannelsId(ctx, orgId, id).Force(force).Execute()

Delete a notification channel



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
	id := int32(56) // int32 | Channel ID
	force := "force_example" // string | Send the literal value true to prune the channel's alert-rule picks and delete in one cascade (any other value behaves as unset) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationChannelsAPI.DeleteOrganizationsOrgIdNotificationChannelsId(context.Background(), orgId, id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.DeleteOrganizationsOrgIdNotificationChannelsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrganizationsOrgIdNotificationChannelsId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsAPI.DeleteOrganizationsOrgIdNotificationChannelsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationsOrgIdNotificationChannelsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **string** | Send the literal value true to prune the channel&#39;s alert-rule picks and delete in one cascade (any other value behaves as unset) | 

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


## GetOrganizationsOrgIdNotificationChannels

> map[string]interface{} GetOrganizationsOrgIdNotificationChannels(ctx, orgId).Execute()

List the org's notification channels



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
	resp, r, err := apiClient.NotificationChannelsAPI.GetOrganizationsOrgIdNotificationChannels(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.GetOrganizationsOrgIdNotificationChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsOrgIdNotificationChannels`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsAPI.GetOrganizationsOrgIdNotificationChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsOrgIdNotificationChannelsRequest struct via the builder pattern


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


## GetOrganizationsOrgIdNotificationChannelsId

> ModelsNotificationChannel GetOrganizationsOrgIdNotificationChannelsId(ctx, orgId, id).Execute()

Get one notification channel



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
	id := int32(56) // int32 | Channel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationChannelsAPI.GetOrganizationsOrgIdNotificationChannelsId(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.GetOrganizationsOrgIdNotificationChannelsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsOrgIdNotificationChannelsId`: ModelsNotificationChannel
	fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsAPI.GetOrganizationsOrgIdNotificationChannelsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsOrgIdNotificationChannelsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsNotificationChannel**](ModelsNotificationChannel.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsOrgIdNotificationChannels

> ModelsNotificationChannel PostOrganizationsOrgIdNotificationChannels(ctx, orgId).Body(body).Execute()

Register a notification channel



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
	body := map[string]interface{}{ ... } // map[string]interface{} | Channel payload: kind=webhook|email, target=URL or address, label=unique display name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationChannelsAPI.PostOrganizationsOrgIdNotificationChannels(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.PostOrganizationsOrgIdNotificationChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsOrgIdNotificationChannels`: ModelsNotificationChannel
	fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsAPI.PostOrganizationsOrgIdNotificationChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsOrgIdNotificationChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Channel payload: kind&#x3D;webhook|email, target&#x3D;URL or address, label&#x3D;unique display name | 

### Return type

[**ModelsNotificationChannel**](ModelsNotificationChannel.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsOrgIdNotificationChannelsIdTest

> map[string]interface{} PostOrganizationsOrgIdNotificationChannelsIdTest(ctx, orgId, id).Execute()

Send a test delivery to the channel



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
	id := int32(56) // int32 | Channel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationChannelsAPI.PostOrganizationsOrgIdNotificationChannelsIdTest(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.PostOrganizationsOrgIdNotificationChannelsIdTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrganizationsOrgIdNotificationChannelsIdTest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsAPI.PostOrganizationsOrgIdNotificationChannelsIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsOrgIdNotificationChannelsIdTestRequest struct via the builder pattern


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


## PutOrganizationsOrgIdNotificationChannelsId

> ModelsNotificationChannel PutOrganizationsOrgIdNotificationChannelsId(ctx, orgId, id).Body(body).Execute()

Update a notification channel



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
	id := int32(56) // int32 | Channel ID
	body := map[string]interface{}{ ... } // map[string]interface{} | Optional fields: target=URL or address, enabled=true|false, label=unique display name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationChannelsAPI.PutOrganizationsOrgIdNotificationChannelsId(context.Background(), orgId, id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.PutOrganizationsOrgIdNotificationChannelsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrganizationsOrgIdNotificationChannelsId`: ModelsNotificationChannel
	fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsAPI.PutOrganizationsOrgIdNotificationChannelsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Channel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrganizationsOrgIdNotificationChannelsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | Optional fields: target&#x3D;URL or address, enabled&#x3D;true|false, label&#x3D;unique display name | 

### Return type

[**ModelsNotificationChannel**](ModelsNotificationChannel.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

