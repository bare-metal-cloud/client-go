# EngineContextFacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | Arch is the node architecture: amd64 | arm64. | [optional] 
**ClaimedHosts** | Pointer to [**[]EngineHostClaim**](EngineHostClaim.md) | ClaimedHosts are the route hostnames other installs on this cluster already claim — the fact the hostname-uniqueness preflight (EDGE.md §4 rule 1) checks route hosts against; a collision refuses naming the conflicting install. The HOST keeps this registry authoritative (record on provision, release on uninstall); omitempty keeps legacy facts byte-identical. | [optional] 
**ClusterSize** | Pointer to **int32** | ClusterSize is the node count. | [optional] 
**DataDiskFs** | Pointer to [**EngineDataDiskFS**](EngineDataDiskFS.md) | DataDiskFS is the host data-disk filesystem (the frozen DataDiskFS vocabulary). Continuous FILE sync requires a CoW snapshot layer — on plain ext4 the resolve-time check refuses it (STORAGE.md §3). | [optional] 
**Destinations** | Pointer to **[]string** | Destinations are the backup destinations the platform&#39;s destination registry offers as viable for this install (user&#39;s buckets, probed reachable). A durability.destination outside this set is refused at resolve (ErrCapabilityUnavailable). | [optional] 
**Domains** | Pointer to **[]string** | Domains are the manageable zones for env_local hostnames. | [optional] 
**Gateway** | Pointer to [**EngineGatewayFact**](EngineGatewayFact.md) | Gateway is the cluster&#39;s edge-gateway publish fact (story 130.2, EDGE.md §2) — the flavor the cluster runs and the address its edge publishes. The web-class compile targets it instead of guessing: the dns-apply instructions point here, dns-verify expects it, and route-apply renders for its flavor. Nil when the cluster has no gateway addon yet — a web compile refuses without it. omitempty so pre-130.2 persisted facts re-marshal byte-identically. | [optional] 
**NodeTier** | Pointer to [**EngineNodeTier**](EngineNodeTier.md) | NodeTier is the node&#39;s memory tier (the frozen NodeTier vocabulary). The 1gb tier narrows durability to off/scheduled only (STORAGE.md §7). | [optional] 
**StorageClasses** | Pointer to [**[]EngineStorageClassFact**](EngineStorageClassFact.md) | StorageClasses are the storage classes the platform declared on this cluster (STORAGE.md §7 tiering is compiled FROM these). A storage.class value outside this set is refused at resolve (ErrCapabilityUnavailable). | [optional] 

## Methods

### NewEngineContextFacts

`func NewEngineContextFacts() *EngineContextFacts`

NewEngineContextFacts instantiates a new EngineContextFacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineContextFactsWithDefaults

`func NewEngineContextFactsWithDefaults() *EngineContextFacts`

NewEngineContextFactsWithDefaults instantiates a new EngineContextFacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *EngineContextFacts) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *EngineContextFacts) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *EngineContextFacts) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *EngineContextFacts) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetClaimedHosts

`func (o *EngineContextFacts) GetClaimedHosts() []EngineHostClaim`

GetClaimedHosts returns the ClaimedHosts field if non-nil, zero value otherwise.

### GetClaimedHostsOk

`func (o *EngineContextFacts) GetClaimedHostsOk() (*[]EngineHostClaim, bool)`

GetClaimedHostsOk returns a tuple with the ClaimedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedHosts

`func (o *EngineContextFacts) SetClaimedHosts(v []EngineHostClaim)`

SetClaimedHosts sets ClaimedHosts field to given value.

### HasClaimedHosts

`func (o *EngineContextFacts) HasClaimedHosts() bool`

HasClaimedHosts returns a boolean if a field has been set.

### GetClusterSize

`func (o *EngineContextFacts) GetClusterSize() int32`

GetClusterSize returns the ClusterSize field if non-nil, zero value otherwise.

### GetClusterSizeOk

`func (o *EngineContextFacts) GetClusterSizeOk() (*int32, bool)`

GetClusterSizeOk returns a tuple with the ClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSize

`func (o *EngineContextFacts) SetClusterSize(v int32)`

SetClusterSize sets ClusterSize field to given value.

### HasClusterSize

`func (o *EngineContextFacts) HasClusterSize() bool`

HasClusterSize returns a boolean if a field has been set.

### GetDataDiskFs

`func (o *EngineContextFacts) GetDataDiskFs() EngineDataDiskFS`

GetDataDiskFs returns the DataDiskFs field if non-nil, zero value otherwise.

### GetDataDiskFsOk

`func (o *EngineContextFacts) GetDataDiskFsOk() (*EngineDataDiskFS, bool)`

GetDataDiskFsOk returns a tuple with the DataDiskFs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskFs

`func (o *EngineContextFacts) SetDataDiskFs(v EngineDataDiskFS)`

SetDataDiskFs sets DataDiskFs field to given value.

### HasDataDiskFs

`func (o *EngineContextFacts) HasDataDiskFs() bool`

HasDataDiskFs returns a boolean if a field has been set.

### GetDestinations

`func (o *EngineContextFacts) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *EngineContextFacts) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *EngineContextFacts) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *EngineContextFacts) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetDomains

`func (o *EngineContextFacts) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *EngineContextFacts) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *EngineContextFacts) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *EngineContextFacts) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetGateway

`func (o *EngineContextFacts) GetGateway() EngineGatewayFact`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *EngineContextFacts) GetGatewayOk() (*EngineGatewayFact, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *EngineContextFacts) SetGateway(v EngineGatewayFact)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *EngineContextFacts) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNodeTier

`func (o *EngineContextFacts) GetNodeTier() EngineNodeTier`

GetNodeTier returns the NodeTier field if non-nil, zero value otherwise.

### GetNodeTierOk

`func (o *EngineContextFacts) GetNodeTierOk() (*EngineNodeTier, bool)`

GetNodeTierOk returns a tuple with the NodeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTier

`func (o *EngineContextFacts) SetNodeTier(v EngineNodeTier)`

SetNodeTier sets NodeTier field to given value.

### HasNodeTier

`func (o *EngineContextFacts) HasNodeTier() bool`

HasNodeTier returns a boolean if a field has been set.

### GetStorageClasses

`func (o *EngineContextFacts) GetStorageClasses() []EngineStorageClassFact`

GetStorageClasses returns the StorageClasses field if non-nil, zero value otherwise.

### GetStorageClassesOk

`func (o *EngineContextFacts) GetStorageClassesOk() (*[]EngineStorageClassFact, bool)`

GetStorageClassesOk returns a tuple with the StorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClasses

`func (o *EngineContextFacts) SetStorageClasses(v []EngineStorageClassFact)`

SetStorageClasses sets StorageClasses field to given value.

### HasStorageClasses

`func (o *EngineContextFacts) HasStorageClasses() bool`

HasStorageClasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


