# HandlersDeleteClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgeIrreversible** | Pointer to **bool** | AcknowledgeIrreversible is the AC5 confirm-again flag (the server&#39;s irreversibility_warning 409 carries the verbatim text; re-submit with this true). | [optional] 
**ImpactFingerprint** | Pointer to **string** | ImpactFingerprint is the epic-#3964 story-C1 execute-time re-verification handle: the fingerprint GET /clusters/:id/delete-impact returned. When provided and the consequences moved since that preflight, the delete refuses with 409 consequence_changed carrying the fresh impact. Absent &#x3D; the legacy contract, unchanged. | [optional] 
**VmDisposition** | Pointer to **string** |  | [optional] 
**VmDispositions** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewHandlersDeleteClusterRequest

`func NewHandlersDeleteClusterRequest() *HandlersDeleteClusterRequest`

NewHandlersDeleteClusterRequest instantiates a new HandlersDeleteClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersDeleteClusterRequestWithDefaults

`func NewHandlersDeleteClusterRequestWithDefaults() *HandlersDeleteClusterRequest`

NewHandlersDeleteClusterRequestWithDefaults instantiates a new HandlersDeleteClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgeIrreversible

`func (o *HandlersDeleteClusterRequest) GetAcknowledgeIrreversible() bool`

GetAcknowledgeIrreversible returns the AcknowledgeIrreversible field if non-nil, zero value otherwise.

### GetAcknowledgeIrreversibleOk

`func (o *HandlersDeleteClusterRequest) GetAcknowledgeIrreversibleOk() (*bool, bool)`

GetAcknowledgeIrreversibleOk returns a tuple with the AcknowledgeIrreversible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeIrreversible

`func (o *HandlersDeleteClusterRequest) SetAcknowledgeIrreversible(v bool)`

SetAcknowledgeIrreversible sets AcknowledgeIrreversible field to given value.

### HasAcknowledgeIrreversible

`func (o *HandlersDeleteClusterRequest) HasAcknowledgeIrreversible() bool`

HasAcknowledgeIrreversible returns a boolean if a field has been set.

### GetImpactFingerprint

`func (o *HandlersDeleteClusterRequest) GetImpactFingerprint() string`

GetImpactFingerprint returns the ImpactFingerprint field if non-nil, zero value otherwise.

### GetImpactFingerprintOk

`func (o *HandlersDeleteClusterRequest) GetImpactFingerprintOk() (*string, bool)`

GetImpactFingerprintOk returns a tuple with the ImpactFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactFingerprint

`func (o *HandlersDeleteClusterRequest) SetImpactFingerprint(v string)`

SetImpactFingerprint sets ImpactFingerprint field to given value.

### HasImpactFingerprint

`func (o *HandlersDeleteClusterRequest) HasImpactFingerprint() bool`

HasImpactFingerprint returns a boolean if a field has been set.

### GetVmDisposition

`func (o *HandlersDeleteClusterRequest) GetVmDisposition() string`

GetVmDisposition returns the VmDisposition field if non-nil, zero value otherwise.

### GetVmDispositionOk

`func (o *HandlersDeleteClusterRequest) GetVmDispositionOk() (*string, bool)`

GetVmDispositionOk returns a tuple with the VmDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmDisposition

`func (o *HandlersDeleteClusterRequest) SetVmDisposition(v string)`

SetVmDisposition sets VmDisposition field to given value.

### HasVmDisposition

`func (o *HandlersDeleteClusterRequest) HasVmDisposition() bool`

HasVmDisposition returns a boolean if a field has been set.

### GetVmDispositions

`func (o *HandlersDeleteClusterRequest) GetVmDispositions() map[string]string`

GetVmDispositions returns the VmDispositions field if non-nil, zero value otherwise.

### GetVmDispositionsOk

`func (o *HandlersDeleteClusterRequest) GetVmDispositionsOk() (*map[string]string, bool)`

GetVmDispositionsOk returns a tuple with the VmDispositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmDispositions

`func (o *HandlersDeleteClusterRequest) SetVmDispositions(v map[string]string)`

SetVmDispositions sets VmDispositions field to given value.

### HasVmDispositions

`func (o *HandlersDeleteClusterRequest) HasVmDispositions() bool`

HasVmDispositions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


