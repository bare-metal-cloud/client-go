# HandlersStackReviewPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changed** | Pointer to **bool** |  | [optional] 
**Delta** | Pointer to [**HandlersStackDeltaPayload**](HandlersStackDeltaPayload.md) |  | [optional] 
**Members** | Pointer to [**[]HandlersStackMemberPayload**](HandlersStackMemberPayload.md) |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**Strategy** | Pointer to [**HandlersStackStrategyPayload**](HandlersStackStrategyPayload.md) |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**Waves** | Pointer to [**[]HandlersStackWavePayload**](HandlersStackWavePayload.md) |  | [optional] 
**WiringDelta** | Pointer to [**HandlersStackWiringDeltaPayload**](HandlersStackWiringDeltaPayload.md) |  | [optional] 

## Methods

### NewHandlersStackReviewPayload

`func NewHandlersStackReviewPayload() *HandlersStackReviewPayload`

NewHandlersStackReviewPayload instantiates a new HandlersStackReviewPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersStackReviewPayloadWithDefaults

`func NewHandlersStackReviewPayloadWithDefaults() *HandlersStackReviewPayload`

NewHandlersStackReviewPayloadWithDefaults instantiates a new HandlersStackReviewPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanged

`func (o *HandlersStackReviewPayload) GetChanged() bool`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *HandlersStackReviewPayload) GetChangedOk() (*bool, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *HandlersStackReviewPayload) SetChanged(v bool)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *HandlersStackReviewPayload) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### GetDelta

`func (o *HandlersStackReviewPayload) GetDelta() HandlersStackDeltaPayload`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *HandlersStackReviewPayload) GetDeltaOk() (*HandlersStackDeltaPayload, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *HandlersStackReviewPayload) SetDelta(v HandlersStackDeltaPayload)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *HandlersStackReviewPayload) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetMembers

`func (o *HandlersStackReviewPayload) GetMembers() []HandlersStackMemberPayload`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *HandlersStackReviewPayload) GetMembersOk() (*[]HandlersStackMemberPayload, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *HandlersStackReviewPayload) SetMembers(v []HandlersStackMemberPayload)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *HandlersStackReviewPayload) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetStack

`func (o *HandlersStackReviewPayload) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *HandlersStackReviewPayload) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *HandlersStackReviewPayload) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *HandlersStackReviewPayload) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetStrategy

`func (o *HandlersStackReviewPayload) GetStrategy() HandlersStackStrategyPayload`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *HandlersStackReviewPayload) GetStrategyOk() (*HandlersStackStrategyPayload, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *HandlersStackReviewPayload) SetStrategy(v HandlersStackStrategyPayload)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *HandlersStackReviewPayload) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetTarget

`func (o *HandlersStackReviewPayload) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *HandlersStackReviewPayload) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *HandlersStackReviewPayload) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *HandlersStackReviewPayload) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetWaves

`func (o *HandlersStackReviewPayload) GetWaves() []HandlersStackWavePayload`

GetWaves returns the Waves field if non-nil, zero value otherwise.

### GetWavesOk

`func (o *HandlersStackReviewPayload) GetWavesOk() (*[]HandlersStackWavePayload, bool)`

GetWavesOk returns a tuple with the Waves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaves

`func (o *HandlersStackReviewPayload) SetWaves(v []HandlersStackWavePayload)`

SetWaves sets Waves field to given value.

### HasWaves

`func (o *HandlersStackReviewPayload) HasWaves() bool`

HasWaves returns a boolean if a field has been set.

### GetWiringDelta

`func (o *HandlersStackReviewPayload) GetWiringDelta() HandlersStackWiringDeltaPayload`

GetWiringDelta returns the WiringDelta field if non-nil, zero value otherwise.

### GetWiringDeltaOk

`func (o *HandlersStackReviewPayload) GetWiringDeltaOk() (*HandlersStackWiringDeltaPayload, bool)`

GetWiringDeltaOk returns a tuple with the WiringDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiringDelta

`func (o *HandlersStackReviewPayload) SetWiringDelta(v HandlersStackWiringDeltaPayload)`

SetWiringDelta sets WiringDelta field to given value.

### HasWiringDelta

`func (o *HandlersStackReviewPayload) HasWiringDelta() bool`

HasWiringDelta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


