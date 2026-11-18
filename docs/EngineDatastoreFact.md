# EngineDatastoreFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to [**EngineDatastoreKind**](EngineDatastoreKind.md) | Kind is the datastore kind (the frozen DatastoreKind vocabulary). | [optional] 
**ServerFlags** | Pointer to **[]string** | ServerFlags are the k3s server flags as the platform observed them, canonicalized by NormalizeServerFlags (sorted \&quot;flag&#x3D;value\&quot; pairs) — the flag-match preflight compares the observed set against this one and refuses drift. | [optional] 
**Servers** | Pointer to **[]string** | Servers are the cluster&#39;s server (control-plane) node names — the quorum fact the conversion preflight reads. | [optional] 

## Methods

### NewEngineDatastoreFact

`func NewEngineDatastoreFact() *EngineDatastoreFact`

NewEngineDatastoreFact instantiates a new EngineDatastoreFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineDatastoreFactWithDefaults

`func NewEngineDatastoreFactWithDefaults() *EngineDatastoreFact`

NewEngineDatastoreFactWithDefaults instantiates a new EngineDatastoreFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *EngineDatastoreFact) GetKind() EngineDatastoreKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EngineDatastoreFact) GetKindOk() (*EngineDatastoreKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EngineDatastoreFact) SetKind(v EngineDatastoreKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EngineDatastoreFact) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetServerFlags

`func (o *EngineDatastoreFact) GetServerFlags() []string`

GetServerFlags returns the ServerFlags field if non-nil, zero value otherwise.

### GetServerFlagsOk

`func (o *EngineDatastoreFact) GetServerFlagsOk() (*[]string, bool)`

GetServerFlagsOk returns a tuple with the ServerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFlags

`func (o *EngineDatastoreFact) SetServerFlags(v []string)`

SetServerFlags sets ServerFlags field to given value.

### HasServerFlags

`func (o *EngineDatastoreFact) HasServerFlags() bool`

HasServerFlags returns a boolean if a field has been set.

### GetServers

`func (o *EngineDatastoreFact) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *EngineDatastoreFact) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *EngineDatastoreFact) SetServers(v []string)`

SetServers sets Servers field to given value.

### HasServers

`func (o *EngineDatastoreFact) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


