# HandlersRoleRecommendationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **int32** |  | [optional] 
**VmIds** | **[]int32** |  | 

## Methods

### NewHandlersRoleRecommendationPayload

`func NewHandlersRoleRecommendationPayload(vmIds []int32, ) *HandlersRoleRecommendationPayload`

NewHandlersRoleRecommendationPayload instantiates a new HandlersRoleRecommendationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRoleRecommendationPayloadWithDefaults

`func NewHandlersRoleRecommendationPayloadWithDefaults() *HandlersRoleRecommendationPayload`

NewHandlersRoleRecommendationPayloadWithDefaults instantiates a new HandlersRoleRecommendationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *HandlersRoleRecommendationPayload) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *HandlersRoleRecommendationPayload) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *HandlersRoleRecommendationPayload) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *HandlersRoleRecommendationPayload) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetVmIds

`func (o *HandlersRoleRecommendationPayload) GetVmIds() []int32`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *HandlersRoleRecommendationPayload) GetVmIdsOk() (*[]int32, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *HandlersRoleRecommendationPayload) SetVmIds(v []int32)`

SetVmIds sets VmIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


