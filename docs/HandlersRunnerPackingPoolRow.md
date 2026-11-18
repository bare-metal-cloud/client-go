# HandlersRunnerPackingPoolRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavors** | Pointer to [**[]HandlersRunnerPackingPoolFlavorRow**](HandlersRunnerPackingPoolFlavorRow.md) | Flavors carries the per-flavor counts (newest flavor first by name — the grouping is deterministic, the page adds no sort). | [optional] 
**PoolId** | Pointer to **string** | PoolID is the packing host&#39;s id (the compound slice provider id&#39;s parent path — the census&#39;s grouping key). | [optional] 
**SliceCount** | Pointer to **int32** | SliceCount is the host&#39;s live slices across flavors. | [optional] 

## Methods

### NewHandlersRunnerPackingPoolRow

`func NewHandlersRunnerPackingPoolRow() *HandlersRunnerPackingPoolRow`

NewHandlersRunnerPackingPoolRow instantiates a new HandlersRunnerPackingPoolRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerPackingPoolRowWithDefaults

`func NewHandlersRunnerPackingPoolRowWithDefaults() *HandlersRunnerPackingPoolRow`

NewHandlersRunnerPackingPoolRowWithDefaults instantiates a new HandlersRunnerPackingPoolRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavors

`func (o *HandlersRunnerPackingPoolRow) GetFlavors() []HandlersRunnerPackingPoolFlavorRow`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *HandlersRunnerPackingPoolRow) GetFlavorsOk() (*[]HandlersRunnerPackingPoolFlavorRow, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *HandlersRunnerPackingPoolRow) SetFlavors(v []HandlersRunnerPackingPoolFlavorRow)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *HandlersRunnerPackingPoolRow) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### GetPoolId

`func (o *HandlersRunnerPackingPoolRow) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *HandlersRunnerPackingPoolRow) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *HandlersRunnerPackingPoolRow) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *HandlersRunnerPackingPoolRow) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetSliceCount

`func (o *HandlersRunnerPackingPoolRow) GetSliceCount() int32`

GetSliceCount returns the SliceCount field if non-nil, zero value otherwise.

### GetSliceCountOk

`func (o *HandlersRunnerPackingPoolRow) GetSliceCountOk() (*int32, bool)`

GetSliceCountOk returns a tuple with the SliceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceCount

`func (o *HandlersRunnerPackingPoolRow) SetSliceCount(v int32)`

SetSliceCount sets SliceCount field to given value.

### HasSliceCount

`func (o *HandlersRunnerPackingPoolRow) HasSliceCount() bool`

HasSliceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


