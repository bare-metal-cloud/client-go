# \VaultAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccessGrantsGrantId**](VaultAPI.md#DeleteAccessGrantsGrantId) | **Delete** /access-grants/{grant_id} | Revoke a service-account access grant now
[**DeleteOrgsOrgIdVaultAccessGrantsId**](VaultAPI.md#DeleteOrgsOrgIdVaultAccessGrantsId) | **Delete** /orgs/{org_id}/vault/access-grants/{id} | Revoke an org vault access grant now
[**DeleteOrgsOrgIdVaultRentalsId**](VaultAPI.md#DeleteOrgsOrgIdVaultRentalsId) | **Delete** /orgs/{org_id}/vault/rentals/{id} | Revoke a vault rental (issue #4027)
[**DeleteOrgsOrgIdVaultSharesShareId**](VaultAPI.md#DeleteOrgsOrgIdVaultSharesShareId) | **Delete** /orgs/{org_id}/vault/shares/{share_id} | End a share I sent
[**GetAccessGrants**](VaultAPI.md#GetAccessGrants) | **Get** /access-grants | List active access grants in my org
[**GetApiV1VaultJitRequests**](VaultAPI.md#GetApiV1VaultJitRequests) | **Get** /api/v1/vault/jit-requests | List JIT requests (admin)
[**GetApiV1VaultJitStatusRequestToken**](VaultAPI.md#GetApiV1VaultJitStatusRequestToken) | **Get** /api/v1/vault/jit-status/{request_token} | Get JIT request status
[**GetOrgsOrgIdVaultAccessGrants**](VaultAPI.md#GetOrgsOrgIdVaultAccessGrants) | **Get** /orgs/{org_id}/vault/access-grants | List org vault access grants
[**GetOrgsOrgIdVaultAccessGrantsMe**](VaultAPI.md#GetOrgsOrgIdVaultAccessGrantsMe) | **Get** /orgs/{org_id}/vault/access-grants/me | Check my org vault access
[**GetOrgsOrgIdVaultDeliveries**](VaultAPI.md#GetOrgsOrgIdVaultDeliveries) | **Get** /orgs/{org_id}/vault/deliveries | Read the org vault&#39;s delivery ledger (issue #4030, A7)
[**GetOrgsOrgIdVaultKeys**](VaultAPI.md#GetOrgsOrgIdVaultKeys) | **Get** /orgs/{org_id}/vault/keys | Read the org vault&#39;s dual key slots (issue #4026)
[**GetOrgsOrgIdVaultMachines**](VaultAPI.md#GetOrgsOrgIdVaultMachines) | **Get** /orgs/{org_id}/vault/machines | List the org vault&#39;s machines (issue #4030, A11)
[**GetOrgsOrgIdVaultReceivedShares**](VaultAPI.md#GetOrgsOrgIdVaultReceivedShares) | **Get** /orgs/{org_id}/vault/received-shares | List shares sent to me in an org
[**GetOrgsOrgIdVaultRentalsId**](VaultAPI.md#GetOrgsOrgIdVaultRentalsId) | **Get** /orgs/{org_id}/vault/rentals/{id} | Read a vault rental&#39;s status (issue #4027)
[**GetOrgsOrgIdVaultSalt**](VaultAPI.md#GetOrgsOrgIdVaultSalt) | **Get** /orgs/{org_id}/vault/salt | Read the org vault&#39;s KDF salt (transitional, issue #4025)
[**GetOrgsOrgIdVaultShares**](VaultAPI.md#GetOrgsOrgIdVaultShares) | **Get** /orgs/{org_id}/vault/shares | List shares I have sent in an org
[**GetServiceAccountsSaIdAccessGrants**](VaultAPI.md#GetServiceAccountsSaIdAccessGrants) | **Get** /service-accounts/{sa_id}/access-grants | List access grants for a service account
[**GetServiceAccountsSaIdAccessGrantsMe**](VaultAPI.md#GetServiceAccountsSaIdAccessGrantsMe) | **Get** /service-accounts/{sa_id}/access-grants/me | Check my access to a service account
[**PostApiV1VaultJitApprove**](VaultAPI.md#PostApiV1VaultJitApprove) | **Post** /api/v1/vault/jit-approve | Approve or deny JIT access request
[**PostApiV1VaultJitRequest**](VaultAPI.md#PostApiV1VaultJitRequest) | **Post** /api/v1/vault/jit-request | Create JIT access request
[**PostOrgsOrgIdVaultAccessGrants**](VaultAPI.md#PostOrgsOrgIdVaultAccessGrants) | **Post** /orgs/{org_id}/vault/access-grants | Grant a member org-wide vault access
[**PostOrgsOrgIdVaultMachinesAgentIdRotateHeldSecrets**](VaultAPI.md#PostOrgsOrgIdVaultMachinesAgentIdRotateHeldSecrets) | **Post** /orgs/{org_id}/vault/machines/{agent_id}/rotate-held-secrets | Rotate the secrets a machine holds (issue #4030, A11 — the tab&#39;s only write)
[**PostOrgsOrgIdVaultRentals**](VaultAPI.md#PostOrgsOrgIdVaultRentals) | **Post** /orgs/{org_id}/vault/rentals | Mint an operation-scoped vault secret rental (issue #4027, D3)
[**PostOrgsOrgIdVaultRotatePassword**](VaultAPI.md#PostOrgsOrgIdVaultRotatePassword) | **Post** /orgs/{org_id}/vault/rotate-password | Rotate org vault password
[**PostOrgsOrgIdVaultShares**](VaultAPI.md#PostOrgsOrgIdVaultShares) | **Post** /orgs/{org_id}/vault/shares | Share a vault secret with a recipient
[**PostServiceAccountsSaIdAccessGrants**](VaultAPI.md#PostServiceAccountsSaIdAccessGrants) | **Post** /service-accounts/{sa_id}/access-grants | Grant a member time-bound access to a service account
[**PutOrgsOrgIdVaultSlotsSlot**](VaultAPI.md#PutOrgsOrgIdVaultSlotsSlot) | **Put** /orgs/{org_id}/vault/slots/{slot} | Replace a wrapped org vault slot (issue #4026, A9; #4075 renamed the route)



## DeleteAccessGrantsGrantId

> map[string]string DeleteAccessGrantsGrantId(ctx, grantId).Execute()

Revoke a service-account access grant now

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
	grantId := int32(56) // int32 | Access grant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.DeleteAccessGrantsGrantId(context.Background(), grantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.DeleteAccessGrantsGrantId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccessGrantsGrantId`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.DeleteAccessGrantsGrantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **int32** | Access grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessGrantsGrantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgsOrgIdVaultAccessGrantsId

> map[string]string DeleteOrgsOrgIdVaultAccessGrantsId(ctx, orgId, id).Execute()

Revoke an org vault access grant now



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
	id := int32(56) // int32 | Access grant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.DeleteOrgsOrgIdVaultAccessGrantsId(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.DeleteOrgsOrgIdVaultAccessGrantsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdVaultAccessGrantsId`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.DeleteOrgsOrgIdVaultAccessGrantsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**id** | **int32** | Access grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdVaultAccessGrantsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgsOrgIdVaultRentalsId

> DeleteOrgsOrgIdVaultRentalsId(ctx, orgId, id).Execute()

Revoke a vault rental (issue #4027)



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
	orgId := int32(56) // int32 | Org ID
	id := "id_example" // string | Rental ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VaultAPI.DeleteOrgsOrgIdVaultRentalsId(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.DeleteOrgsOrgIdVaultRentalsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 
**id** | **string** | Rental ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdVaultRentalsIdRequest struct via the builder pattern


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


## DeleteOrgsOrgIdVaultSharesShareId

> HandlersSecretShareView DeleteOrgsOrgIdVaultSharesShareId(ctx, orgId, shareId).Execute()

End a share I sent

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
	shareId := "shareId_example" // string | Share ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.DeleteOrgsOrgIdVaultSharesShareId(context.Background(), orgId, shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.DeleteOrgsOrgIdVaultSharesShareId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdVaultSharesShareId`: HandlersSecretShareView
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.DeleteOrgsOrgIdVaultSharesShareId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**shareId** | **string** | Share ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdVaultSharesShareIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersSecretShareView**](HandlersSecretShareView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessGrants

> []HandlersActiveGrantItem GetAccessGrants(ctx).Execute()

List active access grants in my org



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
	resp, r, err := apiClient.VaultAPI.GetAccessGrants(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetAccessGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessGrants`: []HandlersActiveGrantItem
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetAccessGrants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessGrantsRequest struct via the builder pattern


### Return type

[**[]HandlersActiveGrantItem**](HandlersActiveGrantItem.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1VaultJitRequests

> []HandlersJitRequestListItem GetApiV1VaultJitRequests(ctx).Status(status).Execute()

List JIT requests (admin)



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
	status := "status_example" // string | pending (default) | approved | denied | expired | all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.GetApiV1VaultJitRequests(context.Background()).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetApiV1VaultJitRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1VaultJitRequests`: []HandlersJitRequestListItem
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetApiV1VaultJitRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1VaultJitRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | pending (default) | approved | denied | expired | all | 

### Return type

[**[]HandlersJitRequestListItem**](HandlersJitRequestListItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1VaultJitStatusRequestToken

> map[string]interface{} GetApiV1VaultJitStatusRequestToken(ctx, requestToken).Execute()

Get JIT request status



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
	requestToken := "requestToken_example" // string | Request token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.GetApiV1VaultJitStatusRequestToken(context.Background(), requestToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetApiV1VaultJitStatusRequestToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1VaultJitStatusRequestToken`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetApiV1VaultJitStatusRequestToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestToken** | **string** | Request token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1VaultJitStatusRequestTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdVaultAccessGrants

> []HandlersGrantWithStatus GetOrgsOrgIdVaultAccessGrants(ctx, orgId).Execute()

List org vault access grants



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
	resp, r, err := apiClient.VaultAPI.GetOrgsOrgIdVaultAccessGrants(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetOrgsOrgIdVaultAccessGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdVaultAccessGrants`: []HandlersGrantWithStatus
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetOrgsOrgIdVaultAccessGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdVaultAccessGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HandlersGrantWithStatus**](HandlersGrantWithStatus.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdVaultAccessGrantsMe

> map[string]interface{} GetOrgsOrgIdVaultAccessGrantsMe(ctx, orgId).Execute()

Check my org vault access

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
	resp, r, err := apiClient.VaultAPI.GetOrgsOrgIdVaultAccessGrantsMe(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetOrgsOrgIdVaultAccessGrantsMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdVaultAccessGrantsMe`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetOrgsOrgIdVaultAccessGrantsMe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdVaultAccessGrantsMeRequest struct via the builder pattern


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


## GetOrgsOrgIdVaultDeliveries

> HandlersOrgVaultDeliveriesResponse GetOrgsOrgIdVaultDeliveries(ctx, orgId).Days(days).Execute()

Read the org vault's delivery ledger (issue #4030, A7)



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
	orgId := int32(56) // int32 | Org ID
	days := int32(56) // int32 | Look-back window in days (default 7, max 30) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.GetOrgsOrgIdVaultDeliveries(context.Background(), orgId).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetOrgsOrgIdVaultDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdVaultDeliveries`: HandlersOrgVaultDeliveriesResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetOrgsOrgIdVaultDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdVaultDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **int32** | Look-back window in days (default 7, max 30) | 

### Return type

[**HandlersOrgVaultDeliveriesResponse**](HandlersOrgVaultDeliveriesResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdVaultKeys

> HandlersOrgVaultKeysResponse GetOrgsOrgIdVaultKeys(ctx, orgId).Execute()

Read the org vault's dual key slots (issue #4026)



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
	orgId := int32(56) // int32 | Org ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.GetOrgsOrgIdVaultKeys(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetOrgsOrgIdVaultKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdVaultKeys`: HandlersOrgVaultKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetOrgsOrgIdVaultKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdVaultKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersOrgVaultKeysResponse**](HandlersOrgVaultKeysResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdVaultMachines

> HandlersOrgVaultMachinesResponse GetOrgsOrgIdVaultMachines(ctx, orgId).Execute()

List the org vault's machines (issue #4030, A11)



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
	orgId := int32(56) // int32 | Org ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.GetOrgsOrgIdVaultMachines(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetOrgsOrgIdVaultMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdVaultMachines`: HandlersOrgVaultMachinesResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetOrgsOrgIdVaultMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdVaultMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersOrgVaultMachinesResponse**](HandlersOrgVaultMachinesResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdVaultReceivedShares

> map[string]interface{} GetOrgsOrgIdVaultReceivedShares(ctx, orgId).Execute()

List shares sent to me in an org

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
	resp, r, err := apiClient.VaultAPI.GetOrgsOrgIdVaultReceivedShares(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetOrgsOrgIdVaultReceivedShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdVaultReceivedShares`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetOrgsOrgIdVaultReceivedShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdVaultReceivedSharesRequest struct via the builder pattern


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


## GetOrgsOrgIdVaultRentalsId

> HandlersVaultRentalStatusResponse GetOrgsOrgIdVaultRentalsId(ctx, orgId, id).Execute()

Read a vault rental's status (issue #4027)



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
	orgId := int32(56) // int32 | Org ID
	id := "id_example" // string | Rental ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.GetOrgsOrgIdVaultRentalsId(context.Background(), orgId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetOrgsOrgIdVaultRentalsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdVaultRentalsId`: HandlersVaultRentalStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetOrgsOrgIdVaultRentalsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 
**id** | **string** | Rental ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdVaultRentalsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersVaultRentalStatusResponse**](HandlersVaultRentalStatusResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdVaultSalt

> map[string]string GetOrgsOrgIdVaultSalt(ctx, orgId).Execute()

Read the org vault's KDF salt (transitional, issue #4025)



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
	orgId := int32(56) // int32 | Org ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.GetOrgsOrgIdVaultSalt(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetOrgsOrgIdVaultSalt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdVaultSalt`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetOrgsOrgIdVaultSalt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdVaultSaltRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdVaultShares

> map[string]interface{} GetOrgsOrgIdVaultShares(ctx, orgId).SecretType(secretType).SecretId(secretId).Execute()

List shares I have sent in an org

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
	secretType := "secretType_example" // string | Filter by secret type (optional)
	secretId := int32(56) // int32 | Filter by secret ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.GetOrgsOrgIdVaultShares(context.Background(), orgId).SecretType(secretType).SecretId(secretId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetOrgsOrgIdVaultShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdVaultShares`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetOrgsOrgIdVaultShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdVaultSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secretType** | **string** | Filter by secret type | 
 **secretId** | **int32** | Filter by secret ID | 

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


## GetServiceAccountsSaIdAccessGrants

> []HandlersGrantWithStatus GetServiceAccountsSaIdAccessGrants(ctx, saId).Execute()

List access grants for a service account

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
	saId := int32(56) // int32 | Service account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.GetServiceAccountsSaIdAccessGrants(context.Background(), saId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetServiceAccountsSaIdAccessGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccountsSaIdAccessGrants`: []HandlersGrantWithStatus
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetServiceAccountsSaIdAccessGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saId** | **int32** | Service account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountsSaIdAccessGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HandlersGrantWithStatus**](HandlersGrantWithStatus.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccountsSaIdAccessGrantsMe

> map[string]interface{} GetServiceAccountsSaIdAccessGrantsMe(ctx, saId).Execute()

Check my access to a service account

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
	saId := int32(56) // int32 | Service account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.GetServiceAccountsSaIdAccessGrantsMe(context.Background(), saId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.GetServiceAccountsSaIdAccessGrantsMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccountsSaIdAccessGrantsMe`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.GetServiceAccountsSaIdAccessGrantsMe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saId** | **int32** | Service account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountsSaIdAccessGrantsMeRequest struct via the builder pattern


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


## PostApiV1VaultJitApprove

> map[string]interface{} PostApiV1VaultJitApprove(ctx).Request(request).Execute()

Approve or deny JIT access request



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
	request := *openapiclient.NewHandlersJITApproveRequest() // HandlersJITApproveRequest | JIT approve/deny details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PostApiV1VaultJitApprove(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PostApiV1VaultJitApprove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1VaultJitApprove`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PostApiV1VaultJitApprove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1VaultJitApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersJITApproveRequest**](HandlersJITApproveRequest.md) | JIT approve/deny details | 

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


## PostApiV1VaultJitRequest

> map[string]interface{} PostApiV1VaultJitRequest(ctx).Request(request).Execute()

Create JIT access request



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
	request := *openapiclient.NewHandlersJITCreateRequest() // HandlersJITCreateRequest | JIT request details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PostApiV1VaultJitRequest(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PostApiV1VaultJitRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiV1VaultJitRequest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PostApiV1VaultJitRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApiV1VaultJitRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersJITCreateRequest**](HandlersJITCreateRequest.md) | JIT request details | 

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


## PostOrgsOrgIdVaultAccessGrants

> ModelsAuthorizedAccess PostOrgsOrgIdVaultAccessGrants(ctx, orgId).Request(request).Execute()

Grant a member org-wide vault access



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
	request := *openapiclient.NewHandlersGrantAccessRequest() // HandlersGrantAccessRequest | Grant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PostOrgsOrgIdVaultAccessGrants(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PostOrgsOrgIdVaultAccessGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdVaultAccessGrants`: ModelsAuthorizedAccess
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PostOrgsOrgIdVaultAccessGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdVaultAccessGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersGrantAccessRequest**](HandlersGrantAccessRequest.md) | Grant | 

### Return type

[**ModelsAuthorizedAccess**](ModelsAuthorizedAccess.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdVaultMachinesAgentIdRotateHeldSecrets

> HandlersOrgVaultRotateResponse PostOrgsOrgIdVaultMachinesAgentIdRotateHeldSecrets(ctx, orgId, agentId).Request(request).Execute()

Rotate the secrets a machine holds (issue #4030, A11 — the tab's only write)



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
	orgId := int32(56) // int32 | Org ID
	agentId := int32(56) // int32 | Agent (machine) ID
	request := *openapiclient.NewHandlersRotateHeldSecretsRequest() // HandlersRotateHeldSecretsRequest | The jobs-may-fail acknowledgement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PostOrgsOrgIdVaultMachinesAgentIdRotateHeldSecrets(context.Background(), orgId, agentId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PostOrgsOrgIdVaultMachinesAgentIdRotateHeldSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdVaultMachinesAgentIdRotateHeldSecrets`: HandlersOrgVaultRotateResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PostOrgsOrgIdVaultMachinesAgentIdRotateHeldSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 
**agentId** | **int32** | Agent (machine) ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdVaultMachinesAgentIdRotateHeldSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**HandlersRotateHeldSecretsRequest**](HandlersRotateHeldSecretsRequest.md) | The jobs-may-fail acknowledgement | 

### Return type

[**HandlersOrgVaultRotateResponse**](HandlersOrgVaultRotateResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdVaultRentals

> HandlersVaultRentalStatusResponse PostOrgsOrgIdVaultRentals(ctx, orgId).Request(request).Execute()

Mint an operation-scoped vault secret rental (issue #4027, D3)



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
	orgId := int32(56) // int32 | Org ID
	request := *openapiclient.NewHandlersMintVaultRentalRequest() // HandlersMintVaultRentalRequest | Operation + target + released payload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PostOrgsOrgIdVaultRentals(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PostOrgsOrgIdVaultRentals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdVaultRentals`: HandlersVaultRentalStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PostOrgsOrgIdVaultRentals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdVaultRentalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersMintVaultRentalRequest**](HandlersMintVaultRentalRequest.md) | Operation + target + released payload | 

### Return type

[**HandlersVaultRentalStatusResponse**](HandlersVaultRentalStatusResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdVaultRotatePassword

> HandlersOrgVaultRotatePasswordResponse PostOrgsOrgIdVaultRotatePassword(ctx, orgId).Request(request).Execute()

Rotate org vault password



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
	orgId := int32(56) // int32 | Org ID
	request := *openapiclient.NewHandlersOrgVaultRotatePasswordRequest() // HandlersOrgVaultRotatePasswordRequest | Current + new passwords (optionally the client-wrapped new slot)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PostOrgsOrgIdVaultRotatePassword(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PostOrgsOrgIdVaultRotatePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdVaultRotatePassword`: HandlersOrgVaultRotatePasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PostOrgsOrgIdVaultRotatePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdVaultRotatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersOrgVaultRotatePasswordRequest**](HandlersOrgVaultRotatePasswordRequest.md) | Current + new passwords (optionally the client-wrapped new slot) | 

### Return type

[**HandlersOrgVaultRotatePasswordResponse**](HandlersOrgVaultRotatePasswordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdVaultShares

> HandlersSecretShareView PostOrgsOrgIdVaultShares(ctx, orgId).Request(request).Execute()

Share a vault secret with a recipient



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
	request := *openapiclient.NewHandlersCreateShareRequest() // HandlersCreateShareRequest | Share

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PostOrgsOrgIdVaultShares(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PostOrgsOrgIdVaultShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdVaultShares`: HandlersSecretShareView
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PostOrgsOrgIdVaultShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdVaultSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersCreateShareRequest**](HandlersCreateShareRequest.md) | Share | 

### Return type

[**HandlersSecretShareView**](HandlersSecretShareView.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceAccountsSaIdAccessGrants

> ModelsAuthorizedAccess PostServiceAccountsSaIdAccessGrants(ctx, saId).Request(request).Execute()

Grant a member time-bound access to a service account

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
	saId := int32(56) // int32 | Service account ID
	request := *openapiclient.NewHandlersGrantAccessRequest() // HandlersGrantAccessRequest | Grant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PostServiceAccountsSaIdAccessGrants(context.Background(), saId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PostServiceAccountsSaIdAccessGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceAccountsSaIdAccessGrants`: ModelsAuthorizedAccess
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PostServiceAccountsSaIdAccessGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saId** | **int32** | Service account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceAccountsSaIdAccessGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersGrantAccessRequest**](HandlersGrantAccessRequest.md) | Grant | 

### Return type

[**ModelsAuthorizedAccess**](ModelsAuthorizedAccess.md)

### Authorization

[CookieAuth](../README.md#CookieAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdVaultSlotsSlot

> HandlersOrgVaultReplaceRecoveryResponse PutOrgsOrgIdVaultSlotsSlot(ctx, orgId, slot).Request(request).Execute()

Replace a wrapped org vault slot (issue #4026, A9; #4075 renamed the route)



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
	orgId := int32(56) // int32 | Org ID
	slot := "slot_example" // string | Slot to replace (only the recovery slot is replaceable)
	request := *openapiclient.NewHandlersOrgVaultReplaceRecoveryRequest() // HandlersOrgVaultReplaceRecoveryRequest | New recovery slot (client-wrapped) + ack

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PutOrgsOrgIdVaultSlotsSlot(context.Background(), orgId, slot).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PutOrgsOrgIdVaultSlotsSlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdVaultSlotsSlot`: HandlersOrgVaultReplaceRecoveryResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PutOrgsOrgIdVaultSlotsSlot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 
**slot** | **string** | Slot to replace (only the recovery slot is replaceable) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdVaultSlotsSlotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**HandlersOrgVaultReplaceRecoveryRequest**](HandlersOrgVaultReplaceRecoveryRequest.md) | New recovery slot (client-wrapped) + ack | 

### Return type

[**HandlersOrgVaultReplaceRecoveryResponse**](HandlersOrgVaultReplaceRecoveryResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

