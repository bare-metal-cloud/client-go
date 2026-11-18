# HandlersRunnerPoolSettingsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumJobAgeSeconds** | Pointer to **int32** | MinimumJobAgeSeconds is the intake&#39;s claim backoff (0 disables it; default 30). | [optional] 
**RowExists** | Pointer to **bool** | RowExists is whether the org carries a settings row (the warm opt-in&#39;s record). False &#x3D; the defaults are the answer. | [optional] 
**WarmFloors** | Pointer to [**[]ModelsRunnerOrgWarmFloor**](ModelsRunnerOrgWarmFloor.md) | WarmFloors is the per-flavor floor list (empty &#x3D; the pool maintains nothing for this org). | [optional] 
**WarmPoolGate** | Pointer to [**HandlersRunnerPoolGateWire**](HandlersRunnerPoolGateWire.md) | WarmPoolGate is the SERVER-recorded warm-gate decision for the org (feature_gates bmc_runner_warm_pool, fail-closed off). The row&#39;s existence AND the gate must both read on before the pool maintains anything — the UI&#39;s gate-off empty state names this. | [optional] 

## Methods

### NewHandlersRunnerPoolSettingsBody

`func NewHandlersRunnerPoolSettingsBody() *HandlersRunnerPoolSettingsBody`

NewHandlersRunnerPoolSettingsBody instantiates a new HandlersRunnerPoolSettingsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerPoolSettingsBodyWithDefaults

`func NewHandlersRunnerPoolSettingsBodyWithDefaults() *HandlersRunnerPoolSettingsBody`

NewHandlersRunnerPoolSettingsBodyWithDefaults instantiates a new HandlersRunnerPoolSettingsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimumJobAgeSeconds

`func (o *HandlersRunnerPoolSettingsBody) GetMinimumJobAgeSeconds() int32`

GetMinimumJobAgeSeconds returns the MinimumJobAgeSeconds field if non-nil, zero value otherwise.

### GetMinimumJobAgeSecondsOk

`func (o *HandlersRunnerPoolSettingsBody) GetMinimumJobAgeSecondsOk() (*int32, bool)`

GetMinimumJobAgeSecondsOk returns a tuple with the MinimumJobAgeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumJobAgeSeconds

`func (o *HandlersRunnerPoolSettingsBody) SetMinimumJobAgeSeconds(v int32)`

SetMinimumJobAgeSeconds sets MinimumJobAgeSeconds field to given value.

### HasMinimumJobAgeSeconds

`func (o *HandlersRunnerPoolSettingsBody) HasMinimumJobAgeSeconds() bool`

HasMinimumJobAgeSeconds returns a boolean if a field has been set.

### GetRowExists

`func (o *HandlersRunnerPoolSettingsBody) GetRowExists() bool`

GetRowExists returns the RowExists field if non-nil, zero value otherwise.

### GetRowExistsOk

`func (o *HandlersRunnerPoolSettingsBody) GetRowExistsOk() (*bool, bool)`

GetRowExistsOk returns a tuple with the RowExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowExists

`func (o *HandlersRunnerPoolSettingsBody) SetRowExists(v bool)`

SetRowExists sets RowExists field to given value.

### HasRowExists

`func (o *HandlersRunnerPoolSettingsBody) HasRowExists() bool`

HasRowExists returns a boolean if a field has been set.

### GetWarmFloors

`func (o *HandlersRunnerPoolSettingsBody) GetWarmFloors() []ModelsRunnerOrgWarmFloor`

GetWarmFloors returns the WarmFloors field if non-nil, zero value otherwise.

### GetWarmFloorsOk

`func (o *HandlersRunnerPoolSettingsBody) GetWarmFloorsOk() (*[]ModelsRunnerOrgWarmFloor, bool)`

GetWarmFloorsOk returns a tuple with the WarmFloors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmFloors

`func (o *HandlersRunnerPoolSettingsBody) SetWarmFloors(v []ModelsRunnerOrgWarmFloor)`

SetWarmFloors sets WarmFloors field to given value.

### HasWarmFloors

`func (o *HandlersRunnerPoolSettingsBody) HasWarmFloors() bool`

HasWarmFloors returns a boolean if a field has been set.

### GetWarmPoolGate

`func (o *HandlersRunnerPoolSettingsBody) GetWarmPoolGate() HandlersRunnerPoolGateWire`

GetWarmPoolGate returns the WarmPoolGate field if non-nil, zero value otherwise.

### GetWarmPoolGateOk

`func (o *HandlersRunnerPoolSettingsBody) GetWarmPoolGateOk() (*HandlersRunnerPoolGateWire, bool)`

GetWarmPoolGateOk returns a tuple with the WarmPoolGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmPoolGate

`func (o *HandlersRunnerPoolSettingsBody) SetWarmPoolGate(v HandlersRunnerPoolGateWire)`

SetWarmPoolGate sets WarmPoolGate field to given value.

### HasWarmPoolGate

`func (o *HandlersRunnerPoolSettingsBody) HasWarmPoolGate() bool`

HasWarmPoolGate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


