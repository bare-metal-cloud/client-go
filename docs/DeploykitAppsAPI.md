# \DeploykitAppsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClustersIdAppsInstall**](DeploykitAppsAPI.md#DeleteClustersIdAppsInstall) | **Delete** /clusters/{id}/apps/{install} | Tear down an app install (the engine&#39;s keep-data uninstall)
[**DeleteClustersIdAppsInstallComponentsComponent**](DeploykitAppsAPI.md#DeleteClustersIdAppsInstallComponentsComponent) | **Delete** /clusters/{id}/apps/{install}/components/{component} | Remove one component from a stack install (keep-data by default)
[**GetClustersIdApps**](DeploykitAppsAPI.md#GetClustersIdApps) | **Get** /clusters/{id}/apps | List a cluster&#39;s app installs (the engine records)
[**GetClustersIdAppsDestinations**](DeploykitAppsAPI.md#GetClustersIdAppsDestinations) | **Get** /clusters/{id}/apps/destinations | List the org&#39;s viable backup destinations for app installs
[**GetClustersIdAppsInstallBlueprintDrift**](DeploykitAppsAPI.md#GetClustersIdAppsInstallBlueprintDrift) | **Get** /clusters/{id}/apps/{install}/blueprint/drift | Check a blueprint stack&#39;s upstream drift (the standing hint)
[**GetClustersIdAppsInstallBlueprintUpdate**](DeploykitAppsAPI.md#GetClustersIdAppsInstallBlueprintUpdate) | **Get** /clusters/{id}/apps/{install}/blueprint/update | Review a blueprint stack&#39;s update (the compiled-artifact diff)
[**GetClustersIdAppsInstallCliRedisWs**](DeploykitAppsAPI.md#GetClustersIdAppsInstallCliRedisWs) | **Get** /clusters/{id}/apps/{install}/cli/redis/ws | Open the redis service CLI terminal
[**GetClustersIdAppsInstallDurability**](DeploykitAppsAPI.md#GetClustersIdAppsInstallDurability) | **Get** /clusters/{id}/apps/{install}/durability | Get an install&#39;s durability read model
[**GetClustersIdAppsInstallExportDumpJobDownload**](DeploykitAppsAPI.md#GetClustersIdAppsInstallExportDumpJobDownload) | **Get** /clusters/{id}/apps/{install}/export-dump/{job}/download | Download one completed install dump export
[**GetClustersIdAppsInstallForm**](DeploykitAppsAPI.md#GetClustersIdAppsInstallForm) | **Get** /clusters/{id}/apps/{install}/form | Render the deploy form (FormSpec) for an install&#39;s bundle
[**GetClustersIdAppsInstallInspect**](DeploykitAppsAPI.md#GetClustersIdAppsInstallInspect) | **Get** /clusters/{id}/apps/{install}/inspect | Inspect an install&#39;s declared insight collectors live
[**GetClustersIdAppsInstallPromoteTargets**](DeploykitAppsAPI.md#GetClustersIdAppsInstallPromoteTargets) | **Get** /clusters/{id}/apps/{install}/promote/targets | List the org&#39;s eligible promotion targets
[**GetClustersIdAppsInstallStatus**](DeploykitAppsAPI.md#GetClustersIdAppsInstallStatus) | **Get** /clusters/{id}/apps/{install}/status | Get one install&#39;s record status and progress events
[**GetClustersIdAppsInstallUpgradePark**](DeploykitAppsAPI.md#GetClustersIdAppsInstallUpgradePark) | **Get** /clusters/{id}/apps/{install}/upgrade/park | Get an install&#39;s active upgrade park
[**GetClustersIdAppsInstallVerifyDb**](DeploykitAppsAPI.md#GetClustersIdAppsInstallVerifyDb) | **Get** /clusters/{id}/apps/{install}/verify-db | Verify an install component&#39;s database answers
[**GetClustersIdAppsVersions**](DeploykitAppsAPI.md#GetClustersIdAppsVersions) | **Get** /clusters/{id}/apps/versions | List the catalog versions of one bundle this engine can run
[**PostClustersIdAppsInstall**](DeploykitAppsAPI.md#PostClustersIdAppsInstall) | **Post** /clusters/{id}/apps/{install} | Start an app install on a cluster (the resolved form payload)
[**PostClustersIdAppsInstallBlueprintDriftDismiss**](DeploykitAppsAPI.md#PostClustersIdAppsInstallBlueprintDriftDismiss) | **Post** /clusters/{id}/apps/{install}/blueprint/drift/dismiss | Dismiss a blueprint stack&#39;s drift hint instance
[**PostClustersIdAppsInstallExportDump**](DeploykitAppsAPI.md#PostClustersIdAppsInstallExportDump) | **Post** /clusters/{id}/apps/{install}/export-dump | Dump (export) ONE stack-member or engine-install database
[**PostClustersIdAppsInstallPromote**](DeploykitAppsAPI.md#PostClustersIdAppsInstallPromote) | **Post** /clusters/{id}/apps/{install}/promote | Promote an install&#39;s current CI digest to another install
[**PostClustersIdAppsInstallResolve**](DeploykitAppsAPI.md#PostClustersIdAppsInstallResolve) | **Post** /clusters/{id}/apps/{install}/resolve | Resolve the form against user input and platform context facts
[**PostClustersIdAppsInstallRestore**](DeploykitAppsAPI.md#PostClustersIdAppsInstallRestore) | **Post** /clusters/{id}/apps/{install}/restore | Restore an install from a recorded backup point
[**PostClustersIdAppsInstallRestoreDump**](DeploykitAppsAPI.md#PostClustersIdAppsInstallRestoreDump) | **Post** /clusters/{id}/apps/{install}/restore-dump | Restore ONE stack-member or engine-install database from a raw dump
[**PostClustersIdAppsInstallRollback**](DeploykitAppsAPI.md#PostClustersIdAppsInstallRollback) | **Post** /clusters/{id}/apps/{install}/rollback | Re-pin an install to a previous revision&#39;s CI build digest
[**PostClustersIdAppsInstallSlotsSlotRotate**](DeploykitAppsAPI.md#PostClustersIdAppsInstallSlotsSlotRotate) | **Post** /clusters/{id}/apps/{install}/slots/{slot}/rotate | Rotate one managed slot&#39;s credential (the slot and every component of the stack re-apply)
[**PostClustersIdAppsInstallUpgrade**](DeploykitAppsAPI.md#PostClustersIdAppsInstallUpgrade) | **Post** /clusters/{id}/apps/{install}/upgrade | Upgrade an app install to a target catalog version
[**PostClustersIdAppsInstallUpgradeParkConfirm**](DeploykitAppsAPI.md#PostClustersIdAppsInstallUpgradeParkConfirm) | **Post** /clusters/{id}/apps/{install}/upgrade/park/confirm | Confirm (grant) an install&#39;s parked upgrade
[**PostClustersIdAppsInstallUpgradeParkWithdraw**](DeploykitAppsAPI.md#PostClustersIdAppsInstallUpgradeParkWithdraw) | **Post** /clusters/{id}/apps/{install}/upgrade/park/withdraw | Withdraw an install&#39;s parked upgrade
[**PutClustersIdAppsInstallDurability**](DeploykitAppsAPI.md#PutClustersIdAppsInstallDurability) | **Put** /clusters/{id}/apps/{install}/durability | Set an install&#39;s durability policy



## DeleteClustersIdAppsInstall

> map[string]interface{} DeleteClustersIdAppsInstall(ctx, id, install).Execute()

Tear down an app install (the engine's keep-data uninstall)



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.DeleteClustersIdAppsInstall(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.DeleteClustersIdAppsInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClustersIdAppsInstall`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.DeleteClustersIdAppsInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClustersIdAppsInstallRequest struct via the builder pattern


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


## DeleteClustersIdAppsInstallComponentsComponent

> map[string]interface{} DeleteClustersIdAppsInstallComponentsComponent(ctx, id, install, component).Body(body).Execute()

Remove one component from a stack install (keep-data by default)



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Stack install name
	component := "component_example" // string | Component name (as the capture declares it)
	body := *openapiclient.NewHandlersAppsStackRemoveRequest() // HandlersAppsStackRemoveRequest | Data-fate toggles (both default to retain) and, for a purge, the fresh confirm (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.DeleteClustersIdAppsInstallComponentsComponent(context.Background(), id, install, component).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.DeleteClustersIdAppsInstallComponentsComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteClustersIdAppsInstallComponentsComponent`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.DeleteClustersIdAppsInstallComponentsComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Stack install name | 
**component** | **string** | Component name (as the capture declares it) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClustersIdAppsInstallComponentsComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**HandlersAppsStackRemoveRequest**](HandlersAppsStackRemoveRequest.md) | Data-fate toggles (both default to retain) and, for a purge, the fresh confirm | 

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


## GetClustersIdApps

> []HandlersDeploykitInstallSummary GetClustersIdApps(ctx, id).Execute()

List a cluster's app installs (the engine records)



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
	id := "id_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdApps(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdApps`: []HandlersDeploykitInstallSummary
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HandlersDeploykitInstallSummary**](HandlersDeploykitInstallSummary.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsDestinations

> []HandlersInstallDestinationEntry GetClustersIdAppsDestinations(ctx, id).Execute()

List the org's viable backup destinations for app installs



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
	id := "id_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsDestinations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsDestinations`: []HandlersInstallDestinationEntry
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdAppsDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HandlersInstallDestinationEntry**](HandlersInstallDestinationEntry.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsInstallBlueprintDrift

> HandlersBlueprintDriftResponse GetClustersIdAppsInstallBlueprintDrift(ctx, id, install).Execute()

Check a blueprint stack's upstream drift (the standing hint)



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallBlueprintDrift(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallBlueprintDrift``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallBlueprintDrift`: HandlersBlueprintDriftResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdAppsInstallBlueprintDrift`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallBlueprintDriftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersBlueprintDriftResponse**](HandlersBlueprintDriftResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsInstallBlueprintUpdate

> HandlersBlueprintUpdateResponse GetClustersIdAppsInstallBlueprintUpdate(ctx, id, install).Execute()

Review a blueprint stack's update (the compiled-artifact diff)



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallBlueprintUpdate(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallBlueprintUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallBlueprintUpdate`: HandlersBlueprintUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdAppsInstallBlueprintUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallBlueprintUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersBlueprintUpdateResponse**](HandlersBlueprintUpdateResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsInstallCliRedisWs

> GetClustersIdAppsInstallCliRedisWs(ctx, id, install).Execute()

Open the redis service CLI terminal



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
	id := int32(56) // int32 | Cluster ID
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallCliRedisWs(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallCliRedisWs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallCliRedisWsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsInstallDurability

> HandlersDurabilitySection GetClustersIdAppsInstallDurability(ctx, id, install).Limit(limit).Execute()

Get an install's durability read model



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	limit := int32(56) // int32 | Maximum run-ledger rows served (1-200, default 50) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallDurability(context.Background(), id, install).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallDurability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallDurability`: HandlersDurabilitySection
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdAppsInstallDurability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallDurabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Maximum run-ledger rows served (1-200, default 50) | 

### Return type

[**HandlersDurabilitySection**](HandlersDurabilitySection.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsInstallExportDumpJobDownload

> map[string]interface{} GetClustersIdAppsInstallExportDumpJobDownload(ctx, id, install, job).Execute()

Download one completed install dump export



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	job := "job_example" // string | The export job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallExportDumpJobDownload(context.Background(), id, install, job).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallExportDumpJobDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallExportDumpJobDownload`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdAppsInstallExportDumpJobDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 
**job** | **string** | The export job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallExportDumpJobDownloadRequest struct via the builder pattern


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


## GetClustersIdAppsInstallForm

> GetClustersIdAppsInstallForm(ctx, id, install).Bundle(bundle).Version(version).Digest(digest).Execute()

Render the deploy form (FormSpec) for an install's bundle

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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	bundle := "bundle_example" // string | Bundle name — the pre-create branch: render the form for this catalog bundle before any install record exists (the panel wizard); omit to render the install record's deployed bundle (optional)
	version := "version_example" // string | Bundle version (with bundle) (optional)
	digest := "digest_example" // string | Pinned bundle digest (with bundle and version) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallForm(context.Background(), id, install).Bundle(bundle).Version(version).Digest(digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bundle** | **string** | Bundle name — the pre-create branch: render the form for this catalog bundle before any install record exists (the panel wizard); omit to render the install record&#39;s deployed bundle | 
 **version** | **string** | Bundle version (with bundle) | 
 **digest** | **string** | Pinned bundle digest (with bundle and version) | 

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


## GetClustersIdAppsInstallInspect

> EdgeinsightsInsightPanelBundle GetClustersIdAppsInstallInspect(ctx, id, install).Execute()

Inspect an install's declared insight collectors live



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallInspect(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallInspect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallInspect`: EdgeinsightsInsightPanelBundle
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdAppsInstallInspect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallInspectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EdgeinsightsInsightPanelBundle**](EdgeinsightsInsightPanelBundle.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsInstallPromoteTargets

> map[string]interface{} GetClustersIdAppsInstallPromoteTargets(ctx, id, install).Execute()

List the org's eligible promotion targets



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
	id := "id_example" // string | Cluster ID (the source install's)
	install := "install_example" // string | Source install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallPromoteTargets(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallPromoteTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallPromoteTargets`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdAppsInstallPromoteTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID (the source install&#39;s) | 
**install** | **string** | Source install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallPromoteTargetsRequest struct via the builder pattern


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


## GetClustersIdAppsInstallStatus

> HandlersDeploykitInstallStatusResponse GetClustersIdAppsInstallStatus(ctx, id, install).JobId(jobId).Limit(limit).Execute()

Get one install's record status and progress events



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	jobId := "jobId_example" // string | The dispatch's job id (the 202's job_id): surfaces the archived task's error when no record exists (optional)
	limit := int32(56) // int32 | Maximum progress events served (1-500, default 100); the newest events are kept (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallStatus(context.Background(), id, install).JobId(jobId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallStatus`: HandlersDeploykitInstallStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdAppsInstallStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jobId** | **string** | The dispatch&#39;s job id (the 202&#39;s job_id): surfaces the archived task&#39;s error when no record exists | 
 **limit** | **int32** | Maximum progress events served (1-500, default 100); the newest events are kept | 

### Return type

[**HandlersDeploykitInstallStatusResponse**](HandlersDeploykitInstallStatusResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsInstallUpgradePark

> HandlersUpgradeParkView GetClustersIdAppsInstallUpgradePark(ctx, id, install).Execute()

Get an install's active upgrade park



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallUpgradePark(context.Background(), id, install).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallUpgradePark``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallUpgradePark`: HandlersUpgradeParkView
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdAppsInstallUpgradePark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallUpgradeParkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersUpgradeParkView**](HandlersUpgradeParkView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsInstallVerifyDb

> EngineportsVerifyResult GetClustersIdAppsInstallVerifyDb(ctx, id, install).Component(component).Execute()

Verify an install component's database answers



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	component := "component_example" // string | Component name (the bundle's stateful component)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsInstallVerifyDb(context.Background(), id, install).Component(component).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsInstallVerifyDb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersIdAppsInstallVerifyDb`: EngineportsVerifyResult
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.GetClustersIdAppsInstallVerifyDb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsInstallVerifyDbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **component** | **string** | Component name (the bundle&#39;s stateful component) | 

### Return type

[**EngineportsVerifyResult**](EngineportsVerifyResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersIdAppsVersions

> GetClustersIdAppsVersions(ctx, id).Bundle(bundle).Execute()

List the catalog versions of one bundle this engine can run

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
	id := "id_example" // string | Cluster ID
	bundle := "bundle_example" // string | Bundle name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploykitAppsAPI.GetClustersIdAppsVersions(context.Background(), id).Bundle(bundle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.GetClustersIdAppsVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersIdAppsVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundle** | **string** | Bundle name | 

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


## PostClustersIdAppsInstall

> map[string]interface{} PostClustersIdAppsInstall(ctx, id, install).Body(body).Execute()

Start an app install on a cluster (the resolved form payload)



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersDeploykitResolveRequest() // HandlersDeploykitResolveRequest | Resolved form payload (bundle ref, target, input, facts)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstall(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstall`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersDeploykitResolveRequest**](HandlersDeploykitResolveRequest.md) | Resolved form payload (bundle ref, target, input, facts) | 

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


## PostClustersIdAppsInstallBlueprintDriftDismiss

> map[string]interface{} PostClustersIdAppsInstallBlueprintDriftDismiss(ctx, id, install).Body(body).Execute()

Dismiss a blueprint stack's drift hint instance



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersBlueprintDismissRequest() // HandlersBlueprintDismissRequest | The dismissed instance's key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallBlueprintDriftDismiss(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallBlueprintDriftDismiss``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallBlueprintDriftDismiss`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstallBlueprintDriftDismiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallBlueprintDriftDismissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersBlueprintDismissRequest**](HandlersBlueprintDismissRequest.md) | The dismissed instance&#39;s key | 

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


## PostClustersIdAppsInstallExportDump

> map[string]interface{} PostClustersIdAppsInstallExportDump(ctx, id, install).Body(body).Execute()

Dump (export) ONE stack-member or engine-install database



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersDumpExportRequest() // HandlersDumpExportRequest | destination_ref; object_key optional (defaults under managed-db-dumps/) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallExportDump(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallExportDump``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallExportDump`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstallExportDump`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallExportDumpRequest struct via the builder pattern


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


## PostClustersIdAppsInstallPromote

> map[string]interface{} PostClustersIdAppsInstallPromote(ctx, id, install).Body(body).Execute()

Promote an install's current CI digest to another install



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
	id := "id_example" // string | Cluster ID (the source install's)
	install := "install_example" // string | Source install name
	body := *openapiclient.NewHandlersAppsPromoteRequest() // HandlersAppsPromoteRequest | The target install and cluster

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallPromote(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallPromote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallPromote`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstallPromote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID (the source install&#39;s) | 
**install** | **string** | Source install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallPromoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersAppsPromoteRequest**](HandlersAppsPromoteRequest.md) | The target install and cluster | 

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


## PostClustersIdAppsInstallResolve

> PostClustersIdAppsInstallResolve(ctx, id, install).Body(body).Execute()

Resolve the form against user input and platform context facts

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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersDeploykitResolveRequest() // HandlersDeploykitResolveRequest | Resolve request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallResolve(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersDeploykitResolveRequest**](HandlersDeploykitResolveRequest.md) | Resolve request | 

### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClustersIdAppsInstallRestore

> map[string]interface{} PostClustersIdAppsInstallRestore(ctx, id, install).Body(body).Execute()

Restore an install from a recorded backup point



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersAppsRestoreRequest() // HandlersAppsRestoreRequest | The point's archive handle, the target mode, and (on the confirmed re-post) the fresh confirm

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallRestore(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallRestore`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstallRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersAppsRestoreRequest**](HandlersAppsRestoreRequest.md) | The point&#39;s archive handle, the target mode, and (on the confirmed re-post) the fresh confirm | 

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


## PostClustersIdAppsInstallRestoreDump

> map[string]interface{} PostClustersIdAppsInstallRestoreDump(ctx, id, install).Body(body).Execute()

Restore ONE stack-member or engine-install database from a raw dump



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersDumpRestoreRequest() // HandlersDumpRestoreRequest | destination_ref, object_key, optional digest; confirm on the re-post

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallRestoreDump(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallRestoreDump``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallRestoreDump`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstallRestoreDump`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallRestoreDumpRequest struct via the builder pattern


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


## PostClustersIdAppsInstallRollback

> map[string]interface{} PostClustersIdAppsInstallRollback(ctx, id, install).Body(body).Execute()

Re-pin an install to a previous revision's CI build digest



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersAppsRollbackRequest() // HandlersAppsRollbackRequest | The revision number to re-pin to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallRollback(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallRollback`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstallRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersAppsRollbackRequest**](HandlersAppsRollbackRequest.md) | The revision number to re-pin to | 

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


## PostClustersIdAppsInstallSlotsSlotRotate

> map[string]interface{} PostClustersIdAppsInstallSlotsSlotRotate(ctx, id, install, slot).Body(body).Execute()

Rotate one managed slot's credential (the slot and every component of the stack re-apply)



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Stack install name
	slot := "slot_example" // string | The slot's service name (the managed_slots key)
	body := *openapiclient.NewHandlersAppsSlotRotateRequest() // HandlersAppsSlotRotateRequest | The fresh confirm (omit for the dialog) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallSlotsSlotRotate(context.Background(), id, install, slot).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallSlotsSlotRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallSlotsSlotRotate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstallSlotsSlotRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Stack install name | 
**slot** | **string** | The slot&#39;s service name (the managed_slots key) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallSlotsSlotRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**HandlersAppsSlotRotateRequest**](HandlersAppsSlotRotateRequest.md) | The fresh confirm (omit for the dialog) | 

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


## PostClustersIdAppsInstallUpgrade

> map[string]interface{} PostClustersIdAppsInstallUpgrade(ctx, id, install).Body(body).Execute()

Upgrade an app install to a target catalog version



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersAppsUpgradeRequest() // HandlersAppsUpgradeRequest | Target version (and, for a migrating upgrade, the fresh confirm)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallUpgrade(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallUpgrade`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstallUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersAppsUpgradeRequest**](HandlersAppsUpgradeRequest.md) | Target version (and, for a migrating upgrade, the fresh confirm) | 

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


## PostClustersIdAppsInstallUpgradeParkConfirm

> map[string]interface{} PostClustersIdAppsInstallUpgradeParkConfirm(ctx, id, install).Body(body).Execute()

Confirm (grant) an install's parked upgrade



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersParkActionRequest() // HandlersParkActionRequest | The park_id (and plan_fingerprint) the panel rendered

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallUpgradeParkConfirm(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallUpgradeParkConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallUpgradeParkConfirm`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstallUpgradeParkConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallUpgradeParkConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersParkActionRequest**](HandlersParkActionRequest.md) | The park_id (and plan_fingerprint) the panel rendered | 

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


## PostClustersIdAppsInstallUpgradeParkWithdraw

> map[string]interface{} PostClustersIdAppsInstallUpgradeParkWithdraw(ctx, id, install).Body(body).Execute()

Withdraw an install's parked upgrade



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := *openapiclient.NewHandlersParkActionRequest() // HandlersParkActionRequest | The park_id the panel rendered

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PostClustersIdAppsInstallUpgradeParkWithdraw(context.Background(), id, install).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PostClustersIdAppsInstallUpgradeParkWithdraw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClustersIdAppsInstallUpgradeParkWithdraw`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PostClustersIdAppsInstallUpgradeParkWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostClustersIdAppsInstallUpgradeParkWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersParkActionRequest**](HandlersParkActionRequest.md) | The park_id the panel rendered | 

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


## PutClustersIdAppsInstallDurability

> HandlersDurabilitySection PutClustersIdAppsInstallDurability(ctx, id, install).Body(body).XOrgID(xOrgID).XProjectId(xProjectId).Execute()

Set an install's durability policy



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
	id := "id_example" // string | Cluster ID
	install := "install_example" // string | Install name
	body := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Durability answers keyed by bundle field id (durability.mode, durability.schedule, durability.destination, durability.retention); string values. Optional policy_gen integer: the gen the client loaded (the GET read model exposes it) — a different gen answers 409
	xOrgID := int32(56) // int32 | Organization selector (PAT callers; session callers resolve the org from the session) (optional)
	xProjectId := "xProjectId_example" // string | Project selector: the project owning the cluster (a PAT caller without it resolves project 0, and the scoped cluster lookup answers 404) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploykitAppsAPI.PutClustersIdAppsInstallDurability(context.Background(), id, install).Body(body).XOrgID(xOrgID).XProjectId(xProjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploykitAppsAPI.PutClustersIdAppsInstallDurability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutClustersIdAppsInstallDurability`: HandlersDurabilitySection
	fmt.Fprintf(os.Stdout, "Response from `DeploykitAppsAPI.PutClustersIdAppsInstallDurability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cluster ID | 
**install** | **string** | Install name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutClustersIdAppsInstallDurabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | Durability answers keyed by bundle field id (durability.mode, durability.schedule, durability.destination, durability.retention); string values. Optional policy_gen integer: the gen the client loaded (the GET read model exposes it) — a different gen answers 409 | 
 **xOrgID** | **int32** | Organization selector (PAT callers; session callers resolve the org from the session) | 
 **xProjectId** | **string** | Project selector: the project owning the cluster (a PAT caller without it resolves project 0, and the scoped cluster lookup answers 404) | 

### Return type

[**HandlersDurabilitySection**](HandlersDurabilitySection.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

