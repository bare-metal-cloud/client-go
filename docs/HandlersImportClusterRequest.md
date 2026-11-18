# HandlersImportClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdoptIngress** | Pointer to **bool** | AdoptIngress is the explicit consent flag (story 133.3, GH #2200 AC 2, optional): when the target cluster runs a live traefik holding ingress, importing requires the owner&#39;s explicit acceptance that the platform&#39;s caddy gateway addon takes ingress over (traefik stops serving it). Without the flag a traefik-holding cluster REFUSES the import with named, guided alternatives — never a silent ingress takeover. Honorable only while the adoption is lossless (traefik serves no Traefik-native CRD routes the addon cannot replace). | [optional] 
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

### GetAdoptIngress

`func (o *HandlersImportClusterRequest) GetAdoptIngress() bool`

GetAdoptIngress returns the AdoptIngress field if non-nil, zero value otherwise.

### GetAdoptIngressOk

`func (o *HandlersImportClusterRequest) GetAdoptIngressOk() (*bool, bool)`

GetAdoptIngressOk returns a tuple with the AdoptIngress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptIngress

`func (o *HandlersImportClusterRequest) SetAdoptIngress(v bool)`

SetAdoptIngress sets AdoptIngress field to given value.

### HasAdoptIngress

`func (o *HandlersImportClusterRequest) HasAdoptIngress() bool`

HasAdoptIngress returns a boolean if a field has been set.

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


