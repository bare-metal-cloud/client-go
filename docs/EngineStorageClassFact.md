# EngineStorageClassFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowVolumeExpansion** | Pointer to **bool** | AllowVolumeExpansion is the class&#39;s allowVolumeExpansion flag as the platform observed it on the StorageClass object (the dk64 ledger&#39;s standing interpretation). POINTER on purpose: an omitempty plain bool drops false on the wire, making capability-FALSE wire-indistinguishable from UNKNOWN — nil &#x3D; unknown (collected before the flag existed, or no live read backed it), false &#x3D; refused capability, true &#x3D; expandable. The restore gate&#39;s phase-0 live read overwrites a carried nil before any consumer runs (an absent field on the live object maps to false, never pointer-nil), so no gate phase ever sees UNKNOWN from a live-verified path; a compile-side refusal on nil applies only where no live read exists, and the member&#39;s ABSENCE from the payload (not a nil flag inside a present fact) is the MISSING-FACT case. Additive field (omitempty): facts written before dk64 Story A decode with it nil and re-marshal byte-identically. | [optional] 
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

### GetAllowVolumeExpansion

`func (o *EngineStorageClassFact) GetAllowVolumeExpansion() bool`

GetAllowVolumeExpansion returns the AllowVolumeExpansion field if non-nil, zero value otherwise.

### GetAllowVolumeExpansionOk

`func (o *EngineStorageClassFact) GetAllowVolumeExpansionOk() (*bool, bool)`

GetAllowVolumeExpansionOk returns a tuple with the AllowVolumeExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVolumeExpansion

`func (o *EngineStorageClassFact) SetAllowVolumeExpansion(v bool)`

SetAllowVolumeExpansion sets AllowVolumeExpansion field to given value.

### HasAllowVolumeExpansion

`func (o *EngineStorageClassFact) HasAllowVolumeExpansion() bool`

HasAllowVolumeExpansion returns a boolean if a field has been set.

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


