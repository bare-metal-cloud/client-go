# HandlersPeerProbePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **int32** | NetworkID is the network the mesh probe measures (required for the mesh plane; rejected for direct). | [optional] 
**Plane** | Pointer to **string** | Plane defaults to \&quot;direct\&quot;; \&quot;mesh\&quot; requires the network_id. | [optional] 
**Retest** | Pointer to **bool** | Retest is the AC-5 affordance: explicitly invalidate the cached failed verdict and probe again (rate-limit counted). | [optional] 
**VmIds** | **[]int32** |  | 

## Methods

### NewHandlersPeerProbePayload

`func NewHandlersPeerProbePayload(vmIds []int32, ) *HandlersPeerProbePayload`

NewHandlersPeerProbePayload instantiates a new HandlersPeerProbePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersPeerProbePayloadWithDefaults

`func NewHandlersPeerProbePayloadWithDefaults() *HandlersPeerProbePayload`

NewHandlersPeerProbePayloadWithDefaults instantiates a new HandlersPeerProbePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *HandlersPeerProbePayload) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *HandlersPeerProbePayload) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *HandlersPeerProbePayload) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *HandlersPeerProbePayload) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPlane

`func (o *HandlersPeerProbePayload) GetPlane() string`

GetPlane returns the Plane field if non-nil, zero value otherwise.

### GetPlaneOk

`func (o *HandlersPeerProbePayload) GetPlaneOk() (*string, bool)`

GetPlaneOk returns a tuple with the Plane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlane

`func (o *HandlersPeerProbePayload) SetPlane(v string)`

SetPlane sets Plane field to given value.

### HasPlane

`func (o *HandlersPeerProbePayload) HasPlane() bool`

HasPlane returns a boolean if a field has been set.

### GetRetest

`func (o *HandlersPeerProbePayload) GetRetest() bool`

GetRetest returns the Retest field if non-nil, zero value otherwise.

### GetRetestOk

`func (o *HandlersPeerProbePayload) GetRetestOk() (*bool, bool)`

GetRetestOk returns a tuple with the Retest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetest

`func (o *HandlersPeerProbePayload) SetRetest(v bool)`

SetRetest sets Retest field to given value.

### HasRetest

`func (o *HandlersPeerProbePayload) HasRetest() bool`

HasRetest returns a boolean if a field has been set.

### GetVmIds

`func (o *HandlersPeerProbePayload) GetVmIds() []int32`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *HandlersPeerProbePayload) GetVmIdsOk() (*[]int32, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *HandlersPeerProbePayload) SetVmIds(v []int32)`

SetVmIds sets VmIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


