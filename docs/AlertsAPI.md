# \AlertsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVirtualMachinesIdAlertRulesRuleId**](AlertsAPI.md#DeleteVirtualMachinesIdAlertRulesRuleId) | **Delete** /virtual-machines/{id}/alert-rules/{rule_id} | Delete an alert rule on a VM
[**GetAlerts**](AlertsAPI.md#GetAlerts) | **Get** /alerts | List the account&#39;s fired alert events
[**GetAlertsBucket**](AlertsAPI.md#GetAlertsBucket) | **Get** /alerts/bucket | Storage bucket alert status
[**GetAlertsVolume**](AlertsAPI.md#GetAlertsVolume) | **Get** /alerts/volume | Storage volume alert status
[**GetVirtualMachinesIdAlertRules**](AlertsAPI.md#GetVirtualMachinesIdAlertRules) | **Get** /virtual-machines/{id}/alert-rules | List a VM&#39;s alert rules
[**PostOrgsOrgIdAlertsIdAcknowledge**](AlertsAPI.md#PostOrgsOrgIdAlertsIdAcknowledge) | **Post** /orgs/{org_id}/alerts/{id}/acknowledge | Acknowledge an org alert event
[**PostOrgsOrgIdAlertsIdResolve**](AlertsAPI.md#PostOrgsOrgIdAlertsIdResolve) | **Post** /orgs/{org_id}/alerts/{id}/resolve | Resolve an org alert event
[**PostVirtualMachinesIdAlertRules**](AlertsAPI.md#PostVirtualMachinesIdAlertRules) | **Post** /virtual-machines/{id}/alert-rules | Create an alert rule on a VM
[**PutVirtualMachinesIdAlertRulesRuleId**](AlertsAPI.md#PutVirtualMachinesIdAlertRulesRuleId) | **Put** /virtual-machines/{id}/alert-rules/{rule_id} | Update an alert rule on a VM



## DeleteVirtualMachinesIdAlertRulesRuleId

> map[string]interface{} DeleteVirtualMachinesIdAlertRulesRuleId(ctx, id, ruleId).Execute()

Delete an alert rule on a VM

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
	id := int32(56) // int32 | VM ID
	ruleId := int32(56) // int32 | Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteVirtualMachinesIdAlertRulesRuleId(context.Background(), id, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteVirtualMachinesIdAlertRulesRuleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVirtualMachinesIdAlertRulesRuleId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteVirtualMachinesIdAlertRulesRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 
**ruleId** | **int32** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualMachinesIdAlertRulesRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetAlerts

> HandlersAccountAlertEventsResponse GetAlerts(ctx).XOrgID(xOrgID).XProjectId(xProjectId).Status(status).Limit(limit).Execute()

List the account's fired alert events



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
	xOrgID := int32(56) // int32 | Organization selector (PAT callers; session callers resolve the org from the session) (optional)
	xProjectId := "xProjectId_example" // string | Project selector: a project id or the value all (a PAT caller without it: 428 project_required) (optional)
	status := "status_example" // string | Filter: active returns only unacknowledged events (optional)
	limit := int32(56) // int32 | Page size (default 100, max 500) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlerts(context.Background()).XOrgID(xOrgID).XProjectId(xProjectId).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlerts`: HandlersAccountAlertEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xOrgID** | **int32** | Organization selector (PAT callers; session callers resolve the org from the session) | 
 **xProjectId** | **string** | Project selector: a project id or the value all (a PAT caller without it: 428 project_required) | 
 **status** | **string** | Filter: active returns only unacknowledged events | 
 **limit** | **int32** | Page size (default 100, max 500) | 

### Return type

[**HandlersAccountAlertEventsResponse**](HandlersAccountAlertEventsResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsBucket

> HandlersStorageAlertStatus GetAlertsBucket(ctx).TargetId(targetId).TargetName(targetName).Execute()

Storage bucket alert status



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
	targetId := int32(56) // int32 | S3 connection that holds the bucket
	targetName := "targetName_example" // string | Bucket name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertsBucket(context.Background()).TargetId(targetId).TargetName(targetName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertsBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertsBucket`: HandlersStorageAlertStatus
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertsBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **int32** | S3 connection that holds the bucket | 
 **targetName** | **string** | Bucket name | 

### Return type

[**HandlersStorageAlertStatus**](HandlersStorageAlertStatus.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertsVolume

> HandlersStorageAlertStatus GetAlertsVolume(ctx).TargetId(targetId).TargetName(targetName).Execute()

Storage volume alert status



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
	targetId := int32(56) // int32 | Cluster that holds the volume
	targetName := "targetName_example" // string | Volume as namespace/pvc

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertsVolume(context.Background()).TargetId(targetId).TargetName(targetName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertsVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertsVolume`: HandlersStorageAlertStatus
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertsVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetId** | **int32** | Cluster that holds the volume | 
 **targetName** | **string** | Volume as namespace/pvc | 

### Return type

[**HandlersStorageAlertStatus**](HandlersStorageAlertStatus.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualMachinesIdAlertRules

> []ModelsAlertRule GetVirtualMachinesIdAlertRules(ctx, id).Execute()

List a VM's alert rules

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
	id := int32(56) // int32 | VM ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetVirtualMachinesIdAlertRules(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetVirtualMachinesIdAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualMachinesIdAlertRules`: []ModelsAlertRule
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetVirtualMachinesIdAlertRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualMachinesIdAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelsAlertRule**](ModelsAlertRule.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdAlertsIdAcknowledge

> ModelsAlertEvent PostOrgsOrgIdAlertsIdAcknowledge(ctx, orgId, id).Execute()

Acknowledge an org alert event



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
	id := int32(56) // int32 | Alert event ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.PostOrgsOrgIdAlertsIdAcknowledge(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.PostOrgsOrgIdAlertsIdAcknowledge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdAlertsIdAcknowledge`: ModelsAlertEvent
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.PostOrgsOrgIdAlertsIdAcknowledge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Alert event ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdAlertsIdAcknowledgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsAlertEvent**](ModelsAlertEvent.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdAlertsIdResolve

> ModelsAlertEvent PostOrgsOrgIdAlertsIdResolve(ctx, orgId, id).Execute()

Resolve an org alert event



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
	id := int32(56) // int32 | Alert event ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.PostOrgsOrgIdAlertsIdResolve(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.PostOrgsOrgIdAlertsIdResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdAlertsIdResolve`: ModelsAlertEvent
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.PostOrgsOrgIdAlertsIdResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Alert event ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdAlertsIdResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelsAlertEvent**](ModelsAlertEvent.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVirtualMachinesIdAlertRules

> ModelsAlertRule PostVirtualMachinesIdAlertRules(ctx, id).Request(request).Execute()

Create an alert rule on a VM

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
	id := int32(56) // int32 | VM ID
	request := *openapiclient.NewModelsAlertRule() // ModelsAlertRule | Rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.PostVirtualMachinesIdAlertRules(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.PostVirtualMachinesIdAlertRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVirtualMachinesIdAlertRules`: ModelsAlertRule
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.PostVirtualMachinesIdAlertRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVirtualMachinesIdAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ModelsAlertRule**](ModelsAlertRule.md) | Rule | 

### Return type

[**ModelsAlertRule**](ModelsAlertRule.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVirtualMachinesIdAlertRulesRuleId

> ModelsAlertRule PutVirtualMachinesIdAlertRulesRuleId(ctx, id, ruleId).Request(request).Execute()

Update an alert rule on a VM

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
	id := int32(56) // int32 | VM ID
	ruleId := int32(56) // int32 | Rule ID
	request := *openapiclient.NewModelsAlertRule() // ModelsAlertRule | Rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.PutVirtualMachinesIdAlertRulesRuleId(context.Background(), id, ruleId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.PutVirtualMachinesIdAlertRulesRuleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutVirtualMachinesIdAlertRulesRuleId`: ModelsAlertRule
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.PutVirtualMachinesIdAlertRulesRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | VM ID | 
**ruleId** | **int32** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVirtualMachinesIdAlertRulesRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**ModelsAlertRule**](ModelsAlertRule.md) | Rule | 

### Return type

[**ModelsAlertRule**](ModelsAlertRule.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

