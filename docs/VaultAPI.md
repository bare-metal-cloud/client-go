# \VaultAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1VaultJitRequests**](VaultAPI.md#GetApiV1VaultJitRequests) | **Get** /api/v1/vault/jit-requests | List JIT requests (admin)
[**GetApiV1VaultJitStatusRequestToken**](VaultAPI.md#GetApiV1VaultJitStatusRequestToken) | **Get** /api/v1/vault/jit-status/{request_token} | Get JIT request status
[**PostApiV1VaultJitApprove**](VaultAPI.md#PostApiV1VaultJitApprove) | **Post** /api/v1/vault/jit-approve | Approve or deny JIT access request
[**PostApiV1VaultJitRequest**](VaultAPI.md#PostApiV1VaultJitRequest) | **Post** /api/v1/vault/jit-request | Create JIT access request
[**PostOrgsOrgIdVaultInit**](VaultAPI.md#PostOrgsOrgIdVaultInit) | **Post** /orgs/{org_id}/vault/init | Initialize org vault
[**PostOrgsOrgIdVaultRotatePassword**](VaultAPI.md#PostOrgsOrgIdVaultRotatePassword) | **Post** /orgs/{org_id}/vault/rotate-password | Rotate org vault password
[**PostOrgsOrgIdVaultUnlock**](VaultAPI.md#PostOrgsOrgIdVaultUnlock) | **Post** /orgs/{org_id}/vault/unlock | Unlock org vault



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


## PostOrgsOrgIdVaultInit

> HandlersOrgVaultInitResponse PostOrgsOrgIdVaultInit(ctx, orgId).Request(request).Execute()

Initialize org vault



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
	request := *openapiclient.NewHandlersOrgVaultInitRequest() // HandlersOrgVaultInitRequest | Org vault password

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PostOrgsOrgIdVaultInit(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PostOrgsOrgIdVaultInit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdVaultInit`: HandlersOrgVaultInitResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PostOrgsOrgIdVaultInit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdVaultInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersOrgVaultInitRequest**](HandlersOrgVaultInitRequest.md) | Org vault password | 

### Return type

[**HandlersOrgVaultInitResponse**](HandlersOrgVaultInitResponse.md)

### Authorization

No authorization required

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
	request := *openapiclient.NewHandlersOrgVaultRotatePasswordRequest() // HandlersOrgVaultRotatePasswordRequest | Current + new passwords

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

 **request** | [**HandlersOrgVaultRotatePasswordRequest**](HandlersOrgVaultRotatePasswordRequest.md) | Current + new passwords | 

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


## PostOrgsOrgIdVaultUnlock

> HandlersOrgVaultUnlockResponse PostOrgsOrgIdVaultUnlock(ctx, orgId).Request(request).Execute()

Unlock org vault



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
	request := *openapiclient.NewHandlersOrgVaultUnlockRequest() // HandlersOrgVaultUnlockRequest | Org vault password

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultAPI.PostOrgsOrgIdVaultUnlock(context.Background(), orgId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultAPI.PostOrgsOrgIdVaultUnlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdVaultUnlock`: HandlersOrgVaultUnlockResponse
	fmt.Fprintf(os.Stdout, "Response from `VaultAPI.PostOrgsOrgIdVaultUnlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Org ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdVaultUnlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**HandlersOrgVaultUnlockRequest**](HandlersOrgVaultUnlockRequest.md) | Org vault password | 

### Return type

[**HandlersOrgVaultUnlockResponse**](HandlersOrgVaultUnlockResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

