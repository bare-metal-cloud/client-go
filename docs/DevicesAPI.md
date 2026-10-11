# \DevicesAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDevicesDeviceId**](DevicesAPI.md#DeleteDevicesDeviceId) | **Delete** /devices/{device_id} | Revoke a specific device
[**GetDevices**](DevicesAPI.md#GetDevices) | **Get** /devices | List all devices for the current user
[**GetDevicesCurrent**](DevicesAPI.md#GetDevicesCurrent) | **Get** /devices/current | Get current device information
[**PostDevicesRevokeAll**](DevicesAPI.md#PostDevicesRevokeAll) | **Post** /devices/revoke-all | Revoke all devices for the current user
[**PutDevicesCurrentTrustWithKey**](DevicesAPI.md#PutDevicesCurrentTrustWithKey) | **Put** /devices/current/trust-with-key | Mark current device as trusted by proving an unlocked org vault session
[**PutDevicesDeviceIdTrust**](DevicesAPI.md#PutDevicesDeviceIdTrust) | **Put** /devices/{device_id}/trust | Update trust status of a device



## DeleteDevicesDeviceId

> map[string]interface{} DeleteDevicesDeviceId(ctx, deviceId).Execute()

Revoke a specific device



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
	deviceId := "deviceId_example" // string | Device ID to revoke

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.DeleteDevicesDeviceId(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DeleteDevicesDeviceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDevicesDeviceId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.DeleteDevicesDeviceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | Device ID to revoke | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDevicesDeviceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> map[string]interface{} GetDevices(ctx).Execute()

List all devices for the current user



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
	resp, r, err := apiClient.DevicesAPI.GetDevices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevices`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


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


## GetDevicesCurrent

> map[string]interface{} GetDevicesCurrent(ctx).Execute()

Get current device information



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
	resp, r, err := apiClient.DevicesAPI.GetDevicesCurrent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevicesCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevicesCurrent`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevicesCurrent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesCurrentRequest struct via the builder pattern


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


## PostDevicesRevokeAll

> map[string]interface{} PostDevicesRevokeAll(ctx).Execute()

Revoke all devices for the current user



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
	resp, r, err := apiClient.DevicesAPI.PostDevicesRevokeAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.PostDevicesRevokeAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDevicesRevokeAll`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.PostDevicesRevokeAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevicesRevokeAllRequest struct via the builder pattern


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


## PutDevicesCurrentTrustWithKey

> map[string]interface{} PutDevicesCurrentTrustWithKey(ctx).Request(request).Execute()

Mark current device as trusted by proving an unlocked org vault session



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
	request := *openapiclient.NewHandlersTrustCurrentDeviceRequest("EncryptionKey_example") // HandlersTrustCurrentDeviceRequest | encryption_key is required and validated, but the trust gate is the org vault session (X-Org-Vault-Token header)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.PutDevicesCurrentTrustWithKey(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.PutDevicesCurrentTrustWithKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDevicesCurrentTrustWithKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.PutDevicesCurrentTrustWithKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDevicesCurrentTrustWithKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**HandlersTrustCurrentDeviceRequest**](HandlersTrustCurrentDeviceRequest.md) | encryption_key is required and validated, but the trust gate is the org vault session (X-Org-Vault-Token header) | 

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


## PutDevicesDeviceIdTrust

> map[string]interface{} PutDevicesDeviceIdTrust(ctx, deviceId).Request(request).Execute()

Update trust status of a device



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
	deviceId := "deviceId_example" // string | Device ID
	request := map[string]bool{"key": false} // map[string]bool | Trust status (trusted: true/false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.PutDevicesDeviceIdTrust(context.Background(), deviceId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.PutDevicesDeviceIdTrust``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDevicesDeviceIdTrust`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.PutDevicesDeviceIdTrust`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDevicesDeviceIdTrustRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | **map[string]bool** | Trust status (trusted: true/false) | 

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

