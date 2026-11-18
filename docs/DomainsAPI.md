# \DomainsAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgsOrgIdArrivalsInstance**](DomainsAPI.md#DeleteOrgsOrgIdArrivalsInstance) | **Delete** /orgs/{org_id}/arrivals/{instance} | Revoke a tunnel addon&#39;s management custody
[**DeleteOrgsOrgIdDomainsMigrationConsentsId**](DomainsAPI.md#DeleteOrgsOrgIdDomainsMigrationConsentsId) | **Delete** /orgs/{org_id}/domains/migration-consents/{id} | Revoke a live DNS record-migration consent
[**DeleteOrgsOrgIdDomainsZonesId**](DomainsAPI.md#DeleteOrgsOrgIdDomainsZonesId) | **Delete** /orgs/{org_id}/domains/zones/{id} | Remove a registered zone
[**GetOrgsOrgIdArrivals**](DomainsAPI.md#GetOrgsOrgIdArrivals) | **Get** /orgs/{org_id}/arrivals | List the org&#39;s arrival custody rows
[**GetOrgsOrgIdArrivalsInstanceReservedDomains**](DomainsAPI.md#GetOrgsOrgIdArrivalsInstanceReservedDomains) | **Get** /orgs/{org_id}/arrivals/{instance}/reserved-domains | List a ngrok addon&#39;s reserved domains (the 154.6 leg)
[**GetOrgsOrgIdArrivalsInstanceStrandedRecords**](DomainsAPI.md#GetOrgsOrgIdArrivalsInstanceStrandedRecords) | **Get** /orgs/{org_id}/arrivals/{instance}/stranded-records | Preview the org&#39;s stranded tunnel records (the bulk recovery&#39;s dry run)
[**GetOrgsOrgIdDomainsMigrationConsents**](DomainsAPI.md#GetOrgsOrgIdDomainsMigrationConsents) | **Get** /orgs/{org_id}/domains/migration-consents | List the org&#39;s live DNS record-migration consents
[**GetOrgsOrgIdDomainsZones**](DomainsAPI.md#GetOrgsOrgIdDomainsZones) | **Get** /orgs/{org_id}/domains/zones | List the org&#39;s registered zones (paginated, with health)
[**GetOrgsOrgIdDomainsZonesIdAvailability**](DomainsAPI.md#GetOrgsOrgIdDomainsZonesIdAvailability) | **Get** /orgs/{org_id}/domains/zones/{id}/availability | Check a hostname&#39;s availability in a registered zone
[**PostOrgsOrgIdArrivals**](DomainsAPI.md#PostOrgsOrgIdArrivals) | **Post** /orgs/{org_id}/arrivals | Register a tunnel addon&#39;s management custody (the install leg)
[**PostOrgsOrgIdArrivalsInstanceRotate**](DomainsAPI.md#PostOrgsOrgIdArrivalsInstanceRotate) | **Post** /orgs/{org_id}/arrivals/{instance}/rotate | Rotate a tunnel addon&#39;s management custody
[**PostOrgsOrgIdArrivalsInstanceStrandedRecordsRetarget**](DomainsAPI.md#PostOrgsOrgIdArrivalsInstanceStrandedRecordsRetarget) | **Post** /orgs/{org_id}/arrivals/{instance}/stranded-records/retarget | Approve and dispatch the bulk stranded-record re-target
[**PostOrgsOrgIdDomainsMigrationConsents**](DomainsAPI.md#PostOrgsOrgIdDomainsMigrationConsents) | **Post** /orgs/{org_id}/domains/migration-consents | Record a DNS record-migration consent (org admin)
[**PostOrgsOrgIdDomainsValidate**](DomainsAPI.md#PostOrgsOrgIdDomainsValidate) | **Post** /orgs/{org_id}/domains/validate | Validate a DNS provider token live (the connect checklist)
[**PostOrgsOrgIdDomainsZones**](DomainsAPI.md#PostOrgsOrgIdDomainsZones) | **Post** /orgs/{org_id}/domains/zones | Register discovered zones (the connect leg)
[**PostOrgsOrgIdDomainsZonesIdRotate**](DomainsAPI.md#PostOrgsOrgIdDomainsZonesIdRotate) | **Post** /orgs/{org_id}/domains/zones/{id}/rotate | Rotate a zone&#39;s credential (re-seal or merge)



## DeleteOrgsOrgIdArrivalsInstance

> map[string]interface{} DeleteOrgsOrgIdArrivalsInstance(ctx, orgId, instance).Provider(provider).Force(force).Execute()

Revoke a tunnel addon's management custody



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
	orgId := "orgId_example" // string | Org ID
	instance := "instance_example" // string | Addon instance id
	provider := "provider_example" // string | Provider (default cloudflare) (optional)
	force := true // bool | Force the revoke past the host-claim guard (org admin/owner only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DeleteOrgsOrgIdArrivalsInstance(context.Background(), orgId, instance).Provider(provider).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteOrgsOrgIdArrivalsInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdArrivalsInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DeleteOrgsOrgIdArrivalsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**instance** | **string** | Addon instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdArrivalsInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provider** | **string** | Provider (default cloudflare) | 
 **force** | **bool** | Force the revoke past the host-claim guard (org admin/owner only) | 

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


## DeleteOrgsOrgIdDomainsMigrationConsentsId

> map[string]interface{} DeleteOrgsOrgIdDomainsMigrationConsentsId(ctx, orgId, id).Execute()

Revoke a live DNS record-migration consent



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
	orgId := "orgId_example" // string | Org ID
	id := "id_example" // string | Consent ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DeleteOrgsOrgIdDomainsMigrationConsentsId(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteOrgsOrgIdDomainsMigrationConsentsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdDomainsMigrationConsentsId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DeleteOrgsOrgIdDomainsMigrationConsentsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**id** | **string** | Consent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdDomainsMigrationConsentsIdRequest struct via the builder pattern


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


## DeleteOrgsOrgIdDomainsZonesId

> map[string]interface{} DeleteOrgsOrgIdDomainsZonesId(ctx, orgId, id).Force(force).Execute()

Remove a registered zone



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
	orgId := "orgId_example" // string | Org ID
	id := "id_example" // string | Zone ID
	force := true // bool | Force the named cascade (org admin/owner only) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.DeleteOrgsOrgIdDomainsZonesId(context.Background(), orgId, id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteOrgsOrgIdDomainsZonesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdDomainsZonesId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.DeleteOrgsOrgIdDomainsZonesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**id** | **string** | Zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdDomainsZonesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | Force the named cascade (org admin/owner only) | 

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


## GetOrgsOrgIdArrivals

> []ModelsArrivalProviderCredential GetOrgsOrgIdArrivals(ctx, orgId).Execute()

List the org's arrival custody rows



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
	orgId := "orgId_example" // string | Org ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.GetOrgsOrgIdArrivals(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetOrgsOrgIdArrivals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdArrivals`: []ModelsArrivalProviderCredential
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetOrgsOrgIdArrivals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdArrivalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelsArrivalProviderCredential**](ModelsArrivalProviderCredential.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdArrivalsInstanceReservedDomains

> []DnsregistryReservedDomainRow GetOrgsOrgIdArrivalsInstanceReservedDomains(ctx, orgId, instance).Execute()

List a ngrok addon's reserved domains (the 154.6 leg)



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
	orgId := "orgId_example" // string | Org ID
	instance := "instance_example" // string | Addon instance id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.GetOrgsOrgIdArrivalsInstanceReservedDomains(context.Background(), orgId, instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetOrgsOrgIdArrivalsInstanceReservedDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdArrivalsInstanceReservedDomains`: []DnsregistryReservedDomainRow
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetOrgsOrgIdArrivalsInstanceReservedDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**instance** | **string** | Addon instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdArrivalsInstanceReservedDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DnsregistryReservedDomainRow**](DnsregistryReservedDomainRow.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdArrivalsInstanceStrandedRecords

> EngineportsStrandRetargetPreview GetOrgsOrgIdArrivalsInstanceStrandedRecords(ctx, orgId, instance).ClusterId(clusterId).Execute()

Preview the org's stranded tunnel records (the bulk recovery's dry run)



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
	orgId := "orgId_example" // string | Org ID
	instance := "instance_example" // string | Addon instance id
	clusterId := int32(56) // int32 | Cluster row id of the arrival instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.GetOrgsOrgIdArrivalsInstanceStrandedRecords(context.Background(), orgId, instance).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetOrgsOrgIdArrivalsInstanceStrandedRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdArrivalsInstanceStrandedRecords`: EngineportsStrandRetargetPreview
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetOrgsOrgIdArrivalsInstanceStrandedRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**instance** | **string** | Addon instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdArrivalsInstanceStrandedRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterId** | **int32** | Cluster row id of the arrival instance | 

### Return type

[**EngineportsStrandRetargetPreview**](EngineportsStrandRetargetPreview.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdDomainsMigrationConsents

> []ModelsDnsMigrationConsent GetOrgsOrgIdDomainsMigrationConsents(ctx, orgId).ZoneId(zoneId).Host(host).Execute()

List the org's live DNS record-migration consents



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
	orgId := "orgId_example" // string | Org ID
	zoneId := int32(56) // int32 | Narrow to one registered zone id (optional)
	host := "host_example" // string | Narrow to one hostname (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.GetOrgsOrgIdDomainsMigrationConsents(context.Background(), orgId).ZoneId(zoneId).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetOrgsOrgIdDomainsMigrationConsents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdDomainsMigrationConsents`: []ModelsDnsMigrationConsent
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetOrgsOrgIdDomainsMigrationConsents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdDomainsMigrationConsentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneId** | **int32** | Narrow to one registered zone id | 
 **host** | **string** | Narrow to one hostname | 

### Return type

[**[]ModelsDnsMigrationConsent**](ModelsDnsMigrationConsent.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdDomainsZones

> map[string]interface{} GetOrgsOrgIdDomainsZones(ctx, orgId).Page(page).PerPage(perPage).Execute()

List the org's registered zones (paginated, with health)



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
	orgId := "orgId_example" // string | Org ID
	page := int32(56) // int32 | Page (default 1) (optional)
	perPage := int32(56) // int32 | Zones per page (default 50, max 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.GetOrgsOrgIdDomainsZones(context.Background(), orgId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetOrgsOrgIdDomainsZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdDomainsZones`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetOrgsOrgIdDomainsZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdDomainsZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page (default 1) | 
 **perPage** | **int32** | Zones per page (default 50, max 100) | 

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


## GetOrgsOrgIdDomainsZonesIdAvailability

> DnsregistryAvailabilityResult GetOrgsOrgIdDomainsZonesIdAvailability(ctx, orgId, id).Host(host).Execute()

Check a hostname's availability in a registered zone



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
	orgId := "orgId_example" // string | Org ID
	id := "id_example" // string | Zone ID
	host := "host_example" // string | Hostname to check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.GetOrgsOrgIdDomainsZonesIdAvailability(context.Background(), orgId, id).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.GetOrgsOrgIdDomainsZonesIdAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdDomainsZonesIdAvailability`: DnsregistryAvailabilityResult
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.GetOrgsOrgIdDomainsZonesIdAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**id** | **string** | Zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdDomainsZonesIdAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **host** | **string** | Hostname to check | 

### Return type

[**DnsregistryAvailabilityResult**](DnsregistryAvailabilityResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdArrivals

> ModelsArrivalProviderCredential PostOrgsOrgIdArrivals(ctx, orgId).Body(body).Execute()

Register a tunnel addon's management custody (the install leg)



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
	orgId := "orgId_example" // string | Org ID
	body := *openapiclient.NewHandlersArrivalRegisterRequest() // HandlersArrivalRegisterRequest | Provider, instance id, and the management token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.PostOrgsOrgIdArrivals(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.PostOrgsOrgIdArrivals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdArrivals`: ModelsArrivalProviderCredential
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.PostOrgsOrgIdArrivals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdArrivalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersArrivalRegisterRequest**](HandlersArrivalRegisterRequest.md) | Provider, instance id, and the management token | 

### Return type

[**ModelsArrivalProviderCredential**](ModelsArrivalProviderCredential.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdArrivalsInstanceRotate

> map[string]interface{} PostOrgsOrgIdArrivalsInstanceRotate(ctx, orgId, instance).Body(body).Execute()

Rotate a tunnel addon's management custody



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
	orgId := "orgId_example" // string | Org ID
	instance := "instance_example" // string | Addon instance id
	body := *openapiclient.NewHandlersArrivalRegisterRequest() // HandlersArrivalRegisterRequest | Provider, instance, and the new token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.PostOrgsOrgIdArrivalsInstanceRotate(context.Background(), orgId, instance).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.PostOrgsOrgIdArrivalsInstanceRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdArrivalsInstanceRotate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.PostOrgsOrgIdArrivalsInstanceRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**instance** | **string** | Addon instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdArrivalsInstanceRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersArrivalRegisterRequest**](HandlersArrivalRegisterRequest.md) | Provider, instance, and the new token | 

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


## PostOrgsOrgIdArrivalsInstanceStrandedRecordsRetarget

> map[string]interface{} PostOrgsOrgIdArrivalsInstanceStrandedRecordsRetarget(ctx, orgId, instance).Body(body).Execute()

Approve and dispatch the bulk stranded-record re-target



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
	orgId := "orgId_example" // string | Org ID
	instance := "instance_example" // string | Addon instance id
	body := *openapiclient.NewHandlersStrandRetargetRequest() // HandlersStrandRetargetRequest | The cluster id and the reviewed records

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.PostOrgsOrgIdArrivalsInstanceStrandedRecordsRetarget(context.Background(), orgId, instance).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.PostOrgsOrgIdArrivalsInstanceStrandedRecordsRetarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdArrivalsInstanceStrandedRecordsRetarget`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.PostOrgsOrgIdArrivalsInstanceStrandedRecordsRetarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**instance** | **string** | Addon instance id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdArrivalsInstanceStrandedRecordsRetargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersStrandRetargetRequest**](HandlersStrandRetargetRequest.md) | The cluster id and the reviewed records | 

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


## PostOrgsOrgIdDomainsMigrationConsents

> ModelsDnsMigrationConsent PostOrgsOrgIdDomainsMigrationConsents(ctx, orgId).Body(body).Execute()

Record a DNS record-migration consent (org admin)



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
	orgId := "orgId_example" // string | Org ID
	body := *openapiclient.NewHandlersMigrationConsentCreateRequest() // HandlersMigrationConsentCreateRequest | Zone, host, the record's current owner, and the target cluster

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.PostOrgsOrgIdDomainsMigrationConsents(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.PostOrgsOrgIdDomainsMigrationConsents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdDomainsMigrationConsents`: ModelsDnsMigrationConsent
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.PostOrgsOrgIdDomainsMigrationConsents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdDomainsMigrationConsentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersMigrationConsentCreateRequest**](HandlersMigrationConsentCreateRequest.md) | Zone, host, the record&#39;s current owner, and the target cluster | 

### Return type

[**ModelsDnsMigrationConsent**](ModelsDnsMigrationConsent.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdDomainsValidate

> DnsregistryValidateResult PostOrgsOrgIdDomainsValidate(ctx, orgId).Body(body).Execute()

Validate a DNS provider token live (the connect checklist)



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
	orgId := "orgId_example" // string | Org ID
	body := *openapiclient.NewHandlersDomainsValidateRequest() // HandlersDomainsValidateRequest | Provider and token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.PostOrgsOrgIdDomainsValidate(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.PostOrgsOrgIdDomainsValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdDomainsValidate`: DnsregistryValidateResult
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.PostOrgsOrgIdDomainsValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdDomainsValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersDomainsValidateRequest**](HandlersDomainsValidateRequest.md) | Provider and token | 

### Return type

[**DnsregistryValidateResult**](DnsregistryValidateResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdDomainsZones

> DnsregistryRegisterResult PostOrgsOrgIdDomainsZones(ctx, orgId).Body(body).Execute()

Register discovered zones (the connect leg)



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
	orgId := "orgId_example" // string | Org ID
	body := *openapiclient.NewHandlersDomainsRegisterRequest() // HandlersDomainsRegisterRequest | Provider, token, and the discovered zone ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.PostOrgsOrgIdDomainsZones(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.PostOrgsOrgIdDomainsZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdDomainsZones`: DnsregistryRegisterResult
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.PostOrgsOrgIdDomainsZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdDomainsZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersDomainsRegisterRequest**](HandlersDomainsRegisterRequest.md) | Provider, token, and the discovered zone ids | 

### Return type

[**DnsregistryRegisterResult**](DnsregistryRegisterResult.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdDomainsZonesIdRotate

> ModelsDnsZone PostOrgsOrgIdDomainsZonesIdRotate(ctx, orgId, id).Body(body).Execute()

Rotate a zone's credential (re-seal or merge)



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
	orgId := "orgId_example" // string | Org ID
	id := "id_example" // string | Zone ID
	body := *openapiclient.NewHandlersDomainsRotateRequest() // HandlersDomainsRotateRequest | The new token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainsAPI.PostOrgsOrgIdDomainsZonesIdRotate(context.Background(), orgId, id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.PostOrgsOrgIdDomainsZonesIdRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdDomainsZonesIdRotate`: ModelsDnsZone
	fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.PostOrgsOrgIdDomainsZonesIdRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Org ID | 
**id** | **string** | Zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdDomainsZonesIdRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HandlersDomainsRotateRequest**](HandlersDomainsRotateRequest.md) | The new token | 

### Return type

[**ModelsDnsZone**](ModelsDnsZone.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

