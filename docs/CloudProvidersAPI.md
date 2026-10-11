# \CloudProvidersAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudProvidersAwsAmi**](CloudProvidersAPI.md#GetCloudProvidersAwsAmi) | **Get** /cloud-providers/aws/ami | Get available AWS AMIs for a region
[**GetCloudProvidersAwsInstanceTypes**](CloudProvidersAPI.md#GetCloudProvidersAwsInstanceTypes) | **Get** /cloud-providers/aws/instance-types | Get available AWS instance types
[**GetCloudProvidersAwsKeyPairs**](CloudProvidersAPI.md#GetCloudProvidersAwsKeyPairs) | **Get** /cloud-providers/aws/key-pairs | Get AWS key pairs for credentials
[**GetCloudProvidersAwsRegions**](CloudProvidersAPI.md#GetCloudProvidersAwsRegions) | **Get** /cloud-providers/aws/regions | Get available AWS regions
[**GetCloudProvidersAwsSecurityGroups**](CloudProvidersAPI.md#GetCloudProvidersAwsSecurityGroups) | **Get** /cloud-providers/aws/security-groups | Get AWS security groups for credentials
[**GetCloudProvidersAzureRegions**](CloudProvidersAPI.md#GetCloudProvidersAzureRegions) | **Get** /cloud-providers/azure/regions | Get available Azure regions
[**GetCloudProvidersAzureVmSizes**](CloudProvidersAPI.md#GetCloudProvidersAzureVmSizes) | **Get** /cloud-providers/azure/vm-sizes | Get available Azure VM sizes
[**GetCloudProvidersGcpDiskTypes**](CloudProvidersAPI.md#GetCloudProvidersGcpDiskTypes) | **Get** /cloud-providers/gcp/disk-types | Get available GCP disk types
[**GetCloudProvidersGcpImages**](CloudProvidersAPI.md#GetCloudProvidersGcpImages) | **Get** /cloud-providers/gcp/images | Get available GCP images
[**GetCloudProvidersGcpMachineTypes**](CloudProvidersAPI.md#GetCloudProvidersGcpMachineTypes) | **Get** /cloud-providers/gcp/machine-types | Get available GCP machine types
[**GetCloudProvidersGcpZones**](CloudProvidersAPI.md#GetCloudProvidersGcpZones) | **Get** /cloud-providers/gcp/zones | Get available GCP zones



## GetCloudProvidersAwsAmi

> map[string]interface{} GetCloudProvidersAwsAmi(ctx).Region(region).Execute()

Get available AWS AMIs for a region



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
	region := "region_example" // string | AWS region (default: us-east-1) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersAwsAmi(context.Background()).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersAwsAmi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersAwsAmi`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersAwsAmi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersAwsAmiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | AWS region (default: us-east-1) | 

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


## GetCloudProvidersAwsInstanceTypes

> map[string]interface{} GetCloudProvidersAwsInstanceTypes(ctx).Execute()

Get available AWS instance types



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
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersAwsInstanceTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersAwsInstanceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersAwsInstanceTypes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersAwsInstanceTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersAwsInstanceTypesRequest struct via the builder pattern


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


## GetCloudProvidersAwsKeyPairs

> map[string]interface{} GetCloudProvidersAwsKeyPairs(ctx).CredentialId(credentialId).Region(region).Execute()

Get AWS key pairs for credentials



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
	credentialId := int32(56) // int32 | Cloud Provider Credential ID
	region := "region_example" // string | AWS region

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersAwsKeyPairs(context.Background()).CredentialId(credentialId).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersAwsKeyPairs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersAwsKeyPairs`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersAwsKeyPairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersAwsKeyPairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialId** | **int32** | Cloud Provider Credential ID | 
 **region** | **string** | AWS region | 

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


## GetCloudProvidersAwsRegions

> map[string]interface{} GetCloudProvidersAwsRegions(ctx).Execute()

Get available AWS regions



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
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersAwsRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersAwsRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersAwsRegions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersAwsRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersAwsRegionsRequest struct via the builder pattern


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


## GetCloudProvidersAwsSecurityGroups

> map[string]interface{} GetCloudProvidersAwsSecurityGroups(ctx).CredentialId(credentialId).Region(region).Execute()

Get AWS security groups for credentials



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
	credentialId := int32(56) // int32 | Cloud Provider Credential ID
	region := "region_example" // string | AWS region

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersAwsSecurityGroups(context.Background()).CredentialId(credentialId).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersAwsSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersAwsSecurityGroups`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersAwsSecurityGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersAwsSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialId** | **int32** | Cloud Provider Credential ID | 
 **region** | **string** | AWS region | 

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


## GetCloudProvidersAzureRegions

> map[string]interface{} GetCloudProvidersAzureRegions(ctx).Execute()

Get available Azure regions



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
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersAzureRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersAzureRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersAzureRegions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersAzureRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersAzureRegionsRequest struct via the builder pattern


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


## GetCloudProvidersAzureVmSizes

> map[string]interface{} GetCloudProvidersAzureVmSizes(ctx).Execute()

Get available Azure VM sizes



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
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersAzureVmSizes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersAzureVmSizes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersAzureVmSizes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersAzureVmSizes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersAzureVmSizesRequest struct via the builder pattern


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


## GetCloudProvidersGcpDiskTypes

> map[string]interface{} GetCloudProvidersGcpDiskTypes(ctx).Execute()

Get available GCP disk types



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
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersGcpDiskTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersGcpDiskTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersGcpDiskTypes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersGcpDiskTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersGcpDiskTypesRequest struct via the builder pattern


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


## GetCloudProvidersGcpImages

> map[string]interface{} GetCloudProvidersGcpImages(ctx).Execute()

Get available GCP images



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
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersGcpImages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersGcpImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersGcpImages`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersGcpImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersGcpImagesRequest struct via the builder pattern


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


## GetCloudProvidersGcpMachineTypes

> map[string]interface{} GetCloudProvidersGcpMachineTypes(ctx).CredentialId(credentialId).Zone(zone).Execute()

Get available GCP machine types



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
	credentialId := int32(56) // int32 | Cloud Provider Credential ID
	zone := "zone_example" // string | GCP zone

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersGcpMachineTypes(context.Background()).CredentialId(credentialId).Zone(zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersGcpMachineTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersGcpMachineTypes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersGcpMachineTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersGcpMachineTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialId** | **int32** | Cloud Provider Credential ID | 
 **zone** | **string** | GCP zone | 

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


## GetCloudProvidersGcpZones

> map[string]interface{} GetCloudProvidersGcpZones(ctx).CredentialId(credentialId).Execute()

Get available GCP zones



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
	credentialId := int32(56) // int32 | Cloud Provider Credential ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudProvidersAPI.GetCloudProvidersGcpZones(context.Background()).CredentialId(credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersAPI.GetCloudProvidersGcpZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudProvidersGcpZones`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudProvidersAPI.GetCloudProvidersGcpZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudProvidersGcpZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialId** | **int32** | Cloud Provider Credential ID | 

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

