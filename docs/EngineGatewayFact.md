# EngineGatewayFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavor** | Pointer to **string** | Flavor is the gateway flavor the cluster chose ONCE (v1 vocabulary: caddy). The choice is sticky per cluster (EDGE.md §4 rule 3) — v1 renders caddy only; a migration is a documented re-deploy. | [optional] 
**PublishAddress** | Pointer to **string** | PublishAddress is the address the gateway&#39;s edge publishes: an IP (A-record form) or a hostname (CNAME form — tunnel topologies). The record type the dns steps use FOLLOWS this form. | [optional] 

## Methods

### NewEngineGatewayFact

`func NewEngineGatewayFact() *EngineGatewayFact`

NewEngineGatewayFact instantiates a new EngineGatewayFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineGatewayFactWithDefaults

`func NewEngineGatewayFactWithDefaults() *EngineGatewayFact`

NewEngineGatewayFactWithDefaults instantiates a new EngineGatewayFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavor

`func (o *EngineGatewayFact) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *EngineGatewayFact) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *EngineGatewayFact) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *EngineGatewayFact) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetPublishAddress

`func (o *EngineGatewayFact) GetPublishAddress() string`

GetPublishAddress returns the PublishAddress field if non-nil, zero value otherwise.

### GetPublishAddressOk

`func (o *EngineGatewayFact) GetPublishAddressOk() (*string, bool)`

GetPublishAddressOk returns a tuple with the PublishAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishAddress

`func (o *EngineGatewayFact) SetPublishAddress(v string)`

SetPublishAddress sets PublishAddress field to given value.

### HasPublishAddress

`func (o *EngineGatewayFact) HasPublishAddress() bool`

HasPublishAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


