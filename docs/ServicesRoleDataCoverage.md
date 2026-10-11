# ServicesRoleDataCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stale** | Pointer to **bool** |  | [optional] 
**StaleNote** | Pointer to **string** | StaleNote is the ENGINE&#39;s honest escalation copy wired through when Stale is true (review MED-2) — BE-owned, rendered verbatim; empty (omitempty) on the fresh path. | [optional] 
**VmsWithData** | Pointer to **int32** |  | [optional] 
**WindowHours** | Pointer to **int32** |  | [optional] 

## Methods

### NewServicesRoleDataCoverage

`func NewServicesRoleDataCoverage() *ServicesRoleDataCoverage`

NewServicesRoleDataCoverage instantiates a new ServicesRoleDataCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesRoleDataCoverageWithDefaults

`func NewServicesRoleDataCoverageWithDefaults() *ServicesRoleDataCoverage`

NewServicesRoleDataCoverageWithDefaults instantiates a new ServicesRoleDataCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStale

`func (o *ServicesRoleDataCoverage) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ServicesRoleDataCoverage) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ServicesRoleDataCoverage) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ServicesRoleDataCoverage) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetStaleNote

`func (o *ServicesRoleDataCoverage) GetStaleNote() string`

GetStaleNote returns the StaleNote field if non-nil, zero value otherwise.

### GetStaleNoteOk

`func (o *ServicesRoleDataCoverage) GetStaleNoteOk() (*string, bool)`

GetStaleNoteOk returns a tuple with the StaleNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleNote

`func (o *ServicesRoleDataCoverage) SetStaleNote(v string)`

SetStaleNote sets StaleNote field to given value.

### HasStaleNote

`func (o *ServicesRoleDataCoverage) HasStaleNote() bool`

HasStaleNote returns a boolean if a field has been set.

### GetVmsWithData

`func (o *ServicesRoleDataCoverage) GetVmsWithData() int32`

GetVmsWithData returns the VmsWithData field if non-nil, zero value otherwise.

### GetVmsWithDataOk

`func (o *ServicesRoleDataCoverage) GetVmsWithDataOk() (*int32, bool)`

GetVmsWithDataOk returns a tuple with the VmsWithData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmsWithData

`func (o *ServicesRoleDataCoverage) SetVmsWithData(v int32)`

SetVmsWithData sets VmsWithData field to given value.

### HasVmsWithData

`func (o *ServicesRoleDataCoverage) HasVmsWithData() bool`

HasVmsWithData returns a boolean if a field has been set.

### GetWindowHours

`func (o *ServicesRoleDataCoverage) GetWindowHours() int32`

GetWindowHours returns the WindowHours field if non-nil, zero value otherwise.

### GetWindowHoursOk

`func (o *ServicesRoleDataCoverage) GetWindowHoursOk() (*int32, bool)`

GetWindowHoursOk returns a tuple with the WindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowHours

`func (o *ServicesRoleDataCoverage) SetWindowHours(v int32)`

SetWindowHours sets WindowHours field to given value.

### HasWindowHours

`func (o *ServicesRoleDataCoverage) HasWindowHours() bool`

HasWindowHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


