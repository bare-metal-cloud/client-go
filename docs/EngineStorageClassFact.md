# EngineStorageClassFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the class name (\&quot;local\&quot; | \&quot;local-cow\&quot; | \&quot;lvm-thin\&quot; | \&quot;longhorn-3rep\&quot;). | [optional] 
**Replication** | Pointer to **int32** | Replication is the class&#39;s replica factor (1 | 3); replicated classes are refused below vendor minimums (STORAGE.md §7). | [optional] 
**SnapshotLayer** | Pointer to **bool** | SnapshotLayer marks CoW snapshot capability — required for continuous FILE sync. | [optional] 

## Methods

### NewEngineStorageClassFact

`func NewEngineStorageClassFact() *EngineStorageClassFact`

NewEngineStorageClassFact instantiates a new EngineStorageClassFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineStorageClassFactWithDefaults

`func NewEngineStorageClassFactWithDefaults() *EngineStorageClassFact`

NewEngineStorageClassFactWithDefaults instantiates a new EngineStorageClassFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EngineStorageClassFact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineStorageClassFact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineStorageClassFact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EngineStorageClassFact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplication

`func (o *EngineStorageClassFact) GetReplication() int32`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *EngineStorageClassFact) GetReplicationOk() (*int32, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *EngineStorageClassFact) SetReplication(v int32)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *EngineStorageClassFact) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetSnapshotLayer

`func (o *EngineStorageClassFact) GetSnapshotLayer() bool`

GetSnapshotLayer returns the SnapshotLayer field if non-nil, zero value otherwise.

### GetSnapshotLayerOk

`func (o *EngineStorageClassFact) GetSnapshotLayerOk() (*bool, bool)`

GetSnapshotLayerOk returns a tuple with the SnapshotLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotLayer

`func (o *EngineStorageClassFact) SetSnapshotLayer(v bool)`

SetSnapshotLayer sets SnapshotLayer field to given value.

### HasSnapshotLayer

`func (o *EngineStorageClassFact) HasSnapshotLayer() bool`

HasSnapshotLayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


