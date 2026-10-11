# NetworkbackendNetworkCandidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NetworkId** | Pointer to **int32** |  | [optional] 
**OnThisNetwork** | Pointer to **bool** | OnThisNetwork badges selected-network members (\&quot;on N1\&quot; chip). False for networkless candidates the wizard will PRE-JOIN. | [optional] 
**Rejection** | Pointer to [**VmpreflightRejection**](VmpreflightRejection.md) | Rejection is the 124.3 server-computed preflight verdict for this VM — the machine contract the picker cards render inline (disabled-with- reason; the contamination card keys Adoptable). Nil/omitted &#x3D; the VM passes list-level classification; the create-time gate re-verifies with a live probe (the list is a hint, the gate is the truth). | [optional] 
**SpecsUnknown** | Pointer to **bool** | SpecsUnknown mirrors 124.2&#39;s specs contract: true while the VM&#39;s health and specs are not yet agent-reported (\&quot;Waiting for agent data\&quot;). | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**VmId** | Pointer to **int32** |  | [optional] 

## Methods

### NewNetworkbackendNetworkCandidate

`func NewNetworkbackendNetworkCandidate() *NetworkbackendNetworkCandidate`

NewNetworkbackendNetworkCandidate instantiates a new NetworkbackendNetworkCandidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkbackendNetworkCandidateWithDefaults

`func NewNetworkbackendNetworkCandidateWithDefaults() *NetworkbackendNetworkCandidate`

NewNetworkbackendNetworkCandidateWithDefaults instantiates a new NetworkbackendNetworkCandidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *NetworkbackendNetworkCandidate) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NetworkbackendNetworkCandidate) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NetworkbackendNetworkCandidate) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *NetworkbackendNetworkCandidate) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetName

`func (o *NetworkbackendNetworkCandidate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkbackendNetworkCandidate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkbackendNetworkCandidate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkbackendNetworkCandidate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkId

`func (o *NetworkbackendNetworkCandidate) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *NetworkbackendNetworkCandidate) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *NetworkbackendNetworkCandidate) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *NetworkbackendNetworkCandidate) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetOnThisNetwork

`func (o *NetworkbackendNetworkCandidate) GetOnThisNetwork() bool`

GetOnThisNetwork returns the OnThisNetwork field if non-nil, zero value otherwise.

### GetOnThisNetworkOk

`func (o *NetworkbackendNetworkCandidate) GetOnThisNetworkOk() (*bool, bool)`

GetOnThisNetworkOk returns a tuple with the OnThisNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnThisNetwork

`func (o *NetworkbackendNetworkCandidate) SetOnThisNetwork(v bool)`

SetOnThisNetwork sets OnThisNetwork field to given value.

### HasOnThisNetwork

`func (o *NetworkbackendNetworkCandidate) HasOnThisNetwork() bool`

HasOnThisNetwork returns a boolean if a field has been set.

### GetRejection

`func (o *NetworkbackendNetworkCandidate) GetRejection() VmpreflightRejection`

GetRejection returns the Rejection field if non-nil, zero value otherwise.

### GetRejectionOk

`func (o *NetworkbackendNetworkCandidate) GetRejectionOk() (*VmpreflightRejection, bool)`

GetRejectionOk returns a tuple with the Rejection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejection

`func (o *NetworkbackendNetworkCandidate) SetRejection(v VmpreflightRejection)`

SetRejection sets Rejection field to given value.

### HasRejection

`func (o *NetworkbackendNetworkCandidate) HasRejection() bool`

HasRejection returns a boolean if a field has been set.

### GetSpecsUnknown

`func (o *NetworkbackendNetworkCandidate) GetSpecsUnknown() bool`

GetSpecsUnknown returns the SpecsUnknown field if non-nil, zero value otherwise.

### GetSpecsUnknownOk

`func (o *NetworkbackendNetworkCandidate) GetSpecsUnknownOk() (*bool, bool)`

GetSpecsUnknownOk returns a tuple with the SpecsUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecsUnknown

`func (o *NetworkbackendNetworkCandidate) SetSpecsUnknown(v bool)`

SetSpecsUnknown sets SpecsUnknown field to given value.

### HasSpecsUnknown

`func (o *NetworkbackendNetworkCandidate) HasSpecsUnknown() bool`

HasSpecsUnknown returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkbackendNetworkCandidate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkbackendNetworkCandidate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkbackendNetworkCandidate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkbackendNetworkCandidate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVmId

`func (o *NetworkbackendNetworkCandidate) GetVmId() int32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *NetworkbackendNetworkCandidate) GetVmIdOk() (*int32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *NetworkbackendNetworkCandidate) SetVmId(v int32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *NetworkbackendNetworkCandidate) HasVmId() bool`

HasVmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


