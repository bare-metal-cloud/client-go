# HandlersRunnerPackingPoolFlavorRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Busy** | Pointer to **int32** | Busy is the flavor&#39;s slices serving a job (the boundJobId set). | [optional] 
**ClassName** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** | Total is the flavor&#39;s live slices on the host. | [optional] 
**Warm** | Pointer to **int32** | Warm is the flavor&#39;s slices held as warm stock (idle past their job or created ahead of demand). | [optional] 

## Methods

### NewHandlersRunnerPackingPoolFlavorRow

`func NewHandlersRunnerPackingPoolFlavorRow() *HandlersRunnerPackingPoolFlavorRow`

NewHandlersRunnerPackingPoolFlavorRow instantiates a new HandlersRunnerPackingPoolFlavorRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerPackingPoolFlavorRowWithDefaults

`func NewHandlersRunnerPackingPoolFlavorRowWithDefaults() *HandlersRunnerPackingPoolFlavorRow`

NewHandlersRunnerPackingPoolFlavorRowWithDefaults instantiates a new HandlersRunnerPackingPoolFlavorRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusy

`func (o *HandlersRunnerPackingPoolFlavorRow) GetBusy() int32`

GetBusy returns the Busy field if non-nil, zero value otherwise.

### GetBusyOk

`func (o *HandlersRunnerPackingPoolFlavorRow) GetBusyOk() (*int32, bool)`

GetBusyOk returns a tuple with the Busy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusy

`func (o *HandlersRunnerPackingPoolFlavorRow) SetBusy(v int32)`

SetBusy sets Busy field to given value.

### HasBusy

`func (o *HandlersRunnerPackingPoolFlavorRow) HasBusy() bool`

HasBusy returns a boolean if a field has been set.

### GetClassName

`func (o *HandlersRunnerPackingPoolFlavorRow) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *HandlersRunnerPackingPoolFlavorRow) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *HandlersRunnerPackingPoolFlavorRow) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *HandlersRunnerPackingPoolFlavorRow) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetTotal

`func (o *HandlersRunnerPackingPoolFlavorRow) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HandlersRunnerPackingPoolFlavorRow) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HandlersRunnerPackingPoolFlavorRow) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HandlersRunnerPackingPoolFlavorRow) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWarm

`func (o *HandlersRunnerPackingPoolFlavorRow) GetWarm() int32`

GetWarm returns the Warm field if non-nil, zero value otherwise.

### GetWarmOk

`func (o *HandlersRunnerPackingPoolFlavorRow) GetWarmOk() (*int32, bool)`

GetWarmOk returns a tuple with the Warm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarm

`func (o *HandlersRunnerPackingPoolFlavorRow) SetWarm(v int32)`

SetWarm sets Warm field to given value.

### HasWarm

`func (o *HandlersRunnerPackingPoolFlavorRow) HasWarm() bool`

HasWarm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


