# \RunnersMachinePolicyAPI

All URIs are relative to *http://localhost:3001/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgsOrgIdRunnersMachinePolicy**](RunnersMachinePolicyAPI.md#GetOrgsOrgIdRunnersMachinePolicy) | **Get** /orgs/{org_id}/runners/machine-policy | Get the org&#39;s machine policy
[**PutOrgsOrgIdRunnersMachinePolicy**](RunnersMachinePolicyAPI.md#PutOrgsOrgIdRunnersMachinePolicy) | **Put** /orgs/{org_id}/runners/machine-policy | Set the org&#39;s machine policy (admin)



## GetOrgsOrgIdRunnersMachinePolicy

> HandlersRunnerMachinePolicyBody GetOrgsOrgIdRunnersMachinePolicy(ctx, orgId).Execute()

Get the org's machine policy



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
	resp, r, err := apiClient.RunnersMachinePolicyAPI.GetOrgsOrgIdRunnersMachinePolicy(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinePolicyAPI.GetOrgsOrgIdRunnersMachinePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgsOrgIdRunnersMachinePolicy`: HandlersRunnerMachinePolicyBody
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinePolicyAPI.GetOrgsOrgIdRunnersMachinePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgsOrgIdRunnersMachinePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlersRunnerMachinePolicyBody**](HandlersRunnerMachinePolicyBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgsOrgIdRunnersMachinePolicy

> HandlersRunnerMachinePolicyBody PutOrgsOrgIdRunnersMachinePolicy(ctx, orgId).Body(body).Execute()

Set the org's machine policy (admin)



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
	body := *openapiclient.NewHandlersRunnerMachinePolicyBody() // HandlersRunnerMachinePolicyBody | the five controls

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersMachinePolicyAPI.PutOrgsOrgIdRunnersMachinePolicy(context.Background(), orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersMachinePolicyAPI.PutOrgsOrgIdRunnersMachinePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrgsOrgIdRunnersMachinePolicy`: HandlersRunnerMachinePolicyBody
	fmt.Fprintf(os.Stdout, "Response from `RunnersMachinePolicyAPI.PutOrgsOrgIdRunnersMachinePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **int32** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgsOrgIdRunnersMachinePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlersRunnerMachinePolicyBody**](HandlersRunnerMachinePolicyBody.md) | the five controls | 

### Return type

[**HandlersRunnerMachinePolicyBody**](HandlersRunnerMachinePolicyBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

