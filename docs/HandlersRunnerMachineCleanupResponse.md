# HandlersRunnerMachineCleanupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CleanupDiskPercent** | Pointer to **int32** | CleanupDiskPercent: 0 &#x3D; not configured; 10-90 &#x3D; auto cleanup when the disk crosses this percent full. | [optional] 
**Delivery** | Pointer to **string** | Delivery states when the setting reaches the machine (the dual-path contract, on the wire so the panel can say it honestly). | [optional] 

## Methods

### NewHandlersRunnerMachineCleanupResponse

`func NewHandlersRunnerMachineCleanupResponse() *HandlersRunnerMachineCleanupResponse`

NewHandlersRunnerMachineCleanupResponse instantiates a new HandlersRunnerMachineCleanupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerMachineCleanupResponseWithDefaults

`func NewHandlersRunnerMachineCleanupResponseWithDefaults() *HandlersRunnerMachineCleanupResponse`

NewHandlersRunnerMachineCleanupResponseWithDefaults instantiates a new HandlersRunnerMachineCleanupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanupDiskPercent

`func (o *HandlersRunnerMachineCleanupResponse) GetCleanupDiskPercent() int32`

GetCleanupDiskPercent returns the CleanupDiskPercent field if non-nil, zero value otherwise.

### GetCleanupDiskPercentOk

`func (o *HandlersRunnerMachineCleanupResponse) GetCleanupDiskPercentOk() (*int32, bool)`

GetCleanupDiskPercentOk returns a tuple with the CleanupDiskPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupDiskPercent

`func (o *HandlersRunnerMachineCleanupResponse) SetCleanupDiskPercent(v int32)`

SetCleanupDiskPercent sets CleanupDiskPercent field to given value.

### HasCleanupDiskPercent

`func (o *HandlersRunnerMachineCleanupResponse) HasCleanupDiskPercent() bool`

HasCleanupDiskPercent returns a boolean if a field has been set.

### GetDelivery

`func (o *HandlersRunnerMachineCleanupResponse) GetDelivery() string`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *HandlersRunnerMachineCleanupResponse) GetDeliveryOk() (*string, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *HandlersRunnerMachineCleanupResponse) SetDelivery(v string)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *HandlersRunnerMachineCleanupResponse) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


