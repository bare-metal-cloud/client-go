# EngineArrivalFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**EngineArrivalHealth**](EngineArrivalHealth.md) | Health is the arrival&#39;s health (the frozen ArrivalHealth vocabulary). Health is defined per shape: a SaaS connector arrival reports its connector deployment&#39;s ready replicas equal to desired; an API-less arrival reports a reachability probe of its far end. | [optional] 
**Instance** | Pointer to **string** | Instance is the arrival addon instance id (matches the route binding&#39;s ingress.instance). | [optional] 
**Kind** | Pointer to [**BundleArrivalKind**](BundleArrivalKind.md) | Kind is the arrival family (the bundle.ArrivalKind vocabulary — one vocabulary across the route binding and the facts). | [optional] 
**Provider** | Pointer to **string** | Provider names the arrival provider behind the instance (story 154.6 — the compile&#39;s shape stamp reads it: the step sequence, the reservation slot, and the co-tenant plan warning are per-provider declarations, and the facts are where the LIVE provider is named). Additive omitempty: pre-154.6 facts marshal and parse unchanged, and the compile treats a missing provider as the remote-config posture (cloudflare&#39;s landed behavior) — never a guess. | [optional] 
**PublishAddress** | Pointer to **string** | PublishAddress is the address traffic arrives at: the gateway&#39;s publish address for public; the full arrival target for a tunnel (the tunnel CNAME for cloudflare — including the tunnel id, never a domain suffix). The compile targets THIS address for a tunnel-arrival-bound route, never the gateway publish. | [optional] 

## Methods

### NewEngineArrivalFact

`func NewEngineArrivalFact() *EngineArrivalFact`

NewEngineArrivalFact instantiates a new EngineArrivalFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineArrivalFactWithDefaults

`func NewEngineArrivalFactWithDefaults() *EngineArrivalFact`

NewEngineArrivalFactWithDefaults instantiates a new EngineArrivalFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *EngineArrivalFact) GetHealth() EngineArrivalHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *EngineArrivalFact) GetHealthOk() (*EngineArrivalHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *EngineArrivalFact) SetHealth(v EngineArrivalHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *EngineArrivalFact) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetInstance

`func (o *EngineArrivalFact) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *EngineArrivalFact) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *EngineArrivalFact) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *EngineArrivalFact) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetKind

`func (o *EngineArrivalFact) GetKind() BundleArrivalKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EngineArrivalFact) GetKindOk() (*BundleArrivalKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EngineArrivalFact) SetKind(v BundleArrivalKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EngineArrivalFact) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProvider

`func (o *EngineArrivalFact) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EngineArrivalFact) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EngineArrivalFact) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EngineArrivalFact) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublishAddress

`func (o *EngineArrivalFact) GetPublishAddress() string`

GetPublishAddress returns the PublishAddress field if non-nil, zero value otherwise.

### GetPublishAddressOk

`func (o *EngineArrivalFact) GetPublishAddressOk() (*string, bool)`

GetPublishAddressOk returns a tuple with the PublishAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishAddress

`func (o *EngineArrivalFact) SetPublishAddress(v string)`

SetPublishAddress sets PublishAddress field to given value.

### HasPublishAddress

`func (o *EngineArrivalFact) HasPublishAddress() bool`

HasPublishAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


