# \RunnersCloudKeysAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgsOrgIdRunnersCloudKeysKeyId**](RunnersCloudKeysAPI.md#DeleteOrgsOrgIdRunnersCloudKeysKeyId) | **Delete** /orgs/{org_id}/runners/cloud-keys/{key_id} | Delete a stored cloud key (admin)
[**DeleteOrgsOrgIdRunnersCloudKeysKeyIdAttach**](RunnersCloudKeysAPI.md#DeleteOrgsOrgIdRunnersCloudKeysKeyIdAttach) | **Delete** /orgs/{org_id}/runners/cloud-keys/{key_id}/attach | Detach a cloud key (admin)
[**GetOrgsOrgIdRunnersCloudKeys**](RunnersCloudKeysAPI.md#GetOrgsOrgIdRunnersCloudKeys) | **Get** /orgs/{org_id}/runners/cloud-keys | List the org&#39;s cloud keys
[**PostOrgsOrgIdRunnersCloudKeys**](RunnersCloudKeysAPI.md#PostOrgsOrgIdRunnersCloudKeys) | **Post** /orgs/{org_id}/runners/cloud-keys | Store a cloud key after a live verification (admin)
[**PostOrgsOrgIdRunnersCloudKeysKeyIdAttach**](RunnersCloudKeysAPI.md#PostOrgsOrgIdRunnersCloudKeysKeyIdAttach) | **Post** /orgs/{org_id}/runners/cloud-keys/{key_id}/attach | Attach a cloud key to the org&#39;s machine-create path (admin)



## DeleteOrgsOrgIdRunnersCloudKeysKeyId

> map[string]string DeleteOrgsOrgIdRunnersCloudKeysKeyId(ctx, orgId, keyId).Execute()

Delete a stored cloud key (admin)



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
	keyId := int32(56) // int32 | Cloud key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersCloudKeysAPI.DeleteOrgsOrgIdRunnersCloudKeysKeyId(context.Background(), orgId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersCloudKeysAPI.DeleteOrgsOrgIdRunnersCloudKeysKeyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdRunnersCloudKeysKeyId`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RunnersCloudKeysAPI.DeleteOrgsOrgIdRunnersCloudKeysKeyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**keyId** | **int32** | Cloud key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdRunnersCloudKeysKeyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgsOrgIdRunnersCloudKeysKeyIdAttach

> HandlersRunnerCloudKeyRow DeleteOrgsOrgIdRunnersCloudKeysKeyIdAttach(ctx, orgId, keyId).Execute()

Detach a cloud key (admin)



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
	keyId := int32(56) // int32 | Cloud key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersCloudKeysAPI.DeleteOrgsOrgIdRunnersCloudKeysKeyIdAttach(context.Background(), orgId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersCloudKeysAPI.DeleteOrgsOrgIdRunnersCloudKeysKeyIdAttach``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgsOrgIdRunnersCloudKeysKeyIdAttach`: HandlersRunnerCloudKeyRow
	fmt.Fprintf(os.Stdout, "Response from `RunnersCloudKeysAPI.DeleteOrgsOrgIdRunnersCloudKeysKeyIdAttach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**keyId** | **int32** | Cloud key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgsOrgIdRunnersCloudKeysKeyIdAttachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerCloudKeyRow**](HandlersRunnerCloudKeyRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgsOrgIdRunnersCloudKeys

> HandlersRunnerCloudKeysResponse GetOrgsOrgIdRunnersCloudKeys(ctx, orgId).Execute()

List the org's cloud keys



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
	resp, r, err := apiClient.RunnersCloudKeysAPI.GetOrgsOrgIdRunnersCloudKeys(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersCloudKeysAPI.GetOrgsOrgIdRunnersCloudKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersCloudKeys`: HandlersRunnerCloudKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersCloudKeysAPI.GetOrgsOrgIdRunnersCloudKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersCloudKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerCloudKeysResponse**](HandlersRunnerCloudKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersCloudKeys

> HandlersRunnerCloudKeyRow PostOrgsOrgIdRunnersCloudKeys(ctx, orgId).Body(body).Execute()

Store a cloud key after a live verification (admin)



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
	body := *openapiclient.NewHandlersRunnerCloudKeySaveBody() // HandlersRunnerCloudKeySaveBody | the key's grammar

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersCloudKeysAPI.PostOrgsOrgIdRunnersCloudKeys(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersCloudKeysAPI.PostOrgsOrgIdRunnersCloudKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersCloudKeys`: HandlersRunnerCloudKeyRow
	fmt.Fprintf(os.Stdout, "Response from `RunnersCloudKeysAPI.PostOrgsOrgIdRunnersCloudKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersCloudKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRunnerCloudKeySaveBody**](HandlersRunnerCloudKeySaveBody.md) | the key&#39;s grammar | 

### Return type

[**HandlersRunnerCloudKeyRow**](HandlersRunnerCloudKeyRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersCloudKeysKeyIdAttach

> HandlersRunnerCloudKeyRow PostOrgsOrgIdRunnersCloudKeysKeyIdAttach(ctx, orgId, keyId).Execute()

Attach a cloud key to the org's machine-create path (admin)



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
	keyId := int32(56) // int32 | Cloud key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersCloudKeysAPI.PostOrgsOrgIdRunnersCloudKeysKeyIdAttach(context.Background(), orgId, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersCloudKeysAPI.PostOrgsOrgIdRunnersCloudKeysKeyIdAttach``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgsOrgIdRunnersCloudKeysKeyIdAttach`: HandlersRunnerCloudKeyRow
	fmt.Fprintf(os.Stdout, "Response from `RunnersCloudKeysAPI.PostOrgsOrgIdRunnersCloudKeysKeyIdAttach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 
**keyId** | **int32** | Cloud key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersCloudKeysKeyIdAttachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlersRunnerCloudKeyRow**](HandlersRunnerCloudKeyRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

