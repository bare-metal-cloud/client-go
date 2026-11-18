# AddresschangeExecuteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcomes** | Pointer to [**[]AddresschangeOutcome**](AddresschangeOutcome.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to **[]string** |  | [optional] 
**Waiting** | Pointer to [**[]AddresschangeRecord**](AddresschangeRecord.md) | Waiting carries every record NOT proven answered at the end of the arm — the honest remainder the surface stays open on. | [optional] 
**Zones** | Pointer to [**[]AddresschangeZone**](AddresschangeZone.md) |  | [optional] 

## Methods

### NewAddresschangeExecuteResult

`func NewAddresschangeExecuteResult() *AddresschangeExecuteResult`

NewAddresschangeExecuteResult instantiates a new AddresschangeExecuteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddresschangeExecuteResultWithDefaults

`func NewAddresschangeExecuteResultWithDefaults() *AddresschangeExecuteResult`

NewAddresschangeExecuteResultWithDefaults instantiates a new AddresschangeExecuteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcomes

`func (o *AddresschangeExecuteResult) GetOutcomes() []AddresschangeOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *AddresschangeExecuteResult) GetOutcomesOk() (*[]AddresschangeOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *AddresschangeExecuteResult) SetOutcomes(v []AddresschangeOutcome)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *AddresschangeExecuteResult) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### GetState

`func (o *AddresschangeExecuteResult) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddresschangeExecuteResult) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddresschangeExecuteResult) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddresschangeExecuteResult) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSteps

`func (o *AddresschangeExecuteResult) GetSteps() []string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *AddresschangeExecuteResult) GetStepsOk() (*[]string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *AddresschangeExecuteResult) SetSteps(v []string)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *AddresschangeExecuteResult) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetWaiting

`func (o *AddresschangeExecuteResult) GetWaiting() []AddresschangeRecord`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *AddresschangeExecuteResult) GetWaitingOk() (*[]AddresschangeRecord, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *AddresschangeExecuteResult) SetWaiting(v []AddresschangeRecord)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *AddresschangeExecuteResult) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.

### GetZones

`func (o *AddresschangeExecuteResult) GetZones() []AddresschangeZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AddresschangeExecuteResult) GetZonesOk() (*[]AddresschangeZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AddresschangeExecuteResult) SetZones(v []AddresschangeZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AddresschangeExecuteResult) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


