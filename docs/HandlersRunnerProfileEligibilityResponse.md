# HandlersRunnerProfileEligibilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineId** | Pointer to **string** |  | [optional] 
**Profiles** | Pointer to [**[]HandlersRunnerProfileEligibilityVerdict**](HandlersRunnerProfileEligibilityVerdict.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlersRunnerProfileEligibilityResponse

`func NewHandlersRunnerProfileEligibilityResponse() *HandlersRunnerProfileEligibilityResponse`

NewHandlersRunnerProfileEligibilityResponse instantiates a new HandlersRunnerProfileEligibilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlersRunnerProfileEligibilityResponseWithDefaults

`func NewHandlersRunnerProfileEligibilityResponseWithDefaults() *HandlersRunnerProfileEligibilityResponse`

NewHandlersRunnerProfileEligibilityResponseWithDefaults instantiates a new HandlersRunnerProfileEligibilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineId

`func (o *HandlersRunnerProfileEligibilityResponse) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *HandlersRunnerProfileEligibilityResponse) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *HandlersRunnerProfileEligibilityResponse) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *HandlersRunnerProfileEligibilityResponse) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetProfiles

`func (o *HandlersRunnerProfileEligibilityResponse) GetProfiles() []HandlersRunnerProfileEligibilityVerdict`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *HandlersRunnerProfileEligibilityResponse) GetProfilesOk() (*[]HandlersRunnerProfileEligibilityVerdict, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *HandlersRunnerProfileEligibilityResponse) SetProfiles(v []HandlersRunnerProfileEligibilityVerdict)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *HandlersRunnerProfileEligibilityResponse) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetSummary

`func (o *HandlersRunnerProfileEligibilityResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *HandlersRunnerProfileEligibilityResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *HandlersRunnerProfileEligibilityResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *HandlersRunnerProfileEligibilityResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


