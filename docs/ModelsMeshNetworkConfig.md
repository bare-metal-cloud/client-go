# ModelsMeshNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DerpEnabled** | Pointer to **bool** |  | [optional] 
**HeadscaleUserId** | Pointer to **string** |  | [optional] 
**IpRange** | Pointer to **string** | CGNAT range (100.64.0.0/10) | [optional] 

## Methods

### NewModelsMeshNetworkConfig

`func NewModelsMeshNetworkConfig() *ModelsMeshNetworkConfig`

NewModelsMeshNetworkConfig instantiates a new ModelsMeshNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsMeshNetworkConfigWithDefaults

`func NewModelsMeshNetworkConfigWithDefaults() *ModelsMeshNetworkConfig`

NewModelsMeshNetworkConfigWithDefaults instantiates a new ModelsMeshNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDerpEnabled

`func (o *ModelsMeshNetworkConfig) GetDerpEnabled() bool`

GetDerpEnabled returns the DerpEnabled field if non-nil, zero value otherwise.

### GetDerpEnabledOk

`func (o *ModelsMeshNetworkConfig) GetDerpEnabledOk() (*bool, bool)`

GetDerpEnabledOk returns a tuple with the DerpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerpEnabled

`func (o *ModelsMeshNetworkConfig) SetDerpEnabled(v bool)`

SetDerpEnabled sets DerpEnabled field to given value.

### HasDerpEnabled

`func (o *ModelsMeshNetworkConfig) HasDerpEnabled() bool`

HasDerpEnabled returns a boolean if a field has been set.

### GetHeadscaleUserId

`func (o *ModelsMeshNetworkConfig) GetHeadscaleUserId() string`

GetHeadscaleUserId returns the HeadscaleUserId field if non-nil, zero value otherwise.

### GetHeadscaleUserIdOk

`func (o *ModelsMeshNetworkConfig) GetHeadscaleUserIdOk() (*string, bool)`

GetHeadscaleUserIdOk returns a tuple with the HeadscaleUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadscaleUserId

`func (o *ModelsMeshNetworkConfig) SetHeadscaleUserId(v string)`

SetHeadscaleUserId sets HeadscaleUserId field to given value.

### HasHeadscaleUserId

`func (o *ModelsMeshNetworkConfig) HasHeadscaleUserId() bool`

HasHeadscaleUserId returns a boolean if a field has been set.

### GetIpRange

`func (o *ModelsMeshNetworkConfig) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *ModelsMeshNetworkConfig) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *ModelsMeshNetworkConfig) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *ModelsMeshNetworkConfig) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


