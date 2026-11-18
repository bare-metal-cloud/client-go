# \ExploreAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAdminCustomersExplore**](ExploreAPI.md#PostAdminCustomersExplore) | **Post** /admin/customers/explore | Open an explore-as session (platform admin)
[**PostExploreDeElevate**](ExploreAPI.md#PostExploreDeElevate) | **Post** /explore/de-elevate | Drop an explore session back to Observe
[**PostExploreElevate**](ExploreAPI.md#PostExploreElevate) | **Post** /explore/elevate | Elevate an explore session to Operate
[**PostExploreExit**](ExploreAPI.md#PostExploreExit) | **Post** /explore/exit | Exit an explore session



## PostAdminCustomersExplore

> map[string]interface{} PostAdminCustomersExplore(ctx).Body(body).Execute()

Open an explore-as session (platform admin)



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
	body := map[string]interface{}{ ... } // map[string]interface{} | org_id and subject_user_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExploreAPI.PostAdminCustomersExplore(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExploreAPI.PostAdminCustomersExplore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdminCustomersExplore`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ExploreAPI.PostAdminCustomersExplore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAdminCustomersExploreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | org_id and subject_user_id | 

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


## PostExploreDeElevate

> map[string]interface{} PostExploreDeElevate(ctx).XExploreToken(xExploreToken).Execute()

Drop an explore session back to Observe



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
	xExploreToken := "xExploreToken_example" // string | Explore record token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExploreAPI.PostExploreDeElevate(context.Background()).XExploreToken(xExploreToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExploreAPI.PostExploreDeElevate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExploreDeElevate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ExploreAPI.PostExploreDeElevate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExploreDeElevateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xExploreToken** | **string** | Explore record token | 

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


## PostExploreElevate

> map[string]interface{} PostExploreElevate(ctx).XExploreToken(xExploreToken).Body(body).Execute()

Elevate an explore session to Operate



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
	xExploreToken := "xExploreToken_example" // string | Explore record token
	body := map[string]interface{}{ ... } // map[string]interface{} | reason and totp_code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExploreAPI.PostExploreElevate(context.Background()).XExploreToken(xExploreToken).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExploreAPI.PostExploreElevate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExploreElevate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ExploreAPI.PostExploreElevate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExploreElevateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xExploreToken** | **string** | Explore record token | 
 **body** | **map[string]interface{}** | reason and totp_code | 

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


## PostExploreExit

> map[string]interface{} PostExploreExit(ctx).XExploreToken(xExploreToken).Execute()

Exit an explore session



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
	xExploreToken := "xExploreToken_example" // string | Explore record token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExploreAPI.PostExploreExit(context.Background()).XExploreToken(xExploreToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExploreAPI.PostExploreExit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExploreExit`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ExploreAPI.PostExploreExit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExploreExitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xExploreToken** | **string** | Explore record token | 

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

