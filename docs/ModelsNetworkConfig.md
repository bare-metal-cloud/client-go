# ModelsNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastNodeLeftAt** | Pointer to **string** | LastNodeLeftAt is stamped by the story 124.14 membership-free write when the write removed a network&#39;s LAST network_nodes row and nothing (no node, no cluster binding) remains — the candidate marker the network reconciler&#39;s last-node reap pass (ReapLastNodeLeftNetworks) tears the network down from, after a grace window. Stored INSIDE the Config JSONB on purpose: the column is already JSONB, so the new key needs NO schema migration and legacy rows simply read it as nil. | [optional] 
**MeshConfig** | Pointer to [**ModelsMeshNetworkConfig**](ModelsMeshNetworkConfig.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ModelsNetworkType**](ModelsNetworkType.md) |  | [optional] 

## Methods

### NewModelsNetworkConfig

`func NewModelsNetworkConfig() *ModelsNetworkConfig`

NewModelsNetworkConfig instantiates a new ModelsNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsNetworkConfigWithDefaults

`func NewModelsNetworkConfigWithDefaults() *ModelsNetworkConfig`

NewModelsNetworkConfigWithDefaults instantiates a new ModelsNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastNodeLeftAt

`func (o *ModelsNetworkConfig) GetLastNodeLeftAt() string`

GetLastNodeLeftAt returns the LastNodeLeftAt field if non-nil, zero value otherwise.

### GetLastNodeLeftAtOk

`func (o *ModelsNetworkConfig) GetLastNodeLeftAtOk() (*string, bool)`

GetLastNodeLeftAtOk returns a tuple with the LastNodeLeftAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNodeLeftAt

`func (o *ModelsNetworkConfig) SetLastNodeLeftAt(v string)`

SetLastNodeLeftAt sets LastNodeLeftAt field to given value.

### HasLastNodeLeftAt

`func (o *ModelsNetworkConfig) HasLastNodeLeftAt() bool`

HasLastNodeLeftAt returns a boolean if a field has been set.

### GetMeshConfig

`func (o *ModelsNetworkConfig) GetMeshConfig() ModelsMeshNetworkConfig`

GetMeshConfig returns the MeshConfig field if non-nil, zero value otherwise.

### GetMeshConfigOk

`func (o *ModelsNetworkConfig) GetMeshConfigOk() (*ModelsMeshNetworkConfig, bool)`

GetMeshConfigOk returns a tuple with the MeshConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshConfig

`func (o *ModelsNetworkConfig) SetMeshConfig(v ModelsMeshNetworkConfig)`

SetMeshConfig sets MeshConfig field to given value.

### HasMeshConfig

`func (o *ModelsNetworkConfig) HasMeshConfig() bool`

HasMeshConfig returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsNetworkConfig) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsNetworkConfig) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsNetworkConfig) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsNetworkConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ModelsNetworkConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsNetworkConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsNetworkConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsNetworkConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ModelsNetworkConfig) GetType() ModelsNetworkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelsNetworkConfig) GetTypeOk() (*ModelsNetworkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelsNetworkConfig) SetType(v ModelsNetworkType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelsNetworkConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


