# \RunnersProfilesAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersProfiles**](RunnersProfilesAPI.md#GetOrgsOrgIdRunnersProfiles) | **Get** /orgs/{org_id}/runners/profiles | List the org&#39;s runner profiles — the seeded set, with each profile&#39;s eligible machine count
[**PostOrgsOrgIdRunnersProfiles**](RunnersProfilesAPI.md#PostOrgsOrgIdRunnersProfiles) | **Post** /orgs/{org_id}/runners/profiles | Raw profile creation — typed refusal (the profile set is fixed in v1)



## GetOrgsOrgIdRunnersProfiles

> HandlersRunnerProfilesResponse GetOrgsOrgIdRunnersProfiles(ctx, orgId).Execute()

List the org's runner profiles — the seeded set, with each profile's eligible machine count



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
	resp, r, err := apiClient.RunnersProfilesAPI.GetOrgsOrgIdRunnersProfiles(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersProfilesAPI.GetOrgsOrgIdRunnersProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersProfiles`: HandlersRunnerProfilesResponse
	fmt.Fprintf(os.Stdout, "Response from `RunnersProfilesAPI.GetOrgsOrgIdRunnersProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerProfilesResponse**](HandlersRunnerProfilesResponse.md)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgsOrgIdRunnersProfiles

> PostOrgsOrgIdRunnersProfiles(ctx, orgId).Execute()

Raw profile creation — typed refusal (the profile set is fixed in v1)



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
	r, err := apiClient.RunnersProfilesAPI.PostOrgsOrgIdRunnersProfiles(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersProfilesAPI.PostOrgsOrgIdRunnersProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgsOrgIdRunnersProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CookieAuth](../README.md#CookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

