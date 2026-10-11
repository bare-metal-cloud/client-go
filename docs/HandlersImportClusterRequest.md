# HandlersImportClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kubeconfig** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewHandlersImportClusterRequest

`func NewHandlersImportClusterRequest(name string, ) *HandlersImportClusterRequest`

NewHandlersImportClusterRequest instantiates a new HandlersImportClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersImportClusterRequestWithDefaults

`func NewHandlersImportClusterRequestWithDefaults() *HandlersImportClusterRequest`

NewHandlersImportClusterRequestWithDefaults instantiates a new HandlersImportClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeconfig

`func (o *HandlersImportClusterRequest) GetKubeconfig() string`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *HandlersImportClusterRequest) GetKubeconfigOk() (*string, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *HandlersImportClusterRequest) SetKubeconfig(v string)`

SetKubeconfig sets Kubeconfig field to given value.

### HasKubeconfig

`func (o *HandlersImportClusterRequest) HasKubeconfig() bool`

HasKubeconfig returns a boolean if a field has been set.

### GetName

`func (o *HandlersImportClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlersImportClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlersImportClusterRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


