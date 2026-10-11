# HandlersInstallCommandPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Required for K8s (validated in handler) | [optional] 
**ClusterType** | **string** | Required: \&quot;kubernetes\&quot;, \&quot;bare-metal\&quot;, or \&quot;imported\&quot; | 
**Provider** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**VmName** | Pointer to **string** | For VM installations (only valid with bare-metal/imported) | [optional] 

## Methods

### NewHandlersInstallCommandPostRequest

`func NewHandlersInstallCommandPostRequest(clusterType string, ) *HandlersInstallCommandPostRequest`

NewHandlersInstallCommandPostRequest instantiates a new HandlersInstallCommandPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersInstallCommandPostRequestWithDefaults

`func NewHandlersInstallCommandPostRequestWithDefaults() *HandlersInstallCommandPostRequest`

NewHandlersInstallCommandPostRequestWithDefaults instantiates a new HandlersInstallCommandPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *HandlersInstallCommandPostRequest) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HandlersInstallCommandPostRequest) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HandlersInstallCommandPostRequest) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HandlersInstallCommandPostRequest) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetClusterType

`func (o *HandlersInstallCommandPostRequest) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *HandlersInstallCommandPostRequest) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *HandlersInstallCommandPostRequest) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.


### GetProvider

`func (o *HandlersInstallCommandPostRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *HandlersInstallCommandPostRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *HandlersInstallCommandPostRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *HandlersInstallCommandPostRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *HandlersInstallCommandPostRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HandlersInstallCommandPostRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HandlersInstallCommandPostRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HandlersInstallCommandPostRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVmName

`func (o *HandlersInstallCommandPostRequest) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *HandlersInstallCommandPostRequest) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *HandlersInstallCommandPostRequest) SetVmName(v string)`

SetVmName sets VmName field to given value.

### HasVmName

`func (o *HandlersInstallCommandPostRequest) HasVmName() bool`

HasVmName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


